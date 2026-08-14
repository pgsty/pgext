## Usage

Sources:

- [pg_oidc_validator 1.1.0 README](https://github.com/percona/pg_oidc_validator/blob/1.1.0/README.md)
- [pg_oidc_validator 1.1.0 Keycloak example](https://github.com/percona/pg_oidc_validator/tree/1.1.0/examples/keycloak)
- [pg_oidc_validator 1.1.0 validator source](https://github.com/percona/pg_oidc_validator/blob/1.1.0/src/pg_oidc_validator.cpp)
- [PostgreSQL 18 OAuth authentication](https://www.postgresql.org/docs/18/auth-oauth.html)
- [PostgreSQL 18 libpq OAuth support](https://www.postgresql.org/docs/18/libpq-oauth.html)

`pg_oidc_validator` 1.1.0 is a PostgreSQL 18 OAuth validator module that validates JWT access tokens against an OpenID Connect provider. It is a server library with no control file or SQL extension, so do not run `CREATE EXTENSION`.

### Configure the Server

Load the module in `postgresql.conf`, then restart PostgreSQL:

```ini
oauth_validator_libraries = 'pg_oidc_validator'
```

Add an OAuth rule to `pg_hba.conf`; the issuer and required scope must match the provider. Use `hostssl` outside a strictly local test:

```text
hostssl  all  all  127.0.0.1/32  oauth  issuer=https://id.example.com/realms/postgres scope="openid postgres" validator=pg_oidc_validator
```

Reload PostgreSQL after HBA or validator-setting changes; adding the module to `oauth_validator_libraries` itself requires a restart.

The default authenticated identity claim is `sub`. To return another stable string claim for role matching, configure:

```ini
pg_oidc_validator.authn_field = 'email'
```

Version 1.1.0 also provides `pg_oidc_validator.discovery_url_override`. It changes where discovery metadata and JWKS are fetched without changing the issuer used to validate the JWT `iss` claim; this is useful when an OIDC provider has different internal and external URLs. Both validator settings are reloadable with `SIGHUP`.

Without `map=` in the HBA rule, the selected claim must exactly equal the requested PostgreSQL role. Use a named `pg_ident.conf` mapping when provider identities and database roles differ; the validator does not create roles.

### Connect with libpq

An OAuth-capable libpq client can start the provider's device authorization flow:

```bash
psql 'host=127.0.0.1 dbname=app user=alice oauth_issuer=https://id.example.com/realms/postgres oauth_client_id=postgres-client'
```

Use `oauth_client_secret` only when the registered client requires it. The client identifier, requested scope, issuer, and provider configuration must agree.

### Provider and Security Boundaries

- Keycloak must enable the OAuth 2 device flow for command-line clients.
- Microsoft Entra ID requires a tenant-specific v2 issuer and custom scopes; use the full scope name in `pg_hba.conf`.
- Google is not usable through libpq's built-in device flow, though custom clients may work.
- Dex does not emit OAuth scopes; an explicitly empty `scope=""` disables scope validation, which weakens the normal check.
- The client `oauth_issuer` must exactly match the HBA issuer and the discovery document. Treat the issuer and any `pg_oidc_validator.discovery_url_override` endpoint as trusted security boundaries, and require verified TLS for database and provider connections.
- Token validation does not replace PostgreSQL grants, role membership, or row-level security.
- Pigsty RPM packages are limited to EL10; DEB packages cover the supported Debian and Ubuntu targets. PostgreSQL 18 is required.
