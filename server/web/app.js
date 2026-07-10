'use strict';
/* ================================================================
   PGEXT.CLOUD — web app (served by `pgext serve`)
   data source: /api/v1 (bootstrap for the catalog, per-extension
   endpoints for detail hydration). No embedded data.
   ================================================================ */

/* ---------------- globals (filled by boot) ---------------- */
let BOOT = null;      // /api/v1/bootstrap payload
let EXT = [];         // slim extension records
let byName = new Map();
let CATS = {};        // code -> {code, zh, en, count}
let CAT_ORDER = [];   // canonical category order (by category id)
let PGS = [];         // active pg majors desc
let OSIDX = {};       // os -> canonical index
let UFIELD = [];      // dots for the universe field
let N_ALL = 0, N_AVAIL = 0, N_VENDOR = 0;
let META = {};

const F = { LEAD: 1, CONTRIB: 2, BIN: 4, LIB: 8, DDL: 16, LOAD: 32, TRUSTED: 64, RELOC: 128 };
const FULLC = new Map(), MXC = new Map(), FILEC = new Map(), DOCC = new Map();

/* bootstrap row columns — keep in sync with handleBootstrap in server/api.go:
   0 name 1 cat 2 avail 3 repo 4 license 5 lang 6 version 7 stars
   8 en 9 zh 10 type 11 vendor 12 kernel 13 pg[] 14 flags 15 docbits 16 commit */
function decodeBoot(b) {
  BOOT = b;
  EXT = b.rows.map((r, i) => ({
    i, name: r[0], cat: r[1], avail: !!r[2], repo: r[3] || 'n/a', license: r[4] || 'Unknown',
    lang: r[5] || '', ver: r[6] || '', stars: r[7], en: r[8] || '', zh: r[9] || '',
    type: r[10] || '', vendor: r[11] || '', kernel: r[12] || '',
    pg: r[13] || [], flags: r[14] | 0, docbits: r[15] | 0, commit: r[16] || ''
  }));
  byName = new Map(EXT.map(e => [e.name, e]));
  CATS = {}; CAT_ORDER = [];
  for (const c of b.cats) { CATS[c.name] = { code: c.name, zh: c.zh_desc, en: c.en_desc, count: c.count }; CAT_ORDER.push(c.name); }
  PGS = b.pg || [];
  OSIDX = {}; (b.os || []).forEach((os, i) => { OSIDX[os] = i; });
  N_ALL = b.counts.total; N_AVAIL = b.counts.packaged; N_VENDOR = b.counts.vendor;
  META = { generated: (b.generated || '').slice(0, 10) };
  UFIELD = [];
  for (const c of CAT_ORDER) {
    const members = EXT.filter(e => e.cat === c);
    members.sort((a, b2) => (b2.stars || 0) - (a.stars || 0) || a.name.localeCompare(b2.name));
    UFIELD.push(...members);
  }
}

