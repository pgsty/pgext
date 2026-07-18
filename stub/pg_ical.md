## Usage

Sources:

- [Current upstream feature and usage documentation](https://github.com/thomastthai/pg_ical/blob/b68e3ce2cb90686739e831f052eeed7324f8caa7/README.md)
- [Extension control file](https://github.com/thomastthai/pg_ical/blob/b68e3ce2cb90686739e831f052eeed7324f8caa7/pg_ical.control)

`pg_ical` adds functions for RFC 5545 recurrence rules backed by `libical`. It parses RRULE fields and expands recurring events, including recurrence dates, exclusion dates, time zones, durations, and recurrence identifiers.

```sql
CREATE EXTENSION pg_ical;
SELECT get_freq('FREQ=WEEKLY;INTERVAL=1;BYDAY=MO,TH');
SELECT unnest(get_occurrences(
  'FREQ=DAILY;COUNT=3',
  '2026-01-01 09:00:00+00'::timestamptz
));
```

The reviewed version is `0.2.0` and upstream explicitly does not support the deprecated EXRULE feature. Calendar correctness depends on RFC interpretation, `libical` version, daylight-saving transitions, inclusive boundaries, and timestamp type. Test real time zones and recurrence limits, and cap expansions generated from untrusted rules.
