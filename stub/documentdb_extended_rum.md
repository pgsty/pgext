

## Usage

> [documentdb_extended_rum: DocumentDB Extended RUM index access method](https://github.com/documentdb/documentdb)

The `documentdb_extended_rum` extension provides an enhanced RUM (Recursive Union Merge) index access method for DocumentDB on PostgreSQL. It improves query performance for document-based workloads.

### Overview

This extension extends the RUM index type to better support BSON document indexing within DocumentDB. It provides optimized index access methods for:

- Full-text search on document fields
- Compound index operations on BSON data
- Efficient range queries and sorting on indexed document properties

### Prerequisites

Requires `documentdb_core` to be installed.

### Enabling

```sql
CREATE EXTENSION documentdb_extended_rum;
```

The extended RUM indexes are automatically utilized by the DocumentDB query planner when appropriate for document query patterns.
