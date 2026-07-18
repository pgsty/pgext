## Usage

Sources:

- [Upstream README](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/README.md)
- [Extension control file](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority.control)
- [SQL installation script](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority--1.0.sql)
- [C implementation](https://github.com/sshutdownow/pg_setpriority/blob/d36af737087d590ff8d94cecbf5b24e02ffdb66a/pg_setpriority.c)

`pg_setpriority` version `1.0` wraps the operating system's process-priority calls. `setpriority(integer)` changes the nice value of the current PostgreSQL backend, while `getpriority(integer)` reads a process priority and defaults to the current backend.

### Example

```sql
CREATE EXTENSION pg_setpriority;
SELECT setpriority(10);
SELECT getpriority();
```

The change lasts for the backend process, so it can affect later sessions when a pool reuses that backend. Ordinary OS users can generally lower priority but cannot restore it or set a negative nice value without extra capability. `setpriority(integer)` returns the raw success/error code instead of raising, so check it explicitly. The install script leaves execution available by default; revoke `EXECUTE` from application roles and expose only a tightly controlled administrative wrapper.
