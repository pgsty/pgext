## Usage

Sources:

- [Upstream README](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/README.md)
- [Version 1.0 install SQL](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/pgreadfile--1.0.sql)
- [File I/O implementation](https://github.com/AbdulYadi/pgreadfile/blob/280468c96e641b47591c581e2729864df8aa0de1/pgreadfile.c)

pgreadfile 1.0 exposes unrestricted server-side file reading and writing through pgreadfile(text) and pgwritefile(text, bytea). The published extension is unsafe for a multi-user or production database.

### Revoke access during installation

```sql
BEGIN;
CREATE EXTENSION pgreadfile;
REVOKE EXECUTE ON FUNCTION pgreadfile(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION pgwritefile(text, bytea) FROM PUBLIC;
COMMIT;
```

Do not grant the write function to application roles. If a reviewed administrative workflow must read a fixed approved file, grant only the read function to a tightly controlled role and validate the path outside this function.

### Caveats

- The install SQL leaves both functions executable by public. Any database user can otherwise read every server file visible to the PostgreSQL operating-system account and overwrite or truncate every writable file.
- Both functions are falsely declared immutable even though they depend on and mutate external state. Planner constant folding or expression reuse can execute them at surprising times; the write declaration is especially hazardous.
- The read function loads the entire file into backend memory without a size limit. A large file can exhaust memory.
- File-open failures and some write failures are not reported reliably; short writes are ignored. Symbolic links and path races are not controlled.
- Upstream documents compilation only against PostgreSQL 11 and provides no tests, permission model, or current compatibility evidence. Prefer maintained core facilities with predefined roles and explicit path controls.
