

## Usage

> [pg_duration: ISO 8601 duration type for PostgreSQL](https://github.com/jkosh44/pg_duration)

The `pg_duration` extension provides a `duration` type for storing elapsed time as microseconds in 8 bytes, simpler and more consistently comparable than the built-in `interval` type.

```sql
CREATE EXTENSION pg_duration;
```

### Data Type

The `duration` type represents absolute elapsed time without calendar components (no months or days). Valid input accepts any PostgreSQL interval syntax that does not exceed hourly units.

```sql
SELECT '01:30:00'::duration;
SELECT '2 hours 30 minutes'::duration;
```

### Operators

- **Arithmetic**: `+`, `-` between durations; `*`, `/` by `float8`; unary `-`
- **Comparison**: `<`, `<=`, `>`, `>=`, `=`, `<>`

### Functions

```sql
-- Construct from components
SELECT make_duration(hours => 2, mins => 30, secs => 15.5);

-- Check for infinity
SELECT isfinite('01:00:00'::duration);

-- Truncate to precision
SELECT date_trunc('hour', '02:45:30'::duration);

-- Extract subfield
SELECT date_part('minute', '02:45:30'::duration);
SELECT extract_duration('hour', '02:45:30'::duration);
```

### Type Conversions

Durations cast implicitly to `interval`. Converting `interval` to `duration` requires explicit casting.

### Aggregates and Indexing

Standard aggregates (`avg`, `count`, `max`, `min`, `sum`) and both B-tree and hash indexes are supported.
