## Usage

Sources:

- [Official extension control file (cat_tools.control)](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/cat_tools.control)
- [Official extension SQL (cat_tools--0.1.0--0.1.3.sql)](https://api.pgxn.org/src/cat_tools/cat_tools-0.2.1/sql/cat_tools--0.1.0--0.1.3.sql)

`cat_tools` — Tools for interfacing with the catalog. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION cat_tools;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `__cat_tools.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `__cat_tools.exec(sql text)` is an extension function and returns `void`.
- `pg_temp.create_function(function_name text , args text , options text , body text , grants text DEFAULT NULL)` is an extension function and returns `void`.
- `pg_temp.exec(sql text)` is an extension function and returns `void`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
