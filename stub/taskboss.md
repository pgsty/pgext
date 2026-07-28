## Usage

Sources:

- [Official upstream README](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/README.md)
- [Official extension control file (taskboss.control)](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/taskboss.control)
- [Official implementation source](https://github.com/sashaaro/taskboss/blob/fcdc0d35ee6d8a9bffc46d3ba7082e6852242dcf/src/lib.rs)

`taskboss` — A native PostgreSQL job-queue extension written in Rust using pgrx. Inspired by pg-boss. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION taskboss;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `create_queue` is an extension function.
- `delete_queue` is an extension function.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
