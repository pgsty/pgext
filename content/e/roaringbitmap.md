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
| **3570** | {{< badge content="roaringbitmap" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | `1.1.0` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_roaringbitmap` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "pg_roaringbitmap_18" "green" >}} {{< bg "17" "pg_roaringbitmap_17" "green" >}} {{< bg "16" "pg_roaringbitmap_16" "green" >}} {{< bg "15" "pg_roaringbitmap_15" "green" >}} {{< bg "14" "pg_roaringbitmap_14" "green" >}} {{< bg "13" "pg_roaringbitmap_13" "green" >}} | `pg_roaringbitmap_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "postgresql-18-roaringbitmap" "green" >}} {{< bg "17" "postgresql-17-roaringbitmap" "green" >}} {{< bg "16" "postgresql-16-roaringbitmap" "green" >}} {{< bg "15" "postgresql-15-roaringbitmap" "green" >}} {{< bg "14" "postgresql-14-roaringbitmap" "green" >}} {{< bg "13" "postgresql-13-roaringbitmap" "green" >}} | `postgresql-$v-roaringbitmap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 5" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.0" "pg_roaringbitmap_13 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1.0" "postgresql-18-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-17-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-16-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-15-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-14-roaringbitmap : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "postgresql-13-roaringbitmap : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.3 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 170.9 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 168.0 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.2 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.4 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.0 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 69.9 KiB | [pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_18-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.5 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.1 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 465.0 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 421.9 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 416.0 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 384.8 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 409.3 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 378.1 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.0 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 167.9 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.2 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.9 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.0 KiB | [pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_17-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.0 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.0 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.1 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.5 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 448.3 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 413.2 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 410.4 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 378.1 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.0 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 167.9 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.1 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 140.3 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.1 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.9 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.0 KiB | [pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_16-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.6 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.0 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.8 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.5 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 447.7 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 413.3 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 410.3 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 379.1 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.6 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.5 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.8 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.6 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.6 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.3 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.6 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.8 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.4 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.6 KiB | [pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_15-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.7 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.5 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.5 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.6 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 158.1 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 464.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 423.5 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 467.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 425.3 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 499.7 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 455.0 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 462.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 422.1 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.5 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.5 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.8 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.7 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.4 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.4 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.2 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.3 KiB | [pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_14-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 180.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.8 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.6 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.7 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 464.0 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 423.6 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 466.9 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 425.2 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 498.1 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 455.4 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 462.1 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 421.9 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_13` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.1 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 173.7 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 173.0 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 161.8 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 108.7 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.8 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.6 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_13` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.8 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.4 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 162.4 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.5 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.3 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.8 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.9 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 143.6 KiB | [pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_13-0.5.5-1PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.2 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_13` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.3 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 180.7 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.2 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.4 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_13` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 158.0 KiB | [pg_roaringbitmap_13-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_13` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.7 KiB | [pg_roaringbitmap_13-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_13` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.9 KiB | [pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_roaringbitmap_13-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 464.6 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 424.3 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 467.2 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 424.8 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 497.5 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 455.3 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 461.0 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 421.6 KiB | [postgresql-13-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-13-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChenHuajun/pg_roaringbitmap" title="Repository" icon="github" subtitle="github.com/ChenHuajun/pg_roaringbitmap" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_roaringbitmap-0.5.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_roaringbitmap;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
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
