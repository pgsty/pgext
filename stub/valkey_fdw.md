## Usage

Sources:

- [Upstream README](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/README.md)
- [Extension control file](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/valkey_fdw.control)
- [PGXN package metadata](https://github.com/hyperspike/valkey_fdw/blob/b9c0c8fd576097776b17e9de215ca97a34e87a84/META.json)

`valkey_fdw` is a foreign data wrapper for reading and modifying data in a Valkey key/value server. Version `1.0` requires PostgreSQL 17 or later and the `libvalkey` C client library. No preload setting is required.

```sql
CREATE EXTENSION valkey_fdw;

CREATE SERVER valkey_server
FOREIGN DATA WRAPPER valkey_fdw
OPTIONS (address '127.0.0.1', port '6379');

CREATE FOREIGN TABLE valkey_kv (
    key text,
    val text
)
SERVER valkey_server
OPTIONS (database '0');

SELECT * FROM valkey_kv WHERE key = 'session:42';
```

Use a user mapping with the `password` option when the server requires authentication. Foreign tables can also represent Valkey hashes, lists, sets, and sorted sets with the documented `tabletype`, `tablekeyprefix`, `tablekeyset`, and `singleton_key` options.

Valkey does not provide PostgreSQL-style MVCC or SQL cursors, so a scan is not an atomic snapshot: keys can disappear while tuples are being built, and cursor scans can return an element more than once. Only one equality condition on the `key` column can be pushed down. Upstream also requires a source branch matching the PostgreSQL major version because FDW APIs change.
