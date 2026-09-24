## Usage

Sources:

- [PGXN 0.1.8 README](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/README.md)
- [Control file](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/pg_reactive.control)
- [Versioned extension SQL](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/pg_reactive--0.1.8.sql)
- [Module initialization](https://api.pgxn.org/src/pg_reactive/pg_reactive-0.1.8/src/pg_reactive.c)

`pg_reactive` turns PostgreSQL SELECT queries into subscriptions and sends changed result rows through LISTEN/NOTIFY. Version 0.1.8 supports PostgreSQL 15–18, requires preloading and a restart, and exposes privileged subscription APIs in the `pgr` schema.

### Enable and Subscribe

Add the library to the existing preload list, then restart PostgreSQL. The installation SQL also uses `plpgsql`, normally installed by default.

```conf
shared_preload_libraries = 'pg_reactive'
pg_reactive.max_subscriptions = 1024
```

Run this example as the extension owner or a trusted administrator, with each statement committed normally:

```sql
CREATE EXTENSION pg_reactive;

CREATE TABLE public.pgr_demo (
    id bigint PRIMARY KEY,
    status text NOT NULL,
    total numeric
);

LISTEN pgr;
SELECT pgr.subscribe(
    'open_orders',
    $$SELECT id, total FROM public.pgr_demo WHERE status = 'open'$$
);

INSERT INTO public.pgr_demo VALUES (1, 'open', 42);
SELECT * FROM pgr.subscriptions;
SELECT * FROM pgr.stats();
```

The extension tracks table and column dependencies, installs statement triggers, and computes set differences against UNLOGGED snapshots. The `delta` mode returns inserted/deleted rows; `notify` sends invalidation messages so the client can fetch results itself.

```sql
SELECT pgr.subscribe(
    'orders_hint',
    $$SELECT id FROM public.pgr_demo$$,
    'notify'
);

SELECT pgr.unsubscribe('orders_hint');
SELECT pgr.unsubscribe('open_orders');
```

### Interfaces and Configuration

- `pgr.subscribe(query_id, query, mode, audience)` validates a SELECT and records its subscription; the last two arguments default to delta mode and a null audience.
- `pgr.unsubscribe(query_id)` removes the subscription, triggers, snapshot and durable catalog entry.
- `pgr.get_subscriptions()` and `pgr.subscriptions` expose active shared-memory state; `pgr.stats()` reports counters.
- `pgr.subscription_meta(query_id)` reads committed mode, audience and generation metadata for a proxy.
- `pgr.persisted_subscriptions` is the durable subscription catalog.

`pg_reactive.max_subscriptions`, `pg_reactive.async_recompute` and `pg_reactive.database` require restart. Asynchronous recomputation is off by default; its worker uses the configured database. `pg_reactive.batch_invalidation` defaults to on and batches recomputation at transaction pre-commit. `pg_reactive.notify_channel` defaults to `pgr`; these last two settings are superuser-controlled at runtime.

### Recovery and Security Boundaries

Shared memory is lost on restart. A trusted startup procedure must run the following once in each affected database; inspect warnings for subscriptions that could not be restored:

```sql
SELECT pgr.restore_subscriptions();
```

The 0.1.8 SQL script revokes public access to subscription functions and the view. Granting only `pgr.subscribe()` is insufficient because it calls restricted helpers and a restricted sequence. Do not expose arbitrary subscription SQL to untrusted roles; upstream specifies a purpose-built wrapper with fixed query templates and server-derived identifiers. Audience metadata is for enforcement by a trusted proxy, not per-subscription authorization on the shared notification channel.

Clients must re-fetch after an overflow or snapshot-layout change, a missed sequence, or reconnection. Notifications share one channel and are not a durable delivery queue. Registration generations identify re-subscriptions; clients/proxies must discard obsolete generations. Recomputing a large result set adds work to writes even when few result rows change.

The reviewed PGXN archive contains the control file, SQL and C source. Its advertised GitHub repository returned 404 during verification; use the versioned PGXN sources above for this release.
