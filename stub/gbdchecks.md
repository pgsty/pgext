## Usage

Sources:

- [Pinned official README](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/README.md)
- [Pinned extension SQL](https://github.com/huseynsnmz/gbdchecks/blob/ff7090d1a2a7064df6d43deeab4f10cf7c09ff0d/sql/gbdchecks--1.1.sql)

`gbdchecks` installs a collection of diagnostic views and privilege-reporting functions. It covers query activity, bloat, locks, cache hit ratio, indexes, sequential scans, vacuum status, and effective privileges; it is a snapshot-oriented troubleshooting aid, not a scheduler or alerting service.

### Core Workflow

The extension uses the fixed `gbdchecks` schema and requires `pg_stat_statements`:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION gbdchecks;

SELECT * FROM gbdchecks.check_blocked_statements;
SELECT * FROM gbdchecks.check_unused_indexes;
SELECT * FROM gbdchecks.all_privs('app_user');
```

For `gbdchecks.check_queries`, configure `pg_stat_statements` in `shared_preload_libraries` and restart PostgreSQL before collecting statistics. The other views can still be used without that query-statistics view.

### Diagnostic Views

- Workload and waits: `check_queries`, `check_blocked_statements`, and `check_locks`.
- Storage and maintenance: `check_bloat`, `check_index_sizes`, and `check_vacuum_stats`.
- Access patterns: `check_hit_ratio`, `check_index_usage`, `check_seq_scans`, and `check_unused_indexes`.

These views expose live catalog/statistics data. A low scan count, an estimated bloat ratio, or an autovacuum expectation is evidence to investigate, not by itself an instruction to drop an index or rewrite a table.

### Privilege Reports

Functions accept a role name as `text`: `table_privs`, `database_privs`, `tablespace_privs`, `fdw_privs`, `fsrv_privs`, `language_privs`, `schema_privs`, `view_privs`, and `sequence_privs`. `all_privs(text)` combines their output.

### Security and Compatibility

Installation revokes public access to the extension schema, functions, and views, then grants usage/execution/select to the creating role. Grant only the required privileges to other roles, for example:

```sql
GRANT USAGE ON SCHEMA gbdchecks TO monitor;
GRANT SELECT ON ALL TABLES IN SCHEMA gbdchecks TO monitor;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA gbdchecks TO monitor;
```

The project is archived, and its SQL references historical `pg_stat_statements` columns such as `total_time` and `mean_time`. Those names changed in newer PostgreSQL/extension releases, so validate creation and every view against the target server version before operational use.
