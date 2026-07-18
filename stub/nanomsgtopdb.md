## Usage

Sources:

- [Official extension control file](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/nanomsgtopdb.control)
- [Official upstream documentation](https://github.com/sangli00/nanomsgtopdb/blob/87df402495454362ccb4ed6ab58c33e7683e7401/README.md)

`nanomsgtopdb` — Receive nanomsg messages into a PipelineDB stream.

The reviewed catalog snapshot records version `1.0`, kind `preload`, and implementation language `C`.

```sql
CREATE EXTENSION "nanomsgtopdb";
SELECT extversion
FROM pg_extension
WHERE extname = 'nanomsgtopdb';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
