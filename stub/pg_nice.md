## Usage

Sources:

- [Extension control file](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice.control)
- [Version 0.0.1 extension SQL](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/pg_nice--0.0.1.sql)
- [C implementation](https://github.com/roa/pg_nice/blob/9394ae17ae7b577daba2f44b146a69b8ca41e100/src/nice.c)

`pg_nice` 0.0.1 lowers the Linux I/O scheduling priority of the current PostgreSQL backend by executing `ionice -c 3 -p <backend-pid>`. It does not change CPU nice level, other sessions, or the whole PostgreSQL instance.

### Core Workflow

```sql
CREATE EXTENSION pg_nice;

SELECT pg_backend_pid(), pg_nice();
```

`pg_nice() RETURNS integer` is declared `VOLATILE STRICT` and returns the exit status from the C library's `system()` call. A return value of zero normally means the external command succeeded; nonzero values are raw process status rather than a PostgreSQL error or a portable diagnostic code.

### Requirements and Effects

The database server must run on Linux, have an `ionice` executable discoverable in the PostgreSQL service account's `PATH`, and permit that account to change the current process to idle I/O class 3. The change applies to the backend process for the remainder of its life, so it can outlive the transaction or request that invoked the function.

Because one backend serves one database connection, connection-pool reuse can give later work the lowered I/O priority. Open a dedicated session for low-priority work if that distinction matters.

### Caveats

The implementation invokes a shell synchronously from inside PostgreSQL. Although the command contains only the backend's numeric PID, it still blocks the backend while spawning an external process and depends on operating-system behavior outside PostgreSQL. It is not portable to macOS, BSD, or Windows and has no fallback when `ionice` is absent.

Test the returned status and actual scheduler state on the deployment kernel. Restrict installation and execution to administrators who understand the service account and connection-pool consequences; use workload management outside this small extension when durable, observable resource governance is required.
