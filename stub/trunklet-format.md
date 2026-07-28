## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/README.md)
- [Official extension control file (trunklet-format.control)](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/trunklet-format.control)
- [Official extension SQL (trunklet-format.sql)](https://api.pgxn.org/src/trunklet-format/trunklet-format-0.2.0/sql/trunklet-format.sql)

`trunklet-format` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it for the corresponding SQL or database utility workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "trunklet-format";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.0`.
- Install the confirmed extension dependencies first: `trunklet`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
