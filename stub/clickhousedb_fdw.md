## Usage

Sources:

- [Official README](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/README.md)
- [Official user guide](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/docs/USER_GUIDE.md)
- [Extension control file](https://github.com/Percona-Lab/clickhousedb_fdw/blob/c6b386f93d97430ea4da1fd680f005a84a2cf3ec/clickhousedb_fdw.control)

`clickhousedb_fdw` exposes ClickHouse tables through PostgreSQL's foreign-data-wrapper API. It supports remote reads and inserts, with aggregate and join pushdown, but it is not a general read/write compatibility layer: upstream documents no `UPDATE` or `DELETE` support.

### Core Workflow

Install and configure the ClickHouse ODBC driver on every PostgreSQL host before creating the extension. The server options identify the ClickHouse database, ODBC driver, and endpoint.

```sql
CREATE EXTENSION clickhousedb_fdw;

CREATE SERVER clickhouse_svr
  FOREIGN DATA WRAPPER clickhousedb_fdw
  OPTIONS (
    dbname 'default',
    driver '/usr/lib/libclickhouseodbc.so',
    host '127.0.0.1'
  );

CREATE USER MAPPING FOR CURRENT_USER
  SERVER clickhouse_svr;
```

Define a foreign table whose columns and types match the remote ClickHouse table, then query or insert through it.

```sql
CREATE FOREIGN TABLE events_ch (
  event_time timestamp,
  user_id bigint,
  payload text
)
SERVER clickhouse_svr
OPTIONS (table_name 'events');

SELECT user_id, count(*)
FROM events_ch
GROUP BY user_id;

INSERT INTO events_ch VALUES (clock_timestamp(), 42, 'created');
```

Use verbose plans to see the SQL sent to ClickHouse and confirm that filters, aggregates, or joins were actually pushed down.

```sql
EXPLAIN (VERBOSE)
SELECT user_id, count(*)
FROM events_ch
WHERE event_time >= timestamp '2026-01-01'
GROUP BY user_id;
```

### Important Objects

- `clickhousedb_fdw`: the foreign-data wrapper installed by the extension.
- Server options `dbname`, `driver`, and `host`: select the ClickHouse database, ODBC driver library, and endpoint.
- Foreign-table option `table_name`: maps a PostgreSQL foreign table to its ClickHouse table.
- Remote `SELECT` and `INSERT`: the documented DML surface; aggregate and join pushdown are planner optimizations, not guarantees for every query shape.

### Operational Notes

The reviewed upstream revision labels the project as under heavy development and says it was tested with PostgreSQL 11–13. Validate the exact PostgreSQL, ClickHouse, ODBC-driver, and type-conversion combination before deployment. Complex joins and some ClickHouse settings are explicitly incomplete. Network, ODBC, and remote execution errors occur inside the PostgreSQL statement, so set appropriate statement timeouts and test failure recovery. Do not assume transaction, constraint, locking, or rollback semantics match local PostgreSQL tables.
