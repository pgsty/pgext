/* ================================================================
   Catalog home, universe dot-field canvas, dimension index pages,
   navigation landing pages and About view.
   ================================================================ */
import { esc, fmtInt, extHref, catHref, dimHref, licenseHue, langHue, dimSeg } from './util.js';
import { t, bi, catName } from './i18n.js';
import { mdCodeHTML } from './markdown.js';
import { catalog, S, effSort, parsePGs, catDesc } from './state.js';
import { DIMS, DIM_GROUPS, dimValueName, dimValueHTML, dimValues, targetAvailable, scopeMatch, runFilter, projectGroups, buildSQL } from './filter.js';
import { tileHTML, packageTileHTML, rowHTML, packageRowHTML, extTableHTML, skel, notFoundHTML } from './components.js';
import { osNavHTML } from './matrix.js';

/* ---------------- home ---------------- */
export function homeHTML() {
  const stats = [
    [catalog.nAll, 'hero.s1'],
    [catalog.nAvail, 'hero.s2'],
    [catalog.nPkgs, 'hero.s3'],
    [catalog.oss.length, 'hero.s4'],
    [catalog.pgs.length, 'hero.s5'],
    [catalog.nSlots, 'hero.s6']
  ].map(([n, key]) => '<li data-tip="' + esc(t(key + '.tip')) + '"><span class="num">' + fmtInt(n) + '</span><span class="lbl">' + t(key) + '</span></li>').join('');
  return '<section class="hero wrap">'
    + '<p class="eyebrow">' + t('hero.eyebrow') + '</p>'
    + '<h1 class="oneline">' + t('hero.title') + '</h1>'
    + '<p class="hero-sub oneline">' + t('hero.sub', { all: fmtInt(catalog.nAll), avail: fmtInt(catalog.nAvail), os: fmtInt(catalog.oss.length) }) + '</p>'
    + '<ul class="hero-stats">' + stats + '</ul>'
    + '<div class="universe"><canvas id="ufield" height="80" aria-hidden="true" style="width:100%;display:block;cursor:crosshair"></canvas></div>'
    + '<div class="console">'
    + '<div class="searchbox"><span class="prompt">pgext=#</span>'
    + '<input id="q" type="search" autocomplete="off" spellcheck="false" placeholder="' + esc(t('search.ph', { n: fmtInt(catalog.nAll) })) + '" value="' + esc(S.q) + '" aria-label="search extensions">'
    + '<kbd>/</kbd></div>'
    + '</div>'
    + '<div id="dynamic">' + dynamicHTML() + '</div>'
    + '</section>';
}

