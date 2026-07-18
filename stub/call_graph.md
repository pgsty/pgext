## Usage

Sources:

- [Official extension control file](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/call_graph.control)
- [Official upstream documentation](https://github.com/johto/call_graph/blob/9bcba8927be7374263db9d0dd55ec6f8015664d9/README)

`call_graph` — Track PostgreSQL function calls automatically and aggregate them into call graphs and optional table-usage statistics.

The reviewed catalog snapshot records version `1.4`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "call_graph";
SELECT extversion
FROM pg_extension
WHERE extname = 'call_graph';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
