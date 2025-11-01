---
title: "pg_polyline"
linkTitle: "pg_polyline"
description: "Fast Google Encoded Polyline encoding & decoding for postgres"
weight: 1570
categories: ["GIS"]
width: full
---

Fast Google Encoded Polyline encoding & decoding for postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1570** | {{< badge content="pg_polyline" link="https://github.com/yihong0618/pg_polyline" >}} | {{< ext "pg_polyline" >}} | `0.0.1` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_polyline" >}} | `0.0.1` | {{< bg "18" "pg_polyline_18" "green" >}} {{< bg "17" "pg_polyline_17" "green" >}} {{< bg "16" "pg_polyline_16" "green" >}} {{< bg "15" "pg_polyline_15" "green" >}} {{< bg "14" "pg_polyline_14" "green" >}} {{< bg "13" "pg_polyline_13" "green" >}} | `pg_polyline_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_polyline" >}} | `0.0.1` | {{< bg "18" "postgresql-18-pg-polyline" "green" >}} {{< bg "17" "postgresql-17-pg-polyline" "green" >}} {{< bg "16" "postgresql-16-pg-polyline" "green" >}} {{< bg "15" "postgresql-15-pg-polyline" "green" >}} {{< bg "14" "postgresql-14-pg-polyline" "green" >}} {{< bg "13" "postgresql-13-pg-polyline" "green" >}} | `postgresql-$v-pg-polyline` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `el10.aarch64`    | {{< bg "PIGSTY 0.0.1" "pg_polyline_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "pg_polyline_13 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-polyline : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-polyline : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-pg-polyline : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-polyline : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.0.1" "postgresql-17-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-16-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-15-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-14-pg-polyline : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.1" "postgresql-13-pg-polyline : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_18` | 0.0.1 | `el8.x86_64` | pigsty | 312.0 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_18-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_18` | 0.0.1 | `el8.aarch64` | pigsty | 205.2 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_18-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_18` | 0.0.1 | `el9.x86_64` | pigsty | 326.6 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_18-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_18` | 0.0.1 | `el9.aarch64` | pigsty | 218.7 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_18-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_18` | 0.0.1 | `el10.x86_64` | pigsty | 326.2 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_18-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_18` | 0.0.1 | `el10.aarch64` | pigsty | 218.7 KiB | [pg_polyline_18-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_18-0.0.1-2PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_17` | 0.0.1 | `el8.x86_64` | pigsty | 311.9 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_17-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el8.aarch64` | pigsty | 205.2 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_17-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el9.x86_64` | pigsty | 326.6 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_17-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el9.aarch64` | pigsty | 219.2 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_17-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el10.x86_64` | pigsty | 326.2 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_17-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el10.aarch64` | pigsty | 218.3 KiB | [pg_polyline_17-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_17-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.5 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_16` | 0.0.1 | `el8.x86_64` | pigsty | 311.9 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_16-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el8.aarch64` | pigsty | 205.1 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_16-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el9.x86_64` | pigsty | 326.7 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_16-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el9.aarch64` | pigsty | 219.1 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_16-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el10.x86_64` | pigsty | 326.3 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_16-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el10.aarch64` | pigsty | 218.3 KiB | [pg_polyline_16-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_16-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.5 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.6 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_15` | 0.0.1 | `el8.x86_64` | pigsty | 311.7 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_15-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el8.aarch64` | pigsty | 205.2 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_15-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el9.x86_64` | pigsty | 326.5 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_15-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el9.aarch64` | pigsty | 218.7 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_15-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el10.x86_64` | pigsty | 325.7 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_15-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el10.aarch64` | pigsty | 218.3 KiB | [pg_polyline_15-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_15-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.3 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_14` | 0.0.1 | `el8.x86_64` | pigsty | 311.3 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_14-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el8.aarch64` | pigsty | 205.2 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_14-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el9.x86_64` | pigsty | 326.1 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_14-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el9.aarch64` | pigsty | 219.1 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_14-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el10.x86_64` | pigsty | 325.8 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_14-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el10.aarch64` | pigsty | 218.6 KiB | [pg_polyline_14-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_14-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.4 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.4 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.5 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_polyline_13` | 0.0.1 | `el8.x86_64` | pigsty | 311.8 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_13-0.0.1-2PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el8.aarch64` | pigsty | 205.1 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_13-0.0.1-2PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el9.x86_64` | pigsty | 326.8 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_13-0.0.1-2PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el9.aarch64` | pigsty | 218.7 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_13-0.0.1-2PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el10.x86_64` | pigsty | 326.3 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_polyline_13-0.0.1-2PIGSTY.el10.x86_64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el10.aarch64` | pigsty | 218.7 KiB | [pg_polyline_13-0.0.1-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_polyline_13-0.0.1-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.3 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.4 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.6 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/yihong0618/pg_polyline" title="Repository" icon="github" subtitle="github.com/yihong0618/pg_polyline" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_polyline-0.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_polyline; # get pg_polyline source code
pig build dep pg_polyline; # install build dependencies
pig build pkg pg_polyline; # build extension rpm or deb
pig build ext pg_polyline; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_polyline; # install by extension name, for the current active PG version
pig ext install pg_polyline; # install via package alias, for the active PG version
pig ext install pg_polyline -v 18;   # install for PG 18
pig ext install pg_polyline -v 17;   # install for PG 17
pig ext install pg_polyline -v 16;   # install for PG 16
pig ext install pg_polyline -v 15;   # install for PG 15
pig ext install pg_polyline -v 14;   # install for PG 14
pig ext install pg_polyline -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_polyline;
```

