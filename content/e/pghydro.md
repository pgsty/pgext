---
title: "pghydro"
linkTitle: "pghydro"
description: "Drainage network analysis core for PostgreSQL and PostGIS"
weight: 1600
categories: ["GIS"]
width: full
---

[**pghydro**](https://github.com/pghydro/pghydro) : Drainage network analysis core for PostgreSQL and PostGIS


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1600** | {{< badge content="pghydro" link="https://github.com/pghydro/pghydro" >}} | {{< ext "pghydro" >}} | `6.6` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pghydro` |
|   **Requires**    | {{< ext "plpgsql" >}} {{< ext "postgis" >}} |
|    **Siblings**   | {{< ext "pgh_raster" >}} {{< ext "pgh_hgm" >}} {{< ext "pgh_output" >}} {{< ext "pgh_output_en_au" >}} {{< ext "pgh_output_pt_br" >}} {{< ext "pgh_consistency" >}} |

> [!Note] Lead row; package also ships pgh_raster, pgh_hgm, pgh_output, pgh_output_en_au, pgh_output_pt_br, and pgh_consistency.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pghydro` | `plpgsql`, `postgis` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "pghydro_18" "green" >}} {{< bg "17" "pghydro_17" "green" >}} {{< bg "16" "pghydro_16" "green" >}} {{< bg "15" "pghydro_15" "green" >}} {{< bg "14" "pghydro_14" "green" >}} | `pghydro_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `6.6` | {{< bg "18" "postgresql-18-pghydro" "green" >}} {{< bg "17" "postgresql-17-pghydro" "green" >}} {{< bg "16" "postgresql-16-pghydro" "green" >}} {{< bg "15" "postgresql-15-pghydro" "green" >}} {{< bg "14" "postgresql-14-pghydro" "green" >}} | `postgresql-$v-pghydro` | `postgresql-$v-postgis-3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 6.6" "pghydro_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "pghydro_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 6.6" "postgresql-18-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-17-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-16-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-15-pghydro : AVAIL 1" "green" >}} | {{< bg "PIGSTY 6.6" "postgresql-14-pghydro : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pghydro_18` | `6.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.0 KiB | [pghydro_18-6.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pghydro_18-6.6-1PIGSTY.el8.x86_64.rpm) |
| `pghydro_18` | `6.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.9 KiB | [pghydro_18-6.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pghydro_18-6.6-1PIGSTY.el8.aarch64.rpm) |
| `pghydro_18` | `6.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.2 KiB | [pghydro_18-6.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pghydro_18-6.6-1PIGSTY.el9.x86_64.rpm) |
| `pghydro_18` | `6.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.1 KiB | [pghydro_18-6.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pghydro_18-6.6-1PIGSTY.el9.aarch64.rpm) |
| `pghydro_18` | `6.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.4 KiB | [pghydro_18-6.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pghydro_18-6.6-1PIGSTY.el10.x86_64.rpm) |
| `pghydro_18` | `6.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [pghydro_18-6.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pghydro_18-6.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pghydro` | `6.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 135.5 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pghydro` | `6.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.5 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pghydro` | `6.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.5 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pghydro` | `6.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 135.5 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pghydro` | `6.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 135.8 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pghydro` | `6.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.8 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pghydro` | `6.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 135.7 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pghydro` | `6.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.7 KiB | [postgresql-18-pghydro_6.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-18-pghydro_6.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pghydro_17` | `6.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.0 KiB | [pghydro_17-6.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pghydro_17-6.6-1PIGSTY.el8.x86_64.rpm) |
| `pghydro_17` | `6.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.9 KiB | [pghydro_17-6.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pghydro_17-6.6-1PIGSTY.el8.aarch64.rpm) |
| `pghydro_17` | `6.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.2 KiB | [pghydro_17-6.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pghydro_17-6.6-1PIGSTY.el9.x86_64.rpm) |
| `pghydro_17` | `6.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.1 KiB | [pghydro_17-6.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pghydro_17-6.6-1PIGSTY.el9.aarch64.rpm) |
| `pghydro_17` | `6.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.4 KiB | [pghydro_17-6.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pghydro_17-6.6-1PIGSTY.el10.x86_64.rpm) |
| `pghydro_17` | `6.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [pghydro_17-6.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pghydro_17-6.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pghydro` | `6.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 135.5 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pghydro` | `6.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.5 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pghydro` | `6.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.5 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pghydro` | `6.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 135.5 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pghydro` | `6.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 135.7 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pghydro` | `6.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.7 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pghydro` | `6.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 135.7 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pghydro` | `6.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.7 KiB | [postgresql-17-pghydro_6.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-17-pghydro_6.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pghydro_16` | `6.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.0 KiB | [pghydro_16-6.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pghydro_16-6.6-1PIGSTY.el8.x86_64.rpm) |
| `pghydro_16` | `6.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.9 KiB | [pghydro_16-6.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pghydro_16-6.6-1PIGSTY.el8.aarch64.rpm) |
| `pghydro_16` | `6.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.2 KiB | [pghydro_16-6.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pghydro_16-6.6-1PIGSTY.el9.x86_64.rpm) |
| `pghydro_16` | `6.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.1 KiB | [pghydro_16-6.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pghydro_16-6.6-1PIGSTY.el9.aarch64.rpm) |
| `pghydro_16` | `6.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.4 KiB | [pghydro_16-6.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pghydro_16-6.6-1PIGSTY.el10.x86_64.rpm) |
| `pghydro_16` | `6.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [pghydro_16-6.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pghydro_16-6.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pghydro` | `6.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 135.5 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pghydro` | `6.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.5 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pghydro` | `6.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.5 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pghydro` | `6.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 135.5 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pghydro` | `6.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 135.7 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pghydro` | `6.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.7 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pghydro` | `6.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 135.8 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pghydro` | `6.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.8 KiB | [postgresql-16-pghydro_6.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-16-pghydro_6.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pghydro_15` | `6.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.0 KiB | [pghydro_15-6.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pghydro_15-6.6-1PIGSTY.el8.x86_64.rpm) |
| `pghydro_15` | `6.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.9 KiB | [pghydro_15-6.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pghydro_15-6.6-1PIGSTY.el8.aarch64.rpm) |
| `pghydro_15` | `6.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.2 KiB | [pghydro_15-6.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pghydro_15-6.6-1PIGSTY.el9.x86_64.rpm) |
| `pghydro_15` | `6.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.1 KiB | [pghydro_15-6.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pghydro_15-6.6-1PIGSTY.el9.aarch64.rpm) |
| `pghydro_15` | `6.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.4 KiB | [pghydro_15-6.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pghydro_15-6.6-1PIGSTY.el10.x86_64.rpm) |
| `pghydro_15` | `6.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [pghydro_15-6.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pghydro_15-6.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pghydro` | `6.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 135.5 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pghydro` | `6.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.5 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pghydro` | `6.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.5 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pghydro` | `6.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 135.5 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pghydro` | `6.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 135.7 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pghydro` | `6.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.7 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pghydro` | `6.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 135.7 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pghydro` | `6.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.7 KiB | [postgresql-15-pghydro_6.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-15-pghydro_6.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pghydro_14` | `6.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.0 KiB | [pghydro_14-6.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pghydro_14-6.6-1PIGSTY.el8.x86_64.rpm) |
| `pghydro_14` | `6.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 144.9 KiB | [pghydro_14-6.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pghydro_14-6.6-1PIGSTY.el8.aarch64.rpm) |
| `pghydro_14` | `6.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.2 KiB | [pghydro_14-6.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pghydro_14-6.6-1PIGSTY.el9.x86_64.rpm) |
| `pghydro_14` | `6.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 138.1 KiB | [pghydro_14-6.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pghydro_14-6.6-1PIGSTY.el9.aarch64.rpm) |
| `pghydro_14` | `6.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 138.4 KiB | [pghydro_14-6.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pghydro_14-6.6-1PIGSTY.el10.x86_64.rpm) |
| `pghydro_14` | `6.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [pghydro_14-6.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pghydro_14-6.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pghydro` | `6.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 135.5 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pghydro` | `6.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 135.5 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pghydro` | `6.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 135.5 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pghydro` | `6.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 135.5 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pghydro` | `6.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 135.7 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pghydro` | `6.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 135.7 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pghydro` | `6.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 135.7 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pghydro` | `6.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 135.7 KiB | [postgresql-14-pghydro_6.6-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pghydro/postgresql-14-pghydro_6.6-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pghydro/pghydro" title="Repository" icon="github" subtitle="github.com/pghydro/pghydro" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pghydro-6.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pghydro;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pghydro;		# install via package name, for the active PG version

pig install pghydro -v 18;   # install for PG 18
pig install pghydro -v 17;   # install for PG 17
pig install pghydro -v 16;   # install for PG 16
pig install pghydro -v 15;   # install for PG 15
pig install pghydro -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pghydro CASCADE; -- requires plpgsql, postgis
```

## Usage

> Sources: [GitHub repo](https://github.com/pghydro/pghydro), [README](https://raw.githubusercontent.com/pghydro/pghydro/master/README.md), [releases](https://github.com/pghydro/pghydro/releases)
> Lead extension for the PgHydro suite.

PgHydro extends PostGIS and PostgreSQL for drainage network analysis and water-resources decision making. The project covers river network modeling, flow-direction analysis, Otto Pfafstetter basin coding, upstream and downstream stretch selection, distance-to-mouth calculations, upstream area analysis, river orders, and basin levels.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION postgis_raster;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_raster;
CREATE EXTENSION pgh_hgm;
CREATE EXTENSION pgh_consistency;
CREATE EXTENSION pgh_output;
```

### Components

- `pghydro` is the core drainage-network analysis extension.
- `pgh_raster` uses hydrological products derived from a digital elevation model.
- `pgh_hgm` combines `pghydro` and `pgh_raster` for hydrogeomorphological analysis.
- `pgh_output` provides reporting objects.
- `pgh_consistency` adds Pfafstetter consistency checks.

### Requirements

- PostgreSQL 9.1 or newer.
- PostGIS 3.x.
- PostGIS Raster.

### Notes

- The README says the master branch tracks the latest minor release, 6.6.
- The CSV lead row is the core `pghydro` package, but the repository also ships companion extensions in the same release tree.
