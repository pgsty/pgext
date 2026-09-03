## Usage

Sources:

- [Official README](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/README.md)
- [Extension control file](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/rrule_safe.control)
- [pgrx manifest](https://codeberg.org/Natureshadow/pg-rrule-safe/src/tag/0.2.1/Cargo.toml)

`rrule_safe` evaluates iCalendar recurrence rules inside PostgreSQL with a Rust implementation intended as a memory-safe replacement for the unmaintained `pg_rrule` library.

### Enablement

Release 0.2.1 supports PostgreSQL 13–18 through pgrx 0.18.0. Install its files for the exact PostgreSQL major, then create the extension as a superuser:

```sql
CREATE EXTENSION rrule_safe;
```

The control file is non-relocatable, untrusted, and loads `rrule_safe`; it does not require preloading or a server restart.

### Recurrence Expansion

`get_occurrences` accepts an RRULE string, a start timestamp, and an optional inclusive upper bound. Timestamp and timestamp-with-time-zone overloads return arrays of matching occurrences.

```sql
SELECT get_occurrences(
  'FREQ=WEEKLY;COUNT=4',
  '2026-01-01 09:00:00 Europe/Berlin'::timestamptz
);

SELECT get_occurrences(
  'FREQ=DAILY',
  '2026-01-01 09:00:00'::timestamp,
  '2026-01-07 09:00:00'::timestamp
);
```

For compatibility, the extension defines `rrule` as a domain over `text`; it does not reproduce `pg_rrule`'s stored custom type. Migrate persisted legacy values explicitly.

### Safety Boundary

Unbounded or malformed recurrence rules can be expensive. `rrule_safe.limit` caps the number of returned occurrences and defaults to 65,535. Set a lower session value for user-controlled rules, and always supply an upper bound where the application has one.

