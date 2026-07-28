## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/konfigraf/konfigraf-0.0.2/konfigraf.control)
- [Official project page](https://pgxn.org/dist/konfigraf/)

`konfigraf` — Git based application configuration. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION konfigraf;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
