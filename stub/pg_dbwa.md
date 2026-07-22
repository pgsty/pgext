## Usage

Sources:

- [Official repository README](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/README.md)
- [Official extension control file](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/pg_dbwa.control)
- [Official version 0.3.1 SQL](https://github.com/alexandersamoylov/pg_dbwa/blob/3a8ea55aafb18ed4337be03a3f896e211b40a988/sql/pg_dbwa--0.3.1.sql)

`pg_dbwa` stores periodic `pg_stat_statements` snapshots and reports workload deltas for local or remote PostgreSQL databases. It is an old, SQL-only workload analyzer whose schema and dependencies require careful compatibility and security review before use.

### Core Workflow

Version 0.3.1 declares hard dependencies on `dblink`, `pg_stat_statements`, `pg_eyes`, and `pg_prttn_tools`, and creates its objects in the `dbwa` schema:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION dblink;
CREATE EXTENSION pg_eyes;
CREATE EXTENSION pg_prttn_tools;
CREATE EXTENSION pg_dbwa;

SELECT dbwa.get_stat('local');
SELECT *
FROM dbwa.show_top_queryes(
    'local',
    timestamp '2026-07-22 00:00:00',
    timestamp '2026-07-22 01:00:00',
    20
);
```

`dbwa.cluster_config` records each monitored cluster and connection information. `dbwa.cluster_operation` selects snapshot operations. `dbwa.get_stat(clustername)` runs enabled operations and records their status. `dbwa.show_top_queryes()` ranks deltas over a time window, while `dbwa.show_stat_query()` returns interval deltas for one `(userid, dbid, queryid)`.

### Security and Compatibility

The reviewed SQL stores remote usernames and passwords in plain columns and constructs a `dblink` connection string from them. Restrict all `dbwa` tables and functions, protect backups, and prefer a minimally privileged monitoring role. Several functions are `SECURITY DEFINER`, concatenate configured operation names into dynamic SQL, and are owned by `postgres`; untrusted users must not be able to alter their configuration or search path.

The 0.3.1 install script uses legacy constructs such as `WITH (OIDS=FALSE)` and old `pg_stat_statements` columns. Verify installation against the exact PostgreSQL and dependency versions instead of assuming modern-server compatibility. Snapshot tables grow over time; schedule retention and account for the monitoring query load.
