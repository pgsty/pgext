---
title: "pg_polyline"
linkTitle: "pg_polyline"
description: "Fast Google Encoded Polyline encoding & decoding for postgres"
weight: 1570
categories: ["Gis"]
width: full
---

Fast Google Encoded Polyline encoding & decoding for postgres

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1570** | {{< badge content="pg_polyline" link="https://github.com/yihong0618/pg_polyline" >}} | {{< ext "pg_polyline" "pg_polyline" >}} | `0.0.1` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |

> [!Note] pgrx=0.12.7


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_polyline" >}} | `0.0.1` | {{< badge content="18" color="red" alt="pg_polyline_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_polyline_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_polyline" >}} | `0.0.1` | {{< badge content="18" color="red" alt="postgresql-18-pg-polyline" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pg-polyline` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pg_polyline_18" >}}     | {{< pkg "pg_polyline_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_17-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_polyline_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_16-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_polyline_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_15-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_polyline_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_14-0.0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pg_polyline_18" >}}     | {{< pkg "pg_polyline_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_17-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_polyline_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_16-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_polyline_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_15-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_polyline_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_14-0.0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pg_polyline_18" >}}     | {{< pkg "pg_polyline_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_17-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_polyline_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_16-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_polyline_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_15-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_polyline_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_14-0.0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pg_polyline_18" >}}     | {{< pkg "pg_polyline_17" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_17-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_polyline_16" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_16-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_polyline_15" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_15-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_polyline_14" "0.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_14-0.0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-polyline" >}}     | {{< pkg "postgresql-17-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pg-polyline" "0.0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_polyline_17` | 0.0.1 | `el8.x86_64` | pigsty | 220.7 KiB | [pg_polyline_17-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_17-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el8.aarch64` | pigsty | 204.9 KiB | [pg_polyline_17-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_17-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el9.aarch64` | pigsty | 218.9 KiB | [pg_polyline_17-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_17-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_17` | 0.0.1 | `el9.x86_64` | pigsty | 225.0 KiB | [pg_polyline_17-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_17-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.5 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-17-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_polyline_16` | 0.0.1 | `el8.x86_64` | pigsty | 220.7 KiB | [pg_polyline_16-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_16-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el8.aarch64` | pigsty | 204.9 KiB | [pg_polyline_16-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_16-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el9.x86_64` | pigsty | 225.2 KiB | [pg_polyline_16-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_16-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_16` | 0.0.1 | `el9.aarch64` | pigsty | 219.0 KiB | [pg_polyline_16-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_16-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.5 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.6 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-16-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_polyline_15` | 0.0.1 | `el8.x86_64` | pigsty | 220.7 KiB | [pg_polyline_15-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_15-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el8.aarch64` | pigsty | 204.8 KiB | [pg_polyline_15-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_15-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el9.x86_64` | pigsty | 225.1 KiB | [pg_polyline_15-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_15-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_15` | 0.0.1 | `el9.aarch64` | pigsty | 218.9 KiB | [pg_polyline_15-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_15-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.7 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.3 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-15-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_polyline_14` | 0.0.1 | `el8.x86_64` | pigsty | 220.7 KiB | [pg_polyline_14-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_14-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el8.aarch64` | pigsty | 205.0 KiB | [pg_polyline_14-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_14-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el9.x86_64` | pigsty | 225.1 KiB | [pg_polyline_14-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_14-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_polyline_14` | 0.0.1 | `el9.aarch64` | pigsty | 218.7 KiB | [pg_polyline_14-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_14-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.4 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.2 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.4 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.5 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-14-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_polyline_13` | 0.0.1 | `el8.aarch64` | pigsty | 204.9 KiB | [pg_polyline_13-0.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_polyline_13-0.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el8.x86_64` | pigsty | 220.7 KiB | [pg_polyline_13-0.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_polyline_13-0.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el9.aarch64` | pigsty | 219.0 KiB | [pg_polyline_13-0.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_polyline_13-0.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_polyline_13` | 0.0.1 | `el9.x86_64` | pigsty | 225.1 KiB | [pg_polyline_13-0.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_polyline_13-0.0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pg-polyline` | 0.0.1 | `d12.aarch64` | pigsty | 158.3 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `d12.x86_64` | pigsty | 178.6 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u22.aarch64` | pigsty | 183.6 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u22.x86_64` | pigsty | 194.4 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u24.aarch64` | pigsty | 182.2 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pg-polyline` | 0.0.1 | `u24.x86_64` | pigsty | 193.7 KiB | [postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-polyline/postgresql-13-pg-polyline_0.0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/yihong0618/pg_polyline" title="Repository" icon="github" subtitle="github.com/yihong0618/pg_polyline" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_polyline-0.0.1.tar.gz" >}}
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

