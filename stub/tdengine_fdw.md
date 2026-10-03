## Usage

Sources:

- [Official README](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/README.md)
- [Extension control file](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/tdengine_fdw.control)
- [Installation SQL](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/tdengine_fdw--1.0.sql)
- [Build configuration](https://github.com/Amonginin/TDengine_fdw/blob/508b3f53b02add0a534e8dc47288969a950f715e/Makefile)

`tdengine_fdw` accesses TDengine time-series data through WebSocket connections. The reviewed implementation supports PostgreSQL 15/16 and documents TDengine 3.4+ with a compatible C client.

### Core Workflow

Use per-user mappings for credentials and align foreign-table columns with the remote table.

```sql
CREATE EXTENSION tdengine_fdw;
CREATE SERVER tdengine_svr FOREIGN DATA WRAPPER tdengine_fdw
  OPTIONS (host 'localhost', port '6041', dbname 'mydb');
CREATE USER MAPPING FOR CURRENT_USER SERVER tdengine_svr
  OPTIONS (user 'app_user', password 'replace-with-password');
CREATE FOREIGN TABLE sensor_data (ts timestamptz, temperature float4)
  SERVER tdengine_svr OPTIONS (table 'sensor_data');
SELECT * FROM sensor_data WHERE temperature > 25;
```

### Capabilities and Limits

Supports reads, predicate pushdown, insertion and `IMPORT FOREIGN SCHEMA`. Table options include `table`, `tags` and `schemaless`. The latter exposes tag/field data through JSONB. The client dependency is `libtaos >= 3.3.6.0`; the documented connection port is `6041`. DELETE, aggregate pushdown and other PostgreSQL majors are listed as unfinished, so do not infer support. External writes need independent durability and rollback planning.
