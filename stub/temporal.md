## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/temporal/)
- [Official function and operator reference](https://pgxn.org/dist/temporal/doc/html/reference.html)
- [Official tutorial](https://pgxn.org/dist/temporal/doc/html/tutorial.html)
- [Version 0.7.1 source archive](https://api.pgxn.org/dist/temporal/0.7.1/temporal-0.7.1.zip)

`temporal` provides a fixed 16-byte `period` type for timestamp-with-time-zone intervals, plus constructors, relation operators, and B-tree and GiST indexing. It is a legacy extension released in 2011, before PostgreSQL's built-in range and multirange types.

### Core Workflow

Create the extension and store half-open periods with the default two-argument constructor.

```sql
CREATE EXTENSION temporal;

CREATE TABLE meeting (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  during period NOT NULL
);

INSERT INTO meeting(during)
VALUES (period(
  timestamptz '2026-07-22 09:00+08',
  timestamptz '2026-07-22 10:00+08'
));
```

Containment and overlap operators are supported by the default GiST operator class.

```sql
CREATE INDEX meeting_during_gist
ON meeting USING gist (during);

SELECT *
FROM meeting
WHERE during @> clock_timestamp();

SELECT *
FROM meeting
WHERE during && period(
  timestamptz '2026-07-22 09:30+08',
  timestamptz '2026-07-22 11:00+08'
);
```

### Important Objects

- `period`: canonical output is a lower-inclusive, upper-exclusive timestamp interval.
- `period(timestamptz)`, `period(timestamptz, timestamptz)`: point and default half-open constructors.
- `period_oo`, `period_oc`, `period_co`, `period_cc`: constructors selecting open/closed endpoints.
- `empty_period()`, `is_empty(period)`: empty-value construction and testing.
- `first`, `last`, `prior`, `next`, `length`: endpoint and duration accessors.
- `@>`, `<@`, `&&`, `<<`, `>>`, `&<`, `&>`: contains, contained-by, overlap, and relative-position operators.
- `period_intersect`, `period_union`, `minus`: set-like operations on periods.
- `gist_period_ops`, `btree_period_ops`: default GiST and B-tree operator classes.

### Semantics and Migration Boundary

`period_union` raises an error for disjoint, non-adjacent inputs. `minus` returns one period and raises an error when subtraction would split the result into two periods. Verify endpoint semantics carefully when using the explicit open/closed constructors.

Version `0.7.1` changed only PGXN metadata from `0.7.0`; its implementation dates to PostgreSQL versions far older than current releases. The type is not SQL-standard system/application `PERIOD` syntax and provides no temporal primary-key or foreign-key semantics. For new designs, prefer built-in `tstzrange` and multiranges unless compatibility with stored `period` values is required. Before migrating, compare empty values, inclusive endpoints, adjacency, union/minus errors, binary storage, and index behavior rather than casting blindly.
