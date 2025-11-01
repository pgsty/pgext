---
title: "pgextwlist"
linkTitle: "pgextwlist"
description: "PostgreSQL Extension Whitelisting"
weight: 7180
categories: ["SEC"]
width: full
---

PostgreSQL Extension Whitelisting


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7180** | {{< badge content="pgextwlist" link="https://github.com/dimitri/pgextwlist" >}} | {{< ext "pgextwlist" >}} | `1.19` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "ddlx" >}} {{< ext "pgdd" >}} {{< ext "pg_permissions" >}} {{< ext "adminpack" >}} {{< ext "pgaudit" >}} {{< ext "set_user" >}} {{< ext "pg_catcheck" >}} {{< ext "noset" >}} |

> [!Note] missing pg18 on el


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgextwlist" >}} | `1.19` | {{< bg "18" "pgextwlist_18*" "green" >}} {{< bg "17" "pgextwlist_17*" "green" >}} {{< bg "16" "pgextwlist_16*" "green" >}} {{< bg "15" "pgextwlist_15*" "green" >}} {{< bg "14" "pgextwlist_14*" "green" >}} {{< bg "13" "pgextwlist_13*" "green" >}} | `pgextwlist_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgextwlist" >}} | `1.19` | {{< bg "18" "postgresql-18-pgextwlist" "green" >}} {{< bg "17" "postgresql-17-pgextwlist" "green" >}} {{< bg "16" "postgresql-16-pgextwlist" "green" >}} {{< bg "15" "postgresql-15-pgextwlist" "green" >}} {{< bg "14" "postgresql-14-pgextwlist" "green" >}} {{< bg "13" "postgresql-13-pgextwlist" "green" >}} | `postgresql-$v-pgextwlist` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 2" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.19" "pgextwlist_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.19" "pgextwlist_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.19" "postgresql-18-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-17-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-16-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-15-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-14-pgextwlist : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.19" "postgresql-13-pgextwlist : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_18` | 1.19 | `el8.x86_64` | pigsty | 20.3 KiB | [pgextwlist_18-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_18-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_18` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_18-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_18` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_18-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_18-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_18` | 1.19 | `el9.aarch64` | pigsty | 20.0 KiB | [pgextwlist_18-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_18-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_18` | 1.19 | `el10.x86_64` | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_18-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_18` | 1.19 | `el10.aarch64` | pigsty | 20.2 KiB | [pgextwlist_18-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_18-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 29.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.7 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 29.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.8 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 30.1 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 29.3 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 29.2 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.6 KiB | [postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-18-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_17` | 1.19 | `el8.x86_64` | pigsty | 20.3 KiB | [pgextwlist_17-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_17-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_17` | 1.17 | `el8.x86_64` | pigsty | 19.9 KiB | [pgextwlist_17-1.17-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_17-1.17-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_17` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_17-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_17-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_17` | 1.17 | `el8.aarch64` | pigsty | 19.6 KiB | [pgextwlist_17-1.17-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_17-1.17-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_17` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_17-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_17-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_17` | 1.17 | `el9.x86_64` | pigsty | 20.0 KiB | [pgextwlist_17-1.17-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_17-1.17-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_17` | 1.19 | `el9.aarch64` | pigsty | 20.0 KiB | [pgextwlist_17-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_17-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_17` | 1.17 | `el9.aarch64` | pigsty | 19.5 KiB | [pgextwlist_17-1.17-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_17-1.17-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_17` | 1.19 | `el10.x86_64` | pigsty | 20.3 KiB | [pgextwlist_17-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_17-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_17` | 1.19 | `el10.aarch64` | pigsty | 20.2 KiB | [pgextwlist_17-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_17-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 29.1 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.7 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 29.0 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.8 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 38.3 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 37.6 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 29.2 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.6 KiB | [postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-17-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_16` | 1.19 | `el8.x86_64` | pigsty | 20.3 KiB | [pgextwlist_16-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_16-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_16` | 1.17 | `el8.x86_64` | pigsty | 19.9 KiB | [pgextwlist_16-1.17-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_16-1.17-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_16` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_16-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_16` | 1.17 | `el8.aarch64` | pigsty | 19.6 KiB | [pgextwlist_16-1.17-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_16-1.17-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_16` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_16-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_16-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_16` | 1.17 | `el9.x86_64` | pigsty | 20.0 KiB | [pgextwlist_16-1.17-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_16-1.17-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_16` | 1.19 | `el9.aarch64` | pigsty | 20.0 KiB | [pgextwlist_16-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_16-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_16` | 1.17 | `el9.aarch64` | pigsty | 19.5 KiB | [pgextwlist_16-1.17-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_16-1.17-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_16` | 1.19 | `el10.x86_64` | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_16-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_16` | 1.19 | `el10.aarch64` | pigsty | 20.2 KiB | [pgextwlist_16-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_16-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 29.1 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.7 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 29.0 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.8 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 37.7 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 37.0 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 29.2 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.6 KiB | [postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-16-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_15` | 1.19 | `el8.x86_64` | pigsty | 20.3 KiB | [pgextwlist_15-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_15-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_15` | 1.17 | `el8.x86_64` | pigsty | 19.9 KiB | [pgextwlist_15-1.17-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_15-1.17-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_15` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_15-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_15-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_15` | 1.17 | `el8.aarch64` | pigsty | 19.5 KiB | [pgextwlist_15-1.17-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_15-1.17-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_15` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_15-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_15-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_15` | 1.17 | `el9.x86_64` | pigsty | 20.3 KiB | [pgextwlist_15-1.17-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_15-1.17-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_15` | 1.19 | `el9.aarch64` | pigsty | 20.1 KiB | [pgextwlist_15-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_15-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_15` | 1.17 | `el9.aarch64` | pigsty | 19.6 KiB | [pgextwlist_15-1.17-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_15-1.17-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_15` | 1.19 | `el10.x86_64` | pigsty | 20.4 KiB | [pgextwlist_15-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_15-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_15` | 1.19 | `el10.aarch64` | pigsty | 20.3 KiB | [pgextwlist_15-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_15-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 28.9 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.4 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 28.9 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.7 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 37.5 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 36.8 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 29.0 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.5 KiB | [postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-15-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_14` | 1.19 | `el8.x86_64` | pigsty | 20.3 KiB | [pgextwlist_14-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_14-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_14` | 1.17 | `el8.x86_64` | pigsty | 19.9 KiB | [pgextwlist_14-1.17-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_14-1.17-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_14` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_14-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_14-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_14` | 1.17 | `el8.aarch64` | pigsty | 19.5 KiB | [pgextwlist_14-1.17-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_14-1.17-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_14` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_14-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_14-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_14` | 1.17 | `el9.x86_64` | pigsty | 20.3 KiB | [pgextwlist_14-1.17-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_14-1.17-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_14` | 1.19 | `el9.aarch64` | pigsty | 20.0 KiB | [pgextwlist_14-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_14-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_14` | 1.17 | `el9.aarch64` | pigsty | 19.7 KiB | [pgextwlist_14-1.17-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_14-1.17-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_14` | 1.19 | `el10.x86_64` | pigsty | 20.3 KiB | [pgextwlist_14-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_14-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_14` | 1.19 | `el10.aarch64` | pigsty | 20.2 KiB | [pgextwlist_14-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_14-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 28.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.4 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 28.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.5 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 37.5 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 36.8 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 29.0 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.4 KiB | [postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-14-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgextwlist_13` | 1.19 | `el8.x86_64` | pigsty | 20.1 KiB | [pgextwlist_13-1.19-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_13-1.19-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_13` | 1.17 | `el8.x86_64` | pigsty | 19.7 KiB | [pgextwlist_13-1.17-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgextwlist_13-1.17-1PIGSTY.el8.x86_64.rpm) |
| `pgextwlist_13` | 1.19 | `el8.aarch64` | pigsty | 20.2 KiB | [pgextwlist_13-1.19-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_13-1.19-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_13` | 1.17 | `el8.aarch64` | pigsty | 19.5 KiB | [pgextwlist_13-1.17-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgextwlist_13-1.17-1PIGSTY.el8.aarch64.rpm) |
| `pgextwlist_13` | 1.19 | `el9.x86_64` | pigsty | 20.4 KiB | [pgextwlist_13-1.19-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_13-1.19-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_13` | 1.17 | `el9.x86_64` | pigsty | 20.2 KiB | [pgextwlist_13-1.17-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgextwlist_13-1.17-1PIGSTY.el9.x86_64.rpm) |
| `pgextwlist_13` | 1.19 | `el9.aarch64` | pigsty | 20.0 KiB | [pgextwlist_13-1.19-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_13-1.19-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_13` | 1.17 | `el9.aarch64` | pigsty | 19.6 KiB | [pgextwlist_13-1.17-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgextwlist_13-1.17-1PIGSTY.el9.aarch64.rpm) |
| `pgextwlist_13` | 1.19 | `el10.x86_64` | pigsty | 20.3 KiB | [pgextwlist_13-1.19-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgextwlist_13-1.19-1PIGSTY.el10.x86_64.rpm) |
| `pgextwlist_13` | 1.19 | `el10.aarch64` | pigsty | 20.2 KiB | [pgextwlist_13-1.19-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgextwlist_13-1.19-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pgextwlist` | 1.19 | `d12.x86_64` | pgdg | 28.7 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `d12.aarch64` | pgdg | 28.0 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `d13.x86_64` | pgdg | 28.7 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `d13.aarch64` | pgdg | 28.0 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `u22.x86_64` | pgdg | 36.6 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `u22.aarch64` | pgdg | 36.0 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `u24.x86_64` | pgdg | 28.8 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgextwlist` | 1.19 | `u24.aarch64` | pgdg | 28.0 KiB | [postgresql-13-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgextwlist/postgresql-13-pgextwlist_1.19-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/dimitri/pgextwlist" title="Repository" icon="github" subtitle="github.com/dimitri/pgextwlist" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgextwlist-1.19.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgextwlist; # get pgextwlist source code
pig build dep pgextwlist; # install build dependencies
pig build pkg pgextwlist; # build extension rpm or deb
pig build ext pgextwlist; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgextwlist; # install by extension name, for the current active PG version
pig ext install pgextwlist; # install via package alias, for the active PG version
pig ext install pgextwlist -v 18;   # install for PG 18
pig ext install pgextwlist -v 17;   # install for PG 17
pig ext install pgextwlist -v 16;   # install for PG 16
pig ext install pgextwlist -v 15;   # install for PG 15
pig ext install pgextwlist -v 14;   # install for PG 14
pig ext install pgextwlist -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgextwlist;
```

