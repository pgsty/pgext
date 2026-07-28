## Usage

Sources:

- [Official upstream README](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/README.md)
- [Official extension control file (postgraph.control)](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/postgraph.control)
- [Official implementation source](https://github.com/postgraphdb/postgraph/blob/72a080d3367aea6e0ffffaed54a9d0b025caee32/src/backend/postgraph.c)

`postgraph` — PostGraph is a multi-model, graph centric query engine build on Postgres. PostGraph is designed to work fast with your OLTP, OLAP and AI Applications. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION postgraph;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
