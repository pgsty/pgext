## Usage

Sources:

- [Official upstream README](https://gitlab.com/daamien/pgrx-tuto/-/blob/main/README.md)
- [Official extension control file](https://gitlab.com/daamien/pgrx-tuto/-/blob/main/pg_zxcvbn/pg_zxcvbn.control)
- [Official project page](https://gitlab.com/daamien/pgrx-tuto)

`pg_zxcvbn` — Ce dépot contient une série d'extensions PostgreSQL developpées avec PGRX. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_zxcvbn;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
