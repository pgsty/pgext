## Usage

Sources:

- [Official upstream README](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/README.md)
- [Official extension control file (pg_gen_query.control)](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/pg_gen_query.control)
- [Official extension SQL (pg_gen_query--1.0.sql)](https://github.com/nimit/whatsinmypg/blob/6728d32cfa680120edac83597d7c319dba4f00f8/sql/pg_gen_query--1.0.sql)

`pg_gen_query` — This PostgreSQL extension provides a function, pg_gen_query, that converts natural language input into an equivalent SQL command. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_gen_query;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_gen_query(query text)` is an extension function and returns `TEXT`.
- `regen_schema_cache()` is an extension function and returns `void`.
- `regen_schema_cache_trigger()` is an extension function and returns `event_trigger`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
