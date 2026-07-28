## Usage

Sources:

- [Official database.dev package page](https://database.dev/olirice/asciiplot)

`olirice-asciiplot` — A Toy ASCII Plotting Library. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "olirice-asciiplot";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `scatter_internal(state scatter_state)` is an extension function and returns `TEXT`.
- `scatter_sfunc(state scatter_state, x numeric, y numeric, title TEXT, height INTEGER, width INTEGER)` is an extension function and returns `scatter_state`.
- `scatter` is an aggregate exposed by the extension.
- `scatter_state` is an extension-defined type.

### Requirements and Caveats

- The catalog records version `0.0.1`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
