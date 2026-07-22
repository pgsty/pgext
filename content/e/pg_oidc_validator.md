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
| **7170** | {{< badge content="pg_oidc_validator" link="https://github.com/percona/pg_oidc_validator" >}} | {{< ext "pg_oidc_validator" >}} | `0.2` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oidc_validator" >}} {{< ext "pg_session_jwt" >}} {{< ext "pgjwt" >}} {{< ext "login_hook" >}} {{< ext "auth_delay" >}} |

> [!Note] Configure oauth_validator_libraries='pg_oidc_validator'. RPM is available on EL10 only; EL8/EL9 RPMs were excluded after libstdc++ ABI smoke failures. DEB covers all supported Debian/Ubuntu targets.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_oidc_validator` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2` | {{< bg "18" "pg_oidc_validator_18" "green" >}} {{< bg "17" "pg_oidc_validator_17" "red" >}} {{< bg "16" "pg_oidc_validator_16" "red" >}} {{< bg "15" "pg_oidc_validator_15" "red" >}} {{< bg "14" "pg_oidc_validator_14" "red" >}} | `pg_oidc_validator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2` | {{< bg "18" "postgresql-18-pg-oidc-validator" "green" >}} {{< bg "17" "postgresql-17-pg-oidc-validator" "red" >}} {{< bg "16" "postgresql-16-pg-oidc-validator" "red" >}} {{< bg "15" "postgresql-15-pg-oidc-validator" "red" >}} {{< bg "14" "postgresql-14-pg-oidc-validator" "red" >}} | `postgresql-$v-pg-oidc-validator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_oidc_validator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_oidc_validator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_oidc_validator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_oidc_validator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2" "pg_oidc_validator_18 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2" "pg_oidc_validator_18 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_oidc_validator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_oidc_validator_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2" "postgresql-18-pg-oidc-validator : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-17-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-oidc-validator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-oidc-validator : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_oidc_validator_18` | `0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.8 KiB | [pg_oidc_validator_18-0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_oidc_validator_18-0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_oidc_validator_18` | `0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 173.1 KiB | [pg_oidc_validator_18-0.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_oidc_validator_18-0.2-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_oidc_validator_18` | `0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 127.8 KiB | [pg_oidc_validator_18-0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_oidc_validator_18-0.2-1PIGSTY.el10.aarch64.rpm) |
| `pg_oidc_validator_18` | `0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 154.9 KiB | [pg_oidc_validator_18-0.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_oidc_validator_18-0.2-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.9 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 94.1 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 115.5 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 100.4 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 105.3 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 96.9 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 107.1 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 98.9 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u26.x86_64](/os/u26.x86_64) | pigsty | 119.6 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-oidc-validator` | `0.2` | [u26.aarch64](/os/u26.aarch64) | pigsty | 104.8 KiB | [postgresql-18-pg-oidc-validator_0.2-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-oidc-validator/postgresql-18-pg-oidc-validator_0.2-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/percona/pg_oidc_validator" title="Repository" icon="github" subtitle="github.com/percona/pg_oidc_validator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_oidc_validator-0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_oidc_validator;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

