---
title: "pgnodemx"
linkTitle: "pgnodemx"
description: "Capture node OS metrics via SQL queries"
weight: 6310
categories: ["STAT"]
width: full
---

Capture node OS metrics via SQL queries


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6310** | {{< badge content="pgnodemx" link="https://github.com/CrunchyData/pgnodemx" >}} | {{< ext "pgnodemx" >}} | `1.7` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "system_stats" >}} {{< ext "pg_wait_sampling" >}} {{< ext "pgsentinel" >}} {{< ext "pgmeminfo" >}} {{< ext "pgfincore" >}} {{< ext "prioritize" >}} {{< ext "pg_buffercache" >}} |
|    **Siblings**   | {{< ext "pg_proctab" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgnodemx" >}} | `1.7` | {{< bg "18" "pgnodemx_18" "green" >}} {{< bg "17" "pgnodemx_17" "green" >}} {{< bg "16" "pgnodemx_16" "green" >}} {{< bg "15" "pgnodemx_15" "green" >}} {{< bg "14" "pgnodemx_14" "green" >}} {{< bg "13" "pgnodemx_13" "green" >}} | `pgnodemx_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgnodemx" >}} | `1.7` | {{< bg "18" "postgresql-18-pgnodemx" "green" >}} {{< bg "17" "postgresql-17-pgnodemx" "green" >}} {{< bg "16" "postgresql-16-pgnodemx" "green" >}} {{< bg "15" "postgresql-15-pgnodemx" "green" >}} {{< bg "14" "postgresql-14-pgnodemx" "green" >}} {{< bg "13" "postgresql-13-pgnodemx" "green" >}} | `postgresql-$v-pgnodemx` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.7" "pgnodemx_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.7" "postgresql-18-pgnodemx : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-17-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-16-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-15-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-14-pgnodemx : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7" "postgresql-13-pgnodemx : AVAIL 2" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_18` | 1.7 | `el8.x86_64` | pigsty | 37.5 KiB | [pgnodemx_18-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_18-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el8.x86_64` | pgdg | 41.8 KiB | [pgnodemx_18-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgnodemx_18-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el8.aarch64` | pigsty | 37.0 KiB | [pgnodemx_18-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_18-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_18` | 1.7 | `el8.aarch64` | pgdg | 41.1 KiB | [pgnodemx_18-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgnodemx_18-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_18` | 1.7 | `el9.x86_64` | pigsty | 35.8 KiB | [pgnodemx_18-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_18-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el9.x86_64` | pgdg | 41.6 KiB | [pgnodemx_18-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgnodemx_18-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el9.aarch64` | pigsty | 35.5 KiB | [pgnodemx_18-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_18-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_18` | 1.7 | `el9.aarch64` | pgdg | 41.2 KiB | [pgnodemx_18-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgnodemx_18-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_18` | 1.7 | `el10.x86_64` | pigsty | 36.0 KiB | [pgnodemx_18-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_18-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el10.x86_64` | pgdg | 42.2 KiB | [pgnodemx_18-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgnodemx_18-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_18` | 1.7 | `el10.aarch64` | pigsty | 35.6 KiB | [pgnodemx_18-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_18-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_18` | 1.7 | `el10.aarch64` | pgdg | 41.7 KiB | [pgnodemx_18-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgnodemx_18-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 82.3 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 81.4 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 82.2 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 81.7 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 81.9 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 80.9 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 81.8 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 80.8 KiB | [postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-18-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_17` | 1.7 | `el8.x86_64` | pigsty | 37.5 KiB | [pgnodemx_17-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_17-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el8.x86_64` | pgdg | 41.8 KiB | [pgnodemx_17-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgnodemx_17-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el8.aarch64` | pigsty | 37.0 KiB | [pgnodemx_17-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_17-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_17` | 1.7 | `el8.aarch64` | pgdg | 41.2 KiB | [pgnodemx_17-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgnodemx_17-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_17` | 1.7 | `el9.x86_64` | pigsty | 35.8 KiB | [pgnodemx_17-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_17-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el9.x86_64` | pgdg | 41.6 KiB | [pgnodemx_17-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgnodemx_17-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el9.aarch64` | pigsty | 35.5 KiB | [pgnodemx_17-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_17-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_17` | 1.7 | `el9.aarch64` | pgdg | 41.2 KiB | [pgnodemx_17-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgnodemx_17-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_17` | 1.7 | `el10.x86_64` | pigsty | 36.0 KiB | [pgnodemx_17-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_17-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el10.x86_64` | pgdg | 42.3 KiB | [pgnodemx_17-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgnodemx_17-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_17` | 1.7 | `el10.aarch64` | pigsty | 35.6 KiB | [pgnodemx_17-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_17-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_17` | 1.7 | `el10.aarch64` | pgdg | 41.6 KiB | [pgnodemx_17-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgnodemx_17-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgnodemx` | 1.7 | `d12.x86_64` | pigsty | 91.0 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 82.2 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `d12.aarch64` | pigsty | 89.8 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 81.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 82.2 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 81.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u22.x86_64` | pigsty | 96.1 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 88.9 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u22.aarch64` | pigsty | 95.5 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 87.5 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u24.x86_64` | pigsty | 87.6 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 81.8 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u24.aarch64` | pigsty | 87.5 KiB | [postgresql-17-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 80.7 KiB | [postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-17-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_16` | 1.7 | `el8.x86_64` | pigsty | 37.5 KiB | [pgnodemx_16-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_16-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el8.x86_64` | pgdg | 41.8 KiB | [pgnodemx_16-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgnodemx_16-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el8.aarch64` | pigsty | 37.0 KiB | [pgnodemx_16-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_16-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_16` | 1.7 | `el8.aarch64` | pgdg | 41.1 KiB | [pgnodemx_16-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgnodemx_16-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_16` | 1.7 | `el9.x86_64` | pigsty | 35.8 KiB | [pgnodemx_16-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_16-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el9.x86_64` | pgdg | 41.6 KiB | [pgnodemx_16-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgnodemx_16-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el9.aarch64` | pigsty | 35.4 KiB | [pgnodemx_16-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_16-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_16` | 1.7 | `el9.aarch64` | pgdg | 41.1 KiB | [pgnodemx_16-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgnodemx_16-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_16` | 1.7 | `el10.x86_64` | pigsty | 36.0 KiB | [pgnodemx_16-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_16-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el10.x86_64` | pgdg | 42.2 KiB | [pgnodemx_16-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgnodemx_16-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_16` | 1.7 | `el10.aarch64` | pigsty | 35.6 KiB | [pgnodemx_16-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_16-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_16` | 1.7 | `el10.aarch64` | pgdg | 41.7 KiB | [pgnodemx_16-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgnodemx_16-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgnodemx` | 1.7 | `d12.x86_64` | pigsty | 91.0 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 82.2 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `d12.aarch64` | pigsty | 89.9 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 81.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 82.2 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 81.6 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u22.x86_64` | pigsty | 96.1 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 88.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u22.aarch64` | pigsty | 95.5 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 87.5 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u24.x86_64` | pigsty | 87.6 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 81.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u24.aarch64` | pigsty | 87.5 KiB | [postgresql-16-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 80.8 KiB | [postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-16-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_15` | 1.7 | `el8.x86_64` | pigsty | 38.7 KiB | [pgnodemx_15-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_15-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el8.x86_64` | pgdg | 43.1 KiB | [pgnodemx_15-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgnodemx_15-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el8.aarch64` | pigsty | 38.1 KiB | [pgnodemx_15-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_15-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_15` | 1.7 | `el8.aarch64` | pgdg | 42.2 KiB | [pgnodemx_15-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgnodemx_15-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_15` | 1.7 | `el9.x86_64` | pigsty | 38.0 KiB | [pgnodemx_15-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_15-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el9.x86_64` | pgdg | 43.9 KiB | [pgnodemx_15-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgnodemx_15-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el9.aarch64` | pigsty | 37.5 KiB | [pgnodemx_15-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_15-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_15` | 1.7 | `el9.aarch64` | pgdg | 43.4 KiB | [pgnodemx_15-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgnodemx_15-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_15` | 1.7 | `el10.x86_64` | pigsty | 38.0 KiB | [pgnodemx_15-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_15-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el10.x86_64` | pgdg | 44.1 KiB | [pgnodemx_15-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgnodemx_15-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_15` | 1.7 | `el10.aarch64` | pigsty | 37.7 KiB | [pgnodemx_15-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_15-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_15` | 1.7 | `el10.aarch64` | pgdg | 44.2 KiB | [pgnodemx_15-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgnodemx_15-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgnodemx` | 1.7 | `d12.x86_64` | pigsty | 92.4 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 83.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `d12.aarch64` | pigsty | 91.2 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 82.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 83.7 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 83.0 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u22.x86_64` | pigsty | 98.1 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 90.6 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u22.aarch64` | pigsty | 97.8 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 89.4 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u24.x86_64` | pigsty | 89.3 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 83.5 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u24.aarch64` | pigsty | 89.4 KiB | [postgresql-15-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 82.4 KiB | [postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-15-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_14` | 1.7 | `el8.x86_64` | pigsty | 38.6 KiB | [pgnodemx_14-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_14-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el8.x86_64` | pgdg | 43.0 KiB | [pgnodemx_14-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgnodemx_14-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el8.aarch64` | pigsty | 38.0 KiB | [pgnodemx_14-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_14-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_14` | 1.7 | `el8.aarch64` | pgdg | 42.2 KiB | [pgnodemx_14-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgnodemx_14-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_14` | 1.7 | `el9.x86_64` | pigsty | 37.9 KiB | [pgnodemx_14-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_14-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el9.x86_64` | pgdg | 43.8 KiB | [pgnodemx_14-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgnodemx_14-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el9.aarch64` | pigsty | 37.4 KiB | [pgnodemx_14-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_14-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_14` | 1.7 | `el9.aarch64` | pgdg | 43.4 KiB | [pgnodemx_14-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgnodemx_14-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_14` | 1.7 | `el10.x86_64` | pigsty | 37.9 KiB | [pgnodemx_14-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_14-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el10.x86_64` | pgdg | 44.2 KiB | [pgnodemx_14-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgnodemx_14-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_14` | 1.7 | `el10.aarch64` | pigsty | 37.5 KiB | [pgnodemx_14-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_14-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_14` | 1.7 | `el10.aarch64` | pgdg | 44.0 KiB | [pgnodemx_14-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgnodemx_14-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgnodemx` | 1.7 | `d12.x86_64` | pigsty | 92.0 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 83.3 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `d12.aarch64` | pigsty | 90.7 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 82.3 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 83.5 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 82.7 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u22.x86_64` | pigsty | 97.8 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 90.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u22.aarch64` | pigsty | 97.4 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 89.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u24.x86_64` | pigsty | 88.9 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 83.0 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u24.aarch64` | pigsty | 88.9 KiB | [postgresql-14-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 82.1 KiB | [postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-14-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgnodemx_13` | 1.7 | `el8.x86_64` | pigsty | 38.3 KiB | [pgnodemx_13-1.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgnodemx_13-1.7-1PIGSTY.el8.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el8.x86_64` | pgdg | 42.7 KiB | [pgnodemx_13-1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgnodemx_13-1.7-1PGDG.rhel8.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el8.aarch64` | pigsty | 38.0 KiB | [pgnodemx_13-1.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgnodemx_13-1.7-1PIGSTY.el8.aarch64.rpm) |
| `pgnodemx_13` | 1.7 | `el8.aarch64` | pgdg | 42.2 KiB | [pgnodemx_13-1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgnodemx_13-1.7-1PGDG.rhel8.aarch64.rpm) |
| `pgnodemx_13` | 1.7 | `el9.x86_64` | pigsty | 37.8 KiB | [pgnodemx_13-1.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgnodemx_13-1.7-1PIGSTY.el9.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el9.x86_64` | pgdg | 44.1 KiB | [pgnodemx_13-1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgnodemx_13-1.7-1PGDG.rhel9.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el9.aarch64` | pigsty | 37.4 KiB | [pgnodemx_13-1.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgnodemx_13-1.7-1PIGSTY.el9.aarch64.rpm) |
| `pgnodemx_13` | 1.7 | `el9.aarch64` | pgdg | 43.4 KiB | [pgnodemx_13-1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgnodemx_13-1.7-1PGDG.rhel9.aarch64.rpm) |
| `pgnodemx_13` | 1.7 | `el10.x86_64` | pigsty | 37.9 KiB | [pgnodemx_13-1.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgnodemx_13-1.7-1PIGSTY.el10.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el10.x86_64` | pgdg | 44.4 KiB | [pgnodemx_13-1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgnodemx_13-1.7-1PGDG.rhel10.x86_64.rpm) |
| `pgnodemx_13` | 1.7 | `el10.aarch64` | pigsty | 37.5 KiB | [pgnodemx_13-1.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgnodemx_13-1.7-1PIGSTY.el10.aarch64.rpm) |
| `pgnodemx_13` | 1.7 | `el10.aarch64` | pgdg | 44.0 KiB | [pgnodemx_13-1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgnodemx_13-1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgnodemx` | 1.7 | `d12.x86_64` | pigsty | 92.1 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `d12.x86_64` | pgdg | 83.7 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `d12.aarch64` | pigsty | 90.6 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `d12.aarch64` | pgdg | 82.0 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `d13.x86_64` | pgdg | 83.5 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `d13.aarch64` | pgdg | 82.5 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u22.x86_64` | pigsty | 97.8 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u22.x86_64` | pgdg | 90.4 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u22.aarch64` | pigsty | 97.3 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u22.aarch64` | pgdg | 88.9 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u24.x86_64` | pigsty | 89.3 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u24.x86_64` | pgdg | 83.3 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u24.aarch64` | pigsty | 88.8 KiB | [postgresql-13-pgnodemx_1.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pgnodemx` | 1.7 | `u24.aarch64` | pgdg | 82.0 KiB | [postgresql-13-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgnodemx/postgresql-13-pgnodemx_1.7-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgnodemx" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgnodemx" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgnodemx-1.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgnodemx; # get pgnodemx source code
pig build dep pgnodemx; # install build dependencies
pig build pkg pgnodemx; # build extension rpm or deb
pig build ext pgnodemx; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgnodemx; # install by extension name, for the current active PG version
pig ext install pgnodemx; # install via package alias, for the active PG version
pig ext install pgnodemx -v 18;   # install for PG 18
pig ext install pgnodemx -v 17;   # install for PG 17
pig ext install pgnodemx -v 16;   # install for PG 16
pig ext install pgnodemx -v 15;   # install for PG 15
pig ext install pgnodemx -v 14;   # install for PG 14
pig ext install pgnodemx -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgnodemx;
```

