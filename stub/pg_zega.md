## Usage

Sources:

- [Official source](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/README.md)
- [Official source](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/postgres/pg_zega.control)
- [Official source](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/LICENSE)
- [Official source](https://github.com/zegadb/pg_zega/blob/f264168306e81a8307e3ce609a1d31620315dbfd/postgres/pg_zega--0.3.sql)

pg_zega 0.3 synchronizes supported tables and DDL into an external Zega graph. It requires PostgreSQL 17, libcurl and superuser administration. PostgreSQL credentials and the Zega bearer token are distinct secrets. Configure zega.url and the protected zega.password setting before tracking data, reload and use a new session.

### Track and import

```sql
CREATE EXTENSION pg_zega;
CREATE TABLE users(user_id integer PRIMARY KEY, name text NOT NULL);
SELECT zega.track('users');
INSERT INTO users VALUES (10, 'Nia');
SELECT zega.import('users');
```

Tracking adds a row trigger but does not copy existing rows. A node table requires one primary-key column; supported fields are integer, floating-point, boolean and text-like types. A ze_-prefixed edge table needs exactly two disjoint foreign-key constraints and a primary key comprising their columns. Unsupported types, including numeric, timestamp, uuid and json, are rejected.

### Delivery and maintenance

Default precommit mode waits for Zega; rejection aborts the PostgreSQL transaction. This is not a general distributed-transaction guarantee. Journal mode commits locally and delivers asynchronously using Zega receipts; it requires preload/restart, logical WAL, a delivery worker and a replication slot created before enabling the mode. Monitor journal lag and retained WAL.

zega.import locks every tracked table against writes and can replace graph nodes/edges in its scope. Run it without other row changes in the transaction; imported engine IDs change. The DDL event trigger rejects changes incompatible with existing graph values. Precommit mode does not require preload; journal mode needs the additional configuration in the official guide.
