## Usage

Sources:

- [Official README](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/README.md)
- [Official Rust package manifest](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/Cargo.toml)
- [Official extension control file](https://github.com/Smertos/stopgap/blob/e076ae165ba35c6f2d484acd943451373376e03e/crates/stopgap/stopgap.control)

`stopgap` is a versioned deployment and invocation layer for TypeScript/JavaScript function bundles executed by the companion `plts` procedural language. It keeps environments, deployment history, activation, diffing, and rollback in PostgreSQL while application code is addressed by a path such as `api.coolApi.myFn`.

### Install and Invoke

Both extensions are untrusted and require superuser installation. Install `plts` first, then `stopgap`:

```sql
CREATE EXTENSION plts;
CREATE EXTENSION stopgap;

SELECT stopgap.call_fn(
  'api.coolApi.myFn',
  '{"id": 1}'::jsonb
);
```

The current upstream direction keeps named TypeScript exports under a project-local `stopgap/` tree and recommends the CLI for deployment validation and ergonomics:

```bash
stopgap deploy --db "$STOPGAP_DB" --env prod --from-schema app --label v1.0
stopgap status --db "$STOPGAP_DB" --env prod
stopgap diff --db "$STOPGAP_DB" --env prod --from-schema app
stopgap rollback --db "$STOPGAP_DB" --env prod --steps 1
```

### Runtime and Management Surface

- `stopgap.call_fn(path, args)` resolves a deployed named export and returns non-null JavaScript values as `jsonb`; JavaScript `undefined` and `null` become SQL `NULL`.
- `stopgap.status`, `stopgap.deployments`, `stopgap.diff`, and `stopgap.rollback` inspect and manage environment state.
- `stopgap.query` handlers run with a read-only database context and reject `db.exec`; `stopgap.mutation` and regular `plts` handlers receive a read-write context.
- Database operations initiated through the handler context run in the same PostgreSQL transaction as the calling SQL statement.

### Security and Change Boundaries

- TypeScript/JavaScript runs inside the database trust boundary. Review bundle provenance, dependency locking, SQL privileges, schema access, resource limits, error handling, and observability before executing application-supplied code.
- Keep deploy, activation, and rollback privileges separate from ordinary invocation. Treat an environment change as a database release, with review, backups, smoke tests, and a tested rollback target.
- A function can hold locks and consume backend CPU or memory for the duration of the calling transaction. Set timeouts and keep network or long-running work out of hot transactions.
- The reviewed README says the product was being course-corrected in March 2026 and labels several SQL signatures as target shapes during migration. Pin version `0.1.3`, verify the installed signatures, and do not automate against roadmap text alone.
- Compiler and runtime semantics are shared with the installed `plts` version. Test TypeScript type checking, module resolution, serialization, transaction rollback, and upgrades as one compatibility unit.
