## Usage

Sources:

- [Official function documentation](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/doc/generate_date_series.md)
- [Extension SQL](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/sql/generate_date_series.sql)
- [C implementation](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/src/generate_date_series.c)

`generate_date_series` adds an inclusive, set-returning date generator whose step is an integer number of days. It avoids converting dates through timestamps or intervals when a query needs a simple ascending or descending calendar-date sequence.

### Core Workflow

```sql
CREATE EXTENSION generate_date_series;

SELECT date_val
FROM generate_date_series('2026-07-01'::date, '2026-07-07'::date) AS d(date_val);

SELECT date_val
FROM generate_date_series('2026-07-31'::date, '2026-07-01'::date, -7) AS d(date_val);
```

The two-argument call defaults to a one-day step. The three-argument form accepts positive or negative day counts and includes the end date only when the sequence lands on it.

### Function Contract

- `generate_date_series(start_date date, end_date date, step_days integer DEFAULT 1) -> setof date` is the only user-facing function.
- A positive step produces rows while the current date is less than or equal to the end date.
- A negative step produces rows while the current date is greater than or equal to the end date.
- A zero step raises `step size cannot equal zero`; a step pointing away from the end returns no rows.

### Operational Notes

Version `1.0.0` is relocatable and declares no preload requirement. The upstream build targets PostgreSQL `9.2` or later. The function is `IMMUTABLE` and `STRICT`, so any SQL `NULL` argument returns no rows.

The extension advances the internal PostgreSQL date value by the integer step. Bound generated ranges in user-facing queries to avoid unexpectedly large result sets, and test behavior near PostgreSQL's supported date limits before relying on extreme values.
