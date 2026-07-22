## Usage

Sources:

- [Official README](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/README.md)
- [Official extension control file](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/pg_chainsync.control)
- [Official Rust package manifest](https://github.com/usagi-coffee/pg_chainsync/blob/c59b7a879dd427bb04f855c29a10979561f4fa08/Cargo.toml)

`pg_chainsync` is a proof-of-concept PostgreSQL background worker for fetching blockchain blocks, events, and scheduled tasks and passing them to user-defined SQL handlers. The reviewed upstream README warns that breaking changes should be expected, so use it for controlled experimentation rather than assuming a stable ingestion contract.

### Preload and Lifecycle

The worker requires `pg_chainsync.so` in `shared_preload_libraries` and a PostgreSQL restart. Set `chainsync.database` to the database in which the worker should operate, then create the extension there:

```conf
shared_preload_libraries = 'pg_chainsync.so'
chainsync.database = 'chains'
```

```sql
CREATE EXTENSION pg_chainsync;

SELECT chainsync.restart();
SELECT chainsync.stop();
```

`chainsync.restart()` reloads registered jobs into the worker; `chainsync.stop()` stops their execution.

### Register a Block Handler

A handler receives a typed chain value plus the job's `jsonb` metadata. For an EVM block watcher:

```sql
CREATE TABLE evm_blocks (
  block_number bigint PRIMARY KEY,
  block_hash text NOT NULL
);

CREATE FUNCTION store_block(block chainsync.EvmBlock, job jsonb)
RETURNS evm_blocks
LANGUAGE sql
AS $$
  INSERT INTO evm_blocks(block_number, block_hash)
  VALUES (block.number, block.hash)
  ON CONFLICT (block_number) DO UPDATE SET block_hash = EXCLUDED.block_hash
  RETURNING *
$$;

SELECT chainsync.register(
  'ethereum_blocks',
  '{
    "ws": "wss://provider.example/ws",
    "evm": {"block_handler": "store_block"}
  }'::jsonb
);

SELECT chainsync.restart();
```

The README also documents `chainsync.EvmLog` event handlers, one-shot work, cron tasks, preloaded tasks, and a block-await hook used to coordinate event processing with block availability. Confirm the installed `chainsync.register` signature and job JSON schema because this project is explicitly unstable.

### Configuration and Reliability

- `chainsync.evm_ws_permits` bounds concurrent tasks sharing an EVM WebSocket provider.
- `chainsync.evm_blocktick_reset` controls when reduced block ranges are reset.
- `chainsync.svm_rpc_permits` bounds concurrent SVM RPC fetches within a task.
- `chainsync.svm_signatures_buffer` limits buffered SVM signatures.
- Provider rate limits, disconnects, historical range limits, chain finality, and reorganizations are external failure modes. Persist checkpoints and chain identifiers, and design handlers to be idempotent under retries and replay.
- A successful callback is not by itself proof of finality. Define a confirmation policy and a reorg correction path before using synchronized data for irreversible decisions.
- Handler functions execute database writes under the worker's privileges. Constrain ownership and `search_path`, schema-qualify objects, validate job JSON, protect provider credentials, and test transaction/error behavior.
