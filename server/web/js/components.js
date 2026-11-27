/* ================================================================
   Shared icons, cards, tables, navigation, footer, theme, tooltip
   and interactive hovercard components.
   ================================================================ */
import { esc, fmtNum, pgRange, extHref, pkgHref, catHref, dimHref, catVar, licenseHue, langHue, repoHue, hueVar } from './util.js';
import { getLang, t, bi, catName, desc } from './i18n.js';
import { mdSafeURL } from './markdown.js';
import { F, FULLC, catalog, j } from './state.js';
import { DIMS, DIM_GROUPS } from './filter.js';

/* ---------------- icons ---------------- */
export const GH_SVG = '<svg width="13" height="13" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27s1.36.09 2 .27c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0 0 16 8c0-4.42-3.58-8-8-8Z"/></svg>';
export const EXTLINK_SVG = '<svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.1" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M13.5 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-6.5"/><path d="M14 3h7v7M21 3l-9.5 9.5"/></svg>';
export const PIGSTY_LOGO = '<svg class="pigsty-logo" width="21" height="21" viewBox="0 0 24 24" aria-hidden="true"><g stroke="#fff" stroke-width=".67" stroke-linecap="round" stroke-linejoin="round">'
  + '<path d="M7.67 11.97 9.83 8.22h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#bbbbbb" fill-opacity=".95"/>'
  + '<path d="M7.67 19.47l2.16-3.75h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#de372c" fill-opacity=".85"/>'
  + '<path d="M14.16 15.75 16.33 12h4.33l2.16 3.75-2.16 3.75h-4.33Z" fill="#424242" fill-opacity=".9"/>'
  + '<path d="M14.16 8.23l2.17-3.75h4.33l2.16 3.75-2.16 3.75h-4.33Z" fill="#ffa269" fill-opacity=".9"/>'
  + '<path d="M7.67 4.5 9.83.75h4.33l2.17 3.75-2.17 3.75H9.83Z" fill="#419edb" fill-opacity=".9"/>'
  + '<path d="M1.15 8.23l2.16-3.75h4.33l2.17 3.75-2.17 3.75H3.31Z" fill="#2f6793" fill-opacity=".9"/>'
  + '<path d="M1.18 15.74l2.17-3.75h4.33l2.17 3.75-2.17 3.75H3.35Z" fill="#53ac79" fill-opacity=".9"/>'
  + '</g></svg>';
export const ICON_DICE = '<svg width="17" height="17" viewBox="0 0 18 18" fill="none" stroke="currentColor" stroke-width="1.25" stroke-linejoin="round" aria-hidden="true">'
  + '<path d="M9 1.5 15.6 5.2v7.4L9 16.5 2.4 12.6V5.2L9 1.5Z"/><path d="M2.4 5.2 9 8.9l6.6-3.7M9 8.9v7.6"/>'
  + '<circle cx="9" cy="5.2" r=".85" fill="currentColor" stroke="none"/>'
  + '<circle cx="5" cy="10.1" r=".8" fill="currentColor" stroke="none"/><circle cx="6.7" cy="12.7" r=".8" fill="currentColor" stroke="none"/>'
  + '<circle cx="11.2" cy="12.5" r=".8" fill="currentColor" stroke="none"/><circle cx="12.4" cy="11.2" r=".8" fill="currentColor" stroke="none"/><circle cx="13.6" cy="9.9" r=".8" fill="currentColor" stroke="none"/></svg>';
export const ICON_LANG = '<svg width="17" height="17" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">'
  + '<text x="0" y="11.5" font-size="11" font-family="system-ui,sans-serif">文</text>'
  + '<text x="10.5" y="19" font-size="10.5" font-weight="700" font-family="system-ui,sans-serif">A</text></svg>';
export const ICON_SUN = '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" aria-hidden="true">'
  + '<circle cx="12" cy="12" r="4.4"/>'
  + '<path d="M12 2.6v2.1M12 19.3v2.1M2.6 12h2.1M19.3 12h2.1M5.2 5.2l1.5 1.5M17.3 17.3l1.5 1.5M18.8 5.2l-1.5 1.5M6.7 17.3l-1.5 1.5"/></svg>';
