/* ================================================================
   PGEXT.CLOUD — SPA entry point, History API router, event binding
   and warm-start bootstrap.
   ================================================================ */
import { esc, fmtInt, extHref, catVar, debounce, copyText } from './util.js';
import { getLang, setLangValue, t, desc } from './i18n.js';
import { catalog, matrixState, S, parsePGs, decodeBoot, parseRoute, pushState, readState } from './state.js';
import { DIMS, navSuggest } from './filter.js';
import { navHTML, footerHTML, notFoundHTML, showTip, hideTip, hideHovercard, bindHovercard, currentTheme, themeIcon, applyTheme, setThemeOverride } from './components.js';
import { MATRIX_REPO_VIEWS, globalMatrixShellHTML, hydrateGlobalMatrix, hydrateSliceMatrix, hydrateRepoMatrix, destroyMatrixView, rerenderMatrixView } from './matrix.js';
import { extHTML, pkgHTML, hydrateExt, hydratePkg } from './detail.js';
import { homeHTML, dynamicHTML, drawField, setFieldHover, fieldHit, navPageHTML, browseHTML, dimHTML, aboutHTML } from './views.js';

/* ---------------- Google Analytics (GA4) ---------------- */
const GA_ID = 'G-5HDLR6CMB4';
const GA_ON = typeof location !== 'undefined' && !['localhost', '127.0.0.1', '[::1]'].includes(location.hostname);
let gaLastPath = null;
if (typeof window !== 'undefined') window.dataLayer = window.dataLayer || [];
function gtag() { window.dataLayer.push(arguments); }
if (GA_ON && typeof document !== 'undefined') {
  gtag('js', new Date());
  gtag('config', GA_ID, { send_page_view: false });
  const ga = document.createElement('script');
  ga.async = true;
  ga.src = 'https://www.googletagmanager.com/gtag/js?id=' + GA_ID;
  document.head.appendChild(ga);
}
function gaPageView(path) {
  if (!GA_ON || path === gaLastPath) return;
  gaLastPath = path;
  gtag('event', 'page_view', { page_location: location.href, page_title: document.title });
}

/* ---------------- router ---------------- */
const DIM_ALIASES = { cate: 'category', cat: 'category', lic: 'license', language: 'lang', pkg: 'package', type: 'kind' };
let currentPath = null;

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

