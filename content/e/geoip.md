---
title: "geoip"
linkTitle: "geoip"
description: "IP-based geolocation query"
weight: 1560
categories: ["GIS"]
width: full
---

[**geoip**](https://github.com/tvondra/geoip) : IP-based geolocation query


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1560** | {{< badge content="geoip" link="https://github.com/tvondra/geoip" >}} | {{< ext "geoip" >}} | `0.3.0` | {{< category "GIS" >}} | {{< license "BSD 2-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `geoip` |
|   **Requires**    | {{< ext "ip4r" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "tzf" >}} {{< ext "country" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |

> [!Note] no pg17 on el9, no legacy branch on el8


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `geoip` | `ip4r` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "geoip_18" "green" >}} {{< bg "17" "geoip_17" "green" >}} {{< bg "16" "geoip_16" "green" >}} {{< bg "15" "geoip_15" "green" >}} {{< bg "14" "geoip_14" "green" >}} | `geoip_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-geoip" "green" >}} {{< bg "17" "postgresql-17-geoip" "green" >}} {{< bg "16" "postgresql-16-geoip" "green" >}} {{< bg "15" "postgresql-15-geoip" "green" >}} {{< bg "14" "postgresql-14-geoip" "green" >}} | `postgresql-$v-geoip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-geoip : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-geoip : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-geoip : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 KiB | [geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-geoip` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-geoip` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 KiB | [geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-geoip` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-geoip` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 KiB | [geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-geoip` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-geoip` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_15` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_15` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_15` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.0 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 KiB | [geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_15` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 10.8 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-geoip` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-geoip` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_14` | `0.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 11.3 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_14` | `0.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 11.2 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_14` | `0.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 11.0 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.6 KiB | [geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_14` | `0.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 10.8 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-geoip` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-geoip` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/geoip" title="Repository" icon="github" subtitle="github.com/tvondra/geoip" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="geoip-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg geoip;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install geoip;		# install via package name, for the active PG version

pig install geoip -v 18;   # install for PG 18
pig install geoip -v 17;   # install for PG 17
pig install geoip -v 16;   # install for PG 16
pig install geoip -v 15;   # install for PG 15
pig install geoip -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION geoip CASCADE; -- requires ip4r
```



## Usage

> [geoip: IP-based geolocation for PostgreSQL](https://github.com/tvondra/geoip)

This extension provides IP-based geolocation — you provide an IPv4 or IPv6 address and the extension looks up country, city, GPS coordinates, ASN and more. It requires the `ip4r` extension and GeoLite2 data from [MaxMind](https://www.maxmind.com).

```sql
CREATE EXTENSION ip4r;
CREATE EXTENSION geoip;
```

### Functions

| Function | Description |
|----------|-------------|
| `geoip_country_code(ip4\|ip6)` | Returns country code (2 chars) |
| `geoip_country(ip4\|ip6)` | Returns all country info (code, name, network) |
| `geoip_city_location(ip4\|ip6)` | Returns just the location ID (INT) |
| `geoip_city(ip4\|ip6)` | Returns all city info (GPS, ZIP code, etc.) |
| `geoip_asn(ip4\|ip6)` | Returns ASN name and IP range |

### Examples

```sql
SELECT geoip_country_code('78.45.133.255'::ip4);
-- CZ

SELECT * FROM geoip.geoip_city('78.45.133.255'::ip4);
--  geoname_id | country_iso_code | city_name | postal_code | ...
-- ------------+------------------+-----------+-------------+----
--     3066399 | CZ               | Sardice   | 696 13      | ...

SELECT * FROM geoip.geoip_country('78.45.133.255'::ip4);
--     network     | country_iso_code | country_name
-- ----------------+------------------+--------------
--  78.45.128.0/17 | CZ               | Czechia

SELECT * FROM geoip.geoip_asn('78.45.133.255'::ip4);
--    network    | asn_number |      asn_name
-- --------------+------------+---------------------
--  78.44.0.0/15 |       6830 | Liberty Global B.V.
```


## Loading Data

The extension requires GeoLite2 CSV data from MaxMind. Download the City, Country, and ASN datasets in CSV format from [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geoip2/geolite2/), then load:

```bash
cat GeoLite2-Country-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-Country-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_country_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Locations-en.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_locations FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-City-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_city_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv4.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'

cat GeoLite2-ASN-Blocks-IPv6.csv | \
  psql $DBNAME -c 'COPY geoip.geoip_asn_blocks FROM stdin WITH (FORMAT CSV, HEADER)'
```

The "locations" files have multiple language variants — pick the one that works for you.
