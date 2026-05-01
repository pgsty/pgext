## Usage

Sources: [README](https://github.com/danolivo/pg_pathcheck/blob/main/README.md), [0.9.1 PG17/18 release](https://github.com/danolivo/pg_pathcheck/releases/tag/v0.9.1_pg17-18), [PGXN metadata](https://api.pgxn.org/dist/pg_pathcheck.json), [source](https://api.pgxn.org/src/pg_pathcheck/pg_pathcheck-0.9.1/pg_pathcheck.c)

`pg_pathcheck` is a PostgreSQL planner diagnostics module that validates reachable planner `Path` trees and reports pointers that look freed, corrupt, or re-used for the wrong relation. It is a preload-only shared library: it registers planner hooks and custom GUCs, but it does not create SQL functions, operators, views, or tables.

### Loading The Module

Build `pg_pathcheck` against the PostgreSQL source line you want to test. The upstream README documents separate long-running branches for PostgreSQL 17/18 and PostgreSQL master/19devel; version `0.9.1` is published for the PG17/18 branch and the PGXN metadata also describes the 0.9.1 source package.

Add the module to `shared_preload_libraries` and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_pathcheck'
```

There is no `CREATE EXTENSION pg_pathcheck` step. After preload, run ordinary SQL, `EXPLAIN`, regression tests, or PostgreSQL test suites; `pg_pathcheck` checks planner paths as queries are planned.

For one temporary cluster:

```bash
initdb -D pgdata
echo "shared_preload_libraries = 'pg_pathcheck'" >> pgdata/postgresql.conf
pg_ctl -D pgdata -l pgdata/logfile start

psql -c 'EXPLAIN SELECT ...'
```

For PostgreSQL test clusters spawned by `make check-world`, use `TEMP_CONFIG`:

```bash
cat > /tmp/pg_pathcheck.conf <<'EOF'
shared_preload_libraries = 'pg_pathcheck'
EOF

TEMP_CONFIG=/tmp/pg_pathcheck.conf make check-world
```

### Configuration

`pg_pathcheck.elevel` controls the severity used when a bad `Path` is detected:

```sql
SET pg_pathcheck.elevel = 'log';
SET pg_pathcheck.elevel = 'warning';  -- default
SET pg_pathcheck.elevel = 'error';
SET pg_pathcheck.elevel = 'panic';
```

Use `warning` or `log` for broad coverage while tests continue, `error` to stop on the first offending query, and `panic` only when a backend crash and core dump are useful for post-mortem debugging.

`pg_pathcheck.stage_checks` enables additional per-stage checks:

```sql
SET pg_pathcheck.stage_checks = off;  -- default
SET pg_pathcheck.stage_checks = on;
```

When enabled, the module checks pathlists at base-rel, join-rel, and upper-rel hook boundaries so a finding can be tied to a narrower planner stage. Leave it off for routine runs because the extra walks add planner overhead, especially for join-heavy queries.

### What Gets Checked

The module walks planner roots, relation pathlists, partial pathlists, cheapest path slots, parameterized paths, subquery subroots, subplan subroots, and nested `Path` fields such as join children, append children, sort subpaths, bitmap paths, and similar compound-path members.

It reports two main classes of problem:

- Invalid `NodeTag`: the pointer no longer looks like a PostgreSQL `Path` node.
- Parent mismatch: a valid-looking `Path` on a base or join relation points at a different `RelOptInfo`, which can indicate same-size memory reuse after a stale path pointer survived.

A typical finding includes the bad tag or mismatch, the containing slot such as `pathlist`, `partial_pathlist`, `cheapest_total_path`, or a nested path field, the relation names that could be resolved, pathlist detail, and the query string available through PostgreSQL debug state.

### Caveats

`pg_pathcheck` is a debug aid for PostgreSQL planner development and extension testing, not a user-facing SQL extension. A PostgreSQL cassert/debug build gives better signal because freed memory is easier to recognize. The upstream README notes coverage differences between the PG17/18 branch and the master branch: PG17/18 checks before later planning stages such as `create_plan` and `setrefs.c`, while master can use the newer planner shutdown hook.
