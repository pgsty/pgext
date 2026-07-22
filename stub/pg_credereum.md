## Usage

Sources:

- [Official README](https://github.com/postgrespro/pg_credereum/blob/master/README.md)
- [Official control file](https://github.com/postgrespro/pg_credereum/blob/master/pg_credereum.control)

`pg_credereum` is a PostgreSQL 10 prototype for cryptographically verifiable row-change history. Clients sign transaction changesets, a background worker packs them into Merkle-tree blocks, and block hashes can optionally be published to an Ethereum smart contract. Upstream explicitly says it is not production-ready.

### Core Workflow

Preload the library, choose the single database and extension schema to manage, then restart the server:

```conf
shared_preload_libraries = 'pg_credereum'
pg_credereum.database = 'appdb'
pg_credereum.schema = 'public'
pg_credereum.block_period = 1000
```

Create the extension and attach its accumulator trigger to tables with exactly one `bigint` primary-key column named `id`:

```sql
CREATE EXTENSION pg_credereum;

CREATE TABLE t (id bigint PRIMARY KEY, value integer NOT NULL);
CREATE TRIGGER t_after
AFTER INSERT OR UPDATE OR DELETE ON t
FOR EACH ROW EXECUTE PROCEDURE credereum_acc_trigger();
REVOKE TRUNCATE ON t FROM PUBLIC;
```

Within a transaction, perform DML, call `credereum_get_changeset()` and verify the returned proof, then sign it with `credereum_sign_transaction(pubkey text, sign bytea)` before committing.

### API and Configuration

- `credereum_get_changeset()`: returns Merkle-proof rows for changes made by the current transaction.
- `credereum_sign_transaction(text, bytea)`: records the client's public key and signature.
- `credereum_merkle_proof(varbit[])`: returns history/proof data for selected Merkle keys.
- `credereum_tx_log` and `credereum_block`: transaction and block history.
- `pg_credereum.block_retry_period`: retry interval, default `5000` ms.
- `pg_credereum.eth_end_point`, `pg_credereum.eth_source_addr`, and `pg_credereum.eth_contract_addr`: optional Ethereum publication settings.

### Caveats

The extension requires PostgreSQL 10 plus its documented Ethereum, CURL, and Jansson build dependencies. It manages only the database named by `pg_credereum.database`. `TRUNCATE` is not tracked, and the same row cannot be modified twice within one block; such a second modification raises a unique-constraint error. A shorter `block_period` can reduce that collision window.

If publishing a block hash to Ethereum fails, the worker skips that hash and tries the next block, so the trusted store can legitimately have gaps. Treat this project as a research prototype, independently verify every proof/signature path, and never use its existence alone as evidence that data is untampered.
