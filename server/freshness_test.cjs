const assert = require('node:assert/strict');
const vm = require('node:vm');
const { test } = require('node:test');
const { createAppContext } = require('./test_app.cjs');

test('detail caches expire independently of bootstrap metadata', () => {
  let now = 1000;
  const ctx = createAppContext({ Date: { now: () => now } });
  vm.runInContext("DOCC.set('demo|en', {content:'old'}); CHLOGC.set('demo', {entries:[]})", ctx);
  assert.equal(vm.runInContext("DOCC.get('demo|en').content", ctx), 'old');
  now += 5 * 60 * 1000;
  assert.equal(vm.runInContext("DOCC.get('demo|en')", ctx), undefined);
  assert.equal(vm.runInContext("CHLOGC.get('demo')", ctx), undefined);
});

test('a changed bootstrap clears detail caches and matrix validation', () => {
  const ctx = createAppContext();
  const boot = { version: 'before', rows: [], cats: [], pg: [], os: [], counts: {} };
  ctx.decodeBoot(boot);
  vm.runInContext("FULLC.set('demo', {}); MATRIX_CHECKED=true; RMATRIX_CHECKED.pgdg=true", ctx);
  ctx.decodeBoot({ ...boot, version: 'after' });
  assert.equal(vm.runInContext("FULLC.get('demo')", ctx), undefined);
  assert.equal(vm.runInContext('MATRIX_CHECKED', ctx), false);
  assert.equal(vm.runInContext('RMATRIX_CHECKED.pgdg', ctx), undefined);
});

test('PG/OS matrix entry revalidates a localStorage payload', async () => {
  const old = { format: 'matrix-row.v2', generated: 'same-time', pg: [18], os: ['el9.x86_64'],
    stats: { cells: 1 }, rows: [{ p: 'demo', e: 'demo', c: 'P', v: ['1.9'] }] };
  const fresh = { ...old, rows: [{ ...old.rows[0], v: ['1.10'] }] };
  let calls = 0;
  const ctx = createAppContext({
    localStorage: { getItem: key => key === 'pgext.matrix.v2' ? JSON.stringify(old) : null, setItem() {} },
    fetch: async () => { calls++; return { ok: true, json: async () => fresh }; },
  });
  assert.equal((await ctx.ensureGlobalMatrix()).rows[0].v[0], '1.10');
  await ctx.ensureGlobalMatrix();
  assert.equal(calls, 1);
});

test('global matrix accepts changed versions with unchanged timestamp and size', async () => {
  const old = { format: 'matrix-row.v2', generated: 'same-time', stats: { cells: 1 }, rows: [{ p: 'demo', v: ['1.9'] }] };
  const fresh = { ...old, rows: [{ p: 'demo', v: ['1.10'] }] };
  const ctx = createAppContext({ fetch: async () => ({ ok: true, json: async () => fresh }), document: { getElementById: () => null } });
  ctx.oldMatrix = old;
  vm.runInContext('GMATRIX=oldMatrix', ctx);
  await ctx.hydrateGlobalMatrix();
  assert.equal(vm.runInContext('GMATRIX.rows[0].v[0]', ctx), '1.10');
});
