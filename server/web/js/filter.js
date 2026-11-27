/* ================================================================
   Client-side filter, ranking, SQL preview and dimension engine
   (mirrors server/api.go Filter semantics).
   ================================================================ */
import { esc, fmtInt, osLabel, starOf } from './util.js';
import { getLang, catName } from './i18n.js';
import { F, B, R, catalog, S, effSort, parsePGs } from './state.js';

export const DIMS = {
  category: { key: 'cate', label: 'dim.category', d: 'dim.category.d', field: 'category' },
  tag: { key: 'tag', label: 'dim.tag', d: 'dim.tag.d', field: 'tags[]' },
  package: { key: 'pkg', label: 'dim.package', d: 'dim.package.d', field: 'pkg · lead_ext' },
  kind: { key: 'kind', label: 'dim.type', d: 'dim.type.d', field: 'kind' },
  lifecycle: { key: 'lifecycle', label: 'dim.lifecycle', d: 'dim.lifecycle.d', field: 'lifecycle' },
  license: { key: 'license', label: 'dim.license', d: 'dim.license.d', field: 'license' },
  lang: { key: 'lang', label: 'dim.lang', d: 'dim.lang.d', field: 'lang' },
  distribution: { key: 'scope', label: 'dim.distribution', d: 'dim.distribution.d', field: 'packaged · contrib · vendor · kernel' },
  repo: { key: 'repo', label: 'dim.repo', d: 'dim.repo.d', field: 'rpm_repo · deb_repo' },
  pg: { key: 'pg', label: 'dim.pg', d: 'dim.pg.d', field: 'pg_ver[]' },
  os: { key: 'os', label: 'dim.os', d: 'dim.os.d', field: 'pgext.pkg targets' },
  build: { key: 'build', label: 'dim.build', d: 'dim.build.d', field: 'rpm_build · deb_build · pgrx_ver' },
  pgrx: { key: 'pgrx', label: 'dim.pgrx', d: 'dim.pgrx.d', field: 'pgrx_ver' },
  capability: { key: 'capability', label: 'dim.capability', d: 'dim.capability.d', field: 'has_* · need_* · trusted · relocatable' },
  docs: { key: 'docs', label: 'dim.docs', d: 'dim.docs.d', field: 'pgext.doc en_doc · zh_doc' },
  relation: { key: 'relation', label: 'dim.relation', d: 'dim.relation.d', field: 'requires[] · required_by[] · see_also[]' },
  vendor: { key: 'vendor', label: 'dim.vendor', d: 'dim.vendor.d', field: 'vendor' },
  kernel: { key: 'kernel', label: 'dim.kernel', d: 'dim.kernel.d', field: 'kernel' },
  activity: { key: 'active', label: 'dim.activity', d: 'dim.activity.d', field: 'last_active' }
};

export const DIM_GROUPS = [
  ['browse.catalog', ['tag', 'package', 'kind', 'lifecycle', 'license', 'lang']],
  ['browse.delivery', ['distribution', 'repo', 'pg', 'os', 'build', 'pgrx']],
  ['browse.attrs', ['capability', 'docs', 'relation', 'vendor', 'kernel', 'activity']]
];

export const DIM_VALUE_NAMES = {
  distribution: {
    packaged: ['Packaged', '已打包'], source: ['Source-only', '仅源码'], vendor: ['Vendor catalog', '厂商目录'],
    kernel: ['Kernel-specific', '内核特定'], contrib: ['PostgreSQL contrib', 'PG 自带']
  },
  build: {
    rpm: ['RPM recipe', 'RPM 配方'], deb: ['DEB recipe', 'DEB 配方'], pgrx: ['pgrx / Rust', 'pgrx / Rust'],
    source: ['Upstream source', '上游源码'], packaged: ['Binary package', '二进制包']
  },
  capability: {
    binary: ['Ships executables', '携带命令行工具'], library: ['Ships libraries', '携带动态库'],
    'create-extension': ['CREATE EXTENSION', 'CREATE EXTENSION'], preload: ['Explicit loading', '显式加载'],
    trusted: ['Trusted install', '非超户可装'], relocatable: ['Relocatable schema', '模式可迁移']
  },
  docs: {
    bilingual: ['English + Chinese', '中英双语'], 'english-only': ['English only', '仅英文'],
    'chinese-only': ['Chinese only', '仅中文'], undocumented: ['Usage manual missing', '尚缺用法手册']
  },
  relation: {
    requires: ['Has dependencies', '存在前置依赖'], 'required-by': ['Has dependants', '存在下游依赖'],
    'see-also': ['Has related entries', '存在相关扩展'], 'no-relations': ['No recorded relations', '暂无关联记录']
  },
  activity: { unknown: ['Unknown', '未知'] }
};

