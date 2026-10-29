## Usage

Sources:

- [Official documentation](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/docs/index.md)
- [Official README](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/README.md)
- [Control file](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/pgedge_safesession.control)
- [Hook and GUC implementation](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/src/pgedge_safesession.c)
- [Optional SQL registration](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/sql/pgedge_safesession--1.0.sql)
- [PostgreSQL CI matrix](https://github.com/pgEdge/pgedge-safesession/blob/b406cd1e82cf21f30959351691c56de046562678/.github/workflows/ci.yml)

`pgedge_safesession` uses SQL version 1.0 and enforces read-only sessions for configured PostgreSQL roles. This documentation describes the pinned source revision, including fixes made after the 1.0 tag. Upstream requires PostgreSQL 14 or later and CI covers 14–18. It supplements ordinary least-privilege grants; superuser logins are exempt.

### Preload and Configure

The hooks must load at server startup:

```ini
shared_preload_libraries = 'pgedge_safesession'
```

Restart PostgreSQL after installing the matching library. As a superuser, configure the restricted roles in the server configuration, rather than only in one session:

```sql
ALTER SYSTEM SET pgedge_safesession.roles =
  'readonly_user,reporting_role';
SELECT pg_reload_conf();
```

The module works without SQL registration; optional registration adds its catalog entry and the status helper `pgedge_safesession_is_restricted()`. Restrictions anchor on `session_user` and actual role grants. `SET ROLE`, temporary privilege elevation and a superuser-owned `SECURITY DEFINER` function do not let an already restricted login escape. Changes that make a session restricted invalidate cached plans.

### Protection Controls

All settings are `SUSET`. `pgedge_safesession.block_dml`, `block_ddl` and `block_c_functions` default on. `block_all_c_functions` defaults off; enabling it also blocks stable and immutable native extension functions. Read-only transaction state is enforced unconditionally on each restricted statement; there is no `force_read_only` setting in this source.

Default protection blocks writes, DDL, COPY FROM, COPY TO PROGRAM, strong table locks, protected read-only GUC changes, side-effecting volatile functions and untrusted procedural calls. Ordinary SELECT, non-writing EXPLAIN ANALYZE, safe transaction/configuration operations and pooler resets remain usable. Plain EXPLAIN plans without executing the statement. Read-only functions in trusted languages can run, while writes within them are intercepted. A conservative allow-list exempts only reviewed harmless volatile built-ins.

### Security Boundaries

This does not filter rows, mask readable data, authenticate users or constrain exempt superusers. IMMUTABLE/STABLE function labels remain a trust boundary unless all native functions are blocked. Direct `fmgr` calls made by a type's input/output functions or an index operator class can bypass expression checks; restricted roles cannot create such objects, so review pre-existing objects and their privileges. Interactions with other hook extensions, including Spock, are not exhaustively tested. Validate the exact server build, pooling, prepared statements, role changes and failover before relying on the policy.
