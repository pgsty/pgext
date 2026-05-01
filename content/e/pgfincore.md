---
title: "pgfincore"
linkTitle: "pgfincore"
description: "examine and manage the os buffer cache"
weight: 5060
categories: ["ADMIN"]
width: full
---

[**pgfincore**](https://github.com/klando/pgfincore) : examine and manage the os buffer cache


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5060** | {{< badge content="pgfincore" link="https://github.com/klando/pgfincore" >}} | {{< ext "pgfincore" >}} | `1.3.1` | {{< category "ADMIN" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} {{< ext "fio" >}} {{< ext "pg_prewarm" >}} {{< ext "pgmeminfo" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |

> [!Note] pg18 el fixed by vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgfincore` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.1` | {{< bg "18" "pgfincore_18" "green" >}} {{< bg "17" "pgfincore_17" "green" >}} {{< bg "16" "pgfincore_16" "green" >}} {{< bg "15" "pgfincore_15" "green" >}} {{< bg "14" "pgfincore_14" "green" >}} | `pgfincore_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.1` | {{< bg "18" "postgresql-18-pgfincore" "green" >}} {{< bg "17" "postgresql-17-pgfincore" "green" >}} {{< bg "16" "postgresql-16-pgfincore" "green" >}} {{< bg "15" "postgresql-15-pgfincore" "green" >}} {{< bg "14" "postgresql-14-pgfincore" "green" >}} | `postgresql-$v-pgfincore` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.4" "pgfincore_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.4" "pgfincore_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.4" "pgfincore_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.4" "pgfincore_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "pgfincore_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.1" "pgfincore_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "pgfincore_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.3.1" "postgresql-18-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-17-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-16-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-15-pgfincore : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.1" "postgresql-14-pgfincore : AVAIL 2" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgfincore_18` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.6 KiB | [pgfincore_18-1.3.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgfincore_18-1.3.1-1PIGSTY.el8.x86_64.rpm) |
| `pgfincore_18` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pgfincore_18-1.3.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgfincore_18-1.3.1-1PIGSTY.el8.aarch64.rpm) |
| `pgfincore_18` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.7 KiB | [pgfincore_18-1.3.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgfincore_18-1.3.1-1PIGSTY.el9.x86_64.rpm) |
| `pgfincore_18` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.3 KiB | [pgfincore_18-1.3.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgfincore_18-1.3.1-1PIGSTY.el9.aarch64.rpm) |
| `pgfincore_18` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.8 KiB | [pgfincore_18-1.3.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgfincore_18-1.3.1-1PIGSTY.el10.x86_64.rpm) |
| `pgfincore_18` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.6 KiB | [pgfincore_18-1.3.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgfincore_18-1.3.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgfincore` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.8 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg12+1_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.4 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg12+1_arm64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 29.0 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg13+1_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.4 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg13+1_arm64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 27.6 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 27.1 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.4 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.7 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 26.9 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.6 KiB | [postgresql-18-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 26.2 KiB | [postgresql-18-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb) |
| `postgresql-18-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.2 KiB | [postgresql-18-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-18-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgfincore_17` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.5 KiB | [pgfincore_17-1.3.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgfincore_17-1.3.1-3PGDG.rhel8.x86_64.rpm) |
| `pgfincore_17` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.2 KiB | [pgfincore_17-1.3.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgfincore_17-1.3.1-3PGDG.rhel8.aarch64.rpm) |
| `pgfincore_17` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.8 KiB | [pgfincore_17-1.3.1-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgfincore_17-1.3.1-3PGDG.rhel9.x86_64.rpm) |
| `pgfincore_17` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.2 KiB | [pgfincore_17-1.3.1-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgfincore_17-1.3.1-3PGDG.rhel9.aarch64.rpm) |
| `pgfincore_17` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.1 KiB | [pgfincore_17-1.3.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgfincore_17-1.3.1-4PGDG.rhel10.x86_64.rpm) |
| `pgfincore_17` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgfincore_17-1.3.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgfincore_17-1.3.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgfincore` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.6 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg12+1_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.0 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg12+1_arm64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.7 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg13+1_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.0 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg13+1_arm64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 31.7 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 31.1 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.1 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.3 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 26.7 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.4 KiB | [postgresql-17-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 26.2 KiB | [postgresql-17-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb) |
| `postgresql-17-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.3 KiB | [postgresql-17-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-17-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgfincore_16` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgfincore_16-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgfincore_16-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_16` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgfincore_16-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgfincore_16-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_16` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.5 KiB | [pgfincore_16-1.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgfincore_16-1.3.1-1PGDG.rhel9.x86_64.rpm) |
| `pgfincore_16` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 22.8 KiB | [pgfincore_16-1.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgfincore_16-1.3.1-1PGDG.rhel9.aarch64.rpm) |
| `pgfincore_16` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.1 KiB | [pgfincore_16-1.3.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgfincore_16-1.3.1-4PGDG.rhel10.x86_64.rpm) |
| `pgfincore_16` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgfincore_16-1.3.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgfincore_16-1.3.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgfincore` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.5 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg12+1_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.0 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg12+1_arm64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.7 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg13+1_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.0 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg13+1_arm64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 31.3 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 30.7 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.1 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.3 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 26.7 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 26.2 KiB | [postgresql-16-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb) |
| `postgresql-16-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.3 KiB | [postgresql-16-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-16-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgfincore_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgfincore_15-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfincore_15-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_15` | `1.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgfincore_15-1.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgfincore_15-1.2.4-1.rhel8.x86_64.rpm) |
| `pgfincore_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgfincore_15-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfincore_15-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_15` | `1.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.7 KiB | [pgfincore_15-1.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgfincore_15-1.2.4-1.rhel8.aarch64.rpm) |
| `pgfincore_15` | `1.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pgfincore_15-1.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgfincore_15-1.2.4-1.rhel9.x86_64.rpm) |
| `pgfincore_15` | `1.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgfincore_15-1.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgfincore_15-1.2.4-1.rhel9.aarch64.rpm) |
| `pgfincore_15` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.1 KiB | [pgfincore_15-1.3.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgfincore_15-1.3.1-4PGDG.rhel10.x86_64.rpm) |
| `pgfincore_15` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgfincore_15-1.3.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgfincore_15-1.3.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgfincore` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.6 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg12+1_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 28.0 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg12+1_arm64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.7 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg13+1_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.0 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg13+1_arm64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 31.3 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 30.7 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.1 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.3 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 26.7 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.4 KiB | [postgresql-15-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 26.3 KiB | [postgresql-15-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb) |
| `postgresql-15-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.3 KiB | [postgresql-15-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-15-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgfincore_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.3 KiB | [pgfincore_14-1.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.3.1-1PGDG.rhel8.x86_64.rpm) |
| `pgfincore_14` | `1.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 24.0 KiB | [pgfincore_14-1.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.2.4-1.rhel8.x86_64.rpm) |
| `pgfincore_14` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 41.0 KiB | [pgfincore_14-1.2.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgfincore_14-1.2.2-3.rhel8.x86_64.rpm) |
| `pgfincore_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 24.0 KiB | [pgfincore_14-1.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfincore_14-1.3.1-1PGDG.rhel8.aarch64.rpm) |
| `pgfincore_14` | `1.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 23.7 KiB | [pgfincore_14-1.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgfincore_14-1.2.4-1.rhel8.aarch64.rpm) |
| `pgfincore_14` | `1.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 23.7 KiB | [pgfincore_14-1.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgfincore_14-1.2.4-1.rhel9.x86_64.rpm) |
| `pgfincore_14` | `1.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 23.1 KiB | [pgfincore_14-1.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgfincore_14-1.2.4-1.rhel9.aarch64.rpm) |
| `pgfincore_14` | `1.3.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 24.1 KiB | [pgfincore_14-1.3.1-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgfincore_14-1.3.1-4PGDG.rhel10.x86_64.rpm) |
| `pgfincore_14` | `1.3.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 23.7 KiB | [pgfincore_14-1.3.1-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgfincore_14-1.3.1-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgfincore` | `1.3.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 28.6 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg12+1_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 27.9 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg12+1_arm64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 28.7 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg13+1_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 28.0 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg13+1_arm64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 31.3 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 30.7 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 27.1 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 26.3 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 26.6 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 26.4 KiB | [postgresql-14-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 26.2 KiB | [postgresql-14-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-3.pgdg26.04+1_arm64.deb) |
| `postgresql-14-pgfincore` | `1.3.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.3 KiB | [postgresql-14-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgfincore/postgresql-14-pgfincore_1.3.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/klando/pgfincore" title="Repository" icon="github" subtitle="github.com/klando/pgfincore" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgfincore-1.3.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgfincore;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgfincore;		# install via package name, for the active PG version

pig install pgfincore -v 18;   # install for PG 18
pig install pgfincore -v 17;   # install for PG 17
pig install pgfincore -v 16;   # install for PG 16
pig install pgfincore -v 15;   # install for PG 15
pig install pgfincore -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgfincore;
```




## Usage

> [pgfincore: examine and manage the os buffer cache](https://github.com/klando/pgfincore)

pgfincore provides functions to inspect and manage OS page cache contents for PostgreSQL relations using `mincore` and `POSIX_FADVISE`.

### Inspect Cache State

```sql
SELECT * FROM pgfincore('pgbench_accounts');
```

Returns per-segment info: `relpath`, `segment`, `os_page_size`, `rel_os_pages`, `pages_mem` (pages in OS cache), `group_mem`, `os_pages_free`, `pages_dirty`, `group_dirty`.

Use `pgfincore('relation', true)` to include the `databit` varbit map for snapshot/restore.

### System Info

```sql
SELECT * FROM pgsysconf();          -- os_page_size, os_pages_free, os_total_pages
SELECT * FROM pgsysconf_pretty();   -- same with human-readable output
```

### Preload into OS Cache

```sql
SELECT * FROM pgfadvise_willneed('pgbench_accounts');
```

### Evict from OS Cache

```sql
SELECT * FROM pgfadvise_dontneed('pgbench_accounts');
```

### Other POSIX_FADVISE Flags

```sql
SELECT * FROM pgfadvise_normal('relation');
SELECT * FROM pgfadvise_sequential('relation');
SELECT * FROM pgfadvise_random('relation');
```

### Snapshot and Restore Cache State

```sql
-- Snapshot
CREATE TABLE pgfincore_snapshot AS
  SELECT 'pgbench_accounts'::text AS relname, *, now() AS date_snapshot
  FROM pgfincore('pgbench_accounts', true);

-- Restore
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true,
               (SELECT databit FROM pgfincore_snapshot
                WHERE relname = 'pgbench_accounts' AND segment = 0));
```

### Direct Page Cache Control

```sql
-- Load first 3 pages, unload next 3
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true, B'111000');
-- Load only
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, false, B'111000');
-- Unload only
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, false, true, B'111000');
```
