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
| **9830** | {{< badge content="pg_bulkload" link="https://github.com/ossc-db/pg_bulkload" >}} | {{< ext "pg_bulkload" >}} | `3.1.22` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "file_fdw" >}} {{< ext "aws_s3" >}} {{< ext "db_migrator" >}} {{< ext "pg_fact_loader" >}} {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "pglogical" >}} |

> [!Note] pg18 breaks


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.22` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_bulkload` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.1.22` | {{< bg "18" "pg_bulkload_18*" "red" >}} {{< bg "17" "pg_bulkload_17*" "green" >}} {{< bg "16" "pg_bulkload_16*" "green" >}} {{< bg "15" "pg_bulkload_15*" "green" >}} {{< bg "14" "pg_bulkload_14*" "green" >}} {{< bg "13" "pg_bulkload_13*" "green" >}} | `pg_bulkload_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.22` | {{< bg "18" "postgresql-18-pg-bulkload" "red" >}} {{< bg "17" "postgresql-17-pg-bulkload" "green" >}} {{< bg "16" "postgresql-16-pg-bulkload" "green" >}} {{< bg "15" "postgresql-15-pg-bulkload" "green" >}} {{< bg "14" "postgresql-14-pg-bulkload" "green" >}} {{< bg "13" "postgresql-13-pg-bulkload" "green" >}} | `postgresql-$v-pg-bulkload` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_13 : AVAIL 6" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_13 : AVAIL 4" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "pg_bulkload_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_14 : AVAIL 4" "green" >}} | {{< bg "PIGSTY 3.1.22" "pg_bulkload_13 : AVAIL 4" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PGDG 3.1.22" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_bulkload_18 : MISS 0" "red" >}}      | {{< bg "PGDG 3.1.22" "pg_bulkload_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.1.22" "pg_bulkload_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-bulkload : MISS 0" "red" >}}      | {{< bg "PIGSTY 3.1.22" "postgresql-17-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-16-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-15-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-14-pg-bulkload : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.22" "postgresql-13-pg-bulkload : AVAIL 1" "green" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_17` | `3.1.22` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.8 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.0 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.0 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.4 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.2 KiB | [pg_bulkload_17-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_bulkload_17-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.0 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_17-3.1.22-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.3 KiB | [pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_bulkload_17-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `pg_bulkload_17` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.1 KiB | [pg_bulkload_17-3.1.22-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_17-3.1.22-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [d12.x86_64](/os/d12.x86_64) | pigsty | 238.4 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [d12.aarch64](/os/d12.aarch64) | pigsty | 229.8 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [d13.x86_64](/os/d13.x86_64) | pigsty | 239.6 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.3 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [u22.x86_64](/os/u22.x86_64) | pigsty | 297.0 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [u22.aarch64](/os/u22.aarch64) | pigsty | 291.9 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [u24.x86_64](/os/u24.x86_64) | pigsty | 240.0 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-bulkload` | `3.1.22` | [u24.aarch64](/os/u24.aarch64) | pigsty | 236.5 KiB | [postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-17-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_16` | `3.1.22` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.2 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_16-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.5 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_16-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.4 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.3 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_16-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.7 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_bulkload_16-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.7 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_16-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.9 KiB | [pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_bulkload_16-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.4 KiB | [pg_bulkload_16-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_bulkload_16-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.2 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_16-3.1.22-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.6 KiB | [pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_bulkload_16-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `pg_bulkload_16` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.5 KiB | [pg_bulkload_16-3.1.22-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_16-3.1.22-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [d12.x86_64](/os/d12.x86_64) | pigsty | 239.1 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.2 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.5 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [d13.aarch64](/os/d13.aarch64) | pigsty | 231.8 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [u22.x86_64](/os/u22.x86_64) | pigsty | 295.4 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [u22.aarch64](/os/u22.aarch64) | pigsty | 289.6 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [u24.x86_64](/os/u24.x86_64) | pigsty | 241.3 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-bulkload` | `3.1.22` | [u24.aarch64](/os/u24.aarch64) | pigsty | 237.0 KiB | [postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-16-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_15` | `3.1.22` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.0 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_15-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.2 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.7 KiB | [pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_bulkload_15-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.2 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_15-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.1 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.6 KiB | [pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_bulkload_15-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.1 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_15-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.6 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_bulkload_15-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.6 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_15-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.8 KiB | [pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_bulkload_15-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.4 KiB | [pg_bulkload_15-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_bulkload_15-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.2 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_15-3.1.22-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.4 KiB | [pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_bulkload_15-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `pg_bulkload_15` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.3 KiB | [pg_bulkload_15-3.1.22-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_15-3.1.22-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [d12.x86_64](/os/d12.x86_64) | pigsty | 239.2 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.9 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [d13.x86_64](/os/d13.x86_64) | pigsty | 240.8 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [d13.aarch64](/os/d13.aarch64) | pigsty | 232.5 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [u22.x86_64](/os/u22.x86_64) | pigsty | 296.2 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [u22.aarch64](/os/u22.aarch64) | pigsty | 290.7 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [u24.x86_64](/os/u24.x86_64) | pigsty | 241.8 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-bulkload` | `3.1.22` | [u24.aarch64](/os/u24.aarch64) | pigsty | 237.4 KiB | [postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-15-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_14` | `3.1.22` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.9 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_14-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el8.x86_64](/os/el8.x86_64) | pgdg | 265.7 KiB | [pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_bulkload_14-3.1.19-1.rhel8.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.2 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_14-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.1 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.6 KiB | [pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.4 KiB | [pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_bulkload_14-3.1.19-2.rhel8.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el9.x86_64](/os/el9.x86_64) | pigsty | 61.0 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_14-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.7 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_bulkload_14-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_14-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.8 KiB | [pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.5 KiB | [pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.6 KiB | [pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_bulkload_14-3.1.19-2.rhel9.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 63.3 KiB | [pg_bulkload_14-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_bulkload_14-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.2 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_14-3.1.22-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.2 KiB | [pg_bulkload_14-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_bulkload_14-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `pg_bulkload_14` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.1 KiB | [pg_bulkload_14-3.1.22-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_14-3.1.22-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [d12.x86_64](/os/d12.x86_64) | pigsty | 238.4 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [d12.aarch64](/os/d12.aarch64) | pigsty | 230.3 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [d13.x86_64](/os/d13.x86_64) | pigsty | 239.3 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [d13.aarch64](/os/d13.aarch64) | pigsty | 230.9 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [u22.x86_64](/os/u22.x86_64) | pigsty | 295.0 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [u22.aarch64](/os/u22.aarch64) | pigsty | 290.0 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [u24.x86_64](/os/u24.x86_64) | pigsty | 240.1 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-bulkload` | `3.1.22` | [u24.aarch64](/os/u24.aarch64) | pigsty | 235.7 KiB | [postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-14-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_bulkload_13` | `3.1.22` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.0 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_bulkload_13-3.1.22-1PIGSTY.el8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.21` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.5 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.21-1PGDG.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.20` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.9 KiB | [pg_bulkload_13-3.1.20-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.20-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.4 KiB | [pg_bulkload_13-3.1.18-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.18-2.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.18` | [el8.x86_64](/os/el8.x86_64) | pgdg | 260.3 KiB | [pg_bulkload_13-3.1.18-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.18-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.17` | [el8.x86_64](/os/el8.x86_64) | pgdg | 522.1 KiB | [pg_bulkload_13-3.1.17-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_bulkload_13-3.1.17-1.rhel8.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el8.aarch64](/os/el8.aarch64) | pigsty | 59.8 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_bulkload_13-3.1.22-1PIGSTY.el8.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.21` | [el8.aarch64](/os/el8.aarch64) | pgdg | 70.7 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.21-1PGDG.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.20` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.2 KiB | [pg_bulkload_13-3.1.20-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.20-1.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.19` | [el8.aarch64](/os/el8.aarch64) | pgdg | 71.0 KiB | [pg_bulkload_13-3.1.19-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_bulkload_13-3.1.19-2.rhel8.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el9.x86_64](/os/el9.x86_64) | pigsty | 60.8 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_bulkload_13-3.1.22-1PIGSTY.el9.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.21` | [el9.x86_64](/os/el9.x86_64) | pgdg | 69.3 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_bulkload_13-3.1.21-1PGDG.rhel9.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el9.aarch64](/os/el9.aarch64) | pigsty | 58.5 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_bulkload_13-3.1.22-1PIGSTY.el9.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.21` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.5 KiB | [pg_bulkload_13-3.1.21-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.21-1PGDG.rhel9.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.20` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.3 KiB | [pg_bulkload_13-3.1.20-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.20-1.rhel9.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.19` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.4 KiB | [pg_bulkload_13-3.1.19-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_bulkload_13-3.1.19-2.rhel9.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pgdg | 62.8 KiB | [pg_bulkload_13-3.1.22-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_bulkload_13-3.1.22-2PGDG.rhel10.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el10.x86_64](/os/el10.x86_64) | pigsty | 61.6 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_bulkload_13-3.1.22-1PIGSTY.el10.x86_64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pgdg | 60.2 KiB | [pg_bulkload_13-3.1.22-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_bulkload_13-3.1.22-2PGDG.rhel10.aarch64.rpm) |
| `pg_bulkload_13` | `3.1.22` | [el10.aarch64](/os/el10.aarch64) | pigsty | 59.0 KiB | [pg_bulkload_13-3.1.22-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_bulkload_13-3.1.22-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [d12.x86_64](/os/d12.x86_64) | pigsty | 236.6 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [d12.aarch64](/os/d12.aarch64) | pigsty | 227.8 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [d13.x86_64](/os/d13.x86_64) | pigsty | 237.5 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [d13.aarch64](/os/d13.aarch64) | pigsty | 228.8 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [u22.x86_64](/os/u22.x86_64) | pigsty | 289.5 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [u22.aarch64](/os/u22.aarch64) | pigsty | 284.1 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [u24.x86_64](/os/u24.x86_64) | pigsty | 237.8 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-bulkload` | `3.1.22` | [u24.aarch64](/os/u24.aarch64) | pigsty | 234.0 KiB | [postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-bulkload/postgresql-13-pg-bulkload_3.1.22-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ossc-db/pg_bulkload" title="Repository" icon="github" subtitle="github.com/ossc-db/pg_bulkload" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_bulkload-VERSION3_1_22.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_bulkload;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_bulkload;		# install via package name, for the active PG version

pig install pg_bulkload -v 17;   # install for PG 17
pig install pg_bulkload -v 16;   # install for PG 16
pig install pg_bulkload -v 15;   # install for PG 15
pig install pg_bulkload -v 14;   # install for PG 14
pig install pg_bulkload -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_bulkload;
```
