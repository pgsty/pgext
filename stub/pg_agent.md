## Usage

Sources:

- [README.md](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/README.md)
- [pg_agent.control](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/pg_agent.control)
- [pg_agent--1.0.sql](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/pg_agent--1.0.sql)
- [Makefile](https://github.com/obartunov/pg_agent/blob/6acd7cdce9ccb9a2a560492cc44394bf430f61f6/Makefile)

`pg_agent` is a SQL-only prototype. Version 1.0 implements a privilege-aware catalog summary and creates policy, annotation, cache and log tables. Statement checking, execution policy enforcement and audit hooks remain proposed features.

### Core workflow

Install the extension, then inspect the catalog as the role whose visibility matters:

```sql
CREATE EXTENSION pg_agent;
SELECT pg_agent.catalog_json(ARRAY['public']::name[]);
SELECT pg_agent.check_statement('SELECT 1');
```

The last call returns `ok=false` with `reason=not_implemented`; it does not validate a statement. `relation_json` returns only a fixed status and the relation name. Do not treat these functions as a SQL sandbox or an authorization decision.

### Catalog and permissions

`catalog_json` runs as the invoker and filters schemas, relations and columns using PostgreSQL privilege checks. Optional arguments control comments, indexes, foreign keys and estimated row counts. Catalog comments are returned as data marked untrusted as instructions. Visibility of a relation is not a guarantee that every row is visible under row security.

The fixed `pg_agent` schema contains `object_annotation`, `role_policy`, `statement_log`, `cache_state` and `catalog_cache`. The installation revokes public access to these tables. An administrator must grant any required schema and object access explicitly. Writing a policy row alone does not enforce it.

### Operational boundaries

The control file is nonrelocatable and sets `superuser=false`; installation still needs permission to create its schema and objects. The SQL-only build has no shared library, preload requirement or restart requirement. Upstream does not publish a PostgreSQL-major support matrix. Keep ordinary role grants, timeouts and read-only transaction controls in place; the prototype is not a replacement for them.
