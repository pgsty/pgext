## Usage

Sources:

- [Upstream README](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/README.md)
- [Cargo package metadata](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/Cargo.toml)
- [PostgreSQL OAuth validator callback implementation](https://github.com/UnAfraid/pg_oidc_validator_rust/blob/master/src/ffi.rs)

`oidc_validator` is a PostgreSQL 18 OAuth bearer-token validator written in Rust. It validates JWT access tokens through OIDC discovery and the issuer's JWKS endpoint, then returns the token subject as the authenticated identity. Upstream version `0.1.0` is tested on Ubuntu and Arch Linux x86-64.

This is a headless authentication module, not a SQL extension. Build and install `oidc_validator.so`, then configure PostgreSQL to load it through the PostgreSQL 18 OAuth validator setting:

```conf
oauth_validator_libraries = 'oidc_validator'
```

Add an OAuth rule to `pg_hba.conf` and provide the validator's environment variables to the PostgreSQL server process:

```conf
host all all 0.0.0.0/0 oauth issuer="https://issuer.example" scope="openid profile"
```

```shell
POSTGRES_OIDC_ISSUER=https://issuer.example
POSTGRES_OIDC_CLIENT_ID=postgres
POSTGRES_OIDC_AUDIENCE=postgres
```

Restart PostgreSQL after changing the validator library setting. The upstream build requires PostgreSQL 18 configured with OpenSSL and libcurl. Authentication performs network requests for OIDC discovery and JWKS data, so the server must be able to reach the issuer over TLS. The validator accepts JWT-shaped tokens only and does not provide SQL objects.
