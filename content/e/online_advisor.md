---
title: "online_advisor"
linkTitle: "online_advisor"
description: "Suggest missing indexes and extended statistics online"
weight: 5270
categories: ["ADMIN"]
width: full
---

[**online_advisor**](https://github.com/knizhnik/online_advisor) : Suggest missing indexes and extended statistics online


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5270** | {{< badge content="online_advisor" link="https://github.com/knizhnik/online_advisor" >}} | {{< ext "online_advisor" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "index_advisor" >}} {{< ext "hypopg" >}} {{< ext "pg_qualstats" >}} |

> [!Note] Requires shared_preload_libraries=online_advisor on PostgreSQL 14-16; PGSTY backports upstream PG18 hook compatibility.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `online_advisor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "online_advisor_18" "green" >}} {{< bg "17" "online_advisor_17" "green" >}} {{< bg "16" "online_advisor_16" "green" >}} {{< bg "15" "online_advisor_15" "green" >}} {{< bg "14" "online_advisor_14" "green" >}} | `online_advisor_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-online-advisor" "green" >}} {{< bg "17" "postgresql-17-online-advisor" "green" >}} {{< bg "16" "postgresql-16-online-advisor" "green" >}} {{< bg "15" "postgresql-15-online-advisor" "green" >}} {{< bg "14" "postgresql-14-online-advisor" "green" >}} | `postgresql-$v-online-advisor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "online_advisor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "online_advisor_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-online-advisor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-online-advisor : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `online_advisor_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [online_advisor_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/online_advisor_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `online_advisor_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.4 KiB | [online_advisor_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/online_advisor_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `online_advisor_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.5 KiB | [online_advisor_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/online_advisor_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `online_advisor_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [online_advisor_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/online_advisor_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `online_advisor_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.7 KiB | [online_advisor_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/online_advisor_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `online_advisor_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.0 KiB | [online_advisor_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/online_advisor_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-online-advisor` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.6 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.4 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.6 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.7 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.0 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.9 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.7 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-online-advisor` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-18-online-advisor_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-18-online-advisor_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `online_advisor_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [online_advisor_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/online_advisor_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `online_advisor_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.4 KiB | [online_advisor_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/online_advisor_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `online_advisor_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.5 KiB | [online_advisor_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/online_advisor_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `online_advisor_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.6 KiB | [online_advisor_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/online_advisor_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `online_advisor_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.7 KiB | [online_advisor_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/online_advisor_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `online_advisor_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.0 KiB | [online_advisor_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/online_advisor_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-online-advisor` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.5 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.4 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.6 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.1 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.9 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.7 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-online-advisor` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 31.1 KiB | [postgresql-17-online-advisor_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-17-online-advisor_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `online_advisor_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.7 KiB | [online_advisor_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/online_advisor_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `online_advisor_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.6 KiB | [online_advisor_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/online_advisor_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `online_advisor_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.7 KiB | [online_advisor_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/online_advisor_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `online_advisor_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.9 KiB | [online_advisor_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/online_advisor_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `online_advisor_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [online_advisor_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/online_advisor_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `online_advisor_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.2 KiB | [online_advisor_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/online_advisor_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-online-advisor` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.6 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.0 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.1 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.4 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.2 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.7 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-online-advisor` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 30.9 KiB | [postgresql-16-online-advisor_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-16-online-advisor_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `online_advisor_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.7 KiB | [online_advisor_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/online_advisor_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `online_advisor_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.6 KiB | [online_advisor_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/online_advisor_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `online_advisor_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.7 KiB | [online_advisor_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/online_advisor_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `online_advisor_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.9 KiB | [online_advisor_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/online_advisor_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `online_advisor_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.8 KiB | [online_advisor_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/online_advisor_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `online_advisor_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.2 KiB | [online_advisor_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/online_advisor_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-online-advisor` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.7 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.7 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.7 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.0 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.9 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.3 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.9 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.2 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.8 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-online-advisor` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 30.9 KiB | [postgresql-15-online-advisor_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-15-online-advisor_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `online_advisor_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.6 KiB | [online_advisor_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/online_advisor_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `online_advisor_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.5 KiB | [online_advisor_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/online_advisor_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `online_advisor_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.6 KiB | [online_advisor_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/online_advisor_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `online_advisor_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.8 KiB | [online_advisor_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/online_advisor_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `online_advisor_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.7 KiB | [online_advisor_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/online_advisor_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `online_advisor_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.2 KiB | [online_advisor_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/online_advisor_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-online-advisor` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.4 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.5 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.5 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.7 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 37.7 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 37.0 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.7 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.0 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 31.5 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-online-advisor` | `1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 30.7 KiB | [postgresql-14-online-advisor_1.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/o/online-advisor/postgresql-14-online-advisor_1.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/knizhnik/online_advisor" title="Repository" icon="github" subtitle="github.com/knizhnik/online_advisor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="online_advisor-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg online_advisor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install online_advisor;		# install via package name, for the active PG version

pig install online_advisor -v 18;   # install for PG 18
pig install online_advisor -v 17;   # install for PG 17
pig install online_advisor -v 16;   # install for PG 16
pig install online_advisor -v 15;   # install for PG 15
pig install online_advisor -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'online_advisor';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION online_advisor;
```

## Usage

Sources:

- [Official README for version 1.0](https://github.com/knizhnik/online_advisor/blob/1.0/README.md)
- [Extension control file](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.control)
- [Version 1.0 SQL objects](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor--1.0.sql)
- [Sample preload configuration](https://github.com/knizhnik/online_advisor/blob/1.0/online_advisor.conf)

`online_advisor` observes PostgreSQL execution plans and workload timing, then recommends indexes, extended statistics, or prepared statements. It reports candidates only; it never creates an index or statistics object automatically.

### Core Workflow

Preload the library and restart PostgreSQL:

```conf
shared_preload_libraries = 'online_advisor'
```

Create and activate the extension in each database whose workload should be observed:

```sql
CREATE EXTENSION online_advisor;

-- Calling an extension function activates collection in this database.
SELECT get_executor_stats();
```

After representative workload has run, inspect the recommendations:

```sql
SELECT * FROM proposed_indexes;
SELECT * FROM proposed_statistics;
SELECT * FROM get_executor_stats();

-- Keep separate index candidates instead of combining compatible clauses.
SELECT * FROM propose_indexes(combine => false);
```

Review each generated `create_index` or `create_statistics` statement before applying it. Run `ANALYZE` after creating an index or statistics object so the planner can use current statistics.

### Objects and Settings

- `proposed_indexes`: view over `propose_indexes(combine, reset)` with filtering volume, call count, elapsed time, and a candidate `CREATE INDEX` statement.
- `proposed_statistics`: view over `propose_statistics(combine, reset)` with misestimation, call count, elapsed time, and a candidate `CREATE STATISTICS` statement.
- `get_executor_stats(reset)`: returns aggregate planning and execution time, query count, and planning-overhead ratios.
- `online_advisor.filtered_threshold`: minimum filtered-row count considered for an index proposal; default `1000`.
- `online_advisor.misestimation_threshold`: actual-to-estimated row ratio considered for statistics; default `10`.
- `online_advisor.min_rows`: minimum returned rows for misestimation analysis; default `1000`.
- `online_advisor.max_index_proposals` and `online_advisor.max_stat_proposals`: proposal capacities; set them before the extension is activated.
- `online_advisor.do_instrumentation`, `online_advisor.log_duration`, and `online_advisor.prepare_threshold`: control collection and prepared-statement advice.

### Caveats

- Instrumentation adds workload overhead; measure it on the target system and disable collection when it is not needed.
- The index heuristic does not reason about operator ordering in compound indexes, join indexes, or indexes used only to avoid sorting.
- The extension does not estimate the benefit of a proposed index. Use plan review or a hypothetical-index tool before building expensive indexes.
- Advice is database-local and depends on the workload observed since activation or reset.
