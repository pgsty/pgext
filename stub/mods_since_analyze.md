## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/mods_since_analyze/mods_since_analyze-1.0.0/README.md)
- [Official extension SQL (mods_since_analyze.sql)](https://api.pgxn.org/src/mods_since_analyze/mods_since_analyze-1.0.0/mods_since_analyze.sql)

`mods_since_analyze` — mods_since_analyze is a PostgreSQL extension that exposes the estimation of number of changed tuples since last analyze through the function pg_stat_get_mod_since_analyze(). Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

The reviewed distribution uses a legacy SQL or non-control installation layout, so it does not establish a modern standalone `CREATE EXTENSION` and upgrade workflow. Follow the pinned upstream installation mechanism and verify the installed objects in an isolated database.

### Important Objects

- `pg_stat_get_mod_since_analyze(oid)` is an extension function and returns `bigint`.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
