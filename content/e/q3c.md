---
title: "q3c"
linkTitle: "q3c"
description: "q3c sky indexing plugin"
weight: 1540
categories: ["GIS"]
width: full
---

q3c sky indexing plugin


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1540** | {{< badge content="q3c" link="https://github.com/segasai/q3c" >}} | {{< ext "q3c" >}} | `2.0.1` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "earthdistance" >}} {{< ext "pg_sphere" >}} {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/q3c" >}} | `2.0.1` | {{< bg "18" "q3c_18*" "green" >}} {{< bg "17" "q3c_17*" "green" >}} {{< bg "16" "q3c_16*" "green" >}} {{< bg "15" "q3c_15*" "green" >}} {{< bg "14" "q3c_14*" "green" >}} | `q3c_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/q3c" >}} | `2.0.1` | {{< bg "18" "postgresql-18-q3c" "green" >}} {{< bg "17" "postgresql-17-q3c" "green" >}} {{< bg "16" "postgresql-16-q3c" "green" >}} {{< bg "15" "postgresql-15-q3c" "green" >}} {{< bg "14" "postgresql-14-q3c" "green" >}} | `postgresql-$v-q3c` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.0.1" "q3c_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_18` | 2.0.1 | `el8.x86_64` | pigsty | 96.6 KiB | [q3c_18-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_18-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_18` | 2.0.1 | `el8.x86_64` | pgdg | 103.5 KiB | [q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_18` | 2.0.1 | `el8.aarch64` | pgdg | 97.8 KiB | [q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_18` | 2.0.1 | `el9.x86_64` | pigsty | 131.5 KiB | [q3c_18-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_18-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_18` | 2.0.1 | `el9.x86_64` | pgdg | 108.7 KiB | [q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_18` | 2.0.1 | `el9.aarch64` | pigsty | 125.8 KiB | [q3c_18-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_18-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_18` | 2.0.1 | `el9.aarch64` | pgdg | 105.3 KiB | [q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-q3c` | 2.0.1 | `d12.x86_64` | pgdg | 140.8 KiB | [postgresql-18-q3c_2.0.1-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg12+1_amd64.deb) |
| `postgresql-18-q3c` | 2.0.1 | `d12.aarch64` | pgdg | 157.1 KiB | [postgresql-18-q3c_2.0.1-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg12+1_arm64.deb) |
| `postgresql-18-q3c` | 2.0.1 | `u22.x86_64` | pgdg | 153.8 KiB | [postgresql-18-q3c_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-18-q3c` | 2.0.1 | `u22.aarch64` | pgdg | 124.0 KiB | [postgresql-18-q3c_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-18-q3c` | 2.0.1 | `u24.x86_64` | pgdg | 155.3 KiB | [postgresql-18-q3c_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-q3c` | 2.0.1 | `u24.aarch64` | pgdg | 124.4 KiB | [postgresql-18-q3c_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_17` | 2.0.1 | `el8.x86_64` | pigsty | 96.6 KiB | [q3c_17-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_17-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_17` | 2.0.1 | `el8.x86_64` | pgdg | 103.5 KiB | [q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_17` | 2.0.1 | `el8.aarch64` | pigsty | 90.8 KiB | [q3c_17-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_17-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_17` | 2.0.1 | `el8.aarch64` | pgdg | 97.8 KiB | [q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_17` | 2.0.1 | `el9.x86_64` | pigsty | 131.5 KiB | [q3c_17-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_17-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_17` | 2.0.1 | `el9.x86_64` | pgdg | 101.8 KiB | [q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_17` | 2.0.1 | `el9.aarch64` | pigsty | 125.9 KiB | [q3c_17-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_17-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_17` | 2.0.1 | `el9.aarch64` | pgdg | 105.4 KiB | [q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-q3c` | 2.0.1 | `d12.x86_64` | pgdg | 150.9 KiB | [postgresql-17-q3c_2.0.1-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg12+1_amd64.deb) |
| `postgresql-17-q3c` | 2.0.1 | `d12.aarch64` | pgdg | 159.7 KiB | [postgresql-17-q3c_2.0.1-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg12+1_arm64.deb) |
| `postgresql-17-q3c` | 2.0.1 | `u22.x86_64` | pgdg | 153.6 KiB | [postgresql-17-q3c_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-17-q3c` | 2.0.1 | `u22.aarch64` | pgdg | 165.4 KiB | [postgresql-17-q3c_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-17-q3c` | 2.0.1 | `u24.x86_64` | pgdg | 142.7 KiB | [postgresql-17-q3c_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-q3c` | 2.0.1 | `u24.aarch64` | pgdg | 144.2 KiB | [postgresql-17-q3c_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_16` | 2.0.1 | `el8.x86_64` | pigsty | 96.6 KiB | [q3c_16-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_16-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_16` | 2.0.1 | `el8.x86_64` | pgdg | 103.5 KiB | [q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_16` | 2.0.1 | `el8.aarch64` | pigsty | 90.8 KiB | [q3c_16-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_16-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_16` | 2.0.1 | `el8.aarch64` | pgdg | 97.8 KiB | [q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_16` | 2.0.1 | `el9.x86_64` | pigsty | 131.6 KiB | [q3c_16-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_16-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_16` | 2.0.1 | `el9.x86_64` | pgdg | 103.3 KiB | [q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_16` | 2.0.1 | `el9.aarch64` | pigsty | 125.9 KiB | [q3c_16-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_16-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_16` | 2.0.1 | `el9.aarch64` | pgdg | 105.4 KiB | [q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-q3c` | 2.0.1 | `d12.x86_64` | pgdg | 134.7 KiB | [postgresql-16-q3c_2.0.1-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg12+1_amd64.deb) |
| `postgresql-16-q3c` | 2.0.1 | `d12.aarch64` | pgdg | 160.4 KiB | [postgresql-16-q3c_2.0.1-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg12+1_arm64.deb) |
| `postgresql-16-q3c` | 2.0.1 | `u22.x86_64` | pgdg | 160.8 KiB | [postgresql-16-q3c_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-16-q3c` | 2.0.1 | `u22.aarch64` | pgdg | 154.0 KiB | [postgresql-16-q3c_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-16-q3c` | 2.0.1 | `u24.x86_64` | pgdg | 149.6 KiB | [postgresql-16-q3c_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-q3c` | 2.0.1 | `u24.aarch64` | pgdg | 144.3 KiB | [postgresql-16-q3c_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_15` | 2.0.1 | `el8.x86_64` | pigsty | 95.4 KiB | [q3c_15-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_15-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_15` | 2.0.1 | `el8.x86_64` | pgdg | 102.3 KiB | [q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_15` | 2.0.1 | `el8.aarch64` | pigsty | 90.3 KiB | [q3c_15-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_15-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_15` | 2.0.1 | `el8.aarch64` | pgdg | 97.2 KiB | [q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_15` | 2.0.1 | `el9.x86_64` | pigsty | 159.6 KiB | [q3c_15-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_15-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_15` | 2.0.1 | `el9.x86_64` | pgdg | 109.3 KiB | [q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_15` | 2.0.1 | `el9.aarch64` | pigsty | 127.2 KiB | [q3c_15-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_15-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_15` | 2.0.1 | `el9.aarch64` | pgdg | 103.5 KiB | [q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-q3c` | 2.0.1 | `d12.x86_64` | pgdg | 139.2 KiB | [postgresql-15-q3c_2.0.1-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg12+1_amd64.deb) |
| `postgresql-15-q3c` | 2.0.1 | `d12.aarch64` | pgdg | 158.9 KiB | [postgresql-15-q3c_2.0.1-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg12+1_arm64.deb) |
| `postgresql-15-q3c` | 2.0.1 | `u22.x86_64` | pgdg | 153.7 KiB | [postgresql-15-q3c_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-15-q3c` | 2.0.1 | `u22.aarch64` | pgdg | 164.1 KiB | [postgresql-15-q3c_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-15-q3c` | 2.0.1 | `u24.x86_64` | pgdg | 132.1 KiB | [postgresql-15-q3c_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-q3c` | 2.0.1 | `u24.aarch64` | pgdg | 146.3 KiB | [postgresql-15-q3c_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_14` | 2.0.1 | `el8.x86_64` | pigsty | 95.4 KiB | [q3c_14-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_14-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_14` | 2.0.1 | `el8.x86_64` | pgdg | 102.3 KiB | [q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_14` | 2.0.1 | `el8.aarch64` | pigsty | 90.3 KiB | [q3c_14-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_14-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_14` | 2.0.1 | `el8.aarch64` | pgdg | 97.2 KiB | [q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_14` | 2.0.1 | `el9.x86_64` | pigsty | 132.7 KiB | [q3c_14-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_14-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_14` | 2.0.1 | `el9.x86_64` | pgdg | 109.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_14` | 2.0.1 | `el9.aarch64` | pigsty | 106.2 KiB | [q3c_14-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_14-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_14` | 2.0.1 | `el9.aarch64` | pgdg | 103.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-14-q3c` | 2.0.1 | `d12.x86_64` | pgdg | 143.4 KiB | [postgresql-14-q3c_2.0.1-5.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg12+1_amd64.deb) |
| `postgresql-14-q3c` | 2.0.1 | `d12.aarch64` | pgdg | 158.3 KiB | [postgresql-14-q3c_2.0.1-5.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg12+1_arm64.deb) |
| `postgresql-14-q3c` | 2.0.1 | `u22.x86_64` | pgdg | 163.5 KiB | [postgresql-14-q3c_2.0.1-5.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg22.04+1_amd64.deb) |
| `postgresql-14-q3c` | 2.0.1 | `u22.aarch64` | pgdg | 146.2 KiB | [postgresql-14-q3c_2.0.1-5.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg22.04+1_arm64.deb) |
| `postgresql-14-q3c` | 2.0.1 | `u24.x86_64` | pgdg | 150.9 KiB | [postgresql-14-q3c_2.0.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-q3c` | 2.0.1 | `u24.aarch64` | pgdg | 155.4 KiB | [postgresql-14-q3c_2.0.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/segasai/q3c" title="Repository" icon="github" subtitle="github.com/segasai/q3c" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="q3c-2.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get q3c; # get q3c source code
pig build dep q3c; # install build dependencies
pig build pkg q3c; # build extension rpm or deb
pig build ext q3c; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install q3c; # install by extension name, for the current active PG version
pig ext install q3c; # install via package alias, for the active PG version
pig ext install q3c -v 18;   # install for PG 18
pig ext install q3c -v 17;   # install for PG 17
pig ext install q3c -v 16;   # install for PG 16
pig ext install q3c -v 15;   # install for PG 15
pig ext install q3c -v 14;   # install for PG 14
pig ext install q3c -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION q3c;
```

