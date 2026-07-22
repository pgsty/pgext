---
title: "oidc_validator"
linkTitle: "oidc_validator"
description: "PostgreSQL 18 OIDC bearer-token validator plugin written in Rust"
weight: 7180
categories: ["SEC"]
width: full
---

[**pg_oidc_validator_rust**](https://github.com/UnAfraid/pg_oidc_validator_rust) : PostgreSQL 18 OIDC bearer-token validator plugin written in Rust


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7180** | {{< badge content="oidc_validator" link="https://github.com/UnAfraid/pg_oidc_validator_rust" >}} | {{< ext "oidc_validator" "pg_oidc_validator_rust" >}} | `0.1.0` | {{< category "SEC" >}} | {{< license "LicenseRef-Upstream-No-License" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_oidc_validator" >}} {{< ext "pg_session_jwt" >}} {{< ext "pgjwt" >}} {{< ext "login_hook" >}} {{< ext "auth_delay" >}} |

> [!Note] Configure oauth_validator_libraries='oidc_validator'. Built from upstream commit b65bbbe288f84fab91d58b8304e8a526d1326af5; upstream publishes no license grant.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_oidc_validator_rust` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_oidc_validator_rust_18" "green" >}} {{< bg "17" "pg_oidc_validator_rust_17" "red" >}} {{< bg "16" "pg_oidc_validator_rust_16" "red" >}} {{< bg "15" "pg_oidc_validator_rust_15" "red" >}} {{< bg "14" "pg_oidc_validator_rust_14" "red" >}} | `pg_oidc_validator_rust_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pg-oidc-validator-rust" "green" >}} {{< bg "17" "postgresql-17-pg-oidc-validator-rust" "red" >}} {{< bg "16" "postgresql-16-pg-oidc-validator-rust" "red" >}} {{< bg "15" "postgresql-15-pg-oidc-validator-rust" "red" >}} {{< bg "14" "postgresql-14-pg-oidc-validator-rust" "red" >}} | `postgresql-$v-pg-oidc-validator-rust` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_oidc_validator_rust_18 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_oidc_validator_rust_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_rust_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pg-oidc-validator-rust : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator-rust : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator-rust : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_oidc_validator_rust_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.1 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_oidc_validator_rust_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.0 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_oidc_validator_rust_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.1 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_oidc_validator_rust_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.1 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_oidc_validator_rust_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.1 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_oidc_validator_rust_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.1 MiB | [pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_oidc_validator_rust_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.4 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.4 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.6 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-oidc-validator-rust` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator-rust/postgresql-18-pg-oidc-validator-rust_0.1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/UnAfraid/pg_oidc_validator_rust" title="Repository" icon="github" subtitle="github.com/UnAfraid/pg_oidc_validator_rust" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_oidc_validator_rust-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_oidc_validator_rust;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_oidc_validator_rust;		# install via package name, for the active PG version
pig install oidc_validator;		# install by extension name, for the current active PG version

pig install oidc_validator -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'oidc_validator';
```


This extension does not need `CREATE EXTENSION` DDL command



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
