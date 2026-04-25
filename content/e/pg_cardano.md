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
| **2920** | {{< badge content="pg_cardano" link="https://github.com/Fell-x27/pg_cardano" >}} | {{< ext "pg_cardano" >}} | `1.2.0` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `pg_cardano` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "pg_cardano_18" "green" >}} {{< bg "17" "pg_cardano_17" "green" >}} {{< bg "16" "pg_cardano_16" "green" >}} {{< bg "15" "pg_cardano_15" "green" >}} {{< bg "14" "pg_cardano_14" "red" >}} | `pg_cardano_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.0` | {{< bg "18" "postgresql-18-pg-cardano" "green" >}} {{< bg "17" "postgresql-17-pg-cardano" "green" >}} {{< bg "16" "postgresql-16-pg-cardano" "green" >}} {{< bg "15" "postgresql-15-pg-cardano" "green" >}} {{< bg "14" "postgresql-14-pg-cardano" "red" >}} | `postgresql-$v-pg-cardano` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_cardano_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pg_cardano_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-cardano : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-pg-cardano : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 518.4 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_18-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.9 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_18-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.6 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_18-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 404.4 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_18-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 537.0 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_18-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 404.0 KiB | [pg_cardano_18-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_18-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-cardano` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 434.7 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 305.4 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 434.2 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 305.3 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 483.1 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 351.7 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 478.2 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-cardano` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 346.1 KiB | [postgresql-18-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-18-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 518.0 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_17-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.3 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_17-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 535.9 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_17-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 404.0 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_17-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 536.1 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_17-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.5 KiB | [pg_cardano_17-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_17-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-cardano` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 434.1 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 305.0 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 434.0 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 304.8 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 482.7 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 351.4 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 477.8 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cardano` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 345.8 KiB | [postgresql-17-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-17-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 518.9 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_16-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.8 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_16-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.6 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_16-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 404.5 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_16-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 537.0 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_16-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 404.1 KiB | [pg_cardano_16-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_16-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-cardano` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 434.7 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 305.3 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 434.6 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 305.1 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 483.7 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 351.4 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 478.1 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cardano` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 346.0 KiB | [postgresql-16-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-16-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cardano_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 518.0 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cardano_15-1.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_cardano_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 379.4 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cardano_15-1.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_cardano_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 536.0 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cardano_15-1.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_cardano_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 403.7 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cardano_15-1.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_cardano_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 536.3 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cardano_15-1.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_cardano_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 403.5 KiB | [pg_cardano_15-1.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cardano_15-1.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-cardano` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 433.6 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 305.3 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 433.5 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 304.7 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 483.0 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 351.0 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 477.5 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cardano` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 345.6 KiB | [postgresql-15-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cardano/postgresql-15-pg-cardano_1.2.0-1PIGSTY~noble_arm64.deb) |

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

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Fell-x27/pg_cardano" title="Repository" icon="github" subtitle="github.com/Fell-x27/pg_cardano" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cardano-1.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_cardano;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cardano;		# install via package name, for the active PG version

pig install pg_cardano -v 18;   # install for PG 18
pig install pg_cardano -v 17;   # install for PG 17
pig install pg_cardano -v 16;   # install for PG 16
pig install pg_cardano -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cardano;
```

## Usage

Sources: [README](https://github.com/Fell-x27/pg_cardano/blob/master/README.md), [Cargo.toml version 1.2.0](https://github.com/Fell-x27/pg_cardano/blob/master/Cargo.toml)

`pg_cardano` is a Rust extension for Cardano-oriented binary and crypto utilities inside PostgreSQL. The upstream repo does not publish GitHub releases or tags, but the current crate version on the default branch is `1.2.0`.

```sql
CREATE EXTENSION pg_cardano;
```

### Core Capabilities

- Base58 encode/decode.
- Bech32 encode/decode.
- CBOR to and from `jsonb`, with both simple and extended pipelines.
- Blake2b hashing.
- Ed25519 signing and signature verification.
- dRep ID helpers for CIP-105 and CIP-129.
- Shelley address builders and parsers.
- Asset-name decoding and CIP-88 pool key registration verification.

### Common Patterns

Base58 and Bech32:

```sql
SELECT cardano.base58_encode('Cardano'::bytea);
SELECT cardano.base58_decode('3Z6ioYHE3x');
SELECT cardano.bech32_encode('ada', 'is amazing'::bytea);
SELECT cardano.bech32_decode_prefix('ada1d9ejqctdv9axjmn8dypl4d');
SELECT cardano.bech32_decode_data('ada1d9ejqctdv9axjmn8dypl4d');
```

CBOR conversion:

```sql
SELECT cardano.cbor_decode_jsonb('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb('{"ada": "is amazing!", "bytes": "\\xdeadbeef"}'::jsonb);
SELECT cardano.cbor_decode_jsonb_ext('\xa3636164616b...'::bytea);
SELECT cardano.cbor_encode_jsonb_ext('{"type":"map","value":[...]}'::jsonb);
```

Hashing and signatures:

```sql
SELECT cardano.blake2b_hash('Cardano is amazing!'::bytea, 32);
SELECT cardano.ed25519_verify_signature(
  '\x432753BD...'::bytea,
  'Cardano is amazing!'::bytea,
  '\x74265f96...'::bytea
);
```

Address and governance helpers:

```sql
SELECT cardano.tools_drep_id_encode_cip105('\x28111ae1...'::bytea, FALSE);
SELECT cardano.tools_drep_id_encode_cip129('\x28111ae1...'::bytea, TRUE);
SELECT cardano.tools_shelley_address_build(
  '\x7415251f...'::bytea, FALSE,
  '\x7c3ae2f2...'::bytea, FALSE,
  0
);
SELECT cardano.tools_shelley_addr_get_type('addr_test1vp6p2fgl...');
```

### Caveats

- Upstream documents PostgreSQL 12+ and Linux builds via `pgrx`; the crate features currently target PostgreSQL 15 through 18, with `pg18` as the default feature.
- The simple CBOR helpers are intentionally lossy for some CBOR constructs; use the `*_ext` functions when exact structural round-tripping matters.
