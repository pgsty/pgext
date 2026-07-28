## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/Readme.md)
- [Official extension control file (mvtbl.control)](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/mvtbl.control)
- [Official extension SQL (mvtbl--0.0.1.sql)](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/mvtbl--0.0.1.sql)

`mvtbl` — A postgres Extension to easily move tables around tablespaces. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION mvtbl;

SELECT pg_size_pretty(mvtbl('test','mvtbl_test_tblspace'));
 pg_size_pretty
----------------
 123 MB
(1 row)

SELECT pg_size_pretty(mvtbl('public.test','pg_default'));
 pg_size_pretty
----------------
 123 MB
(1 row)
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mvtbl(tbl text, tblspace text)` is an extension function and returns `bigint`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.2`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
