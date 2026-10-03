## Usage

Sources:

- [Official README.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/README.md)
- [Official pg_replayfeed.control](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/pg_replayfeed.control)
- [Official pg_replayfeed--2.0.sql](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/sql/pg_replayfeed--2.0.sql)
- [Official semantics.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/docs/semantics.md)
- [Official security.md](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/docs/security.md)
- [Official ci.yml](https://github.com/pg-replayfeed/pg_replayfeed/blob/82bd85807cf70947d2a9471c556fd4c8b1ecf6bb/.github/workflows/ci.yml)

`pg_replayfeed` 2.0 exposes key-only change envelopes from logical decoding. A committed transaction produces a JSON envelope; consumers checkpoint its commit LSN and handle possible redelivery. The reviewed development revision has a PostgreSQL 16–18 CI matrix.

### Register and Read a Feed

Set logical WAL and sufficient replication slots and senders; restart when changing these startup settings. The extension does not require shared preload.

```conf
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

```sql
CREATE EXTENSION pg_replayfeed;
CREATE TABLE feed_demo (id bigint PRIMARY KEY, tenant text, value text);
SELECT replayfeed.create_feed('demo', 'feed_demo', 'id', 'tenant');
SELECT pg_create_logical_replication_slot('demo_slot', 'pg_replayfeed');
INSERT INTO feed_demo VALUES (1, 'team_a', 'first');
SELECT * FROM pg_logical_slot_peek_changes('demo_slot', NULL, NULL);
```

Run this administrative example as a superuser. `create_feed` accepts an ordinary logged table and a NOT NULL key backed by a single-column unique index; it sets `REPLICA IDENTITY FULL`. Inspect `replayfeed.feed` and remove a registration with `replayfeed.drop_feed`. Dropping a feed does not delete replication slots.

### Consumers, Privileges and Retention

`pg_logical_slot_get_changes` consumes and advances a slot; `pg_logical_slot_peek_changes` does not. Replication clients can instead stream with the PostgreSQL replication protocol. Acknowledge only after durable processing and deduplicate by commit position. Key changes, tenant moves and truncation require the consumer to follow the documented envelope semantics.

The fixed `replayfeed` schema contains privileged SECURITY DEFINER administration functions with PUBLIC execution revoked. Delegated administrators need explicit schema/function grants and the checked source-table privileges. A replication role also needs catalog visibility. Feed selection is a filter, not tenant authorization; downstream consumers must enforce isolation.

Monitor slot lag and WAL disk use. An inactive slot retains WAL until advanced or removed; remove the demonstration slot after use:

```sql
SELECT pg_drop_replication_slot('demo_slot');
SELECT replayfeed.drop_feed('demo');
```
