---
title: "pg_mooncake"
linkTitle: "pg_mooncake"
description: "Columnstore Table in Postgres"
weight: 2440
categories: ["OLAP"]
width: full
---

[**pg_mooncake**](https://github.com/Mooncake-Labs/pg_mooncake) : Columnstore Table in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2440** | {{< badge content="pg_mooncake" link="https://github.com/Mooncake-Labs/pg_mooncake" >}} | {{< ext "pg_mooncake" >}} | `0.2.0` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---Ld--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_duckdb" >}} |
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_analytics" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "pg_parquet" >}} {{< ext "orioledb" >}} {{< ext "timescaledb" >}} |

> [!Note] unpublished release


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_mooncake` | `pg_duckdb` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pg_mooncake_18" "green" >}} {{< bg "17" "pg_mooncake_17" "green" >}} {{< bg "16" "pg_mooncake_16" "green" >}} {{< bg "15" "pg_mooncake_15" "green" >}} {{< bg "14" "pg_mooncake_14" "green" >}} {{< bg "13" "pg_mooncake_13" "red" >}} | `pg_mooncake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pg-mooncake" "green" >}} {{< bg "17" "postgresql-17-pg-mooncake" "green" >}} {{< bg "16" "postgresql-16-pg-mooncake" "green" >}} {{< bg "15" "postgresql-15-pg-mooncake" "green" >}} {{< bg "14" "postgresql-14-pg-mooncake" "green" >}} {{< bg "13" "postgresql-13-pg-mooncake" "red" >}} | `postgresql-$v-pg-mooncake` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_mooncake_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-mooncake : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.1 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.2 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_mooncake_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 MiB | [pg_mooncake_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.7 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.7 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.6 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.1 MiB | [postgresql-18-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.1 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.2 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_mooncake_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 MiB | [pg_mooncake_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.7 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.8 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.6 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.1 MiB | [postgresql-17-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.1 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.2 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_16-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_mooncake_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 MiB | [pg_mooncake_16-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_16-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.7 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.8 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.6 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.1 MiB | [postgresql-16-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.1 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.2 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_15-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_mooncake_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 MiB | [pg_mooncake_15-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_15-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.7 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.8 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.7 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.1 MiB | [postgresql-15-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.1 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_mooncake_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.8 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_mooncake_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.5 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_mooncake_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.2 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_mooncake_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_14-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_mooncake_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.0 MiB | [pg_mooncake_14-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_14-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.8 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.7 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.7 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.7 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.1 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.6 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.1 MiB | [postgresql-14-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Mooncake-Labs/pg_mooncake" title="Repository" icon="github" subtitle="github.com/Mooncake-Labs/pg_mooncake" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mooncake-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_mooncake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_mooncake;		# install via package name, for the active PG version

pig install pg_mooncake -v 18;   # install for PG 18
pig install pg_mooncake -v 17;   # install for PG 17
pig install pg_mooncake -v 16;   # install for PG 16
pig install pg_mooncake -v 15;   # install for PG 15
pig install pg_mooncake -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_duckdb, pg_mooncake';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_mooncake CASCADE; -- requires pg_duckdb
```


## Usage

[`pg_mooncake`](https://github.com/Mooncake-Labs/pg_mooncake) 0.2.0 (unpublished yet) is rewritten in Rust and designed as a sub-extension of `pg_duckdb`.

pg_mooncake docs: https://docs.mooncake.dev/


### Quick Setup

Install pg_duckdb and pg_mooncake with pig:

```bash
pig repo set
pig install pg_duckdb pg_mooncake
```

Edit postgresql.conf, then restart to take effect

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
duckdb.allow_community_extensions = true
wal_level = logical
```



### Hello Worlds

- [Tutorial](https://docs.mooncake.dev/pg/get-started/Hello-world)

```sql
-- create the extension alone with pg_duckdb
CREATE EXTENSION pg_mooncake CASCADE;

-- Next, create a regular Postgres table trades:
CREATE TABLE trades(
  id bigint PRIMARY KEY,
  symbol text,
  time timestamp,
  price real
);

-- Then, create a columnstore mirror trades_iceberg that stays in sync with trades:
CALL mooncake.create_table('trades_iceberg', 'trades');

-- Now, insert some data into trades:
INSERT INTO trades VALUES
    (1,  'AMD', '2024-06-05 10:00:00', 119),
    (2, 'AMZN', '2024-06-05 10:05:00', 207),
    (3, 'AAPL', '2024-06-05 10:10:00', 203),
    (4, 'AMZN', '2024-06-05 10:15:00', 210);

-- Finally, query it with duckdb
EXPLAIN
    SELECT avg(price) FROM trades_iceberg WHERE symbol = 'AMZN';
```

You will see the DuckDBScan in the execution plan:

```bash
                         QUERY PLAN
------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0)
   DuckDB Execution Plan:

 ┌───────────────────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │    Aggregates: avg(#0)    │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         PROJECTION        │
 │    ────────────────────   │
 │   CAST(price AS DOUBLE)   │
 │                           │
 │          ~0 rows          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │       MOONCAKE_SCAN       │
 │    ────────────────────   │
 │   Table: trades_iceberg   │
 │     Projections: price    │
 │                           │
 │          Filters:         │
 │       symbol='AMZN'       │
 │                           │
 │          ~0 rows          │
 └───────────────────────────┘
```
