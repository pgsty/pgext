---
title: "ddl_historization"
linkTitle: "ddl_historization"
description: "Historize the ddl changes inside PostgreSQL database"
weight: 4310
categories: ["UTIL"]
width: full
---

[**ddl_historization**](https://github.com/rodo/pg_ddl_historization) : Historize the ddl changes inside PostgreSQL database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4310** | {{< badge content="ddl_historization" link="https://github.com/rodo/pg_ddl_historization" >}} | {{< ext "ddl_historization" >}} | `0.0.7` | {{< category "UTIL" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|    **Need By**    | {{< ext "schedoc" >}} |
|   **See Also**    | {{< ext "pg_readme" >}} {{< ext "data_historization" >}} {{< ext "table_version" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} {{< ext "zstd" >}} {{< ext "http" >}} {{< ext "pg_net" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `ddl_historization` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "ddl_historization_18" "green" >}} {{< bg "17" "ddl_historization_17" "green" >}} {{< bg "16" "ddl_historization_16" "green" >}} {{< bg "15" "ddl_historization_15" "green" >}} {{< bg "14" "ddl_historization_14" "green" >}} {{< bg "13" "ddl_historization_13" "green" >}} | `ddl_historization_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.0.7` | {{< bg "18" "postgresql-18-ddl-historization" "green" >}} {{< bg "17" "postgresql-17-ddl-historization" "green" >}} {{< bg "16" "postgresql-16-ddl-historization" "green" >}} {{< bg "15" "postgresql-15-ddl-historization" "green" >}} {{< bg "14" "postgresql-14-ddl-historization" "green" >}} {{< bg "13" "postgresql-13-ddl-historization" "green" >}} | `postgresql-$v-ddl-historization` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "ddl_historization_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-18-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-17-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-16-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-15-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-14-ddl-historization : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.0.7" "postgresql-13-ddl-historization : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_18` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_18-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_18` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_18-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_18` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_18-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_18` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_18-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_18` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_18-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_18` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_18-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_18-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-18-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-18-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_17` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_17` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_17` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_17` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_17` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_17-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_17` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_17-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_17-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-17-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_16` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_16` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_16` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_16` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_16` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_16-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_16` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_16-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_16-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-16-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_15` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_15` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_15` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_15` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_15` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_15-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_15` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_15-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_15-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-15-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_14` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_14` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_14` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_14` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_14` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_14-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_14` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_14-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_14-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-14-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddl_historization_13` | `0.0.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.0 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddl_historization_13-0.0.7-1PIGSTY.el8.x86_64.rpm) |
| `ddl_historization_13` | `0.0.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.0 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddl_historization_13-0.0.7-1PIGSTY.el8.aarch64.rpm) |
| `ddl_historization_13` | `0.0.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddl_historization_13-0.0.7-1PIGSTY.el9.x86_64.rpm) |
| `ddl_historization_13` | `0.0.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.8 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddl_historization_13-0.0.7-1PIGSTY.el9.aarch64.rpm) |
| `ddl_historization_13` | `0.0.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.9 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddl_historization_13-0.0.7-1PIGSTY.el10.x86_64.rpm) |
| `ddl_historization_13` | `0.0.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.9 KiB | [ddl_historization_13-0.0.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddl_historization_13-0.0.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-ddl-historization` | `0.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.0 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-ddl-historization` | `0.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.7 KiB | [postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddl-historization/postgresql-13-ddl-historization_0.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rodo/pg_ddl_historization" title="Repository" icon="github" subtitle="github.com/rodo/pg_ddl_historization" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_ddl_historization-0.0.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg ddl_historization;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ddl_historization;		# install via package name, for the active PG version

pig install ddl_historization -v 18;   # install for PG 18
pig install ddl_historization -v 17;   # install for PG 17
pig install ddl_historization -v 16;   # install for PG 16
pig install ddl_historization -v 15;   # install for PG 15
pig install ddl_historization -v 14;   # install for PG 14
pig install ddl_historization -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ddl_historization CASCADE; -- requires plpgsql
```
