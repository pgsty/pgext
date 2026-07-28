## Usage

Sources:

- [Official upstream README](https://github.com/benizar/pg_sakila_db/blob/85f43e570893f27577ba273b8b9853a2de7438b5/README.md)
- [Official extension control file (pg_sakila_db.control)](https://github.com/benizar/pg_sakila_db/blob/85f43e570893f27577ba273b8b9853a2de7438b5/pg_sakila_db.control)

`pg_sakila_db` — Introduction Postgres xtensions Contributing Getting started Install the extension Dependencies Related projects TODOs. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_sakila_db;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
