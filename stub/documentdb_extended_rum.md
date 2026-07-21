## Usage

Sources:

- [DocumentDB Extended RUM README](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/README.md)
- [`documentdb_extended_rum` control file](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)
- [Access-method SQL definitions](https://github.com/documentdb/documentdb/blob/v0.114-0/pg_documentdb_extended_rum/sql/documentdb_extended_rum--0.106-0.sql)
- [DocumentDB v0.114-0 changelog](https://github.com/documentdb/documentdb/blob/v0.114-0/CHANGELOG.md)

`documentdb_extended_rum` is DocumentDB's extended RUM index access method. It is an implementation component selected by DocumentDB's indexing layer, not a general-purpose application index or a replacement for installing `documentdb`.

### Configure and Install

The library can only be initialized from `shared_preload_libraries`. Preload it after the base DocumentDB libraries and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_cron, pg_documentdb_core, pg_documentdb, pg_documentdb_extended_rum'
documentdb.alternate_index_handler_name = 'extended_rum'
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
- `documentdb.alternate_index_handler_name = 'extended_rum'` directs the DocumentDB index layer to the adapter.
- The implementation is a RUM fork whose on-disk layout and content are designed to remain backward compatible with upstream RUM while changing query and volatile paths for document workloads.

### Operational Boundaries

Install and upgrade this component with matching `documentdb` and `documentdb_core` binaries. Do not build indexes with its internal operator classes directly unless following upstream development guidance; create and manage indexes through the DocumentDB APIs so metadata stays consistent.

The v0.114-0 changelog describes a RUM WAL page-reuse marker and targeted posting-tree pruning, but both are feature-flagged and disabled by default pending stabilization. They are not default user-visible capabilities of this release.
