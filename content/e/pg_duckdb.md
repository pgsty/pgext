---
title: "pg_duckdb"
linkTitle: "pg_duckdb"
description: "DuckDB Embedded in Postgres"
weight: 2430
categories: ["OLAP"]
width: full
---

DuckDB Embedded in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2430** | {{< badge content="pg_duckdb" link="https://github.com/duckdb/pg_duckdb" >}} | {{< ext "pg_duckdb" >}} | `1.0.0` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_mooncake" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_analytics" >}} {{< ext "pg_parquet" >}} {{< ext "columnar" >}} {{< ext "citus" >}} {{< ext "citus_columnar" >}} {{< ext "orioledb" >}} |

> [!Note] invalidate duckdb_fdw, broken on el8 due to c++ too low


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_duckdb" >}} | `1.0.0` | {{< bg "18" "pg_duckdb_18*" "green" >}} {{< bg "17" "pg_duckdb_17*" "green" >}} {{< bg "16" "pg_duckdb_16*" "green" >}} {{< bg "15" "pg_duckdb_15*" "green" >}} {{< bg "14" "pg_duckdb_14*" "green" >}} {{< bg "13" "pg_duckdb_13*" "red" >}} | `pg_duckdb_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_duckdb" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pg-duckdb" "green" >}} {{< bg "17" "postgresql-17-pg-duckdb" "green" >}} {{< bg "16" "postgresql-16-pg-duckdb" "green" >}} {{< bg "15" "postgresql-15-pg-duckdb" "green" >}} {{< bg "14" "postgresql-14-pg-duckdb" "green" >}} {{< bg "13" "postgresql-13-pg-duckdb" "red" >}} | `postgresql-$v-pg-duckdb` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : BREAK 1" "orange" >}}      |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : BREAK 1" "orange" >}}      |      {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : BREAK 1" "orange" >}}      |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.0.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_duckdb_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duckdb_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duckdb : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duckdb : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.1" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.1" "postgresql-14-pg-duckdb : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-duckdb : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_18` | 1.0.0 | `el8.x86_64` | pigsty | 15.4 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_18` | 1.0.0 | `el8.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_18` | 1.0.0 | `el9.x86_64` | pigsty | 15.8 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_18` | 1.0.0 | `el9.aarch64` | pigsty | 14.2 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_18` | 1.0.0 | `el10.x86_64` | pigsty | 15.0 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_18` | 1.0.0 | `el10.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_17` | 1.0.0 | `el8.x86_64` | pigsty | 15.4 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_17` | 1.0.0 | `el8.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_17` | 1.0.0 | `el9.x86_64` | pigsty | 15.8 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_17` | 0.3.1 | `el9.x86_64` | pigsty | 14.2 MiB | [pg_duckdb_17-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_17-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_17` | 1.0.0 | `el9.aarch64` | pigsty | 14.2 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_17` | 0.3.1 | `el9.aarch64` | pigsty | 12.9 MiB | [pg_duckdb_17-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_17-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_17` | 1.0.0 | `el10.x86_64` | pigsty | 15.0 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_17` | 1.0.0 | `el10.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `d12.x86_64` | pigsty | 12.3 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `d12.aarch64` | pigsty | 10.7 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `u22.x86_64` | pigsty | 14.2 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `u22.aarch64` | pigsty | 13.0 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `u24.x86_64` | pigsty | 14.1 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-duckdb` | 0.3.1 | `u24.aarch64` | pigsty | 13.0 MiB | [postgresql-17-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_16` | 1.0.0 | `el8.x86_64` | pigsty | 15.4 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_16` | 1.0.0 | `el8.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_16` | 1.0.0 | `el9.x86_64` | pigsty | 15.8 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_16` | 0.3.1 | `el9.x86_64` | pigsty | 14.2 MiB | [pg_duckdb_16-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_16-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_16` | 1.0.0 | `el9.aarch64` | pigsty | 14.2 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_16` | 0.3.1 | `el9.aarch64` | pigsty | 12.9 MiB | [pg_duckdb_16-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_16-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_16` | 1.0.0 | `el10.x86_64` | pigsty | 15.0 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_16` | 1.0.0 | `el10.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `d12.x86_64` | pigsty | 12.3 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `d12.aarch64` | pigsty | 10.7 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `u22.x86_64` | pigsty | 14.1 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `u22.aarch64` | pigsty | 13.0 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `u24.x86_64` | pigsty | 14.1 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-duckdb` | 0.3.1 | `u24.aarch64` | pigsty | 12.9 MiB | [postgresql-16-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_15` | 1.0.0 | `el8.x86_64` | pigsty | 15.4 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_15` | 1.0.0 | `el8.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_15` | 1.0.0 | `el9.x86_64` | pigsty | 15.8 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_15` | 0.3.1 | `el9.x86_64` | pigsty | 14.2 MiB | [pg_duckdb_15-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_15-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_15` | 1.0.0 | `el9.aarch64` | pigsty | 14.2 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_15` | 0.3.1 | `el9.aarch64` | pigsty | 12.9 MiB | [pg_duckdb_15-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_15-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_15` | 1.0.0 | `el10.x86_64` | pigsty | 15.1 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_15` | 1.0.0 | `el10.aarch64` | pigsty | 13.4 MiB | [pg_duckdb_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `d12.x86_64` | pigsty | 12.4 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `d12.aarch64` | pigsty | 10.7 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `u22.x86_64` | pigsty | 14.2 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `u22.aarch64` | pigsty | 13.0 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `u24.x86_64` | pigsty | 14.1 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-duckdb` | 0.3.1 | `u24.aarch64` | pigsty | 13.0 MiB | [postgresql-15-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_14` | 1.0.0 | `el8.x86_64` | pigsty | 15.4 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_duckdb_14` | 1.0.0 | `el8.aarch64` | pigsty | 13.3 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_duckdb_14` | 1.0.0 | `el9.x86_64` | pigsty | 15.8 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_14` | 0.3.1 | `el9.x86_64` | pigsty | 14.2 MiB | [pg_duckdb_14-0.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_14-0.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_duckdb_14` | 1.0.0 | `el9.aarch64` | pigsty | 14.2 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_14` | 0.3.1 | `el9.aarch64` | pigsty | 12.9 MiB | [pg_duckdb_14-0.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_14-0.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_duckdb_14` | 1.0.0 | `el10.x86_64` | pigsty | 15.1 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_duckdb_14` | 1.0.0 | `el10.aarch64` | pigsty | 13.4 MiB | [pg_duckdb_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `d12.x86_64` | pigsty | 12.4 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `d12.aarch64` | pigsty | 10.7 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `u22.x86_64` | pigsty | 14.1 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `u22.aarch64` | pigsty | 13.0 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `u24.x86_64` | pigsty | 14.1 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-duckdb` | 0.3.1 | `u24.aarch64` | pigsty | 13.0 MiB | [postgresql-14-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-14-pg-duckdb_0.3.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/duckdb/pg_duckdb" title="Repository" icon="github" subtitle="github.com/duckdb/pg_duckdb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_duckdb-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_duckdb; # get pg_duckdb source code
pig build dep pg_duckdb; # install build dependencies
pig build pkg pg_duckdb; # build extension rpm or deb
pig build ext pg_duckdb; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_duckdb; # install by extension name, for the current active PG version
pig ext install pg_duckdb; # install via package alias, for the active PG version
pig ext install pg_duckdb -v 18;   # install for PG 18
pig ext install pg_duckdb -v 17;   # install for PG 17
pig ext install pg_duckdb -v 16;   # install for PG 16
pig ext install pg_duckdb -v 15;   # install for PG 15
pig ext install pg_duckdb -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_duckdb;
```



--------

## Usage

pg_duckdb embeds DuckDB's columnar-vectorized analytics engine directly into PostgreSQL, enabling high-performance analytics within PostgreSQL while maintaining full SQL compatibility and ACID properties.

### Query Acceleration

Force analytical queries to use DuckDB's vectorized execution engine for dramatic performance improvements.

```sql
-- Enable DuckDB execution for current session
SET duckdb.force_execution = true;

