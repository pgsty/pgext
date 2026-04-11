
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pgcalendar;
> INSERT INTO pgcalendar.events (name, description, category)
> VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
> SELECT * FROM pgcalendar.get_event_projections(1, '2024-01-01'::date, '2024-01-07'::date);
> ```
>
> Source: [README](https://github.com/h4kbas/pgcalendar)

`pgcalendar` is a recurring calendar extension for PostgreSQL. It models events, schedules, exceptions, and projections, and generates calendar occurrences across arbitrary date ranges.

## Data Model

The README describes four main concepts:

- `events` as logical objects such as meetings or tasks
- `schedules` as non-overlapping recurrence definitions
- `exceptions` as per-occurrence cancellations or modifications
- `projections` as the actual generated calendar occurrences

## Quick Start

Create an event:

```sql
INSERT INTO pgcalendar.events (name, description, category)
VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');
```

Create a schedule:

```sql
INSERT INTO pgcalendar.schedules (
    event_id, start_date, end_date, recurrence_type, recurrence_interval
) VALUES (
    1, '2024-01-01 09:00:00', '2024-01-07 23:59:59', 'daily', 1
);
```

Get projections:

```sql
SELECT * FROM pgcalendar.get_event_projections(
    1, '2024-01-01'::date, '2024-01-07'::date
);
```

## Recurrence Types

The README shows schedule examples for:

- daily recurrence
- weekly recurrence with `recurrence_day_of_week`
- monthly recurrence with `recurrence_day_of_month`
- yearly recurrence with `recurrence_month` and `recurrence_day_of_month`

## Exceptions

Exceptions can cancel or modify a single occurrence:

```sql
INSERT INTO pgcalendar.exceptions (
    schedule_id, exception_date, exception_type, notes
) VALUES (
    1, '2024-01-15', 'cancelled', 'Holiday - meeting cancelled'
);
```

Modified occurrences can also change date and time.

## Functions and Views

The README documents:

- `get_event_projections(event_id, start_date, end_date)`
- `get_events_detailed(start_date, end_date)`
- `transition_event_schedule(...)`
- `check_schedule_overlap(event_id, start_date, end_date)`
- `pgcalendar.event_calendar`

`transition_event_schedule(...)` safely switches an event to a new schedule definition, while `check_schedule_overlap(...)` validates that new schedules do not overlap existing ones.
