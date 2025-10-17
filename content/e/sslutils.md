---
title: "sslutils"
linkTitle: "sslutils"
description: "A Postgres extension for managing SSL certificates through SQL"
weight: 7200
categories: ["Sec"]
width: full
---

A Postgres extension for managing SSL certificates through SQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7200** | {{< badge content="sslutils" link="https://github.com/EnterpriseDB/sslutils" >}} | {{< ext "sslutils" "sslutils" >}} | `1.4` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "sslinfo" >}} {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "pg_tde" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} |

> [!Note] no pg15,14,13,12 on el9, no pg12 on el8


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/sslutils" >}} | `1.4` | {{< badge content="18" color="red" alt="sslutils_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `sslutils_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/sslutils" >}} | `1.4` | {{< badge content="18" color="red" alt="postgresql-18-sslutils" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-sslutils` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "sslutils_18" >}}     | {{< pkg "sslutils_17" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/sslutils_17-1.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "sslutils_16" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/sslutils_16-1.4-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "sslutils_15" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sslutils_15-1.3-4.rhel8.x86_64.rpm" >}} | {{< pkg "sslutils_14" "1.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sslutils_14-1.3-4.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "sslutils_18" >}}     | {{< pkg "sslutils_17" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/sslutils_17-1.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "sslutils_16" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/sslutils_16-1.4-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "sslutils_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_15-1.4-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "sslutils_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_14-1.4-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "sslutils_18" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sslutils_18-1.4-2PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "sslutils_17" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sslutils_17-1.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "sslutils_16" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sslutils_16-1.4-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "sslutils_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_15-1.4-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "sslutils_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_14-1.4-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "sslutils_18" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sslutils_18-1.4-2PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "sslutils_17" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sslutils_17-1.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "sslutils_16" "1.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sslutils_16-1.4-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "sslutils_15" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_15-1.4-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "sslutils_14" "1.4" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_14-1.4-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-sslutils" >}}     | {{< pkg "postgresql-17-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-sslutils" "1.4" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_18` | 1.4 | `el9.x86_64` | pigsty | 24.2 KiB | [sslutils_18-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_18-1.4-1PIGSTY.el9.x86_64.rpm) |
| `sslutils_18` | 1.4 | `el9.aarch64` | pgdg | 23.2 KiB | [sslutils_18-1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sslutils_18-1.4-2PGDG.rhel9.aarch64.rpm) |
| `sslutils_18` | 1.4 | `el9.aarch64` | pigsty | 23.0 KiB | [sslutils_18-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_18-1.4-1PIGSTY.el9.aarch64.rpm) |
| `sslutils_18` | 1.4 | `el9.x86_64` | pgdg | 24.5 KiB | [sslutils_18-1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sslutils_18-1.4-2PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_17` | 1.4 | `el8.x86_64` | pigsty | 23.7 KiB | [sslutils_17-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_17-1.4-1PIGSTY.el8.x86_64.rpm) |
| `sslutils_17` | 1.4 | `el8.aarch64` | pgdg | 23.4 KiB | [sslutils_17-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/sslutils_17-1.4-1PGDG.rhel8.aarch64.rpm) |
| `sslutils_17` | 1.4 | `el8.aarch64` | pigsty | 22.6 KiB | [sslutils_17-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_17-1.4-1PIGSTY.el8.aarch64.rpm) |
| `sslutils_17` | 1.4 | `el8.x86_64` | pgdg | 24.5 KiB | [sslutils_17-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/sslutils_17-1.4-1PGDG.rhel8.x86_64.rpm) |
| `sslutils_17` | 1.4 | `el9.aarch64` | pigsty | 23.1 KiB | [sslutils_17-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_17-1.4-1PIGSTY.el9.aarch64.rpm) |
| `sslutils_17` | 1.4 | `el9.x86_64` | pigsty | 24.3 KiB | [sslutils_17-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_17-1.4-1PIGSTY.el9.x86_64.rpm) |
| `sslutils_17` | 1.4 | `el9.x86_64` | pgdg | 24.4 KiB | [sslutils_17-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sslutils_17-1.4-1PGDG.rhel9.x86_64.rpm) |
| `sslutils_17` | 1.4 | `el9.aarch64` | pgdg | 23.3 KiB | [sslutils_17-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sslutils_17-1.4-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-sslutils` | 1.4 | `d12.aarch64` | pigsty | 35.5 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-sslutils` | 1.4 | `d12.x86_64` | pigsty | 37.1 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-sslutils` | 1.4 | `u22.aarch64` | pigsty | 41.6 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-sslutils` | 1.4 | `u22.x86_64` | pigsty | 42.8 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-sslutils` | 1.4 | `u24.x86_64` | pigsty | 39.4 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-sslutils` | 1.4 | `u24.aarch64` | pigsty | 38.1 KiB | [postgresql-17-sslutils_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-17-sslutils_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_16` | 1.4 | `el8.x86_64` | pigsty | 23.7 KiB | [sslutils_16-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_16-1.4-1PIGSTY.el8.x86_64.rpm) |
| `sslutils_16` | 1.4 | `el8.aarch64` | pgdg | 23.4 KiB | [sslutils_16-1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/sslutils_16-1.4-1PGDG.rhel8.aarch64.rpm) |
| `sslutils_16` | 1.4 | `el8.aarch64` | pigsty | 22.6 KiB | [sslutils_16-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_16-1.4-1PIGSTY.el8.aarch64.rpm) |
| `sslutils_16` | 1.4 | `el8.x86_64` | pgdg | 24.5 KiB | [sslutils_16-1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/sslutils_16-1.4-1PGDG.rhel8.x86_64.rpm) |
| `sslutils_16` | 1.4 | `el9.x86_64` | pgdg | 24.4 KiB | [sslutils_16-1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sslutils_16-1.4-1PGDG.rhel9.x86_64.rpm) |
| `sslutils_16` | 1.4 | `el9.aarch64` | pgdg | 23.3 KiB | [sslutils_16-1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sslutils_16-1.4-1PGDG.rhel9.aarch64.rpm) |
| `sslutils_16` | 1.4 | `el9.aarch64` | pigsty | 23.1 KiB | [sslutils_16-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_16-1.4-1PIGSTY.el9.aarch64.rpm) |
| `sslutils_16` | 1.4 | `el9.x86_64` | pigsty | 24.2 KiB | [sslutils_16-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_16-1.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-sslutils` | 1.4 | `d12.x86_64` | pigsty | 37.0 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-sslutils` | 1.4 | `d12.aarch64` | pigsty | 35.4 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-sslutils` | 1.4 | `u22.aarch64` | pigsty | 41.6 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-sslutils` | 1.4 | `u22.x86_64` | pigsty | 42.8 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-sslutils` | 1.4 | `u24.x86_64` | pigsty | 39.4 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-sslutils` | 1.4 | `u24.aarch64` | pigsty | 38.0 KiB | [postgresql-16-sslutils_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-16-sslutils_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_15` | 1.4 | `el8.aarch64` | pigsty | 22.6 KiB | [sslutils_15-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_15-1.4-1PIGSTY.el8.aarch64.rpm) |
| `sslutils_15` | 1.4 | `el8.x86_64` | pigsty | 23.7 KiB | [sslutils_15-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_15-1.4-1PIGSTY.el8.x86_64.rpm) |
| `sslutils_15` | 1.3 | `el8.x86_64` | pgdg | 49.4 KiB | [sslutils_15-1.3-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sslutils_15-1.3-4.rhel8.x86_64.rpm) |
| `sslutils_15` | 1.4 | `el9.x86_64` | pigsty | 24.3 KiB | [sslutils_15-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_15-1.4-1PIGSTY.el9.x86_64.rpm) |
| `sslutils_15` | 1.4 | `el9.aarch64` | pigsty | 23.2 KiB | [sslutils_15-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_15-1.4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-sslutils` | 1.4 | `d12.aarch64` | pigsty | 35.4 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-sslutils` | 1.4 | `d12.x86_64` | pigsty | 36.9 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-sslutils` | 1.4 | `u22.x86_64` | pigsty | 42.8 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-sslutils` | 1.4 | `u22.aarch64` | pigsty | 41.6 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-sslutils` | 1.4 | `u24.x86_64` | pigsty | 39.4 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-sslutils` | 1.4 | `u24.aarch64` | pigsty | 38.1 KiB | [postgresql-15-sslutils_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-15-sslutils_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_14` | 1.4 | `el8.aarch64` | pigsty | 22.6 KiB | [sslutils_14-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_14-1.4-1PIGSTY.el8.aarch64.rpm) |
| `sslutils_14` | 1.4 | `el8.x86_64` | pigsty | 23.7 KiB | [sslutils_14-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_14-1.4-1PIGSTY.el8.x86_64.rpm) |
| `sslutils_14` | 1.3 | `el8.x86_64` | pgdg | 48.9 KiB | [sslutils_14-1.3-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sslutils_14-1.3-4.rhel8.x86_64.rpm) |
| `sslutils_14` | 1.4 | `el9.aarch64` | pigsty | 23.1 KiB | [sslutils_14-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_14-1.4-1PIGSTY.el9.aarch64.rpm) |
| `sslutils_14` | 1.4 | `el9.x86_64` | pigsty | 24.3 KiB | [sslutils_14-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_14-1.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-sslutils` | 1.4 | `d12.aarch64` | pigsty | 35.4 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-sslutils` | 1.4 | `d12.x86_64` | pigsty | 37.0 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-sslutils` | 1.4 | `u22.x86_64` | pigsty | 42.8 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-sslutils` | 1.4 | `u22.aarch64` | pigsty | 41.6 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-sslutils` | 1.4 | `u24.x86_64` | pigsty | 39.4 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-sslutils` | 1.4 | `u24.aarch64` | pigsty | 38.1 KiB | [postgresql-14-sslutils_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-14-sslutils_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `sslutils_13` | 1.4 | `el8.aarch64` | pigsty | 22.6 KiB | [sslutils_13-1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sslutils_13-1.4-1PIGSTY.el8.aarch64.rpm) |
| `sslutils_13` | 1.4 | `el8.x86_64` | pigsty | 23.6 KiB | [sslutils_13-1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sslutils_13-1.4-1PIGSTY.el8.x86_64.rpm) |
| `sslutils_13` | 1.3 | `el8.x86_64` | pgdg | 48.3 KiB | [sslutils_13-1.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/sslutils_13-1.3-2.rhel8.x86_64.rpm) |
| `sslutils_13` | 1.4 | `el9.aarch64` | pigsty | 23.2 KiB | [sslutils_13-1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sslutils_13-1.4-1PIGSTY.el9.aarch64.rpm) |
| `sslutils_13` | 1.4 | `el9.x86_64` | pigsty | 24.3 KiB | [sslutils_13-1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sslutils_13-1.4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-sslutils` | 1.4 | `d12.aarch64` | pigsty | 35.1 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-sslutils` | 1.4 | `d12.x86_64` | pigsty | 37.0 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-sslutils` | 1.4 | `u22.x86_64` | pigsty | 42.5 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-sslutils` | 1.4 | `u22.aarch64` | pigsty | 41.5 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-sslutils` | 1.4 | `u24.x86_64` | pigsty | 39.4 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-sslutils` | 1.4 | `u24.aarch64` | pigsty | 37.9 KiB | [postgresql-13-sslutils_1.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sslutils/postgresql-13-sslutils_1.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/sslutils" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/sslutils" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="sslutils-1.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get sslutils; # get sslutils source code
pig build dep sslutils; # install build dependencies
pig build pkg sslutils; # build extension rpm or deb
pig build ext sslutils; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install sslutils; # install by extension name, for the current active PG version
pig ext install sslutils; # install via package alias, for the active PG version
pig ext install sslutils -v 18;   # install for PG 18
pig ext install sslutils -v 17;   # install for PG 17
pig ext install sslutils -v 16;   # install for PG 16
pig ext install sslutils -v 15;   # install for PG 15
pig ext install sslutils -v 14;   # install for PG 14
pig ext install sslutils -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION sslutils;
```

