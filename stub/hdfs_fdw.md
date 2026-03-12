


## Usage

> [hdfs_fdw: Foreign data wrapper for remote HDFS servers](https://github.com/EnterpriseDB/hdfs_fdw)

### Create Server

```sql
CREATE EXTENSION hdfs_fdw;

CREATE SERVER hdfs_server FOREIGN DATA WRAPPER hdfs_fdw
  OPTIONS (host '127.0.0.1', port '10000', client_type 'hiveserver2');
```

**Server Options:** `host` (default `localhost`), `port` (default `10000`), `client_type` (`hiveserver2` or `spark`, default `hiveserver2`), `auth_type` (`NOSASL` or `LDAP`), `connect_timeout` (default 300), `fetch_size` (default 10000), `log_remote_sql` (default `false`), `use_remote_estimate` (default `false`), `enable_join_pushdown` (default `true`), `enable_aggregate_pushdown` (default `true`), `enable_order_by_pushdown` (default `true`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR postgres SERVER hdfs_server
  OPTIONS (username 'hive_user', password 'hive_password');
```

For NOSASL authentication, omit the OPTIONS clause entirely.

### Create Foreign Table

```sql
CREATE FOREIGN TABLE weblogs (
  client_ip text,
  http_status_code text,
  uri text,
  request_count bigint
)
SERVER hdfs_server
OPTIONS (dbname 'default', table_name 'weblogs');
```

**Table Options:** `dbname` (default `default`), `table_name` (defaults to foreign table name), `enable_join_pushdown`, `enable_aggregate_pushdown`, `enable_order_by_pushdown`.

### Query

```sql
SELECT client_ip, count(*) FROM weblogs GROUP BY client_ip ORDER BY count(*) DESC LIMIT 10;
```

### Spark Example

```sql
CREATE SERVER spark_server FOREIGN DATA WRAPPER hdfs_fdw
  OPTIONS (host '127.0.0.1', port '10000', client_type 'spark');

CREATE USER MAPPING FOR postgres SERVER spark_server
  OPTIONS (username 'spark_user', password 'spark_pass');

CREATE FOREIGN TABLE spark_table (
  id int,
  name text,
  value double precision
)
SERVER spark_server
OPTIONS (dbname 'default', table_name 'my_table');
```

### Pushdown Features

hdfs_fdw pushes down WHERE clauses, JOINs, aggregate functions, ORDER BY, and LIMIT/OFFSET to the remote Hive/Spark server. Control pushdown at the session level:

```sql
SET hdfs_fdw.enable_join_pushdown = on;
SET hdfs_fdw.enable_aggregate_pushdown = on;
SET hdfs_fdw.enable_order_by_pushdown = on;
SET hdfs_fdw.enable_limit_pushdown = on;
```
