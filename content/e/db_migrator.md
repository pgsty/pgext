---
title: "db_migrator"
linkTitle: "db_migrator"
description: "Tools to migrate other databases to PostgreSQL"
weight: 9540
categories: ["ETL"]
width: full
---

[**db_migrator**](https://github.com/cybertec-postgresql/db_migrator) : Tools to migrate other databases to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9540** | {{< badge content="db_migrator" link="https://github.com/cybertec-postgresql/db_migrator" >}} | {{< ext "db_migrator" >}} | `1.0.0` | {{< category "ETL" >}} | {{< license "BSD 3-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mysql_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "orafce" >}} {{< ext "pg_bulkload" >}} {{< ext "jdbc_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "pgtt" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `db_migrator` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "db_migrator_18" "green" >}} {{< bg "17" "db_migrator_17" "green" >}} {{< bg "16" "db_migrator_16" "green" >}} {{< bg "15" "db_migrator_15" "green" >}} {{< bg "14" "db_migrator_14" "green" >}} {{< bg "13" "db_migrator_13" "green" >}} | `db_migrator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-db-migrator" "green" >}} {{< bg "17" "postgresql-17-db-migrator" "green" >}} {{< bg "16" "postgresql-16-db-migrator" "green" >}} {{< bg "15" "postgresql-15-db-migrator" "green" >}} {{< bg "14" "postgresql-14-db-migrator" "green" >}} {{< bg "13" "postgresql-13-db-migrator" "green" >}} | `postgresql-$v-db-migrator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "db_migrator_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-db-migrator : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-db-migrator : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `db_migrator_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 26.4 KiB | [db_migrator_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/db_migrator_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `db_migrator_13` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 26.4 KiB | [db_migrator_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/db_migrator_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `db_migrator_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 25.3 KiB | [db_migrator_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/db_migrator_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `db_migrator_13` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 25.3 KiB | [db_migrator_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/db_migrator_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `db_migrator_13` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 25.4 KiB | [db_migrator_13-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/db_migrator_13-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `db_migrator_13` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 25.4 KiB | [db_migrator_13-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/db_migrator_13-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-db-migrator` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.1 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.1 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.1 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.1 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.3 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 21.3 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 21.3 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-db-migrator` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 21.3 KiB | [postgresql-13-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/db-migrator/postgresql-13-db-migrator_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/db_migrator" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/db_migrator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="db_migrator-RELEASE_1_0_0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg db_migrator;		# build rpm / deb with pig
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
pig install db_migrator -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION db_migrator;
```
