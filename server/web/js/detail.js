/* ================================================================
   Extension (/ext/{name}) and Package (/pkg/{pkg}) detail views
   and lazy hydration (single detail fetch — no N+1 sibling waterfall).
   ================================================================ */
import { esc, fmtNum, fmtInt, fmtSize, pgRange, extHref, pkgHref, catHref, dimHref, catVar, licenseHue, langHue, repoHue, hueVar, sqlIdent, shellArg } from './util.js';
import { getLang, t, bi, catName, desc } from './i18n.js';
import { mdSafeURL, renderMD, mdCodeHTML } from './markdown.js';
import { FULLC, MXC, FILEC, DOCC, CHLOGC, catalog, j, fill } from './state.js';
import { ICON_STAR, ICON_FORK, ICON_EYE, ICON_OUT, ICON_LINK_LG, GH_SVG_LG, hueTag, attrChips, pageTocHTML, skel, hydrateErr, notFoundHTML } from './components.js';

let hydSeq = 0;

/* ---------------- shared hero & table fragments ---------------- */
const labeledBadge = (label, href, value, hue) =>
  '<a class="badge lbl' + (hue ? ' hued' : '') + '" ' + (hue ? hueVar(hue) + ' ' : '') + 'href="' + esc(href) + '">'
  + '<span class="bk">' + label + (getLang() === 'zh' ? '：' : ':') + '</span><span class="bv">' + esc(value) + '</span></a>';

export function heroDimBadgesHTML(e) {
  return '<a class="badge cat lbl" href="' + catHref(e.cat) + '" ' + catVar(e.cat) + ' data-tip="' + esc(catName(e.cat)) + '"><span class="bk">' + t('spec.category')
    + (getLang() === 'zh' ? '：' : ':') + '</span><span class="bv">' + e.cat + '</span></a>'
    + (e.license !== 'Unknown' ? labeledBadge(t('spec.license'), '/license/' + encodeURIComponent(e.license), e.license, licenseHue(e.license)) : '')
    + (e.lang ? labeledBadge(t('spec.language'), '/lang/' + encodeURIComponent(e.lang), e.lang, langHue(e.lang)) : '');
}

