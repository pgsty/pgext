## Usage

Sources:

- [Official extension control file](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/postfga.control)
- [Official upstream documentation](https://github.com/carped99/postfga/blob/92626b6120ab526acd00ff457cfc2071538836c6/TESTING.md)

`postfga` — A preload PostgreSQL extension integrating OpenFGA authorization through an FDW, C functions, gRPC background work and shared-memory caching.

The reviewed catalog snapshot records version `1.0.0`, kind `preload`, and implementation language `C`.
The curated compatibility set is `18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postfga";
SELECT extversion
FROM pg_extension
WHERE extname = 'postfga';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
