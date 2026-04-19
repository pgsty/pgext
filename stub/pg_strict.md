## Usage

Source: [README](https://github.com/spa5k/pg_strict/blob/master/README.md), [Release v1.0.5](https://github.com/spa5k/pg_strict/releases/tag/v1.0.5), [API source](https://github.com/spa5k/pg_strict/blob/master/src/api.rs)

`pg_strict` blocks or warns on `UPDATE` and `DELETE` statements that omit a `WHERE` clause. It installs a `post_parse_analyze_hook`, so it must be loaded from `shared_preload_libraries`.

### Required setup

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_strict'

CREATE EXTENSION pg_strict;
```

### GUCs

- `pg_strict.require_where_on_update`
- `pg_strict.require_where_on_delete`

Each setting supports `off`, `warn`, and `on`.

```sql
SET pg_strict.require_where_on_update = 'on';
SET pg_strict.require_where_on_delete = 'warn';
```

### Helper functions

```sql
SELECT pg_strict_version();
SELECT * FROM pg_strict_config();

SELECT pg_strict_check_where_clause('DELETE FROM t', 'DELETE');
SELECT pg_strict_validate_update('UPDATE t SET x = 1 WHERE id = 42');
SELECT pg_strict_validate_delete('DELETE FROM t WHERE id = 42');

SELECT pg_strict_enable_update();
SELECT pg_strict_warn_delete();
SELECT pg_strict_disable_delete();
```

- `pg_strict_set_update_mode(mode)` and `pg_strict_set_delete_mode(mode)` provide generic mode setters.
- `SET LOCAL` works for one-off bulk operations inside a transaction.

### Caveats

- Enforcement is presence-based, not intent-based: any non-null `WHERE` clause satisfies the rule.
- Only `UPDATE` and `DELETE` are checked.
- Current upstream release is `1.0.5`; the Pigsty note about `pgrx` 0.17.0 is packaging/build metadata, not a documented user-facing feature change.
