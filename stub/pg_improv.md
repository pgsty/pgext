## Usage

Sources:

- [Release changelog](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.1/CHANGELOG.md)

- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.1/crates/pg_improv/README.md)
- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.1/crates/pg_improv/pg_improv.control)
- [Official source](https://api.pgxn.org/src/pg_improv/pg_improv-0.12.1/crates/pg_improv/Cargo.toml)

`pg_improv` 0.12.1 binds PostgreSQL tables to multidimensional models, evaluates formulas and materializes results. It supports PostgreSQL 15–18 and requires superuser installation. This example uses manual synchronization, which works without preloading the library.

### Core Workflow

```sql
CREATE EXTENSION pg_improv;
CREATE TABLE sales (quarter text, price float8, qty float8);
INSERT INTO sales VALUES ('Q1', 10, 100), ('Q2', 12, 200);
SELECT improv.create_model('plan');
SELECT improv.set_sync_mode('plan', 'manual');
SELECT improv.bind_category('plan', 'Time', 'sales', 'quarter');
SELECT improv.bind_measure('plan', 'Price', 'sales', 'price');
SELECT improv.bind_measure('plan', 'Quantity', 'sales', 'qty');
SELECT improv.define('plan', 'Revenue = Price * Quantity');
SELECT improv.recalc('plan');
SELECT * FROM improv_plan.revenue;
SELECT * FROM improv.status('plan');
```

Bind categories before their measures. Duplicate coordinates require an explicit aggregate; null values do not create model cells. Result tables live in a model-specific schema. `improv.define` accepts `replace => true` to replace a formula, and `improv.recalc` returns the number of changed result rows.

### Synchronization Modes

`improv.set_sync_mode` chooses manual, trigger, commit or logical synchronization. Trigger mode works without preloading; commit mode needs session or server preloading. The default logical mode needs `shared_preload_libraries`, a server restart, `wal_level=logical` and the appropriate `improv.database` worker setting. Logical results arrive asynchronously; use explicit recalculation when a transaction must immediately see current results.

Monitor `improv.status` and cap retained WAL with `max_slot_wal_keep_size`. Logical replication slots can retain WAL while a worker stalls. Binding tables can create publications and change replica identity; account for those effects before enabling logical mode.

### Formula and Maintenance Boundaries

The `improv` language supports scalar formulas as well as model expressions, but not arbitrary SQL calls or external I/O. Formula errors normally become null; `improv.raise_errors` changes that policy. Numeric values and large integers pass through floating-point evaluation, so precision beyond 2^53 needs care. `improv.max_cells` limits loaded input cells.

Back up model definitions and result state with database dumps and rehearse restoration; logical mode needs a fresh slot and recomputation after restore. Library upgrades require restarting preloaded processes and updating the extension SQL. Model computation is distinct from an application transaction's source-of-truth tables.

### Ownership and Database-Wide Logical Sync since 0.11.0

Each model now reads and writes as its owner; row-level security protects its catalog, items and cells. Only a superuser or a role allowed to act as the owner may change or recalculate that model. Grant improv_user to a model creator and grant the source-table privileges it needs; the administrator role does not automatically become every model owner. A shared result table still needs ordinary grants. what_if can reveal owner-visible data, so grant access deliberately. Logical models in one database now share one slot and publication. Upgrade with the matching library and SQL scripts, restore roles before model dumps, and monitor retained WAL.

### 0.12.1 Release Note

The 0.12.0 SQL migration adds `improv.measure.format` for persisted display formatting; update the extension before saving formats. Release 0.12.1 changes the version from 0.12.0 without additional code changes and restores RHEL 8 (glibc 2.28) compatible artifacts. Upstream still uses pgrx 0.17.0. Logical synchronization still requires `shared_preload_libraries`, `wal_level=logical` and a restart; trigger and manual modes do not.
