## Usage

Sources:

- [Official UPID README](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/README.md)
- [PostgreSQL extension implementation](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/src/lib.rs)
- [PostgreSQL extension manifest](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_pg/Cargo.toml)
- [Core identifier implementation](https://github.com/carderne/upid/blob/d820a74227adb803d378d89da9eacba78c1a5e09/upid_rs/src/lib.rs)

`upid_pg` adds the `upid` type: a sortable 128-bit identifier containing a short semantic prefix, a timestamp, random bits, and a format version. It is useful when UUID-sized keys should remain globally generated while still revealing a compact object kind such as `user`.

### Generate and store UPIDs

```sql
CREATE EXTENSION upid_pg;

CREATE TABLE users (
    id   upid NOT NULL DEFAULT gen_upid('user') PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO users (name) VALUES ('Ada')
RETURNING id,
          upid_to_prefix(id),
          upid_to_timestamp(id);
```

The binary layout places a 40-bit Unix timestamp first for approximate chronological ordering, followed by 64 random bits and 24 bits for the prefix and version. Timestamp precision is `256` milliseconds, so generation order within the same time bucket is random rather than monotonic.

### Conversion functions

- `gen_upid(prefix)` generates a new value using the current UTC time.
- `upid_to_prefix(upid)` and `upid_to_timestamp(upid)` extract the embedded prefix and approximate timestamp.
- `upid_to_bytea(upid)` exposes the 16-byte representation.
- `upid_to_uuid(upid)` and `upid_from_uuid(uuid)` preserve the same 128 bits.

The extension also defines implicit casts between `uuid` and `upid`, and from `upid` to `bytea` or `timestamp`:

```sql
SELECT upid_to_uuid(gen_upid('user'));
SELECT upid_from_uuid(gen_random_uuid());
```

Converting an arbitrary UUID to `upid` is reversible, but the decoded prefix and timestamp have no semantic meaning.

### Prefix and maturity boundaries

Pass exactly four lowercase Latin letters when generating identifiers. The implementation silently normalizes other input: a short prefix is right-padded with `z`, a long prefix is truncated to four characters, and characters outside its alphabet are replaced with `z`. Validate prefixes in application code if silent normalization is undesirable.

The embedded timestamp is metadata, not an authentication or trustworthy creation-time proof. The 64 random bits provide probabilistic uniqueness, not a database-enforced collision guarantee. Catalog version `0.0.0` denotes an alpha-stage API; the reviewed pgrx manifest has build features for PostgreSQL 15 through 18, but actual package availability and upgrade compatibility must be checked for the target server. Pin the revision and test type I/O, indexes, casts, dump/restore, and upgrades before adoption.
