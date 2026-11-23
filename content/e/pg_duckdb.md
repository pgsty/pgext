---
title: "pg_duckdb"
linkTitle: "pg_duckdb"
description: "DuckDB Embedded in Postgres"
weight: 2430
categories: ["OLAP"]
languages: ["C++"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_duckdb**](https://github.com/duckdb/pg_duckdb) : DuckDB Embedded in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2430** | {{< badge content="pg_duckdb" link="https://github.com/duckdb/pg_duckdb" >}} | {{< ext "pg_duckdb" >}} | `1.2.1` | {{< category "OLAP" >}} | {{< license "MIT" >}} | {{< language "C++" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `duckdb` |
|    **Need By**    | {{< ext "pg_mooncake" >}} |
|   **See Also**    | {{< ext "pg_mooncake" >}} {{< ext "pg_ducklake" >}} {{< ext "duckdb_fdw" >}} {{< ext "pg_parquet" >}} {{< ext "pg_lake" >}} {{< ext "pg_clickhouse" >}} {{< ext "aws_s3" >}} {{< ext "file_fdw" >}} {{< ext "pg_bulkload" >}} {{< ext "columnar" >}} |

> [!Note] DuckDB 2.0.0; preload required; conflicts with duckdb_fdw. pg_mooncake 0.3.0 from the PGSTY fork is compatible; the old 0.2.0 community plugin is not.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_duckdb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "19" "pg_duckdb_19" "green" >}} {{< bg "18" "pg_duckdb_18" "green" >}} {{< bg "17" "pg_duckdb_17" "green" >}} {{< bg "16" "pg_duckdb_16" "green" >}} {{< bg "15" "pg_duckdb_15" "green" >}} | `pg_duckdb_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.1` | {{< bg "19" "postgresql-19-pg-duckdb" "green" >}} {{< bg "18" "postgresql-18-pg-duckdb" "green" >}} {{< bg "17" "postgresql-17-pg-duckdb" "green" >}} {{< bg "16" "postgresql-16-pg-duckdb" "green" >}} {{< bg "15" "postgresql-15-pg-duckdb" "green" >}} | `postgresql-$v-pg-duckdb` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_duckdb_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_duckdb_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_duckdb_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_15 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_15 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_15 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.2.1" "pg_duckdb_15 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.1" "postgresql-19-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-pg-duckdb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-pg-duckdb : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_19` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.8 MiB | [pg_duckdb_19-1.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_19-1.2.1-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_19` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.3 MiB | [pg_duckdb_19-1.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_19-1.2.1-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_19` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.5 MiB | [pg_duckdb_19-1.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_19-1.2.1-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_19` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.2 MiB | [pg_duckdb_19-1.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_19-1.2.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.4 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.8 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.8 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.3 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.3 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.8 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.2 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-pg-duckdb` | `1.2.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.5 MiB | [postgresql-19-pg-duckdb_1.2.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-19-pg-duckdb_1.2.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.8 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_18-1.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_duckdb_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.3 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_18-1.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_duckdb_18` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.7 MiB | [pg_duckdb_18-1.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_18-1.2.1-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.2 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_18-1.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_18` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.4 MiB | [pg_duckdb_18-1.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_18-1.2.1-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.9 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_18-1.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_18` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.6 MiB | [pg_duckdb_18-1.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_18-1.2.1-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_18-1.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_18` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.2 MiB | [pg_duckdb_18-1.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_18-1.2.1-1PGSTY.el10.aarch64.rpm) |
| `pg_duckdb_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 MiB | [pg_duckdb_18-1.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_18-1.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.6 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.9 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.1 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.2 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.0 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-duckdb` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.3 MiB | [postgresql-18-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-18-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.7 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_17-1.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_duckdb_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.3 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_17-1.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_duckdb_17` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.8 MiB | [pg_duckdb_17-1.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_17-1.2.1-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.2 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_17-1.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_17` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.4 MiB | [pg_duckdb_17-1.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_17-1.2.1-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_17-1.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_17` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.6 MiB | [pg_duckdb_17-1.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_17-1.2.1-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_17-1.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_17` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.4 MiB | [pg_duckdb_17-1.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_17-1.2.1-1PGSTY.el10.aarch64.rpm) |
| `pg_duckdb_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.0 MiB | [pg_duckdb_17-1.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_17-1.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.6 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.6 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.6 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.1 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.0 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-duckdb` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.2 MiB | [postgresql-17-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-17-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.7 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_16-1.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_duckdb_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.2 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_16-1.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_duckdb_16` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.7 MiB | [pg_duckdb_16-1.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_16-1.2.1-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_16-1.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_16` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.4 MiB | [pg_duckdb_16-1.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_16-1.2.1-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.9 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_16-1.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_16` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.5 MiB | [pg_duckdb_16-1.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_16-1.2.1-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.3 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_16-1.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_16` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.3 MiB | [pg_duckdb_16-1.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_16-1.2.1-1PGSTY.el10.aarch64.rpm) |
| `pg_duckdb_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 MiB | [pg_duckdb_16-1.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_16-1.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.3 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.6 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.5 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.6 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.1 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.7 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.0 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-duckdb` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.2 MiB | [postgresql-16-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-16-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duckdb_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.7 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duckdb_15-1.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_duckdb_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.2 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duckdb_15-1.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_duckdb_15` | `1.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.7 MiB | [pg_duckdb_15-1.2.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_15-1.2.1-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.2 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duckdb_15-1.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_duckdb_15` | `1.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.5 MiB | [pg_duckdb_15-1.2.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_15-1.2.1-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.9 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duckdb_15-1.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_duckdb_15` | `1.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.6 MiB | [pg_duckdb_15-1.2.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_15-1.2.1-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 26.2 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_duckdb_15-1.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_duckdb_15` | `1.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.3 MiB | [pg_duckdb_15-1.2.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_15-1.2.1-1PGSTY.el10.aarch64.rpm) |
| `pg_duckdb_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.9 MiB | [pg_duckdb_15-1.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_duckdb_15-1.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 20.2 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.3 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.2 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.5 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.4 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.6 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.1 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 20.5 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 23.0 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-duckdb` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 21.2 MiB | [postgresql-15-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-duckdb/postgresql-15-pg-duckdb_1.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/duckdb/pg_duckdb" title="Repository" icon="github" subtitle="github.com/duckdb/pg_duckdb" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_duckdb-1.2.1.tar.gz curl-8.23.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_duckdb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_duckdb;		# install via package name, for the active PG version

pig install pg_duckdb -v 19;   # install for PG 19
pig install pg_duckdb -v 18;   # install for PG 18
pig install pg_duckdb -v 17;   # install for PG 17
pig install pg_duckdb -v 16;   # install for PG 16
pig install pg_duckdb -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_duckdb';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_duckdb;
```

## Usage

Sources:

- [README.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/README.md)
- [pg_duckdb.control](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/pg_duckdb.control)
- [settings.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/settings.md)
- [functions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/functions.md)
- [transactions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/transactions.md)
- [gotchas_and_syntax.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/gotchas_and_syntax.md)
- [v1.2.1](https://github.com/duckdb/pg_duckdb/releases/tag/v1.2.1)

`pg_duckdb` embeds DuckDB's analytical engine in PostgreSQL and can query ordinary PostgreSQL tables and external files. Release 1.2.1 embeds DuckDB 2.0.1; upstream lists PostgreSQL 14–18. It is not relocatable and needs cluster preload.

### Enable and Compare a Query

Add the library to the existing setting, then restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_duckdb'
```

As a superuser, install it and inspect an analytical query:

```sql
CREATE EXTENSION pg_duckdb;
CREATE TABLE orders (category text, amount numeric);
INSERT INTO orders VALUES ('a', 10), ('a', 20), ('b', 5);
SET duckdb.force_execution = true;
EXPLAIN SELECT category, sum(amount) FROM orders GROUP BY category;
SELECT category, sum(amount) FROM orders GROUP BY category;
SET duckdb.force_execution = false;
```

Ordinary queries use PostgreSQL unless forced; DuckDB-only functions or remote storage operations choose DuckDB automatically. Inspect the plan and benchmark representative data. `read_parquet`, `read_csv` and `read_json` read external data, while `duckdb.create_simple_secret` configures storage credentials. Keep secrets out of shared SQL history and verify the access boundary before granting execution.

### Privileges and Configuration

`duckdb.postgres_role` selects the role allowed to execute DuckDB, manage secrets and use MotherDuck data; by default only superusers may do so. A configured role can be created during installation. Grant membership deliberately because it exposes shared DuckDB secrets and data. Local filesystem access is restricted separately through PostgreSQL server-file roles and DuckDB filesystem settings. Community extension loading is an explicit administrator choice.

### Transaction and Compatibility Limits

A transaction may not write both PostgreSQL and DuckDB tables, or mix DDL for both engines. Keep `duckdb.unsafe_allow_mixed_transactions` disabled: bypassing the rule can commit one side only and lose consistency. Execution inside PostgreSQL functions is also disabled by default; the unsafe setting for it has known crash risks.

Version 1.2.1 supports nested fixed-length arrays and improves numeric and floating conversions. It fixes parallel-worker memory reads, cancellation cleanup and secret-bearing log/statistics handling. These fixes do not remove the need to restrict credentials and review type/query compatibility. Install matching libraries, follow preload restart requirements, and update the SQL extension in each database before relying on new behavior.
