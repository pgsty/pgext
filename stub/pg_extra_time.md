
## Usage

> Sources: [pg_extra_time upstream README](https://github.com/bigsmoke/pg_extra_time), [PGXN pg_extra_time](https://pgxn.org/dist/pg_extra_time/), [local metadata](../db/extension.csv).

`pg_extra_time` provides small SQL functions and casts for date/time, interval, and range calculations that are awkward with PostgreSQL core functions alone.

```sql
CREATE EXTENSION pg_extra_time;
```

### Convert to Seconds (float)

Use `to_float(...)` or explicit casts to `float`/`double precision` for timestamps, timestamp ranges, and intervals. Timestamp values are measured from the Unix epoch; ranges and intervals are measured by duration in seconds.

```sql
SELECT to_float('1970-01-01 00:00:00+0'::timestamptz);  -- 0.0
SELECT to_float('1970-01-01 00:00:00+0'::timestamp);    -- 0.0
SELECT to_float('1 day 1 sec'::interval);                -- 86401.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tstzrange);  -- 130.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tsrange);    -- 130.0
```

Cast syntax also works:

```sql
SELECT '1970-01-01 01:03:01+00'::timestamptz::float;    -- 3181.00
SELECT '1 day 1 sec 200 ms'::interval::float;            -- 86401.2
SELECT '[epoch,1970-01-01T01:03:01+00]'::tstzrange::float;  -- 3181.00
```

### Convert to Days

Use `days(...)` when fractions matter and `whole_days(...)` when an integer number of complete days is needed.

```sql
SELECT days('[2024-06-06,2024-06-09)'::daterange);       -- 2
SELECT days('[2024-06-06,2024-06-08 06:00]'::tstzrange);  -- 3.25 (fractional days)
SELECT whole_days('[2024-06-06,2024-06-08 18:00]'::tstzrange);  -- 2
SELECT days('10 days 12 hours'::interval);                -- 10.5
SELECT whole_days('10 days 20 hours'::interval);          -- 10
```

`whole_days(interval)` handles negative intervals by applying the sign after flooring the absolute day count.

### Count Date Parts

`date_part_parts(part, subpart, timestamp with time zone, timezone)` returns how many smaller date parts exist in a larger date part at a given timestamp and timezone. This helps with calculations where a day is not always 24 hours because of DST.

```sql
SELECT date_part_parts('month', 'days', '2024-02-12'::timestamptz, 'UTC');  -- 29
SELECT date_part_parts('year', 'days', '2024-08-23'::timestamptz, 'UTC');   -- 366
```

### Build And Split Ranges

Use `make_tstzrange` or `make_tsrange` to build ranges from a timestamp and interval, including negative intervals.

```sql
SELECT make_tstzrange('2024-01-05 00:00+00'::timestamptz, '-4 days'::interval);
SELECT make_tsrange('2024-01-01 00:00'::timestamp, '4 days'::interval, '[)');
```

`each_subperiod(tstzrange, interval, round_remainder integer DEFAULT 0)` splits a timestamp range into interval-sized chunks. The remainder policy is: `1` rounds up to a full chunk, `0` keeps a partial final chunk, and `-1` discards the remainder.

```sql
SELECT *
FROM each_subperiod(
  '[2023-01-01,2023-04-02)'::tstzrange,
  '1 month'::interval,
  0
);
```

### Extract And Remainder Intervals

`to_interval(tstzrange)` extracts an interval from a timestamp range using month, day, and microsecond units. `to_interval(tstzrange, interval[])` accepts explicit units in greatest-first order and rounds down by discarding the remainder.

```sql
SELECT to_interval('[2024-01-01,2024-01-05]'::tstzrange);  -- 4 days
SELECT to_interval(
  '[2024-01-01,2024-04-13 01:10]'::tstzrange,
  ARRAY['1 mon'::interval, '1 day'::interval, '1 hour'::interval]
);
```

Use `%` or `modulo(...)` when the remainder matters.

```sql
SELECT '10 seconds 100 milliseconds'::interval % '3 seconds'::interval;
SELECT '[2024-01-01,2024-01-10)'::tstzrange % '4 days'::interval;
```

### Caveats

`to_float(tstzrange)` and `to_float(tsrange)` return positive or negative infinity for unbounded ranges and `0` for empty ranges. Integer casts are intentionally not provided; use `whole_days(...)` when you need integer days. Deprecated aliases such as `extract_days(interval)` and `extract_interval(tstzrange, interval[])` remain for compatibility, but upstream recommends `whole_days(...)` and `to_interval(...)` instead.

### Reference

Common public functions:

| Function | Use |
|----------|-----|
| `current_timezone()` | Return the active `pg_timezone_names` row |
| `date_part_parts(...)` | Count smaller date parts inside larger date parts |
| `days(...)` | Fractional or integer day count, depending on input type |
| `whole_days(...)` | Whole days from intervals or timestamp ranges |
| `to_float(...)` | Seconds from timestamps, timestamp ranges, or intervals |
| `to_interval(...)` | Interval extracted from a `tstzrange` |
| `make_tsrange(...)` / `make_tstzrange(...)` | Build ranges from timestamp plus interval |
| `each_subperiod(...)` | Split a `tstzrange` into subranges |
| `modulo(...)` / `%` | Remainder after dividing intervals or ranges |