export function ghCardHTML(e, full, fallbackName) {
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

export function heroDatesHTML(full) {
  const rows = [
    [bi('Repo Created', '仓库创建'), full.repo_created_at],
    [bi('Last Release', '最近发布'), full.last_release],
    [bi('Last Commit', '最近提交'), full.last_commit],
    [bi('Last Modified', '最近更新'), full.last_active]
  ].filter(row => row[1]);
  if (!rows.length) return '';
  return '<dl class="hero-dates">' + rows.map(([k, v]) => '<div><dt>' + k + '</dt><dd>' + esc(String(v).slice(0, 10)) + '</dd></div>').join('') + '</dl>';
}

export function heroSideMainHTML(e, full, fallbackName) {
  return ghCardHTML(e, full, fallbackName) + (full ? heroDatesHTML(full) : '');
}

export function heroLinkBadgesHTML(full) {
  return (full.tags || []).slice(0, 10).map(tag =>
    '<a class="tg hero-tag" href="/?tag=' + encodeURIComponent(tag) + '">' + esc(tag) + '</a>').join('');
}

const boolHTML = value => value
  ? '<span class="bool yes">✓ ' + bi('yes', '是') + '</span>'
  : '<span class="bool no">— ' + bi('no', '否') + '</span>';

function pgBadgesHTML(values) {
  const pgs = (values || []).map(Number).filter(Boolean).sort((a, b) => b - a);
  return pgs.length ? '<span class="pg-badges">' + pgs.map(pg => '<a href="/pg/' + pg + '">PG' + pg + '</a>').join('') + '</span>' : '—';
}

function pgMajorsCellHTML(supported) {
  const on = new Set((supported || []).map(Number));
  return '<span class="pgm-cell">' + catalog.pgs.map(pg =>
    '<a class="pgm ' + (on.has(pg) ? 'on' : 'off') + '" href="/pg/' + pg + '">' + pg + '</a>').join('') + '</span>';
}

function depsCellHTML(values, linked) {
  const list = (values || []).map(v => linked && catalog.byName.has(v) ? '<a href="' + extHref(v) + '">' + esc(v) + '</a>' : esc(v));
  return list.length ? list.join(list.length >= 2 ? '<br>' : ', ') : '—';
}

function repoCellHTML(repo) {
  if (!repo || repo === 'n/a') return '—';
  const upper = String(repo).toUpperCase();
  return hueTag('/repo/' + encodeURIComponent(upper), upper, repoHue(upper));
}

const buildCellHTML = recipe => recipe == null ? '—'
  : recipe ? '<span class="bool yes">✓ ' + bi('Yes', '支持') + '</span>'
    : '<span class="bool no">— ' + bi('Miss', '暂无') + '</span>';

/* ---------------- install & build ---------------- */
export function packageTabsHTML(full, fallbackPkg) {
  const word = value => /^[A-Za-z0-9_.+:*$-]+$/.test(String(value)) ? String(value) : shellArg(value);
  const align = rows => { const w = Math.max(...rows.map(([c]) => c.length)); return rows.map(([c, n]) => (c + ';').padEnd(w + 4) + '# ' + n).join('\n'); };
  const majors = list => [...new Set(list || [])].filter(pg => !catalog.pgs.length || catalog.pgs.includes(pg)).sort((a, b) => b - a);
  const pkgName = full.pkg || fallbackPkg || full.name;
  const pigPGs = majors([...(full.rpm_pg || []), ...(full.deb_pg || [])]);
  const pigText = align([['pig install ' + word(pkgName), bi('for the active PG major', '为当前活跃 PG 大版本安装')]]
    .concat(pigPGs.map(pg => ['pig install ' + word(pkgName) + ' -v ' + pg, bi('install for PG ', '为 PG ') + pg + bi('', ' 安装')])));
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

export function installHTML(e, full) {
  const ident = sqlIdent(e.name);
  const requires = full.requires || [];
  const cascade = requires.length ? ' CASCADE' : '';
  const createSQL = full.need_ddl ? 'CREATE EXTENSION ' + ident + cascade + ';' : '';
  const libs = (full.preload_libs || []).length ? full.preload_libs : ((full.libs || []).length ? full.libs : [e.name]);
  const repoName = String(full.repo || '').toUpperCase();
  const pgdgOnly = repoName === 'PGDG' && ![full.rpm_repo, full.deb_repo].some(repo => String(repo || '').toUpperCase() === 'PIGSTY');
  const repoCmd = full.contrib || !full.packaged ? '' : (pgdgOnly ? 'pig repo add pgdg -u' : 'pig repo add pgsql -u');
  const sharedPreload = full.need_load && full.kind === 'preload';
  const loadConfig = sharedPreload ? "shared_preload_libraries = '" + libs.join(', ') + "'" : '';

  const upstream = mdSafeURL(full.repo_url || full.doc_url || full.url);
  const packageBody = full.contrib
    ? '<p>' + bi('Included with the matching PostgreSQL contrib/server packages — no separate extension package is required.', '随对应版本的 PostgreSQL contrib / 服务端软件包一并交付，无需安装单独的扩展软件包。') + '</p>'
    : full.packaged
      ? packageTabsHTML(full, e.pkg || e.name)
      : '<p>' + bi('No public binary package is recorded. Follow the upstream build or provider instructions.', '没有公开二进制软件包记录，请遵循上游构建或服务商说明。')
        + (upstream ? ' <a href="' + esc(upstream) + '" target="_blank" rel="noopener">' + esc(upstream.replace(/^https?:\/\//, '')) + ' ↗</a>' : '') + '</p>';
  const repoBody = full.contrib
    ? '<p>' + bi('Bundled with the matching PostgreSQL distribution; no separate repository setup is required.', '随对应 PostgreSQL 发行包交付，无需单独配置扩展仓库。') + '</p>'
    : repoCmd ? '<p>' + bi('Enable and refresh the repository selected by the catalog before installing.', '安装前启用并刷新目录所选的软件仓库。') + '</p>' + mdCodeHTML('bash', repoCmd)
      : '<p>' + bi('No public package repository is recorded. Use the build manual and upstream source instead.', '没有公开软件仓库记录，请改用构建手册与上游源码。') + '</p>';
  const loadBody = sharedPreload
    ? '<p>' + bi('Append these libraries to the existing setting, then restart PostgreSQL. Do not overwrite other preload entries.', '将下列库追加到现有配置并重启 PostgreSQL；不要覆盖已有预加载项。') + '</p>' + mdCodeHTML('postgresql.conf', loadConfig)
    : full.need_load
      ? '<p>' + bi('An explicit loading or enablement step is required. Follow the ', '需要显式加载或启用。请参考')
        + '<a href="#ext-usage">' + t('ext.docs') + '</a>'
        + bi(' guide for the exact configuration and restart requirements.', '文档中的具体配置与重启要求。') + '</p>'
      : '<p class="step-done">✓ ' + t('ext.noload') + '</p>';
  const enableBody = createSQL
    ? '<p>' + (requires.length ? bi('CASCADE also creates declared extension dependencies: ', 'CASCADE 会同时创建声明的扩展依赖：') + '<code>' + esc(requires.join(', ')) + '</code>' : bi('Create the SQL objects in each target database.', '在每一个目标数据库中创建扩展 SQL 对象。')) + '</p>' + mdCodeHTML('sql', createSQL)
    : '<p class="step-done">✓ ' + t('ext.noddl') + '</p>';
  return '<div class="install-steps">'
    + '<section class="install-step"><h3>' + t('ext.step.repo') + '</h3>' + repoBody + '</section>'
    + '<section class="install-step package-step"><h3>' + t('ext.step.package') + '</h3>' + packageBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.load') + '</h3>' + loadBody + '</section>'
    + '<section class="install-step"><h3>' + t('ext.step.enable') + '</h3>' + enableBody + '</section></div>';
}

export function sourceArchiveURL(source) {
  if (!source) return '';
  if (/^https?:\/\//i.test(source)) return source;
  const path = String(source).replace(/^\/+/, '').split('/').map(encodeURIComponent).join('/');
  return (getLang() === 'zh' ? 'https://repo.pigsty.cc/ext/src/' : 'https://repo.pigsty.io/ext/src/') + path;
}

export function buildHTML(full) {
  if (full.contrib) return '<div class="build-note contrib"><b>PostgreSQL contrib</b><p>' + t('ext.contribbuild') + '</p>' + pgBadgesHTML(full.pg_ver) + '</div>';
  const types = [full.rpm_build ? 'RPM' : '', full.deb_build ? 'DEB' : ''].filter(Boolean);
  const sources = [...new Set(String(full.tarball || '').trim().split(/\s+/).filter(Boolean))];
  const upstream = mdSafeURL(full.repo_url || full.url);
  const docs = mdSafeURL(full.doc_url);
  const links = [];
  if (upstream) links.push('<a href="' + esc(upstream) + '" target="_blank" rel="noopener"><span>' + bi('Upstream source', '上游源码') + '</span><b>' + esc(upstream.replace(/^https?:\/\//, '')) + ' ↗</b></a>');
  for (const source of sources) links.push('<a href="' + esc(sourceArchiveURL(source)) + '" target="_blank" rel="noopener"><span>' + t('ext.sourcearchive') + '</span><b>' + esc(source.split('/').pop()) + ' ↓</b></a>');
  if (docs) links.push('<a href="' + esc(docs) + '" target="_blank" rel="noopener"><span>' + bi('Upstream build docs', '上游构建文档') + '</span><b>' + esc(docs.replace(/^https?:\/\//, '')) + ' ↗</b></a>');
  let guide = '';
  if (types.length) {
    const name = full.pkg || full.name;
    const pkg = /^[A-Za-z0-9_][A-Za-z0-9_.+-]*$/.test(name) ? name : shellArg(name);
    const setup = ['pgext env repo', 'pgext env tool', 'pgext env spec'];
    if (full.pgrx_ver) {
      setup.push('pgext env rust');
      if (/^\d+(\.\d+)*$/.test(full.pgrx_ver)) setup.push('pgext env pgrx --pgrx ' + full.pgrx_ver);
    }
    guide = '<div class="build-recipe"><p>' + t('ext.buildrecipe', { types: '<b>' + types.join(' / ') + '</b>' }) + '</p>'
      + '<p>' + bi('Prepare the builder once (system packages require root or sudo):', '首次准备构建环境（系统包操作需要 root 或 sudo）：') + '</p>'
      + mdCodeHTML('bash', setup.join('\n'))
      + (full.pgrx_ver ? '<p>pgrx · <code>' + esc(full.pgrx_ver) + '</code> — ' + bi('use the toolchain required by the recipe.', '使用配方要求的工具链。') + '</p>' : '')
      + '<p>' + bi('Preview, then run get → dep → mak:', '先预览，再执行 get → dep → mak：') + '</p>'
      + mdCodeHTML('bash', 'pgext pkg ' + pkg + ' --plan\npgext pkg ' + pkg)
      + '<p>' + bi('Targets come from the recipe and installed PostgreSQL versions. Select majors with <code>--pg</code>; published RPM/DEB coverage does not constrain source builds. Packages and debug files are written under <code>~/ext/pkg/</code>, with logs and manifests under <code>~/ext/log/</code>.',
        '构建目标由配方与已安装的 PostgreSQL 版本决定，可用 <code>--pg</code> 指定；已发布 RPM/DEB 的覆盖范围不限制源码构建。软件包与调试文件输出至 <code>~/ext/pkg/</code>，日志与清单位于 <code>~/ext/log/</code>。') + '</p>'
      + '<p><a href="https://github.com/pgsty/pgext/blob/main/docs/build-workflow.md" target="_blank" rel="noopener">'
      + bi('PGEXT build workflow', 'PGEXT 构建指南') + ' ↗</a></p></div>';
  } else {
    guide = '<div class="build-note"><b>' + bi('Upstream build', '上游构建') + '</b><p>' + t('ext.buildnone') + '</p>'
      + (full.pgrx_ver ? '<span class="badge">pgrx · ' + esc(full.pgrx_ver) + '</span>' : '') + '</div>';
  }
  return (links.length ? '<div class="source-cards">' + links.join('') + '</div>' : '') + guide;
}

/* ---------------- matrices, files & metadata ---------------- */
export function fullMatrixHTML(m, e, options) {
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
        return '<td><button class="cellv ' + cls + '"' + target + ' data-tip="' + esc(c.name || '') + ' · ' + c.count + ' builds">' + esc(String(c.org || '✓').toUpperCase()) + ' ' + esc(c.version || '') + '</button></td>';
      }
      counts.na++;
      return '<td><span class="cellv st-na">N/A</span></td>';
    }).join('');
    rows += '<tr class="' + archCls + (famStart ? ' fam-start' : '') + '"><td class="oslab"><a href="/os/' + encodeURIComponent(os) + '"><b>' + esc(osname) + '</b><span class="oslab-dot">.</span><span class="oslab-arch">' + esc(arch) + '</span></a></td>' + cells + '</tr>';
  }
  const total = m.pg.length * m.os.length;
  const avail = counts.pgdg + counts.pigsty;
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
    + chip('st-na', 'N/A', 'mx.legend.na', counts.na) + '</div>';
  return summary + '<div class="matrix-scroll"><table class="fmx"><thead><tr><th class="corner">OS / PG</th>' + ths + '</tr></thead><tbody>'
    + rows + '</tbody></table></div>' + legend;
}

export function filesHTML(f) {
  if (!f || !f.files || !f.files.length) return '<p class="empty-note">' + t('files.none') + '</p>';
  const pgs = [...new Set(f.files.map(x => x.pg))].sort((a, b) => b - a);
  const tabs = pgs.map((pg, i) => '<button data-ftab="' + pg + '" aria-selected="' + (i === 0) + '">PG ' + pg + '</button>').join('');
  const panes = pgs.map((pg, i) => {
    const rows = f.files.filter(x => x.pg === pg)
      .sort((a, b) => (catalog.osIdx[a.os] ?? 99) - (catalog.osIdx[b.os] ?? 99) || a.name.localeCompare(b.name));
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

export function overviewHTML(e, full) {
  return metadataHTML(full)
    + (full.comment ? '<blockquote class="catalog-quote">' + esc(full.comment) + '</blockquote>' : '');
}

export function metadataHTML(full) {
  const row = (label, value, mono) => '<div class="meta-row"><dt>' + label + '</dt><dd' + (mono ? ' class="mono"' : '') + '>' + (value || '') + '</dd></div>';
  const panel = (title, body, cls) => '<section class="meta-panel' + (cls ? ' ' + cls : '') + '"><h3>' + title + '</h3>' + body + '</section>';
  const dl = rows => '<dl>' + rows.join('') + '</dl>';
  const catalogRows = [
    row('ID', esc(full.id), true),
    row(bi('Version', '版本'), esc(full.version || ''), true),
    row(bi('PG Ver', 'PG 版本'), esc(pgRange(full.pg_ver)), true),
    row(bi('Package', '软件包'), '<a href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a>'
      + (full.family_size > 1 ? ' · ' + full.family_size + bi(' ext', ' 扩展') : ''), true),
    row(bi('Lead', '主扩展'), '<a href="' + extHref(full.lead_ext || full.name) + '">' + esc(full.lead_ext || full.name) + '</a>', true),
    row(bi('Kind', '形态'), full.kind ? '<span data-tip="' + esc(t('type.' + full.kind)) + '">' + esc(full.kind) + '</span>' : '', true)
  ];
  const attrRow = (en, zh, tipEN, tipZH, value) =>
    '<div class="meta-row attr-row" data-tip="' + esc(bi(tipEN, tipZH)) + '"><dt>' + bi(en, zh) + '</dt><dd>' + boolHTML(value) + '</dd></div>';
  const runtimeRows = [
    attrRow('Has Binary', '带有二进制', 'Ships executable command-line tools', '该扩展附带可执行的命令行工具', full.has_bin),
    attrRow('Has Library', '带有动态库', 'Ships a shared library (.so)', '该扩展附带动态链接库（.so 文件）', full.has_lib),
    attrRow('Explicit Loading', '需显式加载', 'Requires LOAD, preload, or another enablement step; see the usage guide', '需要 LOAD、预加载或其他启用步骤，具体方式见用法文档', full.need_load),
    attrRow('Need CREATE', '用 DDL 创建', 'Enabled per database with CREATE EXTENSION', '需执行 CREATE EXTENSION 语句创建后方可使用', full.need_ddl),
    attrRow('Need DBSU', '需超级用户', 'Only a database superuser can install it', '需要数据库超级用户权限才能安装', !full.trusted),
    attrRow('Relocatable', '模式可修改', 'Its objects can be installed into a schema of your choice', '扩展对象可以安装到指定的模式（schema）中', full.relocatable)
  ];
  const resources = [
    [bi('Homepage', '主页'), full.home_url],
    [bi('Repository', '仓库'), full.repo_url],
    [bi('License', '许可证'), full.license_url],
    ['Control', full.control_url],
    [bi('Docs', '文档'), full.doc_url],
    [bi('Author', '作者'), full.author_url],
    ['PGXN', full.pgxn_url],
    ['Cargo', full.cargo_url]
  ].map(([label, url]) => [label, mdSafeURL(url)]).filter(([, url]) => url).slice(0, 6);
  while (resources.length < 6) resources.push(null);
  const resourcesHTML = dl(resources.map(item => item
    ? row(item[0], '<a class="res-link" href="' + esc(item[1]) + '" target="_blank" rel="noopener"><span>'
      + esc(item[1].replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</span>' + ICON_OUT + '</a>', true)
    : '<div class="meta-row"><dt>&nbsp;</dt><dd></dd></div>'));
  const json = JSON.stringify(full, null, 2);
  const extra = '<details class="extra-meta"><summary>' + t('spec.extra') + '</summary>'
    + '<div class="extra-body"><button class="chip" data-copy="' + esc(json) + '">' + bi('copy JSON', '复制 JSON') + '</button>'
    + '<pre>' + esc(json) + '</pre></div></details>';
  return '<div class="metadata-grid tri">' + panel(t('ext.catalog'), dl(catalogRows)) + panel(t('ext.runtime'), dl(runtimeRows))
    + panel(t('ext.projectlinks'), resourcesHTML, 'links') + '</div>' + extra;
}

export function packageVersionsHTML(full) {
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

export function pkgDefinitionsHTML(full) {
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

export function depGraphHTML(full) {
  const up = full.requires || [];
  const down = full.required_by || [];
  const box = name => {
    const x = catalog.byName.get(name);
    return '<a class="dep-item" ' + (x ? catVar(x.cat) + ' data-hover-ext="' + esc(x.name) + '" ' : '') + 'href="' + extHref(name) + '">' + esc(name) + '</a>';
  };
  const col = (names, kind) => {
    if (!names.length) return '<div class="dep-none">' + t(kind === 'up' ? 'dep.noup' : 'dep.nodown') + '</div>';
    const items = names.slice(0, 12).map(box).join('')
      + (names.length > 12 ? '<span class="dep-item dep-more">+' + (names.length - 12) + '</span>' : '');
    return '<div class="dep-col dep-' + kind + '">' + items + '</div>';
  };
  const e = catalog.byName.get(full.name);
  return '<h3 class="section-subhead">' + t('ext.deps') + '</h3>'
    + '<div class="dep-graph' + (up.length ? ' has-up' : '') + (down.length ? ' has-down' : '') + '">'
    + '<span class="dep-cap dep-cap-up">' + bi('Upstream', '上游 Upstream') + '</span>'
    + '<span class="dep-cap dep-cap-down">' + bi('Downstream', '下游 Downstream') + '</span>'
    + col(up, 'up')
    + '<div class="dep-mid"><span class="dep-self" ' + (e ? catVar(e.cat) : '') + '>' + esc(full.name) + '</span></div>'
    + col(down, 'down') + '</div>';
}

export function seeAlsoHTML(full) {
  const names = full.see_also || [];
  if (!names.length) return '';
  const cards = names.map(name => {
    const e = catalog.byName.get(name);
    return '<a href="' + extHref(name) + '" ' + (e ? catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '"' : '') + '><code>' + esc(name) + '</code>'
      + '<span>' + esc(e ? desc(e) : bi('catalog reference', '目录引用')) + '</span>'
      + (e ? '<small>' + esc(e.cat) + ' · ' + esc(e.kind || '') + '</small>' : '') + '</a>';
  }).join('');
  return '<h3 class="section-subhead">See Also</h3><div class="relation-grid">' + cards + '</div>';
}

export function pkgFamilyTableHTML(members, lead, fulls) {
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

function familyMapFrom(full, members) {
  const map = new Map();
  if (full && Array.isArray(full.family)) {
    for (const m of full.family) map.set(m.name, m);
  }
  for (const m of members) {
    if (!map.has(m.name)) {
      const cached = FULLC.get(m.name);
      if (cached) map.set(m.name, cached);
    }
  }
  return map;
}

export function relationshipHTML(full, members, lead, fulls) {
  return '<h3 class="section-subhead sp-head"><a class="sp-pkg" href="' + pkgHref(full.pkg) + '">' + esc(full.pkg) + '</a></h3>'
    + pkgFamilyTableHTML(members, lead, fulls)
    + depGraphHTML(full)
    + seeAlsoHTML(full);
}

export function packageInstallHTML(pkg, full) {
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

export function changelogHTML(cl) {
  const entries = cl.entries || [];
  const zh = getLang() === 'zh';
  const verCell = e => {
    const o = e.old || '', n = e.new || '';
    if (o && n && o !== n) return '<span class="chlog-old">' + esc(o) + '</span><span class="chlog-arrow">→</span><b class="chlog-new">' + esc(n) + '</b>';
    if (!o && n) return '<span class="chlog-tag">' + bi('NEW', '新增') + '</span><b class="chlog-new">' + esc(n) + '</b>';
    if (n || o) return '<span class="chlog-same">' + esc(n || o) + '</span>';
    return '<span class="chlog-old">—</span>';
  };
  const note = e => { const s = zh ? (e.note_zh || e.note_en) : (e.note_en || e.note_zh); return s ? esc(s) : ''; };
  const rows = entries.map(e => '<tr><td class="mono chlog-ds">' + esc(e.ds) + '</td>'
    + '<td class="mono chlog-ver">' + verCell(e) + '</td>'
    + '<td class="chlog-note">' + note(e) + '</td></tr>').join('');
  const oldest = entries[entries.length - 1];
  return '<p class="section-lede">' + entries.length
    + bi(' release entries since ', ' 条发布记录，始于 ') + esc(oldest.ds)
    + bi(', distilled from the RPM / DEB release notes.', '，摘自 RPM / DEB 仓库发布说明。') + '</p>'
    + '<div class="version-table chlog-table"><div class="rows-scroll"><table><thead><tr>'
    + '<th>' + bi('date', '日期') + '</th><th>' + bi('version', '版本变更') + '</th><th>' + bi('note', '备注') + '</th>'
    + '</tr></thead><tbody>' + rows + '</tbody></table></div></div>';
}

/* ---------------- pages & hydration ---------------- */
export function extHTML(name) {
  const e = catalog.byName.get(name);
  if (!e) return notFoundHTML(name);
  const zh = getLang() === 'zh';
  const providerNote = (e.kernel || e.vendor) && !e.avail
    ? '<div class="notice provider">☁ ' + t('ext.providerNote', { provider: esc([e.vendor, e.kernel].filter(Boolean).join(' · ')) }) + '</div>' : '';
  const lifecycleNote = ['deprecated', 'archived', 'abandoned'].includes(e.lifecycle)
    ? '<div class="notice lifecycle">⚠ ' + t('ext.lifecycleNote', { state: '<b>' + esc(e.lifecycle) + '</b>' }) + '</div>' : '';
  const stateChip = e.avail
    ? '<span class="state-chip ok">● ' + t('state.avail') + '</span>'
    : '<span class="state-chip">○ ' + t('state.na') + '</span>';
  const toc = pageTocHTML([
    ['ext-overview', 'ext.overview'], ['ext-relations', 'ext.relations'], ['ext-packages', 'ext.veravail'],
    ['ext-build', 'ext.build'], ['ext-install', 'ext.install'], ['ext-usage', 'ext.docs']
  ].map(([id, key]) => [id, t(key)]));

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
    + (e.kernel ? '<span class="badge lbl"><span class="bk">' + t('spec.kernel') + (zh ? '：' : ':') + '</span><span class="bv">' + esc(e.kernel) + '</span></span>' : '')
    + (e.vendor ? '<span class="badge lbl"><span class="bk">' + t('spec.vendor') + (zh ? '：' : ':') + '</span><span class="bv">' + esc(e.vendor) + '</span></span>' : '')
    + '</div>'
    + '<div class="badge-row" id="d-links"></div>'
    + lifecycleNote + providerNote + '</div>'
    + '<aside class="ext-hero-side" id="d-side-main">' + heroSideMainHTML(e, FULLC.get(e.name), e.name) + '</aside>'
    + '</div></header>'
    + toc
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

export function pkgHTML(pkg) {
  const members = catalog.byPkg.get(pkg);
  if (!members || !members.length) return notFoundHTML(pkg);
  const lead = catalog.byName.get(members[0].lead) || members[0];
  const packaged = members.some(e => e.avail);
  const toc = pageTocHTML([
    ['pkg-family', t('ext.family')], ['pkg-packages', t('ext.pkgs')],
    ['pkg-downloads', t('ext.downloads')], ['pkg-build', t('ext.build')],
    ['pkg-install', t('ext.install')], ['pkg-changelog', t('ext.changelog')]
  ]);
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
    + toc
    + '<section class="section pkg-section" id="pkg-family"><h2>' + t('ext.family') + '</h2><p class="section-lede">'
    + members.length + bi(' extension definitions delivered by this package.', ' 个由该软件包交付的扩展定义。') + '</p><div id="p-family">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-packages"><h2>' + t('ext.pkgs') + '</h2>'
    + '<div id="p-definitions">' + skel(3) + '</div><div id="p-matrix" style="margin-top:12px">' + skel(5) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-downloads"><h2>' + t('ext.downloads') + '</h2><div id="p-files">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-build"><h2>' + t('ext.build') + '</h2><div id="p-build">' + skel(3) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-install"><h2>' + t('ext.install') + '</h2><div id="p-install">' + skel(4) + '</div></section>'
    + '<section class="section pkg-section" id="pkg-changelog"><h2>' + t('ext.changelog') + '</h2><div id="p-changelog">' + skel(3) + '</div></section>'
    + '</article>';
}

function recordVisit(enc, chipId, tok) {
  fetch('/api/v1/ext/' + enc + '/visit', { method: 'POST' })
    .then(res => res.ok ? res.json() : null)
    .then(v => {
      if (!v || tok !== hydSeq) return;
      const chip = document.getElementById(chipId);
      if (!chip) return;
      chip.hidden = false;
      chip.innerHTML = ICON_EYE + fmtInt(v.visits);
      chip.dataset.tip = fmtInt(v.visits) + ' ' + t('ext.visits');
    }).catch(() => {});
}

async function fetchCached(cache, key, url) {
  let val = cache.get(key);
  if (!val) { val = await j(url); cache.set(key, val); }
  return val;
}

export async function hydrateExt(name) {
  const tok = ++hydSeq;
  const e = catalog.byName.get(name);
  if (!e) return;
  const enc = encodeURIComponent(name);
  const fullP = fetchCached(FULLC, name, '/api/v1/ext/' + enc);
  const matrixP = fetchCached(MXC, e.pkg, '/api/v1/ext/' + enc + '/matrix');

  fullP.then(full => {
    if (tok !== hydSeq) return;
    fill('d-overview', overviewHTML(e, full));
    fill('d-versions', packageVersionsHTML(full));
    fill('d-build', buildHTML(full));
    fill('d-side-main', heroSideMainHTML(e, full, e.name));
    fill('d-links', heroLinkBadgesHTML(full));
    fill('d-install', installHTML(e, full));
    const members = catalog.byPkg.get(e.pkg) || [e];
    const lead = catalog.byName.get(members[0].lead) || members[0];
    fill('d-deps', relationshipHTML(full, members, lead, familyMapFrom(full, members)));
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['d-overview', 'd-deps', 'd-versions', 'd-build', 'd-install']) fill(id, hydrateErr(err));
  });

  matrixP.then(matrix => {
    if (tok !== hydSeq) return;
    fill('d-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, e) : '<p class="empty-note">' + t('files.none') + '</p>');
  }).catch(err => { if (tok === hydSeq) fill('d-matrix', hydrateErr(err)); });

  recordVisit(enc, 'd-visits', tok);

  if (e.docbits) {
    (async () => {
      const wantZh = getLang() === 'zh' && (e.docbits & 2);
      const lang = wantZh ? 'zh' : 'en';
      try {
        const d = await fetchCached(DOCC, name + '|' + lang, '/api/v1/ext/' + enc + '/doc?lang=' + lang);
        if (tok !== hydSeq) return;
        const note = (getLang() === 'zh' && !wantZh) ? '<p class="empty-note" style="margin-bottom:12px">' + t('ext.doconlyen') + '</p>' : '';
        const content = String(d.content || '').replace(/^\s*(?:#{1,2}\s+(?:usage|用法)\s*\n+|(?:usage|用法)\s*\n[-=]{3,}\s*\n+)/i, '');
        fill('d-doc', note + '<article class="prose usage-prose">' + renderMD(content, { usage: true }) + '</article>');
      } catch (err) { if (tok === hydSeq) fill('d-doc', hydrateErr(err)); }
    })();
  }
}

function hidePkgChangelog() {
  const sec = document.getElementById('pkg-changelog');
  if (sec) sec.style.display = 'none';
  const btn = document.querySelector('.page-toc button[data-scroll="pkg-changelog"]');
  if (btn) btn.style.display = 'none';
}

export async function hydratePkg(pkg) {
  const members = catalog.byPkg.get(pkg);
  if (!members || !members.length) return;
  const lead = catalog.byName.get(members[0].lead) || members[0];
  const tok = ++hydSeq;
  const enc = encodeURIComponent(lead.name);
  const fullP = fetchCached(FULLC, lead.name, '/api/v1/ext/' + enc);
  const matrixP = fetchCached(MXC, pkg, '/api/v1/ext/' + enc + '/matrix');

  recordVisit(enc, 'p-visits', tok);

  fetchCached(CHLOGC, pkg, '/api/v1/ext/' + enc + '/changelog').then(cl => {
    if (tok !== hydSeq) return;
    if (cl && cl.entries && cl.entries.length) fill('p-changelog', changelogHTML(cl));
    else hidePkgChangelog();
  }).catch(() => { if (tok === hydSeq) hidePkgChangelog(); });

  fullP.then(full => {
    if (tok !== hydSeq) return;
    fill('p-side-main', heroSideMainHTML(lead, full, pkg));
    fill('p-links', heroLinkBadgesHTML(full));
    fill('p-family', pkgFamilyTableHTML(members, lead, familyMapFrom(full, members)));
    fill('p-definitions', pkgDefinitionsHTML(full));
    fill('p-build', buildHTML(full));
    fill('p-install', packageInstallHTML(pkg, full));
  }).catch(err => {
    if (tok !== hydSeq) return;
    for (const id of ['p-family', 'p-definitions', 'p-build', 'p-install']) fill(id, hydrateErr(err));
  });

  matrixP.then(matrix => {
    if (tok !== hydSeq) return;
    fill('p-matrix', matrix.cells && matrix.cells.length ? fullMatrixHTML(matrix, lead, { pkg }) : '<p class="empty-note">' + t('files.none') + '</p>');
  }).catch(err => { if (tok === hydSeq) fill('p-matrix', hydrateErr(err)); });

  try {
    const files = await fetchCached(FILEC, pkg, '/api/v1/ext/' + enc + '/files');
    if (tok === hydSeq) fill('p-files', filesHTML(files));
  } catch (err) { if (tok === hydSeq) fill('p-files', hydrateErr(err)); }
}
