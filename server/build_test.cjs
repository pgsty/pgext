const assert = require('node:assert/strict');
const vm = require('node:vm');
const { test } = require('node:test');
const { createAppContext } = require('./test_app.cjs');

for (const lang of ['en', 'zh']) {
  test(`PGEXT recipe guide and separate source inputs (${lang})`, () => {
    const ctx = createAppContext();
    vm.runInContext(`LANG='${lang}'`, ctx);
    const html = ctx.buildHTML({ name: 'h3_postgis', pkg: 'pg_h3', rpm_build: true, deb_build: true,
      rpm_pg: [99], deb_pg: [98], tarball: 'h3-pg.tar.gz h3.tar.gz h3.tar.gz' });
    assert.match(html, /pgext pkg pg_h3 --plan\npgext pkg pg_h3/);
    assert.match(html, /pgext env repo\npgext env tool\npgext env spec/);
    assert.doesNotMatch(html, /pig build|PG 99|PG 98/);
    const domain = lang === 'zh' ? 'cc' : 'io';
    assert.match(html, new RegExp('href="https://repo.pigsty.' + domain + '/ext/src/h3-pg.tar.gz"'));
    assert.equal((html.match(/href="[^\"]*\/h3.tar.gz"/g) || []).length, 1);
    assert.doesNotMatch(html, /%20/);
  });
}

test('recipe flags control build guidance independently of published packages', () => {
  const ctx = createAppContext();
  assert.match(ctx.buildHTML({ name: 'source', pkg: 'source', packaged: false, rpm_build: true }), /pgext pkg source/);
  assert.doesNotMatch(ctx.buildHTML({ name: 'provider', pkg: 'provider', packaged: false }), /pgext pkg/);
  assert.doesNotMatch(ctx.buildHTML({ name: 'pg_stat_statements', contrib: true, rpm_build: true }), /pgext pkg/);
  assert.match(ctx.buildHTML({ name: 'rust', pkg: 'rust', deb_build: true, pgrx_ver: '0.19.3' }), /pgext env pgrx --pgrx 0.19.3/);
  assert.doesNotMatch(ctx.buildHTML({ name: 'rust', pkg: 'rust', deb_build: true, pgrx_ver: '=0.16.1' }), /--pgrx =/);
});

test('about setup commands have real line breaks and describe both tools', () => {
  const html = createAppContext().aboutHTML();
  assert.match(html, /pgext pkg pgvector --pg 18/);
  assert.match(html, /pig repo set/);
  assert.doesNotMatch(html, /\\npig/);
});

test('file renderer preserves server package-version order', () => {
  const ctx = createAppContext();
  const files = ['1.10-1', '1.9-1'].map(ver => ({ pg: 18, os: 'el9.x86_64', name: 'example_18', ver, file: ver, url: '/file' }));
  const html = ctx.filesHTML({ files });
  assert.match(html, /<tr><td[^>]*>[\s\S]*?<td>1\.10-1<\/td>/);
  assert.match(html, /<tr class="older">[\s\S]*?<td>1\.9-1<\/td>/);
});
