## Usage

Sources:

- [Project README](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/README.md)
- [Implemented function list in Rust](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/src/lib.rs)
- [Function documentation](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/docs/FUNCTIONS.md)
- [Extension control file](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/mumak.control)
- [Extension Cargo manifest](https://github.com/txpipe/mumak/blob/4da43f3cb21d986280faf2ad373db954418dc888/extension/Cargo.toml)

`mumak` 0.0.6 is a pgrx extension for inspecting Cardano CBOR inside PostgreSQL. It lets an indexer keep blocks, transactions, or UTxOs as `bytea` and derive hashes, addresses, fees, lovelace and asset amounts, metadata, Plutus data, epochs, slots, and other fields in SQL without first normalizing every structure.

### Verify installation and query CBOR

```sql
CREATE EXTENSION mumak;
SELECT hello_extension();

SELECT tx_hash(tx_cbor), tx_fee(tx_cbor), tx_outputs_json(tx_cbor)
FROM cardano_transactions;
```

Implemented function families include `block_tx_count(bytea)`, `block_number(bytea)`, `tx_hash(bytea)`, `tx_inputs(bytea)`, `tx_outputs(bytea)`, `tx_is_valid(bytea)`, `utxo_lovelace(integer,bytea)`, and address/Bech32 conversion helpers. Use the Rust implementation as the API authority: the top-level README contains planned names that do not exactly match all functions shipped in 0.0.6.

### Input and deployment cautions

Many decode failures return sentinel results such as an empty value, false, zero, or -1 instead of raising an error. Validate CBOR and its era/network context before treating those results as blockchain facts. The control file requires superuser installation. The reviewed Cargo features cover PostgreSQL 11 through 16, and the project publishes only source and limited Linux artifact guidance; build and regression-test the exact server/architecture combination.
