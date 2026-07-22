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
   24 target_idx[] 25 family_size 26 comment 27 relationbits 28 pgrx_ver */
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
    relationbits: r[27] | 0, pgrx: r[28] || ''
  }));
  byName = new Map(EXT.map(e => [e.name, e]));
  byPkg = new Map();
  for (const e of EXT) { if (!byPkg.has(e.pkg)) byPkg.set(e.pkg, []); byPkg.get(e.pkg).push(e); }
  CATS = {}; CAT_ORDER = [];
  for (const c of b.cats) { CATS[c.name] = { code: c.name, zh: c.zh_desc, en: c.en_desc, count: c.count }; CAT_ORDER.push(c.name); }
  PGS = b.pg || [];
  OSS = b.os || [];
  OSIDX = {}; OSS.forEach((os, i) => { OSIDX[os] = i; });
  N_ALL = b.counts.total; N_AVAIL = b.counts.packaged; N_PROJECTS = b.counts.projects || byPkg.size;
  N_SOURCE = b.counts.source_only || 0; N_VENDOR = b.counts.vendor || 0; N_KERNEL = b.counts.kernel || 0;
  N_CONTRIB = b.counts.contrib || 0; N_DOCS = b.counts.docs || 0;
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
  'nav.matrix': ['Matrix', '全局矩阵'],
  'nav.browse': ['Browse', '多维索引'],
  'nav.about': ['About', '关于'],
  'nav.lang': ['中文', 'EN'],
  'hero.eyebrow': ['pgext.cloud · the postgresql extension catalog', 'pgext.cloud · PostgreSQL 扩展目录'],
  'hero.title': ['Everything <em>PostgreSQL</em> can become.', 'PostgreSQL 的<em>一切可能</em>。'],
  'hero.sub': ['Search <b>{all}</b> extensions across <b>{projects}</b> upstream projects, with exact package availability for every supported PostgreSQL, OS, and architecture target.',
               '检索来自 <b>{projects}</b> 个上游项目的 <b>{all}</b> 个扩展，并精确查询每个 PostgreSQL、操作系统与架构目标上的软件包可用性。'],
  'hero.s1': ['extensions catalogued', '个扩展收录在册'],
  'hero.s2': ['upstream projects', '个上游项目'],
  'hero.s3': ['packaged extensions', '个已打包扩展'],
  'hero.s4': ['usage documents', '份用法文档'],
  'field.note': ['Every dot is one extension, colored by category — hover to identify, click to open.', '每一个点都是一个扩展，颜色代表分类——悬停查看，点击进入。'],
  'search.ph': ['Search {n} extensions — name, package, tags, or try tag:vector build:pgrx', '搜索 {n} 个扩展——名称、项目、标签，或试试 tag:vector build:pgrx'],
  'search.copy': ['click to copy query', '点击复制查询'],
  'search.copied': ['copied ✓', '已复制 ✓'],
  'chip.all': ['All', '全部'],
  'chip.packaged': ['Packaged', '已打包'],
  'chip.source': ['Source-only', '仅源码'],
  'chip.kernel': ['Kernel-specific', '内核特定'],
  'chip.vendor': ['Vendor catalog', '厂商目录'],
  'chip.contrib': ['PostgreSQL contrib', 'PG 自带'],
  'sel.cat': ['category', '分类'],
  'sel.license': ['license', '许可证'],
  'sel.lang': ['language', '语言'],
  'sel.pg': ['pg version', 'PG 版本'],
  'sel.os': ['binary target', '二进制目标'],
  'sel.kind': ['extension kind', '扩展形态'],
  'sel.lifecycle': ['lifecycle', '生命周期'],
  'filter.scope': ['catalog scope', '收录范围'],
  'filter.category': ['category', '分类'],
  'filter.pg': ['PostgreSQL · all selected', 'PostgreSQL · 同时支持所选版本'],
  'filter.kind': ['extension kind', '扩展形态'],
  'filter.repo': ['package source', '软件包来源'],
  'filter.lifecycle': ['lifecycle', '生命周期'],
  'filter.lang': ['implementation', '实现语言'],
  'filter.os': ['binary target', '二进制目标'],
  'filter.license': ['common licenses', '常用许可证'],
  'filter.any': ['Any', '不限'],
  'filter.pg.hint': ['Multiple selections use <code>pg_ver @&gt; \'{…}\'</code>: every selected major must be supported.', '多选采用 <code>pg_ver @&gt; \'{…}\'</code>：扩展必须同时支持全部所选大版本。'],
  'filter.license.all': ['all {n} licenses', '全部 {n} 种许可证'],
  'filter.dimensions': ['all dimensions', '全部维度'],
  'filter.active': ['Active filters', '当前筛选'],
  'filter.results.ext': ['matching extensions', '个匹配扩展'],
  'filter.results.pkg': ['matching projects', '个匹配项目'],
  'filter.entity': ['catalog', '目录对象'],
  'filter.layout': ['layout', '呈现方式'],
  'sort.stars': ['sort: stars', '排序：星标'],
  'sort.name': ['sort: name', '排序：名称'],
  'sort.updated': ['sort: updated', '排序：更新'],
  'view.ext': ['EXT', 'EXT'],
  'view.pkg': ['PKG', 'PKG'],
  'view.card': ['Card', '卡片'],
  'view.list': ['List', '列表'],
  'wall.empty': ['No extension matches. Loosen a filter, or try <code>vector</code>, <code>gis</code>, <code>parquet</code>.',
                 '没有匹配的扩展。放宽筛选条件，或试试 <code>vector</code>、<code>gis</code>、<code>parquet</code>。'],
  'rows.name': ['name', '名称'], 'rows.cat': ['category', '分类'], 'rows.ver': ['version', '版本'],
  'rows.license': ['license', '许可证'], 'rows.lang': ['lang', '语言'], 'rows.pg': ['pg', 'PG'],
  'rows.stars': ['stars', '星标'], 'rows.desc': ['description', '描述'],
  'rows.package': ['package', '项目包'], 'rows.lead': ['lead extension', '主扩展'],
  'rows.extensions': ['extensions', '扩展数'], 'rows.packaged': ['packaged', '已打包'],
  'ext.crumb': ['extensions', '扩展'],
  'ext.overview': ['Overview', '概览'],
  'ext.metadata': ['Metadata', '元数据'],
  'ext.relations': ['Related extensions', '关联扩展'],
  'ext.versions': ['Versions & package definitions', '版本与包定义'],
  'ext.build': ['Build manual', '构建手册'],
  'ext.install': ['Install', '安装'],
  'ext.avail': ['Availability', '可用性矩阵'],
  'ext.pkgs': ['Packages & Downloads', '安装包与下载'],
  'ext.downloads': ['Downloads', '软件下载'],
  'ext.deps': ['Dependencies & Relations', '依赖与关联'],
  'ext.docs': ['Usage', '用法'],
  'ext.onpage': ['On this page', '本页目录'],
  'ext.catalog': ['Catalog identity', '目录标识'],
  'ext.runtime': ['Runtime behavior', '运行时行为'],
  'ext.packaging': ['Packaging metadata', '打包元数据'],
  'ext.upstream': ['Upstream & activity', '上游与活跃度'],
  'ext.projectlinks': ['Project resources', '项目资源'],
  'ext.freshness': ['Catalog freshness', '目录时效'],
  'ext.packageintro': ['The extension definition, RPM and DEB rows below follow the same model used by <code>pgext gen io/cc</code>. Package patterns use <code>$v</code> as the PostgreSQL major placeholder.',
                       '下面的扩展定义、RPM 与 DEB 行沿用 <code>pgext gen io/cc</code> 的内容模型；包名中的 <code>$v</code> 代表 PostgreSQL 大版本。'],
  'ext.buildnone': ['The catalog has source metadata but no reproducible RPM/DEB build recipe. Follow the upstream build documentation.',
                    '目录已收录源码信息，但没有可复现的 RPM/DEB 构建配方；请遵循上游构建文档。'],
  'ext.contribbuild': ['This extension is delivered with PostgreSQL contrib; build it together with the matching PostgreSQL server packages.',
                       '这是 PostgreSQL contrib 扩展，应随对应版本的 PostgreSQL 服务端软件包一起构建和交付。'],
  'ext.buildrecipe': ['A Pigsty package recipe is recorded for {types}. Build it with the same package-family name used by the catalog.',
                      '目录记录了 {types} 的 Pigsty 构建配方，请使用目录中的包族名称进行构建。'],
  'ext.sourcearchive': ['Source archive', '源码归档'],
  'ext.installguide': ['Install, load & enable', '安装、加载与启用'],
  'ext.installlede': ['Choose an exact PostgreSQL and OS target. The plan is generated from live <code>pgext.pkg</code> rows, then completed with preload, CREATE EXTENSION and verification steps from <code>pgext.universe</code>.',
                      '选择精确的 PostgreSQL 与操作系统目标。安装计划来自实时 <code>pgext.pkg</code> 行，并结合 <code>pgext.universe</code> 生成预加载、CREATE EXTENSION 与验证步骤。'],
  'ext.step.repo': ['1 · Repository', '1 · 软件仓库'],
  'ext.step.package': ['2 · Install package', '2 · 安装软件包'],
  'ext.step.load': ['3 · Load library', '3 · 加载动态库'],
  'ext.step.enable': ['4 · Enable extension', '4 · 启用扩展'],
  'ext.step.verify': ['5 · Verify', '5 · 验证结果'],
  'ext.noload': ['No shared preload is required.', '无需配置 shared_preload_libraries。'],
  'ext.noddl': ['This entry does not create SQL objects with CREATE EXTENSION.', '该条目无需通过 CREATE EXTENSION 创建 SQL 对象。'],
  'ext.docmeta': ['{sections} sections · about {minutes} min read · {lang} document', '{sections} 节 · 约 {minutes} 分钟阅读 · {lang} 文档'],
  'ext.docsource': ['Curated Markdown stored in pgext.doc', '来自 pgext.doc 的策展 Markdown 文档'],
  'ext.showall': ['Show all {n}', '显示全部 {n} 项'],
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
  'ext.providerNote': ['This catalog entry is associated with {provider}. No public binary package is currently recorded; consult the upstream or provider documentation.',
                       '该目录项关联 {provider}，目前没有公开二进制包记录；请参考上游或服务商文档。'],
  'ext.lifecycleNote': ['Upstream lifecycle is {state}. Evaluate maintenance and compatibility before production use.',
                        '上游生命周期状态为 {state}，用于生产前请评估维护状态与兼容性。'],
  'ext.family': ['Package family', '同包扩展族'],
  'spec.id': ['id', 'ID'],
  'spec.version': ['version', '版本'], 'spec.state': ['state', '状态'], 'spec.category': ['category', '分类'],
  'spec.license': ['license', '许可证'], 'spec.language': ['language', '语言'], 'spec.repo': ['repo', '仓库'],
  'spec.package': ['package', '包名'], 'spec.lead': ['lead ext', '主扩展'],
  'spec.schemas': ['schemas', '模式'], 'spec.pg': ['pg support', 'PG 支持'],
  'spec.rpm': ['rpm', 'RPM'], 'spec.rpmdeps': ['rpm deps', 'RPM 依赖'],
  'spec.deb': ['deb', 'DEB'], 'spec.debdeps': ['deb deps', 'DEB 依赖'],
  'spec.source': ['source', '源码包'], 'spec.vendor': ['vendor', '厂商'],
  'spec.tags': ['tags', '标签'],
  'spec.kind': ['kind', '形态'], 'spec.lifecycle': ['lifecycle', '生命周期'],
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
  'mx.legend.pgdg': ['available · PGDG repo', '可用 · PGDG 仓库'],
  'mx.legend.pigsty': ['available · Pigsty repo', '可用 · Pigsty 仓库'],
  'mx.legend.miss': ['MISS · package missing', 'MISS · 软件包缺失'],
  'mx.legend.na': ['N/A · invalid combination', 'N/A · 无效组合'],
  'files.none': ['No binary artifacts recorded for this extension.', '没有该扩展的二进制包记录。'],
  'files.showall': ['show all {n} builds', '显示全部 {n} 个构建'],
  'files.collapse': ['latest builds only', '只看最新构建'],
  'files.count': ['{n} artifacts for PG {pg}', 'PG {pg} 共 {n} 个包文件'],
  'files.os': ['os', '系统'], 'files.pkg': ['package', '包'], 'files.ver': ['version', '版本'],
  'files.org': ['org', '来源'], 'files.size': ['size', '大小'], 'files.file': ['file', '文件'],
  'files.sha': ['sha256', 'SHA256'],
  'link.home': ['homepage', '主页'], 'link.repo': ['repository', '仓库'], 'link.license': ['license', '许可证'],
  'link.docs': ['documentation', '官方文档'], 'link.pgxn': ['PGXN', 'PGXN'],
  'link.control': ['control file', 'control 文件'], 'link.author': ['author', '作者'], 'link.cargo': ['cargo', 'cargo'],
  'type.standard': ['standard — CREATE EXTENSION, no preload', 'standard——直接 CREATE EXTENSION，无需预加载'],
  'type.preload': ['preload — needs shared_preload_libraries', 'preload——需要 shared_preload_libraries'],
  'type.puresql': ['puresql — SQL objects only, no binary', 'puresql——纯 SQL 对象，无二进制'],
  'type.headless': ['headless — library only, no SQL objects', 'headless——只有库，无 SQL 对象'],
  'matrix.ext': ['CREATE EXTENSION', 'CREATE EXTENSION'],
  'gmx.eyebrow': ['package intelligence · pgext.matrix', '软件包情报 · pgext.matrix'],
  'gmx.title': ['Global Build Matrix', '全局构建矩阵'],
  'gmx.lede': ['One operational view of every package, Linux target, and supported PostgreSQL major. Each colored square is one exact build combination from the latest CI materialization.',
               '用一张表总览每个扩展包、Linux 目标与 PostgreSQL 大版本。每个彩色方格都代表最近一次 CI 物化结果中的一个精确构建组合。'],
  'gmx.packages': ['package families', '个扩展包族'],
  'gmx.targets': ['OS targets', '个 OS 目标'],
  'gmx.pg': ['PG majors', '个 PG 大版本'],
  'gmx.cells': ['build cells', '个构建格子'],
  'gmx.search': ['Filter package or extension…', '筛选扩展包或扩展名…'],
  'gmx.showing': ['{rows} packages · {cells} visible cells', '{rows} 个扩展包 · {cells} 个可见格子'],
  'gmx.source': ['Precomputed in {source} from CI-ingested repository metadata and atomically refreshed with the package catalog. Hover a cell for package, repository, version, and artifact count.',
                 '数据由 CI 回传的软件仓库元数据预计算至 {source}，并与软件包目录原子刷新。悬停格子可查看包名、仓库、版本与制品数量。'],
  'gmx.pgdg': ['PGDG', 'PGDG'],
  'gmx.pigsty': ['Pigsty', 'Pigsty'],
  'gmx.missing': ['Missing', '缺失'],
  'gmx.na': ['N/A', 'N/A'],
  'gmx.empty': ['No package row matches this filter.', '没有扩展包行符合当前筛选条件。'],
  'gmx.hint': ['Click a status to isolate it · click again to reset', '点击状态可单独查看 · 再次点击恢复全部'],
  'gmx.pkg': ['PACKAGE / EXTENSION', '扩展包 / 扩展'],
  'gmx.api': ['JSON API', 'JSON API'],
  'cat.crumb': ['categories', '分类'],
  'cat.featured': ['Featured', '精选'],
  'cat.all': ['All {n} extensions', '全部 {n} 个扩展'],
  'cat.open': ['Open in search ↗', '在搜索中打开 ↗'],
  'browse.title': ['Browse by dimension', '多维索引'],
  'browse.lede': ['One catalog, many ways in. Slice the extension universe by any dimension — every value below is a live filter.',
                  '一份目录，多个入口。任选维度切分扩展宇宙——下面每个值都是一个即时筛选器。'],
  'browse.catalog': ['Identity & classification', '身份与分类'],
  'browse.delivery': ['Packaging & delivery', '构建与交付'],
  'browse.runtime': ['Runtime & documentation', '运行时与文档'],
  'browse.ecosystem': ['Ecosystem & activity', '生态与活跃度'],
  'dim.category': ['Categories', '分类'],
  'dim.tag': ['Tags', '主题标签'],
  'dim.package': ['Package Families', '项目包族'],
  'dim.license': ['Licenses', '许可证'],
  'dim.lang': ['Languages', '实现语言'],
  'dim.repo': ['Repositories', '仓库来源'],
  'dim.distribution': ['Distribution Channels', '分发渠道'],
  'dim.os': ['Binary Targets', '二进制目标'],
  'dim.build': ['Build Toolchains', '构建链'],
  'dim.pgrx': ['pgrx Versions', 'pgrx 版本'],
  'dim.capability': ['Runtime Capabilities', '运行时能力'],
  'dim.docs': ['Documentation Coverage', '文档覆盖'],
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
  'about.title': ['About this catalog', '关于本目录'],
  'about.lede': ['PGEXT.CLOUD is the census of the PostgreSQL extension ecosystem — served live from the pgext catalog database by a single Go binary.',
                 'PGEXT.CLOUD 是 PostgreSQL 扩展生态的全量普查——由单个 Go 二进制从 pgext 目录数据库实时供给。'],
  'about.p1': ['Every entry is drawn from the <code>pgext</code> catalog: upstream repositories and RPM/DEB indexes are inspected, while project activity, lifecycle, dependencies, kernels, and vendors remain distinct catalog dimensions. Counts and documentation coverage come from the live snapshot.',
               '每个条目都来自 <code>pgext</code> 目录数据库：抓取上游仓库、解析 RPM/DEB 包索引，并记录项目活跃度、生命周期、依赖、内核与厂商关系。当前统计与文档覆盖均由实时快照生成。'],
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
  'about.colophon': ['Snapshot loaded {date} · refreshed automatically · authenticated reload is available only when enabled by the operator.',
                     '快照载入于 {date} · 自动刷新 · 手工刷新仅在运维方启用认证令牌后可用。'],
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
const extHref = n => '/e/' + encodeURIComponent(n);
const pkgHref = n => '/p/' + encodeURIComponent(n);
const catHref = n => '/c/' + encodeURIComponent(n);
const dimHref = n => '/dim/' + encodeURIComponent(n);
const catVar = c => 'style="--seg:var(--c-' + esc(c) + ')"';
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
    if (usageMode) level = level > 2 ? level - 1 : 2;
    const id = mdSlug(title, used);
    const clean = mdPlain(title);
    toc.push({ level, title: clean, id });
    out.push('<h' + level + ' id="' + id + '">' + mdInline(title) + '</h' + level + '>');
  };
  let i = 0, paragraphCount = 0;
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
        + '<pre><code class="language-' + esc(language) + '">' + esc(raw) + '</code></pre></div>');
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
      out.push('<blockquote><p>' + mdInline(quote.join(' ')) + '</p></blockquote>');
      continue;
    }

    const para = [line.trim()]; i++;
    while (i < lines.length && lines[i].trim() && !isBlock(lines[i], lines[i + 1])) {
      if (i + 1 < lines.length && /^\s*(?:={3,}|-{3,})\s*$/.test(lines[i + 1])) break;
      para.push(lines[i++].trim());
    }
    const paragraph = para.join(' ');
    let cls = '';
    if (usageMode && /^\s*(?:sources?|来源|资料来源)\s*[:：]/i.test(paragraph)) cls = ' class="doc-sources"';
    else if (usageMode && paragraphCount === 0) cls = ' class="doc-lead"';
    out.push('<p' + cls + '>' + mdInline(paragraph) + '</p>');
    if (!cls.includes('doc-sources')) paragraphCount++;
  }
  const words = src.trim() ? src.trim().split(/\s+/).length : 0;
  const cjk = (src.match(/[\u3400-\u9fff]/g) || []).length;
  const minutes = Math.max(1, Math.ceil(LANG === 'zh' ? Math.max(cjk / 500, words / 260) : words / 220));
  return { html: out.join('\n'), toc, minutes };
}

