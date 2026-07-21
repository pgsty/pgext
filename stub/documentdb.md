## Usage

Sources:

- [DocumentDB v0.114-0 README](https://github.com/documentdb/documentdb/blob/v0.114-0/README.md)
- [DocumentDB v0.114-0 changelog](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)
- [`documentdb` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb/documentdb.control)
- [Official preload helper](https://github.com/documentdb/documentdb/blob/v0.114-0/scripts/preload_libraries.sh)

`documentdb` is the public PostgreSQL API extension for DocumentDB, an open-source MongoDB-compatible document database built on PostgreSQL. It stores BSON documents and implements CRUD, aggregation, full-text, geospatial, and vector workflows. MongoDB drivers require the separate DocumentDB gateway; installing this extension alone exposes the PostgreSQL API, not a wire-protocol listener.

### Configure and Install

The official deployment helper preloads the core and API libraries with `pg_cron`. Restart PostgreSQL after changing this setting:

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb'
```

Install the public extension and its declared dependencies:

```sql
CREATE EXTENSION documentdb CASCADE;
```

`CASCADE` can install `documentdb_core`, `pg_cron`, `tsm_system_rows`, `vector`, and `postgis` when their files are present. Installation is superuser-only and non-relocatable.

### Native SQL Workflow

The SQL surface uses a database name, collection name, and BSON command document:

```sql
SELECT documentdb_api.create_collection('appdb', 'people');

SELECT documentdb_api.insert_one(
  'appdb',
  'people',
  '{"_id": 1, "name": "Ada", "team": "storage"}',
  NULL
);

SELECT document
FROM documentdb_api_catalog.bson_aggregation_find(
  'appdb',
  '{"find":"people","filter":{"team":"storage"}}'
);
```

For application compatibility, run the gateway and use a supported MongoDB driver against its configured TLS endpoint. The gateway translates wire-protocol commands into this PostgreSQL API.

### Important Objects

- `documentdb_api` contains collection-management and command functions such as `create_collection` and `insert_one`.
- `documentdb_api_catalog.bson_aggregation_find` executes a MongoDB-style find specification and returns BSON documents.
- `documentdb_core.bson` is the storage and interchange type supplied by `documentdb_core`.
- DocumentDB roles and internal schemas separate public read/write operations from administrative and implementation objects.
- `documentdb.enableNonBlockingUniqueIndexBuild` controls the v0.114 path for background unique ordered-index builds and is enabled by default in that release.

### Version and Operational Notes

The v0.114-0 tagged changelog enables schema validation by default, fixes validator propagation and caching, and enables non-blocking unique ordered-index builds. It also records gateway configuration, connectivity-check, TLS, and credential-handling improvements. Two RUM optimizations in that changelog remain feature-flagged and disabled by default; do not describe them as active behavior.

MongoDB compatibility is not identical to every MongoDB server version. Test operators, index behavior, transactions, schema validation, authentication, and driver behavior used by the application. Match `documentdb`, `documentdb_core`, gateway, and optional distributed/index components to the same release line.
