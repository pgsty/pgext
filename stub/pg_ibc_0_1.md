## Usage

Sources:

- [Official extension control file](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/pg_ibc_0_1.control)
- [Official Rust package manifest](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/Cargo.toml)
- [Official SQL API implementation](https://github.com/unionlabs/pg_ibc/blob/14abdfdfc6e0fb63500147ed16de14195b0610dc/src/lib.rs)

`pg_ibc_0_1` 0.1.1 provides immutable PostgreSQL functions used by Union's IBC indexing pipeline. It decodes transfer packets and UCS03 acknowledgements to JSONB, computes packet hashes, predicts wrapped-token addresses on several chains, and formats EIP-55 addresses. It does not connect to a blockchain, submit transactions, or validate consensus state.

### Core Workflow

Install the versioned extension name, then decode bytes already collected by an indexer. For Cosmos RPC data, the input is JSON bytes; setting the extension format to JSON converts a string-valued memo to nested JSON when possible.

```sql
CREATE EXTENSION pg_ibc_0_1;

SELECT decode_transfer_packet_0_1(
    convert_to('{"sender":"union1example","memo":"{\"forward\":{}}"}', 'UTF8'),
    'cosmos',
    false,
    'json'
);

SELECT erc55_to_checksum_0_1(
    decode('52908400098527886e0f7030069857d2e4169ee7', 'hex')
);
```

For EVM transfer packets, pass ABI-encoded packet bytes, `evm`, and either `string` or `json` for the extension field. With `throws = false`, ordinary EVM or Cosmos decode failures return JSON null; with `throws = true`, they raise an error.

### Packet and Acknowledgement Functions

- `decode_transfer_packet_0_1(bytea, text, boolean, text)` decodes the older EVM or Cosmos transfer packet shape.
- `decode_packet_0_1(bytea, text)` decodes a packet for channel version `ucs03-zkgm-0` and returns a tagged JSONB result.
- `decode_ack_0_1(bytea, bytea, text)` decodes an acknowledgement for channel version `ucs03-zkgm-0`.
- `decode_packet_ack_0_1` and `decode_packet_ack_0_2` hash packet fields and decode an optional acknowledgement; the latter accepts height and timestamp as text so values are not limited to signed 64-bit SQL integers.
- `decode_packet_ack_0_3` accepts packet data plus a supplied 32-byte packet hash and returns either decoded data or structured hashing/decoding error details.

### Address-Derivation Functions

- `erc55_to_checksum_0_1` requires exactly 20 address bytes and returns an EIP-55 string.
- `create3_0_1` predicts an EVM CREATE3 wrapped-token address.
- `instantiate2_0_1` predicts a CosmWasm Instantiate2 address using the checksum embedded by this release.
- `predict_cosmos_wrapper_0_1`, `predict_osmosis_wrapper_0_1`, and `predict_aptos_wrapper_0_1` derive chain-specific wrapper identifiers from channel paths and token bytes.

### Boundaries

The package exposes a snapshot of Union protocol formats. The packet decoders currently recognize only `evm`, `cosmos`, and channel version `ucs03-zkgm-0`; future protocol versions require a new extension or function version. An unsupported RPC type reaches an unimplemented Rust branch and raises an error even when `throws` is false.

All inputs are raw `bytea` and identifiers; the extension does not verify that a packet, acknowledgement, deployer, creator, or channel exists on-chain. Store the original bytes and chain metadata so results can be recomputed after upgrading. Version 0.1.1 declares build features only for PostgreSQL 14, 15, and 16.
