## Usage

Sources:

- [Official extension control file](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf.control)
- [Official upstream documentation](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/README.md)

`shbf` — C extension providing Bloom filter, shifting Bloom filter, and count-min sketch PostgreSQL data types and functions.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "shbf";
SELECT extversion
FROM pg_extension
WHERE extname = 'shbf';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
