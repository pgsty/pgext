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
| **8760** | {{< badge content="rdf_fdw" link="https://github.com/jimjonesbr/rdf_fdw" >}} | {{< ext "rdf_fdw" >}} | `2.6.0` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "postgres_fdw" >}} {{< ext "sparql" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `rdf_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6.0` | {{< bg "18" "rdf_fdw_18" "green" >}} {{< bg "17" "rdf_fdw_17" "green" >}} {{< bg "16" "rdf_fdw_16" "green" >}} {{< bg "15" "rdf_fdw_15" "green" >}} {{< bg "14" "rdf_fdw_14" "green" >}} | `rdf_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.6.0` | {{< bg "18" "postgresql-18-rdf-fdw" "green" >}} {{< bg "17" "postgresql-17-rdf-fdw" "green" >}} {{< bg "16" "postgresql-16-rdf-fdw" "green" >}} {{< bg "15" "postgresql-15-rdf-fdw" "green" >}} {{< bg "14" "postgresql-14-rdf-fdw" "green" >}} | `postgresql-$v-rdf-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.6.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_18` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 150.9 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_18-2.6.0-1PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_18` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 141.7 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_18-2.6.0-1PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_18` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.3 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_18-2.6.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_18` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 140.3 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_18-2.6.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_18` | `2.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 146.0 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_18-2.6.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_18` | `2.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 142.1 KiB | [rdf_fdw_18-2.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_18-2.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 350.0 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 339.2 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 348.8 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 339.5 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 367.8 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 362.2 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 354.2 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 350.0 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 352.5 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.6.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 347.9 KiB | [postgresql-18-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_17` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 151.1 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_17-2.6.0-1PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_17` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 141.8 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_17-2.6.0-1PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_17` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 145.3 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_17-2.6.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_17` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 140.2 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_17-2.6.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_17` | `2.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 146.0 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_17-2.6.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_17` | `2.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 142.0 KiB | [rdf_fdw_17-2.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_17-2.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 348.3 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 339.5 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 348.2 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 339.4 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 388.9 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 383.9 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 353.6 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 349.5 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 351.3 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.6.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 346.8 KiB | [postgresql-17-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_16` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 151.2 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_16-2.6.0-1PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_16` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 141.7 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_16-2.6.0-1PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_16` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 144.9 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_16-2.6.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_16` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 140.6 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_16-2.6.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_16` | `2.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 146.0 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_16-2.6.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_16` | `2.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 142.0 KiB | [rdf_fdw_16-2.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_16-2.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 348.1 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 339.0 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 348.0 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 340.3 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 387.0 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 381.6 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 353.3 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 349.4 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 351.6 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.6.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 346.8 KiB | [postgresql-16-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_15` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 152.2 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_15-2.6.0-1PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_15` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 143.1 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_15-2.6.0-1PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_15` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 147.3 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_15-2.6.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_15` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 142.9 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_15-2.6.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_15` | `2.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 148.0 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_15-2.6.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_15` | `2.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 144.3 KiB | [rdf_fdw_15-2.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_15-2.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 350.4 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 340.2 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 350.7 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 341.1 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 388.2 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 383.1 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 353.7 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 350.2 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 352.7 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.6.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 347.6 KiB | [postgresql-15-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_14` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 152.3 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_14-2.6.0-1PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_14` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 143.0 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_14-2.6.0-1PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_14` | `2.6.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 147.2 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_14-2.6.0-1PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_14` | `2.6.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 142.8 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_14-2.6.0-1PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_14` | `2.6.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 147.9 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_14-2.6.0-1PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_14` | `2.6.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 144.2 KiB | [rdf_fdw_14-2.6.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_14-2.6.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 350.3 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 339.9 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 349.7 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 341.0 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 387.9 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 382.8 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 353.9 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 350.1 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 352.4 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.6.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 347.5 KiB | [postgresql-14-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.6.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/rdf_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/rdf_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="rdf_fdw-2.6.0.tar.gz" >}}
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

Sources:

- [PGXN rdf_fdw 2.6.0](https://pgxn.org/dist/rdf_fdw/2.6.0/)
- [rdf_fdw README](https://github.com/jimjonesbr/rdf_fdw)
- [rdf_fdw CHANGELOG](https://github.com/jimjonesbr/rdf_fdw/blob/master/CHANGELOG.md)
- [rdf_fdw control file](https://pgxn.org/dist/rdf_fdw/2.6.0/)
- [Local package metadata](../db/extension.csv)

`rdf_fdw` is a PostgreSQL foreign data wrapper for querying RDF triplestores over SPARQL endpoints. It exposes SPARQL result variables as foreign-table columns, supports pushdown for common SQL clauses, includes a native `rdfnode` type for RDF terms, provides SPARQL 1.1 helper functions, and can perform SPARQL `INSERT`, `UPDATE`, and `DELETE` through writable foreign tables.

v2.6.0 adds Bearer-token authentication through `USER MAPPING`, a `max_response_size` server option to cap HTTP response bodies, BCE date/timestamp cast handling, and many `rdfnode` parser/comparison fixes. v2.5 added `request_timeout` and `readonly` options.

### Create the Extension

```sql
CREATE EXTENSION IF NOT EXISTS rdf_fdw;

SELECT rdf_fdw_version();
SELECT * FROM rdf_fdw_settings();
```

To install or update to the exact SQL version:

```sql
CREATE EXTENSION rdf_fdw WITH VERSION '2.6';
ALTER EXTENSION rdf_fdw UPDATE TO '2.6';
```

### Register a SPARQL Endpoint

```sql
CREATE SERVER dbpedia
FOREIGN DATA WRAPPER rdf_fdw
OPTIONS (
  endpoint          'https://dbpedia.org/sparql',
  enable_pushdown   'true',
  request_timeout   '60',
  max_response_size '104857600',
  readonly          'true'
);
```

Useful server options include:

- `endpoint`: SPARQL endpoint URL; required.
- `batch_size`: number of rows per SPARQL UPDATE batch.
- `enable_pushdown`: enables SQL-to-SPARQL pushdown.
- `format`: expected SPARQL result MIME type.
- `http_proxy`: proxy URL; proxy credentials belong in `USER MAPPING`.
- `connect_timeout`: connection timeout.
- `request_timeout`: complete HTTP request timeout.
- `max_response_size`: maximum response body size in bytes; `0` means unlimited.
- `readonly`: prevents `INSERT`, `UPDATE`, and `DELETE` before requests reach the endpoint.
- `request_redirect` and `request_max_redirect`: redirect behavior.

Use `max_response_size` for public or untrusted endpoints because `rdf_fdw` loads retrieved RDF data into memory before converting it for PostgreSQL.

### User Mapping

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  user 'sparql_user',
  password 'secret'
);
```

v2.6.0 adds Bearer-token authentication:

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (
  token 'eyJhbGciOi...'
);
```

Proxy credentials are also `USER MAPPING` options:

```sql
CREATE USER MAPPING FOR app_user
SERVER dbpedia
OPTIONS (
  proxy_user 'proxy-user',
  proxy_password 'proxy-secret'
);
```

### Foreign Tables with rdfnode Columns

Declare foreign-table columns as `rdfnode` to preserve RDF terms, IRIs, blank nodes, language tags, and XSD datatypes.

```sql
CREATE FOREIGN TABLE dbpedia_films (
  film rdfnode OPTIONS (variable '?film'),
  name rdfnode OPTIONS (variable '?name'),
  year rdfnode OPTIONS (variable '?year')
)
SERVER dbpedia
OPTIONS (
  sparql $$
    SELECT ?film ?name ?year
    WHERE {
      ?film a dbo:Film ;
            rdfs:label ?name ;
            dbo:releaseDate ?year .
      FILTER (lang(?name) = 'en')
    }
  $$
);
```

Native PostgreSQL column types are deprecated for RDF values in v2.6.0. Existing native-typed tables continue to work, but they emit warnings and lose RDF term details.

### Querying and Pushdown

```sql
SELECT film, sparql.lex(name) AS title
FROM dbpedia_films
WHERE name = '"The Matrix"@en'::rdfnode
ORDER BY year
LIMIT 10;

EXPLAIN (VERBOSE, COSTS OFF)
SELECT *
FROM dbpedia_films
WHERE film = '<http://dbpedia.org/resource/The_Matrix>'::rdfnode;
```

`rdf_fdw` can push down `WHERE`, `LIMIT`, `ORDER BY`, `DISTINCT`, and supported comparisons/functions. Use `EXPLAIN VERBOSE` to inspect the generated remote SPARQL.

### Prefix Management

`rdf_fdw` provides catalog tables and helper functions under the `sparql` schema for reusable SPARQL prefixes:

```sql
SELECT sparql.add_context('default', 'Default SPARQL prefix context');
SELECT sparql.add_prefix('default', 'rdf',  'http://www.w3.org/1999/02/22-rdf-syntax-ns#');
SELECT sparql.add_prefix('default', 'rdfs', 'http://www.w3.org/2000/01/rdf-schema#');
SELECT sparql.add_prefix('default', 'xsd',  'http://www.w3.org/2001/XMLSchema#');
```

### Data Modification

Writable foreign tables can translate PostgreSQL `INSERT`, `UPDATE`, and `DELETE` into SPARQL UPDATE requests when the foreign table has the required SPARQL update pattern.

```sql
ALTER FOREIGN TABLE dbpedia_films OPTIONS (ADD readonly 'false');

INSERT INTO dbpedia_films(film, name)
VALUES (
  '<http://example.org/film/1>'::rdfnode,
  '"Example Film"@en'::rdfnode
);
```

Use `readonly = true` at the server or table level when an endpoint should never receive writes.

### Clone a Foreign Table

```sql
CALL rdf_fdw_clone_table(
  foreign_table := 'dbpedia_films',
  target_table  := 'dbpedia_films_local',
  fetch_size    := 1000,
  create_table  := true
);
```

`rdf_fdw_clone_table()` copies data from a foreign table into a local table in batches. v2.5 fixed several round-trip issues for RDF terms during cloning.

### SPARQL Functions

The `sparql` schema implements many SPARQL 1.1 functions and aggregates, including:

- aggregates such as `sparql.sum`, `sparql.avg`, `sparql.min`, `sparql.max`, `sparql.group_concat`, and `sparql.sample`
- RDF term helpers such as `sparql.isiri`, `sparql.isblank`, `sparql.isliteral`, `sparql.datatype`, `sparql.iri`, `sparql.strdt`, and `sparql.strlang`
- string functions such as `sparql.strlen`, `sparql.substr`, `sparql.ucase`, `sparql.lcase`, `sparql.contains`, and `sparql.replace`
- numeric, date/time, hash, and custom convenience functions

### Caveats

- PostgreSQL 9.5+ is the upstream baseline, but Pigsty packages target modern PostgreSQL majors listed in local metadata.
- Retrieved RDF data is accumulated in memory before conversion. Set `max_response_size`, use `LIMIT`, and keep remote result sets bounded.
- Prefer `rdfnode` columns. Native PostgreSQL typed columns are deprecated for RDF terms and will lose IRI/language/datatype information.
- Store secrets in `USER MAPPING`; do not put proxy credentials or endpoint tokens into `SERVER` options.
- Public SPARQL endpoints can be slow or rate-limited. Use `connect_timeout`, `request_timeout`, retries, and local materialization when needed.
