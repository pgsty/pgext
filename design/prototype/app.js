'use strict';
/* ================================================================
   PGEXT.CLOUD — prototype SPA
   data: window.__EXT__ (positional rows), __CATS__, __DOCS__, __META__
   ================================================================ */

/* ---------------- data decode ---------------- */
const CAT_ORDER = ['TIME','GIS','RAG','FTS','OLAP','FEAT','LANG','TYPE','UTIL','FUNC','ADMIN','STAT','SEC','FDW','SIM','ETL'];
const F = { LEAD:1, CONTRIB:2, BIN:4, LIB:8, DDL:16, LOAD:32, TRUSTED:64, RELOC:128 };
const PGS = [18,17,16,15,14];

const nums = a => (a || []).map(Number).filter(n => Number.isFinite(n));
const EXT = window.__EXT__.map((r, i) => ({
  i, name: r[0], pkg: r[1] || r[0], lead: r[2] || r[0], cat: r[3], avail: !!r[4],
  url: r[5] && r[5].startsWith('gh:') ? 'https://github.com/' + r[5].slice(3) : r[5],
  license: r[6], lang: r[7], repo: r[8], ver: r[9], flags: r[10] | 0,
  schemas: r[11] || [], pg: nums(r[12]), requires: r[13] || [], seeAlso: r[14] || [],
  rpm: r[17] ? { ver: r[15] || r[9], repo: r[16] || r[8], pkg: r[17], pg: nums(r[18] || r[12]) } : null,
  deb: r[21] ? { ver: r[19] || r[9], repo: r[20] || r[8], pkg: r[21], pg: nums(r[22] || r[12]) } : null,
  type: r[23], kernel: r[24], vendor: r[25], stars: r[26], commit: r[27],
  en: r[28] || '', zh: r[29] || ''
}));

const byName = new Map(EXT.map(e => [e.name, e]));
const requireBy = new Map();   // name -> [names]
const byPkg = new Map();       // pkg  -> [ext]
for (const e of EXT) {
  for (const req of e.requires) {
    if (!requireBy.has(req)) requireBy.set(req, []);
    requireBy.get(req).push(e.name);
  }
  if (!byPkg.has(e.pkg)) byPkg.set(e.pkg, []);
  byPkg.get(e.pkg).push(e);
}

const CATS = {};               // code -> {code, count, zh, en}
for (const [, name, zh, en] of window.__CATS__) CATS[name] = { code: name, zh, en, count: 0 };
for (const e of EXT) if (CATS[e.cat]) CATS[e.cat].count++;

const DOCS = window.__DOCS__ || {};
const META = window.__META__ || {};

const N_ALL = EXT.length;
const N_AVAIL = EXT.filter(e => e.avail).length;
const N_VENDOR = EXT.filter(e => e.kernel || e.vendor).length;

/* universe field: dot order = canonical category, stars desc inside */
const UFIELD = [];
for (const c of CAT_ORDER) {
  const members = EXT.filter(e => e.cat === c);
  members.sort((a, b) => (b.stars || 0) - (a.stars || 0) || a.name.localeCompare(b.name));
  UFIELD.push(...members);
}

/* ---------------- i18n ---------------- */
const CAT_NAMES = {
  TIME: ['Time-Series & Temporal', '时序与时态'],
  GIS:  ['Geospatial', '地理空间'],
  RAG:  ['AI & Vectors', 'AI 与向量'],
  FTS:  ['Full-Text Search', '全文检索'],
  OLAP: ['Analytics & Columnar', '分析与列存'],
  FEAT: ['New Capabilities', '功能特性'],
  LANG: ['Procedural Languages', '过程语言'],
  TYPE: ['Data Types', '数据类型'],
  UTIL: ['Utilities', '实用工具'],
  FUNC: ['Functions & Aggregates', '函数与聚合'],
  ADMIN:['Administration', '管理运维'],
  STAT: ['Observability', '监控统计'],
  SEC:  ['Security', '安全'],
  FDW:  ['Foreign Data Wrappers', '外部数据源'],
  SIM:  ['Compatibility', '兼容仿真'],
  ETL:  ['Replication & ETL', '复制与数据流']
};

