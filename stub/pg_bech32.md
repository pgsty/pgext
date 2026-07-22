## Usage

Sources:

- [Pinned upstream README](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/README.md)
- [Pinned Rust SQL implementation](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/src/lib.rs)
- [Pinned Rust package manifest](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/Cargo.toml)
- [Pinned extension control template](https://github.com/KaiserKarel/pg_bech32/blob/a4eb4aaae249540b9839cd8735cb450bad1ee8eb/pg_bech32.control)

`pg_bech32` version `0.0.0` adds Bech32, Bech32m, and checksum-free Bech32 encoding to PostgreSQL, plus decoding for checksummed Bech32 strings. It accepts binary payloads as `bytea` and returns decoded data together with the human-readable prefix.

### Core Workflow

```sql
CREATE EXTENSION pg_bech32;

SELECT bech32_encode(
  'union',
  decode('644a2606654a7c0e70bf343ae6b828d3fe448447', 'hex'),
  'bech32'
);

SELECT (bech32_decode(
  'union1v39zvpn9ff7quu9lxsawdwpg60lyfpz8pmhfey'
)).*;

SELECT bech32_encode_lower(
  'example',
  convert_to('payload', 'UTF8'),
  'bech32m'
);
```

### Important Objects

- `bech32_encode(hrp, bytea, mode)`: encodes with mode `bech32`, `bech32m`, or `nochecksum`.
- `bech32_encode_lower(hrp, bytea, mode)`: produces a lower-case encoded string using the same modes.
- `bech32_decode(text)`: validates and decodes a checksummed Bech32 or Bech32m string.
- Composite type `bech32`: contains fields `hrp text` and `data bytea`; the unquoted upstream type declaration is folded to lower case by PostgreSQL.

### Operational Notes

The decoder does not return which checksum variant was used, and the documented decoder path is for checksummed Bech32/Bech32m rather than `nochecksum`; keep protocol/version context outside the value. Invalid HRPs, strings, checksums, or mode names raise errors from the Rust implementation. The control template marks installation as superuser-only and non-relocatable.

The Cargo feature matrix targets PostgreSQL 11–16 using pgrx `0.11.3`; later PostgreSQL releases are not confirmed by this revision. Version `0.0.0` and the minimal upstream documentation indicate an early project surface, so verify canonical casing, maximum lengths, protocol-specific witness/version rules, and round trips with authoritative test vectors before using encoded values as durable identifiers.
