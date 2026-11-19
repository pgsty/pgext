## Usage

Sources:

- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/README.md)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/fdw/steampipe_postgres_fdw.control)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/fdw/steampipe_postgres_fdw--1.0.sql)
- [Official source](https://github.com/turbot/steampipe-postgres-fdw/blob/v2.3.1/CHANGELOG.md)

steampipe_postgres_fdw exposes Steampipe plugins as PostgreSQL foreign tables. Distribution 2.3.1 retains extension SQL version 1.0. The core FDW connects to plugin services over gRPC; plugin-specific standalone extensions are separate builds.

### Enable and connect

```sql
CREATE EXTENSION steampipe_postgres_fdw;
CREATE SERVER fdw_aws
  FOREIGN DATA WRAPPER steampipe_postgres_fdw
  OPTIONS (wrapper 'aws');
```

The SQL file names the wrapper steampipe_postgres_fdw; older architecture examples abbreviate it as fdw. Configure a compatible plugin service and its credentials before importing or declaring provider tables. The core library requires administrator installation but no preload; match the server ABI and the selected deployment's documented plugin setup.

### Version 2.3.0

This release fixes crashes and incorrect results for runtime parameter quals, including LATERAL references and subqueries. A scan interrupted by cancellation, statement_timeout or a plugin error now fails instead of treating a partial row set as complete. Invalid/duplicate connection names or malformed configuration files are skipped with warnings; inspect those warnings rather than assuming every configured connection loaded.

The standalone FDW is marked parallel safe, and GOGC/GOMEMLIMIT are read from the PostgreSQL server environment. These are process settings, not SQL GUCs. The legacy cache_off command is corrected. Replace matching libraries and reconnect/restart processes that loaded old code; extversion 1.0 alone cannot identify the runtime fix. Remote APIs still impose their own permissions, pagination, consistency and rate limits.

### Distribution 2.3.1

The 2.3.1 patch backports Darwin cross-compilation and image-publishing changes. It retains SQL version 1.0 and adds no new SQL workflow; the 2.3.0 runtime fixes above still apply.
