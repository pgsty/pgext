## Usage

Sources:

- [docs/operate/postgres.md](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/docs/operate/postgres.md)
- [integrations/tapi-postgres/tapi_pg--0.1.sql](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg--0.1.sql)
- [integrations/tapi-postgres/tapi_pg.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg.c)
- [integrations/tapi-postgres/ddl.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/ddl.c)
- [integrations/tapi-postgres/txn.c](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/txn.c)
- [LICENSE](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/LICENSE)
- [integrations/tapi-postgres/tapi_pg.control](https://github.com/mumoshu/tapirs/blob/546de00b0dc5d89a2f4121815e6cf7b6e41876fd/integrations/tapi-postgres/tapi_pg.control)

`tapi_pg` 0.1 links a C foreign data wrapper to the Rust Tapirs C API. It needs a running cluster, shared preload and a server restart before SQL use. Configure `tapirs.cluster_config` with the deployment's discovery endpoint and catalog settings.

### Core Workflow

```ini
shared_preload_libraries = 'tapi_pg'
```

```sql
CREATE EXTENSION tapi_pg;
SELECT tapi_pg_status();
CREATE SERVER tapirs_store FOREIGN DATA WRAPPER tapi_pg;
CREATE FOREIGN TABLE tapirs_users (id integer, name text)
    SERVER tapirs_store
    OPTIONS (keyspace 'app', "table" 'users', primary_key 'id');
```

### Operational Boundaries

`tapi_pg_status()` checks cluster reachability. Foreign tables require `keyspace`, `table` and `primary_key` options; optional `secondary_indexes` are declared when creating the table, not through CREATE INDEX. Transactions use Tapirs optimistic concurrency; serialization failures require application retry. Savepoints, two-phase commit and automatic remote table renaming are not supported by the reviewed implementation. The integration does not wire cluster TLS, so its network boundary matters. Installation requires superuser privileges and a binary matching the PostgreSQL major; upstream does not declare a general major-version range. The integration README overstates table renaming; the operator guide and DDL code document it as deferred.
