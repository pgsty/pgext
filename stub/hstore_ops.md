## Usage

Sources:

- [Official extension control file](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/hstore_ops.control)
- [Official upstream documentation](https://github.com/postgrespro/hstore_ops/blob/a9b00d87ee32171a7f7ef46101cfad9e90b0ff2e/README.md)

`hstore_ops` — Alternative GIN operator classes for hstore, optimized for smaller indexes and faster containment queries.

The reviewed catalog snapshot records version `1.1`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `hstore`.

```sql
CREATE EXTENSION "hstore_ops";
SELECT extversion
FROM pg_extension
WHERE extname = 'hstore_ops';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
