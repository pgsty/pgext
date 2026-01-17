---
title: "wrappers"
linkTitle: "wrappers"
description: "Foreign data wrappers developed by Supabase"
weight: 8500
categories: ["FDW"]
width: full
---

[**wrappers**](https://github.com/supabase/wrappers) : Foreign data wrappers developed by Supabase


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8500** | {{< badge content="wrappers" link="https://github.com/supabase/wrappers" >}} | {{< ext "wrappers" >}} | `0.5.7` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} |

> [!Note] manual updated pgrx by Vonng


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `wrappers` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.7` | {{< bg "18" "wrappers_18" "green" >}} {{< bg "17" "wrappers_17" "green" >}} {{< bg "16" "wrappers_16" "green" >}} {{< bg "15" "wrappers_15" "green" >}} {{< bg "14" "wrappers_14" "green" >}} {{< bg "13" "wrappers_13" "red" >}} | `wrappers_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.5.7` | {{< bg "18" "postgresql-18-wrappers" "green" >}} {{< bg "17" "postgresql-17-wrappers" "green" >}} {{< bg "16" "postgresql-16-wrappers" "green" >}} {{< bg "15" "postgresql-15-wrappers" "green" >}} {{< bg "14" "postgresql-14-wrappers" "green" >}} {{< bg "13" "postgresql-13-wrappers" "red" >}} | `postgresql-$v-wrappers` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "wrappers_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "wrappers_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.5.7" "postgresql-14-wrappers : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-wrappers : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_18` | `0.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 240.1 KiB | [wrappers_18-0.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_18-0.5.7-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_18` | `0.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.9 KiB | [wrappers_18-0.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_18-0.5.7-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_18` | `0.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 251.0 KiB | [wrappers_18-0.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_18-0.5.7-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_18` | `0.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 166.4 KiB | [wrappers_18-0.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_18-0.5.7-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_18` | `0.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 250.6 KiB | [wrappers_18-0.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_18-0.5.7-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_18` | `0.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 166.6 KiB | [wrappers_18-0.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_18-0.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-wrappers` | `0.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 200.1 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.3 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 200.1 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 122.3 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 224.0 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.3 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.8 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-wrappers` | `0.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.9 KiB | [postgresql-18-wrappers_0.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.5.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_17` | `0.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 240.0 KiB | [wrappers_17-0.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_17-0.5.7-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_17` | `0.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.8 KiB | [wrappers_17-0.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_17-0.5.7-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_17` | `0.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 250.9 KiB | [wrappers_17-0.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_17-0.5.7-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_17` | `0.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 166.4 KiB | [wrappers_17-0.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_17-0.5.7-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_17` | `0.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 250.8 KiB | [wrappers_17-0.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_17-0.5.7-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_17` | `0.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 166.6 KiB | [wrappers_17-0.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_17-0.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-wrappers` | `0.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 200.2 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.4 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 200.3 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 122.3 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 223.9 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.3 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.8 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-wrappers` | `0.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 140.0 KiB | [postgresql-17-wrappers_0.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.5.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_16` | `0.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 240.0 KiB | [wrappers_16-0.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_16-0.5.7-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_16` | `0.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.8 KiB | [wrappers_16-0.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_16-0.5.7-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_16` | `0.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 250.9 KiB | [wrappers_16-0.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_16-0.5.7-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_16` | `0.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 166.4 KiB | [wrappers_16-0.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_16-0.5.7-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_16` | `0.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 250.8 KiB | [wrappers_16-0.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_16-0.5.7-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_16` | `0.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 166.6 KiB | [wrappers_16-0.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_16-0.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-wrappers` | `0.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 200.0 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.4 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 200.3 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 122.3 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 223.8 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.3 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.8 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-wrappers` | `0.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 140.0 KiB | [postgresql-16-wrappers_0.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.5.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_15` | `0.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.8 KiB | [wrappers_15-0.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_15-0.5.7-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_15` | `0.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.8 KiB | [wrappers_15-0.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_15-0.5.7-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_15` | `0.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 250.4 KiB | [wrappers_15-0.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_15-0.5.7-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_15` | `0.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 166.4 KiB | [wrappers_15-0.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_15-0.5.7-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_15` | `0.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 250.4 KiB | [wrappers_15-0.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_15-0.5.7-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_15` | `0.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 166.6 KiB | [wrappers_15-0.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_15-0.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-wrappers` | `0.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 200.0 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.3 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 200.0 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 122.3 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 223.7 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.3 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.6 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-wrappers` | `0.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.8 KiB | [postgresql-15-wrappers_0.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.5.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_14` | `0.5.7` | [el8.x86_64](/os/el8.x86_64) | pigsty | 239.8 KiB | [wrappers_14-0.5.7-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_14-0.5.7-1PIGSTY.el8.x86_64.rpm) |
| `wrappers_14` | `0.5.7` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.9 KiB | [wrappers_14-0.5.7-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_14-0.5.7-1PIGSTY.el8.aarch64.rpm) |
| `wrappers_14` | `0.5.7` | [el9.x86_64](/os/el9.x86_64) | pigsty | 250.3 KiB | [wrappers_14-0.5.7-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_14-0.5.7-1PIGSTY.el9.x86_64.rpm) |
| `wrappers_14` | `0.5.7` | [el9.aarch64](/os/el9.aarch64) | pigsty | 166.4 KiB | [wrappers_14-0.5.7-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_14-0.5.7-1PIGSTY.el9.aarch64.rpm) |
| `wrappers_14` | `0.5.7` | [el10.x86_64](/os/el10.x86_64) | pigsty | 250.4 KiB | [wrappers_14-0.5.7-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_14-0.5.7-1PIGSTY.el10.x86_64.rpm) |
| `wrappers_14` | `0.5.7` | [el10.aarch64](/os/el10.aarch64) | pigsty | 166.6 KiB | [wrappers_14-0.5.7-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_14-0.5.7-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-wrappers` | `0.5.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 200.1 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 122.3 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 200.1 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 122.3 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 223.7 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 141.1 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 221.6 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-wrappers` | `0.5.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.9 KiB | [postgresql-14-wrappers_0.5.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.5.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/wrappers" title="Repository" icon="github" subtitle="github.com/supabase/wrappers" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="wrappers-0.5.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg wrappers;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install wrappers;		# install via package name, for the active PG version

pig install wrappers -v 18;   # install for PG 18
pig install wrappers -v 17;   # install for PG 17
pig install wrappers -v 16;   # install for PG 16
pig install wrappers -v 15;   # install for PG 15
pig install wrappers -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION wrappers;
```