export function dimValueName(dim, value) {
  const raw = String(value == null ? '' : value);
  const pair = DIM_VALUE_NAMES[dim] && DIM_VALUE_NAMES[dim][raw];
  if (pair) return pair[getLang() === 'zh' ? 1 : 0];
  if (dim === 'pg') return 'PG ' + raw;
  if (dim === 'os') return osLabel(raw);
  return raw;
}

export function dimValueHTML(dim, value) {
  const raw = String(value);
  if (dim === 'category') return '<span class="dim-category"><code>' + esc(raw) + '</code><b>' + esc(catName(raw)) + '</b></span>';
  const name = dimValueName(dim, raw);
  if (name !== raw && dim !== 'pg' && dim !== 'os' && dim !== 'activity') {
    return '<span class="dim-value-name"><b>' + esc(name) + '</b><code>' + esc(raw) + '</code></span>';
  }
  return '<span class="mono">' + esc(name) + '</span>';
}

export function targetAvailable(e, pg, os) {
  const { oss, osIdx, pgs } = catalog;
  if (!e.targets || !e.targets.length || !oss.length) return false;
  const oi = os ? osIdx[os] : -1;
  if (os && oi == null) return false;
  for (const idx of e.targets) {
    const pi = Math.floor(idx / oss.length), ti = idx % oss.length;
    if (os && ti !== oi) continue;
    if (pg && pgs[pi] !== pg) continue;
    return true;
  }
  return false;
}

export const targetsAvailable = (e, pgs, os) =>
  pgs.length ? pgs.every(pg => targetAvailable(e, pg, os)) : targetAvailable(e, 0, os);

export function capabilityMatches(e, value) {
  switch (String(value || '').toLowerCase()) {
    case 'binary': case 'bin': return Boolean(e.flags & F.BIN);
    case 'library': case 'lib': return Boolean(e.flags & F.LIB);
    case 'create-extension': case 'ddl': return Boolean(e.flags & F.DDL);
    case 'preload': case 'load': return Boolean(e.flags & F.LOAD);
    case 'trusted': return Boolean(e.flags & F.TRUSTED);
    case 'relocatable': return Boolean(e.flags & F.RELOC);
    default: return false;
  }
}

export function buildMatches(e, value) {
  switch (String(value || '').toLowerCase()) {
    case 'rpm': return Boolean(e.buildbits & B.RPM);
    case 'deb': return Boolean(e.buildbits & B.DEB);
    case 'pgrx': return Boolean(e.buildbits & B.PGRX);
    case 'source': return Boolean(e.buildbits & B.SOURCE);
    case 'packaged': case 'binary-package': return e.avail;
    default: return false;
  }
}

export function docsClass(e) {
  if ((e.docbits & 3) === 3) return 'bilingual';
  if (e.docbits & 1) return 'english-only';
  if (e.docbits & 2) return 'chinese-only';
  return 'undocumented';
}

export function relationMatches(e, value) {
  switch (String(value || '').toLowerCase()) {
    case 'requires': return Boolean(e.relationbits & R.REQUIRES);
    case 'required-by': return Boolean(e.relationbits & R.REQUIRED_BY);
    case 'see-also': return Boolean(e.relationbits & R.SEE_ALSO);
    case 'no-relations': return e.relationbits === 0;
    default: return false;
  }
}

export function activeYear(e) {
  const value = String(e.active || '').slice(0, 4);
  return /^\d{4}$/.test(value) ? value : 'unknown';
}

export function scopeMatch(e, scope) {
  switch (scope) {
    case 'packaged': return e.avail;
    case 'unpacked': return !e.avail;
    case 'source': return !e.avail && Boolean(e.buildbits & B.SOURCE);
    case 'kernel': return Boolean(e.kernel);
    case 'vendor': return Boolean(e.vendor);
    case 'contrib': return Boolean(e.flags & F.CONTRIB);
    default: return true;
  }
}

const OP_RE = /^(cat|cate|category|repo|license|lang|lng|pg|os|target|type|kind|life|lifecycle|kernel|vendor|tag|tags|pkg|package|project|cap|capability|feature|build|builder|doc|docs|rel|relation|dependency|pgrx|active|year|is):(.+)$/i;

