## Usage

Sources:

- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [extensions/pg_sheet/pg_sheet.control](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/pg_sheet.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/README.md)
- [extensions/pg_sheet/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/Cargo.toml)
- [extensions/pg_sheet/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/LICENSE)
- [extensions/pg_sheet/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/ac623fe0885b79a2517eeefb9a1f4c92bddcc114/extensions/pg_sheet/README.md)

`pg_sheet` 0.3.1 adds editable overlays to existing entity tables through `pgsheet`, retaining source data and producing merged views with overlay values and formulas.

### Core Workflow

```sql
CREATE EXTENSION pg_sheet;
CREATE TABLE sheet_entities (id uuid PRIMARY KEY, title text);
INSERT INTO sheet_entities VALUES ('11111111-1111-1111-1111-111111111111', 'Example');
SELECT pgsheet.create_sheet('review', 'public', 'sheet_entities');
SELECT pgsheet.add_column('review', 'notes', 'text');
SELECT pgsheet.set_value('review', '11111111-1111-1111-1111-111111111111', 'notes', 'Checked');
SELECT pgsheet.get_data('review', 100, 0);
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. `create_sheet`, `add_column`, `set_value` and `get_data` form the core workflow. Column formulas are translated to SQL; cell formulas are stored for client-side HyperFormula evaluation. Snapshots/diffs, audit history and cell locks support collaboration. Restrict source filters, formulas and type definitions to trusted authors; these produce SQL. The example uses UUID entity identifiers. Overlay restoration changes overlay state, not the underlying source table. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

### Current Release and Upgrade

Version 0.3.1 fixes `create_sheet` when the optional `app_user` role is absent. It grants that role access only if it exists; application permissions remain an explicit deployment decision. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.4.1; upstream bottles do not imply Pigsty package availability.
