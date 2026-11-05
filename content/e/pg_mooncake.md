---
title: "pg_mooncake"
linkTitle: "pg_mooncake"
description: "Columnstore Table in Postgres"
weight: 2440
categories: ["OLAP"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_mooncake**](https://github.com/pgsty/pg_mooncake) : Columnstore Table in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2440** | {{< badge content="pg_mooncake" link="https://github.com/pgsty/pg_mooncake" >}} | {{< ext "pg_mooncake" >}} | `0.3.0` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pg_duckdb" >}} |
|   **See Also**    | {{< ext "pg_duckdb" >}} {{< ext "pg_ducklake" >}} {{< ext "columnar" >}} {{< ext "citus_columnar" >}} {{< ext "pg_lake" >}} {{< ext "pg_parquet" >}} {{< ext "orioledb" >}} {{< ext "storage_engine" >}} {{< ext "pg_sorted_heap" >}} {{< ext "duckdb_fdw" >}} |

> [!Note] PGSTY-maintained 0.3.0 fork for pg_duckdb 1.2.0 / DuckDB 2.0.0; bundles local Mooncake v0.0.2 and BUSL-1.1 Moonlink. Requires unsigned extensions and a matching local extension_repository. All 80 RPM/DEB PG14-18 combinations verified and locally indexed; CDN sync pending. No upstream 0.2.0 SQL upgrade script; U24 ARM64 PG14 PL/pgSQL dynamic EXPLAIN remains a known pg_duckdb error; S3/crash/stress paths unverified.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_mooncake` | `pg_duckdb` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "pg_mooncake_18" "green" >}} {{< bg "17" "pg_mooncake_17" "green" >}} {{< bg "16" "pg_mooncake_16" "green" >}} {{< bg "15" "pg_mooncake_15" "green" >}} {{< bg "14" "pg_mooncake_14" "green" >}} | `pg_mooncake_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-pg-mooncake" "green" >}} {{< bg "17" "postgresql-17-pg-mooncake" "green" >}} {{< bg "16" "postgresql-16-pg-mooncake" "green" >}} {{< bg "15" "postgresql-15-pg-mooncake" "green" >}} {{< bg "14" "postgresql-14-pg-mooncake" "green" >}} | `postgresql-$v-pg-mooncake` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "pg_mooncake_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-pg-mooncake : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-pg-mooncake : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.5 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_18-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_mooncake_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_18-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_mooncake_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_18-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_mooncake_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_18-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_mooncake_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_18-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_mooncake_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_18-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_18-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.5 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.8 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-mooncake` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 15.8 MiB | [postgresql-18-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-18-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.0 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_17-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_mooncake_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_17-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_mooncake_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_17-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_mooncake_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_17-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_mooncake_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_17-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_mooncake_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_17-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_17-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.5 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.9 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.6 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.9 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-mooncake` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 15.8 MiB | [postgresql-17-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-17-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.5 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_16-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_mooncake_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_16-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_mooncake_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_16-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_mooncake_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_16-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_mooncake_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_16-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_mooncake_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_16-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_16-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.5 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.9 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-mooncake` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 15.8 MiB | [postgresql-16-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-16-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.5 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_15-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_mooncake_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_15-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_mooncake_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_15-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_mooncake_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_15-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_mooncake_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_15-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_mooncake_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_15-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_15-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.5 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.8 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-mooncake` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 15.8 MiB | [postgresql-15-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-15-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_mooncake_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_mooncake_14-0.3.0-1PGSTY.el8.x86_64.rpm) |
| `pg_mooncake_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_mooncake_14-0.3.0-1PGSTY.el8.aarch64.rpm) |
| `pg_mooncake_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_mooncake_14-0.3.0-1PGSTY.el9.x86_64.rpm) |
| `pg_mooncake_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_mooncake_14-0.3.0-1PGSTY.el9.aarch64.rpm) |
| `pg_mooncake_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.4 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_mooncake_14-0.3.0-1PGSTY.el10.x86_64.rpm) |
| `pg_mooncake_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.7 MiB | [pg_mooncake_14-0.3.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_mooncake_14-0.3.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 15.3 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.5 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 15.4 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.6 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.6 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 15.9 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-mooncake` | `0.3.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 15.8 MiB | [postgresql-14-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-mooncake/postgresql-14-pg-mooncake_0.3.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsty/pg_mooncake" title="Repository" icon="github" subtitle="github.com/pgsty/pg_mooncake" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_mooncake-0.3.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_mooncake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_mooncake;		# install via package name, for the active PG version

pig install pg_mooncake -v 18;   # install for PG 18
pig install pg_mooncake -v 17;   # install for PG 17
pig install pg_mooncake -v 16;   # install for PG 16
pig install pg_mooncake -v 15;   # install for PG 15
pig install pg_mooncake -v 14;   # install for PG 14

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_duckdb, pg_mooncake';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_mooncake CASCADE; -- requires pg_duckdb
```

## Usage

Sources:

