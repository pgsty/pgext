## Usage

Sources:

- [Upstream README and SQL signature](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/README.md)
- [Extension control file](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/pg_ecdsa_verify.control)
- [Cargo package and PostgreSQL feature matrix](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/Cargo.toml)
- [Rust implementation](https://github.com/joelonsql/pg_ecdsa_verify/blob/6cca981b2019192945613096eebd89f62c0c5b23/src/lib.rs)

`pg_ecdsa_verify` provides one Rust/pgrx function for ECDSA signature verification. `ecdsa_verify.ecdsa_verify(bytea, bytea, bytea, text, text)` hashes the payload, verifies the signature against a public key, and returns a Boolean. It supports the `secp256r1` and `secp256k1` curves and only the `sha256` hash selector.

### Verify stored signed data

```sql
CREATE EXTENSION pg_ecdsa_verify;

SELECT ecdsa_verify.ecdsa_verify(
  public_key,
  payload,
  signature,
  'sha256',
  'secp256r1'
)
FROM signed_messages
WHERE message_id = 42;
```

The public key is a raw 64-byte concatenation of the X and Y coordinates, and the signature is a raw 64-byte concatenation of the R and S values. The extension intentionally does not generate keys or create signatures, so private keys need not enter PostgreSQL.

The implementation slices both byte arrays without checking their lengths and panics for unsupported curve/hash names. Validate `octet_length(public_key) = 64` and `octet_length(signature) = 64` before calling it, and treat malformed input as an error path rather than a false signature. The catalog and Cargo package agree on version `1.2.4`; Cargo features cover PostgreSQL 13 through 18. Installation is superuser-only and creates the fixed `ecdsa_verify` schema.
