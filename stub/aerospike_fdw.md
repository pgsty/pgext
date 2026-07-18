## Usage

Sources:

- [README at the reviewed commit](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/README)
- [Extension control file](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw.control)
- [Extension install SQL](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw--1.0.sql)
- [FDW implementation](https://github.com/sanglijunfengDev/aerospike_fdw/blob/61f0e5143d73dc1d5ce15249daab98ca8e7a46ea/aerospike_fdw.c)

`aerospike_fdw` is an experimental foreign data wrapper for reading records from an Aerospike namespace and set and inserting records through a PostgreSQL foreign table. The server connection is process-global: `aerospike.as_server_ip` and `aerospike.as_server_port` are postmaster-start settings, so the library must be preloaded and the server restarted.

### Configuration and Foreign Table

The upstream build requires the Aerospike C client and its Lua support files. Its implementation also expects Lua files under the hard-coded `/home/pg/lua` tree.

```conf
shared_preload_libraries = 'aerospike_fdw'
aerospike.as_server_ip = '127.0.0.1'
aerospike.as_server_port = 3000
```

After restarting PostgreSQL, create the wrapper and map an Aerospike set. The `key` option names the foreign-table column used as the Aerospike record key.

```sql
CREATE EXTENSION aerospike_fdw;

CREATE SERVER aerospike_server
  FOREIGN DATA WRAPPER aerospike_fdw;

CREATE FOREIGN TABLE aerospike_demo (
  id bigint,
  value text
)
SERVER aerospike_server
OPTIONS (namespace 'test', set 'demo', key 'id');
```

### Caveats

- The reviewed FDW routine exposes scans and inserts, but not updates or deletes. Its validator accepts only the foreign-table options `namespace`, `set`, and `key` and does not verify that all three are present.
- The library connects to Aerospike during PostgreSQL startup and calls process termination on some connection failures. Test startup and failure handling away from production.
- Scan execution generates and uploads Aerospike Lua UDF files and uses hard-coded server-side paths. Those paths, file permissions, and UDF privileges are deployment requirements, not ordinary SQL configuration.
- The upstream source dates from 2016 and publishes no current PostgreSQL or Aerospike compatibility matrix. Treat version `1.0` as legacy experimental code and validate it against the exact server and client versions before use.
