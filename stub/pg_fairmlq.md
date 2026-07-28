## Usage

Sources:

- [Official upstream README](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/README.md)
- [Official extension control file (pg_fairmlq.control)](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/pg_fairmlq.control)
- [Official extension SQL (pg_fairmlq--0.1.0.sql)](https://github.com/sam-harri/pg_fairmlq/blob/a1df8e9f4bc51282f87e4144973716c33826c136/sql/pg_fairmlq--0.1.0.sql)

`pg_fairmlq` — The tests require the pgtap extension to run. Use it when an application needs this specific database capability. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_fairmlq;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `plpgsql`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
