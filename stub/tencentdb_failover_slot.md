## Usage

Sources:

- [Official TencentDB Failover Slot documentation](https://cloud.tencent.com/document/product/409/86587)
- [Official TencentDB for PostgreSQL product page](https://cloud.tencent.com/product/postgres)

`tencentdb_failover_slot` is the TencentDB for PostgreSQL provider extension for synchronizing logical replication-slot state to a standby so logical replication can continue after an HA switchover. It is a TencentDB kernel capability, not a portable community extension, and it supports logical slots only.

### Create a Failover Slot

TencentDB creates logical slots as failover slots by default when the instance parameter `tencentdb_force_enable_failover_slot` is enabled:

```sql
CREATE EXTENSION tencentdb_failover_slot;

SELECT *
FROM pg_create_logical_replication_slot(
  'app_slot',
  'pgoutput'
);

SELECT * FROM pg_failover_slots;
SELECT *
FROM pg_replication_slots
WHERE slot_name = 'app_slot';
```

If the default-forcing parameter is disabled, create one explicitly with `pg_create_logical_failover_slot(slotname, pluginname)`. Tencent recommends `pgoutput` as the decoding plugin.

### Lifecycle Operations

- `pg_create_logical_failover_slot()` explicitly creates a failover logical slot.
- `pg_failover_slots` lists the names of current failover slots; join or query `pg_replication_slots` for the normal slot details.
- `transform_slot_to_nonfailover()` converts an inactive failover slot to a normal slot.
- `pg_drop_replication_slot()` drops either kind of slot.

Slots are instance-wide, while the extension is database-local. Create `tencentdb_failover_slot` again in another database before calling its functions there.

### Failover Policy and Caveats

The TencentDB parameter `failover_slot_timeline_diverged_option` controls a case where standby WAL reception lags behind logical consumption during HA. Its default `error` pauses logical replication and reports an error for operator action. `rewind` resumes from the switchover point and can therefore imply a continuity/data-loss tradeoff; select it only with an application-level reconciliation plan.

Monitor retained WAL, slot activity, consumer LSNs, standby replay, and disk usage before and after failover. Test planned and unplanned switchovers with the real subscriber, confirm duplicate/gap handling, and remove abandoned slots promptly. Physical replication slots are not supported by this feature, and behavior, availability, permissions, backups, and upgrades remain subject to the TencentDB service and region.
