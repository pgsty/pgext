import assert from 'node:assert/strict';
import { test } from 'node:test';
import { setLangValue } from './web/js/i18n.js';
import { installHTML } from './web/js/detail.js';

setLangValue('en');

for (const [name, fields, expected] of [
  ['PGDG only', { repo: 'PGDG', rpm_repo: 'PGDG', deb_repo: 'PGDG' }, 'pig repo add pgdg -u'],
  ['pgexporter_ext', { repo: 'PGDG', rpm_repo: 'PGDG', deb_repo: 'PIGSTY' }, 'pig repo add pgsql -u'],
  ['Pigsty RPM', { repo: 'PGDG', rpm_repo: 'PIGSTY', deb_repo: 'PGDG' }, 'pig repo add pgsql -u'],
  ['legacy PGDG', { repo: 'PGDG' }, 'pig repo add pgdg -u'],
  ['Pigsty', { repo: 'PIGSTY' }, 'pig repo add pgsql -u'],
  ['contrib', { repo: 'CONTRIB', contrib: true }, null],
  ['source only', { packaged: false }, null],
]) {
  test(name, () => {
    const full = { packaged: true, ...fields };
    const before = structuredClone(full);
    const html = installHTML({ name: 'pgexporter_ext' }, full);
    const commands = html.match(/pig repo add \w+ -u/g) || [];
    assert.deepEqual(commands, expected ? [expected] : []);
    assert.deepEqual(full, before, 'renderer must preserve the curated supplier');
  });
}

for (const kind of ['headless', 'standard', '']) {
  test(`explicit ${kind || 'unknown'} loading does not invent a shared preload`, () => {
    const html = installHTML({ name: 'plugin' }, { kind, need_load: true, libs: ['plugin'], need_ddl: false });
    assert.doesNotMatch(html, /shared_preload_libraries\s*=/);
    assert.match(html, /href="#ext-usage"/);
  });
}

test('known shared preload uses dependency libraries', () => {
  const html = installHTML({ name: 'demo' }, { kind: 'preload', need_load: true, preload_libs: ['dependency', 'demo'] });
  assert.match(html, /data-copy="shared_preload_libraries = &#39;dependency, demo&#39;"/);
});