export function dynamicHTML() {
  const { f, list } = runFilter();
  const { sql, html: sqlHTML } = buildSQL(f, list.length);
  const sortOpts = [['name', t('sort.name')], ['stars', t('sort.stars')], ['updated', t('sort.updated')]].map(([v, l]) =>
    '<option value="' + v + '"' + (effSort() === v ? ' selected' : '') + '>' + l + '</option>').join('');

  const scopedEXT = S.scope ? catalog.ext.filter(e => scopeMatch(e, S.scope)) : catalog.ext;
  const countMap = dim => new Map(dimValues(dim, scopedEXT).map(x => [String(x.v), x.n]));
  const button = (key, value, label, on, count, cls, style, tip) => {
    const data = key === 'pg' ? ' data-pg-toggle="' + esc(value) + '"' : ' data-fkey="' + key + '" data-fval="' + esc(value) + '"';
    return '<button class="facet-btn' + (cls ? ' ' + cls : '') + '"' + data + ' aria-pressed="' + on + '"'
      + (style ? ' style="' + style + '"' : '') + (tip ? ' data-tip="' + esc(tip) + '"' : '') + '>'
      + (cls === 'category' || cls === 'hued' ? '<i></i>' : '') + '<span>' + label + '</span>'
      + (count == null ? '' : '<small>' + fmtInt(count) + '</small>') + '</button>';
  };
  const huedButton = (key, value, label, on, count, hue) =>
    button(key, value, label, on, count, hue ? 'hued' : '', hue ? '--seg:var(--h-' + hue + ')' : '');
  const any = (key, on) => button(key, '', t('filter.any'), on, null, 'any');
  const row = (label, values, cls, href) => '<div class="facet-row' + (cls ? ' ' + cls : '') + '"><div class="facet-name">'
    + (href ? '<a class="facet-head" href="' + href + '"><b>' + label + '</b><i>→</i></a>' : '<b>' + label + '</b>')
    + '</div><div class="facet-values">' + values + '</div></div>';

  const scopeValues = [
    ['', t('chip.all'), catalog.nAll], ['packaged', t('chip.packaged'), catalog.nAvail], ['unpacked', t('chip.unpacked'), catalog.nAll - catalog.nAvail],
    ['kernel', t('chip.kernel'), catalog.nKernel], ['vendor', t('chip.vendor'), catalog.nVendor], ['contrib', t('chip.contrib'), catalog.nContrib]
  ].map(([v, label, count]) => v === 'contrib'
    ? huedButton('scope', v, label, S.scope === v, count, 'pgnavy')
    : button('scope', v, label, S.scope === v, count)).join('');

  const catCounts = countMap('category');
  const categoryValues = any('cat', !S.cat) + catalog.catOrder.map(c =>
    button('cat', c, '<code>' + c + '</code>', S.cat === c, catCounts.get(c) || 0, 'category', '--seg:var(--c-' + c + ')', catName(c))).join('');

  const pgCounts = countMap('pg'), selectedPGs = parsePGs(S.pg);
  const pgValues = [...new Set([...catalog.pgs, ...selectedPGs])].sort((a, b) => b - a);
  const pgButtons = button('pg', '', t('filter.any'), !selectedPGs.length, null, 'any')
    + pgValues.map(pg => button('pg', String(pg), 'PG ' + pg, selectedPGs.includes(pg), pgCounts.get(String(pg)) || 0)).join('');

  const langCounts = countMap('lang');
  const langButtons = any('lang', !S.lang) + [...langCounts.entries()]
    .map(([v, count]) => huedButton('lang', v, esc(v), S.lang === v, count, langHue(v))).join('');

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
    results = extTableHTML(list);
  } else {
    results = '<ul class="wall ext-wall">' + list.map(tileHTML).join('') + '</ul>';
  }

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

/* ---------------- universe dot-field canvas ---------------- */
let ucells = null;
let rafPending = false;

export function drawField(forceColors = false) {
  const cv = document.getElementById('ufield');
  if (!cv) return;
  const dpr = window.devicePixelRatio || 1;
  const W = cv.clientWidth || cv.parentElement.clientWidth;
  const pitch = W > 900 ? 8 : W > 560 ? 7 : 6;
  const dot = pitch - 2;
  const cols = Math.max(24, Math.floor(W / pitch));

  // Rebuild cell coordinates and CSS color palette only when layout, catalog
  // or theme changes — hover redraws reuse both without getComputedStyle.
  if (!ucells || forceColors || ucells.W !== W || ucells.dpr !== dpr || ucells.src !== catalog.ufield) {
    const root = getComputedStyle(document.documentElement);
    const colors = {};
    for (const c of catalog.catOrder) colors[c] = root.getPropertyValue('--c-' + c).trim() || '#888';
    const cells = [];
    const byPos = new Map();
    let cell = 0, prevCat = null;
    for (const e of catalog.ufield) {
      if (prevCat && e.cat !== prevCat) cell += 1;
      const col = cell % cols, row = Math.floor(cell / cols);
      const item = { col, row, e };
      cells.push(item);
      byPos.set(row * cols + col, e);
      prevCat = e.cat; cell++;
    }
    const H = (Math.floor(cell / cols) + 1) * pitch;
    cv.width = W * dpr; cv.height = H * dpr;
    cv.style.height = H + 'px';
    ucells = { cells, byPos, colors, pitch, dot, cols, W, H, dpr, src: catalog.ufield, hoverName: ucells ? ucells.hoverName : null };
  }

  const ctx = cv.getContext('2d');
  ctx.setTransform(dpr, 0, 0, dpr, 0, 0);
  ctx.clearRect(0, 0, ucells.W, ucells.H);
  const hov = ucells.hoverName;
  for (const c of ucells.cells) {
    ctx.fillStyle = ucells.colors[c.e.cat];
    const x = c.col * ucells.pitch, y = c.row * ucells.pitch;
    if (hov && c.e.name === hov) {
      ctx.globalAlpha = 1;
      ctx.fillRect(x - 1, y - 1, ucells.dot + 2, ucells.dot + 2);
    } else {
      ctx.globalAlpha = c.e.avail ? 0.95 : 0.45;
      ctx.fillRect(x, y, ucells.dot, ucells.dot);
    }
  }
  ctx.globalAlpha = 1;
}

