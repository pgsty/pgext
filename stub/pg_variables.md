
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pg_variables;
> SELECT pgv_set('vars', 'int1', 101);
> SELECT pgv_get('vars', 'int1', NULL::int);
> ```
>
> Source: [README](https://github.com/postgrespro/pg_variables)

`pg_variables` provides session-wide variables for PostgreSQL. Variables are grouped into packages, live only for the current session, and can be configured as transactional or non-transactional.

## Basic Behavior

By default, variables are not transactional and are not affected by `BEGIN`, `COMMIT`, or `ROLLBACK`. The optional `is_transactional` argument on `pgv_set()` changes that behavior.

```sql
SELECT pgv_set('vars', 'int1', 101);
SELECT pgv_get('vars', 'int1', NULL::int);
```

Transactional example:

```sql
BEGIN;
SELECT pgv_set('vars', 'trans_int', 101, true);
SAVEPOINT sp1;
SELECT pgv_set('vars', 'trans_int', 102, true);
ROLLBACK TO sp1;
COMMIT;
SELECT pgv_get('vars', 'trans_int', NULL::int);
```

## Packages

Variables are grouped into packages so multiple named variables can coexist and whole groups can be removed together. The README notes that empty packages are deleted automatically.

## Core Functions

### Scalar and Array Variables

The generic API is:

```sql
pgv_set(package text, name text, value anynonarray, is_transactional bool default false)
pgv_get(package text, name text, var_type anynonarray, strict bool default true)

pgv_set(package text, name text, value anyarray, is_transactional bool default false)
pgv_get(package text, name text, var_type anyarray, strict bool default true)
```

`pgv_get()` checks both existence and type. If the package or variable is missing, behavior depends on `strict`.

### Record Collections

The README also documents record-oriented operations such as:

- `pgv_insert()`
- `pgv_update()`
- `pgv_delete()`
- `pgv_select()`

These functions work with collections of records stored under a package and variable name.

## Deprecated Helpers

The project still ships older type-specific helpers like:

- `pgv_set_int()` / `pgv_get_int()`
- `pgv_set_text()` / `pgv_get_text()`
- `pgv_set_numeric()` / `pgv_get_numeric()`
- `pgv_set_timestamp()` / `pgv_get_timestamp()`
- `pgv_set_timestamptz()` / `pgv_get_timestamptz()`
- `pgv_set_date()` / `pgv_get_date()`
- `pgv_set_jsonb()` / `pgv_get_jsonb()`

The README labels these as deprecated in favor of the generic `pgv_set()` / `pgv_get()` API.
