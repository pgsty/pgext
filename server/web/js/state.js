/* ================================================================
   In-memory catalog snapshot, detail TTL caches, URL filter state
   and JSON fetch helpers.
   ================================================================ */
import { getLang } from './i18n.js';

export const F = { LEAD: 1, CONTRIB: 2, BIN: 4, LIB: 8, DDL: 16, LOAD: 32, TRUSTED: 64, RELOC: 128 };
export const B = { RPM: 1, DEB: 2, PGRX: 4, SOURCE: 8 };
export const R = { REQUIRES: 1, REQUIRED_BY: 2, SEE_ALSO: 4 };

// Revisited detail pages must pick up doc-only and package-only updates too.
export class DetailCache {
  constructor(ttlMs = 5 * 60 * 1000) { this.items = new Map(); this.ttlMs = ttlMs; }
  get(key) {
    const entry = this.items.get(key);
    if (!entry) return undefined;
    if (Date.now() >= entry.expires) { this.items.delete(key); return undefined; }
    return entry.value;
  }
  set(key, value) { this.items.set(key, { value, expires: Date.now() + this.ttlMs }); }
  clear() { this.items.clear(); }
}

export const FULLC = new DetailCache();
export const MXC = new DetailCache();
export const FILEC = new DetailCache();
export const DOCC = new DetailCache();
export const CHLOGC = new DetailCache();

export const catalog = {
  boot: null,
  ext: [],
  byName: new Map(),
  byPkg: new Map(),
  cats: {},
  catOrder: [],
  pgs: [],
  oss: [],
  osIdx: {},
  ufield: [],
  nAll: 0, nAvail: 0, nProjects: 0, nSource: 0, nVendor: 0, nKernel: 0, nContrib: 0, nDocs: 0,
  nPkgs: 0, nSlots: 0,
  meta: {}
};

export const matrixState = {
  data: {},     // view ('' | 'pgdg' | 'pigsty') -> payload
  checked: {},  // view -> true once revalidated this session
  inflight: {}, // view -> Promise
  view: null,   // active canvas view controller { render, state, destroy }
  hydSeq: 0,
  resetChecks() {
    for (const k of Object.keys(this.checked)) delete this.checked[k];
  }
};

/* bootstrap row columns — keep in sync with marshalBootstrap in server/api.go:
   0 name 1 cat 2 avail 3 repo 4 license 5 lang 6 version 7 stars
   8 en 9 zh 10 kind 11 vendor 12 kernel 13 pg[] 14 flags 15 docbits 16 commit
   17 pkg 18 lead 19 lifecycle 20 tags[] 21 active 22 checked 23 buildbits
   24 target_idx[] 25 family_size 26 comment 27 relationbits 28 pgrx_ver 29 repo_url 30 url 31 license_url
   32 rpm_pkg 33 deb_pkg */
export function decodeBoot(b) {
  if (catalog.boot && catalog.boot.version !== b.version) {
    for (const cache of [FULLC, MXC, FILEC, DOCC, CHLOGC]) cache.clear();
    matrixState.resetChecks();
  }
  catalog.boot = b;
  const ext = b.rows.map((r, i) => ({
    i, name: r[0], cat: r[1], avail: !!r[2], repo: r[3] || 'n/a', license: r[4] || 'Unknown',
    lang: r[5] || '', ver: r[6] || '', stars: r[7], en: r[8] || '', zh: r[9] || '',
    kind: r[10] || '', vendor: r[11] || '', kernel: r[12] || '',
    pg: r[13] || [], flags: r[14] | 0, docbits: r[15] | 0, commit: r[16] || '',
    pkg: r[17] || r[0], lead: r[18] || r[0], lifecycle: r[19] || '', tags: r[20] || [],
    active: r[21] || r[16] || '', checked: r[22] || '', buildbits: r[23] | 0,
    targets: r[24] || [], familySize: r[25] || 1, comment: r[26] || '',
    relationbits: r[27] | 0, pgrx: r[28] || '', repoUrl: r[29] || '', url: r[30] || '', licenseUrl: r[31] || '',
    rpmPkg: r[32] || '', debPkg: r[33] || ''
  }));
  const byName = new Map(ext.map(e => [e.name, e]));
  const byPkg = new Map();
  for (const e of ext) {
    if (!byPkg.has(e.pkg)) byPkg.set(e.pkg, []);
    byPkg.get(e.pkg).push(e);
  }
  const cats = {}, catOrder = [];
  for (const c of b.cats) {
    cats[c.name] = { code: c.name, zh: c.zh_desc, en: c.en_desc, count: c.count };
    catOrder.push(c.name);
  }
  const pgs = b.pg || [], oss = b.os || [], osIdx = {};
  oss.forEach((os, i) => { osIdx[os] = i; });

  // Counts derivable from the rows are derived so figures never disagree with
  // the decoded catalog; only pkg-table aggregates rely on the server.
  let nAvail = 0, nSource = 0, nVendor = 0, nKernel = 0, nContrib = 0;
  for (const e of ext) {
    if (e.avail) nAvail++;
    else if (e.buildbits & B.SOURCE) nSource++;
    if (e.vendor) nVendor++;
    if (e.kernel) nKernel++;
    if (e.flags & F.CONTRIB) nContrib++;
  }
  const ufield = [];
  for (const c of catOrder) {
    const members = ext.filter(e => e.cat === c);
    members.sort((a, b2) => (b2.stars ?? -1) - (a.stars ?? -1) || a.name.localeCompare(b2.name));
    ufield.push(...members);
  }

  Object.assign(catalog, {
    ext, byName, byPkg, cats, catOrder, pgs, oss, osIdx, ufield,
    nAll: ext.length, nAvail, nProjects: byPkg.size, nSource, nVendor, nKernel, nContrib,
    nDocs: (b.counts && b.counts.docs) || 0,
    nPkgs: (b.counts && b.counts.packages) ?? 0,
    nSlots: (b.counts && b.counts.build_slots) ?? 0,
    meta: { generated: (b.generated || '').slice(0, 10) }
  });
}

