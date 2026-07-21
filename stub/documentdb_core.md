## Usage

Sources:

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [`documentdb_core` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/documentdb_core.control)
- [BSON SQL definitions](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_core/sql/udfs/bson_io/bson_io--latest.sql)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb_core` is the low-level BSON type and operator layer used by DocumentDB. It is normally installed as a dependency of `documentdb`; by itself it does not provide collection CRUD, the MongoDB wire protocol, or the gateway.

### Configure and Install

`pg_documentdb_core` must be loaded through `shared_preload_libraries`, followed by a PostgreSQL restart:

```conf
shared_preload_libraries = 'pg_documentdb_core'
```

For a complete single-node stack, the official helper also preloads `pg_cron` and `pg_documentdb`. Install the parent extension in normal deployments:

```sql
CREATE EXTENSION documentdb CASCADE;
```

Direct installation is useful only for low-level BSON work:

```sql
CREATE EXTENSION documentdb_core;
```

The extension is superuser-only and non-relocatable.

### BSON Workflow

```sql
SELECT '{"name":"Ada","score":42}'::documentdb_core.bson;

SELECT documentdb_core.bson_get_value_text(
  '{"name":"Ada","score":42}'::documentdb_core.bson,
  'name'
);
```

Use explicit schema qualification unless `documentdb_core` is in `search_path`.

### Important Objects

- `documentdb_core.bson` stores BSON documents.
- `documentdb_core.bsonquery` represents BSON query values used by the DocumentDB planner and operator layer.
- `documentdb_core.bsonsequence` represents sequences of BSON values.
- `bson_get_value` and `bson_get_value_text`, also exposed through `->` and `->>`, extract a path from a BSON document.
- `bson_from_bytea`, `bson_to_bytea`, `bson_json_to_bson`, and `bson_to_json_string` support serialization boundaries.
- `bson_btree_ops` and `bson_hash_ops` provide comparison and hashing support required by higher layers.

### Operational Boundaries

BSON comparison, indexing, and numeric semantics follow DocumentDB's implementation and should not be assumed to match PostgreSQL `jsonb`. Most objects are infrastructure for `documentdb`; applications seeking collections and MongoDB commands should use the parent extension or gateway rather than building directly on internal types.

Version 0.114-0 keeps `documentdb_core` aligned with the rest of the DocumentDB stack. The upstream changelog does not identify a separate end-user core API migration for this release, so no new standalone workflow is claimed.
