---
title: "pg_oidc_validator"
linkTitle: "pg_oidc_validator"
description: "OAuth and OIDC token validator for PostgreSQL 18"
weight: 7170
categories: ["SEC"]
width: full
---

[**pg_oidc_validator**](https://github.com/percona/pg_oidc_validator) : OAuth and OIDC token validator for PostgreSQL 18


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7170** | {{< badge content="pg_oidc_validator" link="https://github.com/percona/pg_oidc_validator" >}} | {{< ext "pg_oidc_validator" >}} | `1.1.0` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oidc_validator" >}} {{< ext "pg_session_jwt" >}} {{< ext "pgjwt" >}} {{< ext "login_hook" >}} {{< ext "sslinfo" >}} {{< ext "sslutils" >}} {{< ext "pgsodium" >}} {{< ext "pguecc" >}} |

> [!Note] Configure oauth_validator_libraries=pg_oidc_validator; 1.1.0 adds discovery_url_override; RPM is available on EL10 only while DEB covers all supported Debian and Ubuntu targets.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_oidc_validator` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "pg_oidc_validator_18" "green" >}} {{< bg "17" "pg_oidc_validator_17" "red" >}} {{< bg "16" "pg_oidc_validator_16" "red" >}} {{< bg "15" "pg_oidc_validator_15" "red" >}} {{< bg "14" "pg_oidc_validator_14" "red" >}} | `pg_oidc_validator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-pg-oidc-validator" "green" >}} {{< bg "17" "postgresql-17-pg-oidc-validator" "red" >}} {{< bg "16" "postgresql-16-pg-oidc-validator" "red" >}} {{< bg "15" "postgresql-15-pg-oidc-validator" "red" >}} {{< bg "14" "postgresql-14-pg-oidc-validator" "red" >}} | `postgresql-$v-pg-oidc-validator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_oidc_validator_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_oidc_validator_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pg_oidc_validator_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pg_oidc_validator_18 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_oidc_validator_18 : AVAIL 3" "green" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_oidc_validator_18 : AVAIL 3" "green" >}} | {{< bg "N/A" "pg_oidc_validator_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_oidc_validator_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-oidc-validator : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-oidc-validator : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_oidc_validator_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 142.6 KiB | [pg_oidc_validator_18-1.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_oidc_validator_18-1.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_oidc_validator_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 173.5 KiB | [pg_oidc_validator_18-1.0.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_oidc_validator_18-1.0.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_oidc_validator_18` | `0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 173.1 KiB | [pg_oidc_validator_18-0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_oidc_validator_18-0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_oidc_validator_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 129.5 KiB | [pg_oidc_validator_18-1.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_oidc_validator_18-1.1.0-1PGSTY.el10.aarch64.rpm) |
| `pg_oidc_validator_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 155.4 KiB | [pg_oidc_validator_18-1.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_oidc_validator_18-1.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_oidc_validator_18` | `0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 154.9 KiB | [pg_oidc_validator_18-0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_oidc_validator_18-0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.8 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 94.5 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 116.6 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 101.3 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 106.3 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 98.0 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 107.8 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.6 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 120.9 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 105.2 KiB | [postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_1.1.0-1PGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_oidc_validator" title="Repository" icon="github" subtitle="github.com/percona/pg_oidc_validator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_oidc_validator-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_oidc_validator;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_oidc_validator;		# install via package name, for the active PG version

pig install pg_oidc_validator -v 18;   # install for PG 18

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_oidc_validator';
```


This extension does not need `CREATE EXTENSION` DDL command



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
