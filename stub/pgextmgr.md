## Usage

Sources:

- [Official repository README](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/README.md)
- [Extension control file](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/pgextmgr.control)
- [Manager implementation](https://github.com/cmu-db/pgextmgrext/blob/a412421903e0ddb5fb51a195ab92f70f698bb945/pgextmgr/src/lib.rs)

`pgextmgr` is an experimental pgrx hook manager that lets specially instrumented PostgreSQL extension plugins register ordered planner, executor, and query-rewriter callbacks. It can inspect those registrations and enable or disable a plugin at runtime. It is a prototype framework for compatible plugins, not a general manager for arbitrary installed extensions.

### Core Workflow

The upstream test configuration preloads the manager and its instrumented plugins so their hooks are installed during server startup:

```conf
shared_preload_libraries = 'pgextmgr,pgext_pg_stat_statements,pgext_pg_hint_plan,pgext_pg_poop'
```

Restart PostgreSQL after changing `shared_preload_libraries`, then create `pgextmgr` and the plugins built for its ABI. Inspect their status and hook order before changing it:

```sql
CREATE EXTENSION pgextmgr;

SELECT * FROM pgextmgr.all();
SELECT * FROM pgextmgr.hooks();

SELECT pgextmgr.disable('pgext_pg_poop');
SELECT pgextmgr.enable('pgext_pg_poop');
```

`pgextmgr.all()` reports each registered plugin's order and enabled/disabled state. `pgextmgr.hooks()` reports the ordered plugin chain for each supported hook. The state-changing functions return the number of affected registrations.

### Function Index

- `pgextmgr.all()` lists registered plugins and their current state.
- `pgextmgr.hooks()` lists the manager's planner, executor, and rewriter callback order.
- `pgextmgr.enable(text)` and `pgextmgr.disable(text)` change one known plugin.
- `pgextmgr.enable_all()` and `pgextmgr.disable_all()` change every registered plugin.

### Operational Notes

`pgextmgr` installs into the fixed `pgextmgr` schema and requires superuser privileges to create. The reviewed source targets the PostgreSQL 15-era pgrx environment. Its plugin quickstart uses custom linker flags and the framework's `__pgext_before_init` and `__pgext_after_init` symbols; an ordinary extension that was not built for this ABI will not appear as a manageable plugin.

Enable/disable state is held in backend process memory rather than a durable catalog. Confirm behavior after reconnects and restarts, and test hook ordering and failure handling with the exact plugin set. The upstream README presents the framework as work in progress, so use it for controlled research rather than assuming production compatibility or stable APIs.
