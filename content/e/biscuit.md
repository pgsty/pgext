---
title: "biscuit"
linkTitle: "biscuit"
description: "IAM-LIKE pattern matching with bitmap indexing"
weight: 2170
categories: ["FTS"]
width: full
---

[**pg_biscuit**](https://github.com/CrystallineCore/pg_biscuit) : IAM-LIKE pattern matching with bitmap indexing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2170** | {{< badge content="biscuit" link="https://github.com/CrystallineCore/pg_biscuit" >}} | {{< ext "biscuit" "pg_biscuit" >}} | `2.2.2` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_textsearch" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_biscuit` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.2` | {{< bg "18" "pg_biscuit_18" "green" >}} {{< bg "17" "pg_biscuit_17" "green" >}} {{< bg "16" "pg_biscuit_16" "green" >}} {{< bg "15" "pg_biscuit_15" "red" >}} {{< bg "14" "pg_biscuit_14" "red" >}} {{< bg "13" "pg_biscuit_13" "red" >}} | `pg_biscuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.2` | {{< bg "18" "postgresql-18-biscuit" "green" >}} {{< bg "17" "postgresql-17-biscuit" "green" >}} {{< bg "16" "postgresql-16-biscuit" "green" >}} {{< bg "15" "postgresql-15-biscuit" "red" >}} {{< bg "14" "postgresql-14-biscuit" "red" >}} {{< bg "13" "postgresql-13-biscuit" "red" >}} | `postgresql-$v-biscuit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "pg_biscuit_16 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_biscuit_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.2" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-biscuit : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_biscuit_18` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_biscuit_18-2.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_biscuit_18` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.0 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_biscuit_18-2.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 65.8 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_biscuit_18-2.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.6 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_biscuit_18-2.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 68.5 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_biscuit_18-2.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 65.6 KiB | [pg_biscuit_18-2.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_biscuit_18-2.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-biscuit` | `2.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 154.0 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.3 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 155.3 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.4 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 157.5 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 155.2 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 157.1 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-biscuit` | `2.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 153.6 KiB | [postgresql-18-biscuit_2.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_biscuit_17` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.3 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_biscuit_17-2.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_biscuit_17` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.9 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_biscuit_17-2.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 65.5 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_biscuit_17-2.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.4 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_biscuit_17-2.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 68.5 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_biscuit_17-2.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 65.6 KiB | [pg_biscuit_17-2.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_biscuit_17-2.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-biscuit` | `2.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 154.0 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.0 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 155.1 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.0 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.2 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.2 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 157.2 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-biscuit` | `2.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 153.2 KiB | [postgresql-17-biscuit_2.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_biscuit_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_biscuit_16-2.2.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_biscuit_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.0 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_biscuit_16-2.2.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 65.6 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_biscuit_16-2.2.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.5 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_biscuit_16-2.2.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 68.6 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_biscuit_16-2.2.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 65.6 KiB | [pg_biscuit_16-2.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_biscuit_16-2.2.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-biscuit` | `2.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 154.0 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 149.0 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 155.1 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.0 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.1 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.0 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 157.2 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-biscuit` | `2.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 153.3 KiB | [postgresql-16-biscuit_2.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrystallineCore/pg_biscuit" title="Repository" icon="github" subtitle="github.com/CrystallineCore/pg_biscuit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="Biscuit-2.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_biscuit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_biscuit;		# install via package name, for the active PG version
pig install biscuit;		# install by extension name, for the current active PG version

pig install biscuit -v 18;   # install for PG 18
pig install biscuit -v 17;   # install for PG 17
pig install biscuit -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION biscuit CASCADE; -- requires plpgsql
```
