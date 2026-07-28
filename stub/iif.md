## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/iif/iif-0.0.1/README.md)
- [Official extension control file (iif.control)](https://api.pgxn.org/src/iif/iif-0.0.1/iif.control)
- [Official extension SQL (iif--0.0.1.sql)](https://api.pgxn.org/src/iif/iif-0.0.1/iif--0.0.1.sql)

`iif` — A sample extension for adding a function _iif_ to Postgres. Use it when porting or emulating the corresponding database API. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION iif;

SELECT iif(1<0,1,2);
 iif
-----
   2
(1 row)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `iif(boolean, anyelement, anyelement)` is an extension function and returns `anyelement`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
