---
title: "supautils"
linkTitle: "supautils"
description: "Extension that secures a cluster on a cloud environment"
weight: 7010
categories: ["SEC"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**supautils**](https://github.com/supabase/supautils) : Extension that secures a cluster on a cloud environment


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7010** | {{< badge content="supautils" link="https://github.com/supabase/supautils" >}} | {{< ext "supautils" >}} | `3.4.4` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_command_fw" >}} {{< ext "pgextwlist" >}} {{< ext "block_copy_command" >}} {{< ext "pg_kpart" >}} {{< ext "noset" >}} {{< ext "sepgsql" >}} |

> [!Note] Hook library only; no CREATE EXTENSION objects.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `supautils` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.4` | {{< bg "18" "supautils_18" "green" >}} {{< bg "17" "supautils_17" "green" >}} {{< bg "16" "supautils_16" "green" >}} {{< bg "15" "supautils_15" "green" >}} {{< bg "14" "supautils_14" "green" >}} | `supautils_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.4.4` | {{< bg "18" "postgresql-18-supautils" "green" >}} {{< bg "17" "postgresql-17-supautils" "green" >}} {{< bg "16" "postgresql-16-supautils" "green" >}} {{< bg "15" "postgresql-15-supautils" "green" >}} {{< bg "14" "postgresql-14-supautils" "green" >}} | `postgresql-$v-supautils` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.4.4" "postgresql-14-supautils : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_18` | `3.4.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 102.0 KiB | [supautils_18-3.4.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_18-3.4.4-1PGSTY.el8.x86_64.rpm) |
| `supautils_18` | `3.4.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 99.6 KiB | [supautils_18-3.4.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_18-3.4.4-1PGSTY.el8.aarch64.rpm) |
| `supautils_18` | `3.4.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 102.3 KiB | [supautils_18-3.4.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_18-3.4.4-1PGSTY.el9.x86_64.rpm) |
| `supautils_18` | `3.4.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.3 KiB | [supautils_18-3.4.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_18-3.4.4-1PGSTY.el9.aarch64.rpm) |
| `supautils_18` | `3.4.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 103.3 KiB | [supautils_18-3.4.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_18-3.4.4-1PGSTY.el10.x86_64.rpm) |
| `supautils_18` | `3.4.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 101.1 KiB | [supautils_18-3.4.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_18-3.4.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-supautils` | `3.4.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 95.0 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 92.9 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 95.0 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.1 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 101.4 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 100.0 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 99.1 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~noble_amd64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 97.4 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~noble_arm64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.0 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-supautils` | `3.4.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 97.3 KiB | [postgresql-18-supautils_3.4.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-18-supautils_3.4.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_17` | `3.4.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 101.8 KiB | [supautils_17-3.4.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_17-3.4.4-1PGSTY.el8.x86_64.rpm) |
| `supautils_17` | `3.4.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 99.5 KiB | [supautils_17-3.4.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_17-3.4.4-1PGSTY.el8.aarch64.rpm) |
| `supautils_17` | `3.4.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 102.2 KiB | [supautils_17-3.4.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_17-3.4.4-1PGSTY.el9.x86_64.rpm) |
| `supautils_17` | `3.4.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.2 KiB | [supautils_17-3.4.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_17-3.4.4-1PGSTY.el9.aarch64.rpm) |
| `supautils_17` | `3.4.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 103.1 KiB | [supautils_17-3.4.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_17-3.4.4-1PGSTY.el10.x86_64.rpm) |
| `supautils_17` | `3.4.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 101.0 KiB | [supautils_17-3.4.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_17-3.4.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-supautils` | `3.4.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 94.9 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 92.7 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 94.9 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.0 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 127.8 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.8 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 99.0 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~noble_amd64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 97.4 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~noble_arm64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 98.9 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-supautils` | `3.4.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 97.2 KiB | [postgresql-17-supautils_3.4.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-17-supautils_3.4.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_16` | `3.4.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 102.0 KiB | [supautils_16-3.4.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_16-3.4.4-1PGSTY.el8.x86_64.rpm) |
| `supautils_16` | `3.4.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 99.7 KiB | [supautils_16-3.4.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_16-3.4.4-1PGSTY.el8.aarch64.rpm) |
| `supautils_16` | `3.4.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 102.4 KiB | [supautils_16-3.4.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_16-3.4.4-1PGSTY.el9.x86_64.rpm) |
| `supautils_16` | `3.4.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 100.4 KiB | [supautils_16-3.4.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_16-3.4.4-1PGSTY.el9.aarch64.rpm) |
| `supautils_16` | `3.4.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 103.3 KiB | [supautils_16-3.4.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_16-3.4.4-1PGSTY.el10.x86_64.rpm) |
| `supautils_16` | `3.4.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 101.2 KiB | [supautils_16-3.4.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_16-3.4.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-supautils` | `3.4.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 95.0 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 92.8 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 95.1 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 93.0 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.1 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 99.1 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~noble_amd64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 97.5 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~noble_arm64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 99.1 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-supautils` | `3.4.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 97.3 KiB | [postgresql-16-supautils_3.4.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-16-supautils_3.4.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_15` | `3.4.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 103.2 KiB | [supautils_15-3.4.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_15-3.4.4-1PGSTY.el8.x86_64.rpm) |
| `supautils_15` | `3.4.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 100.9 KiB | [supautils_15-3.4.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_15-3.4.4-1PGSTY.el8.aarch64.rpm) |
| `supautils_15` | `3.4.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 104.5 KiB | [supautils_15-3.4.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_15-3.4.4-1PGSTY.el9.x86_64.rpm) |
| `supautils_15` | `3.4.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 102.5 KiB | [supautils_15-3.4.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_15-3.4.4-1PGSTY.el9.aarch64.rpm) |
| `supautils_15` | `3.4.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.1 KiB | [supautils_15-3.4.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_15-3.4.4-1PGSTY.el10.x86_64.rpm) |
| `supautils_15` | `3.4.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 103.2 KiB | [supautils_15-3.4.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_15-3.4.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-supautils` | `3.4.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 96.5 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 94.1 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 96.5 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.5 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 127.5 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.2 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.4 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~noble_amd64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.2 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~noble_arm64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 100.4 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-supautils` | `3.4.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 99.1 KiB | [postgresql-15-supautils_3.4.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-15-supautils_3.4.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_14` | `3.4.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 103.1 KiB | [supautils_14-3.4.4-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_14-3.4.4-1PGSTY.el8.x86_64.rpm) |
| `supautils_14` | `3.4.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 100.8 KiB | [supautils_14-3.4.4-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_14-3.4.4-1PGSTY.el8.aarch64.rpm) |
| `supautils_14` | `3.4.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 104.5 KiB | [supautils_14-3.4.4-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_14-3.4.4-1PGSTY.el9.x86_64.rpm) |
| `supautils_14` | `3.4.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 102.4 KiB | [supautils_14-3.4.4-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_14-3.4.4-1PGSTY.el9.aarch64.rpm) |
| `supautils_14` | `3.4.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.0 KiB | [supautils_14-3.4.4-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_14-3.4.4-1PGSTY.el10.x86_64.rpm) |
| `supautils_14` | `3.4.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 103.3 KiB | [supautils_14-3.4.4-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_14-3.4.4-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-supautils` | `3.4.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 96.5 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 94.0 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 96.5 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.5 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 121.2 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 118.8 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 100.3 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~noble_amd64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 99.1 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~noble_arm64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 100.4 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-supautils` | `3.4.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 99.1 KiB | [postgresql-14-supautils_3.4.4-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/supautils/postgresql-14-supautils_3.4.4-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/supautils" title="Repository" icon="github" subtitle="github.com/supabase/supautils" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="supautils-3.4.4.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg supautils;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install supautils;		# install via package name, for the active PG version

pig install supautils -v 18;   # install for PG 18
pig install supautils -v 17;   # install for PG 17
pig install supautils -v 16;   # install for PG 16
pig install supautils -v 15;   # install for PG 15
pig install supautils -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'supautils';
```


This extension does not need `CREATE EXTENSION` DDL command



## Usage

Sources:

- [v3.4.4 README](https://github.com/supabase/supautils/blob/v3.4.4/README.md)
- [v3.4.4 release](https://github.com/supabase/supautils/releases/tag/v3.4.4)
- [Version restriction implementation](https://github.com/supabase/supautils/blob/v3.4.4/src/extensions.c)

`supautils` is a loadable library that unlocks selected superuser-only PostgreSQL features for non-superusers through configuration. Upstream emphasizes that it adds no tables, functions, or security labels to the database.

### Load it

Cluster-wide:

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

Per role:

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### Privileged role capabilities

The README documents a privileged proxy role that can create publications, foreign data wrappers, event triggers, and privileged extensions without granting `SUPERUSER`.

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

For event triggers, the README says privileged-role triggers run for non-superusers, skip superusers, and also skip reserved roles. It also documents one limitation: those triggers do not fire while creating publications, foreign data wrappers, or extensions.

### Important configuration knobs

- `supautils.superuser`
- `supautils.privileged_role`
- `supautils.privileged_role_allowed_configs`
- `supautils.privileged_extensions`
- `supautils.extension_custom_scripts_path`
- `supautils.constrained_extensions`
- `supautils.extensions_parameter_overrides`
- `supautils.policy_grants`
- `supautils.drop_trigger_grants`
- `supautils.reserved_roles`
- `supautils.reserved_memberships`
- `supautils.hint_roles`
- `supautils.log_skipped_evtrigs`

### Useful examples

Allow a non-superuser to create specific privileged extensions:

```ini
supautils.privileged_extensions = 'hstore'
```

Allow a role to manage RLS policies on tables it does not own:

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

Force an extension into a specific schema on `CREATE EXTENSION`:

```ini
supautils.extensions_parameter_overrides = '{ "pg_cron": { "schema": "pg_catalog" } }'
```

Protect managed-service roles from `CREATEROLE` users:

```ini
supautils.reserved_roles = 'connector, storage_admin'
supautils.reserved_memberships = 'pg_read_server_files'
```

### Version Selection and Operational Boundaries

`supautils.restrict_extension_versions` controls explicit version clauses for non-superusers: `off` allows them, `warn` ignores them and selects the control-file default with a warning, and `error` rejects them. This applies to both extension creation and upgrades; superusers and the configured proxy superuser are exempt. Omitting an explicit version remains allowed subject to normal privilege checks.

Cluster preload requires a restart; role-specific session preload applies to new connections. Do not run CREATE EXTENSION for supautils itself. Source release 3.4.4 is a library update and has no SQL extension-update step. It avoids ACCESS EXCLUSIVE locks during allowlisted-table policy checks and restores the caller's role on every exit from an elevated region.

Review allowed extensions and custom scripts as trusted code because their operations run with delegated superuser privileges. Enhanced privilege hints do not work for views on PostgreSQL 18 according to the tagged README. Test role transitions, event-trigger ownership and reserved-role protections before broadening grants.
