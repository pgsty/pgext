## Usage

Sources:

- [pg_rational v0.0.3 README](https://github.com/begriffs/pg_rational/blob/v0.0.3/README.md)
- [pg_rational v0.0.3 control file](https://github.com/begriffs/pg_rational/blob/v0.0.3/pg_rational.control)
- [Changes through v0.0.3](https://github.com/begriffs/pg_rational/compare/v0.0.2...v0.0.3)

`pg_rational` provides exact fractional arithmetic in a fixed 64-bit PostgreSQL type. Use `rational` for values that must remain exact and for user-defined row ordering where new positions need to be inserted between existing positions without renumbering the table.

### Exact Arithmetic

```sql
CREATE EXTENSION pg_rational;

SELECT 1::rational / 3 * 3 = 1;
SELECT '1/3'::rational + '2/7'::rational;
SELECT rational_simplify('36/12');
```

The extension detects arithmetic overflow instead of silently wrapping. `ratt` is a helper type for tuple coercion:

```sql
SELECT 1 + (i, i + 1)::ratt
FROM generate_series(1, 5) AS i;
```

Conversions are available between integer values, floating-point values, and rationals. Converting a float finds a rational approximation; converting a rational to float loses exactness.

### Stable User-Defined Ordering

```sql
CREATE SEQUENCE todos_seq AS integer;

CREATE TABLE todos (
  prio rational UNIQUE DEFAULT nextval('todos_seq')::integer,
  what text NOT NULL
);

INSERT INTO todos (what)
VALUES ('install extension'), ('read about it'), ('try it');

UPDATE todos
SET prio = rational_intermediate(1, 2)
WHERE what = 'try it';

SELECT * FROM todos ORDER BY prio;
```

Use an `integer` sequence and cast `nextval()` explicitly. The extension intentionally has no implicit `bigint`-to-`rational` conversion because its numerator is limited to the PostgreSQL `integer` range.

### Indexes, Aggregates, and Caveats

- `rational` supports btree and hash operator classes, so it can be used in ordered and equality indexes.
- The extension supplies `min(rational)`, `max(rational)`, and `sum(rational)` aggregates in addition to arithmetic and comparison operators.
- `rational_intermediate(lower, upper)` walks a Stern-Brocot tree to find a fraction between its arguments. Extremely narrow ranges take longer, and v0.0.3 has no maximum-depth parameter; do not expose attacker-controlled pathological bounds without a statement timeout.
- Values are exact only while arithmetic stays within the type's numerator and denominator limits. Handle overflow errors rather than falling back silently to floating point.
- Version 0.0.3 is primarily a build-compatibility and documentation release; the user-facing rational arithmetic surface remains stable.
