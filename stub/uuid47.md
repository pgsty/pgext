## Usage

Sources:

- [Official extension control file](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/pgext/uuid47/uuid47.control)
- [Official upstream documentation](https://github.com/stateless-me/uuidv47/blob/81b4025249cdd99e31ee633f756c26b3b8268e73/README.md)
- [Official project or provider page](https://uuidv47.stateless.me/)

`uuid47` — UUIDv7 storage with keyed UUIDv4-looking façade I/O

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "uuid47";
SELECT extversion
FROM pg_extension
WHERE extname = 'uuid47';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