function titleFor(path) {
  if (path.startsWith('/ext/')) return decodeURIComponent(path.slice(5)) + ' · PGEXT.CLOUD';
  if (path.startsWith('/pkg/')) return decodeURIComponent(path.slice(5)) + ' package · PGEXT.CLOUD';
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

function route() {
  const { path, params } = parseRoute();
  if (path.startsWith('/e/')) return navigateTo('/ext/' + path.slice(3) + location.search, true);
  if (path.startsWith('/p/')) return navigateTo('/pkg/' + path.slice(3) + location.search, true);
  if (path.startsWith('/c/')) return navigateTo('/cate/' + path.slice(3) + location.search, true);
  if (path.startsWith('/cat/')) return navigateTo('/cate/' + path.slice(5) + location.search, true);
  if (path === '/browse') return navigateTo('/list' + location.search, true);
  if (path.startsWith('/dim/')) return navigateTo('/list/' + path.slice(5) + location.search, true);

  hideHovercard();
  const app = document.getElementById('app');
  const nav = document.getElementById('nav');
  const pathChanged = path !== currentPath;
  currentPath = path;
  if (path !== '/matrix') { destroyMatrixView(); matrixState.hydSeq++; }

  let active = 'home';
  if (path === '/' || path === '') {
    readState(params);
    app.innerHTML = homeHTML();
    drawField(true);
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
    hydrateGlobalMatrix(() => currentPath === '/matrix');
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

  const zh = getLang() === 'zh';
  let description = zh ? 'PostgreSQL 扩展、项目包族、依赖关系与 PG/OS 精确可用性目录。' : 'Search PostgreSQL extensions, package families, dependencies, and exact PG/OS availability.';
  if (path.startsWith('/ext/')) { const e = catalog.byName.get(decodeURIComponent(path.slice(5))); if (e) description = desc(e); }
  if (path.startsWith('/pkg/')) { const family = catalog.byPkg.get(decodeURIComponent(path.slice(5))); if (family && family[0]) description = desc(catalog.byName.get(family[0].lead) || family[0]); }
  if (path === '/matrix') description = zh ? '跨扩展包、操作系统与 PostgreSQL 大版本的全局构建矩阵。' : 'The global build matrix across extension packages, operating systems, and PostgreSQL majors.';
  if (path.startsWith('/repo/')) {
    const repoName = decodeURIComponent(path.slice(6)).toUpperCase();
    if (MATRIX_REPO_VIEWS[repoName]) description = zh
      ? '假设只剩 ' + repoName + ' 仓库时的独立构建矩阵：它实际交付哪些扩展包，又会缺失多少构建槽位。'
      : 'The standalone ' + repoName + ' build matrix: exactly what this repository ships, and how much would go missing without the other one.';
  }
  const meta = document.querySelector('meta[name="description"]');
  const ogTitle = document.querySelector('meta[property="og:title"]');
  const ogDesc = document.querySelector('meta[property="og:description"]');
  if (meta) meta.content = description;
  if (ogTitle) ogTitle.content = document.title;
  if (ogDesc) ogDesc.content = description;
  gaPageView(path);
}

function renderDynamic() {
  const d = document.getElementById('dynamic');
  if (d) { d.innerHTML = dynamicHTML(); pushState(); }
}

function bindNavSearch() {
  const input = document.getElementById('nav-q');
  const box = document.getElementById('nav-sugg');
  if (!input || !box) return;
  let items = [], active = -1;
  const close = () => { box.hidden = true; box.innerHTML = ''; items = []; active = -1; };
  const paint = () => box.querySelectorAll('.ns-item').forEach((el, i) => el.classList.toggle('on', i === active));
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
  q.addEventListener('input', debounce(() => { S.q = q.value; renderDynamic(); }, 130));
  q.addEventListener('keydown', ev => {
    if (ev.key === 'Escape') { q.value = ''; S.q = ''; renderDynamic(); }
  });
}

/* ---------------- theme & language controls ---------------- */
function cycleTheme() {
  const next = currentTheme() === 'dark' ? 'light' : 'dark';
  setThemeOverride(null);
  try { localStorage.setItem('pgext.theme', next); } catch (e) {}
  applyTheme();
  const b = document.getElementById('theme-toggle');
  if (b) b.innerHTML = themeIcon();
  drawField(true);
  rerenderMatrixView();
}

function setLang(next) {
  if ((next !== 'zh' && next !== 'en') || next === getLang()) return;
  setLangValue(next);
  try { localStorage.setItem('pgext.lang', next); } catch (e) {}
  document.documentElement.lang = next === 'zh' ? 'zh-CN' : 'en';
  const y = window.scrollY;
  currentPath = null;
  route();
  window.scrollTo(0, y);
}

/* ---------------- global events ---------------- */
function attachEvents() {
  bindHovercard();
  window.addEventListener('popstate', route);
  window.addEventListener('hashchange', route);
  window.addEventListener('resize', debounce(() => { drawField(false); rerenderMatrixView(); }, 150));
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
    const el = ev.target.closest('[data-fkey],[data-pg-toggle],[data-entity],[data-layout],[data-sql],[data-copy],[data-scroll],[data-itab],[data-ftab],[data-fall],[data-target-ext],[data-target-pkg],[data-lang-toggle],[data-lang-set],#random-ext,#theme-toggle,#ufield');
    if (!el) return;
    if (el.dataset.langSet) return setLang(el.dataset.langSet);
    if (el.dataset.langToggle !== undefined) return setLang(getLang() === 'zh' ? 'en' : 'zh');
    if (el.id === 'theme-toggle') return cycleTheme();
    if (el.id === 'random-ext') {
      const packed = catalog.ext.filter(x => x.avail);
      const rest = catalog.ext.filter(x => !x.avail);
      const pool = Math.random() * (packed.length * 5 + rest.length) < packed.length * 5 && packed.length ? packed : rest;
      const pick = (pool.length ? pool : catalog.ext)[Math.floor(Math.random() * (pool.length || catalog.ext.length))];
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
    if (el.dataset.fkey) {
      const k = el.dataset.fkey, v = el.dataset.fval;
      S[k] = S[k] === v ? '' : v;
      renderDynamic(); return;
    }
    if (el.dataset.pgToggle !== undefined) {
      const pg = Number.parseInt(el.dataset.pgToggle, 10);
      const selected = new Set(parsePGs(S.pg));
      if (!pg) selected.clear();
      else if (selected.has(pg)) selected.delete(pg);
      else selected.add(pg);
      S.pg = [...selected].sort((a, b) => b - a).join(',');
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
    }
  });

  document.addEventListener('change', ev => {
    const el = ev.target.closest('[data-skey]');
    if (!el) return;
    S[el.dataset.skey] = el.value;
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
      const target = (path === '/' || path === '') ? document.getElementById('q') : document.getElementById('nav-q');
      if (target) target.focus();
    }
  });

  document.addEventListener('mousemove', ev => {
    const cv = ev.target.id === 'ufield' ? ev.target : null;
    if (cv) {
      const e = fieldHit(ev);
      if (e) {
        setFieldHover(e.name);
        showTip('<b>' + esc(e.name) + '</b> <span style="color:var(--c-' + e.cat + ')">' + e.cat + '</span><br><span class="d">' + esc(desc(e)) + '</span>', ev.clientX, ev.clientY);
        cv.style.cursor = 'pointer';
      } else {
        setFieldHover(null);
        hideTip();
        cv.style.cursor = 'crosshair';
      }
      return;
    }
    const el = ev.target.closest('[data-tip]');
    if (el && el.dataset.tip) showTip(esc(el.dataset.tip), ev.clientX, ev.clientY);
    else hideTip();
  });
  document.addEventListener('mouseout', ev => { if (!ev.relatedTarget) hideTip(); });
}

