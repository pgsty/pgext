## Usage

Sources:

- [Upstream source 1](https://github.com/justusepostgres/pg_broker/blob/18f080f57de5e8c2bff310f93bc3366a1f0e7ade/README.md)
- [Upstream source 2](https://github.com/justusepostgres/pg_broker/blob/18f080f57de5e8c2bff310f93bc3366a1f0e7ade/crates/extension/pg_broker.control)

`pg_broker` SQL version 0.1.1 supplies a durable partitioned event log behind a separate Kafka protocol gateway.

### Core Workflow

Deploy matching gateway, binary and SQL artifacts. The gateway calls `public.broker_call(jsonb)` to create topics, append/fetch records and manage group offsets. The `broker` schema stores retained records and coordinator state; consuming a record does not delete it.

```sql
CREATE EXTENSION pg_broker;
SELECT extversion FROM pg_extension WHERE extname = 'pg_broker';
```

### Operational Boundaries

The supported build targets PostgreSQL 17 and requires privileged installation. Grant the gateway only its documented entry-point privileges. Successful produce acknowledgments follow a PostgreSQL commit, not a Kafka replica acknowledgment. This MVP has a limited Kafka API, no authentication/TLS, idempotent producer or Kafka transaction guarantee. Retention can remove data needed by lagging groups; monitor it separately.
