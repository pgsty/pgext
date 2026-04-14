---
title: "supautils"
linkTitle: "supautils"
description: "Extension that secures a cluster on a cloud environment"
weight: 7010
categories: ["SEC"]
width: full
---

[**supautils**](https://github.com/supabase/supautils) : Extension that secures a cluster on a cloud environment


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7010** | {{< badge content="supautils" link="https://github.com/supabase/supautils" >}} | {{< ext "supautils" >}} | `3.2.1` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "pg_session_jwt" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `supautils` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "supautils_18" "green" >}} {{< bg "17" "supautils_17" "green" >}} {{< bg "16" "supautils_16" "green" >}} {{< bg "15" "supautils_15" "green" >}} {{< bg "14" "supautils_14" "green" >}} | `supautils_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.2.1` | {{< bg "18" "postgresql-18-supautils" "green" >}} {{< bg "17" "postgresql-17-supautils" "green" >}} {{< bg "16" "postgresql-16-supautils" "green" >}} {{< bg "15" "postgresql-15-supautils" "green" >}} {{< bg "14" "postgresql-14-supautils" "green" >}} | `postgresql-$v-supautils` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "supautils_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.2.1" "postgresql-14-supautils : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_18` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.3 KiB | [supautils_18-3.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_18-3.2.1-1PIGSTY.el8.x86_64.rpm) |
| `supautils_18` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.4 KiB | [supautils_18-3.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_18-3.2.1-1PIGSTY.el8.aarch64.rpm) |
| `supautils_18` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.5 KiB | [supautils_18-3.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_18-3.2.1-1PIGSTY.el9.x86_64.rpm) |
| `supautils_18` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [supautils_18-3.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_18-3.2.1-1PIGSTY.el9.aarch64.rpm) |
| `supautils_18` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [supautils_18-3.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_18-3.2.1-1PIGSTY.el10.x86_64.rpm) |
| `supautils_18` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.8 KiB | [supautils_18-3.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_18-3.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-supautils` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.3 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.1 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.5 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.7 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.4 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.2 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-supautils` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.2 KiB | [postgresql-18-supautils_3.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_17` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.3 KiB | [supautils_17-3.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_17-3.2.1-1PIGSTY.el8.x86_64.rpm) |
| `supautils_17` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.4 KiB | [supautils_17-3.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_17-3.2.1-1PIGSTY.el8.aarch64.rpm) |
| `supautils_17` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.5 KiB | [supautils_17-3.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_17-3.2.1-1PIGSTY.el9.x86_64.rpm) |
| `supautils_17` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [supautils_17-3.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_17-3.2.1-1PIGSTY.el9.aarch64.rpm) |
| `supautils_17` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 31.0 KiB | [supautils_17-3.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_17-3.2.1-1PIGSTY.el10.x86_64.rpm) |
| `supautils_17` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.9 KiB | [supautils_17-3.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_17-3.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-supautils` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.3 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.1 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.4 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.7 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.4 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.2 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-supautils` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.3 KiB | [postgresql-17-supautils_3.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_16` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 32.5 KiB | [supautils_16-3.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_16-3.2.1-1PIGSTY.el8.x86_64.rpm) |
| `supautils_16` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.4 KiB | [supautils_16-3.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_16-3.2.1-1PIGSTY.el8.aarch64.rpm) |
| `supautils_16` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 30.5 KiB | [supautils_16-3.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_16-3.2.1-1PIGSTY.el9.x86_64.rpm) |
| `supautils_16` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 29.3 KiB | [supautils_16-3.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_16-3.2.1-1PIGSTY.el9.aarch64.rpm) |
| `supautils_16` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 30.9 KiB | [supautils_16-3.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_16-3.2.1-1PIGSTY.el10.x86_64.rpm) |
| `supautils_16` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.8 KiB | [supautils_16-3.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_16-3.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-supautils` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.3 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.1 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.5 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.4 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.7 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 25.4 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.2 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-supautils` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.3 KiB | [postgresql-16-supautils_3.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_15` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.1 KiB | [supautils_15-3.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_15-3.2.1-1PIGSTY.el8.x86_64.rpm) |
| `supautils_15` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 32.0 KiB | [supautils_15-3.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_15-3.2.1-1PIGSTY.el8.aarch64.rpm) |
| `supautils_15` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.3 KiB | [supautils_15-3.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_15-3.2.1-1PIGSTY.el9.x86_64.rpm) |
| `supautils_15` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [supautils_15-3.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_15-3.2.1-1PIGSTY.el9.aarch64.rpm) |
| `supautils_15` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.8 KiB | [supautils_15-3.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_15-3.2.1-1PIGSTY.el10.x86_64.rpm) |
| `supautils_15` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.7 KiB | [supautils_15-3.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_15-3.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-supautils` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.0 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.8 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.2 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.1 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.8 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.0 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.3 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-supautils` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.9 KiB | [postgresql-15-supautils_3.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_14` | `3.2.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 33.1 KiB | [supautils_14-3.2.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_14-3.2.1-1PIGSTY.el8.x86_64.rpm) |
| `supautils_14` | `3.2.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 31.9 KiB | [supautils_14-3.2.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_14-3.2.1-1PIGSTY.el8.aarch64.rpm) |
| `supautils_14` | `3.2.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 32.2 KiB | [supautils_14-3.2.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_14-3.2.1-1PIGSTY.el9.x86_64.rpm) |
| `supautils_14` | `3.2.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 31.1 KiB | [supautils_14-3.2.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_14-3.2.1-1PIGSTY.el9.aarch64.rpm) |
| `supautils_14` | `3.2.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 32.7 KiB | [supautils_14-3.2.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_14-3.2.1-1PIGSTY.el10.x86_64.rpm) |
| `supautils_14` | `3.2.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 31.8 KiB | [supautils_14-3.2.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_14-3.2.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-supautils` | `3.2.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 26.0 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.7 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 26.2 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.1 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.8 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.9 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 27.3 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-supautils` | `3.2.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.9 KiB | [postgresql-14-supautils_3.2.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.2.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/supautils" title="Repository" icon="github" subtitle="github.com/supabase/supautils" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="supautils-3.2.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg supautils;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

> [supautils: Extension that secures a cluster on a cloud environment](https://github.com/supabase/supautils)

`supautils` is a loadable library that securely allows creating event triggers, publications, and extensions for non-superusers. It is completely managed by configuration -- no tables, functions, or security labels are added to your database.

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'supautils'
supautils.privileged_role = 'your_privileged_role'
```

Or enable per-role:

```sql
ALTER ROLE role1 SET session_preload_libraries TO 'supautils';
```

### Key GUC Parameters

| Parameter | Description |
|-----------|-------------|
| `supautils.privileged_role` | Proxy role for superuser operations |
| `supautils.superuser` | The actual superuser (defaults to bootstrap user) |
| `supautils.privileged_extensions` | Extensions allowed for non-superuser installation |
| `supautils.privileged_role_allowed_configs` | Superuser-only settings the privileged role may change |
| `supautils.reserved_roles` | Roles protected from mutation by CREATEROLE users |
| `supautils.reserved_memberships` | Role memberships restricted from being granted |
| `supautils.constrained_extensions` | JSON defining resource constraints for extensions |
| `supautils.extensions_parameter_overrides` | JSON overriding CREATE EXTENSION parameters |
| `supautils.policy_grants` | JSON granting RLS policy management to non-owners |
| `supautils.drop_trigger_grants` | JSON granting trigger drop permission to non-owners |

### Non-Superuser Publications

```sql
SET ROLE privileged_role;
CREATE PUBLICATION p FOR ALL TABLES;
DROP PUBLICATION p;
```

### Privileged Extensions

```ini
supautils.privileged_extensions = 'hstore'
```

Non-superusers can then create extensions that normally require superuser:

```sql
CREATE EXTENSION hstore;
```

### Reserved Roles

```ini
supautils.reserved_roles = 'connector, storage_admin'
```

Users with CREATEROLE cannot ALTER or DROP these roles.

### Table Ownership Bypass (RLS Policy Management)

```ini
supautils.policy_grants = '{ "my_role": ["public.not_my_table"] }'
```

Allows `my_role` to manage RLS policies on tables it does not own.
