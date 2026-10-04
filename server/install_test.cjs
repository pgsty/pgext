const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const vm = require('node:vm');
const { test } = require('node:test');

// Exercise the actual install renderer without starting the browser application.
const source = fs.readFileSync(path.join(__dirname, 'web/app.js'), 'utf8');
const renderer = source.slice(source.indexOf('function installHTML('), source.indexOf('/* full availability matrix:'));
const context = vm.createContext({
  sqlIdent: value => value,
  mdCodeHTML: (language, value) => '<code>' + value + '</code>',
  mdSafeURL: value => value,
  packageTabsHTML: () => '',
  bi: en => en,
  t: value => value,
  esc: value => value,
});
vm.runInContext(renderer, context);

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
    const html = context.installHTML({ name: 'pgexporter_ext' }, full);
    const commands = html.match(/pig repo add \w+ -u/g) || [];
    assert.deepEqual(commands, expected ? [expected] : []);
    assert.deepEqual(full, before, 'renderer must preserve the curated supplier');
  });
}
