## Usage

Sources:

- [Official README](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/README.md)
- [Version 1.0 SQL API](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp--1.0.sql)
- [Filter and GUC implementation](https://github.com/CrunchyData/pgseccomp/blob/3ed6d0780dc5d9a5204e57ce7628544a3610e218/pgseccomp.c)

`pgseccomp` applies Linux seccomp syscall filters to the PostgreSQL postmaster and client backends. It can allow, audit-log, return an error for, or kill a process on selected syscalls, with a cluster-wide base policy and stricter session overlays.

### Core Workflow

The library requires Linux kernel 3.5 or later and libseccomp 2.4 or later. Add it to `shared_preload_libraries`, start with logging or a permissive default, restart PostgreSQL, and create the extension before inspecting the effective policy.

```sql
CREATE EXTENSION pgseccomp;
SELECT * FROM seccomp_filter() ORDER BY context, syscallnum;
```

Cluster settings use `pgseccomp.global_syscall_allow`, `pgseccomp.global_syscall_log`, `pgseccomp.global_syscall_error`, `pgseccomp.global_syscall_kill`, and `pgseccomp.global_syscall_default`. Corresponding `pgseccomp.session_syscall_*` settings overlay client backends; `pgseccomp.session_roles` selects role-specific lists.

`set_client_filter(default_action, allow_list, log_list, error_list, kill_list)` can apply a same-or-more-restrictive overlay once in a session. NULL list arguments use configured session lists, while `*` inherits the global list for that action. Execution on both SQL functions is revoked from `PUBLIC` by the installation script.

### Safety and Recovery

An incorrect rule can prevent PostgreSQL from opening files, allocating memory, communicating, forking workers, or shutting down cleanly; the `kill` action sends `SIGSYS`. Build an observed syscall inventory for the exact PostgreSQL build and workload, introduce `log` before `error` or `kill`, retain console-level recovery access, and test startup, authentication, backup, replication, extensions, maintenance, and failover.

Seccomp policies can only become more restrictive after loading; a session cannot use the SQL function to relax the inherited filter. Audit events are normally read from the operating system's audit log, which database administrators may not be permitted to access. Coordinate policy and diagnostics with the host security team rather than treating this extension as a complete sandbox.
