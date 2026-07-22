## Usage

Sources:

- [Official README](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/README.md)
- [Benchmark orchestration SQL](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/FUNCTIONS/measure.sql)
- [Query-timing SQL](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/FUNCTIONS/time_query.sql)
- [C implementation](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.c)
- [Extension control file](https://github.com/joelonsql/pg-timeit/blob/d83ab65bc9ae7b919bd99de1ebf3e2177bdea607/timeit.control)

`timeit` benchmarks PostgreSQL built-in internal C functions with high-resolution wall time and, on supported x86/Linux systems, CPU cycles. It repeatedly calls the named internal symbol, doubles the iteration count, and uses linear regression to estimate per-call cost. This is an invasive diagnostic tool for an isolated test server, not an application timing API.

### Benchmark a Built-In Function

Restrict the schema before allowing any other role to connect, then benchmark a known-safe built-in symbol with text forms of its arguments:

```sql
CREATE EXTENSION timeit;

REVOKE ALL ON SCHEMA timeit FROM PUBLIC;
REVOKE EXECUTE ON ALL FUNCTIONS IN SCHEMA timeit FROM PUBLIC;
GRANT USAGE ON SCHEMA timeit TO performance_lab;
GRANT EXECUTE ON FUNCTION timeit.t(text, text[], integer, double precision, integer, interval, integer)
    TO performance_lab;

SELECT timeit.t('numeric_add', ARRAY['1.5', '2.5']);
SELECT timeit.t('pg_sleep', ARRAY['0.01'], 2, 0.99, 10, interval '1 second');
```

The default measurement starts with one invocation and doubles until the recent sample fits the requested regression threshold or the timeout is observed. The timeout is checked between completed batches; it cannot interrupt one slow or stuck invocation, and the total number of calls can grow quickly.

### Main Objects

- `timeit.t(...)`: return estimated wall time per call as human-readable text.
- `timeit.c(...)`: return estimated CPU cycles per call as `bigint`.
- `timeit.measure(...)`: return sample arrays, R-squared, slope, intercept, and final iteration count; `measure_type` is `time` or `cycles`.
- `timeit.eval(text, text[])`: invoke the internal function once and return its result as text.
- `timeit.time_query(text)`: plan and execute a read-only SQL statement and report planning/execution time plus Linux performance counters.
- `timeit.pretty_time`, `timeit.round_to_sig_figs`, and `timeit.compute_regression_metrics`: formatting and regression helpers.

The `core_id` argument defaults to `-1`, letting the operating system schedule the backend. Selecting a core uses Linux CPU-affinity calls. Cycle measurements require a compatible CPU, while `timeit.time_query` requires Linux `perf_event_open` access permitted by the host kernel and container policy.

### Safety and Measurement Caveats

The C code resolves a built-in symbol with `fmgr_internal_function`, converts input text using the function's declared argument types, and calls it directly through PostgreSQL's function manager. It does not perform normal SQL name resolution or an explicit `EXECUTE` privilege check on the selected built-in. By default, extension functions are executable by `PUBLIC`; revoke that access and grant only narrowly selected wrappers to a trusted laboratory role.

Do not pass arbitrary function names. Built-ins can be volatile, sleep, allocate heavily, change state, expect special call context, use pseudo-types, or fail when invoked directly. Repetition magnifies every side effect and can block or crash the backend. `timeit.time_query` asks SPI for a read-only execution, but a read-only `SELECT` can still call volatile functions; accept only reviewed queries.

Measurements include cache state, CPU frequency, interrupts, backend allocation, and instrumentation overhead. Pinning a core changes backend affinity, perf counters may be unavailable or multiplexed, and concurrent activity distorts the regression. Use a disposable database, warm up deliberately, repeat runs, inspect the returned fit, and treat results as comparative observations on that exact host rather than universal performance guarantees.
