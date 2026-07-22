## Usage

Sources:

- [Official PGXN usage documentation](https://pgxn.org/dist/tab_tier/doc/tab_tier.html)
- [Official PGXN distribution page](https://pgxn.org/dist/tab_tier/)

`tab_tier` is a historical PL/pgSQL partition-maintenance and archival toolkit. Unlike trigger-routed partitioning, recent rows remain in a root table; scheduled jobs later move older rows into inheritance children, and optional long-term storage can be a local or foreign table.

```sql
CREATE EXTENSION tab_tier;
SELECT tab_tier.register_tier_root('comm', 'yell', 'created_dt');
SELECT tab_tier.bootstrap_tier_parts('comm', 'yell');
SELECT tab_tier.migrate_tier_data('comm', 'yell', '201301');
```

Schedule `cap_tier_partitions()` more frequently than `part_period`, then run `migrate_all_tiers()` so an eligible destination exists before data movement. The default design keeps `root_retain` in the root, uses `part_period` for child boundaries, and archives after `lts_threshold`; inspect `tier_root`, `tier_part`, and `tier_config` for effective state.

Migration copies rows and then deletes them from the root. Bootstrap large tables in resumable per-partition transactions: `flush_tier_data()` and `flush_all_tiers()` can move all eligible rows in one very large transaction, and any error rolls that work back. Test locks, WAL, replication lag, disk headroom, constraints, indexes, triggers, late/future timestamps, and concurrent writes.

Archival can move data to `lts_target` and drop the old child. Verify row counts and durable backups before enabling it. The extension uses old inheritance partitioning, broad administrative privileges through `tab_tier_role`, and 2014-era assumptions; prefer maintained declarative-partition tooling for new systems.
