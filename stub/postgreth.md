## Usage

Sources:

- [Official README](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/README.md)
- [Official Rust implementation](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/src/lib.rs)
- [Official package and PostgreSQL feature matrix](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/Cargo.toml)
- [Official control file](https://github.com/KaiserKarel/postgreth/blob/c86650deaaf1f86f7ae9931c339e24e44143a0b0/postgreth.control)

`postgreth` is an experimental pgrx extension for decoding Ethereum ABI values and logs and for working with Ethereum bloom filters. It does not connect to an Ethereum node, verify chain data, or synchronize tables; callers must supply already-fetched bytes, ABI JSON, and log JSON.

### Core Workflow

For a single ABI-encoded scalar, provide its Solidity type and bytes:

```sql
CREATE EXTENSION postgreth;

SELECT item_to_jsonb(
  'uint256',
  decode(repeat('00', 31) || '2a', 'hex'),
  false
);
```

`item_to_jsonb(text,bytea,boolean)` returns decoded `jsonb`. The final flag prepends an ABI offset word for tuple-like values that require it; it is not a general repair option for malformed input.

### Types and Functions

- `Bloom`, `Address`, and `B256` represent a 256-byte bloom filter, 20-byte address, and 32-byte value.
- `covers(Bloom,Bloom)`, `contains_input(Bloom,bytea)`, and `contains_input_hashed(Bloom,bytea)` test bloom membership/coverage.
- `m3_2048(Bloom,bytea)` and `m3_2048_hashed(Bloom,bytea)` add raw or pre-hashed input to a bloom filter.
- `contains_raw_log(Bloom,Address,B256[])` tests an address and topics against a bloom filter.
- `log_to_jsonb(text,text)` decodes receipt-log JSON with an event ABI and returns keyed `jsonb`.

Bloom membership is probabilistic: a positive result can be a false positive and must not replace exact log matching or chain verification.

### Safety and Compatibility

Several parsing paths use Rust `expect`/`unwrap`. Invalid ABI JSON, malformed log JSON, undecodable bytes, an invalid Solidity type, or a hash whose length is not exactly 32 bytes can raise an error or panic through the extension boundary. Validate untrusted data before invoking the functions and fuzz-test failure behavior on the target backend.

The control file requires superuser installation and marks the extension non-relocatable. The reviewed crate is version `0.0.1`, uses pgrx `0.11.3`, defaults to PostgreSQL 13, and defines build features for PostgreSQL 11–16. Those feature declarations are not release certification; build and test the selected feature against the exact server.

Upstream explicitly says the project is not for public use and lacks extensive testing, CI/CD, and a release process. Treat it as source for controlled experiments or an internally audited fork, not a production-ready blockchain trust boundary.
