## Usage

Sources:

- [Official README for v1.1.3](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/README.md)
- [Official extension SQL for v1.1.3](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/sql/pg_ethiopian_calendar--1.1.sql)
- [Official control file for v1.1.3](https://github.com/HuluWZ/pg-ethiopian-calendar/blob/v1.1.3/pg_ethiopian_calendar.control)

`pg_ethiopian_calendar` converts Gregorian PostgreSQL timestamps to Ethiopian calendar dates and back. The extension SQL version is `1.1`; the reviewed upstream distribution tag is `v1.1.3`. Use it when an application needs Ethiopia's 13-month civil calendar, while keeping ordinary PostgreSQL timestamps as the storage and interchange boundary.

### Core Workflow

```sql
CREATE EXTENSION pg_ethiopian_calendar;

SELECT to_ethiopian_date(TIMESTAMP '2026-01-01 00:00:00');
SELECT from_ethiopian_date('2018-04-23');

CREATE TABLE events (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  occurred_at timestamp NOT NULL,
  ethiopian_at timestamp
    GENERATED ALWAYS AS (to_ethiopian_timestamp(occurred_at)) STORED
);
```

The text conversion uses Ethiopian `YYYY-MM-DD` notation. Validate application input and decide explicitly whether the text-returning or timestamp-returning form belongs in each interface.

### Functions

- `to_ethiopian_date(timestamp)` returns Ethiopian date text; `from_ethiopian_date(text)` parses that text back to a Gregorian timestamp.
- `to_ethiopian_datetime(timestamp)` returns a `timestamptz`, while `to_ethiopian_timestamp(timestamp)` preserves a timestamp-style result.
- `current_ethiopian_date()` reports the current Ethiopian date.
- `pg_ethiopian_to_date`, `pg_ethiopian_from_date`, `pg_ethiopian_to_datetime`, and `pg_ethiopian_to_timestamp` are aliases for the corresponding conversion functions.

The conversion functions are declared `IMMUTABLE` and `STRICT`; `current_ethiopian_date()` is `STABLE`. That makes deterministic conversions usable in generated expressions and indexes, subject to normal PostgreSQL type and time-zone semantics.

### Operational Notes

The input of `to_ethiopian_datetime(timestamp)` is `timestamp without time zone` but its result is `timestamptz`; test the session time zone used by the application before relying on that form. Prefer the timestamp-preserving function when implicit time-zone interpretation is undesirable.

The repository also ships a separate `sql/ethiopian_calendar_plpgsql.sql` implementation with a different API. It is not installed by this extension's control/SQL path, so examples for its no-argument helpers must not be assumed to exist after `CREATE EXTENSION`.

Upstream documents builds across PostgreSQL 11–17. Verify the actual package against the target PostgreSQL release, especially for newer servers, and test calendar boundaries and invalid input before production use.
