---
title: "pg_bulkload"
linkTitle: "pg_bulkload"
description: "pg_bulkload is a high speed data loading utility for PostgreSQL"
weight: 9830
categories: ["ETL"]
languages: ["C"]
licenses: ["BSD-3-Clause"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_bulkload**](https://github.com/ossc-db/pg_bulkload) : pg_bulkload is a high speed data loading utility for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9830** | {{< badge content="pg_bulkload" link="https://github.com/ossc-db/pg_bulkload" >}} | {{< ext "pg_bulkload" >}} | `3.1.23` | {{< category "ETL" >}} | {{< license "BSD-3-Clause" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "file_fdw" >}} {{< ext "aws_s3" >}} {{< ext "pg_parquet" >}} {{< ext "pg_lake" >}} {{< ext "pg_ducklake" >}} {{< ext "pg_fact_loader" >}} {{< ext "pg_csv" >}} {{< ext "omni_csv" >}} {{< ext "pg_duckdb" >}} {{< ext "duckdb_fdw" >}} |

> [!Note] pg18 fixed by vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.23` | {{< bg "19" "" "red" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_bulkload` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.23` | {{< bg "19" "pg_bulkload_19" "red" >}} {{< bg "18" "pg_bulkload_18" "green" >}} {{< bg "17" "pg_bulkload_17" "green" >}} {{< bg "16" "pg_bulkload_16" "green" >}} {{< bg "15" "pg_bulkload_15" "green" >}} | `pg_bulkload_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.23` | {{< bg "19" "postgresql-19-pg-bulkload" "red" >}} {{< bg "18" "postgresql-18-pg-bulkload" "green" >}} {{< bg "17" "postgresql-17-pg-bulkload" "green" >}} {{< bg "16" "postgresql-16-pg-bulkload" "green" >}} {{< bg "15" "postgresql-15-pg-bulkload" "green" >}} | `postgresql-$v-pg-bulkload` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "pg_bulkload_19 : N/A 0" "gray" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.1.23" "pg_bulkload_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-19-pg-bulkload : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_18` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.4 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bulkload_18-3.1.23-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 210.6 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_18-3.1.23-1PGSTY.el8.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bulkload_18-3.1.23-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.6 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bulkload_18-3.1.23-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 203.5 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_18-3.1.23-1PGSTY.el8.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.5 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.3 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bulkload_18-3.1.23-6PGDG.rhel9.9.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 207.6 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_18-3.1.23-1PGSTY.el9.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.7 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bulkload_18-3.1.23-6PGDG.rhel9.9.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 202.7 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_18-3.1.23-1PGSTY.el9.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.2 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bulkload_18-3.1.23-6PGDG.rhel10.3.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 208.4 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_18-3.1.23-1PGSTY.el10.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-6PGDG.rhel10.3.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_bulkload_18-3.1.23-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.6 KiB | [pg_bulkload_18-3.1.23-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_18-3.1.23-1PGSTY.el10.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel10.1.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 228.4 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 220.3 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 230.0 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 222.4 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 240.1 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 234.7 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 230.7 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 227.5 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 229.7 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 225.6 KiB | [postgresql-18-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_17` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.1 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_bulkload_17-3.1.23-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 209.7 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_17-3.1.23-1PGSTY.el8.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.3 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_bulkload_17-3.1.23-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 202.5 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_17-3.1.23-1PGSTY.el8.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.0 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_bulkload_17-3.1.23-6PGDG.rhel9.9.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 206.9 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_17-3.1.23-1PGSTY.el9.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.4 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_bulkload_17-3.1.23-6PGDG.rhel9.9.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 201.7 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_17-3.1.23-1PGSTY.el9.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.0 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bulkload_17-3.1.23-6PGDG.rhel10.3.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 207.3 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_17-3.1.23-1PGSTY.el10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.2 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.23-6PGDG.rhel10.3.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_bulkload_17-3.1.23-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.23-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 202.8 KiB | [pg_bulkload_17-3.1.23-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_17-3.1.23-1PGSTY.el10.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.5 KiB | [pg_bulkload_17-3.1.23-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.23-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.3 KiB | [pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 227.7 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 219.7 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 229.2 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 221.0 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 286.9 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 281.9 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 230.1 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 227.0 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 229.2 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 225.0 KiB | [postgresql-17-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_16` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.5 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.23-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 211.0 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_16-3.1.23-1PGSTY.el8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.7 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.23-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 203.3 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_16-3.1.23-1PGSTY.el8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.3 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bulkload_16-3.1.23-6PGDG.rhel9.9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 207.3 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_16-3.1.23-1PGSTY.el9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.7 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bulkload_16-3.1.23-6PGDG.rhel9.9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 202.4 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_16-3.1.23-1PGSTY.el9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.3 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bulkload_16-3.1.23-6PGDG.rhel10.3.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 208.6 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_16-3.1.23-1PGSTY.el10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.5 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.23-6PGDG.rhel10.3.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.6 KiB | [pg_bulkload_16-3.1.23-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.23-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.6 KiB | [pg_bulkload_16-3.1.23-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_16-3.1.23-1PGSTY.el10.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.8 KiB | [pg_bulkload_16-3.1.23-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.23-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 229.1 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 220.3 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 230.4 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 221.8 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 285.7 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 279.9 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 231.5 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 227.7 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 230.3 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 225.4 KiB | [postgresql-16-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_15` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 69.3 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.23-6PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 211.2 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_15-3.1.23-1PGSTY.el8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.7 KiB | [pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 64.5 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.23-6PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 203.4 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_15-3.1.23-1PGSTY.el8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.1 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.6 KiB | [pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.2 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bulkload_15-3.1.23-6PGDG.rhel9.9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 207.6 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_15-3.1.23-1PGSTY.el9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pgdg | 60.7 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.23-6PGDG.rhel9.9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 202.8 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_15-3.1.23-1PGSTY.el9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.2 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bulkload_15-3.1.23-6PGDG.rhel10.3.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 209.1 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_15-3.1.23-1PGSTY.el10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.4 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.23-6PGDG.rhel10.3.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 61.3 KiB | [pg_bulkload_15-3.1.23-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.23-6PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 203.9 KiB | [pg_bulkload_15-3.1.23-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_15-3.1.23-1PGSTY.el10.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.7 KiB | [pg_bulkload_15-3.1.23-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.23-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.4 KiB | [pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 229.6 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 221.1 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 231.1 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 222.6 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 286.5 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 281.1 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 231.9 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 228.1 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 230.8 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 226.0 KiB | [postgresql-15-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_bulkload" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_bulkload" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bulkload-VERSION3_1_23.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_bulkload;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_bulkload;		# install via package name, for the active PG version

pig install pg_bulkload -v 18;   # install for PG 18
pig install pg_bulkload -v 17;   # install for PG 17
pig install pg_bulkload -v 16;   # install for PG 16
pig install pg_bulkload -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_bulkload;
```




## Usage

> [pg_bulkload: pg_bulkload is a high speed data loading utility for PostgreSQL](https://github.com/ossc-db/pg_bulkload)

A high-speed data loading tool for PostgreSQL that bypasses shared buffers for massive data loads, with built-in ETL features for input validation and data transformation.

### Basic Usage

Load data using a control file:

```bash
pg_bulkload sample_csv.ctl
```

Output:

```text
NOTICE: BULK LOAD START
NOTICE: BULK LOAD END
    0 Rows skipped.
    8 Rows successfully loaded.
    0 Rows not loaded due to parse errors.
    0 Rows not loaded due to duplicate errors.
    0 Rows replaced with new rows.
```

### Control File Example

```ini
# sample_csv.ctl
OUTPUT = my_table
INPUT = /path/to/data.csv
TYPE = CSV
DELIMITER = ,
QUOTE = "\""
ESCAPE = "\""
NULL = ""
SKIP = 1              # skip header row
PARSE_ERRORS = 100    # allow up to 100 parse errors
DUPLICATE_ERRORS = 0  # reject on duplicate key errors
ON_DUPLICATE_KEEP = NEW  # or OLD
TRUNCATE = NO
```

### Loading Modes

- **DIRECT**: Bypasses shared buffers, writes directly to data files (fastest)
- **PARALLEL**: Uses multiple processes for loading
- **CSV/BINARY/FIXED**: Supports various input formats

### SQL Interface

```sql
-- Load data from within SQL
SELECT pg_bulkload(
    'OUTPUT = my_table, INPUT = /path/to/data.csv, TYPE = CSV'
);
```

### Key Features

- Bypasses PostgreSQL shared buffers for maximum throughput
- Input data validation with configurable error thresholds
- Duplicate key handling (keep new, keep old, or reject)
- CSV, fixed-length, and binary input formats
- Skip rows, filter functions for data transformation
- Parallel loading support

### Documentation

Full documentation: http://ossc-db.github.io/pg_bulkload/index.html
