## Usage

Sources:

- [Official 0.85 source archive](https://momjian.us/download/pgcryptokey/pgcryptokey-0.85.tar.gz)
- [Official project directory](https://momjian.us/download/pgcryptokey/)

`pgcryptokey` manages encryption data keys wrapped by an access password. It stores wrapped keys in a database table and integrates with `pgcrypto` for encryption, key rotation and re-encryption.

### Install and Unlock

```sql
CREATE EXTENSION pgcryptokey CASCADE;
```

The dependency is `pgcrypto`. Source release 0.85 uses SQL extension version 1.0 and needs superuser installation. In client mode, first establish the documented session access password using `get_shared_key()` and `set_session_access_password(encrypted_password)`. The shared-key exchange supports SSL or Unix-domain socket connections only; the encrypted password argument is hex-encoded.

Boot mode instead preloads `pgcryptokey_acpass`, runs the protected server-side password acquisition script and requires a restart. It makes the access password server-wide and read-only. Choose one mode; boot and client modes cannot be combined while the server is running.

### Create and Use a Key

After unlocking access to the keys:

```sql
SELECT create_cryptokey('app-key', 32);
SELECT set_cryptokey('app-key');

CREATE TEMP TABLE secrets(ciphertext bytea);
INSERT INTO secrets VALUES
  (pgp_sym_encrypt('example', get_cryptokey('app-key')));
SELECT pgp_sym_decrypt(ciphertext, get_cryptokey('app-key'))
FROM secrets;
```

The key length is in bytes. Keys may be selected by name or integer key ID; name lookup refers to an active, non-superseded key.

### Rotate and Re-encrypt

`supersede_cryptokey(name, byte_len)` or its key-ID overload creates a replacement and returns its ID. The old and new keys initially use the same access password. Use the integer key-ID overload `change_key_access_password(key_id, new_encrypted_password)` to change that wrapping password. The session must already have its shared key and current access password set; the new password must be encrypted with the shared key and hex-encoded.

In source release 0.85, the name overload calls an undefined `change_access_password` function and cannot complete the password change. Use the integer overload above.

`reencrypt_data(data, old_key_id, new_key_id)` and `reencrypt_data_bytea(data, old_key_id, new_key_id)` migrate encrypted values. Preserve old key IDs alongside ciphertext and verify re-encryption before calling `drop_cryptokey(name)` or its key-ID overload; removing a key can make remaining ciphertext unreadable.

### Security Boundary

Protect the key table, function grants, access-password acquisition script and backups. Upstream warns that all users can view the boot-time `pgcryptokey.access_password`; table privileges are still required to use wrapped keys. `get_cryptokey(name)` returns raw key material, so do not expose its result through ordinary queries, logs or application traces. This design does not isolate keys from a trusted database administrator.
