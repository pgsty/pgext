

## Usage

> [pg_rrule: iCalendar RRULE recurrence rule type for PostgreSQL](https://github.com/petropavel13/pg_rrule)

The `pg_rrule` extension provides an RRULE data type for parsing and expanding iCalendar recurrence rules (RFC 5545).

```sql
CREATE EXTENSION pg_rrule;
```

### Parameter Extraction

```sql
-- Get frequency
SELECT get_freq('FREQ=WEEKLY;INTERVAL=1;WKST=MO;UNTIL=20200101T045102Z'::rrule);
-- WEEKLY

-- Get days of week (numeric array)
SELECT get_byday('FREQ=WEEKLY;BYDAY=MO,TH,SU'::rrule);
-- {2,5,1}
```

### Occurrence Generation

```sql
-- Generate occurrences from an RRULE
SELECT unnest(get_occurrences(
    'FREQ=DAILY;INTERVAL=1;UNTIL=20200105T000000Z'::rrule,
    '2020-01-01 00:00:00'::timestamp
));
```

The `get_occurrences()` function expands RRULE definitions into concrete timestamp sequences, supporting both timezone-aware and naive timestamp parameters.

### Dependencies

Requires `libical` for iCalendar standard compliance.
