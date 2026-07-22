## Usage

Sources:

- [Official README](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/README.md)
- [Official extension control file](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar.control)
- [Official extension SQL](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar--1.0.0.sql)

`geekspeak_calendar` 1.0.0 is an application-specific SQL extension for the GeekSpeak podcast database. It stores recurring recording windows and exposes a rolling `calendar` view that combines those schedules with episode, participant, person, and location data supplied by the required `geekspeak` extension.

### Core Workflow

Install the dependency and extension in the intended GeekSpeak schema, then add a recording schedule and query the generated calendar:

```sql
CREATE EXTENSION geekspeak;
CREATE EXTENSION geekspeak_calendar WITH SCHEMA gs;
SET search_path = gs, public;

INSERT INTO recording_schedules ("start", "end", location, cancellations)
VALUES (
  '2026-07-25 09:00:00+00',
  '2026-12-26 10:00:00+00',
  NULL,
  ARRAY['2026-08-01'::date]
);

SELECT title, description, participants, location, "start", "end"
FROM calendar
ORDER BY "start";
```

The time difference between `start` and `end` defines each recording's duration. Their dates delimit the recurring schedule, while `cancellations` removes individual dates and `location` supplies a default location.

### Important Objects

- `recording_schedules` stores the first start, final end, default location, and cancelled dates for a weekly schedule.
- `sessions_similarity_gist` is an exclusion constraint that rejects overlapping schedule windows with the same weekday and overlapping time-of-day ranges.
- `calendar` expands schedules around the current date, joins pending episode data, and returns title, description, participants, location, geometry, modification time, start, and end.

### Operational Notes

This is not a general iCalendar parser or exporter. The view is coupled to the `geekspeak` schema objects `episodes`, `participants`, `people`, and `locations`, and the control file declares `geekspeak` as a required extension. The rolling view uses `now()` and a fixed `generate_series(0, 50)` horizon, so its output changes with time and covers roughly the next 50 weekly occurrences plus the preceding week. Review the upstream SQL against the exact `geekspeak` schema before deployment.
