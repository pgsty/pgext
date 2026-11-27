## Usage

Sources:

- [README.md](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/README.md)
- [Cargo.toml](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/Cargo.toml)
- [pg_bitcoin.control](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/pg_bitcoin.control)
- [src/lib.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/lib.rs)
- [src/config.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/config.rs)
- [src/client_wrapper.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/client_wrapper.rs)
- [src/electrum.rs](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/src/electrum.rs)
- [tests/pg_regress/sql/setup.sql](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/tests/pg_regress/sql/setup.sql)
- [LICENSE](https://github.com/LPbigFish/pg_bitcoin/blob/61671288be0cabfe50f25491940c73f57d20fa6e/LICENSE)

`pg_bitcoin` exposes Bitcoin address balances, history and unspent outputs through an Electrum server. The reviewed source declares version 0.0.0; this is an early source version, without a published release. Its manifest declares PostgreSQL 16 and 18, with 18 as the default.

### Enable and Configure

A superuser installs this untrusted, non-relocatable extension. Ordinary use does not require shared preload. Start a reachable Electrum indexer for the selected network, then configure the backend:

```sql
CREATE EXTENSION pg_bitcoin;
SET pg_bitcoin.network_type = 'regtest';
SET pg_bitcoin.electrum_url = 'tcp://127.0.0.1:60401';
SELECT bitcoin_ping();
```

Both settings require superuser privileges. Their defaults are `regtest` and `tcp://127.0.0.1:60401`. Supported network values are `mainnet`, `regtest`, `testnet`, `testnet4` and `signet`. The address must belong to the selected network; changing either setting recreates the cached connection on its next use.

### SQL Interface

`get_confirmed_balance(text)` returns the confirmed balance. `get_balance(text)` and `get_balance(text[])` return address, confirmed and unconfirmed amounts. `get_history(text)` returns height, transaction hash and optional fee; its array overload also identifies the address. `list_unspent(text)` returns height, transaction hash, output position and value. Supply actual addresses for the configured network.

Amounts are integer satoshi values. These are synchronous queries to the indexer, not a local blockchain replica or wallet. `bitcoin_ping()` returns false on a connection or ping error; the data functions raise PostgreSQL errors for invalid addresses and transport failures.

### Operational Boundaries

Each backend caches an Electrum client. The client configuration uses a five-second timeout and one retry. Remote data and latency depend on the selected indexer, and the sample TCP endpoint is unencrypted. Restrict SQL access and choose the endpoint for the intended trust boundary. The extension provides query functions; it does not hold private keys or submit transactions. No SQL upgrade scripts are supplied in this revision.
