## Usage

Sources:

- [Pinned recurrence-rule SQL implementation](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/sql/pgsql_rrule.sql)
- [Pinned extension control file](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/pgsql_rrule.control)
- [Pinned upstream README](https://github.com/maternity/pgsql_rrule/blob/ce231b90c1d610b7880d25ce66d8af329da1f27e/README.md)

`pgsql_rrule` is a pure-SQL implementation of data types and functions for parsing and expanding RFC 5545 recurrence rules. It defines composite type `rrule`, frequency and weekday types, parser `rrule(text)`, set-returning `rrule_expand(...)`, and `get_occurrences(...)` compatibility functions.

```sql
CREATE EXTENSION pgsql_rrule;

SELECT *
FROM rrule_expand(
  rrule('FREQ=WEEKLY;BYDAY=MO,WE'),
  DATE '2026-07-01',
  DATE '2026-07-31'
);
```

The parser validates combinations of fields, but the reviewed expansion code is intentionally incomplete. It raises errors for unsupported frequencies and explicitly rejects weekly `BYSETPOS` and weekly `INTERVAL > 1`; test every recurrence shape used by the application, including inclusivity, time zones, fractional seconds, malformed input, and large/unbounded result sets.

The control file records `0.0.4`, generated from the repository's unversioned SQL source, while the project metadata marks that line unstable. Pin the exact commit rather than relying only on a version label, compare results with a maintained RFC 5545 implementation, and enforce an application-level upper bound on expansion.
