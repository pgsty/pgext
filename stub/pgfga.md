## Usage

Sources:

- [Official README](https://github.com/craigpastro/pgfga/blob/main/README.md)
- [SQL-facing implementation](https://github.com/craigpastro/pgfga/blob/main/src/lib.rs)
- [Authorization evaluator](https://github.com/craigpastro/pgfga/blob/main/src/check.rs)
- [Storage implementation](https://github.com/craigpastro/pgfga/blob/main/src/storage.rs)

`pgfga` 0.1.0 is an experimental pgrx implementation of relationship-based access control inspired by Zanzibar. It stores authorization schemas and relationship tuples, then evaluates checks through the fixed pgfga schema. Upstream labels it work in progress with missing validation; do not use it as the sole production authorization boundary.

### Core Workflow

```sql
CREATE EXTENSION pgfga;

SELECT pgfga.create_schema(
  '{"namespaces":{"document":{"relations":{"viewer":[{"namespace":"user"}]},"permissions":{"can_view":{"computedUserset":"viewer"}}},"user":{"relations":{},"permissions":{}}}}'::json
) AS schema_id \gset

SELECT pgfga.create_tuple(
    :'schema_id'::uuid,
    'document', 'policy-1', 'viewer',
    'user', 'anya', ''
);

SELECT pgfga.check(
    :'schema_id'::uuid,
    'document', 'policy-1', 'can_view',
    'user', 'anya', ''
);
```

The psql command stores the UUID returned by schema creation in a variable for the following calls.

### Function Index

- `pgfga.create_schema(json)` stores a namespace/relation/permission model and returns its UUID.
- `pgfga.read_schema(uuid)` and `pgfga.read_schemas()` inspect stored models.
- `pgfga.create_tuple(...)`, `pgfga.read_tuples(...)`, and `pgfga.delete_tuple(...)` manage relationship tuples. Empty filter strings in reads match any value.
- `pgfga.check(...)` evaluates whether a subject has a relation or permission on a resource; the evaluator supports computed usersets, tuple-to-userset traversal, union, intersection, and exclusion.

### Critical Caveats

- Upstream explicitly reports no schema/tuple validation, and its roadmap calls for proper indexes and iterator-based reads. Invalid tuples can therefore be persisted and large checks can be inefficient.
- The extension is superuser-only, non-relocatable, and fixed to the pgfga schema. Protect all mutation functions separately from read/check access.
- Authorization data and application data are not automatically coupled by constraints or a distributed policy transaction. Define lifecycle cleanup, audit, backup, and fail-closed behavior in the application.
- The upstream repository is archived and the README calls the project experimental. Freeze and review the exact source revision if evaluating it.
