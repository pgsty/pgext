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
| **7010** | {{< badge content="supautils" link="https://github.com/supabase/supautils" >}} | {{< ext "supautils" >}} | `3.0.2` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "pg_session_jwt" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `supautils` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.2` | {{< bg "18" "supautils_18" "green" >}} {{< bg "17" "supautils_17" "green" >}} {{< bg "16" "supautils_16" "green" >}} {{< bg "15" "supautils_15" "green" >}} {{< bg "14" "supautils_14" "green" >}} {{< bg "13" "supautils_13" "green" >}} | `supautils_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.2` | {{< bg "18" "postgresql-18-supautils" "green" >}} {{< bg "17" "postgresql-17-supautils" "green" >}} {{< bg "16" "postgresql-16-supautils" "green" >}} {{< bg "15" "postgresql-15-supautils" "green" >}} {{< bg "14" "postgresql-14-supautils" "green" >}} {{< bg "13" "postgresql-13-supautils" "green" >}} | `postgresql-$v-supautils` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "supautils_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-18-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.0.2" "postgresql-13-supautils : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_18` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [supautils_18-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_18-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_18` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.3 KiB | [supautils_18-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_18-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_18` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.1 KiB | [supautils_18-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_18-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_18` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.6 KiB | [supautils_18-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_18-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_18` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.5 KiB | [supautils_18-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_18-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_18` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.3 KiB | [supautils_18-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_18-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.6 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.1 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.7 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.8 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.2 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.6 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.2 KiB | [postgresql-18-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-18-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_17` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.7 KiB | [supautils_17-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_17-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_17` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.3 KiB | [supautils_17-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_17-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_17` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.1 KiB | [supautils_17-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_17-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_17` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [supautils_17-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_17-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_17` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.6 KiB | [supautils_17-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_17-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_17` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.3 KiB | [supautils_17-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_17-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.6 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.1 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.7 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.3 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.9 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.3 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.7 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.2 KiB | [postgresql-17-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_16` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.8 KiB | [supautils_16-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_16-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_16` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.4 KiB | [supautils_16-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_16-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_16` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.2 KiB | [supautils_16-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_16-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_16` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.7 KiB | [supautils_16-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_16-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_16` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.6 KiB | [supautils_16-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_16-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_16` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.4 KiB | [supautils_16-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_16-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 22.7 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.1 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 22.7 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.4 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 23.9 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 23.3 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.7 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.3 KiB | [postgresql-16-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_15` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [supautils_15-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_15-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_15` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.8 KiB | [supautils_15-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_15-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_15` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.6 KiB | [supautils_15-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_15-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_15` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.9 KiB | [supautils_15-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_15-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_15` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.6 KiB | [supautils_15-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_15-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_15` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.2 KiB | [supautils_15-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_15-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.6 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.8 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.9 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.2 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.2 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.3 KiB | [postgresql-15-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_14` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 30.3 KiB | [supautils_14-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_14-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_14` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.8 KiB | [supautils_14-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_14-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_14` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.6 KiB | [supautils_14-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_14-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_14` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.8 KiB | [supautils_14-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_14-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_14` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.6 KiB | [supautils_14-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_14-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_14` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.3 KiB | [supautils_14-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_14-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.2 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.6 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.8 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.9 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.2 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.2 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.3 KiB | [postgresql-14-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_13` | `3.0.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.6 KiB | [supautils_13-3.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_13-3.0.2-1PIGSTY.el8.x86_64.rpm) |
| `supautils_13` | `3.0.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 29.7 KiB | [supautils_13-3.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_13-3.0.2-1PIGSTY.el8.aarch64.rpm) |
| `supautils_13` | `3.0.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 29.5 KiB | [supautils_13-3.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_13-3.0.2-1PIGSTY.el9.x86_64.rpm) |
| `supautils_13` | `3.0.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 28.6 KiB | [supautils_13-3.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_13-3.0.2-1PIGSTY.el9.aarch64.rpm) |
| `supautils_13` | `3.0.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 29.4 KiB | [supautils_13-3.0.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/supautils_13-3.0.2-1PIGSTY.el10.x86_64.rpm) |
| `supautils_13` | `3.0.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 29.1 KiB | [supautils_13-3.0.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/supautils_13-3.0.2-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-supautils` | `3.0.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.1 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.4 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 23.2 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 22.7 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 24.8 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.0 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 24.2 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-supautils` | `3.0.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 24.1 KiB | [postgresql-13-supautils_3.0.2-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-13-supautils_3.0.2-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/supautils" title="Repository" icon="github" subtitle="github.com/supabase/supautils" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="supautils-3.0.2.tar.gz" >}}
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
pig install supautils -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'supautils';
```


This extension does not need `CREATE EXTENSION` DDL command


