## Usage

Sources:

- [README](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/README.md)
- [Control file](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/pg_cryptocurrency.control)
- [Cargo.toml](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/Cargo.toml)
- [src/lib.rs](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/src/lib.rs)
- [docs/SECURITY.md](https://github.com/RustedBytes/pg-cryptocurrency/blob/308046c977bbd3674adce767890a2d3d6f346dbd/docs/SECURITY.md)

`pg_cryptocurrency` supplies exact blockchain asset amounts, network-qualified identities, addresses, and unsigned 256-bit integers on PostgreSQL 14–18. It performs no blockchain access or network requests.

### Core Workflow

```sql
CREATE EXTENSION pg_cryptocurrency;
SELECT '1.25 BTC'::crypto_amount;
SELECT '1500 USDC@ethereum'::crypto_amount;
SELECT crypto_to_units('10 ETH');
SELECT crypto_from_units(1000000000000000001::bigint, 'ETH');
SELECT '1 BTC'::crypto_amount + '0.25 BTC';
```

### Objects and Identity

`crypto_amount` binds an exact amount to an asset. `crypto_asset` identifies native currencies and tokens; `crypto_address` validates network-qualified addresses, and `uint256` stores unsigned 256-bit values. `crypto_register_asset` registers a token using its symbol, network, contract, and decimals. `crypto_amount_value`, `crypto_amount_asset`, and `crypto_amount_network` inspect a value.

A ticker alone is not token identity: network and contract matter. Mixed-asset arithmetic fails rather than converting implicitly. Conversion to smallest units checks the asset scale and amount bounds. Address validation does not prove ownership or on-chain existence.

### Operation

The extension is relocatable and requires no preload. Administrative registry operations need explicit privilege review. The optional adapter for `pg_money` is separate from the core type workflow; installing this extension does not establish a wallet, custody service, or exchange-rate source.
