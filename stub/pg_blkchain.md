## Usage

Sources:

- [Official README](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/README.md)
- [Official extension SQL](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/pg_blkchain--0.0.1.sql)
- [Official C implementation](https://github.com/blkchain/pg_blkchain/blob/d90127e749a72c918fe3499cfc4d26a731ba30d5/pg_blkchain.c)

`pg_blkchain` is an experimental C extension for decoding and constructing legacy Bitcoin blocks, transactions, inputs, outputs, and scripts inside PostgreSQL. Version `0.0.1` is explicitly work in progress and the reviewed source stopped in 2017; do not treat its parsing or signature verification as a current consensus or security implementation.

### Core Workflow

Install the extension, keep raw serialized transactions in `bytea`, and expand their outputs for analysis:

```sql
CREATE EXTENSION pg_blkchain;

SELECT r.id, (get_vout(r.tx)).*
FROM raw_transactions AS r
WHERE r.id = 37898;
```

Scripts can be expanded into opcode records:

```sql
SELECT op_sym, op, encode(data, 'hex') AS data_hex
FROM parse_script(transaction_output_script);
```

The second example assumes `transaction_output_script` is a `bytea` value supplied by the caller.

### Important Objects

- `get_block(bytea)` returns `CBlock`; `get_tx(bytea)` returns `CTx`.
- `get_vin(bytea)` and `get_vout(bytea)` return sets of `CTxIn` and `CTxOut` records.
- `parse_script(bytea)` returns `CScriptOp` records containing the symbolic opcode, numeric opcode, and pushed data.
- `verify_sig(txTo bytea, txFrom bytea, int)` attempts to verify one transaction input signature.
- `build_vin(...)` and `build_vout(...)` are aggregates that serialize input and output components.
- `get_vin_outpt_arr(bytea)`, `get_vin_outpt_jsonb(bytea)`, and `get_vin_outpt_bytea(bytea)` expose previous-output references in several representations.

### Operational Caveats

The upstream warning says to use the project at your own risk. The code predates major Bitcoin and PostgreSQL changes, parses untrusted binary data in the database backend, and provides no maintained compatibility or consensus-validation guarantee. A malformed input or C bug can terminate a backend. The README's signature example also needs `pgcrypto` only for its double-SHA256 lookup query; it is not declared as an extension dependency. Use an actively maintained Bitcoin library outside PostgreSQL for validation, and keep this extension limited to isolated historical-data experiments.
