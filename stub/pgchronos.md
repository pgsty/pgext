## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/worden341/pgchronos/blob/16731fb0cd98a915ef8732523dffddbef09f49ce/README.md)
- [Version 1.0.2 SQL objects](https://github.com/worden341/pgchronos/blob/16731fb0cd98a915ef8732523dffddbef09f49ce/sql/pgchronos--1.0.2.sql)
- [PGXN distribution](https://pgxn.org/dist/pgchronos/)

`pgchronos` is a pure-SQL demonstration of set operations over arrays of `daterange` and `tstzrange`. It supplies union, intersection, difference, reduction, and containment functions and corresponding `+`, `*`, `-`, `<@`, and `@>` operators.

```sql
CREATE EXTENSION pgchronos;
SELECT ARRAY[daterange('2026-01-01', '2026-02-01')]
       - ARRAY[daterange('2026-01-10', '2026-01-15')];
```

The result is an array of non-overlapping remaining ranges. Union and `reduce` combine overlapping or adjacent ranges, while intersection retains common segments. Confirm boundary inclusivity, empty and infinite bounds, nulls, ordering, duplicates, daylight-saving behavior, and normalization for the exact range subtype.

Upstream explicitly corrects an earlier “stable” label: the entire 1.x line is an alpha-quality conceptual demonstration and is not suitable for production. The repository is abandoned. Prefer maintained multirange features in supported PostgreSQL releases or thoroughly reviewed application SQL. If preserving old data, create property-based tests for algebraic identities and migrate operator-dependent expressions carefully because custom operators can change query meaning globally.
