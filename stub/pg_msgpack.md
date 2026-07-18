## Usage

Sources:

- [Official extension control file](https://github.com/choplin/pg_msgpack/blob/7cf18817053f4a42df80c34d816003505de8b1c1/pg_msgpack.control)

`pg_msgpack` — MessagePack data type with JSON and bytea casts and access operators

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pg_msgpack";
SELECT extversion
FROM pg_extension
WHERE extname = 'pg_msgpack';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
