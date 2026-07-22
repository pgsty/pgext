---
title: "pgmonitor"
linkTitle: "pgmonitor"
description: "Collector-friendly metric views and background refresh worker"
weight: 6070
categories: ["STAT"]
width: full
---

[**pgmonitor**](https://github.com/CrunchyData/pgmonitor-extension) : Collector-friendly metric views and background refresh worker


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6070** | {{< badge content="pgmonitor" link="https://github.com/CrunchyData/pgmonitor-extension" >}} | {{< ext "pgmonitor" >}} | `2.2.0` | {{< category "STAT" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgnodemx" >}} {{< ext "system_stats" >}} {{< ext "pg_stat_monitor" >}} {{< ext "pg_profile" >}} |

> [!Note] Metric objects work without preloading; the optional background worker requires shared_preload_libraries=pgmonitor_bgw.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmonitor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "pgmonitor_18" "green" >}} {{< bg "17" "pgmonitor_17" "green" >}} {{< bg "16" "pgmonitor_16" "green" >}} {{< bg "15" "pgmonitor_15" "green" >}} {{< bg "14" "pgmonitor_14" "green" >}} | `pgmonitor_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "postgresql-18-pgmonitor" "green" >}} {{< bg "17" "postgresql-17-pgmonitor" "green" >}} {{< bg "16" "postgresql-16-pgmonitor" "green" >}} {{< bg "15" "postgresql-15-pgmonitor" "green" >}} {{< bg "14" "postgresql-14-pgmonitor" "green" >}} | `postgresql-$v-pgmonitor` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "pgmonitor_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-18-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-17-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-16-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-15-pgmonitor : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.2.0" "postgresql-14-pgmonitor : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmonitor_18` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmonitor_18-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmonitor_18` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.3 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmonitor_18-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmonitor_18` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmonitor_18-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmonitor_18` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.5 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmonitor_18-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmonitor_18` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmonitor_18-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmonitor_18` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pgmonitor_18-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmonitor_18-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmonitor` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.7 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.5 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 36.8 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.6 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 38.9 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.4 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.0 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.8 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.2 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmonitor` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 37.9 KiB | [postgresql-18-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-18-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmonitor_17` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmonitor_17-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmonitor_17` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.3 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmonitor_17-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmonitor_17` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmonitor_17-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmonitor_17` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.5 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmonitor_17-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmonitor_17` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmonitor_17-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmonitor_17` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pgmonitor_17-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmonitor_17-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmonitor` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.7 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.5 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 36.7 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.6 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 43.4 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 42.9 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.9 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.8 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.2 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmonitor` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.0 KiB | [postgresql-17-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-17-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmonitor_16` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmonitor_16-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmonitor_16` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.2 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmonitor_16-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmonitor_16` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmonitor_16-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmonitor_16` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.5 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmonitor_16-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmonitor_16` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmonitor_16-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmonitor_16` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pgmonitor_16-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmonitor_16-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmonitor` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.7 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.5 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 36.7 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.6 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.9 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 42.4 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 37.9 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.8 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.2 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmonitor` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.0 KiB | [postgresql-16-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-16-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmonitor_15` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmonitor_15-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmonitor_15` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.2 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmonitor_15-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmonitor_15` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.8 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmonitor_15-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmonitor_15` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.5 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmonitor_15-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmonitor_15` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmonitor_15-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmonitor_15` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pgmonitor_15-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmonitor_15-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmonitor` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.7 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.5 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 36.7 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.6 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 42.9 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 42.5 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.0 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.8 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.2 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmonitor` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.0 KiB | [postgresql-15-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-15-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmonitor_14` | `2.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.1 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmonitor_14-2.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmonitor_14` | `2.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.2 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmonitor_14-2.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmonitor_14` | `2.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 31.9 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmonitor_14-2.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmonitor_14` | `2.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.6 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmonitor_14-2.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmonitor_14` | `2.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.0 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmonitor_14-2.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmonitor_14` | `2.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.9 KiB | [pgmonitor_14-2.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmonitor_14-2.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmonitor` | `2.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 36.7 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 36.5 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 36.7 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 36.5 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 41.5 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 41.1 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 38.0 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 37.8 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 38.2 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmonitor` | `2.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 38.0 KiB | [postgresql-14-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmonitor/postgresql-14-pgmonitor_2.2.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrunchyData/pgmonitor-extension" title="Repository" icon="github" subtitle="github.com/CrunchyData/pgmonitor-extension" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmonitor-extension-2.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmonitor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmonitor;		# install via package name, for the active PG version

pig install pgmonitor -v 18;   # install for PG 18
pig install pgmonitor -v 17;   # install for PG 17
pig install pgmonitor -v 16;   # install for PG 16
pig install pgmonitor -v 15;   # install for PG 15
pig install pgmonitor -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgmonitor_bgw';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmonitor;
```

## Usage

Sources:

- [pgmonitor-extension v2.2.0 README](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/README.md)
- [pgmonitor v2.2.0 control file](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/pgmonitor.control)
- [pgmonitor-extension v2.2.0 release notes](https://github.com/CrunchyData/pgmonitor-extension/releases/tag/v2.2.0)

pgmonitor exposes PostgreSQL monitoring metrics through a curated set of views, materialized views, and tables for external collectors. Its SQL metrics work without a background worker; the optional pgmonitor_bgw worker periodically refreshes materialized data.

### Create the Extension

Create a dedicated schema and install pgmonitor there:

    CREATE SCHEMA pgmonitor_ext;
    CREATE EXTENSION pgmonitor SCHEMA pgmonitor_ext;

Grant collectors only the access they need to the metric objects. Some underlying PostgreSQL statistics remain subject to built-in role and row-visibility rules.

### Collect Metrics

External agents can select the active objects described by the extension's configuration tables:

    SELECT *
    FROM pgmonitor_ext.metric_views
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_matviews
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_tables
    WHERE active;

These tables describe metric name, activation, scope, and refresh interval. Query the installed definitions rather than assuming every metric is enabled on every PostgreSQL version.

The metric surface includes activity, database and table statistics, locks, replication, WAL and archive status, vacuum progress, settings, checkpoints, and extension-specific views when their dependencies are available.

### Refresh Materialized Metrics Manually

Without the background worker, invoke the refresh procedure for the configured schema and metric:

    CALL pgmonitor_ext.refresh_metrics(
      'pgmonitor_ext',
      'pg_stat_statements'
    );

Use names returned by metric_matviews; do not assume the example metric is installed or active. The extension retains a legacy refresh function for compatibility, but new integrations should use the documented procedure.

### Optional Background Worker

To schedule refreshes inside PostgreSQL:

    shared_preload_libraries = 'pgmonitor_bgw'
    pgmonitor_bgw.dbname = 'postgres,app'
    pgmonitor_bgw.role = 'postgres'
    pgmonitor_bgw.interval = 30

Restart PostgreSQL after changing shared_preload_libraries. pgmonitor_bgw.dbname is required and lists the databases to maintain. Upstream v2.2 currently requires the worker role to be a superuser; use the narrowest controlled role and protect its credentials and settings.

### Object Index

- metric_views: directly queried metric views and their collection metadata.
- metric_matviews: materialized metrics and refresh intervals.
- metric_tables: table-backed metrics and maintenance metadata.
- refresh_metrics(schema, name): refresh procedure for one configured metric.
- pgmonitor_bgw.dbname: databases processed by the optional worker.
- pgmonitor_bgw.role: role used for refresh work.
- pgmonitor_bgw.interval: worker loop interval in seconds.

### Version 2.2 and Caveats

Version 2.2 removes the settings-checksum metric, fixes the legacy refresh path on PostgreSQL 13, and reduces routine log noise.

- Metric queries add load to shared statistics, catalogs, and extension objects. Set collection intervals from measured cost.
- A healthy collector connection does not prove materialized views are fresh; monitor their timestamps and worker logs.
- The extension supplies database metrics, not host, filesystem, or process metrics.
- Installing pgmonitor does not automatically configure Prometheus, exporters, dashboards, or alert rules.
