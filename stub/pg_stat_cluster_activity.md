## Usage

Sources:

- [Official upstream README](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/README)
- [Official extension control file (pg_stat_cluster_activity.control)](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/pg_stat_cluster_activity/pg_stat_cluster_activity.control)
- [Official extension SQL (pg_stat_cluster_activity--1.0.sql)](https://github.com/andy31002/opentenbase/blob/ff795be78c8583b4129fa9b745597c4fa2e122c8/contrib/pg_stat_cluster_activity/pg_stat_cluster_activity--1.0.sql)

`pg_stat_cluster_activity` — track execution statistics in whole cluster scope. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_stat_cluster_activity;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_cancel_session(text)` is an extension function.
- `pg_signal_session(text, integer, bool)` is an extension function.
- `pg_terminate_session(text)` is an extension function.
- `pg_stat_cluster_activity` is an extension-defined view.
- `pg_stat_cluster_activity_cn` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
