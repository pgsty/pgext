## Usage

Sources:

- [Official README.md](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/README.md)
- [Official pg_keyspace.control](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/pg_keyspace.control)
- [Official Cargo.toml](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/Cargo.toml)
- [Official lib.rs](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/src/lib.rs)
- [Official pg_keyspace--0.5.0--0.6.0.sql](https://github.com/supatype/postgres/blob/3bf5432da3c3bd228ba1977bf65e777d46a0c107/extensions/pg_keyspace/extension/sql/pg_keyspace--0.5.0--0.6.0.sql)

`pg_keyspace` 0.6.0 runs a RESP2/RESP3 keyspace in PostgreSQL background workers and provides an optional PostgREST row cache. Upstream builds and tests this implementation on PostgreSQL 17. The control and SQL version is 0.6.0; the Rust package manifest still says 0.1.0.

### Standalone Keyspace

Preload and restart, then install as a superuser. Standalone mode explicitly disables the optional masking integration:

```conf
shared_preload_libraries = 'pg_keyspace'
pg_keyspace.port = 6381
pg_keyspace.require_mask = off
```

```sql
CREATE EXTENSION pg_keyspace;
```

```sh
redis-cli -p 6381 SET demo:item example
redis-cli -p 6381 GET demo:item
```

```sql
SELECT convert_from(supacache.get('demo:item'), 'UTF8');
SELECT * FROM supacache.stats();
```

This minimal example assumes an isolated test listener. Configure RESP credentials, ACLs, tenant scoping and TLS before wider access. `supacache.set_credential` manages RESP credentials; these are not simply PostgreSQL connection authentication. `supatype_mask` and `pg_guard` are not required dependencies.

### Durability and Row Caching

`pg_keyspace.durability` selects ephemeral, relaxed, durable or replicated operation. Their acknowledgement and crash-recovery guarantees differ; persistence uses `supacache.kv`. Choose and test the required mode before storing data that cannot be reconstructed.

The optional row cache uses planner hooks and the supplied keys-only decoding plugin. Read-through is opt-in and requires the invalidation decoder. It preserves a bounded-staleness cache contract, not unconditional transaction-snapshot equivalence. RLS checks above a cached row do not make stale row data fresh.

### Operational Boundaries

Streams, Lua scripting and blocking list operations are not implemented. Multiple workers have distinct keyspaces and require deliberate client routing and recovery configuration. Background workers do not serve RESP on a streaming standby. Cross-instance pub/sub relay is opt-in and at-most-once; it is not keyspace replication. Rehearse upgrades and read the version-specific limitations before replacing an existing Redis service.
