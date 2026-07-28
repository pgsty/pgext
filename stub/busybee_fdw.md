## Usage

Sources:

- [Official upstream README](https://gitlab.com/byard1/busybee_fdw/-/blob/master/README.md)
- [Official extension control file](https://gitlab.com/byard1/busybee_fdw/-/blob/master/busybee_fdw.control)
- [Official project page](https://gitlab.com/byard1/busybee_fdw)

`busybee_fdw` — A MQTT Foreign Data Wrapper for PostgreSQL. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION busybee_fdw;

CREATE SERVER busybee_server FOREIGN DATA WRAPPER busybee_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
