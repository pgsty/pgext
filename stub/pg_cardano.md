


## Usage

> [pg_cardano: Cardano blockchain data types and cryptographic functions for PostgreSQL](https://github.com/Fell-x27/pg_cardano)

The `pg_cardano` extension provides cryptographic and utility functions for working with Cardano blockchain data in PostgreSQL, including Base58, Bech32, CBOR, Blake2b, Ed25519, and Shelley address utilities.

```sql
CREATE EXTENSION pg_cardano;
```

### Base58 Encoding/Decoding

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
-- '3Z6ioYHE3x'

SELECT cardano.base58_decode('3Z6ioYHE3x');
-- '\x43617264616e6f'
```

### Bech32 Encoding/Decoding

```sql
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
-- 'ada1d9ejqctdv9axjmn8dypl4d'

SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
-- 'ada'

SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
-- '\x697320616d617a696e67'
```

### CBOR Encoding/Decoding

Simple CBOR (lightweight, sufficient for most Cardano tasks):

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
```

Extended CBOR (full support, no limitations):

```sql
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

### Blake2b Hashing

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
-- '\x2244d5c9699fa93b...'
```

### Ed25519 Signing and Verification

```sql
SELECT cardano.ed25519_sign_message(
    '\x43D68AEC...'::bytea,       -- signing key
    'Cardano is amazing!'::bytea   -- message
);

SELECT cardano.ed25519_verify_signature(
    '\x432753BD...'::bytea,        -- verification key
    'Cardano is amazing!'::bytea,  -- message
    '\x74265f96...'::bytea         -- signature
);
-- true
```

### dRep View ID Builders (CIP-105/CIP-129)

```sql
SELECT cardano.tools_drep_id_encode_cip105('\x28111ae1...'::bytea, FALSE);
-- 'drep_vkh19qg34ctllr7lh48nnj4akfc978qzqzuwzkgsdu6r3zc42lnl6a0'

SELECT cardano.tools_drep_id_encode_cip129('\x28111ae1...'::bytea, TRUE);
-- 'drep1yv5pzxhp0lu0m7757ww2hke8qhcuqgqt3c2ezphngwytz4gj324g7'
```

### Shelley Address Utilities

```sql
-- Build a base address
SELECT cardano.tools_shelley_address_build(
    '\x7415251f...'::bytea, FALSE,  -- payment cred, is_script
    '\x7c3ae2f2...'::bytea, FALSE,  -- stake cred, is_script
    0                                -- network id
);

-- Extract payment/stake credentials
SELECT cardano.tools_shelley_addr_extract_payment_cred('addr_test1qp6p2fgl...');
SELECT cardano.tools_shelley_addr_extract_stake_cred('addr_test1qp6p2fgl...');

-- Get address type
SELECT cardano.tools_shelley_addr_get_type('addr_test1vp6p2fgl...');
-- 'PMT_KEY:NONE'
```

### Asset Name Conversion

```sql
SELECT cardano.tools_read_asset_name('hello'::bytea);
-- 'hello' (valid UTF-8 returned as string)

SELECT cardano.tools_read_asset_name('\xdeadbeef'::bytea);
-- 'deadbeef' (non-UTF-8 returned as hex)
```

### CIP-88 Pool Key Registration Verification

```sql
SELECT cardano.tools_verify_cip88_pool_key_registration('\xa119...'::bytea);
-- true
```
