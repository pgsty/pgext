## Usage

Sources:

- [Official README at the catalog revision](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/README.md)
- [Extension SQL at the catalog revision](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms--1.0.sql)

`pg_algorithms` 1.0 exposes demonstration implementations of bubble sort and quicksort for integer arrays. Upstream describes it as a pet project; use it for learning or small experiments, not as a production sorting primitive.

### Core Workflow

```sql
CREATE EXTENSION pg_algorithms;

SELECT bubble_sort(ARRAY[5, 1, 4, 2, 3]);
SELECT quick_sort(ARRAY[5, 1, 4, 2, 3]);

-- PostgreSQL's native relational sort is the normal production choice.
SELECT array_agg(v ORDER BY v)
FROM unnest(ARRAY[5, 1, 4, 2, 3]) AS u(v);
```

### Function Index

- `bubble_sort(integer[])` returns an ascending one-dimensional integer array using bubble sort.
- `quick_sort(integer[])` returns an ascending one-dimensional integer array using a recursive first-element-pivot quicksort.
- Both functions are strict and stable. The SQL surface exposes integer arrays only.

### Caveats

- NULL array elements and multidimensional arrays are rejected by the C implementation.
- Bubble sort has quadratic runtime. The chosen quicksort pivot also reaches quadratic runtime and deep recursion on already ordered or adversarial input.
- Prefer PostgreSQL's native ORDER BY for real workloads; it has mature memory, disk-spill, collation, and planner integration.