/* ---------------- landing state / filter engine ---------------- */
const FILTER_KEYS = ['q', 'cat', 'repo', 'license', 'lng', 'pg', 'os', 'kind', 'lifecycle', 'scope',
  'vendor', 'kernel', 'tag', 'pkg', 'capability', 'build', 'docs', 'relation', 'pgrx', 'active'];
const DEFAULT_STATE = {
  q: '', cat: '', repo: '', license: '', lng: '', pg: '', os: '', kind: '', lifecycle: '', scope: '',
  vendor: '', kernel: '', tag: '', pkg: '', capability: '', build: '', docs: '', relation: '', pgrx: '', active: '',
  sort: 'stars', entity: 'ext', layout: 'card'
};
let S = { ...DEFAULT_STATE };

function migrateLegacyHash() {
  const raw = location.hash.startsWith('#/') ? location.hash.slice(1) : '';
  if (!raw) return false;
  const qi = raw.indexOf('?');
  let path = qi >= 0 ? raw.slice(0, qi) : raw;
  const query = qi >= 0 ? raw.slice(qi + 1) : '';
  if (path.startsWith('/ext/')) path = '/e/' + path.slice(5);
  else if (path.startsWith('/pkg/')) path = '/p/' + path.slice(5);
  else if (path.startsWith('/cat/')) path = '/c/' + path.slice(5);
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
  for (const k of FILTER_KEYS) if (S[k]) p.set(k, S[k]);
  if (S.sort !== 'stars') p.set('sort', S.sort);
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
    cat: S.cat, repo: S.repo, license: S.license, lng: S.lng, pg: S.pg, os: S.os,
    kind: S.kind, lifecycle: S.lifecycle, scope: S.scope, vendorQ: S.vendor.toLowerCase(), kernelQ: S.kernel.toLowerCase(),
    tag: S.tag, pkg: S.pkg, capability: S.capability, build: S.build, docs: S.docs,
    relation: S.relation, pgrx: S.pgrx, active: S.active, words: []
  };
  for (const tok of S.q.trim().split(/\s+/).filter(Boolean)) {
    const m = tok.match(/^(cat|category|repo|license|lang|lng|pg|os|target|type|kind|life|lifecycle|kernel|vendor|tag|tags|pkg|package|project|cap|capability|feature|build|builder|doc|docs|rel|relation|dependency|pgrx|active|year|is):(.+)$/i);
    if (!m) { f.words.push(tok.toLowerCase()); continue; }
    const key = m[1].toLowerCase(), val = m[2];
    if (key === 'cat' || key === 'category') f.cat = val.toUpperCase();
    else if (key === 'repo') f.repo = val.toUpperCase();
    else if (key === 'license') f.license = val;
    else if (key === 'lang' || key === 'lng') f.lng = val;
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
      if (v === 'source' || v === 'source-only' || v === 'unpackaged') f.scope = 'source';
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

function runFilter() {
  const f = effectiveFilters();
  const pgs = parsePGs(f.pg);
  const list = [];
  for (const e of EXT) {
    if (f.cat && e.cat !== f.cat) continue;
    if (f.repo && e.repo !== f.repo) continue;
    if (f.license && e.license.toLowerCase() !== f.license.toLowerCase()) continue;
    if (f.lng && e.lang.toLowerCase() !== f.lng.toLowerCase()) continue;
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
    if (f.scope === 'packaged' && !e.avail) continue;
    if (f.scope === 'source' && (e.avail || !(e.buildbits & B.SOURCE))) continue;
    if (f.scope === 'kernel' && !e.kernel) continue;
    if (f.scope === 'vendor' && !e.vendor) continue;
    if (f.scope === 'contrib' && !(e.flags & F.CONTRIB)) continue;
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
  list.sort((a, b) => {
    if (hasQ && b[0] !== a[0]) return b[0] - a[0];
    if (S.sort === 'name') return a[1].name.localeCompare(b[1].name);
    if (S.sort === 'updated') return (b[1].active || '').localeCompare(a[1].active || '') || (b[1].stars || 0) - (a[1].stars || 0);
    return (b[1].stars || 0) - (a[1].stars || 0) || a[1].name.localeCompare(b[1].name);
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
  if (f.lng) wh.push('lang = ' + lit(f.lng));
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
  if (f.scope === 'source') wh.push("NOT packaged AND (repo_url IS NOT NULL OR tarball IS NOT NULL)");
  if (f.scope === 'kernel') wh.push('kernel IS NOT NULL');
  if (f.scope === 'vendor') wh.push('vendor IS NOT NULL');
  if (f.scope === 'contrib') wh.push('contrib');
  const sql = 'SELECT * FROM pgext.universe' + (wh.length ? ' WHERE ' + wh.join(' AND ') : '') + ' ORDER BY stars DESC NULLS LAST;';
  const kw = s => '<span class="kw">' + s + '</span>';
  let html = kw('SELECT') + ' * ' + kw('FROM') + ' pgext.universe';
  if (wh.length) html += ' ' + kw('WHERE') + ' ' + esc(wh.join(' AND ')).replace(/&#39;([^&]*)&#39;/g, '<span class="lit">&#39;$1&#39;</span>');
  html += ' ' + kw('ORDER BY') + ' stars <span class="kw">DESC NULLS LAST</span>;';
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
    + (!e.avail ? '<span class="status source">source</span>' : '')
    + (e.lifecycle ? '<span class="status life">' + esc(e.lifecycle) + '</span>' : '')
    + '<span class="lic">' + esc(e.license === 'Unknown' ? '' : e.license) + '</span></span>'
    + '</a></li>';
}

function projectGroups(list) {
  const groups = new Map();
  for (const e of list) {
    if (!groups.has(e.pkg)) groups.set(e.pkg, []);
    groups.get(e.pkg).push(e);
  }
  return [...groups.entries()].map(([pkg, matched]) => {
    const all = byPkg.get(pkg) || matched;
    const leadName = (all[0] && all[0].lead) || pkg;
    const lead = byName.get(leadName) || matched[0] || all[0];
    return { pkg, matched, all, lead };
  });
}

function packageTileHTML(g) {
  const e = g.lead;
  const packaged = g.all.filter(x => x.avail).length;
  const cats = [...new Set(g.all.map(x => x.cat))];
  const matched = g.matched.length === g.all.length ? '' : ' · ' + g.matched.length + ' matched';
  return '<li><a class="tile project" href="' + pkgHref(g.pkg) + '" ' + catVar(e.cat) + '>'
    + '<span class="tile-head"><span class="tile-name">' + esc(g.pkg) + '</span>'
    + (e.stars ? '<span class="tile-star">★ ' + fmtNum(e.stars) + '</span>' : '') + '</span>'
    + '<span class="tile-desc">' + esc(desc(e)) + '</span>'
    + '<span class="family-preview">' + g.all.slice(0, 4).map(x => '<code>' + esc(x.name) + '</code>').join('')
    + (g.all.length > 4 ? '<code>+' + (g.all.length - 4) + '</code>' : '') + '</span>'
    + '<span class="tile-meta"><span class="cat">' + esc(cats.slice(0, 3).join(' · ')) + '</span>'
    + '<span>' + g.all.length + (LANG === 'zh' ? ' 个扩展' : ' ext') + matched + '</span>'
    + '<span class="lic">' + packaged + (LANG === 'zh' ? ' 已打包' : ' packaged') + '</span></span>'
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

function packageRowHTML(g) {
  const e = g.lead;
  const cats = [...new Set(g.all.map(x => x.cat))];
  const pgs = [...new Set(g.all.flatMap(x => x.pg))].sort((a, b) => b - a);
  const packaged = g.all.filter(x => x.avail).length;
  const matched = g.matched.length === g.all.length ? fmtInt(g.all.length) : fmtInt(g.matched.length) + ' / ' + fmtInt(g.all.length);
  return '<tr ' + catVar(e.cat) + '><td class="r-name"><a href="' + pkgHref(g.pkg) + '">' + esc(g.pkg) + '</a></td>'
    + '<td class="r-mono"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a></td>'
    + '<td class="r-desc">' + esc(desc(e)) + '</td>'
    + '<td class="r-cat">' + esc(cats.join(' · ')) + '</td>'
    + '<td class="r-num">' + matched + '</td><td class="r-num">' + fmtInt(packaged) + '</td>'
    + '<td class="r-mono">' + esc(pgRange(pgs)) + '</td><td class="r-num">' + (e.stars ? fmtNum(e.stars) : '') + '</td></tr>';
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
    + '<a class="brand" href="/"><span class="brand-mark">\\dx</span><span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<nav class="nav-links">'
    + '<a href="/" aria-current="' + (active === 'home') + '">' + t('nav.ext') + '</a>'
    + '<a href="/matrix" aria-current="' + (active === 'matrix') + '">' + t('nav.matrix') + '</a>'
    + '<a href="/browse" aria-current="' + (active === 'browse') + '">' + t('nav.browse') + '</a>'
    + '<a href="/about" aria-current="' + (active === 'about') + '">' + t('nav.about') + '</a>'
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
    [fmtInt(N_PROJECTS), t('hero.s2')],
    [fmtInt(N_AVAIL), t('hero.s3')],
    [fmtInt(N_DOCS), t('hero.s4')]
  ].map(([n, l]) => '<li><span class="num">' + n + '</span><span class="lbl">' + l + '</span></li>').join('');
  const legend = CAT_ORDER.map(c =>
    '<li style="--seg:var(--c-' + c + ')">'
    + '<button data-cat-go="' + c + '" data-tip="' + esc(catDesc(c)) + '">' + c
    + '<span class="cnt">' + CATS[c].count + '</span></button></li>'
  ).join('');
  return '<section class="hero wrap">'
    + '<p class="eyebrow">' + t('hero.eyebrow') + '</p>'
    + '<h1>' + t('hero.title') + '</h1>'
    + '<p class="hero-sub">' + t('hero.sub', { all: fmtInt(N_ALL), projects: fmtInt(N_PROJECTS), avail: fmtInt(N_AVAIL) }) + '</p>'
    + '<div class="console">'
    + '<div class="searchbox"><span class="prompt">pgext=#</span>'
    + '<input id="q" type="search" autocomplete="off" spellcheck="false" placeholder="' + esc(t('search.ph', { n: fmtInt(N_ALL) })) + '" value="' + esc(S.q) + '" aria-label="search extensions">'
    + '<kbd>/</kbd></div>'
    + '</div>'
    + '<ul class="hero-stats">' + stats + '</ul>'
    + '<div class="universe"><canvas id="ufield" height="80" aria-hidden="true" style="width:100%;display:block;cursor:crosshair"></canvas>'
    + '<ul class="ubar-legend">' + legend + '</ul>'
    + '<p class="universe-note">' + t('field.note') + '</p></div>'
    + '<div id="dynamic">' + dynamicHTML() + '</div>'
    + '</section>';
}

function dynamicHTML() {
  const { f, list } = runFilter();
  const { sql, html: sqlHTML } = buildSQL(f, list.length);
  const sortOpts = [['stars', t('sort.stars')], ['name', t('sort.name')], ['updated', t('sort.updated')]].map(([v, l]) =>
    '<option value="' + v + '"' + (S.sort === v ? ' selected' : '') + '>' + l + '</option>').join('');

  const countMap = dim => new Map(dimValues(dim).map(x => [String(x.v), x.n]));
  const button = (key, value, label, on, count, cls, style) => {
    const data = key === 'pg' ? ' data-pg-toggle="' + esc(value) + '"' : ' data-fkey="' + key + '" data-fval="' + esc(value) + '"';
    return '<button class="facet-btn' + (cls ? ' ' + cls : '') + '"' + data + ' aria-pressed="' + on + '"' + (style ? ' style="' + style + '"' : '') + '>'
      + (cls === 'category' ? '<i></i>' : '') + '<span>' + label + '</span>'
      + (count == null ? '' : '<small>' + fmtInt(count) + '</small>') + '</button>';
  };
  const any = (key, on) => button(key, '', t('filter.any'), on, null, 'any');
  const row = (label, values, hint, cls) => '<div class="facet-row' + (cls ? ' ' + cls : '') + '"><div class="facet-name"><b>' + label + '</b>'
    + (hint ? '<span>' + hint + '</span>' : '') + '</div><div class="facet-values">' + values + '</div></div>';

  const scopeValues = [
    ['', t('chip.all'), N_ALL], ['packaged', t('chip.packaged'), N_AVAIL], ['source', t('chip.source'), N_SOURCE],
    ['kernel', t('chip.kernel'), N_KERNEL], ['vendor', t('chip.vendor'), N_VENDOR], ['contrib', t('chip.contrib'), N_CONTRIB]
  ].map(([v, label, count]) => button('scope', v, label, S.scope === v, count)).join('');
  const categoryValues = any('cat', !S.cat) + CAT_ORDER.map(c => button('cat', c, '<code>' + c + '</code><em>' + esc(catName(c)) + '</em>', S.cat === c, CATS[c].count, 'category', '--seg:var(--c-' + c + ')')).join('');

  const pgCounts = countMap('pg'), selectedPGs = parsePGs(S.pg);
  const pgValues = [...new Set([...PGS, ...selectedPGs])].sort((a, b) => b - a);
  const pgButtons = button('pg', '', t('filter.any'), !selectedPGs.length, null, 'any')
    + pgValues.map(pg => button('pg', String(pg), 'PG ' + pg, selectedPGs.includes(pg), pgCounts.get(String(pg)) || 0)).join('');

  const kindCounts = countMap('kind');
  const kindButtons = any('kind', !S.kind) + [...kindCounts.entries()].map(([v, count]) => button('kind', v, esc(v), S.kind === v, count)).join('');
  const repoCounts = countMap('repo');
  const repoButtons = any('repo', !S.repo) + [...repoCounts.entries()].filter(([v]) => v && v !== 'n/a')
    .map(([v, count]) => button('repo', v, esc(v), S.repo === v, count)).join('');
  const lifecycleCounts = countMap('lifecycle');
  const lifecycleButtons = any('lifecycle', !S.lifecycle) + [...lifecycleCounts.entries()]
    .map(([v, count]) => button('lifecycle', v, esc(v), S.lifecycle === v, count)).join('');
  const langCounts = countMap('lang');
  const langButtons = any('lng', !S.lng) + [...langCounts.entries()]
    .map(([v, count]) => button('lng', v, esc(v), S.lng === v, count)).join('');
  const osCounts = countMap('os');
  const osButtons = any('os', !S.os) + OSS.map(v => button('os', v, esc(osLabel(v)), S.os === v, osCounts.get(v) || 0)).join('');
  const licenses = dimValues('license');
  const visibleLicenses = licenses.slice(0, 10);
  if (S.license && !visibleLicenses.some(x => x.v === S.license)) {
    const selected = licenses.find(x => x.v === S.license);
    if (selected) visibleLicenses.push(selected);
  }
  const licenseButtons = any('license', !S.license) + visibleLicenses.map(x => button('license', x.v, esc(x.v), S.license === x.v, x.n)).join('')
    + '<a class="facet-more" href="' + dimHref('license') + '">' + t('filter.license.all', { n: fmtInt(licenses.length) }) + ' →</a>';
  const facets = '<section class="facet-panel" aria-label="filters">'
    + row(t('filter.scope'), scopeValues, '', 'scope-facet')
    + row(t('filter.category'), categoryValues, '', 'category-facet')
    + row(t('filter.pg'), pgButtons, t('filter.pg.hint'), 'pg-facet')
    + row(t('filter.kind'), kindButtons)
    + row(t('filter.repo'), repoButtons)
    + row(t('filter.lifecycle'), lifecycleButtons)
    + row(t('filter.lang'), langButtons)
    + row(t('filter.os'), osButtons)
    + row(t('filter.license'), licenseButtons)
    + '<footer><a class="dimension-link" href="/browse">＋ ' + t('filter.dimensions') + '</a><span>pgext.universe · pgext.pkg · pgext.doc</span></footer></section>';

  const projects = S.entity === 'pkg' ? projectGroups(list) : [];
  const itemTotal = S.entity === 'pkg' ? projects.length : list.length;
  let results;
  if (!itemTotal) {
    results = '<div class="result-empty">' + t('wall.empty') + '</div>';
  } else if (S.entity === 'pkg' && S.layout === 'card') {
    results = '<ul class="wall project-wall">' + projects.map(packageTileHTML).join('') + '</ul>';
  } else if (S.entity === 'pkg') {
    results = '<div class="rows"><div class="rows-scroll"><table><thead><tr>'
      + '<th>' + t('rows.package') + '</th><th>' + t('rows.lead') + '</th><th>' + t('rows.desc') + '</th><th>' + t('rows.cat') + '</th>'
      + '<th>' + t('rows.extensions') + '</th><th>' + t('rows.packaged') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
      + '</tr></thead><tbody>' + projects.map(packageRowHTML).join('') + '</tbody></table></div></div>';
  } else if (S.layout === 'list') {
    results = '<div class="rows"><div class="rows-scroll"><table><thead><tr>'
      + '<th>' + t('rows.name') + '</th><th>' + t('rows.cat') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th>'
      + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
      + '</tr></thead><tbody>' + list.map(rowHTML).join('') + '</tbody></table></div></div>';
  } else {
    results = '<ul class="wall">' + list.map(tileHTML).join('') + '</ul>';
  }

  const advanced = [
    ['vendor', 'vendor'], ['kernel', 'kernel'], ['tag', 'tag'], ['pkg', 'package'],
    ['capability', 'capability'], ['build', 'build'], ['docs', 'docs'], ['relation', 'relation'],
    ['pgrx', 'pgrx'], ['active', 'activity']
  ].filter(([key]) => S[key]);
  const activeFilters = advanced.length
    ? '<div class="active-filters"><span>' + t('filter.active') + '</span>' + advanced.map(([key, dim]) =>
      '<button data-fkey="' + key + '" data-fval="' + esc(S[key]) + '"><b>' + t(DIMS[dim].label) + '</b>'
      + '<code>' + esc(dimValueName(dim, S[key])) + '</code><i>×</i></button>').join('') + '</div>'
    : '';
  const toolbar = '<div class="catalog-toolbar"><div class="result-summary"><strong>' + fmtInt(itemTotal) + '</strong><span>'
    + t(S.entity === 'pkg' ? 'filter.results.pkg' : 'filter.results.ext') + '</span></div><div class="catalog-controls">'
    + '<select data-skey="sort" aria-label="sort">' + sortOpts + '</select>'
    + '<div class="control-cluster"><span>' + t('filter.entity') + '</span><span class="viewtoggle">'
    + '<button data-entity="ext" aria-pressed="' + (S.entity === 'ext') + '">' + t('view.ext') + '</button>'
    + '<button data-entity="pkg" aria-pressed="' + (S.entity === 'pkg') + '">' + t('view.pkg') + '</button></span></div>'
    + '<div class="control-cluster"><span>' + t('filter.layout') + '</span><span class="viewtoggle">'
    + '<button data-layout="card" aria-pressed="' + (S.layout === 'card') + '">' + t('view.card') + '</button>'
    + '<button data-layout="list" aria-pressed="' + (S.layout === 'list') + '">' + t('view.list') + '</button></span></div></div></div>';

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

function manualOutlineHTML(items, extraClass) {
  return '<nav class="manual-outline ' + (extraClass || '') + '" aria-label="' + esc(t('ext.onpage')) + '"><span>' + t('ext.onpage') + '</span>'
    + items.map(([id, label]) => '<button data-scroll="' + id + '">' + label + '</button>').join('') + '</nav>';
}

function extOutlineHTML() {
  return manualOutlineHTML([
    ['ext-overview', 'ext.overview'], ['ext-metadata', 'ext.metadata'], ['ext-relations', 'ext.relations'],
    ['ext-packages', 'ext.pkgs'], ['ext-build', 'ext.build'], ['ext-install', 'ext.installguide'], ['ext-usage', 'ext.docs']
  ].map(([id, key]) => [id, t(key)]), 'ext-outline');
}

function extHTML(name) {
  const e = byName.get(name);
  if (!e) return notFoundHTML(name);
  const related = EXT.filter(x => x.cat === e.cat && x.name !== e.name && x.avail)
    .sort((a, b) => (b.stars || 0) - (a.stars || 0)).slice(0, 4);
  const providerNote = (e.kernel || e.vendor) && !e.avail
    ? '<div class="notice provider">☁ ' + t('ext.providerNote', { provider: esc([e.vendor, e.kernel].filter(Boolean).join(' · ')) }) + '</div>' : '';
  const lifecycleNote = ['deprecated', 'archived', 'abandoned'].includes(e.lifecycle)
    ? '<div class="notice lifecycle">⚠ ' + t('ext.lifecycleNote', { state: '<b>' + esc(e.lifecycle) + '</b>' }) + '</div>' : '';

  return '<article class="page wrap manual-page ext-page">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="' + catHref(e.cat) + '" style="color:var(--c-' + e.cat + ')">' + e.cat + '</a><span class="sep">/</span>'
    + '<span class="here">' + esc(e.name) + '</span></nav>'
    + '<header class="detail-hero ext-detail-hero"><div class="detail-kicker">EXTENSION</div><div class="ext-head"><h1>' + esc(e.name) + '</h1>' + (e.ver ? '<span class="ver">v' + esc(e.ver) + '</span>' : '') + '</div>'
    + '<p class="ext-tagline">' + esc(desc(e)) + '</p>'
    + '<div class="badge-row">'
    + '<a class="badge cat" href="' + catHref(e.cat) + '" ' + catVar(e.cat) + '><span class="dot"></span>' + e.cat + ' · ' + esc(catName(e.cat)) + '</a>'
    + '<a class="badge" href="/?license=' + encodeURIComponent(e.license) + '">' + esc(e.license) + '</a>'
    + '<a class="badge" href="/?lng=' + encodeURIComponent(e.lang) + '">' + esc(e.lang) + '</a>'
    + (e.repo !== 'n/a' ? '<a class="badge" href="/?repo=' + encodeURIComponent(e.repo) + '">' + esc(e.repo) + '</a>' : '')
    + (e.kind ? '<a class="badge" href="/?kind=' + encodeURIComponent(e.kind) + '" data-tip="' + esc(t('type.' + e.kind)) + '">' + esc(e.kind) + '</a>' : '')
    + (e.lifecycle ? '<a class="badge" href="/?lifecycle=' + encodeURIComponent(e.lifecycle) + '">' + esc(e.lifecycle) + '</a>' : '')
    + (e.kernel ? '<span class="badge">kernel · ' + esc(e.kernel) + '</span>' : '')
    + (e.vendor ? '<span class="badge">☁ ' + esc(e.vendor) + '</span>' : '')
    + '</div>'
    + '<div class="badge-row">' + attrBadges(e) + '</div>'
    + '<div class="ghstats" id="d-gh">'
    + (e.stars != null ? '<span>★ <b>' + fmtInt(e.stars) + '</b></span>' : '')
    + (e.active ? '<span>' + t('commit.freshness', { d: '<b>' + esc(e.active) + '</b>' }) + '</span>' : '')
    + (e.familySize > 1 ? '<a href="' + pkgHref(e.pkg) + '">' + esc(e.pkg) + ' · ' + e.familySize + (LANG === 'zh' ? ' 个同包扩展' : ' extensions') + ' ↗</a>' : '')
    + '</div>' + lifecycleNote + providerNote + '</header>'
    + extOutlineHTML()
    + '<div class="ext-grid ext-manual"><main class="ext-main">'
    + '<section class="section ext-section" id="ext-overview"><h2>' + t('ext.overview') + '</h2><div id="d-overview">' + skel(5) + '</div></section>'
    + '<section class="section ext-section" id="ext-metadata"><h2>' + t('ext.metadata') + '</h2><div id="d-metadata">' + skel(7) + '</div></section>'
    + '<section class="section ext-section" id="ext-relations"><h2>' + t('ext.relations') + '</h2><div id="d-deps">' + skel(4) + '</div></section>'
    + '<section class="section ext-section" id="ext-packages"><h2>' + t('ext.pkgs') + '</h2>'
    + '<p class="section-lede">' + t('ext.packageintro') + '</p><div id="d-versions">' + skel(3) + '</div>'
    + '<h3 class="section-subhead">' + t('ext.avail') + '</h3><div id="d-matrix">' + skel(5) + '</div>'
    + '<h3 class="section-subhead">' + t('ext.downloads') + '</h3><div id="d-files">' + skel(4) + '</div></section>'
    + '<section class="section ext-section" id="ext-build"><h2>' + t('ext.build') + '</h2><div id="d-build">' + skel(3) + '</div></section>'
    + '<section class="section ext-section" id="ext-install"><h2>' + t('ext.installguide') + '</h2><p class="section-lede">' + t('ext.installlede') + '</p><div id="d-install">' + skel(6) + '</div></section>'
    + '</main><aside class="ext-rail"><div id="d-spec">' + skel(8) + '</div></aside></div>'
    + '<section class="section ext-section usage-section" id="ext-usage"><h2>' + t('ext.docs') + '</h2><div id="d-doc">' + (e.docbits ? skel(8) : '<div class="docs-missing">' + t('ext.docsnone') + '</div>') + '</div></section>'
    + (related.length ? '<section class="section related-section"><h2>' + t('ext.related', { cat: e.cat }) + '</h2><ul class="related">' + related.map(tileHTML).join('') + '</ul></section>' : '')
    + '</article>';
}

function sqlIdent(name) {
  return /^[a-z_][a-z0-9_$]*$/.test(name) ? name : '"' + String(name).replaceAll('"', '""') + '"';
}

function shellArg(value) {
  return "'" + String(value).replaceAll("'", "'\"'\"'") + "'";
}

function installPrefs(m, override) {
  const cells = (m && m.cells || []).filter(c => c.state === 'AVAIL' && c.name);
  let pg = (override && override.pg) || INSTALL_PREF.pg || '', os = (override && override.os) || INSTALL_PREF.os || '';
  try {
    const storedPG = localStorage.getItem('pgext.target.pg'), storedOS = localStorage.getItem('pgext.target.os');
    if (!pg && storedPG) pg = storedPG;
    if (!os && storedOS) os = storedOS;
  } catch (err) {}
  if (pg && os) return { pg: parseInt(pg, 10), os, cells };
  const pgRank = new Map((m && m.pg || PGS).map((v, i) => [v, i]));
  const osRank = new Map((m && m.os || OSS).map((v, i) => [v, i]));
  const preferred = ['el9.x86_64', 'd12.x86_64', 'u24.x86_64', 'el8.x86_64'];
  cells.sort((a, b) => (pgRank.get(a.pg) ?? 99) - (pgRank.get(b.pg) ?? 99)
    || (preferred.indexOf(a.os) < 0 ? 99 : preferred.indexOf(a.os)) - (preferred.indexOf(b.os) < 0 ? 99 : preferred.indexOf(b.os))
    || (osRank.get(a.os) ?? 99) - (osRank.get(b.os) ?? 99));
  return cells.length ? { pg: cells[0].pg, os: cells[0].os, cells } : { pg: (m && m.pg && m.pg[0]) || PGS[0] || 18, os: '', cells };
}

function installHTML(e, full, matrix, override) {
  const tabs = [];
  const pref = installPrefs(matrix, override);
  const cell = pref.cells.find(c => c.pg === pref.pg && c.os === pref.os);
  const ident = sqlIdent(e.name);
  const requires = full.requires || [];
  const cascade = requires.length ? ' CASCADE' : '';
  const createSQL = full.need_ddl ? 'CREATE EXTENSION ' + ident + cascade + ';' : '';
  const verifySQL = full.need_ddl ? "SELECT extname, extversion FROM pg_extension WHERE extname = '" + e.name.replaceAll("'", "''") + "';" : '';
  const libs = (full.preload_libs || []).length ? full.preload_libs : ((full.libs || []).length ? full.libs : [e.name]);
  const repoName = String((cell && cell.org) || full.repo || '').toUpperCase();
  const repoCmd = full.contrib || !full.packaged ? '' : (repoName === 'PGDG' ? 'pig repo add pgdg -u' : 'pig repo add pgsql -u');
  const loadConfig = full.need_load ? "shared_preload_libraries = '" + libs.join(', ') + "'" : '';
  const stepsPlain = [];
  if (repoCmd) stepsPlain.push(repoCmd);

  if (cell) {
    const manager = cell.os.startsWith('el') ? 'dnf' : 'apt-get';
    const installCmd = manager === 'dnf' ? 'sudo dnf install -y ' + cell.name : 'sudo apt-get install -y ' + cell.name;
    const target = 'PG ' + cell.pg + ' · ' + osLabel(cell.os) + ' · ' + (cell.org || 'repository') + ' · ' + (cell.version || full.version || '');
    const activation = [];
    if (full.need_load) activation.push('# postgresql.conf: append while preserving existing preload entries', loadConfig, '# restart PostgreSQL after changing shared_preload_libraries');
    if (createSQL) activation.push('psql -d mydb -c ' + shellArg(createSQL));
    if (verifySQL) activation.push('psql -d mydb -c ' + shellArg(verifySQL));
    const planPlain = ['# Exact catalog target: ' + target, ...stepsPlain, installCmd, ...activation].join('\n');
    const planHTML = '<span class="cmt"># Exact catalog target: ' + esc(target) + '</span>\n' + esc([...stepsPlain, installCmd, ...activation].join('\n'));
    tabs.push(['plan', planHTML, planPlain]);
    tabs.push([manager, '<span class="cmt"># ' + esc(target) + '</span>\n' + esc(installCmd), installCmd]);
    const pigCmd = 'pig ext install -y ' + (full.pkg || e.pkg || e.name) + ' -v ' + cell.pg;
    tabs.push(['pig', '<span class="cmt"># project package · target PG ' + cell.pg + '</span>\n' + esc(pigCmd), pigCmd]);
  } else if (full.contrib) {
    tabs.push(['contrib', '<span class="cmt"># Included with the matching PostgreSQL contrib/server packages.\n# No separate extension package is recorded.</span>', '']);
  } else if (full.packaged) {
    tabs.push(['unavailable', '<span class="cmt"># No AVAIL package is recorded for PG ' + esc(pref.pg) + (pref.os ? ' · ' + esc(osLabel(pref.os)) : '') + '.\n# Choose another target or inspect the availability matrix below.</span>', '']);
  } else {
    const upstream = full.repo_url || full.doc_url || full.url || '';
    tabs.push(['source', '<span class="cmt"># No public binary package is recorded.\n# Follow the upstream build or provider instructions.</span>'
      + (upstream ? '\n' + esc(upstream) : ''), upstream]);
  }

  const pgs = (matrix && matrix.pg || PGS);
  const oss = (matrix && matrix.os || OSS);
  const controls = matrix && matrix.cells && matrix.cells.length ? '<div class="install-target">'
    + '<label>PostgreSQL <select data-install-env="pg" data-install-ext="' + esc(e.name) + '">'
    + pgs.map(pg => '<option value="' + pg + '"' + (pg === pref.pg ? ' selected' : '') + '>PG ' + pg + '</option>').join('') + '</select></label>'
    + '<label>' + (LANG === 'zh' ? '系统与架构' : 'OS & architecture') + ' <select data-install-env="os" data-install-ext="' + esc(e.name) + '">'
    + oss.map(os => '<option value="' + esc(os) + '"' + (os === pref.os ? ' selected' : '') + '>' + esc(osLabel(os)) + '</option>').join('') + '</select></label>'
    + '<span class="target-state ' + (cell ? 'ok' : 'bad') + '">' + (cell ? '● AVAIL' : '○ MISS') + '</span></div>' : '';
  const heads = tabs.map(([n], i) => '<button role="tab" aria-selected="' + (i === 0) + '" data-itab="' + i + '">' + esc(n) + '</button>').join('');
  const bodies = tabs.map(([, html, copy], i) =>
    '<div class="install-body" data-ipane="' + i + '"' + (i ? ' hidden' : '') + '>'
    + (copy ? '<button class="copy-btn" data-copy="' + esc(copy) + '">copy</button>' : '')
    + '<pre>' + html + '</pre></div>').join('');
  const code = (language, value) => '<div class="md-code compact"><div class="md-codebar"><span>' + language + '</span><button data-copy="' + esc(value) + '">copy</button></div><pre><code>' + esc(value) + '</code></pre></div>';
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
  const verifyBody = verifySQL ? code('sql', verifySQL)
    : full.need_load ? code('sql', 'SHOW shared_preload_libraries;')
      : '<p>' + bi('Use the upstream project health check for this headless/provider entry.', '对于此 headless/服务商条目，请使用上游项目提供的健康检查。') + '</p>';
  return '<div class="install-steps">'
    + '<section class="install-step"><h3>' + t('ext.step.repo') + '</h3>' + repoBody + '</section>'
    + '<section class="install-step package-step"><h3>' + t('ext.step.package') + '</h3>' + controls + '<div class="install"><div class="install-tabs" role="tablist">' + heads + '</div>' + bodies + '</div></section>'
    + '<section class="install-step"><h3>' + t('ext.step.load') + '</h3>' + loadBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.enable') + '</h3>' + enableBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.verify') + '</h3>' + verifyBody + '</section></div>';
}

/* full availability matrix: os rows × pg columns from /matrix */
function fullMatrixHTML(m, e, options) {
  const opts = options || {};
  const byKey = {};
  for (const c of m.cells) byKey[c.os + '|' + c.pg] = c;
  const famOf = os => os.startsWith('el') ? 'EL' : os.startsWith('d') ? 'Debian' : 'Ubuntu';
  const ths = m.pg.map(pg => '<th>PG ' + pg + '</th>').join('');
  // synthetic first row: CREATE EXTENSION availability from pg_ver
  const extRow = '<tr><td class="oslab"><b>' + t('matrix.ext') + '</b> · v' + esc(e.ver || '—') + '</td>'
    + m.pg.map(pg => (e.pg || []).includes(pg)
      ? '<td><span class="cellv org-other" data-tip="' + esc(e.name) + '">✓</span></td>'
      : '<td><span class="cellv st-na">N/A</span></td>').join('') + '</tr>';
  let prevFam = null, rows = '';
  for (const os of m.os) {
    const fam = famOf(os);
    const famStart = fam !== prevFam;
    prevFam = fam;
    const [osname, arch] = os.split('.');
    const cells = m.pg.map(pg => {
      const c = byKey[os + '|' + pg];
      if (!c || c.state === 'N/A') return '<td><span class="cellv st-na">N/A</span></td>';
      if (c.state === 'MISS') return '<td><span class="cellv st-miss">MISS</span></td>';
      if (c.state === 'AVAIL') {
        const org = (c.org || '').toLowerCase();
        const cls = org === 'pgdg' ? 'org-pgdg' : org === 'pigsty' ? 'org-pigsty' : 'org-other';
        const target = opts.pkg ? ' data-target-pkg="' + esc(opts.pkg) + '"' : ' data-target-ext="' + esc(e.name) + '"';
        return '<td><button class="cellv ' + cls + '"' + target + ' data-target-pg="' + c.pg + '" data-target-os="' + esc(c.os) + '" data-tip="' + esc(c.name || '') + ' · ' + c.count + ' builds">' + esc(c.org || '✓') + ' ' + esc(c.version || '') + '</button></td>';
      }
      return '<td><span class="cellv st-na">N/A</span></td>';
    }).join('');
    rows += '<tr' + (famStart ? ' class="fam-start"' : '') + '><td class="oslab"><b>' + esc(osname) + '</b> · ' + esc(arch) + '</td>' + cells + '</tr>';
  }
  const legend = '<div class="mx-legend">'
    + '<span><span class="cellv org-pgdg">pgdg</span> ' + t('mx.legend.pgdg') + '</span>'
    + '<span><span class="cellv org-pigsty">pigsty</span> ' + t('mx.legend.pigsty') + '</span>'
    + '<span><span class="cellv st-miss">MISS</span> ' + t('mx.legend.miss') + '</span>'
    + '<span><span class="cellv st-na">N/A</span> ' + t('mx.legend.na') + '</span>'
    + '</div>';
  return '<div class="matrix-scroll"><table class="fmx"><thead><tr><th></th>' + ths + '</tr></thead><tbody>'
    + (opts.includeExtension === false ? '' : extRow) + rows + '</tbody></table></div>' + legend;
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
  return pgs.length ? '<span class="pg-badges">' + pgs.map(pg => '<span>PG' + pg + '</span>').join('') + '</span>' : '—';
}

function overviewHTML(e, full) {
  const fact = (label, value, note, mono) => '<div class="detail-fact"><span>' + label + '</span><strong' + (mono ? ' class="mono"' : '') + '>' + value + '</strong>'
    + (note ? '<small>' + note + '</small>' : '') + '</div>';
  const distribution = full.contrib ? bi('PostgreSQL contrib', 'PostgreSQL 自带')
    : full.packaged ? bi('Prebuilt packages', '预编译软件包')
      : (full.repo_url || full.tarball) ? bi('Source catalog', '源码目录') : bi('Provider catalog', '服务商目录');
  const docLangs = [full.has_en_doc ? 'EN' : '', full.has_zh_doc ? '中文' : ''].filter(Boolean).join(' + ') || '—';
  const flags = [
    ['attr.bin', full.has_bin], ['attr.lib', full.has_lib], ['attr.load', full.need_load], ['attr.ddl', full.need_ddl],
    ['attr.trusted', full.trusted], ['attr.reloc', full.relocatable], ['attr.contrib', full.contrib]
  ];
  return '<div class="overview-copy"><p>' + esc(LANG === 'zh' ? (full.zh_desc || full.en_desc || e.zh || e.en) : (full.en_desc || full.zh_desc || e.en || e.zh)) + '</p>'
    + (full.comment ? '<div class="catalog-note"><b>' + bi('Catalog note', '目录备注') + '</b><span>' + esc(full.comment) + '</span></div>' : '') + '</div>'
    + '<div class="detail-facts">'
    + fact(bi('Project package', '项目包族'), '<a href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a>', full.family_size + bi(' extension definitions', ' 个扩展定义'), true)
    + fact(bi('Catalog version', '目录版本'), esc(full.version || '—'), bi('extension control version', '扩展 control 版本'), true)
    + fact(bi('PostgreSQL', 'PostgreSQL'), esc(pgRange(full.pg_ver)), bi('declared compatibility', '声明兼容范围'), true)
    + fact(bi('Distribution', '交付方式'), esc(distribution), full.repo && full.repo !== 'n/a' ? esc(full.repo) : '', false)
    + fact(bi('Extension kind', '扩展形态'), esc(full.kind || '—'), esc(t('type.' + (full.kind || ''))), true)
    + fact(bi('Usage manual', '用法手册'), esc(docLangs), full.has_en_doc || full.has_zh_doc ? t('ext.docsource') : t('ext.docsnone'), true)
    + '</div><div class="attribute-grid">' + flags.map(([key, value]) => '<div class="attribute ' + (value ? 'on' : 'off') + '"><span>' + (value ? '✓' : '—') + '</span><div><b>' + t(key) + '</b><small>' + boolHTML(value) + '</small></div></div>').join('') + '</div>';
}

function metadataHTML(full) {
  const row = (label, value, mono) => '<div class="meta-row"><dt>' + label + '</dt><dd' + (mono ? ' class="mono"' : '') + '>' + (value || '—') + '</dd></div>';
  const panel = (title, rows) => '<section class="meta-panel"><h3>' + title + '</h3><dl>' + rows.join('') + '</dl></section>';
  const names = values => (values || []).length ? esc(values.join(', ')) : '—';
  const pkgLine = (name, version, repo) => [name, version ? 'v' + version : '', repo].filter(Boolean).map(esc).join(' · ') || '—';
  const catalog = [
    row('ID', esc(full.id), true), row(bi('name', '扩展名'), esc(full.name), true),
    row(bi('package family', '项目包族'), '<a href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a>', true),
    row(bi('lead extension', '主扩展'), full.lead_ext === full.name ? esc(full.name) : '<a href="' + extHref(full.lead_ext) + '">' + esc(full.lead_ext) + '</a>', true),
    row(bi('category', '分类'), '<a href="' + catHref(full.category) + '">' + esc(full.category) + ' · ' + esc(catName(full.category)) + '</a>'),
    row(bi('kind', '形态'), esc(full.kind || '—'), true), row(bi('lifecycle', '生命周期'), esc(full.lifecycle || '—'), true),
    row(bi('tags', '标签'), names(full.tags), true)
  ];
  const runtime = [
    row(bi('PG support', 'PG 支持'), pgBadgesHTML(full.pg_ver)), row(bi('schemas', '模式'), names(full.schemas), true),
    row(bi('libraries', '动态库'), names(full.libs), true), row('shared_preload_libraries', names(full.preload_libs), true),
    row(bi('ships executables', '携带可执行文件'), boolHTML(full.has_bin)), row(bi('ships libraries', '携带动态库'), boolHTML(full.has_lib)),
    row(bi('needs preload', '需要预加载'), boolHTML(full.need_load)), row('CREATE EXTENSION', boolHTML(full.need_ddl)),
    row(bi('trusted', '非超户可安装'), boolHTML(full.trusted)), row(bi('relocatable', '模式可迁移'), boolHTML(full.relocatable)),
    row('contrib', boolHTML(full.contrib))
  ];
  const packaging = [
    row('EXT', pkgLine(full.pkg, full.version, full.repo), true), row(bi('source archive', '源码归档'), esc(full.tarball || '—'), true),
    row('pgrx', esc(full.pgrx_ver || '—'), true),
    row('RPM', pkgLine(full.rpm_pkg, full.rpm_ver, full.rpm_repo), true), row('RPM PG', pgBadgesHTML(full.rpm_pg)),
    row(bi('RPM dependencies', 'RPM 依赖'), names(full.rpm_deps), true), row(bi('RPM build recipe', 'RPM 构建配方'), boolHTML(full.rpm_build)),
    row('DEB', pkgLine(full.deb_pkg, full.deb_ver, full.deb_repo), true), row('DEB PG', pgBadgesHTML(full.deb_pg)),
    row(bi('DEB dependencies', 'DEB 依赖'), names(full.deb_deps), true), row(bi('DEB build recipe', 'DEB 构建配方'), boolHTML(full.deb_build))
  ];
  const upstream = [
    row(bi('language', '实现语言'), esc(full.lang || '—'), true), row(bi('license', '许可证'), esc(full.license || '—'), true),
    row(bi('kernel', '专属内核'), esc(full.kernel || '—'), true), row(bi('vendor', '服务商'), esc(full.vendor || '—'), true),
    row('GitHub', ['★ ' + fmtInt(full.stars), '⑂ ' + fmtInt(full.forks), '👁 ' + fmtInt(full.watchers)].join(' · '), true),
    row(bi('repository created', '仓库创建'), esc(full.repo_created_at || '—'), true), row(bi('last release', '最近发布'), esc(full.last_release || '—'), true),
    row(bi('last commit', '最近提交'), esc(full.last_commit || '—'), true), row(bi('last active', '最近活跃'), esc(full.last_active || '—'), true),
    row(bi('checked at', '核验日期'), esc(full.checked_at || '—'), true), row(bi('catalog mtime', '目录更新时间'), esc(full.mtime || '—'), true)
  ];
  const extra = full.extra && typeof full.extra === 'object' && Object.keys(full.extra).length
    ? '<details class="extra-meta"><summary>' + t('spec.extra') + '</summary><pre>' + esc(JSON.stringify(full.extra, null, 2)) + '</pre></details>' : '';
  return '<div class="metadata-grid">' + panel(t('ext.catalog'), catalog) + panel(t('ext.runtime'), runtime)
    + panel(t('ext.packaging'), packaging) + panel(t('ext.upstream'), upstream) + '</div>' + extra;
}

function packageVersionsHTML(full) {
  const deps = values => (values || []).length ? values.map(x => '<a href="' + extHref(x) + '">' + esc(x) + '</a>').join(', ') : '—';
  const row = (type, repo, version, pgs, pattern, dependencies, recipe) => '<tr><td><b>' + type + '</b></td><td>' + esc(repo || '—') + '</td>'
    + '<td class="mono">' + esc(version || '—') + '</td><td>' + pgBadgesHTML(pgs) + '</td><td class="mono">' + esc(pattern || '—') + '</td>'
    + '<td>' + deps(dependencies) + '</td><td>' + (recipe == null ? '—' : boolHTML(recipe)) + '</td></tr>';
  let rows = row('EXT', full.repo, full.version, full.pg_ver, full.pkg, full.requires, null);
  if (!full.contrib && (full.rpm_pkg || full.rpm_ver || (full.rpm_pg || []).length)) rows += row('RPM', full.rpm_repo, full.rpm_ver, full.rpm_pg, full.rpm_pkg, full.rpm_deps, full.rpm_build);
  if (!full.contrib && (full.deb_pkg || full.deb_ver || (full.deb_pg || []).length)) rows += row('DEB', full.deb_repo, full.deb_ver, full.deb_pg, full.deb_pkg, full.deb_deps, full.deb_build);
  return '<div class="version-table"><div class="rows-scroll"><table><thead><tr><th>' + bi('type', '类型') + '</th><th>' + bi('repo', '仓库')
    + '</th><th>' + bi('version', '版本') + '</th><th>PG</th><th>' + bi('package pattern', '包名模式') + '</th><th>' + bi('dependencies', '依赖')
    + '</th><th>' + bi('recipe', '配方') + '</th></tr></thead><tbody>' + rows + '</tbody></table></div></div>';
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
      + '<div class="md-code"><div class="md-codebar"><span>bash</span><button data-copy="' + esc(command) + '">copy</button></div><pre><code>' + esc(command) + '</code></pre></div></div>';
  } else {
    guide = '<div class="build-note"><b>' + bi('Upstream build', '上游构建') + '</b><p>' + t('ext.buildnone') + '</p>'
      + (full.pgrx_ver ? '<span class="badge">pgrx · ' + esc(full.pgrx_ver) + '</span>' : '') + '</div>';
  }
  return (links.length ? '<div class="source-cards">' + links.join('') + '</div>' : '') + guide;
}

function relationLinks(names, role) {
  const values = names || [];
  if (!values.length) return '<span class="none">' + t('ext.none') + '</span>';
  return '<div class="relation-grid ' + role + '">' + values.map(name => {
    const e = byName.get(name);
    return '<a href="' + extHref(name) + '" ' + (e ? catVar(e.cat) : '') + '><code>' + esc(name) + '</code>'
      + (e ? '<span>' + esc(desc(e)) + '</span><small>' + esc(e.cat) + ' · ' + esc(e.kind || '') + '</small>' : '<span>' + bi('catalog reference', '目录引用') + '</span>') + '</a>';
  }).join('') + '</div>';
}

function depsHTML(full) {
  const requiredBy = full.required_by || full.require_by || [];
  const family = [full.name, ...(full.siblings || [])].map(name => byName.get(name)).filter(Boolean).sort((a, b) => a.name.localeCompare(b.name));
  const group = (title, names, role, hint) => '<section class="relation-group"><header><div><h3>' + title + '</h3><p>' + hint + '</p></div><span>' + (names || []).length + '</span></header>' + relationLinks(names, role) + '</section>';
  const familyRows = family.map(e => '<tr><td><a href="' + extHref(e.name) + '"><code>' + esc(e.name) + '</code></a>' + (e.name === full.lead_ext ? '<span class="lead-mark">lead</span>' : '') + '</td>'
    + '<td>' + esc(e.kind || '—') + '</td><td>' + boolHTML(Boolean(e.flags & F.LOAD)) + '</td><td>' + boolHTML(Boolean(e.flags & F.DDL)) + '</td><td>' + esc(desc(e)) + '</td></tr>').join('');
  const familyTable = family.length > 1 ? '<section class="relation-group family"><header><div><h3>' + t('ext.siblings') + '</h3><p><a href="' + pkgHref(full.pkg) + '"><code>' + esc(full.pkg) + '</code> · ' + family.length + ' ' + bi('extension definitions', '个扩展定义') + ' ↗</a></p></div><span>' + family.length + '</span></header>'
    + '<div class="version-table"><div class="rows-scroll"><table><thead><tr><th>' + bi('extension', '扩展') + '</th><th>' + bi('kind', '形态') + '</th><th>load</th><th>create</th><th>' + bi('description', '描述') + '</th></tr></thead><tbody>' + familyRows + '</tbody></table></div></div></section>' : '';
  return '<div class="relation-summary"><span><b>' + (full.requires || []).length + '</b>' + t('ext.requires') + '</span><i>→</i><span class="self"><b>' + esc(full.name) + '</b>' + esc(full.kind) + '</span><i>→</i><span><b>' + requiredBy.length + '</b>' + t('ext.requiredby') + '</span></div>'
    + '<div class="relation-groups">'
    + group(t('ext.requires'), full.requires || [], 'requires', bi('Installed first; CREATE EXTENSION may use CASCADE.', '需要先安装；CREATE EXTENSION 可使用 CASCADE。'))
    + group(t('ext.requiredby'), requiredBy, 'required-by', bi('Catalog entries that declare this extension as a dependency.', '在目录中声明依赖本扩展的条目。'))
    + group(t('ext.seealso'), full.see_also || [], 'see-also', bi('Curated alternatives and adjacent capabilities.', '人工整理的替代方案与相邻能力。'))
    + familyTable + '</div>';
}

function specHTML(e, full) {
  const links = full.doc_links || {};
  const resources = [
    ['link.home', full.home_url || links.home_url], ['link.repo', full.repo_url || links.repo_url],
    ['link.docs', full.doc_url || links.doc_url], ['link.pgxn', full.pgxn_url || links.pgxn_url],
    ['link.license', full.license_url || links.license_url], ['link.control', full.control_url || links.control_url],
    ['link.author', full.author_url || links.author_url], ['link.cargo', full.cargo_url || links.cargo_url]
  ].map(([key, url]) => [key, mdSafeURL(url)]).filter(([, url]) => url);
  const resourceHTML = resources.length ? '<div class="rail-card"><h3>' + t('ext.projectlinks') + '</h3><div class="resource-links">'
    + resources.map(([k, u]) => '<a href="' + esc(u) + '" target="_blank" rel="noopener"><span>' + t(k) + '</span><b>' + esc(u.replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</b><i>↗</i></a>').join('') + '</div></div>' : '';
  const freshness = [
    [bi('release', '发布'), full.last_release], [bi('commit', '提交'), full.last_commit], [bi('active', '活跃'), full.last_active],
    [bi('checked', '核验'), full.checked_at], [bi('catalog', '目录'), full.mtime]
  ].filter(([, value]) => value);
  const freshHTML = '<div class="rail-card"><h3>' + t('ext.freshness') + '</h3><dl class="fresh-list">'
    + freshness.map(([label, value]) => '<div><dt>' + label + '</dt><dd>' + esc(value) + '</dd></div>').join('') + '</dl>'
    + '<div class="rail-stats"><span>★ <b>' + fmtInt(full.stars) + '</b></span><span>⑂ <b>' + fmtInt(full.forks) + '</b></span><span>👁 <b>' + fmtInt(full.watchers) + '</b></span></div></div>';
  const identity = '<div class="rail-card identity"><h3>' + t('ext.catalog') + '</h3><code>' + esc(full.name) + '</code><p>' + esc(full.pkg) + ' · ' + esc(full.kind) + '</p>'
    + '<div class="rail-status"><span class="' + (full.packaged ? 'ok' : '') + '">' + (full.packaged ? '● ' + t('state.avail') : '○ ' + t('state.na')) + '</span><span>PG ' + esc(pgRange(full.pg_ver)) + '</span></div></div>';
  return identity + resourceHTML + freshHTML;
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
    fill('d-metadata', metadataHTML(full));
    fill('d-deps', depsHTML(full));
    fill('d-versions', packageVersionsHTML(full));
    fill('d-build', buildHTML(full));
    fill('d-spec', specHTML(e, full));
    if (full.repo_url || full.url) {
      const gh = document.getElementById('d-gh');
      const url = mdSafeURL(full.repo_url || full.url);
      if (url && gh && !gh.querySelector('.upstream-link')) gh.insertAdjacentHTML('beforeend', '<a class="upstream-link" href="' + esc(url) + '" target="_blank" rel="noopener">' + esc(url.replace(/^https?:\/\//, '')) + ' ↗</a>');
    }
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['d-overview', 'd-metadata', 'd-deps', 'd-versions', 'd-build', 'd-spec']) fill(id, hydrateErr(err));
  });

  matrixP.then(matrix => {
    if (tok !== hydSeq) return;
    fill('d-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, e) : '<p class="files-note">' + t('files.none') + '</p>');
  }).catch(err => { if (tok === hydSeq) fill('d-matrix', hydrateErr(err)); });

  Promise.all([fullP, matrixP]).then(([full, matrix]) => {
    if (tok === hydSeq) fill('d-install', installHTML(e, full, matrix));
  }).catch(err => { if (tok === hydSeq) fill('d-install', hydrateErr(err)); });

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
        const content = String(d.content || '').replace(/^\s*(?:#{1,2}\s+(?:usage|用法)\s*\n+|(?:usage|用法)\s*\n[-=]{3,}\s*\n+)/i, '');
        const rendered = renderMD(content, { usage: true });
        const outline = rendered.toc.length ? '<nav class="doc-toc" aria-label="' + esc(t('ext.docs')) + '"><span>' + t('ext.onpage') + '</span>'
          + '<div class="doc-toc-items">' + rendered.toc.map(h => '<button class="lv' + h.level + '" data-scroll="' + esc(h.id) + '">' + esc(h.title) + '</button>').join('') + '</div></nav>' : '';
        const meta = '<div class="doc-head"><div><b>' + t('ext.docsource') + '</b><span>'
          + t('ext.docmeta', { sections: rendered.toc.length, minutes: rendered.minutes, lang: lang === 'zh' ? '中文' : 'English' }) + '</span></div>'
          + '<span class="badge">' + (e.docbits === 3 ? 'EN + 中文' : (lang === 'zh' ? '中文' : 'EN')) + '</span></div>';
        fill('d-doc', note + meta + '<div class="doc-layout">' + outline + '<article class="prose usage-prose">' + rendered.html + '</article></div>');
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
function pkgOutlineHTML() {
  return manualOutlineHTML([
    ['pkg-overview', bi('Overview', '概览')],
    ['pkg-version', bi('Package definitions', '软件包定义')],
    ['pkg-availability', t('ext.avail')],
    ['pkg-downloads', t('ext.downloads')],
    ['pkg-build', t('ext.build')],
    ['pkg-install', t('ext.install')],
    ['pkg-extensions', t('ext.family')]
  ], 'pkg-outline');
}

function pkgOverviewHTML(pkg, members, full, matrix) {
  const cells = (matrix && matrix.cells || []).filter(c => c.state === 'AVAIL');
  const formats = [full.rpm_pkg || full.rpm_ver ? 'RPM' : '', full.deb_pkg || full.deb_ver ? 'DEB' : ''].filter(Boolean);
  if (full.contrib) formats.push('contrib');
  if (!formats.length) formats.push(bi('source', '源码'));
  const repos = [...new Set([full.rpm_repo, full.deb_repo, full.repo].filter(x => x && x !== 'n/a'))];
  const pgs = [...new Set(cells.map(c => c.pg))].sort((a, b) => b - a);
  const oss = new Set(cells.map(c => c.os));
  const fact = (label, value, note) => '<div class="pkg-fact"><span>' + label + '</span><strong>' + value + '</strong><small>' + note + '</small></div>';
  return '<div class="pkg-overview-copy"><p>' + esc(LANG === 'zh' ? (full.zh_desc || full.en_desc) : (full.en_desc || full.zh_desc)) + '</p></div>'
    + '<div class="pkg-facts">'
    + fact(bi('Extension family', '扩展族'), fmtInt(members.length), bi('extension definitions delivered by this project', '个由该项目交付的扩展定义'))
    + fact(bi('Distribution formats', '交付格式'), esc(formats.join(' + ')), esc(repos.join(' · ') || bi('upstream source', '上游源码')))
    + fact(bi('PostgreSQL coverage', 'PostgreSQL 覆盖'), esc(pgs.length ? pgRange(pgs) : pgRange(full.pg_ver)), pgs.length + bi(' packaged majors', ' 个已打包大版本'))
    + fact(bi('Binary targets', '二进制目标'), fmtInt(cells.length), fmtInt(oss.size) + bi(' OS / architecture targets', ' 个操作系统与架构目标'))
    + '</div>'
    + (full.tarball ? '<div class="pkg-source-note"><span>' + t('ext.sourcearchive') + '</span><code>' + esc(full.tarball) + '</code></div>' : '');
}

function pkgDefinitionsHTML(full) {
  const deps = values => (values || []).length ? esc(values.join(', ')) : '—';
  const row = (type, repo, version, pgs, pattern, dependencies, recipe) => '<tr><td><b>' + type + '</b></td><td>' + esc(repo || '—') + '</td>'
    + '<td class="mono">' + esc(version || '—') + '</td><td>' + pgBadgesHTML(pgs) + '</td><td class="mono">' + esc(pattern || '—') + '</td>'
    + '<td class="mono">' + deps(dependencies) + '</td><td>' + (recipe == null ? '—' : boolHTML(recipe)) + '</td></tr>';
  let rows = row('EXT', full.repo, full.version, full.pg_ver, full.tarball || full.pkg, [], null);
  if (!full.contrib && (full.rpm_pkg || full.rpm_ver || (full.rpm_pg || []).length)) rows += row('RPM', full.rpm_repo, full.rpm_ver, full.rpm_pg, full.rpm_pkg, full.rpm_deps, full.rpm_build);
  if (!full.contrib && (full.deb_pkg || full.deb_ver || (full.deb_pg || []).length)) rows += row('DEB', full.deb_repo, full.deb_ver, full.deb_pg, full.deb_pkg, full.deb_deps, full.deb_build);
  return '<div class="version-table pkg-definitions"><div class="rows-scroll"><table><thead><tr><th>' + bi('type', '类型') + '</th><th>' + bi('repo', '仓库')
    + '</th><th>' + bi('version', '版本') + '</th><th>PG</th><th>' + bi('package / source', '包名 / 源码') + '</th><th>' + bi('system dependencies', '系统依赖')
    + '</th><th>' + bi('build recipe', '构建配方') + '</th></tr></thead><tbody>' + rows + '</tbody></table></div></div>';
}

function pkgExtensionsHTML(members, full) {
  const lead = full.lead_ext || full.name;
  return '<div class="pkg-extension-grid">' + members.slice().sort((a, b) => a.name.localeCompare(b.name)).map(e =>
    '<a class="pkg-extension-card" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '><span class="pkg-extension-top"><code>' + esc(e.name) + '</code>'
    + (e.name === lead ? '<b>LEAD</b>' : '') + '</span><p>' + esc(desc(e)) + '</p><small><span>' + esc(e.cat) + '</span><span>' + esc(e.kind || '—') + '</span><span>PG ' + esc(pgRange(e.pg)) + '</span></small></a>'
  ).join('') + '</div>';
}

function packageInstallHTML(pkg, full, matrix, override) {
  const pref = installPrefs(matrix, override);
  const cell = pref.cells.find(c => c.pg === pref.pg && c.os === pref.os);
  const repoName = String((cell && cell.org) || full.repo || '').toUpperCase();
  const repoCmd = full.contrib || !full.packaged ? '' : (repoName === 'PGDG' ? 'pig repo add pgdg -u' : 'pig repo add pgsql -u');
  const word = value => /^[A-Za-z0-9_.+:-]+$/.test(String(value)) ? String(value) : shellArg(value);
  const tabs = [];
  if (cell) {
    const manager = cell.os.startsWith('el') ? 'dnf' : 'apt-get';
    const systemCmd = (manager === 'dnf' ? 'sudo dnf install -y ' : 'sudo apt-get install -y ') + word(cell.name);
    const pigCurrent = 'pig install ' + word(pkg);
    const pigExact = 'pig ext install -y ' + word(pkg) + ' -v ' + cell.pg;
    const target = 'PG ' + cell.pg + ' · ' + osLabel(cell.os) + ' · ' + (cell.org || 'repository') + ' · ' + (cell.version || full.version || '');
    const plan = ['# ' + target, repoCmd, systemCmd].filter(Boolean).join('\n');
    tabs.push(['install', esc(pigCurrent), pigCurrent]);
    tabs.push(['pig', '<span class="cmt"># ' + esc(target) + '</span>\n' + esc(pigExact), pigExact]);
    tabs.push([manager, '<span class="cmt"># ' + esc(target) + '</span>\n' + esc(systemCmd), systemCmd]);
    tabs.push(['plan', esc(plan), plan]);
  } else if (full.contrib) {
    tabs.push(['contrib', '<span class="cmt"># ' + esc(bi('Delivered by the matching PostgreSQL contrib/server package.', '由对应 PostgreSQL contrib / server 软件包交付。')) + '</span>', '']);
  } else if (full.packaged) {
    tabs.push(['unavailable', '<span class="cmt"># ' + esc(bi('No AVAIL package exists for this exact target. Choose another PG or OS above.', '该精确目标没有 AVAIL 软件包，请在上方改选 PG 或操作系统。')) + '</span>', '']);
  } else {
    const upstream = mdSafeURL(full.repo_url || full.doc_url || full.url);
    tabs.push(['source', '<span class="cmt"># ' + esc(bi('No public binary package is recorded. Build from the upstream source.', '没有公开二进制包记录，请从上游源码构建。')) + '</span>' + (upstream ? '\n' + esc(upstream) : ''), upstream]);
  }
  const pgs = matrix && matrix.pg || PGS;
  const oss = matrix && matrix.os || OSS;
  const controls = pref.cells.length ? '<div class="install-target pkg-install-target">'
    + '<label>PostgreSQL <select data-pkg-install-env="pg" data-install-pkg="' + esc(pkg) + '">' + pgs.map(pg => '<option value="' + pg + '"' + (pg === pref.pg ? ' selected' : '') + '>PG ' + pg + '</option>').join('') + '</select></label>'
    + '<label>' + bi('OS & architecture', '系统与架构') + ' <select data-pkg-install-env="os" data-install-pkg="' + esc(pkg) + '">' + oss.map(os => '<option value="' + esc(os) + '"' + (os === pref.os ? ' selected' : '') + '>' + esc(osLabel(os)) + '</option>').join('') + '</select></label>'
    + '<span class="target-state ' + (cell ? 'ok' : 'bad') + '">' + (cell ? '● AVAIL' : '○ MISS') + '</span></div>' : '';
  const heads = tabs.map(([name], i) => '<button role="tab" aria-selected="' + (i === 0) + '" data-itab="' + i + '">' + esc(name) + '</button>').join('');
  const panes = tabs.map(([, html, copy], i) => '<div class="install-body" data-ipane="' + i + '"' + (i ? ' hidden' : '') + '>'
    + (copy ? '<button class="copy-btn" data-copy="' + esc(copy) + '">copy</button>' : '') + '<pre>' + html + '</pre></div>').join('');
  const repo = repoCmd ? '<div class="pkg-repo-command"><span>' + bi('Repository setup', '软件仓库准备') + '</span><code>' + esc(repoCmd) + '</code><button data-copy="' + esc(repoCmd) + '">copy</button></div>' : '';
  return '<div class="pkg-install-panel">' + repo + controls + '<div class="install"><div class="install-tabs" role="tablist">' + heads + '</div>' + panes + '</div>'
    + '<p class="pkg-install-next">' + bi('After the package is installed, open an extension below for its preload and CREATE EXTENSION steps.', '软件包安装完成后，请从下方选择具体扩展，继续查看预加载与 CREATE EXTENSION 步骤。') + '</p></div>';
}

function pkgHTML(pkg) {
  const members = byPkg.get(pkg);
  if (!members || !members.length) return notFoundHTML(pkg);
  const lead = byName.get(members[0].lead) || members[0];
  const packaged = members.filter(e => e.avail).length;
  const categories = [...new Set(members.map(e => e.cat))];
  return '<article class="page wrap manual-page pkg-page">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span><span>' + bi('packages', '软件包') + '</span><span class="sep">/</span><span class="here">' + esc(pkg) + '</span></nav>'
    + '<header class="detail-hero pkg-detail-hero"><div class="detail-kicker">PACKAGE</div><div class="pkg-title-row"><h1><code>' + esc(pkg) + '</code></h1><span>' + members.length + bi(' extensions', ' 个扩展') + '</span></div><p class="lede">' + esc(desc(lead)) + '</p>'
    + '<div class="badge-row"><span class="badge">' + members.length + (LANG === 'zh' ? ' 个扩展' : ' extensions') + '</span>'
    + '<span class="badge flag-on">' + packaged + (LANG === 'zh' ? ' 个已打包' : ' packaged') + '</span>'
    + categories.map(c => '<a class="badge cat" href="' + catHref(c) + '" ' + catVar(c) + '><span class="dot"></span>' + esc(c) + '</a>').join('')
    + '<a class="badge" href="' + extHref(lead.name) + '">' + bi('lead extension', '主扩展') + ' · ' + esc(lead.name) + ' ↗</a><span id="p-upstream"></span></div></header>'
    + pkgOutlineHTML()
    + '<section class="section pkg-section" id="pkg-overview"><h2>' + bi('Package overview', '软件包概览') + '</h2><div id="p-overview">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-version"><h2>' + bi('Package definitions', '软件包定义') + '</h2><p class="section-lede">' + bi('Source, RPM and DEB definitions follow the same package model used by pgext gen io/cc.', '源码、RPM 与 DEB 定义沿用 pgext gen io/cc 的软件包模型。') + '</p><div id="p-definitions">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-availability"><h2>' + t('ext.avail') + '</h2><div id="p-matrix">' + skel(5) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-downloads"><h2>' + t('ext.downloads') + '</h2><div id="p-files">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-build"><h2>' + t('ext.build') + '</h2><div id="p-build">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-install"><h2>' + t('ext.install') + '</h2><p class="section-lede">' + bi('Choose an exact target or use pig for the active PostgreSQL installation.', '选择精确目标，或使用 pig 为当前活动 PostgreSQL 安装。') + '</p><div id="p-install">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-extensions"><h2>' + t('ext.family') + '</h2>' + pkgExtensionsHTML(members, lead) + '</section>'
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
  Promise.all([fullP, matrixP]).then(([full, matrix]) => {
    if (tok !== hydSeq) return;
    fill('p-overview', pkgOverviewHTML(pkg, members, full, matrix));
    fill('p-definitions', pkgDefinitionsHTML(full));
    fill('p-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, lead, { pkg, includeExtension: false }) : '<p class="empty-panel">' + t('files.none') + '</p>');
    fill('p-build', buildHTML(full));
    fill('p-install', packageInstallHTML(pkg, full, matrix));
    const upstream = mdSafeURL(full.repo_url || full.url);
    if (upstream) fill('p-upstream', '<a class="badge" href="' + esc(upstream) + '" target="_blank" rel="noopener">' + bi('upstream source', '上游源码') + ' ↗</a>');
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['p-overview', 'p-definitions', 'p-matrix', 'p-build', 'p-install']) fill(id, hydrateErr(err));
  });
  try {
    let files = FILEC.get(lead.name);
    if (!files) { files = await j('/api/v1/ext/' + enc + '/files'); FILEC.set(lead.name, files); }
    if (tok === hydSeq) fill('p-files', filesHTML(files));
  } catch (err) { if (tok === hydSeq) fill('p-files', hydrateErr(err)); }
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
    '<a href="' + catHref(x) + '" style="--seg:var(--c-' + x + ');flex:' + CATS[x].count + ' 1 0" aria-current="' + (x === code) + '" data-tip="' + x + ' · ' + CATS[x].count + '"></a>').join('');
  const fcards = featured.map(e =>
    '<li><a class="fcard" href="' + extHref(e.name) + '" ' + catVar(e.cat) + '>'
    + '<span class="name">' + esc(e.name) + '<span class="ver">' + esc(e.ver || '') + '</span>'
    + (e.stars ? '<span class="star">★ ' + fmtNum(e.stars) + '</span>' : '') + '</span>'
    + '<span class="desc">' + esc(desc(e)) + '</span>'
    + '<span class="meta"><span>' + esc(e.license) + '</span><span>' + esc(e.lang) + '</span><span>PG ' + esc(pgRange(e.pg)) + '</span>'
    + (e.docbits ? '<span>📖 docs</span>' : '') + '</span>'
    + '</a></li>').join('');
  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="/">' + t('ext.crumb') + '</a><span class="sep">/</span>'
    + '<a href="' + dimHref('category') + '">' + t('cat.crumb') + '</a><span class="sep">/</span><span class="here">' + code + '</span></nav>'
    + '<div class="cat-strip">' + strip + '</div>'
    + '<header class="page-head cat-hero" style="--seg:var(--c-' + code + ');margin-top:22px">'
    + '<span class="code">' + code + '</span>'
    + '<h1>' + esc(catName(code)) + '</h1>'
    + '<p class="lede">' + esc(catDesc(code)) + '</p>'
    + '<p class="cat-count" style="margin-top:8px">' + fmtInt(members.length) + ' extensions · '
    + fmtInt(members.filter(e => e.avail).length) + ' packaged · <a href="/?cat=' + code + '">' + t('cat.open') + '</a></p>'
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
  category: { key: 'cat', label: 'dim.category', d: 'dim.category.d', field: 'category' },
  tag: { key: 'tag', label: 'dim.tag', d: 'dim.tag.d', field: 'tags[]' },
  package: { key: 'pkg', label: 'dim.package', d: 'dim.package.d', field: 'pkg · lead_ext' },
  kind: { key: 'kind', label: 'dim.type', d: 'dim.type.d', field: 'kind' },
  lifecycle: { key: 'lifecycle', label: 'dim.lifecycle', d: 'dim.lifecycle.d', field: 'lifecycle' },
  license: { key: 'license', label: 'dim.license', d: 'dim.license.d', field: 'license' },
  lang: { key: 'lng', label: 'dim.lang', d: 'dim.lang.d', field: 'lang' },
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

const DIM_GROUPS = [
  ['browse.catalog', ['category', 'tag', 'package', 'kind', 'lifecycle', 'license', 'lang']],
  ['browse.delivery', ['distribution', 'repo', 'pg', 'os', 'build', 'pgrx']],
  ['browse.runtime', ['capability', 'docs', 'relation']],
  ['browse.ecosystem', ['vendor', 'kernel', 'activity']]
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

function dimValues(dim) {
  const cnt = new Map();
  const add = (k, e) => { if (!k) return; if (!cnt.has(k)) cnt.set(k, { n: 0, ex: [] }); const o = cnt.get(k); o.n++; if (o.ex.length < 4) o.ex.push(e.name); };
  for (const e of EXT) {
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
      const seg = dim === 'category' ? 'var(--c-' + v.v + ')' : 'var(--accent)';
      return '<i style="--w:' + v.n + ';--seg:' + seg + ';opacity:' + (0.25 + 0.7 * v.n / max).toFixed(2) + '"></i>';
    }).join('');
    return '<li><a class="dim-card" href="' + dimHref(dim) + '">'
      + '<span class="dim-field">pgext.universe · <code>' + esc(cfg.field) + '</code></span>'
      + '<span class="t">' + t(cfg.label) + '<span class="n">' + fmtInt(all.length) + '</span></span>'
      + '<span class="d">' + t(cfg.d) + '</span>'
      + '<span class="preview">' + preview + '</span></a></li>';
  };
  const groups = DIM_GROUPS.map(([title, dims]) => '<section class="dimension-group"><header><h2>' + t(title) + '</h2><span>' + dims.length + '</span></header>'
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
    const isCat = dim === 'category';
    const href = '/?' + cfg.key + '=' + encodeURIComponent(o.v);
    const seg = isCat ? 'style="--seg:var(--c-' + o.v + ')"' : '';
    const sample = o.ex.map(name => '<a href="' + extHref(name) + '">' + esc(name) + '</a>').join('');
    return '<tr ' + seg + ' data-dim-row data-dim-value="' + esc((o.v + ' ' + dimValueName(dim, o.v)).toLowerCase()) + '"><td class="v"><a href="' + href + '">' + dimValueHTML(dim, o.v) + '</a></td>'
      + '<td class="bar"><i style="width:' + Math.max(2, Math.round(100 * o.n / max)) + '%"></i></td>'
      + '<td class="n">' + fmtInt(o.n) + '</td>'
      + '<td class="ex">' + sample + '</td></tr>';
  }).join('');
  return '<article class="page wrap">'
    + '<nav class="crumbs"><a href="/browse">' + t('nav.browse') + '</a><span class="sep">/</span><span class="here">' + t(cfg.label) + '</span></nav>'
    + '<header class="page-head dimension-head"><span class="dim-source">pgext.universe · <code>' + esc(cfg.field) + '</code></span>'
    + '<h1>' + t(cfg.label) + '</h1><p class="lede">' + t(cfg.d) + '</p></header>'
    + '<div class="dim-toolbar"><label><span>⌕</span><input type="search" data-dim-search placeholder="' + esc(t('dim.search', { n: fmtInt(vals.length) })) + '" autocomplete="off"></label>'
    + '<span data-dim-count>' + t('dim.showing', { shown: fmtInt(vals.length), all: fmtInt(vals.length) }) + '</span></div>'
    + '<div class="dimtable"><table><thead><tr><th>' + t('dim.value') + '</th><th></th><th style="text-align:right">' + t('dim.count') + '</th><th>' + t('dim.sample') + '</th></tr></thead>'
    + '<tbody>' + rows + '</tbody></table></div></article>';
}

/* ---------------- view: global build matrix ---------------- */
const GMX_META = [
  ['B', 'gmx.pgdg', 'pgdg'],
  ['G', 'gmx.pigsty', 'pigsty'],
  ['R', 'gmx.missing', 'missing'],
  ['.', 'gmx.na', 'na']
];

function globalMatrixShellHTML() {
  return '<article class="gmx-page">'
    + '<header class="gmx-hero gmx-frame"><p class="eyebrow">' + t('gmx.eyebrow') + '</p>'
    + '<div class="gmx-titleline"><div><h1>' + t('gmx.title') + '</h1><p>' + t('gmx.lede') + '</p></div>'
    + '<a class="gmx-api" href="/api/v1/matrix" target="_blank" rel="noopener">' + t('gmx.api') + ' ↗</a></div></header>'
    + '<div class="gmx-frame" id="gmx-root">' + skel(7) + '</div></article>';
}

// Decode one positional cell lazily. Available-cell details use indexes into
// the row-local name/version dictionaries: [nameIndex, versionIndex, count].
function globalMatrixCell(row, index) {
  const code = (row.c || '').charAt(index) || '.';
  const detail = Array.isArray(row.d) ? row.d[index] : null;
  let state = '', org = '';
  if (code === 'B') { state = 'AVAIL'; org = 'PGDG'; }
  else if (code === 'G') { state = 'AVAIL'; org = 'PIGSTY'; }
  else if (code === 'R') state = 'MISS';
  else if (code === '.') state = 'N/A';
  if (detail) {
    state = detail[3] || state;
    org = detail[4] || org;
  }
  return {
    code, state, org,
    name: detail && detail[0] >= 0 ? (row.n[detail[0]] || '') : '',
    version: detail && detail[1] >= 0 ? (row.v[detail[1]] || '') : '',
    count: detail ? detail[2] : 0
  };
}

function globalMatrixStatusLabel(cell) {
  if (!cell) return '—';
  const item = GMX_META.find(x => x[0] === cell.code);
  return item ? t(item[1]) : 'N/A';
}

function globalMatrixClass(code) {
  const item = GMX_META.find(x => x[0] === code);
  return item ? item[2] : 'na';
}

function globalMatrixHTML(data) {
  const stats = data.stats || { rows: 0, os: 0, pg: 0, cells: 0, counts: {} };
  const statItems = [
    [stats.rows, 'gmx.packages'], [stats.os, 'gmx.targets'],
    [stats.pg, 'gmx.pg'], [stats.cells, 'gmx.cells']
  ].map(x => '<div><strong>' + fmtInt(x[0]) + '</strong><span>' + t(x[1]) + '</span></div>').join('');
  const legend = GMX_META.map(item => {
    const count = (stats.counts && stats.counts[item[0]]) || 0;
    return '<button type="button" class="gmx-legend-item gmx-' + item[2] + '" data-gmx-code="' + item[0]
      + '" aria-pressed="false"><i></i><span>' + t(item[1]) + '</span><b>' + fmtInt(count) + '</b></button>';
  }).join('');
  const cols = (data.os || []).length * (data.pg || []).length;
  const osHeads = (data.os || []).map(osName => {
    const bits = osName.split('.');
    return '<span class="gmx-oshead" style="grid-column:span ' + (data.pg || []).length + '"><b>'
      + esc(bits[0]) + '</b><small>' + esc(bits.slice(1).join('.') || '') + '</small></span>';
  }).join('');
  let pgHeads = '';
  for (const osName of (data.os || [])) {
    for (let i = 0; i < (data.pg || []).length; i++) {
      pgHeads += '<span class="gmx-pghead' + (i === 0 ? ' group-start' : '') + '">' + esc(data.pg[i]) + '</span>';
    }
  }
  return '<section class="gmx-stats" aria-label="Matrix summary">' + statItems + '</section>'
    + '<section class="gmx-panel">'
    + '<div class="gmx-toolbar"><label class="gmx-search"><span>\\</span><input id="gmx-q" type="search" autocomplete="off" spellcheck="false" placeholder="'
    + esc(t('gmx.search')) + '"></label><span class="gmx-showing" id="gmx-showing"></span></div>'
    + '<div class="gmx-legend" aria-label="Matrix status filters">' + legend + '</div>'
    + '<div class="gmx-hint">' + t('gmx.hint') + '</div>'
    + '<div class="gmx-viewport" id="gmx-viewport" role="grid" aria-rowcount="' + stats.rows + '" aria-colcount="' + (cols + 1) + '">'
    + '<div class="gmx-stage" style="--gmx-columns:' + cols + '">'
    + '<div class="gmx-grid-head"><div class="gmx-headrow"><span class="gmx-corner">' + t('gmx.pkg') + '</span>' + osHeads + '</div>'
    + '<div class="gmx-headrow gmx-pg-row"><span class="gmx-corner gmx-corner-sub">OS → · PG →</span>' + pgHeads + '</div></div>'
    + '<div class="gmx-grid-body" id="gmx-body"></div></div></div>'
    + '<p class="gmx-source"><span class="gmx-source-copy">' + t('gmx.source', { source: '<code>' + esc(data.source || 'pgext.matrix') + '</code>' })
    + '</span><span class="gmx-snapshot">snapshot ' + esc((data.generated || '').replace('T', ' ').slice(0, 19)) + '</span></p>'
    + '</section>';
}

function setupGlobalMatrix(data) {
  const viewport = document.getElementById('gmx-viewport');
  const body = document.getElementById('gmx-body');
  const input = document.getElementById('gmx-q');
  const showing = document.getElementById('gmx-showing');
  if (!viewport || !body || !input) return;
  const matrixPage = document.querySelector('.gmx-page');
  const rowHeight = Number.parseFloat(getComputedStyle(matrixPage).getPropertyValue('--gmx-row')) || 20;
  const cellCount = (data.os || []).length * (data.pg || []).length;
  const overscan = 10;
  const state = { rows: [], activeCode: '', start: -1, end: -1, raf: 0 };

  function render(force) {
    if (!viewport.isConnected) return;
    const bodyY = Math.max(0, viewport.scrollTop - body.offsetTop);
    const start = Math.max(0, Math.floor(bodyY / rowHeight) - overscan);
    const visible = Math.ceil(viewport.clientHeight / rowHeight) + overscan * 2;
    const end = Math.min(state.rows.length, start + visible);
    if (!force && start === state.start && end === state.end) return;
    state.start = start; state.end = end;
    body.style.height = (state.rows.length * rowHeight) + 'px';
    if (!state.rows.length) {
      body.innerHTML = '<div class="gmx-empty">' + t('gmx.empty') + '</div>';
      return;
    }
    const html = [];
    const pgCount = (data.pg || []).length;
    for (let vi = start; vi < end; vi++) {
      const ref = state.rows[vi], row = ref.row;
      let cells = '';
      for (let ci = 0; ci < cellCount; ci++) {
        const cell = globalMatrixCell(row, ci);
        const pi = ci % pgCount;
        const oi = Math.floor(ci / pgCount);
        const aria = row.p + ' · ' + data.os[oi] + ' · PG ' + data.pg[pi] + ' · ' + globalMatrixStatusLabel(cell);
        cells += '<span class="gmx-cell gmx-' + globalMatrixClass(cell.code) + (pi === 0 ? ' group-start' : '')
          + '" role="gridcell" data-gmx-row="' + ref.index + '" data-gmx-cell="' + ci + '" aria-label="' + esc(aria) + '"></span>';
      }
      const ext = row.e && row.e !== row.p ? '<small>' + esc(row.e) + '</small>' : '';
      html.push('<div class="gmx-row" role="row" style="top:' + (vi * rowHeight) + 'px"><a class="gmx-row-name" href="'
        + pkgHref(row.p) + '"><b>' + esc(row.p) + '</b>' + ext + '</a>' + cells + '</div>');
    }
    body.innerHTML = html.join('');
  }

  function schedule(force) {
    if (state.raf) cancelAnimationFrame(state.raf);
    state.raf = requestAnimationFrame(() => { state.raf = 0; render(force); });
  }

  function applyFilter() {
    const query = input.value.trim().toLowerCase();
    state.rows = data.rows.map((row, index) => ({ row, index })).filter(ref => {
      const row = ref.row;
      if (query && !(row.p + ' ' + row.e).toLowerCase().includes(query)) return false;
      return !state.activeCode || (row.c || '').includes(state.activeCode);
    });
    showing.textContent = t('gmx.showing', {
      rows: fmtInt(state.rows.length), cells: fmtInt(state.rows.length * data.os.length * data.pg.length)
    });
    document.querySelectorAll('[data-gmx-code]').forEach(button => {
      button.setAttribute('aria-pressed', button.dataset.gmxCode === state.activeCode ? 'true' : 'false');
    });
    viewport.scrollTop = 0;
    state.start = state.end = -1;
    schedule(true);
  }

  viewport.addEventListener('scroll', () => { hideTip(); schedule(false); }, { passive: true });
  viewport.addEventListener('click', ev => {
    const cell = ev.target.closest('.gmx-cell');
    if (!cell) return;
    const row = data.rows[Number.parseInt(cell.dataset.gmxRow, 10)];
    if (row) navigateTo(pkgHref(row.p));
  });
  input.addEventListener('input', debounce(applyFilter, 90));
  document.querySelectorAll('[data-gmx-code]').forEach(button => {
    button.addEventListener('click', () => {
      state.activeCode = state.activeCode === button.dataset.gmxCode ? '' : button.dataset.gmxCode;
      applyFilter();
    });
  });
  GMATRIX_VIEW = { render: () => schedule(true), state };
  applyFilter();
}

async function hydrateGlobalMatrix() {
  const token = ++matrixHydSeq;
  try {
    if (!GMATRIX) GMATRIX = await j('/api/v1/matrix');
    if (token !== matrixHydSeq || currentPath !== '/matrix') return;
    const root = document.getElementById('gmx-root');
    if (!root) return;
    root.innerHTML = globalMatrixHTML(GMATRIX);
    setupGlobalMatrix(GMATRIX);
  } catch (err) {
    if (token === matrixHydSeq) {
      const root = document.getElementById('gmx-root');
      if (root) root.innerHTML = hydrateErr(err);
    }
  }
}

function globalMatrixCellInfo(el) {
  if (!GMATRIX || !el) return null;
  const ri = Number.parseInt(el.dataset.gmxRow, 10);
  const ci = Number.parseInt(el.dataset.gmxCell, 10);
  const row = GMATRIX.rows[ri];
  if (!row || ci < 0 || ci >= (row.c || '').length) return null;
  const pgCount = GMATRIX.pg.length;
  return { row, cell: globalMatrixCell(row, ci), os: GMATRIX.os[Math.floor(ci / pgCount)], pg: GMATRIX.pg[ci % pgCount] };
}

function globalMatrixTipHTML(info) {
  const cell = info.cell;
  let details = '<span class="d">' + esc(info.row.e) + ' · ' + esc(info.os) + ' · PG ' + esc(info.pg) + '</span>';
  if (cell.name) details += '<br><span class="k">package</span> ' + esc(cell.name);
  if (cell.org) details += '<br><span class="k">repo</span> ' + esc(cell.org);
  if (cell.version) details += '<br><span class="k">version</span> ' + esc(cell.version);
  if (cell.count) details += '<br><span class="k">artifacts</span> ' + fmtInt(cell.count);
  return '<b>' + esc(info.row.p) + '</b> <span class="gmx-tip-state gmx-' + globalMatrixClass(cell.code) + '">'
    + esc(globalMatrixStatusLabel(cell)) + '</span><br>' + details;
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
    + '<li><span class="tag">GET</span><code>/api/v1/matrix</code></li>'
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
  const { path, params } = parseRoute();
  if (path.startsWith('/ext/')) { navigateTo('/e/' + path.slice(5) + location.search, true); return; }
  if (path.startsWith('/pkg/')) { navigateTo('/p/' + path.slice(5) + location.search, true); return; }
  if (path.startsWith('/cat/')) { navigateTo('/c/' + path.slice(5) + location.search, true); return; }
  const app = document.getElementById('app');
  const nav = document.getElementById('nav');
  const pathChanged = path !== currentPath;
  currentPath = path;
  if (path !== '/matrix') { GMATRIX_VIEW = null; matrixHydSeq++; }
  let active = 'home';
  if (path === '/' || path === '') {
    readState(params);
    app.innerHTML = homeHTML();
    drawField();
    bindSearch();
  } else if (path.startsWith('/e/')) {
    const name = decodeURIComponent(path.slice(3));
    app.innerHTML = extHTML(name);
    hydrateExt(name);
    active = '';
  } else if (path.startsWith('/p/')) {
    const pkg = decodeURIComponent(path.slice(3));
    app.innerHTML = pkgHTML(pkg);
    hydratePkg(pkg);
    active = '';
  } else if (path.startsWith('/c/')) {
    app.innerHTML = catHTML(decodeURIComponent(path.slice(3)).toUpperCase()); active = '';
  } else if (path === '/matrix') {
    app.innerHTML = globalMatrixShellHTML(); active = 'matrix';
    hydrateGlobalMatrix();
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
  const meta = document.querySelector('meta[name="description"]');
  const ogTitle = document.querySelector('meta[property="og:title"]');
  const ogDesc = document.querySelector('meta[property="og:description"]');
  let description = LANG === 'zh' ? 'PostgreSQL 扩展、项目包族、依赖关系与 PG/OS 精确可用性目录。' : 'Search PostgreSQL extensions, package families, dependencies, and exact PG/OS availability.';
  if (path.startsWith('/e/')) { const e = byName.get(decodeURIComponent(path.slice(3))); if (e) description = desc(e); }
  if (path.startsWith('/p/')) { const family = byPkg.get(decodeURIComponent(path.slice(3))); if (family && family[0]) description = desc(byName.get(family[0].lead) || family[0]); }
  if (path === '/matrix') description = LANG === 'zh' ? '跨扩展包、操作系统与 PostgreSQL 大版本的 32,000 格全局构建矩阵。' : 'A 32,000-cell global build matrix across extension packages, operating systems, and PostgreSQL majors.';
  if (meta) meta.content = description;
  if (ogTitle) ogTitle.content = document.title;
  if (ogDesc) ogDesc.content = description;
}
function titleFor(path) {
  if (path.startsWith('/e/')) { const n = decodeURIComponent(path.slice(3)); return n + ' · PGEXT.CLOUD'; }
  if (path.startsWith('/p/')) { const n = decodeURIComponent(path.slice(3)); return n + ' package · PGEXT.CLOUD'; }
  if (path.startsWith('/c/')) return decodeURIComponent(path.slice(3)).toUpperCase() + ' · PGEXT.CLOUD';
  if (path === '/matrix') return t('gmx.title') + ' · PGEXT.CLOUD';
  if (path === '/browse') return t('nav.browse') + ' · PGEXT.CLOUD';
  if (path.startsWith('/dim/')) {
    const dim = decodeURIComponent(path.slice(5));
    return (DIMS[dim] ? t(DIMS[dim].label) : dim) + ' · PGEXT.CLOUD';
  }
  if (path === '/about') return t('nav.about') + ' · PGEXT.CLOUD';
  return 'PGEXT.CLOUD — The PostgreSQL Extension Catalog';
}

function renderDynamic() {
  const d = document.getElementById('dynamic');
  if (d) { d.innerHTML = dynamicHTML(); pushState(); }
}
function rerenderInstall(name, override) {
  const e = byName.get(name), full = FULLC.get(name), matrix = MXC.get(name);
  if (!e || !full || !matrix) return;
  const pgSelect = document.querySelector('[data-install-env="pg"][data-install-ext="' + CSS.escape(name) + '"]');
  const osSelect = document.querySelector('[data-install-env="os"][data-install-ext="' + CSS.escape(name) + '"]');
  const selected = { pg: pgSelect ? pgSelect.value : '', os: osSelect ? osSelect.value : '', ...(override || {}) };
  if (selected.pg) INSTALL_PREF.pg = selected.pg;
  if (selected.os) INSTALL_PREF.os = selected.os;
  fill('d-install', installHTML(e, full, matrix, selected));
}
function rerenderPkgInstall(pkg, override) {
  const members = byPkg.get(pkg);
  if (!members || !members.length) return;
  const lead = byName.get(members[0].lead) || members[0];
  const full = FULLC.get(lead.name), matrix = MXC.get(lead.name);
  if (!full || !matrix) return;
  const pgSelect = document.querySelector('[data-pkg-install-env="pg"][data-install-pkg="' + CSS.escape(pkg) + '"]');
  const osSelect = document.querySelector('[data-pkg-install-env="os"][data-install-pkg="' + CSS.escape(pkg) + '"]');
  const selected = { pg: pgSelect ? pgSelect.value : '', os: osSelect ? osSelect.value : '', ...(override || {}) };
  if (selected.pg) INSTALL_PREF.pg = selected.pg;
  if (selected.os) INSTALL_PREF.os = selected.os;
  fill('p-install', packageInstallHTML(pkg, full, matrix, selected));
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
    const el = ev.target.closest('[data-fkey],[data-skey],[data-pg-toggle],[data-entity],[data-layout],[data-sql],[data-copy],[data-scroll],[data-itab],[data-ftab],[data-fall],[data-cat-go],[data-target-ext],[data-target-pkg],#lang-toggle,#theme-toggle,#ufield');
    if (!el) return;
    if (el.id === 'lang-toggle') return toggleLang();
    if (el.id === 'theme-toggle') return cycleTheme();
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
    const pkgEnv = ev.target.closest('[data-pkg-install-env]');
    if (pkgEnv) {
      INSTALL_PREF[pkgEnv.dataset.pkgInstallEnv] = pkgEnv.value;
      try { localStorage.setItem('pgext.target.' + pkgEnv.dataset.pkgInstallEnv, pkgEnv.value); } catch (err) {}
      rerenderPkgInstall(pkgEnv.dataset.installPkg, { [pkgEnv.dataset.pkgInstallEnv]: pkgEnv.value });
      return;
    }
    const env = ev.target.closest('[data-install-env]');
    if (env) {
      INSTALL_PREF[env.dataset.installEnv] = env.value;
      try { localStorage.setItem('pgext.target.' + env.dataset.installEnv, env.value); } catch (err) {}
      rerenderInstall(env.dataset.installExt, { [env.dataset.installEnv]: env.value });
      return;
    }
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
      if (path !== '/' && path !== '') { navigateTo('/'); setTimeout(() => { const q = document.getElementById('q'); if (q) q.focus(); }, 60); }
      else { const q = document.getElementById('q'); if (q) q.focus(); }
    }
  });

  document.addEventListener('mousemove', ev => {
    const matrixCell = ev.target.closest && ev.target.closest('.gmx-cell');
    if (matrixCell) {
      const info = globalMatrixCellInfo(matrixCell);
      if (info) showTip(globalMatrixTipHTML(info), ev.clientX, ev.clientY);
      return;
    }
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
