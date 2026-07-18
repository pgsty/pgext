## Usage

Sources:

- [Official extension control file](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/orb_fdw.control)
- [Official upstream documentation](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/README.md)
- [Official Rust package manifest](https://github.com/Jayko001/orb_fdw/blob/75c05395fac96a558f91e4ca51d1627140dd6c98/Cargo.toml)

`orb_fdw` — A foreign data wrapper for Orb

The reviewed catalog snapshot records version `0.13.3`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "orb_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'orb_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