export function effectiveFilters() {
  const f = {
    cat: S.cat, repo: S.repo, license: S.license, lang: S.lang, pg: S.pg, os: S.os,
    kind: S.kind, lifecycle: S.lifecycle, scope: S.scope, vendorQ: S.vendor.toLowerCase(), kernelQ: S.kernel.toLowerCase(),
    tag: S.tag, pkg: S.pkg, capability: S.capability, build: S.build, docs: S.docs,
    relation: S.relation, pgrx: S.pgrx, active: S.active, words: []
  };
  for (const tok of S.q.trim().split(/\s+/).filter(Boolean)) {
    const m = tok.match(OP_RE);
    if (!m) { f.words.push(tok.toLowerCase()); continue; }
    const key = m[1].toLowerCase(), val = m[2];
    if (key === 'cat' || key === 'cate' || key === 'category') f.cat = val.toUpperCase();
    else if (key === 'repo') f.repo = val.toUpperCase();
    else if (key === 'license') f.license = val;
    else if (key === 'lang' || key === 'lng') f.lang = val;
    else if (key === 'pg') f.pg = val;
    else if (key === 'os' || key === 'target') f.os = val;
    else if (key === 'type' || key === 'kind') f.kind = val.toLowerCase();
    else if (key === 'life' || key === 'lifecycle') f.lifecycle = val.toLowerCase();
    else if (key === 'kernel') f.kernelQ = val.toLowerCase();
    else if (key === 'vendor') f.vendorQ = val.toLowerCase();
    else if (key === 'tag' || key === 'tags') f.tag = val;
    else if (key === 'pkg' || key === 'package' || key === 'project') f.pkg = val;
    else if (key === 'cap' || key === 'capability' || key === 'feature') f.capability = val.toLowerCase();
    else if (key === 'build' || key === 'builder') f.build = val.toLowerCase();
    else if (key === 'doc' || key === 'docs') f.docs = val.toLowerCase();
    else if (key === 'rel' || key === 'relation' || key === 'dependency') f.relation = val.toLowerCase();
    else if (key === 'pgrx') f.pgrx = val;
    else if (key === 'active' || key === 'year') f.active = val.toLowerCase();
    else if (key === 'is') {
      const v = val.toLowerCase();
      if (v === 'packaged') f.scope = 'packaged';
      if (v === 'unpacked' || v === 'unpackaged') f.scope = 'unpacked';
      if (v === 'source' || v === 'source-only') f.scope = 'source';
      if (v === 'kernel') f.scope = 'kernel';
      if (v === 'vendor' || v === 'cloud') f.scope = 'vendor';
      if (v === 'contrib') f.scope = 'contrib';
      if (['binary', 'library', 'create-extension', 'preload', 'trusted', 'relocatable'].includes(v)) f.capability = v;
    }
  }
  return f;
}

