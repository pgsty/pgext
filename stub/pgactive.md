

## Usage

> [pgactive: Active-Active Replication Extension for PostgreSQL](https://github.com/aws/pgactive)

The `pgactive` extension provides active-active (multi-master) replication for PostgreSQL, where multiple database instances can accept writes and replicate changes to each other.

### Enabling

```sql
CREATE EXTENSION pgactive;
```

### Overview

pgactive enables active-active replication using logical replication as its foundation. Multiple databases in a cluster can accept changes and replicate them bidirectionally.

### Key Concepts

- **Active-Active**: All nodes accept reads and writes simultaneously
- **Asynchronous**: Changes are replicated asynchronously between nodes
- **Conflict Resolution**: Applications must handle conflicting changes from multiple writers
- **Logical Replication**: Uses PostgreSQL's logical decoding to interpret and apply changes

### Use Cases

- Multi-region high availability database clusters
- Reducing write latency between applications and databases
- Blue/green application updates
- Data migration between systems that must both remain writable

### Design Considerations

Applications using pgactive must be designed to handle:

- **Write conflicts**: Concurrent modifications to the same row on different nodes
- **Replication lag**: Changes may not be immediately visible on all nodes
- **Feature limitations**: Some features like auto-incrementing sequences require special handling across nodes

### Notes

- Built on PostgreSQL's native logical replication infrastructure
- Requires PostgreSQL 10+ (native logical replication support)
- Refer to the project documentation for detailed setup and conflict resolution strategies
