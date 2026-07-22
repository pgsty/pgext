## Usage

Sources:

- [Official README](https://github.com/extio1/pg_isolate/blob/main/README.md)
- [Postmaster hook implementation](https://github.com/extio1/pg_isolate/blob/main/src/main.c)
- [Configuration parser](https://github.com/extio1/pg_isolate/blob/main/src/config/group_config.c)
- [Setuid helper implementation](https://github.com/extio1/pg_isolate/blob/main/src/cgroup_initialize.c)

`pg_isolate` 1.0 is a Linux-only preload prototype that moves PostgreSQL backends into cgroup v2 directories according to database and user rules. It installs no SQL extension objects. The implementation requires a setuid-root helper and contains unsafe, host-specific code paths; do not deploy it on a production server without a full security redesign.

### Intended Configuration

The module reads a relative `pg_isolate.json` file and selects the first matching database/user group, with an optional default:

```json
{
  "reporting": {
    "database": "analytics",
    "user": "reporter"
  },
  "default": {}
}
```

It is loaded at postmaster start:

```conf
shared_preload_libraries = 'pg_isolate'
```

There is no `CREATE EXTENSION` step. The upstream build target compiles `cgroup_initialize` into PostgreSQL's binary directory, changes its owner to root, and sets the setuid bit before the server starts.

### Effective Behavior

At startup the helper creates a hard-coded `/sys/fs/cgroup/postgres` hierarchy, enables the memory and CPU controllers, and moves the postmaster into a per-instance subtree. The authentication hook then moves a matching backend and logs connection metadata. Although the JSON parser accepts `memory_hard`, `memory_soft`, and `cpu_weight`, the inspected code only creates directories and does not apply those resource values.

### Critical Caveats

The parser copies unbounded JSON strings into small fixed buffers. The helper writes a debug file under a developer-specific home directory, and the background-worker hook contains a hard-coded `test_module` branch that unshares a mount namespace and attempts `pivot_root` into `/home/extio1/alpine`. Configuration lookup is relative to the postmaster working directory despite the README's PGDATA wording. These behaviors, the setuid helper, version-sensitive hooks, cgroup v2 assumptions, and verbose authentication logging make the current source unsuitable for production isolation.
