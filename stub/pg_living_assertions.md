## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/README.md)
- [pg_living_assertions.control](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/CHANGELOG.md)
- [pg_living_assertions--0.5.9--0.5.10.sql](https://api.pgxn.org/src/pg_living_assertions/pg_living_assertions-0.5.10/pg_living_assertions--0.5.9--0.5.10.sql)

`pg_living_assertions` 0.5.10 stores named SQL assertions, results and replacement history. Checks run on demand; they are not constraints evaluated on every write. This pure SQL extension needs no preload.

### Core Workflow

```sql
CREATE EXTENSION pg_living_assertions;
SELECT living_assertions.declare(
  'simple_check', 'one equals one',
  $$SELECT 1 = 1 AS holds, 'arithmetic check'::text AS detail$$);
SELECT living_assertions.run('simple_check');
SELECT name, state, age FROM living_assertions.status;
```

### Operational Boundaries

A check returns one boolean `holds` and optional text `detail`. `living_assertions.run_all()` evaluates the registry; `living_assertions.state()`, `living_assertions.status` and `living_assertions.stale()` distinguish failures, unknown/erroring results and stale or absent checks. `living_assertions.declare_unchanged()` fingerprints a canonicalized expression; a value disappearing to NULL is broken. Checks run as their recorded author in a read-only, rolled-back subtransaction with the author search path and temporary schemas last. The caller must be allowed to `SET ROLE` to that author. Session advisory locks acquired during evaluation are released, but external effects and temporary-sequence changes are not a sandboxed resource. Only trusted authors should register SQL. Registration and checks are server-dated, tables reject TRUNCATE, and retirement/replacement is restricted to the author or an authorized role. Upgrade through the complete matching SQL chain; old assertions retain their existing fingerprint algorithm. Current releases use Apache-2.0; releases through 0.5.3 retain their original PostgreSQL License.

### Role Isolation in 0.5.10

Version 0.5.10 evaluates a check in a temporary `SECURITY DEFINER` function owned by its recorded author when role isolation is needed. PostgreSQL then refuses attempts to reset the role or session authorization inside that frame; the frame is created and rolled back with the read-only check. The author needs `TEMP` on the database when a frame is required, and the caller still needs authority to act as that author. A `SECURITY DEFINER` caller can now run another author's checks under that authority. This closes the role-reset escape in 0.5.8/0.5.9; it does not make external effects safe.
