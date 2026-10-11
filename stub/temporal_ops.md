## Usage

Sources:

- [README.md](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/README.md)
- [temporal_ops--1.0.0.sql](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops--1.0.0.sql)
- [temporal_ops.c](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops.c)
- [temporal_ops.control](https://github.com/pjungwir/temporal_ops/blob/516fe77bf732e1858329dc583a7cfbea43a4e81b/temporal_ops.control)

`temporal_ops` provides range-aware joins: a semijoin keeps the covered portions of a left row, an antijoin keeps uncovered portions, and an outer join also represents unmatched time. It combines C query generation with SQL-callable wrappers.

### Core Workflow

```sql
CREATE EXTENSION temporal_ops;
CREATE TABLE temporal_left (id integer, valid_at daterange);
CREATE TABLE temporal_right (id integer, valid_at daterange);
INSERT INTO temporal_left VALUES (1, '[2026-01-01,2026-02-01)');
INSERT INTO temporal_right VALUES (1, '[2026-01-10,2026-01-20)');
SELECT (j.item).id, j.valid_at
FROM temporal_semijoin('temporal_left', 'id', 'valid_at',
                       'temporal_right', 'id', 'valid_at')
AS j(item temporal_left, valid_at daterange);
```

### Operational Boundaries

This source uses the PostgreSQL 19 `SupportRequestInlineInFrom` planner API; older-major compatibility is not claimed. Installation needs administrator privileges for C and leakproof functions. `temporal_semijoin`, `temporal_antijoin` and `temporal_outer_join` accept relation names as `regclass`, matching key columns, and range-column names. Array overloads support composite keys. Returned records need an explicit caller-supplied row definition. The `_sql` variants expose generated SQL. Relations must exist and the caller must be allowed to read them. Union, difference and aggregation are listed as future work, not implemented operators. No project license declaration was found.
