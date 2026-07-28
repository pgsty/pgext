## Usage

Sources:

- [Official upstream README](https://gitlab.com/ottohahn/dbstats/-/blob/master/README.md)
- [Official extension control file](https://gitlab.com/ottohahn/dbstats/-/blob/master/dbstats.control)
- [Official project page](https://gitlab.com/ottohahn/dbstats)

`dbstats` — A comprehensive PL/pgSQL PostgreSQL extension for rapid data analytics, statistical analysis, and data quality assessment directly within your database. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dbstats;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
