## Usage

Sources:

- [Argm README at the reviewed commit](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/README.md)
- [argm.control at the reviewed commit](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/argm.control)
- [Version 1.1.2 install SQL](https://github.com/bashtanov/argm/blob/b8b2db36d0a57987d413f909e82369f9677e384a/argm--1.1.2.sql)
- [Argm on PGXN](https://pgxn.org/dist/argm/)

`argm` provides the `argmax`, `argmin`, and `anyold` aggregates. The first two return the value from the row whose sortable key tuple is lexicographically greatest or least within a group. `anyold` returns the first non-null value seen by the aggregate and is useful when a selected attribute is functionally determined by the grouping key.

### Basic Example

```sql
CREATE EXTENSION argm;

SELECT
  argmax(label, score) AS highest_label,
  argmin(label, score) AS lowest_label,
  anyold(label) AS one_label
FROM (VALUES
  ('alpha', 10),
  ('beta', 20),
  ('gamma', NULL)
) AS sample(label, score);
```

Unlike a `DISTINCT ON` formulation, these aggregates can participate in the same `GROUP BY` operation as other aggregates and may use hash aggregation.

### Caveats

- If several rows have identical key tuples, `argmax` or `argmin` chooses one corresponding value arbitrarily.
- Values may use any PostgreSQL type, but ordering keys must be sortable. The upstream implementation applies one ordering direction and collation to the complete key tuple.
- Current source and catalog identify `1.1.2`; the latest visible Git tag is `1.1.1`, and PGXN still presents `1.0.2`. Do not infer current behavior from the older PGXN package.