export function runFilter() {
  const f = effectiveFilters();
  const pgs = parsePGs(f.pg);
  const fRepo = f.repo.toLowerCase(), fLic = f.license.toLowerCase(), fLang = f.lang.toLowerCase();
  const fKind = f.kind.toLowerCase(), fLife = f.lifecycle.toLowerCase(), fTag = f.tag.toLowerCase();
  const fPkg = f.pkg.toLowerCase(), fDocs = f.docs.toLowerCase(), fPgrx = f.pgrx.toLowerCase(), fAct = f.active.toLowerCase();
  const list = [];
  for (const e of catalog.ext) {
    if (f.cat && e.cat !== f.cat) continue;
    if (fRepo && e.repo.toLowerCase() !== fRepo) continue;
    if (fLic && e.license.toLowerCase() !== fLic) continue;
    if (fLang && e.lang.toLowerCase() !== fLang) continue;
    if (pgs.length && !pgs.every(pg => e.pg.includes(pg))) continue;
    if (f.os && !targetsAvailable(e, pgs, f.os)) continue;
    if (fKind && e.kind.toLowerCase() !== fKind) continue;
    if (fLife && e.lifecycle.toLowerCase() !== fLife) continue;
    if (fTag && !e.tags.some(tag => tag.toLowerCase() === fTag)) continue;
    if (fPkg && e.pkg.toLowerCase() !== fPkg) continue;
    if (f.capability && !capabilityMatches(e, f.capability)) continue;
    if (f.build && !buildMatches(e, f.build)) continue;
    if (fDocs && docsClass(e) !== fDocs) continue;
    if (f.relation && !relationMatches(e, f.relation)) continue;
    if (fPgrx && e.pgrx.toLowerCase() !== fPgrx) continue;
    if (fAct && activeYear(e) !== fAct) continue;
    if (f.scope && !scopeMatch(e, f.scope)) continue;
    if (f.vendorQ && !e.vendor.toLowerCase().includes(f.vendorQ)) continue;
    if (f.kernelQ && !e.kernel.toLowerCase().includes(f.kernelQ)) continue;
    let score = 0, drop = false;
    for (const w of f.words) {
      const nm = e.name.toLowerCase();
      if (nm === w) score += 120;
      else if (nm.startsWith(w)) score += 60;
      else if (nm.includes(w) || e.pkg.toLowerCase().includes(w)) score += 30;
      else if (e.tags.some(tag => tag.toLowerCase().includes(w))) score += 18;
      else if (e.en.toLowerCase().includes(w) || (e.zh && e.zh.includes(w)) || e.comment.toLowerCase().includes(w)) score += 10;
      else if (e.cat.toLowerCase() === w || e.vendor.toLowerCase().includes(w) || e.kernel.toLowerCase().includes(w) || e.lifecycle.toLowerCase().includes(w)) score += 8;
      else { drop = true; break; }
    }
    if (drop) continue;
    list.push([score, e]);
  }
  const hasQ = f.words.length > 0;
  const sort = effSort();
  list.sort((a, b) => {
    if (hasQ && b[0] !== a[0]) return b[0] - a[0];
    if (sort === 'stars') return starOf(b[1]) - starOf(a[1]) || a[1].name.localeCompare(b[1].name);
    if (sort === 'updated') return (b[1].active || '').localeCompare(a[1].active || '') || starOf(b[1]) - starOf(a[1]) || a[1].name.localeCompare(b[1].name);
    return a[1].name.localeCompare(b[1].name);
  });
  return { f, list: list.map(x => x[1]) };
}

export function projectGroups(list) {
  const groups = new Map();
  for (const e of list) {
    if (!groups.has(e.pkg)) groups.set(e.pkg, []);
    groups.get(e.pkg).push(e);
  }
  const projects = [...groups.entries()].map(([pkg, matched]) => {
    const all = catalog.byPkg.get(pkg) || matched;
    const leadName = (all[0] && all[0].lead) || pkg;
    const lead = catalog.byName.get(leadName) || matched[0] || all[0];
    return { pkg, matched, all, lead };
  });
  const sort = effSort();
  projects.sort((a, b) => {
    if (sort === 'stars') return starOf(b.lead) - starOf(a.lead) || a.pkg.localeCompare(b.pkg);
    if (sort === 'updated') return (b.lead.active || '').localeCompare(a.lead.active || '') || a.pkg.localeCompare(b.pkg);
    return a.pkg.localeCompare(b.pkg);
  });
  return projects;
}

