/* ================================================================
   Global, sliced (/pg, /os) and single-repo what-if build matrices.
   Payload format matrix-row.v2 (see server/matrix.go).
   ================================================================ */
import { esc, fmtInt, pkgHref, debounce } from './util.js';
import { t, bi } from './i18n.js';
import { catalog, matrixState, j } from './state.js';
import { skel, hydrateErr, showTip, hideTip } from './components.js';

export const MATRIX_CACHE_KEY = 'pgext.matrix.v2';
export const MATRIX_REPO_VIEWS = { PGDG: 'pgdg', PIGSTY: 'pigsty' };

export const GMX_META = [
  ['G', 'gmx.pgdg', 'pgdg'],
  ['P', 'gmx.pigsty', 'pigsty'],
  ['M', 'gmx.missing', 'missing'],
  ['N', 'gmx.na', 'na']
];
const GMX_CLASS = { G: 'pgdg', P: 'pigsty', M: 'missing', N: 'na' };
const gmxClassOf = ch => GMX_CLASS[String(ch).toUpperCase()] || 'na';
const GMX_CELL = 15, GMX_ROWH = 15;

export const matrixUsable = data =>
  Boolean(data && data.format === 'matrix-row.v2' && Array.isArray(data.rows) && data.rows.length);

export function destroyMatrixView() {
  if (matrixState.view) { matrixState.view.destroy(); matrixState.view = null; }
}

export function rerenderMatrixView() {
  if (matrixState.view) matrixState.view.render();
}

/* Unified cache-first loader for global ('') and single-repo ('pgdg'|'pigsty')
   matrices: renders localStorage immediately when warm, then revalidates once
   per session (or after any catalog bootstrap version change). */
export async function loadMatrix(view = '', opts = {}) {
  const storageKey = view ? MATRIX_CACHE_KEY + '.' + view : MATRIX_CACHE_KEY;
  const apiURL = view ? '/api/v1/matrix?repo=' + view : '/api/v1/matrix';
  const valid = data => matrixUsable(data) && (!view || data.repo === view);

  if (!matrixState.data[view]) {
    try {
      if (typeof localStorage !== 'undefined') {
        const cached = JSON.parse(localStorage.getItem(storageKey) || 'null');
        if (valid(cached)) matrixState.data[view] = cached;
      }
    } catch (e) {}
  }
  if (matrixState.data[view] && opts.onPaint) opts.onPaint(matrixState.data[view]);
  if (matrixState.checked[view] && matrixState.data[view]) return matrixState.data[view];

  try {
    if (!matrixState.inflight[view]) {
      matrixState.inflight[view] = j(apiURL).finally(() => { delete matrixState.inflight[view]; });
    }
    const fresh = await matrixState.inflight[view];
    if (!valid(fresh)) throw new Error('Invalid matrix response');
    matrixState.data[view] = fresh;
    matrixState.checked[view] = true;
    try { if (typeof localStorage !== 'undefined') localStorage.setItem(storageKey, JSON.stringify(fresh)); } catch (e) {}
    if (opts.onPaint) opts.onPaint(fresh);
    return fresh;
  } catch (err) {
    if (matrixState.data[view]) return matrixState.data[view];
    throw err;
  }
}

export const ensureGlobalMatrix = () => loadMatrix('');

export function globalMatrixShellHTML() {
  const combos = catalog.oss.length * catalog.pgs.length;
  return '<article class="page wrap page-wide gmx-page">'
    + '<header class="page-head"><p class="eyebrow">' + t('gmx.eyebrow') + '</p>'
    + '<div class="gmx-titleline"><div><h1>' + t('gmx.title') + '</h1>'
    + '<p class="gmx-lede">' + t('gmx.lede', { pkgs: fmtInt(catalog.nPkgs), combos: fmtInt(combos) }) + '</p></div>'
    + '<a class="gmx-api" href="/api/v1/matrix" target="_blank" rel="noopener">' + t('gmx.api') + ' ↗</a></div></header>'
    + '<div id="gmx-root">' + skel(7) + '</div></article>';
}

export function matrixBlockHTML(data, rows) {
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
  const legendMeta = data.repo ? GMX_META.filter(item => counts[item[0]] || item[0] === 'M' || item[0] === 'N') : GMX_META;
  const legend = legendMeta.map(item =>
    '<button type="button" class="gmx-legend-item gmx-' + item[2] + '" data-gmx-code="' + item[0]
    + '" aria-pressed="false"><i></i><span>' + t(item[1]) + '</span><b>' + fmtInt(counts[item[0]] || 0) + '</b></button>').join('');
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
    + '<p class="gmx-source"><span class="gmx-source-copy">' + t('gmx.source', { source: '<code>' + esc(data.source || 'pgext.pkg') + '</code>' })
    + '</span><span class="gmx-snapshot">snapshot ' + esc((data.generated || '').replace('T', ' ').slice(0, 19)) + '</span></p>';
}

export function gmxTipHTML(row, ch, os, pg, version) {
  const slim = catalog.byName.get(row.e) || catalog.byName.get(row.p);
  const isDeb = os[0] === 'd' || os[0] === 'u';
  const pattern = slim ? (isDeb ? slim.debPkg : slim.rpmPkg) : '';
  const pkgName = pattern ? pattern.split(/\s+/)[0].replaceAll('$v', String(pg)) : '';
  const meta = GMX_META.find(x => x[0] === ch);
  return '<b>' + esc(row.p) + '</b> <span class="gmx-tip-state gmx-' + gmxClassOf(ch) + '">' + (meta ? t(meta[1]) : 'N/A') + '</span><br>'
    + '<span class="d">' + esc(row.e) + ' · ' + esc(os) + ' · PG ' + esc(String(pg)) + '</span>'
    + (pkgName ? '<br><span class="k">package</span> ' + esc(pkgName) : '')
    + (version ? '<br><span class="k">version</span> ' + esc(version) : '');
}

