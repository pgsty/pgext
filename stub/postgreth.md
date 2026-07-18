## Usage

Sources:

- [Official extension control file](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/postgreth.control)
- [Official upstream documentation](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/README.md)
- [Official Rust package manifest](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/Cargo.toml)

`postgreth` — A pgrx extension providing Ethereum bloom/address/hash types plus Solidity ABI item and event-log decoding to JSONB.

The reviewed catalog snapshot records version `0.0.1`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgreth";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgreth';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
