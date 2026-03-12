

## Usage

> [documentdb_core: Core API surface for DocumentDB for PostgreSQL](https://github.com/documentdb/documentdb)

DocumentDB provides MongoDB-compatible document database functionality built on PostgreSQL. The `documentdb_core` extension introduces BSON datatype support and operations for native Postgres.

### BSON Data Type

The extension adds a native BSON (Binary JSON) data type to PostgreSQL, enabling storage and manipulation of MongoDB-style documents.

### Basic Document Operations

Documents are managed through MongoDB-compatible CRUD operations via the DocumentDB API layer:

```python
import pymongo

client = pymongo.MongoClient(
    'mongodb://user:pass@localhost:10260/?tls=true&tlsAllowInvalidCertificates=true'
)

db = client["myDatabase"]
collection = db.create_collection("myCollection")

# Insert documents
collection.insert_one({
    'name': 'John Doe',
    'email': 'john@email.com',
    'address': '123 Main St'
})

collection.insert_many([
    {'name': 'Jane Smith', 'email': 'jane@email.com'},
    {'name': 'Alice Johnson', 'email': 'alice@email.com'}
])

# Query documents
for doc in collection.find():
    print(doc)

single = collection.find_one({'name': 'John Doe'})
```

### Aggregation Pipelines

```python
pipeline = [
    {'$match': {'name': 'Alice Johnson'}},
    {'$project': {'_id': 0, 'name': 1, 'email': 1}}
]

results = collection.aggregate(pipeline)
for doc in results:
    print(doc)
```

### Components

- **documentdb_core**: BSON datatype support and operations for native Postgres
- **documentdb (pg_documentdb)**: Public API surface providing CRUD functionality
- **pg_documentdb_gw**: Gateway protocol translation layer (MongoDB wire protocol to PostgreSQL)

The extension supports full-text search, geospatial queries, and vector search on BSON documents.
