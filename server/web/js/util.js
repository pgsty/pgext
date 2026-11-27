/* ================================================================
   Pure formatting, escaping, URL, hue and clipboard utilities.
   ================================================================ */

const ESC_MAP = { '&': '&amp;', '<': '&lt;', '>': '&gt;', '"': '&quot;', "'": '&#39;' };
export const esc = s => String(s == null ? '' : s).replace(/[&<>"']/g, m => ESC_MAP[m]);
export const fmtNum = n => n == null ? '—' : n >= 1000 ? (n / 1000).toFixed(n >= 10000 ? 0 : 1) + 'k' : String(n);
export const fmtInt = n => n == null ? '—' : String(n).replace(/\B(?=(\d{3})+(?!\d))/g, ',');
export const fmtSize = b => b == null || b === 0 ? '—' : b >= 1048576 ? (b / 1048576).toFixed(1) + ' MiB' : (b / 1024).toFixed(0) + ' KiB';
export const starOf = e => (e && e.stars != null ? e.stars : -1);

export const pgRange = arr => {
  if (!arr || !arr.length) return '—';
  const s = [...arr].sort((a, b) => a - b);
  return s.length > 1 ? s[0] + '–' + s[s.length - 1] : String(s[0]);
};

export const extHref = n => '/ext/' + encodeURIComponent(n);
export const pkgHref = n => '/pkg/' + encodeURIComponent(n);
export const catHref = n => '/cate/' + encodeURIComponent(n);
export const dimHref = n => '/list/' + encodeURIComponent(n);
export const catVar = c => 'style="--seg:var(--c-' + esc(c) + ')"';

/* Dimension hues: cool = permissive/common, warm = copyleft/modern,
   red = restrictive/rare, violet = commercial/exotic. */
export function licenseHue(v) {
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

export const langHue = v => LANG_HUES[v] || (v ? 'red' : '');
export const repoHue = v => REPO_HUES[v] || '';
export const hueVar = h => h ? 'style="--seg:var(--h-' + h + ')"' : '';

export function dimSeg(dim, v) {
  if (dim === 'category') return 'var(--c-' + v + ')';
  const h = dim === 'license' ? licenseHue(v) : dim === 'lang' ? langHue(v) : dim === 'repo' ? repoHue(v) : '';
  return h ? 'var(--h-' + h + ')' : '';
}

export const debounce = (fn, ms) => { let h; return (...a) => { clearTimeout(h); h = setTimeout(() => fn(...a), ms); }; };

export const osLabel = os => {
  const [base, arch] = String(os || '').split('.');
  const family = base.startsWith('el') ? 'EL ' + base.slice(2)
    : base.startsWith('d') ? 'Debian ' + base.slice(1)
    : base.startsWith('u') ? 'Ubuntu ' + base.slice(1) : base;
  return family + (arch ? ' · ' + arch : '');
};

export const sqlIdent = name => /^[a-z_][a-z0-9_$]*$/.test(name) ? name : '"' + String(name).replaceAll('"', '""') + '"';
export const shellArg = value => "'" + String(value).replaceAll("'", "'\"'\"'") + "'";

export function copyText(txt, cb) {
  const done = ok => cb && cb(ok);
  if (typeof navigator !== 'undefined' && navigator.clipboard && navigator.clipboard.writeText) {
    navigator.clipboard.writeText(txt).then(() => done(true), () => done(false));
  } else if (typeof document !== 'undefined') {
    const ta = document.createElement('textarea');
    ta.value = txt; ta.style.position = 'fixed'; ta.style.opacity = '0';
    document.body.appendChild(ta); ta.select();
    let ok = false; try { ok = document.execCommand('copy'); } catch (e) {}
    ta.remove(); done(ok);
  } else {
    done(false);
  }
}