export function setFieldHover(name) {
  if (!ucells || ucells.hoverName === name) return;
  ucells.hoverName = name;
  if (rafPending) return;
  rafPending = true;
  requestAnimationFrame(() => { rafPending = false; drawField(false); });
}

export function fieldHit(ev) {
  if (!ucells) return null;
  const cv = document.getElementById('ufield');
  if (!cv) return null;
  const r = cv.getBoundingClientRect();
  const col = Math.floor((ev.clientX - r.left) / ucells.pitch);
  const row = Math.floor((ev.clientY - r.top) / ucells.pitch);
  if (col < 0 || col >= ucells.cols || row < 0) return null;
  return ucells.byPos.get(row * ucells.cols + col) || null;
}

/* ---------------- dimension navigation pages ---------------- */
export const NAV_DIMS = {
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
    match: (e, v) => e.repo.toUpperCase() === v,
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
    lede: () => bi('Extensions installable on this Linux target — a per-platform digest of the build matrix.',
                   '可在该 Linux 目标上安装的扩展——按平台切片的构建矩阵摘要。')
  }
};

export function navValueHref(dim, value) {
  const m = { category: '/cate/', lang: '/lang/', repo: '/repo/', license: '/license/', pg: '/pg/', os: '/os/' }[dim];
  return m ? m + encodeURIComponent(value) : '/?' + DIMS[dim].key + '=' + encodeURIComponent(value);
}

export function navPageHTML(kind, raw) {
  const cfg = NAV_DIMS[kind];
  if (!cfg) return notFoundHTML(raw);
  const requested = cfg.canon(raw);
  const vals = dimValues(cfg.dim);
  const hit = vals.find(x => x.v.toLowerCase() === requested.toLowerCase());
  const v = hit ? hit.v : requested;
  const members = catalog.ext.filter(e => cfg.match(e, v)).sort((a, b) => a.name.localeCompare(b.name));
  if (!members.length) return notFoundHTML(v);
  const featured = members.filter(e => e.avail)
    .sort((a, b) => ((b.docbits ? 1 : 0) - (a.docbits ? 1 : 0)) || (b.stars ?? -1) - (a.stars ?? -1))
    .slice(0, 4);
  const isCat = kind === 'category';
  const strip = isCat ? '<div class="cat-strip">' + catalog.catOrder.map(x =>
    '<a href="' + catHref(x) + '" style="--seg:var(--c-' + x + ');flex:' + (catalog.cats[x].count || 1) + ' 1 0" aria-current="' + (x === v) + '" data-tip="' + x + ' · ' + catalog.cats[x].count + '"></a>').join('') + '</div>' : '';
  const navVals = isCat
    ? catalog.catOrder.map(c => ({ v: c, n: (vals.find(x => x.v === c) || {}).n || 0 }))
    : kind === 'pg'
      ? catalog.pgs.map(pg => ({ v: String(pg), n: (vals.find(x => x.v === String(pg)) || {}).n || 0 }))
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
  const body = kind === 'pg' || kind === 'os'
    ? '<div class="section"><h2>' + t('gmx.title') + '</h2><div id="slice-matrix">' + skel(6) + '</div></div>'
    : kind === 'repo'
      ? '<div class="section"><h2>' + t('gmx.title') + '</h2><p class="gmx-solo" id="repo-solo" hidden></p><div id="repo-matrix">' + skel(6) + '</div></div>'
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
    + body + '</article>';
}

