## Usage

Sources:

- [Official README](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/README.md)
- [FDW option and callback implementation](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/hive_fdw.c)
- [Import foreign schema documentation](https://github.com/pgsql-io/hive_fdw/blob/600587892b0a811c129c81d19e01b3b4dd1c87b2/IMPORT_FOREIGN_SCHEMA.md)

`hive_fdw` version `4.0` provides read access from PostgreSQL to HiveServer2 through JDBC. It starts a JVM inside the PostgreSQL process and requires OpenJDK 11 plus the matching Hive client JARs.

### Runtime Prerequisites

Set `PGHOME`, `HIVECLIENT_JAR_HOME`, and `HIVE_FDW_CLASSPATH` in the PostgreSQL server environment before startup. They are read when the JVM initializes, so changing them requires restarting the server process that hosts the FDW.

### Core Workflow

```sql
CREATE EXTENSION hive_fdw;

CREATE SERVER hive_serv FOREIGN DATA WRAPPER hive_fdw
OPTIONS (host 'localhost', port '10000');

CREATE USER MAPPING FOR CURRENT_USER SERVER hive_serv
OPTIONS (username 'hive_user', password 'secret');

CREATE FOREIGN TABLE orders (id integer)
SERVER hive_serv
OPTIONS (schema 'example', table 'orders');

SELECT * FROM orders LIMIT 5;
```

The implementation accepts `host` and `port` on the server, `username` and `password` on the user mapping, and `schema`, `table`, or a custom `query` on a foreign table. The README's prose calls the table options `schema_name` and `table_name`, but version `4.0` source and examples validate the actual spellings `schema` and `table`.

### Import and Pushdown

`IMPORT FOREIGN SCHEMA` can create local foreign-table definitions from Hive metadata. The FDW also implements join pushdown. Inspect plans with `EXPLAIN` and validate type mappings before relying on pushdown or imported definitions.

### Operational Notes

The callback table has no insert, update, or delete methods, so treat foreign tables as read-only. User-mapping passwords are stored in PostgreSQL catalogs, and each query crosses the JDBC/JVM boundary; restrict catalog access, use a least-privileged Hive account, and plan for remote latency and HiveServer2 failures.
