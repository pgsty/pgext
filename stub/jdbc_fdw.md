


## Usage

> [jdbc_fdw: Foreign data wrapper for remote servers available over JDBC](https://github.com/pgspider/jdbc_fdw)

### Create Server

```sql
CREATE EXTENSION jdbc_fdw;

CREATE SERVER jdbc_server FOREIGN DATA WRAPPER jdbc_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://remotehost:5432/mydb',
    jarfile '/usr/share/java/postgresql.jar',
    maxheapsize '256'
  );
```

**Server Options:** `drivername` (required, JDBC driver class), `url` (required, JDBC connection URL), `jarfile` (required, absolute path to JDBC driver JAR), `querytimeout` (query timeout in seconds), `maxheapsize` (JVM heap size in MB, minimum 1).

### Create User Mapping

```sql
CREATE USER MAPPING FOR CURRENT_USER SERVER jdbc_server
  OPTIONS (username 'dbuser', password 'dbpass');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE remote_table (
  id integer OPTIONS (key 'true'),
  name text,
  value numeric
)
SERVER jdbc_server
OPTIONS (table_name 'schema.tablename');
```

Set `key 'true'` on primary key columns to enable UPDATE and DELETE operations.

### Query Remote Data

```sql
SELECT * FROM remote_table WHERE id > 100;
```

### Execute Arbitrary SQL with jdbc_exec

The `jdbc_exec` function executes SQL against the remote database and returns result sets:

```sql
SELECT * FROM jdbc_exec('jdbc_server', 'SELECT id, name FROM remote_schema.remote_table WHERE status = 1')
  AS t(id integer, name text);
```

This is useful for executing queries that go beyond the foreign table definition, including DDL or complex queries on the remote server.
