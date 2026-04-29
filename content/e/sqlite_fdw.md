---
title: "sqlite_fdw"
linkTitle: "sqlite_fdw"
description: "SQLite Foreign Data Wrapper"
weight: 8640
categories: ["FDW"]
width: full
---

[**sqlite_fdw**](https://github.com/pgspider/sqlite_fdw) : SQLite Foreign Data Wrapper


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8640** | {{< badge content="sqlite_fdw" link="https://github.com/pgspider/sqlite_fdw" >}} | {{< ext "sqlite_fdw" >}} | `2.5.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "file_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "duckdb_fdw" >}} |

> [!Note] break on el8 due to sqlite-lib version low


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.5.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `sqlite_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.5.0` | {{< bg "18" "sqlite_fdw_18" "red" >}} {{< bg "17" "sqlite_fdw_17" "green" >}} {{< bg "16" "sqlite_fdw_16" "green" >}} {{< bg "15" "sqlite_fdw_15" "green" >}} {{< bg "14" "sqlite_fdw_14" "green" >}} | `sqlite_fdw_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.5.0` | {{< bg "18" "postgresql-18-sqlite-fdw" "red" >}} {{< bg "17" "postgresql-17-sqlite-fdw" "green" >}} {{< bg "16" "postgresql-16-sqlite-fdw" "green" >}} {{< bg "15" "postgresql-15-sqlite-fdw" "green" >}} {{< bg "14" "postgresql-14-sqlite-fdw" "green" >}} | `postgresql-$v-sqlite-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.4.0" "sqlite_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.4.0" "sqlite_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.4.0" "sqlite_fdw_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_14 : AVAIL 5" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_15 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_14 : AVAIL 5" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_14 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "sqlite_fdw_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.5.0" "sqlite_fdw_14 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.5.0" "postgresql-17-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-sqlite-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-sqlite-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-sqlite-fdw : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-sqlite-fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-sqlite-fdw : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sqlite_fdw_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.9 KiB | [sqlite_fdw_17-2.4.0-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/sqlite_fdw_17-2.4.0-4PGDG.rhel8.x86_64.rpm) |
| `sqlite_fdw_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.3 KiB | [sqlite_fdw_17-2.4.0-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/sqlite_fdw_17-2.4.0-4PGDG.rhel8.aarch64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 65.7 KiB | [sqlite_fdw_17-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sqlite_fdw_17-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 64.9 KiB | [sqlite_fdw_17-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sqlite_fdw_17-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.9 KiB | [sqlite_fdw_17-2.4.0-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sqlite_fdw_17-2.4.0-4PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 64.1 KiB | [sqlite_fdw_17-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sqlite_fdw_17-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 63.3 KiB | [sqlite_fdw_17-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sqlite_fdw_17-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 55.3 KiB | [sqlite_fdw_17-2.4.0-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sqlite_fdw_17-2.4.0-4PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 67.1 KiB | [sqlite_fdw_17-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sqlite_fdw_17-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 66.8 KiB | [sqlite_fdw_17-2.5.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/sqlite_fdw_17-2.5.0-2PGDG.rhel10.x86_64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 65.1 KiB | [sqlite_fdw_17-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sqlite_fdw_17-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `sqlite_fdw_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.6 KiB | [sqlite_fdw_17-2.5.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/sqlite_fdw_17-2.5.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 153.7 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 148.9 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 154.1 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.0 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 188.4 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 185.7 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 159.5 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-sqlite-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 156.8 KiB | [postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-17-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sqlite_fdw_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.6 KiB | [sqlite_fdw_16-2.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/sqlite_fdw_16-2.4.0-1PGDG.rhel8.x86_64.rpm) |
| `sqlite_fdw_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.1 KiB | [sqlite_fdw_16-2.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/sqlite_fdw_16-2.4.0-1PGDG.rhel8.aarch64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.8 KiB | [sqlite_fdw_16-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sqlite_fdw_16-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 63.9 KiB | [sqlite_fdw_16-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sqlite_fdw_16-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.4 KiB | [sqlite_fdw_16-2.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sqlite_fdw_16-2.4.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.2 KiB | [sqlite_fdw_16-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sqlite_fdw_16-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.5 KiB | [sqlite_fdw_16-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sqlite_fdw_16-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.8 KiB | [sqlite_fdw_16-2.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sqlite_fdw_16-2.4.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 66.1 KiB | [sqlite_fdw_16-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sqlite_fdw_16-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 65.8 KiB | [sqlite_fdw_16-2.5.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/sqlite_fdw_16-2.5.0-2PGDG.rhel10.x86_64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [sqlite_fdw_16-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sqlite_fdw_16-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `sqlite_fdw_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.8 KiB | [sqlite_fdw_16-2.5.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/sqlite_fdw_16-2.5.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 151.7 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 147.4 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 152.0 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 148.0 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 183.5 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 181.0 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 157.0 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-sqlite-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 155.1 KiB | [postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-16-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sqlite_fdw_15` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.0 KiB | [sqlite_fdw_15-2.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sqlite_fdw_15-2.4.0-1PGDG.rhel8.x86_64.rpm) |
| `sqlite_fdw_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.4 KiB | [sqlite_fdw_15-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sqlite_fdw_15-2.3.0-1.rhel8.x86_64.rpm) |
| `sqlite_fdw_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [sqlite_fdw_15-2.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/sqlite_fdw_15-2.2.0-1.rhel8.x86_64.rpm) |
| `sqlite_fdw_15` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.4 KiB | [sqlite_fdw_15-2.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/sqlite_fdw_15-2.4.0-1PGDG.rhel8.aarch64.rpm) |
| `sqlite_fdw_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.6 KiB | [sqlite_fdw_15-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/sqlite_fdw_15-2.3.0-1.rhel8.aarch64.rpm) |
| `sqlite_fdw_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 155.8 KiB | [sqlite_fdw_15-2.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/sqlite_fdw_15-2.2.0-1.rhel8.aarch64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 67.0 KiB | [sqlite_fdw_15-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sqlite_fdw_15-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.1 KiB | [sqlite_fdw_15-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sqlite_fdw_15-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_15` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 58.1 KiB | [sqlite_fdw_15-2.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sqlite_fdw_15-2.4.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_15` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.5 KiB | [sqlite_fdw_15-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sqlite_fdw_15-2.3.0-1.rhel9.x86_64.rpm) |
| `sqlite_fdw_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 162.0 KiB | [sqlite_fdw_15-2.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sqlite_fdw_15-2.2.0-1.rhel9.x86_64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 65.7 KiB | [sqlite_fdw_15-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sqlite_fdw_15-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 64.6 KiB | [sqlite_fdw_15-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sqlite_fdw_15-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_15` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.6 KiB | [sqlite_fdw_15-2.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sqlite_fdw_15-2.4.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_15` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.3 KiB | [sqlite_fdw_15-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sqlite_fdw_15-2.3.0-1.rhel9.aarch64.rpm) |
| `sqlite_fdw_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 159.6 KiB | [sqlite_fdw_15-2.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sqlite_fdw_15-2.2.0-1.rhel9.aarch64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 68.0 KiB | [sqlite_fdw_15-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sqlite_fdw_15-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.6 KiB | [sqlite_fdw_15-2.5.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/sqlite_fdw_15-2.5.0-2PGDG.rhel10.x86_64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 66.7 KiB | [sqlite_fdw_15-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sqlite_fdw_15-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `sqlite_fdw_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.0 KiB | [sqlite_fdw_15-2.5.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/sqlite_fdw_15-2.5.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 152.6 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 148.2 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 152.9 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 148.8 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 185.2 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 183.0 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 158.6 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-sqlite-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 157.2 KiB | [postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-15-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sqlite_fdw_14` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 58.1 KiB | [sqlite_fdw_14-2.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sqlite_fdw_14-2.4.0-1PGDG.rhel8.x86_64.rpm) |
| `sqlite_fdw_14` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.4 KiB | [sqlite_fdw_14-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sqlite_fdw_14-2.3.0-1.rhel8.x86_64.rpm) |
| `sqlite_fdw_14` | `2.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 157.0 KiB | [sqlite_fdw_14-2.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sqlite_fdw_14-2.1.1-1.rhel8.x86_64.rpm) |
| `sqlite_fdw_14` | `2.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 154.8 KiB | [sqlite_fdw_14-2.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sqlite_fdw_14-2.1.0-1.rhel8.x86_64.rpm) |
| `sqlite_fdw_14` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.5 KiB | [sqlite_fdw_14-2.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/sqlite_fdw_14-2.4.0-1PGDG.rhel8.aarch64.rpm) |
| `sqlite_fdw_14` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.7 KiB | [sqlite_fdw_14-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/sqlite_fdw_14-2.3.0-1.rhel8.aarch64.rpm) |
| `sqlite_fdw_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 156.4 KiB | [sqlite_fdw_14-2.2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/sqlite_fdw_14-2.2.0-1.rhel8.aarch64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 67.1 KiB | [sqlite_fdw_14-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sqlite_fdw_14-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.1 KiB | [sqlite_fdw_14-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sqlite_fdw_14-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_14` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 58.2 KiB | [sqlite_fdw_14-2.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sqlite_fdw_14-2.4.0-1PGDG.rhel9.x86_64.rpm) |
| `sqlite_fdw_14` | `2.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 53.6 KiB | [sqlite_fdw_14-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sqlite_fdw_14-2.3.0-1.rhel9.x86_64.rpm) |
| `sqlite_fdw_14` | `2.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 159.0 KiB | [sqlite_fdw_14-2.1.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sqlite_fdw_14-2.1.1-1.rhel9.x86_64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 65.7 KiB | [sqlite_fdw_14-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sqlite_fdw_14-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 64.7 KiB | [sqlite_fdw_14-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sqlite_fdw_14-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_14` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 56.7 KiB | [sqlite_fdw_14-2.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sqlite_fdw_14-2.4.0-1PGDG.rhel9.aarch64.rpm) |
| `sqlite_fdw_14` | `2.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 52.2 KiB | [sqlite_fdw_14-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sqlite_fdw_14-2.3.0-1.rhel9.aarch64.rpm) |
| `sqlite_fdw_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 160.4 KiB | [sqlite_fdw_14-2.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sqlite_fdw_14-2.2.0-1.rhel9.aarch64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 68.4 KiB | [sqlite_fdw_14-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/sqlite_fdw_14-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.7 KiB | [sqlite_fdw_14-2.5.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/sqlite_fdw_14-2.5.0-2PGDG.rhel10.x86_64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 66.7 KiB | [sqlite_fdw_14-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/sqlite_fdw_14-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `sqlite_fdw_14` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.0 KiB | [sqlite_fdw_14-2.5.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/sqlite_fdw_14-2.5.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 152.8 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 148.5 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 153.3 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 149.2 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 185.4 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 183.4 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 159.1 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-sqlite-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 157.4 KiB | [postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sqlite-fdw/postgresql-14-sqlite-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgspider/sqlite_fdw" title="Repository" icon="github" subtitle="github.com/pgspider/sqlite_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="sqlite_fdw-2.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg sqlite_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install sqlite_fdw;		# install via package name, for the active PG version

pig install sqlite_fdw -v 17;   # install for PG 17
pig install sqlite_fdw -v 16;   # install for PG 16
pig install sqlite_fdw -v 15;   # install for PG 15
pig install sqlite_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION sqlite_fdw;
```




## Usage

> [sqlite_fdw: SQLite Foreign Data Wrapper](https://github.com/pgspider/sqlite_fdw)

### Create Server

```sql
CREATE EXTENSION sqlite_fdw;

CREATE SERVER sqlite_server FOREIGN DATA WRAPPER sqlite_fdw
  OPTIONS (database '/path/to/database.db');
```

**Server Options:** `database` (required, path to SQLite file), `updatable` (default `true`), `truncatable` (default `false`), `keep_connections` (default `true`), `batch_size` (default 1), `force_readonly` (default `false`).

No `CREATE USER MAPPING` is required since SQLite has no authentication model.

### Create Foreign Table

```sql
CREATE FOREIGN TABLE remote_data (
  id integer OPTIONS (key 'true'),
  name text,
  created timestamp OPTIONS (column_type 'INT'),
  data bytea
)
SERVER sqlite_server
OPTIONS (table 'data_table');
```

**Table Options:** `table` (SQLite table name if different from PostgreSQL name), `updatable`, `truncatable`, `batch_size`.

**Column Options:** `column_name` (map to different SQLite column name), `column_type` (SQLite affinity: `INT` for epoch timestamps, `BLOB` for UUIDs), `key` (mark as primary key for UPDATE/DELETE).

### CRUD Operations

```sql
SELECT * FROM remote_data WHERE id > 100;
INSERT INTO remote_data (id, name) VALUES (1, 'test');
UPDATE remote_data SET name = 'updated' WHERE id = 1;
DELETE FROM remote_data WHERE id = 1;
```

### Import Foreign Schema

```sql
IMPORT FOREIGN SCHEMA public
  FROM SERVER sqlite_server INTO local_schema;
```

**Import Options:** `import_default` (default `false`), `import_not_null` (default `true`).

### Data Type Mapping

| SQLite Type | PostgreSQL Type |
|-------------|-----------------|
| int | bigint |
| text, char, clob | text |
| blob | bytea |
| real, float, double | double precision |
| datetime | timestamp |
| uuid | uuid |
| json, jsonb | json, jsonb |

Timestamps can be stored as TEXT (ISO format) or INT (Unix epoch, use `column_type 'INT'`). UUIDs can be stored as TEXT (36 chars) or BLOB (16 bytes). The SQLite database file must be readable (and writable for DML) by the PostgreSQL OS user.
