## Usage

Sources:

- [Official introduction](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/docs/source/intro.rst)
- [Extension control file](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/pgsfti.control)
- [Extension SQL](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/sql/pgsfti--1.0.sql)
- [Type and relation implementation](https://github.com/OnroerendErfgoed/pgSFTI/blob/7a7bab98746a1fe9bd0ddf7ee6452196ae476a6e/src/pgsfti.c)

`pgsfti` adds the `sfti` type for simple trapezoidal fuzzy time intervals. It is useful when a historical or domain event has uncertain support and core boundaries and queries need graded Allen-style temporal relations rather than only sharp PostgreSQL date ranges.

### Core Workflow

```sql
CREATE EXTENSION pgsfti;

CREATE TABLE historical_event (
  event_name text PRIMARY KEY,
  happened sfti NOT NULL
);

INSERT INTO historical_event VALUES
  ('period_a', sfti_makeSFTI(1900, 1925, 1975, 1999)),
  ('period_b', sfti_fuzzify(1980, 1990, 3, 5));

SELECT event_name
FROM historical_event
WHERE happened && sfti_makeSFTI(1970, 1985);

SELECT allen_overlaps(
  sfti_makeSFTI(1900, 1925, 1975, 1999),
  sfti_makeSFTI(1970, 1980, 2000, 2010),
  0.0,
  0.0
);
```

An `sfti` stores five floating-point values: support start, core start, core end, support end, and maximum membership. `sfti_makeSFTI` has overloads for years and dates; `sfti_fuzzify` expands sharp years, dates, or intervals with fuzzy margins. Implicit casts exist from `smallint`, `integer`, and `date`.

Relation functions include the Allen family `allen_before`, `allen_meets`, `allen_overlaps`, `allen_during`, `allen_starts`, `allen_finishes`, `allen_equals`, and their converses, plus `kvd_before`, `kvd_after`, `kvd_during`, `kvd_contains`, and `kvd_intersects`. They return a floating membership value and take two tolerance parameters. Boolean operators `<`, `>`, `=`, `&&`, `@`, and `~` provide fixed-threshold shortcuts.

### Representation Caveats

Dates are converted to floating year positions rather than stored as calendar dates; there is no year zero, so 50 BC maps to -49. Conversion and output use floating-point values and cannot preserve every calendar representation exactly. The input function only parses five numbers and does not enforce support/core ordering or a membership range, so validate `support_start <= core_start <= core_end <= support_end` and a domain-appropriate membership value in table constraints or application code.

The SQL defines operators but no index operator class, so relation predicates should be expected to scan unless another indexed coarse filter is added. No preload is required. Test version 1.0 on the exact PostgreSQL major, and preserve the extension binary during dump/restore because tables depend on its custom base type.
