## Usage

Sources:

- [Official extension README](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/README.md)
- [Official extension source](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/src/lib.rs)
- [Official changelog](https://github.com/Creatif/creatif-backend/blob/7b0466a2b0f505c824d98c2276107531ff1af2c3/pgx_ulid/CHANGELOG.md)

`ulid` provides a binary PostgreSQL ULID type, ordinary and monotonic generators, ordering and hashing, plus implicit casts to UUID and timestamps. Use it for sortable 128-bit identifiers rendered as 26-character Crockford Base32 strings.

### Core Workflow

Ordinary generation needs no preload setting.

```sql
CREATE EXTENSION ulid;

CREATE TABLE users (
  id ulid PRIMARY KEY DEFAULT gen_ulid(),
  name text NOT NULL,
  created_at timestamp
    GENERATED ALWAYS AS (id::timestamp) STORED
);

INSERT INTO users(name) VALUES ('Alice') RETURNING id, created_at;

SELECT *
FROM users
WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV'::ulid;
```

To guarantee monotonic generation within the shared generator state, preload the library and restart PostgreSQL before calling `gen_monotonic_ulid`.

```conf
shared_preload_libraries = 'ulid'
```

```sql
SELECT gen_monotonic_ulid();
```

### Main Objects

- `ulid` is a 128-bit type with text input/output, comparison, hashing, and ordering support.
- `gen_ulid` creates a new time-based identifier with a random component.
- `gen_monotonic_ulid` increments the prior shared value when multiple IDs are generated in the same millisecond.
- `ulid_from_uuid` and `ulid_to_uuid` back implicit casts between `uuid` and `ulid`.
- `ulid_to_timestamp` extracts the millisecond timestamp; `timestamp_to_ulid` produces the lower bound whose random portion is zero.

```sql
SELECT *
FROM users
WHERE id BETWEEN '2026-07-01'::timestamp::ulid
             AND '2026-07-02'::timestamp::ulid;
```

### Operational Notes

Only the monotonic generator depends on shared memory and takes an exclusive LWLock; the rest of the extension works without preloading. Monotonic overflow raises an error. Version `0.1.3` added timestamp-to-ULID casting, and the reviewed manifest exposes PostgreSQL `14`, `15`, and `16` build features. Test implicit-cast behavior in mixed `uuid`, `ulid`, and timestamp expressions so convenience casts do not hide unintended conversions.
