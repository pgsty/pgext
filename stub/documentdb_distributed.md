

## Usage

> [documentdb_distributed: Multi-Node API surface for DocumentDB](https://github.com/documentdb/documentdb)

The `documentdb_distributed` extension provides multi-node distributed capabilities for DocumentDB on PostgreSQL. It extends the core DocumentDB functionality to support horizontal scaling across multiple PostgreSQL nodes.

### Overview

This extension works in conjunction with `documentdb_core` to provide distributed document database operations. It enables:

- Sharding of document collections across multiple nodes
- Distributed query execution for MongoDB-compatible operations
- Horizontal scaling for large document workloads

### Prerequisites

The `documentdb_distributed` extension requires:

- `documentdb_core` extension installed and configured
- A multi-node PostgreSQL cluster (typically using Citus for distribution)

### Enabling

```sql
CREATE EXTENSION documentdb_distributed;
```

The distributed layer transparently routes CRUD operations and aggregation pipelines across the cluster nodes while maintaining MongoDB wire protocol compatibility.
