


## Usage

Sources: [README](https://github.com/documentdb/documentdb/blob/v0.113-0/README.md), [CHANGELOG](https://github.com/documentdb/documentdb/blob/v0.113-0/CHANGELOG.md), [documentdb.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb/documentdb.control), [documentdb_core.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_core/documentdb_core.control), [documentdb_extended_rum.control](https://github.com/documentdb/documentdb/blob/v0.113-0/pg_documentdb_extended_rum/documentdb_extended_rum.control)

`documentdb` is a MongoDB-compatible document database implemented as PostgreSQL extensions. It adds BSON storage and APIs in PostgreSQL, plus an optional gateway layer for MongoDB wire-protocol clients. FerretDB 2.0+ can use DocumentDB as its backend.

### Components

The public extension surface is split across related extensions:

- `documentdb_core`: BSON datatype and low-level BSON operations.
- `documentdb`: public API for document CRUD and query behavior.
- `documentdb_extended_rum`: extended RUM access method used by DocumentDB indexing.
- `pg_documentdb_gw`: gateway protocol layer used by the local Docker image and MongoDB-compatible clients.

Install the SQL extension in each database that needs the API:

```sql
CREATE EXTENSION IF NOT EXISTS documentdb CASCADE;
```

The `documentdb.control` file for `0.113-0` declares dependencies on `documentdb_core`, `pg_cron`, `tsm_system_rows`, `vector`, and `postgis`. The gateway container listens on a MongoDB-compatible port; the README examples use `10260` to avoid colliding with local MongoDB services.

### MongoDB Client Example

```python
import pymongo

client = pymongo.MongoClient(
    "mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true"
)

db = client["quickStartDatabase"]
coll = db.create_collection("quickStartCollection")

coll.insert_one({"name": "Alice", "email": "alice@example.com"})
print(coll.find_one({"name": "Alice"}))
```

The upstream README also demonstrates aggregation pipelines through normal MongoDB drivers:

```python
pipeline = [
    {"$match": {"name": "Alice"}},
    {"$project": {"_id": 0, "name": 1, "email": 1}},
]

for doc in coll.aggregate(pipeline):
    print(doc)
```

### Version Notes

This project's CSV tracks DocumentDB `0.113` for PostgreSQL 15-18. The upstream tag is `v0.113-0`; control files report `default_version = '0.113-0'`.

The `0.111` through `0.113` changelog entries are mostly query-planner, collation, and index correctness work:

- `0.113-0` adds opt-in collation support for non-unique ordered indexes with `$in` and `$nin`, and supports pruning dead index entries on ordered TTL indexes behind feature flags.
- `0.112-0` removes the legacy composite-returning `bson_update_document` UDF path, expands non-unique ordered-index collation support, and improves `$group` and accumulator execution.
- `0.111-0` adds background init job infrastructure, more `$group` validation, collation/index pushdown improvements, and several crash fixes.

### Caveats

- DocumentDB is a multi-extension stack; `CREATE EXTENSION documentdb CASCADE` is the normal entry point, but operational deployments also need the gateway/runtime pieces if MongoDB wire compatibility is required.
- Some features listed in the changelog are gated by `documentdb.*` feature flags. Verify flag defaults in the exact installed build before documenting behavior as always-on.
- `documentdb_extended_rum` is relocatable, but `documentdb` and `documentdb_core` are not.
