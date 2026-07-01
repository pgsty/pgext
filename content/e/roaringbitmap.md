---
title: "roaringbitmap"
linkTitle: "roaringbitmap"
description: "support for Roaring Bitmaps"
weight: 3630
categories: ["TYPE"]
width: full
---

[**pg_roaringbitmap**](https://github.com/ChenHuajun/pg_roaringbitmap) : support for Roaring Bitmaps


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3630** | {{< badge content="roaringbitmap" link="https://github.com/ChenHuajun/pg_roaringbitmap" >}} | {{< ext "roaringbitmap" "pg_roaringbitmap" >}} | `1.2.0` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


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
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_roaringbitmap` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.0` | {{< bg "18" "pg_roaringbitmap_18" "green" >}} {{< bg "17" "pg_roaringbitmap_17" "green" >}} {{< bg "16" "pg_roaringbitmap_16" "green" >}} {{< bg "15" "pg_roaringbitmap_15" "green" >}} {{< bg "14" "pg_roaringbitmap_14" "green" >}} | `pg_roaringbitmap_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2.0` | {{< bg "18" "postgresql-18-roaringbitmap" "green" >}} {{< bg "17" "postgresql-17-roaringbitmap" "green" >}} {{< bg "16" "postgresql-16-roaringbitmap" "green" >}} {{< bg "15" "postgresql-15-roaringbitmap" "green" >}} {{< bg "14" "postgresql-14-roaringbitmap" "green" >}} | `postgresql-$v-roaringbitmap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 6" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 6" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 7" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 7" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 7" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_17 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_16 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_15 : AVAIL 7" "green" >}} | {{< bg "PIGSTY 1.2.0" "pg_roaringbitmap_14 : AVAIL 7" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-18-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-17-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-16-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-15-roaringbitmap : AVAIL 3" "green" >}} | {{< bg "PIGSTY 1.2.0" "postgresql-14-roaringbitmap : AVAIL 3" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 183.6 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 172.2 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.3 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 170.9 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 153.9 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.3 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.2 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.2 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.5 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 103.2 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.2 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.4 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.4 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.0 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.5 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 120.2 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.3 KiB | [pg_roaringbitmap_18-1.2.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roaringbitmap_18-1.2.0-2PIGSTY.el10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.5 KiB | [pg_roaringbitmap_18-1.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.9 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_18` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.4 KiB | [pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_roaringbitmap_18-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 465.7 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.9 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.5 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 424.5 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 423.2 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.1 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 468.7 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.7 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 465.0 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 427.9 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.9 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 421.9 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 448.9 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 420.3 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 416.0 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 418.4 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.6 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 384.8 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 436.6 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 412.0 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 409.3 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 408.2 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 380.6 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 378.1 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 431.8 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 409.0 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 406.0 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 401.9 KiB | [postgresql-18-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-18-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 377.2 KiB | [postgresql-18-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-roaringbitmap` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 373.1 KiB | [postgresql-18-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-18-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 183.6 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 172.2 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.0 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 153.9 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.3 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.2 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.2 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 118.2 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.5 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 103.2 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.2 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.9 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.7 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.5 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 120.2 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.4 KiB | [pg_roaringbitmap_17-1.2.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roaringbitmap_17-1.2.0-2PIGSTY.el10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.6 KiB | [pg_roaringbitmap_17-1.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_17` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_roaringbitmap_17-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 466.7 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 463.5 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.0 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 424.3 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 421.6 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.0 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 467.9 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 465.4 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.1 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 426.6 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.8 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.5 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.6 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 450.8 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 448.3 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 446.8 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 416.7 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 413.2 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 436.6 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 412.6 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 410.4 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 408.2 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 380.6 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 378.1 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 431.7 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 408.9 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 405.7 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 401.6 KiB | [postgresql-17-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-17-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 376.7 KiB | [postgresql-17-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-roaringbitmap` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 373.7 KiB | [postgresql-17-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-17-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 183.6 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 172.2 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 171.0 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 159.1 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 153.9 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.2 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.1 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 133.7 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 116.3 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 118.2 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.4 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 116.2 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 84.8 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 72.0 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 103.2 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.2 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.4 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 103.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 102.9 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 72.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 69.9 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 118.5 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 120.2 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.3 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 118.4 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 86.6 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 82.3 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 104.4 KiB | [pg_roaringbitmap_16-1.2.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roaringbitmap_16-1.2.0-2PIGSTY.el10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 106.6 KiB | [pg_roaringbitmap_16-1.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.8 KiB | [pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 73.5 KiB | [pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_16` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 71.2 KiB | [pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_roaringbitmap_16-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 465.3 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 463.5 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 462.6 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 424.2 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 421.7 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 420.0 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 468.7 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.6 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 464.8 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 426.4 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.9 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 423.5 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 480.5 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 451.2 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 447.7 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 446.7 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 416.6 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 413.3 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 436.6 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 412.6 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 410.3 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 408.2 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 380.8 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 379.1 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 431.7 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 408.1 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 405.7 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 401.6 KiB | [postgresql-16-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-16-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 376.8 KiB | [postgresql-16-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-roaringbitmap` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 373.6 KiB | [postgresql-16-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-16-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.9 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 175.9 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.6 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 159.1 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.7 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.8 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 178.0 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 176.6 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.8 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.6 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.6 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 157.5 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 156.3 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.7 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.8 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 154.4 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.7 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 182.7 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.8 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.5 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.5 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.5 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 159.8 KiB | [pg_roaringbitmap_15-1.2.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roaringbitmap_15-1.2.0-2PIGSTY.el10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 158.8 KiB | [pg_roaringbitmap_15-1.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.7 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.6 KiB | [pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 158.1 KiB | [pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_15` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_roaringbitmap_15-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 469.4 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 466.3 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 464.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 427.8 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 425.9 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 423.5 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 471.5 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 468.2 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 467.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 429.4 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 427.0 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 425.3 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 536.5 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 500.2 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 499.7 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 495.9 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 458.6 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 455.0 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 492.4 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 462.9 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 462.4 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 457.6 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 424.1 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 422.1 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 489.8 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 462.2 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 460.1 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 452.9 KiB | [postgresql-15-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-15-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 421.3 KiB | [postgresql-15-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-roaringbitmap` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 418.3 KiB | [postgresql-15-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-15-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_roaringbitmap_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 189.9 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 175.9 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 174.5 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 162.2 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 109.1 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 159.1 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.7 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.8 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 135.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 178.1 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 176.5 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.2 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 175.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 174.9 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 161.6 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 114.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 157.4 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 156.3 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.4 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 155.2 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 105.3 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel9.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 182.7 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 182.0 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.4 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 180.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 181.8 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 167.3 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 121.8 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.x86_64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 159.4 KiB | [pg_roaringbitmap_14-1.2.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_roaringbitmap_14-1.2.0-2PIGSTY.el10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.2.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 158.8 KiB | [pg_roaringbitmap_14-1.2.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.2.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.7 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.6 KiB | [pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.1.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 157.7 KiB | [pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-1.0.0-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 145.4 KiB | [pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.5-1PGDG.rhel10.aarch64.rpm) |
| `pg_roaringbitmap_14` | `0.5.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.0 KiB | [pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_roaringbitmap_14-0.5.4-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 468.6 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 467.1 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 464.0 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 427.3 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 426.1 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 423.6 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 470.6 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 468.5 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 466.9 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 429.6 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 427.0 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 425.2 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 536.6 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 502.2 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 498.1 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 495.8 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 458.1 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 455.4 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 492.3 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 463.7 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 462.1 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 457.5 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~noble_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 423.3 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 421.9 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 489.7 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~resolute_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 462.1 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 459.9 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 452.8 KiB | [postgresql-14-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/roaringbitmap/postgresql-14-roaringbitmap_1.2.0-2PIGSTY~resolute_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.2.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 421.0 KiB | [postgresql-14-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.2.0-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-roaringbitmap` | `1.1.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 419.0 KiB | [postgresql-14-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-roaringbitmap/postgresql-14-roaringbitmap_1.1.0-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ChenHuajun/pg_roaringbitmap" title="Repository" icon="github" subtitle="github.com/ChenHuajun/pg_roaringbitmap" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_roaringbitmap-1.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_roaringbitmap;		# build rpm
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

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION roaringbitmap;
```




## Usage

Sources:

- [PGXN pg_roaringbitmap 1.2.0](https://pgxn.org/dist/pg_roaringbitmap/1.2.0/)
- [pg_roaringbitmap README](https://github.com/ChenHuajun/pg_roaringbitmap)
- [pg_roaringbitmap CHANGELOG](https://github.com/ChenHuajun/pg_roaringbitmap/blob/master/CHANGELOG.md)
- [Local package metadata](../db/extension.csv)

`pg_roaringbitmap` installs the PostgreSQL extension `roaringbitmap`, which provides compressed bitmap types and set-operation functions backed by Roaring Bitmaps. Use it for compact integer-set storage, fast unions/intersections, cohort filters, faceting, and bitmap aggregation.

v1.2.0 adds `rb_runoptimize()` / `rb64_runoptimize()` to shrink bitmap binary size, preserves the legacy `rb_exsit` spelling for backward compatibility, adds PostgreSQL 19 compatibility, and validates untrusted bitmap input in receive functions.

### Create the Extension

```sql
CREATE EXTENSION IF NOT EXISTS roaringbitmap;

SET roaringbitmap.output_format = 'array';
SELECT rb_build(ARRAY[1, 2, 3, 4, 5]);
```

`roaringbitmap.output_format` can be `bytea` or `array`. The default output format is `bytea`, which is better for large bitmaps; `array` is easier to read interactively.

### Data Types

- `roaringbitmap` stores unsigned 32-bit integer sets over the logical range `[0, 4294967296)`.
- `roaringbitmap64` stores unsigned 64-bit integer sets and uses the `rb64_` function family.

```sql
CREATE TABLE cohorts (
  segment text PRIMARY KEY,
  users32 roaringbitmap,
  users64 roaringbitmap64
);
```

### Build and Convert

```sql
INSERT INTO cohorts(segment, users32)
VALUES ('trial', rb_build(ARRAY[1, 2, 3, 100, 200]));

INSERT INTO cohorts(segment, users32)
SELECT 'active', rb_build_agg(id)
FROM generate_series(1, 100000) AS id;

SELECT rb_cardinality(users32) FROM cohorts WHERE segment = 'active';
SELECT rb_to_array(users32) FROM cohorts WHERE segment = 'trial';
SELECT rb_iterate(users32) FROM cohorts WHERE segment = 'trial';
```

For 64-bit values, use `rb64_build()`, `rb64_build_agg()`, `rb64_to_array()`, and `rb64_iterate()`.

### Set Operations

```sql
SELECT rb_build(ARRAY[1,2,3]) | rb_build(ARRAY[3,4,5]);  -- union
SELECT rb_build(ARRAY[1,2,3]) & rb_build(ARRAY[3,4,5]);  -- intersection
SELECT rb_build(ARRAY[1,2,3]) # rb_build(ARRAY[3,4,5]);  -- xor
SELECT rb_build(ARRAY[1,2,3]) - rb_build(ARRAY[3,4,5]);  -- difference

SELECT rb_build(ARRAY[1,2,3]) | 9;                       -- add element
SELECT rb_build(ARRAY[1,2,3]) - 2;                       -- remove element
```

Containment and overlap operators are available:

```sql
SELECT rb_build(ARRAY[1,2,3]) @> rb_build(ARRAY[2,3]);
SELECT rb_build(ARRAY[2,3]) <@ rb_build(ARRAY[1,2,3]);
SELECT rb_build(ARRAY[1,2,3]) && rb_build(ARRAY[3,4,5]);
```

### Cardinality and Range Helpers

```sql
SELECT rb_and_cardinality(a.users32, b.users32);
SELECT rb_or_cardinality(a.users32, b.users32);
SELECT rb_xor_cardinality(a.users32, b.users32);
SELECT rb_andnot_cardinality(a.users32, b.users32);

SELECT rb_range(users32, 100, 1000);
SELECT rb_range_cardinality(users32, 100, 1000);
SELECT rb_fill(users32, 100, 200);
SELECT rb_clear(users32, 100, 200);
SELECT rb_flip(users32, 100, 200);

SELECT rb_min(users32), rb_max(users32), rb_rank(users32, 500), rb_index(users32, 500);
SELECT rb_jaccard_dist(a.users32, b.users32);
```

The 64-bit range helpers use the `rb64_` prefix. Since v1.1.0, `range_end = 0` means unlimited for several `rb64_` range/select functions.

### Aggregate Functions

```sql
SELECT rb_build_agg(user_id) FROM events;
SELECT rb_or_agg(users32) FROM cohorts;
SELECT rb_and_agg(users32) FROM cohorts;
SELECT rb_xor_agg(users32) FROM cohorts;

SELECT rb64_build_agg(user_id::bigint) FROM events;
SELECT rb64_or_agg(users64) FROM cohorts;
```

### Optimize Serialized Size

```sql
UPDATE cohorts
SET users32 = rb_runoptimize(users32);

UPDATE cohorts
SET users64 = rb64_runoptimize(users64);
```

`rb_runoptimize()` and `rb64_runoptimize()` can reduce serialized bitmap size for suitable data distributions. Measure before using them in hot write paths.

### Caveats

- Pigsty uses extension file name `roaringbitmap.md`; the upstream package name is `pg_roaringbitmap`.
- Source builds require PostgreSQL headers and CRoaring dependencies. The README notes regression testing before release covers PostgreSQL 13 and above.
- `Makefile_native` can compile with SIMD instructions and may be faster on matching CPUs, but binaries built that way can crash with `SIGILL` on machines without the required CPU features.
- Use `bytea` output for large bitmaps and `array` output for human inspection.
