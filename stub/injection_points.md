## Usage

Sources:

- [Official upstream README](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/README.md)
- [Official extension control file (injection_points.control)](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/crates/contrib/injection_points/extension/injection_points.control)
- [Official extension SQL (injection_points--1.0.sql)](https://github.com/malisper/pgrust/blob/ed269002a1730e18446e716d2e9abbd0a4f00c82/crates/contrib/injection_points/extension/injection_points--1.0.sql)

`injection_points` — Postgres rewritten in Rust, now passing 100% of the Postgres regression tests. Use it when an application needs this specific database capability. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION injection_points;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `injection_points_attach(IN point_name TEXT, IN action text)` is an extension function and returns `void`.
- `injection_points_cached(IN point_name TEXT, IN arg TEXT DEFAULT NULL)` is an extension function and returns `void`.
- `injection_points_detach(IN point_name TEXT)` is an extension function and returns `void`.
- `injection_points_load(IN point_name TEXT)` is an extension function and returns `void`.
- `injection_points_run(IN point_name TEXT, IN arg TEXT DEFAULT NULL)` is an extension function and returns `void`.
- `injection_points_set_local()` is an extension function and returns `void`.
- `injection_points_stats_drop()` is an extension function and returns `void`.
- `injection_points_stats_fixed(OUT numattach int8, OUT numdetach int8, OUT numrun int8, OUT numcached int8, OUT numloaded int8)` is an extension function and returns `record`.
- `injection_points_stats_numcalls(IN point_name TEXT)` is an extension function and returns `bigint`.
- `injection_points_wakeup(IN point_name TEXT)` is an extension function and returns `void`.
- `removable_cutoff(rel regclass)` is an extension function and returns `xid8`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
