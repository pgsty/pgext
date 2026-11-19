## Usage

Sources:

- [Official source](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/README.md)
- [Official source](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/pg_governor.control)
- [Official source](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/sql/pg_governor--0.1.sql)
- [Official source](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/LICENSE)

`pg_governor` 0.1 limits concurrent top-level queries and connections for directly assigned role groups. It supports PostgreSQL 15–17. Admission priority controls queue order, not CPU or I/O allocation.

### Configure a workload group

Preload the library and restart. Create the extension in the database selected by `pg_governor.database`; the example assumes role `alice` already exists.

```ini
shared_preload_libraries = 'pg_governor'
pg_governor.database = 'postgres'
```

```sql
CREATE EXTENSION pg_governor;
SELECT governor.create_group('reporting', 2, 5000, 65536, 20, 3);
SELECT governor.assign_role('alice', 'reporting');
SELECT * FROM governor.stats;
SELECT * FROM governor.instance_status();
```

### Limits and administration

The example allows two active queries, a 5000 ms wait, 65536 KB of `work_mem`, 20 connections and priority 3. Higher priority wins; equal priority is FIFO. Zero active/connection limits mean unlimited; zero wait means no timeout. `pg_governor.max_active_total` and `pg_governor.max_connections_total` add instance ceilings for managed roles.

`governor.alter_group`, `governor.assign_role`, `governor.unassign_role`, `governor.drop_group` and `governor.reload` change configuration. Administrative functions are revoked from `PUBLIC`; statistics are publicly readable. Cumulative counters are periodically saved by the worker; a crash can lose the most recent interval.

### Operational boundaries

`governor.reload()` changes shared state nontransactionally. After rolling back a configuration change, call it again to align shared memory with the tables. Only directly assigned roles are controlled; role membership does not inherit a group. Nested queries and queries while a cursor is open are not independently counted. The shared queue holds at most 1024 waiting queries. Drop groups before removing the extension: `DROP EXTENSION` does not clear shared state. This is not a CPU, memory or I/O isolation mechanism.
