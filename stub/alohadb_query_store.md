## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_query_store.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_query_store/alohadb_query_store.control)
- [Official extension SQL (alohadb_query_store--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_query_store/alohadb_query_store--1.0.sql)

`alohadb_query_store` — AlohaDB Query Store and Index Advisor. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_query_store;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `autovacuum_suggestions()` is an extension function and returns `TABLE`.
- `index_advisor_recommend()` is an extension function and returns `TABLE`.
- `index_advisor_unused_indexes()` is an extension function and returns `TABLE`.
- `query_store_entries()` is an extension function and returns `TABLE`.
- `query_store_reset()` is an extension function and returns `void`.
- `query_store_stats()` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
