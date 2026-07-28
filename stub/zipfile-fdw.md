## Usage

Sources:

- [Official upstream README](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/README.md)
- [Official extension control file (zipfile-fdw.control)](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/zipfile-fdw.control)
- [Official implementation source](https://github.com/beargiles/zipfile-fdw/blob/dfdfd5f3be8724ef338d7caa72ba1ecb4d4a5754/src/zipfile-fdw.c)

`zipfile-fdw` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "zipfile-fdw";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
