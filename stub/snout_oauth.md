## Usage

Sources:

- [README.md](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/README.md)
- [Cargo.toml](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/Cargo.toml)
- [src/lib.rs](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/src/lib.rs)
- [snout_oauth.control](https://github.com/snoutdata/snout-oauth/blob/997f4874afd74d5d287bc01dda0473fc0d324f00/snout_oauth.control)

`snout_oauth` 0.1.0 is a PostgreSQL 18 OAuth validator for Snout-issued database tokens. It validates JWTs against a local public JWKS file; it neither contacts an issuer over HTTP nor installs SQL objects.

### Core Workflow

```ini
oauth_validator_libraries = 'snout_oauth'
snout_oauth.issuer = 'https://auth.example.com'
snout_oauth.audience = 'project-ref'
snout_oauth.keys_file = '/etc/postgresql/oauth-jwks.json'
```

### Operational Boundaries

Set the issuer, project audience and public-key file before enabling OAuth authentication. The file must contain public EC P-256 keys; private keys are rejected. Tokens must pass issuer, audience, signature, expiry and lifetime checks, carry `token_use` equal to `db`, and bind `db_role` to the requested database role. This is not a general-purpose validator for arbitrary OIDC tokens.

Reload server configuration and authentication rules after changes. The library is loaded on demand by `oauth_validator_libraries`; shared preload is optional, and no `CREATE EXTENSION` step is supported. The control file is for development tooling only. Role mapping can be delegated because the validator checks the database role itself. Restrict matching HBA rules to intended networks and users.

Key files are reread when their identity or metadata changes; replace them atomically for rotation. Default clock leeway is 30 seconds and maximum token lifetime is one day. Empty required settings fail authentication. PostgreSQL 18 OAuth-capable clients are required.
