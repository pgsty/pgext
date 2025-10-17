---
title: "pgsodium"
linkTitle: "pgsodium"
description: "Postgres extension for libsodium functions"
weight: 7020
categories: ["Sec"]
width: full
---

Postgres extension for libsodium functions

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7020** | {{< badge content="pgsodium" link="https://github.com/michelp/pgsodium" >}} | {{< ext "pgsodium" "pgsodium" >}} | `3.1.9` | {{< category "SEC" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "supabase_vault" >}} |
|   **See Also**    | {{< ext "pgsmcrypto" >}} {{< ext "pgcryptokey" >}} {{< ext "pgcrypto" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "sslutils" >}} {{< ext "faker" >}} |

> [!Note] +fix missing pg17


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgsodium" >}} | `3.1.9` | {{< badge content="18" color="red" alt="pgsodium_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgsodium_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgsodium" >}} | `3.1.9` | {{< badge content="18" color="red" alt="postgresql-18-pgsodium" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pgsodium` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgsodium_18" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsodium_18-3.1.9-4PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsodium_17" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsodium_17-3.1.9-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsodium_16" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsodium_16-3.1.9-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsodium_15" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.9-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgsodium_14" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.9-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgsodium_18" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsodium_18-3.1.9-4PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsodium_17" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsodium_17-3.1.9-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsodium_16" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsodium_16-3.1.9-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsodium_15" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.9-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgsodium_14" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.9-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgsodium_18" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsodium_18-3.1.9-4PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsodium_17" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsodium_17-3.1.9-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsodium_16" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsodium_16-3.1.9-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsodium_15" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.9-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgsodium_14" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.9-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgsodium_18" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsodium_18-3.1.9-4PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsodium_17" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsodium_17-3.1.9-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsodium_16" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsodium_16-3.1.9-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsodium_15" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.9-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgsodium_14" "3.1.9" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.9-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pgsodium" >}}     | {{< pkg "postgresql-17-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pgsodium" "3.1.9" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_18` | 3.1.9 | `el8.x86_64` | pigsty | 53.4 KiB | [pgsodium_18-3.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsodium_18-3.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pgsodium_18` | 3.1.9 | `el8.aarch64` | pigsty | 49.9 KiB | [pgsodium_18-3.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsodium_18-3.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pgsodium_18` | 3.1.9 | `el8.aarch64` | pgdg | 67.9 KiB | [pgsodium_18-3.1.9-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgsodium_18-3.1.9-4PGDG.rhel8.aarch64.rpm) |
| `pgsodium_18` | 3.1.9 | `el8.x86_64` | pgdg | 71.4 KiB | [pgsodium_18-3.1.9-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgsodium_18-3.1.9-4PGDG.rhel8.x86_64.rpm) |
| `pgsodium_18` | 3.1.9 | `el9.aarch64` | pigsty | 51.4 KiB | [pgsodium_18-3.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsodium_18-3.1.9-1PIGSTY.el9.aarch64.rpm) |
| `pgsodium_18` | 3.1.9 | `el9.x86_64` | pigsty | 54.3 KiB | [pgsodium_18-3.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsodium_18-3.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pgsodium_18` | 3.1.9 | `el9.x86_64` | pgdg | 70.8 KiB | [pgsodium_18-3.1.9-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgsodium_18-3.1.9-4PGDG.rhel9.x86_64.rpm) |
| `pgsodium_18` | 3.1.9 | `el9.aarch64` | pgdg | 67.5 KiB | [pgsodium_18-3.1.9-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgsodium_18-3.1.9-4PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_17` | 3.1.9 | `el8.x86_64` | pigsty | 53.4 KiB | [pgsodium_17-3.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgsodium_17-3.1.9-1PIGSTY.el8.x86_64.rpm) |
| `pgsodium_17` | 3.1.9 | `el8.aarch64` | pgdg | 67.8 KiB | [pgsodium_17-3.1.9-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgsodium_17-3.1.9-3PGDG.rhel8.aarch64.rpm) |
| `pgsodium_17` | 3.1.9 | `el8.aarch64` | pigsty | 49.9 KiB | [pgsodium_17-3.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgsodium_17-3.1.9-1PIGSTY.el8.aarch64.rpm) |
| `pgsodium_17` | 3.1.9 | `el8.x86_64` | pgdg | 71.3 KiB | [pgsodium_17-3.1.9-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgsodium_17-3.1.9-3PGDG.rhel8.x86_64.rpm) |
| `pgsodium_17` | 3.1.9 | `el9.aarch64` | pgdg | 68.5 KiB | [pgsodium_17-3.1.9-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgsodium_17-3.1.9-3PGDG.rhel9.aarch64.rpm) |
| `pgsodium_17` | 3.1.9 | `el9.x86_64` | pgdg | 71.5 KiB | [pgsodium_17-3.1.9-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgsodium_17-3.1.9-3PGDG.rhel9.x86_64.rpm) |
| `pgsodium_17` | 3.1.9 | `el9.x86_64` | pigsty | 54.3 KiB | [pgsodium_17-3.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgsodium_17-3.1.9-1PIGSTY.el9.x86_64.rpm) |
| `pgsodium_17` | 3.1.9 | `el9.aarch64` | pigsty | 51.4 KiB | [pgsodium_17-3.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgsodium_17-3.1.9-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pgsodium` | 3.1.9 | `d12.aarch64` | pigsty | 188.5 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgsodium` | 3.1.9 | `d12.x86_64` | pigsty | 192.1 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgsodium` | 3.1.9 | `u22.aarch64` | pigsty | 201.5 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgsodium` | 3.1.9 | `u22.x86_64` | pigsty | 204.9 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgsodium` | 3.1.9 | `u24.x86_64` | pigsty | 191.8 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgsodium` | 3.1.9 | `u24.aarch64` | pigsty | 188.5 KiB | [postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-17-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_16` | 3.1.9 | `el8.aarch64` | pgdg | 67.6 KiB | [pgsodium_16-3.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsodium_16-3.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_16` | 3.1.9 | `el8.x86_64` | pgdg | 71.1 KiB | [pgsodium_16-3.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsodium_16-3.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_16` | 3.1.8 | `el8.aarch64` | pgdg | 65.5 KiB | [pgsodium_16-3.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgsodium_16-3.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_16` | 3.1.8 | `el8.x86_64` | pgdg | 69.1 KiB | [pgsodium_16-3.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgsodium_16-3.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_16` | 3.1.9 | `el9.x86_64` | pgdg | 71.3 KiB | [pgsodium_16-3.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsodium_16-3.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_16` | 3.1.9 | `el9.aarch64` | pgdg | 68.3 KiB | [pgsodium_16-3.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsodium_16-3.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_16` | 3.1.8 | `el9.aarch64` | pgdg | 66.0 KiB | [pgsodium_16-3.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgsodium_16-3.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_16` | 3.1.8 | `el9.x86_64` | pgdg | 69.1 KiB | [pgsodium_16-3.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgsodium_16-3.1.8-1PGDG.rhel9.x86_64.rpm) |
| `postgresql-16-pgsodium` | 3.1.9 | `d12.aarch64` | pigsty | 184.6 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgsodium` | 3.1.9 | `d12.x86_64` | pigsty | 188.6 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgsodium` | 3.1.9 | `u22.x86_64` | pigsty | 200.2 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgsodium` | 3.1.9 | `u22.aarch64` | pigsty | 197.0 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgsodium` | 3.1.9 | `u24.x86_64` | pigsty | 186.7 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgsodium` | 3.1.9 | `u24.aarch64` | pigsty | 183.1 KiB | [postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-16-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_15` | 3.1.9 | `el8.x86_64` | pgdg | 69.8 KiB | [pgsodium_15-3.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.9 | `el8.aarch64` | pgdg | 66.1 KiB | [pgsodium_15-3.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.8 | `el8.x86_64` | pgdg | 67.6 KiB | [pgsodium_15-3.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.8 | `el8.aarch64` | pgdg | 64.0 KiB | [pgsodium_15-3.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.7 | `el8.x86_64` | pgdg | 67.1 KiB | [pgsodium_15-3.1.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.7-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.7 | `el8.aarch64` | pgdg | 63.6 KiB | [pgsodium_15-3.1.7-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.7-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.6 | `el8.aarch64` | pgdg | 63.1 KiB | [pgsodium_15-3.1.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.6-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.6 | `el8.x86_64` | pgdg | 66.7 KiB | [pgsodium_15-3.1.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.6-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.5 | `el8.x86_64` | pgdg | 66.5 KiB | [pgsodium_15-3.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.5-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.5 | `el8.aarch64` | pgdg | 63.0 KiB | [pgsodium_15-3.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.5-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.4 | `el8.x86_64` | pgdg | 66.1 KiB | [pgsodium_15-3.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.4-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.4 | `el8.aarch64` | pgdg | 62.5 KiB | [pgsodium_15-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.4-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.1 | `el8.x86_64` | pgdg | 65.4 KiB | [pgsodium_15-3.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.1-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.1 | `el8.aarch64` | pgdg | 61.9 KiB | [pgsodium_15-3.1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.1-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.1.0 | `el8.x86_64` | pgdg | 64.8 KiB | [pgsodium_15-3.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.1.0-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.0 | `el8.aarch64` | pgdg | 61.2 KiB | [pgsodium_15-3.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.1.0-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.0.6 | `el8.aarch64` | pgdg | 55.5 KiB | [pgsodium_15-3.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.0.6-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.0.6 | `el8.x86_64` | pgdg | 57.0 KiB | [pgsodium_15-3.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.0.6-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.0.5 | `el8.x86_64` | pgdg | 56.6 KiB | [pgsodium_15-3.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.0.5-1.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.0.5 | `el8.aarch64` | pgdg | 55.2 KiB | [pgsodium_15-3.0.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgsodium_15-3.0.5-1.rhel8.aarch64.rpm) |
| `pgsodium_15` | 3.0.4 | `el8.x86_64` | pgdg | 53.3 KiB | [pgsodium_15-3.0.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgsodium_15-3.0.4-2.rhel8.x86_64.rpm) |
| `pgsodium_15` | 3.1.9 | `el9.x86_64` | pgdg | 69.9 KiB | [pgsodium_15-3.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.9 | `el9.aarch64` | pgdg | 66.8 KiB | [pgsodium_15-3.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.8 | `el9.aarch64` | pgdg | 64.5 KiB | [pgsodium_15-3.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.8 | `el9.x86_64` | pgdg | 67.5 KiB | [pgsodium_15-3.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.7 | `el9.x86_64` | pgdg | 67.1 KiB | [pgsodium_15-3.1.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.7-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.7 | `el9.aarch64` | pgdg | 64.0 KiB | [pgsodium_15-3.1.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.7-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.6 | `el9.aarch64` | pgdg | 63.6 KiB | [pgsodium_15-3.1.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.6-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.6 | `el9.x86_64` | pgdg | 66.7 KiB | [pgsodium_15-3.1.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.6-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.5 | `el9.aarch64` | pgdg | 63.5 KiB | [pgsodium_15-3.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.5-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.5 | `el9.x86_64` | pgdg | 66.5 KiB | [pgsodium_15-3.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.5-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.4 | `el9.x86_64` | pgdg | 66.1 KiB | [pgsodium_15-3.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.4-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.4 | `el9.aarch64` | pgdg | 63.1 KiB | [pgsodium_15-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.4-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.1 | `el9.x86_64` | pgdg | 65.4 KiB | [pgsodium_15-3.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.1-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.1 | `el9.aarch64` | pgdg | 62.4 KiB | [pgsodium_15-3.1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.1-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.1.0 | `el9.x86_64` | pgdg | 64.8 KiB | [pgsodium_15-3.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.1.0-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.1.0 | `el9.aarch64` | pgdg | 61.8 KiB | [pgsodium_15-3.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.1.0-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.0.6 | `el9.aarch64` | pgdg | 56.3 KiB | [pgsodium_15-3.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.0.6-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.0.6 | `el9.x86_64` | pgdg | 57.4 KiB | [pgsodium_15-3.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.0.6-1.rhel9.x86_64.rpm) |
| `pgsodium_15` | 3.0.5 | `el9.aarch64` | pgdg | 56.0 KiB | [pgsodium_15-3.0.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgsodium_15-3.0.5-1.rhel9.aarch64.rpm) |
| `pgsodium_15` | 3.0.5 | `el9.x86_64` | pgdg | 57.0 KiB | [pgsodium_15-3.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgsodium_15-3.0.5-1.rhel9.x86_64.rpm) |
| `postgresql-15-pgsodium` | 3.1.9 | `d12.aarch64` | pigsty | 179.9 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgsodium` | 3.1.9 | `d12.x86_64` | pigsty | 183.8 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgsodium` | 3.1.9 | `u22.x86_64` | pigsty | 195.7 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgsodium` | 3.1.9 | `u22.aarch64` | pigsty | 192.2 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgsodium` | 3.1.9 | `u24.x86_64` | pigsty | 181.2 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgsodium` | 3.1.9 | `u24.aarch64` | pigsty | 177.6 KiB | [postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-15-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_14` | 3.1.9 | `el8.aarch64` | pgdg | 66.1 KiB | [pgsodium_14-3.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.9 | `el8.x86_64` | pgdg | 69.8 KiB | [pgsodium_14-3.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.8 | `el8.x86_64` | pgdg | 67.6 KiB | [pgsodium_14-3.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.8 | `el8.aarch64` | pgdg | 64.0 KiB | [pgsodium_14-3.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.7 | `el8.x86_64` | pgdg | 67.2 KiB | [pgsodium_14-3.1.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.7-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.7 | `el8.aarch64` | pgdg | 63.5 KiB | [pgsodium_14-3.1.7-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.7-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.6 | `el8.aarch64` | pgdg | 63.1 KiB | [pgsodium_14-3.1.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.6-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.6 | `el8.x86_64` | pgdg | 66.7 KiB | [pgsodium_14-3.1.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.6-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.5 | `el8.x86_64` | pgdg | 66.5 KiB | [pgsodium_14-3.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.5-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.5 | `el8.aarch64` | pgdg | 62.9 KiB | [pgsodium_14-3.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.5-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.4 | `el8.aarch64` | pgdg | 62.5 KiB | [pgsodium_14-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.4-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.4 | `el8.x86_64` | pgdg | 66.1 KiB | [pgsodium_14-3.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.4-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.1 | `el8.aarch64` | pgdg | 61.8 KiB | [pgsodium_14-3.1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.1-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.1.1 | `el8.x86_64` | pgdg | 65.4 KiB | [pgsodium_14-3.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.1-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.0 | `el8.x86_64` | pgdg | 64.8 KiB | [pgsodium_14-3.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.1.0-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.0 | `el8.aarch64` | pgdg | 61.2 KiB | [pgsodium_14-3.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.1.0-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.0.6 | `el8.aarch64` | pgdg | 55.6 KiB | [pgsodium_14-3.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.0.6-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.0.6 | `el8.x86_64` | pgdg | 56.9 KiB | [pgsodium_14-3.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.0.6-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.0.5 | `el8.x86_64` | pgdg | 56.6 KiB | [pgsodium_14-3.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.0.5-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.0.5 | `el8.aarch64` | pgdg | 55.2 KiB | [pgsodium_14-3.0.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgsodium_14-3.0.5-1.rhel8.aarch64.rpm) |
| `pgsodium_14` | 3.0.4 | `el8.x86_64` | pgdg | 53.2 KiB | [pgsodium_14-3.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.0.4-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.0.2 | `el8.x86_64` | pgdg | 52.7 KiB | [pgsodium_14-3.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.0.2-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.0.0 | `el8.x86_64` | pgdg | 52.5 KiB | [pgsodium_14-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-3.0.0-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 2.0.2 | `el8.x86_64` | pgdg | 48.6 KiB | [pgsodium_14-2.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgsodium_14-2.0.2-1.rhel8.x86_64.rpm) |
| `pgsodium_14` | 3.1.9 | `el9.x86_64` | pgdg | 69.5 KiB | [pgsodium_14-3.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.9 | `el9.aarch64` | pgdg | 66.8 KiB | [pgsodium_14-3.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.8 | `el9.x86_64` | pgdg | 67.6 KiB | [pgsodium_14-3.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.8 | `el9.aarch64` | pgdg | 64.7 KiB | [pgsodium_14-3.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.7 | `el9.x86_64` | pgdg | 67.1 KiB | [pgsodium_14-3.1.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.7-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.7 | `el9.aarch64` | pgdg | 64.2 KiB | [pgsodium_14-3.1.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.7-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.6 | `el9.aarch64` | pgdg | 63.8 KiB | [pgsodium_14-3.1.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.6-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.6 | `el9.x86_64` | pgdg | 66.6 KiB | [pgsodium_14-3.1.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.6-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.5 | `el9.x86_64` | pgdg | 66.5 KiB | [pgsodium_14-3.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.5-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.5 | `el9.aarch64` | pgdg | 63.6 KiB | [pgsodium_14-3.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.5-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.4 | `el9.aarch64` | pgdg | 63.2 KiB | [pgsodium_14-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.4-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.4 | `el9.x86_64` | pgdg | 66.1 KiB | [pgsodium_14-3.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.4-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.1 | `el9.aarch64` | pgdg | 62.4 KiB | [pgsodium_14-3.1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.1-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.1.1 | `el9.x86_64` | pgdg | 65.4 KiB | [pgsodium_14-3.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.1-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.0 | `el9.x86_64` | pgdg | 64.8 KiB | [pgsodium_14-3.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.1.0-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.1.0 | `el9.aarch64` | pgdg | 61.8 KiB | [pgsodium_14-3.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.1.0-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.0.6 | `el9.x86_64` | pgdg | 57.4 KiB | [pgsodium_14-3.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.0.6-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.0.6 | `el9.aarch64` | pgdg | 56.3 KiB | [pgsodium_14-3.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.0.6-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 3.0.5 | `el9.x86_64` | pgdg | 57.1 KiB | [pgsodium_14-3.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-3.0.5-1.rhel9.x86_64.rpm) |
| `pgsodium_14` | 3.0.5 | `el9.aarch64` | pgdg | 56.0 KiB | [pgsodium_14-3.0.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgsodium_14-3.0.5-1.rhel9.aarch64.rpm) |
| `pgsodium_14` | 2.0.2 | `el9.x86_64` | pgdg | 49.3 KiB | [pgsodium_14-2.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgsodium_14-2.0.2-1.rhel9.x86_64.rpm) |
| `postgresql-14-pgsodium` | 3.1.9 | `d12.aarch64` | pigsty | 179.8 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgsodium` | 3.1.9 | `d12.x86_64` | pigsty | 183.6 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgsodium` | 3.1.9 | `u22.aarch64` | pigsty | 192.1 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgsodium` | 3.1.9 | `u22.x86_64` | pigsty | 195.6 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgsodium` | 3.1.9 | `u24.x86_64` | pigsty | 181.0 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgsodium` | 3.1.9 | `u24.aarch64` | pigsty | 177.4 KiB | [postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-14-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgsodium_13` | 3.1.9 | `el8.aarch64` | pgdg | 66.1 KiB | [pgsodium_13-3.1.9-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.9-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.9 | `el8.x86_64` | pgdg | 69.8 KiB | [pgsodium_13-3.1.9-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.9-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.8 | `el8.aarch64` | pgdg | 64.1 KiB | [pgsodium_13-3.1.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.8-1PGDG.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.8 | `el8.x86_64` | pgdg | 67.7 KiB | [pgsodium_13-3.1.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.8-1PGDG.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.7 | `el8.aarch64` | pgdg | 63.6 KiB | [pgsodium_13-3.1.7-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.7-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.7 | `el8.x86_64` | pgdg | 67.2 KiB | [pgsodium_13-3.1.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.7-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.6 | `el8.aarch64` | pgdg | 63.1 KiB | [pgsodium_13-3.1.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.6-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.6 | `el8.x86_64` | pgdg | 66.7 KiB | [pgsodium_13-3.1.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.6-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.5 | `el8.aarch64` | pgdg | 62.9 KiB | [pgsodium_13-3.1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.5-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.5 | `el8.x86_64` | pgdg | 66.6 KiB | [pgsodium_13-3.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.5-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.4 | `el8.x86_64` | pgdg | 66.1 KiB | [pgsodium_13-3.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.4-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.4 | `el8.aarch64` | pgdg | 62.5 KiB | [pgsodium_13-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.4-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.1 | `el8.aarch64` | pgdg | 61.8 KiB | [pgsodium_13-3.1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.1-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.1.1 | `el8.x86_64` | pgdg | 65.5 KiB | [pgsodium_13-3.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.1-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.0 | `el8.x86_64` | pgdg | 64.8 KiB | [pgsodium_13-3.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.1.0-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.0 | `el8.aarch64` | pgdg | 61.2 KiB | [pgsodium_13-3.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.1.0-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.0.6 | `el8.aarch64` | pgdg | 55.6 KiB | [pgsodium_13-3.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.0.6-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.0.6 | `el8.x86_64` | pgdg | 56.9 KiB | [pgsodium_13-3.0.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.0.6-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.0.5 | `el8.x86_64` | pgdg | 56.6 KiB | [pgsodium_13-3.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.0.5-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.0.5 | `el8.aarch64` | pgdg | 55.2 KiB | [pgsodium_13-3.0.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgsodium_13-3.0.5-1.rhel8.aarch64.rpm) |
| `pgsodium_13` | 3.0.4 | `el8.x86_64` | pgdg | 53.2 KiB | [pgsodium_13-3.0.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.0.4-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.0.2 | `el8.x86_64` | pgdg | 52.7 KiB | [pgsodium_13-3.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.0.2-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.0.0 | `el8.x86_64` | pgdg | 52.4 KiB | [pgsodium_13-3.0.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-3.0.0-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 2.0.2 | `el8.x86_64` | pgdg | 48.5 KiB | [pgsodium_13-2.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgsodium_13-2.0.2-1.rhel8.x86_64.rpm) |
| `pgsodium_13` | 3.1.9 | `el9.aarch64` | pgdg | 66.8 KiB | [pgsodium_13-3.1.9-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.9-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.9 | `el9.x86_64` | pgdg | 69.4 KiB | [pgsodium_13-3.1.9-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.9-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.8 | `el9.x86_64` | pgdg | 67.5 KiB | [pgsodium_13-3.1.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.8-1PGDG.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.8 | `el9.aarch64` | pgdg | 64.6 KiB | [pgsodium_13-3.1.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.8-1PGDG.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.7 | `el9.aarch64` | pgdg | 64.2 KiB | [pgsodium_13-3.1.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.7-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.7 | `el9.x86_64` | pgdg | 67.1 KiB | [pgsodium_13-3.1.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.7-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.6 | `el9.x86_64` | pgdg | 66.6 KiB | [pgsodium_13-3.1.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.6-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.6 | `el9.aarch64` | pgdg | 63.7 KiB | [pgsodium_13-3.1.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.6-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.5 | `el9.x86_64` | pgdg | 66.4 KiB | [pgsodium_13-3.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.5-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.5 | `el9.aarch64` | pgdg | 63.6 KiB | [pgsodium_13-3.1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.5-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.4 | `el9.x86_64` | pgdg | 66.1 KiB | [pgsodium_13-3.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.4-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.4 | `el9.aarch64` | pgdg | 63.2 KiB | [pgsodium_13-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.4-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.1 | `el9.aarch64` | pgdg | 62.4 KiB | [pgsodium_13-3.1.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.1-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.1 | `el9.x86_64` | pgdg | 65.4 KiB | [pgsodium_13-3.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.1-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.1.0 | `el9.aarch64` | pgdg | 61.8 KiB | [pgsodium_13-3.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.1.0-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.1.0 | `el9.x86_64` | pgdg | 64.7 KiB | [pgsodium_13-3.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.1.0-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.0.6 | `el9.x86_64` | pgdg | 57.4 KiB | [pgsodium_13-3.0.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.0.6-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 3.0.6 | `el9.aarch64` | pgdg | 56.3 KiB | [pgsodium_13-3.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.0.6-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.0.5 | `el9.aarch64` | pgdg | 56.0 KiB | [pgsodium_13-3.0.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgsodium_13-3.0.5-1.rhel9.aarch64.rpm) |
| `pgsodium_13` | 3.0.5 | `el9.x86_64` | pgdg | 57.2 KiB | [pgsodium_13-3.0.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-3.0.5-1.rhel9.x86_64.rpm) |
| `pgsodium_13` | 2.0.2 | `el9.x86_64` | pgdg | 49.3 KiB | [pgsodium_13-2.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgsodium_13-2.0.2-1.rhel9.x86_64.rpm) |
| `postgresql-13-pgsodium` | 3.1.9 | `d12.x86_64` | pigsty | 179.7 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgsodium` | 3.1.9 | `d12.aarch64` | pigsty | 175.6 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgsodium` | 3.1.9 | `u22.x86_64` | pigsty | 191.4 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgsodium` | 3.1.9 | `u22.aarch64` | pigsty | 187.5 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgsodium` | 3.1.9 | `u24.aarch64` | pigsty | 172.2 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgsodium` | 3.1.9 | `u24.x86_64` | pigsty | 175.6 KiB | [postgresql-13-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgsodium/postgresql-13-pgsodium_3.1.9-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michelp/pgsodium" title="Repository" icon="github" subtitle="github.com/michelp/pgsodium" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pgsodium-3.1.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgsodium; # get pgsodium source code
pig build dep pgsodium; # install build dependencies
pig build pkg pgsodium; # build extension rpm or deb
pig build ext pgsodium; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgsodium; # install by extension name, for the current active PG version
pig ext install pgsodium; # install via package alias, for the active PG version
pig ext install pgsodium -v 18;   # install for PG 18
pig ext install pgsodium -v 17;   # install for PG 17
pig ext install pgsodium -v 16;   # install for PG 16
pig ext install pgsodium -v 15;   # install for PG 15
pig ext install pgsodium -v 14;   # install for PG 14
pig ext install pgsodium -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgsodium CASCADE SCHEMA pgsodium;
```

