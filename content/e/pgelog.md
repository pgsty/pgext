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

> [!Note] Requires the dblink extension at runtime in addition to pg_variables.


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
| `pgelog_18` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_18-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_18` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_18-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_18` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_18-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_18-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_18` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_18-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_18-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_18` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_18-1.0.2-1PIGSTY.el10.noarch.rpm) |
| `pgelog_18` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_18-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_18-1.0.2-1PIGSTY.el10.noarch.rpm) |
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
| `pgelog_17` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_17-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_17` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_17-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_17` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_17-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_17-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_17` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_17-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_17-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_17` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_17-1.0.2-1PIGSTY.el10.noarch.rpm) |
| `pgelog_17` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_17-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_17-1.0.2-1PIGSTY.el10.noarch.rpm) |
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
| `pgelog_16` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_16-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_16` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_16-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_16` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_16-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_16-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_16` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_16-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_16-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_16` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_16-1.0.2-1PIGSTY.el10.noarch.rpm) |
| `pgelog_16` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_16-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_16-1.0.2-1PIGSTY.el10.noarch.rpm) |
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
| `pgelog_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_15-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_15` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_15-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_15` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_15-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_15-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_15` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_15-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_15-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_15` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_15-1.0.2-1PIGSTY.el10.noarch.rpm) |
| `pgelog_15` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_15-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_15-1.0.2-1PIGSTY.el10.noarch.rpm) |
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
| `pgelog_14` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgelog_14-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_14` | `1.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgelog_14-1.0.2-1PIGSTY.el8.noarch.rpm) |
| `pgelog_14` | `1.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.6 KiB | [pgelog_14-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgelog_14-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_14` | `1.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.5 KiB | [pgelog_14-1.0.2-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgelog_14-1.0.2-1PIGSTY.el9.noarch.rpm) |
| `pgelog_14` | `1.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgelog_14-1.0.2-1PIGSTY.el10.noarch.rpm) |
| `pgelog_14` | `1.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.7 KiB | [pgelog_14-1.0.2-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgelog_14-1.0.2-1PIGSTY.el10.noarch.rpm) |
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

> Syntax:
>
> ```sql
> CREATE EXTENSION IF NOT EXISTS dblink;
> CREATE EXTENSION IF NOT EXISTS pg_variables;
> CREATE EXTENSION pgelog;
> SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
> ```
>
> Source: [README](https://github.com/anfiau/pgelog)

`pgelog` writes log records into PostgreSQL tables using pseudo-autonomous transactions implemented through `dblink`. The key goal is that log entries survive even when the caller's main transaction rolls back.

## Prerequisites

The README requires:

- PostgreSQL 11 or newer
- `dblink`
- `pg_variables`
- local passwordless `dblink` access, typically via a `peer` local entry in `pg_hba.conf`

It also warns that each session may open one extra connection for `dblink`, so `max_connections` should be sized accordingly.

## Objects

The extension creates:

- `pgelog_params` for configuration
- `pgelog_logs` as the base log table
- `pgelog_vw_logs` as a log view with timing information

The log table/view stores fields such as timestamp, log type, source function, phase, message text, transaction id, SQLSTATE, SQLERRM, and connection name.

## Basic Logging

Write a log entry:

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');
```

Read the latest log:

```sql
SELECT log_stamp, log_info
FROM pgelog_logs
ORDER BY log_stamp DESC
LIMIT 1;
```

## PL/pgSQL Exception Logging

The README includes a larger PL/pgSQL example that catches exceptions, collects diagnostics, writes a `FAIL` log entry through `pgelog_to_log(...)`, and then re-raises the exception. This is the main pattern for capturing rollback-resistant failure logs.

## Configuration

Configuration parameters are managed with:

```sql
SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

The README documents `pgelog_ttl_minutes` and other parameters through the `pgelog_params` table and helper functions.
