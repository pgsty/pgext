## Usage

Sources:

- [pg_algorithms README at the reviewed commit](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/README.md)
- [pg_algorithms.control at the reviewed commit](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms.control)
- [Version 1.0 install SQL](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/pg_algorithms--1.0.sql)
- [Bubble-sort implementation](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/bubble_sort.c)
- [Quick-sort implementation](https://github.com/kostiantyn-nemchenko/pg_algorithms/blob/e258c1cd1710f5867d21c864a01ad90f0f04326c/quick_sort.c)

`pg_algorithms` is a learning project that exposes two C functions for sorting a one-dimensional `int[]`: `bubble_sort` and `quick_sort`. Both return a new ascending array and reject arrays containing `NULL`.

### Sort Integer Arrays

```sql
CREATE EXTENSION pg_algorithms;

SELECT bubble_sort(ARRAY[9, 0, -2, 9, 4]);
SELECT quick_sort(ARRAY[9, 0, -2, 9, 4]);
```

### Caveats

- Upstream explicitly says not to use the project in production. The reviewed source dates from 2018 and provides no modern PostgreSQL compatibility matrix.
- `bubble_sort` has `O(n²)` comparison cost, while the hand-written recursive quick sort can degrade on unfavorable input. PostgreSQL's native ordering facilities are more appropriate for operational workloads.
- Version `1.0` implements only the two functions above, despite the project's broader algorithm name. The SQL signature accepts PostgreSQL `integer[]`, not arbitrary numeric or polymorphic arrays.
