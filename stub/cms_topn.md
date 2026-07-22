## Usage

Sources:

- [Official README](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/README.md)
- [Official extension control file](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/cms_topn.control)
- [Official extension SQL](https://github.com/ozturkosu/cms_topn/blob/78ce0d1e0437c0b35419d963685d5de57a87078e/cms_topn--1.0.0.sql)

`cms_topn` 1.0.0 stores an approximate Count-Min Sketch together with a bounded list of frequent-item candidates. It is useful for compact frequency estimates, approximate top-N reporting, and merging summaries from separate partitions when exact aggregation is too expensive.

### Core Workflow

Build one summary per group with `cms_topn_add_agg`, query candidate frequencies with `topn`, and merge compatible summaries with `cms_topn_union_agg`.

```sql
CREATE EXTENSION cms_topn;

CREATE TABLE daily_hits (
    day date PRIMARY KEY,
    users cms_topn NOT NULL
);

INSERT INTO daily_hits (day, users)
SELECT visited_at::date,
       cms_topn_add_agg(user_id, 10, 0.001, 0.99)
FROM visits
GROUP BY visited_at::date;

SELECT day, item, frequency
FROM daily_hits
CROSS JOIN LATERAL topn(users, NULL::bigint);

SELECT *
FROM topn(
    (SELECT cms_topn_union_agg(users) FROM daily_hits),
    NULL::bigint
);
```

The typed null passed to `topn` determines the SQL type of its `item` column. Use the same input type and sketch parameters for summaries that will be merged.

### Main Objects

- `cms_topn(n, error_bound, confidence)` constructs an empty summary; defaults are `0.001` and `0.99`.
- `cms_topn_add(summary, value)` adds one value; aggregate overloads of `cms_topn_add_agg` build summaries from a column.
- `cms_topn_frequency(summary, value)` returns an estimated frequency.
- `topn(summary, typed_null)` returns candidate `item` and estimated `frequency` rows.
- `cms_topn_union` and `cms_topn_union_agg` combine compatible summaries; `cms_topn_info` reports sketch dimensions and metadata.

Smaller error bounds increase the sketch width; higher confidence increases its depth. Both improve statistical guarantees at a memory and CPU cost.

### Approximation and Compatibility Boundaries

A Count-Min Sketch can overestimate frequency because of hash collisions, and the retained candidate list does not make the answer exact. Validate acceptable error with production-like distributions, especially heavy skew, many distinct values, and merged summaries. Do not use approximate counts for billing, integrity constraints, or other exact decisions.

The source and tests date from 2015. The binary custom type stores internal type-dependent data, so do not assume dump/restore, replication, architecture, or PostgreSQL-major compatibility without testing the exact build. Treat serialized summaries as derived data that can be rebuilt from source rows.
