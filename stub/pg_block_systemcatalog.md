## Usage

Sources:

- [Official README](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/README.md)
- [Extension control file](https://github.com/nuko-yokohama/pg_block_systemcatalog/blob/0849fa12fd6049e422b37c2912ae56f10f6f1f22/pg_block_systemcatalog.control)

`pg_block_systemcatalog` rejects ordinary users' queries that reference PostgreSQL system catalogs, statistics views, or `information_schema`. Superusers remain exempt, and one configured role can be used as an allow-list role for monitoring accounts.

### Core Workflow

Preload the library and optionally name one group role in `postgresql.conf`, then restart PostgreSQL and create the extension.

```conf
shared_preload_libraries = 'pg_block_systemcatalog'
pg_block_systemcatalog.allow_role = 'monitor_catalog'
```

```sql
CREATE EXTENSION pg_block_systemcatalog;

CREATE ROLE monitor_catalog;
CREATE ROLE monitor LOGIN;
GRANT monitor_catalog, pg_monitor TO monitor;
```

`pg_block_systemcatalog.allow_role` accepts one role. Grant that role to each account that needs catalog access; if it is empty, only privileged users are allowed. Separately revoke access to application objects from monitoring accounts when duties must remain segregated.

### Security Boundary

This is a query-reference guard, not a complete metadata sandbox. Upstream demonstrates that a scalar call to a C statistics function such as `pg_stat_get_activity()` may still return information even when selecting that function in `FROM` is blocked. Other C functions can also read catalogs internally without exposing the catalog relation to the query tree.

The only documented validation environment is PostgreSQL 10 on CentOS 7. The module inspects planned or parsed query structures and can affect administration tools, drivers, ORMs, monitoring, migrations, and extension code that consult catalogs. Test every required workflow and every hook-using extension on the exact server major; do not treat membership in the allow-list role as a substitute for normal PostgreSQL privileges.
