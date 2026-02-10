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
| **8680** | {{< badge content="nominatim_fdw" link="https://github.com/jimjonesbr/nominatim_fdw" >}} | {{< ext "nominatim_fdw" >}} | `1.1.0` | {{< category "FDW" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `nominatim_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.0` | {{< bg "18" "nominatim_fdw_18" "green" >}} {{< bg "17" "nominatim_fdw_17" "green" >}} {{< bg "16" "nominatim_fdw_16" "green" >}} {{< bg "15" "nominatim_fdw_15" "green" >}} {{< bg "14" "nominatim_fdw_14" "green" >}} {{< bg "13" "nominatim_fdw_13" "red" >}} | `nominatim_fdw_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "postgresql-18-nominatim-fdw" "green" >}} {{< bg "17" "postgresql-17-nominatim-fdw" "green" >}} {{< bg "16" "postgresql-16-nominatim-fdw" "green" >}} {{< bg "15" "postgresql-15-nominatim-fdw" "green" >}} {{< bg "14" "postgresql-14-nominatim-fdw" "green" >}} {{< bg "13" "postgresql-13-nominatim-fdw" "red" >}} | `postgresql-$v-nominatim-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.1.0" "nominatim_fdw_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "nominatim_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-18-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-17-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-16-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-15-nominatim-fdw : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "postgresql-14-nominatim-fdw : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-nominatim-fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_18` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_18` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/nominatim_fdw_18-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.7 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.9 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.6 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 56.6 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 56.0 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.8 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-nominatim-fdw` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.7 KiB | [postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-18-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_17` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_17` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/nominatim_fdw_17-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.5 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.6 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.7 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.3 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 64.0 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.5 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-nominatim-fdw` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.3 KiB | [postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-17-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_16` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_16` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.6 KiB | [nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/nominatim_fdw_16-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.7 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.9 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.1 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.8 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.5 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-nominatim-fdw` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.3 KiB | [postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-16-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.3 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_15` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/nominatim_fdw_15-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.8 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.1 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.8 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.5 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-nominatim-fdw` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.4 KiB | [postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-15-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `nominatim_fdw_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 29.2 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel8.10.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel9.7.aarch64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.9 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.x86_64.rpm) |
| `nominatim_fdw_14` | `1.1.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 30.5 KiB | [nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/nominatim_fdw_14-1.1.0-1PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 52.4 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 51.7 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 52.4 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 51.8 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.0 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 63.6 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 54.5 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-nominatim-fdw` | `1.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 54.3 KiB | [postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/nominatim-fdw/postgresql-14-nominatim-fdw_1.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jimjonesbr/nominatim_fdw" title="Repository" icon="github" subtitle="github.com/jimjonesbr/nominatim_fdw" >}}
{{< /cards >}}


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
