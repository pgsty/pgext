## Usage

Sources:

- [Official upstream README](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/README.md)
- [Official extension control file (pgleak.control)](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/pgleak.control)
- [Official extension SQL (pgleak--0.0.1.sql)](https://github.com/eatonphil/pgleak/blob/543555544763d5fe746bdf7869c6d9a90581527d/pgleak--0.0.1.sql)

`pgleak` — Create and start a new Postgres database:. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgleak;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgleak.leak_sum(INT, INT)` is an extension function and returns `INT`.
- `pgleak.leak_sum_malloc(INT, INT)` is an extension function and returns `INT`.
- `pgleak` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