/* ---------------- boot ---------------- */
const BOOT_CACHE_KEY = 'pgext.boot.v5';

function bootError(err) {
  const app = document.getElementById('app');
  app.innerHTML = '<div class="boot"><div class="boot-err">'
    + '<span class="e">psql: error:</span> ' + esc(t('boot.err')) + '\n'
    + '<span class="muted">DETAIL:  ' + esc((err && err.message) || 'network error') + '</span>\n'
    + '<button class="chip" id="boot-retry">↻ ' + esc(t('boot.retry')) + '</button>'
    + '</div></div>';
  document.getElementById('boot-retry').addEventListener('click', boot);
}

async function boot() {
  applyTheme();
  document.documentElement.lang = getLang() === 'zh' ? 'zh-CN' : 'en';
  let cached = null;
  try { cached = JSON.parse(localStorage.getItem(BOOT_CACHE_KEY) || 'null'); } catch (e) {}
  if (cached && Array.isArray(cached.rows) && cached.rows.length && cached.rows[0].length >= 34) {
    try { decodeBoot(cached); route(); } catch (e) { cached = null; }
  } else { cached = null; }
  try {
    const headers = { Accept: 'application/json' };
    if (cached && cached.version) headers['If-None-Match'] = '"b34-' + cached.version + '"';
    const res = await fetch('/api/v1/bootstrap?fmt=b34', { headers });
    if (res.status === 304) return;
    if (!res.ok) throw new Error('HTTP ' + res.status);
    const b = await res.json();
    const hadCache = Boolean(cached);
    decodeBoot(b);
    try { localStorage.setItem(BOOT_CACHE_KEY, JSON.stringify(b)); } catch (e) {}
    const y = window.scrollY;
    currentPath = null;
    route();
    if (hadCache) window.scrollTo(0, y);
  } catch (err) {
    if (!cached) bootError(err);
  }
}

attachEvents();
boot();
