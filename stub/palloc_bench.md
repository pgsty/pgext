## Usage

Sources:

- [Official extension control file](https://github.com/tvondra/palloc_bench/blob/7f5c9caa357aa44c49b7af6167389aeaa54b71d3/palloc_bench.control)

`palloc_bench` — Benchmark function for PostgreSQL palloc memory allocation

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "palloc_bench";
SELECT extversion
FROM pg_extension
WHERE extname = 'palloc_bench';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
