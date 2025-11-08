---
title: "geoip"
linkTitle: "geoip"
description: "IP-based geolocation query"
weight: 1560
categories: ["GIS"]
width: full
---

IP-based geolocation query


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1560** | {{< badge content="geoip" link="https://github.com/tvondra/geoip" >}} | {{< ext "geoip" >}} | `0.3.0` | {{< category "GIS" >}} | {{< license "BSD 2-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "ip4r" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "tzf" >}} {{< ext "country" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |

> [!Note] no pg17 on el9, no pg13 on el8


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/geoip" >}} | `0.3.0` | {{< bg "18" "geoip_18" "green" >}} {{< bg "17" "geoip_17" "green" >}} {{< bg "16" "geoip_16" "green" >}} {{< bg "15" "geoip_15" "green" >}} {{< bg "14" "geoip_14" "green" >}} {{< bg "13" "geoip_13" "green" >}} | `geoip_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/geoip" >}} | `0.3.0` | {{< bg "18" "postgresql-18-geoip" "green" >}} {{< bg "17" "postgresql-17-geoip" "green" >}} {{< bg "16" "postgresql-16-geoip" "green" >}} {{< bg "15" "postgresql-15-geoip" "green" >}} {{< bg "14" "postgresql-14-geoip" "green" >}} {{< bg "13" "postgresql-13-geoip" "green" >}} | `postgresql-$v-geoip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 2" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 2" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 2" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "geoip_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "geoip_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-geoip : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-geoip : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_18` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_18` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_18` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_18` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_18` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_18` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-18-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-18-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_17` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_17` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_17` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_17` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_17` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_17` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_16` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_16` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_16` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_16` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_16` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_16` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_15` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_15` | 0.2.4 | `el8.x86_64` | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_15` | 0.2.4 | `el8.aarch64` | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_15` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_15` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_15` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_15` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_15` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_14` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_14` | 0.2.4 | `el8.x86_64` | pgdg | 11.3 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_14` | 0.2.4 | `el8.aarch64` | pgdg | 11.2 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_14` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_14` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_14` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_14` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_14` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `geoip_13` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_13` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_13` | 0.2.4 | `el8.aarch64` | pgdg | 11.3 KiB | [geoip_13-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/geoip_13-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_13` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_13` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_13-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/geoip_13-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_13` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_13` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_13-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/geoip_13-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_13` | 0.3.0 | `el10.x86_64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/geoip_13-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `geoip_13` | 0.3.0 | `el10.aarch64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/geoip_13-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `d13.x86_64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `d13.aarch64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/geoip" title="Repository" icon="github" subtitle="github.com/tvondra/geoip" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="geoip-0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get geoip; # get geoip source code
pig build dep geoip; # install build dependencies
pig build pkg geoip; # build extension rpm or deb
pig build ext geoip; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install geoip; # install by extension name, for the current active PG version
pig ext install geoip; # install via package alias, for the active PG version
pig ext install geoip -v 18;   # install for PG 18
pig ext install geoip -v 17;   # install for PG 17
pig ext install geoip -v 16;   # install for PG 16
pig ext install geoip -v 15;   # install for PG 15
pig ext install geoip -v 14;   # install for PG 14
pig ext install geoip -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION geoip CASCADE SCHEMA geoip;
```

