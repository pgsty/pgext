---
title: "uri"
linkTitle: "uri"
description: "URI Data type for PostgreSQL"
weight: 3840
categories: ["Type"]
width: full
---

URI Data type for PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3840** | {{< badge content="uri" link="https://github.com/petere/pguri" >}} | {{< ext "uri" "pg_uri" >}} | `1.20151224` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] +int flag


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/uri" >}} | `1.20151224` | {{< badge content="18" color="red" alt="pg_uri_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_uri_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/uri" >}} | `1.20151224` | {{< badge content="18" color="red" alt="postgresql-18-pg-uri" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-uri` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_uri_18" >}}     | {{< pkg "pg_uri_17" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_17-1.20151224-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_uri_16" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_16-1.20151224-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_uri_15" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_15-1.20151224-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_uri_14" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_14-1.20151224-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_uri_18" >}}     | {{< pkg "pg_uri_17" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_17-1.20151224-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_uri_16" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_16-1.20151224-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_uri_15" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_15-1.20151224-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_uri_14" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_14-1.20151224-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_uri_18" >}}     | {{< pkg "pg_uri_17" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_17-1.20151224-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_uri_16" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_16-1.20151224-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_uri_15" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_15-1.20151224-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_uri_14" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_14-1.20151224-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_uri_18" >}}     | {{< pkg "pg_uri_17" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_17-1.20151224-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_uri_16" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_16-1.20151224-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_uri_15" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_15-1.20151224-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_uri_14" "1.20151224" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_14-1.20151224-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-uri" >}}     | {{< pkg "postgresql-17-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-uri" "1.20151224" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_uri_17` | 1.20151224 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_uri_17-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_17-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_17` | 1.20151224 | `el8.aarch64` | pigsty | 18.3 KiB | [pg_uri_17-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_17-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_17` | 1.20151224 | `el9.aarch64` | pigsty | 18.4 KiB | [pg_uri_17-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_17-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_17` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_17-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_17-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 22.2 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.9 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.2 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.3 KiB | [postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-17-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_uri_16` | 1.20151224 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_uri_16-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_16-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_16` | 1.20151224 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_uri_16-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_16-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_16` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_16-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_16-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_16` | 1.20151224 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_uri_16-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_16-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 22.2 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.9 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.2 KiB | [postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-16-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_uri_15` | 1.20151224 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_uri_15-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_15-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_15` | 1.20151224 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_uri_15-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_15-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_15` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_15-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_15-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_15` | 1.20151224 | `el9.aarch64` | pigsty | 18.4 KiB | [pg_uri_15-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_15-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.9 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 22.2 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 23.0 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.2 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.3 KiB | [postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-15-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_uri_14` | 1.20151224 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_uri_14-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_14-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_14` | 1.20151224 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_uri_14-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_14-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_14` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_14-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_14-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `pg_uri_14` | 1.20151224 | `el9.aarch64` | pigsty | 18.4 KiB | [pg_uri_14-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_14-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 22.2 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.9 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 23.1 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 22.9 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.8 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.2 KiB | [postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-14-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_uri_13` | 1.20151224 | `el8.aarch64` | pigsty | 18.2 KiB | [pg_uri_13-1.20151224-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_uri_13-1.20151224-1PIGSTY.el8.aarch64.rpm) |
| `pg_uri_13` | 1.20151224 | `el8.x86_64` | pigsty | 18.5 KiB | [pg_uri_13-1.20151224-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_uri_13-1.20151224-1PIGSTY.el8.x86_64.rpm) |
| `pg_uri_13` | 1.20151224 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_uri_13-1.20151224-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_uri_13-1.20151224-1PIGSTY.el9.aarch64.rpm) |
| `pg_uri_13` | 1.20151224 | `el9.x86_64` | pigsty | 18.9 KiB | [pg_uri_13-1.20151224-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_uri_13-1.20151224-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-uri` | 1.20151224 | `d12.aarch64` | pigsty | 21.8 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `d12.x86_64` | pigsty | 22.1 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u22.aarch64` | pigsty | 22.7 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u22.x86_64` | pigsty | 22.9 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u24.aarch64` | pigsty | 22.4 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-uri` | 1.20151224 | `u24.x86_64` | pigsty | 22.7 KiB | [postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uri/postgresql-13-pg-uri_1.20151224-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pguri" title="Repository" icon="github" subtitle="github.com/petere/pguri" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pguri-1.20151224.tar.gz" >}}
{{< /cards >}}


```bash
pig build get uri; # get uri source code
pig build dep uri; # install build dependencies
pig build pkg uri; # build extension rpm or deb
pig build ext uri; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install uri; # install by extension name, for the current active PG version
pig ext install pg_uri; # install via package alias, for the active PG version
pig ext install uri -v 18;   # install for PG 18
pig ext install uri -v 17;   # install for PG 17
pig ext install uri -v 16;   # install for PG 16
pig ext install uri -v 15;   # install for PG 15
pig ext install uri -v 14;   # install for PG 14
pig ext install uri -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION uri;
```