-- Analytical queries automatically use DuckDB engine
SELECT 
    product_category,
    AVG(price) as avg_price,
    COUNT(*) as total_items
FROM products 
GROUP BY product_category
ORDER BY avg_price DESC;

-- Check execution plan
EXPLAIN SELECT count(*) FROM large_table;
-- Shows: Custom Scan (DuckDBScan) with DuckDB execution plan
```

### Data Lake Integration

Query cloud storage and modern data formats directly from PostgreSQL without ETL.

```sql
-- Read Parquet files from S3
SELECT customer_id, SUM(amount) as total_spent
FROM read_parquet('s3://data-lake/transactions/*.parquet')
WHERE purchase_date >= '2024-01-01'
GROUP BY customer_id;

-- Query CSV files with schema inference
SELECT * FROM read_csv('https://example.com/sales_data.csv')
WHERE region = 'North America';

-- Read JSON files with automatic flattening
SELECT event_type, COUNT(*) as event_count
FROM read_json('s3://events/user_events.json')
GROUP BY event_type;

-- Scan Apache Iceberg tables
SELECT * FROM iceberg_scan('s3://warehouse/orders/metadata.json')
WHERE order_date >= CURRENT_DATE - INTERVAL '30 days';

-- Query Delta Lake tables
SELECT product_id, SUM(quantity) as total_sales
FROM delta_scan('s3://delta-lake/sales/')
GROUP BY product_id;
```

### Hybrid Storage Tables

Create DuckDB-backed tables for analytical workloads with columnar storage and compression.

```sql
-- Create analytics table using DuckDB storage
CREATE TABLE sales_analytics (
    transaction_id BIGINT,
    product_id INT,
    customer_id INT,
    sale_date DATE,
    amount DECIMAL(10,2),
    quantity INT
) USING duckdb;

