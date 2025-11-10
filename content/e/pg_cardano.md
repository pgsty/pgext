---
title: "pg_cardano"
linkTitle: "pg_cardano"
description: "A suite of Cardano-related tools"
weight: 2920
categories: ["FEAT"]
width: full
---

[**pg_cardano**](https://github.com/Fell-x27/pg_cardano) : A suite of Cardano-related tools


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2920** | {{< badge content="pg_cardano" link="https://github.com/Fell-x27/pg_cardano" >}} | {{< ext "pg_cardano" >}} | `1.1.1` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] PG18 fix by https://github.com/Vonng/pg_cardano


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_cardano` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "pg_cardano_18" "green" >}} {{< bg "17" "pg_cardano_17" "green" >}} {{< bg "16" "pg_cardano_16" "green" >}} {{< bg "15" "pg_cardano_15" "green" >}} {{< bg "14" "pg_cardano_14" "green" >}} {{< bg "13" "pg_cardano_13" "green" >}} | `pg_cardano_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-pg-cardano" "green" >}} {{< bg "17" "postgresql-17-pg-cardano" "green" >}} {{< bg "16" "postgresql-16-pg-cardano" "green" >}} {{< bg "15" "postgresql-15-pg-cardano" "green" >}} {{< bg "14" "postgresql-14-pg-cardano" "green" >}} {{< bg "13" "postgresql-13-pg-cardano" "green" >}} | `postgresql-$v-pg-cardano` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-13-pg-cardano : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_18` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 523.1 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_18-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_18` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 378.9 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_18-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_18` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 540.7 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_18-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_18` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.5 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_18-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_18` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 539.5 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_18-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_18` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.5 KiB | [pg_cardano_18-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_18-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 445.2 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.7 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.3 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 312.0 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 493.9 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 356.7 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.3 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.1 KiB | [postgresql-18-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_17` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.9 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_17-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_17` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.5 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_17-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_17` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 540.1 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_17-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_17` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.9 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_17-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_17` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 539.7 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_17-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_17` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.5 KiB | [pg_cardano_17-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_17-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 444.9 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.7 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.3 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 311.8 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 493.7 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.2 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.0 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.0 KiB | [postgresql-17-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_16` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.7 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_16-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_16` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 378.9 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_16-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_16` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.7 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_16-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_16` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.4 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_16-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_16` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 539.3 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_16-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_16` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.4 KiB | [pg_cardano_16-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_16-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 445.4 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 312.0 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.8 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 311.9 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 494.1 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.1 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.3 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.5 KiB | [postgresql-16-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.7 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_15-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.0 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_15-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.4 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_15-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.6 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_15-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_15` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 538.9 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_15-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_15` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.7 KiB | [pg_cardano_15-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_15-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 445.8 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.5 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.5 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 312.2 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 494.0 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.4 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.4 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.5 KiB | [postgresql-15-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 522.9 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_14-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.5 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_14-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 539.5 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_14-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.9 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_14-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_14` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 538.9 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_14-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_14` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.4 KiB | [pg_cardano_14-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_14-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 445.3 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.5 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 445.1 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 311.6 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 494.2 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 357.0 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.4 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 351.9 KiB | [postgresql-14-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-14-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_13` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 523.7 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_13-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_13` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.0 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_13-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_13` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 540.5 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_13-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_13` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 402.5 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_13-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_13` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 540.0 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_13-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_13` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.7 KiB | [pg_cardano_13-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_13-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-cardano` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 446.1 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 311.8 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 446.4 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 311.7 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 494.4 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 356.9 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 488.8 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-cardano` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 352.2 KiB | [postgresql-13-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-13-pg-cardano_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_cardano" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_cardano" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cardano-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_cardano;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cardano;		# install via package name, for the active PG version

pig install pg_cardano -v 18;   # install for PG 18
pig install pg_cardano -v 17;   # install for PG 17
pig install pg_cardano -v 16;   # install for PG 16
pig install pg_cardano -v 15;   # install for PG 15
pig install pg_cardano -v 14;   # install for PG 14
pig install pg_cardano -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cardano;
```
