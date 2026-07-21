## Usage

Sources:

- [Official README](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/README.md)
- [Validator configuration source](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/src/config.rs)
- [PostgreSQL OAuth callback implementation](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/b65bbbe288f84fab91d58b8304e8a526d1326af5/src/ffi.rs)
- [PostgreSQL 18 OAuth authentication documentation](https://www.postgresql.org/docs/18/auth-oauth.html)

`oidc_validator` is a PostgreSQL 18 OAuth validator module written in Rust. It validates JWT access tokens against an OpenID Connect issuer and returns the token subject as the authenticated identity. It is a headless authentication library, not a SQL extension, so it creates no SQL objects and does not use `CREATE EXTENSION`.

### Core Workflow

Install `oidc_validator.so` in PostgreSQL's library directory, then configure the PostgreSQL 18 validator module:

```conf
oauth_validator_libraries = 'oidc_validator'
```

Add an OAuth rule to `pg_hba.conf`:

```conf
host all all 0.0.0.0/0 oauth issuer="https://issuer.example" scope="openid profile"
```

Provide the validator configuration to the PostgreSQL server process:

```shell
POSTGRES_OIDC_ISSUER=https://issuer.example
POSTGRES_OIDC_CLIENT_ID=postgres
POSTGRES_OIDC_AUDIENCE=postgres
```

Restart PostgreSQL after changing `oauth_validator_libraries` or the server-process environment. OAuth clients can then authenticate through a matching `pg_hba.conf` rule.

### Configuration Index

- `POSTGRES_OIDC_ISSUER`: issuer URL without the well-known discovery suffix.
- `POSTGRES_OIDC_CLIENT_ID`: OIDC application client ID.
- `POSTGRES_OIDC_AUDIENCE`: required token audience, commonly the client ID.
- `oauth_validator_libraries`: PostgreSQL 18 setting that loads the trusted validator module.

### Requirements and Caveats

- Upstream version `0.1.0` targets PostgreSQL 18 and requires PostgreSQL to be built with OpenSSL and libcurl.
- Only JWT-shaped bearer tokens are accepted. Opaque access tokens are rejected.
- Validation performs OIDC discovery and JWKS retrieval, so the PostgreSQL server must be able to reach the issuer over TLS.
- The callback currently ignores the requested PostgreSQL role and authorizes based on successful token validation. Do not set `delegate_ident_mapping=1` with this implementation; keep PostgreSQL's standard exact-name or `pg_ident.conf` mapping so the returned token subject is checked against the requested role.
