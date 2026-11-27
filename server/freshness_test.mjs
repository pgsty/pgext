import assert from 'node:assert/strict';
import { test } from 'node:test';
import { DetailCache, decodeBoot, FULLC, DOCC, CHLOGC, matrixState } from './web/js/state.js';
import { ensureGlobalMatrix, hydrateGlobalMatrix, MATRIX_CACHE_KEY } from './web/js/matrix.js';

test('detail caches expire independently of bootstrap metadata', () => {
  const realNow = Date.now;
  let now = 1000;
  Date.now = () => now;
  try {
    const doc = new DetailCache();
    const chlog = new DetailCache();
    doc.set('demo|en', { content: 'old' });
    chlog.set('demo', { entries: [] });
    assert.equal(doc.get('demo|en').content, 'old');
    now += 5 * 60 * 1000;
    assert.equal(doc.get('demo|en'), undefined);
    assert.equal(chlog.get('demo'), undefined);
  } finally {
    Date.now = realNow;
  }
});

test('a changed bootstrap clears detail caches and matrix validation', () => {
  const boot = { version: 'before', rows: [], cats: [], pg: [], os: [], counts: {} };
  decodeBoot(boot);
  FULLC.set('demo', {});
  DOCC.set('demo|en', {});
  CHLOGC.set('demo', {});
  matrixState.checked[''] = true;
  matrixState.checked.pgdg = true;
  decodeBoot({ ...boot, version: 'after' });
  assert.equal(FULLC.get('demo'), undefined);
  assert.equal(DOCC.get('demo|en'), undefined);
  assert.equal(CHLOGC.get('demo'), undefined);
  assert.equal(matrixState.checked[''], undefined);
  assert.equal(matrixState.checked.pgdg, undefined);
});

test('PG/OS matrix entry revalidates a localStorage payload', async () => {
  const old = {
    format: 'matrix-row.v2', generated: 'same-time', pg: [18], os: ['el9.x86_64'],
    stats: { cells: 1 }, rows: [{ p: 'demo', e: 'demo', c: 'P', v: ['1.9'] }]
  };
  const fresh = { ...old, rows: [{ ...old.rows[0], v: ['1.10'] }] };
  let calls = 0;
  const origStorage = globalThis.localStorage;
  const origFetch = globalThis.fetch;
  delete matrixState.data[''];
  delete matrixState.checked[''];
  globalThis.localStorage = { getItem: key => key === MATRIX_CACHE_KEY ? JSON.stringify(old) : null, setItem() {} };
  globalThis.fetch = async () => { calls++; return { ok: true, json: async () => fresh }; };
  try {
    assert.equal((await ensureGlobalMatrix()).rows[0].v[0], '1.10');
    await ensureGlobalMatrix();
    assert.equal(calls, 1);
  } finally {
    globalThis.localStorage = origStorage;
    globalThis.fetch = origFetch;
  }
});

test('global matrix accepts changed versions with unchanged timestamp and size', async () => {
  const old = { format: 'matrix-row.v2', generated: 'same-time', stats: { cells: 1 }, rows: [{ p: 'demo', v: ['1.9'] }] };
  const fresh = { ...old, rows: [{ p: 'demo', v: ['1.10'] }] };
  const origFetch = globalThis.fetch;
  const origDoc = globalThis.document;
  matrixState.data[''] = old;
  delete matrixState.checked[''];
  globalThis.fetch = async () => ({ ok: true, json: async () => fresh });
  globalThis.document = { getElementById: () => null };
  try {
    await hydrateGlobalMatrix(() => true);
    assert.equal(matrixState.data[''].rows[0].v[0], '1.10');
  } finally {
    globalThis.fetch = origFetch;
    globalThis.document = origDoc;
  }
});
