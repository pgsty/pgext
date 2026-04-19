## Usage

Sources: [README](https://github.com/h4kbas/pgcalendar/blob/master/README.md), [repo](https://github.com/h4kbas/pgcalendar)

`pgcalendar` is a recurring calendar extension for PostgreSQL. The upstream README models recurring schedules with four main pieces: `events`, `schedules`, `exceptions`, and generated projections.

### Create events and schedules

```sql
CREATE EXTENSION pgcalendar;

INSERT INTO pgcalendar.events (name, description, category)
VALUES ('Daily Standup', 'Team daily standup meeting', 'meeting');

INSERT INTO pgcalendar.schedules (
    event_id, start_date, end_date, recurrence_type, recurrence_interval
) VALUES (
    1, '2024-01-01 09:00:00', '2024-01-07 23:59:59', 'daily', 1
);
```

The README shows `daily`, `weekly`, `monthly`, and `yearly` recurrences, with extra columns such as `recurrence_day_of_week`, `recurrence_day_of_month`, and `recurrence_month` depending on the recurrence type.

### Query projections

```sql
SELECT * FROM pgcalendar.get_event_projections(
    1, '2024-01-01'::date, '2024-01-07'::date
);

SELECT * FROM pgcalendar.get_events_detailed(
    '2024-01-01'::date, '2024-01-31'::date
);
```

The README also uses the `pgcalendar.event_calendar` view as a quick verification target.

### Exceptions and schedule transitions

```sql
INSERT INTO pgcalendar.exceptions (
    schedule_id, exception_date, exception_type, notes
) VALUES (
    1, '2024-01-15', 'cancelled', 'Holiday - meeting cancelled'
);

SELECT pgcalendar.transition_event_schedule(
    p_event_id := 1,
    p_new_start_date := '2024-01-15 09:00:00',
    p_new_end_date := '2024-01-31 23:59:59',
    p_recurrence_type := 'weekly',
    p_recurrence_interval := 2,
    p_recurrence_day_of_week := 1,
    p_description := 'Changed to bi-weekly schedule'
);
```

Use `pgcalendar.check_schedule_overlap(...)` before adding a new schedule when you need to verify that date ranges do not overlap.

### Caveat

The upstream README is the only user-facing documentation currently published. It gives clear table and function examples, but it does not add separate versioned release notes for user-visible SQL changes.
