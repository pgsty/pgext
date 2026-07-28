## Usage

Sources:

- [Official upstream README](https://github.com/gabrosys/pg_purr/blob/5064eda546b0b36c9d3ff1877b02b708b18b99b9/README.md)
- [Official extension control file (pg_purr.control)](https://github.com/gabrosys/pg_purr/blob/5064eda546b0b36c9d3ff1877b02b708b18b99b9/pg_purr.control)

`pg_purr` — A PostgreSQL extension (via PL/Python) that brings quantum computing to your database:. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_purr;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.0`.
- Install the confirmed extension dependencies first: `plpython3u`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
