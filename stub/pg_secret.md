## Usage

Sources:

- [Official design notes](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/NOTES)
- [Official extension SQL](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/pg_secret--0.5.sql)
- [Official C implementation](https://github.com/TypedLambda/pg_secret/blob/eef18ef9f504fa55fb19011e9f4844495f89709a/pgsecret.c)

`pg_secret` 0.5 is an experimental order-revealing encryption prototype. It turns an integer or text value into the `secret` domain and supplies comparison operators plus a B-tree operator class. Use it only to study indexed ordering over ciphertext; it is not a general secret-management or production encryption system.

### Core Workflow

`make_secret(key, key, int8)` encrypts a 64-bit integer. The text overload first converts the input to a 64-bit SipHash value, so its encrypted ordering is hash ordering, not lexical text ordering. Both keys must be exactly 16 bytes even though the `key` domain itself does not enforce a length.

```sql
CREATE EXTENSION pg_secret;

CREATE TABLE secure_rankings (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    score secret NOT NULL
);

WITH keys AS (
    SELECT decode('00112233445566778899aabbccddeeff', 'hex')::key AS k1,
           decode('ffeeddccbbaa99887766554433221100', 'hex')::key AS k2
)
INSERT INTO secure_rankings (score)
SELECT make_secret(k1, k2, value::int8)
FROM keys, unnest(ARRAY[10, 30, 20]) AS v(value);

CREATE INDEX secure_rankings_score_idx
    ON secure_rankings USING btree (score);

SELECT id
FROM secure_rankings
ORDER BY score;
```

Applications must retain and protect both keys outside this table. Recreating a comparison boundary requires calling `make_secret` with the same pair of keys.

### Objects

- Domains: `secret` is a `bytea` value constrained to 224 bytes; `key` is an unconstrained `bytea` domain.
- Constructors: `make_secret(key, key, int8)` and `make_secret(key, key, text)`.
- Comparisons: `<`, `<=`, `=`, `>=`, and `>` for `secret`, implemented through `secret_cmp` and the default B-tree operator class `secret_btree_ops`.

There is no decryption function, authenticated-encryption API, key registry, or key-rotation workflow in version 0.5.

### Security and Operational Boundaries

Order-revealing encryption deliberately leaks order, and the upstream notes call this implementation a weaker security model. The server receives the plaintext and both keys; PostgreSQL logging, error reporting, memory handling, or WAL-related behavior may expose sensitive material. The implementation also contains unresolved key-cleanup work.

The source was an early 2019–2020 experiment and depends on bundled cryptographic code plus OpenSSL/GMP-era build assumptions. Perform independent cryptographic review, restrict execution privileges, and test the exact PostgreSQL build. Do not use `pg_secret` as a substitute for audited encryption, KMS integration, or application-side key custody.
