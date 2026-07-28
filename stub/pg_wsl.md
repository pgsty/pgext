## Usage

Sources:

- [Official upstream README](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/README.md)
- [Official extension control file (pg_wsl.control)](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/pg_wsl.control)
- [Official extension SQL (pg_wsl--1.0.sql)](https://github.com/pierreforstmann/pg_wsl/blob/5444e39f6cf9adc378107b4f70e1d5c1000ad8fc/pg_wsl--1.0.sql)

`pg_wsl` — PostgreSQL extension to write from primary to standby log. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_wsl;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
