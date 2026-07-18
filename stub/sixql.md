## Usage

Sources:

- [Official upstream documentation](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/readme.md)
- [Official Rust package manifest](https://github.com/wrnrlr/sixql/blob/4b4231ca0d3ef779c9ccb4bcd0f9f3c4e41ab969/Cargo.toml)

`sixql` — Experimental pgrx extension attempting to render SIXEL plots through PostgreSQL server output for compatible terminals.

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "sixql";
SELECT extversion
FROM pg_extension
WHERE extname = 'sixql';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
