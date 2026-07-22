## Usage

Sources:

- [Extension control file](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.control)
- [Extension SQL](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench--1.0.sql)
- [Benchmark implementation](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.c)

`palloc_bench` is a tiny developer microbenchmark for PostgreSQL memory allocation. It creates nested allocation contexts, repeatedly calls `palloc` and `pfree` in the deepest context, and reports elapsed wall-clock time as a warning; it does not return a structured benchmark result.

### Core Workflow

Use only in an isolated test database with conservative inputs and no competing workload.

```sql
CREATE EXTENSION palloc_bench;

REVOKE ALL ON FUNCTION palloc_bench(integer, integer, integer) FROM PUBLIC;
GRANT EXECUTE ON FUNCTION palloc_bench(integer, integer, integer) TO benchmark_admin;

SET client_min_messages = warning;
SELECT palloc_bench(8, 100000, 64);
```

The parameters are the number of nested memory contexts, the allocation/free iteration count, and the allocation size in bytes. The function returns `void`; read the emitted `WARNING: duration = ... ms` message. Context construction is outside the timed loop, while each iteration includes both allocation and free.

### Safety and Interpretation

The implementation does not validate negative or excessive inputs. A large context count consumes memory and stack, a large iteration count monopolizes a backend, and a negative or huge allocation size can become an invalid or enormous allocation request. The function has no privilege checks and receives public execute permission by default, so revoke it from `PUBLIC` to prevent a trivial resource-exhaustion surface.

The timer uses `gettimeofday`, is not monotonic, and the result includes PostgreSQL build, allocator, CPU, scheduler, cache, and concurrent-load effects. Run repeated controlled trials and do not treat a single warning as a production performance guarantee. The pinned source dates from 2014 and uses old memory-context APIs; compile and test it against the exact PostgreSQL source version before use. No preload is required.
