## Usage

Sources:

- [Official extension control file (pgespresso.control)](https://api.pgxn.org/src/pgespresso/pgespresso-1.2.0/pgespresso.control)
- [Official extension SQL (pgespresso--1.2.sql)](https://api.pgxn.org/src/pgespresso/pgespresso-1.2.0/pgespresso--1.2.sql)

`pgespresso` — Optional Extension for Barman, Backup and Recovery Manager for PostgreSQL. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgespresso;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgespresso_abort_backup()` is an extension function and returns `VOID`.
- `pgespresso_start_backup(label TEXT, fast BOOL)` is an extension function and returns `TEXT`.
- `pgespresso_stop_backup(label_content TEXT)` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
