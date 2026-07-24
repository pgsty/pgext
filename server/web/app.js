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
let byPkg = new Map();
let CATS = {};        // code -> {code, zh, en, count}
let CAT_ORDER = [];   // canonical category order (by category id)
let PGS = [];         // active pg majors desc
let OSS = [];         // active os targets in canonical order
let OSIDX = {};       // os -> canonical index
let UFIELD = [];      // dots for the universe field
let N_ALL = 0, N_AVAIL = 0, N_PROJECTS = 0, N_SOURCE = 0, N_VENDOR = 0, N_KERNEL = 0, N_CONTRIB = 0, N_DOCS = 0;
let N_PKGS = 0, N_SLOTS = 0; // pgext.pkg package families / build slots
let META = {};
let INSTALL_PREF = { pg: '', os: '' };

const F = { LEAD: 1, CONTRIB: 2, BIN: 4, LIB: 8, DDL: 16, LOAD: 32, TRUSTED: 64, RELOC: 128 };
const B = { RPM: 1, DEB: 2, PGRX: 4, SOURCE: 8 };
const R = { REQUIRES: 1, REQUIRED_BY: 2, SEE_ALSO: 4 };
const FULLC = new Map(), MXC = new Map(), FILEC = new Map(), DOCC = new Map();
let GMATRIX = null, GMATRIX_VIEW = null, matrixHydSeq = 0;

/* bootstrap row columns — keep in sync with handleBootstrap in server/api.go:
   0 name 1 cat 2 avail 3 repo 4 license 5 lang 6 version 7 stars
   8 en 9 zh 10 kind 11 vendor 12 kernel 13 pg[] 14 flags 15 docbits 16 commit
   17 pkg 18 lead 19 lifecycle 20 tags[] 21 active 22 checked 23 buildbits
   24 target_idx[] 25 family_size 26 comment 27 relationbits 28 pgrx_ver 29 repo_url 30 url 31 license_url
   32 rpm_pkg 33 deb_pkg */
function decodeBoot(b) {
  BOOT = b;
  EXT = b.rows.map((r, i) => ({
    i, name: r[0], cat: r[1], avail: !!r[2], repo: r[3] || 'n/a', license: r[4] || 'Unknown',
    lang: r[5] || '', ver: r[6] || '', stars: r[7], en: r[8] || '', zh: r[9] || '',
    kind: r[10] || '', type: r[10] || '', vendor: r[11] || '', kernel: r[12] || '',
    pg: r[13] || [], flags: r[14] | 0, docbits: r[15] | 0, commit: r[16] || '',
    pkg: r[17] || r[0], lead: r[18] || r[0], lifecycle: r[19] || '', tags: r[20] || [],
    active: r[21] || r[16] || '', checked: r[22] || '', buildbits: r[23] | 0,
    targets: r[24] || [], familySize: r[25] || 1, comment: r[26] || '',
    relationbits: r[27] | 0, pgrx: r[28] || '', repoUrl: r[29] || '', url: r[30] || '', licenseUrl: r[31] || '',
    rpmPkg: r[32] || '', debPkg: r[33] || ''
  }));
  byName = new Map(EXT.map(e => [e.name, e]));
  byPkg = new Map();
  for (const e of EXT) { if (!byPkg.has(e.pkg)) byPkg.set(e.pkg, []); byPkg.get(e.pkg).push(e); }
  CATS = {}; CAT_ORDER = [];
  for (const c of b.cats) { CATS[c.name] = { code: c.name, zh: c.zh_desc, en: c.en_desc, count: c.count }; CAT_ORDER.push(c.name); }
  PGS = b.pg || [];
  OSS = b.os || [];
  OSIDX = {}; OSS.forEach((os, i) => { OSIDX[os] = i; });
  // Counts derivable from the rows are derived, so figures can never disagree
  // with the decoded catalog itself; only the pkg-table aggregates rely on the
  // server (with a slot fallback of families × pg × os).
  N_ALL = EXT.length; N_PROJECTS = byPkg.size;
  N_AVAIL = 0; N_SOURCE = 0; N_VENDOR = 0; N_KERNEL = 0; N_CONTRIB = 0;
  for (const e of EXT) {
    if (e.avail) N_AVAIL++;
    else if (e.buildbits & B.SOURCE) N_SOURCE++;
    if (e.vendor) N_VENDOR++;
    if (e.kernel) N_KERNEL++;
    if (e.flags & F.CONTRIB) N_CONTRIB++;
  }
  N_DOCS = b.counts.docs || 0;
  N_PKGS = b.counts.packages || 0;
  N_SLOTS = b.counts.build_slots || (N_PKGS * PGS.length * OSS.length);
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
  'nav.matrix': ['Matrix', '构建矩阵'],
  'nav.browse': ['Index', '扩展索引'],
  'nav.about': ['About', '关于'],
  'nav.lang': ['中文', 'EN'],
  'hero.eyebrow': ['pgext.cloud · everything postgresql can become', 'pgext.cloud · PostgreSQL 的一切可能'],
  'hero.title': ['<em>PostgreSQL</em> Extension Catalog', '<em>PostgreSQL</em> 扩展目录'],
  'hero.sub': ['Explore the superpowers of <b>{all}</b> PostgreSQL extensions, and the availability of <b>{avail}</b> extension artifacts across <b>{os}</b> Linux platforms.',
               '探索 <b>{all}</b> 个 PostgreSQL 扩展的超能力，以及 <b>{avail}</b> 个 PG 扩展在 <b>{os}</b> 个 Linux 平台上的可用性。'],
  'hero.s1': ['extensions catalogued', '收录在册扩展'],
  'hero.s2': ['packaged extensions', '已打包扩展'],
  'hero.s3': ['extension packages', '扩展软件包'],
  'hero.s4': ['Linux platforms', 'Linux 大版本'],
  'hero.s5': ['PG majors', 'PG 大版本'],
  'hero.s6': ['build slots', '构建槽位'],
  'search.ph': ['Search {n} extensions — name, package, tags, or try tag:vector build:pgrx', '搜索 {n} 个扩展——名称、项目、标签，或试试 tag:vector build:pgrx'],
  'search.copy': ['click to copy query', '点击复制查询'],
  'search.copied': ['copied ✓', '已复制 ✓'],
  'chip.all': ['All', '全部'],
  'chip.packaged': ['Packaged', '已打包'],
  'chip.unpacked': ['Unpacked', '未交付'],
  'chip.source': ['Source-only', '仅源码'],
  'chip.kernel': ['Fork-Only', '分支限定'],
  'chip.vendor': ['Vendor-Specific', '厂商限定'],
  'chip.contrib': ['Contrib', 'PG 自带'],
  'sel.cat': ['category', '分类'],
  'sel.license': ['license', '许可证'],
  'sel.lang': ['language', '语言'],
  'sel.pg': ['pg version', 'PG 版本'],
  'sel.os': ['binary target', '二进制目标'],
  'sel.kind': ['extension kind', '扩展形态'],
  'sel.lifecycle': ['lifecycle', '生命周期'],
  'filter.scope': ['scope', '收录范围'],
  'filter.category': ['category', '功能分类'],
  'filter.pg': ['PG major', 'PG 版本'],
  'filter.lang': ['language', '实现语言'],
  'filter.license': ['license', '许可证'],
  'filter.license.more': ['all...', 'all...'],
  'filter.any': ['Any', '不限'],
  'filter.dimensions': ['More Dims', '更多维度'],
  'filter.active': ['Active filters', '当前筛选'],
  'filter.results.ext': ['matching extensions', '个匹配扩展'],
  'filter.results.pkg': ['matching projects', '个匹配项目'],
  'sort.label': ['sort', '排序'],
  'sort.stars': ['stars', '星标'],
  'sort.name': ['name', '名称'],
  'sort.updated': ['updated', '更新'],
  'view.ext': ['EXT', 'EXT'],
  'view.pkg': ['PKG', 'PKG'],
  'view.card': ['Card', '卡片'],
  'view.list': ['List', '列表'],
  'wall.empty': ['No extension matches. Loosen a filter, or try <code>vector</code>, <code>gis</code>, <code>parquet</code>.',
                 '没有匹配的扩展。放宽筛选条件，或试试 <code>vector</code>、<code>gis</code>、<code>parquet</code>。'],
  'rows.name': ['name', '名称'], 'rows.cat': ['category', '分类'], 'rows.ver': ['version', '版本'],
  'rows.license': ['license', '许可证'], 'rows.lang': ['lang', '语言'], 'rows.pg': ['pg ver', 'PG 版本'],
  'rows.stars': ['stars', '星标'], 'rows.desc': ['description', '描述'],
  'rows.package': ['package', '扩展包'], 'rows.lead': ['lead extension', '主扩展'],
  'rows.extensions': ['extensions', '扩展数'], 'rows.packaged': ['packaged', '已打包'],
  'ext.crumb': ['extensions', '扩展'],
  'ext.overview': ['Overview', '概览'],
  'ext.relations': ['Relationship', '关联'],
  'ext.build': ['Building', '构建'],
  'ext.install': ['Installation', '安装'],
  'ext.veravail': ['Packages', '制品'],
  'ext.pkgs': ['Packages', '软件包列表'],
  'ext.downloads': ['Downloads', '制成品下载'],
  'ext.installmore': ['For preload, CREATE EXTENSION and verification details, see the extension install guide:',
                      '预加载、CREATE EXTENSION 与验证等更多细节，请参考扩展页面的详细安装说明：'],
  'ext.deps': ['Dependency', '依赖关系'],
  'dep.noup': ['This extension depends on no other extension.', '此扩展不依赖其他扩展'],
  'dep.nodown': ['No extension depends on this one.', '没有扩展依赖此扩展'],
  'ext.docs': ['Usage', '用法'],
  'ext.onpage': ['On this page', '本页目录'],
  'ext.catalog': ['Catalog identity', '目录标识'],
  'ext.runtime': ['Extension Attribute', '扩展属性'],
  'ext.packaging': ['Packaging metadata', '打包元数据'],
  'ext.upstream': ['Upstream & activity', '上游与活跃度'],
  'ext.projectlinks': ['Project Links', '项目链接'],
  'ext.buildnone': ['The catalog has source metadata but no reproducible RPM/DEB build recipe. Follow the upstream build documentation.',
                    '目录已收录源码信息，但没有可复现的 RPM/DEB 构建配方；请遵循上游构建文档。'],
  'ext.contribbuild': ['This extension is delivered with PostgreSQL contrib; build it together with the matching PostgreSQL server packages.',
                       '这是 PostgreSQL contrib 扩展，应随对应版本的 PostgreSQL 服务端软件包一起构建和交付。'],
  'ext.buildrecipe': ['A Pigsty package recipe is recorded for {types}. Build it with the same package-family name used by the catalog.',
                      '目录记录了 {types} 的 Pigsty 构建配方，请使用目录中的包族名称进行构建。'],
  'ext.sourcearchive': ['Source Tarball', '源码包'],
  'ext.installlede': ['Enable the repository, install the package for your PostgreSQL major with <code>pig</code>, <code>apt</code> or <code>dnf</code>, then follow the preload, CREATE EXTENSION and verification steps from <code>pgext.universe</code>.',
                      '启用软件仓库，用 <code>pig</code>、<code>apt</code> 或 <code>dnf</code> 为你的 PostgreSQL 大版本安装软件包，再按 <code>pgext.universe</code> 生成的步骤完成预加载、CREATE EXTENSION 与验证。'],
  'ext.step.repo': ['1 · Repository', '1 · 软件仓库'],
  'ext.step.package': ['2 · Install package', '2 · 安装软件包'],
  'ext.step.load': ['3 · Load library', '3 · 加载动态库'],
  'ext.step.enable': ['4 · Enable extension', '4 · 启用扩展'],
  'ext.noload': ['No shared preload is required.', '无需配置 shared_preload_libraries。'],
  'ext.noddl': ['This entry does not create SQL objects with CREATE EXTENSION.', '该条目无需通过 CREATE EXTENSION 创建 SQL 对象。'],
  'ext.showall': ['Show all {n}', '显示全部 {n} 项'],
  'ext.spec': ['Spec sheet', '规格表'],
  'ext.links': ['Links', '相关链接'],
  'ext.requires': ['requires', '依赖于'],
  'ext.requiredby': ['required by', '被依赖'],
  'ext.seealso': ['see also', '另请参阅'],
  'ext.siblings': ['same package', '同包扩展'],
  'ext.docsnone': ['No curated usage doc for this extension yet — check the upstream repository.',
                   '该扩展暂无人工撰写的用法文档——请参考上游仓库。'],
  'ext.doconlyen': ['Chinese doc not available yet — showing the English one.', '中文文档尚未就绪——先显示英文版。'],
  'ext.vendorNote': ['This extension ships only inside {vendor} managed services — it is not installable from public repositories.',
                     '该扩展仅随 {vendor} 云托管服务提供，无法从公开仓库安装。'],
  'ext.providerNote': ['This catalog entry is associated with {provider}. No public binary package is currently recorded; consult the upstream or provider documentation.',
                       '该目录项关联 {provider}，目前没有公开二进制包记录；请参考上游或服务商文档。'],
  'ext.lifecycleNote': ['Upstream lifecycle is {state}. Evaluate maintenance and compatibility before production use.',
                        '上游生命周期状态为 {state}，用于生产前请评估维护状态与兼容性。'],
  'ext.family': ['Extensions', '扩展列表'],
  'ext.visits': ['visits', '次访问'],
  'spec.category': ['Category', '分类'],
  'spec.license': ['License', '许可'], 'spec.language': ['Language', '语言'],
  'spec.vendor': ['Vendor', '厂商'], 'spec.kernel': ['Kernel', '内核'],
  'spec.kind': ['Kind', '形态'], 'spec.lifecycle': ['Lifecycle', '周期'],
  'spec.kernel': ['kernel', '内核'], 'spec.libs': ['libraries', '动态库'],
  'spec.pgrx': ['pgrx', 'pgrx'], 'spec.checked': ['checked', '核验于'], 'spec.active': ['last active', '最近活跃'],
  'spec.created': ['repo created', '仓库创建'], 'spec.comment': ['catalog note', '目录备注'],
  'spec.github': ['github', 'GitHub'], 'spec.release': ['released', '发布于'],
  'spec.updated': ['updated', '更新于'],
  'spec.mtime': ['catalog mtime', '目录更新时间'],
  'spec.extra': ['extra metadata', '扩展元数据'],
  'attr.ddl': ['CREATE EXTENSION', 'CREATE EXTENSION'],
  'attr.load': ['needs preload', '需要预加载'],
  'attr.lib': ['ships .so library', '携带动态库'],
  'attr.bin': ['ships binaries', '携带命令行工具'],
  'attr.trusted': ['trusted', '非超户可装'],
  'attr.reloc': ['relocatable', '模式可迁移'],
  'attr.contrib': ['postgres contrib', 'PG 自带'],
  'state.avail': ['available', '可用'], 'state.na': ['not packaged', '未打包'],
  'mx.legend.pgdg': ['Available in PGDG Repo', '在 PGDG 仓库中可用'],
  'mx.legend.pigsty': ['Available in PIGSTY Repo', '在 PIGSTY 仓库中可用'],
  'mx.legend.miss': ['Not provided', '扩展缺失'],
  'mx.legend.na': ['Invalid combination', '无效构建组合'],
  'files.none': ['No binary artifacts recorded for this extension.', '没有该扩展的二进制包记录。'],
  'files.showall': ['show all {n} builds', '显示全部 {n} 个构建'],
  'files.collapse': ['latest builds only', '只看最新构建'],
  'files.count': ['{n} artifacts for PG {pg}', 'PG {pg} 共 {n} 个包文件'],
  'files.os': ['platform', '系统'], 'files.pkg': ['package', '包名'], 'files.ver': ['version', '版本'],
  'files.org': ['repo', '仓库'], 'files.size': ['size', '大小'], 'files.file': ['file', '文件'],
  'files.sha': ['sha256', 'SHA256'],
  'link.home': ['homepage', '主页'], 'link.repo': ['repository', '仓库'], 'link.license': ['license', '许可证'],
  'link.docs': ['document', '官方文档'], 'link.pgxn': ['PGXN', 'PGXN'],
  'link.control': ['control', 'Control'], 'link.author': ['author', '作者'], 'link.cargo': ['cargo', 'cargo'],
  'type.standard': ['standard — CREATE EXTENSION, no preload', 'standard——直接 CREATE EXTENSION，无需预加载'],
  'type.preload': ['preload — needs shared_preload_libraries', 'preload——需要 shared_preload_libraries'],
  'type.puresql': ['puresql — SQL objects only, no binary', 'puresql——纯 SQL 对象，无二进制'],
  'type.headless': ['headless — library only, no SQL objects', 'headless——只有库，无 SQL 对象'],
  'gmx.eyebrow': ['package intelligence · pgext.matrix', '软件包情报 · pgext.matrix'],
  'gmx.title': ['Build Matrix', '构建矩阵'],
  'gmx.lede': ['Availability of {pkgs} extension packages across {combos} Linux × PostgreSQL build targets.',
               '列出 {pkgs} 个扩展包在 {combos} 个 Linux × PostgreSQL 组合矩阵下的可用性。'],
  'gmx.search': ['Filter package or extension…', '筛选扩展包或扩展名…'],
  'gmx.showing': ['{rows} packages · {cells} cells', '{rows} 个扩展包 · {cells} 个格子'],
  'gmx.source': ['Computed at catalog load time from pgext.pkg and cached in {source} — the page renders the whole sheet in one pass.',
                 '目录装载时由 pgext.pkg 预计算并缓存于 {source}——页面一次性渲染整张矩阵。'],
  'gmx.pgdg': ['PGDG', 'PGDG'],
  'gmx.pigsty': ['Pigsty', 'Pigsty'],
  'gmx.missing': ['Missing', '缺失'],
  'gmx.na': ['N/A', 'N/A'],
  'gmx.empty': ['No package row matches this filter.', '没有扩展包行符合当前筛选条件。'],
  'gmx.hint': ['Click a status to isolate its rows · click again to reset · click any cell to inspect it',
               '点击状态可单独查看 · 再次点击恢复全部 · 点击任意格子查看详情'],
  'gmx.pkg': ['PACKAGE / PLATFORM', '扩展包 / 平台'],
  'gmx.api': ['JSON API', 'JSON API'],
  'cat.featured': ['Featured', '精选'],
  'cat.all': ['All {n} extensions', '全部 {n} 个扩展'],
  'cat.open': ['Open in search ↗', '在搜索中打开 ↗'],
  'browse.title': ['Extension Index', '扩展索引'],
  'browse.lede': ['One catalog, many ways in. Slice the extension universe by any dimension — every value below is a live filter.',
                  '一份目录，多个入口。任选维度切分扩展宇宙——下面每个值都是一个即时筛选器。'],
  'browse.cats': ['Categories', '功能分类'],
  'browse.catalog': ['Identity & Attributes', '身份与属性'],
  'browse.delivery': ['Packaging & Delivery', '构建与交付'],
  'browse.attrs': ['Attributes & Ecosystem', '属性与生态'],
  'dim.category': ['Categories', '分类'],
  'dim.tag': ['Tags', '主题标签'],
  'dim.package': ['Package Families', '项目包族'],
  'dim.license': ['Licenses', '许可证'],
  'dim.lang': ['Languages', '实现语言'],
  'dim.repo': ['Repositories', '仓库来源'],
  'dim.distribution': ['Catalog Scope', '目录口径'],
  'dim.os': ['Linux Platforms', '操作系统'],
  'dim.build': ['Build Toolchains', '构建链'],
  'dim.pgrx': ['pgrx Versions', 'pgrx 版本'],
  'dim.capability': ['Runtime Capabilities', '运行时能力'],
  'dim.docs': ['Document Coverage', '文档覆盖'],
  'dim.relation': ['Dependency Signals', '依赖关系'],
  'dim.activity': ['Last Active Year', '最近活跃年份'],
  'dim.vendor': ['Cloud Vendors', '云厂商'],
  'dim.kernel': ['PostgreSQL Kernels', 'PostgreSQL 内核'],
  'dim.lifecycle': ['Lifecycles', '生命周期'],
  'dim.pg': ['PostgreSQL Versions', 'PostgreSQL 版本'],
  'dim.type': ['Extension Kinds', '扩展形态'],
  'dim.category.d': ['16 functional domains, from GIS to AI', '16 个功能领域，从 GIS 到 AI'],
  'dim.tag.d': ['curated topics, features, and ecosystem signals', '策展的主题、能力与生态标签'],
  'dim.package.d': ['group extension definitions by upstream project', '按上游项目聚合同包扩展定义'],
  'dim.license.d': ['from PostgreSQL-liberal to commercial', '从 PostgreSQL 式宽松到商业许可'],
  'dim.lang.d': ['what extensions are written in', '扩展用什么写成'],
  'dim.repo.d': ['who packages and distributes it', '谁打包、谁分发'],
  'dim.distribution.d': ['packaged, source-only, contrib, kernel, and vendor catalogs', '已打包、仅源码、PG 自带、内核与厂商目录'],
  'dim.os.d': ['exact operating system and architecture package targets', '精确到操作系统与架构的软件包目标'],
  'dim.build.d': ['RPM, DEB, pgrx, source, and binary delivery paths', 'RPM、DEB、pgrx、源码与二进制交付路径'],
  'dim.pgrx.d': ['Rust extension ABI/toolchain versions recorded upstream', '上游记录的 Rust 扩展 ABI 与工具链版本'],
  'dim.capability.d': ['DDL, preload, libraries, binaries, trust, and relocation', 'DDL、预加载、动态库、命令行、可信与模式迁移'],
  'dim.docs.d': ['bilingual, single-language, and missing usage manuals', '双语、单语与尚缺的用法手册'],
  'dim.relation.d': ['requires, required-by, and see-also graph coverage', '依赖、被依赖与相关扩展图谱覆盖'],
  'dim.activity.d': ['upstream freshness grouped by the last active year', '按最近活跃年份观察上游新鲜度'],
  'dim.vendor.d': ['vendor and managed-service associations', '厂商与托管服务关联'],
  'dim.kernel.d': ['kernel-specific compatibility and availability', '内核特定的兼容与可用关系'],
  'dim.lifecycle.d': ['active, preview, deprecated, archived, abandoned', '活跃、预览、弃用、归档与废弃状态'],
  'dim.pg.d': ['coverage across supported majors', '各大版本的覆盖面'],
  'dim.type.d': ['standard, preload, pure-SQL, headless', 'standard / preload / puresql / headless'],
  'dim.value': ['value', '取值'], 'dim.count': ['extensions', '扩展数'], 'dim.sample': ['examples', '示例'],
  'dim.search': ['Filter {n} values…', '筛选 {n} 个取值…'],
  'dim.showing': ['{shown} of {all} values', '显示 {shown} / {all} 个取值'],
  'notfound.hint': ['HINT:  Check the spelling, or search the catalog below.', 'HINT:  检查拼写，或回到目录搜索。'],
  'notfound.back': ['← back to the catalog', '← 返回目录'],
  'boot.err': ['could not reach the catalog API', '无法连接目录 API'],
  'boot.retry': ['retry', '重试'],
  'hydrate.err': ['failed to load: {msg}', '加载失败：{msg}'],
  'footer.built': ['Served live from the pgext catalog · snapshot {date} · API /api/v1', '由 pgext 目录数据库实时供给 · 快照 {date} · API /api/v1'],
  'commit.freshness': ['last commit {d}', '最近提交 {d}']
};

/* UI language: ?lang= override > stored toggle choice > browser language
   (any zh-* Accept-Language picks Chinese) > English. */
let LANG = 'en';
try {
  const prefs = navigator.languages && navigator.languages.length ? navigator.languages : [navigator.language || ''];
  if (prefs.some(l => /^zh\b/i.test(String(l)))) LANG = 'zh';
} catch (e) {}
try { LANG = localStorage.getItem('pgext.lang') || LANG; } catch (e) {}
try { const q = new URLSearchParams(location.search).get('lang'); if (q === 'zh' || q === 'en') LANG = q; } catch (e) {}
const t = (k, vars) => {
  const pair = I18N[k]; let s = pair ? pair[LANG === 'zh' ? 1 : 0] : k;
  if (vars) for (const [kk, vv] of Object.entries(vars)) s = s.replaceAll('{' + kk + '}', vv);
  return s;
};
const bi = (en, zh) => LANG === 'zh' ? zh : en;
const catName = c => { const p = CAT_NAMES[c]; return p ? p[LANG === 'zh' ? 1 : 0] : c; };
const desc = e => (LANG === 'zh' && e.zh) ? e.zh : (e.en || e.zh || '');
const catDesc = c => { const m = CATS[c]; return m ? (LANG === 'zh' ? m.zh : m.en) : ''; };

