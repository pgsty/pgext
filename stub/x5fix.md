## Usage

Sources:

- [Official repository](https://github.com/darthunix/x5fix)
- [Extension control file at the catalog source revision](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.control)
- [Version 0.1 extension SQL](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix--0.1.sql)
- [Version 0.1 C implementation](https://github.com/darthunix/x5fix/blob/3c28e84d98ec17a0e2cf0da7be5abb3b571334e1/x5fix.c)

x5fix is an abandoned, invasive repair prototype for legacy Greenplum persistent-filespace internals. It is not a normal PostgreSQL extension: its single SQL function reaches into Greenplum shared memory and persistent catalog storage to alter primary and mirror database identifiers and paths.

### Do Not Invoke on a Live Cluster

The cataloged source depends on private Greenplum headers, locks, structures, relation identifiers, and global variables. Those interfaces are version-specific and unavailable in ordinary PostgreSQL. A binary built against a different Greenplum tree can fail to load, crash a backend, corrupt shared state, or damage persistent filespace metadata.

The function is declared immutable and exposed without a SQL-level privilege guard even though it mutates internal storage. Those declarations are incompatible with its behavior. The C implementation also writes the primary path into both persistent tuple path fields instead of using the distinct mirror-path argument, which is a blocking correctness defect in the pinned source.

### Non-Mutating Audit

For an inherited cluster, only inventory the extension and function definition before planning removal or forensic analysis with a matching legacy Greenplum specialist.

```sql
SELECT extversion
FROM pg_extension
WHERE extname = 'x5fix';

SELECT to_regprocedure(
    'add_entry_gp_persistent_filespace_node(oid,smallint,text,smallint,text)'
);
```

Do not call the returned function as a diagnostic. A successful return would not prove filesystem consistency, mirror correctness, or durable catalog repair.

### Recovery Boundary

Use vendor- and version-specific Greenplum recovery procedures, verified backups, and offline catalog inspection instead of this extension. If an old deployment already contains it, revoke execution from all application roles, preserve the exact server binaries and source revision for forensic purposes, and test any remediation on a disposable copy.

The control file declares no preload requirement, but the function initializes and mutates shared-memory structures at call time. That design is another reason not to port or deploy version 0.1 without a complete expert rewrite and fault-injection test program.
