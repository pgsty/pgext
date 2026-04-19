
## Usage

Sources: [README](https://github.com/EnterpriseDB/pg_failover_slots/blob/master/README.md), [releases](https://github.com/EnterpriseDB/pg_failover_slots/releases)

`pg_failover_slots` keeps logical replication slots usable across failover by synchronizing slot definitions and positions from a primary to a standby.

### Enable it on both nodes

```ini
shared_preload_libraries = 'pg_failover_slots'
```

Required standby settings from the README:

```ini
hot_standby_feedback = on
primary_slot_name = 'my_physical_slot'
```

### Main configuration

```ini
pg_failover_slots.synchronize_slot_names = 'name_like:%'
pg_failover_slots.drop_extra_slots = true
pg_failover_slots.primary_dsn = 'host=primary dbname=mydb'
pg_failover_slots.standby_slot_names = 'standby_physical_slot'
pg_failover_slots.standby_slots_min_confirmed = -1
pg_failover_slots.worker_nap_time = 60000
pg_failover_slots.maintenance_db = 'postgres'
```

The README documents `synchronize_slot_names` filters by exact slot name, `LIKE` pattern, or plugin name.

### Check standby readiness before failover

```sql
SELECT slot_name, active
FROM pg_replication_slots
WHERE slot_type = 'logical';
```

On the standby, logical slots are ready only when they exist and show `active = false`. The README says `active = true` means a slot is still being initialized.

### Notes

- PostgreSQL 11+ is required upstream.
- `v1.2.1` is a bug-fix release; no new user-facing SQL or GUC surface was added there.
- `v1.2.0` added PostgreSQL 18 support and clarified `drop_extra_slots` behavior.
