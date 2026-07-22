## Usage

Sources:

- [Official upstream documentation](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/README.md)
- [Official option implementation](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/option.c)
- [Official extension SQL](https://github.com/heimir-sverrisson/jdbc2_fdw/blob/758caf50c674577144b644f36432c4736ec99627/jdbc2_fdw--1.0.sql)

`jdbc2_fdw` is an experimental foreign data wrapper for querying databases through JDBC. It embeds a Java virtual machine in PostgreSQL backend processes and uses a PostgreSQL 9.4-era `postgres_fdw` deparser to push eligible filters to the remote database. The published project describes read-only querying; do not rely on remote `INSERT`, `UPDATE`, or `DELETE` support.

### Core Workflow

Install the extension, define the JDBC driver and connection on a foreign server, keep credentials in a user mapping, and map a remote table:

```sql
CREATE EXTENSION jdbc2_fdw;

CREATE SERVER jdbc_remote
  FOREIGN DATA WRAPPER jdbc2_fdw
  OPTIONS (
    drivername 'org.postgresql.Driver',
    url 'jdbc:postgresql://db.example.com:5432/app',
    jarfile '/opt/jdbc/postgresql.jar',
    querytimeout '30'
  );

CREATE USER MAPPING FOR app_user
  SERVER jdbc_remote
  OPTIONS (username 'remote_user', password 'secret');

CREATE FOREIGN TABLE public.remote_orders (
  id bigint,
  status text
)
SERVER jdbc_remote
OPTIONS (schema_name 'public', table_name 'orders');

SELECT * FROM public.remote_orders WHERE status = 'open';
```

The JDBC driver JAR and embedded JVM libraries must be readable and loadable by every PostgreSQL backend that uses the server.

### Options

Server options include `drivername`, `url`, `querytimeout`, `jarfile`, `maxheapsize`, `use_remote_estimate`, `fdw_startup_cost`, `fdw_tuple_cost`, and `updatable`. User mappings accept `username` and `password`. Foreign tables accept `schema_name`, `table_name`, `updatable`, and `use_remote_estimate`; columns can override the remote name with `column_name`.

### Compatibility and Security

The code is based on old PostgreSQL and JNI interfaces. Verify the exact PostgreSQL, JDK, architecture, and JDBC driver combination before use; a successful extension installation does not prove that the JVM can be created in a backend. Limit catalog visibility and user-mapping administration because passwords are stored in PostgreSQL metadata, require TLS in the JDBC URL, and set a bounded query timeout and JVM heap. Test type conversions, nulls, pushed predicates, cancellation, connection failure, and backend memory use. Treat it as a read-only experimental integration unless the selected fork explicitly documents and tests otherwise.
