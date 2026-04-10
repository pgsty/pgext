---
title: "nominatim_fdw"
linkTitle: "nominatim_fdw"
description: "Nominatim Foreign Data Wrapper for PostgreSQL"
weight: 8680
categories: ["FDW"]
width: full
---

[**nominatim_fdw**](https://github.com/jimjonesbr/nominatim_fdw) : Nominatim Foreign Data Wrapper for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8680** | {{< badge content="nominatim_fdw" link="https://github.com/jimjonesbr/nominatim_fdw" >}} | {{< ext "nominatim_fdw" >}} | `1.2` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `nominatim_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.2` | {{< bg "18" "nominatim_fdw_18" "green" >}} {{< bg "17" "nominatim_fdw_17" "green" >}} {{< bg "16" "nominatim_fdw_16" "green" >}} {{< bg "15" "nominatim_fdw_15" "green" >}} {{< bg "14" "nominatim_fdw_14" "green" >}} | `nominatim_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2` | {{< bg "18" "postgresql-18-nominatim-fdw" "green" >}} {{< bg "17" "postgresql-17-nominatim-fdw" "green" >}} {{< bg "16" "postgresql-16-nominatim-fdw" "green" >}} {{< bg "15" "postgresql-15-nominatim-fdw" "green" >}} {{< bg "14" "postgresql-14-nominatim-fdw" "green" >}} | `postgresql-$v-nominatim-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.2" "nominatim_fdw_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.2" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.2" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_18` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-nominatim-fdw` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.9 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.8 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.8 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 52.1 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 56.6 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.3 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.6 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.4 KiB | [postgresql-18-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_17` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.0 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-nominatim-fdw` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.6 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.6 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.6 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.3 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.2 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.5 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.9 KiB | [postgresql-17-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_16` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.0 KiB | [nominatim_fdw_16-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-16-nominatim-fdw` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.6 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.8 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.5 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.9 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.6 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.4 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.9 KiB | [postgresql-16-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_15` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.7 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-15-nominatim-fdw` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.7 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.8 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.5 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.9 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.5 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.3 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.9 KiB | [postgresql-15-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_14` | `1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.8 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.1 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.4 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.2-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.2-1PGDG.rhel10.1.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-14-nominatim-fdw` | `1.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.8 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.8 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.7 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 63.7 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.4 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.3 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 53.9 KiB | [postgresql-14-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/nominatim_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/nominatim_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="nominatim_fdw-1.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg nominatim_fdw;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install nominatim_fdw;		# install via package name, for the active PG version

pig install nominatim_fdw -v 18;   # install for PG 18
pig install nominatim_fdw -v 17;   # install for PG 17
pig install nominatim_fdw -v 16;   # install for PG 16
pig install nominatim_fdw -v 15;   # install for PG 15
pig install nominatim_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION nominatim_fdw;
```




## Usage

> [nominatim_fdw: Nominatim Foreign Data Wrapper for PostgreSQL](https://github.com/jimjonesbr/nominatim_fdw)

This FDW provides access to [Nominatim](https://nominatim.org/) geocoding services directly from PostgreSQL using SQL functions rather than traditional foreign tables.

### Create Server

```sql
CREATE EXTENSION nominatim_fdw;

CREATE SERVER osm FOREIGN DATA WRAPPER nominatim_fdw
  OPTIONS (url 'https://nominatim.openstreetmap.org');
```

**Server Options:** `url` (required, Nominatim endpoint URL), `http_proxy`, `proxy_user`, `proxy_user_password`, `connect_timeout` (default 300 seconds), `max_connect_retry` (default 3), `max_request_redirect` (0 = unlimited).

### Geocoding (Address to Coordinates)

Structured search:

```sql
SELECT osm_id, ref, lon, lat, boundingbox
FROM nominatim_search(
  server_name => 'osm',
  street => 'Neubrueckenstrasse 63',
  city => 'Muenster',
  country => 'Germany'
);
```

Free-form search:

```sql
SELECT osm_id, display_name, lon, lat
FROM nominatim_search(
  server_name => 'osm',
  q => '1600 Pennsylvania Avenue, Washington DC'
);
```

**Parameters:** `q` (free-form query), `street`, `city`, `county`, `state`, `country`, `postalcode`, `amenity`, `limit` (default 10), `addressdetails`, `extratags`, `namedetails`.

### Reverse Geocoding (Coordinates to Address)

```sql
SELECT osm_id, display_name, boundingbox
FROM nominatim_reverse(
  server_name => 'osm',
  lon => -77.0365,
  lat => 38.8977,
  zoom => 18,
  addressdetails => true
);
```

**Parameters:** `lon`, `lat` (required), `zoom` (default 18), `addressdetails`, `extratags`, `namedetails`, `layer`.

### OSM Object Lookup

```sql
SELECT osm_id, display_name
FROM nominatim_lookup(
  server_name => 'osm',
  osm_ids => 'W121736959,R123456'
);
```

**Parameters:** `osm_ids` (comma-separated list of OSM IDs with type prefix: N=node, W=way, R=relation), `addressdetails`, `extratags`, `namedetails`.

### Version Check

```sql
SELECT nominatim_fdw_version();
```
