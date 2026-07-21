## Usage

Sources:

- [PoWA Archivist 5.2.0 README](https://github.com/powa-team/powa-archivist/blob/REL_5_2_0/README.md)
- [PoWA Archivist 5.2.0 control file](https://github.com/powa-team/powa-archivist/blob/REL_5_2_0/powa.control)
- [PoWA Archivist 5.2.0 release](https://github.com/powa-team/powa-archivist/releases/tag/REL_5_2_0)
- [PoWA documentation](https://powa.readthedocs.io/en/latest/components/powa-archivist/index.html)

PoWA is a workload analysis framework that collects performance statistics and provides real-time charts through a web UI. The core extension (`powa-archivist`) snapshots statistics from multiple stat extensions.

### Architecture

PoWA consists of several components:

- **powa-archivist**: The PostgreSQL extension that collects and stores statistics
- **powa-web**: A web UI for visualizing performance data
- **pg_stat_statements**: Required for query statistics
- **pg_qualstats**: Optional, provides predicate statistics and index suggestions
- **pg_stat_kcache**: Optional, provides OS-level metrics (CPU, I/O)
- **pg_wait_sampling**: Optional, provides wait event sampling

Enable the extension in the repository database. Its control file requires PL/pgSQL, `pg_stat_statements`, and `btree_gist`:

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION btree_gist;
CREATE EXTENSION powa CASCADE;
```

`pg_stat_statements` itself must be configured in `shared_preload_libraries` before PostgreSQL starts.

### Taking Snapshots

PoWA periodically snapshots statistics from all registered extensions:

```sql
-- Manual snapshot
SELECT powa_take_snapshot();

-- Check snapshot status
SELECT * FROM powa_snapshot_metas;
```

### Configuration

```sql
-- Register stat extensions (done automatically on CREATE EXTENSION)
SELECT powa_register_server(hostname => 'localhost', port => 5432);

-- Configure snapshot interval and retention
ALTER EXTENSION powa UPDATE;
```

### Key Tables and Views

```sql
-- View collected query statistics
SELECT * FROM powa_statements;

-- View snapshot history
SELECT * FROM powa_snapshot_metas;
```

### Web UI

The PoWA web interface (installed separately) provides:

- Real-time query performance dashboards
- Per-query drill-down with plan details
- Index suggestions based on `pg_qualstats`
- Wait event analysis
- System resource usage graphs

Documentation: [powa.readthedocs.io](https://powa.readthedocs.io/)

### Version and Deployment Notes

- PoWA Archivist 5.2.0 adds PostgreSQL 19 support, including collectors for the new `pg_stat_recovery` and `pg_stat_lock` statistics views. The core snapshot workflow remains compatible with the 5.1 series.
- `powa` is the database extension; `powa-web` is a separate visualization service, and `powa-collector` is used for remote collection architectures. Installing the extension alone does not deploy the UI.
- Retention and snapshot frequency directly affect repository growth. Monitor the PoWA repository database and size retention for the number of databases, queries, and enabled optional modules.
