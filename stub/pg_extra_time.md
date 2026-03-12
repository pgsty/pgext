


## Usage

> [pg_extra_time: Extra date/time functions and operators for PostgreSQL](https://github.com/bigsmoke/pg_extra_time)

### Convert to Seconds (float)

```sql
SELECT to_float('1970-01-01 00:00:00+0'::timestamptz);  -- 0.0
SELECT to_float('1 day 1 sec'::interval);                -- 86401.0
SELECT to_float('[2024-06-06 05:58:00,2024-06-06 06:00:10]'::tstzrange);  -- 130.0
```

Cast syntax also works:

```sql
SELECT '1970-01-01 01:03:01+00'::timestamptz::float;    -- 3181.00
SELECT '1 day 1 sec 200 ms'::interval::float;            -- 86401.2
```

### Convert to Days

```sql
SELECT days('[2024-06-06,2024-06-08 06:00]'::tstzrange);  -- 3.25 (fractional days)
SELECT whole_days('[2024-06-06,2024-06-08 18:00]'::tstzrange);  -- 2 (whole days only)
SELECT days('10 days 12 hours'::interval);                -- 10.5
SELECT whole_days('10 days 20 hours'::interval);          -- 10
```

### Extract Interval from Range

```sql
SELECT to_interval('[2024-01-01,2024-01-05]'::tstzrange);  -- 4 days
```

### Each Functions (generate series from ranges)

```sql
SELECT * FROM each_subperiod('[2024-01-01,2024-01-05]'::tstzrange, '1 day'::interval);
```

### Operators

```sql
-- Range duration as float (seconds)
SELECT '[epoch,1970-01-01T01:03:01+00]'::tstzrange::float;
-- Interval as float (seconds)
SELECT '10 seconds 100 milliseconds'::interval::float;  -- 10.100
```
