## Usage

Sources:

- [Official extension control file](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/pgtera.control)
- [Official upstream documentation](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/README.md)
- [Official Rust package manifest](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/Cargo.toml)

`pgtera` — render Tera HTML templates inside PostgreSQL

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgtera";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgtera';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
