## Usage

Sources:

- [Official README and API reference](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/README.md)
- [Official extension control file](https://github.com/truthly/pg-webauthn/blob/c1f8099db1cbed3a89cf12a2668ed67bb10ee354/webauthn.control)
- [WebAuthn Level 2 specification](https://www.w3.org/TR/webauthn-2/)

`webauthn` is a pure-SQL PostgreSQL implementation of WebAuthn credential registration and assertion verification. It stores one-time challenges and public-key credentials in the fixed `webauthn` schema, while the application still owns the browser ceremony, HTTPS origin, relying-party policy, authorization, and account-recovery design.

### Prerequisites and Registration

Version `1.6` depends on `pgcrypto`, `pg_ecdsa_verify`, and `cbor`. Install it with dependency cascading only after those extensions and their trust boundary have been reviewed:

```sql
CREATE EXTENSION webauthn CASCADE;

SELECT webauthn.init_credential(
  challenge := gen_random_bytes(32),
  user_name := 'alex@example.com',
  user_id := gen_random_bytes(32),
  user_display_name := 'Alex',
  relying_party_name := 'Example'
);
```

Send the returned `publicKey` object to `navigator.credentials.create()`. Pass the browser response to `webauthn.store_credential()`; on success it stores the public key and returns the registered `user_id`, while `NULL` means failure. Each challenge is single-use.

### Authentication

Start sign-in with a fresh server-generated challenge:

```sql
SELECT webauthn.get_credentials(
  challenge := gen_random_bytes(32),
  user_name := 'alex@example.com',
  relying_party_id := 'example.com'
);
```

Send that result to `navigator.credentials.get()`, then pass the assertion fields to `webauthn.verify_assertion()`. A verified assertion returns the credential's `user_id`; reject `NULL`. Username-less sign-in is available only for credentials created with `require_resident_key := TRUE`, and then `user_handle` identifies the user during verification.

### API Index

- `webauthn.init_credential()` creates registration options and stores the registration challenge.
- `webauthn.store_credential()` verifies and stores a browser-created credential.
- `webauthn.get_credentials()` creates assertion options and stores the sign-in challenge.
- `webauthn.verify_assertion()` verifies the browser assertion and returns its `user_id`.

Both option-producing functions default to a five-minute timeout; upstream restricts custom timeouts to 30 seconds through 10 minutes. Set the relying-party ID deliberately, generate challenges with a cryptographically secure server-side source, restrict direct access to challenge and credential tables, and treat the returned identity as authentication input rather than application authorization. Test supported authenticators, backup and recovery, cloned databases, challenge cleanup, failure paths, and the complete origin/RP policy before production deployment.
