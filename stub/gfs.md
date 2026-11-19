## Usage

Sources:

- [README.md](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/README.md)
- [gfs.control](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/gfs.control)
- [Cargo.toml](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/Cargo.toml)
- [src/lib.rs](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/src/lib.rs)
- [src/sql/schema.sql](https://github.com/Guepard-Corp/gfs/blob/v0.5.3/crates/extensions/gfs/src/sql/schema.sql)

`gfs` uses a planner hook for lazy logical copies of a remote PostgreSQL database. The current Rust module retains extension version 0.0.1 inside GFS distribution 0.5.3. The old C table-access-method implementation in `c-ref/` is reference code, not the current interface. Cargo provides PostgreSQL 16–18 features.

### Enablement and Routing

A superuser installs the non-relocatable extension and its runtime prerequisites:

```sql
CREATE EXTENSION postgres_fdw;
CREATE EXTENSION dblink;
CREATE EXTENSION gfs;
```

Enable `session_preload_libraries = 'gfs'` for the clone database and reconnect, preserving existing entries. CREATE EXTENSION alone does not install the planner hook in subsequent sessions. The GFS bootstrap prepares ordinary local tables plus `postgres_fdw` foreign tables and user mappings; parallel hydration and source diagnostics use `dblink`. Keep connection credentials in restricted mappings.

Register a prepared local/foreign table pair with `gfs.register_clone(local, source_ref, key_col)`. Integer key ranges hydrate missing ranges; selective predicates can hydrate a bounded slice; unbounded joins and aggregates can be redirected to foreign tables. Fully cached tables serve locally. `gfs.warm(local, p_freeze, p_conn)` materializes a table; `gfs.unregister_clone` removes registration. Ordinary remote reads are not a frozen point-in-time database snapshot.

### Administration

`gfs.clone_source`, `gfs.cached`, `gfs.cached_predicate` and `gfs.clones` expose routing/cache state. `gfs.calibrate` measures source-link weights; `gfs.cost` sets policies and `gfs.budget` throttles source contacts. The current SQL also provides source drift checks, `gfs.pull`, resynchronization and explicit freeze helpers; review their destructive/overwrite choices before applying them to a diverged clone.

Treat source references, key columns and predicates as administrator-controlled inputs; the planner/SPI path uses unsafe native FFI. The bootstrap can omit foreign keys and generated columns, so it is not an exact physical replica. The README lists unfinished hardening and an outdated test suite; code presence alone does not certify those paths. With unchanged SQL version 0.0.1, replacing a library does not migrate old prototype catalogs. Plan a controlled clone recreation or a reviewed explicit migration rather than assuming ALTER EXTENSION supplies an upgrade.
