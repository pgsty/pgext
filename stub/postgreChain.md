## Usage

Sources:

- [Official README](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/README.md)
- [Version and PostgreSQL feature matrix](https://github.com/uygunbodur/postgrechain/blob/dafe2307effd636a6954526cd95190779987af07/Cargo.toml)
- [Official RPC implementation](https://github.com/uygunbodur/postgrechain/tree/dafe2307effd636a6954526cd95190779987af07/src/rpc)

`postgreChain` version 0.0.0 exposes synchronous Solana RPC operations as PostgreSQL functions. It can read balances and program accounts, create wallets, request test-network airdrops, and submit transfers; network calls run inside the database backend, so latency, failures, and secret handling directly affect database sessions.

### Core Workflow

```sql
CREATE EXTENSION "postgreChain";

SELECT pc_balance('WalletPublicKeyBase58', 'devnet');

SELECT public_key, data_len, lamports, rent_epoch, executable
FROM pc_get_program_accounts('ProgramPublicKeyBase58', 'testnet');
```

Supported network strings are `mainnet`, `devnet`, `testnet`, and `localhost`. The pinned implementation silently falls back to `devnet` for any other value, so validate the network in application code.

### Main Functions

- `pc_balance(address, network)` returns the account balance in lamports; the implementation returns zero on address or RPC errors.
- `pc_token_account_balance(address, network)` returns an SPL token-account balance and also returns zero on error.
- `pc_get_program_accounts(program_id, network)` returns account keys, data lengths, lamports, rent epochs, and executable flags.
- `pc_create_wallet()` returns a newly generated public key and Base58-encoded secret key.
- `pc_airdrop(address, amount, network)` requests SOL and returns success as boolean.
- `pc_transfer(public_key, secret_key, destination, amount, network)` signs and submits a SOL transfer and returns success as boolean.

### Security and Transaction Boundaries

Restrict execution aggressively. Wallet creation returns the secret key through SQL results, and transfer calls accept it as a SQL argument, which can expose it through client logs, statement logging, activity views, monitoring, or query history. A blockchain transfer or airdrop is an external side effect and is not rolled back with the surrounding PostgreSQL transaction.

RPC calls are synchronous and use fixed public Solana endpoints for the named networks. The catalog version is 0.0.0 and the source uses pgrx 0.11.3 with build features for PostgreSQL 11 through 16; verify the exact compiled package against your server and test on a non-production network before granting access.
