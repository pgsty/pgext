---
title: "tzf"
linkTitle: "tzf"
description: "Fast lookup timezone name by GPS coordinates"
weight: 1680
categories: ["GIS"]
width: full
---

[**pg_tzf**](https://github.com/ringsaturn/pg-tzf) : Fast lookup timezone name by GPS coordinates


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1680** | {{< badge content="tzf" link="https://github.com/ringsaturn/pg-tzf" >}} | {{< ext "tzf" "pg_tzf" >}} | `0.2.3` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "geoip" >}} {{< ext "pg_cron" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_tzf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "pg_tzf_18" "green" >}} {{< bg "17" "pg_tzf_17" "green" >}} {{< bg "16" "pg_tzf_16" "green" >}} {{< bg "15" "pg_tzf_15" "green" >}} {{< bg "14" "pg_tzf_14" "green" >}} {{< bg "13" "pg_tzf_13" "red" >}} | `pg_tzf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.3` | {{< bg "18" "postgresql-18-tzf" "green" >}} {{< bg "17" "postgresql-17-tzf" "green" >}} {{< bg "16" "postgresql-16-tzf" "green" >}} {{< bg "15" "postgresql-15-tzf" "green" >}} {{< bg "14" "postgresql-14-tzf" "green" >}} {{< bg "13" "postgresql-13-tzf" "red" >}} | `postgresql-$v-tzf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "pg_tzf_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_tzf_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-tzf : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-tzf : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-18-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-17-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-16-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-15-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.3" "postgresql-14-tzf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.2" "postgresql-13-tzf : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tzf_18` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tzf_18-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_tzf_18` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tzf_18-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_tzf_18` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.7 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tzf_18-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_tzf_18` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tzf_18-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_tzf_18` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.7 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tzf_18-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_tzf_18` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pg_tzf_18-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tzf_18-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-tzf` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.9 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.8 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.8 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.8 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-tzf` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.8 KiB | [postgresql-18-tzf_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-18-tzf_0.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tzf_17` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tzf_17-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_tzf_17` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tzf_17-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_tzf_17` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.6 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tzf_17-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_tzf_17` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tzf_17-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_tzf_17` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.7 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tzf_17-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_tzf_17` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pg_tzf_17-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tzf_17-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-tzf` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.2 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.2 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.7 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-tzf` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 MiB | [postgresql-17-tzf_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-17-tzf_0.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tzf_16` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tzf_16-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_tzf_16` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tzf_16-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_tzf_16` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.6 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tzf_16-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_tzf_16` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tzf_16-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_tzf_16` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.7 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tzf_16-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_tzf_16` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pg_tzf_16-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tzf_16-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-tzf` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.2 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.2 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.8 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.7 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-tzf` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 MiB | [postgresql-16-tzf_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-16-tzf_0.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tzf_15` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tzf_15-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_tzf_15` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tzf_15-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_tzf_15` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.7 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tzf_15-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_tzf_15` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tzf_15-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_tzf_15` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.7 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tzf_15-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_tzf_15` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pg_tzf_15-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tzf_15-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-tzf` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.2 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.2 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.7 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-tzf` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.6 MiB | [postgresql-15-tzf_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-15-tzf_0.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tzf_14` | `0.2.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tzf_14-0.2.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_tzf_14` | `0.2.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tzf_14-0.2.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_tzf_14` | `0.2.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.7 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tzf_14-0.2.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_tzf_14` | `0.2.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tzf_14-0.2.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_tzf_14` | `0.2.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.7 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tzf_14-0.2.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_tzf_14` | `0.2.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pg_tzf_14-0.2.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tzf_14-0.2.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-tzf` | `0.2.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.2 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.2 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.1 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.8 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.7 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-tzf` | `0.2.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.6 MiB | [postgresql-14-tzf_0.2.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-14-tzf_0.2.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-tzf` | `0.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.2 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-tzf` | `0.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.1 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-tzf` | `0.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.7 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-tzf` | `0.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.6 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-tzf` | `0.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.6 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-tzf` | `0.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 MiB | [postgresql-13-tzf_0.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/tzf/postgresql-13-tzf_0.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ringsaturn/pg-tzf" title="Repository" icon="github" subtitle="github.com/ringsaturn/pg-tzf" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg-tzf-0.2.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tzf;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tzf;		# install via package name, for the active PG version
pig install tzf;		# install by extension name, for the current active PG version

pig install tzf -v 18;   # install for PG 18
pig install tzf -v 17;   # install for PG 17
pig install tzf -v 16;   # install for PG 16
pig install tzf -v 15;   # install for PG 15
pig install tzf -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION tzf;
```
