## Usage

Sources:

- [Official README](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/README.md)
- [Extension control file](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/autoreindex.control)
- [Reindex worker source](https://github.com/okbob/autoreindex/blob/363e5299bda99f9e5166c228fd82c8845bf77500/reindexdb.c)

autoreindex is an unfinished background-worker prototype that was intended to detect bloated indexes and rebuild them. The checked upstream revision does not actually run a reindex command: it finds and logs candidates, then executes a test block containing a sleep. Do not use it as an automatic index-maintenance system.

### Core Workflow

This is a headless module, not a SQL extension. Add `autoreindex` to `shared_preload_libraries` and restart PostgreSQL only in an isolated test instance:

```ini
shared_preload_libraries = 'autoreindex'
```

The controller visits connectable databases and starts workers. The current worker excludes primary and system indexes, logs candidate rows, and then sleeps; there is no completed `CREATE INDEX CONCURRENTLY` or equivalent maintenance path.

### Important Objects

- `autoreindex.diagnostics` controls diagnostic logging.
- `autoreindex.hide_context_log_messages` controls context-message suppression.
- `autoreindex.max_workers` and `autoreindex.max_workers_per_database` limit background workers.
- `autoreindex.naptime` controls the worker interval.

### Operational Notes

Because the library must be preloaded, installation changes postmaster startup behavior even though it does not provide useful reindexing. The prototype is from 2016, has a control file but no extension SQL script, and contains hard-coded test behavior. Keep it out of production and use maintained PostgreSQL maintenance tooling for index-bloat policy and concurrent rebuilds.
