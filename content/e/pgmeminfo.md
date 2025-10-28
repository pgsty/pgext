---
title: "pgmeminfo"
linkTitle: "pgmeminfo"
description: "show memory usage"
weight: 6350
categories: ["STAT"]
width: full
---

show memory usage


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6350** | {{< badge content="pgmeminfo" link="https://github.com/okbob/pgmeminfo" >}} | {{< ext "pgmeminfo" >}} | `1.0.0` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "system_stats" >}} {{< ext "pg_buffercache" >}} {{< ext "pgnodemx" >}} {{< ext "pg_proctab" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} {{< ext "pg_prewarm" >}} |

> [!Note] no pg14,13,12 on el8/9 pgdg repo


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pgmeminfo" >}} | `1.0.0` | {{< bg "18" "pgmeminfo_18*" "green" >}} {{< bg "17" "pgmeminfo_17*" "green" >}} {{< bg "16" "pgmeminfo_16*" "green" >}} {{< bg "15" "pgmeminfo_15*" "green" >}} {{< bg "14" "pgmeminfo_14*" "green" >}} {{< bg "13" "pgmeminfo_13*" "green" >}} | `pgmeminfo_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pgmeminfo" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pgmeminfo" "red" >}} {{< bg "17" "postgresql-17-pgmeminfo" "green" >}} {{< bg "16" "postgresql-16-pgmeminfo" "green" >}} {{< bg "15" "postgresql-15-pgmeminfo" "green" >}} {{< bg "14" "postgresql-14-pgmeminfo" "green" >}} {{< bg "13" "postgresql-13-pgmeminfo" "green" >}} | `postgresql-$v-pgmeminfo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_13 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_13 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_13 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 2" "blue" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.0" "pgmeminfo_13 : AVAIL 2" "green" >}} |
|    `el10.x86_64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_13 : AVAIL 1" "blue" >}} |
|    `el10.aarch64`    | {{< bg "PGDG 1.0.0" "pgmeminfo_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pgmeminfo_13 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgmeminfo : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgmeminfo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pgmeminfo : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pgmeminfo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pgmeminfo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-13-pgmeminfo : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_18` | 1.0.0 | `el8.x86_64` | pgdg | 15.1 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgmeminfo_18-1.0.0-3PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_18` | 1.0.0 | `el8.aarch64` | pgdg | 15.0 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgmeminfo_18-1.0.0-3PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_18` | 1.0.0 | `el9.x86_64` | pgdg | 15.4 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgmeminfo_18-1.0.0-3PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_18` | 1.0.0 | `el9.aarch64` | pgdg | 15.0 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgmeminfo_18-1.0.0-3PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_18` | 1.0.0 | `el10.x86_64` | pgdg | 15.8 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgmeminfo_18-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_18` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_18-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgmeminfo_18-1.0.0-3PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_17` | 1.0.0 | `el8.x86_64` | pigsty | 15.0 KiB | [pgmeminfo_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmeminfo_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el8.x86_64` | pgdg | 15.0 KiB | [pgmeminfo_17-1.0.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgmeminfo_17-1.0.0-2PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el8.aarch64` | pigsty | 14.9 KiB | [pgmeminfo_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmeminfo_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el8.aarch64` | pgdg | 15.0 KiB | [pgmeminfo_17-1.0.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgmeminfo_17-1.0.0-2PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el9.x86_64` | pigsty | 15.2 KiB | [pgmeminfo_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmeminfo_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el9.x86_64` | pgdg | 15.3 KiB | [pgmeminfo_17-1.0.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgmeminfo_17-1.0.0-2PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el9.aarch64` | pigsty | 15.0 KiB | [pgmeminfo_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmeminfo_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el9.aarch64` | pgdg | 15.1 KiB | [pgmeminfo_17-1.0.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgmeminfo_17-1.0.0-2PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el10.x86_64` | pgdg | 15.8 KiB | [pgmeminfo_17-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgmeminfo_17-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_17` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_17-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgmeminfo_17-1.0.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `d12.x86_64` | pigsty | 16.2 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `d12.aarch64` | pigsty | 16.2 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `u22.x86_64` | pigsty | 16.9 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `u22.aarch64` | pigsty | 16.9 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `u24.x86_64` | pigsty | 14.9 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmeminfo` | 1.0.0 | `u24.aarch64` | pigsty | 14.9 KiB | [postgresql-17-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-17-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_16` | 1.0.0 | `el8.x86_64` | pigsty | 14.9 KiB | [pgmeminfo_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmeminfo_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el8.x86_64` | pgdg | 15.0 KiB | [pgmeminfo_16-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgmeminfo_16-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el8.aarch64` | pigsty | 14.9 KiB | [pgmeminfo_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmeminfo_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el8.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgmeminfo_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el9.x86_64` | pigsty | 15.2 KiB | [pgmeminfo_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmeminfo_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el9.x86_64` | pgdg | 15.1 KiB | [pgmeminfo_16-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgmeminfo_16-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el9.aarch64` | pigsty | 15.0 KiB | [pgmeminfo_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmeminfo_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgmeminfo_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el10.x86_64` | pgdg | 15.8 KiB | [pgmeminfo_16-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgmeminfo_16-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_16` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_16-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgmeminfo_16-1.0.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `d12.x86_64` | pigsty | 16.2 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `d12.aarch64` | pigsty | 16.1 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `u22.x86_64` | pigsty | 16.8 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `u22.aarch64` | pigsty | 16.8 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `u24.x86_64` | pigsty | 14.9 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmeminfo` | 1.0.0 | `u24.aarch64` | pigsty | 14.9 KiB | [postgresql-16-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-16-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_15` | 1.0.0 | `el8.x86_64` | pigsty | 14.9 KiB | [pgmeminfo_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmeminfo_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el8.x86_64` | pgdg | 15.0 KiB | [pgmeminfo_15-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgmeminfo_15-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el8.aarch64` | pigsty | 14.9 KiB | [pgmeminfo_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmeminfo_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el8.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_15-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgmeminfo_15-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el9.x86_64` | pigsty | 15.2 KiB | [pgmeminfo_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmeminfo_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el9.x86_64` | pgdg | 15.1 KiB | [pgmeminfo_15-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgmeminfo_15-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el9.aarch64` | pigsty | 15.0 KiB | [pgmeminfo_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmeminfo_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_15-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgmeminfo_15-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el10.x86_64` | pgdg | 15.8 KiB | [pgmeminfo_15-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgmeminfo_15-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_15` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_15-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgmeminfo_15-1.0.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `d12.x86_64` | pigsty | 16.2 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `d12.aarch64` | pigsty | 16.2 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `u22.x86_64` | pigsty | 16.9 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `u22.aarch64` | pigsty | 16.8 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `u24.x86_64` | pigsty | 14.9 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmeminfo` | 1.0.0 | `u24.aarch64` | pigsty | 14.9 KiB | [postgresql-15-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-15-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_14` | 1.0.0 | `el8.x86_64` | pigsty | 14.9 KiB | [pgmeminfo_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmeminfo_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el8.x86_64` | pgdg | 14.9 KiB | [pgmeminfo_14-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgmeminfo_14-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el8.aarch64` | pigsty | 14.9 KiB | [pgmeminfo_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmeminfo_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el8.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_14-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgmeminfo_14-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el9.x86_64` | pigsty | 15.2 KiB | [pgmeminfo_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmeminfo_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el9.x86_64` | pgdg | 15.1 KiB | [pgmeminfo_14-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgmeminfo_14-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el9.aarch64` | pigsty | 15.0 KiB | [pgmeminfo_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmeminfo_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_14-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgmeminfo_14-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el10.x86_64` | pgdg | 15.7 KiB | [pgmeminfo_14-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgmeminfo_14-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_14` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_14-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgmeminfo_14-1.0.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `d12.x86_64` | pigsty | 16.1 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `d12.aarch64` | pigsty | 16.1 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `u22.x86_64` | pigsty | 16.8 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `u22.aarch64` | pigsty | 16.8 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `u24.x86_64` | pigsty | 14.8 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmeminfo` | 1.0.0 | `u24.aarch64` | pigsty | 14.8 KiB | [postgresql-14-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-14-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmeminfo_13` | 1.0.0 | `el8.x86_64` | pigsty | 14.9 KiB | [pgmeminfo_13-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmeminfo_13-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el8.x86_64` | pgdg | 14.9 KiB | [pgmeminfo_13-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgmeminfo_13-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el8.aarch64` | pigsty | 14.9 KiB | [pgmeminfo_13-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmeminfo_13-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el8.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_13-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgmeminfo_13-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el9.x86_64` | pigsty | 15.2 KiB | [pgmeminfo_13-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmeminfo_13-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el9.x86_64` | pgdg | 15.1 KiB | [pgmeminfo_13-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgmeminfo_13-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el9.aarch64` | pigsty | 15.0 KiB | [pgmeminfo_13-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmeminfo_13-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el9.aarch64` | pgdg | 14.9 KiB | [pgmeminfo_13-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgmeminfo_13-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el10.x86_64` | pgdg | 15.8 KiB | [pgmeminfo_13-1.0.0-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgmeminfo_13-1.0.0-3PGDG.rhel10.x86_64.rpm) |
| `pgmeminfo_13` | 1.0.0 | `el10.aarch64` | pgdg | 15.7 KiB | [pgmeminfo_13-1.0.0-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgmeminfo_13-1.0.0-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `d12.x86_64` | pigsty | 15.9 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `d12.aarch64` | pigsty | 15.7 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `u22.x86_64` | pigsty | 16.8 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `u22.aarch64` | pigsty | 16.4 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `u24.x86_64` | pigsty | 14.9 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pgmeminfo` | 1.0.0 | `u24.aarch64` | pigsty | 14.8 KiB | [postgresql-13-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmeminfo/postgresql-13-pgmeminfo_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/pgmeminfo" title="Repository" icon="github" subtitle="github.com/okbob/pgmeminfo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmeminfo-VERSION_1_0_0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pgmeminfo; # get pgmeminfo source code
pig build dep pgmeminfo; # install build dependencies
pig build pkg pgmeminfo; # build extension rpm or deb
pig build ext pgmeminfo; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgmeminfo; # install by extension name, for the current active PG version
pig ext install pgmeminfo; # install via package alias, for the active PG version
pig ext install pgmeminfo -v 18;   # install for PG 18
pig ext install pgmeminfo -v 17;   # install for PG 17
pig ext install pgmeminfo -v 16;   # install for PG 16
pig ext install pgmeminfo -v 15;   # install for PG 15
pig ext install pgmeminfo -v 14;   # install for PG 14
pig ext install pgmeminfo -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgmeminfo;
```

