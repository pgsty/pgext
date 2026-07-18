## Usage

Sources:

- [Upstream README](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/README.rst)
- [Extension control file](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/eigen.control)
- [Version 1.1 SQL](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/sql/eigen--1.1.sql)
- [C++ matrix implementation](https://github.com/tlb-lab/pgeigen/blob/af9c58f84d58e629ed2a2329ffcbe8d6d5c95cff/src/matrixxd.cpp)

`eigen` version `1.1` wraps the Eigen C++ library with PostgreSQL domains and functions for three-dimensional vectors, integer/double arrays, and double matrices. It provides elementwise arithmetic, vector dot/cross products, norms and distances, set/similarity operations, matrix products, reductions, and stacking helpers.

### Example

```sql
CREATE EXTENSION eigen;
SELECT vector3d_dot(
  ARRAY[1.0, 0.0, 0.0]::vector3d,
  ARRAY[0.0, 1.0, 0.0]::vector3d
);
SELECT vector3d_norm(ARRAY[3.0, 4.0, 0.0]::vector3d);
```

Inputs are PostgreSQL arrays constrained by domains; enforce dimensions and reject null elements before expensive operations. The source targets PostgreSQL 9.1+ and Eigen 3 but has seen no repository update since 2020 and publishes no current compatibility matrix. It also defines many generic function and operator names, so install in a planned schema and test for conflicts and numerical edge cases.
