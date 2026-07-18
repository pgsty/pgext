## Usage

Sources:

- [Upstream README](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/README.md)
- [Extension control file](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/pcg_random.control)
- [SQL installation script](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/pcg_random--0.0.1.sql)
- [C implementation](https://github.com/cayetanobv/pgsql-pcg-random/blob/cfac9eafc3d34f16f43a398fd2b14b692cc359a3/src/pcg_random.c)

`pcg_random` version `0.0.1` exposes scalar, bounded, array, and bounded-array functions backed by a PCG generator. The optional seed argument selects a sequence; the implementation also mixes time and an address into each call.

### Example

```sql
CREATE EXTENSION pcg_random;
SELECT pcg_random();
SELECT pcg_random_bound(100);
SELECT pcg_random_array(5);
SELECT pcg_random_bound_array(5, 100);
```

This is not a cryptographic random source. Pass a strictly positive bound and a small nonnegative array length; the C code does not adequately validate zero/negative bounds or lengths. More importantly, the install SQL declares nondeterministic functions `IMMUTABLE`, allowing planner constant folding and otherwise incorrect reuse. The source is from 2016 and has no current compatibility matrix; prefer maintained PostgreSQL randomness facilities.
