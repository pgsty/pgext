---
title: "wrappers"
linkTitle: "wrappers"
description: "Foreign data wrappers developed by Supabase"
weight: 8500
categories: ["FDW"]
languages: ["Rust"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**wrappers**](https://github.com/supabase/wrappers) : Foreign data wrappers developed by Supabase


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8500** | {{< badge content="wrappers" link="https://github.com/supabase/wrappers" >}} | {{< ext "wrappers" >}} | `0.6.3` | {{< category "FDW" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "odbc_fdw" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `wrappers` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.3` | {{< bg "18" "wrappers_18" "green" >}} {{< bg "17" "wrappers_17" "green" >}} {{< bg "16" "wrappers_16" "green" >}} {{< bg "15" "wrappers_15" "green" >}} {{< bg "14" "wrappers_14" "green" >}} | `wrappers_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.3` | {{< bg "18" "postgresql-18-wrappers" "green" >}} {{< bg "17" "postgresql-17-wrappers" "green" >}} {{< bg "16" "postgresql-16-wrappers" "green" >}} {{< bg "15" "postgresql-15-wrappers" "green" >}} {{< bg "14" "postgresql-14-wrappers" "green" >}} | `postgresql-$v-wrappers` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "wrappers_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-18-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-17-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-16-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-15-wrappers : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.3" "postgresql-14-wrappers : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_18` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.3 MiB | [wrappers_18-0.6.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_18-0.6.3-1PGSTY.el8.x86_64.rpm) |
| `wrappers_18` | `0.6.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.4 MiB | [wrappers_18-0.6.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_18-0.6.3-1PGSTY.el8.aarch64.rpm) |
| `wrappers_18` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.5 MiB | [wrappers_18-0.6.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_18-0.6.3-1PGSTY.el9.x86_64.rpm) |
| `wrappers_18` | `0.6.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 MiB | [wrappers_18-0.6.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_18-0.6.3-1PGSTY.el9.aarch64.rpm) |
| `wrappers_18` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.3 MiB | [wrappers_18-0.6.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_18-0.6.3-1PGSTY.el10.x86_64.rpm) |
| `wrappers_18` | `0.6.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 MiB | [wrappers_18-0.6.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_18-0.6.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-wrappers` | `0.6.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.9 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.8 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.9 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-wrappers` | `0.6.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 MiB | [postgresql-18-wrappers_0.6.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-18-wrappers_0.6.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_17` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.3 MiB | [wrappers_17-0.6.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_17-0.6.3-1PGSTY.el8.x86_64.rpm) |
| `wrappers_17` | `0.6.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.4 MiB | [wrappers_17-0.6.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_17-0.6.3-1PGSTY.el8.aarch64.rpm) |
| `wrappers_17` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.5 MiB | [wrappers_17-0.6.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_17-0.6.3-1PGSTY.el9.x86_64.rpm) |
| `wrappers_17` | `0.6.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 MiB | [wrappers_17-0.6.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_17-0.6.3-1PGSTY.el9.aarch64.rpm) |
| `wrappers_17` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.3 MiB | [wrappers_17-0.6.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_17-0.6.3-1PGSTY.el10.x86_64.rpm) |
| `wrappers_17` | `0.6.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 MiB | [wrappers_17-0.6.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_17-0.6.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-wrappers` | `0.6.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.9 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.9 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.9 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-wrappers` | `0.6.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 MiB | [postgresql-17-wrappers_0.6.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-17-wrappers_0.6.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_16` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.3 MiB | [wrappers_16-0.6.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_16-0.6.3-1PGSTY.el8.x86_64.rpm) |
| `wrappers_16` | `0.6.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.4 MiB | [wrappers_16-0.6.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_16-0.6.3-1PGSTY.el8.aarch64.rpm) |
| `wrappers_16` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.5 MiB | [wrappers_16-0.6.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_16-0.6.3-1PGSTY.el9.x86_64.rpm) |
| `wrappers_16` | `0.6.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 MiB | [wrappers_16-0.6.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_16-0.6.3-1PGSTY.el9.aarch64.rpm) |
| `wrappers_16` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.3 MiB | [wrappers_16-0.6.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_16-0.6.3-1PGSTY.el10.x86_64.rpm) |
| `wrappers_16` | `0.6.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 MiB | [wrappers_16-0.6.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_16-0.6.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-wrappers` | `0.6.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.9 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.8 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.7 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-wrappers` | `0.6.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 MiB | [postgresql-16-wrappers_0.6.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-16-wrappers_0.6.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_15` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.3 MiB | [wrappers_15-0.6.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_15-0.6.3-1PGSTY.el8.x86_64.rpm) |
| `wrappers_15` | `0.6.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.4 MiB | [wrappers_15-0.6.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_15-0.6.3-1PGSTY.el8.aarch64.rpm) |
| `wrappers_15` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.5 MiB | [wrappers_15-0.6.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_15-0.6.3-1PGSTY.el9.x86_64.rpm) |
| `wrappers_15` | `0.6.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 MiB | [wrappers_15-0.6.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_15-0.6.3-1PGSTY.el9.aarch64.rpm) |
| `wrappers_15` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.3 MiB | [wrappers_15-0.6.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_15-0.6.3-1PGSTY.el10.x86_64.rpm) |
| `wrappers_15` | `0.6.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 MiB | [wrappers_15-0.6.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_15-0.6.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-wrappers` | `0.6.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 23.9 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.9 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.6 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-wrappers` | `0.6.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 MiB | [postgresql-15-wrappers_0.6.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-15-wrappers_0.6.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wrappers_14` | `0.6.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 31.3 MiB | [wrappers_14-0.6.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/wrappers_14-0.6.3-1PGSTY.el8.x86_64.rpm) |
| `wrappers_14` | `0.6.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 30.4 MiB | [wrappers_14-0.6.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/wrappers_14-0.6.3-1PGSTY.el8.aarch64.rpm) |
| `wrappers_14` | `0.6.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.5 MiB | [wrappers_14-0.6.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/wrappers_14-0.6.3-1PGSTY.el9.x86_64.rpm) |
| `wrappers_14` | `0.6.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 30.0 MiB | [wrappers_14-0.6.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/wrappers_14-0.6.3-1PGSTY.el9.aarch64.rpm) |
| `wrappers_14` | `0.6.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 33.3 MiB | [wrappers_14-0.6.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/wrappers_14-0.6.3-1PGSTY.el10.x86_64.rpm) |
| `wrappers_14` | `0.6.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 30.5 MiB | [wrappers_14-0.6.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/wrappers_14-0.6.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-wrappers` | `0.6.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.4 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.0 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.0 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 29.8 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.6 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.8 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 30.4 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-wrappers` | `0.6.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 26.9 MiB | [postgresql-14-wrappers_0.6.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/w/wrappers/postgresql-14-wrappers_0.6.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/wrappers" title="Repository" icon="github" subtitle="github.com/supabase/wrappers" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="wrappers-0.6.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg wrappers;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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




## Usage

Sources:

- [Wrappers v0.6.2 README](https://github.com/supabase/wrappers/blob/v0.6.2/README.md)
- [Official FDW documentation](https://fdw.dev/)
- [v0.6.2 release](https://github.com/supabase/wrappers/releases/tag/v0.6.2)
- [MongoDB FDW documentation](https://fdw.dev/catalog/mongodb/)
- [Security guidance](https://fdw.dev/guides/security/)

`wrappers` is both a Rust framework for writing PostgreSQL foreign data wrappers and a packaged collection of Supabase-maintained FDWs. A single extension installs many wrapper implementations, then each foreign server chooses the specific wrapper type it needs.

```sql
CREATE EXTENSION wrappers;
```

### Typical Workflow

Create a server for one wrapper, then expose remote data through foreign tables:

```sql
CREATE SERVER stripe_server
  FOREIGN DATA WRAPPER stripe_wrapper
  OPTIONS (
    api_key_id 'stripe_api_key',
    api_url 'https://api.stripe.com/v1/'
  );

CREATE FOREIGN TABLE stripe_customers (
  id text,
  email text,
  name text,
  description text,
  created timestamp,
  attrs jsonb
)
  SERVER stripe_server
  OPTIONS (
    object 'customers',
    rowid_column 'id'
  );
```

### What It Covers

Upstream ships wrappers for databases and services such as BigQuery, ClickHouse, DuckDB, DynamoDB, MySQL/Doris, Redis, S3, S3 Vectors, Stripe, Snowflake, Slack, Notion, OpenAPI, Infura, and many others. Read and write support varies by wrapper, but pushdown for `WHERE`, `ORDER BY`, and `LIMIT` is a core framework feature.

### Version 0.6.2

The `v0.6.2` release keeps the same extension model and adds:

- a MongoDB FDW with read and write support
- session-variable credentials for per-request authentication in WASM wrappers
- RFC 8288 `Link` header pagination for the OpenAPI FDW
- runtime, dependency, and wrapper-specific fixes documented in the release notes

Wrapper-specific pages remain the authority for server options, foreign-table columns, pushdown, and write support.

### Caveats

- Wrapper-specific options, supported objects, and write support differ widely; check the official catalog page for the exact FDW you use.
- The docs warn that logical restores can fail when materialized views depend on foreign tables, so avoid that pattern or rely on physical backups.
- Foreign tables do not provide a security boundary by themselves. Keep them in private schemas, grant access deliberately, use the least-privileged remote credentials available, and apply Row Level Security to local tables that expose or cache remote data.
- Keep API keys and tokens in the supported secret store or per-request credential mechanism instead of embedding them in SQL checked into source control.
