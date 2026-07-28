## Usage

Sources:

- [Official upstream README](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/README.md)
- [Official extension control file (pg_ac.control)](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/pg_ac.control)
- [Official extension SQL (pg_ac--0.1.sql)](https://github.com/enerzed/pg_ac/blob/4f618296d20de7d26f638abc3fc5603a20bffc3d/pg_ac--0.1.sql)

`pg_ac` — PostgreSQL 12 or later C compiler with C99 support pg_config available in PATH. Use it for the corresponding text-search, parsing, or linguistic workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_ac;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ac_add(bigint, text)` is an extension function and returns `boolean`.
- `ac_build(text[])` is an extension function and returns `bigint`.
- `ac_build(tsvector)` is an extension function and returns `bigint`.
- `ac_deserialize(bytea)` is an extension function and returns `bigint`.
- `ac_destroy(bigint)` is an extension function and returns `boolean`.
- `ac_fini()` is an extension function and returns `boolean`.
- `ac_init()` is an extension function and returns `boolean`.
- `ac_match(bigint, text)` is an extension function and returns `integer[]`.
- `ac_rank_simple(bigint, text)` is an extension function and returns `real`.
- `ac_remove(bigint, text)` is an extension function and returns `boolean`.
- `ac_search(bigint, text)` is an extension function and returns `boolean`.
- `ac_search(bigint, tsquery)` is an extension function and returns `boolean`.
- `ac_serialize(bigint)` is an extension function and returns `bytea`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
