## Usage

Sources:

- [Official upstream README](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/README.md)
- [Official extension control file (pg_replica.control)](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/pg_replica.control)
- [Official implementation source](https://github.com/hyperiondb/hyperiondb/blob/b0ea5e901de4a5f417721981b3ecbfece1664c3a/src/lib.rs)

`pg_replica` — A PostgreSQL extension that gives a small cluster of **vanilla Postgres** nodes **automatic, consensus-driven failover** — full-cluster replication (tables, roles, DDL, *everything*) with a **built-in Raft group** and **no external dependencies**: no etcd, no Consul, no Kubernetes. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_replica;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `rotate_credential` is an extension function.
- `status()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.7.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
