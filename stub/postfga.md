## Usage

Sources:

- [Official extension SQL](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/sql/postfga--1.0.0.sql)
- [Official preload and initialization source](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/src/init.c)
- [Official runtime configuration source](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/src/guc.c)
- [Official testing guide](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/TESTING.md)

`postfga` version `1.0.0` is a preload-only preview extension that sends fine-grained authorization checks and tuple mutations from PostgreSQL to an OpenFGA server over gRPC. Its background worker, shared-memory queue, network calls, and external side effects make it an infrastructure component, not a transaction-local SQL helper.

### Prerequisites and Configuration

Add the library and the source-confirmed `fga.*` settings to PostgreSQL configuration, then restart the server:

```conf
shared_preload_libraries = 'postfga'
fga.endpoint = 'dns:///openfga:8081'
fga.store_id = '01HXXXXXXXXXXXXXXX'
fga.model_id = ''
fga.cache_enabled = off
```

```sql
CREATE EXTENSION postfga;
```

The current C source defines `fga.endpoint`, `fga.store_id`, `fga.model_id`, `fga.cache_enabled`, `fga.cache_size`, and `fga.cache_ttl_ms`. Parts of the testing guide use stale `postfga.endpoint` names; use the names registered by the reviewed source.

### Core Workflow

Check a relationship, then add or remove tuples through the extension API:

```sql
SELECT fga_check(
  'document', 'budget',
  'user', 'alice',
  'reader'
);

SELECT fga_write_tuple(
  'document', 'budget',
  'user', 'alice',
  'reader'
);

SELECT fga_delete_tuple(
  'document', 'budget',
  'user', 'alice',
  'reader'
);
```

### Important Objects

- `fga_check(...)` returns whether OpenFGA allows the relationship.
- `fga_write_tuple(...)` and `fga_delete_tuple(...)` mutate tuples in the remote store.
- `fga_create_store(name)` and `fga_delete_store(id)` manage OpenFGA stores.
- `fga_stats()` reports extension queue/cache statistics.
- `fga_fdw` is registered, but the reviewed SQL exposes no documented foreign-table workflow.

### Operational Caveats

`postfga` refuses to load outside `shared_preload_libraries`, so enabling or removing it requires a restart. OpenFGA tuple writes are external network side effects and are not rolled back when the surrounding PostgreSQL transaction aborts. Restrict EXECUTE on mutation and store-management functions, secure the gRPC channel and endpoint, and test denial behavior during timeouts, worker failure, queue exhaustion, OpenFGA outages, and stale cache entries. The source and testing guide contain naming and roadmap inconsistencies, including test-only functions and an FDW workflow described as future work. Treat the build as preview code, pin the commit, and validate authorization fail-closed semantics before any sensitive use.