export function dimValues(dim, subset) {
  const cnt = new Map();
  const add = (k, e) => {
    if (!k) return;
    if (!cnt.has(k)) cnt.set(k, { n: 0, ex: [] });
    const o = cnt.get(k);
    o.n++;
    if (o.ex.length < 4) o.ex.push(e.name);
  };
  for (const e of (subset || catalog.ext)) {
    const values = [];
    if (dim === 'category') values.push(e.cat);
    else if (dim === 'tag') values.push(...e.tags);
    else if (dim === 'package') values.push(e.pkg);
    else if (dim === 'license') values.push(e.license);
    else if (dim === 'lang') values.push(e.lang);
    else if (dim === 'repo') values.push(e.repo === 'n/a' ? '' : e.repo);
    else if (dim === 'vendor') values.push(e.vendor);
    else if (dim === 'kernel') values.push(e.kernel);
    else if (dim === 'lifecycle') values.push(e.lifecycle);
    else if (dim === 'kind') values.push(e.kind);
    else if (dim === 'pg') values.push(...e.pg.map(String));
    else if (dim === 'os' && catalog.oss.length) values.push(...e.targets.map(idx => catalog.oss[idx % catalog.oss.length]));
    else if (dim === 'distribution') {
      if (e.avail) values.push('packaged');
      if (!e.avail && (e.buildbits & B.SOURCE)) values.push('source');
      if (e.vendor) values.push('vendor');
      if (e.kernel) values.push('kernel');
      if (e.flags & F.CONTRIB) values.push('contrib');
    } else if (dim === 'build') {
      for (const value of ['rpm', 'deb', 'pgrx', 'source', 'packaged']) if (buildMatches(e, value)) values.push(value);
    } else if (dim === 'capability') {
      for (const value of ['binary', 'library', 'create-extension', 'preload', 'trusted', 'relocatable']) if (capabilityMatches(e, value)) values.push(value);
    } else if (dim === 'docs') values.push(docsClass(e));
    else if (dim === 'relation') {
      for (const value of ['requires', 'required-by', 'see-also']) if (relationMatches(e, value)) values.push(value);
      if (!e.relationbits) values.push('no-relations');
    } else if (dim === 'pgrx') values.push(e.pgrx);
    else if (dim === 'activity') values.push(activeYear(e));
    for (const value of new Set(values.filter(Boolean))) add(value, e);
  }
  const vals = [...cnt.entries()].map(([v, o]) => ({ v, ...o }));
  if (dim === 'category') vals.sort((a, b) => catalog.catOrder.indexOf(a.v) - catalog.catOrder.indexOf(b.v));
  else if (dim === 'pg') vals.sort((a, b) => Number(b.v) - Number(a.v));
  else if (dim === 'os') vals.sort((a, b) => catalog.oss.indexOf(a.v) - catalog.oss.indexOf(b.v));
  else if (dim === 'activity') vals.sort((a, b) => a.v === 'unknown' ? 1 : b.v === 'unknown' ? -1 : b.v.localeCompare(a.v));
  else if (dim === 'pgrx') vals.sort((a, b) => b.v.localeCompare(a.v, undefined, { numeric: true }));
  else vals.sort((a, b) => b.n - a.n || a.v.localeCompare(b.v));
  return vals;
}

