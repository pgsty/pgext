---
title: "pg_geohash"
linkTitle: "pg_geohash"
description: "Handle geohash based functionality for spatial coordinates"
weight: 1590
categories: ["Gis"]
width: full
---

Handle geohash based functionality for spatial coordinates

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1590** | {{< badge content="pg_geohash" link="https://github.com/jistok/pg_geohash" >}} | {{< ext "pg_geohash" "pg_geohash" >}} | `1.0` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "h3" >}} {{< ext "q3c" >}} {{< ext "pg_polyline" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_geohash" >}} | `1.0` | {{< badge content="18" color="red" alt="pg_geohash_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_geohash_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_geohash" >}} | `1.0` | {{< badge content="18" color="red" alt="postgresql-18-pg-geohash" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-geohash` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_geohash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_18-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_geohash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_17-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_geohash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_16-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_geohash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_15-1.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_geohash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_14-1.0-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_geohash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_18-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_geohash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_17-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_geohash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_16-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_geohash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_15-1.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_geohash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_14-1.0-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_geohash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_18-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_geohash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_17-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_geohash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_16-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_geohash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_15-1.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_geohash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_14-1.0-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_geohash_18" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_18-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_geohash_17" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_17-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_geohash_16" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_16-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_geohash_15" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_15-1.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_geohash_14" "1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_14-1.0-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-geohash" >}}     | {{< pkg "postgresql-17-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-geohash" "1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_18` | 1.0 | `el8.aarch64` | pigsty | 15.5 KiB | [pg_geohash_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_18` | 1.0 | `el8.x86_64` | pigsty | 15.9 KiB | [pg_geohash_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_18` | 1.0 | `el9.x86_64` | pigsty | 15.1 KiB | [pg_geohash_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_18` | 1.0 | `el9.aarch64` | pigsty | 14.7 KiB | [pg_geohash_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_18-1.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_17` | 1.0 | `el8.aarch64` | pigsty | 15.5 KiB | [pg_geohash_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_17` | 1.0 | `el8.x86_64` | pigsty | 15.9 KiB | [pg_geohash_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_17` | 1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [pg_geohash_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_17` | 1.0 | `el9.x86_64` | pigsty | 15.1 KiB | [pg_geohash_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-geohash` | 1.0 | `d12.aarch64` | pigsty | 16.8 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-geohash` | 1.0 | `d12.x86_64` | pigsty | 16.9 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-geohash` | 1.0 | `u22.aarch64` | pigsty | 16.2 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-geohash` | 1.0 | `u22.x86_64` | pigsty | 16.6 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-geohash` | 1.0 | `u24.x86_64` | pigsty | 16.3 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-geohash` | 1.0 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-17-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-17-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_16` | 1.0 | `el8.aarch64` | pigsty | 15.6 KiB | [pg_geohash_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_16` | 1.0 | `el8.x86_64` | pigsty | 15.9 KiB | [pg_geohash_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_16` | 1.0 | `el9.x86_64` | pigsty | 15.1 KiB | [pg_geohash_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_16` | 1.0 | `el9.aarch64` | pigsty | 14.8 KiB | [pg_geohash_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-geohash` | 1.0 | `d12.x86_64` | pigsty | 16.9 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-geohash` | 1.0 | `d12.aarch64` | pigsty | 16.8 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-geohash` | 1.0 | `u22.aarch64` | pigsty | 16.2 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-geohash` | 1.0 | `u22.x86_64` | pigsty | 16.6 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-geohash` | 1.0 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-geohash` | 1.0 | `u24.x86_64` | pigsty | 16.3 KiB | [postgresql-16-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-16-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_15` | 1.0 | `el8.x86_64` | pigsty | 16.0 KiB | [pg_geohash_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_15` | 1.0 | `el8.aarch64` | pigsty | 15.7 KiB | [pg_geohash_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_15` | 1.0 | `el9.x86_64` | pigsty | 16.3 KiB | [pg_geohash_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_geohash_15` | 1.0 | `el9.aarch64` | pigsty | 15.9 KiB | [pg_geohash_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-geohash` | 1.0 | `d12.x86_64` | pigsty | 16.9 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-geohash` | 1.0 | `d12.aarch64` | pigsty | 16.8 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-geohash` | 1.0 | `u22.x86_64` | pigsty | 16.6 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-geohash` | 1.0 | `u22.aarch64` | pigsty | 16.2 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-geohash` | 1.0 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-geohash` | 1.0 | `u24.x86_64` | pigsty | 16.3 KiB | [postgresql-15-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-15-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_14` | 1.0 | `el8.x86_64` | pigsty | 16.0 KiB | [pg_geohash_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_14` | 1.0 | `el8.aarch64` | pigsty | 15.7 KiB | [pg_geohash_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_14` | 1.0 | `el9.aarch64` | pigsty | 15.9 KiB | [pg_geohash_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_14` | 1.0 | `el9.x86_64` | pigsty | 16.3 KiB | [pg_geohash_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-14-pg-geohash` | 1.0 | `d12.x86_64` | pigsty | 16.9 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-geohash` | 1.0 | `d12.aarch64` | pigsty | 16.8 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-geohash` | 1.0 | `u22.x86_64` | pigsty | 16.6 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-geohash` | 1.0 | `u22.aarch64` | pigsty | 16.2 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-geohash` | 1.0 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-geohash` | 1.0 | `u24.x86_64` | pigsty | 16.3 KiB | [postgresql-14-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-14-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_geohash_13` | 1.0 | `el8.x86_64` | pigsty | 16.0 KiB | [pg_geohash_13-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_geohash_13-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_geohash_13` | 1.0 | `el8.aarch64` | pigsty | 15.7 KiB | [pg_geohash_13-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_geohash_13-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_geohash_13` | 1.0 | `el9.aarch64` | pigsty | 15.8 KiB | [pg_geohash_13-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_geohash_13-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_geohash_13` | 1.0 | `el9.x86_64` | pigsty | 16.3 KiB | [pg_geohash_13-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_geohash_13-1.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-geohash` | 1.0 | `d12.x86_64` | pigsty | 16.9 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-geohash` | 1.0 | `d12.aarch64` | pigsty | 16.8 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-geohash` | 1.0 | `u22.aarch64` | pigsty | 16.2 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-geohash` | 1.0 | `u22.x86_64` | pigsty | 16.6 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-geohash` | 1.0 | `u24.x86_64` | pigsty | 16.3 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-geohash` | 1.0 | `u24.aarch64` | pigsty | 16.1 KiB | [postgresql-13-pg-geohash_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-geohash/postgresql-13-pg-geohash_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jistok/pg_geohash" title="Repository" icon="github" subtitle="github.com/jistok/pg_geohash" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_geohash-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_geohash; # get pg_geohash source code
pig build dep pg_geohash; # install build dependencies
pig build pkg pg_geohash; # build extension rpm or deb
pig build ext pg_geohash; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_geohash; # install by extension name, for the current active PG version
pig ext install pg_geohash; # install via package alias, for the active PG version
pig ext install pg_geohash -v 18;   # install for PG 18
pig ext install pg_geohash -v 17;   # install for PG 17
pig ext install pg_geohash -v 16;   # install for PG 16
pig ext install pg_geohash -v 15;   # install for PG 15
pig ext install pg_geohash -v 14;   # install for PG 14
pig ext install pg_geohash -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_geohash;
```

