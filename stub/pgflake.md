## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/README.md)
- [Version 0.0.1 SQL objects](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake--0.0.1.sql)
- [C implementation](https://github.com/mausimag/pgflake/blob/f07f73395258fefdae982f7b24393d7f02625e77/pgflake.c)

`pgflake` 0.0.1 intends to generate sortable 64-bit Snowflake-style IDs from 41 time bits, 10 instance bits, and 12 sequence bits. The reviewed implementation does not provide reliable time ordering or cross-session uniqueness, so treat it as an abandoned prototype rather than a production ID generator.

### Demonstration

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

The extractors return the encoded time, instance, and sequence components; they do not repair or validate a generated ID. Keep the `PRIMARY KEY` or another uniqueness constraint when evaluating the prototype.

### Integrity defects

The intended configuration parameters are `pgflake.instance_id` and `pgflake.start_epoch`, but the module registers them through stack-local variables and copies their values only during library initialization. A later `SET` does not update the generation state and may write through invalid storage. Do not rely on runtime GUC changes without correcting the C code.

The implementation converts `clock_gettime` by dividing `tv_nsec` by 1000 rather than 1000000. The encoded value is therefore not milliseconds, can move backward at each second boundary, and overlaps values from other seconds. Its `last_time`, `sequence`, and spinlock are also backend-local, so two sessions using the same instance ID can emit an identical value. Assigning one instance ID per PostgreSQL server does not distinguish those sessions.

Sequence rollover contains a second time-arithmetic error, and backward clocks are not rejected. NTP does not fix these implementation defects. The repository documents PostgreSQL 11.4 and 12 only; do not use this version to generate persistent identifiers unless the code is repaired, concurrency-tested across backends, and reviewed again.
