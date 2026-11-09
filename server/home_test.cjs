const assert = require('node:assert/strict');
const vm = require('node:vm');
const { test } = require('node:test');
const { createAppContext } = require('./test_app.cjs');

function fixture() {
  const row = (name, pkg, packaged, cat, flags = 0) => {
    const r = Array(34).fill(null);
    Object.assign(r, { 0: name, 1: cat, 2: packaged, 4: 'MIT', 5: 'C',
      13: [18, 17], 14: flags, 17: pkg, 18: name });
    return r;
  };
  return {
    rows: [row('spatial', 'spatial', 1, 'GIS'), row('spatial_helper', 'spatial', 1, 'GIS'),
      row('source_only', 'source_only', 0, 'GIS'), row('contrib', 'contrib', 1, 'ADMIN', 2)],
    cats: [{ name: 'GIS', count: 3 }, { name: 'ADMIN', count: 1 }],
    pg: [18, 17], os: ['el9.x86_64', 'el9.aarch64'],
    // Row counts are deliberately stale. Matrix totals cannot be inferred from rows.
    counts: { total: 999, packaged: 600, packages: 2, build_slots: 5 },
  };
}

function renderContext(lang, payload) {
  const context = createAppContext();
  vm.runInContext(`LANG = '${lang}'`, context);
  context.decodeBoot(payload);
  return context;
}

function metrics(html) {
  const stats = html.match(/<ul class="hero-stats">([\s\S]*?)<\/ul>/)[1];
  return [...stats.matchAll(/<span class="num">([^<]+)<\/span>/g)].map(m => m[1]);
}

for (const lang of ['en', 'zh']) {
  test(`home, scope counts and about share live totals (${lang})`, () => {
    const context = renderContext(lang, fixture());
    const html = context.homeHTML();
    assert.deepEqual(metrics(html), ['4', '3', '2', '2', '2', '5']);
    assert.match(html, /class="hero-sub oneline">[^\n]*<b>3<\/b>/);
    assert.match(html, /data-fval="packaged"[^>]*><span>[^<]+<\/span><small>3<\/small>/);
    assert.match(html, lang === 'en' ? /Linux targets/ : /Linux 目标/);
    assert.match(html, lang === 'en' ? /not a count of successful builds/ : /并非成功构建数/);
    const about = context.aboutHTML();
    assert.doesNotMatch(about, /<b>600<\/b>/);
    assert.equal((about.match(/<b>3<\/b>/g) || []).length, 2);

    // Narrowing results must not replace the global hero totals.
    vm.runInContext("S.scope = 'packaged'", context);
    const filtered = context.homeHTML();
    assert.deepEqual(metrics(filtered), metrics(html));
    assert.match(filtered, /class="result-summary"><strong>3<\/strong>/);
    assert.match(filtered, /data-fval="GIS"[^>]*>[\s\S]*?<small>2<\/small>/);
    vm.runInContext("S.entity = 'pkg'", context);
    assert.match(context.homeHTML(), /class="result-summary"><strong>2<\/strong>/);
  });

  test(`fresh bootstrap updates all six metrics (${lang})`, () => {
    const payload = fixture();
    const context = renderContext(lang, payload);
    payload.rows[2][2] = 1;
    payload.rows.push([...payload.rows[2]]);
    payload.rows[4][0] = 'new_extension';
    payload.pg.push(16);
    payload.os.push('u24.aarch64');
    payload.counts = { packages: 3, build_slots: 1234 };
    context.decodeBoot(payload);
    assert.deepEqual(metrics(context.homeHTML()), ['5', '5', '3', '3', '3', '1,234']);
  });
}

for (const slots of [0, undefined]) {
  test(`matrix slots do not assume full coverage when count is ${slots}`, () => {
    const payload = fixture();
    payload.counts.build_slots = slots;
    const context = renderContext('en', payload);
    assert.deepEqual(metrics(context.homeHTML()), ['4', '3', '2', '2', '2', '0']);
  });
}