const I18N = {
  'nav.ext': ['Extensions', '扩展目录'],
  'nav.browse': ['Browse', '多维索引'],
  'nav.about': ['About', '关于'],
  'nav.lang': ['中文', 'EN'],
  'hero.eyebrow': ['pgext.cloud · the postgresql extension catalog', 'pgext.cloud · PostgreSQL 扩展目录'],
  'hero.title': ['Everything <em>PostgreSQL</em> can become.', 'PostgreSQL 的<em>一切可能</em>。'],
  'hero.sub': ['The catalog of the PostgreSQL extension universe — <b>{all}</b> known extensions, <b>{avail}</b> of them packaged, documented, and one <code>CREATE EXTENSION</code> away.',
               'PostgreSQL 扩展宇宙的完整目录——收录 <b>{all}</b> 个已知扩展，其中 <b>{avail}</b> 个已打包、有文档，一条 <code>CREATE EXTENSION</code> 即可拥有。'],
  'hero.s1': ['extensions catalogued', '个扩展收录在册'],
  'hero.s2': ['packaged & installable', '个已打包可安装'],
  'hero.s3': ['categories', '个功能分类'],
  'hero.s4': ['cloud-vendor exclusives', '个云厂商专属'],
  'field.note': ['Every dot is one extension, colored by category — hover to identify, click to open.', '每一个点都是一个扩展，颜色代表分类——悬停查看，点击进入。'],
  'search.ph': ['Search {n} extensions — name, purpose, or try cat:GIS lang:rust pg:18', '搜索 {n} 个扩展——名称、用途，或试试 cat:GIS lang:rust pg:18'],
  'search.copy': ['click to copy query', '点击复制查询'],
  'search.copied': ['copied ✓', '已复制 ✓'],
  'chip.all': ['All', '全部'],
  'chip.packaged': ['Packaged', '已打包'],
  'chip.cloud': ['Cloud-only', '云厂商专属'],
  'sel.cat': ['category', '分类'],
  'sel.license': ['license', '许可证'],
  'sel.lang': ['language', '语言'],
  'sel.pg': ['pg version', 'PG 版本'],
  'sel.type': ['packaging', '形态'],
  'sort.stars': ['sort: stars', '排序：星标'],
  'sort.name': ['sort: name', '排序：名称'],
  'sort.updated': ['sort: updated', '排序：更新'],
  'view.grid': ['grid', '卡片'],
  'view.rows': ['rows', '行'],
  'wall.more': ['Show {n} more', '再显示 {n} 个'],
  'wall.empty': ['No extension matches. Loosen a filter, or try <code>vector</code>, <code>gis</code>, <code>parquet</code>.',
                 '没有匹配的扩展。放宽筛选条件，或试试 <code>vector</code>、<code>gis</code>、<code>parquet</code>。'],
  'rows.name': ['name', '名称'], 'rows.cat': ['category', '分类'], 'rows.ver': ['version', '版本'],
  'rows.license': ['license', '许可证'], 'rows.lang': ['lang', '语言'], 'rows.pg': ['pg', 'PG'],
  'rows.stars': ['stars', '星标'], 'rows.desc': ['description', '描述'], 'rows.repo': ['repo', '仓库'],
  'ext.crumb': ['extensions', '扩展'],
  'ext.install': ['Install', '安装'],
  'ext.avail': ['Availability', '可用性矩阵'],
  'ext.deps': ['Dependencies & Relations', '依赖与关联'],
  'ext.docs': ['Usage', '用法'],
  'ext.related': ['More in {cat}', '更多{cat}扩展'],
  'ext.spec': ['Spec sheet', '规格表'],
  'ext.attrs': ['Attributes', '属性'],
  'ext.requires': ['requires', '依赖于'],
  'ext.requiredby': ['required by', '被依赖'],
  'ext.seealso': ['see also', '另请参阅'],
  'ext.siblings': ['same package', '同包扩展'],
  'ext.none': ['none', '无'],
  'ext.docsmissing': ['Full bilingual usage docs exist for all {n} packaged extensions on the production site — this prototype bundles them for 12 showcase extensions only. Meanwhile: <code>CREATE EXTENSION {name};</code>',
                      '正式站为全部 {n} 个已打包扩展提供完整双语文档——本原型仅内置 12 个演示扩展的文档。先来一条：<code>CREATE EXTENSION {name};</code>'],
  'ext.vendorNote': ['This extension ships only inside {vendor} managed services — it is not installable from public repositories.',
                     '该扩展仅随 {vendor} 云托管服务提供，无法从公开仓库安装。'],
  'spec.version': ['version', '版本'], 'spec.state': ['state', '状态'], 'spec.category': ['category', '分类'],
  'spec.license': ['license', '许可证'], 'spec.language': ['language', '语言'], 'spec.repo': ['repo', '仓库'],
  'spec.package': ['package', '包名'], 'spec.schemas': ['schemas', '模式'], 'spec.pg': ['pg support', 'PG 支持'],
  'spec.rpm': ['rpm', 'RPM'], 'spec.deb': ['deb', 'DEB'], 'spec.updated': ['updated', '更新于'],
  'spec.source': ['source', '源码'], 'spec.vendor': ['vendor', '厂商'], 'spec.kernel': ['kernel', '内核'],
  'spec.stars': ['github', 'GitHub'],
  'attr.ddl': ['CREATE EXTENSION', 'CREATE EXTENSION'],
  'attr.load': ['needs preload', '需要预加载'],
  'attr.lib': ['ships .so library', '携带动态库'],
  'attr.bin': ['ships binaries', '携带命令行工具'],
  'attr.trusted': ['trusted', '非超户可装'],
  'attr.reloc': ['relocatable', '模式可迁移'],
  'attr.lead': ['lead of its package', '包内主扩展'],
  'attr.contrib': ['postgres contrib', 'PG 自带'],
  'yes': ['yes', '是'], 'no': ['no', '否'],
  'state.avail': ['available', '可用'], 'state.na': ['not packaged', '未打包'],
  'matrix.ext': ['CREATE EXTENSION', 'CREATE EXTENSION'],
  'matrix.rpm': ['RPM · EL 8/9/10', 'RPM · EL 8/9/10'],
  'matrix.deb': ['DEB · Debian/Ubuntu', 'DEB · Debian/Ubuntu'],
  'matrix.na': ['not packaged for this family', '该系列未打包'],
  'cat.crumb': ['categories', '分类'],
  'cat.featured': ['Featured', '精选'],
  'cat.all': ['All {n} extensions', '全部 {n} 个扩展'],
  'cat.open': ['Open in search ↗', '在搜索中打开 ↗'],
  'browse.title': ['Browse by dimension', '多维索引'],
  'browse.lede': ['One catalog, many ways in. Slice the extension universe by any dimension — every value below is a live filter.',
                  '一份目录，多个入口。任选维度切分扩展宇宙——下面每个值都是一个即时筛选器。'],
  'dim.category': ['Categories', '分类'],
  'dim.license': ['Licenses', '许可证'],
  'dim.lang': ['Languages', '实现语言'],
  'dim.repo': ['Repositories', '仓库来源'],
  'dim.vendor': ['Cloud Vendors', '云厂商'],
  'dim.pg': ['PostgreSQL Versions', 'PostgreSQL 版本'],
  'dim.type': ['Packaging Types', '扩展形态'],
  'dim.category.d': ['16 functional domains, from GIS to AI', '16 个功能领域，从 GIS 到 AI'],
  'dim.license.d': ['from PostgreSQL-liberal to commercial', '从 PostgreSQL 式宽松到商业许可'],
  'dim.lang.d': ['what extensions are written in', '扩展用什么写成'],
  'dim.repo.d': ['who packages and distributes it', '谁打包、谁分发'],
  'dim.vendor.d': ['extensions exclusive to managed clouds', '云托管服务专属扩展'],
  'dim.pg.d': ['coverage across supported majors', '各大版本的覆盖面'],
  'dim.type.d': ['standard, preload, pure-SQL, headless', 'standard / preload / puresql / headless'],
  'dim.value': ['value', '取值'], 'dim.count': ['extensions', '扩展数'], 'dim.sample': ['examples', '示例'],
  'type.standard': ['standard — CREATE EXTENSION, no preload', 'standard——直接 CREATE EXTENSION，无需预加载'],
  'type.preload': ['preload — needs shared_preload_libraries', 'preload——需要 shared_preload_libraries'],
  'type.puresql': ['puresql — SQL objects only, no binary', 'puresql——纯 SQL 对象，无二进制'],
  'type.headless': ['headless — library only, no SQL objects', 'headless——只有库，无 SQL 对象'],
  'about.title': ['About this catalog', '关于本目录'],
  'about.lede': ['PGEXT.CLOUD is the census of the PostgreSQL extension ecosystem — built from a live metadata pipeline, not a wiki.',
                 'PGEXT.CLOUD 是 PostgreSQL 扩展生态的全量普查——由实时元数据管线生成，而不是一个靠人肉维护的 wiki。'],
  'about.p1': ['Every entry is drawn from the <code>pgext</code> catalog database: upstream repositories are crawled, RPM/DEB package indexes are parsed, GitHub activity is sampled, and 555 packaged extensions carry hand-curated bilingual documentation.',
               '每个条目都来自 <code>pgext</code> 目录数据库：抓取上游仓库、解析 RPM/DEB 包索引、采样 GitHub 活跃度；555 个已打包扩展配有人工撰写的中英双语文档。'],
  'about.p2': ['This page is a static prototype rendered from a database snapshot. The production plan replaces it with <code>pgext server</code> — the same catalog, served live.',
               '本页是由数据库快照渲染的静态原型。正式版将由 <code>pgext server</code> 驱动——同一份目录，实时供给。'],
  'about.roadmap': ['Roadmap', '路线图'],
  'about.r1': ['Availability Matrix — every extension × OS × PG version, explorable', '可用性矩阵——扩展 × 操作系统 × PG 版本，全景可查'],
  'about.r2': ['Galaxy — the dependency graph as a navigable star map', 'Galaxy——把依赖关系画成可漫游的星图'],
  'about.r3': ['Live search API & dynamic pages via pgext server', 'pgext server 提供实时搜索 API 与动态页面'],
  'about.sources': ['Data sources', '数据来源'],
  'about.s1': ['pgext.universe — the curated catalog of 1,633 extensions', 'pgext.universe——1,633 个扩展的策展目录'],
  'about.s2': ['PGDG & Pigsty repositories — RPM / DEB package metadata', 'PGDG 与 Pigsty 仓库——RPM / DEB 包元数据'],
  'about.s3': ['GitHub — stars, forks, freshness signals', 'GitHub——星标、分支与活跃度信号'],
  'about.colophon': ['Prototype snapshot generated {date}. Twelve showcase extensions carry full bilingual docs; the rest render live from catalog metadata.',
                     '原型快照生成于 {date}。12 个演示扩展内置完整双语文档，其余条目由目录元数据实时渲染。'],
  'notfound.hint': ['HINT:  Check the spelling, or search the catalog below.', 'HINT:  检查拼写，或回到目录搜索。'],
  'notfound.back': ['← back to the catalog', '← 返回目录'],
  'footer.built': ['Built from the pgext catalog database · snapshot {date}', '由 pgext 目录数据库生成 · 快照 {date}'],
  'footer.pigsty': ['Battery-Included PostgreSQL Distribution', '开箱即用的 PostgreSQL 发行版'],
  'commit.freshness': ['last commit {d}', '最近提交 {d}']
};

let LANG = 'en';
try { LANG = localStorage.getItem('pgext.lang') || 'en'; } catch (e) {}
try { const q = new URLSearchParams(location.search).get('lang'); if (q === 'zh' || q === 'en') LANG = q; } catch (e) {}
const t = (k, vars) => {
  const pair = I18N[k]; let s = pair ? pair[LANG === 'zh' ? 1 : 0] : k;
  if (vars) for (const [kk, vv] of Object.entries(vars)) s = s.replaceAll('{' + kk + '}', vv);
  return s;
};
const catName = c => { const p = CAT_NAMES[c]; return p ? p[LANG === 'zh' ? 1 : 0] : c; };
const desc = e => (LANG === 'zh' && e.zh) ? e.zh : (e.en || e.zh || '');
const catDesc = c => { const m = CATS[c]; return m ? (LANG === 'zh' ? m.zh : m.en) : ''; };

