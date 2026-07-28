## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/README.md)
- [Official extension control file (json_enhancements_with_hstore.control)](https://api.pgxn.org/src/json_enhancements/json_enhancements-1.0.2/json_enhancements_with_hstore.control)

`json_enhancements_with_hstore` — Json Enhancements for PostgreSQL 9.2 ====================================. Use it when an application needs this specific database capability. The reviewed upstream project is archived or no longer maintained.

### Core Workflow

```sql
CREATE EXTENSION json_enhancements_with_hstore;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `hstore`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
