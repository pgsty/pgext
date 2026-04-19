## Usage

- Sources: [README](https://github.com/postgrespro/pg_variables/blob/master/README.md), [repository tags](https://github.com/postgrespro/pg_variables/tags), [control file](https://github.com/postgrespro/pg_variables/blob/master/pg_variables.control)

`pg_variables` provides session-scoped variables grouped into named packages. Variables live only in the current session and are non-transactional by default unless created with `is_transactional := true`.

### Basic Set And Get

```sql
CREATE EXTENSION pg_variables;

SELECT pgv_set('vars', 'int1', 101);
SELECT pgv_get('vars', 'int1', NULL::int);
```

Transactional variables participate in savepoints and rollbacks:

```sql
BEGIN;
SELECT pgv_set('vars', 'trans_int', 101, true);
SAVEPOINT sp1;
SELECT pgv_set('vars', 'trans_int', 102, true);
ROLLBACK TO sp1;
COMMIT;
```

### Core APIs

The README documents generic scalar and array APIs:

- `pgv_set(package, name, value, is_transactional default false)`
- `pgv_get(package, name, NULL::type, strict default true)`

It also documents record-oriented APIs:

- `pgv_insert()`
- `pgv_update()`
- `pgv_delete()`
- `pgv_select()`

Useful administration helpers include `pgv_exists()`, `pgv_remove()`, `pgv_free()`, `pgv_list()`, and `pgv_stats()`.

### Error And Strictness Behavior

`pgv_get()` checks both existence and type. The README shows that missing packages, missing variables, or mismatched types raise errors unless `strict := false`, in which case `NULL` is returned for missing values.

### Deprecated Helpers And Version Note

Upstream still ships deprecated type-specific helpers such as `pgv_set_int()` / `pgv_get_int()` and `pgv_set_jsonb()` / `pgv_get_jsonb()`, but recommends the generic `pgv_set()` / `pgv_get()` API.

The repository tag is `v1.2.5`, while the current `pg_variables.control` file still declares `default_version = '1.2'`. That matches the packaging note that the release tag advanced without changing the SQL extension version string.
