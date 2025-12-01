---
title: "pgcryptokey"
linkTitle: "pgcryptokey"
description: "cryptographic key management"
weight: 7320
categories: ["SEC"]
width: full
---

[**pgcryptokey**](https://momjian.us/download/pgcryptokey/) : cryptographic key management


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7320** | {{< badge content="pgcryptokey" link="https://momjian.us/download/pgcryptokey/" >}} | {{< ext "pgcryptokey" >}} | `0.85` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcrypto" >}} |
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pg_tde" >}} {{< ext "faker" >}} {{< ext "passwordcheck_cracklib" >}} {{< ext "supautils" >}} {{< ext "supabase_vault" >}} |

> [!Note] missing 14,13,12 on el pgdg repo


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pgcryptokey` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "pgcryptokey_18" "green" >}} {{< bg "17" "pgcryptokey_17" "green" >}} {{< bg "16" "pgcryptokey_16" "green" >}} {{< bg "15" "pgcryptokey_15" "green" >}} {{< bg "14" "pgcryptokey_14" "green" >}} {{< bg "13" "pgcryptokey_13" "green" >}} | `pgcryptokey_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.85` | {{< bg "18" "postgresql-18-pgcryptokey" "green" >}} {{< bg "17" "postgresql-17-pgcryptokey" "green" >}} {{< bg "16" "postgresql-16-pgcryptokey" "green" >}} {{< bg "15" "postgresql-15-pgcryptokey" "green" >}} {{< bg "14" "postgresql-14-pgcryptokey" "green" >}} {{< bg "13" "postgresql-13-pgcryptokey" "green" >}} | `postgresql-$v-pgcryptokey` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.85" "pgcryptokey_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 0.85" "pgcryptokey_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.85" "pgcryptokey_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.85" "postgresql-18-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-17-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-16-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-15-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-14-pgcryptokey : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.85" "postgresql-13-pgcryptokey : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_18` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_18-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_18-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_18` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_18-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_18-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_18` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_18-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_18-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_18` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_18-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_18-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.5 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.5 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-18-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_17` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.1 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgcryptokey_17-0.85-6PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_17-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.2 KiB | [pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgcryptokey_17-0.85-6PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_17-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.5 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgcryptokey_17-0.85-6PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_17-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.4 KiB | [pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgcryptokey_17-0.85-6PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_17-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgcryptokey_17-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_17-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_17-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgcryptokey_17-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_17` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_17-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_17-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-17-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_16` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgcryptokey_16-0.85-5PGDG.rhel8.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_16-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 18.1 KiB | [pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgcryptokey_16-0.85-5PGDG.rhel8.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_16-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgcryptokey_16-0.85-5PGDG.rhel9.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_16-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.0 KiB | [pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgcryptokey_16-0.85-5PGDG.rhel9.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_16-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgcryptokey_16-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_16-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_16-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgcryptokey_16-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_16` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_16-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_16-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-16-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_15` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgcryptokey_15-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_15-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgcryptokey_15-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_15-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pgdg | 22.7 KiB | [pgcryptokey_15-0.85-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgcryptokey_15-0.85-3.rhel9.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_15-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_15-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgcryptokey_15-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.9 KiB | [pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_15-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgcryptokey_15-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_15-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_15-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgcryptokey_15-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_15` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_15-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_15-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.7 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.8 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-15-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_14` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pgdg | 22.6 KiB | [pgcryptokey_14-0.85-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgcryptokey_14-0.85-3.rhel8.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_14-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_14-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgcryptokey_14-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_14-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_14-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.3 KiB | [pgcryptokey_14-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgcryptokey_14-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_14-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgcryptokey_14-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_14-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_14-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgcryptokey_14-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_14` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_14-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_14-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.7 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-14-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcryptokey_13` | `0.85` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_13-0.85-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcryptokey_13-0.85-1PIGSTY.el8.x86_64.rpm) |
| `pgcryptokey_13` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pgdg | 22.4 KiB | [pgcryptokey_13-0.85-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgcryptokey_13-0.85-3.rhel8.aarch64.rpm) |
| `pgcryptokey_13` | `0.85` | [el8.aarch64](/os/el8.aarch64) | pigsty | 17.1 KiB | [pgcryptokey_13-0.85-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcryptokey_13-0.85-1PIGSTY.el8.aarch64.rpm) |
| `pgcryptokey_13` | `0.85` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_13-0.85-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcryptokey_13-0.85-1PIGSTY.el9.x86_64.rpm) |
| `pgcryptokey_13` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.3 KiB | [pgcryptokey_13-0.85-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgcryptokey_13-0.85-3.rhel9.aarch64.rpm) |
| `pgcryptokey_13` | `0.85` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.8 KiB | [pgcryptokey_13-0.85-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcryptokey_13-0.85-1PIGSTY.el9.aarch64.rpm) |
| `pgcryptokey_13` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pgcryptokey_13-0.85-8PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgcryptokey_13-0.85-8PGDG.rhel10.x86_64.rpm) |
| `pgcryptokey_13` | `0.85` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgcryptokey_13-0.85-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcryptokey_13-0.85-1PIGSTY.el10.x86_64.rpm) |
| `pgcryptokey_13` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.9 KiB | [pgcryptokey_13-0.85-8PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgcryptokey_13-0.85-8PGDG.rhel10.aarch64.rpm) |
| `pgcryptokey_13` | `0.85` | [el10.aarch64](/os/el10.aarch64) | pigsty | 17.0 KiB | [pgcryptokey_13-0.85-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcryptokey_13-0.85-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgcryptokey` | `0.85` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.4 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.5 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.4 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.6 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.7 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.6 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgcryptokey` | `0.85` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.4 KiB | [postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgcryptokey/postgresql-13-pgcryptokey_0.85-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://momjian.us/download/pgcryptokey/" title="Repository" icon="link" subtitle="momjian.us/download/pgcryptokey/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcryptokey-0.85.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgcryptokey;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgcryptokey;		# install via package name, for the active PG version

pig install pgcryptokey -v 18;   # install for PG 18
pig install pgcryptokey -v 17;   # install for PG 17
pig install pgcryptokey -v 16;   # install for PG 16
pig install pgcryptokey -v 15;   # install for PG 15
pig install pgcryptokey -v 14;   # install for PG 14
pig install pgcryptokey -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgcryptokey CASCADE; -- requires pgcrypto
```
