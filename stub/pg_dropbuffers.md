## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/README.md)
- [Extension control file at the catalog source revision](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.control)
- [Version 0.0.1 extension SQL](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers--0.0.1.sql)
- [Version 0.0.1 C implementation](https://github.com/rjuju/pg_dropbuffers/blob/8328aed1a7ffc3c1a322fa2decf1be7c63024842/pg_dropbuffers.c)

pg_dropbuffers is a destructive benchmarking aid that evicts buffers for the current database and can also request an operating-system page-cache drop. Upstream explicitly limits it to testing and warns about data loss and severe performance effects; do not install or run it on production systems.

### Isolated Test Workflow

Use only on a disposable cluster after stopping application traffic and preserving any data that matters. Run one cache layer at a time so the benchmark records what was changed.

```sql
CREATE EXTENSION pg_dropbuffers;

SELECT pg_drop_current_db_buffers();
SELECT pg_drop_system_cache();
SELECT pg_drop_caches();
```

The first function flushes and drops shared buffers belonging to the current database. The second calls the operating-system command configured in the C source. The wrapper invokes both in that order.

### Operating-System Requirement

Version 0.0.1 is Linux-specific for system-cache eviction: it executes a fixed sudo command for the kernel cache setting. The PostgreSQL service account must be permitted to run exactly that command without a password, or the function raises an error. This is a privileged host-wide action that affects other databases and processes, not just the current session.

### Safety Boundaries

The upstream README warns that dirty buffers might not be safely preserved before removal. Treat every invocation as capable of losing data or destabilizing a test cluster. Never expose these functions to untrusted SQL users; revoke default execution privileges and grant access only to a dedicated benchmark operator if the extension must exist at all.

Cache eviction does not create a generally reproducible cold-cache benchmark by itself: concurrent activity can repopulate caches immediately, and the host-wide cache call disturbs unrelated workloads. Record cluster state, call order, workload timing, and operating-system conditions with every result.

The extension declares no preload requirement, but it uses PostgreSQL internal buffer APIs and a hard-coded host command. Confirm source and binary compatibility with the exact server and operating system before any disposable test run.
