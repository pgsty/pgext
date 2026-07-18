## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/README.rst)
- [Version 3.5 SQL objects](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/pgq_node--3.5.sql)
- [Extension control file](https://github.com/pgq/pgq-node/blob/f06df0f82f026b4067e2d0252dd02ff60f193c2a/pgq_node.control)

`pgq_node` is SQL support infrastructure for cascaded PGQ queueing. It requires `pgq`, installs as a superuser, is fixed in `pg_catalog`, and is intended to be managed by the surrounding PGQ/Londiste tooling rather than called as an isolated utility library.

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_node;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN ('pgq', 'pgq_node');
```

Use the exact documentation and daemon versions for the complete cascaded queue topology, node registration, worker state, failover, and resynchronization workflow. Do not invent node metadata with ad hoc updates; queue position and topology tables are coordinated protocol state.

The reviewed README only says “support code for cascaded queueing,” and the catalog coverage ends at PostgreSQL 14. Audit all installed `pg_catalog` objects, security-definer functions, grants, and search paths. Test duplicate delivery, ordering, consumer lag, node loss, split brain, retention, backup/restore, and version skew. Use maintained logical replication or a currently supported PGQ stack for new systems.
