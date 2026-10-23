## Usage

Sources:

- [DocumentDB Extended RUM README](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/README.md)
- [`documentdb_extended_rum` control file](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [Access-method SQL definitions](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--0.106-0.sql)
- [DocumentDB v1.0-0 changelog](https://github.com/documentdb/documentdb/blob/v1.0-0/CHANGELOG.md)

`documentdb_extended_rum` is DocumentDB's extended RUM index access method. It is an implementation component selected by DocumentDB's indexing layer, not a general-purpose application index or a replacement for installing `documentdb`.

### Configure and Install

The library can only be initialized from `shared_preload_libraries`. Preload it after the base DocumentDB libraries and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
documentdb.rum_library_load_option = 'require_documentdb_extended_rum'
```

Then install the extension using the same release as the base stack:

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION documentdb_extended_rum;
```

DocumentDB deployment tooling normally owns this configuration. Existing databases should follow the release-specific upgrade procedure rather than switching an index handler ad hoc.

### Important Objects

- `documentdb_extended_rum` is the index access method registered by the extension.
- `documentdb_extended_rum_catalog` contains BSON operator families and classes used by DocumentDB.
- `documentdb.rum_library_load_option = 'require_documentdb_extended_rum'` directs the DocumentDB index layer to the adapter.
- The implementation is a RUM fork whose on-disk layout and content are designed to remain backward compatible with upstream RUM while changing query and volatile paths for document workloads.

### Operational Boundaries

Install and upgrade this component with matching `documentdb` and `documentdb_core` binaries. Do not build indexes with its internal operator classes directly unless following upstream development guidance; create and manage indexes through the DocumentDB APIs so metadata stays consistent.

In 1.0-0, `documentdb.rum_library_load_option` continues to default to `require_documentdb_extended_rum`. The release fixes concurrent root-split handling in vacuum, restores vacuum delays and interrupt checks, and enables targeted posting-tree pruning and empty entry-leaf pruning by default. Apply the matching upgrade SQL and validate existing indexes through supported DocumentDB operations.
