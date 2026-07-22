## Usage

Sources:

- [Official README](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/README.md)
- [Extension control file](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops.control)
- [Version 1.1 extension SQL](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops--1.1.sql)

`hstore_ops` adds two alternative GIN operator classes for the PostgreSQL `hstore` type. Choose it when containment-heavy workloads need a smaller or faster index than the built-in GIN class, or when bytewise comparison avoids expensive collation work.

### Core Workflow

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION hstore_ops;

CREATE TABLE item_attributes (
    item_id bigint PRIMARY KEY,
    attrs hstore NOT NULL
);

CREATE INDEX item_attrs_hash_idx
ON item_attributes USING gin (attrs gin_hstore_hash_ops);

SELECT item_id
FROM item_attributes
WHERE attrs @> 'color=>blue'::hstore;
```

For the alternative bytewise class:

```sql
CREATE INDEX item_attrs_bytea_idx
ON item_attributes USING gin (attrs gin_hstore_bytea_ops);
```

Both classes support the `@>`, `?`, `?|`, and `?&` operators.

### Choosing an Operator Class

- `gin_hstore_hash_ops` stores hashes of keys and values. Upstream reports smaller indexes and faster `@>` searches than the default GIN class, while key-existence operators can be somewhat slower.
- `gin_hstore_bytea_ops`, added in version `1.1`, is a variation of the standard `hstore` GIN class that compares keys bytewise instead of using collation keys. It is intended to improve index work when collation comparison is slow without changing supported operators.

Hash collisions are possible with `gin_hstore_hash_ops`, so PostgreSQL must recheck candidate rows. Measure index size, build time, update cost, and each query shape on representative data rather than assuming one class is universally faster.

### Dependencies and Maintenance

The control file declares `hstore` as a required extension; no preload or restart is required. Upgrading an existing `hstore_ops` `1.0` installation to `1.1` adds `gin_hstore_bytea_ops`:

```sql
ALTER EXTENSION hstore_ops UPDATE TO '1.1';
```

Changing an existing index to another operator class requires rebuilding that index. Plan the lock and disk-space impact as for any large GIN index build.
