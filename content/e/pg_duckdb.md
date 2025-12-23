---
title: "pg_duckdb"
linkTitle: "pg_duckdb"
description: "DuckDB Embedded in Postgres"
weight: 2430
categories: ["OLAP"]
width: full
---

[**pg_duckdb**](https://github.com/duckdb/pg_duckdb) : DuckDB Embedded in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2430** | {{< badge content="pg_duckdb" link="https://github.com/duckdb/pg_duckdb" >}} | {{< ext "pg_duckdb" >}} | `1.1.1` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pg_mooncake" >}} |
|   **See Also**    | {{< ext "pg_mooncake" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_analytics" >}} {{< ext "pg_parquet" >}} {{< ext "columnar" >}} {{< ext "citus" >}} {{< ext "citus_columnar" >}} {{< ext "orioledb" >}} |

> [!Note] conflict with duckdb_fdw


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_duckdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "pg_duckdb_18*" "green" >}} {{< bg "17" "pg_duckdb_17*" "green" >}} {{< bg "16" "pg_duckdb_16*" "green" >}} {{< bg "15" "pg_duckdb_15*" "green" >}} {{< bg "14" "pg_duckdb_14*" "green" >}} {{< bg "13" "pg_duckdb_13*" "red" >}} | `pg_duckdb_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-pg-duckdb" "green" >}} {{< bg "17" "postgresql-17-pg-duckdb" "green" >}} {{< bg "16" "postgresql-16-pg-duckdb" "green" >}} {{< bg "15" "postgresql-15-pg-duckdb" "green" >}} {{< bg "14" "postgresql-14-pg-duckdb" "green" >}} {{< bg "13" "postgresql-13-pg-duckdb" "red" >}} | `postgresql-$v-pg-duckdb` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_18-1.1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_18-1.1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.7 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_18-1.1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.2 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_18-1.1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.1 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_18-1.1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 MiB | [pg_duckdb_18-1.1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_18-1.1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.9 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.3 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.6 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.2 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.8 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.5 MiB | [postgresql-18-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_17-1.1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_17-1.1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.7 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_17-1.1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.2 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_17-1.1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.1 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_17-1.1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 MiB | [pg_duckdb_17-1.1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_17-1.1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.9 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.3 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.0 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.8 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.5 MiB | [postgresql-17-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_16-1.1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_16-1.1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.7 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_16-1.1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.2 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_16-1.1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.1 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_16-1.1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.4 MiB | [pg_duckdb_16-1.1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_16-1.1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.9 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.3 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 15.9 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.8 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.5 MiB | [postgresql-16-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_15-1.1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_15-1.1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_15-1.1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.2 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_15-1.1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.1 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_15-1.1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 MiB | [pg_duckdb_15-1.1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_15-1.1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.9 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.3 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.0 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.5 MiB | [postgresql-15-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.4 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_14-1.1.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.3 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_14-1.1.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.8 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_14-1.1.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.2 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_14-1.1.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.1 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_14-1.1.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.5 MiB | [pg_duckdb_14-1.1.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_14-1.1.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 13.9 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.0 MiB | [postgresql-14-pg-duckdb_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 14.3 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.2 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.0 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.6 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 15.9 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-duckdb` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 14.5 MiB | [postgresql-14-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_1.1.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/duckdb/pg_duckdb" title="Repository" icon="github" subtitle="github.com/duckdb/pg_duckdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_duckdb-1.1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_duckdb;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_duckdb;		# install via package name, for the active PG version

pig install pg_duckdb -v 18;   # install for PG 18
pig install pg_duckdb -v 17;   # install for PG 17
pig install pg_duckdb -v 16;   # install for PG 16
pig install pg_duckdb -v 15;   # install for PG 15
pig install pg_duckdb -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_duckdb';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_duckdb;
```


## Usage

[pg_duckdb docs](https://github.com/duckdb/pg_duckdb/tree/main/docs)

| Topic                                                                                                  | Description                                                |
|:-------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------|
| [**Functions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/functions.md)                       | Complete reference for all available functions             |
| [**Syntax Guide & Gotchas**](https://github.com/duckdb/pg_duckdb/blob/main/docs/gotchas_and_syntax.md) | Quick reference for common SQL patterns and things to know |
| [**Types**](https://github.com/duckdb/pg_duckdb/blob/main/docs/types.md)                               | Supported data types and type mappings                     |
| [**Extensions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/extensions.md)                     | DuckDB extension installation and usage                    |
| [**Settings**](https://github.com/duckdb/pg_duckdb/blob/main/docs/settings.md)                         | Configuration options and parameters                       |
| [**Transactions**](https://github.com/duckdb/pg_duckdb/blob/main/docs/transactions.md)                 | Transaction behavior and limitations                       |



### Quick Setup

Install pg_duckdb with pig:

```bash
pig repo set
pig install pg_duckdb
```

Edit `postgresql.conf`, then restart to take effect

```ini
shared_preload_libraries = 'pg_duckdb'
duckdb.allow_community_extensions = true
```


### Accelerate Query

You can use DuckDB to query existing PostgreSQL table without modifying them:

```sql
-- pgbench -is 1000  # init some test workloads with pgbench
CREATE EXTENSION pg_duckdb;

-- default behavior, common postgres engine
SET duckdb.force_execution = false;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;

-- now the query goes to pg_duckdb
SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
```

The result would be 8s -> 4s on 4c VM on local laptop) :

```
postgres@el9:5432/postgres=# SET duckdb.force_execution = true;
EXPLAIN ANALYZE SELECT count(*) FROM pgbench_accounts;
SET
Time: 0.206 ms
                                              QUERY PLAN
------------------------------------------------------------------------------------------------------
 Custom Scan (DuckDBScan)  (cost=0.00..0.00 rows=0 width=0) (actual time=0.001..0.001 rows=0 loops=1)
   DuckDB Execution Plan:

 ┌─────────────────────────────────────┐
 │┌───────────────────────────────────┐│
 ││    Query Profiling Information    ││
 │└───────────────────────────────────┘│
 └─────────────────────────────────────┘
 EXPLAIN ANALYZE  SELECT count(*) AS count FROM pgduckdb.public.pgbench_accounts
 ┌────────────────────────────────────────────────┐
 │┌──────────────────────────────────────────────┐│
 ││               Total Time: 3.89s              ││
 │└──────────────────────────────────────────────┘│
 └────────────────────────────────────────────────┘
 ┌───────────────────────────┐
 │           QUERY           │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │      EXPLAIN_ANALYZE      │
 │    ────────────────────   │
 │           0 rows          │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │    UNGROUPED_AGGREGATE    │
 │    ────────────────────   │
 │        Aggregates:        │
 │        count_star()       │
 │                           │
 │           1 row           │
 │          (0.00s)          │
 └─────────────┬─────────────┘
 ┌─────────────┴─────────────┐
 │         TABLE_SCAN        │
 │    ────────────────────   │
 │           Table:          │
 │      pgbench_accounts     │
 │                           │
 │      100,000,000 rows     │
 │          (3.88s)          │
 └───────────────────────────┘
```



### Data Lake

Let's play with a local minio instance:

```sql
SELECT duckdb.create_simple_secret(
    type := 'S3', key_id := 's3user_data', secret := 'S3User.Data',
    endpoint := 'https://sss.pigsty:9000', url_style := 'path' 
);
```

