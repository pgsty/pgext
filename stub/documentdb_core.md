## Usage

Sources:

- [DocumentDB v1.0-0 README](https://github.com/documentdb/documentdb/blob/v1.0-0/README.md)
- [`documentdb_core` control file](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_core/documentdb_core.control)
- [BSON SQL definitions](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_core/sql/udfs/bson_io/bson_io--latest.sql)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v1.0-0/scripts/preload_libraries.sh)
- [DocumentDB 1.2 release](https://github.com/documentdb/documentdb/releases/tag/v1.2-0)
- [1.2 control](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb_core/documentdb_core.control)
- [1.1 to 1.2 SQL](https://github.com/documentdb/documentdb/blob/v1.2-0/pg_documentdb_core/sql/documentdb_core--1.1-0--1.2-0.sql)
- [Official upgrade procedure](https://github.com/documentdb/documentdb/blob/v1.2-0/packaging/README.md)

`documentdb_core` is the low-level BSON type and operator layer used by DocumentDB. It is normally installed as a dependency of `documentdb`; by itself it does not provide collection CRUD, the MongoDB wire protocol, or the gateway.

### Configure and Install

`pg_documentdb_core` must be loaded through `shared_preload_libraries`, followed by a PostgreSQL restart:

```ini
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

Use SQL extension version 1.0-0 with matching DocumentDB components. This core layer continues to supply BSON infrastructure; collection and gateway workflows belong to the parent extension.

### DocumentDB 1.2 changes

The 1.2 source declares SQL version `1.2-0` for `documentdb_core`, with the supplied 1.0 → 1.1 → 1.2 update chain. Upgrade it together with the API and index components using the official package or container procedure, after backing up data; a downgrade of updated data to 1.0 is unsupported. `bson_build_document` now accepts `timestamptz` as a BSON date at millisecond precision. Collation remains opt-in through `documentdb_core.enableCollation = on`; explicit `simple` collation is also rejected while that setting is off. Application tests should check comparison-sensitive queries after enabling it.