export function buildSQL(f, n) {
  const wh = [];
  const lit = v => "'" + String(v).replace(/'/g, "''") + "'";
  const hay = "concat_ws(' ', name, pkg, en_desc, zh_desc, array_to_string(tags, ' '))";
  for (const word of f.words) wh.push(hay + ' ILIKE ' + lit('%' + word + '%'));
  if (f.cat) wh.push('category = ' + lit(f.cat));
  if (f.repo) wh.push("COALESCE(extra->>'repo', rpm_repo, deb_repo) = " + lit(f.repo));
  if (f.license) wh.push(f.license.toLowerCase() === 'unknown' ? "COALESCE(NULLIF(license, ''), 'Unknown') = 'Unknown'" : 'license = ' + lit(f.license));
  if (f.lang) wh.push('lang = ' + lit(f.lang));
  const pgs = parsePGs(f.pg);
  if (pgs.length) wh.push("pg_ver @> '{" + pgs.join(',') + "}'::text[]");
  if (f.os) {
    if (pgs.length) {
      for (const pg of pgs) {
        wh.push("EXISTS (SELECT 1 FROM pgext.pkg p WHERE p.pkg = universe.pkg AND p.state = 'AVAIL' AND p.os = " + lit(f.os) + ' AND p.pg = ' + pg + ')');
      }
    } else {
      wh.push("EXISTS (SELECT 1 FROM pgext.pkg p WHERE p.pkg = universe.pkg AND p.state = 'AVAIL' AND p.os = " + lit(f.os) + ')');
    }
  }
  if (f.kind) wh.push('kind = ' + lit(f.kind));
  if (f.lifecycle) wh.push('lifecycle = ' + lit(f.lifecycle));
  if (f.kernelQ) wh.push('kernel ILIKE ' + lit('%' + f.kernelQ + '%'));
  if (f.vendorQ) wh.push('vendor ILIKE ' + lit('%' + f.vendorQ + '%'));
  if (f.tag) wh.push('EXISTS (SELECT 1 FROM unnest(tags) AS t(tag) WHERE lower(t.tag) = lower(' + lit(f.tag) + '))');
  if (f.pkg) wh.push('pkg = ' + lit(f.pkg));
  if (f.capability) {
    const expr = {
      binary: 'has_bin', bin: 'has_bin', library: 'has_lib', lib: 'has_lib',
      'create-extension': 'need_ddl', ddl: 'need_ddl', preload: 'need_load', load: 'need_load',
      trusted: 'trusted', relocatable: 'relocatable'
    }[f.capability];
    if (expr) wh.push(expr);
  }
  if (f.build) {
    const expr = {
      rpm: 'rpm_build', deb: 'deb_build', pgrx: 'pgrx_ver IS NOT NULL',
      source: '(repo_url IS NOT NULL OR tarball IS NOT NULL)', packaged: 'packaged', 'binary-package': 'packaged'
    }[f.build];
    if (expr) wh.push(expr);
  }
  if (f.docs) {
    const en = "EXISTS (SELECT 1 FROM pgext.doc d WHERE d.ext = universe.name AND NULLIF(btrim(d.en_doc), '') IS NOT NULL)";
    const zh = "EXISTS (SELECT 1 FROM pgext.doc d WHERE d.ext = universe.name AND NULLIF(btrim(d.zh_doc), '') IS NOT NULL)";
    if (f.docs === 'bilingual') wh.push(en, zh);
    if (f.docs === 'english-only') wh.push(en, 'NOT ' + zh);
    if (f.docs === 'chinese-only') wh.push(zh, 'NOT ' + en);
    if (f.docs === 'undocumented') wh.push('NOT ' + en, 'NOT ' + zh);
  }
  if (f.relation) {
    const expr = {
      requires: 'COALESCE(cardinality(requires), 0) > 0',
      'required-by': 'COALESCE(cardinality(required_by), 0) > 0',
      'see-also': 'COALESCE(cardinality(see_also), 0) > 0',
      'no-relations': 'COALESCE(cardinality(requires), 0) = 0 AND COALESCE(cardinality(required_by), 0) = 0 AND COALESCE(cardinality(see_also), 0) = 0'
    }[f.relation];
    if (expr) wh.push('(' + expr + ')');
  }
  if (f.pgrx) wh.push('pgrx_ver = ' + lit(f.pgrx));
  if (f.active === 'unknown') wh.push('last_active IS NULL');
  else if (/^\d{4}$/.test(f.active)) wh.push('left(last_active::text, 4) = ' + lit(f.active));
  if (f.scope === 'packaged') wh.push('packaged');
  if (f.scope === 'unpacked') wh.push('NOT packaged');
  if (f.scope === 'source') wh.push('NOT packaged AND (repo_url IS NOT NULL OR tarball IS NOT NULL)');
  if (f.scope === 'kernel') wh.push('kernel IS NOT NULL');
  if (f.scope === 'vendor') wh.push('vendor IS NOT NULL');
  if (f.scope === 'contrib') wh.push('contrib');
  const order = effSort() === 'stars' ? ['stars', ' DESC NULLS LAST']
    : effSort() === 'updated' ? ['last_active', ' DESC NULLS LAST'] : ['name', ''];
  const sql = 'SELECT * FROM pgext.universe' + (wh.length ? ' WHERE ' + wh.join(' AND ') : '') + ' ORDER BY ' + order[0] + order[1] + ';';
  const kw = s => '<span class="kw">' + s + '</span>';
  let html = kw('SELECT') + ' * ' + kw('FROM') + ' pgext.universe';
  if (wh.length) html += ' ' + kw('WHERE') + ' ' + esc(wh.join(' AND ')).replace(/&#39;([^&]*)&#39;/g, '<span class="lit">&#39;$1&#39;</span>');
  html += ' ' + kw('ORDER BY') + ' ' + order[0] + (order[1] ? '<span class="kw">' + order[1] + '</span>' : '') + ';';
  const rows = n === 1 ? '(1 row)' : '(' + fmtInt(n) + ' rows)';
  return { sql, html: '<span class="sql-q">' + html + '</span><span class="rcount">' + rows + '</span>' };
}

export function navSuggest(query) {
  const w = query.trim().toLowerCase();
  if (!w) return [];
  const scored = [];
  for (const e of catalog.ext) {
    const n = e.name.toLowerCase(), p = e.pkg.toLowerCase();
    let score = 0;
    if (n === w || p === w) score = 200;
    else if (n.startsWith(w)) score = 120;
    else if (p.startsWith(w)) score = 90;
    else if (n.includes(w)) score = 60;
    else if (p.includes(w)) score = 40;
    else if (e.tags.some(tag => tag.toLowerCase().startsWith(w))) score = 24;
    else continue;
    score += Math.min((e.stars || 0) / 800, 24) + (e.avail ? 8 : 0);
    scored.push([score, e]);
  }
  scored.sort((a, b) => b[0] - a[0] || a[1].name.localeCompare(b[1].name));
  return scored.slice(0, 8).map(x => x[1]);
}
