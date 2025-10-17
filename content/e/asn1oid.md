---
title: "asn1oid"
linkTitle: "asn1oid"
description: "asn1oid extension"
weight: 3560
categories: ["Type"]
width: full
---

asn1oid extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3560** | {{< badge content="asn1oid" link="https://github.com/df7cb/pgsql-asn1oid" >}} | {{< ext "asn1oid" "asn1oid" >}} | `1.6` | {{< category "TYPE" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/asn1oid" >}} | `1.6` | {{< badge content="18" color="red" alt="asn1oid_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `asn1oid_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/asn1oid" >}} | `1.6` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-asn1oid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "asn1oid_18" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_18-1.6-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "asn1oid_17" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_17-1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "asn1oid_16" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_16-1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "asn1oid_15" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_15-1.5-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "asn1oid_14" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_14-1.5-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "asn1oid_18" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_18-1.6-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "asn1oid_17" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_17-1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "asn1oid_16" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_16-1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "asn1oid_15" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_15-1.5-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "asn1oid_14" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_14-1.5-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "asn1oid_18" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_18-1.6-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "asn1oid_17" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_17-1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "asn1oid_16" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_16-1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "asn1oid_15" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_15-1.5-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "asn1oid_14" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_14-1.5-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "asn1oid_18" "1.6" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_18-1.6-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "asn1oid_17" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_17-1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "asn1oid_16" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_16-1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "asn1oid_15" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_15-1.5-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "asn1oid_14" "1.5" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_14-1.5-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_amd64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_amd64.deb" >}} |
|    `d12.aarch64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_arm64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_arm64.deb" >}} |
|    `u22.x86_64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_amd64.deb" >}} |
|    `u22.aarch64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_arm64.deb" >}} |
|    `u24.x86_64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_amd64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_amd64.deb" >}} |
|    `u24.aarch64`    | {{< pkg "postgresql-18-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-17-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-16-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-15-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_arm64.deb" >}} | {{< pkg "postgresql-14-asn1oid" "1.6" "pgdg" "https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_18` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_18-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_18-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_18` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_18-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_18-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_18` | 1.6 | `el9.aarch64` | pigsty | 13.4 KiB | [asn1oid_18-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_18-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_18` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_18-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_18-1.6-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-18-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.9 KiB | [postgresql-18-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-18-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.8 KiB | [postgresql-18-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-18-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.0 KiB | [postgresql-18-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.1 KiB | [postgresql-18-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 13.1 KiB | [postgresql-18-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-18-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_17` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_17-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_17-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_17` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_17-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_17-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_17` | 1.5 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_17-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_17-1.5-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_17` | 1.5 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_17-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_17-1.5-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_17` | 1.6 | `el9.aarch64` | pigsty | 13.5 KiB | [asn1oid_17-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_17-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_17` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_17-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_17-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_17` | 1.5 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_17-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_17-1.5-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_17` | 1.5 | `el9.aarch64` | pigsty | 13.6 KiB | [asn1oid_17-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_17-1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-17-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-17-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.8 KiB | [postgresql-17-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-17-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.4 KiB | [postgresql-17-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.3 KiB | [postgresql-17-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-17-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-17-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_16` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_16-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_16-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_16` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_16-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_16-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_16` | 1.5 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_16-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_16-1.5-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_16` | 1.5 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_16-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_16-1.5-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_16` | 1.6 | `el9.aarch64` | pigsty | 13.4 KiB | [asn1oid_16-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_16-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_16` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_16-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_16-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_16` | 1.5 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_16-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_16-1.5-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_16` | 1.5 | `el9.aarch64` | pigsty | 13.6 KiB | [asn1oid_16-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_16-1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.8 KiB | [postgresql-16-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-16-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-16-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-16-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.3 KiB | [postgresql-16-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.4 KiB | [postgresql-16-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 12.9 KiB | [postgresql-16-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-16-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_15` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_15-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_15-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_15` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_15-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_15-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_15` | 1.5 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_15-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_15-1.5-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_15` | 1.5 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_15-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_15-1.5-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_15` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_15-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_15-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_15` | 1.6 | `el9.aarch64` | pigsty | 13.5 KiB | [asn1oid_15-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_15-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_15` | 1.5 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_15-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_15-1.5-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_15` | 1.5 | `el9.aarch64` | pigsty | 13.6 KiB | [asn1oid_15-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_15-1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-15-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-15-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.8 KiB | [postgresql-15-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-15-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.4 KiB | [postgresql-15-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.3 KiB | [postgresql-15-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.9 KiB | [postgresql-15-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-15-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_14` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_14-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_14-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_14` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_14-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_14-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_14` | 1.5 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_14-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_14-1.5-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_14` | 1.5 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_14-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_14-1.5-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_14` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_14-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_14-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_14` | 1.6 | `el9.aarch64` | pigsty | 13.5 KiB | [asn1oid_14-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_14-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_14` | 1.5 | `el9.aarch64` | pigsty | 13.6 KiB | [asn1oid_14-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_14-1.5-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_14` | 1.5 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_14-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_14-1.5-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-14-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-14-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.3 KiB | [postgresql-14-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.4 KiB | [postgresql-14-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 13.0 KiB | [postgresql-14-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `asn1oid_13` | 1.6 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_13-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_13-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_13` | 1.6 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_13-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_13-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_13` | 1.5 | `el8.x86_64` | pigsty | 13.5 KiB | [asn1oid_13-1.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_13-1.5-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_13` | 1.5 | `el8.aarch64` | pigsty | 13.5 KiB | [asn1oid_13-1.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_13-1.5-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_13` | 1.6 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_13-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_13-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_13` | 1.6 | `el9.aarch64` | pigsty | 13.4 KiB | [asn1oid_13-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_13-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_13` | 1.5 | `el9.x86_64` | pigsty | 13.7 KiB | [asn1oid_13-1.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_13-1.5-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_13` | 1.5 | `el9.aarch64` | pigsty | 13.5 KiB | [asn1oid_13-1.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_13-1.5-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-asn1oid` | 1.6 | `d12.aarch64` | pgdg | 12.8 KiB | [postgresql-13-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-13-asn1oid` | 1.6 | `d12.x86_64` | pgdg | 12.7 KiB | [postgresql-13-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-13-asn1oid` | 1.6 | `u22.aarch64` | pgdg | 13.2 KiB | [postgresql-13-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-asn1oid` | 1.6 | `u22.x86_64` | pgdg | 13.3 KiB | [postgresql-13-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-asn1oid` | 1.6 | `u24.x86_64` | pgdg | 12.8 KiB | [postgresql-13-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-asn1oid` | 1.6 | `u24.aarch64` | pgdg | 12.9 KiB | [postgresql-13-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/pgsql-asn1oid" title="Repository" icon="github" subtitle="github.com/df7cb/pgsql-asn1oid" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsql-asn1oid-1.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build get asn1oid; # get asn1oid source code
pig build dep asn1oid; # install build dependencies
pig build pkg asn1oid; # build extension rpm or deb
pig build ext asn1oid; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install asn1oid; # install by extension name, for the current active PG version
pig ext install asn1oid; # install via package alias, for the active PG version
pig ext install asn1oid -v 18;   # install for PG 18
pig ext install asn1oid -v 17;   # install for PG 17
pig ext install asn1oid -v 16;   # install for PG 16
pig ext install asn1oid -v 15;   # install for PG 15
pig ext install asn1oid -v 14;   # install for PG 14
pig ext install asn1oid -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION asn1oid;
```

