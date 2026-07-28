## Usage

Sources:

- [Official upstream README](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/README.md)
- [Official extension control file (pg_wal_ext.control)](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/pg_wal_ext.control)
- [Official extension SQL (pg_wal_ext--1.0.sql)](https://github.com/misachi/pg_wal_ext/blob/b8160bbd5956a7346bd2b61d4b426395881a1798/pg_wal_ext--1.0.sql)

`pg_wal_ext` — Read WAL files and generate SQL from WAL. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_wal_ext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_xlog_records(IN xlog_file_path text, OUT page_num int4, OUT txn_id xid, OUT xlog_type text, OUT commit_ts timestamptz, OUT generated_sql text)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
