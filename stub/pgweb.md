## Usage

Sources:

- [Official upstream README](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/README.md)
- [Official extension control file (pgweb.control)](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/pgweb.control)
- [Official extension SQL (pgweb--0.0.1.sql)](https://github.com/eatonphil/pgweb/blob/0944c98ac7b09741830920db413d40cf633538f2/pgweb--0.0.1.sql)

`pgweb` — Create a new Postgres database and run the test script to define a web service and start the server:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgweb;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `handle_hello_world(params JSON)` is an extension function and returns `TEXT`.
- `pgweb.register_get(TEXT, TEXT)` is an extension function and returns `VOID`.
- `pgweb.serve(TEXT, INT)` is an extension function and returns `VOID`.
- `pgweb` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
