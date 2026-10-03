## Usage

Sources:

- [README.md](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/README.md)
- [pg_living_assertions.control](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions.control)
- [pg_living_assertions--0.4.1--0.5.0.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.4.1--0.5.0.sql)
- [pg_living_assertions--0.5.0--0.5.1.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/pg_living_assertions--0.5.0--0.5.1.sql)
- [test/sql/read_only.sql](https://github.com/Manuelreyesbravo/pg_living_assertions/blob/7bf075c6de879b05ec0ed8d135304ece83519317/test/sql/read_only.sql)

`pg_living_assertions` 0.5.1 stores SQL checks, verdicts, verification times and replacement history. These are checks run on demand, not SQL ASSERTION constraints evaluated on every write. It is a pure SQL extension with no preload requirement.

### Register and Verify

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### Results and History

A check must return exactly one row with boolean `holds` and optional text `detail`. `living_assertions.run_all()` evaluates registered checks. `living_assertions.state()` distinguishes holds, broken, unknown, erroring, unchecked, retired and unregistered; `living_assertions.stale()` separates never-checked assertions from old results. `living_assertions.declare_unchanged()` records an expression for later text comparison, so the author must canonicalize its output. Definitions are superseded with a reason; results and registry data are included in database dumps.

### Execution and Privileges

Since 0.5.0 the evaluator runs read-only inside a subtransaction that is always rolled back, preserving its verdict. This repairs the older STABLE-only evaluator, which did not stop side effects through volatile functions. It is not a sandbox for untrusted SQL: temporary-sequence changes, session advisory locks and external effects can survive. Only trusted administrators should register checks; they execute with the privileges of the later caller. Registry tables belong to the extension owner and write functions are revoked from `PUBLIC` by default.

### Upgrade

After installing the matching files, use `ALTER EXTENSION pg_living_assertions UPDATE TO '0.5.1'`. The 0.4.1→0.5.0→0.5.1 chain replaces evaluator functions without changing registry tables. The final patch qualifies row types in `run()` so type-cache invalidation does not resolve them under an assertion's unrelated search path.

Fresh installation also uses an earlier base SQL script followed by the packaged upgrade chain; keep the complete set of matching scripts installed.