export const catDesc = c => {
  const m = catalog.cats[c];
  return m ? (getLang() === 'zh' ? m.zh : m.en) : '';
};

/* ---------------- URL & filter state ---------------- */
export const FILTER_KEYS = ['q', 'repo', 'license', 'lang', 'pg', 'os', 'kind', 'lifecycle', 'scope',
  'vendor', 'kernel', 'tag', 'pkg', 'capability', 'build', 'docs', 'relation', 'pgrx', 'active'];

export const DEFAULT_STATE = {
  q: '', cat: '', repo: '', license: '', lang: '', pg: '', os: '', kind: '', lifecycle: '', scope: '',
  vendor: '', kernel: '', tag: '', pkg: '', capability: '', build: '', docs: '', relation: '', pgrx: '', active: '',
  sort: '', entity: 'ext', layout: 'card'
};

export const S = { ...DEFAULT_STATE };
export function resetState() { Object.assign(S, DEFAULT_STATE); }
export const effSort = () => S.sort || 'stars';

export function parsePGs(value) {
  const seen = new Set();
  for (const raw of String(value || '').split(/[,\s;]+/)) {
    const pg = Number.parseInt(raw, 10);
    if (Number.isInteger(pg) && pg > 0 && pg < 100) seen.add(pg);
  }
  return [...seen].sort((a, b) => b - a);
}

export function migrateLegacyHash() {
  if (typeof location === 'undefined') return false;
  const raw = location.hash.startsWith('#/') ? location.hash.slice(1) : '';
  if (!raw) return false;
  const qi = raw.indexOf('?');
  let path = qi >= 0 ? raw.slice(0, qi) : raw;
  const query = qi >= 0 ? raw.slice(qi + 1) : '';
  if (path.startsWith('/cat/')) path = '/cate/' + path.slice(5);
  history.replaceState(null, '', path + (query ? '?' + query : ''));
  return true;
}

export function parseRoute() {
  migrateLegacyHash();
  const path = location.pathname.length > 1 ? location.pathname.replace(/\/+$/, '') : '/';
  return { path, params: new URLSearchParams(location.search) };
}

export function pushState() {
  const p = new URLSearchParams();
  if (S.cat) p.set('cate', S.cat);
  for (const k of FILTER_KEYS) if (S[k]) p.set(k, S[k]);
  if (S.sort) p.set('sort', S.sort);
  if (S.entity !== 'ext') p.set('entity', S.entity);
  if (S.layout !== 'card') p.set('layout', S.layout);
  const qs = p.toString();
  history.replaceState(null, '', '/' + (qs ? '?' + qs : ''));
}

const PARAM_ALIASES = {
  lang: 'lng', kind: 'type', pkg: 'package', capability: 'cap', docs: 'doc', relation: 'rel', active: 'year'
};

export function readState(params) {
  resetState();
  for (const k of [...FILTER_KEYS, 'sort', 'entity', 'layout']) {
    const v = params.get(k) || (PARAM_ALIASES[k] ? params.get(PARAM_ALIASES[k]) : '');
    if (v) S[k] = v;
  }
  S.cat = (params.get('cate') || params.get('cat') || '').toUpperCase();
  // ?lang=zh|en switches UI language, not the implementation-language filter.
  if (S.lang === 'zh' || S.lang === 'en') S.lang = params.get('lng') || '';
  if (S.scope.toLowerCase() === 'cloud') S.scope = 'vendor';
  else S.scope = S.scope.toLowerCase();
  const legacyView = params.get('view');
  if (legacyView === 'packages') { S.entity = 'pkg'; S.layout = 'card'; }
  if (legacyView === 'rows') { S.entity = 'ext'; S.layout = 'list'; }
  if (legacyView === 'grid') { S.entity = 'ext'; S.layout = 'card'; }
  if (!['ext', 'pkg'].includes(S.entity)) S.entity = 'ext';
  if (!['card', 'list'].includes(S.layout)) S.layout = 'card';
  S.pg = parsePGs(S.pg).join(',');
}

export async function j(url) {
  const res = await fetch(url, { headers: { Accept: 'application/json' } });
  if (!res.ok) {
    let msg = 'HTTP ' + res.status;
    try { const body = await res.json(); if (body.error) msg = body.error; } catch (e) {}
    const err = new Error(msg); err.status = res.status; throw err;
  }
  return res.json();
}

export function fill(id, html) {
  const el = document.getElementById(id);
  if (el) el.innerHTML = html;
}
