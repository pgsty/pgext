## Usage

Sources:

- [README](https://github.com/bunko666-cloud/pg_nanotime/blob/9e1371a0645a381d915e3c485800392d9b78a49d/README.md)
- [Control](https://github.com/bunko666-cloud/pg_nanotime/blob/9e1371a0645a381d915e3c485800392d9b78a49d/pg_nanotime.control)
- [SQL objects](https://github.com/bunko666-cloud/pg_nanotime/blob/9e1371a0645a381d915e3c485800392d9b78a49d/sql/pg_nanotime--0.3.0.sql)
- [Technical reference](https://github.com/bunko666-cloud/pg_nanotime/blob/9e1371a0645a381d915e3c485800392d9b78a49d/docs/TECHNICAL.md)
- [Changelog](https://github.com/bunko666-cloud/pg_nanotime/blob/9e1371a0645a381d915e3c485800392d9b78a49d/CHANGELOG.md)

`pg_nanotime` 0.3.0 provides nanosecond timestamp and interval types on PostgreSQL 17–18 for event timestamps, latency measurements and externally supplied nanosecond data. Nanosecond storage precision does not imply nanosecond clock accuracy.

### Core Workflow

```sql
CREATE EXTENSION pg_nanotime;
CREATE TABLE events (at nanotimestamptz DEFAULT clock_ns(), body text);
INSERT INTO events (body) VALUES ('received');
SELECT at, at::timestamptz FROM events;
```

| Type | Storage and meaning |
| --- | --- |
| `nanotimestamptz` | 16 bytes; an instant displayed in the session time zone, with PostgreSQL's full timestamp range |
| `nanotimestamp` | 16 bytes; a timestamp without a time zone |
| `nanotimestamptz8` | 8 bytes; signed Unix nanoseconds with a narrower 1677–2262 range |
| `nanointerval` | 16 bytes; a fixed elapsed duration, without calendar months |

Conversion to built-in timestamps floors to microseconds, including before the Unix epoch. Casting an instant to bigint yields Unix nanoseconds and fails if the result exceeds bigint range. Do not use double precision to transport large integer nanosecond values.

### Durations, Aggregates and Series

Subtracting timestamps produces a nanosecond duration. Version 0.2.0 adds interval aggregates and series generation:

```sql
WITH samples(ns) AS (VALUES (1250::bigint), (2751::bigint))
SELECT sum(nanointerval_from_ns(ns)) AS total,
       avg(nanointerval_from_ns(ns)) AS mean,
       abs(nanointerval_from_ns(-1250::bigint)) AS magnitude
FROM samples;

SELECT *
FROM generate_series(
  '2026-10-24 00:00:00+00'::nanotimestamptz,
  '2026-10-24 00:00:00.000000003+00'::nanotimestamptz,
  nanointerval_from_ns(1::bigint)
);
```

`sum` and `avg` accumulate exactly, ignore NULL values and return NULL for no input; overflow raises an error. The average rounds to the nearest nanosecond, with ties away from zero. Series include a reachable stop value, support negative steps and reject a zero step.

Use `nanointerval_from_ns(bigint)` for exact integer input and `nanointerval_from_ns(numeric)` for larger values; fractional input rounds to nanoseconds. The double-precision overload becomes inexact beyond 2^53 nanoseconds.

Other useful objects include `time_bucket_ns` for fixed-duration buckets, `extract_ns` for precise fields, `date_trunc` and `timezone`. Explicitly type duration literals: an untyped string in timestamp arithmetic may resolve to PostgreSQL's calendar-aware interval. Timestamp types support B-tree, hash and BRIN minmax indexes; the interval type supports B-tree and hash.

### Clocks and Boundaries

`clock_ns()` reads the wall clock on each call. `now_ns()` caches an extension clock reading for the transaction and `statement_ns()` for the statement; these readings are not the same sampling instant as PostgreSQL's built-in timestamps. Ordinary clocks can move backwards with the system clock.

`clock_ns_monotonic()` produces strictly increasing values within one backend. `clock_ns_global()` does so across backends in one PostgreSQL instance and requires this configuration followed by a server restart:

```ini
shared_preload_libraries = 'pg_nanotime'
```

Preserve other libraries already in the setting. Ordinary types and clocks do not require preloading. Global monotonic state is in shared memory and does not survive restart; it is not a durable sequence or distributed ordering mechanism. Install as a superuser. The extension is relocatable and is not marked trusted.

`nanotime_clock_status()` reports clock and monotonic-state diagnostics. `nanotime.clock_source` defaults to `realtime`; the optional `tai` source depends on the operating system's configured TAI offset. TAI values still use UTC-style output, so do not mix these sources in the same column.

### Upgrade and Migration

Upstream marks 0.3.0 as a PGXN testing release. After replacing library and SQL files, update each database:

```sql
ALTER EXTENSION pg_nanotime UPDATE TO '0.3.0';
```

The column-conversion helper `nanotime_migrate_column` defaults to the 16-byte type, even for a bigint source, and changes the column with ALTER TABLE, rebuilding its indexes. Review the conversion and locking cost before running it. TimescaleDB integration is optional; its integer-time and continuous-aggregate requirements are separate from merely storing nanosecond timestamps.

### 0.3.0 Range Types and Timescale Helper

Version 0.3.0 adds `nanotstzrange` and `nanotsrange` with multirange, GiST and exclusion-constraint support. It also adds `nanotime_timescale_setup()` for the optional TimescaleDB integration and extends the technical helpers for range construction and extraction. Existing timestamp and interval objects keep their prior semantics. Replace matching library and SQL files before running `ALTER EXTENSION pg_nanotime UPDATE TO '0.3.0'`; review any range indexes and the optional TimescaleDB setup separately.
