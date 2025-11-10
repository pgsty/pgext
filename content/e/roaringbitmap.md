---
title: "roaringbitmap"
linkTitle: "roaringbitmap"
description: "support for Roaring Bitmaps"
weight: 3570
categories: ["TYPE"]
width: full
---

[**pg_roaringbitmap**](https://github.com/ChenHuajun/pg_roaringbitmap) : support for Roaring Bitmaps


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3570** | {{< badge content="roaringbitmap" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | `0.5.5` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "pgfaceting" >}} |
|   **See Also**    | {{< ext "rum" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.5.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_roaringbitmap` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.5` | {{< bg "18" "pg_roaringbitmap_18*" "green" >}} {{< bg "17" "pg_roaringbitmap_17*" "green" >}} {{< bg "16" "pg_roaringbitmap_16*" "green" >}} {{< bg "15" "pg_roaringbitmap_15*" "green" >}} {{< bg "14" "pg_roaringbitmap_14*" "green" >}} {{< bg "13" "pg_roaringbitmap_13*" "green" >}} | `pg_roaringbitmap_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.5.5` | {{< bg "18" "postgresql-18-roaringbitmap" "green" >}} {{< bg "17" "postgresql-17-roaringbitmap" "green" >}} {{< bg "16" "postgresql-16-roaringbitmap" "green" >}} {{< bg "15" "postgresql-15-roaringbitmap" "green" >}} {{< bg "14" "postgresql-14-roaringbitmap" "green" >}} {{< bg "13" "postgresql-13-roaringbitmap" "green" >}} | `postgresql-$v-roaringbitmap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_13 : AVAIL 3" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_13 : AVAIL 3" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_13 : AVAIL 3" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_16 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_15 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_14 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 0.5.5" "pg_roaringbitmap_13 : AVAIL 3" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 0.5.5" "pg_roaringbitmap_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.5.5" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.5.5" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_18` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 168.0 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.9 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 429.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 389.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 430.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 391.0 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 368.1 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 338.1 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.6 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 331.3 KiB | [postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_17` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 167.9 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.0 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 429.1 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 389.5 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 430.8 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 391.0 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.8 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 363.6 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.3 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 331.1 KiB | [postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_16` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 167.9 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.0 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 429.6 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 389.7 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 430.8 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 390.8 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 363.7 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 331.4 KiB | [postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_15` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.5 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.6 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.3 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.6 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.6 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.7 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 431.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 391.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 432.4 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 392.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 463.5 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 420.8 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 428.7 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 390.4 KiB | [postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_14` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.5 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.7 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.4 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.3 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 430.7 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 391.6 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 431.5 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 392.9 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 463.1 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 420.7 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 428.1 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 390.5 KiB | [postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_13` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.0 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 161.8 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.7 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.6 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.4 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.5 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.3 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.6 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.2 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.2 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.4 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.9 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 430.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 391.5 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg12+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 431.5 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg13+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 392.5 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg13+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 463.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 420.8 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 427.9 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `0.5.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 390.4 KiB | [postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_0.5.5-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChenHuajun/pg_roaringbitmap" title="Repository" icon="github" subtitle="github.com/ChenHuajun/pg_roaringbitmap" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_roaringbitmap-0.5.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_roaringbitmap;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_roaringbitmap;		# install via package name, for the active PG version
pig install roaringbitmap;		# install by extension name, for the current active PG version

pig install roaringbitmap -v 18;   # install for PG 18
pig install roaringbitmap -v 17;   # install for PG 17
pig install roaringbitmap -v 16;   # install for PG 16
pig install roaringbitmap -v 15;   # install for PG 15
pig install roaringbitmap -v 14;   # install for PG 14
pig install roaringbitmap -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION roaringbitmap;
```
