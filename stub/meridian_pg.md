## Usage

Sources:

- [Project README](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/README.md)
- [Extension implementation](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/lib.rs)
- [Extension control file](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/meridian_pg.control)
- [CRDT SQL functions](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/types/mod.rs)
- [Notification trigger](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/src/trigger.rs)
- [Extension SQL](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/sql/meridian--0.1.0.sql)
- [Package metadata](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/crates/meridian-pg/Cargo.toml)
- [Postgres example schema](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/examples/postgres-articles/schema.sql)
- [MIT license](https://github.com/Chahine-tech/meridian/blob/d19739734eb222e307baa5ca4252631c2a06aed1/LICENSE)

`meridian_pg` 0.1.0 stores Meridian CRDT states in PostgreSQL `bytea` columns and provides SQL functions for counters, sets, registers, read-only views of large documents, and change notification. It is source-only, uses pgrx 0.13, and supports the PostgreSQL 14–17 feature gates declared by its package.

### Core Workflow

```sql
CREATE EXTENSION meridian_pg;

CREATE TABLE articles (
    id       text PRIMARY KEY,
    views    bytea,
    likes    bytea,
    tags     bytea,
    headline bytea
);

INSERT INTO articles (id) VALUES ('rust-intro');

UPDATE articles
SET views = meridian.gcounter_increment(views, 3, 42),
    likes = meridian.pncounter_decrement(likes, 1, 42)
WHERE id = 'rust-intro';

SELECT meridian.gcounter_value(views),
       meridian.pncounter_value(likes)
FROM articles
WHERE id = 'rust-intro';
```

`orset_add` accepts a JSON element string and a node/sequence identity; `orset_elements` returns a JSON array. `lww_set` accepts a JSON value, wall-clock milliseconds, and an author ID, while `lww_value` and `lww_updated_at_ms` read the winning state. Merge functions combine counter and set states.

### Notification and Large Documents

Attach `meridian.notify_trigger(namespace, id_prefix, primary_key_column, crdt_columns...)` after inserts or updates. It emits one `pg_notify` event on channel `meridian_ops` for each changed, nonempty CRDT column and leaves the row unchanged. PostgreSQL notification payloads above the extension's 7,800-byte guard are skipped with a warning rather than delivered partially.

Use the Meridian server's pg-sync logical-replication path for RGA and Tree documents that can exceed that boundary. The extension provides SQL read functions `rga_text`, `rga_len`, `tree_json`, and `tree_node_count`; writes for those CRDT types are intended to come through Meridian clients, not SQL.

### Boundaries

The control file sets `relocatable = false` and `superuser = false`. Preserve a stable CRDT type for each column and stable client/node IDs across writes; decoding or serializing an invalid state can warn and return a default or empty state. The extension itself does not authenticate Meridian clients or guarantee that a PostgreSQL table is reachable from a server. Review trigger privileges, notification exposure, replication-slot retention, and the Meridian server's namespace authorization before production use.
