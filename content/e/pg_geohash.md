---
title: "pg_geohash"
linkTitle: "pg_geohash"
description: "Handle geohash based functionality for spatial coordinates"
weight: 1590
categories: ["GIS"]
width: full
---

[**pg_geohash**](https://github.com/jistok/pg_geohash) : Handle geohash based functionality for spatial coordinates


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1590** | {{< badge content="pg_geohash" link="https://github.com/jistok/pg_geohash" >}} | {{< ext "pg_geohash" >}} | `1.0` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "pg_polyline" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_geohash` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_geohash_18*" "green" >}} {{< bg "17" "pg_geohash_17*" "green" >}} {{< bg "16" "pg_geohash_16*" "green" >}} {{< bg "15" "pg_geohash_15*" "green" >}} {{< bg "14" "pg_geohash_14*" "green" >}} {{< bg "13" "pg_geohash_13*" "green" >}} | `pg_geohash_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-geohash" "green" >}} {{< bg "17" "postgresql-17-pg-geohash" "green" >}} {{< bg "16" "postgresql-16-pg-geohash" "green" >}} {{< bg "15" "postgresql-15-pg-geohash" "green" >}} {{< bg "14" "postgresql-14-pg-geohash" "green" >}} {{< bg "13" "postgresql-13-pg-geohash" "green" >}} | `postgresql-$v-pg-geohash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_geohash_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-geohash : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-13-pg-geohash : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [pg_geohash_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_geohash_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.9 KiB | [pg_geohash_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_geohash_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_geohash_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_geohash_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.3 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-18-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-18-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [pg_geohash_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_geohash_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.9 KiB | [pg_geohash_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_geohash_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_geohash_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.0 KiB | [pg_geohash_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [pg_geohash_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.9 KiB | [pg_geohash_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.9 KiB | [pg_geohash_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.7 KiB | [pg_geohash_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.9 KiB | [pg_geohash_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.9 KiB | [pg_geohash_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.3 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_geohash_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [pg_geohash_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_geohash_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_geohash_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_geohash_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.1 KiB | [pg_geohash_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.1 KiB | [pg_geohash_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [pg_geohash_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_geohash_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_geohash_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_geohash_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.1 KiB | [pg_geohash_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_geohash_13` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.1 KiB | [pg_geohash_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_13` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [pg_geohash_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_13` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_geohash_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_13` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [pg_geohash_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_13` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.1 KiB | [pg_geohash_13-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_geohash_13-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_geohash_13` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.1 KiB | [pg_geohash_13-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_geohash_13-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-geohash` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.4 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.2 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.5 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.2 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.2 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 15.8 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.3 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-geohash` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.0 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jistok/pg_geohash" title="Repository" icon="github" subtitle="github.com/jistok/pg_geohash" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_geohash-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_geohash;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_geohash;		# install via package name, for the active PG version

pig install pg_geohash -v 18;   # install for PG 18
pig install pg_geohash -v 17;   # install for PG 17
pig install pg_geohash -v 16;   # install for PG 16
pig install pg_geohash -v 15;   # install for PG 15
pig install pg_geohash -v 14;   # install for PG 14
pig install pg_geohash -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_geohash;
```
