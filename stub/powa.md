

## Usage

> [powa: PostgreSQL Workload Analyzer](https://github.com/powa-team/powa)

PoWA is a workload analysis framework that collects performance statistics and provides real-time charts through a web UI. The core extension (`powa-archivist`) snapshots statistics from multiple stat extensions.

### Architecture

PoWA consists of several components:

- **powa-archivist**: The PostgreSQL extension that collects and stores statistics
- **powa-web**: A web UI for visualizing performance data
- **pg_stat_statements**: Required for query statistics
- **pg_qualstats**: Optional, provides predicate statistics and index suggestions
- **pg_stat_kcache**: Optional, provides OS-level metrics (CPU, I/O)
- **pg_wait_sampling**: Optional, provides wait event sampling

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