export const ICON_MOON = '<svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">'
  + '<path d="M20.6 14.6A8.7 8.7 0 0 1 9.4 3.4a8.7 8.7 0 1 0 11.2 11.2Z"/></svg>';
export const ICON_STAR = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 .25a.75.75 0 0 1 .673.418l1.882 3.815 4.21.612a.75.75 0 0 1 .416 1.279l-3.046 2.97.719 4.192a.751.751 0 0 1-1.088.791L8 12.347l-3.766 1.98a.75.75 0 0 1-1.088-.79l.72-4.194L.818 6.374a.75.75 0 0 1 .416-1.28l4.21-.611L7.327.668A.75.75 0 0 1 8 .25Zm0 2.445L6.615 5.5a.75.75 0 0 1-.564.41l-3.097.45 2.24 2.184a.75.75 0 0 1 .216.664l-.528 3.084 2.769-1.456a.75.75 0 0 1 .698 0l2.77 1.456-.53-3.084a.75.75 0 0 1 .216-.664l2.24-2.183-3.096-.45a.75.75 0 0 1-.564-.41L8 2.694Z"/></svg>';
export const ICON_FORK = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M5 5.372v.878c0 .414.336.75.75.75h4.5a.75.75 0 0 0 .75-.75v-.878a2.25 2.25 0 1 1 1.5 0v.878a2.25 2.25 0 0 1-2.25 2.25h-1.5v2.128a2.251 2.251 0 1 1-1.5 0V8.5h-1.5A2.25 2.25 0 0 1 3.5 6.25v-.878a2.25 2.25 0 1 1 1.5 0ZM5 3.25a.75.75 0 1 0-1.5 0 .75.75 0 0 0 1.5 0Zm6.75.75a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Zm-3 8.75a.75.75 0 1 0-1.5 0 .75.75 0 0 0 1.5 0Z"/></svg>';
export const ICON_EYE = '<svg width="14" height="14" viewBox="0 0 16 16" fill="currentColor" aria-hidden="true"><path d="M8 2c1.981 0 3.671.992 4.933 2.078 1.27 1.091 2.187 2.345 2.637 3.023a1.62 1.62 0 0 1 0 1.798c-.45.678-1.367 1.932-2.637 3.023C11.67 13.008 9.981 14 8 14c-1.981 0-3.671-.992-4.933-2.078C1.797 10.831.88 9.577.43 8.899a1.62 1.62 0 0 1 0-1.798c.45-.678 1.367-1.932 2.637-3.023C4.33 2.992 6.019 2 8 2ZM1.679 7.932a.12.12 0 0 0 0 .136c.411.622 1.241 1.75 2.366 2.717C5.176 11.758 6.527 12.5 8 12.5c1.473 0 2.825-.742 3.955-1.715 1.124-.967 1.954-2.096 2.366-2.717a.12.12 0 0 0 0-.136c-.412-.621-1.242-1.75-2.366-2.717C10.824 4.242 9.473 3.5 8 3.5c-1.473 0-2.825.742-3.955 1.715-1.124.967-1.954 2.096-2.366 2.717ZM8 10a2 2 0 1 1-.001-3.999A2 2 0 0 1 8 10Z"/></svg>';
export const ICON_OUT = '<svg class="ext-mark" width="9" height="9" viewBox="0 0 14 14" fill="none" stroke="currentColor" stroke-width="1.6" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M6 2.8H3.4A1.4 1.4 0 0 0 2 4.2v6.4A1.4 1.4 0 0 0 3.4 12h6.4a1.4 1.4 0 0 0 1.4-1.4V8M8.2 2h3.8v3.8M11.7 2.3 6.6 7.4"/></svg>';
export const ICON_LINK = '<svg width="13" height="13" viewBox="0 0 16 16" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M6.6 9.4 9.4 6.6M5 7.5 3.4 9.1a2.6 2.6 0 0 0 3.7 3.7l1.6-1.6M11 8.5l1.6-1.6a2.6 2.6 0 0 0-3.7-3.7L7.3 4.8"/></svg>';
export const GH_SVG_LG = GH_SVG.replace('width="13" height="13"', 'width="26" height="26"');
export const ICON_LINK_LG = ICON_LINK.replace('width="13" height="13"', 'width="24" height="24"');

