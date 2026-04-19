## Usage

Sources: [README](https://github.com/Fell-x27/pg_cardano/blob/master/README.md), [Cargo.toml version 1.2.0](https://github.com/Fell-x27/pg_cardano/blob/master/Cargo.toml)

`pg_cardano` is a Rust extension for Cardano-oriented binary and crypto utilities inside PostgreSQL. The upstream repo does not publish GitHub releases or tags, but the current crate version on the default branch is `1.2.0`.

```sql
CREATE EXTENSION pg_cardano;
```

### Core Capabilities

- Base58 encode/decode.
- Bech32 encode/decode.
- CBOR to and from `jsonb`, with both simple and extended pipelines.
- Blake2b hashing.
- Ed25519 signing and signature verification.
- dRep ID helpers for CIP-105 and CIP-129.
- Shelley address builders and parsers.
- Asset-name decoding and CIP-88 pool key registration verification.

### Common Patterns

Base58 and Bech32:

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
SELECT cardano.base58_decode('3Z6ioYHE3x');
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
```

CBOR conversion:

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

Hashing and signatures:

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
SELECT cardano.ed25519_verify_signature(
  '\x432753BD...'::bytea,
  'Cardano is amazing!'::bytea,
  '\x74265f96...'::bytea
);
```

Address and governance helpers:

```sql
SELECT cardano.tools_drep_id_encode_cip105('\x28111ae1...'::bytea, FALSE);
SELECT cardano.tools_drep_id_encode_cip129('\x28111ae1...'::bytea, TRUE);
SELECT cardano.tools_shelley_address_build(
  '\x7415251f...'::bytea, FALSE,
  '\x7c3ae2f2...'::bytea, FALSE,
  0
);
SELECT cardano.tools_shelley_addr_get_type('addr_test1vp6p2fgl...');
```

### Caveats

- Upstream documents PostgreSQL 12+ and Linux builds via `pgrx`; the crate features currently target PostgreSQL 15 through 18, with `pg18` as the default feature.
- The simple CBOR helpers are intentionally lossy for some CBOR constructs; use the `*_ext` functions when exact structural round-tripping matters.
