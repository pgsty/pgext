---
title: "pg_rational"
linkTitle: "pg_rational"
description: "bigint fractions"
weight: 3720
categories: ["Type"]
width: full
---

bigint fractions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3720** | {{< badge content="pg_rational" link="https://github.com/begriffs/pg_rational" >}} | {{< ext "pg_rational" "pg_rational" >}} | `0.0.2` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_rational" >}} | `0.0.2` | {{< badge content="18" color="red" alt="pg_rational_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_rational_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pg_rational" >}} | `0.0.2` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-rational` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_rational_18" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_18-0.0.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_rational_17" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_17-0.0.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_rational_16" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_16-0.0.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_rational_15" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_15-0.0.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_rational_14" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_14-0.0.2-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_rational_18" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_18-0.0.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_rational_17" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_17-0.0.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_rational_16" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_16-0.0.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_rational_15" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_15-0.0.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_rational_14" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_14-0.0.2-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_rational_18" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_18-0.0.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_rational_17" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_17-0.0.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_rational_16" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_16-0.0.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_rational_15" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_15-0.0.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_rational_14" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_14-0.0.2-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_rational_18" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_18-0.0.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_rational_17" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_17-0.0.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_rational_16" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_16-0.0.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_rational_15" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_15-0.0.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_rational_14" "0.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_14-0.0.2-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-rational" "0.0.2" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_18` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_18` | 0.0.2 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_rational_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_18` | 0.0.2 | `el9.aarch64` | pigsty | 18.2 KiB | [pg_rational_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_rational_18` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-18-rational` | 0.0.2 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-18-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-18-rational` | 0.0.2 | `d12.x86_64` | pgdg | 24.2 KiB | [postgresql-18-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-18-rational` | 0.0.2 | `u22.x86_64` | pgdg | 24.2 KiB | [postgresql-18-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-18-rational` | 0.0.2 | `u22.aarch64` | pgdg | 23.8 KiB | [postgresql-18-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-18-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.2 KiB | [postgresql-18-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.3 KiB | [postgresql-18-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-18-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_17` | 0.0.2 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_rational_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_17` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_17` | 0.0.2 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_rational_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_rational_17` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-rational` | 0.0.2 | `d12.x86_64` | pgdg | 24.1 KiB | [postgresql-17-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-17-rational` | 0.0.2 | `d12.aarch64` | pgdg | 23.9 KiB | [postgresql-17-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-17-rational` | 0.0.2 | `u22.aarch64` | pgdg | 24.6 KiB | [postgresql-17-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-17-rational` | 0.0.2 | `u22.x86_64` | pgdg | 25.1 KiB | [postgresql-17-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-17-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.2 KiB | [postgresql-17-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |
| `postgresql-17-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.1 KiB | [postgresql-17-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-17-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_16` | 0.0.2 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_rational_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_16` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_16` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_rational_16` | 0.0.2 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_rational_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-rational` | 0.0.2 | `d12.aarch64` | pgdg | 23.9 KiB | [postgresql-16-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-16-rational` | 0.0.2 | `d12.x86_64` | pgdg | 24.1 KiB | [postgresql-16-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-16-rational` | 0.0.2 | `u22.aarch64` | pgdg | 24.6 KiB | [postgresql-16-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rational` | 0.0.2 | `u22.x86_64` | pgdg | 25.1 KiB | [postgresql-16-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.1 KiB | [postgresql-16-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.2 KiB | [postgresql-16-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-16-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_15` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_15` | 0.0.2 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_rational_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_15` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_rational_15` | 0.0.2 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_rational_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-rational` | 0.0.2 | `d12.x86_64` | pgdg | 24.2 KiB | [postgresql-15-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-15-rational` | 0.0.2 | `d12.aarch64` | pgdg | 23.9 KiB | [postgresql-15-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-15-rational` | 0.0.2 | `u22.aarch64` | pgdg | 24.6 KiB | [postgresql-15-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rational` | 0.0.2 | `u22.x86_64` | pgdg | 25.1 KiB | [postgresql-15-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.2 KiB | [postgresql-15-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.3 KiB | [postgresql-15-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-15-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_14` | 0.0.2 | `el8.x86_64` | pigsty | 18.6 KiB | [pg_rational_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_14` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_14` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_rational_14` | 0.0.2 | `el9.aarch64` | pigsty | 18.3 KiB | [pg_rational_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-rational` | 0.0.2 | `d12.x86_64` | pgdg | 24.1 KiB | [postgresql-14-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-14-rational` | 0.0.2 | `d12.aarch64` | pgdg | 23.9 KiB | [postgresql-14-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-14-rational` | 0.0.2 | `u22.x86_64` | pgdg | 25.0 KiB | [postgresql-14-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-14-rational` | 0.0.2 | `u22.aarch64` | pgdg | 24.6 KiB | [postgresql-14-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-14-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.2 KiB | [postgresql-14-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |
| `postgresql-14-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.1 KiB | [postgresql-14-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-14-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_rational_13` | 0.0.2 | `el8.aarch64` | pigsty | 18.0 KiB | [pg_rational_13-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_rational_13-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_rational_13` | 0.0.2 | `el8.x86_64` | pigsty | 18.4 KiB | [pg_rational_13-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_rational_13-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_rational_13` | 0.0.2 | `el9.aarch64` | pigsty | 18.4 KiB | [pg_rational_13-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_rational_13-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_rational_13` | 0.0.2 | `el9.x86_64` | pigsty | 18.7 KiB | [pg_rational_13-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_rational_13-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-rational` | 0.0.2 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-13-rational_0.0.2-8.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg12+1_arm64.deb) |
| `postgresql-13-rational` | 0.0.2 | `d12.x86_64` | pgdg | 23.9 KiB | [postgresql-13-rational_0.0.2-8.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg12+1_amd64.deb) |
| `postgresql-13-rational` | 0.0.2 | `u22.aarch64` | pgdg | 24.5 KiB | [postgresql-13-rational_0.0.2-8.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg22.04+1_arm64.deb) |
| `postgresql-13-rational` | 0.0.2 | `u22.x86_64` | pgdg | 24.7 KiB | [postgresql-13-rational_0.0.2-8.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg22.04+1_amd64.deb) |
| `postgresql-13-rational` | 0.0.2 | `u24.x86_64` | pgdg | 24.0 KiB | [postgresql-13-rational_0.0.2-8.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg24.04+1_amd64.deb) |
| `postgresql-13-rational` | 0.0.2 | `u24.aarch64` | pgdg | 24.1 KiB | [postgresql-13-rational_0.0.2-8.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rational/postgresql-13-rational_0.0.2-8.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/begriffs/pg_rational" title="Repository" icon="github" subtitle="github.com/begriffs/pg_rational" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_rational-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_rational; # get pg_rational source code
pig build dep pg_rational; # install build dependencies
pig build pkg pg_rational; # build extension rpm or deb
pig build ext pg_rational; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_rational; # install by extension name, for the current active PG version
pig ext install pg_rational; # install via package alias, for the active PG version
pig ext install pg_rational -v 18;   # install for PG 18
pig ext install pg_rational -v 17;   # install for PG 17
pig ext install pg_rational -v 16;   # install for PG 16
pig ext install pg_rational -v 15;   # install for PG 15
pig ext install pg_rational -v 14;   # install for PG 14
pig ext install pg_rational -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_rational;
```

