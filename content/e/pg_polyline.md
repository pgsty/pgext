---
title: "pg_polyline"
linkTitle: "pg_polyline"
description: "Fast Google Encoded Polyline encoding & decoding for postgres"
weight: 1570
categories: ["GIS"]
width: full
---

[**pg_polyline**](https://github.com/yihong0618/pg_polyline) : Fast Google Encoded Polyline encoding & decoding for postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1570** | {{< badge content="pg_polyline" link="https://github.com/yihong0618/pg_polyline" >}} | {{< ext "pg_polyline" >}} | `0.0.1` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_polyline` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "pg_polyline_18" "green" >}} {{< bg "17" "pg_polyline_17" "green" >}} {{< bg "16" "pg_polyline_16" "green" >}} {{< bg "15" "pg_polyline_15" "green" >}} {{< bg "14" "pg_polyline_14" "green" >}} | `pg_polyline_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-polyline" "green" >}} {{< bg "17" "postgresql-17-pg-polyline" "green" >}} {{< bg "16" "postgresql-16-pg-polyline" "green" >}} {{< bg "15" "postgresql-15-pg-polyline" "green" >}} {{< bg "14" "postgresql-14-pg-polyline" "green" >}} | `postgresql-$v-pg-polyline` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-18-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_18` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 855.9 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_18-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_18` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 768.4 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_18-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_18` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 865.7 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_18-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_18` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 814.0 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_18-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_18` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 865.6 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_18-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_18` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 793.1 KiB | [pg_polyline_18-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_18-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-polyline` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 685.2 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 571.8 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 684.8 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 571.9 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 761.8 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 672.4 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 752.4 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 666.0 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 751.2 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-polyline` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 664.7 KiB | [postgresql-18-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-18-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_17` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 853.3 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_17-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_17` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 765.2 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_17-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_17` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 862.6 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_17-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_17` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 811.8 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_17-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_17` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 863.4 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_17-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_17` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 791.8 KiB | [pg_polyline_17-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_17-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-polyline` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 683.5 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 570.3 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 684.4 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 570.7 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 759.2 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 672.6 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 752.4 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 663.8 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 748.5 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-polyline` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 662.2 KiB | [postgresql-17-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_16` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 851.6 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_16-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_16` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 763.8 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_16-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_16` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 860.7 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_16-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_16` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 809.7 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_16-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_16` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 861.2 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_16-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_16` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 789.7 KiB | [pg_polyline_16-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_16-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-polyline` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 682.9 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 569.5 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 683.2 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 570.0 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 758.1 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 671.0 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 751.5 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 662.5 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 747.5 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-polyline` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 661.1 KiB | [postgresql-16-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_15` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 841.8 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_15-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_15` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 755.0 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_15-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_15` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 850.8 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_15-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_15` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 799.3 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_15-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_15` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 850.6 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_15-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_15` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 786.9 KiB | [pg_polyline_15-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_15-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-polyline` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 676.3 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 565.4 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 677.2 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 565.5 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 752.0 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 666.4 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 744.2 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 657.4 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 740.6 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-polyline` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 655.8 KiB | [postgresql-15-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_14` | `0.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 839.3 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_14-0.0.1-4PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_14` | `0.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 752.0 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_14-0.0.1-4PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_14` | `0.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 848.7 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_14-0.0.1-4PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_14` | `0.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 795.7 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_14-0.0.1-4PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_14` | `0.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 849.0 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_14-0.0.1-4PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_14` | `0.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 786.9 KiB | [pg_polyline_14-0.0.1-4PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_14-0.0.1-4PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-polyline` | `0.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 674.4 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 564.0 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 674.4 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 564.0 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 751.2 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 664.5 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 742.2 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 654.8 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 738.7 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-polyline` | `0.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 653.9 KiB | [postgresql-14-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/yihong0618/pg_polyline" title="Repository" icon="github" subtitle="github.com/yihong0618/pg_polyline" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_polyline-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_polyline;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_polyline;		# install via package name, for the active PG version

pig install pg_polyline -v 18;   # install for PG 18
pig install pg_polyline -v 17;   # install for PG 17
pig install pg_polyline -v 16;   # install for PG 16
pig install pg_polyline -v 15;   # install for PG 15
pig install pg_polyline -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_polyline;
```




## Usage

> [pg_polyline: Google Encoded Polyline encoding & decoding for PostgreSQL](https://github.com/yihong0618/pg_polyline)

Fast Google Encoded Polyline encoding & decoding as a PostgreSQL extension. Built with `pgrx`.

```sql
CREATE EXTENSION pg_polyline;

-- Encode an array of points to a polyline string
SELECT polyline_encode(
  ARRAY[point(-120.2, 38.5), point(-120.95, 40.7), point(-126.453, 43.252)], 6
);
--          polyline_encode
-- ----------------------------------
--  _izlhA~rlgdF_{geC~ywl@_kwzCn`{nI

-- Decode a polyline string back to points
SELECT polyline_decode('_ibE_seK_seK_seK', 6);
--       polyline_decode
-- ---------------------------
--  {"(0.2,0.1)","(0.4,0.3)"}
```

The second parameter is the precision (number of decimal places).
