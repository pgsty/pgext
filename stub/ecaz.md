## Usage

Sources:

- [Official upstream README](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/README.md)
- [Official extension control file (ecaz.control)](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/ecaz.control)
- [Official extension SQL (ecaz--0.1.0--0.1.1.sql)](https://github.com/agent-ix/ecaz/blob/f4bf945395d199297eb5a233e0da806dd0489f29/ecaz--0.1.0--0.1.1.sql)

`ecaz` — Ecaz is a rust based PostgreSQL extension for performant, highly scalable vector storage and retrieval. It supports a broad range of quantization and index options rather than a single fixed architecture. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION ecaz;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ec_distann_active_head_policy(index_regclass regclass)` is an extension function and returns `TABLE`.
- `ec_distann_build_epoch_with_training(index_regclass regclass, epoch bigint, build_id uuid, training_relation regclass)` is an extension function and returns `bytea`.
- `ec_spire_coordinator_index_shape_fingerprint(index_oid regclass)` is an extension function and returns `text`.
- `ec_spire_coordinator_insert_shape_fingerprint(table_oid regclass)` is an extension function and returns `text`.
- `ec_spire_register_placement_batch(index_oid oid, entries ec_spire_placement_entry[])` is an extension function and returns `bigint`.
- `ec_spire_remote_catalog_drop_index_cleanup_event()` is an extension function and returns `event_trigger`.
- `ec_spire_remote_index_shape_fingerprint(index_oid regclass)` is an extension function and returns `text`.
- `ecvector(ecvector, integer, boolean)` is an extension function and returns `ecvector`.
- `ecvector_from_bytea(bytea, integer, boolean)` is an extension function and returns `ecvector`.
- `ecvector_from_real_array(real[], integer, boolean)` is an extension function and returns `ecvector`.
- `ecvector_in(cstring, oid, integer)` is an extension function and returns `ecvector`.
- `ecvector_inner_product(ecvector, ecvector)` is an extension function and returns `float4`.
- `ecvector_negative_inner_product(ecvector, ecvector)` is an extension function and returns `float4`.
- `ecvector_negative_query_inner_product(ecvector, real[])` is an extension function and returns `float4`.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
