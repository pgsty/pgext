## Usage
- GitHub Repo: [`pierreforstmann/pg_query_rewrite`](https://github.com/pierreforstmann/pg_query_rewrite)
- README: [pierreforstmann/pg_query_rewrite/blob/master/README.md](https://github.com/pierreforstmann/pg_query_rewrite/blob/master/README.md)

`pg_query_rewrite` is a PostgreSQL extension that translates an exact source SQL statement into another predefined SQL statement. It must be loaded at server level with `shared_preload_libraries`, then installed in each database.

The README shows the extension as a simple statement rewrite engine backed by shared-memory rules.

### Setup

```conf
shared_preload_libraries = 'pg_query_rewrite'
pg_query_rewrite.max_rules = 10
```

```sql
CREATE EXTENSION pg_query_rewrite;
```

The README says the extension has been successfully tested with PostgreSQL 9.5 through 18.

### Managing Rules

Use the helper functions to manage translations:

```sql
SELECT pgqr_add_rule('select 10;', 'select 11;');
SELECT pgqr_remove_rule('select 10;');
SELECT pgqr_truncate();
SELECT pgqr_rules();
```

The example in the README rewrites `select 10;` to `select 11;` and then shows the rule list after insertion.

### Behavior

- Matching is exact and sensitive to case, spaces, and semicolons.
- Parameterized SQL statements are not supported.
- The maximum SQL statement length is hard-coded at 32K.
- Rules live only in shared memory; the extension does not persist settings on its own.
- `pg_query_rewrite.max_rules` caps the number of SQL statements that can be translated and defaults to `10` when unset.

### Scope

The upstream README is sufficient here: it covers purpose, server-level loading, rule management, a concrete rewrite example, and the main limitations. No separate docs page or homepage was needed.
