## Usage

Sources:

- [Official extension control file](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/geekspeak_calendar.control)
- [Official upstream documentation](https://github.com/ttfkam/pg_geekspeak_calendar/blob/34dcd6de80a4bf3116a19d21d0d222337c1f1d7f/README.md)

`geekspeak_calendar` — Calendar and iCalendar integration for GeekSpeak podcast recording schedules and episodes.

The reviewed catalog snapshot records version `1.0.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `geekspeak`.

```sql
CREATE EXTENSION "geekspeak_calendar";
SELECT extversion
FROM pg_extension
WHERE extname = 'geekspeak_calendar';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
