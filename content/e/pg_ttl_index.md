---
title: "pg_ttl_index"
linkTitle: "pg_ttl_index"
description: "Automatic data expiration with TTL indexes"
weight: 2780
categories: ["FEAT"]
width: full
---

[**pg_ttl_index**](https://github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl) : Automatic data expiration with TTL indexes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2780** | {{< badge content="pg_ttl_index" link="https://github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl" >}} | {{< ext "pg_ttl_index" >}} | `2.0.0` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "temporal_tables" >}} {{< ext "periods" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_partman" >}} {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "timescaledb" >}} |

> [!Note] pg13/14 breaks


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_ttl_index` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "pg_ttl_index_18" "green" >}} {{< bg "17" "pg_ttl_index_17" "green" >}} {{< bg "16" "pg_ttl_index_16" "green" >}} {{< bg "15" "pg_ttl_index_15" "green" >}} {{< bg "14" "pg_ttl_index_14" "red" >}} {{< bg "13" "pg_ttl_index_13" "red" >}} | `pg_ttl_index_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "postgresql-18-ttl-index" "green" >}} {{< bg "17" "postgresql-17-ttl-index" "green" >}} {{< bg "16" "postgresql-16-ttl-index" "green" >}} {{< bg "15" "postgresql-15-ttl-index" "green" >}} {{< bg "14" "postgresql-14-ttl-index" "red" >}} {{< bg "13" "postgresql-13-ttl-index" "red" >}} | `postgresql-$v-ttl-index` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ttl_index_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ttl_index_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pg_ttl_index_15 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_ttl_index_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_ttl_index_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-ttl-index : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-ttl-index : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-14-ttl-index : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-ttl-index : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_18-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.8 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_18-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.1 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_18-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.2 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_18-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ttl_index_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ttl_index_18-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ttl_index_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.4 KiB | [pg_ttl_index_18-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ttl_index_18-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-ttl-index` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.2 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.2 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.3 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.2 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.1 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.2 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-ttl-index` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.0 KiB | [postgresql-18-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-18-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.8 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.1 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.2 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ttl_index_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ttl_index_17-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ttl_index_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.4 KiB | [pg_ttl_index_17-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ttl_index_17-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ttl-index` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.1 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.2 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ttl-index` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.0 KiB | [postgresql-17-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-17-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.8 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.2 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.2 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ttl_index_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.2 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ttl_index_16-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ttl_index_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.4 KiB | [pg_ttl_index_16-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ttl_index_16-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ttl-index` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.2 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.2 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.2 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.2 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.1 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.2 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ttl-index` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.0 KiB | [postgresql-16-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-16-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.7 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_15-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.0 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_15-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_15-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_15-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_ttl_index_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_ttl_index_15-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_ttl_index_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.7 KiB | [pg_ttl_index_15-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_ttl_index_15-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ttl-index` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.4 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.3 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.4 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.5 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.3 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.5 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ttl-index` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.3 KiB | [postgresql-15-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/ttl-index/postgresql-15-ttl-index_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.7 KiB | [pg_ttl_index_14-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_14-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.0 KiB | [pg_ttl_index_14-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_14-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_14-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_14-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_ttl_index_14-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_14-2.0.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_ttl_index_13` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [pg_ttl_index_13-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_ttl_index_13-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_ttl_index_13` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.9 KiB | [pg_ttl_index_13-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_ttl_index_13-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_ttl_index_13` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.5 KiB | [pg_ttl_index_13-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_ttl_index_13-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_ttl_index_13` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_ttl_index_13-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_ttl_index_13-2.0.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl" title="Repository" icon="github" subtitle="github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="postgres-extensions-pg_ttl-2.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_ttl_index;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_ttl_index;		# install via package name, for the active PG version

pig install pg_ttl_index -v 18;   # install for PG 18
pig install pg_ttl_index -v 17;   # install for PG 17
pig install pg_ttl_index -v 16;   # install for PG 16
pig install pg_ttl_index -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_ttl_index';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_ttl_index;
```
