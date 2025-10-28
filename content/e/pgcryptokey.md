---
title: "pgcryptokey"
linkTitle: "pgcryptokey"
description: "cryptographic key management"
weight: 7120
categories: ["SEC"]
width: full
---

cryptographic key management


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7120** | {{< badge content="pgcryptokey" link="https://momjian.us/download/pgcryptokey/" >}} | {{< ext "pgcryptokey" >}} | `0.85` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcrypto" >}} |
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pg_tde" >}} {{< ext "faker" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "supabase_vault" >}} |

> [!Note] missing 14,13,12 on el pgdg repo


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgcryptokey" >}} | `0.85` | {{< bg "18" "pgcryptokey_18" "green" >}} {{< bg "17" "pgcryptokey_17" "green" >}} {{< bg "16" "pgcryptokey_16" "green" >}} {{< bg "15" "pgcryptokey_15" "green" >}} {{< bg "14" "pgcryptokey_14" "green" >}} {{< bg "13" "pgcryptokey_13" "green" >}} | `pgcryptokey_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgcryptokey" >}} | `0.85` | {{< bg "18" "postgresql-18-pgcryptokey" "red" >}} {{< bg "17" "postgresql-17-pgcryptokey" "green" >}} {{< bg "16" "postgresql-16-pgcryptokey" "green" >}} {{< bg "15" "postgresql-15-pgcryptokey" "green" >}} {{< bg "14" "postgresql-14-pgcryptokey" "green" >}} {{< bg "13" "postgresql-13-pgcryptokey" "green" >}} | `postgresql-$v-pgcryptokey` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "pgcryptokey_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    |      {{< bg "MISS" "pgcryptokey_18 : MISS 0" "red" >}}      | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgcryptokey : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgcryptokey : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgcryptokey : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgcryptokey : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_18` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_18` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_18` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_18` | 0.85 | `el9.aarch64` | pigsty | 16.9 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_17` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_17` | 0.85 | `el8.x86_64` | pgdg | 18.1 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_17` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_17` | 0.85 | `el8.aarch64` | pgdg | 18.2 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_17` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_17` | 0.85 | `el9.x86_64` | pgdg | 17.5 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_17` | 0.85 | `el9.aarch64` | pigsty | 17.0 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_17` | 0.85 | `el9.aarch64` | pgdg | 17.4 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_17` | 0.85 | `el10.x86_64` | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_17` | 0.85 | `el10.aarch64` | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgcryptokey` | 0.85 | `d12.x86_64` | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgcryptokey` | 0.85 | `d12.aarch64` | pigsty | 11.7 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgcryptokey` | 0.85 | `u22.x86_64` | pigsty | 11.7 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgcryptokey` | 0.85 | `u22.aarch64` | pigsty | 11.8 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgcryptokey` | 0.85 | `u24.x86_64` | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgcryptokey` | 0.85 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_16` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_16` | 0.85 | `el8.x86_64` | pgdg | 18.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_16` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_16` | 0.85 | `el8.aarch64` | pgdg | 18.1 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_16` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_16` | 0.85 | `el9.x86_64` | pgdg | 17.3 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_16` | 0.85 | `el9.aarch64` | pigsty | 17.0 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_16` | 0.85 | `el9.aarch64` | pgdg | 17.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_16` | 0.85 | `el10.x86_64` | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_16` | 0.85 | `el10.aarch64` | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgcryptokey` | 0.85 | `d12.x86_64` | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgcryptokey` | 0.85 | `d12.aarch64` | pigsty | 11.7 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgcryptokey` | 0.85 | `u22.x86_64` | pigsty | 11.7 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgcryptokey` | 0.85 | `u22.aarch64` | pigsty | 11.8 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgcryptokey` | 0.85 | `u24.x86_64` | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgcryptokey` | 0.85 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_15` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_15` | 0.85 | `el8.x86_64` | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgcryptokey_15-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_15` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_15` | 0.85 | `el8.aarch64` | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgcryptokey_15-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_15` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_15` | 0.85 | `el9.x86_64` | pgdg | 22.7 KiB | [pgcryptokey_15-0.85-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgcryptokey_15-0.85-3.rhel9.x86_64.rpm) |
| `pgcryptokey_15` | 0.85 | `el9.aarch64` | pigsty | 17.0 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_15` | 0.85 | `el9.aarch64` | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgcryptokey_15-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_15` | 0.85 | `el10.x86_64` | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_15` | 0.85 | `el10.aarch64` | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgcryptokey` | 0.85 | `d12.x86_64` | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgcryptokey` | 0.85 | `d12.aarch64` | pigsty | 11.7 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgcryptokey` | 0.85 | `u22.x86_64` | pigsty | 11.7 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgcryptokey` | 0.85 | `u22.aarch64` | pigsty | 11.8 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgcryptokey` | 0.85 | `u24.x86_64` | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgcryptokey` | 0.85 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_14` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_14` | 0.85 | `el8.x86_64` | pgdg | 22.6 KiB | [pgcryptokey_14-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgcryptokey_14-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_14` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_14` | 0.85 | `el8.aarch64` | pgdg | 22.4 KiB | [pgcryptokey_14-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgcryptokey_14-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_14` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_14` | 0.85 | `el9.aarch64` | pigsty | 17.0 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_14` | 0.85 | `el9.aarch64` | pgdg | 22.3 KiB | [pgcryptokey_14-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgcryptokey_14-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_14` | 0.85 | `el10.x86_64` | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_14` | 0.85 | `el10.aarch64` | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgcryptokey` | 0.85 | `d12.x86_64` | pigsty | 11.5 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgcryptokey` | 0.85 | `d12.aarch64` | pigsty | 11.7 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgcryptokey` | 0.85 | `u22.x86_64` | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgcryptokey` | 0.85 | `u22.aarch64` | pigsty | 11.7 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgcryptokey` | 0.85 | `u24.x86_64` | pigsty | 11.5 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgcryptokey` | 0.85 | `u24.aarch64` | pigsty | 11.3 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_13` | 0.85 | `el8.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_13-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_13-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_13` | 0.85 | `el8.aarch64` | pigsty | 17.2 KiB | [pgcryptokey_13-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_13-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_13` | 0.85 | `el8.aarch64` | pgdg | 22.4 KiB | [pgcryptokey_13-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgcryptokey_13-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_13` | 0.85 | `el9.x86_64` | pigsty | 17.1 KiB | [pgcryptokey_13-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_13-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_13` | 0.85 | `el9.aarch64` | pigsty | 17.0 KiB | [pgcryptokey_13-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_13-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_13` | 0.85 | `el9.aarch64` | pgdg | 22.3 KiB | [pgcryptokey_13-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgcryptokey_13-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_13` | 0.85 | `el10.x86_64` | pgdg | 17.9 KiB | [pgcryptokey_13-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgcryptokey_13-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_13` | 0.85 | `el10.aarch64` | pgdg | 17.9 KiB | [pgcryptokey_13-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgcryptokey_13-0.85-8PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgcryptokey` | 0.85 | `d12.x86_64` | pigsty | 11.5 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgcryptokey` | 0.85 | `d12.aarch64` | pigsty | 11.7 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgcryptokey` | 0.85 | `u22.x86_64` | pigsty | 11.7 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgcryptokey` | 0.85 | `u22.aarch64` | pigsty | 11.7 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgcryptokey` | 0.85 | `u24.x86_64` | pigsty | 11.5 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgcryptokey` | 0.85 | `u24.aarch64` | pigsty | 11.4 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://momjian.us/download/pgcryptokey/" title="Repository" icon="link" subtitle="momjian.us/download/pgcryptokey/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcryptokey-0.85.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgcryptokey; # get pgcryptokey source code
pig build dep pgcryptokey; # install build dependencies
pig build pkg pgcryptokey; # build extension rpm or deb
pig build ext pgcryptokey; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgcryptokey; # install by extension name, for the current active PG version
pig ext install pgcryptokey; # install via package alias, for the active PG version
pig ext install pgcryptokey -v 18;   # install for PG 18
pig ext install pgcryptokey -v 17;   # install for PG 17
pig ext install pgcryptokey -v 16;   # install for PG 16
pig ext install pgcryptokey -v 15;   # install for PG 15
pig ext install pgcryptokey -v 14;   # install for PG 14
pig ext install pgcryptokey -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgcryptokey;
```

