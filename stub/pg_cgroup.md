## Usage

Sources:

- [Official README](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/README.md)
- [Extension control file](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.control)
- [Extension SQL](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup--1.0.sql)
- [libcgroup implementation](https://github.com/MasahikoSawada/pg_cgroup/blob/09214a99f5acfced8697ff3c72b0b05dc6bc7bb2/pg_cgroup.c)

`pg_cgroup` exposes Linux libcgroup operations as SQL functions so a PostgreSQL backend can create child resource groups, set controller values, and attach backend processes. It targets cgroup v1-era Linux and was documented with CentOS 7.4, PostgreSQL 10, and libcgroup 0.41.

### Core Workflow

An operating-system administrator must first create a parent cgroup owned by the PostgreSQL OS account. Then preload the module, set that parent, restart, and install the SQL objects.

```conf
shared_preload_libraries = 'pg_cgroup'
pg_cgroup.cgroup_name = 'pgsql'
```

```sql
CREATE EXTENSION pg_cgroup;

REVOKE ALL ON FUNCTION create_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION delete_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION set_resource_value(text, text, text, bigint) FROM PUBLIC;
REVOKE ALL ON FUNCTION attach_resource_group(text, text) FROM PUBLIC;
REVOKE ALL ON FUNCTION attach_resource_group(text, text, integer) FROM PUBLIC;

SELECT create_resource_group('memory', 'reporting');
SELECT set_resource_value('memory', 'reporting', 'memory.limit_in_bytes', 1073741824);
SELECT attach_resource_group('memory', 'reporting');
```

`create_resource_group` and `delete_resource_group` manage a child below the configured parent. `set_resource_value` writes a named controller parameter. `attach_resource_group` attaches the current backend, or its three-argument overload attaches a supplied operating-system PID.

### Privilege and Platform Caveats

These functions run with the PostgreSQL server account's libcgroup privileges and contain no SQL superuser check. Default public execution would let database users alter resource controls or attempt to move other accessible PIDs, so revoke every overload from `PUBLIC` and grant only to a tightly controlled administrative role. Validate controller and parameter names in an allowlisted wrapper; negative `bigint` inputs are converted to unsigned values by the C code.

Several delete and attach return codes are ignored, so a returned `true` does not prove that the kernel operation succeeded. Verify the actual cgroup filesystem after every change. The pinned source predates unified cgroup v2 and modern PostgreSQL APIs; it should not be assumed compatible with cgroup v2, containers, systemd delegation, or current server majors. Test only on a matching isolated host and prefer operating-system resource managers for production enforcement.
