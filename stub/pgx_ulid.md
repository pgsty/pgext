
## Usage

Sources: [README](https://github.com/pksunkara/pgx_ulid/blob/master/README.md), [releases](https://github.com/pksunkara/pgx_ulid/releases)

`pgx_ulid` provides a native `ulid` type, generators, and casts to and from `timestamp` and `uuid`. The README documents binary storage and monotonic ULID support.

### Enable the extension

```sql
CREATE EXTENSION ulid;
-- or CREATE EXTENSION pgx_ulid; if installed manually under that name
```

### Generate ULIDs

```sql
SELECT gen_ulid();
SELECT gen_monotonic_ulid();
```

`gen_monotonic_ulid()` needs:

```conf
shared_preload_libraries = 'pgx_ulid'
```

The README explicitly says this preload requirement only affects `gen_monotonic_ulid()`; the rest of the extension works without it.

### Use `ulid` as a primary key

```sql
CREATE TABLE users (
  id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
  name text NOT NULL
);

SELECT * FROM users
WHERE id = '01ARZ3NDEKTSV4RRFFQ69G5FAV';
```

### Casts and range queries

```sql
ALTER TABLE users
ADD COLUMN created_at timestamp GENERATED ALWAYS AS (id::timestamp) STORED;

SELECT * FROM users
WHERE id BETWEEN '2023-09-15'::timestamp::ulid
            AND '2023-09-16'::timestamp::ulid;
```

The README also documents casts between `ulid` and `uuid`.

### Caveats

- Monotonic ULIDs use shared memory plus an LWLock to keep the last generated value.
- The README notes that monotonic generation can theoretically overflow and raise an error, although it treats this as negligible in practice.
- Release `v0.2.3` is current upstream as of 2026-04-19, but upstream did not publish separate user-facing release notes for it.
