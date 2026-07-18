## Usage

Sources:

- [Upstream compatibility statement](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/README.md)
- [Function implementation](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/src/jsonb_set_lax.c)
- [SQL declaration](https://github.com/adunstan/jsonb_set_lax/blob/86822571fdaca44f2bf9adbaa912bcb05d195e0a/sql/jsonb_set_lax.sql)

`jsonb_set_lax` backports PostgreSQL 13's same-named function to PostgreSQL 9.5 through 12. It changes how an SQL NULL replacement is handled: `use_json_null`, `delete_key`, `return_target`, or `raise_exception` can be selected.

```sql
CREATE EXTENSION jsonb_set_lax;
SELECT jsonb_set_lax('{"a":1}'::jsonb, '{a}', NULL, true, 'use_json_null');
SELECT jsonb_set_lax('{"a":1}'::jsonb, '{a}', NULL, true, 'delete_key');
```

The extension installs into `pg_catalog` and is not relocatable. Its source deliberately refuses PostgreSQL 13 and later, where the core implementation should be used instead. Distinguish SQL NULL from JSON `null`, validate the treatment string, and test missing paths and arrays before migrating logic between this backport and a newer server.
