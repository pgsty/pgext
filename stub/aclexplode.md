## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/README.md)
- [Official extension control file (aclexplode.control)](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/aclexplode.control)
- [Official extension SQL (aclexplode.sql)](https://api.pgxn.org/src/aclexplode/aclexplode-1.0.3/sql/aclexplode.sql)

`aclexplode` — Be sure that you have pg_config installed and in your path. If you used a package management system such as RPM to install PostgreSQL, be sure that the -devel package is also installed. If necessary tell the build process where to find it:. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION aclexplode;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `aclexplode(aclitem[], OUT grantor oid, OUT grantee oid, OUT privilege_type text, OUT is_grantable bool)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.3`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
