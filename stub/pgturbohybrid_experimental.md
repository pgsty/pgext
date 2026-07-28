## Usage

Sources:

- [Official upstream README](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/README.md)
- [Official extension control file (pgturbohybrid_experimental.control)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/pgturbohybrid_experimental.control)
- [Official extension SQL (pgturbohybrid_experimental--0.2.0.sql)](https://github.com/mayflower/pgturbohybrid/blob/cd68155df20545b163dd610fb95e8aa7e62fc108/sql/pgturbohybrid_experimental--0.2.0.sql)

`pgturbohybrid_experimental` — This README helps you understand what pgturbohybrid does, when hybrid search is useful, how to install it, how to create your first index, and how to check whether the fast path is working. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pgturbohybrid_experimental;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `turbohybrid_experimental_compact_code_score(query_codes pg_catalog.int2[], doc_codes pg_catalog.int2[], experimental pg_catalog.bool DEFAULT false, force_kernel pg_catalog.text DEFAULT 'auto')` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector(vector[])` is an extension function and returns `turbohybrid_multivector`.
- `turbohybrid_multivector_context_count(turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_context_maxsim(query turbohybrid_multivector, doc turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_context_offsets(turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_count(turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_dims(turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_distance(turbohybrid_multivector, turbohybrid_query)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_field_ids(turbohybrid_multivector)` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_field_weighted_maxsim(query turbohybrid_multivector, doc turbohybrid_multivector, field_ids pg_catalog.int4[], weights pg_catalog.float4[])` is an extension function and returns `pg_catalog`.
- `turbohybrid_multivector_from_contexts(raw_values pg_catalog.float4[], dim pg_catalog.int4, context_offsets pg_catalog.int4[])` is an extension function and returns `turbohybrid_multivector`.
- `turbohybrid_multivector_from_contexts_and_fields(raw_values pg_catalog.float4[], dim pg_catalog.int4, context_offsets pg_catalog.int4[], field_ids pg_catalog.int4[])` is an extension function and returns `turbohybrid_multivector`.
- `turbohybrid_multivector_from_float4(raw_values pg_catalog.float4[], dim pg_catalog.int4)` is an extension function and returns `turbohybrid_multivector`.
- `turbohybrid_multivector_in(pg_catalog.cstring)` is an extension function and returns `turbohybrid_multivector`.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.0`.
- Install the confirmed extension dependencies first: `pgturbohybrid`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
