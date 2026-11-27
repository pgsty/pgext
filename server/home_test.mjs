import assert from 'node:assert/strict';
import { test } from 'node:test';
import { setLangValue } from './web/js/i18n.js';
import { decodeBoot, S, resetState, readState } from './web/js/state.js';
import { runFilter } from './web/js/filter.js';
import { hovercardHTML, pkgHovercardHTML } from './web/js/components.js';
import { homeHTML, aboutHTML } from './web/js/views.js';

function fixture() {
  const row = (name, pkg, packaged, cat, flags = 0, overrides = {}) => {
    const r = Array(34).fill(null);
    Object.assign(r, { 0: name, 1: cat, 2: packaged, 4: 'MIT', 5: 'C', 13: [18, 17], 14: flags, 17: pkg, 18: name }, overrides);
    return r;
  };
  return {
    rows: [
      row('spatial', 'spatial', 1, 'GIS', 0, { 3: 'PGDG', 7: 0, 10: 'standard' }),
      row('spatial_helper', 'spatial', 1, 'GIS', 0, { 3: 'PGDG', 7: null, 10: 'standard' }),
      row('source_only', 'source_only', 0, 'GIS', 0, { 11: 'AWS', 10: 'puresql' }),
      row('contrib', 'contrib', 1, 'ADMIN', 2, { 3: 'CONTRIB', 7: 50, 10: 'standard' })
    ],
    cats: [{ name: 'GIS', count: 3 }, { name: 'ADMIN', count: 1 }],
    pg: [18, 17], os: ['el9.x86_64', 'el9.aarch64'],
    counts: { total: 999, packaged: 600, packages: 2, build_slots: 5 }
  };
}

function metrics(html) {
  const stats = html.match(/<ul class="hero-stats">([\s\S]*?)<\/ul>/)[1];
  return [...stats.matchAll(/<span class="num">([^<]+)<\/span>/g)].map(m => m[1]);
}

for (const lang of ['en', 'zh']) {
  test(`home, scope counts and about share live totals (${lang})`, () => {
    resetState();
    setLangValue(lang);
    decodeBoot(fixture());
    const html = homeHTML();
    assert.deepEqual(metrics(html), ['4', '3', '2', '2', '2', '5']);
    assert.match(html, /class="hero-sub oneline">[^\n]*<b>3<\/b>/);
    assert.match(html, /data-fval="packaged"[^>]*><span>[^<]+<\/span><small>3<\/small>/);
    assert.match(html, lang === 'en' ? /Linux targets/ : /Linux 目标/);
    assert.match(html, lang === 'en' ? /not a count of successful builds/ : /并非成功构建数/);
    const about = aboutHTML();
    assert.doesNotMatch(about, /<b>600<\/b>/);
    assert.equal((about.match(/<b>3<\/b>/g) || []).length, 2);

    S.scope = 'packaged';
    const filtered = homeHTML();
    assert.deepEqual(metrics(filtered), metrics(html));
    assert.match(filtered, /class="result-summary"><strong>3<\/strong>/);
    assert.match(filtered, /data-fval="GIS"[^>]*>[\s\S]*?<small>2<\/small>/);
    S.entity = 'pkg';
    assert.match(homeHTML(), /class="result-summary"><strong>2<\/strong>/);
  });

  test(`fresh bootstrap updates all six metrics (${lang})`, () => {
    resetState();
    setLangValue(lang);
    const payload = fixture();
    decodeBoot(payload);
    payload.rows[2][2] = 1;
    payload.rows.push([...payload.rows[2]]);
    payload.rows[4][0] = 'new_extension';
    payload.pg.push(16);
    payload.os.push('u24.aarch64');
    payload.counts = { packages: 3, build_slots: 1234 };
    decodeBoot(payload);
    assert.deepEqual(metrics(homeHTML()), ['5', '5', '3', '3', '3', '1,234']);
  });
}

for (const slots of [0, undefined]) {
  test(`matrix slots do not assume full coverage when count is ${slots}`, () => {
    resetState();
    setLangValue('en');
    const payload = fixture();
    payload.counts.build_slots = slots;
    decodeBoot(payload);
    assert.deepEqual(metrics(homeHTML()), ['4', '3', '2', '2', '2', '0']);
  });
}

test('filter and sort parity with Go backend', () => {
  setLangValue('en');
  decodeBoot(fixture());

  // Stars sort places 0-star entries ahead of unknown-star (null) entries.
  resetState();
  assert.deepEqual(runFilter().list.map(e => e.name), ['contrib', 'spatial', 'source_only', 'spatial_helper']);

  // Legacy scope=cloud normalizes to vendor; repo/kind matching is case-insensitive.
  readState(new URLSearchParams('scope=cloud&kind=PURESQL'));
  assert.equal(S.scope, 'vendor');
  assert.deepEqual(runFilter().list.map(e => e.name), ['source_only']);

  readState(new URLSearchParams('repo=pgdg'));
  assert.deepEqual(runFilter().list.map(e => e.name), ['spatial', 'spatial_helper']);
});

test('package hovercard shares body renderer without string slicing', () => {
  setLangValue('en');
  decodeBoot(fixture());
  const e = { name: 'spatial', pkg: 'spatial', cat: 'GIS', avail: true, repo: 'PGDG', license: 'MIT', lang: 'C', ver: '1.0', pg: [18], flags: 16, en: 'desc', active: '2026-01-01' };
  const extCard = hovercardHTML(e, null);
  const pkgCard = pkgHovercardHTML('spatial', [e], e, null);
  assert.match(extCard, /class="hc-title" href="\/ext\/spatial"/);
  assert.match(pkgCard, /class="hc-title" href="\/pkg\/spatial"/);
  assert.match(pkgCard, /class="hc-exts"/);
});
