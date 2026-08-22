/*
 * The home console.
 *
 * One fetch of `catalog/universe.<lang>.json` puts the whole catalogue —
 * every extension, packaged or not — in memory, and every keystroke and facet
 * click filters it locally. There is no request per query and no index to keep
 * in sync: the payload is built at publish time from the same `db/universe.csv`
 * the extension pages are generated from.
 *
 * The payload is positional and interned; the layout is documented in
 * layouts/_partials/pgext/catalog-data.html and the two must be changed
 * together. Field order here mirrors it exactly.
 */
(function () {
  'use strict';

  var NAME = 0, PKG = 1, CAT = 2, LIC = 3, LANG = 4, REPO = 5,
      VER = 6, PGMASK = 7, STARS = 8, FLAGS = 9, DESC = 10, TAGS = 11, URL = 12;

  var F_PACKAGED = 1, F_CONTRIB = 2;
  var PAGE = 60;

  var root = document.getElementById('pgx-console');
  if (!root) return;

  var zh = root.getAttribute('data-pgx-lang') === 'zh';
  var extBase = root.getAttribute('data-pgx-ext-base') || '/e/';

  var el = {
    q: document.getElementById('pgx-q'),
    form: root.querySelector('.pgx-searchbox'),
    sql: document.getElementById('pgx-sql'),
    facets: document.getElementById('pgx-facets'),
    count: document.getElementById('pgx-n'),
    sort: document.getElementById('pgx-sort'),
    results: document.getElementById('pgx-results'),
    more: document.getElementById('pgx-more')
  };

  var DB = null;
  var state = { q: '', cat: '', repo: '', lang: '', lic: '', pg: '', scope: '', sort: 'stars', view: 'cards' };
  var shown = PAGE;
  var matched = [];

  /* ------------------------------------------------------------------ util */

  function esc(s) {
    return String(s == null ? '' : s).replace(/[&<>"']/g, function (m) {
      return { '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' }[m];
    });
  }

  function int(n) { return String(n).replace(/\B(?=(\d{3})+(?!\d))/g, ','); }

  function compact(n) {
    if (n == null || n === 0) return '';
    return n >= 1000 ? (n / 1000).toFixed(n >= 10000 ? 0 : 1) + 'k' : String(n);
  }

  function t(en, cn) { return zh ? cn : en; }

  /* The dimension hue scale, kept in step with pgext/hue.html and with the
     dynamic catalogue's app.js. Same value, same colour, all three places. */
  function licenceHue(v) {
    var s = (v || '').toLowerCase();
    if (!s || s === 'unknown') return '';
    if (s.indexOf('postgresql') === 0) return 'pgblue';
    if (s.indexOf('mit') === 0) return 'blue';
    if (/^(bsd|0bsd|isc|zlib|unlicense|cc0|wtfpl|public)/.test(s)) return 'cyan';
    if (s.indexOf('apache') === 0) return 'green';
    if (/^(mpl|epl|cddl|osl)/.test(s)) return 'teal';
    if (/^(artistic|eupl|cecill)/.test(s)) return 'olive';
    if (s.indexOf('lgpl') === 0) return 'amber';
    if (s.indexOf('agpl') === 0) return 'red';
    if (s.indexOf('gpl') === 0) return 'orange';
    return 'violet';
  }

  var LANG_HUES = {
    c: 'blue', 'c++': 'indigo', sql: 'green', 'pl/pgsql': 'teal', go: 'teal',
    rust: 'orange', python: 'amber', javascript: 'amber', typescript: 'amber',
    java: 'red', r: 'red', ruby: 'red', shell: 'maroon', perl: 'maroon', data: 'slate'
  };
  var REPO_HUES = { pgdg: 'pgblue', pigsty: 'green', contrib: 'pgnavy', mixed: 'teal' };

  function hueOf(dim, value) {
    if (!value) return '';
    if (dim === 'lic') return licenceHue(value);
    if (dim === 'lang') return LANG_HUES[value.toLowerCase()] || 'red';
    if (dim === 'repo') return REPO_HUES[value.toLowerCase()] || '';
    return '';
  }

  function segStyle(dim, value) {
    if (dim === 'cat') return value ? ' style="--pgx-seg: var(--c-' + esc(value) + ')"' : '';
    var h = hueOf(dim, value);
    return h ? ' style="--pgx-seg: var(--h-' + h + ')"' : '';
  }

  /* ---------------------------------------------------------------- filter */

  function pgMaskFor(major) {
    var i = DB.pgs.indexOf(Number(major));
    return i < 0 ? 0 : 1 << i;
  }

  function pgList(mask) {
    var out = [];
    for (var i = 0; i < DB.pgs.length; i++) if (mask & (1 << i)) out.push(DB.pgs[i]);
    return out;
  }

  function run() {
    var words = state.q.toLowerCase().split(/\s+/).filter(Boolean);
    var catIdx = state.cat ? DB.cats.indexOf(state.cat) : -1;
    var repoIdx = state.repo ? DB.repos.indexOf(state.repo) : -1;
    var langIdx = state.lang ? DB.langs.indexOf(state.lang) : -1;
    var licIdx = state.lic ? DB.lics.indexOf(state.lic) : -1;
    var pgBit = state.pg ? pgMaskFor(state.pg) : 0;

    var scored = [];
    for (var i = 0; i < DB.rows.length; i++) {
      var r = DB.rows[i];
      if (catIdx >= 0 && r[CAT] !== catIdx) continue;
      if (repoIdx >= 0 && r[REPO] !== repoIdx) continue;
      if (langIdx >= 0 && r[LANG] !== langIdx) continue;
      if (licIdx >= 0 && r[LIC] !== licIdx) continue;
      if (pgBit && !(r[PGMASK] & pgBit)) continue;
      if (state.scope === 'packaged' && !(r[FLAGS] & F_PACKAGED)) continue;
      if (state.scope === 'unpacked' && (r[FLAGS] & F_PACKAGED)) continue;
      if (state.scope === 'contrib' && !(r[FLAGS] & F_CONTRIB)) continue;

      var score = 0, drop = false;
      for (var w = 0; w < words.length; w++) {
        var word = words[w];
        var nm = r[NAME].toLowerCase();
        if (nm === word) score += 120;
        else if (nm.indexOf(word) === 0) score += 60;
        else if (nm.indexOf(word) >= 0 || (r[PKG] && r[PKG].toLowerCase().indexOf(word) >= 0)) score += 30;
        else if (r[TAGS] && r[TAGS].toLowerCase().indexOf(word) >= 0) score += 18;
        else if (r[DESC] && r[DESC].toLowerCase().indexOf(word) >= 0) score += 10;
        else { drop = true; break; }
      }
      if (drop) continue;
      scored.push([score, i, r]);
    }

    var hasQ = words.length > 0;
    var sort = state.sort;
    scored.sort(function (a, b) {
      if (hasQ && b[0] !== a[0]) return b[0] - a[0];
      if (sort === 'stars') return (b[2][STARS] - a[2][STARS]) || a[2][NAME].localeCompare(b[2][NAME]);
      if (sort === 'name') return a[2][NAME].localeCompare(b[2][NAME]);
      return a[1] - b[1];
    });

    matched = scored.map(function (x) { return x[2]; });
    shown = PAGE;
  }

  /* ---------------------------------------------------------------- render */

  function tag(text, dim, value, href) {
    if (!text) return '';
    var cls = 'pgx-tg' + (dim ? ' pgx-tg--hued' : '');
    return '<a class="' + cls + '"' + segStyle(dim, value) + ' href="' + esc(href) + '">' + esc(text) + '</a>';
  }

  function cardHTML(r) {
    var cat = r[CAT] >= 0 ? DB.cats[r[CAT]] : '';
    var lic = r[LIC] >= 0 ? DB.lics[r[LIC]] : '';
    var lg = r[LANG] >= 0 ? DB.langs[r[LANG]] : '';
    var repo = r[REPO] >= 0 ? DB.repos[r[REPO]] : '';
    var packaged = !!(r[FLAGS] & F_PACKAGED);
    var href = packaged ? extBase + encodeURIComponent(r[NAME]) + '/' : ('https://' + r[URL]);
    var stars = compact(r[STARS]);

    var tags = [];
    if (cat) tags.push(tag(cat, 'cat', cat, langURL('/categories/' + cat.toLowerCase() + '/')));
    if (lg) tags.push(tag(lg, 'lang', lg, langURL('/languages/' + slug(lg) + '/')));
    if (lic) tags.push(tag(lic, 'lic', lic, langURL('/licenses/' + slug(lic) + '/')));
    if (repo) tags.push(tag(repo, 'repo', repo, langURL('/repos/' + repo.toLowerCase() + '/')));

    return '<li class="pgx-card' + (packaged ? ' is-packaged' : '') + '"' + segStyle('cat', cat) + '>'
      + '<a class="pgx-card-cover" href="' + esc(href) + '"' + (packaged ? '' : ' target="_blank" rel="noopener"') + ' aria-label="' + esc(r[NAME]) + '"></a>'
      + '<div class="pgx-card-head"><span class="pgx-card-name">' + esc(r[NAME]) + '</span>'
      + (stars ? '<span class="pgx-card-stars">★ ' + stars + '</span>' : '') + '</div>'
      + '<div class="pgx-card-sub">' + (r[PKG] ? '<span>' + esc(r[PKG]) + '</span>' : '')
      + (r[VER] ? '<span class="ver">' + esc(r[VER]) + '</span>' : '') + '</div>'
      + '<p class="pgx-card-desc">' + esc(r[DESC]) + '</p>'
      + '<div class="pgx-card-tags">' + tags.join('') + '</div>'
      + '</li>';
  }

  function rowHTML(r) {
    var cat = r[CAT] >= 0 ? DB.cats[r[CAT]] : '';
    var lic = r[LIC] >= 0 ? DB.lics[r[LIC]] : '';
    var lg = r[LANG] >= 0 ? DB.langs[r[LANG]] : '';
    var packaged = !!(r[FLAGS] & F_PACKAGED);
    var href = packaged ? extBase + encodeURIComponent(r[NAME]) + '/' : ('https://' + r[URL]);
    // An unpackaged extension has no verified major-version support recorded,
    // which is a different statement from "supported nowhere".
    var pgs = pgList(r[PGMASK]);
    var strip = r[PGMASK] === 0
      ? '<span class="pgx-v pgx-v--gray">&mdash;</span>'
      : DB.pgs.map(function (pg) {
          var on = pgs.indexOf(pg) >= 0;
          return '<span class="pgx-v pgx-v--' + (on ? 'green' : 'red') + '">' + pg + '</span>';
        }).join('');

    return '<tr>'
      + '<td class="r-name"><a href="' + esc(href) + '"' + (packaged ? '' : ' target="_blank" rel="noopener"') + '>' + esc(r[NAME]) + '</a></td>'
      + '<td class="r-cat"' + segStyle('cat', cat) + '>' + esc(cat) + '</td>'
      + '<td class="r-mono">' + esc(r[VER]) + '</td>'
      + '<td class="r-pg"><span class="pgx-vers">' + strip + '</span></td>'
      + '<td class="r-mono">' + esc(lg) + '</td>'
      + '<td class="r-mono">' + esc(lic) + '</td>'
      + '<td class="r-desc">' + esc(r[DESC]) + '</td>'
      + '</tr>';
  }

  function slug(v) {
    return String(v).toLowerCase().replace(/[^a-z0-9+]+/g, '-').replace(/^-+|-+$/g, '');
  }

  function langURL(path) {
    return zh ? '/zh' + path : path;
  }

  function paint() {
    var list = matched.slice(0, shown);
    if (!list.length) {
      el.results.innerHTML = '<p class="pgx-empty">'
        + t('Nothing matches that query.', '没有匹配的扩展。')
        + ' <code>' + esc(state.q) + '</code></p>';
    } else if (state.view === 'table') {
      el.results.innerHTML = '<div class="pgx-rows"><table>'
        + '<thead><tr><th>' + t('Extension', '扩展') + '</th><th>' + t('Category', '分类')
        + '</th><th>' + t('Version', '版本') + '</th><th>PG</th><th>' + t('Language', '语言')
        + '</th><th>' + t('Licence', '许可证') + '</th><th>' + t('Description', '描述') + '</th></tr></thead>'
        + '<tbody>' + list.map(rowHTML).join('') + '</tbody></table></div>';
    } else {
      el.results.innerHTML = '<ul class="pgx-wall">' + list.map(cardHTML).join('') + '</ul>';
    }
    el.count.textContent = int(matched.length);
    el.more.hidden = shown >= matched.length;
    el.more.textContent = t('load more', '加载更多') + ' (' + int(matched.length - shown) + ')';
    paintSQL();
  }

  /* The readout is the same idea as the dynamic catalogue's: show the query
     the console just answered, so the filters stay legible as a statement. */
  function paintSQL() {
    var where = [];
    var lit = function (v) { return "'" + String(v).replace(/'/g, "''") + "'"; };
    state.q.split(/\s+/).filter(Boolean).forEach(function (w) {
      where.push("concat_ws(' ', name, pkg, tags, description) ILIKE " + lit('%' + w + '%'));
    });
    if (state.cat) where.push('category = ' + lit(state.cat));
    if (state.repo) where.push('repo = ' + lit(state.repo));
    if (state.lang) where.push('lang = ' + lit(state.lang));
    if (state.lic) where.push('license = ' + lit(state.lic));
    if (state.pg) where.push("pg_ver @> '{" + state.pg + "}'");
    if (state.scope === 'packaged') where.push('packaged');
    if (state.scope === 'unpacked') where.push('NOT packaged');
    if (state.scope === 'contrib') where.push('contrib');
    var order = state.sort === 'stars' ? 'stars DESC NULLS LAST'
      : state.sort === 'name' ? 'name' : 'id';
    var sql = 'SELECT * FROM pgext.universe'
      + (where.length ? ' WHERE ' + where.join(' AND ') : '')
      + ' ORDER BY ' + order + ';';
    el.sql.hidden = false;
    el.sql.dataset.sql = sql;
    var html = '<span class="kw">SELECT</span> * <span class="kw">FROM</span> pgext.universe';
    if (where.length) {
      html += ' <span class="kw">WHERE</span> '
        + esc(where.join(' AND ')).replace(/&#39;([^&]*)&#39;/g, '<span class="lit">&#39;$1&#39;</span>');
    }
    html += ' <span class="kw">ORDER BY</span> ' + esc(order) + ';';
    el.sql.innerHTML = '<span class="q">' + html + '</span><span class="n">('
      + int(matched.length) + (matched.length === 1 ? ' row' : ' rows') + ')</span>';
  }

  /* Facet buttons for the interned dimensions are built from the payload, so
     a value that leaves the catalogue leaves the console with it. */
  function buildFacets() {
    [['repo', DB.repos, 8], ['lang', DB.langs, 10], ['lic', DB.lics, 12]].forEach(function (spec) {
      var key = spec[0], values = spec[1], limit = spec[2];
      var host = root.querySelector('[data-facet-values="' + key + '"]');
      if (!host) return;
      var counts = {};
      DB.rows.forEach(function (r) {
        var i = key === 'repo' ? r[REPO] : key === 'lang' ? r[LANG] : r[LIC];
        if (i >= 0) counts[values[i]] = (counts[values[i]] || 0) + 1;
      });
      Object.keys(counts)
        .sort(function (a, b) { return counts[b] - counts[a] || a.localeCompare(b); })
        .slice(0, limit)
        .forEach(function (v) {
          var b = document.createElement('button');
          b.type = 'button';
          b.className = 'pgx-fbtn hued';
          b.dataset.facet = key;
          b.dataset.value = v;
          b.setAttribute('aria-pressed', 'false');
          var hue = hueOf(key, v);
          if (hue) b.style.setProperty('--pgx-seg', 'var(--h-' + hue + ')');
          b.innerHTML = '<i aria-hidden="true"></i>' + esc(v) + '<small>' + counts[v] + '</small>';
          host.appendChild(b);
        });
    });
  }

  function syncFacetButtons() {
    root.querySelectorAll('.pgx-fbtn').forEach(function (b) {
      var on = String(state[b.dataset.facet] || '') === b.dataset.value;
      b.setAttribute('aria-pressed', on ? 'true' : 'false');
    });
    root.querySelectorAll('.pgx-view button').forEach(function (b) {
      b.setAttribute('aria-pressed', b.dataset.view === state.view ? 'true' : 'false');
    });
  }

  /* ------------------------------------------------------------ url state */

  function readURL() {
    var p = new URLSearchParams(location.search);
    ['q', 'cat', 'repo', 'lang', 'lic', 'pg', 'scope', 'sort', 'view'].forEach(function (k) {
      var v = p.get(k);
      if (v != null) state[k] = v;
    });
    if (el.q) el.q.value = state.q;
    if (el.sort) el.sort.value = state.sort;
  }

  function writeURL(replace) {
    var p = new URLSearchParams();
    ['q', 'cat', 'repo', 'lang', 'lic', 'pg', 'scope'].forEach(function (k) {
      if (state[k]) p.set(k, state[k]);
    });
    if (state.sort !== 'stars') p.set('sort', state.sort);
    if (state.view !== 'cards') p.set('view', state.view);
    var url = location.pathname + (p.toString() ? '?' + p : '');
    history[replace ? 'replaceState' : 'pushState'](null, '', url);
  }

  function update(replace) {
    run();
    paint();
    syncFacetButtons();
    writeURL(replace !== false);
  }

  /* ----------------------------------------------------------------- boot */

  function bind() {
    el.form.addEventListener('submit', function (e) { e.preventDefault(); });

    var timer = null;
    el.q.addEventListener('input', function () {
      clearTimeout(timer);
      timer = setTimeout(function () { state.q = el.q.value.trim(); update(); }, 120);
    });

    el.facets.addEventListener('click', function (e) {
      var b = e.target.closest('.pgx-fbtn');
      if (!b) return;
      var key = b.dataset.facet;
      state[key] = state[key] === b.dataset.value ? '' : b.dataset.value;
      update();
    });

    el.sort.addEventListener('change', function () { state.sort = el.sort.value; update(); });

    root.querySelector('.pgx-view').addEventListener('click', function (e) {
      var b = e.target.closest('button[data-view]');
      if (!b) return;
      state.view = b.dataset.view;
      update();
    });

    el.more.addEventListener('click', function () { shown += PAGE; paint(); });

    el.sql.addEventListener('click', function () {
      var sql = el.sql.dataset.sql || '';
      if (!navigator.clipboard) return;
      navigator.clipboard.writeText(sql).then(function () {
        el.sql.classList.add('is-copied');
        setTimeout(function () { el.sql.classList.remove('is-copied'); }, 1200);
      });
    });

    // `/` focuses the console, the way it does on the dynamic catalogue —
    // but never while the reader is already typing somewhere else.
    document.addEventListener('keydown', function (e) {
      if (e.key !== '/' || e.metaKey || e.ctrlKey || e.altKey) return;
      var a = document.activeElement;
      if (a && (a.tagName === 'INPUT' || a.tagName === 'TEXTAREA' || a.isContentEditable)) return;
      e.preventDefault();
      el.q.focus();
      el.q.select();
    });

    window.addEventListener('popstate', function () { readURL(); update(false); });
  }

  fetch(root.getAttribute('data-pgx-catalog'), { credentials: 'same-origin' })
    .then(function (r) { return r.json(); })
    .then(function (payload) {
      DB = payload;
      root.classList.add('is-live');
      buildFacets();
      readURL();
      bind();
      update(true);
    })
    .catch(function () {
      // The server-rendered category grid stays; the console simply does not
      // become interactive. Nothing to clean up.
      root.classList.add('is-offline');
    });
})();
