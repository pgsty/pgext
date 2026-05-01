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
| **8760** | {{< badge content="rdf_fdw" link="https://github.com/jimjonesbr/rdf_fdw" >}} | {{< ext "rdf_fdw" >}} | `2.5.0` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "postgres_fdw" >}} {{< ext "sparql" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `rdf_fdw` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "rdf_fdw_18" "green" >}} {{< bg "17" "rdf_fdw_17" "green" >}} {{< bg "16" "rdf_fdw_16" "green" >}} {{< bg "15" "rdf_fdw_15" "green" >}} {{< bg "14" "rdf_fdw_14" "green" >}} | `rdf_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.5.0` | {{< bg "18" "postgresql-18-rdf-fdw" "green" >}} {{< bg "17" "postgresql-17-rdf-fdw" "green" >}} {{< bg "16" "postgresql-16-rdf-fdw" "green" >}} {{< bg "15" "postgresql-15-rdf-fdw" "green" >}} {{< bg "14" "postgresql-14-rdf-fdw" "green" >}} | `postgresql-$v-rdf-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "rdf_fdw_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-18-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-17-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-16-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-15-rdf-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.5.0" "postgresql-14-rdf-fdw : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_18` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.6 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_18-2.5.0-2PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_18` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 137.0 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_18-2.5.0-2PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_18` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.7 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_18-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_18` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 135.8 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_18-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_18` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.7 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_18-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_18` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 137.2 KiB | [rdf_fdw_18-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_18-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 332.2 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.6 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 332.2 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 323.0 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 355.2 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 349.6 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.5 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 337.5 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 339.9 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-rdf-fdw` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 335.3 KiB | [postgresql-18-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-18-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_17` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.7 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_17-2.5.0-2PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_17` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 137.0 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_17-2.5.0-2PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_17` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.8 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_17-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_17` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 135.9 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_17-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_17` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.7 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_17-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_17` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 137.2 KiB | [rdf_fdw_17-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_17-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 331.7 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.4 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.3 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 323.1 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 376.6 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 371.2 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.0 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 336.9 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 338.9 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-rdf-fdw` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 334.4 KiB | [postgresql-17-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-17-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 145.8 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_16-2.5.0-2PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 137.0 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_16-2.5.0-2PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_16` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 139.9 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_16-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_16` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 135.9 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_16-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_16` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 140.4 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_16-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_16` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 137.3 KiB | [rdf_fdw_16-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_16-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 332.7 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 322.2 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.7 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 322.8 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 374.5 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 368.9 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.1 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 337.0 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 339.4 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-rdf-fdw` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 334.9 KiB | [postgresql-16-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-16-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_15` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 147.3 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_15-2.5.0-2PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_15` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 138.3 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_15-2.5.0-2PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_15` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 142.6 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_15-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_15` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.9 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_15-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_15` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 142.7 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_15-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_15` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.7 KiB | [rdf_fdw_15-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_15-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 333.2 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 323.8 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 332.6 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 323.9 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 375.7 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 370.1 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.7 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 338.0 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 340.1 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-rdf-fdw` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 335.7 KiB | [postgresql-15-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-15-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdf_fdw_14` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 147.2 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rdf_fdw_14-2.5.0-2PIGSTY.el8.x86_64.rpm) |
| `rdf_fdw_14` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 138.3 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rdf_fdw_14-2.5.0-2PIGSTY.el8.aarch64.rpm) |
| `rdf_fdw_14` | `2.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 142.5 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rdf_fdw_14-2.5.0-2PIGSTY.el9.x86_64.rpm) |
| `rdf_fdw_14` | `2.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 137.8 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rdf_fdw_14-2.5.0-2PIGSTY.el9.aarch64.rpm) |
| `rdf_fdw_14` | `2.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 142.6 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdf_fdw_14-2.5.0-2PIGSTY.el10.x86_64.rpm) |
| `rdf_fdw_14` | `2.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 139.7 KiB | [rdf_fdw_14-2.5.0-2PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdf_fdw_14-2.5.0-2PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 332.8 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 323.2 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 331.8 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 323.8 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 375.4 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 369.9 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.7 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 338.0 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 340.0 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-rdf-fdw` | `2.5.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 335.4 KiB | [postgresql-14-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/r/rdf-fdw/postgresql-14-rdf-fdw_2.5.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/rdf_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/rdf_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="rdf_fdw-2.5.0.tar.gz" >}}
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

> Sources: [README](https://github.com/jimjonesbr/rdf_fdw/blob/main/README.md), [v2.4 release](https://github.com/jimjonesbr/rdf_fdw/releases/tag/v2.4)

`rdf_fdw` is a foreign data wrapper for RDF triplestores exposed through SPARQL endpoints. It lets PostgreSQL query RDF data with SQL, supports SQL clause pushdown, adds an `rdfnode` type for RDF terms, and includes SPARQL 1.1 function support.

### Register a SPARQL Endpoint

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

In `v2.4`, proxy credentials were moved out of `SERVER` options and into `USER MAPPING` for security.

### User Mapping and Foreign Tables

```sql
CREATE USER MAPPING FOR postgres
SERVER dbpedia
OPTIONS (user 'admin', password 'secret');
```

`rdf_fdw` works by declaring foreign tables that embed SPARQL queries and map result variables to PostgreSQL columns. The README also highlights native RDF node handling through the custom `rdfnode` type.

```sql
CREATE FOREIGN TABLE film (
  film_id text OPTIONS (variable '?film', nodetype 'iri'),
  name text OPTIONS (variable '?name', nodetype 'literal', literal_type 'xsd:string')
)
SERVER dbpedia
OPTIONS (sparql 'SELECT ?film ?name WHERE { ?film dbp:name ?name }');
```

### Pushdown, DML, and Helpers

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

The README lists utility functions including:

- `rdf_fdw_version()`
- `rdf_fdw_settings()`
- `rdf_fdw_clone_table()`

It also documents broader SPARQL function coverage, including aggregates, string functions, numeric functions, date/time functions, hash functions, and custom functions.

### Caveats

The current README warns that retrieved RDF data is loaded into memory before conversion for PostgreSQL, so large result sets require adequate PostgreSQL memory. It also documents PostgreSQL 9.5+ as the supported baseline.
