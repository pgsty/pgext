

## Usage

> [pglogical_origin: Dummy extension for compatibility when upgrading from Postgres 9.4](https://github.com/2ndQuadrant/pglogical)

The `pglogical_origin` extension is a compatibility shim provided alongside pglogical. It exists solely to facilitate upgrades from PostgreSQL 9.4, where replication origin tracking was handled by the pglogical extension itself rather than by PostgreSQL core.

### Enabling

```sql
CREATE EXTENSION pglogical_origin;
```

### Overview

Starting with PostgreSQL 9.5, replication origin tracking became a built-in PostgreSQL feature (`pg_replication_origin`). The `pglogical_origin` extension is a dummy/empty extension that:

- Prevents errors when upgrading databases that previously depended on it
- Provides a smooth migration path from pglogical on PostgreSQL 9.4 to newer versions
- Contains no actual functionality -- all origin tracking is handled by PostgreSQL core

### When to Use

This extension is only needed when:

- Upgrading a database from PostgreSQL 9.4 that used pglogical
- The database has existing references to the `pglogical_origin` extension

For new installations, this extension is not required. Use pglogical directly, which leverages PostgreSQL's built-in replication origin support.