/* ---------------- utils ---------------- */
const esc = s => String(s == null ? '' : s).replace(/[&<>"']/g, m => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[m]));
const fmtNum = n => n == null ? '—' : n >= 1000 ? (n / 1000).toFixed(n >= 10000 ? 0 : 1) + 'k' : String(n);
const fmtInt = n => (n == null ? '—' : String(n).replace(/\B(?=(\d{3})+(?!\d))/g, ','));
const pgRange = arr => { if (!arr || !arr.length) return '—'; const s = [...arr].sort((a, b) => a - b); return s.length > 1 ? s[0] + '–' + s[s.length - 1] : String(s[0]); };
const extHref = n => '#/ext/' + encodeURIComponent(n);
const catVar = c => 'style="--seg:var(--c-' + esc(c) + ')"';
const debounce = (fn, ms) => { let h; return (...a) => { clearTimeout(h); h = setTimeout(() => fn(...a), ms); }; };

function copyText(txt, cb) {
  const done = ok => cb && cb(ok);
  if (navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(txt).then(() => done(true), () => done(false));
  } else {
    const ta = document.createElement('textarea');
    ta.value = txt; ta.style.position = 'fixed'; ta.style.opacity = '0';
    document.body.appendChild(ta); ta.select();
    let ok = false; try { ok = document.execCommand('copy'); } catch (e) {}
    ta.remove(); done(ok);
  }
}

/* ---------------- micro markdown (for stub docs) ---------------- */
function renderMD(src) {
  const out = [];
  const lines = src.split('\n');
  let i = 0, para = [], list = null, table = null, quote = null;
  const flushP = () => { if (para.length) { out.push('<p>' + inline(para.join(' ')) + '</p>'); para = []; } };
  const flushL = () => { if (list) { out.push('<ul>' + list.map(x => '<li>' + inline(x) + '</li>').join('') + '</ul>'); list = null; } };
  const flushT = () => {
    if (!table) return;
    const head = table[0], body = table.slice(2);
    let h = '<div class="rows-scroll"><table><thead><tr>' + head.map(c => '<th>' + inline(c) + '</th>').join('') + '</tr></thead><tbody>';
    for (const row of body) h += '<tr>' + row.map(c => '<td>' + inline(c) + '</td>').join('') + '</tr>';
    out.push(h + '</tbody></table></div>'); table = null;
  };
  const flushQ = () => { if (quote) { out.push('<blockquote>' + inline(quote.join(' ')) + '</blockquote>'); quote = null; } };
  const flushAll = () => { flushP(); flushL(); flushT(); flushQ(); };
  const inline = s => esc(s)
    .replace(/`([^`]+)`/g, (m, c) => '<code>' + c + '</code>')
    .replace(/\*\*([^*]+)\*\*/g, '<b>$1</b>')
    .replace(/\[([^\]]+)\]\((https?:[^)\s]+)\)/g, '<a href="$2" target="_blank" rel="noopener">$1</a>');
  while (i < lines.length) {
    const ln = lines[i];
    if (ln.startsWith('```')) {
      flushAll();
      const buf = []; i++;
      while (i < lines.length && !lines[i].startsWith('```')) { buf.push(lines[i]); i++; }
      out.push('<pre><code>' + esc(buf.join('\n')) + '</code></pre>');
      i++; continue;
    }
    const h = ln.match(/^(#{2,4})\s+(.*)$/);
    if (h) { flushAll(); const lv = h[1].length; out.push('<h' + lv + '>' + inline(h[2]) + '</h' + lv + '>'); i++; continue; }
    if (/^\s*[-*]\s+/.test(ln)) { flushP(); flushT(); flushQ(); (list = list || []).push(ln.replace(/^\s*[-*]\s+/, '')); i++; continue; }
    if (/^\|/.test(ln)) { flushP(); flushL(); flushQ(); (table = table || []).push(ln.replace(/^\||\|$/g, '').split('|').map(s => s.trim())); i++; continue; }
    if (/^>\s?/.test(ln)) { flushP(); flushL(); flushT(); (quote = quote || []).push(ln.replace(/^>\s?/, '')); i++; continue; }
    if (!ln.trim()) { flushAll(); i++; continue; }
    flushL(); flushT(); flushQ(); para.push(ln.trim()); i++;
  }
  flushAll();
  return out.join('\n');
}

/* ---------------- landing state / filter engine ---------------- */
const DEFAULT_STATE = { q: '', cat: '', repo: '', license: '', lng: '', pg: '', type: '', scope: '', sort: 'stars', view: 'grid', limit: 96 };
let S = { ...DEFAULT_STATE };

function parseHash() {
  const raw = location.hash.slice(1) || '/';
  const qi = raw.indexOf('?');
  const path = (qi >= 0 ? raw.slice(0, qi) : raw) || '/';
  const params = new URLSearchParams(qi >= 0 ? raw.slice(qi + 1) : '');
  return { path, params };
}
let syncingHash = false;
function pushState() {
  const p = new URLSearchParams();
  for (const k of ['q', 'cat', 'repo', 'license', 'lng', 'pg', 'type', 'scope']) if (S[k]) p.set(k, S[k]);
  if (S.sort !== 'stars') p.set('sort', S.sort);
  if (S.view !== 'grid') p.set('view', S.view);
  const qs = p.toString();
  syncingHash = true;
  const url = location.pathname + location.search + '#/' + (qs ? '?' + qs : '');
  history.replaceState(null, '', url);
  setTimeout(() => { syncingHash = false; }, 0);
}
function readState(params) {
  S = { ...DEFAULT_STATE };
  for (const k of ['q', 'cat', 'repo', 'license', 'lng', 'pg', 'type', 'scope', 'sort', 'view']) {
    const v = params.get(k); if (v) S[k] = v;
  }
}

/* free text may embed operators: cat:GIS repo:pgdg lang:rust license:mit pg:18 type:preload */
function effectiveFilters() {
  const f = { cat: S.cat, repo: S.repo, license: S.license, lng: S.lng, pg: S.pg, type: S.type, scope: S.scope, words: [] };
  for (const tok of S.q.trim().split(/\s+/).filter(Boolean)) {
    const m = tok.match(/^(cat|category|repo|license|lang|lng|pg|type|vendor|is):(.+)$/i);
    if (!m) { f.words.push(tok.toLowerCase()); continue; }
    const key = m[1].toLowerCase(), val = m[2];
    if (key === 'cat' || key === 'category') f.cat = val.toUpperCase();
    else if (key === 'repo') f.repo = val.toUpperCase();
    else if (key === 'license') f.license = val;
    else if (key === 'lang' || key === 'lng') f.lng = val;
    else if (key === 'pg') f.pg = val;
    else if (key === 'type') f.type = val.toLowerCase();
    else if (key === 'vendor') f.vendorQ = val.toLowerCase();
    else if (key === 'is') { if (val === 'packaged') f.scope = 'packaged'; if (val === 'cloud') f.scope = 'cloud'; if (val === 'contrib') f.repo = 'CONTRIB'; }
  }
  return f;
}

function runFilter() {
  const f = effectiveFilters();
  const pgNum = f.pg ? parseInt(f.pg, 10) : 0;
  const list = [];
  for (const e of EXT) {
    if (f.cat && e.cat !== f.cat) continue;
    if (f.repo && e.repo !== f.repo) continue;
    if (f.license && e.license.toLowerCase() !== f.license.toLowerCase()) continue;
    if (f.lng && e.lang.toLowerCase() !== f.lng.toLowerCase()) continue;
    if (pgNum && !e.pg.includes(pgNum)) continue;
    if (f.type && e.type !== f.type) continue;
    if (f.scope === 'packaged' && !e.avail) continue;
    if (f.scope === 'cloud' && !(e.kernel || e.vendor)) continue;
    if (f.vendorQ && !((e.vendor || '').toLowerCase().includes(f.vendorQ) || (e.kernel || '').toLowerCase().includes(f.vendorQ))) continue;
    let score = 0, drop = false;
    for (const w of f.words) {
      const nm = e.name.toLowerCase(), pk = e.pkg.toLowerCase();
      if (nm === w) score += 120;
      else if (nm.startsWith(w)) score += 60;
      else if (nm.includes(w) || pk.includes(w)) score += 30;
      else if (e.en.toLowerCase().includes(w) || (e.zh && e.zh.includes(w))) score += 10;
      else if (e.cat.toLowerCase() === w || (e.vendor || '').toLowerCase().includes(w) || (e.kernel || '').toLowerCase().includes(w)) score += 8;
      else { drop = true; break; }
    }
    if (drop) continue;
    list.push([score, e]);
  }
  const hasQ = f.words.length > 0;
  list.sort((a, b) => {
    if (hasQ && b[0] !== a[0]) return b[0] - a[0];
    if (S.sort === 'name') return a[1].name.localeCompare(b[1].name);
    if (S.sort === 'updated') return (b[1].commit || '').localeCompare(a[1].commit || '') || (b[1].stars || 0) - (a[1].stars || 0);
    return (b[1].stars || 0) - (a[1].stars || 0) || a[1].name.localeCompare(b[1].name);
  });
  return { f, list: list.map(x => x[1]) };
}

/* filter state → SQL (the live readout) */
function buildSQL(f, n) {
  const wh = [];
  const lit = v => "'" + String(v).replace(/'/g, "''") + "'";
  if (f.words.length === 1) wh.push('name ~ ' + lit(f.words[0]));
  else if (f.words.length > 1) wh.push(f.words.map(w => 'name ~ ' + lit(w)).join(' AND '));
  if (f.cat) wh.push('category = ' + lit(f.cat));
  if (f.repo) wh.push('repo = ' + lit(f.repo));
  if (f.license) wh.push('license = ' + lit(f.license));
  if (f.lng) wh.push('lang = ' + lit(f.lng));
  if (f.pg) wh.push(parseInt(f.pg, 10) + ' = ANY(pg_ver)');
  if (f.type) wh.push('ext_type = ' + lit(f.type));
  if (f.scope === 'packaged') wh.push("state = 'available'");
  if (f.scope === 'cloud') wh.push('ext_vendor IS NOT NULL');
  const sql = 'SELECT * FROM pgext.universe' + (wh.length ? ' WHERE ' + wh.join(' AND ') : '') + ' ORDER BY star_cnt DESC;';
  const kw = s => '<span class="kw">' + s + '</span>';
  let html = kw('SELECT') + ' * ' + kw('FROM') + ' pgext.universe';
  if (wh.length) html += ' ' + kw('WHERE') + ' ' + esc(wh.join(' AND ')).replace(/&#39;([^&]*)&#39;/g, '<span class="lit">&#39;$1&#39;</span>');
  html += ' ' + kw('ORDER BY') + ' star_cnt <span class="kw">DESC</span>;';
  const rows = n === 1 ? '(1 row)' : '(' + fmtInt(n) + ' rows)';
  return { sql, html: '<span class="sql-q">' + html + '</span><span class="rcount">' + rows + '</span>' };
}

/* ---------------- shared components ---------------- */
function tileHTML(e) {
  return '<li><a class="tile" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '>'
    + '<span class="tile-head"><span class="tile-name">' + esc(e.name) + '</span>'
    + (e.stars ? '<span class="tile-star">★ ' + fmtNum(e.stars) + '</span>' : '') + '</span>'
    + '<span class="tile-desc">' + esc(desc(e)) + '</span>'
    + '<span class="tile-meta"><span class="cat">' + esc(e.cat) + '</span>'
    + '<span class="ver">' + esc(e.ver || '') + '</span>'
    + '<span class="lic">' + esc(e.license === 'Unknown' ? '' : e.license) + '</span></span>'
    + '</a></li>';
}

function rowHTML(e) {
  return '<tr ' + catVar(e.cat) + '><td class="r-name"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a></td>'
    + '<td class="r-cat">' + esc(e.cat) + '</td>'
    + '<td class="r-mono">' + esc(e.ver || '—') + '</td>'
    + '<td class="r-desc">' + esc(desc(e)) + '</td>'
    + '<td class="r-mono">' + esc(e.license) + '</td>'
    + '<td class="r-mono">' + esc(e.lang) + '</td>'
    + '<td class="r-mono">' + esc(pgRange(e.pg)) + '</td>'
    + '<td class="r-num">' + (e.stars ? fmtNum(e.stars) : '') + '</td></tr>';
}

function depChips(names, current) {
  const items = (names || []).filter(n => n !== current);
  if (!items.length) return '<span class="none">' + t('ext.none') + '</span>';
  return items.map(n => {
    const e = byName.get(n);
    if (!e) return '<span class="badge">' + esc(n) + '</span>';
    return '<a class="badge" href="' + extHref(n) + '" ' + catVar(e.cat) + ' data-tip="' + esc(desc(e)) + '"><span class="dot"></span>' + esc(n) + '</a>';
  }).join('');
}

function navHTML(active) {
  return '<div class="wrap nav-in">'
    + '<a class="brand" href="#/"><span class="brand-mark">\\dx</span><span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<nav class="nav-links">'
    + '<a href="#/" aria-current="' + (active === 'home') + '">' + t('nav.ext') + '</a>'
    + '<a href="#/browse" aria-current="' + (active === 'browse') + '">' + t('nav.browse') + '</a>'
    + '<a href="#/about" aria-current="' + (active === 'about') + '">' + t('nav.about') + '</a>'
    + '</nav><span class="nav-spacer"></span><div class="nav-actions">'
    + '<button class="nav-btn" id="lang-toggle" aria-label="switch language">' + t('nav.lang') + '</button>'
    + '<button class="nav-btn" id="theme-toggle" aria-label="switch theme">' + themeLabel() + '</button>'
    + '<a class="nav-btn nav-github" href="https://github.com/pgsty/pgext" target="_blank" rel="noopener" aria-label="GitHub">'
    + '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27s1.36.09 2 .27c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0 0 16 8c0-4.42-3.58-8-8-8Z"/></svg></a>'
    + '</div></div>';
}

function footerHTML() {
  const date = META.generated || '';
  return '<div class="wrap footer-in">'
    + '<span class="sig"><span class="p">pgext=#</span> \\q</span>'
    + '<span class="sig">' + t('footer.built', { date }) + '</span>'
    + '<nav><a href="https://github.com/pgsty/pgext" target="_blank" rel="noopener">GitHub</a>'
    + '<a href="https://pigsty.io" target="_blank" rel="noopener">Pigsty — ' + t('footer.pigsty') + '</a>'
    + '<a href="https://ext.pigsty.io" target="_blank" rel="noopener">ext.pigsty.io</a></nav></div>';
}

/* ---------------- view: home ---------------- */
function homeHTML() {
  const stats = [
    [fmtInt(N_ALL), t('hero.s1')],
    [fmtInt(N_AVAIL), t('hero.s2')],
    ['16', t('hero.s3')],
    [fmtInt(N_VENDOR), t('hero.s4')]
  ].map(([n, l]) => '<li><span class="num">' + n + '</span><span class="lbl">' + l + '</span></li>').join('');
  const legend = CAT_ORDER.map(c =>
    '<li style="--seg:var(--c-' + c + ')">'
    + '<button data-cat-go="' + c + '" data-tip="' + esc(catDesc(c)) + '">' + c
    + '<span class="cnt">' + CATS[c].count + '</span></button></li>'
  ).join('');
  return '<section class="hero wrap">'
    + '<p class="eyebrow">' + t('hero.eyebrow') + '</p>'
    + '<h1>' + t('hero.title') + '</h1>'
    + '<p class="hero-sub">' + t('hero.sub', { all: fmtInt(N_ALL), avail: fmtInt(N_AVAIL) }) + '</p>'
    + '<ul class="hero-stats">' + stats + '</ul>'
    + '<div class="universe"><canvas id="ufield" height="80" aria-hidden="true" style="width:100%;display:block;cursor:crosshair"></canvas>'
    + '<ul class="ubar-legend">' + legend + '</ul>'
    + '<p class="universe-note">' + t('field.note') + '</p></div>'
    + '<div class="console">'
    + '<div class="searchbox"><span class="prompt">pgext=#</span>'
    + '<input id="q" type="search" autocomplete="off" spellcheck="false" placeholder="' + esc(t('search.ph', { n: fmtInt(N_ALL) })) + '" value="' + esc(S.q) + '" aria-label="search extensions">'
    + '<kbd>/</kbd></div>'
    + '</div>'
    + '<div id="dynamic">' + dynamicHTML() + '</div>'
    + '</section>';
}

function dynamicHTML() {
  const { f, list } = runFilter();
  const { sql, html: sqlHTML } = buildSQL(f, list.length);
  const chips = [
    ['scope', '', t('chip.all') + ' ' + fmtInt(N_ALL), S.scope === ''],
    ['scope', 'packaged', t('chip.packaged') + ' ' + fmtInt(N_AVAIL), S.scope === 'packaged'],
    ['repo', 'PGDG', 'PGDG', S.repo === 'PGDG'],
    ['repo', 'PIGSTY', 'PIGSTY', S.repo === 'PIGSTY'],
    ['repo', 'CONTRIB', 'CONTRIB', S.repo === 'CONTRIB'],
    ['scope', 'cloud', t('chip.cloud') + ' ' + fmtInt(N_VENDOR), S.scope === 'cloud']
  ].map(([k, v, label, on]) =>
    '<button class="chip" data-fkey="' + k + '" data-fval="' + esc(v) + '" aria-pressed="' + on + '">' + label + '</button>').join('');

  const catOpts = '<option value="">' + t('sel.cat') + '</option>' + CAT_ORDER.map(c =>
    '<option value="' + c + '"' + (S.cat === c ? ' selected' : '') + '>' + c + ' · ' + catName(c) + '</option>').join('');
  const licCount = new Map();
  for (const e of EXT) licCount.set(e.license, (licCount.get(e.license) || 0) + 1);
  const licTop = [...licCount.entries()].sort((a, b) => b[1] - a[1]).slice(0, 14);
  const licOpts = '<option value="">' + t('sel.license') + '</option>' + licTop.map(([l]) =>
    '<option value="' + esc(l) + '"' + (S.license === l ? ' selected' : '') + '>' + esc(l) + '</option>').join('');
  const langs = [...new Set(EXT.map(e => e.lang))].filter(Boolean).sort();
  const langOpts = '<option value="">' + t('sel.lang') + '</option>' + langs.map(l =>
    '<option value="' + esc(l) + '"' + (S.lng === l ? ' selected' : '') + '>' + esc(l) + '</option>').join('');
  const pgOpts = '<option value="">' + t('sel.pg') + '</option>' + PGS.map(v =>
    '<option value="' + v + '"' + (S.pg === String(v) ? ' selected' : '') + '>PG ' + v + '</option>').join('');
  const typeOpts = '<option value="">' + t('sel.type') + '</option>' + ['standard', 'preload', 'puresql', 'headless'].map(v =>
    '<option value="' + v + '"' + (S.type === v ? ' selected' : '') + '>' + v + '</option>').join('');
  const sortOpts = [['stars', t('sort.stars')], ['name', t('sort.name')], ['updated', t('sort.updated')]].map(([v, l]) =>
    '<option value="' + v + '"' + (S.sort === v ? ' selected' : '') + '>' + l + '</option>').join('');

  const shown = list.slice(0, S.limit);
  let results;
  if (!list.length) {
    results = '<div class="result-empty">' + t('wall.empty') + '</div>';
  } else if (S.view === 'rows') {
    results = '<div class="rows"><div class="rows-scroll"><table><thead><tr>'
      + '<th>' + t('rows.name') + '</th><th>' + t('rows.cat') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th>'
      + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
      + '</tr></thead><tbody>' + shown.map(rowHTML).join('') + '</tbody></table></div></div>';
  } else {
    results = '<ul class="wall">' + shown.map(tileHTML).join('') + '</ul>';
  }
  const more = list.length > S.limit
    ? '<button class="chip load-more" data-more>' + t('wall.more', { n: fmtInt(Math.min(list.length - S.limit, 400)) }) + ' ↓</button>' : '';

  return '<button class="sql-readout" data-sql="' + esc(sql) + '" title="' + esc(t('search.copy')) + '">' + sqlHTML + '</button>'
    + '<div class="filterbar">' + chips
    + '<select data-skey="cat" class="' + (S.cat ? 'on' : '') + '">' + catOpts + '</select>'
    + '<select data-skey="license" class="' + (S.license ? 'on' : '') + '">' + licOpts + '</select>'
    + '<select data-skey="lng" class="' + (S.lng ? 'on' : '') + '">' + langOpts + '</select>'
    + '<select data-skey="pg" class="' + (S.pg ? 'on' : '') + '">' + pgOpts + '</select>'
    + '<select data-skey="type" class="' + (S.type ? 'on' : '') + '">' + typeOpts + '</select>'
    + '<span class="spacer"></span>'
    + '<select data-skey="sort">' + sortOpts + '</select>'
    + '<span class="viewtoggle"><button data-view="grid" aria-pressed="' + (S.view === 'grid') + '">' + t('view.grid') + '</button>'
    + '<button data-view="rows" aria-pressed="' + (S.view === 'rows') + '">' + t('view.rows') + '</button></span>'
    + '</div>' + results + more;
}

/* ---------------- universe field canvas ---------------- */
let ucells = null; // [{x,y,e}] geometry cache
function drawField() {
  const cv = document.getElementById('ufield');
  if (!cv) return;
  const dpr = window.devicePixelRatio || 1;
  const W = cv.clientWidth || cv.parentElement.clientWidth;
  const pitch = W > 900 ? 7 : W > 560 ? 6 : 5;
  const dot = pitch - 2;
  const cols = Math.max(24, Math.floor(W / pitch));
  let cell = 0;
  const cells = [];
  const root = getComputedStyle(document.documentElement);
  const colors = {};
  for (const c of CAT_ORDER) colors[c] = root.getPropertyValue('--c-' + c).trim() || '#888';
  let prevCat = null;
  for (const e of UFIELD) {
    if (prevCat && e.cat !== prevCat) cell += 1; // breathing gap between categories
    cells.push({ col: cell % cols, row: Math.floor(cell / cols), e });
    prevCat = e.cat; cell++;
  }
  const rows = Math.floor(cell / cols) + 1;
  const H = rows * pitch;
  cv.width = W * dpr; cv.height = H * dpr;
  cv.style.height = H + 'px';
  const ctx = cv.getContext('2d');
  ctx.scale(dpr, dpr);
  ctx.clearRect(0, 0, W, H);
  const hov = ucells && ucells.hoverName;
  for (const c of cells) {
    ctx.fillStyle = colors[c.e.cat];
    const x = c.col * pitch, y = c.row * pitch;
    if (hov && c.e.name === hov) {
      ctx.globalAlpha = 1;
      ctx.fillRect(x - 1, y - 1, dot + 2, dot + 2);
    } else {
      ctx.globalAlpha = c.e.avail ? 0.95 : 0.45;
      ctx.fillRect(x, y, dot, dot);
    }
  }
  ctx.globalAlpha = 1;
  ucells = { cells, pitch, cols, W, hoverName: hov || null };
}
function fieldHit(ev) {
  if (!ucells) return null;
  const cv = document.getElementById('ufield');
  const r = cv.getBoundingClientRect();
  const x = ev.clientX - r.left, y = ev.clientY - r.top;
  const col = Math.floor(x / ucells.pitch), row = Math.floor(y / ucells.pitch);
  for (const c of ucells.cells) if (c.col === col && c.row === row) return c.e;
  return null;
}

/* ---------------- view: extension detail ---------------- */
function attrBadges(e) {
  const items = [];
  if (e.flags & F.DDL) items.push(['attr.ddl', true]);
  if (e.flags & F.LOAD) items.push(['attr.load', true]);
  if (e.flags & F.TRUSTED) items.push(['attr.trusted', true]);
  if (e.flags & F.RELOC) items.push(['attr.reloc', true]);
  if (e.flags & F.LIB) items.push(['attr.lib', false]);
  if (e.flags & F.BIN) items.push(['attr.bin', false]);
  if (e.flags & F.CONTRIB) items.push(['attr.contrib', false]);
  return items.map(([k, hot]) => '<span class="badge' + (hot ? ' flag-on' : '') + '">' + t(k) + '</span>').join('');
}

function installHTML(e) {
  const tabs = [];
  const pgTop = e.pg.length ? Math.max(...e.pg) : 18;
  const sqlLines = [];
  if (e.flags & F.LOAD) {
    sqlLines.push('<span class="cmt">-- 1. requires preload: add to postgresql.conf, then restart</span>');
    sqlLines.push('shared_preload_libraries = \'' + esc(e.name) + '\'');
    sqlLines.push('');
  }
  if (e.flags & F.DDL) {
    const cascade = e.requires.length ? ' CASCADE' : '';
    sqlLines.push('<span class="kw">CREATE EXTENSION</span> ' + esc(e.name) + cascade + ';');
  } else {
    sqlLines.push('<span class="cmt">-- ' + esc(e.type) + ': ships no SQL objects, nothing to CREATE</span>');
  }
  tabs.push(['SQL', sqlLines.join('\n'), (e.flags & F.DDL) ? 'CREATE EXTENSION ' + e.name + (e.requires.length ? ' CASCADE' : '') + ';' : '']);
  tabs.push(['pig', '<span class="cmt"># works on EL / Debian / Ubuntu alike</span>\npig install ' + esc(e.pkg), 'pig install ' + e.pkg]);
  if (e.rpm) {
    const p = e.rpm.pkg.replaceAll('$v', String(pgTop));
    tabs.push(['yum', '<span class="cmt"># ' + esc(e.rpm.repo) + ' repo · v' + esc(e.rpm.ver) + ' · for PG ' + pgTop + '</span>\nyum install ' + esc(p), 'yum install ' + p]);
  }
  if (e.deb) {
    const p = e.deb.pkg.replaceAll('$v', String(pgTop));
    tabs.push(['apt', '<span class="cmt"># ' + esc(e.deb.repo) + ' repo · v' + esc(e.deb.ver) + ' · for PG ' + pgTop + '</span>\napt install ' + esc(p), 'apt install ' + p]);
  }
  const heads = tabs.map(([n], i) => '<button role="tab" aria-selected="' + (i === 0) + '" data-itab="' + i + '">' + n + '</button>').join('');
  const bodies = tabs.map(([, html, copy], i) =>
    '<div class="install-body" data-ipane="' + i + '"' + (i ? ' hidden' : '') + '>'
    + (copy ? '<button class="copy-btn" data-copy="' + esc(copy) + '">copy</button>' : '')
    + '<pre>' + html + '</pre></div>').join('');
  return '<div class="install"><div class="install-tabs" role="tablist">' + heads + '</div>' + bodies + '</div>';
}

function matrixHTML(e) {
  const cols = PGS.map(v => '<th>PG ' + v + '</th>').join('');
  const row = (lab, sub, pgs, pkgPattern) => {
    const cells = PGS.map(v => {
      if (pgs.includes(v)) {
        const pkg = pkgPattern ? pkgPattern.replaceAll('$v', String(v)) : '';
        return '<td class="ok" ' + (pkg ? 'data-tip="' + esc(pkg) + '"' : '') + '>✓</td>';
      }
      return '<td class="no">·</td>';
    }).join('');
    return '<tr><td class="rowlab">' + lab + (sub ? '<small>' + sub + '</small>' : '') + '</td>' + cells + '</tr>';
  };
  let rows = row(t('matrix.ext'), esc(e.name) + ' v' + esc(e.ver || '—'), e.pg, '');
  rows += e.rpm ? row(t('matrix.rpm'), esc(e.rpm.repo) + ' · v' + esc(e.rpm.ver), e.rpm.pg, e.rpm.pkg)
                : row(t('matrix.rpm'), t('matrix.na'), [], '');
  rows += e.deb ? row(t('matrix.deb'), esc(e.deb.repo) + ' · v' + esc(e.deb.ver), e.deb.pg, e.deb.pkg)
                : row(t('matrix.deb'), t('matrix.na'), [], '');
  return '<div class="matrix-scroll"><table class="matrix"><thead><tr><th></th>' + cols + '</tr></thead><tbody>' + rows + '</tbody></table></div>';
}

function specHTML(e) {
  const r = (k, v, mono) => v ? '<div class="row"><dt>' + t(k) + '</dt><dd' + (mono ? ' class="mono"' : '') + '>' + v + '</dd></div>' : '';
  const licLink = '<a href="#/?license=' + encodeURIComponent(e.license) + '">' + esc(e.license) + '</a>';
  const langLink = '<a href="#/?lng=' + encodeURIComponent(e.lang) + '">' + esc(e.lang) + '</a>';
  const catLink = '<a href="#/cat/' + e.cat + '">' + esc(e.cat) + ' · ' + esc(catName(e.cat)) + '</a>';
  const repoLink = e.repo !== 'n/a' ? '<a href="#/?repo=' + encodeURIComponent(e.repo) + '">' + esc(e.repo) + '</a>' : '';
  return '<div class="spec"><h3>' + t('ext.spec') + '</h3><dl>'
    + r('spec.version', esc(e.ver), true)
    + r('spec.state', e.avail ? '<span style="color:var(--good);font-weight:650">' + t('state.avail') + '</span>' : '<span class="muted">' + t('state.na') + '</span>')
    + r('spec.category', catLink)
    + r('spec.license', licLink)
    + r('spec.language', langLink)
    + r('spec.repo', repoLink)
    + r('spec.package', esc(e.pkg), true)
    + r('spec.pg', esc(pgRange(e.pg)), true)
    + r('spec.schemas', e.schemas.length ? esc(e.schemas.join(', ')) : '', true)
    + r('spec.rpm', e.rpm ? esc(e.rpm.pkg) + ' <span class="muted">v' + esc(e.rpm.ver) + '</span>' : '', true)
    + r('spec.deb', e.deb ? esc(e.deb.pkg) + ' <span class="muted">v' + esc(e.deb.ver) + '</span>' : '', true)
    + r('spec.vendor', e.vendor ? esc(e.vendor) + (e.kernel ? ' <span class="muted">' + esc(e.kernel) + '</span>' : '') : '', true)
    + r('spec.stars', e.stars != null ? '★ ' + fmtInt(e.stars) : '', true)
    + r('spec.updated', esc(e.commit || ''), true)
    + '</dl></div>';
}

function extHTML(name) {
  const e = byName.get(name);
  if (!e) return notFoundHTML(name);
  const siblings = (byPkg.get(e.pkg) || []).map(x => x.name);
  const reqBy = requireBy.get(e.name) || [];
  const doc = DOCS[e.name];
  const docHTML = doc
    ? '<div class="prose">' + renderMD(LANG === 'zh' ? doc[1] : doc[0]) + '</div>'
    : '<div class="docs-missing">' + t('ext.docsmissing', { n: fmtInt(N_AVAIL), name: esc(e.name) }) + '</div>';
  const related = EXT.filter(x => x.cat === e.cat && x.name !== e.name && x.avail)
    .sort((a, b) => (b.stars || 0) - (a.stars || 0)).slice(0, 4);
  const vendorNote = (e.kernel || e.vendor) && !e.avail
    ? '<p class="hero-sub" style="margin-top:10px">☁ ' + t('ext.vendorNote', { vendor: esc(e.vendor || e.kernel) }) + '</p>' : '';

  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="#/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="#/cat/' + e.cat + '" style="color:var(--c-' + e.cat + ')">' + e.cat + '</a><span class="sep">/</span>'
    + '<span class="here">' + esc(e.name) + '</span></nav>'
    + '<header><div class="ext-head"><h1>' + esc(e.name) + '</h1>' + (e.ver ? '<span class="ver">v' + esc(e.ver) + '</span>' : '') + '</div>'
    + '<p class="ext-tagline">' + esc(desc(e)) + '</p>'
    + '<div class="badge-row">'
    + '<a class="badge cat" href="#/cat/' + e.cat + '" ' + catVar(e.cat) + '><span class="dot"></span>' + e.cat + ' · ' + esc(catName(e.cat)) + '</a>'
    + '<a class="badge" href="#/?license=' + encodeURIComponent(e.license) + '">' + esc(e.license) + '</a>'
    + '<a class="badge" href="#/?lng=' + encodeURIComponent(e.lang) + '">' + esc(e.lang) + '</a>'
    + (e.repo !== 'n/a' ? '<a class="badge" href="#/?repo=' + e.repo + '">' + esc(e.repo) + '</a>' : '')
    + '<span class="badge" data-tip="' + esc(t('type.' + e.type)) + '">' + esc(e.type) + '</span>'
    + (e.vendor ? '<span class="badge">☁ ' + esc(e.vendor) + '</span>' : '')
    + '</div>'
    + '<div class="badge-row">' + attrBadges(e) + '</div>'
    + '<div class="ghstats">'
    + (e.stars != null ? '<span>★ <b>' + fmtInt(e.stars) + '</b></span>' : '')
    + (e.commit ? '<span>' + t('commit.freshness', { d: '<b>' + esc(e.commit) + '</b>' }) + '</span>' : '')
    + (e.url ? '<a href="' + esc(e.url) + '" target="_blank" rel="noopener">' + esc(e.url.replace(/^https?:\/\//, '')) + ' ↗</a>' : '')
    + '</div>' + vendorNote + '</header>'
    + '<div class="ext-grid" style="margin-top:26px"><div class="ext-main">'
    + '<div class="section" style="margin-top:0"><h2>' + t('ext.install') + '</h2>' + installHTML(e) + '</div>'
    + '<div class="section"><h2>' + t('ext.avail') + '</h2>' + matrixHTML(e) + '</div>'
    + '<div class="section"><h2>' + t('ext.deps') + '</h2>'
    + '<div class="depgroup"><div class="lab">' + t('ext.requires') + '</div><div class="dep-chips">' + depChips(e.requires, e.name) + '</div></div>'
    + '<div class="depgroup"><div class="lab">' + t('ext.requiredby') + '</div><div class="dep-chips">' + depChips(reqBy, e.name) + '</div></div>'
    + '<div class="depgroup"><div class="lab">' + t('ext.seealso') + '</div><div class="dep-chips">' + depChips(e.seeAlso, e.name) + '</div></div>'
    + (siblings.length > 1 ? '<div class="depgroup"><div class="lab">' + t('ext.siblings') + ' · ' + esc(e.pkg) + '</div><div class="dep-chips">' + depChips(siblings, e.name) + '</div></div>' : '')
    + '</div>'
    + '<div class="section"><h2>' + t('ext.docs') + '</h2>' + docHTML + '</div>'
    + (related.length ? '<div class="section"><h2>' + t('ext.related', { cat: e.cat }) + '</h2><ul class="related">' + related.map(tileHTML).join('') + '</ul></div>' : '')
    + '</div><aside class="ext-rail">' + specHTML(e) + '</aside></div>'
    + '</article>';
}

function notFoundHTML(name) {
  return '<div class="page wrap"><div class="section">'
    + '<pre class="mono" style="font-size:14px;line-height:1.8;color:var(--bad)">ERROR:  extension "' + esc(name) + '" does not exist\n'
    + '<span style="color:var(--ink-3)">' + esc(t('notfound.hint')) + '</span></pre>'
    + '<p style="margin-top:18px"><a href="#/">' + t('notfound.back') + '</a></p>'
    + '</div></div>';
}

/* ---------------- view: category ---------------- */
function catHTML(code) {
  const c = CATS[code];
  if (!c) return notFoundHTML(code);
  const members = EXT.filter(e => e.cat === code).sort((a, b) => (b.stars || 0) - (a.stars || 0) || a.name.localeCompare(b.name));
  const featured = members.filter(e => e.avail)
    .sort((a, b) => ((DOCS[b.name] ? 1 : 0) - (DOCS[a.name] ? 1 : 0)) || (b.stars || 0) - (a.stars || 0))
    .slice(0, 6);
  const strip = CAT_ORDER.map(x =>
    '<a href="#/cat/' + x + '" style="--seg:var(--c-' + x + ');flex:' + CATS[x].count + ' 1 0" aria-current="' + (x === code) + '" data-tip="' + x + ' · ' + CATS[x].count + '"></a>').join('');
  const fcards = featured.map(e =>
    '<li><a class="fcard" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '>'
    + '<span class="name">' + esc(e.name) + '<span class="ver">' + esc(e.ver || '') + '</span>'
    + (e.stars ? '<span class="star">★ ' + fmtNum(e.stars) + '</span>' : '') + '</span>'
    + '<span class="desc">' + esc(desc(e)) + '</span>'
    + '<span class="meta"><span>' + esc(e.license) + '</span><span>' + esc(e.lang) + '</span><span>PG ' + esc(pgRange(e.pg)) + '</span>'
    + (DOCS[e.name] ? '<span>📖 docs</span>' : '') + '</span>'
    + '</a></li>').join('');
  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="#/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="#/dim/category">' + t('cat.crumb') + '</a><span class="sep">/</span><span class="here">' + code + '</span></nav>'
    + '<div class="cat-strip">' + strip + '</div>'
    + '<header class="page-head cat-hero" style="--seg:var(--c-' + code + ');margin-top:22px">'
    + '<span class="code">' + code + '</span>'
    + '<h1>' + esc(catName(code)) + '</h1>'
    + '<p class="lede">' + esc(catDesc(code)) + '</p>'
    + '<p class="cat-count" style="margin-top:8px">' + fmtInt(members.length) + ' extensions · '
    + fmtInt(members.filter(e => e.avail).length) + ' packaged · <a href="#/?cat=' + code + '">' + t('cat.open') + '</a></p>'
    + '</header>'
    + '<div class="section"><h2>' + t('cat.featured') + '</h2><ul class="featured">' + fcards + '</ul></div>'
    + '<div class="section"><h2>' + t('cat.all', { n: fmtInt(members.length) }) + '</h2>'
    + '<div class="rows"><div class="rows-scroll"><table><thead><tr>'
    + '<th>' + t('rows.name') + '</th><th>' + t('rows.cat') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th>'
    + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
    + '</tr></thead><tbody>' + members.map(rowHTML).join('') + '</tbody></table></div></div></div>'
    + '</article>';
}

/* ---------------- view: browse & dimensions ---------------- */
const DIMS = {
  category: { key: 'cat', label: 'dim.category', d: 'dim.category.d' },
  license: { key: 'license', label: 'dim.license', d: 'dim.license.d' },
  lang: { key: 'lng', label: 'dim.lang', d: 'dim.lang.d' },
  repo: { key: 'repo', label: 'dim.repo', d: 'dim.repo.d' },
  vendor: { key: 'vendor', label: 'dim.vendor', d: 'dim.vendor.d' },
  pg: { key: 'pg', label: 'dim.pg', d: 'dim.pg.d' },
  type: { key: 'type', label: 'dim.type', d: 'dim.type.d' }
};

function dimValues(dim) {
  const cnt = new Map();
  const add = (k, e) => { if (!k) return; if (!cnt.has(k)) cnt.set(k, { n: 0, ex: [] }); const o = cnt.get(k); o.n++; if (o.ex.length < 4 && e.stars) o.ex.push(e.name); };
  for (const e of EXT) {
    if (dim === 'category') add(e.cat, e);
    else if (dim === 'license') add(e.license, e);
    else if (dim === 'lang') add(e.lang, e);
    else if (dim === 'repo') add(e.repo === 'n/a' ? '' : e.repo, e);
    else if (dim === 'vendor') add(e.vendor, e);
    else if (dim === 'type') add(e.type, e);
    else if (dim === 'pg') for (const v of e.pg) add('PG ' + v, e);
  }
  let vals = [...cnt.entries()].map(([v, o]) => ({ v, ...o }));
  if (dim === 'category') vals.sort((a, b) => CAT_ORDER.indexOf(a.v) - CAT_ORDER.indexOf(b.v));
  else if (dim === 'pg') vals.sort((a, b) => parseInt(b.v.slice(3), 10) - parseInt(a.v.slice(3), 10));
  else vals.sort((a, b) => b.n - a.n);
  return vals;
}

function browseHTML() {
  const cards = Object.entries(DIMS).map(([dim, cfg]) => {
    const vals = dimValues(dim).slice(0, 16);
    const max = Math.max(...vals.map(v => v.n));
    const preview = vals.map(v => {
      const seg = dim === 'category' ? 'var(--c-' + v.v + ')' : 'var(--accent)';
      return '<i style="--w:' + v.n + ';--seg:' + seg + ';opacity:' + (0.25 + 0.7 * v.n / max).toFixed(2) + '"></i>';
    }).join('');
    return '<li><a class="dim-card" href="#/dim/' + dim + '">'
      + '<span class="t">' + t(cfg.label) + '<span class="n">' + dimValues(dim).length + '</span></span>'
      + '<span class="d">' + t(cfg.d) + '</span>'
      + '<span class="preview">' + preview + '</span></a></li>';
  }).join('');
  return '<article class="page wrap"><header class="page-head">'
    + '<p class="eyebrow">' + t('nav.browse').toLowerCase() + '</p>'
    + '<h1>' + t('browse.title') + '</h1><p class="lede">' + t('browse.lede') + '</p></header>'
    + '<ul class="dims">' + cards + '</ul></article>';
}

function dimHTML(dim) {
  const cfg = DIMS[dim];
  if (!cfg) return notFoundHTML(dim);
  const vals = dimValues(dim);
  const max = Math.max(...vals.map(v => v.n));
  const rows = vals.map(o => {
    const isCat = dim === 'category';
    const href = isCat ? '#/cat/' + o.v
      : dim === 'pg' ? '#/?pg=' + o.v.replace('PG ', '')
      : dim === 'vendor' ? '#/?q=' + encodeURIComponent('vendor:' + o.v)
      : '#/?' + cfg.key + '=' + encodeURIComponent(o.v);
    const seg = isCat ? 'style="--seg:var(--c-' + o.v + ')"' : '';
    const label = isCat ? o.v + ' · ' + catName(o.v) : (dim === 'type' ? '<span class="mono">' + esc(o.v) + '</span>' : esc(o.v));
    return '<tr ' + seg + '><td class="v"><a href="' + href + '">' + label + '</a></td>'
      + '<td class="bar"><i style="width:' + Math.max(2, Math.round(100 * o.n / max)) + '%"></i></td>'
      + '<td class="n">' + fmtInt(o.n) + '</td>'
      + '<td class="ex">' + o.ex.map(esc).join(', ') + '</td></tr>';
  }).join('');
  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="#/browse">' + t('nav.browse') + '</a><span class="sep">/</span><span class="here">' + t(cfg.label) + '</span></nav>'
    + '<header class="page-head"><h1>' + t(cfg.label) + '</h1><p class="lede">' + t(cfg.d) + '</p></header>'
    + '<div class="dimtable"><table><thead><tr><th>' + t('dim.value') + '</th><th></th><th style="text-align:right">' + t('dim.count') + '</th><th>' + t('dim.sample') + '</th></tr></thead>'
    + '<tbody>' + rows + '</tbody></table></div></article>';
}

/* ---------------- view: about ---------------- */
function aboutHTML() {
  return '<article class="page wrap">'
    + '<header class="page-head"><p class="eyebrow">pgext.cloud</p>'
    + '<h1>' + t('about.title') + '</h1><p class="lede">' + t('about.lede') + '</p></header>'
    + '<div class="about-cols">'
    + '<div><h3>' + (LANG === 'zh' ? '这是什么' : 'What this is') + '</h3><p>' + t('about.p1') + '</p><p>' + t('about.p2') + '</p></div>'
    + '<div><h3>' + t('about.sources') + '</h3><ul class="roadmap">'
    + '<li><span class="tag">catalog</span>' + t('about.s1') + '</li>'
    + '<li><span class="tag">packages</span>' + t('about.s2') + '</li>'
    + '<li><span class="tag">github</span>' + t('about.s3') + '</li></ul></div>'
    + '<div><h3>' + t('about.roadmap') + '</h3><ul class="roadmap">'
    + '<li><span class="tag">phase 2</span>' + t('about.r1') + '</li>'
    + '<li><span class="tag">phase 2</span>' + t('about.r2') + '</li>'
    + '<li><span class="tag">phase 3</span>' + t('about.r3') + '</li></ul></div>'
    + '</div>'
    + '<p class="universe-note" style="margin-top:34px">' + t('about.colophon', { date: META.generated || '—' }) + '</p>'
    + '</article>';
}

/* ---------------- router ---------------- */
let currentPath = null;
function route() {
  const { path, params } = parseHash();
  const app = document.getElementById('app');
  const nav = document.getElementById('nav');
  const pathChanged = path !== currentPath;
  currentPath = path;
  let active = 'home';
  if (path === '/' || path === '') {
    readState(params);
    app.innerHTML = homeHTML();
    drawField();
    bindSearch();
  } else if (path.startsWith('/ext/')) {
    app.innerHTML = extHTML(decodeURIComponent(path.slice(5))); active = '';
  } else if (path.startsWith('/cat/')) {
    app.innerHTML = catHTML(decodeURIComponent(path.slice(5)).toUpperCase()); active = '';
  } else if (path === '/browse') {
    app.innerHTML = browseHTML(); active = 'browse';
  } else if (path.startsWith('/dim/')) {
    app.innerHTML = dimHTML(decodeURIComponent(path.slice(5))); active = 'browse';
  } else if (path === '/about') {
    app.innerHTML = aboutHTML(); active = 'about';
  } else {
    app.innerHTML = notFoundHTML(path); active = '';
  }
  nav.innerHTML = navHTML(active);
  document.getElementById('footer').innerHTML = footerHTML();
  if (pathChanged) window.scrollTo(0, 0);
  document.title = titleFor(path);
}
function titleFor(path) {
  if (path.startsWith('/ext/')) { const n = decodeURIComponent(path.slice(5)); return n + ' · PGEXT.CLOUD'; }
  if (path.startsWith('/cat/')) return decodeURIComponent(path.slice(5)).toUpperCase() + ' · PGEXT.CLOUD';
  if (path === '/browse') return t('nav.browse') + ' · PGEXT.CLOUD';
  if (path === '/about') return t('nav.about') + ' · PGEXT.CLOUD';
  return 'PGEXT.CLOUD — The PostgreSQL Extension Catalog';
}

/* re-render just the dynamic region (keeps search input focus) */
function renderDynamic() {
  const d = document.getElementById('dynamic');
  if (d) { d.innerHTML = dynamicHTML(); pushState(); }
}
function bindSearch() {
  const q = document.getElementById('q');
  if (!q) return;
  q.addEventListener('input', debounce(() => {
    S.q = q.value; S.limit = DEFAULT_STATE.limit;
    renderDynamic();
  }, 130));
  q.addEventListener('keydown', ev => {
    if (ev.key === 'Escape') { q.value = ''; S.q = ''; renderDynamic(); }
  });
}

/* ---------------- theme & lang ---------------- */
function themeLabel() {
  let mode = 'auto';
  try { mode = localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {}
  return mode === 'dark' ? '☾ dark' : mode === 'light' ? '☀ light' : '◐ auto';
}
function cycleTheme() {
  let mode = 'auto';
  try { mode = localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {}
  const next = mode === 'auto' ? 'dark' : mode === 'dark' ? 'light' : 'auto';
  try { localStorage.setItem('pgext.theme', next); } catch (e) {}
  applyTheme();
  const b = document.getElementById('theme-toggle');
  if (b) b.textContent = themeLabel();
  drawField();
}
function applyTheme() {
  let mode = 'auto';
  try { mode = localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {}
  if (mode === 'auto') document.documentElement.removeAttribute('data-theme');
  else document.documentElement.setAttribute('data-theme', mode);
}
function toggleLang() {
  LANG = LANG === 'zh' ? 'en' : 'zh';
  try { localStorage.setItem('pgext.lang', LANG); } catch (e) {}
  document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
  const y = window.scrollY;
  currentPath = null; // force full re-render
  route();
  window.scrollTo(0, y); // keep reading position across language switch
}

/* ---------------- tooltip ---------------- */
const tip = () => document.getElementById('tip');
let tipTimer = null;
function showTip(html, x, y) {
  const el = tip(); el.innerHTML = html;
  const pad = 14, vw = window.innerWidth;
  el.style.left = Math.min(x + pad, vw - el.offsetWidth - 10) + 'px';
  el.style.top = (y + pad) + 'px';
  el.classList.add('show');
  requestAnimationFrame(() => {
    el.style.left = Math.min(x + pad, vw - el.offsetWidth - 10) + 'px';
    const vh = window.innerHeight;
    el.style.top = (y + pad + el.offsetHeight > vh ? y - el.offsetHeight - 8 : y + pad) + 'px';
  });
}
function hideTip() { tip().classList.remove('show'); }

/* ---------------- global events ---------------- */
function boot() {
  applyTheme();
  document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
  route();
  window.addEventListener('hashchange', () => { if (!syncingHash) route(); });
  window.addEventListener('resize', debounce(() => drawField(), 150));

  document.addEventListener('click', ev => {
    const el = ev.target.closest('[data-fkey],[data-skey],[data-view],[data-more],[data-sql],[data-copy],[data-itab],[data-cat-go],#lang-toggle,#theme-toggle,#ufield');
    if (!el) return;
    if (el.id === 'lang-toggle') return toggleLang();
    if (el.id === 'theme-toggle') return cycleTheme();
    if (el.id === 'ufield') {
      const e = fieldHit(ev);
      if (e) location.hash = extHref(e.name);
      return;
    }
    if (el.dataset.catGo) {
      S.cat = S.cat === el.dataset.catGo ? '' : el.dataset.catGo;
      S.limit = DEFAULT_STATE.limit; renderDynamic();
      const d = document.getElementById('dynamic');
      if (d) d.scrollIntoView({ behavior: 'smooth', block: 'start' });
      return;
    }
    if (el.dataset.fkey) {
      const k = el.dataset.fkey, v = el.dataset.fval;
      if (k === 'scope') S.scope = S.scope === v ? '' : v;
      else S[k] = S[k] === v ? '' : v;
      S.limit = DEFAULT_STATE.limit; renderDynamic(); return;
    }
    if (el.dataset.view) { S.view = el.dataset.view; renderDynamic(); return; }
    if (el.dataset.more !== undefined) { S.limit += 400; renderDynamic(); return; }
    if (el.dataset.sql) {
      copyText(el.dataset.sql, ok => {
        if (!ok) return;
        const r = el.querySelector('.rcount');
        if (r) { const old = r.textContent; r.textContent = t('search.copied'); r.classList.add('copied'); setTimeout(() => { r.textContent = old; r.classList.remove('copied'); }, 1200); }
      });
      return;
    }
    if (el.dataset.copy) {
      copyText(el.dataset.copy, ok => {
        if (ok) { el.textContent = 'copied ✓'; el.classList.add('ok'); setTimeout(() => { el.textContent = 'copy'; el.classList.remove('ok'); }, 1200); }
      });
      return;
    }
    if (el.dataset.itab !== undefined) {
      const box = el.closest('.install');
      box.querySelectorAll('[data-itab]').forEach(b => b.setAttribute('aria-selected', b === el));
      box.querySelectorAll('[data-ipane]').forEach(p => { p.hidden = p.dataset.ipane !== el.dataset.itab; });
      return;
    }
  });

  document.addEventListener('change', ev => {
    const el = ev.target.closest('[data-skey]');
    if (!el) return;
    S[el.dataset.skey] = el.value;
    S.limit = DEFAULT_STATE.limit;
    renderDynamic();
  });

  document.addEventListener('keydown', ev => {
    if (ev.key === '/' && !ev.metaKey && !ev.ctrlKey && !ev.altKey) {
      const tag = (document.activeElement && document.activeElement.tagName) || '';
      if (tag === 'INPUT' || tag === 'TEXTAREA' || tag === 'SELECT') return;
      ev.preventDefault();
      const { path } = parseHash();
      if (path !== '/' && path !== '') { location.hash = '#/'; setTimeout(() => { const q = document.getElementById('q'); if (q) q.focus(); }, 60); }
      else { const q = document.getElementById('q'); if (q) q.focus(); }
    }
  });

  /* tooltip: [data-tip] hover + universe field */
  document.addEventListener('mousemove', ev => {
    const cv = ev.target.id === 'ufield' ? ev.target : null;
    if (cv) {
      const e = fieldHit(ev);
      if (e) {
        if (ucells) { ucells.hoverName = e.name; }
        drawFieldHover();
        showTip('<b>' + esc(e.name) + '</b> <span style="color:var(--c-' + e.cat + ')">' + e.cat + '</span><br><span class="d">' + esc(desc(e)) + '</span>', ev.clientX, ev.clientY);
        cv.style.cursor = 'pointer';
      } else { if (ucells) ucells.hoverName = null; drawFieldHover(); hideTip(); cv.style.cursor = 'crosshair'; }
      return;
    }
    const el = ev.target.closest('[data-tip]');
    if (el) {
      const txt = el.dataset.tip;
      if (txt) showTip(esc(txt), ev.clientX, ev.clientY);
    } else hideTip();
  });
  document.addEventListener('mouseout', ev => { if (!ev.relatedTarget) hideTip(); });
}

/* light hover redraw with rAF coalescing */
let rafPending = false;
function drawFieldHover() {
  if (rafPending) return;
  rafPending = true;
  requestAnimationFrame(() => { rafPending = false; drawField(); });
}

boot();
