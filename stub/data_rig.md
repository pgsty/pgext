## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/README.md)
- [Version 1.0 SQL objects](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig--1.0.sql)
- [C implementation](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.c)
- [Extension control file](https://github.com/Octonica/data_rig/blob/e099c8f3a7a4a471a448c4574b7d5049ef7617c3/data_rig.control)

`data_rig` is an experimental native type for representing a multidimensional OLAP fact as a sorted set of encoded integers. It supplies containment operators and a default GiST operator class.

### Core Workflow

The SQL script refers to the `cube` type in GiST support signatures, but the control file does not declare `cube` in `requires`. Install it explicitly before `data_rig`; `CREATE EXTENSION data_rig CASCADE` cannot infer this dependency.

```sql
CREATE EXTENSION cube;
CREATE EXTENSION data_rig;

CREATE TABLE fact_demo (id bigint, dimensions fact);
INSERT INTO fact_demo
VALUES (1, fact(ARRAY[to_fact_number(1, 10), to_fact_number(2, 20)]));
CREATE INDEX fact_demo_dimensions_idx
  ON fact_demo USING gist (dimensions gist_fact_ops);

SELECT id
FROM fact_demo
WHERE dimensions @> fact(ARRAY[to_fact_number(1, 10)]);
```

`fact(integer[])` sorts and deduplicates the encoded members and rejects arrays containing NULL. `@>` and `<@` test containment, `fact_intersect(fact, fact)` returns the common members, and `to_fact_number(dimension, value)` packs a dimension and value into an `int4`.

### Limitations and Safety

Text input is unusable: the type's `fact_in(cstring)` function always raises `fact_in() not implemented`, so construct values with `fact(integer[])` rather than literals or text casts. The output routine exists mainly for display and does not provide a round-trip text format.

Upstream provides only a one-line README and no compatibility matrix or meaningful regression suite. The custom type and GiST implementation contain unchecked pointer and array arithmetic; malformed values or index edge cases may return incorrect results or crash a backend. Do not use `data_rig` for production data. If evaluating it, use a disposable database, compare every indexed result with a sequential scan, exercise empty and boundary sets, and be prepared to drop and rebuild indexes after any binary change.
