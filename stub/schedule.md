## Usage

Sources:

- [Official README](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/README.md)
- [Extension SQL](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule--1.0.sql)
- [C type implementation](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/schedule.c)
- [Haskell schedule implementation](https://github.com/albertov/pg_schedule/blob/3d4ed9a6305b7a206f3c234a0e5acf4afa93cbc5/PGSchedule.chs)

`schedule` 1.0 adds a validated cron-formatted base type and functions that test or enumerate matching timestamps. It performs schedule calculations only; it does not execute jobs. All calculations are in UTC, so it is unsuitable for a recurring local wall-clock time that must follow daylight-saving transitions.

### Core Workflow

Create the extension, cast a cron expression to the custom type, and query its UTC occurrences:

```sql
CREATE EXTENSION schedule;

SELECT schedule_contains('0 */6 * * *'::schedule, '2026-07-22 06:00:00+00');
SELECT schedule_next('0 */6 * * *'::schedule, '2026-07-22 06:00:00+00');
SELECT *
FROM schedule_series(
  '0 */6 * * *'::schedule,
  '2026-07-22 00:00:00+00',
  '2026-07-23 00:00:00+00'
);
```

The series includes matching values from the lower bound through the inclusive upper bound.

### Function and Operator Index

- `schedule_contains` tests membership.
- `schedule_next` and `schedule_previous` return strictly later or earlier matches.
- `schedule_floor` and `schedule_ceiling` include the input instant when it matches.
- `schedule_series` returns matches across a bounded interval.

Assignment casts to and from text are provided. Comparison operators and a B-tree operator class compare the stored cron strings lexically, not the semantic set of matching timestamps; equivalent schedules written differently can compare unequal.

### Operational Notes

The type preserves the original validated expression text. `schedule_series` asks the Haskell layer to build the complete timestamp array before PostgreSQL returns rows, so bound ranges and frequency to prevent excessive memory use. The source is abandoned, dates from a PostgreSQL 9.5/Nix-era build, and embeds a Haskell runtime; verify toolchain and PostgreSQL ABI compatibility before any use. Handle no-result and timestamp-range errors, and use a maintained scheduler for actual job execution.
