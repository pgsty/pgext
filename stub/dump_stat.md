## Usage

Sources:

- [Official documentation](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/dump_stat--1.0.sql)
- [Control file](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/anyarray_elemtype.control)
- [Version 1.0 SQL](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/dump_stat--1.0.sql)
- [Native helper](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/anyarray_elemtype.c)

`dump_stat` 1.0 exports optimizer statistics as SQL from the historical Postgres Pro 9.5 branch. This entry documents that exact fork revision; it does not establish compatibility with current PostgreSQL releases.

### Core Workflow

After installing the matching fork's extension files, a suitably privileged administrator can generate and inspect statements:

```sql
CREATE EXTENSION dump_stat;
SELECT * FROM dump_statistic('public'::text);
```

The result is SQL text, not an automatically applied restore. Review it before using it in a compatible test environment. The generated statements update or insert rows in the destination's `pg_catalog.pg_statistic` and resolve table, type, column, and operator identities by name.

### Functions and Boundaries

`dump_statistic()` exports available statistics outside system schemas. `dump_statistic(schema_name text)` restricts the schema; `dump_statistic(schema_name text, table_name text)` and `dump_statistic(relid oid)` select one relation. Additional helpers resolve qualified object names; `anyarray_elemtype()` is implemented in the shared library.

Installation uses a C module plus PL/pgSQL routines, without preload. Reading raw catalog statistics and applying the generated catalog writes require appropriate administrative privileges. Sensitive sample values can appear in the output. Target objects and catalog layouts must match; the source hard-codes the statistics layout of its historical branch, so do not treat the output as a portable cross-version dump or a replacement for normal backups and ANALYZE.
