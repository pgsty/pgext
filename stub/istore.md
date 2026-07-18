## Usage

Sources:

- [Official extension control file](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore.control)
- [Official upstream documentation](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/README.md)
- [Official PGXN distribution page](https://pgxn.org/dist/istore/)

`istore` — Integer-to-integer key/value store for analytic workloads.

The reviewed catalog snapshot records version `0.1.12`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "istore";
SELECT extversion
FROM pg_extension
WHERE extname = 'istore';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
