## Usage

Sources:

- [Upstream README](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/README.md)
- [Installation guide](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/INSTALL.md)
- [Extension control file](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/pgnats.control)
- [Extension Cargo manifest](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/Cargo.toml)
- [Version 1.1.0 upgrade SQL](https://github.com/luxms/pgnats/blob/6928cd1beed1ff8ae68fea0b82529550a93d9325/sql/pgnats--1.0.0--1.1.0.sql)

`pgnats` connects SQL functions to a NATS server for publish/request messaging, JetStream, key-value storage, object storage, and optional subject subscriptions. Preloading is specifically required for `nats_subscribe` and `nats_unsubscribe` background-worker support:

```conf
shared_preload_libraries = 'pgnats.so'
```

Restart PostgreSQL, create the extension and its single allowed NATS foreign server, then call the messaging API:

```sql
CREATE EXTENSION pgnats;

CREATE SERVER nats_fdw_server
  FOREIGN DATA WRAPPER pgnats_fdw
  OPTIONS (host 'nats.example.net', port '4222', capacity '128');

SELECT nats_publish_jsonb(
  'events.order.created',
  '{"id":42}'::jsonb
);

SELECT nats_request_text('service.echo', 'ping', 1000);
```

### Build and privilege boundaries

Version `1.1.0` has Cargo features for PostgreSQL 14 through 18. Build features control whether subscription, key-value, and object-store APIs are present; the default build enables all three. Subscription callbacks must accept one `bytea` argument. The SQL functions can send data, invoke callbacks, and access external NATS storage, so restrict `EXECUTE`, use NATS-side ACLs, and configure TLS certificate paths where required. The extension enforces at most one `pgnats_fdw` server per database.
