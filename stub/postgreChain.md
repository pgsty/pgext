## Usage

Sources:

- [Official extension control file](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/postgreChain.control)
- [Official upstream documentation](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/README.md)
- [Official Rust package manifest](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/Cargo.toml)

`postgreChain` — Query Solana RPC data and perform wallet, airdrop, and transfer operations directly from PostgreSQL

The reviewed catalog snapshot records version `0.0.0`, kind `standard`, and implementation language `Rust`.
The curated compatibility set is `11,12,13,14,15,16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "postgreChain";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgreChain';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
