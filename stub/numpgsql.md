## Usage

Sources:

- [Upstream README](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/README.md)
- [Extension control file](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/numpgsql.control)
- [Array mathematics SQL](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/src/math/math.sql)
- [Random-array SQL](https://github.com/tarkmeper/numpgsql/blob/a5098af9b7c4d88564092c0a4809029aab9f614f/src/random/random.sql)

`numpgsql` version `0.0.1` provides NumPy-inspired operations over PostgreSQL numeric arrays. Its C++/Boost implementation includes elementwise arithmetic and transcendental functions, vector reductions and aggregates, index/boolean slicing, sorting/reversal, and seeded random arrays and distributions.

### Example

```sql
CREATE EXTENSION numpgsql;
SELECT cos(ARRAY[5, 1, 6, 4]::real[]);
SELECT ARRAY[10, 20, 30, 40] @ ARRAY[2, 4];
SELECT random_seed(42);
SELECT random_randn(2, 2);
```

The project lists broadcasting, type conversion, axis aggregation, and documentation as unfinished and has not been updated since 2019. Many functions overload common names and operators, and result types inherit PostgreSQL array element semantics, including integer truncation/overflow behavior. Do not treat it as NumPy-compatible without tests for shape, type, null, overflow, and reproducibility requirements.
