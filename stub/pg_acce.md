## Usage

Sources:

- [Upstream README](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/README.md)
- [Version 1.0 install SQL](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce--1.0.sql)
- [Background-worker implementation](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce.c)
- [Internal API declarations](https://github.com/papercompute/pg_acce/blob/4ba51aa93e7d4d568fbf0ec524051f176fd03c80/pg_acce.h)

pg_acce 1.0 is an unfinished 2015 background-worker experiment. Despite its Accelerated Engine description, the published SQL only exposes a logging function and a function that starts a worker to execute supplied SQL text.

### Lock down installation

Do not call acce_setup. If the extension must be inspected, revoke its default public privileges in the installation transaction:

```sql
BEGIN;
CREATE EXTENSION pg_acce;
REVOKE EXECUTE ON FUNCTION acce_setup(text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION acce_info() FROM PUBLIC;
COMMIT;
```

Treat even an isolated inspection build as hostile experimental code.

### Caveats

- acce_setup accepts arbitrary SQL text, starts a dynamic background worker, and leaves execution public by default. The worker copies the caller name but ignores it when initializing its database connection, creating unsafe and potentially elevated privilege semantics.
- The worker executes the supplied text in read-only SPI mode but logs the complete query and every returned field to the server's standard error. Sensitive query text and data can leak into logs.
- The result path assumes at least one row and one column and dereferences the tuple table unconditionally. Empty or nonconforming results can crash the worker.
- The declared worker limit is not enforced. Repeated calls can consume dynamic worker slots, shared memory, and server resources.
- The code performs manual transaction, snapshot, resource-owner, dynamic shared-memory, and background-worker handling through old private PostgreSQL APIs. It has no meaningful tests or current compatibility evidence.
- The README calls the project a template, its test script calls a nonexistent function, and repository licensing is incomplete. Do not deploy this code or expose acce_setup.
