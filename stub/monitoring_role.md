## Usage

Sources:

- [Official README](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/README.md)
- [Extension SQL for 0.0.1](https://github.com/frost242/monitoring_role/blob/f76f0629de7a8af4e28b8ca92336a5b050ef4e39/monitoring_role--0.0.1.sql)

`monitoring_role` creates narrowly wrapped monitoring functions that a superuser can grant to an otherwise unprivileged role. The wrappers run with definer rights, so the extension should be installed and delegated only by a trusted administrator.

### Core Workflow

Install it into a dedicated schema, grant a monitoring login access through the helper, and put that schema before `pg_catalog` in the login's search path.

```sql
CREATE ROLE monitor LOGIN;
CREATE SCHEMA monitoring;
CREATE EXTENSION monitoring_role SCHEMA monitoring;

SELECT monitoring.grant_monitor('monitor');
ALTER ROLE monitor SET search_path = monitoring, pg_catalog, public;
```

After reconnecting as the delegated role, unqualified monitoring calls resolve to the wrappers.

```sql
SELECT * FROM pg_stat_activity;
SELECT * FROM pg_ls_dir('.');
SELECT * FROM pg_stat_file('PG_VERSION');
SELECT pg_read_file('PG_VERSION');
```

### Exposed Objects and Boundaries

- `pg_stat_activity()` and the matching view expose the complete activity view.
- `pg_ls_dir(text)` and `pg_stat_file(text)` wrap the corresponding catalog functions.
- Both `pg_read_file` overloads accept only the literal path `PG_VERSION`; other paths raise an exception.
- `grant_monitor(text)` grants schema usage, view selection, and execution on the wrappers.

Creating version 0.0.1 requires superuser rights. Execution is revoked from public by the installation script, but `grant_monitor(text)` itself is not a security-definer function; run it as the extension owner or another role able to issue the grants. Review the exposed directory and activity information against your security policy before delegation.
