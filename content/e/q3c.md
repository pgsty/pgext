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
| **1540** | {{< badge content="q3c" link="https://github.com/segasai/q3c" >}} | {{< ext "q3c" >}} | `2.0.2` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "earthdistance" >}} {{< ext "pg_sphere" >}} {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `q3c` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "q3c_18" "green" >}} {{< bg "17" "q3c_17" "green" >}} {{< bg "16" "q3c_16" "green" >}} {{< bg "15" "q3c_15" "green" >}} {{< bg "14" "q3c_14" "green" >}} | `q3c_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.2` | {{< bg "18" "postgresql-18-q3c" "green" >}} {{< bg "17" "postgresql-17-q3c" "green" >}} {{< bg "16" "postgresql-16-q3c" "green" >}} {{< bg "15" "postgresql-15-q3c" "green" >}} {{< bg "14" "postgresql-14-q3c" "green" >}} | `postgresql-$v-q3c` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.2" "q3c_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.0.2" "q3c_14 : AVAIL 3" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.2" "postgresql-18-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-17-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-16-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-15-q3c : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.2" "postgresql-14-q3c : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 99.3 KiB | [q3c_18-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_18-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.7 KiB | [q3c_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/q3c_18-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/q3c_18-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.6 KiB | [q3c_18-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_18-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `q3c_18` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.8 KiB | [q3c_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/q3c_18-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/q3c_18-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_18` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 97.9 KiB | [q3c_18-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_18-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.7 KiB | [q3c_18-2.0.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/q3c_18-2.0.2-1PGDG.rhel9.7.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 108.7 KiB | [q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/q3c_18-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 126.6 KiB | [q3c_18-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_18-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `q3c_18` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 112.4 KiB | [q3c_18-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/q3c_18-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/q3c_18-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_18` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.4 KiB | [q3c_18-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_18-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 115.8 KiB | [q3c_18-2.0.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/q3c_18-2.0.2-1PGDG.rhel10.1.x86_64.rpm) |
| `q3c_18` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_18-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/q3c_18-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.3 KiB | [q3c_18-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_18-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `q3c_18` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 132.4 KiB | [q3c_18-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/q3c_18-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `q3c_18` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.9 KiB | [q3c_18-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/q3c_18-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-q3c` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 126.9 KiB | [postgresql-18-q3c_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 134.6 KiB | [postgresql-18-q3c_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 136.5 KiB | [postgresql-18-q3c_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 157.6 KiB | [postgresql-18-q3c_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 145.1 KiB | [postgresql-18-q3c_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 153.4 KiB | [postgresql-18-q3c_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 127.0 KiB | [postgresql-18-q3c_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 155.2 KiB | [postgresql-18-q3c_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 143.8 KiB | [postgresql-18-q3c_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-q3c` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 148.8 KiB | [postgresql-18-q3c_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-18-q3c_2.0.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 99.3 KiB | [q3c_17-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_17-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.7 KiB | [q3c_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/q3c_17-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/q3c_17-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.6 KiB | [q3c_17-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_17-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `q3c_17` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.8 KiB | [q3c_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/q3c_17-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/q3c_17-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_17` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 97.9 KiB | [q3c_17-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_17-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 136.4 KiB | [q3c_17-2.0.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/q3c_17-2.0.2-1PGDG.rhel9.7.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 101.8 KiB | [q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/q3c_17-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 94.3 KiB | [q3c_17-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_17-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `q3c_17` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 107.7 KiB | [q3c_17-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/q3c_17-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.4 KiB | [q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/q3c_17-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_17` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.4 KiB | [q3c_17-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_17-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 112.7 KiB | [q3c_17-2.0.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/q3c_17-2.0.2-1PGDG.rhel10.1.x86_64.rpm) |
| `q3c_17` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_17-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/q3c_17-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.2 KiB | [q3c_17-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_17-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `q3c_17` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 132.3 KiB | [q3c_17-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/q3c_17-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `q3c_17` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.7 KiB | [q3c_17-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/q3c_17-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-q3c` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 142.4 KiB | [postgresql-17-q3c_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 152.2 KiB | [postgresql-17-q3c_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 149.7 KiB | [postgresql-17-q3c_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 130.3 KiB | [postgresql-17-q3c_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 148.4 KiB | [postgresql-17-q3c_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 136.8 KiB | [postgresql-17-q3c_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 126.9 KiB | [postgresql-17-q3c_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 148.9 KiB | [postgresql-17-q3c_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 140.4 KiB | [postgresql-17-q3c_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-q3c` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 150.2 KiB | [postgresql-17-q3c_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-17-q3c_2.0.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 99.3 KiB | [q3c_16-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_16-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.7 KiB | [q3c_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/q3c_16-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/q3c_16-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.6 KiB | [q3c_16-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_16-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `q3c_16` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.8 KiB | [q3c_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/q3c_16-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/q3c_16-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 97.9 KiB | [q3c_16-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_16-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 136.4 KiB | [q3c_16-2.0.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/q3c_16-2.0.2-1PGDG.rhel9.7.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 103.3 KiB | [q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/q3c_16-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 94.4 KiB | [q3c_16-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_16-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `q3c_16` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.8 KiB | [q3c_16-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/q3c_16-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.4 KiB | [q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/q3c_16-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_16` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.4 KiB | [q3c_16-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_16-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 112.7 KiB | [q3c_16-2.0.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/q3c_16-2.0.2-1PGDG.rhel10.1.x86_64.rpm) |
| `q3c_16` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 127.7 KiB | [q3c_16-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/q3c_16-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.2 KiB | [q3c_16-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_16-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `q3c_16` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 132.4 KiB | [q3c_16-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/q3c_16-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `q3c_16` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.5 KiB | [q3c_16-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/q3c_16-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-q3c` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 132.7 KiB | [postgresql-16-q3c_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 143.5 KiB | [postgresql-16-q3c_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 130.4 KiB | [postgresql-16-q3c_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 159.2 KiB | [postgresql-16-q3c_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 136.1 KiB | [postgresql-16-q3c_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 158.2 KiB | [postgresql-16-q3c_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 126.7 KiB | [postgresql-16-q3c_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 154.6 KiB | [postgresql-16-q3c_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 133.2 KiB | [postgresql-16-q3c_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-q3c` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 165.8 KiB | [postgresql-16-q3c_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-16-q3c_2.0.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.2 KiB | [q3c_15-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_15-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/q3c_15-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.3 KiB | [q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/q3c_15-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.2 KiB | [q3c_15-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_15-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `q3c_15` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.3 KiB | [q3c_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/q3c_15-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/q3c_15-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_15` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 109.3 KiB | [q3c_15-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_15-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.1 KiB | [q3c_15-2.0.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/q3c_15-2.0.2-1PGDG.rhel9.7.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.3 KiB | [q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/q3c_15-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 102.3 KiB | [q3c_15-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_15-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `q3c_15` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 108.7 KiB | [q3c_15-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/q3c_15-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.5 KiB | [q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/q3c_15-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_15` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.1 KiB | [q3c_15-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_15-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 112.2 KiB | [q3c_15-2.0.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/q3c_15-2.0.2-1PGDG.rhel10.1.x86_64.rpm) |
| `q3c_15` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.2 KiB | [q3c_15-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/q3c_15-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.4 KiB | [q3c_15-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_15-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `q3c_15` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 113.0 KiB | [q3c_15-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/q3c_15-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `q3c_15` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.1 KiB | [q3c_15-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/q3c_15-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-q3c` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.2 KiB | [postgresql-15-q3c_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 149.0 KiB | [postgresql-15-q3c_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 148.2 KiB | [postgresql-15-q3c_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 124.4 KiB | [postgresql-15-q3c_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 140.8 KiB | [postgresql-15-q3c_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 162.2 KiB | [postgresql-15-q3c_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 158.2 KiB | [postgresql-15-q3c_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 153.8 KiB | [postgresql-15-q3c_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 156.5 KiB | [postgresql-15-q3c_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-q3c` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 141.0 KiB | [postgresql-15-q3c_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-15-q3c_2.0.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `q3c_14` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 98.2 KiB | [q3c_14-2.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/q3c_14-2.0.2-1PIGSTY.el8.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [q3c_14-2.0.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/q3c_14-2.0.2-1PGDG.rhel8.10.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.3 KiB | [q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/q3c_14-2.0.1-1PGDG.rhel8.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 93.2 KiB | [q3c_14-2.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/q3c_14-2.0.2-1PIGSTY.el8.aarch64.rpm) |
| `q3c_14` | `2.0.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.3 KiB | [q3c_14-2.0.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/q3c_14-2.0.2-1PGDG.rhel8.10.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/q3c_14-2.0.1-1PGDG.rhel8.aarch64.rpm) |
| `q3c_14` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 109.3 KiB | [q3c_14-2.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/q3c_14-2.0.2-1PIGSTY.el9.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 91.6 KiB | [q3c_14-2.0.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/q3c_14-2.0.2-1PGDG.rhel9.7.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 109.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/q3c_14-2.0.1-1PGDG.rhel9.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 101.0 KiB | [q3c_14-2.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/q3c_14-2.0.2-1PIGSTY.el9.aarch64.rpm) |
| `q3c_14` | `2.0.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 108.7 KiB | [q3c_14-2.0.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/q3c_14-2.0.2-1PGDG.rhel9.7.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/q3c_14-2.0.1-1PGDG.rhel9.aarch64.rpm) |
| `q3c_14` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 106.0 KiB | [q3c_14-2.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/q3c_14-2.0.2-1PIGSTY.el10.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 112.2 KiB | [q3c_14-2.0.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/q3c_14-2.0.2-1PGDG.rhel10.1.x86_64.rpm) |
| `q3c_14` | `2.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 92.1 KiB | [q3c_14-2.0.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/q3c_14-2.0.1-1PGDG.rhel10.x86_64.rpm) |
| `q3c_14` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 127.1 KiB | [q3c_14-2.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/q3c_14-2.0.2-1PIGSTY.el10.aarch64.rpm) |
| `q3c_14` | `2.0.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 115.3 KiB | [q3c_14-2.0.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/q3c_14-2.0.2-1PGDG.rhel10.1.aarch64.rpm) |
| `q3c_14` | `2.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.0 KiB | [q3c_14-2.0.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/q3c_14-2.0.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-q3c` | `2.0.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 135.0 KiB | [postgresql-14-q3c_2.0.2-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg12+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 161.5 KiB | [postgresql-14-q3c_2.0.2-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg12+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 127.1 KiB | [postgresql-14-q3c_2.0.2-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg13+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 160.6 KiB | [postgresql-14-q3c_2.0.2-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg13+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 137.8 KiB | [postgresql-14-q3c_2.0.2-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 162.5 KiB | [postgresql-14-q3c_2.0.2-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 145.2 KiB | [postgresql-14-q3c_2.0.2-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 147.0 KiB | [postgresql-14-q3c_2.0.2-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 135.5 KiB | [postgresql-14-q3c_2.0.2-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-q3c` | `2.0.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 137.1 KiB | [postgresql-14-q3c_2.0.2-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-q3c/postgresql-14-q3c_2.0.2-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/segasai/q3c" title="Repository" icon="github" subtitle="github.com/segasai/q3c" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="q3c-2.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg q3c;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install q3c;		# install via package name, for the active PG version

pig install q3c -v 18;   # install for PG 18
pig install q3c -v 17;   # install for PG 17
pig install q3c -v 16;   # install for PG 16
pig install q3c -v 15;   # install for PG 15
pig install q3c -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION q3c;
```



## Usage

> Source: [`segasai/q3c`](https://github.com/segasai/q3c) | [ADASS Paper](http://adsabs.harvard.edu/abs/2006ASPC..351..735K) | [ASCL](https://ascl.net/1905.008)

Q3C (Quad Tree Cube) is a PostgreSQL extension for fast sky-indexing of astronomical catalogues. It enables efficient spatial queries on spherical coordinates (right ascension and declination), including cone searches, ellipse searches, polygon queries, positional cross-matches, and nearest-neighbor lookups.

All angles (ra, dec, distances) are in **degrees**, proper motions in **mas/year**, and epochs in **years** (e.g. 2000.5, 2010.5). All Q3C function names start with the `q3c_` prefix.

### Table Preparation

To use Q3C, create a spatial index on your table with `ra` and `dec` columns (in degrees):

```sql
CREATE INDEX ON mytable (q3c_ang2ipix(ra, dec));
```

Optionally cluster the table by the index to ensure faster queries on large datasets:

```sql
CLUSTER mytable_q3c_ang2ipix_idx ON mytable;
```

Alternatively, reorder the table before indexing:

```sql
CREATE TABLE mytable1 AS SELECT * FROM mytable ORDER BY q3c_ang2ipix(ra, dec);
```

After indexing, analyze the table:

```sql
ANALYZE mytable;
```


## Functions

- `q3c_ang2ipix(ra, dec)` -- returns the ipix value (64-bit integer pixel identifier) for given ra and dec

- `q3c_dist(ra1, dec1, ra2, dec2)` -- returns the distance in degrees between two points

- `q3c_dist_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2)` -- returns distance in degrees between two points, taking proper motion into account. The `cosdec_flag` (0 or 1) indicates whether the proper motion includes the cos(dec) term (1) or not (0).

- `q3c_join(ra1, dec1, ra2, dec2, radius)` -- returns true if (ra1, dec1) is within `radius` spherical distance of (ra2, dec2). Use when the index on `q3c_ang2ipix(ra2, dec2)` is created.

- `q3c_join_pm(ra1, dec1, pmra1, pmdec1, cosdec_flag, epoch1, ra2, dec2, epoch2, max_delta_epoch, radius)` -- like `q3c_join` but takes proper motion into account. `max_delta_epoch` is the maximum epoch difference possible between two tables.

- `q3c_ellipse_join(ra1, dec1, ra2, dec2, major, ratio, pa)` -- like `q3c_join`, except (ra1, dec1) must be within an ellipse with semi-major axis `major`, axis ratio `ratio`, and position angle `pa` (from north through east)

- `q3c_radial_query(ra, dec, center_ra, center_dec, radius)` -- returns true if (ra, dec) is within `radius` degrees of (center_ra, center_dec). Main function for cone searches. Requires index on `q3c_ang2ipix(ra, dec)`.

- `q3c_ellipse_query(ra, dec, center_ra, center_dec, maj_ax, axis_ratio, PA)` -- returns true if (ra, dec) is within the ellipse from (center_ra, center_dec), specified by semi-major axis, axis ratio, and positional angle.

- `q3c_poly_query(ra, dec, poly)` -- returns true if (ra, dec) is within the spherical polygon specified as an array of RA/DEC values or a PostgreSQL polygon type. Uses the index.

- `q3c_ipix2ang(ipix)` -- returns a two-element array of (ra, dec) corresponding to a given ipix

- `q3c_pixarea(ipix, bits)` -- returns the spherical area corresponding to a given ipix at the pixelisation level given by `bits` (1 is smallest, 30 is the cube face)

- `q3c_ipixcenter(ra, dec, bits)` -- returns the ipix value of the pixel center at certain pixel depth covering the specified (ra, dec)

- `q3c_in_poly(ra, dec, poly)` -- returns true/false if point is inside a polygon. Does **NOT** use the q3c index.

- `q3c_version()` -- returns the installed version of Q3C


## Examples

### Cone Search

Query all objects within 0.1 degrees of (ra, dec) = (11, 12):

```sql
SELECT * FROM mytable WHERE q3c_radial_query(ra, dec, 11, 12, 0.1);
```

The column names of the table must come first, and the search location after, otherwise the index will not be used.

Alternative cone search using `q3c_join` (can be faster for small tables):

```sql
SELECT * FROM mytable WHERE q3c_join(11, 12, ra, dec, 0.1);
```

### Ellipse Search

Search for objects within an ellipse centered at (10, 20) with semi-major axis 1 degree, axis ratio 0.5, and PA of 10 degrees:

```sql
SELECT * FROM mytable WHERE q3c_ellipse_query(ra, dec, 10, 20, 1, 0.5, 10);
```

### Polygon Search

Query objects inside a spherical polygon with vertices (0,0), (2,0), (2,1), (0,1):

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, ARRAY[0, 0, 2, 0, 2, 1, 0, 1]);
```

Using PostgreSQL polygon type:

```sql
SELECT * FROM mytable WHERE
    q3c_poly_query(ra, dec, '((0, 0), (2, 0), (2, 1), (0, 1))'::polygon);
```

### Positional Cross-Match

Cross-match `table1` and `table2` within 0.001 degrees. The index must exist on `q3c_ang2ipix(ra, dec)` of `table2`:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, 0.001);
```

The ra/dec columns from the indexed table must be the 3rd and 4th arguments. This returns **all** pairs within the matching distance, not just nearest neighbors.

With per-object error radius:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join(a.ra, a.dec, b.ra, b.dec, a.err);
```

### Ellipse Cross-Match

Cross-match using elliptical error areas (e.g., matching within galaxy elliptical bodies):

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_ellipse_join(a.ra, a.dec, b.ra, b.dec, a.maj_ax, a.axis_ratio, a.PA);
```

### Cross-Match with Proper Motion

Cross-match with proper motion correction. Assumes `table1` has `pmra`, `pmdec` (mas/yr) and `epoch` columns, pmra includes cos(dec) factor, and max epoch difference is 30 years:

```sql
SELECT * FROM table1 AS a, table2 AS b WHERE
    q3c_join_pm(a.ra, a.dec, a.pmra, a.pmdec, 1,
    a.epoch, b.ra, b.dec, b.epoch, 30, 0.001);
```

### Nearest Neighbour (with NULLs for unmatched)

Returns the nearest neighbour for each row, with NULLs if no match exists within 1 arcsecond:

```sql
SELECT t.*, ss.* FROM mytable AS t
LEFT JOIN LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss ON true;
```

### Nearest Neighbour (matched only)

Returns only objects that have neighbours:

```sql
SELECT t.*, ss.* FROM mytable AS t,
LATERAL (
    SELECT s.*
    FROM sdssdr9.phototag AS s
    WHERE q3c_join(t.ra, t.dec, s.ra, s.dec, 1./3600)
    ORDER BY q3c_dist(t.ra, t.dec, s.ra, s.dec) ASC
    LIMIT 1
) AS ss;
```

### Nearest Neighbour (CTE variant)

Uses a CTE with an object ID column (requires an index on the ID column):

```sql
WITH x AS MATERIALIZED (
    SELECT *, (
        SELECT objid FROM sdssdr9.phototag AS p
        WHERE q3c_join(m.ra, m.dec, p.ra, p.dec, 1./3600)
        ORDER BY q3c_dist(m.ra, m.dec, p.ra, p.dec) ASC
        LIMIT 1
    ) AS match_objid
    FROM mytable AS m
)
SELECT * FROM x, sdssdr9.phototag AS s WHERE x.match_objid = s.objid;
```

### Density Estimation

Estimate object density using pixelation depth of 25:

```sql
SELECT (q3c_ipix2ang(i))[1] AS ra,
       (q3c_ipix2ang(i))[2] AS dec,
       c,
       q3c_pixarea(i, 25) AS area
FROM (
    SELECT q3c_ipixcenter(ra, dec, 25) AS i, count(*) AS c
    FROM mytable
    GROUP BY i
) AS x;
```

Note: Q3C does not have uniform pixel areas (unlike HEALPIX).


## Limitations

- Querying very large polygons with diameter greater than ~25 degrees is not supported
- Polygons with more than 100 vertices are not supported


## Performance Tips

- Ensure correct argument order in Q3C functions (e.g., `q3c_radial_query(ra, dec, 120, 3, 1)` not `q3c_radial_query(120, 3, ra, dec, 1)`)
- Use `EXPLAIN` to verify the query plan uses bitmap scans on the Q3C index
- If the planner chooses a bad plan, try: `SET enable_mergejoin TO off; SET enable_seqscan TO off; SET enable_hashjoin TO off;`
- Cluster the table using the Q3C index for best performance
- When combining `q3c_join()` with additional filter clauses, use CTEs with `MATERIALIZED` to avoid plan issues:

```sql
WITH x AS MATERIALIZED (SELECT * FROM t1 WHERE t1.mag < 1),
     y AS (SELECT *, t2.mag AS t2mag FROM x, t2 WHERE q3c_join(x.ra, x.dec, t2.ra, t2.dec, 1./3600))
SELECT * FROM y WHERE t2mag > 33;
```
