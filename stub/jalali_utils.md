## Usage

Sources:

- [Official README](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/README.md)
- [Official extension SQL](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/jalali_utils--0.0.1.sql)
- [Official C implementation](https://github.com/teamappir/jalali_utils/blob/92d99f2051787dbb13a84b6da8f6e21b5e93d870/jalali_utils.c)

`jalali_utils` converts Gregorian PostgreSQL dates and timestamps to Solar Hijri (Jalali) text or integer date parts. The timestamp overloads always interpret the instant in `Asia/Tehran`; they do not use the session `TimeZone`.

### Core Workflow

```sql
CREATE EXTENSION jalali_utils;

SELECT format_jalali(timestamptz '2019-07-07 14:10:52.84937+04:30');
SELECT format_jalali(date '2021-03-20');
SELECT jalali_part('year', date '2021-03-20');
SELECT jalali_part('doy', timestamptz '2019-07-07 14:10:52.84937+04:30');
```

`format_jalali(timestamptz, boolean)` returns `YYYY/MM/DD HH:MM:SS` when the optional flag is true and date-only text when false. `format_jalali(date)` always returns date-only text.

### Function Index

- `format_jalali(timestamptz, boolean DEFAULT true)` formats an instant in Tehran civil time.
- `format_jalali(date)` converts a Gregorian calendar date.
- `jalali_part(text, timestamptz)` extracts a Jalali field from an instant in Tehran time.
- `jalali_part(text, date)` extracts a Jalali field from a date.

Accepted fields are `year`, `month`, `day`, `hour`, `minute`, `second`, `doy`, `dow`, `quarter`, `decade`, and `century`. `dow` uses Saturday as zero.

### Caveats

Version `0.0.1` returns formatted values as `text`, not a sortable or validated Jalali data type. Sort and compare the original PostgreSQL date or timestamp, and use formatting only at presentation boundaries. The hard-coded Tehran zone means timestamp results can cross a calendar-day boundary relative to UTC or the session zone. Verify leap-day, daylight-saving history, negative/infinite dates, and supported PostgreSQL versions for the exact binary before relying on the conversion in financial or legal workflows.
