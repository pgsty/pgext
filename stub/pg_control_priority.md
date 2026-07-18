## Usage

Sources:

- [Upstream README](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/README.md)
- [Version 1.0 install SQL](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/pg_control_priority--1.0.sql)
- [Priority-control implementation](https://github.com/MasaoFujii/pg_control_priority/blob/68bde19357c6eea0250aed40ed426b5a0b3c1d2f/pg_control_priority.c)

pg_control_priority 1.0 reads or raises the POSIX nice value of a PostgreSQL process. A larger numeric nice value means less CPU scheduling priority; this does not control query, lock, storage, or network priority.

### Deprioritize the current test session

```sql
CREATE EXTENSION pg_control_priority;
SELECT pg_get_priority(pg_backend_pid());
SELECT pg_set_priority(pg_backend_pid(), 5);
SELECT pg_get_priority(pg_backend_pid());
```

Run this in a disposable session because the backend remains deprioritized until it exits. The install SQL revokes public execution of both functions, so the calls require superuser access or an explicit grant.

### Startup configuration

To apply the scheduling-priority setting consistently, preload the library and restart:

```conf
shared_preload_libraries = 'pg_control_priority'
pg_control_priority.scheduling_priority = 5
```

### Caveats

- A PostgreSQL process can normally lower its own CPU priority but cannot raise it again without operating-system privilege. On common Unix systems the effective nice range ends at 19 even though the setting accepts 20.
- Supplying another PostgreSQL process identifier can affect client backends, auxiliary processes, or the postmaster when operating-system permissions allow it. Keep execution tightly restricted.
- The code uses PostgreSQL process checks and Unix scheduling APIs. It is not portable to Windows and must be built and tested for the exact PostgreSQL and operating-system target.