-- Insert data (stored in columnar format)
INSERT INTO sales_analytics 
SELECT id, product_id, customer_id, sale_date, price * quantity, quantity
FROM transactions
WHERE sale_date >= '2024-01-01';

-- High-performance analytical queries
SELECT 
    product_id,
    DATE_TRUNC('month', sale_date) as month,
    SUM(amount) as monthly_revenue,
    AVG(quantity) as avg_quantity
FROM sales_analytics
GROUP BY product_id, month
ORDER BY monthly_revenue DESC;
```

### Cloud Storage Security

Manage secure credentials for cloud data access.

```sql
-- Create S3 credentials
SELECT duckdb.create_simple_secret(
    type := 'S3',
    key_id := 'AKIAIOSFODNN7EXAMPLE',
    secret := 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
);

-- Create Azure credentials
SELECT duckdb.create_simple_secret(
    type := 'AZURE',
    connection_string := 'DefaultEndpointsProtocol=https;AccountName=myaccount;AccountKey=mykey'
);

-- Use credential chain for AWS
CREATE SERVER s3_server FOREIGN DATA WRAPPER duckdb_fdw 
OPTIONS (type 'S3', provider 'credential_chain');
```

### MotherDuck Cloud Integration

Connect to MotherDuck cloud analytics platform for distributed analytics.

```sql
-- Enable MotherDuck connection
CALL duckdb.enable_motherduck('your_token_here', 'production_db');

-- Create cloud-synchronized table
CREATE TABLE cloud_metrics(
    timestamp TIMESTAMP,
    metric_name TEXT,
    value DOUBLE,
    tags JSON
) USING duckdb;

