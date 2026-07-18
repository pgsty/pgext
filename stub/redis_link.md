## Usage

Sources:

- [Upstream README](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/README.md)
- [Extension control file](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link.control)
- [SQL installation script](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link--1.0.sql)
- [C implementation](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link.c)

`redis_link` version `1.0` is a Redis foreign data wrapper tested upstream on PostgreSQL 16. It can scan Redis string, hash, and set data, handles Redis Cluster redirections, and supports `INSERT`; it does not implement `UPDATE` or `DELETE` callbacks.

### Example

```sql
CREATE EXTENSION redis_link;
CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_link
  OPTIONS (host_port '127.0.0.1:6379');
CREATE FOREIGN TABLE redis_items (name text, content text, expire integer)
  SERVER redis_server
  OPTIONS (key 'name', value 'content', prefix 'demo:', table_type 'string', ttl 'expire');
SELECT * FROM redis_items;
```

The validator accepts a `password` option, but the reviewed C implementation never consumes it, so do not assume authenticated Redis support. It constructs Redis commands by concatenating identifiers and row values and passes them through hiredis command formatting; spaces, percent signs, or untrusted data can change parsing or trigger memory-safety bugs. Use only on a tightly isolated, trusted Redis namespace and do not expose writes or option creation to application roles.
