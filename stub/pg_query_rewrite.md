## Usage

Source: [README](https://github.com/pierreforstmann/pg_query_rewrite/blob/master/README.md)

`pg_query_rewrite` rewrites an exact source SQL statement into a predefined target statement using rules stored in shared memory.

### Required setup

```sql
-- postgresql.conf
shared_preload_libraries = 'pg_query_rewrite'
pg_query_rewrite.max_rules = 10

CREATE EXTENSION pg_query_rewrite;
```

### Rule management

```sql
SELECT pgqr_add_rule('select 10;', 'select 11;');
SELECT pgqr_rules();
SELECT pgqr_remove_rule('select 10;');
SELECT pgqr_truncate();
```

- `pgqr_add_rule(source, target)`: add a rewrite rule.
- `pgqr_remove_rule(source)`: remove one rule.
- `pgqr_truncate()`: remove all rules.
- `pgqr_rules()`: inspect current shared-memory rules and rewrite counts.

### Caveats

- Matching is exact: case, whitespace, and semicolons must match character-for-character.
- Parameterized SQL is not supported.
- Maximum statement length is hard-coded at 32 KB.
- Rules are not persisted; they vanish on restart unless you reapply them.