-- Query across local and cloud data
SELECT 
    local.user_id,
    local.session_count,
    cloud.conversion_rate
FROM user_sessions local
JOIN ddb$shared_analytics.user_metrics cloud 
  ON local.user_id = cloud.user_id;

-- Force cloud synchronization
SELECT duckdb.force_motherduck_sync();
```

### Extension Management

Leverage DuckDB's extension ecosystem for additional capabilities.

```sql
-- Install extensions (superuser required)
SELECT duckdb.install_extension('iceberg');
SELECT duckdb.install_extension('delta'); 
SELECT duckdb.install_extension('azure');

-- View installed extensions
SELECT name, loaded, installed 
FROM duckdb.extensions();

-- Extensions auto-load when functions are used
SELECT * FROM delta_scan('azure://container/delta-table/');
```

### Direct DuckDB Queries

Execute raw DuckDB queries for advanced functionality.

```sql
-- Execute DuckDB query directly
SELECT duckdb.query('PRAGMA version');

-- Raw query with complex DuckDB syntax
SELECT duckdb.raw_query($$
    WITH monthly_sales AS (
        SELECT DATE_TRUNC('month', sale_date) as month, 
               SUM(amount) as total
        FROM sales_data 
        GROUP BY month
    )
    SELECT month, total,
           LAG(total) OVER (ORDER BY month) as prev_month,
           total - LAG(total) OVER (ORDER BY month) as growth
    FROM monthly_sales
$$);

-- Reset DuckDB instance
SELECT duckdb.recycle_ddb();
```

## Configuration

### Performance Tuning

```sql
-- Memory allocation (default: 80% of system memory)
SET duckdb.max_memory = '8GB';

-- Thread configuration
SET duckdb.threads = 8;

-- PostgreSQL scan optimization
SET duckdb.max_workers_per_postgres_scan = 4;

-- Disable specific execution for debugging
SET duckdb.force_execution = false;
```

### Security Controls

```sql
-- Restrict access to specific PostgreSQL roles
SET duckdb.postgres_role = 'analytics_team';

-- Disable filesystem access
SET duckdb.disabled_filesystems = 'LocalFileSystem';

-- Control extension installation
SET duckdb.autoinstall_known_extensions = false;
SET duckdb.allow_community_extensions = false;

-- External access control
SET duckdb.enable_external_access = true;
```

### Data Type Compatibility

| PostgreSQL Type        | DuckDB Support | Notes                           |
|------------------------|----------------|---------------------------------|
| INTEGER, BIGINT        | Full           | Direct mapping                  |
| REAL, DOUBLE PRECISION | Full           | Direct mapping                  |
| TEXT, VARCHAR          | Full           | Direct mapping                  |
| BOOLEAN                | Full           | Direct mapping                  |
| UUID                   | Full           | Direct mapping                  |
| JSON/JSONB             | Full           | JSONB converted to JSON         |
| DATE, TIMESTAMP        | Full           | Microsecond precision           |
| NUMERIC                | Limited        | May lose precision              |
| ARRAYS (1D)            | Full           | Multi-dimensional not supported |
| ENUM                   | None           | Not supported                   |

## Use Cases

- **Real-time Analytics**: Accelerate analytical queries on transactional data by 10-100x
- **Data Lake Analytics**: Query Parquet, Iceberg, and Delta Lake tables directly
- **Hybrid Workloads**: Combine OLTP applications with high-performance analytics
- **Operational Dashboards**: Build responsive dashboards on live PostgreSQL data
- **ETL Elimination**: Remove data movement overhead for analytical workloads
- **Cloud Data Integration**: Seamlessly access cloud storage from PostgreSQL

pg_duckdb transforms PostgreSQL into a powerful hybrid OLTP/OLAP database, enabling organizations to run analytical workloads at scale without sacrificing transactional consistency or requiring separate analytics infrastructure.