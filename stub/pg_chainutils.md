## Usage

Sources:

- [Upstream usage documentation](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/README.md)
- [Version and dependency manifest](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/Cargo.toml)
- [Extension control file](https://github.com/usagi-coffee/pg_chainutils/blob/98f534c3912aad5c0c23276baf50045325be7dfb/pg_chainutils.control)

`pg_chainutils` is a pgrx extension for parsing and transforming blockchain values inside PostgreSQL. It supplies Ethereum-style `H160`, `H256`, and `U256` types and helpers for event topics and payloads, including ERC-20, ERC-721, Uniswap, Sushiswap, and related formats.

```sql
CREATE EXTENSION pg_chainutils;
SELECT H256.keccak256('Sync(uint112,uint112)');
SELECT H160.from_H256('0x0000000000000000000000001111111111111111111111111111111111111111');
```

Version `0.2.4` is built with pgrx and the reviewed manifest defaults to PostgreSQL 17. Input helpers expect exact ABI encodings; malformed lengths, prefixes, topic ordering, signedness, token decimals, and chain-specific conventions must be validated before ingestion. Treat decoded values as untrusted data and compare critical financial results with an independent decoder.
