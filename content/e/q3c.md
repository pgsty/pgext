---
title: "q3c"
linkTitle: "q3c"
description: "q3c sky indexing plugin"
weight: 1540
categories: ["GIS"]
width: full
---

[**q3c**](https://github.com/segasai/q3c) : q3c sky indexing plugin


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1540** | {{< badge content="q3c" link="https://github.com/segasai/q3c" >}} | {{< ext "q3c" >}} | `2.0.1` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "earthdistance" >}} {{< ext "pg_sphere" >}} {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `q3c` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.1` | {{< bg "18" "q3c_18*" "green" >}} {{< bg "17" "q3c_17*" "green" >}} {{< bg "16" "q3c_16*" "green" >}} {{< bg "15" "q3c_15*" "green" >}} {{< bg "14" "q3c_14*" "green" >}} {{< bg "13" "q3c_13*" "green" >}} | `q3c_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.1` | {{< bg "18" "postgresql-18-q3c" "green" >}} {{< bg "17" "postgresql-17-q3c" "green" >}} {{< bg "16" "postgresql-16-q3c" "green" >}} {{< bg "15" "postgresql-15-q3c" "green" >}} {{< bg "14" "postgresql-14-q3c" "green" >}} {{< bg "13" "postgresql-13-q3c" "green" >}} | `postgresql-$v-q3c` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.1" "q3c_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 2.0.1" "q3c_13 : AVAIL 2" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.1" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-14-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.1" "postgresql-13-q3c : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.7 KiB | [q3c_18-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_18-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.1 KiB | [q3c_18-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_18-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 132.3 KiB | [q3c_18-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_18-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 126.1 KiB | [q3c_18-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_18-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.5 KiB | [q3c_18-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_18-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_18-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/q3c_18-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.9 KiB | [q3c_18-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_18-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [q3c_18-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/q3c_18-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 127.3 KiB | [postgresql-18-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 151.7 KiB | [postgresql-18-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 152.2 KiB | [postgresql-18-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 154.0 KiB | [postgresql-18-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 158.2 KiB | [postgresql-18-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 163.8 KiB | [postgresql-18-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 155.4 KiB | [postgresql-18-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 158.6 KiB | [postgresql-18-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.7 KiB | [q3c_17-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_17-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.1 KiB | [q3c_17-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_17-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 132.3 KiB | [q3c_17-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_17-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.8 KiB | [q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 126.9 KiB | [q3c_17-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_17-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.4 KiB | [q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.7 KiB | [q3c_17-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_17-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_17-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/q3c_17-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.4 KiB | [q3c_17-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_17-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.7 KiB | [q3c_17-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/q3c_17-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 155.8 KiB | [postgresql-17-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 126.7 KiB | [postgresql-17-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 166.7 KiB | [postgresql-17-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 151.5 KiB | [postgresql-17-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 152.3 KiB | [postgresql-17-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 146.2 KiB | [postgresql-17-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 148.7 KiB | [postgresql-17-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 155.8 KiB | [postgresql-17-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.7 KiB | [q3c_16-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_16-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.1 KiB | [q3c_16-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_16-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 132.3 KiB | [q3c_16-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_16-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.3 KiB | [q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 103.8 KiB | [q3c_16-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_16-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.4 KiB | [q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.7 KiB | [q3c_16-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_16-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_16-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/q3c_16-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.3 KiB | [q3c_16-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_16-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.5 KiB | [q3c_16-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/q3c_16-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 131.0 KiB | [postgresql-16-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 140.2 KiB | [postgresql-16-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 133.2 KiB | [postgresql-16-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 130.9 KiB | [postgresql-16-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 162.3 KiB | [postgresql-16-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 154.9 KiB | [postgresql-16-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 129.8 KiB | [postgresql-16-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 159.4 KiB | [postgresql-16-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 97.6 KiB | [q3c_15-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_15-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.3 KiB | [q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.7 KiB | [q3c_15-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_15-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 160.3 KiB | [q3c_15-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_15-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.3 KiB | [q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 107.3 KiB | [q3c_15-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_15-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [q3c_15-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_15-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.2 KiB | [q3c_15-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/q3c_15-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.1 KiB | [q3c_15-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_15-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.1 KiB | [q3c_15-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/q3c_15-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 148.5 KiB | [postgresql-15-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 130.8 KiB | [postgresql-15-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 155.5 KiB | [postgresql-15-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 160.8 KiB | [postgresql-15-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 163.1 KiB | [postgresql-15-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 129.7 KiB | [postgresql-15-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 157.3 KiB | [postgresql-15-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 158.0 KiB | [postgresql-15-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 97.6 KiB | [q3c_14-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_14-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.3 KiB | [q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.6 KiB | [q3c_14-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_14-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 160.3 KiB | [q3c_14-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_14-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 128.1 KiB | [q3c_14-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_14-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.7 KiB | [q3c_14-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_14-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.1 KiB | [q3c_14-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/q3c_14-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.1 KiB | [q3c_14-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_14-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.0 KiB | [q3c_14-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/q3c_14-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 156.5 KiB | [postgresql-14-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 156.1 KiB | [postgresql-14-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 156.8 KiB | [postgresql-14-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 160.1 KiB | [postgresql-14-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 163.4 KiB | [postgresql-14-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 157.4 KiB | [postgresql-14-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 143.2 KiB | [postgresql-14-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 159.1 KiB | [postgresql-14-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_13` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 97.3 KiB | [q3c_13-2.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_13-2.0.1-1PIGSTY.el8.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.1 KiB | [q3c_13-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/q3c_13-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 92.6 KiB | [q3c_13-2.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_13-2.0.1-1PIGSTY.el8.aarch64.rpm) |
| `q3c_13` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [q3c_13-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/q3c_13-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_13` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 160.3 KiB | [q3c_13-2.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_13-2.0.1-1PIGSTY.el9.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.1 KiB | [q3c_13-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/q3c_13-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 112.2 KiB | [q3c_13-2.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_13-2.0.1-1PIGSTY.el9.aarch64.rpm) |
| `q3c_13` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [q3c_13-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/q3c_13-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_13` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.9 KiB | [q3c_13-2.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_13-2.0.1-1PIGSTY.el10.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.1 KiB | [q3c_13-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/q3c_13-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_13` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.1 KiB | [q3c_13-2.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_13-2.0.1-1PIGSTY.el10.aarch64.rpm) |
| `q3c_13` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.0 KiB | [q3c_13-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/q3c_13-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-q3c` | `2.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 138.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg12+1_amd64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 142.5 KiB | [postgresql-13-q3c_2.0.1-6.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg12+1_arm64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 162.4 KiB | [postgresql-13-q3c_2.0.1-6.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg13+1_amd64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 151.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg13+1_arm64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 160.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg22.04+1_amd64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 155.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg22.04+1_arm64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 152.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg24.04+1_amd64.deb) |
| `postgresql-13-q3c` | `2.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 159.2 KiB | [postgresql-13-q3c_2.0.1-6.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-13-q3c_2.0.1-6.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/segasai/q3c" title="Repository" icon="github" subtitle="github.com/segasai/q3c" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="q3c-2.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg q3c;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install q3c;		# install via package name, for the active PG version

pig install q3c -v 18;   # install for PG 18
pig install q3c -v 17;   # install for PG 17
pig install q3c -v 16;   # install for PG 16
pig install q3c -v 15;   # install for PG 15
pig install q3c -v 14;   # install for PG 14
pig install q3c -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION q3c;
```
