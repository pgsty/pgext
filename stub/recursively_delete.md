## Usage

Sources:

- [Pinned upstream README](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/README.md)
- [Version 0.1.5 implementation](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/sql/recursively_delete.sql)
- [Pinned PGXN metadata](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/META.json)
- [Pinned extension control file](https://github.com/trlorenz/PG-recursively_delete/blob/ae9f1f17b51e2307a97c1fe0f38662906cee0053/recursively_delete.control)

recursively_delete version 0.1.5 is a destructive administration function that discovers foreign-key dependents and builds one recursive data-modifying CTE. It handles some circular dependencies and composite primary keys and can print an ASCII damage preview. The upstream author explicitly requires backups and says it is not intended for transactional workloads.

### Preview and rollback a trial

Preview is the default. Run a real trial in an explicit transaction and inspect all affected tables before choosing to commit in a separate, deliberate run:

```sql
CREATE EXTENSION recursively_delete;

SELECT recursively_delete(
  'public.users'::regclass,
  4402,
  false
);

BEGIN;

SELECT recursively_delete(
  'public.users'::regclass,
  4402,
  true
);

ROLLBACK;
```

The return value counts only root rows explicitly deleted, not every dependent row. Text scalar values require an explicit text cast. Composite-key arrays must list values in primary-key index order and all array elements must share one PostgreSQL type.

### Preview side effects and graph limits

Preview mode does not merely calculate SELECT counts. It executes the actual DELETE CTE inside a PL/pgSQL subtransaction and deliberately raises and catches an exception to roll it back. It can still take locks, consume work, advance sequences, write logs, and invoke DELETE triggers whose external side effects are not transactional. Run previews on a restored copy first, with lock and statement timeouts.

Relationships using ON DELETE SET NULL or SET DEFAULT are not recursively deleted. Graph discovery depends on primary keys on involved child tables. Partitioning, inherited tables, row-level security, deferred constraints, statement and row triggers, generated columns, unusual key types, concurrent writes, deep or wide graphs, and quoted identifiers need explicit tests. The final query can become extremely large and hold many row and table locks.

The function runs with caller privileges, but PUBLIC receives execution by default. Revoke it and grant only a dedicated maintenance role that already has the exact DELETE privileges required. The helper functions and view use unqualified names and no fixed search_path; install in a controlled schema and use a trusted search path to prevent name capture.

Upstream reports use only on PostgreSQL 10.10 and 13.1, and its PGXN metadata marks the release unstable. Version 0.1.5 is a preview-grade emergency tool, not a referential-action engine. Prefer explicit, reviewed deletion procedures and declarative foreign-key actions; if this function is unavoidable, take and verify a backup, stop concurrent writers, preserve the preview, and validate row counts and integrity afterward.
