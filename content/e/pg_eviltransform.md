---
title: "pg_eviltransform"
linkTitle: "pg_eviltransform"
description: "Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform"
weight: 1580
categories: ["GIS"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_eviltransform**](https://github.com/aiyou178/pg_eviltransform) : Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1580** | {{< badge content="pg_eviltransform" link="https://github.com/aiyou178/pg_eviltransform" >}} | {{< ext "pg_eviltransform" >}} | `0.0.7` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `eviltransform_internal` |
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "h3" >}} {{< ext "pg_geohash" >}} {{< ext "pg_polyline" >}} {{< ext "earthdistance" >}} {{< ext "qdgc" >}} {{< ext "convert" >}} {{< ext "pgrouting" >}} {{< ext "nominatim_fdw" >}} {{< ext "q3c" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_eviltransform` | `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "pg_eviltransform_18" "green" >}} {{< bg "17" "pg_eviltransform_17" "green" >}} {{< bg "16" "pg_eviltransform_16" "green" >}} {{< bg "15" "pg_eviltransform_15" "green" >}} {{< bg "14" "pg_eviltransform_14" "green" >}} | `pg_eviltransform_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "postgresql-18-eviltransform" "green" >}} {{< bg "17" "postgresql-17-eviltransform" "green" >}} {{< bg "16" "postgresql-16-eviltransform" "green" >}} {{< bg "15" "postgresql-15-eviltransform" "green" >}} {{< bg "14" "postgresql-14-eviltransform" "green" >}} | `postgresql-$v-eviltransform` | `postgresql-$v-postgis` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_18` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 895.7 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_18-0.0.7-1PGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 767.4 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_18-0.0.7-1PGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 912.9 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_18-0.0.7-1PGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 824.4 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_18-0.0.7-1PGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 912.9 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_18-0.0.7-1PGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 799.7 KiB | [pg_eviltransform_18-0.0.7-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_18-0.0.7-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-eviltransform` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 745.5 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 617.1 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 744.9 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 617.1 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 827.6 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 730.4 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 820.8 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~noble_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 721.2 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~noble_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 814.5 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 719.8 KiB | [postgresql-18-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_17` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 891.5 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_17-0.0.7-1PGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 763.8 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_17-0.0.7-1PGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 909.2 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_17-0.0.7-1PGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 820.9 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_17-0.0.7-1PGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 905.0 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_17-0.0.7-1PGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 798.8 KiB | [pg_eviltransform_17-0.0.7-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_17-0.0.7-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-eviltransform` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 742.4 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 614.9 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 742.9 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 614.9 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 825.8 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 727.1 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 818.7 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~noble_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 717.8 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~noble_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 812.3 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 715.7 KiB | [postgresql-17-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_16` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 890.3 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_16-0.0.7-1PGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 762.5 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_16-0.0.7-1PGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 907.7 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_16-0.0.7-1PGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 817.8 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_16-0.0.7-1PGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 907.8 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_16-0.0.7-1PGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 798.6 KiB | [pg_eviltransform_16-0.0.7-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_16-0.0.7-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-eviltransform` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 742.1 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 612.8 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 742.5 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 614.1 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 823.9 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 725.3 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 817.5 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~noble_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 717.0 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~noble_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 813.2 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 715.2 KiB | [postgresql-16-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_15` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 881.4 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_15-0.0.7-1PGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 753.7 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_15-0.0.7-1PGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 897.5 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_15-0.0.7-1PGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 807.8 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_15-0.0.7-1PGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 898.3 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_15-0.0.7-1PGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 796.1 KiB | [pg_eviltransform_15-0.0.7-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_15-0.0.7-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-eviltransform` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 736.9 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 609.0 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 736.7 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 609.3 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 816.8 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 718.5 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 809.2 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~noble_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 709.6 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~noble_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 803.5 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 708.2 KiB | [postgresql-15-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_14` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 879.4 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_14-0.0.7-1PGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 752.1 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_14-0.0.7-1PGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 895.6 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_14-0.0.7-1PGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 805.5 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_14-0.0.7-1PGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 893.6 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_14-0.0.7-1PGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 793.7 KiB | [pg_eviltransform_14-0.0.7-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_14-0.0.7-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-eviltransform` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 733.8 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 607.4 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 734.3 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 607.3 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 812.8 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 715.9 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 807.0 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~noble_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 708.1 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~noble_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u26.x86_64](/os/u26.x86_64) | pigsty | 804.1 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.7` | [u26.aarch64](/os/u26.aarch64) | pigsty | 706.7 KiB | [postgresql-14-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.7-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aiyou178/pg_eviltransform" title="Repository" icon="github" subtitle="github.com/aiyou178/pg_eviltransform" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_eviltransform-0.0.7.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_eviltransform;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_eviltransform;		# install via package name, for the active PG version

pig install pg_eviltransform -v 18;   # install for PG 18
pig install pg_eviltransform -v 17;   # install for PG 17
pig install pg_eviltransform -v 16;   # install for PG 16
pig install pg_eviltransform -v 15;   # install for PG 15
pig install pg_eviltransform -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_eviltransform CASCADE; -- requires postgis
```

## Usage

Sources:

- [Official v0.0.7 README](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/README.md)
- [v0.0.7 release notes](https://github.com/aiyou178/pg_eviltransform/releases/tag/v0.0.7)
- [v0.0.7 control file](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/pg_eviltransform.control)
- [v0.0.7 upgrade SQL](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.7/pg_eviltransform--0.0.6--0.0.7.sql)

`pg_eviltransform` extends PostGIS with coordinate transformations involving China's GCJ-02 and BD-09 systems. Version `0.0.4` also adds exact Jenks natural-break classification through `ST_JenksBins` array and aggregate overloads.

### Coordinate Transformation

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_eviltransform;

-- WGS84 to GCJ-02 using a readable coordinate-system name.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326),
    'GCJ02'
);

-- BD-09 to Web Mercator.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002),
    3857
);
```

Custom SRIDs are `990001` for GCJ-02 and `990002` for BD-09. When neither endpoint uses a custom system, `ST_EvilTransform` delegates to PostGIS `ST_Transform`; otherwise it converts through WGS84 (`4326`) when necessary.

Version 0.0.7 adds the geographic CGCS2000 aliases `CGCS2000`, `CGCS-2000`, `4490`, and `EPSG:4490` to text overloads. PostGIS/PROJ chooses the datum transformation. For a projected CGCS2000 Gauss-Kruger zone, use its explicit integer EPSG SRID.

```sql
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326), 'CGCS2000'
);
```

### Jenks Natural Breaks

```sql
-- Array form; NULL elements are ignored.
SELECT ST_JenksBins(ARRAY[1, 2, NULL, 10, 11]::numeric[], 2);

-- Streaming aggregate form for a large table.
SELECT ST_JenksBins(value, 7)
FROM measurements;

-- Return lower rather than upper bin edges.
SELECT ST_JenksBins(value, 7, true)
FROM measurements;
```

Array inputs support `numeric`, `double precision`, `real`, `bigint`, `integer`, and `smallint`. Aggregate inputs are `numeric` or `double precision`; cast other numeric columns when needed.

### API Index and Caveats

- `ST_EvilTransform(geometry, integer|text)` and `ST_EvilTransform(geometry, text, integer|text)`: four overloads corresponding to the PostGIS `ST_Transform` interface.
- `ST_JenksBins(values[], breaks [, invert])`: classifies an array and returns `double precision[]` edges.
- `ST_JenksBins(value, breaks [, invert])`: streaming aggregate that avoids materializing `array_agg`.
- PostGIS is a runtime prerequisite and must be installed before `pg_eviltransform`.
- Jenks inputs must be finite and `breaks` must be at least one. `numeric` values are converted to finite `f64`, so returned edges are floating-point values.
- When the distinct value count does not exceed `breaks`, the result is the sorted set of unique values; no valid input rows return `NULL`.

### 0.0.7 Compatibility

The 0.0.6-to-0.0.7 upgrade SQL updates the coordinate-name resolution functions. Upstream supports PostgreSQL 14-18 and PostgreSQL 19 beta4 with pgrx 0.19.3. The `ST_JenksBins` interface remains unchanged; package/build metadata is maintained separately from the upstream release.
