---
title: "pgrouting"
linkTitle: "pgrouting"
description: "pgRouting Extension"
weight: 1510
categories: ["GIS"]
width: full
---

[**pgrouting**](https://github.com/pgRouting/pgrouting) : pgRouting Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1510** | {{< badge content="pgrouting" link="https://github.com/pgRouting/pgrouting" >}} | {{< ext "pgrouting" >}} | `4.0.1` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis_topology" >}} {{< ext "mobilitydb" >}} {{< ext "pg_polyline" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgrouting` | `plpgsql`, `postgis` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.0.1` | {{< bg "18" "pgrouting_18" "green" >}} {{< bg "17" "pgrouting_17" "green" >}} {{< bg "16" "pgrouting_16" "green" >}} {{< bg "15" "pgrouting_15" "green" >}} {{< bg "14" "pgrouting_14" "green" >}} | `pgrouting_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.0.1` | {{< bg "18" "postgresql-18-pgrouting" "green" >}} {{< bg "17" "postgresql-17-pgrouting" "green" >}} {{< bg "16" "postgresql-16-pgrouting" "green" >}} {{< bg "15" "postgresql-15-pgrouting" "green" >}} {{< bg "14" "postgresql-14-pgrouting" "green" >}} | `postgresql-$v-pgrouting` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 19" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 15" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 15" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 15" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 15" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.0.1" "pgrouting_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.0.1" "pgrouting_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 4.0.1" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.0.1" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_18` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 904.6 KiB | [pgrouting_18-4.0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgrouting_18-4.0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgrouting_18` | `3.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 943.4 KiB | [pgrouting_18-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgrouting_18-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_18` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 797.7 KiB | [pgrouting_18-4.0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgrouting_18-4.0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgrouting_18` | `3.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 830.6 KiB | [pgrouting_18-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgrouting_18-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_18` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 696.4 KiB | [pgrouting_18-4.0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgrouting_18-4.0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgrouting_18` | `3.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 741.6 KiB | [pgrouting_18-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgrouting_18-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_18` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 648.1 KiB | [pgrouting_18-4.0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgrouting_18-4.0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgrouting_18` | `3.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 693.5 KiB | [pgrouting_18-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgrouting_18-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_18` | `4.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 727.0 KiB | [pgrouting_18-4.0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgrouting_18-4.0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgrouting_18` | `3.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 773.0 KiB | [pgrouting_18-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgrouting_18-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_18` | `4.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 670.8 KiB | [pgrouting_18-4.0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgrouting_18-4.0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgrouting_18` | `3.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 718.9 KiB | [pgrouting_18-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgrouting_18-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgrouting` | `4.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 813.8 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 695.8 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 902.0 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 773.0 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 614.5 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 521.9 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 596.8 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 518.4 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 641.0 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pgrouting` | `4.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 566.9 KiB | [postgresql-18-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_17` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 904.6 KiB | [pgrouting_17-4.0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-4.0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgrouting_17` | `3.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 943.4 KiB | [pgrouting_17-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `3.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 921.1 KiB | [pgrouting_17-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `3.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 965.4 KiB | [pgrouting_17-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 968.8 KiB | [pgrouting_17-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 958.6 KiB | [pgrouting_17-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `3.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 958.2 KiB | [pgrouting_17-3.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 797.7 KiB | [pgrouting_17-4.0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-4.0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgrouting_17` | `3.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 830.6 KiB | [pgrouting_17-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `3.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.2 KiB | [pgrouting_17-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `3.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 849.0 KiB | [pgrouting_17-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 852.4 KiB | [pgrouting_17-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.0 KiB | [pgrouting_17-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `3.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 839.3 KiB | [pgrouting_17-3.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 696.5 KiB | [pgrouting_17-4.0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-4.0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgrouting_17` | `3.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 741.9 KiB | [pgrouting_17-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `3.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 719.1 KiB | [pgrouting_17-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `3.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 748.4 KiB | [pgrouting_17-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 752.6 KiB | [pgrouting_17-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 738.2 KiB | [pgrouting_17-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `3.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 737.4 KiB | [pgrouting_17-3.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 648.4 KiB | [pgrouting_17-4.0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-4.0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgrouting_17` | `3.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 693.0 KiB | [pgrouting_17-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `3.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 671.5 KiB | [pgrouting_17-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `3.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 699.2 KiB | [pgrouting_17-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 701.7 KiB | [pgrouting_17-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.7 KiB | [pgrouting_17-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `3.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.0 KiB | [pgrouting_17-3.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | `4.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 726.9 KiB | [pgrouting_17-4.0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgrouting_17-4.0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgrouting_17` | `3.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 772.5 KiB | [pgrouting_17-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgrouting_17-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_17` | `4.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 670.5 KiB | [pgrouting_17-4.0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgrouting_17-4.0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgrouting_17` | `3.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 719.0 KiB | [pgrouting_17-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgrouting_17-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgrouting` | `4.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 813.7 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 695.5 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 902.8 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 772.8 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 614.6 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 522.0 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 596.8 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 518.6 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 641.2 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pgrouting` | `4.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 566.8 KiB | [postgresql-17-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_16` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 904.6 KiB | [pgrouting_16-4.0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-4.0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgrouting_16` | `3.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 943.4 KiB | [pgrouting_16-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 921.1 KiB | [pgrouting_16-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 965.4 KiB | [pgrouting_16-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 968.8 KiB | [pgrouting_16-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 958.6 KiB | [pgrouting_16-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 956.8 KiB | [pgrouting_16-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | `3.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 939.1 KiB | [pgrouting_16-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_16` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 797.7 KiB | [pgrouting_16-4.0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-4.0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgrouting_16` | `3.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 830.6 KiB | [pgrouting_16-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.1 KiB | [pgrouting_16-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 849.0 KiB | [pgrouting_16-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 852.4 KiB | [pgrouting_16-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.0 KiB | [pgrouting_16-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 837.9 KiB | [pgrouting_16-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | `3.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 818.8 KiB | [pgrouting_16-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_16` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 696.8 KiB | [pgrouting_16-4.0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-4.0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgrouting_16` | `3.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 741.7 KiB | [pgrouting_16-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 719.2 KiB | [pgrouting_16-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 749.5 KiB | [pgrouting_16-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 752.4 KiB | [pgrouting_16-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 738.1 KiB | [pgrouting_16-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 736.8 KiB | [pgrouting_16-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | `3.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 732.6 KiB | [pgrouting_16-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_16` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 648.6 KiB | [pgrouting_16-4.0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-4.0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgrouting_16` | `3.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 693.1 KiB | [pgrouting_16-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 671.6 KiB | [pgrouting_16-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 698.7 KiB | [pgrouting_16-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 702.0 KiB | [pgrouting_16-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.8 KiB | [pgrouting_16-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.7 KiB | [pgrouting_16-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | `3.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.2 KiB | [pgrouting_16-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_16` | `4.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 726.7 KiB | [pgrouting_16-4.0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgrouting_16-4.0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgrouting_16` | `3.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 772.5 KiB | [pgrouting_16-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgrouting_16-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_16` | `4.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 670.6 KiB | [pgrouting_16-4.0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgrouting_16-4.0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgrouting_16` | `3.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 718.9 KiB | [pgrouting_16-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgrouting_16-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgrouting` | `4.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 813.9 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 695.4 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 902.5 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 772.9 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 614.5 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 521.8 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 596.9 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 518.6 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 641.1 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pgrouting` | `4.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 566.9 KiB | [postgresql-16-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_15` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 904.6 KiB | [pgrouting_15-4.0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-4.0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgrouting_15` | `3.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 943.4 KiB | [pgrouting_15-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 921.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 965.4 KiB | [pgrouting_15-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 968.8 KiB | [pgrouting_15-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 958.6 KiB | [pgrouting_15-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 956.8 KiB | [pgrouting_15-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 938.5 KiB | [pgrouting_15-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 918.2 KiB | [pgrouting_15-3.4.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.2-2.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 917.5 KiB | [pgrouting_15-3.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.1-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 915.9 KiB | [pgrouting_15-3.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.0-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 868.7 KiB | [pgrouting_15-3.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.3.4-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 868.4 KiB | [pgrouting_15-3.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.3.3-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 847.9 KiB | [pgrouting_15-3.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.2.2-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `3.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 789.4 KiB | [pgrouting_15-3.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.1.4-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 797.7 KiB | [pgrouting_15-4.0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-4.0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgrouting_15` | `3.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 830.6 KiB | [pgrouting_15-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 849.0 KiB | [pgrouting_15-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 852.4 KiB | [pgrouting_15-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.0 KiB | [pgrouting_15-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 837.9 KiB | [pgrouting_15-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 818.5 KiB | [pgrouting_15-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 821.6 KiB | [pgrouting_15-3.4.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.2-2.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 821.3 KiB | [pgrouting_15-3.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.1-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 819.7 KiB | [pgrouting_15-3.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.0-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 775.5 KiB | [pgrouting_15-3.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.3.4-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | `3.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 710.5 KiB | [pgrouting_15-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.1.4-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 696.6 KiB | [pgrouting_15-4.0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-4.0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgrouting_15` | `3.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 741.7 KiB | [pgrouting_15-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 719.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 749.1 KiB | [pgrouting_15-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 752.0 KiB | [pgrouting_15-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 738.2 KiB | [pgrouting_15-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 736.4 KiB | [pgrouting_15-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 732.8 KiB | [pgrouting_15-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 733.5 KiB | [pgrouting_15-3.4.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.2-2.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 732.1 KiB | [pgrouting_15-3.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.1-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 729.2 KiB | [pgrouting_15-3.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.0-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 698.0 KiB | [pgrouting_15-3.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.3.4-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 698.3 KiB | [pgrouting_15-3.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.3.3-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 684.9 KiB | [pgrouting_15-3.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.2.2-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `3.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 654.4 KiB | [pgrouting_15-3.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.1.4-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 648.3 KiB | [pgrouting_15-4.0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-4.0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgrouting_15` | `3.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 693.5 KiB | [pgrouting_15-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 671.6 KiB | [pgrouting_15-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 698.9 KiB | [pgrouting_15-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 701.7 KiB | [pgrouting_15-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.9 KiB | [pgrouting_15-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 689.2 KiB | [pgrouting_15-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.4 KiB | [pgrouting_15-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 686.7 KiB | [pgrouting_15-3.4.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.2-2.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.0 KiB | [pgrouting_15-3.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.1-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 683.6 KiB | [pgrouting_15-3.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.0-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 654.4 KiB | [pgrouting_15-3.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.3.4-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | `3.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 613.8 KiB | [pgrouting_15-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.1.4-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | `4.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 726.9 KiB | [pgrouting_15-4.0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgrouting_15-4.0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgrouting_15` | `3.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 772.3 KiB | [pgrouting_15-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgrouting_15-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_15` | `4.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 670.8 KiB | [pgrouting_15-4.0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgrouting_15-4.0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgrouting_15` | `3.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 718.8 KiB | [pgrouting_15-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgrouting_15-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgrouting` | `4.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 813.8 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 695.3 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 902.3 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 772.7 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 614.6 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 522.0 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 596.6 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 518.5 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 641.4 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pgrouting` | `4.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 566.8 KiB | [postgresql-15-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_14` | `4.0.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 904.6 KiB | [pgrouting_14-4.0.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-4.0.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pgrouting_14` | `3.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 943.4 KiB | [pgrouting_14-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.7.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 921.1 KiB | [pgrouting_14-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 965.4 KiB | [pgrouting_14-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 968.8 KiB | [pgrouting_14-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.6.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 958.7 KiB | [pgrouting_14-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 956.7 KiB | [pgrouting_14-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 938.4 KiB | [pgrouting_14-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 918.2 KiB | [pgrouting_14-3.4.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.2-2.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 917.5 KiB | [pgrouting_14-3.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 915.9 KiB | [pgrouting_14-3.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 868.8 KiB | [pgrouting_14-3.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.4-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.3.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 868.5 KiB | [pgrouting_14-3.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.3-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 868.2 KiB | [pgrouting_14-3.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.2-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 870.8 KiB | [pgrouting_14-3.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 861.0 KiB | [pgrouting_14-3.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 846.9 KiB | [pgrouting_14-3.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.2.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 788.8 KiB | [pgrouting_14-3.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.1.3-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `3.0.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 780.7 KiB | [pgrouting_14-3.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.0.5-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | `4.0.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 797.6 KiB | [pgrouting_14-4.0.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-4.0.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pgrouting_14` | `3.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 830.5 KiB | [pgrouting_14-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.7.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 810.1 KiB | [pgrouting_14-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 849.0 KiB | [pgrouting_14-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 852.3 KiB | [pgrouting_14-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.6.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 839.9 KiB | [pgrouting_14-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 837.9 KiB | [pgrouting_14-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 818.4 KiB | [pgrouting_14-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 821.8 KiB | [pgrouting_14-3.4.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.2-2.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 821.2 KiB | [pgrouting_14-3.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.1-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 819.6 KiB | [pgrouting_14-3.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.0-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 775.5 KiB | [pgrouting_14-3.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.3.4-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 758.3 KiB | [pgrouting_14-3.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.2.2-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 710.5 KiB | [pgrouting_14-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.1.4-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `3.0.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 702.6 KiB | [pgrouting_14-3.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.0.6-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | `4.0.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 696.2 KiB | [pgrouting_14-4.0.1-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-4.0.1-1PGDG.rhel9.7.x86_64.rpm) |
| `pgrouting_14` | `3.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 741.5 KiB | [pgrouting_14-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.7.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 719.2 KiB | [pgrouting_14-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.7.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 748.6 KiB | [pgrouting_14-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 752.6 KiB | [pgrouting_14-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.6.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 738.3 KiB | [pgrouting_14-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 736.6 KiB | [pgrouting_14-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 732.7 KiB | [pgrouting_14-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 733.4 KiB | [pgrouting_14-3.4.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.2-2.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 732.6 KiB | [pgrouting_14-3.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.1-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 729.1 KiB | [pgrouting_14-3.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.0-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.3.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 697.9 KiB | [pgrouting_14-3.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.4-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.3.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 697.6 KiB | [pgrouting_14-3.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.3-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.3.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 697.2 KiB | [pgrouting_14-3.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.2-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `3.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 699.2 KiB | [pgrouting_14-3.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.1-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | `4.0.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 648.5 KiB | [pgrouting_14-4.0.1-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-4.0.1-1PGDG.rhel9.7.aarch64.rpm) |
| `pgrouting_14` | `3.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 693.7 KiB | [pgrouting_14-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.7.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 671.7 KiB | [pgrouting_14-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.7.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 699.1 KiB | [pgrouting_14-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 701.7 KiB | [pgrouting_14-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.6.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.7 KiB | [pgrouting_14-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 689.0 KiB | [pgrouting_14-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.2 KiB | [pgrouting_14-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 687.4 KiB | [pgrouting_14-3.4.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.2-2.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 688.0 KiB | [pgrouting_14-3.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.1-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 684.3 KiB | [pgrouting_14-3.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.0-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.3.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 654.5 KiB | [pgrouting_14-3.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.3.4-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 642.3 KiB | [pgrouting_14-3.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.2.2-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 613.4 KiB | [pgrouting_14-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.1.4-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `3.0.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 607.8 KiB | [pgrouting_14-3.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.0.6-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | `4.0.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 726.8 KiB | [pgrouting_14-4.0.1-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgrouting_14-4.0.1-1PGDG.rhel10.1.x86_64.rpm) |
| `pgrouting_14` | `3.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 773.0 KiB | [pgrouting_14-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgrouting_14-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_14` | `4.0.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 670.6 KiB | [pgrouting_14-4.0.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgrouting_14-4.0.1-1PGDG.rhel10.1.aarch64.rpm) |
| `pgrouting_14` | `3.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 718.8 KiB | [pgrouting_14-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgrouting_14-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgrouting` | `4.0.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 813.8 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 695.6 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 902.8 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 773.0 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 614.5 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 521.9 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 596.6 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 518.5 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u26.x86_64](/os/u26.x86_64) | pgdg | 641.0 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pgrouting` | `4.0.1` | [u26.aarch64](/os/u26.aarch64) | pgdg | 566.7 KiB | [postgresql-14-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_4.0.1-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgRouting/pgrouting" title="Repository" icon="github" subtitle="github.com/pgRouting/pgrouting" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgrouting;		# install via package name, for the active PG version

pig install pgrouting -v 18;   # install for PG 18
pig install pgrouting -v 17;   # install for PG 17
pig install pgrouting -v 16;   # install for PG 16
pig install pgrouting -v 15;   # install for PG 15
pig install pgrouting -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgrouting CASCADE; -- requires plpgsql, postgis
```



## Usage

> [pgRouting - Routing on PostgreSQL](https://github.com/pgRouting/pgrouting)

pgRouting extends the PostGIS/PostgreSQL geospatial database to provide geospatial routing and other network analysis functionality.

This library contains the following features:

- All Pairs Shortest Path (Floyd-Warshall, Johnson)
- A* algorithm (with bidirectional variant)
- Dijkstra algorithms (cost, cost matrix, driving distance, K shortest paths, via routing, nearest)
- Bidirectional Dijkstra
- Traveling Salesman Problem (TSP)
- Network flow (max flow, Boykov-Kolmogorov, Edmonds-Karp, push-relabel)
- Spanning trees (Kruskal, Prim with BFS/DFS/driving distance variants)
- Graph components (connected, strong, biconnected, articulation points, bridges)
- Turn Restriction Shortest Path (TRSP)
- WithPoints routing (arbitrary locations on edges)
- Graph contraction and utility functions

### Getting Started

Enable the extension (requires PostGIS):

```sql
CREATE EXTENSION pgrouting CASCADE;
```

### Graph Representation

pgRouting represents graphs using SQL queries that return edge data. The standard edge query format:

```sql
SELECT id, source, target, cost, reverse_cost FROM edges;
```

| Column | Type | Description |
|--------|------|-------------|
| `id` | ANY-INTEGER | Edge identifier |
| `source` | ANY-INTEGER | Starting vertex identifier |
| `target` | ANY-INTEGER | Ending vertex identifier |
| `cost` | ANY-NUMERICAL | Weight (source to target); negative values exclude the edge |
| `reverse_cost` | ANY-NUMERICAL | Weight (target to source); default -1 (non-existent) |

### Simple Example Without Geometry

Create a graph and find the shortest path:

```sql
CREATE TABLE wiki (
  id SERIAL,
  source INTEGER,
  target INTEGER,
  cost INTEGER
);

INSERT INTO wiki (source, target, cost) VALUES
  (1, 2, 7),  (1, 3, 9), (1, 6, 14),
  (2, 3, 10), (2, 4, 15),
  (3, 6, 2),  (3, 4, 11),
  (4, 5, 6),
  (5, 6, 9);

SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost FROM wiki',
  1, 5, false);
```

--------

## Function Families

### Dijkstra - Shortest Path

The core routing function. Supports one-to-one, one-to-many, many-to-one, many-to-many, and combinations signatures.

```sql
pgr_dijkstra(Edges SQL, start_vid,  end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vid,  end_vids, [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vid,  [directed])
pgr_dijkstra(Edges SQL, start_vids, end_vids, [directed])
pgr_dijkstra(Edges SQL, Combinations SQL,     [directed])
```

Returns: `(seq, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

**One to One:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10, true);
```

**One to Many:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, ARRAY[10, 17]);
```

**Many to Many (undirected):**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  ARRAY[6, 1], ARRAY[10, 17],
  directed => false);
```

**Combinations:**

```sql
SELECT * FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  'SELECT source, target FROM combinations',
  false);
```

#### Dijkstra Cost

Returns only aggregate cost without path details:

```sql
pgr_dijkstraCost(Edges SQL, start_vid, end_vid, [directed])
```

Returns: `(start_vid, end_vid, agg_cost)`

#### Dijkstra Cost Matrix

Generate a cost matrix for a set of vertices:

```sql
pgr_dijkstraCostMatrix(Edges SQL, vids, [directed])
```

#### Dijkstra Via

Route through an ordered sequence of vertices:

```sql
pgr_dijkstraVia(Edges SQL, via_vertices, [directed, strict, U_turn_on_edge])
```

#### Dijkstra Near

Find the nearest vertex to a set of targets:

```sql
pgr_dijkstraNear(Edges SQL, start_vid, end_vids, [directed])
```

### A* - Shortest Path

Uses the A* heuristic algorithm. Requires additional coordinate columns (`x1`, `y1`, `x2`, `y2`) in the edges query.

```sql
pgr_aStar(Edges SQL, start_vid, end_vid, [directed, heuristic, factor, epsilon])
```

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `directed` | BOOLEAN | `true` | Graph direction |
| `heuristic` | INTEGER | `5` | Distance heuristic (0-5) |
| `factor` | FLOAT | `1` | Units manipulation |
| `epsilon` | FLOAT | `1` | Approximation factor |

```sql
SELECT * FROM pgr_aStar(
  'SELECT id, source, target, cost, reverse_cost, x1, y1, x2, y2 FROM edges',
  6, 12,
  directed => true, heuristic => 2);
```

Also available: `pgr_aStarCost`, `pgr_aStarCostMatrix`

### Bidirectional Algorithms

Bidirectional variants search from both ends simultaneously:

- `pgr_bdDijkstra`, `pgr_bdDijkstraCost`, `pgr_bdDijkstraCostMatrix`
- `pgr_bdAstar`, `pgr_bdAstarCost`, `pgr_bdAstarCostMatrix`

```sql
SELECT * FROM pgr_bdDijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10);
```

### K Shortest Paths (Yen's Algorithm)

Find the K shortest paths between two vertices:

```sql
pgr_KSP(Edges SQL, start_vid, end_vid, K, [directed, heap_paths])
```

Returns: `(seq, path_id, path_seq, start_vid, end_vid, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_KSP(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 17, 2);
```

### Driving Distance

Find all vertices reachable within a given distance:

```sql
pgr_drivingDistance(Edges SQL, start_vid, distance, [directed])
pgr_drivingDistance(Edges SQL, start_vids, distance, [directed, equicost])
```

Returns: `(seq, depth, start_vid, pred, node, edge, cost, agg_cost)`

```sql
SELECT * FROM pgr_drivingDistance(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  11, 3.0);
```

### Traveling Salesman Problem

**Matrix-based TSP:**

```sql
pgr_TSP(Matrix SQL, [start_id, end_id])
```

Returns: `(seq, node, cost, agg_cost)`

```sql
SELECT * FROM pgr_TSP(
  $$SELECT * FROM pgr_dijkstraCostMatrix(
    'SELECT id, source, target, cost, reverse_cost FROM edges',
    ARRAY[1, 3, 5, 6, 7, 8, 9, 10, 11, 15, 16, 17],
    directed => false)$$,
  start_id => 1);
```

**Euclidean TSP** (uses coordinates directly):

```sql
pgr_TSPeuclidean(Coordinates SQL, [start_id, end_id])
```

### Network Flow

Compute maximum flow and related properties:

```sql
-- Maximum flow
pgr_maxFlow(Edges SQL, source, sink)

-- Specific algorithms
pgr_boykovKolmogorov(Edges SQL, source, sink)
pgr_edmondsKarp(Edges SQL, source, sink)
pgr_pushRelabel(Edges SQL, source, sink)

-- Edge-disjoint paths
pgr_edgeDisjointPaths(Edges SQL, source, sink)

-- Maximum cardinality matching
pgr_maxCardinalityMatch(Edges SQL, [directed])
```

Flow edges SQL uses `capacity` and `reverse_capacity` instead of `cost`/`reverse_cost`.

### Spanning Trees

**Kruskal's algorithm:**

```sql
pgr_kruskal(Edges SQL)         -- minimum spanning tree
pgr_kruskalBFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDFS(Edges SQL, root_vid, [max_depth])
pgr_kruskalDD(Edges SQL, root_vid, distance)
```

**Prim's algorithm:**

```sql
pgr_prim(Edges SQL)            -- minimum spanning tree
pgr_primBFS(Edges SQL, root_vid, [max_depth])
pgr_primDFS(Edges SQL, root_vid, [max_depth])
pgr_primDD(Edges SQL, root_vid, distance)
```

### Graph Components

```sql
-- Connected components (undirected)
pgr_connectedComponents(Edges SQL)

-- Strongly connected components (directed)
pgr_strongComponents(Edges SQL)

-- Biconnected components
pgr_biconnectedComponents(Edges SQL)

-- Articulation points (cut vertices)
pgr_articulationPoints(Edges SQL)

-- Bridges (cut edges)
pgr_bridges(Edges SQL)
```

### Turn Restriction Shortest Path (TRSP)

Route with forbidden path restrictions:

```sql
pgr_trsp(Edges SQL, Restrictions SQL, start_vid, end_vid, [directed])
pgr_trspVia(Edges SQL, Restrictions SQL, via_vertices, [directed, strict, U_turn_on_edge])
pgr_trsp_withPoints(Edges SQL, Restrictions SQL, Points SQL, start_pid, end_pid, [options])
```

Restrictions SQL format:

| Column | Type | Description |
|--------|------|-------------|
| `path` | ARRAY[ANY-INTEGER] | Sequence of forbidden edge IDs |
| `cost` | ANY-NUMERICAL | Cost of the forbidden path |

### WithPoints - Routing from Arbitrary Locations

Route between points located on edges (not just vertices):

```sql
pgr_withPoints(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCost(Edges SQL, Points SQL, start_pid, end_pid, [options])
pgr_withPointsCostMatrix(Edges SQL, Points SQL, pids, [options])
pgr_withPointsKSP(Edges SQL, Points SQL, start_pid, end_pid, K, [options])
pgr_withPointsDD(Edges SQL, Points SQL, start_pid, distance, [options])
```

Points SQL format:

| Column | Type | Default | Description |
|--------|------|---------|-------------|
| `pid` | ANY-INTEGER | | Point identifier |
| `edge_id` | ANY-INTEGER | | Closest edge |
| `fraction` | ANY-NUMERICAL | | Position on edge (0-1) |
| `side` | CHAR | `b` | `r`(right), `l`(left), `b`(both) |

### Graph Contraction

Simplify graphs by contracting vertices:

```sql
pgr_contraction(Edges SQL, contraction_order, [options])
```

### Utility Functions

```sql
-- Extract vertices from edge data
pgr_extractVertices(Edges SQL)

-- Find edges near a point
pgr_findCloseEdges(Edges SQL, point, tolerance, [options])

-- Separate crossing geometries
pgr_separateCrossing(Edges SQL)

-- Separate touching geometries
pgr_separateTouching(Edges SQL)

-- Version info
pgr_version()
pgr_full_version()
```

--------

## Working with Geometries

### Building a Routing Topology

Extract vertices from spatial edges and build topology:

```sql
-- Extract vertices from edge geometries
SELECT * INTO vertices
FROM pgr_extractVertices('SELECT id, geom FROM edges ORDER BY id');

-- Set source vertex
UPDATE edges AS e
SET source = v.id, x1 = x, y1 = y
FROM vertices AS v
WHERE ST_StartPoint(e.geom) = v.geom;

-- Set target vertex
UPDATE edges AS e
SET target = v.id, x2 = x, y2 = y
FROM vertices AS v
WHERE ST_EndPoint(e.geom) = v.geom;
```

### Setting Costs from Geometry Length

```sql
UPDATE edges SET
  cost = sign(cost) * ST_Length(geom),
  reverse_cost = sign(reverse_cost) * ST_Length(geom);
```

### Getting Route Geometry

Combine routing results with edge geometries:

```sql
SELECT seq, node, edge, cost, agg_cost, geom
FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edges',
  6, 10
) AS r
LEFT JOIN edges AS e ON r.edge = e.id;
```

--------

## Performance Tips

Bound queries to the area of interest to reduce processed edges:

```sql
SELECT * FROM pgr_dijkstra($$
  SELECT id, source, target, cost, reverse_cost
  FROM edges
  WHERE geom && (
    SELECT ST_Buffer(ST_Union(geom), 1)
    FROM edges WHERE source IN (7) OR target IN (8))$$,
  7, 8);
```

--------

## All Pairs Shortest Path

For computing distances between all pairs of vertices:

```sql
-- Floyd-Warshall (no edge id required)
pgr_floydWarshall(Edges SQL, [directed])

-- Johnson (no edge id required)
pgr_johnson(Edges SQL, [directed])
```

Returns: `(start_vid, end_vid, agg_cost)`

```sql
SELECT * FROM pgr_floydWarshall(
  'SELECT id, source, target, cost, reverse_cost FROM edges');
```
