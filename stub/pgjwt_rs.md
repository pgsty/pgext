## Usage

Sources:

- [Official README](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/README.md)
- [Rust implementation](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/src/lib.rs)
- [Extension SQL](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/sql/pgjwt_rs--0.1.2.sql)
- [Rust package manifest](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/Cargo.toml)
- [Extension control file](https://github.com/vishvish/pgjwt/blob/ddaab1e4d47b78c91812dd757f3b878d43e0bad4/pgjwt_rs.control)

`pgjwt_rs` 0.1.2 verifies asymmetric JWT signatures inside PostgreSQL using RS256 or Ed25519. It deliberately performs cryptographic verification only: expiration, not-before, audience, issuer, subject, nonce, and application authorization remain the caller's responsibility.

### Verify Before Using Claims

Store trusted public keys and the server-selected algorithm separately from untrusted tokens. Join a token to a key through a trusted key identifier or issuer mapping, require `valid`, then apply every claim required by the application. For example, given application-owned `auth.incoming_token` and `auth.jwt_key` tables:

```sql
CREATE EXTENSION pgjwt_rs;

SELECT v.header, v.payload
FROM auth.incoming_token AS i
JOIN auth.jwt_key AS k ON k.key_id = i.key_id
CROSS JOIN LATERAL jwt_verify(i.token, k.public_key_pem, k.algorithm) AS v
WHERE i.id = 42
  AND v.valid
  AND (v.payload->>'iss') = 'https://issuer.example'
  AND (v.payload->>'aud') = 'orders-api'
  AND now() < to_timestamp((v.payload->>'exp')::bigint);
```

Do not choose `k.algorithm` solely from the token header. The verifier checks that the signed header matches the expected algorithm, but the application must supply that expectation from trusted configuration.

### User-Facing Functions

- `jwt_verify_rs256(text, text)`: verify an RS256 signature with an SPKI PEM public key and return `header`, `payload`, and `valid`.
- `jwt_verify_ed25519(text, text)`: verify an Ed25519/EdDSA signature with an SPKI PEM public key.
- `jwt_verify(text, text, text)`: dispatch using `RS256`, `EdDSA`, or `ED25519` supplied by trusted code.
- `jwt_decode_payload(text)`: decode payload JSON without checking a signature; the returned object is marked with `_unverified`.

The public key must use the `PUBLIC KEY` PEM label, not `RSA PUBLIC KEY`, `RSA PRIVATE KEY`, or another container. Tokens larger than 16 KiB and keys larger than 8 KiB are rejected.

### Security Boundaries

`jwt_decode_payload` is only a parsing aid for locating a candidate trusted key. Never authorize, select a tenant, or construct unrestricted key lookups from its result. The verification functions decode and return a payload even when `valid` is false, so callers must test `valid` before reading any claim.

Version 0.1.2 does not enforce `exp`, `nbf`, or `aud`, and it does not require `exp` to exist. Add explicit policy for time claims, issuer, audience, subject, token type, clock skew, key rotation, and revocation. If wrapping these functions in `SECURITY DEFINER`, set a fixed safe `search_path`, schema-qualify objects, and revoke default execution from untrusted roles. The package advertises builds for PostgreSQL 13 through 18; validate the exact pgrx artifact on the target server.