/* ---------------- browse & dimension index ---------------- */
export function browseHTML() {
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
  const catCards = '<section class="dimension-group"><header><h2>' + t('browse.cats') + '</h2><span>' + catalog.catOrder.length + '</span></header>'
    + '<ul class="cat-cards">' + catalog.catOrder.map(c =>
      '<li><a class="cat-card" href="' + catHref(c) + '" style="--seg:var(--c-' + c + ')">'
      + '<code>' + c + '</code><b>' + esc(catName(c)) + '</b><span>' + fmtInt((catalog.cats[c] || {}).count || 0) + '</span></a></li>').join('')
    + '</ul></section>';
  const groups = catCards + DIM_GROUPS.map(([title, dims]) => '<section class="dimension-group"><header><h2>' + t(title) + '</h2><span>' + dims.length + '</span></header>'
    + '<ul class="dims">' + dims.map(card).join('') + '</ul></section>').join('');
  return '<article class="page wrap"><header class="page-head">'
    + '<p class="eyebrow">' + t('nav.browse').toLowerCase() + '</p>'
    + '<h1>' + t('browse.title') + '</h1><p class="lede">' + t('browse.lede') + '</p>'
    + '<div class="browse-stats"><span><b>' + Object.keys(DIMS).length + '</b>' + bi('dimensions', '个维度') + '</span>'
    + '<span><b>' + fmtInt(catalog.nAll) + '</b>' + bi('extensions', '个扩展') + '</span><span><b>' + fmtInt(catalog.nProjects) + '</b>' + bi('projects', '个项目') + '</span></div></header>'
    + groups + '</article>';
}

