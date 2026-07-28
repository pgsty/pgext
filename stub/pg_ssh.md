## Usage

Sources:

- [Official upstream README](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/README.md)
- [Official extension control file (pg_ssh.control)](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/pg_ssh.control)
- [Official implementation source](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/src/lib.rs)

`pg_ssh` — PostgreSQL **18** (13–17 also build via the pgXX features) Rust toolchain (stable) [cargo-pgrx][cpgrx] **0.19.1** — must match the pgrx crate version exactly System packages (Debian/Ubuntu):. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_ssh;

SELECT convert_from(stdout, 'UTF8') AS stdout, exit_code
  FROM ssh.exec('web-1', 'uname -a; uptime');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `exec` is an extension function.
- `keygen` is an extension function.
- `session_close` is an extension function.
- `session_exec` is an extension function.
- `session_open` is an extension function.
- `sessions()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.3.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
