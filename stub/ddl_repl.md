## Usage

Sources:

- [Official README](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/README.md)
- [Extension control file](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.control)
- [Extension SQL](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl--1.0.sql)
- [Hook implementation](https://github.com/realyota/pg_repl/blob/7fe3bbe7b1e23f050ca12400849c4d452b469f7a/ddl_repl.c)

`ddl_repl` is an experimental ProcessUtility hook that forwards utility statements to configured PostgreSQL nodes. It is intended to keep DDL and DCL similar across instances, but it is synchronous text replay rather than logical replication or an atomic distributed transaction.

### Core Workflow

Load the library in every backend that should originate replay; server-wide preload is the practical configuration. Restart after changing `shared_preload_libraries`, install the catalog objects, and register one-way destinations with DSNs that the PostgreSQL server can use.

```conf
shared_preload_libraries = 'ddl_repl'
```

```sql
CREATE EXTENSION ddl_repl;

SELECT ddl_repl.create_node(
  'reporting',
  'host=reporting-db dbname=app user=ddl_repl password=secret',
  'one-way DDL destination'
);

SELECT ddl_repl.set_node_active('reporting', false);
SELECT ddl_repl.drop_node('reporting');
```

`ddl_repl.nodes` stores `node_name`, `dsn`, `description`, `active`, and `creation_date`. `ddl_repl.enabled` controls replay in the current backend; `ddl_repl.only_repl_users` exists in the source and defaults to off. The management surface is `ddl_repl.create_node`, `ddl_repl.set_node_active`, and `ddl_repl.drop_node`.

### Operational Boundaries

The hook executes locally first, then opens libpq connections, starts remote transactions, replays the original query text, and commits each active node. A later connection or commit failure can occur after an earlier node committed, so recovery is manual and the nodes can diverge. Avoid reciprocal node definitions, which can replay the same statement back through another loaded hook.

The filter excludes only a small group of utility statements such as `SET`, `SHOW`, `PREPARE`, `EXECUTE`, and `DEALLOCATE`; it is not a narrow DDL allowlist. DSNs are stored as plain catalog text, and the extension SQL does not harden default privileges. Revoke public access to the schema, table, and management functions, use a least-privilege remote role, protect credentials, and test every supported PostgreSQL major before use. The pinned upstream source dates from 2018 and targets old server internals; treat `ddl_repl` as a lab prototype, not a production consistency mechanism.
