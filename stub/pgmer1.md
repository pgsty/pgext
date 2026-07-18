## Usage

Sources:

- [Official extension control file](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/pgmer1.control)
- [Official upstream documentation](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/README.md)
- [Official Rust package manifest](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/Cargo.toml)

`pgmer1` — MeritRank HTTP connector exposing graph ranking service calls in SQL

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "pgmer1";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgmer1';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
