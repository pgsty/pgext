## Usage

Sources:

- [Extension control file](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup.control)
- [Version 1.0 installation SQL](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/pgrollup--1.0.sql)
- [Upstream project status](https://github.com/mikeizbicki/pgrollup/blob/904f19f5d02fb0901fdb6a7297eceefab549d7ae/README.md)

`pgrollup` version `1.0` is an experimental framework that parses pseudo-DDL and creates tables, triggers, functions, and views for incrementally maintained aggregate results.

### Core Workflow

Install the untrusted `plpython3u` language and the upstream Python package in PostgreSQL's Python environment before creating the extension. Preview generated DDL with `dry_run` first.

```sql
CREATE EXTENSION pgrollup;

SELECT pgrollup_parse($rollup$
CREATE INCREMENTAL MATERIALIZED VIEW sales_rollup AS (
  SELECT region, count(*), sum(amount)
  FROM sales
  GROUP BY region
);
$rollup$, dry_run => true);
```

After reviewing the emitted SQL, rerun with `dry_run => false`. The default maintenance mode is `trigger`; `manual` uses `do_rollup`, and `cron` additionally requires `pg_cron`. Use `rollup_mode` to switch mode, `assert_rollup` to compare the rollup with a recomputation, and `drop_rollup` for managed cleanup.

### Supported Aggregates and Caveats

Built-in algebra definitions cover count, sum, min, max, boolean aggregates, average, variance, and standard deviation. Optional HLL, t-digest, TopN, and DataSketches definitions are enabled only when their extensions already exist.

This project creates dynamic SQL, event triggers, data triggers, and untrusted Python functions. Manual and cron maintenance rely on a suitable single-column sequence or rollup key. The upstream README is mostly an unfinished work list and records parser/query limitations, including unsupported outer joins and subqueries. Treat generated objects as experimental: review every command, test inserts/updates/deletes and schema changes, run `assert_rollup` regularly, and keep a separately computed source of truth.