/* ---------------- utils ---------------- */
const esc = s => String(s == null ? '' : s).replace(/[&<>"']/g, m => ({ '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[m]));
const fmtNum = n => n == null ? '—' : n >= 1000 ? (n / 1000).toFixed(n >= 10000 ? 0 : 1) + 'k' : String(n);
const fmtInt = n => (n == null ? '—' : String(n).replace(/\B(?=(\d{3})+(?!\d))/g, ','));
const fmtSize = b => b == null || b === 0 ? '—' : b >= 1048576 ? (b / 1048576).toFixed(1) + ' MiB' : (b / 1024).toFixed(0) + ' KiB';
const pgRange = arr => { if (!arr || !arr.length) return '—'; const s = [...arr].sort((a, b) => a - b); return s.length > 1 ? s[0] + '–' + s[s.length - 1] : String(s[0]); };
const extHref = n => '/ext/' + encodeURIComponent(n);
const pkgHref = n => '/pkg/' + encodeURIComponent(n);
const catHref = n => '/cate/' + encodeURIComponent(n);
const dimHref = n => '/list/' + encodeURIComponent(n);
const catVar = c => 'style="--seg:var(--c-' + esc(c) + ')"';

/* ---------------- dimension hues ----------------
   License, language, and repository values map onto a shared hue scale
   (--h-* tokens in style.css, aliased to the theme-aware category palette).
   The temperature encodes openness/familiarity: cool blues for permissive
   licenses and core languages, greens for the moderate middle, ambers and
   oranges for copyleft and modern stacks, reds for restrictive or rare,
   violet for commercial/exotic outliers. */
function licenseHue(v) {
  const s = (v || '').toLowerCase();
  if (!s || s === 'unknown') return '';
  if (s.startsWith('postgresql')) return 'pgblue';
  if (s.startsWith('mit')) return 'blue';
  if (/^(bsd|0bsd|isc|zlib|unlicense|cc0|wtfpl|public)/.test(s)) return 'cyan';
  if (s.startsWith('apache')) return 'green';
  if (/^(mpl|epl|cddl|osl)/.test(s)) return 'teal';
  if (/^(artistic|eupl|cecill)/.test(s)) return 'olive';
  if (s.startsWith('lgpl')) return 'amber';
  if (s.startsWith('gpl')) return 'orange';
  if (s.startsWith('agpl')) return 'red';
  return 'violet';
}
const LANG_HUES = {
  C: 'blue', 'C++': 'indigo', SQL: 'green', 'PL/pgSQL': 'teal', Go: 'teal',
  Rust: 'orange', Python: 'amber', JavaScript: 'amber', TypeScript: 'amber',
  Java: 'red', R: 'red', Ruby: 'red', Shell: 'maroon', Perl: 'maroon', Data: 'slate'
};
const REPO_HUES = { PGDG: 'pgblue', PIGSTY: 'green', CONTRIB: 'pgnavy', MIXED: 'teal' };
const langHue = v => LANG_HUES[v] || (v ? 'red' : '');
const repoHue = v => REPO_HUES[v] || '';
const hueVar = h => h ? 'style="--seg:var(--h-' + h + ')"' : '';
// color expression for one dimension value: category palette or hue scale
function dimSeg(dim, v) {
  if (dim === 'category') return 'var(--c-' + v + ')';
  const h = dim === 'license' ? licenseHue(v) : dim === 'lang' ? langHue(v) : dim === 'repo' ? repoHue(v) : '';
  return h ? 'var(--h-' + h + ')' : '';
}
const debounce = (fn, ms) => { let h; return (...a) => { clearTimeout(h); h = setTimeout(() => fn(...a), ms); }; };
const osLabel = os => {
  const [base, arch] = String(os || '').split('.');
  const family = base.startsWith('el') ? 'EL ' + base.slice(2) : base.startsWith('d') ? 'Debian ' + base.slice(1) : base.startsWith('u') ? 'Ubuntu ' + base.slice(1) : base;
  return family + (arch ? ' · ' + arch : '');
};

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

/* ---------------- lightweight code highlighting ----------------
   Self-contained tokenizers for the languages that dominate pgext docs:
   SQL, shell, and key/value config (yaml, postgresql.conf). Rules must not
   contain capturing groups; the first alternative that matches wins. */
const HL_RULES = {
  sql: {
    flags: 'gi',
    rules: [
      ['--[^\\n]*', 'tok-cmt'],
      ['/\\*[\\s\\S]*?\\*/', 'tok-cmt'],
      ["'(?:[^'\\n]|'')*'", 'tok-str'],
      ['\\b(?:select|insert|update|delete|create|drop|alter|extension|table|index|view|materialized|function|procedure|trigger|schema|database|from|where|join|left|right|inner|outer|full|cross|lateral|on|group|order|by|having|limit|offset|as|and|or|not|null|true|false|is|in|exists|between|like|ilike|cascade|with|values|into|set|returning|union|all|distinct|case|when|then|else|end|grant|revoke|comment|show|explain|analyze|vacuum|begin|commit|rollback|if|replace|to|using|primary|key|foreign|references|unique|default|constraint|add|column|type|language|returns|immutable|stable|volatile|security|definer|partition|of|for|each|row|execute|declare|loop|return|raise|notice|copy|owner)\\b', 'tok-kw'],
      ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
    ]
  },
  bash: {
    flags: 'gm',
    rules: [
      ['#[^\\n]*', 'tok-cmt'],
      ["'[^'\\n]*'", 'tok-str'],
      ['"(?:[^"\\\\\\n]|\\\\.)*"', 'tok-str'],
      ['\\$\\{[^}\\n]*\\}|\\$[A-Za-z_][A-Za-z0-9_]*', 'tok-var'],
      ['\\b(?:sudo|dnf|yum|apt|apt-get|zypper|pig|pgext|psql|pg_ctl|pg_dump|pg_restore|initdb|systemctl|service|curl|wget|git|make|cmake|cargo|python3?|pip3?|echo|export|source|cd|cp|mv|rm|mkdir|tar|install|tee|cat|grep|sed|awk|docker)\\b', 'tok-kw'],
      ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
    ]
  },
  conf: {
    flags: 'gm',
    rules: [
      ['#[^\\n]*', 'tok-cmt'],
      ["'[^'\\n]*'", 'tok-str'],
      ['"[^"\\n]*"', 'tok-str'],
      ['^\\s*[A-Za-z_][\\w.-]*(?=\\s*[:=])', 'tok-var'],
      ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
    ]
  }
};

function hlLang(language) {
  const l = String(language || '').toLowerCase();
  if (['sql', 'psql', 'postgresql', 'pgsql', 'plpgsql'].includes(l)) return 'sql';
  if (['bash', 'sh', 'shell', 'zsh', 'console', 'shell-session', 'terminal'].includes(l)) return 'bash';
  if (['yaml', 'yml', 'ini', 'conf', 'toml', 'properties', 'postgresql.conf', 'env'].includes(l)) return 'conf';
  return '';
}

function highlightCode(source, language) {
  const src = String(source == null ? '' : source);
  const spec = HL_RULES[hlLang(language)];
  if (!spec) return esc(src);
  const re = new RegExp(spec.rules.map(r => '(' + r[0] + ')').join('|'), spec.flags);
  let out = '', pos = 0, m;
  while ((m = re.exec(src))) {
    if (m.index > pos) out += esc(src.slice(pos, m.index));
    let cls = '';
    for (let g = 1; g < m.length; g++) if (m[g] !== undefined) { cls = spec.rules[g - 1][1]; break; }
    out += cls ? '<span class="' + cls + '">' + esc(m[0]) + '</span>' : esc(m[0]);
    pos = m.index + m[0].length;
    if (!m[0].length) re.lastIndex++;
  }
  return out + esc(src.slice(pos));
}

/* ---------------- curated markdown (pgext.doc usage manuals) ---------------- */
function mdSafeURL(raw) {
  const url = String(raw || '').trim().replace(/^<|>$/g, '');
  return /^(https?:\/\/|mailto:|\/|#)/i.test(url) ? url : '';
}

function mdInline(source) {
  const tokens = [];
  const keep = html => { const key = '\u0001' + tokens.length + '\u0002'; tokens.push(html); return key; };
  const labelHTML = s => esc(s)
    .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>')
    .replace(/\*([^*]+)\*/g, '<em>$1</em>');
  const linkHTML = (label, rawURL) => {
    const url = mdSafeURL(rawURL);
    if (!url) return labelHTML(label);
    const external = /^https?:\/\//i.test(url);
    return '<a href="' + esc(url) + '"' + (external ? ' target="_blank" rel="noopener"' : '') + '>' + labelHTML(label) + '</a>';
  };
  let text = String(source == null ? '' : source);
  text = text.replace(/`([^`\n]+)`/g, (_, code) => keep('<code>' + esc(code) + '</code>'));
  text = text.replace(/!\[([^\]]*)\]\(([^)\s]+)(?:\s+"[^"]*")?\)/g, (_, alt, rawURL) => {
    const url = mdSafeURL(rawURL);
    return url && /^https?:\/\//i.test(url) ? keep('<img src="' + esc(url) + '" alt="' + esc(alt) + '" loading="lazy">') : esc(alt);
  });
  text = text.replace(/\[([^\]]+)\]\(([^)\s]+)(?:\s+"[^"]*")?\)/g, (_, label, url) => keep(linkHTML(label, url)));
  text = text.replace(/<(https?:\/\/[^>\s]+)>/g, (_, url) => keep(linkHTML(url, url)));
  let html = esc(text)
    .replace(/\*\*\*([^*]+)\*\*\*/g, '<strong><em>$1</em></strong>')
    .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>')
    .replace(/~~([^~]+)~~/g, '<del>$1</del>')
    .replace(/(^|[\s(])\*([^*\n]+)\*(?=$|[\s).,!?:;])/g, '$1<em>$2</em>');
  html = html.replace(/\u0001(\d+)\u0002/g, (_, n) => tokens[Number(n)] || '');
  return html;
}

function mdPlain(source) {
  return String(source || '')
    .replace(/!\[([^\]]*)\]\([^)]*\)/g, '$1')
    .replace(/\[([^\]]+)\]\([^)]*\)/g, '$1')
    .replace(/[`*_~]/g, '').trim();
}

function mdSlug(source, used) {
  let base = mdPlain(source).toLowerCase().normalize('NFKD')
    .replace(/[^\p{L}\p{N}]+/gu, '-').replace(/^-+|-+$/g, '') || 'section';
  const count = used.get(base) || 0;
  used.set(base, count + 1);
  return 'doc-' + base + (count ? '-' + (count + 1) : '');
}

function mdTableRow(line) {
  let source = String(line || '').trim();
  if (source.startsWith('|')) source = source.slice(1);
  if (source.endsWith('|')) source = source.slice(0, -1);
  const cells = []; let cell = '', code = false, escaped = false;
  for (const ch of source) {
    if (escaped) { cell += ch; escaped = false; continue; }
    if (ch === '\\') { escaped = true; cell += ch; continue; }
    if (ch === '`') { code = !code; cell += ch; continue; }
    if (ch === '|' && !code) { cells.push(cell.trim()); cell = ''; continue; }
    cell += ch;
  }
  cells.push(cell.trim());
  return cells;
}

function mdTableDivider(line) {
  const cells = mdTableRow(line);
  return cells.length > 0 && cells.every(c => /^:?-{3,}:?$/.test(c));
}

function mdList(lines, start) {
  const items = [];
  let i = start;
  while (i < lines.length) {
    const m = lines[i].match(/^(\s*)([-+*]|\d+[.)])\s+(.*)$/);
    if (!m) break;
    items.push({ indent: m[1].replace(/\t/g, '    ').length, ordered: /^\d/.test(m[2]), text: m[3] });
    i++;
  }
  let pos = 0;
  const level = indent => {
    const ordered = items[pos].ordered;
    let html = '<' + (ordered ? 'ol' : 'ul') + '>';
    while (pos < items.length && items[pos].indent === indent) {
      const item = items[pos++];
      html += '<li>' + mdInline(item.text);
      if (pos < items.length && items[pos].indent > indent) html += level(items[pos].indent);
      html += '</li>';
      if (pos < items.length && items[pos].indent < indent) break;
    }
    return html + '</' + (ordered ? 'ol' : 'ul') + '>';
  };
  return { html: level(items[0].indent), next: i };
}

function renderMD(source, options) {
  const src = String(source || '').replace(/\r\n?/g, '\n');
  const lines = src.split('\n');
  const out = [], toc = [], used = new Map();
  const usageMode = Boolean(options && options.usage);
  const isBlock = (line, next) => /^\s*```/.test(line) || /^#{1,6}\s+/.test(line)
    || /^\s*(?:[-+*]|\d+[.)])\s+/.test(line) || /^>\s?/.test(line)
    || /^(?:\s*[-*_]){3,}\s*$/.test(line)
    || (line.includes('|') && next != null && mdTableDivider(next));
  const heading = (level, title) => {
    // usage docs render inside a regular page section whose own title is an
    // H2, so every doc heading sits one level below it
    if (usageMode) level = Math.min(Math.max(level, 3), 6);
    const id = mdSlug(title, used);
    const clean = mdPlain(title);
    toc.push({ level, title: clean, id });
    out.push('<h' + level + ' id="' + id + '">' + mdInline(title) + '</h' + level + '>');
  };
  let i = 0;
  while (i < lines.length) {
    const line = lines[i];
    if (!line.trim()) { i++; continue; }

    const fence = line.match(/^\s*```\s*([^\s`]*)?.*$/);
    if (fence) {
      const language = (fence[1] || 'text').toLowerCase();
      const code = []; i++;
      while (i < lines.length && !/^\s*```/.test(lines[i])) code.push(lines[i++]);
      if (i < lines.length) i++;
      const raw = code.join('\n');
      out.push('<div class="md-code"><div class="md-codebar"><span>' + esc(language) + '</span><button data-copy="' + esc(raw) + '">copy</button></div>'
        + '<pre><code class="language-' + esc(language) + '">' + highlightCode(raw, language) + '</code></pre></div>');
      continue;
    }

    const atx = line.match(/^(#{1,6})\s+(.+?)\s*#*$/);
    if (atx) { heading(atx[1].length, atx[2]); i++; continue; }
    if (i + 1 < lines.length && /^\s*(?:={3,}|-{3,})\s*$/.test(lines[i + 1]) && line.trim()) {
      heading(lines[i + 1].includes('=') ? 2 : 3, line.trim()); i += 2; continue;
    }
    if (/^\s*(?:-{3,}|\*{3,}|_{3,})\s*$/.test(line)) { out.push('<hr>'); i++; continue; }

    if (line.includes('|') && i + 1 < lines.length && mdTableDivider(lines[i + 1])) {
      const head = mdTableRow(line); i += 2;
      const body = [];
      while (i < lines.length && lines[i].includes('|') && lines[i].trim()) body.push(mdTableRow(lines[i++]));
      let html = '<div class="md-table"><table><thead><tr>' + head.map(c => '<th>' + mdInline(c) + '</th>').join('') + '</tr></thead><tbody>';
      for (const row of body) html += '<tr>' + head.map((_, n) => '<td>' + mdInline(row[n] || '') + '</td>').join('') + '</tr>';
      out.push(html + '</tbody></table></div>');
      continue;
    }

    if (/^\s*(?:[-+*]|\d+[.)])\s+/.test(line)) {
      const list = mdList(lines, i); out.push(list.html); i = list.next; continue;
    }
    if (/^>\s?/.test(line)) {
      const quote = [];
      while (i < lines.length && /^>\s?/.test(lines[i])) quote.push(lines[i++].replace(/^>\s?/, ''));
      // GitHub-flavored alerts: > [!NOTE] / [!TIP] / [!IMPORTANT] / [!WARNING] /
      // [!CAUTION] — the marker may stand alone or share its line with text.
      const alert = quote[0] && quote[0].match(/^\s*\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]\s*(.*)$/i);
      const inner = source => renderMD(source.join('\n')).html || ('<p>' + mdInline(source.join(' ')) + '</p>');
      if (alert) {
        const kind = alert[1].toUpperCase();
        const label = { NOTE: ['Note', '提示'], TIP: ['Tip', '技巧'], IMPORTANT: ['Important', '重要'], WARNING: ['Warning', '警告'], CAUTION: ['Caution', '当心'] }[kind];
        const icon = { NOTE: 'ⓘ', TIP: '☛', IMPORTANT: '❢', WARNING: '⚠', CAUTION: '⛒' }[kind];
        const rest = alert[2] ? [alert[2], ...quote.slice(1)] : quote.slice(1);
        out.push('<div class="md-alert md-alert-' + kind.toLowerCase() + '"><p class="md-alert-title"><span>' + icon + '</span>'
          + (LANG === 'zh' ? label[1] : label[0]) + '</p>' + inner(rest) + '</div>');
      } else {
        out.push('<blockquote>' + inner(quote) + '</blockquote>');
      }
      continue;
    }

    const para = [line.trim()]; i++;
    while (i < lines.length && lines[i].trim() && !isBlock(lines[i], lines[i + 1])) {
      if (i + 1 < lines.length && /^\s*(?:={3,}|-{3,})\s*$/.test(lines[i + 1])) break;
      para.push(lines[i++].trim());
    }
    out.push('<p>' + mdInline(para.join(' ')) + '</p>');
  }
  const words = src.trim() ? src.trim().split(/\s+/).length : 0;
  const cjk = (src.match(/[\u3400-\u9fff]/g) || []).length;
  const minutes = Math.max(1, Math.ceil(LANG === 'zh' ? Math.max(cjk / 500, words / 260) : words / 220));
  return { html: out.join('\n'), toc, minutes };
}

/* ---------------- landing state / filter engine ---------------- */
// The category state key stays `S.cat`; its URL parameter is spelled `cate`
// (with `cat` accepted for old links) — see pushState/readState.
const FILTER_KEYS = ['q', 'repo', 'license', 'lang', 'pg', 'os', 'kind', 'lifecycle', 'scope',
  'vendor', 'kernel', 'tag', 'pkg', 'capability', 'build', 'docs', 'relation', 'pgrx', 'active'];
const DEFAULT_STATE = {
  q: '', cat: '', repo: '', license: '', lang: '', pg: '', os: '', kind: '', lifecycle: '', scope: '',
  vendor: '', kernel: '', tag: '', pkg: '', capability: '', build: '', docs: '', relation: '', pgrx: '', active: '',
  sort: '', entity: 'ext', layout: 'card'
};
let S = { ...DEFAULT_STATE };
// An empty S.sort means "default order": stars, for both entities.
const effSort = () => S.sort || 'stars';

function migrateLegacyHash() {
  const raw = location.hash.startsWith('#/') ? location.hash.slice(1) : '';
  if (!raw) return false;
  const qi = raw.indexOf('?');
  let path = qi >= 0 ? raw.slice(0, qi) : raw;
  const query = qi >= 0 ? raw.slice(qi + 1) : '';
  if (path.startsWith('/cat/')) path = '/cate/' + path.slice(5);
  history.replaceState(null, '', path + (query ? '?' + query : ''));
  return true;
}

function parseRoute() {
  migrateLegacyHash();
  const path = location.pathname.length > 1 ? location.pathname.replace(/\/+$/, '') : '/';
  return { path, params: new URLSearchParams(location.search) };
}

function navigateTo(url, replace) {
  history[replace ? 'replaceState' : 'pushState'](null, '', url);
  route();
  if (location.hash && !location.hash.startsWith('#/')) {
    requestAnimationFrame(() => {
      const id = decodeURIComponent(location.hash.slice(1));
      const target = document.getElementById(id);
      if (target) target.scrollIntoView({ block: 'start' });
    });
  }
}

function pushState() {
  const p = new URLSearchParams();
  if (S.cat) p.set('cate', S.cat);
  for (const k of FILTER_KEYS) if (S[k]) p.set(k, S[k]);
  if (S.sort) p.set('sort', S.sort);
  if (S.entity !== 'ext') p.set('entity', S.entity);
  if (S.layout !== 'card') p.set('layout', S.layout);
  const qs = p.toString();
  history.replaceState(null, '', '/' + (qs ? '?' + qs : ''));
}
function readState(params) {
  S = { ...DEFAULT_STATE };
  for (const k of [...FILTER_KEYS, 'sort', 'entity', 'layout']) {
    const v = params.get(k); if (v) S[k] = v;
  }
  S.cat = (params.get('cate') || params.get('cat') || '').toUpperCase();
  if (!S.lang && params.get('lng')) S.lang = params.get('lng'); // legacy param name
  if (S.lang === 'zh' || S.lang === 'en') S.lang = ''; // ?lang=zh|en selects the UI language, not a filter
  // Keep old shared links useful while moving from one overloaded view switch
  // to independent entity and presentation controls.
  const legacyView = params.get('view');
  if (legacyView === 'packages') { S.entity = 'pkg'; S.layout = 'card'; }
  if (legacyView === 'rows') { S.entity = 'ext'; S.layout = 'list'; }
  if (legacyView === 'grid') { S.entity = 'ext'; S.layout = 'card'; }
  if (!['ext', 'pkg'].includes(S.entity)) S.entity = 'ext';
  if (!['card', 'list'].includes(S.layout)) S.layout = 'card';
  S.pg = parsePGs(S.pg).join(',');
  try {
    if (parsePGs(S.pg).length === 1) localStorage.setItem('pgext.target.pg', S.pg);
    if (S.os) localStorage.setItem('pgext.target.os', S.os);
  } catch (err) {}
  if (parsePGs(S.pg).length === 1) INSTALL_PREF.pg = S.pg;
  if (S.os) INSTALL_PREF.os = S.os;
}

function parsePGs(value) {
  const seen = new Set();
  for (const raw of String(value || '').split(/[,\s]+/)) {
    const pg = Number.parseInt(raw, 10);
    if (Number.isInteger(pg) && pg > 0 && pg < 100) seen.add(pg);
  }
  return [...seen].sort((a, b) => b - a);
}

