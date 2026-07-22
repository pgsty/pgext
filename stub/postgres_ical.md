## Usage

Sources:

- [Official README](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/README.md)
- [Extension control file](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres_ical.control)
- [SQL-visible Rust implementation](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/src/lib.rs)
- [iCalendar event parser](https://github.com/edgarogh/postgres-ical/blob/8a2a9b7fa32002c53298b723de86045163812f74/postgres-ical-parser/src/parser.rs)

`postgres_ical` 0.1.0 parses RFC 5545 iCalendar `VEVENT` components into PostgreSQL rows. It accepts either in-memory calendar text or a URL fetched by the database backend, making it useful for explicit queries or carefully controlled materialized-view refreshes.

### Parse In-Memory Calendar Text

Select explicit columns because upstream reserves the right to add output columns without treating that as a breaking change:

```sql
CREATE EXTENSION postgres_ical;

SELECT uid, summary, dt_start, dt_start_naive
FROM pg_ical($ical$
BEGIN:VCALENDAR
BEGIN:VEVENT
UID:event-42@example.org
DTSTART:20260722T090000Z
SUMMARY:Review
END:VEVENT
END:VCALENDAR
$ical$);
```

UTC or timezone-qualified values populate the `timestamptz` column such as `dt_start`; floating local values populate the companion `timestamp` column such as `dt_start_naive`.

### Fetch a Calendar URL

```sql
SELECT uid, summary, dt_start, dt_start_naive
FROM pg_ical_curl('https://example.com/calendar.ical');
```

`pg_ical_curl(url text)` streams the response through libcurl and returns the same row shape as `pg_ical(calendar text)`.

### Parsed Surface

The current parser requires `UID` and `DTSTART`. It recognizes `CREATED`, `DESCRIPTION`, `DTSTAMP`, `DTEND`, `LAST-MODIFIED`, `LOCATION`, `SEQUENCE`, and `SUMMARY`; `SEQUENCE` defaults to zero. Unknown properties and non-`VEVENT` components are skipped.

The returned row type exposes more RFC 5545-oriented columns, but the reviewed implementation currently leaves attachment, categories, class, comments, completed, due, duration, geography, percent-complete, priority, resources, and status as NULL or empty values. Every returned `component_type` is `VEVENT`.

### Security and Compatibility Boundaries

`pg_ical_curl` makes outbound requests from the PostgreSQL backend and the reviewed code sets no explicit protocol allowlist, response-size limit, or request timeout. Revoke it from untrusted roles and constrain egress to avoid server-side request forgery and backend exhaustion. Prefer fetching and validating remote data outside the database when the URL is user-controlled.

Malformed events, missing required properties, invalid timezones, and transport failures are unwrapped by the Rust code and surface as PostgreSQL errors. The upstream README and Cargo manifest state that the project is unlicensed/`UNLICENSED`, so source redistribution and production adoption require a separate legal review.
