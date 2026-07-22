## Usage

Sources:

- [Official README](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/README.md)
- [Version 1.0 SQL implementation](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg--1.0.sql)
- [Extension control file](https://github.com/pjungwir/pjpg/blob/9437cb6b5efaa7bb7e67e625226c69962486a9a0/pjpg.control)

`pjpg` is a pure-SQL collection of general utilities, chiefly for interpreting `timestamp without time zone` values in named time zones and constructing day or week `tsrange` windows. It also provides `blank_to_null(text)` for normalizing empty or whitespace-only strings.

### Time-Zone and Range Workflow

```sql
CREATE EXTENSION pjpg;

SELECT in_tz(timestamp '2026-07-22 00:00:00', 'UTC', 'Asia/Shanghai');
SELECT beginning_of_day(timestamp '2026-07-22 14:30:00', 'Asia/Shanghai');
SELECT day_range(timestamp '2026-07-22 14:30:00', 'Asia/Shanghai', 7);
SELECT blank_to_null('   ');
```

The two-argument `in_tz(t, to_tz)` assumes that `t` represents UTC. `beginning_of_week_usa()` starts on Sunday, `beginning_of_week_iso()` starts on Monday, and `beginning_of_week()` aliases the USA convention. `day_range()`, `weekusa_range()`, `weekiso_range()`, and `week_range()` return half-open `[)` ranges with a caller-supplied width.

### Timestamp Semantics

Inputs and outputs deliberately have no time-zone type. They represent local wall-clock readings under the caller's convention, so the database cannot remember which zone was used. Daylight-saving transitions can make a local time ambiguous or nonexistent; test both transition directions for every deployed zone, and prefer `timestamptz` for stored instants.

Time-zone results depend on PostgreSQL's installed tzdata and may change after time-zone database updates. Validate nonpositive range widths, NULL and blank inputs, ISO versus USA week requirements, and downstream exclusion-bound behavior. Schema-qualify these generic function names when `search_path` is not tightly controlled.
