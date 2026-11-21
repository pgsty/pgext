---
title: "pg_parquet"
linkTitle: "pg_parquet"
description: "copy data between Postgres and Parquet"
weight: 2480
categories: ["OLAP"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_parquet**](https://github.com/CrunchyData/pg_parquet/) : copy data between Postgres and Parquet


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2480** | {{< badge content="pg_parquet" link="https://github.com/CrunchyData/pg_parquet/" >}} | {{< ext "pg_parquet" >}} | `0.5.1` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "file_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_bulkload" >}} {{< ext "pg_lake" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_duckdb" >}} {{< ext "pg_mooncake" >}} {{< ext "pg_fact_loader" >}} {{< ext "pg_csv" >}} {{< ext "omni_csv" >}} |

> [!Note] Pigsty fixes streaming COPY temp-file leaks.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_parquet` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "19" "pg_parquet_19" "red" >}} {{< bg "18" "pg_parquet_18" "green" >}} {{< bg "17" "pg_parquet_17" "green" >}} {{< bg "16" "pg_parquet_16" "green" >}} {{< bg "15" "pg_parquet_15" "green" >}} | `pg_parquet_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.1` | {{< bg "19" "postgresql-19-pg-parquet" "red" >}} {{< bg "18" "postgresql-18-pg-parquet" "green" >}} {{< bg "17" "postgresql-17-pg-parquet" "green" >}} {{< bg "16" "postgresql-16-pg-parquet" "green" >}} {{< bg "15" "postgresql-15-pg-parquet" "green" >}} | `postgresql-$v-pg-parquet` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "pg_parquet_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_18 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_17 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_16 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 0.5.1" "pg_parquet_15 : AVAIL 4" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-parquet : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-18-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-17-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-16-pg-parquet : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.1" "postgresql-15-pg-parquet : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 MiB | [pg_parquet_18-0.5.1-4PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_18-0.5.1-4PGSTY.el8.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_18-0.5.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_parquet_18-0.5.1-2PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_18-0.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_parquet_18-0.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.8 MiB | [pg_parquet_18-0.5.1-4PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_18-0.5.1-4PGSTY.el8.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_18-0.5.1-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_parquet_18-0.5.1-2PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_18-0.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_parquet_18-0.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.7 MiB | [pg_parquet_18-0.5.1-4PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_18-0.5.1-4PGSTY.el9.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 MiB | [pg_parquet_18-0.5.1-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_parquet_18-0.5.1-3PGDG.rhel9.9.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.2 MiB | [pg_parquet_18-0.5.1-4PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_18-0.5.1-4PGSTY.el9.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 MiB | [pg_parquet_18-0.5.1-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_parquet_18-0.5.1-3PGDG.rhel9.9.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 MiB | [pg_parquet_18-0.5.1-4PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_18-0.5.1-4PGSTY.el10.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 6.7 MiB | [pg_parquet_18-0.5.1-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_parquet_18-0.5.1-3PGDG.rhel10.3.x86_64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.1 MiB | [pg_parquet_18-0.5.1-4PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_18-0.5.1-4PGSTY.el10.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.2 MiB | [pg_parquet_18-0.5.1-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_parquet_18-0.5.1-3PGDG.rhel10.3.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_18-0.5.1-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_parquet_18-0.5.1-2PGDG.rhel10.2.aarch64.rpm) |
| `pg_parquet_18` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_18-0.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_parquet_18-0.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.4 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.1 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.7 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.1 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.6 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.0 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-parquet` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.6 MiB | [postgresql-18-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-18-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 MiB | [pg_parquet_17-0.5.1-4PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_17-0.5.1-4PGSTY.el8.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_17-0.5.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_parquet_17-0.5.1-2PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_17-0.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_parquet_17-0.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.8 MiB | [pg_parquet_17-0.5.1-4PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_17-0.5.1-4PGSTY.el8.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_17-0.5.1-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_parquet_17-0.5.1-2PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_17-0.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_parquet_17-0.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.7 MiB | [pg_parquet_17-0.5.1-4PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_17-0.5.1-4PGSTY.el9.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 MiB | [pg_parquet_17-0.5.1-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_parquet_17-0.5.1-3PGDG.rhel9.9.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.2 MiB | [pg_parquet_17-0.5.1-4PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_17-0.5.1-4PGSTY.el9.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 MiB | [pg_parquet_17-0.5.1-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_parquet_17-0.5.1-3PGDG.rhel9.9.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 MiB | [pg_parquet_17-0.5.1-4PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_17-0.5.1-4PGSTY.el10.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 6.7 MiB | [pg_parquet_17-0.5.1-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_parquet_17-0.5.1-3PGDG.rhel10.3.x86_64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.1 MiB | [pg_parquet_17-0.5.1-4PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_17-0.5.1-4PGSTY.el10.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.2 MiB | [pg_parquet_17-0.5.1-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_parquet_17-0.5.1-3PGDG.rhel10.3.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_17-0.5.1-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_parquet_17-0.5.1-2PGDG.rhel10.2.aarch64.rpm) |
| `pg_parquet_17` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_17-0.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_parquet_17-0.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.4 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.1 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.7 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.1 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.6 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.0 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-parquet` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.6 MiB | [postgresql-17-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-17-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 MiB | [pg_parquet_16-0.5.1-4PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_16-0.5.1-4PGSTY.el8.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_16-0.5.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_parquet_16-0.5.1-2PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_16-0.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_parquet_16-0.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.8 MiB | [pg_parquet_16-0.5.1-4PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_16-0.5.1-4PGSTY.el8.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_16-0.5.1-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_parquet_16-0.5.1-2PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_16-0.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_parquet_16-0.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.7 MiB | [pg_parquet_16-0.5.1-4PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_16-0.5.1-4PGSTY.el9.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 MiB | [pg_parquet_16-0.5.1-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_parquet_16-0.5.1-3PGDG.rhel9.9.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.2 MiB | [pg_parquet_16-0.5.1-4PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_16-0.5.1-4PGSTY.el9.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 MiB | [pg_parquet_16-0.5.1-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_parquet_16-0.5.1-3PGDG.rhel9.9.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 MiB | [pg_parquet_16-0.5.1-4PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_16-0.5.1-4PGSTY.el10.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 6.7 MiB | [pg_parquet_16-0.5.1-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_parquet_16-0.5.1-3PGDG.rhel10.3.x86_64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.1 MiB | [pg_parquet_16-0.5.1-4PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_16-0.5.1-4PGSTY.el10.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.2 MiB | [pg_parquet_16-0.5.1-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_parquet_16-0.5.1-3PGDG.rhel10.3.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_16-0.5.1-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_parquet_16-0.5.1-2PGDG.rhel10.2.aarch64.rpm) |
| `pg_parquet_16` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_16-0.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_parquet_16-0.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.1 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.6 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.1 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.7 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.0 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-parquet` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.6 MiB | [postgresql-16-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-16-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_parquet_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.8 MiB | [pg_parquet_15-0.5.1-4PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_parquet_15-0.5.1-4PGSTY.el8.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_15-0.5.1-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_parquet_15-0.5.1-2PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.1 MiB | [pg_parquet_15-0.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_parquet_15-0.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.8 MiB | [pg_parquet_15-0.5.1-4PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_parquet_15-0.5.1-4PGSTY.el8.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_15-0.5.1-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_parquet_15-0.5.1-2PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.2 MiB | [pg_parquet_15-0.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_parquet_15-0.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.7 MiB | [pg_parquet_15-0.5.1-4PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_parquet_15-0.5.1-4PGSTY.el9.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 MiB | [pg_parquet_15-0.5.1-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_parquet_15-0.5.1-3PGDG.rhel9.9.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.2 MiB | [pg_parquet_15-0.5.1-4PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_parquet_15-0.5.1-4PGSTY.el9.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 MiB | [pg_parquet_15-0.5.1-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_parquet_15-0.5.1-3PGDG.rhel9.9.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 MiB | [pg_parquet_15-0.5.1-4PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_parquet_15-0.5.1-4PGSTY.el10.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 6.7 MiB | [pg_parquet_15-0.5.1-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_parquet_15-0.5.1-3PGDG.rhel10.3.x86_64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.1 MiB | [pg_parquet_15-0.5.1-4PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_parquet_15-0.5.1-4PGSTY.el10.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.2 MiB | [pg_parquet_15-0.5.1-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_parquet_15-0.5.1-3PGDG.rhel10.3.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_15-0.5.1-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_parquet_15-0.5.1-2PGDG.rhel10.2.aarch64.rpm) |
| `pg_parquet_15` | `0.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 6.4 MiB | [pg_parquet_15-0.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_parquet_15-0.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.3 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 7.3 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.3 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 7.3 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.1 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 8.6 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.1 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 8.6 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.0 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-parquet` | `0.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 8.6 MiB | [postgresql-15-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-parquet/postgresql-15-pg-parquet_0.5.1-4PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pg_parquet/" title="Repository" icon="github" subtitle="github.com/CrunchyData/pg_parquet/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_parquet-0.5.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_parquet;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_parquet;		# install via package name, for the active PG version

pig install pg_parquet -v 18;   # install for PG 18
pig install pg_parquet -v 17;   # install for PG 17
pig install pg_parquet -v 16;   # install for PG 16
pig install pg_parquet -v 15;   # install for PG 15

```


[**Config**](https://pigsty.io/docs/pgsql/ext/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_parquet';
```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_parquet;
```

## Usage

Sources:

- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/README.md)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/pg_parquet.control)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/Cargo.toml)
- [Official source](https://github.com/CrunchyData/pg_parquet/releases/tag/v0.6.0)

`pg_parquet` 0.6.0 reads and writes Parquet through PostgreSQL COPY and inspects file metadata. The source targets PostgreSQL 14–19. It uses server-side local files or S3, Azure, GCS and HTTP(S) endpoints; these locations and credentials are separate from a PostgreSQL backup.

### Enable and Copy

Preserve other preload entries, add the library and restart. A superuser creates the non-relocatable extension. The server process must have access to the chosen local path or remote object. Run the example as an administrator with the relevant COPY privileges:

```ini
shared_preload_libraries = 'pg_parquet'
```

```sql
CREATE EXTENSION pg_parquet;
CREATE TABLE parquet_demo (id bigint, label text);
INSERT INTO parquet_demo VALUES (1, 'first'), (2, 'second');
COPY parquet_demo TO '/tmp/parquet_demo.parquet' WITH (format parquet);
CREATE TABLE parquet_copy (LIKE parquet_demo);
COPY parquet_copy FROM '/tmp/parquet_demo.parquet' WITH (format parquet, match_by name);
SELECT * FROM parquet.schema('/tmp/parquet_demo.parquet');
SELECT * FROM parquet.file_metadata('/tmp/parquet_demo.parquet');
```

Use COPY of a query for a filtered export. STDIN, STDOUT and PROGRAM streams require an explicit Parquet format. External output is not rolled back with the SQL transaction: failed or rolled-back exports can leave files or incomplete uploads.

### Schema and Type Handling

```sql
CREATE TABLE inferred_demo () WITH (load_from = '/tmp/parquet_demo.parquet');
SET pg_parquet.max_parallel_workers = 2;
SELECT * FROM pg_stat_progress_copy;
```

The `load_from` table option infers columns and loads rows. Composite types are created in `parquet_structs`; the creator needs CREATE there, which installation grants to PUBLIC. Generated types survive dropping the table. Dropping the extension with CASCADE can remove dependent columns. Maps require `crunchy_map`; unsupported inferred types need an explicitly defined target table.

`match_by name` matches nested composite fields by name; the default is position. Version 0.6.0 improves dictionary/decimal/INT96 and infinite date/timestamp reads, preserves JSON text and rejects embedded NUL strings. COPY TO rejects multidimensional arrays. PostGIS geometry/geography use Parquet geospatial types; curves and other unsupported geometries must be converted or exported explicitly as binary.

### Files, Performance and Credentials

`parquet.metadata`, `parquet.kv_metadata`, `parquet.column_stats` and `parquet.schema` inspect files. The `bloom_filter` and `dictionary` export options select columns; COPY FROM reads all rows and does not use bloom filters to filter them.

`pg_parquet.max_parallel_workers` defaults to zero. Eligible copies can use parallel workers; query exports, temporary tables, triggers/defaults and other cases can keep the serial path. Import workers decode while the leader inserts. Inspect `pg_stat_progress_copy` and size resources for the actual file and table.

`file_size_bytes` writes a folder of files. Nonempty folders are refused unless `overwrite true` is specified, except that a denied listing cannot establish emptiness and retains the older write behavior. Overwrite can remove files from an earlier copy. Concurrent writers are excluded within one PostgreSQL instance; other instances require external coordination. A single-file destination retains its replacement behavior. Plan cleanup for failed exports.

`pg_parquet.profile_name` selects an object-store identity only if the role is allowed by `parquet.user_profile_mapping`. Only a superuser or `parquet_profile_manager` may manage that mapping; privileges are rechecked for each copy. Keep credentials in protected server-side files/environment. `pg_parquet.allow_credentials_in_parameters` defaults off because SQL credentials can appear in logs and activity views. HTTP exports buffer one file in memory and cap it at 10 MiB; object-store parallel multipart requirements vary, including limitations for stores demanding equally sized parts.

### Upgrade

Deploy files matching the server major, follow the preload maintenance procedure and use the supplied SQL update path. Check existing folder-export jobs, inferred composite-type dependencies, credential mappings and type coercion behavior before moving to 0.6.0.
