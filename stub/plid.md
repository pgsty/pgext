## Usage

Sources:

- [Official README](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/README.md)
- [Rust implementation](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/src/lib.rs)
- [Cargo compatibility features](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/Cargo.toml)
- [Extension control file](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/plid.control)

`plid` is a 128-bit, ULID-inspired identifier type with a one-to-three-letter application prefix. It stores 16 prefix bits, a 48-bit Unix timestamp in milliseconds, and 64 random bits, and prints them as a lowercase prefix, underscore, and 23 Crockford Base32 characters.

### Preload and Create

Monotonic generation uses postmaster shared memory. Add the library to `shared_preload_libraries`, restart PostgreSQL, and then create the extension.

```conf
shared_preload_libraries = 'plid'
```

```sql
CREATE EXTENSION plid;
```

### Core Workflow

Use a stable prefix for each entity type and generate IDs in a column default. Prefixes accept only one to three ASCII letters in `a-z`.

```sql
CREATE TABLE app_user (
  id plid PRIMARY KEY DEFAULT gen_plid_monotonic('usr'),
  name text NOT NULL
);

INSERT INTO app_user(name) VALUES ('Ada') RETURNING id;
```

Extract the embedded timestamp or build an all-ones boundary ID for timestamp range predicates.

```sql
SELECT id, plid_to_timestamptz(id), plid_to_prefix(id)
FROM app_user;

SELECT *
FROM app_user
WHERE id > timestamptz_to_plid('2026-01-01 00:00:00+00', 'usr');
```

### Important Objects

- `plid`: fixed 16-byte type with comparison operators and a B-tree operator class.
- `gen_plid(text)`: generates random bits without monotonic adjustment.
- `gen_plid_monotonic(text)`: increments the random portion for IDs generated in the same millisecond.
- `plid_to_timestamptz(plid)`, `plid_to_timestamp(plid)`: extract the embedded time.
- `timestamptz_to_plid(timestamptz, text)`: creates a boundary value with all random bits set to one.
- `plid_to_prefix(plid)`: extracts the prefix.

### Ordering and Scope

Prefix bits are the most significant bits. Global ordering groups by prefix before timestamp, so chronological range queries are meaningful only when the predicate uses the same prefix. Monotonic state belongs to one PostgreSQL postmaster; it resets on restart and does not coordinate across primary/replica pairs, shards, or independent servers. It is not a global sequence guarantee.

Text input is case-insensitive and output normalizes the prefix to lowercase and the Base32 portion to uppercase. The timestamp must be on or after the Unix epoch and fit the 48-bit millisecond field. The source lists B-tree support but hash-index support as unfinished. The extension is versioned `0.0.0`, requires superuser installation, and its author says it was written as a learning exercise and has not been used by them in production. Treat those maturity signals as a deployment boundary even though the Cargo manifest has build features for PostgreSQL 13–18.
