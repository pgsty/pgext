## Usage

Sources:

- [pg_when 0.1.10 on PGXN](https://pgxn.org/dist/pg_when/0.1.10/)
- [pg_when 0.1.10 README](https://github.com/frectonz/pg-when/blob/0.1.10/README.md)
- [pg_when 0.1.10 Cargo manifest](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/Cargo.toml)
- [pg_when 0.1.10 control file](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/pg_when.control)
- [pg_when 0.1.10 exported functions](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_is.rs)
- [pg_when 0.1.10 relative-date implementation](https://api.pgxn.org/src/pg_when/pg_when-0.1.10/src/when_relative_date.rs)

`pg_when` 0.1.10 parses a constrained natural-language date and time expression and returns either a PostgreSQL `timestamptz` value or a Unix epoch value at a selected precision.

```sql
CREATE EXTENSION pg_when;

SELECT when_is('next friday at 8:00 pm in America/New_York');
SELECT seconds_at('5 days ago at this hour in Asia/Tokyo');
SELECT millis_at('in 2 months at midnight in UTC-8');
SELECT micros_at('December 31, 2026 at evening');
SELECT nanos_at('last monday at 22:30');
```

### Query Shape

A query can contain a date, a time, and a time zone, joined by `at` and `in`:

```sql
SELECT when_is('<date> at <time> in <timezone>');
SELECT when_is('<date>');
SELECT when_is('<date> in <timezone>');
SELECT when_is('<time>');
SELECT when_is('<time> in <timezone>');
SELECT when_is('<date> at <time>');
```

If the time zone is omitted, the parser uses UTC. Supported inputs include relative dates such as `tomorrow`, `last month`, and `5 days ago`; exact dates in common numeric and month-name forms; relative times such as `noon`, `midnight`, and `next hour`; clock times; IANA time-zone names; and UTC offsets.

### Function Index

- `when_is(text)` returns `timestamptz`.
- `seconds_at(text)` returns Unix epoch seconds.
- `millis_at(text)` returns Unix epoch milliseconds.
- `micros_at(text)` returns Unix epoch microseconds.
- `nanos_at(text)` returns Unix epoch nanoseconds.

### Compatibility and Boundaries

- The parser implements the documented grammar; it is not a general-purpose natural-language interpreter.
- Upstream 0.1.10 declares PostgreSQL 13–18 features and pins pgrx 0.18.1. Pigsty packages cover PostgreSQL 14–18 and apply a locked pgrx 0.19.1 compatibility update.
- `pg_when` is not relocatable and its control file requires a superuser for `CREATE EXTENSION`.
- Invalid text raises an error. All five functions are `STRICT`, so a null input returns null; `nanos_at(text)` also errors when the epoch nanoseconds cannot fit in `bigint`.
- The 0.1.10 SQL functions are declared `IMMUTABLE`, but relative expressions such as `now`, `tomorrow`, and `5 days ago` read the wall clock. Do not use relative-input calls in expression indexes or generated columns, and do not rely on them being reevaluated in cached plans; only fully specified date, time, and time-zone inputs are time-independent.
