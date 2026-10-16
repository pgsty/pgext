## Usage

Sources:

- [README.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/README.md)
- [docs/USER_GUIDE.md](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/docs/USER_GUIDE.md)
- [codesystem--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.2.0.sql)
- [codesystem--0.1.0--0.2.0.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem--0.1.0--0.2.0.sql)
- [codesystem.control](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/codesystem.control)
- [.github/workflows/test.yml](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/.github/workflows/test.yml)
- [tools/gss_update.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/gss_update.sql)
- [tools/install_isic_isco.sql](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/tools/install_isic_isco.sql)
- [NOTICE](https://github.com/John-Cameron-Burns/pg-codesystem/blob/482b052ab3ebd2b60816b4479d5cefa63eb7f61d/NOTICE)

`codesystem` stores named classification systems and typed `code` values. Version 0.2.0 supports list hierarchies, revisions, derived standards, crosswalks and pattern-defined identifiers. A pattern validates syntax, not whether an identifier has actually been issued.

### Core Workflow

```sql
CREATE EXTENSION codesystem;
CREATE TABLE classification_source (code text, level text, parent text, label text);
INSERT INTO classification_source VALUES ('A', 'section', '', 'Agriculture');
SELECT define_code_system(
  'toy',
  '[{"name":"section","width":1,"charset":"[A-Z]","link":"none","expected":1}]',
  'SELECT * FROM classification_source', 'Toy classification', 'Example');
SELECT name, current_revision, ncodes FROM code_systems;
SELECT code_label('toy:A'::code);
SELECT define_code_system_pattern('my_id', 'AA-NNN');
SELECT code_from('AB-123', 'my_id');
```

### Operational Boundaries

Upstream CI targets PostgreSQL 13–17. Install this native extension as an administrator. Run `ALTER EXTENSION codesystem UPDATE` only when the SQL extension version changes. Distribution 0.2.1 still uses SQL/control version 0.2.0, so an existing 0.2.0 database needs no extension upgrade.

`define_code_system` validates levels and the input query’s code, level, parent and label columns. SQL literals include a system prefix; `code_from` converts external text. `code_level`, `code_parent`, `code_ancestor`, `code_children` and `code_label` expose the hierarchy. Equal text in different revisions represents different values; `code_current` converts to the current revision when the code remains valid.

`define_code_system` also accepts `refines` and `shared_levels` for derived classifications, queried through `code_rollup`. `define_crosswalk` registers correspondence pairs; `code_map` can return multiple matches, while `code_map_unique` requires exactly one. `define_code_system_pattern` defines bounded templates or regular expressions; `code_part` and `code_parts` return named parts. `code_from_padded` handles lost leading zeros, and `code_first_with_prefix` plus `code_last_with_prefix` provide indexable prefix bounds. A system cannot switch between list and pattern forms.

Registry structural edits are protected; define a new revision when accepted values change. Registry writes need appropriate table permissions, and readers need SELECT on registry tables. Custom-format restores require the upstream `tools/restore.sh` helper to restore registry data before typed table values. Ordinary restore ordering can fail.

### Optional Classification Data Tools

Distribution 0.2.1 includes GSS code patterns and entity-type data. After separately loading `tools/gss_update.sql`, `gss_update_entity_types()` updates entity types from ONS. Loading `tools/install_isic_isco.sql` provides `codesystem_install_isic_rev4()` and `codesystem_install_isco_08()`, which download classifications from the UN and ILO. These helpers require an existing `plpython3u`, superuser installation and database-server access to the corresponding sources. They are not dependencies of the core extension.

Use `dry_run` to inspect fetched data first, and connect to the primary when writing the classification registry. ISIC/ISCO data is not redistributed in the package and remains subject to publisher terms. Upstream `NOTICE` records attribution and licensing for other bundled data.
