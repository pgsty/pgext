


## Usage

> [pgcryptokey: Cryptographic key management for PostgreSQL](https://momjian.us/download/pgcryptokey/)

`pgcryptokey` manages cryptographic data encryption keys within PostgreSQL. Keys are stored encrypted and secured by access passwords, supporting both system-wide and per-session key access.

```sql
CREATE EXTENSION pgcryptokey;
```

### Key Management Functions

| Function | Description |
|----------|-------------|
| `create_cryptokey(name, byte_len)` | Generate a new cryptographic key |
| `set_cryptokey(name)` | Set the active key for operations |
| `get_cryptokey(name)` | Retrieve key material |
| `drop_cryptokey(name)` | Remove a key |
| `supersede_cryptokey()` | Rotate to a new key (same access password) |
| `change_key_access_password()` | Update key authentication credentials |
| `reencrypt_data()` | Re-encrypt data with a different key |

### Session Control

| Function | Description |
|----------|-------------|
| `get_shared_key()` | Establish client/server shared secret (SSL/Unix only) |
| `set_session_access_password()` | Client-supplied password authentication |

### Typical Workflow

```sql
-- Create a key
SELECT create_cryptokey('mykey', 32);

-- Set active key
SELECT set_cryptokey('mykey');

-- Encrypt data using pgcrypto functions with the managed key
UPDATE secrets SET data = pgp_sym_encrypt(plaintext, get_cryptokey('mykey'));

-- Decrypt data
SELECT pgp_sym_decrypt(data, get_cryptokey('mykey')) FROM secrets;

-- Rotate key
SELECT supersede_cryptokey();
```

Access passwords can be configured at database boot time for system-wide access, or per-session by individual clients for granular security control.
