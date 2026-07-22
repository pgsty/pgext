## Usage

Sources:

- [Official README](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/README.md)
- [Official control file](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/pg_snowflake_id.control)
- [Official GUC definitions](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/config.rs)
- [Official SQL function](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/generate_snowflake_id.rs)
- [Official generator implementation](https://github.com/k9982874/pg_snowflake_id/blob/38c33b7fb1e8587be0a24a465f6a679335c22099/src/snowflake.rs)

`pg_snowflake_id` exposes `generate_snowflake_id()` for producing Snowflake-shaped 64-bit integers. The reviewed `0.0.1` implementation keeps one generator per PostgreSQL backend, so it does not provide database-wide uniqueness when concurrent backends use the same node identifiers.

### Core Workflow

Set a data-center and worker identity before the first generated value in every backend, then call the function.

```sql
CREATE EXTENSION pg_snowflake_id;

SET pg_snowflake_id.data_center_id = 1;
SET pg_snowflake_id.worker_id = 7;
SET pg_snowflake_id.epoch = '2021-01-01 00:00:00 UTC';

SELECT generate_snowflake_id();
```

The bit layout is a 41-bit timestamp, 5-bit data-center ID, 5-bit worker ID, and 12-bit sequence. A generator can emit 4096 values per millisecond, then waits for the next millisecond.

### Configuration

- `pg_snowflake_id.data_center_id` accepts `0` through `31` and defaults to `1`.
- `pg_snowflake_id.worker_id` accepts `0` through `31` and defaults to `1`.
- `pg_snowflake_id.epoch` selects the timestamp epoch and defaults to `2021-01-01 00:00:00 UTC`.

All three settings are user-settable. The backend initializes its generator on the first call and caches it, so later setting changes in that same session do not alter the active generator.

### Uniqueness Boundary

PostgreSQL backends do not share the Rust mutex or sequence counter. Two concurrent sessions with the same data-center and worker IDs can therefore generate the same value in the same millisecond. Session-level settings and connection pooling also make it difficult to assign a durable unique identity to every backend.

Do not use `generate_snowflake_id()` as a primary-key default under concurrency unless a reviewed downstream patch or an external allocation scheme establishes uniqueness across all backends and nodes. Also handle the explicit clock-regression error and verify the epoch's usable time range.
