---
title: "db_migrator"
linkTitle: "db_migrator"
description: "Tools to migrate other databases to PostgreSQL"
weight: 9550
categories: ["ETL"]
width: full
---

[**db_migrator**](https://github.com/cybertec-postgresql/db_migrator) : Tools to migrate other databases to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9550** | {{< badge content="db_migrator" link="https://github.com/cybertec-postgresql/db_migrator" >}} | {{< ext "db_migrator" >}} | `1.0.0` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "orafce" >}} {{< ext "pg_bulkload" >}} {{< ext "jdbc_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "pgtt" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `db_migrator` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "db_migrator_18" "green" >}} {{< bg "17" "db_migrator_17" "green" >}} {{< bg "16" "db_migrator_16" "green" >}} {{< bg "15" "db_migrator_15" "green" >}} {{< bg "14" "db_migrator_14" "green" >}} | `db_migrator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-db-migrator" "green" >}} {{< bg "17" "postgresql-17-db-migrator" "green" >}} {{< bg "16" "postgresql-16-db-migrator" "green" >}} {{< bg "15" "postgresql-15-db-migrator" "green" >}} {{< bg "14" "postgresql-14-db-migrator" "green" >}} | `postgresql-$v-db-migrator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-db-migrator : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-db-migrator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-db-migrator : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.2 KiB | [postgresql-18-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-18-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.2 KiB | [postgresql-17-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-17-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.2 KiB | [postgresql-16-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-16-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.2 KiB | [postgresql-15-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-15-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.2 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.2 KiB | [postgresql-14-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-14-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/db_migrator" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/db_migrator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="db_migrator-RELEASE_1_0_0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg db_migrator;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install db_migrator;		# install via package name, for the active PG version

pig install db_migrator -v 18;   # install for PG 18
pig install db_migrator -v 17;   # install for PG 17
pig install db_migrator -v 16;   # install for PG 16
pig install db_migrator -v 15;   # install for PG 15
pig install db_migrator -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION db_migrator;
```



## Usage

> [db_migrator: Tools to migrate other databases to PostgreSQL](https://github.com/cybertec-postgresql/db_migrator)

A framework for migrating databases from other data sources to PostgreSQL using foreign data wrappers and source-specific plugins.

### Enabling

```sql
CREATE EXTENSION db_migrator;
```

### Available Plugins

- **ora_migrator** - Oracle migration
- **mysql_migrator** - MySQL/MariaDB migration
- **mssql_migrator** - Microsoft SQL Server migration

### Complete Migration Example (Oracle)

```sql
-- Setup (as superuser)
CREATE EXTENSION oracle_fdw;
CREATE SERVER oracle FOREIGN DATA WRAPPER oracle_fdw
    OPTIONS (dbserver '//dbserver.mydomain.com/ORADB');
GRANT USAGE ON FOREIGN SERVER oracle TO migrator;
CREATE USER MAPPING FOR migrator SERVER oracle
    OPTIONS (user 'orauser', password 'orapwd');

-- Migrate (as migrator user)
CREATE EXTENSION ora_migrator;

SELECT db_migrate(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{TESTSCHEMA1,TESTSCHEMA2}'
);
```

### Step-by-Step Migration

For more control, execute migration in stages:

```sql
-- 1. Create staging schemas and snapshot metadata
SELECT db_migrate_prepare(
    plugin => 'ora_migrator',
    server => 'oracle',
    only_schemas => '{SCHEMA1}'
);

-- 2. Review and modify staging data
-- Edit pgsql_stage tables to customize type mappings, rename objects, etc.
UPDATE pgsql_stage.tables SET migrate = TRUE WHERE ...;

-- 3. Create schemas and migrate data
SELECT db_migrate_mkforeign(plugin => 'ora_migrator', server => 'oracle');
SELECT db_migrate_tables(plugin => 'ora_migrator');

-- 4. Create constraints and indexes
SELECT db_migrate_constraints(plugin => 'ora_migrator');
SELECT db_migrate_indexes(plugin => 'ora_migrator');

-- 5. Cleanup
SELECT db_migrate_finish();
```

### Key Features

- Migrates tables, sequences, indexes, constraints, views, functions
- Data type mapping from source to PostgreSQL types (customizable)
- Continues on errors, reporting which objects failed
- Uses FDW staging schema for metadata inspection before migration
- Schema and object name translation via plugin functions
