## Usage

Sources:

- [README.md](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/README.md)
- [src/lib.rs](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/src/lib.rs)
- [Cargo.toml](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/Cargo.toml)
- [META.json](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/META.json)
- [encsearch.control](https://github.com/lokeshmulpuru/encsearch/blob/cfa1a2af334ebbd4026a9f467c0ec40bed12a035/encsearch.control)

`encsearch` stores encrypted text alongside keyed search tokens in an `int8[]` column. Its experimental 0.1.0 release supports exact, word and pattern searches through a GIN index. Upstream has tested PostgreSQL 17; the other Cargo feature targets are untested.

### Core Workflow

```ini
session_preload_libraries = 'encsearch'
```

```sql
CREATE EXTENSION encsearch;
CREATE TABLE encrypted_notes (id bigint GENERATED ALWAYS AS IDENTITY, body bytea, idx int8[]);
CREATE INDEX encrypted_notes_tokens ON encrypted_notes USING gin (idx);
INSERT INTO encrypted_notes (body, idx)
VALUES (enc_encrypt('Patient has diabetes'), enc_index('Patient has diabetes'));
SELECT enc_decrypt(body) FROM encrypted_notes
WHERE enc_like(body, idx, '%diab%');
```

### Operational Boundaries

An administrator installs this untrusted extension. Configure `session_preload_libraries` before opening client sessions and before setting `encsearch.key`; otherwise an early key setting remains an unregistered setting visible to other roles. Supply a securely retained key of at least 32 characters to the session or explicit-key function overloads before running the data example. `enc_gen_key()` generates a 64-character key. Losing the key loses access to the plaintext.

`enc_encrypt` and `enc_decrypt` handle text and ciphertext. `enc_index` creates stored tokens; `enc_like`, `enc_ilike` and `enc_eq` filter candidates and check decrypted values. `enc_has_words` tests token containment without decrypting; `enc_q_like`, `enc_q_exact` and `enc_q_words` expose raw query tokens. `enc_algorithm` and `enc_algorithms` inspect available ciphers.

`encsearch.algorithm` defaults to `aes-256-gcm`; deterministic `aes-256-siv` reveals duplicate plaintext. Keep `encsearch.token_hash` identical when building and querying tokens: `hmac-sha256` is the default and `blake3` is the alternative. Changing the token hash requires regenerating tokens; changing keys requires decrypting with the old key and rebuilding ciphertext and tokens with the new key.

Blind indexes reveal shared values, words and character chunks and their frequency. Keys reach the database server and may appear in statement or parameter logs. This is an experimental privacy/performance tradeoff, not protection against a server that has the key.
