## Usage

Sources:

- [Official pg_ortools README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/README.md)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/pg_ortools.control)
- [SQL API implementation](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_ortools/src/lib.rs)

`pg_ortools` version `0.2.0` defines mixed-integer and linear optimization problems in SQL and solves them with HiGHS. Use it for bounded assignment, allocation, feasibility, and objective-optimization workloads that are naturally represented as integer or Boolean variables and linear constraints.

### Core Workflow

```sql
CREATE EXTENSION pg_ortools;

SELECT pgortools.create_problem('example');
SELECT pgortools.add_int_var('example', 'x', 0, 100);
SELECT pgortools.add_int_var('example', 'y', 0, 100);
SELECT pgortools.add_constraint('example', 'x + y <= 150');
SELECT pgortools.maximize('example', '2*x + 3*y');

SELECT pgortools.solve_sync('example');
SELECT pgortools.get_solution('example');
```

For asynchronous work, `solve` returns a job ID; use `solve_status`, `cancel_solve`, `LISTEN pgortools_solve`, and `get_solution` to manage it. `solve_greedy` returns the first feasible solution, while `solve_local`, `solve_with_strategy`, and `solve_auto` expose alternative search strategies. The declarative `solve_assignment` and `parse_assignment` functions build common table-driven assignment models.

### Operational Notes

The extension stores problems, variables, constraints, jobs, and solutions in `pgortools`. Constraint expressions support linear comparisons, but the official README explicitly excludes `!=`. The default solver time limit is controlled by `pg_ortools.solver_time_limit`; worker enablement, polling interval, and worker database are server settings described as restart-sensitive upstream. Confirm that the background worker is running before relying on async jobs. The control file is non-relocatable but not superuser-only, and the install SQL grants public DML on its schema tables; review privileges before multi-tenant use.
