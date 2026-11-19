## Usage

Sources:

- [3.0 changes](https://github.com/s0l0ist/ferroid/blob/3.0.0/CHANGELOG.md)
- [PostgreSQL extension README](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/README.md)
- [Extension control file](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/pg_ferroid.control)
- [Type, cast, and function implementation](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/src/lib.rs)
- [Package metadata](https://github.com/s0l0ist/ferroid/blob/3.0.0/crates/pg-ferroid/Cargo.toml)

`pg_ferroid` 3.0.0 is a pgrx extension that adds a native 16-byte `ulid` type. ULIDs encode a millisecond timestamp plus an 80-bit tail, have a 26-character Base32 representation, and support comparisons, B-tree and hash indexing, validation, and explicit conversions.

### Generate and store ULIDs

```sql
CREATE EXTENSION pg_ferroid;

CREATE TABLE users (
  id ulid PRIMARY KEY DEFAULT gen_ulid_mono(),
  name text NOT NULL
);

INSERT INTO users (name) VALUES ('alice'), ('bob');

SELECT id, id::timestamptz AS created_at, name
FROM users
ORDER BY id;

SELECT ulid_is_valid('01JEPY8K5V3XQZW6M9N7P8Q2RT');
```

`gen_ulid()` generates a time-ordered value with a random tail. `gen_ulid_mono()` increments the tail for calls in the same millisecond, but its monotonic state is per backend or thread, not cluster-wide. Do not interpret it as a global sequence.

### Ordering and information exposure

Explicit casts connect `ulid` with `text`, `bytea` and `timestamptz`. Casting a timestamp to `ulid` sets the random portion to zero, which is useful as a lower range bound but does not generate a normal application identifier.

The embedded timestamp exposes approximate creation time at millisecond precision. Avoid ULIDs where identifier disclosure must not reveal age or ordering. The type uses a native extension representation, so test binary upgrades, logical decoding, dump and restore, drivers, and index rebuilds across PostgreSQL and extension versions. Version 3.0.0 is superuser-only and not trusted according to its control file.

### Version 3.0 boundary

Requires PostgreSQL 15–18; PostgreSQL 13/14 support was removed. Direct casts between `ulid` and `timestamp without time zone` were removed. Specify the zone and convert through `timestamptz`:

```sql
SELECT (timestamp '2025-01-01 12:00:00' AT TIME ZONE 'UTC')::ulid;
```

Timestamp inputs must be finite, at or after the Unix epoch, and fit the 48-bit ULID time range. A backward clock adjustment can delay monotonic generation; cancellation and statement timeouts remain effective. The matching Rust/pgrx build uses pgrx 0.20.0.