async function j(url) {
  const res = await fetch(url, { headers: { Accept: 'application/json' } });
  if (!res.ok) {
    let msg = 'HTTP ' + res.status;
    try { const body = await res.json(); if (body.error) msg = body.error; } catch (e) {}
    const err = new Error(msg); err.status = res.status; throw err;
  }
  return res.json();
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
  'rows.stars': ['stars', '星标'], 'rows.desc': ['description', '描述'],
  'ext.crumb': ['extensions', '扩展'],
  'ext.install': ['Install', '安装'],
  'ext.avail': ['Availability', '可用性矩阵'],
  'ext.pkgs': ['Packages & Downloads', '安装包与下载'],
  'ext.deps': ['Dependencies & Relations', '依赖与关联'],
  'ext.docs': ['Usage', '用法'],
  'ext.related': ['More in {cat}', '更多{cat}扩展'],
  'ext.spec': ['Spec sheet', '规格表'],
  'ext.links': ['Links', '相关链接'],
  'ext.requires': ['requires', '依赖于'],
  'ext.requiredby': ['required by', '被依赖'],
  'ext.seealso': ['see also', '另请参阅'],
  'ext.siblings': ['same package', '同包扩展'],
  'ext.none': ['none', '无'],
  'ext.docsnone': ['No curated usage doc for this extension yet — check the upstream repository.',
                   '该扩展暂无人工撰写的用法文档——请参考上游仓库。'],
  'ext.doconlyen': ['Chinese doc not available yet — showing the English one.', '中文文档尚未就绪——先显示英文版。'],
  'ext.vendorNote': ['This extension ships only inside {vendor} managed services — it is not installable from public repositories.',
                     '该扩展仅随 {vendor} 云托管服务提供，无法从公开仓库安装。'],
  'spec.id': ['id', 'ID'],
  'spec.version': ['version', '版本'], 'spec.state': ['state', '状态'], 'spec.category': ['category', '分类'],
  'spec.license': ['license', '许可证'], 'spec.language': ['language', '语言'], 'spec.repo': ['repo', '仓库'],
  'spec.package': ['package', '包名'], 'spec.lead': ['lead ext', '主扩展'],
  'spec.schemas': ['schemas', '模式'], 'spec.pg': ['pg support', 'PG 支持'],
  'spec.rpm': ['rpm', 'RPM'], 'spec.rpmdeps': ['rpm deps', 'RPM 依赖'],
  'spec.deb': ['deb', 'DEB'], 'spec.debdeps': ['deb deps', 'DEB 依赖'],
  'spec.source': ['source', '源码包'], 'spec.vendor': ['vendor', '厂商'],
  'spec.tags': ['tags', '标签'],
  'spec.github': ['github', 'GitHub'], 'spec.release': ['released', '发布于'],
  'spec.updated': ['updated', '更新于'],
  'attr.ddl': ['CREATE EXTENSION', 'CREATE EXTENSION'],
  'attr.load': ['needs preload', '需要预加载'],
  'attr.lib': ['ships .so library', '携带动态库'],
  'attr.bin': ['ships binaries', '携带命令行工具'],
  'attr.trusted': ['trusted', '非超户可装'],
  'attr.reloc': ['relocatable', '模式可迁移'],
  'attr.contrib': ['postgres contrib', 'PG 自带'],
  'state.avail': ['available', '可用'], 'state.na': ['not packaged', '未打包'],
  'mx.legend.pgdg': ['available · PGDG repo', '可用 · PGDG 仓库'],
  'mx.legend.pigsty': ['available · Pigsty repo', '可用 · Pigsty 仓库'],
  'mx.legend.miss': ['not packaged', '未打包'],
  'mx.legend.warn': ['FORK: renamed / forked build', 'FORK：改名或分叉构建'],
  'mx.legend.bad': ['BREAK / THROW: known broken', 'BREAK / THROW：已知损坏'],
  'files.none': ['No binary artifacts recorded for this extension.', '没有该扩展的二进制包记录。'],
  'files.showall': ['show all {n} builds', '显示全部 {n} 个构建'],
  'files.collapse': ['latest builds only', '只看最新构建'],
  'files.count': ['{n} artifacts for PG {pg}', 'PG {pg} 共 {n} 个包文件'],
  'files.os': ['os', '系统'], 'files.pkg': ['package', '包'], 'files.ver': ['version', '版本'],
  'files.org': ['org', '来源'], 'files.size': ['size', '大小'], 'files.file': ['file', '文件'],
  'link.home': ['homepage', '主页'], 'link.repo': ['repository', '仓库'], 'link.license': ['license', '许可证'],
  'link.control': ['control file', 'control 文件'], 'link.author': ['author', '作者'], 'link.cargo': ['cargo', 'cargo'],
  'type.standard': ['standard — CREATE EXTENSION, no preload', 'standard——直接 CREATE EXTENSION，无需预加载'],
  'type.preload': ['preload — needs shared_preload_libraries', 'preload——需要 shared_preload_libraries'],
  'type.puresql': ['puresql — SQL objects only, no binary', 'puresql——纯 SQL 对象，无二进制'],
  'type.headless': ['headless — library only, no SQL objects', 'headless——只有库，无 SQL 对象'],
  'matrix.ext': ['CREATE EXTENSION', 'CREATE EXTENSION'],
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
  'about.title': ['About this catalog', '关于本目录'],
  'about.lede': ['PGEXT.CLOUD is the census of the PostgreSQL extension ecosystem — served live from the pgext catalog database by a single Go binary.',
                 'PGEXT.CLOUD 是 PostgreSQL 扩展生态的全量普查——由单个 Go 二进制从 pgext 目录数据库实时供给。'],
  'about.p1': ['Every entry is drawn from the <code>pgext</code> catalog: upstream repositories are crawled, RPM/DEB package indexes are parsed, GitHub activity is sampled, and 531 packaged extensions carry hand-curated bilingual documentation.',
               '每个条目都来自 <code>pgext</code> 目录数据库：抓取上游仓库、解析 RPM/DEB 包索引、采样 GitHub 活跃度；531 个已打包扩展配有人工撰写的中英双语文档。'],
  'about.p2': ['This site is <code>pgext serve</code>: web assets embedded in the binary, data queried live, snapshots cached in memory. The same API that renders these pages is public under <code>/api/v1</code>.',
               '本站就是 <code>pgext serve</code>：网页资产内嵌于二进制，数据实时查询、内存快照缓存。渲染这些页面所用的 API 就公开在 <code>/api/v1</code>。'],
  'about.roadmap': ['Roadmap', '路线图'],
  'about.r1': ['Galaxy — the dependency graph as a navigable star map', 'Galaxy——把依赖关系画成可漫游的星图'],
  'about.r2': ['Global availability matrix explorer — every extension × OS × PG at once', '全局可用性矩阵——所有扩展 × 操作系统 × PG 一屏总览'],
  'about.r3': ['Server-side rendering for real URLs & SEO', '服务端渲染真实 URL 与 SEO'],
  'about.sources': ['Data sources', '数据来源'],
  'about.s1': ['pgext.universe — the curated catalog of {n} extensions', 'pgext.universe——{n} 个扩展的策展目录'],
  'about.s2': ['PGDG & Pigsty repositories — RPM / DEB package metadata', 'PGDG 与 Pigsty 仓库——RPM / DEB 包元数据'],
  'about.s3': ['GitHub — stars, forks, freshness signals', 'GitHub——星标、分支与活跃度信号'],
  'about.api': ['Query API', '查询 API'],
  'about.colophon': ['Snapshot loaded {date} · refreshed automatically · POST /api/v1/reload to refresh now.',
                     '快照载入于 {date} · 自动刷新 · POST /api/v1/reload 可立即刷新。'],
  'notfound.hint': ['HINT:  Check the spelling, or search the catalog below.', 'HINT:  检查拼写，或回到目录搜索。'],
  'notfound.back': ['← back to the catalog', '← 返回目录'],
  'boot.err': ['could not reach the catalog API', '无法连接目录 API'],
  'boot.retry': ['retry', '重试'],
  'hydrate.err': ['failed to load: {msg}', '加载失败：{msg}'],
  'footer.built': ['Served live from the pgext catalog · snapshot {date} · API /api/v1', '由 pgext 目录数据库实时供给 · 快照 {date} · API /api/v1'],
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
const fmtSize = b => b == null || b === 0 ? '—' : b >= 1048576 ? (b / 1048576).toFixed(1) + ' MiB' : (b / 1024).toFixed(0) + ' KiB';
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

/* ---------------- micro markdown (usage docs) ---------------- */
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
      const nm = e.name.toLowerCase();
      if (nm === w) score += 120;
      else if (nm.startsWith(w)) score += 60;
      else if (nm.includes(w)) score += 30;
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

function skel(lines) {
  let bars = '';
  for (let i = 0; i < (lines || 3); i++) bars += '<div class="bar" style="width:' + (55 + ((i * 17) % 40)) + '%"></div>';
  return '<div class="skel">' + bars + '</div>';
}
const hydrateErr = err => '<div class="hydrate-err">' + esc(t('hydrate.err', { msg: err && err.message || 'unknown' })) + '</div>';

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
  return '<div class="wrap footer-in">'
    + '<span class="sig"><span class="p">pgext=#</span> \\q</span>'
    + '<span class="sig">' + t('footer.built', { date: META.generated || '' }) + '</span>'
    + '<nav><a href="/api/v1/meta" target="_blank" rel="noopener">API</a>'
    + '<a href="https://github.com/pgsty/pgext" target="_blank" rel="noopener">GitHub</a>'
    + '<a href="https://pigsty.io" target="_blank" rel="noopener">Pigsty — ' + t('footer.pigsty') + '</a></nav></div>';
}

