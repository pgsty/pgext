## Usage

Sources:

- [Official README](https://github.com/pjungwir/range_agg/blob/a9e825be48954ff1d83b619f7b7be7eaf29fc3be/README.md)
- [Official version 1.2.1 SQL](https://github.com/pjungwir/range_agg/blob/a9e825be48954ff1d83b619f7b7be7eaf29fc3be/range_agg--1.2.1.sql)
- [Official PGXN release](https://pgxn.org/dist/range_agg/1.2.1/)
- [PostgreSQL built-in aggregate documentation](https://www.postgresql.org/docs/current/functions-aggregate.html)

`range_agg` version `1.2.1` merges adjacent or overlapping PostgreSQL ranges and can reject gaps or overlaps. It predates PostgreSQL's built-in multirange aggregate of the same name, so modern installations must account for name resolution and different return types.

### Core Workflow

On a server where this extension is appropriate, merge contiguous reservations per room:

```sql
CREATE EXTENSION range_agg;

SELECT room_id, range_agg(booked_during)
FROM reservations
GROUP BY room_id;
```

The one-argument form returns one range and raises an error if it encounters a gap or overlap. To coalesce ranges while permitting gaps, use the two-argument form and unnest the returned array:

```sql
SELECT room_id, merged_range
FROM (
  SELECT room_id, range_agg(booked_during, true) AS merged_ranges
  FROM reservations
  GROUP BY room_id
) AS grouped
CROSS JOIN LATERAL unnest(grouped.merged_ranges) AS merged_range;
```

### Aggregate Forms

- `range_agg(anyrange)` returns one range and rejects gaps and overlaps.
- `range_agg(range, permit_gaps)` returns an array, allows gaps when true, and still rejects overlaps.
- `range_agg(range, permit_gaps, permit_overlaps)` returns an array and independently controls both conditions.
- The two- and three-argument forms are declared for `int4range`, `int8range`, `numrange`, `tsrange`, `tstzrange`, and `daterange`. Custom range types need matching transition/final functions and aggregate declarations.

### Operational Caveats

PostgreSQL 14 and later provide `pg_catalog.range_agg(anyrange)` returning a multirange. That built-in has different semantics and normally wins unqualified name lookup, while this extension's one-argument aggregate returns a single range. Verify the selected function with `regprocedure`, schema-qualify calls if coexistence is intentional, and prefer the maintained built-in on current PostgreSQL. Aggregation errors abort the statement when ranges overlap or contain forbidden gaps. Test empty groups, infinite bounds, discrete versus continuous ranges, and custom canonicalization before relying on coalescing results.
