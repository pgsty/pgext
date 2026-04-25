---
title: "pgelog"
linkTitle: "pgelog"
description: "Extended logging via pseudo-autonomous transactions"
weight: 5870
categories: ["ADMIN"]
width: full
---

[**pgelog**](https://github.com/anfiau/pgelog) : Extended logging via pseudo-autonomous transactions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5870** | {{< badge content="pgelog" link="https://github.com/anfiau/pgelog" >}} | {{< ext "pgelog" >}} | `1.0.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "dblink" >}} {{< ext "pg_variables" >}} |
|   **See Also**    | {{< ext "table_log" >}} {{< ext "pgaudit" >}} {{< ext "logerrors" >}} {{< ext "dblink" >}} |

> [!Note] Release tag 1.0.2 still ships extension SQL version 1.0; requires the dblink extension at runtime in addition to pg_variables.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgelog` | `dblink`, `pg_variables` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "pgelog_18" "green" >}} {{< bg "17" "pgelog_17" "green" >}} {{< bg "16" "pgelog_16" "green" >}} {{< bg "15" "pgelog_15" "green" >}} {{< bg "14" "pgelog_14" "green" >}} | `pgelog_$v` | `postgresql$v-contrib`, `pg_variables_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.2` | {{< bg "18" "postgresql-18-pgelog" "green" >}} {{< bg "17" "postgresql-17-pgelog" "green" >}} {{< bg "16" "postgresql-16-pgelog" "green" >}} {{< bg "15" "postgresql-15-pgelog" "green" >}} {{< bg "14" "postgresql-14-pgelog" "green" >}} | `postgresql-$v-pgelog` | `postgresql-$v-pg-variables` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pgelog_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-18-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-17-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-16-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-15-pgelog : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "postgresql-14-pgelog : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgelog_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pgelog_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pgelog_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pgelog_18` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.5 KiB | [pgelog_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pgelog_18` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pgelog_18` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_18-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pgelog_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_18-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgelog` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.9 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.9 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgelog` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.9 KiB | [postgresql-18-pgelog_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-18-pgelog_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgelog_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pgelog_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pgelog_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pgelog_17` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.5 KiB | [pgelog_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pgelog_17` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pgelog_17` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_17-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pgelog_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_17-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgelog` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.9 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.9 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgelog` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.9 KiB | [postgresql-17-pgelog_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-17-pgelog_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgelog_16` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_16-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pgelog_16` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pgelog_16-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_16-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pgelog_16` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.5 KiB | [pgelog_16-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_16-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pgelog_16` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_16-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_16-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pgelog_16` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_16-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pgelog_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_16-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgelog` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.9 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.9 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgelog` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.9 KiB | [postgresql-16-pgelog_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-16-pgelog_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgelog_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_15-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pgelog_15` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pgelog_15-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_15-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pgelog_15` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_15-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_15-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pgelog_15` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_15-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_15-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pgelog_15` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_15-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pgelog_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_15-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgelog` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.9 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.9 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgelog` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.9 KiB | [postgresql-15-pgelog_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-15-pgelog_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgelog_14` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_14-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pgelog_14` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.6 KiB | [pgelog_14-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_14-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pgelog_14` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_14-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_14-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pgelog_14` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_14-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_14-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pgelog_14` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_14-1.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pgelog_14` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_14-1.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgelog` | `1.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.8 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.8 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.9 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgelog` | `1.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.9 KiB | [postgresql-14-pgelog_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgelog/postgresql-14-pgelog_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/anfiau/pgelog" title="Repository" icon="github" subtitle="github.com/anfiau/pgelog" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgelog-1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgelog;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgelog;		# install via package name, for the active PG version

pig install pgelog -v 18;   # install for PG 18
pig install pgelog -v 17;   # install for PG 17
pig install pgelog -v 16;   # install for PG 16
pig install pgelog -v 15;   # install for PG 15
pig install pgelog -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgelog CASCADE; -- requires dblink, pg_variables
```

## Usage

Source: [README](https://github.com/anfiau/pgelog/blob/master/README.md), [control file](https://github.com/anfiau/pgelog/blob/master/pgelog.control), [extension SQL 1.0.2](https://github.com/anfiau/pgelog/blob/master/pgelog--1.0.2.sql), [Tag v1.0.2](https://github.com/anfiau/pgelog/tree/v1.0.2)

`pgelog` writes rollback-resistant log rows through pseudo-autonomous transactions implemented with `dblink`. It caches the extra connection in `pg_variables` so repeated logging in the same session is cheaper.

### Requirements and install

- PostgreSQL 11+
- `dblink`
- `pg_variables`
- passwordless local `dblink` access, typically via `peer`

```sql
CREATE EXTENSION IF NOT EXISTS dblink;
CREATE EXTENSION IF NOT EXISTS pg_variables;
CREATE EXTENSION pgelog;
```

### Main objects and functions

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');

SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

- `pgelog_logs`: base log table.
- `pgelog_vw_logs`: log view with timing deltas.
- `pgelog_params`: configuration table.
- `pgelog_to_log(...)`: write a log row that survives rollback.
- `pgelog_get_param(...)`, `pgelog_set_param(...)`, `pgelog_delete_param(...)`: manage extension settings.

### Typical use

The official README shows `pgelog_to_log(...)` in PL/pgSQL exception handlers: collect diagnostics with `GET STACKED DIAGNOSTICS`, write a `FAIL` log row, then re-raise the error.

### Caveats

- Each session can open one additional `dblink` connection, so `max_connections` should account for that.
- Upstream `v1.0.2` still ships extension SQL under the same user-visible object family; Pigsty's note about runtime `dblink` plus `pg_variables` requirements is confirmed by the control file and README.
