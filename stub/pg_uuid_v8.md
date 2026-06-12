## Usage

Sources: [pg_uuid_v8 README](https://github.com/ineron/pg_uuid_v8), [SQL definitions](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8--1.0.sql), [control file](https://github.com/ineron/pg_uuid_v8/blob/main/pg_uuid_v8.control).

`pg_uuid_v8` generates UUIDs that look like UUID v4 values while embedding encrypted microsecond timestamps for extraction, sorting, and range predicates. The SQL file exposes both `uuid_stego_*` names and `uuid_v8_*` convenience aliases.

### Generate UUIDs

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_generate();
```

The equivalent lower-level generator is:

```sql
SELECT uuid_stego_generate();
```

Use a default expression when inserting events:

```sql
CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz DEFAULT now()
);
```

### Extract And Query Hidden Timestamps

Extract the embedded timestamp as microseconds since the Unix epoch:

```sql
SELECT uuid_v8_extract_timestamp(id)
FROM events
ORDER BY uuid_v8_extract_timestamp(id)
LIMIT 10;
```

The README recommends functional indexes for time-based lookups:

```sql
CREATE INDEX events_uuid_v8_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

Helper functions convert between timestamps and the integer timestamp format:

```sql
SELECT timestamp_to_stego_time(now());
SELECT stego_time_to_timestamp(uuid_v8_extract_timestamp(id))
FROM events;
```

### Range Helpers And Operators

The SQL definition includes direct range helpers:

```sql
SELECT *
FROM events
WHERE uuid_stego_in_range(
  id,
  now() - interval '24 hours',
  now()
);
```

It also defines timestamp-aware comparison functions and operators for `uuid`:

- `uuid_stego_compare(uuid, uuid)` and `uuid_v8_compare(uuid, uuid)`.
- `uuid_stego_lt`, `uuid_stego_le`, `uuid_stego_gt`, `uuid_stego_ge`.
- Operators `<`, `<=`, `>`, and `>=` compare UUIDs by hidden timestamp.

### Seed And Encryption Mode

Set and inspect the seed:

```sql
SELECT uuid_v8_set_seed('replace-with-a-secret-seed');
SELECT uuid_v8_get_seed();
```

Available encryption modes are `XOR`, `AES128`, and `AES256`:

```sql
SELECT uuid_v8_get_encryption_mode();
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('XOR');
```

For a persistent default, the README documents the `uuid_v8.encryption_mode` GUC:

```sql
ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

### Caveats

- Keep the seed secret; it is required to interpret hidden timestamps.
- UUIDs generated with one seed and encryption mode must be decoded with the same settings.
- Functional indexes on extracted timestamps add storage and update overhead, but are the intended path for efficient time-range predicates.
- Local Pigsty metadata pins this extension to the `public` schema so the UUID comparison operator commutators resolve on PostgreSQL 17 and 18; test operators explicitly if using a different schema in a non-Pigsty build.
