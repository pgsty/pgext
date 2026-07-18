## Usage

Sources:

- [Pinned upstream README](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/README.md)
- [Version 1.0 installation SQL](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump--1.0.sql)
- [Pinned memory-context implementation](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump.c)
- [Pinned extension control file](https://github.com/jatobadb/pg_memorydump/blob/558ad20f2b2826607ec56b881bf3db49a79fc167/pg_contextdump.control)

pg_contextdump version 1.0 exposes the current backend's PostgreSQL memory-context tree. Its view reports context names and types, synthetic parent/child IDs, allocation parameters, block counts, total and free bytes, and a block-size histogram.

### Configuration and query

Upstream requires preloading the library and restarting PostgreSQL before creating the extension:

```conf
shared_preload_libraries = 'pg_contextdump'
```

```sql
CREATE EXTENSION pg_contextdump;

SELECT contextname, contexttype, parent_id, totalsize, freespace
FROM pg_contextdump
ORDER BY parent_id, id;
```

The output describes only the backend running the query, not every session in the cluster. Sizes are bytes, and the IDs are generated for that invocation rather than stable memory-context identifiers.

### Compatibility and exposure

The implementation copies private PostgreSQL memory-context structure layouts and traverses the live context tree with casts. Upstream states compatibility with PostgreSQL 11.5 and says newer versions need testing. These internal structures change across major and minor releases, so a mismatched build can return invalid data or crash a backend. Rebuild and validate against the exact server source version.

The install SQL grants SELECT on the view to PUBLIC even though it revokes direct execution of the underlying function. The view still exposes backend-internal names and allocation behavior to ordinary users. Revoke that view privilege and grant it only to an operational diagnostics role. Run it sparingly because collecting the tree allocates and walks memory inside the session being measured.
