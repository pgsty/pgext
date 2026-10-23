## Usage

Sources:

- [Official source](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/README.md)
- [Official source](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/pg_oauth_rs.control)
- [Official source](https://github.com/irfanfaraaz/pg-oauth-rs/blob/6c9e9a79eb1854dcac5bd03536def51b644e2352/src/entry.rs)

`pg_oauth_rs` 0.2.0 is a pre-alpha OAuth validator for PostgreSQL 18 authentication. It validates signed JWTs through the server's OAuth validator interface; it does not expose a SQL extension API. The native library requires libcurl.

### Enablement

Configure the library and the intended identity provider. These example URLs and audience must be replaced with values issued by your provider:

```ini
oauth_validator_libraries = 'pg_oauth_rs'
pg_oauth_rs.provider = 'generic'
pg_oauth_rs.audience = 'postgres'
pg_oauth_rs.jwks_uri = 'https://identity.example.com/.well-known/jwks.json'
```

Initial setup follows the upstream restart procedure. Configure OAuth authentication in `pg_hba.conf`, including the issuer and required scope, and map the validated identity to a database role through `pg_ident.conf`. The provider, audience and identity mapping must agree; possession of a signed token alone does not establish the intended database authorization.

For diagnostics in an administrative session:

```sql
LOAD 'pg_oauth_rs';
SHOW pg_oauth_rs.provider;
```

### Validation and Identity

JWT signatures use RS256 or ES256. Validation checks issuer, audience, expiry and not-before claims. Provider presets cover generic issuers, Microsoft Entra ID, Okta and Keycloak. The generic identity uses the subject claim; Entra's stable object identifier avoids relying on a renameable user name. Email identities require a verified-email claim. Opaque access tokens are unsupported.

`pg_oauth_rs.identity_claim` selects the mapping claim. `pg_oauth_rs.clock_skew_leeway`, `pg_oauth_rs.jwks_timeout`, `pg_oauth_rs.jwks_refetch_interval`, `pg_oauth_rs.jwks_max_age` and `pg_oauth_rs.jwks_max_stale` control clock tolerance and key fetching/cache lifetime.

### Operational Boundary

Signing keys are cached across backends. Consider the permitted stale-key interval when planning emergency key revocation; an issuer outage and a revoked key are different cases. Keep the key endpoint authenticated with HTTPS. The client must support PostgreSQL OAuth authentication; do not assume older clients or PostgreSQL 19 roadmap features work with this release.
