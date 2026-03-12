


## Usage

> [postgres_fdw: Foreign data wrapper for remote PostgreSQL servers](https://www.postgresql.org/docs/current/postgres-fdw.html)

### Create Server

```sql
CREATE EXTENSION postgres_fdw;

CREATE SERVER remote_pg FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host '192.168.1.10', port '5432', dbname 'remotedb');
```

**Server Options:** Any libpq connection parameter (`host`, `port`, `dbname`, etc.), plus `use_remote_estimate` (default `false`), `fdw_startup_cost` (default `100`), `fdw_tuple_cost` (default `0.2`), `extensions` (comma-separated list of extensions installed on both servers), `fetch_size` (default `100`), `batch_size` (default `1`), `keep_connections` (default `on`), `parallel_commit` (default `false`), `parallel_abort` (default `false`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR local_user SERVER remote_pg
  OPTIONS (user 'remote_user', password 'secret');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE remote_table (
  id integer NOT NULL,
  name text,
  value numeric
)
SERVER remote_pg
OPTIONS (schema_name 'public', table_name 'source_table');
```

**Table/Column Options:** `schema_name` (default: local schema name), `table_name` (default: local table name), `column_name` (per-column, remote column name), `updatable` (default `true`), `truncatable` (default `true`), `async_capable` (default `false`).

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA remote_schema
  FROM SERVER remote_pg
  INTO local_schema;

-- Import specific tables
IMPORT FOREIGN SCHEMA remote_schema
  LIMIT TO (table1, table2)
  FROM SERVER remote_pg
  INTO local_schema;
```

**Import Options:** `import_collate` (default `true`), `import_default` (default `false`), `import_generated` (default `true`), `import_not_null` (default `true`).

### CRUD Operations

```sql
SELECT * FROM remote_table WHERE id > 100;
INSERT INTO remote_table VALUES (1, 'test', 42.0);
UPDATE remote_table SET value = 100 WHERE id = 1;
DELETE FROM remote_table WHERE id = 1;
TRUNCATE remote_table;
```

### Query Optimization

postgres_fdw automatically pushes down WHERE clauses, JOINs between tables on the same server, aggregate functions, ORDER BY, and LIMIT/OFFSET. View the remote query with:

```sql
EXPLAIN VERBOSE SELECT * FROM remote_table WHERE id = 1;
```

Use the `extensions` option to allow function/operator pushdown from those extensions:

```sql
ALTER SERVER remote_pg OPTIONS (ADD extensions 'postgis,hstore');
```

### Asynchronous Execution

Enable concurrent scans across multiple foreign servers:

```sql
ALTER FOREIGN TABLE remote_table OPTIONS (ADD async_capable 'true');
```

### Connection Management

```sql
SELECT * FROM postgres_fdw_get_connections(true);   -- List connections
SELECT postgres_fdw_disconnect('remote_pg');         -- Close specific connection
SELECT postgres_fdw_disconnect_all();                -- Close all connections
```

### Transaction Behavior

Remote transactions use SERIALIZABLE if the local transaction is SERIALIZABLE; otherwise REPEATABLE READ. Two-phase commit is not currently supported.
