---
title: "pg_eviltransform"
linkTitle: "pg_eviltransform"
description: "Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform"
weight: 1580
categories: ["GIS"]
width: full
---

[**pg_eviltransform**](https://github.com/aiyou178/pg_eviltransform) : Coordinate transforms for BD09/GCJ02 via PostGIS ST_Transform


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1580** | {{< badge content="pg_eviltransform" link="https://github.com/aiyou178/pg_eviltransform" >}} | {{< ext "pg_eviltransform" >}} | `0.0.2` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `eviltransform_internal` |
|   **Requires**    | {{< ext "postgis" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "earthdistance" >}} {{< ext "tzf" >}} {{< ext "geoip" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_eviltransform` | `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "pg_eviltransform_18" "green" >}} {{< bg "17" "pg_eviltransform_17" "green" >}} {{< bg "16" "pg_eviltransform_16" "green" >}} {{< bg "15" "pg_eviltransform_15" "green" >}} {{< bg "14" "pg_eviltransform_14" "green" >}} | `pg_eviltransform_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.2` | {{< bg "18" "postgresql-18-eviltransform" "green" >}} {{< bg "17" "postgresql-17-eviltransform" "green" >}} {{< bg "16" "postgresql-16-eviltransform" "green" >}} {{< bg "15" "postgresql-15-eviltransform" "green" >}} {{< bg "14" "postgresql-14-eviltransform" "green" >}} | `postgresql-$v-eviltransform` | `postgresql-$v-postgis` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "pg_eviltransform_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-18-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-17-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-16-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-15-eviltransform : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.2" "postgresql-14-eviltransform : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-eviltransform : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-eviltransform : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-eviltransform : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_18` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 300.7 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_18-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 194.6 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_18-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 315.6 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_18-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 207.5 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_18-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_18` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 316.2 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_18-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_18` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 208.3 KiB | [pg_eviltransform_18-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_18-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-eviltransform` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 249.0 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 150.6 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.7 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.8 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.4 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.5 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 277.8 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-eviltransform` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.1 KiB | [postgresql-18-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-18-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_17` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 300.7 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_17-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 194.6 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_17-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 315.5 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_17-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 207.5 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_17-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_17` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 315.9 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_17-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_17` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 208.1 KiB | [pg_eviltransform_17-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_17-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-eviltransform` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.5 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 150.6 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.6 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.7 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.1 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.5 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 277.7 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-eviltransform` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.2 KiB | [postgresql-17-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-17-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_16` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 300.5 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_16-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 194.6 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_16-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 315.5 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_16-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 207.6 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_16-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_16` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 315.9 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_16-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_16` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 208.5 KiB | [pg_eviltransform_16-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_16-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-eviltransform` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.3 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 150.7 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 247.9 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.8 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.2 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.5 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 277.8 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-eviltransform` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.1 KiB | [postgresql-16-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-16-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_15` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 300.4 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_15-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 194.7 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_15-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 315.5 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_15-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 207.5 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_15-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_15` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 315.8 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_15-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_15` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 208.2 KiB | [pg_eviltransform_15-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_15-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-eviltransform` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 249.1 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 150.6 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.2 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.8 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.2 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.5 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 277.7 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-eviltransform` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.2 KiB | [postgresql-15-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-15-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_eviltransform_14` | `0.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 300.3 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_eviltransform_14-0.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 194.6 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_eviltransform_14-0.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 315.1 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_eviltransform_14-0.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 207.6 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_eviltransform_14-0.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_eviltransform_14` | `0.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 315.7 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_eviltransform_14-0.0.2-1PIGSTY.el10.x86_64.rpm) |
| `pg_eviltransform_14` | `0.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 208.1 KiB | [pg_eviltransform_14-0.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_eviltransform_14-0.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-eviltransform` | `0.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 248.0 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 150.7 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 248.8 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 150.6 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 280.1 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.5 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 277.7 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-eviltransform` | `0.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 173.2 KiB | [postgresql-14-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-eviltransform/postgresql-14-eviltransform_0.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aiyou178/pg_eviltransform" title="Repository" icon="github" subtitle="github.com/aiyou178/pg_eviltransform" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_eviltransform-0.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_eviltransform;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

> [pg_eviltransform: Coordinate transform between WGS84, GCJ02, and BD09](https://github.com/aiyou178/pg_eviltransform)

`pg_eviltransform` extends PostGIS `ST_Transform` with BD09/GCJ02 support for Chinese coordinate systems. It exposes `ST_EvilTransform` with the same overload interface as `ST_Transform`.

Custom SRIDs:
- `990001`: GCJ02
- `990002`: BD09

### Functions

```sql
ST_EvilTransform(geometry, to_srid integer)
ST_EvilTransform(geometry, to_proj text)
ST_EvilTransform(geometry, from_proj text, to_srid integer)
ST_EvilTransform(geometry, from_proj text, to_proj text)
```

If neither side uses custom coordinates, it delegates directly to `ST_Transform`. If BD09/GCJ02 is involved, it transforms via WGS84 (`4326`) when needed.

### Examples

```sql
-- WGS84 to GCJ02 using text literal
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'GCJ02');

-- WGS84 to BD09 using text literal
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'BD09');

-- WGS84 to GCJ02 using numeric SRID
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 990001);

-- BD09 to Web Mercator
SELECT ST_EvilTransform(
  ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002), 3857
);

-- from_proj / to_proj overload
SELECT ST_EvilTransform('POINT(120 30)'::geometry, 'EPSG:4326', 'GCJ02');
```

### Performance

On PG18 with 200,000 rows, `ST_EvilTransform` is ~30-45x faster than the regex-based SQL approach.
