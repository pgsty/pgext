## Usage

Sources:

- [pg_uuid_v8 1.1.0 on PGXN](https://pgxn.org/dist/pg_uuid_v8/1.1.0/)
- [pg_uuid_v8 1.1.0 README](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/README.md)
- [pg_uuid_v8 1.1.0 control file](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8.control)
- [pg_uuid_v8 1.0 base SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0.sql)
- [pg_uuid_v8 1.0 to 1.1 upgrade SQL](https://api.pgxn.org/src/pg_uuid_v8/pg_uuid_v8-1.1.0/pg_uuid_v8--1.0--1.1.sql)
- [Pigsty pg_uuid_v8 package matrix](https://pgext.cloud/ext/pg_uuid_v8)

`pg_uuid_v8` 1.1.0 generates UUID values with UUID-v4 version and variant bits while embedding an obfuscated creation time in the random payload. Its `uuid_v8_*` convenience functions mirror the lower-level `uuid_stego_*` API. Use it when hidden time extraction and time-range indexing are useful, but do not treat the embedded value as an authentication token or a substitute for a separate trusted creation timestamp.

### Generate Values

```sql
CREATE EXTENSION pg_uuid_v8;

SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_set_encryption_mode('AES128');

CREATE TABLE events (
  id uuid PRIMARY KEY DEFAULT uuid_v8_generate(),
  data jsonb,
  created_at timestamptz NOT NULL DEFAULT now()
);

INSERT INTO events(data) VALUES ('{"type":"login"}');
```

The upstream implementation defaults to a published built-in seed and `XOR` mode. Set a deployment-specific secret before generating values. `AES128` and `AES256` are also available, but the same seed and mode must be selected when extracting a value.

### Extract and Index the Hidden Time

```sql
SELECT
  uuid_v8_extract_timestamp(id) AS epoch_microseconds,
  stego_time_to_timestamp(uuid_v8_extract_timestamp(id)) AS created_time
FROM events;

CREATE INDEX events_uuid_time_idx
ON events USING btree (uuid_v8_extract_timestamp(id));

SELECT *
FROM events
WHERE uuid_v8_extract_timestamp(id)
      BETWEEN timestamp_to_stego_time('2026-01-01'::timestamptz)
          AND timestamp_to_stego_time(now())
ORDER BY uuid_v8_extract_timestamp(id);
```

`uuid_v8_extract_timestamp(uuid)` returns a microsecond-scaled `bigint` so it remains compatible with `timestamp_to_stego_time()` and `stego_time_to_timestamp()`. In version 1.1 the internal 48-bit field stores milliseconds, so the returned value has millisecond resolution and its last three decimal digits are zero.

`uuid_stego_in_range()` offers a boolean timestamp-range helper. A functional B-tree index on the extraction function is the explicit and predictable path for indexed time predicates.

### Compare Hidden Times

`uuid_v8_compare(uuid, uuid)` and `uuid_stego_compare(uuid, uuid)` return ordering by extracted hidden time. The extension also defines `<`, `<=`, `>`, and `>=` operators for UUID arguments.

Pigsty packages install these added operators in `public` and qualify their commutator and negator references for PostgreSQL 17 and 18 compatibility. PostgreSQL already has built-in UUID ordering operators, so use the comparison functions or a schema-qualified `OPERATOR(public.<)` expression when hidden-time semantics must be unambiguous.

### Seed and Mode Controls

```sql
SELECT uuid_v8_set_seed('replace-with-a-unique-secret');
SELECT uuid_v8_get_seed();

SELECT uuid_v8_set_encryption_mode('XOR');
SELECT uuid_v8_set_encryption_mode('AES128');
SELECT uuid_v8_set_encryption_mode('AES256');
SELECT uuid_v8_get_encryption_mode();

ALTER SYSTEM SET uuid_v8.encryption_mode = 'AES128';
SELECT pg_reload_conf();
```

The seed is exposed as `uuid_v8.stego_seed` and the mode as `uuid_v8.encryption_mode`. Setter functions change the current session; configuration settings can establish defaults for later sessions. `uuid_v8_get_seed()` returns the active seed, so restrict database access accordingly and never log its result.

### Upgrade and Compatibility Boundaries

```sql
ALTER EXTENSION pg_uuid_v8 UPDATE TO '1.1';
```

Version 1.1 changes timestamp storage from microseconds to milliseconds. The old 48-bit microsecond field rolled over about every 8.9 years and could not reliably recover current absolute dates; the 48-bit millisecond field lasts about 8,925 years. Relative ordering of pre-1.1 values was unaffected, but absolute time extraction and range predicates for those existing values remain unreliable after the upgrade because their encoded representation is not rewritten.

The PGXN metadata targets PostgreSQL 12 or later; current Pigsty packages cover PostgreSQL 14–18. Pigsty packages pin the extension to `public` and make it non-relocatable so the added operators resolve consistently. Keep an ordinary `created_at` column when provenance, auditability, sub-millisecond precision, or migrations across seeds and modes matter.
