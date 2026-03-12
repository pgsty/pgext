


## Usage

> [pguecc: Elliptic curve cryptography (micro-ecc bindings) for PostgreSQL](https://github.com/ameensol/pg-ecdsa)

Requires the `pgcrypto` extension.

### Generate Key Pair

```sql
SELECT ecdsa_make_key('secp256k1');
-- (public_key_hex, private_key_hex)
```

### Sign Data

```sql
SELECT ecdsa_sign(
    '000000000000000000000000000000000000000000',  -- private key (hex)
    '1234',                                          -- data to sign
    'sha256',                                        -- hash function
    'secp160r1'                                      -- curve name
);
```

### Verify Signature

```sql
SELECT ecdsa_verify(
    '<public_key_hex>',
    'hello, world',
    '<signature_hex>',
    'sha256',
    'secp256k1'
);
-- t
```

### Key Validation

```sql
SELECT ecdsa_is_valid_public_key('<public_key_hex>', 'secp256k1');
SELECT ecdsa_is_valid_private_key('<private_key_hex>', 'secp256k1');
SELECT ecdsa_is_valid_curve('secp256k1');  -- true
```

### Supported Curves

`secp160r1`, `secp192r1`, `secp224r1`, `secp256r1`, `secp256k1`

### Raw API

For direct signing without hashing (use with caution):

```sql
SELECT ecdsa_sign_raw(private_key_bytea, hash_bytea, 'secp256k1');
SELECT ecdsa_verify_raw(public_key_bytea, hash_bytea, signature_bytea, 'secp256k1');
```
