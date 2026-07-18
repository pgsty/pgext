## Usage

Sources:

- [Pinned upstream architecture and usage guide](https://github.com/postgrespro/vops/blob/8259b306c6d71b605fd0ac4e1c0b64846358e4c2/README.md)
- [Pinned extension control file](https://github.com/postgrespro/vops/blob/8259b306c6d71b605fd0ac4e1c0b64846358e4c2/vops.control)
- [Official PGXN distribution page](https://pgxn.org/dist/vops/)

`vops` provides vectorized tile types, operators, aggregates, projections, and an FDW for analytic PostgreSQL workloads. Types such as `vops_int4`, `vops_float8`, and `vops_bool` hold up to 64 scalar values per tile, reducing per-row executor overhead while retaining ordinary PostgreSQL heap storage.

Load the query-rewrite hook at startup; the guide also documents explicit `vops_initialize()` for sessions where preloading is unavailable:

```conf
shared_preload_libraries = 'vops'
```

After creating and populating a VOPS tile table, vector predicates use `filter`, `betwixt`, and overloaded boolean operator `&`:

```sql
CREATE EXTENSION vops;

SELECT sum(l_extendedprice * l_discount) AS revenue
FROM vops_lineitem
WHERE filter(
  betwixt(l_shipdate, '1996-01-01', '1997-01-01')
  & betwixt(l_discount, 0.08, 0.1)
  & (l_quantity < 24)
);
```

Tile columns are not scalar columns: null masks, casts, filtering, aggregation, projection refresh, and bulk loading have VOPS-specific semantics. The reviewed guide says joins are not supported by vector operators, and operator precedence requires parentheses around compound vector predicates. Validate query plans and results against a scalar reference table before trusting performance gains.

The pinned control version is `1.1`, while PGXN lists a later `2.0.1` distribution. Pin one source/version line, verify the server compatibility and rewrite-hook behavior, and treat projection refresh consistency, crash recovery, concurrent updates, and unsupported SQL shapes as explicit operational boundaries.
