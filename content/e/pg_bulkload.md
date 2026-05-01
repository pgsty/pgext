---
title: "pg_bulkload"
linkTitle: "pg_bulkload"
description: "pg_bulkload is a high speed data loading utility for PostgreSQL"
weight: 9830
categories: ["ETL"]
width: full
---

[**pg_bulkload**](https://github.com/ossc-db/pg_bulkload) : pg_bulkload is a high speed data loading utility for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9830** | {{< badge content="pg_bulkload" link="https://github.com/ossc-db/pg_bulkload" >}} | {{< ext "pg_bulkload" >}} | `3.1.23` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "file_fdw" >}} {{< ext "aws_s3" >}} {{< ext "db_migrator" >}} {{< ext "pg_fact_loader" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical" >}} |

> [!Note] pg18 fixed by vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.23` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_bulkload` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.1.23` | {{< bg "18" "pg_bulkload_18" "green" >}} {{< bg "17" "pg_bulkload_17" "green" >}} {{< bg "16" "pg_bulkload_16" "green" >}} {{< bg "15" "pg_bulkload_15" "green" >}} {{< bg "14" "pg_bulkload_14" "green" >}} | `pg_bulkload_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.23` | {{< bg "18" "postgresql-18-pg-bulkload" "green" >}} {{< bg "17" "postgresql-17-pg-bulkload" "green" >}} {{< bg "16" "postgresql-16-pg-bulkload" "green" >}} {{< bg "15" "postgresql-15-pg-bulkload" "green" >}} {{< bg "14" "postgresql-14-pg-bulkload" "green" >}} | `postgresql-$v-pg-bulkload` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 4" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.23" "pg_bulkload_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-18-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.23" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_18` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.1 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_18-3.1.23-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pgdg | 68.2 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_bulkload_18-3.1.23-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.4 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_18-3.1.23-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pgdg | 63.5 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.3 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_18-3.1.23-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_bulkload_18-3.1.23-1PGDG.rhel9.7.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.7 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_18-3.1.23-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pgdg | 59.9 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel9.7.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.2 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_18-3.1.23-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.5 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_bulkload_18-3.1.23-1PGDG.rhel10.1.x86_64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.4 KiB | [pg_bulkload_18-3.1.23-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_18-3.1.23-1PIGSTY.el10.aarch64.rpm) |
| `pg_bulkload_18` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_18-3.1.23-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_bulkload_18-3.1.23-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 238.7 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.6 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.3 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.0 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 250.3 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 244.7 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 240.8 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 237.2 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 239.9 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 235.4 KiB | [postgresql-18-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-18-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_17` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.9 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_17-3.1.23-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.0 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_17-3.1.23-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.0 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_17-3.1.23-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.4 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_17-3.1.23-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_17-3.1.23-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.2 KiB | [pg_bulkload_17-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bulkload_17-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.3 KiB | [pg_bulkload_17-3.1.23-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_17-3.1.23-1PIGSTY.el10.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.3 KiB | [pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 237.9 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 229.8 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 239.4 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.3 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 297.0 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 291.9 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 240.0 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 236.5 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 238.9 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 234.6 KiB | [postgresql-17-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_16` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.2 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_16-3.1.23-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_16-3.1.23-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.3 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_16-3.1.23-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.7 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.8 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_16-3.1.23-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.9 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.3 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_16-3.1.23-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.4 KiB | [pg_bulkload_16-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bulkload_16-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.5 KiB | [pg_bulkload_16-3.1.23-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_16-3.1.23-1PIGSTY.el10.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 239.3 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.4 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.5 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.8 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 295.4 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 289.6 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 241.3 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 236.8 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 240.2 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 234.9 KiB | [postgresql-16-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_15` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.1 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_15-3.1.23-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.7 KiB | [pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.2 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_15-3.1.23-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.1 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.6 KiB | [pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.2 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_15-3.1.23-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.6 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.6 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_15-3.1.23-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.8 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.3 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_15-3.1.23-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.4 KiB | [pg_bulkload_15-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bulkload_15-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.3 KiB | [pg_bulkload_15-3.1.23-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_15-3.1.23-1PIGSTY.el10.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.4 KiB | [pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 239.6 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.7 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.7 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.5 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 296.2 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 291.0 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 241.8 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 237.4 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 240.7 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 235.7 KiB | [postgresql-15-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_14` | `3.1.23` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.9 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_14-3.1.23-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 265.7 KiB | [pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.23` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.2 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_14-3.1.23-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.4 KiB | [pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.23` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.0 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_14-3.1.23-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.7 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.23` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_14-3.1.23-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.8 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.5 KiB | [pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.23` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.1 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_14-3.1.23-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.3 KiB | [pg_bulkload_14-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bulkload_14-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.23` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.2 KiB | [pg_bulkload_14-3.1.23-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_14-3.1.23-1PIGSTY.el10.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.2 KiB | [pg_bulkload_14-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bulkload_14-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [d12.x86_64](/os/d12.x86_64) | pigsty | 238.5 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.9 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [d13.x86_64](/os/d13.x86_64) | pigsty | 239.4 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.6 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u22.x86_64](/os/u22.x86_64) | pigsty | 295.1 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u22.aarch64](/os/u22.aarch64) | pigsty | 289.9 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u24.x86_64](/os/u24.x86_64) | pigsty | 240.1 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u24.aarch64](/os/u24.aarch64) | pigsty | 235.7 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u26.x86_64](/os/u26.x86_64) | pigsty | 239.2 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.23` | [u26.aarch64](/os/u26.aarch64) | pigsty | 234.6 KiB | [postgresql-14-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.23-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_bulkload" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_bulkload" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bulkload-VERSION3_1_23.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_bulkload;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_bulkload;		# install via package name, for the active PG version

pig install pg_bulkload -v 18;   # install for PG 18
pig install pg_bulkload -v 17;   # install for PG 17
pig install pg_bulkload -v 16;   # install for PG 16
pig install pg_bulkload -v 15;   # install for PG 15
pig install pg_bulkload -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

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

```
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