/* ---------------- view: home ---------------- */
function homeHTML() {
  const stats = [
    [fmtInt(N_ALL), t('hero.s1')],
    [fmtInt(N_AVAIL), t('hero.s2')],
    [String(CAT_ORDER.length), t('hero.s3')],
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
let ucells = null;
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
    if (prevCat && e.cat !== prevCat) cell += 1;
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

function extHTML(name) {
  const e = byName.get(name);
  if (!e) return notFoundHTML(name);
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
    + (e.type ? '<span class="badge" data-tip="' + esc(t('type.' + e.type)) + '">' + esc(e.type) + '</span>' : '')
    + (e.vendor ? '<span class="badge">☁ ' + esc(e.vendor) + '</span>' : '')
    + '</div>'
    + '<div class="badge-row">' + attrBadges(e) + '</div>'
    + '<div class="ghstats" id="d-gh">'
    + (e.stars != null ? '<span>★ <b>' + fmtInt(e.stars) + '</b></span>' : '')
    + (e.commit ? '<span>' + t('commit.freshness', { d: '<b>' + esc(e.commit) + '</b>' }) + '</span>' : '')
    + '</div>' + vendorNote + '</header>'
    + '<div class="ext-grid" style="margin-top:26px"><div class="ext-main">'
    + '<div class="section" style="margin-top:0"><h2>' + t('ext.install') + '</h2><div id="d-install">' + skel(3) + '</div></div>'
    + '<div class="section"><h2>' + t('ext.avail') + '</h2><div id="d-matrix">' + skel(5) + '</div></div>'
    + '<div class="section"><h2>' + t('ext.pkgs') + '</h2><div id="d-files">' + skel(4) + '</div></div>'
    + '<div class="section"><h2>' + t('ext.deps') + '</h2><div id="d-deps">' + skel(2) + '</div></div>'
    + '<div class="section"><h2>' + t('ext.docs') + '</h2><div id="d-doc">' + (e.docbits ? skel(6) : '<div class="docs-missing">' + t('ext.docsnone') + '</div>') + '</div></div>'
    + (related.length ? '<div class="section"><h2>' + t('ext.related', { cat: e.cat }) + '</h2><ul class="related">' + related.map(tileHTML).join('') + '</ul></div>' : '')
    + '</div><aside class="ext-rail"><div id="d-spec">' + skel(8) + '</div></aside></div>'
    + '</article>';
}

function installHTML(e, full) {
  const tabs = [];
  const pgs = full.pg_ver || e.pg || [];
  const pgTop = pgs.length ? Math.max(...pgs) : (PGS[0] || 18);
  const requires = full.requires || [];
  const sqlLines = [];
  if (e.flags & F.LOAD) {
    sqlLines.push('<span class="cmt">-- 1. requires preload: add to postgresql.conf, then restart</span>');
    sqlLines.push('shared_preload_libraries = \'' + esc(e.name) + '\'');
    sqlLines.push('');
  }
  if (e.flags & F.DDL) {
    const cascade = requires.length ? ' CASCADE' : '';
    sqlLines.push('<span class="kw">CREATE EXTENSION</span> ' + esc(e.name) + cascade + ';');
  } else {
    sqlLines.push('<span class="cmt">-- ' + esc(e.type) + ': ships no SQL objects, nothing to CREATE</span>');
  }
  tabs.push(['SQL', sqlLines.join('\n'), (e.flags & F.DDL) ? 'CREATE EXTENSION ' + e.name + (requires.length ? ' CASCADE' : '') + ';' : '']);
  tabs.push(['pig', '<span class="cmt"># works on EL / Debian / Ubuntu alike</span>\npig install ' + esc(full.pkg || e.name), 'pig install ' + (full.pkg || e.name)]);
  if (full.rpm_pkg) {
    const p = full.rpm_pkg.replaceAll('$v', String(pgTop));
    const deps = (full.rpm_deps || []).length ? '\n<span class="cmt"># deps: ' + esc(full.rpm_deps.join(' ')) + '</span>' : '';
    tabs.push(['yum', '<span class="cmt"># ' + esc(full.rpm_repo || full.repo || '') + ' repo · v' + esc(full.rpm_ver || full.version || '') + ' · for PG ' + pgTop + '</span>' + deps + '\nyum install ' + esc(p), 'yum install ' + p]);
  }
  if (full.deb_pkg) {
    const p = full.deb_pkg.replaceAll('$v', String(pgTop));
    const deps = (full.deb_deps || []).length ? '\n<span class="cmt"># deps: ' + esc(full.deb_deps.join(' ')) + '</span>' : '';
    tabs.push(['apt', '<span class="cmt"># ' + esc(full.deb_repo || full.repo || '') + ' repo · v' + esc(full.deb_ver || full.version || '') + ' · for PG ' + pgTop + '</span>' + deps + '\napt install ' + esc(p), 'apt install ' + p]);
  }
  const heads = tabs.map(([n], i) => '<button role="tab" aria-selected="' + (i === 0) + '" data-itab="' + i + '">' + n + '</button>').join('');
  const bodies = tabs.map(([, html, copy], i) =>
    '<div class="install-body" data-ipane="' + i + '"' + (i ? ' hidden' : '') + '>'
    + (copy ? '<button class="copy-btn" data-copy="' + esc(copy) + '">copy</button>' : '')
    + '<pre>' + html + '</pre></div>').join('');
  return '<div class="install"><div class="install-tabs" role="tablist">' + heads + '</div>' + bodies + '</div>';
}

/* full availability matrix: os rows × pg columns from /matrix */
function fullMatrixHTML(m, e) {
  const byKey = {};
  for (const c of m.cells) byKey[c.os + '|' + c.pg] = c;
  const famOf = os => os.startsWith('el') ? 'EL' : os.startsWith('d') ? 'Debian' : 'Ubuntu';
  const ths = m.pg.map(pg => '<th>PG ' + pg + '</th>').join('');
  // synthetic first row: CREATE EXTENSION availability from pg_ver
  const extRow = '<tr><td class="oslab"><b>' + t('matrix.ext') + '</b> · v' + esc(e.ver || '—') + '</td>'
    + m.pg.map(pg => (e.pg || []).includes(pg)
      ? '<td><span class="cellv org-other" data-tip="' + esc(e.name) + '">✓</span></td>'
      : '<td class="miss">·</td>').join('') + '</tr>';
  let prevFam = null, rows = '';
  for (const os of m.os) {
    const fam = famOf(os);
    const famStart = fam !== prevFam;
    prevFam = fam;
    const [osname, arch] = os.split('.');
    const cells = m.pg.map(pg => {
      const c = byKey[os + '|' + pg];
      if (!c || c.state === 'MISS') return '<td class="miss">·</td>';
      if (c.state === 'AVAIL') {
        const org = (c.org || '').toLowerCase();
        const cls = org === 'pgdg' ? 'org-pgdg' : org === 'pigsty' ? 'org-pigsty' : 'org-other';
        return '<td><span class="cellv ' + cls + '" data-tip="' + esc(c.name || '') + ' · ' + c.count + ' builds">' + esc(c.org || '✓') + ' ' + esc(c.version || '') + '</span></td>';
      }
      const cls = (c.state === 'BREAK' || c.state === 'THROW') ? 'st-bad' : 'st-warn';
      return '<td><span class="cellv ' + cls + '" data-tip="' + esc(c.name || '') + ' · ' + esc(c.state) + '">' + esc(c.state) + '</span></td>';
    }).join('');
    rows += '<tr' + (famStart ? ' class="fam-start"' : '') + '><td class="oslab"><b>' + esc(osname) + '</b> · ' + esc(arch) + '</td>' + cells + '</tr>';
  }
  const legend = '<div class="mx-legend">'
    + '<span><span class="cellv org-pgdg">pgdg</span> ' + t('mx.legend.pgdg') + '</span>'
    + '<span><span class="cellv org-pigsty">pigsty</span> ' + t('mx.legend.pigsty') + '</span>'
    + '<span><span class="miss">·</span> ' + t('mx.legend.miss') + '</span>'
    + '<span><span class="cellv st-warn">FORK</span> ' + t('mx.legend.warn') + '</span>'
    + '<span><span class="cellv st-bad">BREAK</span> ' + t('mx.legend.bad') + '</span>'
    + '</div>';
  return '<div class="matrix-scroll"><table class="fmx"><thead><tr><th></th>' + ths + '</tr></thead><tbody>'
    + extRow + rows + '</tbody></table></div>' + legend;
}

/* package files: per-PG tabs, latest builds first, older collapsible */
function filesHTML(f) {
  if (!f || !f.files || !f.files.length) return '<p class="files-note">' + t('files.none') + '</p>';
  const pgs = [...new Set(f.files.map(x => x.pg))].sort((a, b) => b - a);
  const tabs = pgs.map((pg, i) => '<button data-ftab="' + pg + '" aria-selected="' + (i === 0) + '">PG ' + pg + '</button>').join('');
  const panes = pgs.map((pg, i) => {
    const rows = f.files.filter(x => x.pg === pg)
      .sort((a, b) => (OSIDX[a.os] ?? 99) - (OSIDX[b.os] ?? 99) || a.name.localeCompare(b.name) || b.ver.localeCompare(a.ver));
    const seen = new Set();
    let older = 0;
    const trs = rows.map(x => {
      const key = x.os + '|' + x.name;
      const latest = !seen.has(key);
      seen.add(key);
      if (!latest) older++;
      return '<tr' + (latest ? '' : ' class="older"') + '>'
        + '<td class="f-os">' + esc(x.os) + '</td>'
        + '<td>' + esc(x.name) + '</td>'
        + '<td>' + esc(x.ver) + '</td>'
        + '<td>' + esc(x.org || x.repo) + '</td>'
        + '<td class="f-size">' + fmtSize(x.size) + '</td>'
        + '<td class="f-file"><a href="' + esc(x.url) + '" target="_blank" rel="noopener" data-tip="' + esc(x.file) + '">' + esc(x.file) + '</a></td>'
        + '</tr>';
    }).join('');
    const foot = older
      ? '<div class="files-foot"><button class="chip" data-fall>' + t('files.showall', { n: rows.length }) + '</button>'
      + '<span class="files-note">' + t('files.count', { n: rows.length, pg }) + '</span></div>'
      : '<div class="files-foot"><span class="files-note">' + t('files.count', { n: rows.length, pg }) + '</span></div>';
    return '<div class="files-wrap" data-fpane="' + pg + '"' + (i ? ' hidden' : '') + '>'
      + '<div class="files"><div class="rows-scroll"><table><thead><tr>'
      + '<th>' + t('files.os') + '</th><th>' + t('files.pkg') + '</th><th>' + t('files.ver') + '</th>'
      + '<th>' + t('files.org') + '</th><th style="text-align:right">' + t('files.size') + '</th><th>' + t('files.file') + '</th>'
      + '</tr></thead><tbody>' + trs + '</tbody></table></div></div>' + foot + '</div>';
  }).join('');
  return '<div class="files-tabs" role="tablist">' + tabs + '</div>' + panes;
}

function depsHTML(full) {
  const sib = full.siblings || [];
  return '<div class="depgroup"><div class="lab">' + t('ext.requires') + '</div><div class="dep-chips">' + depChips(full.requires, full.name) + '</div></div>'
    + '<div class="depgroup"><div class="lab">' + t('ext.requiredby') + '</div><div class="dep-chips">' + depChips(full.require_by, full.name) + '</div></div>'
    + '<div class="depgroup"><div class="lab">' + t('ext.seealso') + '</div><div class="dep-chips">' + depChips(full.see_also, full.name) + '</div></div>'
    + (sib.length ? '<div class="depgroup"><div class="lab">' + t('ext.siblings') + ' · ' + esc(full.pkg) + '</div><div class="dep-chips">' + depChips(sib, full.name) + '</div></div>' : '');
}

function specHTML(e, full) {
  const r = (k, v, mono) => v ? '<div class="row"><dt>' + t(k) + '</dt><dd' + (mono ? ' class="mono"' : '') + '>' + v + '</dd></div>' : '';
  const gh = [];
  if (full.star_cnt != null) gh.push('★ ' + fmtInt(full.star_cnt));
  if (full.fork_cnt != null) gh.push('⑂ ' + fmtInt(full.fork_cnt));
  if (full.watch_cnt != null) gh.push('👁 ' + fmtInt(full.watch_cnt));
  const links = full.doc_links || {};
  const linkBadges = [
    ['link.home', links.home_url], ['link.repo', links.repo_url], ['link.license', links.license_url],
    ['link.control', links.control_url], ['link.author', links.author_url], ['link.cargo', links.cargo_url]
  ].filter(([, u]) => u).map(([k, u]) => '<a href="' + esc(u) + '" target="_blank" rel="noopener">' + t(k) + ' ↗</a>').join('');
  return '<div class="spec"><h3>' + t('ext.spec') + '</h3><dl>'
    + r('spec.id', String(full.id), true)
    + r('spec.version', esc(full.version), true)
    + r('spec.state', full.state === 'available' ? '<span style="color:var(--good);font-weight:650">' + t('state.avail') + '</span>' : '<span class="muted">' + t('state.na') + '</span>')
    + r('spec.category', '<a href="#/cat/' + full.category + '">' + esc(full.category) + ' · ' + esc(catName(full.category)) + '</a>')
    + r('spec.license', '<a href="#/?license=' + encodeURIComponent(full.license || '') + '">' + esc(full.license) + '</a>')
    + r('spec.language', '<a href="#/?lng=' + encodeURIComponent(full.lang || '') + '">' + esc(full.lang) + '</a>')
    + r('spec.repo', full.repo && full.repo !== 'n/a' ? '<a href="#/?repo=' + encodeURIComponent(full.repo) + '">' + esc(full.repo) + '</a>' : '')
    + r('spec.package', esc(full.pkg), true)
    + r('spec.lead', full.lead_ext && full.lead_ext !== full.name ? '<a href="' + extHref(full.lead_ext) + '">' + esc(full.lead_ext) + '</a>' : '', true)
    + r('spec.pg', esc(pgRange(full.pg_ver)), true)
    + r('spec.schemas', (full.schemas || []).length ? esc(full.schemas.join(', ')) : '', true)
    + r('spec.rpm', full.rpm_pkg ? esc(full.rpm_pkg) + ' <span class="muted">v' + esc(full.rpm_ver || full.version || '') + ' · ' + esc(full.rpm_repo || full.repo || '') + '</span>' : '', true)
    + r('spec.rpmdeps', (full.rpm_deps || []).length ? esc(full.rpm_deps.join(', ')) : '', true)
    + r('spec.deb', full.deb_pkg ? esc(full.deb_pkg) + ' <span class="muted">v' + esc(full.deb_ver || full.version || '') + ' · ' + esc(full.deb_repo || full.repo || '') + '</span>' : '', true)
    + r('spec.debdeps', (full.deb_deps || []).length ? esc(full.deb_deps.join(', ')) : '', true)
    + r('spec.source', esc(full.source || ''), true)
    + r('spec.vendor', full.ext_vendor ? esc(full.ext_vendor) + (full.ext_kernel ? ' <span class="muted">' + esc(full.ext_kernel) + '</span>' : '') : '', true)
    + r('spec.tags', (full.tags || []).length ? esc(full.tags.join(', ')) : '', true)
    + r('spec.github', gh.join(' · '), true)
    + r('spec.release', esc(full.last_release || ''), true)
    + r('spec.updated', esc(full.last_commit || ''), true)
    + '</dl>'
    + (linkBadges ? '<h3>' + t('ext.links') + '</h3><div class="doclinks">' + linkBadges + '</div>' : '')
    + '</div>';
}

/* hydration: fetch full record, matrix, files, doc — fill sections as they land */
let hydSeq = 0;
function fill(id, html) { const el = document.getElementById(id); if (el) el.innerHTML = html; }

async function hydrateExt(name) {
  const tok = ++hydSeq;
  const e = byName.get(name);
  if (!e) return;
  const enc = encodeURIComponent(name);

  (async () => {
    try {
      let full = FULLC.get(name);
      if (!full) { full = await j('/api/v1/ext/' + enc); FULLC.set(name, full); }
      if (tok !== hydSeq) return;
      fill('d-install', installHTML(e, full));
      fill('d-deps', depsHTML(full));
      fill('d-spec', specHTML(e, full));
      if (full.url) {
        const gh = document.getElementById('d-gh');
        if (gh && !gh.querySelector('a')) {
          gh.insertAdjacentHTML('beforeend', '<a href="' + esc(full.url) + '" target="_blank" rel="noopener">' + esc(full.url.replace(/^https?:\/\//, '')) + ' ↗</a>');
        }
      }
    } catch (err) {
      if (tok !== hydSeq) return;
      fill('d-install', hydrateErr(err)); fill('d-deps', hydrateErr(err)); fill('d-spec', hydrateErr(err));
    }
  })();

  (async () => {
    try {
      let m = MXC.get(name);
      if (!m) { m = await j('/api/v1/ext/' + enc + '/matrix'); MXC.set(name, m); }
      if (tok !== hydSeq) return;
      fill('d-matrix', m.cells && m.cells.length ? fullMatrixHTML(m, e) : '<p class="files-note">' + t('files.none') + '</p>');
    } catch (err) { if (tok === hydSeq) fill('d-matrix', hydrateErr(err)); }
  })();

  (async () => {
    try {
      let fl = FILEC.get(name);
      if (!fl) { fl = await j('/api/v1/ext/' + enc + '/files'); FILEC.set(name, fl); }
      if (tok !== hydSeq) return;
      fill('d-files', filesHTML(fl));
    } catch (err) { if (tok === hydSeq) fill('d-files', hydrateErr(err)); }
  })();

  if (e.docbits) {
    (async () => {
      const wantZh = LANG === 'zh' && (e.docbits & 2);
      const lang = wantZh ? 'zh' : 'en';
      const key = name + '|' + lang;
      try {
        let d = DOCC.get(key);
        if (!d) { d = await j('/api/v1/ext/' + enc + '/doc?lang=' + lang); DOCC.set(key, d); }
        if (tok !== hydSeq) return;
        const note = (LANG === 'zh' && !wantZh) ? '<p class="files-note" style="margin-bottom:10px">' + t('ext.doconlyen') + '</p>' : '';
        fill('d-doc', note + '<div class="prose">' + renderMD(d.content) + '</div>');
      } catch (err) { if (tok === hydSeq) fill('d-doc', hydrateErr(err)); }
    })();
  }
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
    .sort((a, b) => ((b.docbits ? 1 : 0) - (a.docbits ? 1 : 0)) || (b.stars || 0) - (a.stars || 0))
    .slice(0, 6);
  const strip = CAT_ORDER.map(x =>
    '<a href="#/cat/' + x + '" style="--seg:var(--c-' + x + ');flex:' + CATS[x].count + ' 1 0" aria-current="' + (x === code) + '" data-tip="' + x + ' · ' + CATS[x].count + '"></a>').join('');
  const fcards = featured.map(e =>
    '<li><a class="fcard" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '>'
    + '<span class="name">' + esc(e.name) + '<span class="ver">' + esc(e.ver || '') + '</span>'
    + (e.stars ? '<span class="star">★ ' + fmtNum(e.stars) + '</span>' : '') + '</span>'
    + '<span class="desc">' + esc(desc(e)) + '</span>'
    + '<span class="meta"><span>' + esc(e.license) + '</span><span>' + esc(e.lang) + '</span><span>PG ' + esc(pgRange(e.pg)) + '</span>'
    + (e.docbits ? '<span>📖 docs</span>' : '') + '</span>'
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
    + '<li><span class="tag">catalog</span>' + t('about.s1', { n: fmtInt(N_ALL) }) + '</li>'
    + '<li><span class="tag">packages</span>' + t('about.s2') + '</li>'
    + '<li><span class="tag">github</span>' + t('about.s3') + '</li></ul></div>'
    + '<div><h3>' + t('about.roadmap') + '</h3><ul class="roadmap">'
    + '<li><span class="tag">phase 3</span>' + t('about.r1') + '</li>'
    + '<li><span class="tag">phase 3</span>' + t('about.r2') + '</li>'
    + '<li><span class="tag">phase 3</span>' + t('about.r3') + '</li></ul></div>'
    + '<div><h3>' + t('about.api') + '</h3><ul class="roadmap">'
    + '<li><span class="tag">GET</span><code>/api/v1/ext?q=vector&cat=RAG</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/matrix</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/files?pg=18</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/doc?lang=zh</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/dim/license</code></li>'
    + '</ul></div>'
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
    const name = decodeURIComponent(path.slice(5));
    app.innerHTML = extHTML(name);
    hydrateExt(name);
    active = '';
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
  currentPath = null;
  route();
  window.scrollTo(0, y);
}

/* ---------------- tooltip ---------------- */
const tip = () => document.getElementById('tip');
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
function attachEvents() {
  window.addEventListener('hashchange', () => { if (!syncingHash) route(); });
  window.addEventListener('resize', debounce(() => drawField(), 150));

  document.addEventListener('click', ev => {
    const el = ev.target.closest('[data-fkey],[data-skey],[data-view],[data-more],[data-sql],[data-copy],[data-itab],[data-ftab],[data-fall],[data-cat-go],#lang-toggle,#theme-toggle,#ufield');
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
    if (el.dataset.ftab !== undefined) {
      const box = document.getElementById('d-files');
      if (!box) return;
      box.querySelectorAll('[data-ftab]').forEach(b => b.setAttribute('aria-selected', b === el));
      box.querySelectorAll('[data-fpane]').forEach(p => { p.hidden = p.dataset.fpane !== el.dataset.ftab; });
      return;
    }
    if (el.dataset.fall !== undefined) {
      const wrap = el.closest('.files-wrap');
      if (!wrap) return;
      wrap.classList.toggle('show-all');
      const total = wrap.querySelectorAll('tbody tr').length;
      el.textContent = wrap.classList.contains('show-all') ? t('files.collapse') : t('files.showall', { n: total });
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

let rafPending = false;
function drawFieldHover() {
  if (rafPending) return;
  rafPending = true;
  requestAnimationFrame(() => { rafPending = false; drawField(); });
}

/* ---------------- boot ---------------- */
function bootError(err) {
  const app = document.getElementById('app');
  app.innerHTML = '<div class="boot"><div class="boot-err">'
    + '<span class="e">psql: error:</span> ' + esc(t('boot.err')) + '\n'
    + '<span class="muted">DETAIL:  ' + esc(err && err.message || 'network error') + '</span>\n'
    + '<button class="chip" id="boot-retry">↻ ' + esc(t('boot.retry')) + '</button>'
    + '</div></div>';
  document.getElementById('boot-retry').addEventListener('click', boot);
}

async function boot() {
  applyTheme();
  document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
  try {
    const b = await j('/api/v1/bootstrap');
    decodeBoot(b);
  } catch (err) {
    bootError(err);
    return;
  }
  route();
}

attachEvents();
boot();
