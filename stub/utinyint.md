## Usage

Sources:

- [Official extension control file](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/utinyint.control)
- [Official upstream documentation](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/doc/utinyint.md)
- [Official upstream README](https://github.com/blueskylxb/utinyint/blob/8c5d9d832af6d148668d86769fa2b662b6582a19/README.md)

`utinyint` — One-byte unsigned integer type with casts to smallint, integer, and boolean.

The reviewed catalog snapshot records version `0.1.1`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "utinyint";
SELECT extversion
FROM pg_extension
WHERE extname = 'utinyint';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
