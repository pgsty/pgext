---
title: "geoip"
linkTitle: "geoip"
description: "IP-based geolocation query"
weight: 1560
categories: ["Gis"]
width: full
---

IP-based geolocation query

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1560** | {{< badge content="geoip" link="https://github.com/tvondra/geoip" >}} | {{< ext "geoip" "geoip" >}} | `0.3.0` | {{< category "GIS" >}} | {{< license "BSD 2-Clause" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "ip4r" >}} |
|   **See Also**    | {{< ext "postgis" >}} {{< ext "tzf" >}} {{< ext "country" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |

> [!Note] no pg17 on el9, no pg13 on el8


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/geoip" >}} | `0.3.0` | {{< badge content="18" color="red" alt="geoip_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `geoip_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/geoip" >}} | `0.3.0` | {{< badge content="18" color="red" alt="postgresql-18-geoip" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-geoip` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "geoip_18" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "geoip_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "geoip_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "geoip_15" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/geoip_15-0.2.4-3.rhel8.noarch.rpm" >}} | {{< pkg "geoip_14" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/geoip_14-0.2.4-3.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "geoip_18" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "geoip_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "geoip_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "geoip_15" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/geoip_15-0.2.4-3.rhel8.noarch.rpm" >}} | {{< pkg "geoip_14" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/geoip_14-0.2.4-3.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "geoip_18" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "geoip_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "geoip_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "geoip_15" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/geoip_15-0.2.4-3.rhel9.noarch.rpm" >}} | {{< pkg "geoip_14" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/geoip_14-0.2.4-3.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "geoip_18" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "geoip_17" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "geoip_16" "0.3.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "geoip_15" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/geoip_15-0.2.4-3.rhel9.noarch.rpm" >}} | {{< pkg "geoip_14" "0.2.4" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/geoip_14-0.2.4-3.rhel9.noarch.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-geoip" >}}     | {{< pkg "postgresql-17-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-geoip" "0.3.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_18` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_18` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_18` | 0.3.0 | `el9.x86_64` | pigsty | 11.7 KiB | [geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_18` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_17` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_17` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_17` | 0.3.0 | `el9.x86_64` | pigsty | 11.6 KiB | [geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_17` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-17-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_16` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_16` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_16` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_16` | 0.3.0 | `el9.x86_64` | pigsty | 11.6 KiB | [geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-16-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-16-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_15` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_15` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_15` | 0.2.4 | `el8.x86_64` | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | 0.2.4 | `el8.aarch64` | pgdg | 11.3 KiB | [geoip_15-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/geoip_15-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_15` | 0.3.0 | `el9.x86_64` | pigsty | 11.6 KiB | [geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_15` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_15` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_15` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_15-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/geoip_15-0.2.4-3.rhel9.noarch.rpm) |
| `postgresql-15-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-15-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_14` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_14` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_14` | 0.2.4 | `el8.x86_64` | pgdg | 11.3 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | 0.2.4 | `el8.aarch64` | pgdg | 11.2 KiB | [geoip_14-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/geoip_14-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_14` | 0.3.0 | `el9.x86_64` | pigsty | 11.6 KiB | [geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_14` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_14` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_14` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_14-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/geoip_14-0.2.4-3.rhel9.noarch.rpm) |
| `postgresql-14-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-14-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `geoip_13` | 0.3.0 | `el8.x86_64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/geoip_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `geoip_13` | 0.3.0 | `el8.aarch64` | pigsty | 11.7 KiB | [geoip_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/geoip_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `geoip_13` | 0.2.4 | `el8.aarch64` | pgdg | 11.3 KiB | [geoip_13-0.2.4-3.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/geoip_13-0.2.4-3.rhel8.noarch.rpm) |
| `geoip_13` | 0.3.0 | `el9.aarch64` | pigsty | 11.6 KiB | [geoip_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/geoip_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `geoip_13` | 0.3.0 | `el9.x86_64` | pigsty | 11.6 KiB | [geoip_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/geoip_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `geoip_13` | 0.2.4 | `el9.aarch64` | pgdg | 10.8 KiB | [geoip_13-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/geoip_13-0.2.4-3.rhel9.noarch.rpm) |
| `geoip_13` | 0.2.4 | `el9.x86_64` | pgdg | 11.0 KiB | [geoip_13-0.2.4-3.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/geoip_13-0.2.4-3.rhel9.noarch.rpm) |
| `postgresql-13-geoip` | 0.3.0 | `d12.x86_64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `d12.aarch64` | pigsty | 6.3 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u22.aarch64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u22.x86_64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u24.x86_64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-geoip` | 0.3.0 | `u24.aarch64` | pigsty | 6.4 KiB | [postgresql-13-geoip_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/g/geoip/postgresql-13-geoip_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/geoip" title="Repository" icon="github" subtitle="github.com/tvondra/geoip" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="geoip-0.3.0.tar.gz" >}}
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

