## Usage

Sources:

- [Official extension control file](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/plid.control)
- [Official upstream documentation](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/README.md)
- [Official Rust package manifest](https://github.com/k0nserv/plid/blob/2bffd9caefa7eda2cb50955d00c7470498906c55/Cargo.toml)

`plid` — A 128-bit, lexicographically sortable PostgreSQL PLID type with short prefixes, millisecond timestamps, randomness, and optional monotonic generation.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "plid";
SELECT extversion
FROM pg_extension
WHERE extname = 'plid';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
