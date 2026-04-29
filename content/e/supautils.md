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
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-supautils : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-supautils : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

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
{{< tab name="PG17" >}}

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
{{< tab name="PG16" >}}

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
{{< tab name="PG15" >}}

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
{{< tab name="PG14" >}}

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

Sources: [README](https://github.com/supabase/supautils/blob/master/README.md), [homepage](https://supabase.github.io/supautils/), [releases](https://github.com/supabase/supautils/releases)

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

### Release notes

- `v3.2.1` was released on 2026-04-02 and its published notes are maintenance-oriented; no new user-facing SQL surface is described there.
- `v3.2.0` added a hint when a `GRANT` privilege is missing.

### Caveat

This extension is configuration-driven. When documenting it, prefer the GUCs and behavior guarantees in the README over implying database objects that upstream explicitly says it does not create.