/* ---------------- theme ---------------- */
let themeOverride = null;
try {
  if (typeof location !== 'undefined') {
    const q = new URLSearchParams(location.search).get('theme');
    if (q === 'dark' || q === 'light') themeOverride = q;
  }
} catch (e) {}

export function setThemeOverride(v) { themeOverride = v; }

export function currentTheme() {
  let mode = themeOverride || 'auto';
  if (!themeOverride) {
    try { if (typeof window !== 'undefined' && window.localStorage) mode = window.localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {}
  }
  if (mode === 'auto') {
    return typeof window !== 'undefined' && window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
  }
  return mode;
}

export const themeIcon = () => (currentTheme() === 'dark' ? ICON_MOON : ICON_SUN);

export function applyTheme() {
  if (typeof document === 'undefined') return;
  let mode = themeOverride || 'auto';
  if (!themeOverride) {
    try { if (typeof window !== 'undefined' && window.localStorage) mode = window.localStorage.getItem('pgext.theme') || 'auto'; } catch (e) {}
  }
  if (mode === 'auto') document.documentElement.removeAttribute('data-theme');
  else document.documentElement.setAttribute('data-theme', mode);
}

/* ---------------- cards & tables ---------------- */
export function cardBadgeHTML(e) {
  const link = e.repoUrl || e.url;
  const isGH = /github\.com/i.test(e.repoUrl || '');
  if (link) {
    return '<a class="card-gh" href="' + esc(link) + '" target="_blank" rel="noopener" aria-label="' + esc(e.name) + ' upstream">'
      + (isGH ? GH_SVG : EXTLINK_SVG)
      + (e.stars != null ? '<b>' + fmtNum(e.stars) + '</b>' : '') + '</a>';
  }
  return e.stars != null ? '<span class="card-gh plain">' + GH_SVG + '<b>' + fmtNum(e.stars) + '</b></span>' : '';
}

export const hueTag = (href, value, hue, external) =>
  '<a class="tg' + (hue ? ' tg-hue' : '') + '" ' + (hue ? hueVar(hue) + ' ' : '')
  + 'href="' + esc(href) + '"' + (external ? ' target="_blank" rel="noopener"' : '') + '>' + esc(value) + '</a>';

export const licenseTagHref = e => mdSafeURL(e.licenseUrl) || '/license/' + encodeURIComponent(e.license);

export function cardTagsHTML(e) {
  return [
    '<a class="tg tg-cat" href="/?cate=' + encodeURIComponent(e.cat) + '">' + esc(e.cat) + '</a>',
    e.lang ? hueTag('/lang/' + encodeURIComponent(e.lang), e.lang, langHue(e.lang)) : '',
    e.license !== 'Unknown' ? hueTag(licenseTagHref(e), e.license, licenseHue(e.license), !!mdSafeURL(e.licenseUrl)) : '',
    e.repo !== 'n/a' ? hueTag('/repo/' + encodeURIComponent(e.repo), e.repo, repoHue(e.repo)) : '',
    e.vendor ? '<a class="tg tg-dim" href="/?vendor=' + encodeURIComponent(e.vendor) + '">' + esc(e.vendor) + '</a>' : '',
    e.lifecycle && e.lifecycle !== 'active' ? '<a class="tg tg-life" href="/?lifecycle=' + encodeURIComponent(e.lifecycle) + '">' + esc(e.lifecycle) + '</a>' : ''
  ].filter(Boolean).join('');
}

export function tileHTML(e) {
  return '<li><article class="ecard' + (e.avail ? ' packaged' : '') + '" ' + catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '">'
    + '<a class="card-cover" href="' + extHref(e.name) + '" aria-label="' + esc(e.name) + '"></a>'
    + '<header class="card-head"><h3 class="card-name">' + esc(e.name) + '</h3>' + cardBadgeHTML(e) + '</header>'
    + '<p class="card-sub"><a href="' + pkgHref(e.pkg) + '">' + esc(e.pkg) + '</a>'
    + (e.ver ? '<span class="ver">' + esc(e.ver) + '</span>' : '') + '</p>'
    + '<p class="card-desc">' + esc(desc(e)) + '</p>'
    + '<footer class="card-tags">' + cardTagsHTML(e) + '</footer>'
    + '</article></li>';
}

export function packageTileHTML(g) {
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

export function licenseCellHTML(e) {
  const url = mdSafeURL(e.licenseUrl);
  if (e.license === 'Unknown') return '—';
  return url ? '<a href="' + esc(url) + '" target="_blank" rel="noopener">' + esc(e.license) + '</a>' : esc(e.license);
}

export function rowHTML(e) {
  return '<tr ' + catVar(e.cat) + ' data-hover-ext="' + esc(e.name) + '"><td class="r-name"><a href="' + extHref(e.name) + '">' + esc(e.name) + '</a></td>'
    + '<td class="r-mono r-ver">' + esc(e.ver || '—') + '</td>'
    + '<td class="r-desc">' + esc(desc(e)) + '</td>'
    + '<td class="r-cat">' + esc(e.cat) + '</td>'
    + '<td class="r-mono r-lic" title="' + esc(e.license) + '">' + licenseCellHTML(e) + '</td>'
    + '<td class="r-mono">' + esc(e.lang) + '</td>'
    + '<td class="r-mono">' + esc(pgRange(e.pg)) + '</td>'
    + '<td class="r-num">' + (e.stars ? fmtNum(e.stars) : '') + '</td></tr>';
}

export function packageRowHTML(g) {
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

export function extTableHTML(members) {
  return '<div class="rows"><div class="rows-scroll"><table class="ext-table"><thead><tr>'
    + '<th>' + t('rows.name') + '</th><th>' + t('rows.ver') + '</th><th>' + t('rows.desc') + '</th><th>' + t('rows.cat') + '</th>'
    + '<th>' + t('rows.license') + '</th><th>' + t('rows.lang') + '</th><th>' + t('rows.pg') + '</th><th>' + t('rows.stars') + '</th>'
    + '</tr></thead><tbody>' + members.map(rowHTML).join('') + '</tbody></table></div></div>';
}

export function attrChips(e) {
  const chips = [];
  if (e.flags & F.DDL) chips.push(['DDL', 'attr.ddl']);
  if (e.flags & F.LOAD) chips.push(['LOAD', 'attr.load']);
  if (e.flags & F.LIB) chips.push(['LIB', 'attr.lib']);
  if (e.flags & F.BIN) chips.push(['BIN', 'attr.bin']);
  if (e.flags & F.TRUSTED) chips.push(['TRUST', 'attr.trusted']);
  if (e.flags & F.RELOC) chips.push(['RELOC', 'attr.reloc']);
  return chips.map(([code, key]) => '<span class="attr-chip" data-tip="' + esc(t(key)) + '">' + code + '</span>').join('');
}

export function pageTocHTML(items) {
  return '<nav class="page-toc" aria-label="' + esc(t('ext.onpage')) + '"><span>' + t('ext.onpage') + '</span>'
    + items.map(([id, label]) => '<button data-scroll="' + id + '">' + label + '</button>').join('') + '</nav>';
}

export function skel(lines) {
  let bars = '';
  for (let i = 0; i < (lines || 3); i++) bars += '<div class="bar" style="width:' + (55 + ((i * 17) % 40)) + '%"></div>';
  return '<div class="skel">' + bars + '</div>';
}

export const hydrateErr = err => '<div class="hydrate-err">' + esc(t('hydrate.err', { msg: (err && err.message) || 'unknown' })) + '</div>';

export function notFoundHTML(name) {
  return '<div class="page wrap"><div class="section">'
    + '<pre class="mono" style="font-size:14px;line-height:1.8;color:var(--bad)">ERROR:  extension "' + esc(name) + '" does not exist\n'
    + '<span style="color:var(--ink-3)">' + esc(t('notfound.hint')) + '</span></pre>'
    + '<p style="margin-top:18px"><a href="/">' + t('notfound.back') + '</a></p>'
    + '</div></div>';
}

/* ---------------- nav & footer ---------------- */
export function navHTML(active) {
  const zh = getLang() === 'zh';
  return '<div class="wrap nav-in">'
    + '<a class="brand" href="/">' + PIGSTY_LOGO + '<span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<nav class="nav-links">'
    + '<span class="nav-drop"><a href="/matrix" aria-current="' + (active === 'matrix') + '">' + t('nav.matrix') + '</a>'
    + '<div class="nav-menu"><a class="nav-menu-all" href="/matrix">' + bi('Overview', '全景总览') + '</a>'
    + catalog.oss.map(os => '<a href="/os/' + encodeURIComponent(os) + '">' + esc(os) + '</a>').join('')
    + '<span class="nav-menu-sep"></span>'
    + '<div class="nav-menu-pgs">' + catalog.pgs.map(pg => '<a href="/pg/' + pg + '">PG' + pg + '</a>').join('') + '</div>'
    + '<span class="nav-menu-sep"></span>'
    + '<div class="nav-menu-pgs">' + ['PGDG', 'PIGSTY', 'CONTRIB', 'MIXED'].map(r => '<a href="/repo/' + r + '">' + r + '</a>').join('') + '</div>'
    + '</div></span>'
    + '<span class="nav-drop"><a href="/list" aria-current="' + (active === 'browse') + '">' + t('nav.browse') + '</a>'
    + '<div class="nav-menu nav-menu-index">'
    + '<div class="nav-idx-cats">' + catalog.catOrder.map(c => '<a href="' + catHref(c) + '" style="--seg:var(--c-' + c + ')" data-tip="' + esc(catName(c)) + '"><i></i>' + c + '</a>').join('') + '</div>'
    + DIM_GROUPS.map(([, dims]) => '<div class="nav-idx-col">'
      + dims.map(d => '<a href="' + dimHref(d) + '">' + t(DIMS[d].label) + '</a>').join('') + '</div>').join('')
    + '</div></span>'
    + '<a href="/about" aria-current="' + (active === 'about') + '">' + t('nav.about') + '</a>'
    + '<a href="https://doc.pgext.cloud" target="_blank" rel="noopener">' + t('nav.docs') + '</a>'
    + '</nav><span class="nav-spacer"></span><div class="nav-actions">'
    + '<div class="nav-search"><input id="nav-q" type="search" autocomplete="off" spellcheck="false" placeholder="'
    + esc(bi('Press / to search', '按 / 开始搜索')) + '" aria-label="search extensions"><kbd>/</kbd>'
    + '<div class="nav-sugg" id="nav-sugg" hidden></div></div>'
    + '<button class="nav-ico" id="random-ext" aria-label="random extension" data-tip="' + esc(bi('Random extension', '随机看一个扩展')) + '">' + ICON_DICE + '</button>'
    + '<button class="nav-ico" data-lang-toggle aria-label="' + (zh ? 'switch to English' : '切换到中文') + '">' + ICON_LANG + '</button>'
    + '<button class="nav-ico" id="theme-toggle" aria-label="switch theme">' + themeIcon() + '</button>'
    + '<a class="nav-ico nav-github" href="https://github.com/pgsty/pgext" target="_blank" rel="noopener" aria-label="GitHub">' + GH_SVG.replace('width="13" height="13"', 'width="16" height="16"') + '</a>'
    + '</div></div>';
}

export function footerHTML() {
  const zh = getLang() === 'zh';
  const group = (title, links) => '<div class="footer-col"><span>' + title + '</span>'
    + links.map(([href, label, ext]) => '<a href="' + href + '"' + (ext ? ' target="_blank" rel="noopener"' : '') + '>' + label + '</a>').join('') + '</div>';
  return '<div class="wrap footer-in">'
    + '<div class="footer-brand"><a class="brand" href="/">' + PIGSTY_LOGO + '<span class="brand-name">PGEXT<span class="tld">.CLOUD</span></span></a>'
    + '<p>' + bi('PostgreSQL Extension Catalog.', 'PostgreSQL 扩展目录') + '</p>'
    + '<span class="footer-snap">snapshot ' + esc(catalog.meta.generated || '—') + '</span></div>'
    + group(bi('Catalog', '目录'), [['/matrix', t('nav.matrix')], ['/list', t('nav.browse')], ['/about', t('nav.about')]])
    + group(bi('Resources', '资源'), [['/api/v1/meta', 'API', true], ['https://github.com/pgsty/pgext', 'GitHub', true], ['https://github.com/pgsty/pgext/blob/main/db/universe.csv', bi('Raw Data', '原始数据'), true], ['https://pg.center/docs/reference/', bi('PG Center', 'PG 知识图谱'), true], ['https://pgsql.cc/docs/', bi('PGSQL.CC', 'PG 中文文档'), true]])
    + group('Pigsty', [['https://pigsty.io', 'pigsty.io', true], ['https://pigsty.cc', 'pigsty.cc', true], [bi('https://pig.pgsty.com', 'https://pig.pgsty.com/zh'), 'pig.pgsty.com', true], [bi('https://ext.pgsty.com/', 'https://ext.pgsty.com/zh/'), bi('Extension documentation', '扩展文档'), true]])
    + '</div>'
    + '<div class="wrap footer-bottom">'
    + '<span>© 2018-2026 <a href="https://pigsty.io" target="_blank" rel="noopener">Pigsty</a> · <a href="https://vonng.com/en" target="_blank" rel="noopener">Ruohang Feng</a></span>'
    + '<span class="footer-langs"><button data-lang-set="zh" aria-pressed="' + zh + '">中文</button><i>·</i>'
    + '<button data-lang-set="en" aria-pressed="' + !zh + '">English</button></span></div>';
}

/* ---------------- tooltip ---------------- */
const tip = () => document.getElementById('tip');

export function showTip(html, x, y) {
  const el = tip();
  if (!el) return;
  el.innerHTML = html;
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

export function hideTip() {
  const el = tip();
  if (el) el.classList.remove('show');
}

/* ---------------- extension / package hovercard ---------------- */
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

export function hovercardBodyHTML(e, full) {
  const link = (label, url) => url
    ? '<a href="' + esc(url) + '" target="_blank" rel="noopener"><span>' + label + '</span><b>'
      + esc(url.replace(/^https?:\/\//, '').replace(/\/$/, '')) + '</b><i>↗</i></a>' : '';
  const links = full ? [
    link(t('link.repo'), mdSafeURL(full.repo_url)), link(t('link.home'), mdSafeURL(full.home_url)),
    link(t('link.docs'), mdSafeURL(full.doc_url)), link('PGXN', mdSafeURL(full.pgxn_url)),
    link(t('link.license'), mdSafeURL(full.license_url)), link(t('link.control'), mdSafeURL(full.control_url))
  ].filter(Boolean).join('') : link(t('link.repo'), mdSafeURL(e.repoUrl || e.url));
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
  return '<p class="hc-desc">' + esc(desc(e)) + '</p>'
    + '<dl class="hc-facts">' + facts.map(([k, v]) => '<div><dt>' + k + '</dt><dd>' + v + '</dd></div>').join('') + '</dl>'
    + (links ? '<div class="hc-links">' + links + '</div>' : '');
}

export function hovercardHTML(e, full) {
  return '<header ' + catVar(e.cat) + '><a class="hc-title" href="' + extHref(e.name) + '">' + esc(e.name) + '</a>'
    + cardBadgeHTML(e) + '</header>'
    + (e.avail ? '<a class="hc-sub" href="' + pkgHref(e.pkg) + '">' + bi('Package: ', '扩展包：') + esc(e.pkg) + '</a>' : '')
    + hovercardBodyHTML(e, full);
}

export function pkgHovercardHTML(pkg, members, lead, full) {
  const rest = members.filter(x => x.name !== lead.name).sort((a, b) => a.name.localeCompare(b.name));
  const rows = [lead, ...rest].map(x =>
    '<div><a href="' + extHref(x.name) + '">' + esc(x.name) + '</a><span>' + esc(x.ver || '—') + '</span>'
    + '<em>' + attrChips(x) + '</em></div>').join('');
  return '<header ' + catVar(lead.cat) + '><a class="hc-title" href="' + pkgHref(pkg) + '">' + esc(pkg) + '</a>'
    + cardBadgeHTML(lead) + '</header>'
    + hovercardBodyHTML(lead, full)
    + '<div class="hc-exts">' + rows + '</div>';
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

function showHovercard(anchor, key) {
  const isPkg = key.startsWith('p:');
  const name = key.slice(2);
  const members = isPkg ? catalog.byPkg.get(name) : null;
  if (isPkg && (!members || !members.length)) return;
  const lead = isPkg ? (catalog.byName.get(members[0].lead) || members[0]) : catalog.byName.get(name);
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

export function hideHovercard() {
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

export function bindHovercard() {
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
