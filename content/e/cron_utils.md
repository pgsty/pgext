---
title: "cron_utils"
linkTitle: "cron_utils"
description: "Parse cron expressions and compute previous or next trigger times"
weight: 1140
categories: ["TIME"]
width: full
---

[**cron_utils**](https://github.com/Myshkouski/pg-cron-utils) : Parse cron expressions and compute previous or next trigger times


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1140** | {{< badge content="cron_utils" link="https://github.com/Myshkouski/pg-cron-utils" >}} | {{< ext "cron_utils" >}} | `0.1.0` | {{< category "TIME" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_when" >}} {{< ext "pgcalendar" >}} {{< ext "periods" >}} |

> [!Note] The PGXN 0.1.0 distribution is marked unstable; the control file marks the extension relocatable.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `cron_utils` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "cron_utils_18" "green" >}} {{< bg "17" "cron_utils_17" "green" >}} {{< bg "16" "cron_utils_16" "green" >}} {{< bg "15" "cron_utils_15" "green" >}} {{< bg "14" "cron_utils_14" "green" >}} | `cron_utils_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-cron-utils" "green" >}} {{< bg "17" "postgresql-17-cron-utils" "green" >}} {{< bg "16" "postgresql-16-cron-utils" "green" >}} {{< bg "15" "postgresql-15-cron-utils" "green" >}} {{< bg "14" "postgresql-14-cron-utils" "green" >}} | `postgresql-$v-cron-utils` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "cron_utils_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-cron-utils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-cron-utils : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cron_utils_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [cron_utils_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cron_utils_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [cron_utils_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cron_utils_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [cron_utils_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cron_utils_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [cron_utils_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cron_utils_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [cron_utils_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cron_utils_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `cron_utils_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [cron_utils_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cron_utils_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-cron-utils` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-cron-utils` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 KiB | [postgresql-18-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-18-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cron_utils_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [cron_utils_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cron_utils_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [cron_utils_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cron_utils_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [cron_utils_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cron_utils_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [cron_utils_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cron_utils_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [cron_utils_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cron_utils_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `cron_utils_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [cron_utils_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cron_utils_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-cron-utils` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-cron-utils` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 KiB | [postgresql-17-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-17-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cron_utils_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [cron_utils_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cron_utils_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [cron_utils_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cron_utils_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [cron_utils_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cron_utils_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [cron_utils_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cron_utils_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [cron_utils_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cron_utils_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `cron_utils_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [cron_utils_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cron_utils_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-cron-utils` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.4 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.4 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-cron-utils` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 KiB | [postgresql-16-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-16-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cron_utils_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [cron_utils_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cron_utils_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [cron_utils_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cron_utils_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [cron_utils_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cron_utils_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [cron_utils_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cron_utils_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [cron_utils_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cron_utils_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `cron_utils_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [cron_utils_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cron_utils_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-cron-utils` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-cron-utils` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 KiB | [postgresql-15-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-15-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `cron_utils_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.6 KiB | [cron_utils_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/cron_utils_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [cron_utils_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/cron_utils_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `cron_utils_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [cron_utils_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/cron_utils_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [cron_utils_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/cron_utils_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `cron_utils_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.6 KiB | [cron_utils_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/cron_utils_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `cron_utils_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [cron_utils_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/cron_utils_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-cron-utils` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-cron-utils` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.5 KiB | [postgresql-14-cron-utils_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/cron-utils/postgresql-14-cron-utils_0.1.0-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Myshkouski/pg-cron-utils" title="Repository" icon="github" subtitle="github.com/Myshkouski/pg-cron-utils" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="cron_utils-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg cron_utils;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install cron_utils;		# install via package name, for the active PG version

pig install cron_utils -v 18;   # install for PG 18
pig install cron_utils -v 17;   # install for PG 17
pig install cron_utils -v 16;   # install for PG 16
pig install cron_utils -v 15;   # install for PG 15
pig install cron_utils -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION cron_utils;
```

## Usage

Sources:

- [pg_cron_utils 0.1.0 README](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/README.md)
- [Extension control file](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils.control)
- [SQL definitions](https://github.com/Myshkouski/pg-cron-utils/blob/v0.1.0/cron_utils--0.1.0.sql)

`cron_utils` parses five-field cron expressions and calculates trigger timestamps inside PostgreSQL. It is a scheduling utility, not a job runner: use it to preview, validate, or query a schedule, and use a scheduler such as `pg_cron` separately to execute work.

### Core Workflow

```sql
CREATE EXTENSION cron_utils;

-- First trigger at or after the supplied time.
SELECT cron_first_trigger('0 9 * * 1-5', now());

-- Last trigger before the supplied time (strict is true by default).
SELECT cron_last_trigger('0 9 * * 1-5', now());

-- Next five hourly triggers.
SELECT *
FROM cron_iterate_n('0 * * * *', now(), false, 'next', 5);
```

To inspect a bounded window:

```sql
SELECT *
FROM cron_first_last_triggers(
  '0 0 * * *',
  date_trunc('month', now()),
  date_trunc('month', now()) + interval '1 month'
);
```

Either returned boundary can be `NULL` when the expression has no trigger in the window.

### Important Objects

- `cron_parts` is the parsed representation of the minute, hour, day, month, and day-of-week fields.
- `parse_cron(text)` parses `*`, lists, ranges, and step syntax.
- `cron_first_trigger(expr, base_time, strict)` searches forward. With `strict = true`, a trigger exactly at `base_time` is skipped.
- `cron_last_trigger(expr, base_time, strict)` searches backward and defaults to strict matching.
- `cron_first_last_triggers(expr, start_time, end_time)` returns the first and last matches in a window.
- `cron_iterate_n(expr, base_time, strict, direction, max_matches)` returns consecutive matches in the `next` or `prev` direction.

### Semantics and Caveats

Expressions use the standard five fields `minute hour day month dow`; seconds are not accepted. Day-of-week uses `1 = Monday` through `7 = Sunday`. Results are `timestamptz`, so session time zone affects the displayed local time and daylight-saving transitions should be tested for local-time schedules.

The extension is pure SQL/PL/pgSQL, relocatable, and has no `pg_cron` dependency. Its functions are declared immutable and parallel-safe. Version 0.1.0 is marked `unstable` in the control metadata, so pin behavior and retest edge cases before embedding it in a critical scheduler.