export function setupGlobalMatrix(data, opts) {
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
    ctx.fillStyle = color('--line-2') || '#d3d2c9';
    for (let g = 1; g < (data.os || []).length; g++) {
      ctx.fillRect(g * pgCount * GMX_CELL - 1, 0, 1, H);
    }
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
    labels.innerHTML = state.rows.map(ref =>
      '<a class="gmx-row-name" href="' + pkgHref(ref.row.p) + '">' + esc(ref.row.p) + '</a>').join('');
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

  canvas.addEventListener('mousemove', ev => {
    const c = cellAt(ev);
    if (c) showTip(gmxTipHTML(c.row, c.ch, c.os, c.pg, c.version), ev.clientX, ev.clientY);
    else hideTip();
  });
  canvas.addEventListener('mouseleave', hideTip);
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
  matrixState.view = { render: draw, state, destroy() {} };
  applyFilter();
}

export function osNavHTML(current) {
  const bases = [];
  for (const os of catalog.oss) { const b = os.split('.')[0]; if (!bases.includes(b)) bases.push(b); }
  const cell = os => catalog.oss.includes(os)
    ? '<a class="os-cell" href="/os/' + encodeURIComponent(os) + '" aria-current="' + (os === current) + '">' + esc(os) + '</a>'
    : '<span class="os-cell off"></span>';
  return '<nav class="os-nav" style="--os-cols:' + Math.max(bases.length, 1) + '">'
    + bases.map(b => cell(b + '.x86_64')).join('') + bases.map(b => cell(b + '.aarch64')).join('') + '</nav>';
}

export function sliceMatrixHTML(data, fix) {
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

function soloLedeHTML(data, org) {
  const stats = data.stats || {}, counts = stats.counts || {}, g = stats.global || {};
  if (!g.avail) return '';
  const avail = (counts.G || 0) + (counts.P || 0);
  return esc(t('gmx.solo', {
    org,
    pkgs: fmtInt((data.rows || []).length),
    allPkgs: fmtInt(g.rows || 0),
    avail: fmtInt(avail),
    allAvail: fmtInt(g.avail),
    miss: fmtInt(Math.max((g.avail || 0) - avail, 0)),
    gone: fmtInt(Math.max((g.rows || 0) - (data.rows || []).length, 0))
  }));
}

export async function hydrateGlobalMatrix(isActive = () => true) {
  const token = ++matrixState.hydSeq;
  const paint = data => {
    if (token !== matrixState.hydSeq || !isActive()) return;
    const root = document.getElementById('gmx-root');
    if (!root) return;
    destroyMatrixView();
    const rows = data.rows.map((row, index) => ({ row, index }));
    root.innerHTML = matrixBlockHTML(data, rows);
    setupGlobalMatrix(data, { root, rows });
  };
  try {
    await loadMatrix('', { onPaint: paint });
  } catch (err) {
    if (!matrixState.data[''] && token === matrixState.hydSeq) {
      const root = document.getElementById('gmx-root');
      if (root) root.innerHTML = '<div class="gmx-frame">' + hydrateErr(err) + '</div>';
    }
  }
}

export async function hydrateSliceMatrix(kind, value) {
  if (!document.getElementById('slice-matrix')) return;
  try {
    const data = await ensureGlobalMatrix();
    const box = document.getElementById('slice-matrix');
    if (box) box.innerHTML = sliceMatrixHTML(data, kind === 'pg' ? { pg: Number(value) } : { os: String(value).toLowerCase() });
  } catch (err) {
    const box = document.getElementById('slice-matrix');
    if (box) box.innerHTML = hydrateErr(err);
  }
}

export async function hydrateRepoMatrix(org) {
  const box = document.getElementById('repo-matrix');
  if (!box) return;
  const view = MATRIX_REPO_VIEWS[org];
  const token = matrixState.hydSeq;
  if (view) {
    const paint = data => {
      if (token !== matrixState.hydSeq) return;
      const target = document.getElementById('repo-matrix');
      if (!target) return;
      const rows = data.rows.map((row, index) => ({ row, index }));
      target.innerHTML = matrixBlockHTML(data, rows);
      setupGlobalMatrix(data, { root: target, rows });
      const solo = document.getElementById('repo-solo');
      if (solo) { const html = soloLedeHTML(data, org); solo.hidden = !html; solo.innerHTML = html; }
    };
    try {
      await loadMatrix(view, { onPaint: paint });
    } catch (err) {
      if (!matrixState.data[view] && token === matrixState.hydSeq) {
        const target = document.getElementById('repo-matrix');
        if (target) target.innerHTML = hydrateErr(err);
      }
    }
    return;
  }
  try {
    const data = await ensureGlobalMatrix();
    if (token !== matrixState.hydSeq) return;
    const memberPkgs = new Set(catalog.ext.filter(e => e.repo === org).map(e => e.pkg));
    const rows = data.rows.map((row, index) => ({ row, index })).filter(ref => memberPkgs.has(ref.row.p));
    if (!rows.length) {
      box.innerHTML = '<p class="empty-note">' + bi('No build-matrix rows are recorded for this source.', '该来源没有构建矩阵行记录。') + '</p>';
      return;
    }
    box.innerHTML = matrixBlockHTML(data, rows);
    setupGlobalMatrix(data, { root: box, rows });
  } catch (err) {
    box.innerHTML = hydrateErr(err);
  }
}