export function dimHTML(dim) {
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

/* ---------------- about ---------------- */
export function aboutHTML() {
  const out = (href, label, text) => '<li><span class="tag">' + label + '</span><a href="' + href + '" target="_blank" rel="noopener">' + text + ' ↗</a></li>';
  return '<article class="page wrap">'
    + '<header class="page-head"><p class="eyebrow">pgext.cloud</p>'
    + '<h1>' + bi('About this catalog', '关于本目录') + '</h1><p class="lede">'
    + bi('PGEXT.CLOUD is the census of the PostgreSQL extension ecosystem — served live from the pgext catalog database by a single Go binary.',
         'PGEXT.CLOUD 是 PostgreSQL 扩展生态的全量普查——由单个 Go 二进制从 pgext 目录数据库实时供给。') + '</p></header>'
    + '<div class="about-cols">'
    + '<div><h3>' + bi('What this is', '这是什么') + '</h3>'
    + '<p>' + bi('A curated catalog of <b>' + fmtInt(catalog.nAll) + '</b> PostgreSQL extensions: upstream repositories and RPM/DEB indexes are inspected continuously, with categories, dependencies, lifecycles, kernels, vendors and bilingual usage manuals maintained as catalog dimensions. <b>' + fmtInt(catalog.nAvail) + '</b> are catalogued as packaged, with availability tracked across <b>' + fmtInt(catalog.oss.length) + '</b> Linux OS/architecture targets and <b>' + fmtInt(catalog.pgs.length) + '</b> PostgreSQL majors.',
                 '一份收录 <b>' + fmtInt(catalog.nAll) + '</b> 个 PostgreSQL 扩展的策展目录：持续抓取上游仓库与 RPM/DEB 软件源索引，维护分类、依赖、生命周期、内核、厂商与双语用法手册等目录维度。其中 <b>' + fmtInt(catalog.nAvail) + '</b> 个标记为已打包，跟踪 <b>' + fmtInt(catalog.oss.length) + '</b> 个 Linux 系统/架构目标与 <b>' + fmtInt(catalog.pgs.length) + '</b> 个 PG 大版本的软件包可用性。') + '</p>'
    + '<p>' + bi('This site is <code>pgext serve</code>: web assets embedded in one binary, data queried live, snapshots cached in memory. Also read our post that topped Hacker News: ',
                 '本站就是 <code>pgext serve</code>：网页资产内嵌于单个二进制，数据实时查询、内存快照缓存。也欢迎阅读我们登上 Hacker News 头条的博客：')
    + '<a href="https://medium.com/@fengruohang/postgres-is-eating-the-database-world-157c204dcfc4" target="_blank" rel="noopener"><i>PostgreSQL is eating the Database World</i> ↗</a></p>'
    + '<a class="about-eco" href="/matrix"><img src="https://pigsty.io/img/pigsty/ecosystem.png" alt="' + esc(bi('The PostgreSQL extension ecosystem', 'PostgreSQL 扩展生态宇宙')) + '" loading="lazy"></a>'
    + '</div>'
    + '<div><h3>' + bi('Highlights', '亮点') + '</h3><ul class="roadmap about-highlights">'
    + '<li><span class="tag">catalog</span>' + bi('<b>' + fmtInt(catalog.nAvail) + '</b> extensions in the packaged catalog', '<b>' + fmtInt(catalog.nAvail) + '</b> 个扩展标记为已打包') + '</li>'
    + '<li><span class="tag">native</span>' + bi('RPM/DEB packages, properly built, freely composable', 'Linux 原生 RPM/DEB 软件包，规范构建、自由组合') + '</li>'
    + '<li><span class="tag">pig</span><a href="' + bi('https://pig.pgsty.com', 'https://pig.pgsty.com/zh') + '" target="_blank" rel="noopener">' + bi('A handy CLI on apt/dnf: zero-config, out-of-the-box installs', '趁手的 pig 命令行：零配置、开箱即用') + ' ↗</a></li>'
    + '<li><span class="tag">pgdg</span>' + bi('PGDG-compliant — drop-in with the official PostgreSQL kernel', '兼容 PGDG——与官方 PostgreSQL 内核即插即用') + '</li>'
    + '<li><span class="tag">cdn</span>' + bi('Distributed worldwide via Cloudflare CDN, fast and reliable', '通过 Cloudflare CDN 全球分发，快速可靠') + '</li>'
    + '<li><span class="tag">oss</span>' + bi('Reproducible builds on public infra — free for everyone', '可复现构建、公开基础设施，对所有人免费') + '</li>'
    + '</ul></div>'
    + '<div><h3>' + bi('Get started', '快速上手') + '</h3>'
    + mdCodeHTML('bash', ['curl -fsSL https://repo.pigsty.io/pig | bash  # install pig cli', 'pig repo set                  # setup package repositories', 'pig install pg18              # install PostgreSQL 18', 'pig install pg_duckdb -v 18    # install pg_duckdb for PG 18'].join('\n'))
    + '<p><a href="' + bi('https://pig.pgsty.com/start/', 'https://pig.pgsty.com/zh/start/') + '" target="_blank" rel="noopener">' + bi('Read the PIG getting-started guide', '阅读 PIG 快速上手指南') + ' ↗</a></p></div>'
    + '<div><h3>' + bi('Build extensions with PGEXT', '使用 PGEXT 构建扩展') + '</h3>'
    + '<p>' + bi('The same binary also prepares Linux builders and builds native RPM/DEB packages from Pigsty recipes, without PIG or a metadata database.',
      '同一个二进制还可准备 Linux 构建环境，并按 Pigsty 配方构建原生 RPM/DEB 软件包，无需 PIG 或元数据数据库。') + '</p>'
    + mdCodeHTML('bash', ['pgext env repo', 'pgext env tool --pg 18', 'pgext env spec', 'pgext pkg pgvector --pg 18'].join('\n'))
    + '<p><a href="https://github.com/pgsty/pgext/blob/main/docs/build-workflow.md" target="_blank" rel="noopener">' + bi('PGEXT build workflow', 'PGEXT 构建指南') + ' ↗</a></p></div>'
    + '<div><h3>' + bi('Friends', '友情链接') + '</h3><ul class="roadmap">'
    + out('https://pigsty.io', 'pigsty.io', bi('Pigsty — Battery-Included PostgreSQL Distribution', 'Pigsty——开箱即用的 PostgreSQL 发行版'))
    + out('https://pigsty.cc', 'pigsty.cc', bi('Pigsty Chinese site & mirror', 'Pigsty 中文站与镜像'))
    + out(bi('https://pig.pgsty.com', 'https://pig.pgsty.com/zh'), 'pig.pgsty.com', bi('PIG — PostgreSQL package manager', 'PIG——PostgreSQL 包管理器'))
    + out('https://github.com/pgsty/pgext', 'github', bi('pgsty/pgext — the code behind this catalog', 'pgsty/pgext——本目录背后的代码'))
    + '</ul></div>'
    + '<div><h3>' + bi('Data', '数据') + '</h3>'
    + '<p>' + bi("The whole catalog is versioned as plain CSV in the repository's <code>db/</code> directory — take it straight from GitHub instead of crawling these pages.",
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
    + '<p class="universe-note" style="margin-top:34px">' + bi('Snapshot loaded ' + (catalog.meta.generated || '—') + ' · refreshed automatically.',
      '快照载入于 ' + (catalog.meta.generated || '—') + ' · 自动刷新。') + '</p>'
    + '</article>';
}
