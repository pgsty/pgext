## Usage

Sources:

- [Official upstream README](https://gitlab.com/3manuek/dummy_fdw/-/blob/master/README.md)
- [Official extension control file](https://gitlab.com/3manuek/dummy_fdw/-/blob/master/dummy_data.control)
- [Official project page](https://gitlab.com/3manuek/dummy_fdw)

`dummy_data` — A readable, null foreign data wrapper for Postgresql 9.3+. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dummy_data;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
