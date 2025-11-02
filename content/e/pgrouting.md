---
title: "pgrouting"
linkTitle: "pgrouting"
description: "pgRouting Extension"
weight: 1510
categories: ["GIS"]
width: full
---

pgRouting Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1510** | {{< badge content="pgrouting" link="https://github.com/pgRouting/pgrouting" >}} | {{< ext "pgrouting" >}} | `3.8.0` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis_topology" >}} {{< ext "mobilitydb" >}} {{< ext "pg_polyline" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |

> [!Note] pg17 on deb not ready


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgrouting" >}} | `3.8.0` | {{< bg "18" "pgrouting_18*" "green" >}} {{< bg "17" "pgrouting_17*" "green" >}} {{< bg "16" "pgrouting_16*" "green" >}} {{< bg "15" "pgrouting_15*" "green" >}} {{< bg "14" "pgrouting_14*" "green" >}} {{< bg "13" "pgrouting_13*" "green" >}} | `pgrouting_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgrouting" >}} | `3.8.0` | {{< bg "18" "postgresql-18-pgrouting postgresql-18-pgrouting-scripts" "green" >}} {{< bg "17" "postgresql-17-pgrouting postgresql-17-pgrouting-scripts" "green" >}} {{< bg "16" "postgresql-16-pgrouting postgresql-16-pgrouting-scripts" "green" >}} {{< bg "15" "postgresql-15-pgrouting postgresql-15-pgrouting-scripts" "green" >}} {{< bg "14" "postgresql-14-pgrouting postgresql-14-pgrouting-scripts" "green" >}} {{< bg "13" "postgresql-13-pgrouting postgresql-13-pgrouting-scripts" "green" >}} | `postgresql-$v-pgrouting postgresql-$v-pgrouting-scripts` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 14" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 18" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 18" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 12" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 14" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 11" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 14" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 14" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 14" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 12" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 14" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 12" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.8.0" "pgrouting_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "pgrouting_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.8.0" "postgresql-18-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-17-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-16-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-15-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-14-pgrouting : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.8.0" "postgresql-13-pgrouting : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_18` | 3.8.0 | `el8.x86_64` | pgdg | 943.4 KiB | [pgrouting_18-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgrouting_18-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_18` | 3.8.0 | `el8.aarch64` | pgdg | 830.6 KiB | [pgrouting_18-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgrouting_18-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_18` | 3.8.0 | `el9.x86_64` | pgdg | 741.6 KiB | [pgrouting_18-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgrouting_18-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_18` | 3.8.0 | `el9.aarch64` | pgdg | 693.5 KiB | [pgrouting_18-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgrouting_18-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_18` | 3.8.0 | `el10.x86_64` | pgdg | 773.0 KiB | [pgrouting_18-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pgrouting_18-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_18` | 3.8.0 | `el10.aarch64` | pgdg | 718.9 KiB | [pgrouting_18-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pgrouting_18-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 846.1 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 714.7 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.7 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 800.8 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.2 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.8 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 633.1 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.4 KiB | [postgresql-18-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-18-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_17` | 3.8.0 | `el8.x86_64` | pgdg | 943.4 KiB | [pgrouting_17-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.7.3 | `el8.x86_64` | pgdg | 921.1 KiB | [pgrouting_17-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.7.1 | `el8.x86_64` | pgdg | 965.4 KiB | [pgrouting_17-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.7.0 | `el8.x86_64` | pgdg | 968.8 KiB | [pgrouting_17-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.6.3 | `el8.x86_64` | pgdg | 958.6 KiB | [pgrouting_17-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.6.2 | `el8.x86_64` | pgdg | 958.2 KiB | [pgrouting_17-3.6.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgrouting_17-3.6.2-2PGDG.rhel8.x86_64.rpm) |
| `pgrouting_17` | 3.8.0 | `el8.aarch64` | pgdg | 830.6 KiB | [pgrouting_17-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.7.3 | `el8.aarch64` | pgdg | 810.2 KiB | [pgrouting_17-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.7.1 | `el8.aarch64` | pgdg | 849.0 KiB | [pgrouting_17-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.7.0 | `el8.aarch64` | pgdg | 852.4 KiB | [pgrouting_17-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.6.3 | `el8.aarch64` | pgdg | 840.0 KiB | [pgrouting_17-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.6.2 | `el8.aarch64` | pgdg | 839.3 KiB | [pgrouting_17-3.6.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgrouting_17-3.6.2-2PGDG.rhel8.aarch64.rpm) |
| `pgrouting_17` | 3.8.0 | `el9.x86_64` | pgdg | 741.9 KiB | [pgrouting_17-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.7.3 | `el9.x86_64` | pgdg | 719.1 KiB | [pgrouting_17-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.7.1 | `el9.x86_64` | pgdg | 748.4 KiB | [pgrouting_17-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.7.0 | `el9.x86_64` | pgdg | 752.6 KiB | [pgrouting_17-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.6.3 | `el9.x86_64` | pgdg | 738.2 KiB | [pgrouting_17-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.6.2 | `el9.x86_64` | pgdg | 737.4 KiB | [pgrouting_17-3.6.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgrouting_17-3.6.2-2PGDG.rhel9.x86_64.rpm) |
| `pgrouting_17` | 3.8.0 | `el9.aarch64` | pgdg | 693.0 KiB | [pgrouting_17-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.7.3 | `el9.aarch64` | pgdg | 671.5 KiB | [pgrouting_17-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.7.1 | `el9.aarch64` | pgdg | 699.2 KiB | [pgrouting_17-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.7.0 | `el9.aarch64` | pgdg | 701.7 KiB | [pgrouting_17-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.6.3 | `el9.aarch64` | pgdg | 688.7 KiB | [pgrouting_17-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.6.2 | `el9.aarch64` | pgdg | 688.0 KiB | [pgrouting_17-3.6.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgrouting_17-3.6.2-2PGDG.rhel9.aarch64.rpm) |
| `pgrouting_17` | 3.8.0 | `el10.x86_64` | pgdg | 772.5 KiB | [pgrouting_17-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pgrouting_17-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_17` | 3.8.0 | `el10.aarch64` | pgdg | 719.0 KiB | [pgrouting_17-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pgrouting_17-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 845.3 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 714.6 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.4 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 800.8 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.3 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.8 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 633.2 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.5 KiB | [postgresql-17-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-17-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_16` | 3.8.0 | `el8.x86_64` | pgdg | 943.4 KiB | [pgrouting_16-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.7.3 | `el8.x86_64` | pgdg | 921.1 KiB | [pgrouting_16-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.7.1 | `el8.x86_64` | pgdg | 965.4 KiB | [pgrouting_16-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.7.0 | `el8.x86_64` | pgdg | 968.8 KiB | [pgrouting_16-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.6.3 | `el8.x86_64` | pgdg | 958.6 KiB | [pgrouting_16-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.6.0 | `el8.x86_64` | pgdg | 956.8 KiB | [pgrouting_16-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.5.0 | `el8.x86_64` | pgdg | 939.1 KiB | [pgrouting_16-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgrouting_16-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_16` | 3.8.0 | `el8.aarch64` | pgdg | 830.6 KiB | [pgrouting_16-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.7.3 | `el8.aarch64` | pgdg | 810.1 KiB | [pgrouting_16-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.7.1 | `el8.aarch64` | pgdg | 849.0 KiB | [pgrouting_16-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.7.0 | `el8.aarch64` | pgdg | 852.4 KiB | [pgrouting_16-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.6.3 | `el8.aarch64` | pgdg | 840.0 KiB | [pgrouting_16-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.6.0 | `el8.aarch64` | pgdg | 837.9 KiB | [pgrouting_16-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.5.0 | `el8.aarch64` | pgdg | 818.8 KiB | [pgrouting_16-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgrouting_16-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_16` | 3.8.0 | `el9.x86_64` | pgdg | 741.7 KiB | [pgrouting_16-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.7.3 | `el9.x86_64` | pgdg | 719.2 KiB | [pgrouting_16-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.7.1 | `el9.x86_64` | pgdg | 749.5 KiB | [pgrouting_16-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.7.0 | `el9.x86_64` | pgdg | 752.4 KiB | [pgrouting_16-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.6.3 | `el9.x86_64` | pgdg | 738.1 KiB | [pgrouting_16-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.6.0 | `el9.x86_64` | pgdg | 736.8 KiB | [pgrouting_16-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.5.0 | `el9.x86_64` | pgdg | 732.6 KiB | [pgrouting_16-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgrouting_16-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_16` | 3.8.0 | `el9.aarch64` | pgdg | 693.1 KiB | [pgrouting_16-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.7.3 | `el9.aarch64` | pgdg | 671.6 KiB | [pgrouting_16-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.7.1 | `el9.aarch64` | pgdg | 698.7 KiB | [pgrouting_16-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.7.0 | `el9.aarch64` | pgdg | 702.0 KiB | [pgrouting_16-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.6.3 | `el9.aarch64` | pgdg | 688.8 KiB | [pgrouting_16-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.6.0 | `el9.aarch64` | pgdg | 688.7 KiB | [pgrouting_16-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.5.0 | `el9.aarch64` | pgdg | 688.2 KiB | [pgrouting_16-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgrouting_16-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_16` | 3.8.0 | `el10.x86_64` | pgdg | 772.5 KiB | [pgrouting_16-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pgrouting_16-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_16` | 3.8.0 | `el10.aarch64` | pgdg | 718.9 KiB | [pgrouting_16-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pgrouting_16-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 846.3 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 714.6 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.7 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 800.9 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.1 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.9 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 632.9 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.4 KiB | [postgresql-16-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-16-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_15` | 3.8.0 | `el8.x86_64` | pgdg | 943.4 KiB | [pgrouting_15-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.7.3 | `el8.x86_64` | pgdg | 921.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.7.1 | `el8.x86_64` | pgdg | 965.4 KiB | [pgrouting_15-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.7.0 | `el8.x86_64` | pgdg | 968.8 KiB | [pgrouting_15-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.6.3 | `el8.x86_64` | pgdg | 958.6 KiB | [pgrouting_15-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.6.0 | `el8.x86_64` | pgdg | 956.8 KiB | [pgrouting_15-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.5.0 | `el8.x86_64` | pgdg | 938.5 KiB | [pgrouting_15-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.4.2 | `el8.x86_64` | pgdg | 918.2 KiB | [pgrouting_15-3.4.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.2-2.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.4.1 | `el8.x86_64` | pgdg | 917.5 KiB | [pgrouting_15-3.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.1-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.4.0 | `el8.x86_64` | pgdg | 915.9 KiB | [pgrouting_15-3.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.4.0-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.3.4 | `el8.x86_64` | pgdg | 868.7 KiB | [pgrouting_15-3.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.3.4-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.3.3 | `el8.x86_64` | pgdg | 868.4 KiB | [pgrouting_15-3.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.3.3-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.2.2 | `el8.x86_64` | pgdg | 847.9 KiB | [pgrouting_15-3.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.2.2-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.1.4 | `el8.x86_64` | pgdg | 789.4 KiB | [pgrouting_15-3.1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgrouting_15-3.1.4-1.rhel8.x86_64.rpm) |
| `pgrouting_15` | 3.8.0 | `el8.aarch64` | pgdg | 830.6 KiB | [pgrouting_15-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.7.3 | `el8.aarch64` | pgdg | 810.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.7.1 | `el8.aarch64` | pgdg | 849.0 KiB | [pgrouting_15-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.7.0 | `el8.aarch64` | pgdg | 852.4 KiB | [pgrouting_15-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.6.3 | `el8.aarch64` | pgdg | 840.0 KiB | [pgrouting_15-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.6.0 | `el8.aarch64` | pgdg | 837.9 KiB | [pgrouting_15-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.5.0 | `el8.aarch64` | pgdg | 818.5 KiB | [pgrouting_15-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.4.2 | `el8.aarch64` | pgdg | 821.6 KiB | [pgrouting_15-3.4.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.2-2.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.4.1 | `el8.aarch64` | pgdg | 821.3 KiB | [pgrouting_15-3.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.1-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.4.0 | `el8.aarch64` | pgdg | 819.7 KiB | [pgrouting_15-3.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.4.0-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.3.4 | `el8.aarch64` | pgdg | 775.5 KiB | [pgrouting_15-3.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.3.4-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.1.4 | `el8.aarch64` | pgdg | 710.5 KiB | [pgrouting_15-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgrouting_15-3.1.4-1.rhel8.aarch64.rpm) |
| `pgrouting_15` | 3.8.0 | `el9.x86_64` | pgdg | 741.7 KiB | [pgrouting_15-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.7.3 | `el9.x86_64` | pgdg | 719.1 KiB | [pgrouting_15-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.7.1 | `el9.x86_64` | pgdg | 749.1 KiB | [pgrouting_15-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.7.0 | `el9.x86_64` | pgdg | 752.0 KiB | [pgrouting_15-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.6.3 | `el9.x86_64` | pgdg | 738.2 KiB | [pgrouting_15-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.6.0 | `el9.x86_64` | pgdg | 736.4 KiB | [pgrouting_15-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.5.0 | `el9.x86_64` | pgdg | 732.8 KiB | [pgrouting_15-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.4.2 | `el9.x86_64` | pgdg | 733.5 KiB | [pgrouting_15-3.4.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.2-2.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.4.1 | `el9.x86_64` | pgdg | 732.1 KiB | [pgrouting_15-3.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.1-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.4.0 | `el9.x86_64` | pgdg | 729.2 KiB | [pgrouting_15-3.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.4.0-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.3.4 | `el9.x86_64` | pgdg | 698.0 KiB | [pgrouting_15-3.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.3.4-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.3.3 | `el9.x86_64` | pgdg | 698.3 KiB | [pgrouting_15-3.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.3.3-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.2.2 | `el9.x86_64` | pgdg | 684.9 KiB | [pgrouting_15-3.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.2.2-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.1.4 | `el9.x86_64` | pgdg | 654.4 KiB | [pgrouting_15-3.1.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgrouting_15-3.1.4-1.rhel9.x86_64.rpm) |
| `pgrouting_15` | 3.8.0 | `el9.aarch64` | pgdg | 693.5 KiB | [pgrouting_15-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.7.3 | `el9.aarch64` | pgdg | 671.6 KiB | [pgrouting_15-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.7.1 | `el9.aarch64` | pgdg | 698.9 KiB | [pgrouting_15-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.7.0 | `el9.aarch64` | pgdg | 701.7 KiB | [pgrouting_15-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.6.3 | `el9.aarch64` | pgdg | 688.9 KiB | [pgrouting_15-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.6.0 | `el9.aarch64` | pgdg | 689.2 KiB | [pgrouting_15-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.5.0 | `el9.aarch64` | pgdg | 688.4 KiB | [pgrouting_15-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.4.2 | `el9.aarch64` | pgdg | 686.7 KiB | [pgrouting_15-3.4.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.2-2.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.4.1 | `el9.aarch64` | pgdg | 688.0 KiB | [pgrouting_15-3.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.1-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.4.0 | `el9.aarch64` | pgdg | 683.6 KiB | [pgrouting_15-3.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.4.0-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.3.4 | `el9.aarch64` | pgdg | 654.4 KiB | [pgrouting_15-3.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.3.4-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.1.4 | `el9.aarch64` | pgdg | 613.8 KiB | [pgrouting_15-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgrouting_15-3.1.4-1.rhel9.aarch64.rpm) |
| `pgrouting_15` | 3.8.0 | `el10.x86_64` | pgdg | 772.3 KiB | [pgrouting_15-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pgrouting_15-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_15` | 3.8.0 | `el10.aarch64` | pgdg | 718.8 KiB | [pgrouting_15-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pgrouting_15-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 845.3 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 714.6 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.9 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 801.2 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.0 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.8 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 633.4 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.4 KiB | [postgresql-15-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-15-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_14` | 3.8.0 | `el8.x86_64` | pgdg | 943.4 KiB | [pgrouting_14-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.7.3 | `el8.x86_64` | pgdg | 921.1 KiB | [pgrouting_14-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.7.1 | `el8.x86_64` | pgdg | 965.4 KiB | [pgrouting_14-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.7.0 | `el8.x86_64` | pgdg | 968.8 KiB | [pgrouting_14-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.6.3 | `el8.x86_64` | pgdg | 958.7 KiB | [pgrouting_14-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.6.0 | `el8.x86_64` | pgdg | 956.7 KiB | [pgrouting_14-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.5.0 | `el8.x86_64` | pgdg | 938.4 KiB | [pgrouting_14-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.4.2 | `el8.x86_64` | pgdg | 918.2 KiB | [pgrouting_14-3.4.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.2-2.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.4.1 | `el8.x86_64` | pgdg | 917.5 KiB | [pgrouting_14-3.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.4.0 | `el8.x86_64` | pgdg | 915.9 KiB | [pgrouting_14-3.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.4.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.3.4 | `el8.x86_64` | pgdg | 868.8 KiB | [pgrouting_14-3.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.4-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.3.3 | `el8.x86_64` | pgdg | 868.5 KiB | [pgrouting_14-3.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.3-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.3.2 | `el8.x86_64` | pgdg | 868.2 KiB | [pgrouting_14-3.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.2-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.3.1 | `el8.x86_64` | pgdg | 870.8 KiB | [pgrouting_14-3.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.3.0 | `el8.x86_64` | pgdg | 861.0 KiB | [pgrouting_14-3.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.3.0-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.2.1 | `el8.x86_64` | pgdg | 846.9 KiB | [pgrouting_14-3.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.2.1-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.1.3 | `el8.x86_64` | pgdg | 788.8 KiB | [pgrouting_14-3.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.1.3-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.0.5 | `el8.x86_64` | pgdg | 780.7 KiB | [pgrouting_14-3.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgrouting_14-3.0.5-1.rhel8.x86_64.rpm) |
| `pgrouting_14` | 3.8.0 | `el8.aarch64` | pgdg | 830.5 KiB | [pgrouting_14-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.7.3 | `el8.aarch64` | pgdg | 810.1 KiB | [pgrouting_14-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.7.1 | `el8.aarch64` | pgdg | 849.0 KiB | [pgrouting_14-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.7.0 | `el8.aarch64` | pgdg | 852.3 KiB | [pgrouting_14-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.6.3 | `el8.aarch64` | pgdg | 839.9 KiB | [pgrouting_14-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.6.0 | `el8.aarch64` | pgdg | 837.9 KiB | [pgrouting_14-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.5.0 | `el8.aarch64` | pgdg | 818.4 KiB | [pgrouting_14-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.4.2 | `el8.aarch64` | pgdg | 821.8 KiB | [pgrouting_14-3.4.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.2-2.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.4.1 | `el8.aarch64` | pgdg | 821.2 KiB | [pgrouting_14-3.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.1-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.4.0 | `el8.aarch64` | pgdg | 819.6 KiB | [pgrouting_14-3.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.4.0-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.3.4 | `el8.aarch64` | pgdg | 775.5 KiB | [pgrouting_14-3.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.3.4-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.2.2 | `el8.aarch64` | pgdg | 758.3 KiB | [pgrouting_14-3.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.2.2-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.1.4 | `el8.aarch64` | pgdg | 710.5 KiB | [pgrouting_14-3.1.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.1.4-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.0.6 | `el8.aarch64` | pgdg | 702.6 KiB | [pgrouting_14-3.0.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgrouting_14-3.0.6-1.rhel8.aarch64.rpm) |
| `pgrouting_14` | 3.8.0 | `el9.x86_64` | pgdg | 741.5 KiB | [pgrouting_14-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.7.3 | `el9.x86_64` | pgdg | 719.2 KiB | [pgrouting_14-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.7.1 | `el9.x86_64` | pgdg | 748.6 KiB | [pgrouting_14-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.7.0 | `el9.x86_64` | pgdg | 752.6 KiB | [pgrouting_14-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.6.3 | `el9.x86_64` | pgdg | 738.3 KiB | [pgrouting_14-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.6.0 | `el9.x86_64` | pgdg | 736.6 KiB | [pgrouting_14-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.5.0 | `el9.x86_64` | pgdg | 732.7 KiB | [pgrouting_14-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.4.2 | `el9.x86_64` | pgdg | 733.4 KiB | [pgrouting_14-3.4.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.2-2.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.4.1 | `el9.x86_64` | pgdg | 732.6 KiB | [pgrouting_14-3.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.1-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.4.0 | `el9.x86_64` | pgdg | 729.1 KiB | [pgrouting_14-3.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.4.0-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.3.4 | `el9.x86_64` | pgdg | 697.9 KiB | [pgrouting_14-3.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.4-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.3.3 | `el9.x86_64` | pgdg | 697.6 KiB | [pgrouting_14-3.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.3-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.3.2 | `el9.x86_64` | pgdg | 697.2 KiB | [pgrouting_14-3.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.2-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.3.1 | `el9.x86_64` | pgdg | 699.2 KiB | [pgrouting_14-3.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgrouting_14-3.3.1-1.rhel9.x86_64.rpm) |
| `pgrouting_14` | 3.8.0 | `el9.aarch64` | pgdg | 693.7 KiB | [pgrouting_14-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.7.3 | `el9.aarch64` | pgdg | 671.7 KiB | [pgrouting_14-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.7.1 | `el9.aarch64` | pgdg | 699.1 KiB | [pgrouting_14-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.7.0 | `el9.aarch64` | pgdg | 701.7 KiB | [pgrouting_14-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.6.3 | `el9.aarch64` | pgdg | 688.7 KiB | [pgrouting_14-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.6.0 | `el9.aarch64` | pgdg | 689.0 KiB | [pgrouting_14-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.5.0 | `el9.aarch64` | pgdg | 688.2 KiB | [pgrouting_14-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.4.2 | `el9.aarch64` | pgdg | 687.4 KiB | [pgrouting_14-3.4.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.2-2.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.4.1 | `el9.aarch64` | pgdg | 688.0 KiB | [pgrouting_14-3.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.1-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.4.0 | `el9.aarch64` | pgdg | 684.3 KiB | [pgrouting_14-3.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.4.0-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.3.4 | `el9.aarch64` | pgdg | 654.5 KiB | [pgrouting_14-3.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.3.4-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.2.2 | `el9.aarch64` | pgdg | 642.3 KiB | [pgrouting_14-3.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.2.2-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.1.4 | `el9.aarch64` | pgdg | 613.4 KiB | [pgrouting_14-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.1.4-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.0.6 | `el9.aarch64` | pgdg | 607.8 KiB | [pgrouting_14-3.0.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgrouting_14-3.0.6-1.rhel9.aarch64.rpm) |
| `pgrouting_14` | 3.8.0 | `el10.x86_64` | pgdg | 773.0 KiB | [pgrouting_14-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pgrouting_14-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_14` | 3.8.0 | `el10.aarch64` | pgdg | 718.8 KiB | [pgrouting_14-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pgrouting_14-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 845.8 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 714.7 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.7 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 800.9 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.2 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.8 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 633.0 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.4 KiB | [postgresql-14-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-14-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrouting_13` | 3.8.0 | `el8.x86_64` | pgdg | 942.0 KiB | [pgrouting_13-3.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.7.3 | `el8.x86_64` | pgdg | 919.7 KiB | [pgrouting_13-3.7.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.7.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.7.1 | `el8.x86_64` | pgdg | 963.9 KiB | [pgrouting_13-3.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.7.1-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.7.0 | `el8.x86_64` | pgdg | 967.2 KiB | [pgrouting_13-3.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.6.3 | `el8.x86_64` | pgdg | 957.6 KiB | [pgrouting_13-3.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.6.0 | `el8.x86_64` | pgdg | 955.2 KiB | [pgrouting_13-3.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.5.0 | `el8.x86_64` | pgdg | 936.9 KiB | [pgrouting_13-3.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.5.0-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.4.2 | `el8.x86_64` | pgdg | 916.8 KiB | [pgrouting_13-3.4.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.4.2-2.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.4.1 | `el8.x86_64` | pgdg | 915.9 KiB | [pgrouting_13-3.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.4.1-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.4.0 | `el8.x86_64` | pgdg | 914.4 KiB | [pgrouting_13-3.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.4.0-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.3.4 | `el8.x86_64` | pgdg | 867.1 KiB | [pgrouting_13-3.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.3.4-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.3.3 | `el8.x86_64` | pgdg | 866.8 KiB | [pgrouting_13-3.3.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.3.3-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.3.2 | `el8.x86_64` | pgdg | 866.5 KiB | [pgrouting_13-3.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.3.2-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.3.1 | `el8.x86_64` | pgdg | 869.6 KiB | [pgrouting_13-3.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.3.1-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.3.0 | `el8.x86_64` | pgdg | 859.5 KiB | [pgrouting_13-3.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.3.0-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.1.3 | `el8.x86_64` | pgdg | 788.1 KiB | [pgrouting_13-3.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.1.3-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.1.2 | `el8.x86_64` | pgdg | 801.8 KiB | [pgrouting_13-3.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.1.2-1.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.1.0 | `el8.x86_64` | pgdg | 792.4 KiB | [pgrouting_13-3.1.0-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgrouting_13-3.1.0-2.rhel8.x86_64.rpm) |
| `pgrouting_13` | 3.8.0 | `el8.aarch64` | pgdg | 830.3 KiB | [pgrouting_13-3.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.7.3 | `el8.aarch64` | pgdg | 810.3 KiB | [pgrouting_13-3.7.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.7.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.7.1 | `el8.aarch64` | pgdg | 849.1 KiB | [pgrouting_13-3.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.7.1-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.7.0 | `el8.aarch64` | pgdg | 852.0 KiB | [pgrouting_13-3.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.6.3 | `el8.aarch64` | pgdg | 839.6 KiB | [pgrouting_13-3.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.6.0 | `el8.aarch64` | pgdg | 837.7 KiB | [pgrouting_13-3.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.5.0 | `el8.aarch64` | pgdg | 818.5 KiB | [pgrouting_13-3.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.5.0-1.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.4.2 | `el8.aarch64` | pgdg | 821.9 KiB | [pgrouting_13-3.4.2-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.4.2-2.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.4.1 | `el8.aarch64` | pgdg | 821.3 KiB | [pgrouting_13-3.4.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.4.1-1.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.4.0 | `el8.aarch64` | pgdg | 819.6 KiB | [pgrouting_13-3.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.4.0-1.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.3.4 | `el8.aarch64` | pgdg | 775.5 KiB | [pgrouting_13-3.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgrouting_13-3.3.4-1.rhel8.aarch64.rpm) |
| `pgrouting_13` | 3.8.0 | `el9.x86_64` | pgdg | 741.6 KiB | [pgrouting_13-3.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.7.3 | `el9.x86_64` | pgdg | 717.5 KiB | [pgrouting_13-3.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.7.1 | `el9.x86_64` | pgdg | 749.4 KiB | [pgrouting_13-3.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.7.0 | `el9.x86_64` | pgdg | 752.3 KiB | [pgrouting_13-3.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.6.3 | `el9.x86_64` | pgdg | 739.0 KiB | [pgrouting_13-3.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.6.0 | `el9.x86_64` | pgdg | 737.1 KiB | [pgrouting_13-3.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.5.0 | `el9.x86_64` | pgdg | 732.4 KiB | [pgrouting_13-3.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.5.0-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.4.2 | `el9.x86_64` | pgdg | 734.6 KiB | [pgrouting_13-3.4.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.4.2-2.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.4.1 | `el9.x86_64` | pgdg | 733.4 KiB | [pgrouting_13-3.4.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.4.1-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.4.0 | `el9.x86_64` | pgdg | 730.4 KiB | [pgrouting_13-3.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.4.0-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.3.4 | `el9.x86_64` | pgdg | 698.6 KiB | [pgrouting_13-3.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.3.4-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.3.3 | `el9.x86_64` | pgdg | 698.2 KiB | [pgrouting_13-3.3.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.3.3-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.3.2 | `el9.x86_64` | pgdg | 696.8 KiB | [pgrouting_13-3.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.3.2-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.3.1 | `el9.x86_64` | pgdg | 700.3 KiB | [pgrouting_13-3.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgrouting_13-3.3.1-1.rhel9.x86_64.rpm) |
| `pgrouting_13` | 3.8.0 | `el9.aarch64` | pgdg | 693.6 KiB | [pgrouting_13-3.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.7.3 | `el9.aarch64` | pgdg | 671.6 KiB | [pgrouting_13-3.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.7.1 | `el9.aarch64` | pgdg | 698.9 KiB | [pgrouting_13-3.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.7.0 | `el9.aarch64` | pgdg | 701.3 KiB | [pgrouting_13-3.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.6.3 | `el9.aarch64` | pgdg | 688.5 KiB | [pgrouting_13-3.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.6.0 | `el9.aarch64` | pgdg | 689.2 KiB | [pgrouting_13-3.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.5.0 | `el9.aarch64` | pgdg | 687.9 KiB | [pgrouting_13-3.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.5.0-1.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.4.2 | `el9.aarch64` | pgdg | 688.4 KiB | [pgrouting_13-3.4.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.4.2-2.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.4.1 | `el9.aarch64` | pgdg | 687.0 KiB | [pgrouting_13-3.4.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.4.1-1.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.4.0 | `el9.aarch64` | pgdg | 684.4 KiB | [pgrouting_13-3.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.4.0-1.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.3.4 | `el9.aarch64` | pgdg | 655.0 KiB | [pgrouting_13-3.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.3.4-1.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.1.4 | `el9.aarch64` | pgdg | 614.4 KiB | [pgrouting_13-3.1.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgrouting_13-3.1.4-1.rhel9.aarch64.rpm) |
| `pgrouting_13` | 3.8.0 | `el10.x86_64` | pgdg | 772.6 KiB | [pgrouting_13-3.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pgrouting_13-3.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pgrouting_13` | 3.8.0 | `el10.aarch64` | pgdg | 719.0 KiB | [pgrouting_13-3.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pgrouting_13-3.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pgrouting` | 3.8.0 | `d12.x86_64` | pgdg | 846.4 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg12+1_amd64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `d12.aarch64` | pgdg | 715.7 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg12+1_arm64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `d13.x86_64` | pgdg | 935.7 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg13+1_amd64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `d13.aarch64` | pgdg | 800.9 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg13+1_arm64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `u22.x86_64` | pgdg | 642.3 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `u22.aarch64` | pgdg | 544.2 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `u24.x86_64` | pgdg | 633.0 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pgrouting` | 3.8.0 | `u24.aarch64` | pgdg | 546.7 KiB | [postgresql-13-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgrouting/postgresql-13-pgrouting_3.8.0-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgRouting/pgrouting" title="Repository" icon="github" subtitle="github.com/pgRouting/pgrouting" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgrouting; # install by extension name, for the current active PG version
pig ext install pgrouting; # install via package alias, for the active PG version
pig ext install pgrouting -v 18;   # install for PG 18
pig ext install pgrouting -v 17;   # install for PG 17
pig ext install pgrouting -v 16;   # install for PG 16
pig ext install pgrouting -v 15;   # install for PG 15
pig ext install pgrouting -v 14;   # install for PG 14
pig ext install pgrouting -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgrouting;
```

