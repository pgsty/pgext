/* ================================================================
   Lightweight code highlighting and GFM Markdown renderer
   (pgext.doc usage manuals).
   ================================================================ */
import { esc } from './util.js';
import { getLang } from './i18n.js';

function compileHighlight(flags, rules) {
  return {
    re: new RegExp(rules.map(r => '(' + r[0] + ')').join('|'), flags),
    classes: rules.map(r => r[1])
  };
}

const HL_RULES = {
  sql: compileHighlight('gi', [
    ['--[^\\n]*', 'tok-cmt'],
    ['/\\*[\\s\\S]*?\\*/', 'tok-cmt'],
    ["'(?:[^'\\n]|'')*'", 'tok-str'],
    ['\\b(?:select|insert|update|delete|create|drop|alter|extension|table|index|view|materialized|function|procedure|trigger|schema|database|from|where|join|left|right|inner|outer|full|cross|lateral|on|group|order|by|having|limit|offset|as|and|or|not|null|true|false|is|in|exists|between|like|ilike|cascade|with|values|into|set|returning|union|all|distinct|case|when|then|else|end|grant|revoke|comment|show|explain|analyze|vacuum|begin|commit|rollback|if|replace|to|using|primary|key|foreign|references|unique|default|constraint|add|column|type|language|returns|immutable|stable|volatile|security|definer|partition|of|for|each|row|execute|declare|loop|return|raise|notice|copy|owner)\\b', 'tok-kw'],
    ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
  ]),
  bash: compileHighlight('gm', [
    ['#[^\\n]*', 'tok-cmt'],
    ["'[^'\\n]*'", 'tok-str'],
    ['"(?:[^"\\\\\\n]|\\\\.)*"', 'tok-str'],
    ['\\$\\{[^}\\n]*\\}|\\$[A-Za-z_][A-Za-z0-9_]*', 'tok-var'],
    ['\\b(?:sudo|dnf|yum|apt|apt-get|zypper|pig|pgext|psql|pg_ctl|pg_dump|pg_restore|initdb|systemctl|service|curl|wget|git|make|cmake|cargo|python3?|pip3?|echo|export|source|cd|cp|mv|rm|mkdir|tar|install|tee|cat|grep|sed|awk|docker)\\b', 'tok-kw'],
    ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
  ]),
  conf: compileHighlight('gm', [
    ['#[^\\n]*', 'tok-cmt'],
    ["'[^'\\n]*'", 'tok-str'],
    ['"[^"\\n]*"', 'tok-str'],
    ['^\\s*[A-Za-z_][\\w.-]*(?=\\s*[:=])', 'tok-var'],
    ['\\b\\d+(?:\\.\\d+)?\\b', 'tok-num']
  ])
};

const HL_LANGS = {
  sql: 'sql', psql: 'sql', postgresql: 'sql', pgsql: 'sql', plpgsql: 'sql',
  bash: 'bash', sh: 'bash', shell: 'bash', zsh: 'bash', console: 'bash', 'shell-session': 'bash', terminal: 'bash',
  yaml: 'conf', yml: 'conf', ini: 'conf', conf: 'conf', toml: 'conf', properties: 'conf', 'postgresql.conf': 'conf', env: 'conf'
};

export function highlightCode(source, language) {
  const src = String(source == null ? '' : source);
  const spec = HL_RULES[HL_LANGS[String(language || '').toLowerCase()]];
  if (!spec) return esc(src);
  const { re, classes } = spec;
  re.lastIndex = 0;
  let out = '', pos = 0, m;
  while ((m = re.exec(src))) {
    if (m.index > pos) out += esc(src.slice(pos, m.index));
    let cls = '';
    for (let g = 1; g < m.length; g++) {
      if (m[g] !== undefined) { cls = classes[g - 1]; break; }
    }
    out += cls ? '<span class="' + cls + '">' + esc(m[0]) + '</span>' : esc(m[0]);
    pos = m.index + m[0].length;
    if (!m[0].length) re.lastIndex++;
  }
  return out + esc(src.slice(pos));
}

export const mdCodeHTML = (language, value) =>
  '<div class="md-code compact"><div class="md-codebar"><span>' + esc(language) + '</span><button data-copy="' + esc(value) + '">copy</button></div>'
  + '<pre><code>' + highlightCode(value, language) + '</code></pre></div>';

export function mdSafeURL(raw) {
  const url = String(raw || '').trim().replace(/^<|>$/g, '');
  return /^(https?:\/\/|mailto:|\/|#)/i.test(url) ? url : '';
}

export function mdInline(source) {
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
  return html.replace(/\u0001(\d+)\u0002/g, (_, n) => tokens[Number(n)] || '');
}

function mdPlain(source) {
  return String(source || '')
    .replace(/!\[([^\]]*)\]\([^)]*\)/g, '$1')
    .replace(/\[([^\]]+)\]\([^)]*\)/g, '$1')
    .replace(/[`*_~]/g, '').trim();
}

function mdSlug(source, used) {
  const base = mdPlain(source).toLowerCase().normalize('NFKD')
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

const ALERT_META = {
  NOTE: ['Note', '提示', 'ⓘ'],
  TIP: ['Tip', '技巧', '☛'],
  IMPORTANT: ['Important', '重要', '❢'],
  WARNING: ['Warning', '警告', '⚠'],
  CAUTION: ['Caution', '当心', '⛒']
};

export function renderMD(source, options) {
  const src = String(source || '').replace(/\r\n?/g, '\n');
  const lines = src.split('\n');
  const out = [], used = new Map();
  const usageMode = Boolean(options && options.usage);
  const isBlock = (line, next) => /^\s*```/.test(line) || /^#{1,6}\s+/.test(line)
    || /^\s*(?:[-+*]|\d+[.)])\s+/.test(line) || /^>\s?/.test(line)
    || /^(?:\s*[-*_]){3,}\s*$/.test(line)
    || (line.includes('|') && next != null && mdTableDivider(next));
  const heading = (level, title) => {
    if (usageMode) level = Math.min(Math.max(level, 3), 6);
    out.push('<h' + level + ' id="' + mdSlug(title, used) + '">' + mdInline(title) + '</h' + level + '>');
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
      const alert = quote[0] && quote[0].match(/^\s*\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]\s*(.*)$/i);
      const inner = parts => renderMD(parts.join('\n')) || ('<p>' + mdInline(parts.join(' ')) + '</p>');
      if (alert) {
        const kind = alert[1].toUpperCase();
        const [en, zh, icon] = ALERT_META[kind];
        const rest = alert[2] ? [alert[2], ...quote.slice(1)] : quote.slice(1);
        out.push('<div class="md-alert md-alert-' + kind.toLowerCase() + '"><p class="md-alert-title"><span>' + icon + '</span>'
          + (getLang() === 'zh' ? zh : en) + '</p>' + inner(rest) + '</div>');
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
  return out.join('\n');
}
