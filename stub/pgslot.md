## Usage

Sources:

- [pgslot.control](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/pgslot.control)
- [README.md](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/README.md)
- [sql/pgslot--0.5.sql](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/sql/pgslot--0.5.sql)
- [scripts/roles.sql](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/scripts/roles.sql)
- [src/README.md](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/src/README.md)
- [Makefile](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/Makefile)

`pgslot` records replication-slot history and derives WAL retention, consumption rates and downstream pipeline health. Its SQL extension does not advance or remove replication slots.

### Core Workflow

```sql
CREATE EXTENSION pgslot;
SELECT pgslot.collect();
SELECT * FROM pgslot.slot_health;
SELECT * FROM pgslot.wal_summary;
SELECT pgslot.prune(168);
```

### Operational Boundaries

Install with a superuser. The fixed `pgslot` schema exposes `slot_health`, `wal_summary` and `slot_pipeline`; `report_metric` accepts adapter metrics. Apply the upstream role grants: `pgslot_monitor` reads views, `pgslot_collector` runs collection/pruning, and `pgslot_adapter` reports metrics. The raw history tables and privileged functions are revoked from PUBLIC. Schedule `collect()` and `prune(integer)` externally, or use the optional C worker: preload `pgslot`, restart, and set `pgslot.database` and an explicit login `pgslot.role` in the collector group. That worker covers one database; upstream reports testing on PostgreSQL 15. CLI/TUI programs are optional companions.
