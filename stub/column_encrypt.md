


## Usage

Sources: [README](https://github.com/vibhorkum/column_encrypt/blob/master/README.md), [v4.0 release](https://github.com/vibhorkum/column_encrypt/releases/tag/v4.0), [SQL objects](https://github.com/vibhorkum/column_encrypt/blob/master/column_encrypt--4.0.sql)

`column_encrypt` provides transparent column-level encryption for PostgreSQL. It defines `encrypted_text` and `encrypted_bytea` types, encrypts values through type input functions, decrypts through output functions, and manages data-encryption keys through the `encrypt` schema.

### Enable

Load the shared library at server start, restart PostgreSQL, then create the schema and extension:

```conf
shared_preload_libraries = 'column_encrypt'
```

```sql
CREATE EXTENSION pgcrypto;
CREATE SCHEMA IF NOT EXISTS encrypt;
CREATE EXTENSION column_encrypt;
```

Add `encrypt` to `search_path` or schema-qualify the encrypted types and functions.

### Register And Load Keys

```sql
SELECT encrypt.register_key('my-secret-data-key', 'my-master-passphrase');
SELECT encrypt.load_key('my-master-passphrase');

SELECT * FROM encrypt.keys();
SELECT * FROM encrypt.status();
```

The extension uses a two-tier key model with key-encryption keys and data-encryption keys. Ciphertext carries a key-version header so older values can still be decrypted after rotation.

### Encrypt Columns

```sql
CREATE TABLE secure_data (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  ssn encrypt.encrypted_text,
  payload encrypt.encrypted_bytea
);

INSERT INTO secure_data (ssn, payload)
VALUES ('888-999-2045', decode('aabbcc', 'hex'));

SELECT id, ssn FROM secure_data;
```

Without a loaded key, decrypting encrypted values raises an error.

### Key Operations

Common functions include `encrypt.activate_key`, `encrypt.revoke_key`, `encrypt.rotate`, `encrypt.verify`, `encrypt.unload_key`, `encrypt.loaded_cipher_key_versions`, and `encrypt.blind_index`.

Use blind indexes for lookup patterns that cannot expose plaintext values directly:

```sql
SELECT encrypt.blind_index('888-999-2045', 'lookup-hmac-key');
```

### Notes

The extension intentionally rejects binary send/receive for encrypted values. Equality and hash semantics are based on decrypted plaintext; range ordering is not supported. After upgrading from older ciphertext-hash behavior, rebuild hash indexes on encrypted columns.
