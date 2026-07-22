## Usage

Sources:

- [Official eidos_oauth README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/README.md)
- [Extension control file](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/eidos_oauth.control)
- [Validator implementation](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/eidos_oauth/src/validator.rs)

`eidos_oauth` version `0.2.0` validates OAuth bearer JWTs for PostgreSQL 18 connection authentication. It verifies signatures through a remote JWKS endpoint, checks configured issuer and audience claims, returns an authentication identity, and exposes claims as session GUCs for authorization policies.

### Core Workflow

Configure the library at server start, then create the SQL extension and inspect the validator state:

```conf
shared_preload_libraries = 'eidos_oauth'
oauth_validator_libraries = 'eidos_oauth'

eidos_oauth.jwks_url = 'https://auth.example.com/.well-known/jwks.json'
eidos_oauth.issuer = 'https://auth.example.com'
eidos_oauth.audience = 'my-app'
eidos_oauth.role_claim = 'sub'
eidos_oauth.role_prefix = 'oauth_'
eidos_oauth.jwks_cache_seconds = 300
eidos_oauth.jwks_timeout_ms = 5000
```

```sql
CREATE EXTENSION eidos_oauth;

SELECT eidos_oauth.oauth_jwks_status();
SELECT eidos_oauth.oauth_refresh_jwks();
SELECT eidos_oauth.oauth_validate_token('signed.jwt.value');
SELECT current_setting('app.user_sub', true);
```

The SQL helpers are `oauth_validate_token`, `oauth_get_claim`, `oauth_inject_claims`, `oauth_jwks_status`, and `oauth_refresh_jwks`. Connection-time validation uses the configured claim, defaulting to `sub`, as the authenticated identity; PostgreSQL role mapping remains responsible for mapping that identity to a database role.

### Operational Notes

The connection validator API requires PostgreSQL 18. The source describes `oauth_inject_claims` as a testing or pre-18 helper, not as a replacement for PostgreSQL 18 OAuth authentication. The control file is non-relocatable and superuser-only. Configuration changes use SIGHUP-context GUCs, but adding the libraries requires a restart. Protect the JWKS endpoint, test key rotation and outage behavior, restrict refresh and token-testing functions, and write RLS policies that treat injected claims as authentication context rather than trusted application input.