- [0.3.0 setup and tutorial](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/README.md)
- [0.3.0 SQL procedures](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/functions.rs)
- [Dependency guard and DuckDB extension loading](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/sql/bootstrap.sql)
- [Scan synchronization](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/src/duckdb_mooncake.rs)
- [0.3.0 validation and limitations](https://github.com/pgsty/pg_mooncake/blob/v0.3.0/docs/pg-mooncake-0.3.0-validation.md)

`pg_mooncake` 0.3.0 creates read-only columnstore mirrors of PostgreSQL tables for analytical queries. Moonlink replicates committed row changes into Arrow buffers and Parquet/Iceberg storage; DuckDB executes scans and aggregates through `pg_duckdb`. This PGSTY release targets PostgreSQL 14–18 and requires exactly `pg_duckdb` 1.2.0 with DuckDB 2.0.0 and the matching bundled DuckDB Mooncake extension.

### Configure and Enable

Have an administrator configure logical replication and add both libraries to the existing preload list, in this order. Point the repository setting at the installed local DuckDB extension. This example uses the EL PostgreSQL 17 layout:

```ini
shared_preload_libraries = 'pg_duckdb,pg_mooncake'
wal_level = logical
duckdb.allow_community_extensions = true
duckdb.allow_unsigned_extensions = true
mooncake.extension_repository = '/usr/pgsql-17/share/extension/pg_mooncake'
```

For another installation, use the directory under the matching `pg_config --sharedir`, followed by `/extension/pg_mooncake`. The tag falls back to the community repository when this setting is absent; it does not automatically select the bundled directory. Unsigned loading is required for this local binary. Ensure sufficient replication slots, WAL senders, and worker processes, then restart PostgreSQL.

```sql
CREATE EXTENSION pg_mooncake CASCADE;
```

Extension creation requires superuser privileges and checks the installed dependency version. Mirror creation opens separate database connections: source DDL must already be committed, and the connecting role must be able to authenticate, read the source, and perform the required replication/publication setup. Run the example in autocommit mode.

### Write to the Source, Query the Mirror

```sql
CREATE TABLE trades (
    id bigint PRIMARY KEY,
    symbol text,
    traded_at timestamp,
    price real
);
CALL mooncake.create_table('trades_analytics', 'trades');

INSERT INTO trades VALUES
    (1, 'AMZN', '2026-11-05 10:00:00', 207),
    (2, 'AMZN', '2026-11-05 10:05:00', 210),
    (3, 'AAPL', '2026-11-05 10:10:00', 203);
UPDATE trades SET price = 209 WHERE id = 1;
DELETE FROM trades WHERE id = 3;

SELECT symbol, avg(price) FROM trades_analytics GROUP BY symbol;
EXPLAIN SELECT avg(price) FROM trades_analytics WHERE symbol = 'AMZN';
```

The helper copies existing rows and starts logical replication. Send `INSERT`, `UPDATE`, and `DELETE` to the source table; the mirror uses `USING mooncake` and is not a writable replacement. Rolled-back source changes are not replicated. Queries can combine mirror data with regular heap tables through supported DuckDB query plans; inspect `EXPLAIN` for a DuckDB scan.

In this tag, scan and snapshot synchronization use the current backend's last-commit LSN. Do not treat this as a general guarantee that every other session's latest write has been ingested, or that a remote source shares the local WAL position. Monitor replication progress for cross-session or remote freshness requirements.

### Procedures and Maintenance

| Interface | Purpose |
| --- | --- |
| `mooncake.create_table(dst, src, src_uri, table_config)` | Create a mirror; the last two arguments default to NULL. Use the exact named argument `src_uri` for a remote PostgreSQL source. |
| `mooncake.list_tables()` | Return registered tables, commit/flush LSNs, and Iceberg warehouse locations. |
| `mooncake.create_snapshot(dst)` | Flush buffered data into an Iceberg snapshot at the requested synchronization point. |
| `mooncake.optimize_table(dst, mode)` | Request compaction; full mode compacts data and indexes. |
| `mooncake.load_files(dst, files)` | Unsupported in 0.3.0; raises SQLSTATE `0A000`. |

```sql
SELECT * FROM mooncake.list_tables();
CALL mooncake.create_snapshot('trades_analytics');
CALL mooncake.optimize_table('trades_analytics', 'full');
```

In-database reads can include buffered Arrow data, while external Iceberg readers see published snapshots. Monitor logical-slot lag and retained WAL as well as mirror storage. Dropping the destination table invokes Moonlink cleanup; verify that its registration and replication slot are removed. Preserve the source data until cleanup and any migration have been checked.

### Compatibility and Upgrade Boundaries

Mirror setup handles replica identity for source tables without a primary key, including existing leaf partitions. Partitions added later also need `REPLICA IDENTITY FULL` before receiving writes. `TRUNCATE` replication and automatic handling of later partition DDL are outside the supported/tested scope. Do not assume arbitrary schema changes or PostgreSQL-specific functions are supported by the mirror.

The v0.3.0 tag provides no in-place SQL upgrade path from 0.2.0. Use a separately validated fresh installation and rebuild mirrors from retained source tables. Replace the PostgreSQL library and matching DuckDB extension together: binaries built for DuckDB 1.4 cannot be reused with DuckDB 2.0. The tagged validation covers local storage and normal restart; S3, crash recovery, and stress behavior are not established by those results.
