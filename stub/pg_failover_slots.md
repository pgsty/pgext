

## Usage

> [pg_failover_slots: PG Failover Slots extension](https://github.com/EnterpriseDB/pg_failover_slots)

Makes logical replication slots usable across a physical failover by synchronizing slot state from primary to standby.

### Enabling

Add to `postgresql.conf` on both primary and standby:

```ini
shared_preload_libraries = 'pg_failover_slots'
```

Required settings on standby:

```ini
hot_standby_feedback = on
primary_slot_name = 'my_physical_slot'  -- must be non-empty
```

### Configuration Options

```ini
# Which slots to synchronize (default: all logical slots)
pg_failover_slots.synchronize_slot_names = 'name_like:%'

# Synchronize specific slots
pg_failover_slots.synchronize_slot_names = 'my_slot,plugin:test_decoding'

# Drop extra slots on standby not found on primary (default: true)
pg_failover_slots.drop_extra_slots = true

# Connection string to primary (default: uses primary_conninfo)
pg_failover_slots.primary_dsn = 'host=primary dbname=mydb'

# Ensure physical standbys receive data before logical consumers
pg_failover_slots.standby_slot_names = 'standby_physical_slot'

# How many standby slots must confirm (default: -1 = all)
pg_failover_slots.standby_slots_min_confirmed = -1

# Sync interval in ms (default: 60000)
pg_failover_slots.worker_nap_time = 60000
```

### Checking Standby Readiness

Verify all logical slots are synchronized before failover:

```sql
-- On standby: all slots should show active = false
SELECT slot_name, active FROM pg_replication_slots WHERE slot_type = 'logical';

--  slot_name        | active
-- ------------------+--------
--  regression_slot1 | f        -- synchronized, ready
--  regression_slot2 | f        -- synchronized, ready
--  regression_slot3 | t        -- still syncing, NOT ready
```

When all slots show `active = false`, the standby is safe for failover.

### Key Behaviors

- Copies missing replication slots from primary to standby
- Removes extra slots on standby not found on primary
- Periodically synchronizes slot positions
- `standby_slot_names` provides a synchronous replication barrier to prevent data loss on failover
- Requires PostgreSQL 11 or higher
