## Usage

Sources:

- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/README.md)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier.control)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier--1.0.sql)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/pgext/pgtier.c)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/docs/BUILD.md)
- [Official source](https://github.com/deepanshu406/scx_pgtier/blob/901be082f1c71d7b72c36868445ded8640a62c4c/LICENSE)

`pgtier` `1.0` moves PostgreSQL backends between Linux cgroups and changes their scheduling policy for the external scx_pgtier scheduler. It is an early experimental project tested with PostgreSQL 16 in a virtual machine.

### Workflow

An administrator must first prepare cgroup v2 directories, grant the PostgreSQL OS user the required write access, and start the postmaster inside `/sys/fs/cgroup/ufs-pg/main`. The separate scheduler requires a Linux kernel with sched_ext and BTF and is started with root privileges. Preparing the cgroup tree can change the root CPU-controller setting.

Load the module in new sessions and configure existing application roles:

```sql
ALTER SYSTEM SET session_preload_libraries = 'pgtier';
SELECT pg_reload_conf();
ALTER ROLE oltp_app SET pgtier.tier = 'ts';
ALTER ROLE report_user SET pgtier.tier = 'bg';
```

Preserve any existing entries when configuring `session_preload_libraries`. Reconnect after the reload. The tier mapping works through session loading; `CREATE EXTENSION` supplies the optional diagnostic function in a database:

```sql
CREATE EXTENSION pgtier;
SELECT pgtier_hello();
```

### Configuration and boundaries

`pgtier.tier` accepts `none`, `ts` and `bg`; `pgtier.cgroup_root` defaults to `/sys/fs/cgroup/ufs-pg`. Both settings are superuser-only. A superuser can delegate tier changes with `GRANT SET ON PARAMETER pgtier.tier`. Installation of the diagnostic function also requires a superuser.

The diagnostic reports the actual tier, scheduling policy and cgroup. A failed cgroup move or policy change logs a warning and does not stop the query, so do not treat a configured tier as proof of OS enforcement. Changes can leave the backend at an intermediate step.

Time-sensitive work gains priority at the expense of background work. The scheduler's optional lock-holder polling requires a separate database connection with monitoring access; it does not cover spinlocks or LWLocks. No general latency or production-safety guarantee is established. Rehearse cgroup ownership, scheduler shutdown and server restarts in a disposable VM before use.
