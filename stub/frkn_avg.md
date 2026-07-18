## Usage

Sources:

- [Version 0.0.1 SQL definitions](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg--0.0.1.sql)
- [C implementation](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg.c)
- [Extension control file](https://github.com/furkansahin/Aggregate-Average/blob/b1e8d5425e7c62b4296a96cbf699e3db6cd46e2e/frkn_avg.control)

`frkn_avg` defines `my_avg(bigint)` and `my_avg(float)` aggregates backed by a custom `my_state` C type. Both accumulate their sum and count using single-precision fields.

```sql
CREATE EXTENSION frkn_avg;
SELECT my_avg(v::bigint), my_avg(v::float)
FROM generate_series(1, 10) AS g(v);
```

This implementation is experimental and has material correctness hazards: the SQL `bigint` transition is read with a 32-bit C accessor, floating-point input is narrowed, and state precision degrades as values or row counts grow. The declared internal length also differs from the two-float C structure. Do not expose it to untrusted workloads or use it for financial or exact analytics; compare results against PostgreSQL `avg` on representative data first.
