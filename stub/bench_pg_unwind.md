## Usage

Sources:

- [Extension control file](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/bench_pg_unwind.control)
- [SQL API](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/sql/bench_pg_unwind.sql)
- [Rust benchmark implementation](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/src/lib.rs)
- [C boundary implementation](https://github.com/JLockerman/bench_pg_unwind/blob/ff7c71b7c8935810521d7e58cd11d9643c41f1c3/src/bench.c)

bench_pg_unwind 0.1 is a narrow 2020 microbenchmark for comparing repeated Rust-to-C calls with and without a PostgreSQL error-unwind guard. It is diagnostic code, not an application extension.

### Disposable benchmark

Use small, positive iteration counts on an isolated server:

```sql
CREATE EXTENSION bench_pg_unwind;
SELECT bench_cross_fn(1000, 0, 1);
SELECT bench_try_fn(1000, 0, 1);
```

Both functions return the repeated addition result. Each call also emits a WARNING containing total and per-iteration timing, so collect server and client messages as the benchmark output.

### Caveats

- The iteration count directly controls a tight backend CPU loop and has no upper bound. Never expose it to untrusted input.
- Zero or negative iterations make the timing code divide a duration by a nonpositive value and can enter a Rust panic path.
- Addition occurs in signed 64-bit C arithmetic without overflow checks. Choose start, increment, and iteration values whose result remains in range.
- Results include SQL-call, logging, clock, compiler, and build effects; they are not a general measure of PostgreSQL or Rust performance.
- The build depends on a historical Timescale Rust support branch and 2020-era PostgreSQL headers. Upstream provides no tests, compatibility matrix, README, or complete repository license files.
