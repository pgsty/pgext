## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/check_updates/check_updates-1.0.1/README.md)
- [Official extension control file (check_updates.control)](https://api.pgxn.org/src/check_updates/check_updates-1.0.1/check_updates.control)

`check_updates` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION check_updates;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.1`.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
