## Usage

Sources:

- [Official README](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/README.md)
- [Extension SQL](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec--1.0.sql)
- [C implementation](https://github.com/cybertec-postgresql/pg_remote_exec/blob/1ac5d03725b86a1bebcf754b2e7f0212a46218c5/pg_remote_exec.c)

pg_remote_exec executes operating-system shell commands from a PostgreSQL backend. Despite its name, commands run on the database server host, not on a separately configured remote host. This is an extremely privileged escape from SQL to the operating system and should normally remain uninstalled.

### Core Workflow

Only a superuser or a member of `pg_execute_server_program` can call the functions. If an audited administrative procedure has no safer alternative, revoke default access and wrap one fixed command rather than accepting user-supplied text:

```sql
CREATE EXTENSION pg_remote_exec;

REVOKE ALL ON FUNCTION pg_remote_exec(text) FROM PUBLIC;
REVOKE ALL ON FUNCTION pg_remote_exec_fetch(text, boolean) FROM PUBLIC;

SELECT pg_remote_exec_fetch('/usr/local/sbin/report-db-health', false);
```

The extension passes the text to the host shell. SQL parameters do not make shell metacharacters safe, and the command's external effects are not transactional.

### Important Objects

- `pg_remote_exec(text)` calls the C library `system()` function and returns its raw integer status.
- `pg_remote_exec_fetch(text, boolean)` calls `popen()`, returns standard output one line per row, and either raises or suppresses a nonzero close status according to the boolean.
- `pg_execute_server_program` is the predefined PostgreSQL role checked at execution time.

### Security and Operational Notes

Granting execution is effectively granting arbitrary commands with the database-server operating-system account. It can read files, alter data outside PostgreSQL, open network connections, exhaust resources, or persist malware. Stderr is not captured by the fetch function, cancellation/resource cleanup is incomplete, and long-running commands occupy a backend. Prefer a separately authenticated job runner with an allowlist, timeouts, resource limits, and an audit log.
