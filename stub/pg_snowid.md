## Usage

Sources:

- [Official README](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/README.md)
- [Extension control file](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/pg_snowid.control)
- [Function and shared-memory implementation](https://github.com/rixlhq/snowid-postgres/blob/9be325a349805a3c10ebcb37cd00f1e3a1a150d2/src/lib.rs)

pg_snowid generates time-ordered 64-bit Snowflake-like identifiers and compact Base62 representations in PostgreSQL. It keeps per-key generators in shared memory and needs a distinct node ID for each PostgreSQL instance in a distributed deployment.

### Core Workflow

Add `pg_snowid` to `shared_preload_libraries`, restart PostgreSQL, create the extension, and set the node before generating any ID:

```ini
shared_preload_libraries = 'pg_snowid'
```

```sql
CREATE EXTENSION pg_snowid;

SELECT snowid_set_node(5);

CREATE TABLE event (
  id bigint PRIMARY KEY DEFAULT snowid_generate(1),
  payload jsonb NOT NULL
);

INSERT INTO event(payload) VALUES ('{"kind":"created"}');
SELECT id, snowid_get_timestamp(id) FROM event;
```

The generator key passed to `snowid_generate` is an integer-like OID value. Keep a stable, documented key per ID domain; it selects shared generator state but is not encoded in the resulting ID.

### Important Objects

- `snowid_set_node` sets node ID 0 through 1023 before any generator is created; the default is 1.
- `snowid_get_node` returns the active node ID.
- `snowid_generate` and `snowid_generate_base62` generate one numeric or Base62 ID.
- `snowid_try_generate` and `snowid_try_generate_base62` return null instead of moving logical time ahead.
- `snowid_generate_batch` and `snowid_try_generate_batch` reserve up to 100,000 IDs.
- `snowid_get_timestamp` and `snowid_get_timestamp_base62` extract the millisecond timestamp.
- `snowid_stats` reports node and generator state.

### Operational Notes

Under sustained load, normal generation advances a logical timestamp rather than waiting, so an embedded timestamp can be ahead of wall-clock time. The try variants preserve wall-clock behavior but may return null or a short batch. Node ID is not derived automatically; reusing it across instances risks collisions, and changing it after a generator exists requires a PostgreSQL restart. Separate generator keys are independent and the key is absent from the bit layout, so do not assume global uniqueness across keys without an application-level test and constraint strategy.
