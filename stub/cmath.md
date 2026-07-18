## Usage

Sources:

- [Upstream documentation at the reviewed commit](https://github.com/JoanneBogart/postgres-cmath/blob/d2ead80fda9f41099d766190c536d630aa0c8249/README.md)
- [Extension control file](https://github.com/JoanneBogart/postgres-cmath/blob/d2ead80fda9f41099d766190c536d630aa0c8249/cmath.control)

`cmath` exposes ISO C `<math.h>` functions with a `c_` prefix, including single-precision variants and functions not traditionally present in PostgreSQL. Unlike PostgreSQL numeric functions, the C functions commonly return IEEE-754 `NaN` or infinity for domain and overflow cases instead of raising an SQL error.

```sql
CREATE EXTENSION cmath;
SELECT c_log10(100), c_exp(1e100);
SELECT c_sinf(0.1), c_cosf(0.1);
```

Results follow the platform C library, so edge cases can vary across operating systems and architectures. Callers should explicitly reject non-finite values when downstream constraints, JSON serialization, indexes, or analytics require finite numbers. The reviewed version is `1.2`; test numerical behavior on the exact target platform before relying on reproducibility.
