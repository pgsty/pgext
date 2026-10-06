## Usage

Sources:

- [extensions/pg_ortools/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/Cargo.toml)
- [extensions/pg_ortools/src/worker.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/src/worker.rs)
- [crates/pg_bgworker/src/supervision.rs](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/crates/pg_bgworker/src/supervision.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [extensions/pg_ortools/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pgbrew.toml)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/LICENSE)
- [extensions/pg_ortools/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pgbrew.toml)
- [Official pg_ortools README](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/README.md)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/pg_ortools.control)
- [SQL API implementation](https://github.com/matroidbe/pg_extensions-releases/blob/b1987c769e0cc26730ad345bf252e74c93dec069/extensions/pg_ortools/src/lib.rs)

`pg_ortools` version `0.3.2` defines mixed-integer and linear optimization problems in SQL and solves them with HiGHS. Use it for bounded assignment, allocation, feasibility, and objective-optimization workloads that are naturally represented as integer or Boolean variables and linear constraints.

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

### Version 0.3.0 Boundary

This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.

Preload `pg_ortools` and restart for asynchronous jobs. The default build now also includes the Pumpkin CP-SAT scheduling engine; `solve_cp` is time-bounded and cancellable. HiGHS and its SQL workflow remain a separate solver path. The added `eidos_catalog_*` interfaces describe the live optimization catalog.

### Current Release and Upgrade

Extension version 0.3.2 uses `pg_ortools.database` for background workers. Create the extension in that database; a worker now waits instead of repeatedly exiting when it is absent. Changing restart-sensitive worker configuration requires a restart. `pg_ortools.solver_database` remains a deprecated alias. From extension 0.3.0 onward, install matching files and use ALTER EXTENSION UPDATE. Earlier versions still require the migration described above. Withdrawn repository 0.4.0 bottles must be replaced by 0.5.0; upstream bottles do not imply Pigsty package availability.

### Worker Supervision

Version 0.3.2 adds `pgortools.worker_status()` and the superuser-only `pgortools.reset_workers()`. Status reports worker name, state, failures, restarts, PID, time of entry into the current state and last failure. `pg_ortools.max_worker_failures` defaults to 10 consecutive failures; 0 retries indefinitely. Restart delays back off from 5 to 60 seconds, and reaching the limit leaves a failed worker idle until reset.

Install matching libraries, restart PostgreSQL to initialize the new shared memory, and update the extension SQL in the configured worker database. Inspect the failure before resetting; a reset does not repair its cause.

```sql
ALTER EXTENSION pg_ortools UPDATE;
SELECT * FROM pgortools.worker_status();
```
