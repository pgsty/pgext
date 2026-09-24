## Usage

Sources:

- [README](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/README.md)
- [Control file](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/nabla.control)
- [Cargo.toml](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/Cargo.toml)
- [src/lib.rs](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/src/lib.rs)

`nabla` maintains views asynchronously from logical WAL and exposes ordered deltas to subscribers. The reviewed 0.1.0 source targets PostgreSQL 17 and is experimental; view contents can lag committed writes.

### Enablement

Configure `wal_level = logical`, add `nabla` to `shared_preload_libraries`, set `nabla.database` to an existing database, and reserve replication slots. Restart the server, then install as a superuser. The worker maintains views in one configured database.

```sql
CREATE EXTENSION nabla;
CREATE TABLE orders (id bigserial PRIMARY KEY, k int, amount numeric, status text);
ALTER TABLE orders REPLICA IDENTITY FULL;
SELECT nabla.create_view('orders_by_k',
  'SELECT k, count(*) AS n, sum(amount) AS total FROM orders WHERE status = ''paid'' GROUP BY k');
INSERT INTO orders(k, amount, status) VALUES (1, 10, 'paid');
SELECT nabla.await_ready('orders_by_k');
SELECT nabla.wait_for('orders_by_k', pg_current_wal_lsn());
SELECT * FROM orders_by_k;
```

### Subscriptions and Recovery

`nabla.current_seq` gives a delta cursor; `nabla.changes` reads changes after a sequence within an epoch. Notifications are wakeups, not a durable delivery mechanism. A subscriber must use the snapshot/cursor protocol and re-snapshot when the epoch changes or retained history is unavailable. `nabla.refresh` rebuilds a view; `nabla.drop_view` removes its registration.

Use only query forms accepted by the pinned source and monitor worker errors and view state. The simple filtered and grouped workflow above avoids more complex join restrictions. Replication slots can retain substantial WAL; exceeding the configured retention boundary can make a view stale and require rebuilding. Base-row replica identity must provide the old values needed by incremental maintenance.
