## Usage

Sources:

- [Extension initialization and configuration](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw.c)
- [Foreign-table implementation](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/fdw_driver.c)
- [Extension SQL](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw--1.0.sql)

`hbase_fdw` version `1.0` is an experimental, read-only foreign data wrapper that starts an embedded JVM and queries HBase through its Java client. The repository contains no user README, so the interface below is limited to behavior visible in the extension source.

### Server Setup

The module registers shared memory and a background worker only while PostgreSQL processes `shared_preload_libraries`. Configure the JVM before restarting the server:

```conf
shared_preload_libraries = 'hbase_fdw'
hbase_fdw.java_home = '/path/to/java-home'
hbase_fdw.classpath = '/path/to/hbase-client.jar:/path/to/dependencies/*'
```

The implementation searches `hbase_fdw.java_home` for an amd64 `libjvm.so`; verify that the chosen Java layout and HBase client JAR set match these source assumptions.

### Core Workflow

```sql
CREATE EXTENSION hbase_fdw;
CREATE SERVER hbase_srv FOREIGN DATA WRAPPER hbase_fdw;

CREATE FOREIGN TABLE events (
  row_key text OPTIONS (hbase_type 'row_key'),
  payload text OPTIONS (
    hbase_type 'column', family 'f', qualifier 'payload'
  )
) SERVER hbase_srv OPTIONS (hbase_table 'events');

SELECT * FROM events WHERE row_key = 'event-42';
```

`hbase_table` selects the remote table and defaults to the foreign-table name. Column option `hbase_type` accepts `row_key`, `family`, or `column`; a column value can additionally use `family` and `qualifier` to identify the HBase cell.

### Limitations

The FDW exposes scan callbacks only and pushes down only equality on the row-key column. It defines no server or user-mapping connection options; HBase connectivity comes from the Java client configuration available to the embedded JVM. The worker is configured not to restart after failure. Treat this old, minimally documented code as evaluation-only and test server startup, JVM loading, HBase authentication, type conversion, and failure recovery before use.
