## Usage

Sources:

- [Official README](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/README.md)
- [Official control file](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/geqo_extension.control)
- [Official module entry point](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/src/geqo_main.c)
- [Official SQL test](https://github.com/parkag/geqo/blob/0f5bea3a679c9719bd8d48d97e1ea84f1d77bbed/test/sql/01_geqo_set_parameters.sql)

`geqo_extension` 0.0.1 is a 2015 experiment that extracts PostgreSQL’s Genetic Query Optimizer into a loadable planner-hook library. It has no extension SQL objects and is used with `LOAD`, not `CREATE EXTENSION`. Its purpose is planner research, not a portable replacement for the server’s built-in GEQO.

### Session Workflow

Load the library in a disposable session, set planner parameters, and compare plans without and with the external hook.

```sql
BEGIN;

LOAD 'geqo_extension';
SET geqo_extension = on;
SET geqo_threshold = 2;
SET geqo_effort = 5;

EXPLAIN (COSTS ON)
SELECT *
FROM a
JOIN b USING (id)
JOIN c USING (id);

ROLLBACK;
```

`LOAD` installs the process-local `join_search_hook`; ending the backend removes it. The user-settable GUC `geqo_extension` selects the extracted optimizer when enabled.

### Planner Behavior

The module copies PostgreSQL optimizer source and uses built-in GEQO parameters such as `geqo_effort`, `geqo_pool_size`, `geqo_generations`, `geqo_selection_bias`, and `geqo_seed`. Plan quality and planning time can vary because the search is genetic and parameter-sensitive.

In the pinned source, the first branch of the hook invokes the external optimizer whenever `geqo_extension` is true. Consequently, the later `geqo_threshold` branch is unreachable in that state; do not assume the threshold protects small joins. Confirm behavior from `EXPLAIN` for every tested query shape.

### Compatibility and Operational Boundaries

Planner hooks and copied optimizer internals are tightly coupled to the PostgreSQL source version. This code predates modern PostgreSQL releases and has no maintained compatibility matrix. Compile only against the matching server headers and run planner/regression tests before loading it.

Only one `join_search_hook` chain is available, and ordering with other planner extensions matters. A bad plan or module bug affects queries in the loaded backend. Use isolated benchmarks, keep a control session, and never infer production improvement from a small synthetic join.
