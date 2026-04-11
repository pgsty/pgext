---
title: "rdf_fdw"
linkTitle: "rdf_fdw"
description: "Foreign data wrapper for RDF triplestores over SPARQL endpoints"
weight: 8760
categories: ["FDW"]
width: full
---

[**rdf_fdw**](https://github.com/jimjonesbr/rdf_fdw) : Foreign data wrapper for RDF triplestores over SPARQL endpoints


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8760** | {{< badge content="rdf_fdw" link="https://github.com/jimjonesbr/rdf_fdw" >}} | {{< ext "rdf_fdw" >}} | `2.4.0` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "postgres_fdw" >}} {{< ext "sparql" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `rdf_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.0` | {{< bg "18" "rdf_fdw_18" "green" >}} {{< bg "17" "rdf_fdw_17" "green" >}} {{< bg "16" "rdf_fdw_16" "green" >}} {{< bg "15" "rdf_fdw_15" "green" >}} {{< bg "14" "rdf_fdw_14" "green" >}} | `rdf_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.0` | {{< bg "18" "postgresql-18-rdf-fdw" "green" >}} {{< bg "17" "postgresql-17-rdf-fdw" "green" >}} {{< bg "16" "postgresql-16-rdf-fdw" "green" >}} {{< bg "15" "postgresql-15-rdf-fdw" "green" >}} {{< bg "14" "postgresql-14-rdf-fdw" "green" >}} | `postgresql-$v-rdf-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "rdf_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "rdf_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdf_fdw_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_18` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.3 KiB | [rdf_fdw_18-2.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_18-2.4.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_18` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 134.5 KiB | [rdf_fdw_18-2.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_18-2.4.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_18` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 139.3 KiB | [rdf_fdw_18-2.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_18-2.4.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.8 KiB | [rdf_fdw_18-2.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_18-2.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.7 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 320.5 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 328.8 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 321.2 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 352.9 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 347.7 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 338.9 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 334.8 KiB | [postgresql-18-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.3 KiB | [rdf_fdw_17-2.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_17-2.4.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 134.8 KiB | [rdf_fdw_17-2.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_17-2.4.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_17` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 139.2 KiB | [rdf_fdw_17-2.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_17-2.4.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.6 KiB | [rdf_fdw_17-2.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_17-2.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.4 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 320.6 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 329.9 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 320.5 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 374.4 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 369.6 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 339.2 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 334.6 KiB | [postgresql-17-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 138.4 KiB | [rdf_fdw_16-2.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_16-2.4.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 134.8 KiB | [rdf_fdw_16-2.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_16-2.4.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_16` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 139.0 KiB | [rdf_fdw_16-2.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_16-2.4.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.8 KiB | [rdf_fdw_16-2.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_16-2.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.2 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 321.3 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 329.4 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 320.5 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 372.3 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 367.5 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 339.1 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 334.6 KiB | [postgresql-16-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_15` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 140.9 KiB | [rdf_fdw_15-2.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_15-2.4.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_15` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 136.7 KiB | [rdf_fdw_15-2.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_15-2.4.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_15` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.3 KiB | [rdf_fdw_15-2.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_15-2.4.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_15` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.3 KiB | [rdf_fdw_15-2.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_15-2.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 330.5 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.6 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 330.7 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 322.0 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 373.6 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 369.1 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 340.0 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 336.4 KiB | [postgresql-15-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_14` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 140.6 KiB | [rdf_fdw_14-2.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_14-2.4.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_14` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 136.7 KiB | [rdf_fdw_14-2.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_14-2.4.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_14` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 141.4 KiB | [rdf_fdw_14-2.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_14-2.4.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_14` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 138.4 KiB | [rdf_fdw_14-2.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_14-2.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 330.3 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 321.2 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.0 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 322.2 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 373.4 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 368.9 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 339.9 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 336.4 KiB | [postgresql-14-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/rdf_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/rdf_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="rdf_fdw-2.4.0.zip" >}}
{{< /cards >}}


```bash
pig build pkg rdf_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install rdf_fdw;		# install via package name, for the active PG version

pig install rdf_fdw -v 18;   # install for PG 18
pig install rdf_fdw -v 17;   # install for PG 17
pig install rdf_fdw -v 16;   # install for PG 16
pig install rdf_fdw -v 15;   # install for PG 15
pig install rdf_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION rdf_fdw;
```


## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION rdf_fdw;
> CREATE SERVER dbpedia FOREIGN DATA WRAPPER rdf_fdw
>   OPTIONS (endpoint 'https://dbpedia.org/sparql');
> ```
>
> Source: [README](https://github.com/jimjonesbr/rdf_fdw)

`rdf_fdw` is a foreign data wrapper for RDF triplestores exposed through SPARQL endpoints. It lets PostgreSQL query RDF data with SQL, supports SQL clause pushdown, adds an `rdfnode` type for RDF terms, and includes SPARQL 1.1 function support.

## Server Setup

Register a SPARQL endpoint with `CREATE SERVER`:

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (endpoint 'https://dbpedia.org/sparql');
```

The README documents server options such as:

- `endpoint` (required)
- `batch_size`
- `enable_pushdown`
- `format`
- `http_proxy`
- `connect_timeout`

Proxy credentials belong in a user mapping.

## Foreign Tables

`rdf_fdw` works by declaring foreign tables that embed SPARQL queries and map result variables to PostgreSQL columns. The README also highlights native RDF node handling through the custom `rdfnode` type.

## Pushdown and DML

The upstream docs specifically call out pushdown for:

- `WHERE`
- `LIMIT`
- `ORDER BY`
- `DISTINCT`

They also document data modification support:

- `INSERT`
- `UPDATE`
- `DELETE`

Batching for SPARQL UPDATE traffic is controlled with the `batch_size` option.

## Helper Functions

The README lists utility functions including:

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

It also documents broader SPARQL function coverage, including aggregates, string functions, numeric functions, date/time functions, hash functions, and custom functions.

## Notes

The current README warns that retrieved RDF data is loaded into memory before conversion for PostgreSQL, so large result sets require adequate PostgreSQL memory. It also documents PostgreSQL 9.5+ as the supported baseline.
