## Usage

Sources:

- [pg_oidc_validator 0.2 README](https://github.com/percona/pg_oidc_validator/blob/0.2/README.md)
- [Keycloak example for 0.2](https://github.com/percona/pg_oidc_validator/tree/0.2/examples/keycloak)

pg_oidc_validator is an OAuth validator module for PostgreSQL 18 that validates libpq OAuth bearer tokens against an OpenID Connect issuer. Use it when PostgreSQL clients authenticate through an OIDC provider; it is loaded by the server and does not define a SQL extension, so do not run CREATE EXTENSION.

The project describes the module as experimental and not ready for production. Test the exact identity provider, client, and PostgreSQL build before relying on it.

### Configure the Server

Load the validator and restart PostgreSQL:

    oauth_validator_libraries = 'pg_oidc_validator'

Add an oauth rule to pg_hba.conf. The issuer and scope must match the provider:

    host  all  all  127.0.0.1/32  oauth  issuer=https://id.example.com/realms/postgres scope="openid postgres"

Reload pg_hba.conf after editing it. The validator checks the token issuer, audience, scope, signature, and expiry according to the provider metadata discovered from the issuer.

By default the PostgreSQL role is matched against the JWT sub claim. To authenticate by another claim, such as email, set:

    pg_oidc_validator.authn_field = 'email'

This setting changes the identity claim used for role matching; it does not create or provision database roles.

### Connect with libpq

A libpq client that supports OAuth can initiate the device-authorization flow:

    psql "host=127.0.0.1 dbname=app user=alice +      oauth_issuer=https://id.example.com/realms/postgres +      oauth_client_id=postgres-client"

Use oauth_client_secret only when the registered client requires one. The client identifier, redirect/device-flow settings, audience, and requested scopes must agree with the identity-provider configuration.

### Configuration Index

- oauth_validator_libraries: server-level list of OAuth validator modules; adding pg_oidc_validator requires a restart.
- pg_oidc_validator.authn_field: JWT claim compared with the requested PostgreSQL role; defaults to sub.
- pg_hba.conf oauth method: selects OAuth authentication and supplies the accepted issuer and scope.
- oauth_issuer, oauth_client_id, oauth_client_secret: libpq connection parameters used to obtain a token.

### Provider and Security Boundaries

- The upstream 0.2 documentation targets PostgreSQL 18 and requires an OAuth-capable libpq client.
- The validator supports common OIDC providers, but the README explicitly calls out Google as unsupported and describes provider-specific setup for Microsoft Entra ID.
- Token validation is only one part of authorization. PostgreSQL role membership and object privileges still control database access.
- Protect client secrets and provider credentials outside connection strings where possible, and validate TLS trust for the issuer.

