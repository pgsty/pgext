## Usage

Sources:

- [pg_cron_utils 0.1.0 README](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/README.md)
- [Extension control file](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils.control)
- [SQL definitions](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils--0.1.0.sql)

`cron_utils` parses five-field cron expressions and calculates trigger timestamps inside PostgreSQL. It is a scheduling utility, not a job runner: use it to preview, validate, or query a schedule, and use a scheduler such as `pg_cron` separately to execute work.

### Core Workflow

```sql
CREATE EXTENSION cron_utils;

-- First trigger at or after the supplied time.
SELECT cron_first_trigger('0 9 * * 1-5', now());

-- Last trigger before the supplied time (strict is true by default).
SELECT cron_last_trigger('0 9 * * 1-5', now());

-- Next five hourly triggers.
SELECT *
FROM cron_iterate_n('0 * * * *', now(), false, 'next', 5);
```

To inspect a bounded window:

```sql
SELECT *
FROM cron_first_last_triggers(
  '0 0 * * *',
  date_trunc('month', now()),
  date_trunc('month', now()) + interval '1 month'
);
```

Either returned boundary can be `NULL` when the expression has no trigger in the window.

### Important Objects

- `cron_parts` is the parsed representation of the minute, hour, day, month, and day-of-week fields.
- `parse_cron(text)` parses `*`, lists, ranges, and step syntax.
- `cron_first_trigger(expr, base_time, strict)` searches forward. With `strict = true`, a trigger exactly at `base_time` is skipped.
- `cron_last_trigger(expr, base_time, strict)` searches backward and defaults to strict matching.
- `cron_first_last_triggers(expr, start_time, end_time)` returns the first and last matches in a window.
- `cron_iterate_n(expr, base_time, strict, direction, max_matches)` returns consecutive matches in the `next` or `prev` direction.

### Semantics and Caveats

Expressions use the standard five fields `minute hour day month dow`; seconds are not accepted. Day-of-week uses `1 = Monday` through `7 = Sunday`. Results are `timestamptz`, so session time zone affects the displayed local time and daylight-saving transitions should be tested for local-time schedules.

The extension is pure SQL/PL/pgSQL, relocatable, and has no `pg_cron` dependency. Its functions are declared immutable and parallel-safe. Version 0.1.0 is marked `unstable` in the control metadata, so pin behavior and retest edge cases before embedding it in a critical scheduler.
