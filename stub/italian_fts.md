## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/README.rst)
- [Official extension control file (italian_fts.control)](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/italian_fts.control)
- [Official extension SQL (italian_fts.sql)](https://api.pgxn.org/src/italian_fts/italian_fts-1.2.1/italian_fts.sql)

`italian_fts` — This package can be used to install and configure the ISpell dictionary in PostgreSQL 8.3 and later. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION italian_fts;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
