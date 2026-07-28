## Usage

Sources:

- [Official upstream README](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/README.md)
- [Official extension control file (pg_documentdb_gw_host.control)](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/pg_documentdb_gw_host/pg_documentdb_gw_host.control)
- [Official implementation source](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/pg_documentdb_gw_host/src/lib.rs)

`pg_documentdb_gw_host` — pg documentdb gw host: Created by pgrx. Use it when porting or emulating the corresponding database API. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_documentdb_gw_host;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- Install the confirmed extension dependencies first: `documentdb`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