function effectiveFilters() {
  const f = {
    cat: S.cat, repo: S.repo, license: S.license, lang: S.lang, pg: S.pg, os: S.os,
    kind: S.kind, lifecycle: S.lifecycle, scope: S.scope, vendorQ: S.vendor.toLowerCase(), kernelQ: S.kernel.toLowerCase(),
    tag: S.tag, pkg: S.pkg, capability: S.capability, build: S.build, docs: S.docs,
    relation: S.relation, pgrx: S.pgrx, active: S.active, words: []
  };
  for (const tok of S.q.trim().split(/\s+/).filter(Boolean)) {
    const m = tok.match(/^(cat|cate|category|repo|license|lang|lng|pg|os|target|type|kind|life|lifecycle|kernel|vendor|tag|tags|pkg|package|project|cap|capability|feature|build|builder|doc|docs|rel|relation|dependency|pgrx|active|year|is):(.+)$/i);
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

function targetAvailable(e, pg, os) {
  if (!e.targets || !e.targets.length || !OSS.length) return false;
  const oi = os ? OSIDX[os] : -1;
  if (os && oi == null) return false;
  for (const idx of e.targets) {
    const pi = Math.floor(idx / OSS.length), ti = idx % OSS.length;
    if (os && ti !== oi) continue;
    if (pg && PGS[pi] !== pg) continue;
    return true;
  }
  return false;
}

function targetsAvailable(e, pgs, os) {
  return pgs.length ? pgs.every(pg => targetAvailable(e, pg, os)) : targetAvailable(e, 0, os);
}

function capabilityMatches(e, value) {
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

function buildMatches(e, value) {
  switch (String(value || '').toLowerCase()) {
    case 'rpm': return Boolean(e.buildbits & B.RPM);
    case 'deb': return Boolean(e.buildbits & B.DEB);
    case 'pgrx': return Boolean(e.buildbits & B.PGRX);
    case 'source': return Boolean(e.buildbits & B.SOURCE);
    case 'packaged': case 'binary-package': return e.avail;
    default: return false;
  }
}

function docsClass(e) {
  if ((e.docbits & 3) === 3) return 'bilingual';
  if (e.docbits & 1) return 'english-only';
  if (e.docbits & 2) return 'chinese-only';
  return 'undocumented';
}

function relationMatches(e, value) {
  switch (String(value || '').toLowerCase()) {
    case 'requires': return Boolean(e.relationbits & R.REQUIRES);
    case 'required-by': return Boolean(e.relationbits & R.REQUIRED_BY);
    case 'see-also': return Boolean(e.relationbits & R.SEE_ALSO);
    case 'no-relations': return e.relationbits === 0;
    default: return false;
  }
}

function activeYear(e) {
  const value = String(e.active || '').slice(0, 4);
  return /^\d{4}$/.test(value) ? value : 'unknown';
}

// Catalog scope: packaged/unpacked split the universe by deliverability;
// 'source' survives as a legacy alias (unpackaged with a buildable source).
function scopeMatch(e, scope) {
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

function runFilter() {
  const f = effectiveFilters();
  const pgs = parsePGs(f.pg);
  const list = [];
  for (const e of EXT) {
    if (f.cat && e.cat !== f.cat) continue;
    if (f.repo && e.repo !== f.repo) continue;
    if (f.license && e.license.toLowerCase() !== f.license.toLowerCase()) continue;
    if (f.lang && e.lang.toLowerCase() !== f.lang.toLowerCase()) continue;
    if (pgs.length && !pgs.every(pg => e.pg.includes(pg))) continue;
    if (f.os && !targetsAvailable(e, pgs, f.os)) continue;
    if (f.kind && e.kind !== f.kind) continue;
    if (f.lifecycle && e.lifecycle.toLowerCase() !== f.lifecycle.toLowerCase()) continue;
    if (f.tag && !e.tags.some(tag => tag.toLowerCase() === f.tag.toLowerCase())) continue;
    if (f.pkg && e.pkg.toLowerCase() !== f.pkg.toLowerCase()) continue;
    if (f.capability && !capabilityMatches(e, f.capability)) continue;
    if (f.build && !buildMatches(e, f.build)) continue;
    if (f.docs && docsClass(e) !== f.docs) continue;
    if (f.relation && !relationMatches(e, f.relation)) continue;
    if (f.pgrx && e.pgrx.toLowerCase() !== f.pgrx.toLowerCase()) continue;
    if (f.active && activeYear(e) !== f.active) continue;
    if (f.scope && !scopeMatch(e, f.scope)) continue;
    if (f.vendorQ && !(e.vendor || '').toLowerCase().includes(f.vendorQ)) continue;
    if (f.kernelQ && !(e.kernel || '').toLowerCase().includes(f.kernelQ)) continue;
    let score = 0, drop = false;
    for (const w of f.words) {
      const nm = e.name.toLowerCase();
      if (nm === w) score += 120;
      else if (nm.startsWith(w)) score += 60;
      else if (nm.includes(w) || e.pkg.toLowerCase().includes(w)) score += 30;
      else if (e.tags.some(tag => tag.toLowerCase().includes(w))) score += 18;
      else if (e.en.toLowerCase().includes(w) || (e.zh && e.zh.includes(w)) || e.comment.toLowerCase().includes(w)) score += 10;
      else if (e.cat.toLowerCase() === w || (e.vendor || '').toLowerCase().includes(w) || (e.kernel || '').toLowerCase().includes(w) || e.lifecycle.toLowerCase().includes(w)) score += 8;
      else { drop = true; break; }
    }
    if (drop) continue;
    list.push([score, e]);
  }
  const hasQ = f.words.length > 0;
  const sort = effSort();
  list.sort((a, b) => {
    if (hasQ && b[0] !== a[0]) return b[0] - a[0];
    if (sort === 'stars') return (b[1].stars || 0) - (a[1].stars || 0) || a[1].name.localeCompare(b[1].name);
    if (sort === 'updated') return (b[1].active || '').localeCompare(a[1].active || '') || (b[1].stars || 0) - (a[1].stars || 0);
    return a[1].name.localeCompare(b[1].name);
  });
  return { f, list: list.map(x => x[1]) };
}

function buildSQL(f, n) {
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
        wh.push('EXISTS (SELECT 1 FROM pgext.pkg p WHERE p.pkg = universe.pkg AND p.state = \'AVAIL\' AND p.os = ' + lit(f.os) + ' AND p.pg = ' + pg + ')');
      }
    } else {
      wh.push('EXISTS (SELECT 1 FROM pgext.pkg p WHERE p.pkg = universe.pkg AND p.state = \'AVAIL\' AND p.os = ' + lit(f.os) + ')');
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
  else if (/^\d{4}$/.test(f.active)) wh.push("left(last_active::text, 4) = " + lit(f.active));
  if (f.scope === 'packaged') wh.push('packaged');
  if (f.scope === 'unpacked') wh.push('NOT packaged');
  if (f.scope === 'source') wh.push("NOT packaged AND (repo_url IS NOT NULL OR tarball IS NOT NULL)");
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

/* ---------------- shared components ---------------- */
const GH_SVG = '<svg width="13" height="13" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27s1.36.09 2 .27c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0 0 16 8c0-4.42-3.58-8-8-8Z"/></svg>';

/* Card corner badge — one rule everywhere (ext cards, pkg cards, hovercard):
   a GitHub repo shows the GitHub mark with the star count; any other primary
   URL shows a generic external-link icon. The badge always links out. */
const EXTLINK_SVG = '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.1" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">'
  + '<path d="M13.5 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6.5"/>'
  + '<path d="M14 3h7v7M21 3l-9.5 9.5"/></svg>';

function cardBadgeHTML(e) {
  const link = e.repoUrl || e.url;
  const isGH = /github\.com/i.test(e.repoUrl || '');
  if (link) {
    return '<a class="card-gh" href="' + esc(link) + '" target="_blank" rel="noopener" aria-label="' + esc(e.name) + ' upstream">'
      + (isGH ? GH_SVG : EXTLINK_SVG)
      + (e.stars != null ? '<b>' + fmtNum(e.stars) + '</b>' : '') + '</a>';
  }
  return e.stars != null ? '<span class="card-gh plain">' + GH_SVG + '<b>' + fmtNum(e.stars) + '</b></span>' : '';
}

// Card tags: category, language, license, package source, vendor, and a
// non-active lifecycle. The repo value already covers contrib delivery, and
// unpackaged entries read from the missing emphasis border — no extra tags.
// Language / license / repo tags carry their dimension hue (see dimSeg) and
// navigate to their /lang /repo listing pages; the license tag opens the
// official license text when the catalog records one.
const hueTag = (href, value, hue, external) => '<a class="tg' + (hue ? ' tg-hue' : '') + '" ' + (hue ? hueVar(hue) + ' ' : '')
  + 'href="' + esc(href) + '"' + (external ? ' target="_blank" rel="noopener"' : '') + '>' + esc(value) + '</a>';
const licenseTagHref = e => mdSafeURL(e.licenseUrl) || '/license/' + encodeURIComponent(e.license);
function cardTagsHTML(e) {
  return [
    '<a class="tg tg-cat" href="/?cate=' + encodeURIComponent(e.cat) + '">' + esc(e.cat) + '</a>',
    e.lang ? hueTag('/lang/' + encodeURIComponent(e.lang), e.lang, langHue(e.lang)) : '',
    e.license !== 'Unknown' ? hueTag(licenseTagHref(e), e.license, licenseHue(e.license), !!mdSafeURL(e.licenseUrl)) : '',
    e.repo !== 'n/a' ? hueTag('/repo/' + encodeURIComponent(e.repo), e.repo, repoHue(e.repo)) : '',
    e.vendor ? '<a class="tg tg-dim" href="/?vendor=' + encodeURIComponent(e.vendor) + '">' + esc(e.vendor) + '</a>' : '',
    e.lifecycle && e.lifecycle !== 'active' ? '<a class="tg tg-life" href="/?lifecycle=' + encodeURIComponent(e.lifecycle) + '">' + esc(e.lifecycle) + '</a>' : ''
  ].filter(Boolean).join('');
}

/* Extension card: the whole card is clickable (cover link) while every inner
   element — the corner badge, the package subtitle and the dimension tags —
   stays independently clickable above it. Packaged (pgext.extension) entries
   get an emphasized border so out-of-the-box extensions read at a glance. */
function tileHTML(e) {
  return '<li><article class="ecard' + (e.avail ? ' packaged' : '') + '" ' + catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '">'
    + '<a class="card-cover" href="' + extHref(e.name) + '" aria-label="' + esc(e.name) + '"></a>'
    + '<header class="card-head"><h3 class="card-name">' + esc(e.name) + '</h3>' + cardBadgeHTML(e) + '</header>'
    + '<p class="card-sub"><a href="' + pkgHref(e.pkg) + '">' + esc(e.pkg) + '</a>'
    + (e.ver ? '<span class="ver">' + esc(e.ver) + '</span>' : '') + '</p>'
    + '<p class="card-desc">' + esc(desc(e)) + '</p>'
    + '<footer class="card-tags">' + cardTagsHTML(e) + '</footer>'
    + '</article></li>';
}

function projectGroups(list) {
  const groups = new Map();
  for (const e of list) {
    if (!groups.has(e.pkg)) groups.set(e.pkg, []);
    groups.get(e.pkg).push(e);
  }
  const projects = [...groups.entries()].map(([pkg, matched]) => {
    const all = byPkg.get(pkg) || matched;
    const leadName = (all[0] && all[0].lead) || pkg;
    const lead = byName.get(leadName) || matched[0] || all[0];
    return { pkg, matched, all, lead };
  });
  // the PKG entity sorts by package identity, not by the member extensions
  // the groups were discovered through
  const sort = effSort();
  projects.sort((a, b) => {
    if (sort === 'stars') return (b.lead.stars || 0) - (a.lead.stars || 0) || a.pkg.localeCompare(b.pkg);
    if (sort === 'updated') return (b.lead.active || '').localeCompare(a.lead.active || '') || a.pkg.localeCompare(b.pkg);
    return a.pkg.localeCompare(b.pkg);
  });
  return projects;
}

/* Package card anatomy:
     PKG NAME                 badge (GitHub star / external link)
     lead extension · version           — gray subtitle, links to /ext
     description
     tags (from the lead extension)
   The full extension family lives in the package hovercard and /pkg page. */
function packageTileHTML(g) {
  const e = g.lead;
  const packaged = g.all.some(x => x.avail);
  return '<li><article class="ecard pkg-ecard' + (packaged ? ' packaged' : '') + '" ' + catVar(e.cat) + ' data-hover-pkg="' + esc(g.pkg) + '">'
    + '<a class="card-cover" href="' + pkgHref(g.pkg) + '" aria-label="' + esc(g.pkg) + '"></a>'
    + '<header class="card-head"><h3 class="card-name">' + esc(g.pkg) + '</h3>' + cardBadgeHTML(e) + '</header>'
    + '<p class="card-sub"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a>'
    + (e.ver ? '<span class="ver">' + esc(e.ver) + '</span>' : '') + '</p>'
    + '<p class="card-desc">' + esc(desc(e)) + '</p>'
    + '<footer class="card-tags">' + cardTagsHTML(e) + '</footer>'
    + '</article></li>';
}

function licenseCellHTML(e) {
  const url = mdSafeURL(e.licenseUrl);
  if (e.license === 'Unknown') return '—';
  return url
    ? '<a href="' + esc(url) + '" target="_blank" rel="noopener">' + esc(e.license) + '</a>'
    : esc(e.license);
}

function rowHTML(e) {
  return '<tr ' + catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '"><td class="r-name"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a></td>'
    + '<td class="r-mono r-ver">' + esc(e.ver || '—') + '</td>'
    + '<td class="r-desc">' + esc(desc(e)) + '</td>'
    + '<td class="r-cat">' + esc(e.cat) + '</td>'
    + '<td class="r-mono r-lic" title="' + esc(e.license) + '">' + licenseCellHTML(e) + '</td>'
    + '<td class="r-mono">' + esc(e.lang) + '</td>'
    + '<td class="r-mono">' + esc(pgRange(e.pg)) + '</td>'
    + '<td class="r-num">' + (e.stars ? fmtNum(e.stars) : '') + '</td></tr>';
}

/* Package row: the first cell stacks the package name over its delivered
   extensions (one per line, each navigating to /ext) when there is more than
   one; single-extension packages stay a plain row. */
function packageRowHTML(g) {
  const e = g.lead;
  const cats = [...new Set(g.all.map(x => x.cat))];
  const pgs = [...new Set(g.all.flatMap(x => x.pg))].sort((a, b) => b - a);
  const members = g.all.slice().sort((a, b) => a.name.localeCompare(b.name));
  const subs = g.all.length > 1
    ? '<div class="r-exts">' + members.map(x => '<a href="' + extHref(x.name) + '">' + esc(x.name) + '</a>').join('') + '</div>' : '';
  return '<tr ' + catVar(e.cat) + ' data-hover-pkg="' + esc(g.pkg) + '"><td class="r-name r-pkg"><a class="r-pkg-name" href="' + pkgHref(g.pkg) + '">' + esc(g.pkg) + '</a>' + subs + '</td>'
    + '<td class="r-desc">' + esc(desc(e)) + '</td>'
    + '<td class="r-cat">' + esc(cats.join(' · ')) + '</td>'
    + '<td class="r-mono r-lic" title="' + esc(e.license) + '">' + licenseCellHTML(e) + '</td>'
    + '<td class="r-mono">' + esc(pgRange(pgs)) + '</td><td class="r-num">' + (e.stars ? fmtNum(e.stars) : '') + '</td></tr>';
}

function skel(lines) {
  let bars = '';
  for (let i = 0; i < (lines || 3); i++) bars += '<div class="bar" style="width:' + (55 + ((i * 17) % 40)) + '%"></div>';
  return '<div class="skel">' + bars + '</div>';
}
const hydrateErr = err => '<div class="hydrate-err">' + esc(t('hydrate.err', { msg: err && err.message || 'unknown' })) + '</div>';

/* the Pigsty honeycomb mark, inlined from pigsty.io */
const PIGSTY_LOGO = '<svg class="pigsty-logo" width="21" height="21" viewBox="0 0 24 24" aria-hidden="true"><g stroke="#fff" stroke-width=".67" stroke-linecap="round" stroke-linejoin="round">'
  + '<path d="M7.67 11.97 9.83 8.22h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#bbbbbb" fill-opacity=".95"/>'
  + '<path d="M7.67 19.47l2.16-3.75h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#de372c" fill-opacity=".85"/>'
  + '<path d="M14.16 15.75 16.33 12h4.33l2.16 3.75-2.16 3.75h-4.33Z" fill="#424242" fill-opacity=".9"/>'
  + '<path d="M14.16 8.23l2.17-3.75h4.33l2.16 3.75-2.16 3.75h-4.33Z" fill="#ffa269" fill-opacity=".9"/>'
  + '<path d="M7.67 4.5 9.83.75h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#419edb" fill-opacity=".9"/>'
  + '<path d="M1.15 8.23l2.16-3.75h4.33l2.17 3.75-2.17 3.75H3.31Z" fill="#2f6793" fill-opacity=".9"/>'
  + '<path d="M1.18 15.74l2.17-3.75h4.33l2.17 3.75-2.17 3.75H3.35Z" fill="#53ac79" fill-opacity=".9"/>'
  + '</g></svg>';

/* an isometric die — one click, one (packaged-weighted) random extension */
const ICON_DICE = '<svg width="17" height="17" viewBox="0 0 18 18" fill="none" stroke="currentColor" stroke-width="1.25" stroke-linejoin="round" aria-hidden="true">'
  + '<path d="M9 1.5 15.6 5.2v7.4L9 16.5 2.4 12.6V5.2L9 1.5Z"/>'
  + '<path d="M2.4 5.2 9 8.9l6.6-3.7M9 8.9v7.6"/>'
  + '<circle cx="9" cy="5.2" r=".85" fill="currentColor" stroke="none"/>'
  + '<circle cx="5" cy="10.1" r=".8" fill="currentColor" stroke="none"/><circle cx="6.7" cy="12.7" r=".8" fill="currentColor" stroke="none"/>'
  + '<circle cx="11.2" cy="12.5" r=".8" fill="currentColor" stroke="none"/><circle cx="12.4" cy="11.2" r=".8" fill="currentColor" stroke="none"/><circle cx="13.6" cy="9.9" r=".8" fill="currentColor" stroke="none"/></svg>';

// Icon-only nav actions: language (文/A), theme (sun/moon showing the current
// mode, click flips it), and GitHub — bare glyphs, no button chrome.
const ICON_LANG = '<svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">'
  + '<text x="0" y="11.5" font-size="11" font-family="system-ui,sans-serif">文</text>'
  + '<text x="10.5" y="19" font-size="10.5" font-weight="700" font-family="system-ui,sans-serif">A</text></svg>';
const ICON_SUN = '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" aria-hidden="true">'
  + '<circle cx="12" cy="12" r="4.4"/>'
  + '<path d="M12 2.6v2.1M12 19.3v2.1M2.6 12h2.1M19.3 12h2.1M5.2 5.2l1.5 1.5M17.3 17.3l1.5 1.5M18.8 5.2l-1.5 1.5M6.7 17.3l-1.5 1.5"/></svg>';
const ICON_MOON = '<svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">'
  + '<path d="M20.6 14.6A8.7 8.7 0 0 1 9.4 3.4a8.7 8.7 0 1 0 11.2 11.2Z"/></svg>';

function navHTML(active) {
  return '<div class="wrap nav-in">'
    + '<a class="brand" href="/">' + PIGSTY_LOGO + '<span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<nav class="nav-links">'
    + '<span class="nav-drop"><a href="/matrix" aria-current="' + (active === 'matrix') + '">' + t('nav.matrix') + '</a>'
    + '<div class="nav-menu"><a class="nav-menu-all" href="/matrix">' + bi('Overview', '全景总览') + '</a>'
    + OSS.map(os => '<a href="/os/' + encodeURIComponent(os) + '">' + esc(os) + '</a>').join('')
    + '<span class="nav-menu-sep"></span>'
    + '<div class="nav-menu-pgs">' + PGS.map(pg => '<a href="/pg/' + pg + '">PG' + pg + '</a>').join('') + '</div>'
    + '<span class="nav-menu-sep"></span>'
    + '<div class="nav-menu-pgs">' + ['PGDG', 'PIGSTY', 'CONTRIB', 'MIXED'].map(r => '<a href="/repo/' + r + '">' + r + '</a>').join('') + '</div>'
    + '</div></span>'
    + '<span class="nav-drop"><a href="/list" aria-current="' + (active === 'browse') + '">' + t('nav.browse') + '</a>'
    + '<div class="nav-menu nav-menu-index">'
    + '<div class="nav-idx-cats">' + CAT_ORDER.map(c => '<a href="' + catHref(c) + '" style="--seg:var(--c-' + c + ')" data-tip="' + esc(catName(c)) + '"><i></i>' + c + '</a>').join('') + '</div>'
    + DIM_GROUPS.map(([, dims]) => '<div class="nav-idx-col">'
      + dims.map(d => '<a href="' + dimHref(d) + '">' + t(DIMS[d].label) + '</a>').join('') + '</div>').join('')
    + '</div></span>'
    + '<a href="/about" aria-current="' + (active === 'about') + '">' + t('nav.about') + '</a>'
    + '</nav><span class="nav-spacer"></span><div class="nav-actions">'
    + '<div class="nav-search"><input id="nav-q" type="search" autocomplete="off" spellcheck="false" placeholder="'
    + esc(bi('Press / to search', '按 / 开始搜索')) + '" aria-label="search extensions"><kbd>/</kbd>'
    + '<div class="nav-sugg" id="nav-sugg" hidden></div></div>'
    + '<button class="nav-ico" id="random-ext" aria-label="random extension" data-tip="' + esc(bi('Random extension', '随机看一个扩展')) + '">' + ICON_DICE + '</button>'
    + '<button class="nav-ico" data-lang-toggle aria-label="' + (LANG === 'zh' ? 'switch to English' : '切换到中文') + '">' + ICON_LANG + '</button>'
    + '<button class="nav-ico" id="theme-toggle" aria-label="switch theme">' + themeIcon() + '</button>'
    + '<a class="nav-ico nav-github" href="https://github.com/pgsty/pgext" target="_blank" rel="noopener" aria-label="GitHub">' + GH_SVG.replace('width="13" height="13"', 'width="16" height="16"') + '</a>'
    + '</div></div>';
}

/* footer, after pigsty.io but leaner: brand + a few links, then a copyright
   line with the language switcher tucked into the bottom-right corner */
function footerHTML() {
  const group = (title, links) => '<div class="footer-col"><span>' + title + '</span>'
    + links.map(([href, label, ext]) => '<a href="' + href + '"' + (ext ? ' target="_blank" rel="noopener"' : '') + '>' + label + '</a>').join('') + '</div>';
  return '<div class="wrap footer-in">'
    + '<div class="footer-brand"><a class="brand" href="/">' + PIGSTY_LOGO + '<span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<p>' + bi('PostgreSQL Extension Catalog.', 'PostgreSQL 扩展目录') + '</p>'
    + '<span class="footer-snap">snapshot ' + esc(META.generated || '—') + '</span></div>'
    + group(bi('Catalog', '目录'), [['/matrix', t('nav.matrix')], ['/list', t('nav.browse')], ['/about', t('nav.about')]])
    + group(bi('Resources', '资源'), [['/api/v1/meta', 'API', true], ['https://github.com/pgsty/pgext', 'GitHub', true], ['https://github.com/pgsty/pgext/blob/main/db/universe.csv', bi('Raw Data', '原始数据'), true]])
    + group('Pigsty', [['https://pigsty.io', 'pigsty.io', true], ['https://pigsty.cc', 'pigsty.cc', true], ['https://pgext.cloud', 'pgext.cloud', true]])
    + '</div>'
    + '<div class="wrap footer-bottom">'
    + '<span>© 2018-2026 <a href="https://pigsty.io" target="_blank" rel="noopener">Pigsty</a> · <a href="https://vonng.com/en" target="_blank" rel="noopener">Ruohang Feng</a></span>'
    + '<span class="footer-langs"><button data-lang-set="zh" aria-pressed="' + (LANG === 'zh') + '">中文</button><i>·</i>'
    + '<button data-lang-set="en" aria-pressed="' + (LANG === 'en') + '">English</button></span></div>';
}

/* ---------------- view: home ---------------- */
function homeHTML() {
  const stats = [
    [fmtInt(N_ALL), t('hero.s1')],
    [fmtInt(N_AVAIL), t('hero.s2')],
    [fmtInt(N_PKGS), t('hero.s3')],
    [fmtInt(OSS.length), t('hero.s4')],
    [fmtInt(PGS.length), t('hero.s5')],
    [fmtInt(N_SLOTS), t('hero.s6')]
  ].map(([n, l]) => '<li><span class="num">' + n + '</span><span class="lbl">' + l + '</span></li>').join('');
  return '<section class="hero wrap">'
    + '<p class="eyebrow">' + t('hero.eyebrow') + '</p>'
    + '<h1 class="oneline">' + t('hero.title') + '</h1>'
    + '<p class="hero-sub oneline">' + t('hero.sub', { all: fmtInt(N_ALL), avail: fmtInt(N_AVAIL), os: fmtInt(OSS.length) }) + '</p>'
    + '<ul class="hero-stats">' + stats + '</ul>'
    + '<div class="universe"><canvas id="ufield" height="80" aria-hidden="true" style="width:100%;display:block;cursor:crosshair"></canvas></div>'
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
  const sortOpts = [['name', t('sort.name')], ['stars', t('sort.stars')], ['updated', t('sort.updated')]].map(([v, l]) =>
    '<option value="' + v + '"' + (effSort() === v ? ' selected' : '') + '>' + l + '</option>').join('');

  // The SCOPE row cascades: every count below it is computed over the
  // scope-filtered subset, so narrowing the scope renarrates the whole panel.
  const scopedEXT = S.scope ? EXT.filter(e => scopeMatch(e, S.scope)) : EXT;
  const countMap = dim => new Map(dimValues(dim, scopedEXT).map(x => [String(x.v), x.n]));
  const button = (key, value, label, on, count, cls, style, tip) => {
    const data = key === 'pg' ? ' data-pg-toggle="' + esc(value) + '"' : ' data-fkey="' + key + '" data-fval="' + esc(value) + '"';
    return '<button class="facet-btn' + (cls ? ' ' + cls : '') + '"' + data + ' aria-pressed="' + on + '"'
      + (style ? ' style="' + style + '"' : '') + (tip ? ' data-tip="' + esc(tip) + '"' : '') + '>'
      + (cls === 'category' || cls === 'hued' ? '<i></i>' : '') + '<span>' + label + '</span>'
      + (count == null ? '' : '<small>' + fmtInt(count) + '</small>') + '</button>';
  };
  // facet button carrying its dimension hue as a colored dot
  const huedButton = (key, value, label, on, count, hue) =>
    button(key, value, label, on, count, hue ? 'hued' : '', hue ? '--seg:var(--h-' + hue + ')' : '');
  const any = (key, on) => button(key, '', t('filter.any'), on, null, 'any');
  // Each facet header links to its full index page under /list.
  const row = (label, values, cls, href) => '<div class="facet-row' + (cls ? ' ' + cls : '') + '"><div class="facet-name">'
    + (href ? '<a class="facet-head" href="' + href + '"><b>' + label + '</b><i>→</i></a>' : '<b>' + label + '</b>')
    + '</div><div class="facet-values">' + values + '</div></div>';

  const scopeValues = [
    ['', t('chip.all'), N_ALL], ['packaged', t('chip.packaged'), N_AVAIL], ['unpacked', t('chip.unpacked'), N_ALL - N_AVAIL],
    ['kernel', t('chip.kernel'), N_KERNEL], ['vendor', t('chip.vendor'), N_VENDOR], ['contrib', t('chip.contrib'), N_CONTRIB]
  ].map(([v, label, count]) => v === 'contrib'
    ? huedButton('scope', v, label, S.scope === v, count, 'pgnavy')
    : button('scope', v, label, S.scope === v, count)).join('');

  // Categories: Any stands alone on the left; the 16 codes align 8 × 2.
  const catCounts = countMap('category');
  const categoryValues = any('cat', !S.cat) + CAT_ORDER.map(c =>
    button('cat', c, '<code>' + c + '</code>', S.cat === c, catCounts.get(c) || 0, 'category', '--seg:var(--c-' + c + ')', catName(c))).join('');

  const pgCounts = countMap('pg'), selectedPGs = parsePGs(S.pg);
  const pgValues = [...new Set([...PGS, ...selectedPGs])].sort((a, b) => b - a);
  const pgButtons = button('pg', '', t('filter.any'), !selectedPGs.length, null, 'any')
    + pgValues.map(pg => button('pg', String(pg), 'PG ' + pg, selectedPGs.includes(pg), pgCounts.get(String(pg)) || 0)).join('');

  const langCounts = countMap('lang');
  const langButtons = any('lang', !S.lang) + [...langCounts.entries()]
    .map(([v, count]) => huedButton('lang', v, esc(v), S.lang === v, count, langHue(v))).join('');
  // Licenses: values with ≤2 members collapse into the index-page link.
  // Like the category row, Any stands alone on the left so the license
  // chips wrap as one aligned block on the right.
  const allLicenses = dimValues('license', scopedEXT);
  const licenses = allLicenses.filter(x => x.n > 2);
  if (S.license && !licenses.some(x => x.v === S.license)) {
    licenses.push(allLicenses.find(x => x.v === S.license) || { v: S.license, n: 0 });
  }
  const licenseButtons = any('license', !S.license)
    + '<span class="facet-sub">' + licenses.map(x => huedButton('license', x.v, esc(x.v), S.license === x.v, x.n, licenseHue(x.v))).join('')
    + '<a class="facet-more" href="' + dimHref('license') + '">' + t('filter.license.more', { n: fmtInt(allLicenses.length) }) + '</a></span>';
  const facets = '<section class="facet-panel" aria-label="filters">'
    + row(t('filter.scope'), scopeValues, 'scope-facet', dimHref('distribution'))
    + row(t('filter.category'), categoryValues, 'category-facet', dimHref('category'))
    + row(t('filter.license'), licenseButtons, 'license-facet', dimHref('license'))
    + row(t('filter.lang'), langButtons, 'lang-facet', dimHref('lang'))
    + row(t('filter.pg'), pgButtons, 'pg-facet', dimHref('pg')) + '</section>';

  const projects = S.entity === 'pkg' ? projectGroups(list) : [];
  const itemTotal = S.entity === 'pkg' ? projects.length : list.length;
  let results;
  if (!itemTotal) {
    results = '<div class="result-empty">' + t('wall.empty') + '</div>';
  } else if (S.entity === 'pkg' && S.layout === 'card') {
    results = '<ul class="wall ext-wall pkg-wall">' + projects.map(packageTileHTML).join('') + '</ul>';
  } else if (S.entity === 'pkg') {
    results = '<div class="rows"><div class="rows-scroll"><table class="ext-table pkg-table"><thead><tr>'
      + '<th>' + t('rows.package') + '</th><th>' + t('rows.desc') + '</th><th>' + t('rows.cat') + '</th>'
      + '<th>' + t('rows.license') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
      + '</tr></thead><tbody>' + projects.map(packageRowHTML).join('') + '</tbody></table></div></div>';
  } else if (S.layout === 'list') {
    results = '<div class="rows"><div class="rows-scroll"><table class="ext-table"><thead><tr>'
      + '<th>' + t('rows.name') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th><th>' + t('rows.cat') + '</th>'
      + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
      + '</tr></thead><tbody>' + list.map(rowHTML).join('') + '</tbody></table></div></div>';
  } else {
    results = '<ul class="wall ext-wall">' + list.map(tileHTML).join('') + '</ul>';
  }

  // Dimensions without a dedicated facet row surface here when active, so
  // filters arriving via card tags or shared links stay visible and clearable.
  const advanced = [
    ['repo', 'repo'], ['lifecycle', 'lifecycle'], ['os', 'os'], ['kind', 'kind'],
    ['vendor', 'vendor'], ['kernel', 'kernel'], ['tag', 'tag'], ['pkg', 'package'],
    ['capability', 'capability'], ['build', 'build'], ['docs', 'docs'], ['relation', 'relation'],
    ['pgrx', 'pgrx'], ['active', 'activity']
  ].filter(([key]) => S[key]);
  const activeFilters = advanced.length
    ? '<div class="active-filters"><span>' + t('filter.active') + '</span>' + advanced.map(([key, dim]) => {
      const seg = dimSeg(dim, S[key]);
      return '<button data-fkey="' + key + '" data-fval="' + esc(S[key]) + '"><b>' + t(DIMS[dim].label) + '</b>'
        + '<code' + (seg ? ' style="color:' + seg + '"' : '') + '>' + esc(dimValueName(dim, S[key])) + '</code><i>×</i></button>';
    }).join('') + '</div>'
    : '';
  const toolbar = '<div class="catalog-toolbar"><div class="result-summary"><strong>' + fmtInt(itemTotal) + '</strong><span>'
    + t(S.entity === 'pkg' ? 'filter.results.pkg' : 'filter.results.ext') + '</span></div><div class="catalog-controls">'
    + '<a class="dimension-link" href="/list">＋ ' + t('filter.dimensions') + '</a>'
    + '<label class="sortctl"><span>' + t('sort.label') + '</span>'
    + '<select data-skey="sort" aria-label="sort">' + sortOpts + '</select></label>'
    + '<span class="viewtoggle" role="group" aria-label="catalog entity">'
    + '<button data-entity="ext" aria-pressed="' + (S.entity === 'ext') + '">' + t('view.ext') + '</button>'
    + '<button data-entity="pkg" aria-pressed="' + (S.entity === 'pkg') + '">' + t('view.pkg') + '</button></span>'
    + '<span class="viewtoggle" role="group" aria-label="layout">'
    + '<button data-layout="card" aria-pressed="' + (S.layout === 'card') + '">' + t('view.card') + '</button>'
    + '<button data-layout="list" aria-pressed="' + (S.layout === 'list') + '">' + t('view.list') + '</button></span></div></div>';

  return '<button class="sql-readout" data-sql="' + esc(sql) + '" title="' + esc(t('search.copy')) + '">' + sqlHTML + '</button>'
    + activeFilters + facets + toolbar + results;
}

/* ---------------- universe field canvas ---------------- */
let ucells = null;
function drawField() {
  const cv = document.getElementById('ufield');
  if (!cv) return;
  const dpr = window.devicePixelRatio || 1;
  const W = cv.clientWidth || cv.parentElement.clientWidth;
  const pitch = W > 900 ? 8 : W > 560 ? 7 : 6;
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
// compact attribute chips for family tables and the package hovercard
function attrChips(e) {
  const chips = [];
  if (e.flags & F.DDL) chips.push(['DDL', 'attr.ddl']);
  if (e.flags & F.LOAD) chips.push(['LOAD', 'attr.load']);
  if (e.flags & F.LIB) chips.push(['LIB', 'attr.lib']);
  if (e.flags & F.BIN) chips.push(['BIN', 'attr.bin']);
  if (e.flags & F.TRUSTED) chips.push(['TRUST', 'attr.trusted']);
  if (e.flags & F.RELOC) chips.push(['RELOC', 'attr.reloc']);
  return chips.map(([code, key]) => '<span class="attr-chip" data-tip="' + esc(t(key)) + '">' + code + '</span>').join('');
}

/* small fixed "on this page" index in the right viewport gutter; it only
   appears when the screen is wide enough that it never covers the content */
function pageTocHTML(items) {
  return '<nav class="page-toc" aria-label="' + esc(t('ext.onpage')) + '"><span>' + t('ext.onpage') + '</span>'
    + items.map(([id, label]) => '<button data-scroll="' + id + '">' + label + '</button>').join('') + '</nav>';
}

function extTocHTML() {
  return pageTocHTML([
    ['ext-overview', 'ext.overview'], ['ext-relations', 'ext.relations'], ['ext-packages', 'ext.veravail'],
    ['ext-build', 'ext.build'], ['ext-install', 'ext.install'], ['ext-usage', 'ext.docs']
  ].map(([id, key]) => [id, t(key)]));
}

function extHTML(name) {
  const e = byName.get(name);
  if (!e) return notFoundHTML(name);
  const providerNote = (e.kernel || e.vendor) && !e.avail
    ? '<div class="notice provider">☁ ' + t('ext.providerNote', { provider: esc([e.vendor, e.kernel].filter(Boolean).join(' · ')) }) + '</div>' : '';
  const lifecycleNote = ['deprecated', 'archived', 'abandoned'].includes(e.lifecycle)
    ? '<div class="notice lifecycle">⚠ ' + t('ext.lifecycleNote', { state: '<b>' + esc(e.lifecycle) + '</b>' }) + '</div>' : '';
  const stateChip = e.avail
    ? '<span class="state-chip ok">● ' + t('state.avail') + '</span>'
    : '<span class="state-chip">○ ' + t('state.na') + '</span>';

  return '<article class="page wrap manual-page ext-page">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="' + catHref(e.cat) + '" style="color:var(--c-' + e.cat + ')">' + e.cat + '</a><span class="sep">/</span>'
    + '<span class="here">' + esc(e.name) + '</span></nav>'
    + '<header class="detail-hero ext-detail-hero"><div class="ext-hero-grid"><div class="ext-hero-main">'
    + '<div class="detail-kicker-row"><span class="detail-kicker">EXTENSION</span>' + stateChip
    + '<span class="hero-visits" id="d-visits" hidden></span></div>'
    + '<div class="ext-head"><h1><a class="title-link" href="' + pkgHref(e.pkg) + '" data-tip="' + esc(bi('package · ', '软件包 · ') + e.pkg) + '">' + esc(e.name) + '</a></h1>'
    + (e.ver ? '<span class="ver">v' + esc(e.ver) + '</span>' : '') + '</div>'
    + '<p class="ext-tagline">' + esc(desc(e)) + '</p>'
    + '<div class="badge-row">' + heroDimBadgesHTML(e)
    + (e.lifecycle && e.lifecycle !== 'active' ? labeledBadge(t('spec.lifecycle'), '/?lifecycle=' + encodeURIComponent(e.lifecycle), e.lifecycle) : '')
    + (e.kernel ? '<span class="badge lbl"><span class="bk">' + t('spec.kernel') + (LANG === 'zh' ? '：' : ':') + '</span><span class="bv">' + esc(e.kernel) + '</span></span>' : '')
    + (e.vendor ? '<span class="badge lbl"><span class="bk">' + t('spec.vendor') + (LANG === 'zh' ? '：' : ':') + '</span><span class="bv">' + esc(e.vendor) + '</span></span>' : '')
    + '</div>'
    + '<div class="badge-row" id="d-links"></div>'
    + lifecycleNote + providerNote + '</div>'
    + '<aside class="ext-hero-side" id="d-side-main">' + heroSideMainHTML(e, FULLC.get(e.name), e.name) + '</aside>'
    + '</div></header>'
    + extTocHTML()
    + '<section class="section ext-section" id="ext-overview"><h2>' + t('ext.overview') + '</h2><div id="d-overview">' + skel(5) + '</div></section>'
    + '<section class="section ext-section" id="ext-relations"><h2>' + t('ext.relations') + '</h2><div id="d-deps">' + skel(4) + '</div></section>'
    + '<section class="section ext-section" id="ext-packages"><h2>' + t('ext.veravail') + '</h2>'
    + '<div id="d-versions">' + skel(3) + '</div>'
    + '<div id="d-matrix" style="margin-top:12px">' + skel(5) + '</div></section>'
    + '<section class="section ext-section" id="ext-build"><h2>' + t('ext.build') + '</h2><div id="d-build">' + skel(3) + '</div></section>'
    + '<section class="section ext-section" id="ext-install"><h2>' + t('ext.install') + '</h2><p class="section-lede">' + t('ext.installlede') + '</p><div id="d-install">' + skel(6) + '</div></section>'
    + '<section class="section ext-section" id="ext-usage"><h2>' + t('ext.docs') + '</h2><div id="d-doc">' + (e.docbits ? skel(8) : '<div class="docs-missing">' + t('ext.docsnone') + '</div>') + '</div></section>'
    + '</article>';
}

function sqlIdent(name) {
  return /^[a-z_][a-z0-9_$]*$/.test(name) ? name : '"' + String(name).replaceAll('"', '""') + '"';
}

function shellArg(value) {
  return "'" + String(value).replaceAll("'", "'\"'\"'") + "'";
}

const mdCodeHTML = (language, value) => '<div class="md-code compact"><div class="md-codebar"><span>' + language + '</span><button data-copy="' + esc(value) + '">copy</button></div><pre><code>' + highlightCode(value, language) + '</code></pre></div>';

/* pig / apt / dnf install tabs — target-free (the availability matrix already
   answers which PG/OS combos exist): one code box per package manager, one
   line per packaged PostgreSQL major, $v patterns expanded from the catalog.
   Shared verbatim between the extension install guide and the package page. */
function packageTabsHTML(full, fallbackPkg) {
  const word = value => /^[A-Za-z0-9_.+:*$-]+$/.test(String(value)) ? String(value) : shellArg(value);
  const align = rows => { const w = Math.max(...rows.map(([c]) => c.length)); return rows.map(([c, n]) => (c + ';').padEnd(w + 4) + '# ' + n).join('\n'); };
  const majors = list => [...new Set(list || [])].filter(pg => !PGS.length || PGS.includes(pg)).sort((a, b) => b - a);
  const pkgName = full.pkg || fallbackPkg || full.name;
  const pigPGs = majors([...(full.rpm_pg || []), ...(full.deb_pg || [])]);
  const pigText = align([['pig install ' + word(pkgName), bi('for the active PG major', '为当前活跃 PG 大版本安装')]]
    .concat(pigPGs.map(pg => (['pig install ' + word(pkgName) + ' -v ' + pg, bi('install for PG ', '为 PG ') + pg + bi('', ' 安装')]))));
  const patternLines = (mgr, pattern, pgs, none) => pattern && pgs.length
    ? align(pgs.map(pg => [mgr + ' ' + pattern.replaceAll('$v', String(pg)).split(/\s+/).map(word).join(' '), 'PG ' + pg]))
    : '# ' + none;
  const aptText = patternLines('sudo apt install', full.deb_pkg, majors(full.deb_pg), bi('no DEB package is recorded for this extension', '该扩展没有 DEB 软件包记录'));
  const dnfText = patternLines('sudo dnf install', full.rpm_pkg, majors(full.rpm_pg), bi('no RPM package is recorded for this extension', '该扩展没有 RPM 软件包记录'));
  const tabs = [['pig', mdCodeHTML('bash', pigText)], ['apt', mdCodeHTML('bash', aptText)], ['dnf', mdCodeHTML('bash', dnfText)]];
  const heads = tabs.map(([n], i) => '<button role="tab" aria-selected="' + (i === 0) + '" data-itab="' + i + '">' + esc(n) + '</button>').join('');
  const panes = tabs.map(([, body], i) => '<div data-ipane="' + i + '"' + (i ? ' hidden' : '') + '>' + body + '</div>').join('');
  return '<div class="install install-plain"><div class="install-tabs" role="tablist">' + heads + '</div>' + panes + '</div>';
}

function installHTML(e, full) {
  const ident = sqlIdent(e.name);
  const requires = full.requires || [];
  const cascade = requires.length ? ' CASCADE' : '';
  const createSQL = full.need_ddl ? 'CREATE EXTENSION ' + ident + cascade + ';' : '';
  const libs = (full.preload_libs || []).length ? full.preload_libs : ((full.libs || []).length ? full.libs : [e.name]);
  const repoName = String(full.repo || '').toUpperCase();
  const repoCmd = full.contrib || !full.packaged ? '' : (repoName === 'PGDG' ? 'pig repo add pgdg -u' : 'pig repo add pgsql -u');
  const loadConfig = full.need_load ? "shared_preload_libraries = '" + libs.join(', ') + "'" : '';
  const code = mdCodeHTML;

  const upstream = mdSafeURL(full.repo_url || full.doc_url || full.url);
  const packageBody = full.contrib
    ? '<p>' + bi('Included with the matching PostgreSQL contrib/server packages — no separate extension package is required.', '随对应版本的 PostgreSQL contrib / 服务端软件包一并交付，无需安装单独的扩展软件包。') + '</p>'
    : full.packaged
      ? packageTabsHTML(full, e.pkg || e.name)
      : '<p>' + bi('No public binary package is recorded. Follow the upstream build or provider instructions.', '没有公开二进制软件包记录，请遵循上游构建或服务商说明。')
        + (upstream ? ' <a href="' + esc(upstream) + '" target="_blank" rel="noopener">' + esc(upstream.replace(/^https?:\/\//, '')) + ' ↗</a>' : '') + '</p>';
  const repoBody = full.contrib
    ? '<p>' + bi('Bundled with the matching PostgreSQL distribution; no separate repository setup is required.', '随对应 PostgreSQL 发行包交付，无需单独配置扩展仓库。') + '</p>'
    : repoCmd ? '<p>' + bi('Enable and refresh the repository selected by the catalog before installing.', '安装前启用并刷新目录所选的软件仓库。') + '</p>' + code('bash', repoCmd)
      : '<p>' + bi('No public package repository is recorded. Use the build manual and upstream source instead.', '没有公开软件仓库记录，请改用构建手册与上游源码。') + '</p>';
  const loadBody = full.need_load
    ? '<p>' + bi('Append these libraries to the existing setting, then restart PostgreSQL. Do not overwrite other preload entries.', '将下列库追加到现有配置并重启 PostgreSQL；不要覆盖已有预加载项。') + '</p>' + code('postgresql.conf', loadConfig)
    : '<p class="step-done">✓ ' + t('ext.noload') + '</p>';
  const enableBody = createSQL
    ? '<p>' + (requires.length ? bi('CASCADE also creates declared extension dependencies: ', 'CASCADE 会同时创建声明的扩展依赖：') + '<code>' + esc(requires.join(', ')) + '</code>' : bi('Create the SQL objects in each target database.', '在每一个目标数据库中创建扩展 SQL 对象。')) + '</p>' + code('sql', createSQL)
    : '<p class="step-done">✓ ' + t('ext.noddl') + '</p>';
  return '<div class="install-steps">'
    + '<section class="install-step"><h3>' + t('ext.step.repo') + '</h3>' + repoBody + '</section>'
    + '<section class="install-step package-step"><h3>' + t('ext.step.package') + '</h3>' + packageBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.load') + '</h3>' + loadBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.enable') + '</h3>' + enableBody + '</section></div>';
}

/* full availability matrix: os rows × pg columns from /matrix */
function fullMatrixHTML(m, e, options) {
  const opts = options || {};
  const byKey = {};
  for (const c of m.cells) byKey[c.os + '|' + c.pg] = c;
  const famOf = os => os.startsWith('el') ? 'EL' : os.startsWith('d') ? 'Debian' : 'Ubuntu';
  const ths = m.pg.map(pg => '<th><a href="/pg/' + pg + '">PG ' + pg + '</a></th>').join('');
  const counts = { pgdg: 0, pigsty: 0, miss: 0, na: 0 };
  let prevFam = null, rows = '';
  for (const os of m.os) {
    const fam = famOf(os);
    const famStart = fam !== prevFam;
    prevFam = fam;
    const [osname, arch] = os.split('.');
    const archCls = arch === 'aarch64' || arch === 'arm64' ? 'arch-arm' : 'arch-x86';
    const cells = m.pg.map(pg => {
      const c = byKey[os + '|' + pg];
      if (!c || c.state === 'N/A') { counts.na++; return '<td><span class="cellv st-na">N/A</span></td>'; }
      if (c.state === 'MISS') { counts.miss++; return '<td><span class="cellv st-miss">MISS</span></td>'; }
      if (c.state === 'AVAIL') {
        const org = (c.org || '').toLowerCase();
        const cls = org === 'pgdg' ? 'org-pgdg' : org === 'pigsty' ? 'org-pigsty' : 'org-other';
        if (org === 'pgdg') counts.pgdg++; else counts.pigsty++;
        const target = opts.pkg ? ' data-target-pkg="' + esc(opts.pkg) + '"' : ' data-target-ext="' + esc(e.name) + '"';
        return '<td><button class="cellv ' + cls + '"' + target + ' data-target-pg="' + c.pg + '" data-target-os="' + esc(c.os) + '" data-tip="' + esc(c.name || '') + ' · ' + c.count + ' builds">' + esc(String(c.org || '✓').toUpperCase()) + ' ' + esc(c.version || '') + '</button></td>';
      }
      counts.na++;
      return '<td><span class="cellv st-na">N/A</span></td>';
    }).join('');
    // rows alternate x86_64 / aarch64: the arch class zebra-stripes the pair;
    // the full target spelling stays intact ("el8.x86_64"), one type style,
    // with the aarch64 suffix a shade lighter
    rows += '<tr class="' + archCls + (famStart ? ' fam-start' : '') + '"><td class="oslab"><a href="/os/' + encodeURIComponent(os) + '"><b>' + esc(osname) + '</b><span class="oslab-dot">.</span><span class="oslab-arch">' + esc(arch) + '</span></a></td>' + cells + '</tr>';
  }
  const total = m.pg.length * m.os.length;
  const avail = counts.pgdg + counts.pigsty;
  // the extension page carries a summary line pointing at the package page's
  // full artifact listing; the package page is already there
  const pkgLink = '<a href="' + pkgHref(e.pkg) + '#pkg-packages">' + bi('Package: ', '软件包：') + esc(e.pkg) + '</a>';
  const summary = opts.pkg ? '' : '<p class="mx-summary">'
    + bi('<b>' + fmtInt(avail) + '</b> of <b>' + fmtInt(total) + '</b> build targets carry a binary — see the full artifacts at ',
         '共有 <b>' + fmtInt(avail) + '</b> / <b>' + fmtInt(total) + '</b> 个构建组合可用，完整制品信息请参考 ')
    + pkgLink + '</p>';
  const chip = (cls, label, key, n) => '<span><span class="cellv ' + cls + '">' + label + '</span> ' + t(key) + ' <b>' + fmtInt(n) + '</b></span>';
  const legend = '<div class="mx-legend">'
    + chip('org-pgdg', 'PGDG', 'mx.legend.pgdg', counts.pgdg)
    + chip('org-pigsty', 'PIGSTY', 'mx.legend.pigsty', counts.pigsty)
    + chip('st-miss', 'MISS', 'mx.legend.miss', counts.miss)
    + chip('st-na', 'N/A', 'mx.legend.na', counts.na)
    + '</div>';
  return summary + '<div class="matrix-scroll"><table class="fmx"><thead><tr><th class="corner">OS / PG</th>' + ths + '</tr></thead><tbody>'
    + rows + '</tbody></table></div>' + legend;
}

/* package files: per-PG tabs, latest builds first, older collapsible */
function filesHTML(f) {
  if (!f || !f.files || !f.files.length) return '<p class="empty-note">' + t('files.none') + '</p>';
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
        + '<td class="f-os"><a href="/os/' + encodeURIComponent(x.os) + '">' + esc(x.os) + '</a></td>'
        + '<td>' + esc(x.name) + '</td>'
        + '<td>' + esc(x.ver) + '</td>'
        + '<td class="f-org">' + (x.org || x.repo
          ? '<a href="/repo/' + encodeURIComponent(String(x.org || x.repo).toUpperCase()) + '">' + esc(x.org || x.repo) + '</a>' : '—') + '</td>'
        + '<td class="f-size">' + fmtSize(x.size) + '</td>'
        + '<td class="f-sha">' + (x.sha256 ? '<button data-copy="' + esc(x.sha256) + '" data-tip="' + esc(x.sha256) + '">' + esc(x.sha256.slice(0, 10)) + '…</button>' : '—') + '</td>'
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
      + '<th>' + t('files.org') + '</th><th style="text-align:right">' + t('files.size') + '</th><th>' + t('files.sha') + '</th><th>' + t('files.file') + '</th>'
      + '</tr></thead><tbody>' + trs + '</tbody></table></div></div>' + foot + '</div>';
  }).join('');
  return '<div class="file-browser"><div class="files-tabs" role="tablist">' + tabs + '</div>' + panes + '</div>';
}

function boolHTML(value) {
  return value
    ? '<span class="bool yes">✓ ' + bi('yes', '是') + '</span>'
    : '<span class="bool no">— ' + bi('no', '否') + '</span>';
}

function pgBadgesHTML(values) {
  const pgs = (values || []).map(Number).filter(Boolean).sort((a, b) => b - a);
  return pgs.length ? '<span class="pg-badges">' + pgs.map(pg => '<a href="/pg/' + pg + '">PG' + pg + '</a>').join('') + '</span>' : '—';
}

/* PG-major cell for the definitions tables: every active major on one line,
   bare numbers — green when the package supports it, red when it does not. */
function pgMajorsCellHTML(supported) {
  const on = new Set((supported || []).map(Number));
  return '<span class="pgm-cell">' + PGS.map(pg =>
    '<a class="pgm ' + (on.has(pg) ? 'on' : 'off') + '" href="/pg/' + pg + '">' + pg + '</a>').join('') + '</span>';
}

/* dependency cell: multiple entries break one per line */
function depsCellHTML(values, linked) {
  const list = (values || []).map(v => linked && byName.has(v) ? '<a href="' + extHref(v) + '">' + esc(v) + '</a>' : esc(v));
  if (!list.length) return '—';
  return list.join(list.length >= 2 ? '<br>' : ', ');
}

// repository cell: the standard hued badge, navigating to /repo/{V}
const repoCellHTML = repo => {
  if (!repo || repo === 'n/a') return '—';
  const upper = String(repo).toUpperCase();
  return hueTag('/repo/' + encodeURIComponent(upper), upper, repoHue(upper));
};

const buildCellHTML = recipe => recipe == null ? '—'
  : recipe ? '<span class="bool yes">✓ ' + bi('Yes', '支持') + '</span>'
    : '<span class="bool no">— ' + bi('Miss', '暂无') + '</span>';

/* Overview: three equal six-row panels — catalog identity, runtime behavior
   and project resources — with the catalog comment quoted underneath.
   GitHub counters and dates live in the hero side, packaging in the tables. */
function overviewHTML(e, full) {
  return metadataHTML(full)
    + (full.comment ? '<blockquote class="catalog-quote">' + esc(full.comment) + '</blockquote>' : '');
}

function metadataHTML(full) {
  const row = (label, value, mono) => '<div class="meta-row"><dt>' + label + '</dt><dd' + (mono ? ' class="mono"' : '') + '>' + (value || '') + '</dd></div>';
  const panel = (title, body, cls) => '<section class="meta-panel' + (cls ? ' ' + cls : '') + '"><h3>' + title + '</h3>' + body + '</section>';
  const dl = rows => '<dl>' + rows.join('') + '</dl>';
  const catalog = [
    row('ID', esc(full.id), true),
    row(bi('Version', '版本'), esc(full.version || ''), true),
    row(bi('PG Ver', 'PG 版本'), esc(pgRange(full.pg_ver)), true),
    row(bi('Package', '软件包'), '<a href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a>'
      + (full.family_size > 1 ? ' · ' + full.family_size + bi(' ext', ' 扩展') : ''), true),
    row(bi('Lead', '主扩展'), '<a href="' + extHref(full.lead_ext || full.name) + '">' + esc(full.lead_ext || full.name) + '</a>', true),
    row(bi('Kind', '形态'), full.kind ? '<span data-tip="' + esc(t('type.' + full.kind)) + '">' + esc(full.kind) + '</span>' : '', true)
  ];
  // one attribute per row, value centered, a plain-words explanation on hover
  const attrRow = (en, zh, tipEN, tipZH, value) =>
    '<div class="meta-row attr-row" data-tip="' + esc(bi(tipEN, tipZH)) + '"><dt>' + bi(en, zh) + '</dt><dd>' + boolHTML(value) + '</dd></div>';
  const runtime = [
    attrRow('Has Binary', '带有二进制', 'Ships executable command-line tools', '该扩展附带可执行的命令行工具', full.has_bin),
    attrRow('Has Library', '带有动态库', 'Ships a shared library (.so)', '该扩展附带动态链接库（.so 文件）', full.has_lib),
    attrRow('Need Preload', '需动态加载', 'Must be listed in shared_preload_libraries; changing that requires a restart', '需加入 shared_preload_libraries 预加载，修改后需重启数据库', full.need_load),
    attrRow('Need CREATE', '用 DDL 创建', 'Enabled per database with CREATE EXTENSION', '需执行 CREATE EXTENSION 语句创建后方可使用', full.need_ddl),
    attrRow('Need DBSU', '需超级用户', 'Only a database superuser can install it', '需要数据库超级用户权限才能安装', !full.trusted),
    attrRow('Relocatable', '模式可修改', 'Its objects can be installed into a schema of your choice', '扩展对象可以安装到指定的模式（schema）中', full.relocatable)
  ];
  // fixed six link slots: available URLs fill from the top, blanks keep the
  // three panels the same 6-row height; hosts never wrap — ellipsis + mark
  const links = full.doc_links || {};
  const resources = [
    [bi('Homepage', '主页'), full.home_url || links.home_url],
    [bi('Repository', '仓库'), full.repo_url || links.repo_url],
    [bi('License', '许可证'), full.license_url || links.license_url],
    ['Control', full.control_url || links.control_url],
    [bi('Docs', '文档'), full.doc_url || links.doc_url],
    [bi('Author', '作者'), full.author_url || links.author_url],
    ['PGXN', full.pgxn_url || links.pgxn_url],
    ['Cargo', full.cargo_url || links.cargo_url]
  ].map(([label, url]) => [label, mdSafeURL(url)]).filter(([, url]) => url).slice(0, 6);
  while (resources.length < 6) resources.push(null);
  const resourcesHTML = dl(resources.map(item => item
    ? row(item[0], '<a class="res-link" href="' + esc(item[1]) + '" target="_blank" rel="noopener"><span>'
      + esc(item[1].replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</span>' + ICON_OUT + '</a>', true)
    : '<div class="meta-row"><dt>&nbsp;</dt><dd></dd></div>'));
  // the full universe record, pretty-printed and one click away from the
  // clipboard — the whole truth for anyone who wants to take it home
  const json = JSON.stringify(full, null, 2);
  const extra = '<details class="extra-meta"><summary>' + t('spec.extra') + '</summary>'
    + '<div class="extra-body"><button class="chip" data-copy="' + esc(json) + '">' + bi('copy JSON', '复制 JSON') + '</button>'
    + '<pre>' + esc(json) + '</pre></div></details>';
  return '<div class="metadata-grid tri">' + panel(t('ext.catalog'), dl(catalog)) + panel(t('ext.runtime'), dl(runtime))
    + panel(t('ext.projectlinks'), resourcesHTML, 'links') + '</div>' + extra;
}

function packageVersionsHTML(full) {
  const row = (type, repo, version, pgs, pattern, dependencies, recipe, linked) => '<tr><td><b>' + type + '</b></td><td>' + repoCellHTML(repo) + '</td>'
    + '<td class="mono">' + esc(version || '—') + '</td><td>' + pgMajorsCellHTML(pgs) + '</td><td class="mono">' + esc(pattern || '—') + '</td>'
    + '<td class="mono r-deps">' + depsCellHTML(dependencies, linked) + '</td><td>' + buildCellHTML(recipe) + '</td></tr>';
  let rows = row('EXT', full.repo, full.version, full.pg_ver, full.pkg, full.requires, null, true);
  if (!full.contrib && (full.rpm_pkg || full.rpm_ver || (full.rpm_pg || []).length)) rows += row('RPM', full.rpm_repo, full.rpm_ver, full.rpm_pg, full.rpm_pkg, full.rpm_deps, full.rpm_build, false);
  if (!full.contrib && (full.deb_pkg || full.deb_ver || (full.deb_pg || []).length)) rows += row('DEB', full.deb_repo, full.deb_ver, full.deb_pg, full.deb_pkg, full.deb_deps, full.deb_build, false);
  return '<div class="version-table"><div class="rows-scroll"><table><thead><tr><th>' + bi('type', '类型') + '</th><th>' + bi('repo', '仓库')
    + '</th><th>' + bi('version', '版本') + '</th><th>' + bi('pg major versions', 'PG 大版本') + '</th><th>' + bi('package pattern', '包名模式') + '</th><th>' + bi('dependencies', '依赖')
    + '</th><th>' + bi('build', '构建') + '</th></tr></thead><tbody>' + rows + '</tbody></table></div></div>';
}

function sourceArchiveURL(source) {
  if (!source) return '';
  if (/^https?:\/\//i.test(source)) return source;
  const path = String(source).replace(/^\/+/, '').split('/').map(encodeURIComponent).join('/');
  return (LANG === 'zh' ? 'https://repo.pigsty.cc/ext/src/' : 'https://repo.pigsty.io/ext/src/') + path;
}

function buildHTML(full) {
  if (full.contrib) return '<div class="build-note contrib"><b>PostgreSQL contrib</b><p>' + t('ext.contribbuild') + '</p>' + pgBadgesHTML(full.pg_ver) + '</div>';
  const types = [full.rpm_build ? 'RPM' : '', full.deb_build ? 'DEB' : ''].filter(Boolean);
  const archive = sourceArchiveURL(full.tarball || full.source);
  const upstream = mdSafeURL(full.repo_url || full.url);
  const docs = mdSafeURL(full.doc_url);
  const links = [];
  if (upstream) links.push('<a href="' + esc(upstream) + '" target="_blank" rel="noopener"><span>' + bi('Upstream source', '上游源码') + '</span><b>' + esc(upstream.replace(/^https?:\/\//, '')) + ' ↗</b></a>');
  if (archive) links.push('<a href="' + esc(archive) + '" target="_blank" rel="noopener"><span>' + t('ext.sourcearchive') + '</span><b>' + esc((full.tarball || full.source).split('/').pop()) + ' ↓</b></a>');
  if (docs) links.push('<a href="' + esc(docs) + '" target="_blank" rel="noopener"><span>' + bi('Upstream build docs', '上游构建文档') + '</span><b>' + esc(docs.replace(/^https?:\/\//, '')) + ' ↗</b></a>');
  let guide = '';
  if (types.length) {
    const command = 'pig build pkg ' + full.pkg;
    guide = '<div class="build-recipe"><p>' + t('ext.buildrecipe', { types: '<b>' + types.join(' / ') + '</b>' }) + '</p>'
      + '<div class="build-targets"><span>RPM ' + pgBadgesHTML(full.rpm_pg) + '</span><span>DEB ' + pgBadgesHTML(full.deb_pg) + '</span>'
      + (full.pgrx_ver ? '<span><code>pgrx ' + esc(full.pgrx_ver) + '</code></span>' : '') + '</div>'
      + '<div class="md-code"><div class="md-codebar"><span>bash</span><button data-copy="' + esc(command) + '">copy</button></div><pre><code>' + highlightCode(command, 'bash') + '</code></pre></div></div>';
  } else {
    guide = '<div class="build-note"><b>' + bi('Upstream build', '上游构建') + '</b><p>' + t('ext.buildnone') + '</p>'
      + (full.pgrx_ver ? '<span class="badge">pgrx · ' + esc(full.pgrx_ver) + '</span>' : '') + '</div>';
  }
  return (links.length ? '<div class="source-cards">' + links.join('') + '</div>' : '') + guide;
}

/* Dependency analysis: a mind-map style panel — upstream dependencies branch
   into the extension from the left, downstream dependants branch out to the
   right, with connector lines meeting at the center box. Empty sides state
   the fact in plain words. */
function depGraphHTML(full) {
  const up = full.requires || [];
  const down = full.required_by || full.require_by || [];
  const box = name => {
    const x = byName.get(name);
    return '<a class="dep-item" ' + (x ? catVar(x.cat) + ' data-hover-ext="' + esc(x.name) + '" ' : '') + 'href="' + extHref(name) + '">' + esc(name) + '</a>';
  };
  // The captions sit at the far edges, absolutely positioned so they never
  // shift the columns — item centers and the center box must share one axis
  // for the connector stubs and rails to join into continuous lines.
  const col = (names, kind) => {
    if (!names.length) return '<div class="dep-none">' + t(kind === 'up' ? 'dep.noup' : 'dep.nodown') + '</div>';
    const items = names.slice(0, 12).map(box).join('')
      + (names.length > 12 ? '<span class="dep-item dep-more">+' + (names.length - 12) + '</span>' : '');
    return '<div class="dep-col dep-' + kind + '">' + items + '</div>';
  };
  const e = byName.get(full.name);
  return '<h3 class="section-subhead">' + t('ext.deps') + '</h3>'
    + '<div class="dep-graph' + (up.length ? ' has-up' : '') + (down.length ? ' has-down' : '') + '">'
    + '<span class="dep-cap dep-cap-up">' + bi('Upstream', '上游 Upstream') + '</span>'
    + '<span class="dep-cap dep-cap-down">' + bi('Downstream', '下游 Downstream') + '</span>'
    + col(up, 'up')
    + '<div class="dep-mid"><span class="dep-self" ' + (e ? catVar(e.cat) : '') + '>' + esc(full.name) + '</span></div>'
    + col(down, 'down') + '</div>';
}

function seeAlsoHTML(full) {
  const names = full.see_also || [];
  if (!names.length) return '';
  const cards = names.map(name => {
    const e = byName.get(name);
    return '<a href="' + extHref(name) + '" ' + (e ? catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '"' : '') + '><code>' + esc(name) + '</code>'
      + '<span>' + esc(e ? desc(e) : bi('catalog reference', '目录引用')) + '</span>'
      + (e ? '<small>' + esc(e.cat) + ' · ' + esc(e.kind || '') + '</small>' : '') + '</a>';
  }).join('');
  return '<h3 class="section-subhead">See Also</h3><div class="relation-grid">' + cards + '</div>';
}

/* Relationship: the same-package family table, the dependency analysis
   graph, then curated see-also entries. */
function relationshipHTML(full, members, lead, fulls) {
  return '<h3 class="section-subhead sp-head"><a class="sp-pkg" href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a></h3>'
    + pkgFamilyTableHTML(members, lead, fulls)
    + depGraphHTML(full)
    + seeAlsoHTML(full);
}

/* hydration: fetch full record, matrix, files, doc — fill sections as they land */
let hydSeq = 0;
function fill(id, html) { const el = document.getElementById(id); if (el) el.innerHTML = html; }

async function hydrateExt(name) {
  const tok = ++hydSeq;
  const e = byName.get(name);
  if (!e) return;
  const enc = encodeURIComponent(name);
  const fullP = (async () => {
    let full = FULLC.get(name);
    if (!full) { full = await j('/api/v1/ext/' + enc); FULLC.set(name, full); }
    return full;
  })();
  const matrixP = (async () => {
    let matrix = MXC.get(name);
    if (!matrix) { matrix = await j('/api/v1/ext/' + enc + '/matrix'); MXC.set(name, matrix); }
    return matrix;
  })();

  fullP.then(full => {
    if (tok !== hydSeq) return;
    fill('d-overview', overviewHTML(e, full));
    fill('d-versions', packageVersionsHTML(full));
    fill('d-build', buildHTML(full));
    fill('d-side-main', heroSideMainHTML(e, full, e.name));
    fill('d-links', heroLinkBadgesHTML(full));
    fill('d-install', installHTML(e, full));
    // Relationship renders immediately from cached records, then refines the
    // family table once every member's full record (schema, requires) landed.
    const members = byPkg.get(e.pkg) || [e];
    const lead = byName.get(members[0].lead) || members[0];
    const cached = new Map(members.map(m => [m.name, FULLC.get(m.name)]).filter(pair => pair[1]));
    fill('d-deps', relationshipHTML(full, members, lead, cached));
    if (members.length > 1 && members.length <= 40) {
      Promise.all(members.map(async m => {
        let f = FULLC.get(m.name);
        if (!f) {
          try { f = await j('/api/v1/ext/' + encodeURIComponent(m.name)); FULLC.set(m.name, f); } catch (err) { f = null; }
        }
        return [m.name, f];
      })).then(pairs => {
        if (tok !== hydSeq) return;
        fill('d-deps', relationshipHTML(full, members, lead, new Map(pairs.filter(pair => pair[1]))));
      });
    }
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['d-overview', 'd-deps', 'd-versions', 'd-build', 'd-install']) fill(id, hydrateErr(err));
  });

  matrixP.then(matrix => {
    if (tok !== hydSeq) return;
    fill('d-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, e) : '<p class="empty-note">' + t('files.none') + '</p>');
  }).catch(err => { if (tok === hydSeq) fill('d-matrix', hydrateErr(err)); });

  // Page hit counter: fire-and-forget increment, surfaced as a quiet
  // eye + number under the hero dates.
  fetch('/api/v1/ext/' + enc + '/visit', { method: 'POST' })
    .then(res => res.ok ? res.json() : null)
    .then(v => {
      if (!v || tok !== hydSeq) return;
      const chip = document.getElementById('d-visits');
      if (!chip) return;
      chip.hidden = false;
      chip.innerHTML = ICON_EYE + fmtInt(v.visits);
      chip.dataset.tip = fmtInt(v.visits) + ' ' + t('ext.visits');
    }).catch(() => {});

  if (e.docbits) {
    (async () => {
      const wantZh = LANG === 'zh' && (e.docbits & 2);
      const lang = wantZh ? 'zh' : 'en';
      const key = name + '|' + lang;
      try {
        let d = DOCC.get(key);
        if (!d) { d = await j('/api/v1/ext/' + enc + '/doc?lang=' + lang); DOCC.set(key, d); }
        if (tok !== hydSeq) return;
        const note = (LANG === 'zh' && !wantZh) ? '<p class="empty-note" style="margin-bottom:12px">' + t('ext.doconlyen') + '</p>' : '';
        const content = String(d.content || '').replace(/^\s*(?:#{1,2}\s+(?:usage|用法)\s*\n+|(?:usage|用法)\s*\n[-=]{3,}\s*\n+)/i, '');
        const rendered = renderMD(content, { usage: true });
        fill('d-doc', note + '<article class="prose usage-prose">' + rendered.html + '</article>');
      } catch (err) { if (tok === hydSeq) fill('d-doc', hydrateErr(err)); }
    })();
  }
}

function notFoundHTML(name) {
  return '<div class="page wrap"><div class="section">'
    + '<pre class="mono" style="font-size:14px;line-height:1.8;color:var(--bad)">ERROR:  extension "' + esc(name) + '" does not exist\n'
    + '<span style="color:var(--ink-3)">' + esc(t('notfound.hint')) + '</span></pre>'
    + '<p style="margin-top:18px"><a href="/">' + t('notfound.back') + '</a></p>'
    + '</div></div>';
}

/* ---------------- view: package / upstream project ---------------- */
function pkgTocHTML() {
  return pageTocHTML([
    ['pkg-family', t('ext.family')],
    ['pkg-packages', t('ext.pkgs')],
    ['pkg-downloads', t('ext.downloads')],
    ['pkg-build', t('ext.build')],
    ['pkg-install', t('ext.install')]
  ]);
}

/* authentic octicons for the GitHub card: star / repo-forked / eye */
const ICON_STAR = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 .25a.75.75 0 0 1 .673.418l1.882 3.815 4.21.612a.75.75 0 0 1 .416 1.279l-3.046 2.97.719 4.192a.751.751 0 0 1-1.088.791L8 12.347l-3.766 1.98a.75.75 0 0 1-1.088-.79l.72-4.194L.818 6.374a.75.75 0 0 1 .416-1.28l4.21-.611L7.327.668A.75.75 0 0 1 8 .25Zm0 2.445L6.615 5.5a.75.75 0 0 1-.564.41l-3.097.45 2.24 2.184a.75.75 0 0 1 .216.664l-.528 3.084 2.769-1.456a.75.75 0 0 1 .698 0l2.77 1.456-.53-3.084a.75.75 0 0 1 .216-.664l2.24-2.183-3.096-.45a.75.75 0 0 1-.564-.41L8 2.694Z"/></svg>';
const ICON_FORK = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M5 5.372v.878c0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75v-.878a2.25 2.25 0 1 1 1.5 0v.878a2.25 2.25 0 0 1-2.25 2.25h-1.5v2.128a2.251 2.251 0 1 1-1.5 0V8.5h-1.5A2.25 2.25 0 0 1 3.5 6.25v-.878a2.25 2.25 0 1 1 1.5 0ZM5 3.25a.75.75 0 1 0-1.5 0 .75.75 0 0 0 1.5 0Zm6.75.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Zm-3 8.75a.75.75 0 1 0-1.5 0 .75.75 0 0 0 1.5 0Z"/></svg>';
const ICON_EYE = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 2c1.981 0 3.671.992 4.933 2.078 1.27 1.091 2.187 2.345 2.637 3.023a1.62 1.62 0 0 1 0 1.798c-.45.678-1.367 1.932-2.637 3.023C11.67 13.008 9.981 14 8 14c-1.981 0-3.671-.992-4.933-2.078C1.797 10.831.88 9.577.43 8.899a1.62 1.62 0 0 1 0-1.798c.45-.678 1.367-1.932 2.637-3.023C4.33 2.992 6.019 2 8 2ZM1.679 7.932a.12.12 0 0 0 0 .136c.411.622 1.241 1.75 2.366 2.717C5.176 11.758 6.527 12.5 8 12.5c1.473 0 2.825-.742 3.955-1.715 1.124-.967 1.954-2.096 2.366-2.717a.12.12 0 0 0 0-.136c-.412-.621-1.242-1.75-2.366-2.717C10.824 4.242 9.473 3.5 8 3.5c-1.473 0-2.825.742-3.955 1.715-1.124.967-1.954 2.096-2.366 2.717ZM8 10a2 2 0 1 1-.001-3.999A2 2 0 0 1 8 10Z"/></svg>';

/* stroke glyphs + the outbound mark (a square with an escaping arrow) */
const _stroke = (d, size) => '<svg width="' + (size || 13) + '" height="' + (size || 13) + '" viewBox="0 0 16 16" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="' + d + '"/></svg>';
const ICON_OUT = '<svg class="ext-mark" width="9" height="9" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M6 2.8H3.4A1.4 1.4 0 0 0 2 4.2v6.4A1.4 1.4 0 0 0 3.4 12h6.4a1.4 1.4 0 0 0 1.4-1.4V8M8.2 2h3.8v3.8M11.7 2.3 6.6 7.4"/></svg>';
const ICON_LINK = _stroke('M6.6 9.4 9.4 6.6M5 7.5 3.4 9.1a2.6 2.6 0 0 0 3.7 3.7l1.6-1.6M11 8.5l1.6-1.6a2.6 2.6 0 0 0-3.7-3.7L7.3 4.8');
const GH_SVG_LG = GH_SVG.replace('width="13" height="13"', 'width="26" height="26"');
const ICON_LINK_LG = ICON_LINK.replace('width="13" height="13"', 'width="24" height="24"');

/* labeled hero badge: a quiet gray "Label:" in front of the value —
   Category: STAT · License: PostgreSQL · Language: C */
const labeledBadge = (label, href, value, hue) =>
  '<a class="badge lbl' + (hue ? ' hued' : '') + '" ' + (hue ? hueVar(hue) + ' ' : '') + 'href="' + esc(href) + '">'
  + '<span class="bk">' + label + (LANG === 'zh' ? '：' : ':') + '</span><span class="bv">' + esc(value) + '</span></a>';

/* the shared dimension badge row: category / license / language — the
   category badge carries only its short code, the full name as a tip */
function heroDimBadgesHTML(e) {
  return '<a class="badge cat lbl" href="' + catHref(e.cat) + '" ' + catVar(e.cat) + ' data-tip="' + esc(catName(e.cat)) + '"><span class="bk">' + t('spec.category')
    + (LANG === 'zh' ? '：' : ':') + '</span><span class="bv">' + e.cat + '</span></a>'
    + (e.license !== 'Unknown' ? labeledBadge(t('spec.license'), '/license/' + encodeURIComponent(e.license), e.license, licenseHue(e.license)) : '')
    + (e.lang ? labeledBadge(t('spec.language'), '/lang/' + encodeURIComponent(e.lang), e.lang, langHue(e.lang)) : '');
}

/* project card in the hero side: a large mark on the left (octocat for
   GitHub, a link glyph otherwise), the repo path over its counters on the
   right, and a boxed-arrow outbound mark in the corner. */
function ghCardHTML(e, full, fallbackName) {
  const url = mdSafeURL((full && full.repo_url) || e.repoUrl || (full && full.home_url) || e.url);
  if (!url) return '';
  const isGH = /github\.com/i.test(url);
  const name = isGH
    ? url.replace(/^https?:\/\/(www\.)?github\.com\//i, '').replace(/\/+$/, '')
    : (fallbackName || e.name);
  const stat = (icon, value, label) => '<span data-tip="' + label + '">' + icon + '<b>' + (value == null ? '—' : fmtNum(value)) + '</b></span>';
  const sub = isGH
    ? '<span class="gh-stats">'
      + stat(ICON_STAR, full ? full.stars : e.stars, 'Star') + stat(ICON_FORK, full && full.forks, 'Fork')
      + stat(ICON_EYE, full && full.watchers, 'Watch') + '</span>'
    : '<span class="gh-stats gh-url">' + esc(url.replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</span>';
  return '<a class="gh-card" href="' + esc(url) + '" target="_blank" rel="noopener">'
    + '<span class="gh-logo">' + (isGH ? GH_SVG_LG : ICON_LINK_LG) + '</span>'
    + '<span class="gh-main"><b class="gh-name">' + esc(name) + '</b>' + sub + '</span>'
    + '<i class="gh-out">' + ICON_OUT + '</i></a>';
}

/* the four key project dates in a 2 × 2 grid, shown only when recorded */
function heroDatesHTML(full) {
  const rows = [
    [bi('Repo Created', '仓库创建'), full.repo_created_at],
    [bi('Last Release', '最近发布'), full.last_release],
    [bi('Last Commit', '最近提交'), full.last_commit],
    [bi('Last Modified', '最近更新'), full.last_active]
  ].filter(row => row[1]);
  if (!rows.length) return '';
  return '<dl class="hero-dates">' + rows.map(([k, v]) => '<div><dt>' + k + '</dt><dd>' + esc(String(v).slice(0, 10)) + '</dd></div>').join('') + '</dl>';
}

function heroSideMainHTML(e, full, fallbackName) {
  return ghCardHTML(e, full, fallbackName) + (full ? heroDatesHTML(full) : '');
}

/* deduped outbound URL badges under the hero, each with its glyph and a
   small outbound mark: Home · Repo · Docs · PGXN · License */
function heroLinkBadgesHTML(full) {
  return (full.tags || []).slice(0, 10).map(tag =>
    '<a class="tg hero-tag" href="/?tag=' + encodeURIComponent(tag) + '">' + esc(tag) + '</a>').join('');
}

// Package definitions: RPM and DEB rows only. Contrib families ship with the
// server packages; families with neither format have no prebuilt binaries.
function pkgDefinitionsHTML(full) {
  const row = (type, repo, version, pgs, pattern, dependencies, recipe) => '<tr><td><b>' + type + '</b></td><td>' + repoCellHTML(repo) + '</td>'
    + '<td class="mono">' + esc(version || '—') + '</td><td>' + pgMajorsCellHTML(pgs) + '</td><td class="mono">' + esc(pattern || '—') + '</td>'
    + '<td class="mono r-deps">' + depsCellHTML(dependencies, false) + '</td><td>' + buildCellHTML(recipe) + '</td></tr>';
  if (full.contrib) return '<p class="empty-note">' + t('ext.contribbuild') + '</p>';
  let rows = '';
  if (full.rpm_pkg || full.rpm_ver || (full.rpm_pg || []).length) rows += row('RPM', full.rpm_repo, full.rpm_ver, full.rpm_pg, full.rpm_pkg, full.rpm_deps, full.rpm_build);
  if (full.deb_pkg || full.deb_ver || (full.deb_pg || []).length) rows += row('DEB', full.deb_repo, full.deb_ver, full.deb_pg, full.deb_pkg, full.deb_deps, full.deb_build);
  if (!rows) return '<p class="empty-note">' + bi('No prebuilt binary package is available for this extension yet.', '该扩展尚未提供已构建的二进制软件包。') + '</p>';
  return '<div class="version-table pkg-definitions"><div class="rows-scroll"><table><thead><tr><th>' + bi('type', '类型') + '</th><th>' + bi('repo', '仓库')
    + '</th><th>' + bi('version', '版本') + '</th><th>' + bi('pg major versions', 'PG 大版本') + '</th><th>' + bi('package / source', '包名 / 源码') + '</th><th>' + bi('system dependencies', '系统依赖')
    + '</th><th>' + bi('build', '构建') + '</th></tr></thead><tbody>' + rows + '</tbody></table></div></div>';
}

/* Extension family lives right under the hero as a table: one row per
   delivered extension — name / version / description / attribute chips /
   schema / requires. Schema and requires hydrate from each member's full
   record; every dependency that exists in the catalog navigates to its page. */
function pkgFamilyTableHTML(members, lead, fulls) {
  const ordered = [lead, ...members.filter(m => m.name !== lead.name).sort((a, b) => a.name.localeCompare(b.name))];
  const rows = ordered.map(e => {
    const full = (fulls && fulls.get(e.name)) || {};
    const isLead = ordered.length > 1 && e.name === lead.name;
    return '<tr ' + catVar(e.cat) + (isLead ? ' class="lead-row"' : '') + ' data-hover-ext="' + esc(e.name) + '">'
      + '<td class="r-name"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a></td>'
      + '<td class="r-mono r-ver">' + esc(e.ver || '') + '</td>'
      + '<td class="r-desc">' + esc(desc(e)) + '</td>'
      + '<td class="r-attr">' + attrChips(e) + '</td>'
      + '<td class="r-mono">' + (full.schemas || []).map(esc).join((full.schemas || []).length >= 2 ? '<br>' : ', ') + '</td>'
      + '<td class="r-mono r-req">' + ((full.requires || []).length ? depsCellHTML(full.requires, true) : '') + '</td></tr>';
  }).join('');
  return '<div class="rows"><div class="rows-scroll"><table class="ext-table family-table"><thead><tr>'
    + '<th>' + t('rows.name') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th>'
    + '<th>' + bi('attribute', '属性') + '</th><th>' + bi('schemas', '模式') + '</th><th>' + bi('requires', '依赖') + '</th>'
    + '</tr></thead><tbody>' + rows + '</tbody></table></div></div>';
}

/* Package install: the same pig / apt / dnf command tabs as the extension
   page, plus a pointer to the lead extension's full install guide. Contrib
   and source-only packages state the situation instead of rendering an empty
   command box. */
function packageInstallHTML(pkg, full) {
  const guide = '<p class="install-more">' + t('ext.installmore') + ' <a href="' + extHref(full.name) + '#ext-install">'
    + esc(full.name) + ' · ' + t('ext.install') + ' →</a></p>';
  if (full.contrib) {
    return '<p class="empty-note">' + bi('Included with the matching PostgreSQL contrib/server packages — no separate extension package is required.',
      '随对应版本的 PostgreSQL contrib / 服务端软件包一并交付，无需安装单独的扩展软件包。') + '</p>' + guide;
  }
  if (!full.packaged) {
    const upstream = mdSafeURL(full.repo_url || full.doc_url || full.url);
    return '<p class="empty-note">' + bi('No public binary package is recorded. Follow the upstream build or provider instructions.',
      '没有公开二进制软件包记录，请遵循上游构建或服务商说明。')
      + (upstream ? ' <a href="' + esc(upstream) + '" target="_blank" rel="noopener">' + esc(upstream.replace(/^https?:\/\//, '')) + ' ↗</a>' : '') + '</p>' + guide;
  }
  return packageTabsHTML(full, pkg) + guide;
}

function pkgHTML(pkg) {
  const members = byPkg.get(pkg);
  if (!members || !members.length) return notFoundHTML(pkg);
  const lead = byName.get(members[0].lead) || members[0];
  const packaged = members.some(e => e.avail);
  return '<article class="page wrap manual-page pkg-page">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span><a href="' + dimHref('package') + '">' + bi('packages', '软件包') + '</a><span class="sep">/</span><span class="here">' + esc(pkg) + '</span></nav>'
    + '<header class="detail-hero pkg-detail-hero"><div class="ext-hero-grid"><div class="ext-hero-main">'
    + '<div class="detail-kicker-row"><span class="detail-kicker">PACKAGE</span>'
    + (packaged ? '<span class="state-chip ok">● ' + t('state.avail') + '</span>' : '<span class="state-chip">○ ' + t('state.na') + '</span>')
    + '<span class="hero-visits" id="p-visits" hidden></span></div>'
    + '<div class="ext-head"><h1><a class="title-link" href="' + extHref(lead.name) + '" data-tip="' + esc(bi('lead extension · ', '主扩展 · ') + lead.name) + '">' + esc(pkg) + '</a></h1>'
    + (lead.ver ? '<span class="ver">v' + esc(lead.ver) + '</span>' : '') + '</div>'
    + '<p class="ext-tagline">' + esc(desc(lead)) + '</p>'
    + '<div class="badge-row">' + heroDimBadgesHTML(lead) + '</div>'
    + '<div class="badge-row" id="p-links"></div></div>'
    + '<aside class="ext-hero-side pkg-hero-side" id="p-side-main">' + heroSideMainHTML(lead, FULLC.get(lead.name), pkg) + '</aside>'
    + '</div></header>'
    + pkgTocHTML()
    + '<section class="section pkg-section" id="pkg-family"><h2>' + t('ext.family') + '</h2><p class="section-lede">'
    + members.length + bi(' extension definitions delivered by this package.', ' 个由该软件包交付的扩展定义。') + '</p><div id="p-family">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-packages"><h2>' + t('ext.pkgs') + '</h2>'
    + '<div id="p-definitions">' + skel(3) + '</div><div id="p-matrix" style="margin-top:12px">' + skel(5) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-downloads"><h2>' + t('ext.downloads') + '</h2><div id="p-files">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-build"><h2>' + t('ext.build') + '</h2><div id="p-build">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-install"><h2>' + t('ext.install') + '</h2><div id="p-install">' + skel(4) + '</div></section>'
    + '</article>';
}

async function hydratePkg(pkg) {
  const members = byPkg.get(pkg);
  if (!members || !members.length) return;
  const lead = byName.get(members[0].lead) || members[0];
  const tok = ++hydSeq;
  const enc = encodeURIComponent(lead.name);
  const fullP = (async () => {
    let full = FULLC.get(lead.name);
    if (!full) { full = await j('/api/v1/ext/' + enc); FULLC.set(lead.name, full); }
    return full;
  })();
  const matrixP = (async () => {
    let matrix = MXC.get(lead.name);
    if (!matrix) { matrix = await j('/api/v1/ext/' + enc + '/matrix'); MXC.set(lead.name, matrix); }
    return matrix;
  })();
  // A package profile visit counts against its lead extension — the package
  // page maintains no counter of its own; the same quiet eye shows the total.
  fetch('/api/v1/ext/' + enc + '/visit', { method: 'POST' })
    .then(res => res.ok ? res.json() : null)
    .then(v => {
      if (!v || tok !== hydSeq) return;
      const chip = document.getElementById('p-visits');
      if (!chip) return;
      chip.hidden = false;
      chip.innerHTML = ICON_EYE + fmtInt(v.visits);
      chip.dataset.tip = fmtInt(v.visits) + ' ' + t('ext.visits');
    }).catch(() => {});

  // the family table needs every member's full record for schema + requires
  Promise.all(members.map(async m => {
    if (m.name === lead.name) return [m.name, await fullP.catch(() => null)];
    let f = FULLC.get(m.name);
    if (!f) {
      try { f = await j('/api/v1/ext/' + encodeURIComponent(m.name)); FULLC.set(m.name, f); } catch (err) { f = null; }
    }
    return [m.name, f];
  })).then(pairs => {
    if (tok !== hydSeq) return;
    fill('p-family', pkgFamilyTableHTML(members, lead, new Map(pairs.filter(([, f]) => f))));
  });
  fullP.then(full => {
    if (tok !== hydSeq) return;
    fill('p-side-main', heroSideMainHTML(lead, full, pkg));
    fill('p-links', heroLinkBadgesHTML(full));
    fill('p-definitions', pkgDefinitionsHTML(full));
    fill('p-build', buildHTML(full));
    fill('p-install', packageInstallHTML(pkg, full));
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['p-definitions', 'p-build', 'p-install']) fill(id, hydrateErr(err));
  });
  matrixP.then(matrix => {
    if (tok !== hydSeq) return;
    fill('p-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, lead, { pkg }) : '<p class="empty-note">' + t('files.none') + '</p>');
  }).catch(err => { if (tok === hydSeq) fill('p-matrix', hydrateErr(err)); });
  try {
    let files = FILEC.get(lead.name);
    if (!files) { files = await j('/api/v1/ext/' + enc + '/files'); FILEC.set(lead.name, files); }
    if (tok === hydSeq) fill('p-files', filesHTML(files));
  } catch (err) { if (tok === hydSeq) fill('p-files', hydrateErr(err)); }
}

/* ---------------- view: dimension navigation pages ----------------
   One landing page per value of the primary dimensions — /cate/{CODE},
   /lang/{value}, /repo/{value}, /license/{value}. Each page carries a
   horizontal nav across every value of its dimension, featured cards
   (identical to the home wall) and the full table. */
const NAV_DIMS = {
  category: {
    prefix: '/cate/', param: 'cate', dim: 'category',
    canon: v => v.toUpperCase(),
    match: (e, v) => e.cat === v,
    title: v => catName(v),
    lede: v => catDesc(v)
  },
  lang: {
    prefix: '/lang/', param: 'lang', dim: 'lang',
    canon: v => v,
    match: (e, v) => e.lang.toLowerCase() === v.toLowerCase(),
    title: v => v,
    lede: () => bi('Extensions implemented in this language.', '以该语言实现的扩展。')
  },
  repo: {
    prefix: '/repo/', param: 'repo', dim: 'repo',
    canon: v => v.toUpperCase(),
    match: (e, v) => e.repo === v,
    title: v => v,
    lede: () => bi('Extensions delivered by this package source.', '由该软件源打包分发的扩展。')
  },
  license: {
    prefix: '/license/', param: 'license', dim: 'license',
    canon: v => v,
    match: (e, v) => e.license.toLowerCase() === v.toLowerCase(),
    title: v => v,
    lede: () => bi('Extensions released under this license.', '以该许可证发布的扩展。')
  },
  pg: {
    prefix: '/pg/', param: 'pg', dim: 'pg',
    canon: v => String(Number.parseInt(v, 10) || v),
    match: (e, v) => e.pg.includes(Number(v)),
    title: v => 'PostgreSQL ' + v,
    lede: v => bi('Extensions declaring support for PostgreSQL ' + v + ' — a per-major digest of the build matrix.',
                  '声明支持 PostgreSQL ' + v + ' 的扩展——按大版本切片的构建矩阵摘要。')
  },
  os: {
    prefix: '/os/', param: 'os', dim: 'os',
    canon: v => v.toLowerCase(),
    match: (e, v) => targetAvailable(e, 0, v),
    title: v => v,
    lede: v => bi('Extensions installable on this Linux target — a per-platform digest of the build matrix.',
                  '可在该 Linux 目标上安装的扩展——按平台切片的构建矩阵摘要。')
  }
};

function navValueHref(dim, value) {
  const m = { category: '/cate/', lang: '/lang/', repo: '/repo/', license: '/license/', pg: '/pg/', os: '/os/' }[dim];
  return m ? m + encodeURIComponent(value) : '/?' + DIMS[dim].key + '=' + encodeURIComponent(value);
}

function extTableHTML(members) {
  return '<div class="rows"><div class="rows-scroll"><table class="ext-table"><thead><tr>'
    + '<th>' + t('rows.name') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th><th>' + t('rows.cat') + '</th>'
    + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
    + '</tr></thead><tbody>' + members.map(rowHTML).join('') + '</tbody></table></div></div>';
}

function navPageHTML(kind, raw) {
  const cfg = NAV_DIMS[kind];
  if (!cfg) return notFoundHTML(raw);
  const requested = cfg.canon(raw);
  const vals = dimValues(cfg.dim);
  const hit = vals.find(x => x.v.toLowerCase() === requested.toLowerCase());
  const v = hit ? hit.v : requested;
  const members = EXT.filter(e => cfg.match(e, v)).sort((a, b) => a.name.localeCompare(b.name));
  if (!members.length) return notFoundHTML(v);
  const featured = members.filter(e => e.avail)
    .sort((a, b) => ((b.docbits ? 1 : 0) - (a.docbits ? 1 : 0)) || (b.stars || 0) - (a.stars || 0))
    .slice(0, 4);
  const isCat = kind === 'category';
  const strip = isCat ? '<div class="cat-strip">' + CAT_ORDER.map(x =>
    '<a href="' + catHref(x) + '" style="--seg:var(--c-' + x + ');flex:' + (CATS[x].count || 1) + ' 1 0" aria-current="' + (x === v) + '" data-tip="' + x + ' · ' + CATS[x].count + '"></a>').join('') + '</div>' : '';
  const navVals = isCat
    ? CAT_ORDER.map(c => ({ v: c, n: (vals.find(x => x.v === c) || {}).n || 0 }))
    : kind === 'pg'
      ? PGS.map(pg => ({ v: String(pg), n: (vals.find(x => x.v === String(pg)) || {}).n || 0 })) // active majors only
      : vals;
  const nav = kind === 'os' ? osNavHTML(v) : '<nav class="value-nav">' + navVals.map(o => {
    const seg = dimSeg(cfg.dim, o.v);
    return '<a class="facet-btn' + (isCat ? ' category' : seg ? ' hued' : '') + '" href="' + cfg.prefix + encodeURIComponent(o.v) + '"'
      + ' aria-pressed="' + (o.v === v) + '"' + (seg ? ' style="--seg:' + seg + '"' : '') + '>'
      + (seg ? '<i></i>' : '') + '<span>' + (isCat ? '<code>' + o.v + '</code>' : esc(o.v)) + '</span>'
      + (isCat ? '' : '<small>' + fmtInt(o.n) + '</small>') + '</a>';
  }).join('') + '</nav>';
  const lede = cfg.lede(v);
  const heroSeg = dimSeg(cfg.dim, v) || 'var(--accent)';
  // /pg, /os and /repo are build matrices — no featured wall, no plain table
  const body = kind === 'pg' || kind === 'os'
    ? '<div class="section"><h2>' + t('gmx.title') + '</h2><div id="slice-matrix">' + skel(6) + '</div></div>'
    : kind === 'repo'
      ? '<div class="section"><h2>' + t('gmx.title') + '</h2><div id="repo-matrix">' + skel(6) + '</div></div>'
      : (kind !== 'lang' && featured.length ? '<div class="section"><h2>' + t('cat.featured') + '</h2><ul class="wall ext-wall">' + featured.map(tileHTML).join('') + '</ul></div>' : '')
      + '<div class="section"><h2>' + t('cat.all', { n: fmtInt(members.length) }) + '</h2>' + extTableHTML(members) + '</div>';
  return '<article class="page wrap' + (kind === 'repo' || kind === 'pg' || kind === 'os' ? ' page-wide' : '') + '">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="' + dimHref(cfg.dim) + '">' + t(DIMS[cfg.dim].label) + '</a><span class="sep">/</span><span class="here">' + esc(v) + '</span></nav>'
    + strip + nav
    + '<header class="page-head cat-hero" style="--seg:' + heroSeg + ';margin-top:18px">'
    + '<span class="code">' + (isCat ? v : esc(t(DIMS[cfg.dim].label))) + '</span>'
    + '<h1>' + esc(cfg.title(v)) + '</h1>'
    + (lede ? '<p class="lede">' + esc(lede) + '</p>' : '')
    + '<p class="cat-count" style="margin-top:8px">' + fmtInt(members.length) + bi(' extensions', ' 个扩展') + ' · '
    + fmtInt(members.filter(e => e.avail).length) + bi(' packaged', ' 个已打包')
    + ' · <a href="/?' + cfg.param + '=' + encodeURIComponent(v) + '">' + t('cat.open') + '</a></p>'
    + '</header>'
    + body
    + '</article>';
}

/* /os value navigation: two aligned rows spanning the full width — x86_64
   targets on top, aarch64 underneath, one column per distro release. */
function osNavHTML(current) {
  const bases = [];
  for (const os of OSS) { const b = os.split('.')[0]; if (!bases.includes(b)) bases.push(b); }
  const cell = os => OSS.includes(os)
    ? '<a class="os-cell" href="/os/' + encodeURIComponent(os) + '" aria-current="' + (os === current) + '">' + esc(os) + '</a>'
    : '<span class="os-cell off"></span>';
  return '<nav class="os-nav" style="--os-cols:' + Math.max(bases.length, 1) + '">'
    + bases.map(b => cell(b + '.x86_64')).join('') + bases.map(b => cell(b + '.aarch64')).join('') + '</nav>';
}

/* /pg/{major} and /os/{target} slice the global matrix: one row per package;
   /pg fixes the PostgreSQL major and fans out the platforms, /os fixes the
   platform and fans out the majors. Cells carry the exact version, colored
   by state. */
function sliceMatrixHTML(data, fix) {
  const pgCount = data.pg.length;
  const fixedPG = fix.pg != null;
  const free = fixedPG ? data.os : data.pg;
  const osIdx = fixedPG ? -1 : data.os.indexOf(fix.os);
  const pgIdx = fixedPG ? data.pg.indexOf(fix.pg) : -1;
  if ((fixedPG && pgIdx < 0) || (!fixedPG && osIdx < 0)) return '<p class="empty-note">' + t('gmx.empty') + '</p>';
  const ths = fixedPG
    ? data.os.map(os => {
      const [base, arch] = os.split('.');
      return '<th class="os-th-cell"><a class="os-th" href="/os/' + encodeURIComponent(os) + '"><b>' + esc(base) + '</b><span>' + esc(arch) + '</span></a></th>';
    }).join('')
    : data.pg.map(pg => '<th><a href="/pg/' + pg + '">PG ' + pg + '</a></th>').join('');
  let rows = '';
  for (const row of data.rows) {
    let cells = '';
    for (let ci = 0; ci < free.length; ci++) {
      const idx = fixedPG ? ci * pgCount + pgIdx : osIdx * pgCount + ci;
      const ch = ((row.c || '').charAt(idx) || 'N').toUpperCase();
      const version = row.i && row.v && row.i.charAt(idx) !== '.' ? (row.v[parseInt(row.i.charAt(idx), 36)] || '') : '';
      const cls = ch === 'G' ? 'org-pgdg' : ch === 'P' ? 'org-pigsty' : ch === 'M' ? 'st-miss' : 'st-na';
      cells += '<td><span class="cellv ' + cls + '">' + (ch === 'G' || ch === 'P' ? (esc(version) || '✓') : ch === 'M' ? 'MISS' : 'N/A') + '</span></td>';
    }
    rows += '<tr><td class="oslab"><a href="' + pkgHref(row.p) + '"><b>' + esc(row.p) + '</b></a></td>' + cells + '</tr>';
  }
  const legend = '<div class="mx-legend">'
    + '<span><span class="cellv org-pgdg">PGDG</span> ' + t('mx.legend.pgdg') + '</span>'
    + '<span><span class="cellv org-pigsty">PIGSTY</span> ' + t('mx.legend.pigsty') + '</span>'
    + '<span><span class="cellv st-miss">MISS</span> ' + t('mx.legend.miss') + '</span>'
    + '<span><span class="cellv st-na">N/A</span> ' + t('mx.legend.na') + '</span>'
    + '</div>';
  return '<div class="matrix-scroll"><table class="fmx slice-mx"><thead><tr><th class="corner">'
    + bi(fixedPG ? 'PACKAGE / PLATFORM' : 'PACKAGE / PG', fixedPG ? '扩展包 / 平台' : '扩展包 / PG') + '</th>' + ths + '</tr></thead><tbody>'
    + rows + '</tbody></table></div>' + legend;
}

async function hydrateSliceMatrix(kind, value) {
  if (!document.getElementById('slice-matrix')) return;
  try {
    await ensureGlobalMatrix();
    const box = document.getElementById('slice-matrix');
    if (box) box.innerHTML = sliceMatrixHTML(GMATRIX, kind === 'pg' ? { pg: Number(value) } : { os: String(value).toLowerCase() });
  } catch (err) {
    const box = document.getElementById('slice-matrix');
    if (box) box.innerHTML = hydrateErr(err);
  }
}

// load the global matrix payload once: warm localStorage copy, then network
async function ensureGlobalMatrix() {
  if (!GMATRIX) {
    try {
      const cached = JSON.parse(localStorage.getItem(MATRIX_CACHE_KEY) || 'null');
      if (matrixUsable(cached)) GMATRIX = cached;
    } catch (e) {}
  }
  if (!GMATRIX) {
    GMATRIX = await j('/api/v1/matrix');
    try { localStorage.setItem(MATRIX_CACHE_KEY, JSON.stringify(GMATRIX)); } catch (e) {}
  }
  return GMATRIX;
}

/* /repo/{ORG} embeds the same matrix block as /matrix, restricted to the
   packages this source delivers. */
async function hydrateRepoMatrix(org) {
  const box = document.getElementById('repo-matrix');
  if (!box) return;
  try {
    await ensureGlobalMatrix();
    const memberPkgs = new Set(EXT.filter(e => e.repo === org).map(e => e.pkg));
    const rows = GMATRIX.rows.map((row, index) => ({ row, index })).filter(ref => memberPkgs.has(ref.row.p));
    if (!rows.length) {
      box.innerHTML = '<p class="empty-note">' + bi('No build-matrix rows are recorded for this source.', '该来源没有构建矩阵行记录。') + '</p>';
      return;
    }
    box.innerHTML = matrixBlockHTML(GMATRIX, rows);
    setupGlobalMatrix(GMATRIX, { root: box, rows });
  } catch (err) {
    box.innerHTML = hydrateErr(err);
  }
}

/* ---------------- view: browse & dimensions ---------------- */
const DIMS = {
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

// three six-dimension groups; categories stand alone as the first section
const DIM_GROUPS = [
  ['browse.catalog', ['tag', 'package', 'kind', 'lifecycle', 'license', 'lang']],
  ['browse.delivery', ['distribution', 'repo', 'pg', 'os', 'build', 'pgrx']],
  ['browse.attrs', ['capability', 'docs', 'relation', 'vendor', 'kernel', 'activity']]
];

const DIM_VALUE_NAMES = {
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
    'create-extension': ['CREATE EXTENSION', 'CREATE EXTENSION'], preload: ['Needs preload', '需要预加载'],
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

function dimValueName(dim, value) {
  const raw = String(value == null ? '' : value);
  const pair = DIM_VALUE_NAMES[dim] && DIM_VALUE_NAMES[dim][raw];
  if (pair) return pair[LANG === 'zh' ? 1 : 0];
  if (dim === 'pg') return 'PG ' + raw;
  if (dim === 'os') return osLabel(raw);
  return raw;
}

function dimValueHTML(dim, value) {
  const raw = String(value);
  if (dim === 'category') return '<span class="dim-category"><code>' + esc(raw) + '</code><b>' + esc(catName(raw)) + '</b></span>';
  const name = dimValueName(dim, raw);
  if (name !== raw && dim !== 'pg' && dim !== 'os' && dim !== 'activity') {
    return '<span class="dim-value-name"><b>' + esc(name) + '</b><code>' + esc(raw) + '</code></span>';
  }
  return '<span class="mono">' + esc(name) + '</span>';
}

// dimValues aggregates one dimension over the whole catalog, or over a subset
// (the home facet panel passes the scope-filtered list for cascading counts).
function dimValues(dim, subset) {
  const cnt = new Map();
  const add = (k, e) => { if (!k) return; if (!cnt.has(k)) cnt.set(k, { n: 0, ex: [] }); const o = cnt.get(k); o.n++; if (o.ex.length < 4) o.ex.push(e.name); };
  for (const e of (subset || EXT)) {
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
    else if (dim === 'os' && OSS.length) values.push(...e.targets.map(idx => OSS[idx % OSS.length]));
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
  let vals = [...cnt.entries()].map(([v, o]) => ({ v, ...o }));
  if (dim === 'category') vals.sort((a, b) => CAT_ORDER.indexOf(a.v) - CAT_ORDER.indexOf(b.v));
  else if (dim === 'pg') vals.sort((a, b) => Number(b.v) - Number(a.v));
  else if (dim === 'os') vals.sort((a, b) => OSS.indexOf(a.v) - OSS.indexOf(b.v));
  else if (dim === 'activity') vals.sort((a, b) => a.v === 'unknown' ? 1 : b.v === 'unknown' ? -1 : b.v.localeCompare(a.v));
  else if (dim === 'pgrx') vals.sort((a, b) => b.v.localeCompare(a.v, undefined, { numeric: true }));
  else vals.sort((a, b) => b.n - a.n || a.v.localeCompare(b.v));
  return vals;
}

function browseHTML() {
  const card = dim => {
    const cfg = DIMS[dim];
    const all = dimValues(dim);
    const vals = all.slice(0, 16);
    const max = Math.max(1, ...vals.map(v => v.n));
    const preview = vals.map(v => {
      const seg = dimSeg(dim, v.v) || 'var(--accent)';
      return '<i style="--w:' + v.n + ';--seg:' + seg + ';opacity:' + (0.25 + 0.7 * v.n / max).toFixed(2) + '"></i>';
    }).join('');
    return '<li><a class="dim-card" href="' + dimHref(dim) + '">'
      + '<span class="dim-field">pgext.universe · <code>' + esc(cfg.field) + '</code></span>'
      + '<span class="t">' + t(cfg.label) + '<span class="n">' + fmtInt(all.length) + '</span></span>'
      + '<span class="d">' + t(cfg.d) + '</span>'
      + '<span class="preview">' + preview + '</span></a></li>';
  };
  // categories lead as their own section: 16 compact cards in an 8 × 2 grid
  const catCards = '<section class="dimension-group"><header><h2>' + t('browse.cats') + '</h2><span>' + CAT_ORDER.length + '</span></header>'
    + '<ul class="cat-cards">' + CAT_ORDER.map(c =>
      '<li><a class="cat-card" href="' + catHref(c) + '" style="--seg:var(--c-' + c + ')">'
      + '<code>' + c + '</code><b>' + esc(catName(c)) + '</b><span>' + fmtInt((CATS[c] || {}).count || 0) + '</span></a></li>').join('')
    + '</ul></section>';
  const groups = catCards + DIM_GROUPS.map(([title, dims]) => '<section class="dimension-group"><header><h2>' + t(title) + '</h2><span>' + dims.length + '</span></header>'
    + '<ul class="dims">' + dims.map(card).join('') + '</ul></section>').join('');
  return '<article class="page wrap"><header class="page-head">'
    + '<p class="eyebrow">' + t('nav.browse').toLowerCase() + '</p>'
    + '<h1>' + t('browse.title') + '</h1><p class="lede">' + t('browse.lede') + '</p>'
    + '<div class="browse-stats"><span><b>' + Object.keys(DIMS).length + '</b>' + bi('dimensions', '个维度') + '</span>'
    + '<span><b>' + fmtInt(N_ALL) + '</b>' + bi('extensions', '个扩展') + '</span><span><b>' + fmtInt(N_PROJECTS) + '</b>' + bi('projects', '个项目') + '</span></div></header>'
    + groups + '</article>';
}

function dimHTML(dim) {
  const cfg = DIMS[dim];
  if (!cfg) return notFoundHTML(dim);
  const vals = dimValues(dim);
  const max = Math.max(1, ...vals.map(v => v.n));
  const rows = vals.map(o => {
    const href = navValueHref(dim, o.v);
    const segColor = dimSeg(dim, o.v);
    const seg = segColor ? 'style="--seg:' + segColor + '"' : '';
    const sample = o.ex.map(name => '<a href="' + extHref(name) + '">' + esc(name) + '</a>').join('');
    return '<tr ' + seg + ' data-dim-row data-dim-value="' + esc((o.v + ' ' + dimValueName(dim, o.v)).toLowerCase()) + '"><td class="v"><a href="' + href + '">' + dimValueHTML(dim, o.v) + '</a></td>'
      + '<td class="bar"><i style="width:' + Math.max(2, Math.round(100 * o.n / max)) + '%"></i></td>'
      + '<td class="n">' + fmtInt(o.n) + '</td>'
      + '<td class="ex">' + sample + '</td></tr>';
  }).join('');
  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="/list">' + t('nav.browse') + '</a><span class="sep">/</span><span class="here">' + t(cfg.label) + '</span></nav>'
    + '<header class="page-head dimension-head"><span class="dim-source">pgext.universe · <code>' + esc(cfg.field) + '</code></span>'
    + '<h1>' + t(cfg.label) + '</h1><p class="lede">' + t(cfg.d) + '</p></header>'
    + '<div class="dim-toolbar"><label><span>⌕</span><input type="search" data-dim-search placeholder="' + esc(t('dim.search', { n: fmtInt(vals.length) })) + '" autocomplete="off"></label>'
    + '<span data-dim-count>' + t('dim.showing', { shown: fmtInt(vals.length), all: fmtInt(vals.length) }) + '</span></div>'
    + '<div class="dimtable"><table><thead><tr><th>' + t('dim.value') + '</th><th></th><th style="text-align:right">' + t('dim.count') + '</th><th>' + t('dim.sample') + '</th></tr></thead>'
    + '<tbody>' + rows + '</tbody></table></div></article>';
}

/* ---------------- view: global build matrix ----------------
   Payload format matrix-row.v2 (see server/matrix.go): one row per package,
   c = one status byte per cell — G AVAIL·PGDG / P AVAIL·Pigsty / M MISS /
   N N/A — plus v/i, a row-local version dictionary with its per-cell base36
   index. Rendering follows the four states only. */
const GMX_META = [
  ['G', 'gmx.pgdg', 'pgdg'],
  ['P', 'gmx.pigsty', 'pigsty'],
  ['M', 'gmx.missing', 'missing'],
  ['N', 'gmx.na', 'na']
];
const GMX_CLASS = { G: 'pgdg', P: 'pigsty', M: 'missing', N: 'na' };
const gmxClassOf = ch => GMX_CLASS[String(ch).toUpperCase()] || 'na';

function globalMatrixShellHTML() {
  const combos = (OSS.length || 16) * (PGS.length || 5);
  return '<article class="page wrap page-wide gmx-page">'
    + '<header class="page-head"><p class="eyebrow">' + t('gmx.eyebrow') + '</p>'
    + '<div class="gmx-titleline"><div><h1>' + t('gmx.title') + '</h1>'
    + '<p class="gmx-lede">' + t('gmx.lede', { pkgs: fmtInt(N_PKGS), combos: fmtInt(combos) }) + '</p></div>'
    + '<a class="gmx-api" href="/api/v1/matrix" target="_blank" rel="noopener">' + t('gmx.api') + ' ↗</a></div></header>'
    + '<div id="gmx-root">' + skel(7) + '</div></article>';
}

/* one matrix block — toolbar, legend, readout and the canvas sheet inside a
   rounded card. Shared verbatim by /matrix (all rows) and /repo (member
   rows); counts are computed over the rows it actually shows. */
function matrixBlockHTML(data, rows) {
  const pgCount = (data.pg || []).length;
  const cols = (data.os || []).length * pgCount;
  const counts = { G: 0, P: 0, M: 0, N: 0 };
  for (const ref of rows) {
    const c = ref.row.c || '';
    for (let i = 0; i < c.length; i++) {
      const ch = c[i] >= 'a' ? String.fromCharCode(c.charCodeAt(i) - 32) : c[i];
      if (counts[ch] != null) counts[ch]++;
    }
  }
  const legend = GMX_META.map(item =>
    '<button type="button" class="gmx-legend-item gmx-' + item[2] + '" data-gmx-code="' + item[0]
    + '" aria-pressed="false"><i></i><span>' + t(item[1]) + '</span><b>' + fmtInt(counts[item[0]] || 0) + '</b></button>').join('');
  // one full target name per OS group, PG majors underneath — both navigate
  const osHeads = (data.os || []).map(osName =>
    '<a class="gmx-oshead" style="grid-column:span ' + pgCount + '" href="/os/' + encodeURIComponent(osName) + '">' + esc(osName) + '</a>').join('');
  let pgHeads = '';
  for (let o = 0; o < (data.os || []).length; o++) {
    for (let i = 0; i < pgCount; i++) {
      pgHeads += '<a class="gmx-pghead' + (i === 0 ? ' group-start' : '') + '" href="/pg/' + data.pg[i] + '">' + esc(data.pg[i]) + '</a>';
    }
  }
  return '<div class="gmx-toolbar">'
    + '<label class="gmx-search"><span>\\</span><input type="search" autocomplete="off" spellcheck="false" placeholder="'
    + esc(t('gmx.search')) + '"></label>'
    + '<div class="gmx-legend" aria-label="Matrix status filters">' + legend + '</div>'
    + '<span class="gmx-showing"></span></div>'
    + '<p class="gmx-hint">' + t('gmx.hint') + '</p>'
    + '<div class="gmx-readout" hidden></div>'
    + '<div class="mx-card"><div class="mx-scroll">'
    + '<div class="gmx-sheet" style="--gmx-columns:' + cols + '">'
    + '<div class="gmx-headrow gmx-os-row"><span class="gmx-corner">' + t('gmx.pkg') + '</span>' + osHeads + '</div>'
    + '<div class="gmx-headrow gmx-pg-row"><span class="gmx-corner gmx-corner-sub">PG →</span>' + pgHeads + '</div>'
    + '<div class="gmx-grid"><div class="gmx-labels"></div><canvas class="gmx-canvas"></canvas></div>'
    + '</div></div>'
    + '<div class="gmx-empty" hidden>' + t('gmx.empty') + '</div></div>'
    + '<p class="gmx-source"><span class="gmx-source-copy">' + t('gmx.source', { source: '<code>' + esc(data.source || 'pgext.matrix_cache') + '</code>' })
    + '</span><span class="gmx-snapshot">snapshot ' + esc((data.generated || '').replace('T', ' ').slice(0, 19)) + '</span></p>';
}

/* hover tip for one cell — everything comes from data already in hand: the
   package pattern from the bootstrap record ($v → the PG major), the exact
   version from the row-local dictionary. */
function gmxTipHTML(row, ch, os, pg, version) {
  const slim = byName.get(row.e) || byName.get(row.p);
  const isDeb = os[0] === 'd' || os[0] === 'u';
  const pattern = slim ? (isDeb ? slim.debPkg : slim.rpmPkg) : '';
  const pkgName = pattern ? pattern.split(/\s+/)[0].replaceAll('$v', String(pg)) : '';
  const meta = GMX_META.find(x => x[0] === ch);
  return '<b>' + esc(row.p) + '</b> <span class="gmx-tip-state gmx-' + gmxClassOf(ch) + '">' + (meta ? t(meta[1]) : 'N/A') + '</span><br>'
    + '<span class="d">' + esc(row.e) + ' · ' + esc(os) + ' · PG ' + esc(String(pg)) + '</span>'
    + (pkgName ? '<br><span class="k">package</span> ' + esc(pkgName) : '')
    + (version ? '<br><span class="k">version</span> ' + esc(version) : '');
}

/* The cell field is ONE canvas painted in a single pass — the same technique
   as the home-page universe dots, so 31k slots draw in a few milliseconds
   with zero DOM layout cost. Cell size is FIXED: a narrow viewport scrolls
   the card horizontally instead of recomputing anything. Hovering a cell
   tips its package name, repo and exact version from data already in hand. */
const GMX_CELL = 15, GMX_ROWH = 15;
function setupGlobalMatrix(data, opts) {
  const root = opts.root;
  const labels = root.querySelector('.gmx-labels');
  const canvas = root.querySelector('.gmx-canvas');
  const input = root.querySelector('.gmx-search input');
  const showing = root.querySelector('.gmx-showing');
  const empty = root.querySelector('.gmx-empty');
  const readout = root.querySelector('.gmx-readout');
  if (!labels || !canvas || !input) return;
  const pgCount = (data.pg || []).length;
  const cellCount = (data.os || []).length * pgCount;
  const base = opts.rows;
  const state = { rows: [], activeCode: '' };

  function draw() {
    const rows = state.rows;
    const cs = getComputedStyle(document.documentElement);
    const color = name => cs.getPropertyValue(name).trim();
    const colors = {
      G: color('--gmx-pgdg') || '#5b84ae', P: color('--gmx-pigsty') || '#639b72',
      M: color('--gmx-missing') || '#c26d5c', N: color('--gmx-na') || '#e3e2db'
    };
    const dpr = window.devicePixelRatio || 1;
    const W = cellCount * GMX_CELL, H = Math.max(rows.length * GMX_ROWH, 1);
    canvas.width = Math.round(W * dpr);
    canvas.height = Math.round(H * dpr);
    canvas.style.width = W + 'px';
    canvas.style.height = H + 'px';
    const ctx = canvas.getContext('2d');
    ctx.scale(dpr, dpr);
    for (let vi = 0; vi < rows.length; vi++) {
      const codes = rows[vi].row.c || '';
      const y = vi * GMX_ROWH;
      for (let ci = 0; ci < cellCount; ci++) {
        const ch = codes.charAt(ci) || 'N';
        ctx.fillStyle = colors[ch >= 'a' ? String.fromCharCode(ch.charCodeAt(0) - 32) : ch] || colors.N;
        ctx.fillRect(ci * GMX_CELL, y, GMX_CELL - 1, GMX_ROWH - 1);
      }
    }
    // heavier separators between OS groups
    ctx.fillStyle = color('--line-2') || '#d3d2c9';
    for (let g = 1; g < (data.os || []).length; g++) {
      ctx.fillRect(g * pgCount * GMX_CELL - 1, 0, 1, H);
    }
  }

  function renderLabels() {
    labels.innerHTML = state.rows.map(ref =>
      '<a class="gmx-row-name" href="' + pkgHref(ref.row.p) + '">' + esc(ref.row.p) + '</a>').join('');
  }

  function applyFilter() {
    const query = input.value.trim().toLowerCase();
    const code = state.activeCode;
    state.rows = base.filter(ref =>
      (!query || (ref.row.p + ' ' + ref.row.e).toLowerCase().includes(query))
      && (!code || (ref.row.c || '').toUpperCase().includes(code)));
    if (showing) showing.textContent = t('gmx.showing', { rows: fmtInt(state.rows.length), cells: fmtInt(state.rows.length * cellCount) });
    if (empty) empty.hidden = state.rows.length > 0;
    root.querySelectorAll('[data-gmx-code]').forEach(button => {
      button.setAttribute('aria-pressed', button.dataset.gmxCode === code ? 'true' : 'false');
    });
    renderLabels();
    draw();
  }

  const cellAt = ev => {
    const rect = canvas.getBoundingClientRect();
    const ci = Math.floor((ev.clientX - rect.left) / GMX_CELL);
    const vi = Math.floor((ev.clientY - rect.top) / GMX_ROWH);
    const ref = state.rows[vi];
    if (!ref || ci < 0 || ci >= cellCount) return null;
    const row = ref.row;
    return {
      row, ci,
      ch: (row.c.charAt(ci) || 'N').toUpperCase(),
      os: (data.os || [])[Math.floor(ci / pgCount)] || '',
      pg: (data.pg || [])[ci % pgCount],
      version: row.i && row.v && row.i.charAt(ci) !== '.' ? (row.v[parseInt(row.i.charAt(ci), 36)] || '') : ''
    };
  };

  // hover tip: O(1) per mousemove, nothing fetched
  canvas.addEventListener('mousemove', ev => {
    const c = cellAt(ev);
    if (c) showTip(gmxTipHTML(c.row, c.ch, c.os, c.pg, c.version), ev.clientX, ev.clientY);
    else hideTip();
  });
  canvas.addEventListener('mouseleave', hideTip);

  // clicking pins the same details into the readout strip
  canvas.addEventListener('click', ev => {
    const c = cellAt(ev);
    if (!c || !readout) return;
    const meta = GMX_META.find(x => x[0] === c.ch);
    readout.hidden = false;
    readout.innerHTML = '<span class="gmx-tip-state gmx-' + gmxClassOf(c.ch) + '">' + (meta ? t(meta[1]) : 'N/A') + '</span>'
      + '<a href="' + pkgHref(c.row.p) + '"><b>' + esc(c.row.p) + '</b></a>'
      + (c.row.e && c.row.e !== c.row.p ? '<span class="gmx-ro-dim">' + esc(c.row.e) + '</span>' : '')
      + '<code>' + esc(c.os) + '</code><code>PG ' + esc(String(c.pg)) + '</code>'
      + (c.version ? '<code class="gmx-ro-ver">' + esc(c.version) + '</code>' : '');
  });
  input.addEventListener('input', debounce(applyFilter, 90));
  root.querySelectorAll('[data-gmx-code]').forEach(button => {
    button.addEventListener('click', () => {
      state.activeCode = state.activeCode === button.dataset.gmxCode ? '' : button.dataset.gmxCode;
      applyFilter();
    });
  });
  GMATRIX_VIEW = { render: draw, state, destroy() {} };
  applyFilter();
}

/* Matrix hydration is cache-first: the last payload is kept in localStorage
   (a few KB, parses in ~1ms) so revisits render instantly; one background
   fetch per session revalidates by the materialization timestamp. */
const MATRIX_CACHE_KEY = 'pgext.matrix.v2';
let MATRIX_CHECKED = false;
const matrixUsable = data => Boolean(data && data.format === 'matrix-row.v2' && Array.isArray(data.rows) && data.rows.length);

function renderGlobalMatrix() {
  const root = document.getElementById('gmx-root');
  if (!root || !GMATRIX) return;
  if (GMATRIX_VIEW) { GMATRIX_VIEW.destroy(); GMATRIX_VIEW = null; }
  const rows = GMATRIX.rows.map((row, index) => ({ row, index }));
  root.innerHTML = matrixBlockHTML(GMATRIX, rows);
  setupGlobalMatrix(GMATRIX, { root, rows });
}

async function hydrateGlobalMatrix() {
  const token = ++matrixHydSeq;
  if (!GMATRIX) {
    try {
      const cached = JSON.parse(localStorage.getItem(MATRIX_CACHE_KEY) || 'null');
      if (matrixUsable(cached)) GMATRIX = cached;
    } catch (e) {}
  }
  if (GMATRIX) renderGlobalMatrix();
  if (MATRIX_CHECKED && GMATRIX) return;
  try {
    const fresh = await j('/api/v1/matrix');
    MATRIX_CHECKED = true;
    if (!matrixUsable(fresh)) return;
    const changed = !GMATRIX || fresh.generated !== GMATRIX.generated
      || !GMATRIX.stats || !fresh.stats || fresh.stats.cells !== GMATRIX.stats.cells;
    if (!changed) return;
    GMATRIX = fresh;
    try { localStorage.setItem(MATRIX_CACHE_KEY, JSON.stringify(fresh)); } catch (e) {}
    if (token === matrixHydSeq && currentPath === '/matrix') renderGlobalMatrix();
  } catch (err) {
    if (!GMATRIX && token === matrixHydSeq) {
      const root = document.getElementById('gmx-root');
      if (root) root.innerHTML = '<div class="gmx-frame">' + hydrateErr(err) + '</div>';
    }
  }
}

/* ---------------- view: about ---------------- */
function aboutHTML() {
  const out = (href, label, text) => '<li><span class="tag">' + label + '</span><a href="' + href + '" target="_blank" rel="noopener">' + text + ' ↗</a></li>';
  return '<article class="page wrap">'
    + '<header class="page-head"><p class="eyebrow">pgext.cloud</p>'
    + '<h1>' + bi('About this catalog', '关于本目录') + '</h1><p class="lede">'
    + bi('PGEXT.CLOUD is the census of the PostgreSQL extension ecosystem — served live from the pgext catalog database by a single Go binary.',
         'PGEXT.CLOUD 是 PostgreSQL 扩展生态的全量普查——由单个 Go 二进制从 pgext 目录数据库实时供给。') + '</p></header>'
    + '<div class="about-cols">'
    + '<div><h3>' + bi('What this is', '这是什么') + '</h3>'
    + '<p>' + bi('A curated catalog of <b>' + fmtInt(N_ALL) + '</b> PostgreSQL extensions: upstream repositories and RPM/DEB indexes are inspected continuously, with categories, dependencies, lifecycles, kernels, vendors and bilingual usage manuals maintained as catalog dimensions. <b>' + fmtInt(N_AVAIL) + '</b> of them ship as prebuilt packages across <b>' + fmtInt(OSS.length) + '</b> Linux platforms and <b>' + fmtInt(PGS.length) + '</b> PostgreSQL majors.',
                 '一份收录 <b>' + fmtInt(N_ALL) + '</b> 个 PostgreSQL 扩展的策展目录：持续抓取上游仓库与 RPM/DEB 软件源索引，维护分类、依赖、生命周期、内核、厂商与双语用法手册等目录维度。其中 <b>' + fmtInt(N_AVAIL) + '</b> 个提供预编译软件包，覆盖 <b>' + fmtInt(OSS.length) + '</b> 个 Linux 平台与 <b>' + fmtInt(PGS.length) + '</b> 个 PG 大版本。') + '</p>'
    + '<p>' + bi('This site is <code>pgext serve</code>: web assets embedded in one binary, data queried live, snapshots cached in memory. Also read our post that topped Hacker News: ',
                 '本站就是 <code>pgext serve</code>：网页资产内嵌于单个二进制，数据实时查询、内存快照缓存。也欢迎阅读我们登上 Hacker News 头条的博客：')
    + '<a href="https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4" target="_blank" rel="noopener"><i>PostgreSQL is eating the Database World</i> ↗</a></p>'
    + '<a class="about-eco" href="/matrix"><img src="https://pigsty.io/img/pigsty/ecosystem.png" alt="' + esc(bi('The PostgreSQL extension ecosystem', 'PostgreSQL 扩展生态宇宙')) + '" loading="lazy"></a>'
    + '</div>'
    + '<div><h3>' + bi('Highlights', '亮点') + '</h3><ul class="roadmap about-highlights">'
    + '<li><span class="tag">catalog</span>' + bi('<b>' + fmtInt(N_AVAIL) + '</b> packaged extensions — the largest catalog in the Postgres ecosystem', '<b>' + fmtInt(N_AVAIL) + '</b> 个已打包扩展——PG 生态最大的扩展目录') + '</li>'
    + '<li><span class="tag">native</span>' + bi('RPM/DEB packages, properly built, freely composable', 'Linux 原生 RPM/DEB 软件包，规范构建、自由组合') + '</li>'
    + '<li><span class="tag">pig</span>' + bi('A handy CLI on apt/dnf: zero-config, out-of-the-box installs', '趁手的 pig 命令行：零配置、开箱即用') + '</li>'
    + '<li><span class="tag">pgdg</span>' + bi('PGDG-compliant — drop-in with the official PostgreSQL kernel', '兼容 PGDG——与官方 PostgreSQL 内核即插即用') + '</li>'
    + '<li><span class="tag">cdn</span>' + bi('Distributed worldwide via Cloudflare CDN, fast and reliable', '通过 Cloudflare CDN 全球分发，快速可靠') + '</li>'
    + '<li><span class="tag">oss</span>' + bi('Reproducible builds on public infra — free for everyone', '可复现构建、公开基础设施，对所有人免费') + '</li>'
    + '</ul></div>'
    + '<div><h3>' + bi('Get started', '快速上手') + '</h3>'
    + mdCodeHTML('bash', 'curl -fsSL https://repo.pigsty.io/pig | bash  # install pig cli\\npig repo set                  # setup upstream repository on your linux\\npig install pg18              # install PostgreSQL 18 kernel from PGDG\\npig install pg_duckdb -v 18   # install pg_duckdb extension for PG 18') + '</div>'
    + '<div><h3>' + bi('Friends', '友情链接') + '</h3><ul class="roadmap">'
    + out('https://pigsty.io', 'pigsty.io', bi('Pigsty — Battery-Included PostgreSQL Distribution', 'Pigsty——开箱即用的 PostgreSQL 发行版'))
    + out('https://pigsty.cc', 'pigsty.cc', bi('Pigsty Chinese site & mirror', 'Pigsty 中文站与镜像'))
    + out('https://github.com/pgsty/pgext', 'github', bi('pgsty/pgext — the code behind this catalog', 'pgsty/pgext——本目录背后的代码'))
    + '</ul></div>'
    + '<div><h3>' + bi('Data', '数据') + '</h3>'
    + '<p>' + bi('The whole catalog is versioned as plain CSV in the repository\'s <code>db/</code> directory — take it straight from GitHub instead of crawling these pages.',
                 '整套目录数据以 CSV 形式版本化存放在仓库的 <code>db/</code> 目录——直接去 GitHub 获取即可，无需爬取本站页面。') + '</p>'
    + '<p><a href="https://github.com/pgsty/pgext/tree/main/db" target="_blank" rel="noopener">github.com/pgsty/pgext/tree/main/db ↗</a></p></div>'
    + '<div><h3>' + bi('Query API', '查询 API') + '</h3><ul class="roadmap">'
    + '<li><span class="tag">GET</span><code>/api/v1/ext?q=vector&cate=RAG</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/matrix</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/files?pg=18</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/ext/postgis/doc?lang=zh</code></li>'
    + '<li><span class="tag">GET</span><code>/api/v1/matrix</code></li>'
    + '</ul></div>'
    + '</div>'
    + '<p class="universe-note" style="margin-top:34px">' + bi('Snapshot loaded ' + (META.generated || '—') + ' · refreshed automatically.',
      '快照载入于 ' + (META.generated || '—') + ' · 自动刷新。') + '</p>'
    + '</article>';
}

/* ---------------- router ---------------- */
// /list/{dim} accepts a few spelling variants for shareable URLs.
const DIM_ALIASES = { cate: 'category', cat: 'category', lic: 'license', language: 'lang', pkg: 'package', type: 'kind' };
let currentPath = null;
function route() {
  const { path, params } = parseRoute();
  // Canonical routes are /ext/{name}, /pkg/{name} and /list[/{dim}]; the
  // legacy short forms redirect in place so old links keep resolving.
  if (path.startsWith('/e/')) { navigateTo('/ext/' + path.slice(3) + location.search, true); return; }
  if (path.startsWith('/p/')) { navigateTo('/pkg/' + path.slice(3) + location.search, true); return; }
  if (path.startsWith('/c/')) { navigateTo('/cate/' + path.slice(3) + location.search, true); return; }
  if (path.startsWith('/cat/')) { navigateTo('/cate/' + path.slice(5) + location.search, true); return; }
  if (path === '/browse') { navigateTo('/list' + location.search, true); return; }
  if (path.startsWith('/dim/')) { navigateTo('/list/' + path.slice(5) + location.search, true); return; }
  hideHovercard();
  const app = document.getElementById('app');
  const nav = document.getElementById('nav');
  const pathChanged = path !== currentPath;
  currentPath = path;
  if (path !== '/matrix' && GMATRIX_VIEW) { GMATRIX_VIEW.destroy(); GMATRIX_VIEW = null; }
  if (path !== '/matrix') matrixHydSeq++;
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
  } else if (path.startsWith('/pkg/')) {
    const pkg = decodeURIComponent(path.slice(5));
    app.innerHTML = pkgHTML(pkg);
    hydratePkg(pkg);
    active = '';
  } else if (path.startsWith('/cate/')) {
    app.innerHTML = navPageHTML('category', decodeURIComponent(path.slice(6))); active = '';
  } else if (path.startsWith('/lang/')) {
    app.innerHTML = navPageHTML('lang', decodeURIComponent(path.slice(6))); active = '';
  } else if (path.startsWith('/repo/')) {
    const repoValue = decodeURIComponent(path.slice(6));
    app.innerHTML = navPageHTML('repo', repoValue); active = '';
    hydrateRepoMatrix(repoValue.toUpperCase());
  } else if (path.startsWith('/license/')) {
    app.innerHTML = navPageHTML('license', decodeURIComponent(path.slice(9))); active = '';
  } else if (path.startsWith('/pg/')) {
    const pgValue = decodeURIComponent(path.slice(4));
    app.innerHTML = navPageHTML('pg', pgValue); active = 'matrix';
    hydrateSliceMatrix('pg', pgValue);
  } else if (path.startsWith('/os/')) {
    const osValue = decodeURIComponent(path.slice(4));
    app.innerHTML = navPageHTML('os', osValue); active = 'matrix';
    hydrateSliceMatrix('os', osValue);
  } else if (path === '/matrix') {
    app.innerHTML = globalMatrixShellHTML(); active = 'matrix';
    hydrateGlobalMatrix();
  } else if (path === '/list') {
    app.innerHTML = browseHTML(); active = 'browse';
  } else if (path.startsWith('/list/')) {
    const raw = decodeURIComponent(path.slice(6));
    app.innerHTML = dimHTML(DIM_ALIASES[raw] || raw); active = 'browse';
  } else if (path === '/about') {
    app.innerHTML = aboutHTML(); active = 'about';
  } else {
    app.innerHTML = notFoundHTML(path); active = '';
  }
  nav.innerHTML = navHTML(active);
  bindNavSearch();
  document.getElementById('footer').innerHTML = footerHTML();
  if (pathChanged) window.scrollTo(0, 0);
  document.title = titleFor(path);
  const meta = document.querySelector('meta[name="description"]');
  const ogTitle = document.querySelector('meta[property="og:title"]');
  const ogDesc = document.querySelector('meta[property="og:description"]');
  let description = LANG === 'zh' ? 'PostgreSQL 扩展、项目包族、依赖关系与 PG/OS 精确可用性目录。' : 'Search PostgreSQL extensions, package families, dependencies, and exact PG/OS availability.';
  if (path.startsWith('/ext/')) { const e = byName.get(decodeURIComponent(path.slice(5))); if (e) description = desc(e); }
  if (path.startsWith('/pkg/')) { const family = byPkg.get(decodeURIComponent(path.slice(5))); if (family && family[0]) description = desc(byName.get(family[0].lead) || family[0]); }
  if (path === '/matrix') description = LANG === 'zh' ? '跨扩展包、操作系统与 PostgreSQL 大版本的全局构建矩阵。' : 'The global build matrix across extension packages, operating systems, and PostgreSQL majors.';
  if (meta) meta.content = description;
  if (ogTitle) ogTitle.content = document.title;
  if (ogDesc) ogDesc.content = description;
}
function titleFor(path) {
  if (path.startsWith('/ext/')) { const n = decodeURIComponent(path.slice(5)); return n + ' · PGEXT.CLOUD'; }
  if (path.startsWith('/pkg/')) { const n = decodeURIComponent(path.slice(5)); return n + ' package · PGEXT.CLOUD'; }
  if (path.startsWith('/cate/')) return decodeURIComponent(path.slice(6)).toUpperCase() + ' · PGEXT.CLOUD';
  if (path.startsWith('/lang/')) return decodeURIComponent(path.slice(6)) + ' · PGEXT.CLOUD';
  if (path.startsWith('/repo/')) return decodeURIComponent(path.slice(6)).toUpperCase() + ' · PGEXT.CLOUD';
  if (path.startsWith('/license/')) return decodeURIComponent(path.slice(9)) + ' · PGEXT.CLOUD';
  if (path.startsWith('/pg/')) return 'PostgreSQL ' + decodeURIComponent(path.slice(4)) + ' · PGEXT.CLOUD';
  if (path.startsWith('/os/')) return decodeURIComponent(path.slice(4)) + ' · PGEXT.CLOUD';
  if (path === '/matrix') return t('gmx.title') + ' · PGEXT.CLOUD';
  if (path === '/list') return t('browse.title') + ' · PGEXT.CLOUD';
  if (path.startsWith('/list/')) {
    const raw = decodeURIComponent(path.slice(6));
    const dim = DIM_ALIASES[raw] || raw;
    return (DIMS[dim] ? t(DIMS[dim].label) : dim) + ' · PGEXT.CLOUD';
  }
  if (path === '/about') return t('nav.about') + ' · PGEXT.CLOUD';
  return 'PostgreSQL Extension Catalog';
}

function renderDynamic() {
  const d = document.getElementById('dynamic');
  if (d) { d.innerHTML = dynamicHTML(); pushState(); }
}
// The install guide is target-free; matrix-cell clicks only refresh it (and
// scroll to it) so the package hovertip target preference still lands there.
function rerenderInstall(name) {
  const e = byName.get(name), full = FULLC.get(name);
  if (!e || !full) return;
  fill('d-install', installHTML(e, full));
}
function rerenderPkgInstall(pkg) {
  const members = byPkg.get(pkg);
  if (!members || !members.length) return;
  const lead = byName.get(members[0].lead) || members[0];
  const full = FULLC.get(lead.name);
  if (!full) return;
  fill('p-install', packageInstallHTML(pkg, full));
}
/* ---------------- nav quick search ----------------
   Instant name completion over the in-memory catalog: score by how close the
   query sits to the extension (or package) name, nudge by popularity, take
   the top handful. No engine, just one pass over ~1,600 records. */
function navSuggest(query) {
  const w = query.trim().toLowerCase();
  if (!w) return [];
  const scored = [];
  for (const e of EXT) {
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

function bindNavSearch() {
  const input = document.getElementById('nav-q');
  const box = document.getElementById('nav-sugg');
  if (!input || !box) return;
  let items = [], active = -1;
  const close = () => { box.hidden = true; box.innerHTML = ''; items = []; active = -1; };
  const paint = () => {
    box.querySelectorAll('.ns-item').forEach((el, i) => el.classList.toggle('on', i === active));
  };
  const render = () => {
    items = navSuggest(input.value);
    active = items.length ? 0 : -1;
    if (!items.length) { close(); return; }
    box.hidden = false;
    box.innerHTML = items.map(e =>
      '<a class="ns-item" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '><i></i><b>' + esc(e.name) + '</b>'
      + (e.pkg !== e.name ? '<code>' + esc(e.pkg) + '</code>' : '')
      + '<span>' + esc(desc(e)) + '</span></a>').join('');
    paint();
  };
  input.addEventListener('input', debounce(render, 60));
  input.addEventListener('focus', () => { if (input.value.trim()) render(); });
  input.addEventListener('blur', () => setTimeout(close, 160));
  input.addEventListener('keydown', ev => {
    if (ev.key === 'ArrowDown' || ev.key === 'ArrowUp') {
      if (!items.length) return;
      ev.preventDefault();
      active = (active + (ev.key === 'ArrowDown' ? 1 : items.length - 1)) % items.length;
      paint();
    } else if (ev.key === 'Enter') {
      const pick = items[active] || items[0];
      if (pick) { navigateTo(extHref(pick.name)); input.value = ''; close(); input.blur(); }
    } else if (ev.key === 'Escape') {
      input.value = ''; close(); input.blur();
    }
  });
}

function bindSearch() {
  const q = document.getElementById('q');
  if (!q) return;
  q.addEventListener('input', debounce(() => {
    S.q = q.value;
    renderDynamic();
  }, 130));
  q.addEventListener('keydown', ev => {
    if (ev.key === 'Escape') { q.value = ''; S.q = ''; renderDynamic(); }
  });
}

/* ---------------- theme & lang ---------------- */
// currentTheme resolves the effective theme, including the OS preference when
// nothing is stored; the toggle simply flips between light and dark.
function currentTheme() {
  let mode = THEME_OVERRIDE || 'auto';
  if (!THEME_OVERRIDE) { try { mode = localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {} }
  if (mode === 'auto') {
    return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  }
  return mode;
}
function themeIcon() { return currentTheme() === 'dark' ? ICON_MOON : ICON_SUN; }
function cycleTheme() {
  const next = currentTheme() === 'dark' ? 'light' : 'dark';
  THEME_OVERRIDE = null; // manual toggle always wins over the ?theme= override
  try { localStorage.setItem('pgext.theme', next); } catch (e) {}
  applyTheme();
  const b = document.getElementById('theme-toggle');
  if (b) b.innerHTML = themeIcon();
  drawField();
  if (GMATRIX_VIEW) GMATRIX_VIEW.render(); // repaint the matrix canvas in the new palette
}
// ?theme=dark|light applies for the visit without touching the stored preference.
let THEME_OVERRIDE = null;
try { const q = new URLSearchParams(location.search).get('theme'); if (q === 'dark' || q === 'light') THEME_OVERRIDE = q; } catch (e) {}
function applyTheme() {
  let mode = THEME_OVERRIDE || 'auto';
  if (!THEME_OVERRIDE) { try { mode = localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {} }
  if (mode === 'auto') document.documentElement.removeAttribute('data-theme');
  else document.documentElement.setAttribute('data-theme', mode);
}
function setLang(next) {
  if ((next !== 'zh' && next !== 'en') || next === LANG) return;
  LANG = next;
  try { localStorage.setItem('pgext.lang', LANG); } catch (e) {}
  document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
  const y = window.scrollY;
  currentPath = null;
  route();
  window.scrollTo(0, y);
}
function toggleLang() { setLang(LANG === 'zh' ? 'en' : 'zh'); }

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

/* ---------------- extension hovercard ----------------
   Any element with data-hover-ext (cards, table rows) pops a rich, interactive
   detail card after a short dwell. The card itself accepts the pointer, so its
   links are clickable; the full record hydrates lazily via FULLC. */
const HOVER_DELAY = 180, HOVER_GRACE = 260;
let hcEl = null, hcName = null, hcPendingName = null, hcShowTimer = 0, hcHideTimer = 0, hcSeq = 0;

function hovercardEl() {
  if (!hcEl) {
    hcEl = document.createElement('aside');
    hcEl.id = 'hovercard';
    document.body.appendChild(hcEl);
    hcEl.addEventListener('mouseenter', () => clearTimeout(hcHideTimer));
    hcEl.addEventListener('mouseleave', scheduleHovercardHide);
  }
  return hcEl;
}

/* Hovercard anatomy: linked title + GitHub badge, package subtitle (packaged
   entries only), summary, then a fixed 2 × 4 fact grid — version/category,
   pg availability/license, kind/language, last modified/repo — plus lazily
   hydrated upstream links. */
function hovercardHTML(e, full) {
  // Links render as a compact label | URL table, one row per destination.
  const link = (label, url) => url
    ? '<a href="' + esc(url) + '" target="_blank" rel="noopener"><span>' + label + '</span><b>'
      + esc(url.replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</b><i>↗</i></a>' : '';
  const links = full ? [
    link(t('link.repo'), mdSafeURL(full.repo_url)), link(t('link.home'), mdSafeURL(full.home_url)),
    link(t('link.docs'), mdSafeURL(full.doc_url)), link('PGXN', mdSafeURL(full.pgxn_url)),
    link(t('link.license'), mdSafeURL(full.license_url)), link(t('link.control'), mdSafeURL(full.control_url))
  ].filter(Boolean).join('') : link(t('link.repo'), mdSafeURL(e.repoUrl || e.url));
  // dimension values carry their hue so the fact grid mirrors the card tags
  const hueText = (v, h) => h ? '<b style="color:var(--h-' + h + ')">' + esc(v) + '</b>' : esc(v);
  const facts = [
    [bi('version', '版本'), esc(e.ver || '—')],
    [bi('category', '分类'), '<b style="color:var(--c-' + esc(e.cat) + ')">' + esc(e.cat) + '</b>'],
    [bi('pg avail', 'PG 可用性'), esc(pgRange(e.pg))],
    [bi('license', '许可证'), hueText(e.license, licenseHue(e.license))],
    [bi('kind', '形态'), esc(e.kind || '—')],
    [bi('language', '语言'), e.lang ? hueText(e.lang, langHue(e.lang)) : '—'],
    [bi('modified', '最后修改'), esc((e.active || '').slice(0, 10) || '—')],
    [bi('provider', '仓库'), e.repo === 'n/a' ? '—' : hueText(e.repo, repoHue(e.repo))]
  ];
  return '<header ' + catVar(e.cat) + '><a class="hc-title" href="' + extHref(e.name) + '">' + esc(e.name) + '</a>'
    + cardBadgeHTML(e) + '</header>'
    + (e.avail ? '<a class="hc-sub" href="' + pkgHref(e.pkg) + '">' + bi('Package: ', '扩展包：') + esc(e.pkg) + '</a>' : '')
    + '<p class="hc-desc">' + esc(desc(e)) + '</p>'
    + '<dl class="hc-facts">' + facts.map(([k, v]) => '<div><dt>' + k + '</dt><dd>' + v + '</dd></div>').join('') + '</dl>'
    + (links ? '<div class="hc-links">' + links + '</div>' : '');
}

function positionHovercard(anchor) {
  const el = hovercardEl();
  if (!anchor.isConnected) { hideHovercard(); return; }
  const r = anchor.getBoundingClientRect();
  const w = el.offsetWidth, h = el.offsetHeight;
  const vw = window.innerWidth, vh = window.innerHeight;
  let x = r.right + 10, y = r.top;
  if (x + w > vw - 8) x = r.left - w - 10;
  if (x < 8) { x = Math.min(Math.max(8, r.left), Math.max(8, vw - w - 8)); y = r.bottom + 8; }
  if (y + h > vh - 8) y = Math.max(8, vh - h - 8);
  el.style.left = Math.round(x) + 'px';
  el.style.top = Math.round(y) + 'px';
}

/* Package hovercard: package identity header, then the lead's summary,
   facts and links, and finally the extension family as a compact headerless
   table — name (→ /ext) | version | attribute chips — matching hc-facts. */
function pkgHovercardHTML(pkg, members, lead, full) {
  const rest = members.filter(x => x.name !== lead.name).sort((a, b) => a.name.localeCompare(b.name));
  const rows = [lead, ...rest].map(x =>
    '<div><a href="' + extHref(x.name) + '">' + esc(x.name) + '</a><span>' + esc(x.ver || '—') + '</span>'
    + '<em>' + attrChips(x) + '</em></div>').join('');
  const body = hovercardHTML(lead, full);
  // reuse the extension hovercard body but swap the header for the package
  // identity, and append the family table at the very bottom
  const afterSub = body.indexOf('<p class="hc-desc"');
  return '<header ' + catVar(lead.cat) + '><a class="hc-title" href="' + pkgHref(pkg) + '">' + esc(pkg) + '</a>'
    + cardBadgeHTML(lead) + '</header>'
    + (afterSub >= 0 ? body.slice(afterSub) : '')
    + '<div class="hc-exts">' + rows + '</div>';
}

// key is "e:<extension>" or "p:<package>"
function showHovercard(anchor, key) {
  const isPkg = key.startsWith('p:');
  const name = key.slice(2);
  const members = isPkg ? byPkg.get(name) : null;
  if (isPkg && (!members || !members.length)) return;
  const lead = isPkg ? (byName.get(members[0].lead) || members[0]) : byName.get(name);
  if (!lead) return;
  const el = hovercardEl();
  hcName = key;
  const render = full => isPkg ? pkgHovercardHTML(name, members, lead, full) : hovercardHTML(lead, full);
  el.innerHTML = render(FULLC.get(lead.name));
  el.classList.add('show');
  positionHovercard(anchor);
  hideTip();
  const seq = ++hcSeq;
  if (!FULLC.get(lead.name)) {
    j('/api/v1/ext/' + encodeURIComponent(lead.name)).then(full => {
      FULLC.set(lead.name, full);
      if (seq === hcSeq && hcName === key && el.classList.contains('show')) {
        el.innerHTML = render(full);
        positionHovercard(anchor);
      }
    }).catch(() => {});
  }
}

function hideHovercard() {
  clearTimeout(hcShowTimer);
  clearTimeout(hcHideTimer);
  hcName = null; hcPendingName = null;
  if (hcEl) hcEl.classList.remove('show');
}

function scheduleHovercardHide() {
  clearTimeout(hcHideTimer);
  hcHideTimer = setTimeout(hideHovercard, HOVER_GRACE);
}

const hoverKeyOf = el => el.dataset.hoverExt ? 'e:' + el.dataset.hoverExt : 'p:' + el.dataset.hoverPkg;

function bindHovercard() {
  if (window.matchMedia && !window.matchMedia('(hover: hover) and (pointer: fine)').matches) return;
  document.addEventListener('mouseover', ev => {
    const trigger = ev.target.closest('[data-hover-ext],[data-hover-pkg]');
    if (!trigger) return;
    clearTimeout(hcHideTimer);
    const key = hoverKeyOf(trigger);
    if ((hcName === key && hcEl && hcEl.classList.contains('show')) || hcPendingName === key) return;
    hcPendingName = key;
    clearTimeout(hcShowTimer);
    hcShowTimer = setTimeout(() => { hcPendingName = null; showHovercard(trigger, key); }, HOVER_DELAY);
  });
  document.addEventListener('mouseout', ev => {
    const trigger = ev.target.closest('[data-hover-ext],[data-hover-pkg]');
    if (!trigger) return;
    if (ev.relatedTarget && (trigger.contains(ev.relatedTarget) || (hcEl && hcEl.contains(ev.relatedTarget)))) return;
    hcPendingName = null;
    clearTimeout(hcShowTimer);
    scheduleHovercardHide();
  });
  window.addEventListener('scroll', () => { if (hcName) hideHovercard(); }, { passive: true });
}

/* ---------------- global events ---------------- */
function attachEvents() {
  bindHovercard();
  window.addEventListener('popstate', route);
  window.addEventListener('hashchange', route);
  window.addEventListener('resize', debounce(() => { drawField(); if (GMATRIX_VIEW) GMATRIX_VIEW.render(); }, 150));
  window.addEventListener('scroll', hideTip, { passive: true });

  document.addEventListener('click', ev => {
    const link = ev.target.closest('a[href]');
    if (!link || link.target || link.hasAttribute('download') || ev.defaultPrevented || ev.button !== 0 || ev.metaKey || ev.ctrlKey || ev.shiftKey || ev.altKey) return;
    const url = new URL(link.href, location.href);
    if (url.origin !== location.origin || url.pathname.startsWith('/api/')) return;
    if (url.hash && url.pathname === location.pathname && url.search === location.search && !url.hash.startsWith('#/')) return;
    ev.preventDefault();
    navigateTo(url.pathname + url.search + url.hash);
  });

  document.addEventListener('click', ev => {
    const el = ev.target.closest('[data-fkey],[data-skey],[data-pg-toggle],[data-entity],[data-layout],[data-sql],[data-copy],[data-scroll],[data-itab],[data-ftab],[data-fall],[data-cat-go],[data-target-ext],[data-target-pkg],[data-lang-toggle],[data-lang-set],#random-ext,#theme-toggle,#ufield');
    if (!el) return;
    if (el.dataset.langSet) return setLang(el.dataset.langSet);
    if (el.dataset.langToggle !== undefined) return toggleLang();
    if (el.id === 'theme-toggle') return cycleTheme();
    if (el.id === 'random-ext') {
      // packaged extensions are five times as likely per entry
      const packed = EXT.filter(x => x.avail);
      const rest = EXT.filter(x => !x.avail);
      const pool = Math.random() * (packed.length * 5 + rest.length) < packed.length * 5 && packed.length ? packed : rest;
      const pick = (pool.length ? pool : EXT)[Math.floor(Math.random() * (pool.length || EXT.length))];
      if (pick) navigateTo(extHref(pick.name));
      return;
    }
    if (el.dataset.scroll) {
      const target = document.getElementById(el.dataset.scroll);
      if (target) target.scrollIntoView({ behavior: 'smooth', block: 'start' });
      return;
    }
    if (el.id === 'ufield') {
      const e = fieldHit(ev);
      if (e) navigateTo(extHref(e.name));
      return;
    }
    if (el.dataset.catGo) {
      S.cat = S.cat === el.dataset.catGo ? '' : el.dataset.catGo;
      renderDynamic();
      const d = document.getElementById('dynamic');
      if (d) d.scrollIntoView({ behavior: 'smooth', block: 'start' });
      return;
    }
    if (el.dataset.fkey) {
      const k = el.dataset.fkey, v = el.dataset.fval;
      if (k === 'scope') S.scope = S.scope === v ? '' : v;
      else S[k] = S[k] === v ? '' : v;
      if (k === 'os' && S.os) {
        INSTALL_PREF.os = S.os;
        try { localStorage.setItem('pgext.target.os', S.os); } catch (err) {}
      }
      renderDynamic(); return;
    }
    if (el.dataset.pgToggle !== undefined) {
      const pg = Number.parseInt(el.dataset.pgToggle, 10);
      const selected = new Set(parsePGs(S.pg));
      if (!pg) selected.clear();
      else if (selected.has(pg)) selected.delete(pg);
      else selected.add(pg);
      S.pg = [...selected].sort((a, b) => b - a).join(',');
      if (selected.size === 1) {
        INSTALL_PREF.pg = S.pg;
        try { localStorage.setItem('pgext.target.pg', S.pg); } catch (err) {}
      }
      renderDynamic(); return;
    }
    if (el.dataset.entity) { S.entity = el.dataset.entity; renderDynamic(); return; }
    if (el.dataset.layout) { S.layout = el.dataset.layout; renderDynamic(); return; }
    if (el.dataset.sql) {
      copyText(el.dataset.sql, ok => {
        if (!ok) return;
        const r = el.querySelector('.rcount');
        if (r) { const old = r.textContent; r.textContent = t('search.copied'); r.classList.add('copied'); setTimeout(() => { r.textContent = old; r.classList.remove('copied'); }, 1200); }
      });
      return;
    }
    if (el.dataset.copy) {
      const old = el.textContent;
      copyText(el.dataset.copy, ok => {
        if (ok) { el.textContent = 'copied ✓'; el.classList.add('ok'); setTimeout(() => { el.textContent = old; el.classList.remove('ok'); }, 1200); }
      });
      return;
    }
    if (el.dataset.targetExt || el.dataset.targetPkg) {
      INSTALL_PREF.pg = el.dataset.targetPg;
      INSTALL_PREF.os = el.dataset.targetOs;
      try {
        localStorage.setItem('pgext.target.pg', el.dataset.targetPg);
        localStorage.setItem('pgext.target.os', el.dataset.targetOs);
      } catch (err) {}
      if (el.dataset.targetExt) rerenderInstall(el.dataset.targetExt, { pg: el.dataset.targetPg, os: el.dataset.targetOs });
      else rerenderPkgInstall(el.dataset.targetPkg, { pg: el.dataset.targetPg, os: el.dataset.targetOs });
      const install = document.getElementById(el.dataset.targetExt ? 'd-install' : 'p-install');
      if (install) install.scrollIntoView({ behavior: 'smooth', block: 'center' });
      return;
    }
    if (el.dataset.itab !== undefined) {
      const box = el.closest('.install');
      box.querySelectorAll('[data-itab]').forEach(b => b.setAttribute('aria-selected', b === el));
      box.querySelectorAll('[data-ipane]').forEach(p => { p.hidden = p.dataset.ipane !== el.dataset.itab; });
      return;
    }
    if (el.dataset.ftab !== undefined) {
      const box = el.closest('.file-browser');
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
    if ((el.dataset.skey === 'pg' || el.dataset.skey === 'os') && el.value) {
      INSTALL_PREF[el.dataset.skey] = el.value;
      try { localStorage.setItem('pgext.target.' + el.dataset.skey, el.value); } catch (err) {}
    }
    renderDynamic();
  });

  document.addEventListener('input', ev => {
    const input = ev.target.closest('[data-dim-search]');
    if (!input) return;
    const query = input.value.trim().toLowerCase();
    const rows = [...document.querySelectorAll('[data-dim-row]')];
    let shown = 0;
    for (const row of rows) {
      row.hidden = query !== '' && !row.dataset.dimValue.includes(query);
      if (!row.hidden) shown++;
    }
    const count = document.querySelector('[data-dim-count]');
    if (count) count.textContent = t('dim.showing', { shown: fmtInt(shown), all: fmtInt(rows.length) });
  });

  document.addEventListener('keydown', ev => {
    if (ev.key === '/' && !ev.metaKey && !ev.ctrlKey && !ev.altKey) {
      const tag = (document.activeElement && document.activeElement.tagName) || '';
      if (tag === 'INPUT' || tag === 'TEXTAREA' || tag === 'SELECT') return;
      ev.preventDefault();
      const { path } = parseRoute();
      // the home hero has its own search box — everywhere else the slash
      // lights up the quick search in the nav bar
      const target = (path === '/' || path === '') ? document.getElementById('q') : document.getElementById('nav-q');
      if (target) target.focus();
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

/* Boot with a warm-start cache: the last bootstrap payload is kept in
   localStorage so revisits render instantly from the cached catalog, then a
   conditional fetch (If-None-Match) revalidates it in the background. All
   filtering, sorting and search run client-side over this dataset. */
const BOOT_CACHE_KEY = 'pgext.boot.v5';

async function boot() {
  applyTheme();
  document.documentElement.lang = LANG === 'zh' ? 'zh-CN' : 'en';
  let cached = null;
  try { cached = JSON.parse(localStorage.getItem(BOOT_CACHE_KEY) || 'null'); } catch (e) {}
  if (cached && Array.isArray(cached.rows) && cached.rows.length && cached.rows[0].length >= 34) {
    try { decodeBoot(cached); route(); } catch (e) { cached = null; }
  } else { cached = null; }
  try {
    const headers = { Accept: 'application/json' };
    if (cached && cached.version) headers['If-None-Match'] = '"b34-' + cached.version + '"';
    // The fmt query parameter mirrors the server's payload-format version so
    // the browser HTTP cache can never serve an older payload layout to a
    // newer app.js (the format also salts the ETag server-side).
    const res = await fetch('/api/v1/bootstrap?fmt=b34', { headers });
    if (res.status === 304) return; // cached catalog is current
    if (!res.ok) throw new Error('HTTP ' + res.status);
    const b = await res.json();
    const hadCache = Boolean(cached);
    decodeBoot(b);
    try { localStorage.setItem(BOOT_CACHE_KEY, JSON.stringify(b)); } catch (e) {}
    const y = window.scrollY;
    currentPath = null; // re-render the live route with fresh data
    route();
    if (hadCache) window.scrollTo(0, y);
  } catch (err) {
    if (!cached) bootError(err);
  }
}

attachEvents();
boot();
