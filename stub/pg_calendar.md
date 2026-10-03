## Usage

Sources:

- [extensions/pg_calendar/pg_calendar.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/pg_calendar.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_calendar/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/Cargo.toml)
- [extensions/pg_calendar/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_calendar/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/seed.rs)
- [extensions/pg_calendar/src/compute.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/compute.rs)
- [extensions/pg_calendar/src/pattern.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_calendar/src/pattern.rs)

`pg_calendar` 0.3.0 manages named working calendars, weekly patterns, holidays and explicit date exceptions in the `pgcalendar` schema.

### Core Workflow

```sql
CREATE EXTENSION pg_calendar;
SELECT pgcalendar.seed_iso();
SELECT pgcalendar.is_working_day('iso', DATE '2026-10-02');
SELECT pgcalendar.add_working_days('iso', DATE '2026-10-02', 3);
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. `create_calendar`, `set_working_days`, `add_holiday` and `add_exception` define calendars; `is_working_day`, `add_working_days` and interval-counting functions query them. Exceptions override holidays and weekly patterns. Seed functions create working-week conventions, not authoritative worldwide holiday data. The `pg_calendar.enabled` setting is superuser controlled. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
