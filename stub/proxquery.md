## Usage

Sources:

- [Official upstream README](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/README.md)
- [Official extension control file (proxquery.control)](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/proxquery.control)
- [Official implementation source](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/src/lib.rs)

`proxquery` — proxquery is a PostgreSQL extension that adds the @~@ operator with a more flexible term-proximity search syntax on top of tsvector. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION proxquery;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `proxquery_build_query` is an extension function.
- `proxquery_is_analyzer` is an extension function.
- `proxquery_recheck` is an extension function.
- `ts_prox_chain` is an extension function.
- `ts_prox_not_within` is an extension function.
- `ts_prox_positions` is an extension function.
- `ts_prox_positions_prefix` is an extension function.
- `ts_prox_pre` is an extension function.
- `ts_prox_query_exact_cfg_droppable` is an extension function.
- `ts_prox_query_exact_string` is an extension function.
- `ts_prox_query_native_string` is an extension function.
- `ts_prox_query_skeleton` is an extension function.
- `ts_prox_query_support` is an extension function.
- `ts_prox_recheck` is an extension function.

### Requirements and Caveats

- The catalog records version `0.5.2`.
- The control file marks the extension as relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
