## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/README.md)
- [Version 0.0.1 SQL objects](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake--0.0.1.sql)
- [C implementation](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake.c)

`pgflake` generates sortable 64-bit Snowflake-style IDs: 41 bits of millisecond time relative to a custom epoch, 10 bits of instance ID, and 12 bits of per-millisecond sequence. Configure `pgflake.instance_id` and `pgflake.start_epoch` consistently before generating IDs.

```sql
CREATE EXTENSION pgflake;
CREATE TABLE event (
  event_id bigint PRIMARY KEY DEFAULT pgflake_generate(),
  payload jsonb NOT NULL
);
SELECT pgflake_extract_time(event_id),
       pgflake_extract_instance(event_id),
       pgflake_extract_sequence(event_id)
FROM event;
```

Every concurrently active generator must have a unique `pgflake.instance_id`; reuse can create collisions. Every node must also share the same `pgflake.start_epoch`, and the 41-bit time field lasts about 69 years from that epoch.

This abandoned 0.0.1 implementation was documented for PostgreSQL 11.4 and 12. Keep clocks synchronized with NTP and test backward-clock handling, sequence exhaustion above 4096 IDs per millisecond, failover identity assignment, backend concurrency, epoch overflow, backup/restore, and upgrades. A database uniqueness constraint remains necessary even when the generator is correctly configured.
