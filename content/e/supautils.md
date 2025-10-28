---
title: "supautils"
linkTitle: "supautils"
description: "Extension that secures a cluster on a cloud environment"
weight: 7010
categories: ["SEC"]
width: full
---

Extension that secures a cluster on a cloud environment


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7010** | {{< badge content="supautils" link="https://github.com/supabase/supautils" >}} | {{< ext "supautils" >}} | `3.0.1` | {{< category "SEC" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sL---" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="No" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "passwordcheck_cracklib" >}} {{< ext "pgsodium" >}} {{< ext "supabase_vault" >}} {{< ext "pg_session_jwt" >}} {{< ext "anon" >}} {{< ext "pg_tde" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgaudit" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/supautils" >}} | `3.0.1` | {{< bg "18" "supautils_18" "green" >}} {{< bg "17" "supautils_17" "green" >}} {{< bg "16" "supautils_16" "green" >}} {{< bg "15" "supautils_15" "green" >}} {{< bg "14" "supautils_14" "green" >}} {{< bg "13" "supautils_13" "green" >}} | `supautils_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/supautils" >}} | `3.0.1` | {{< bg "18" "postgresql-18-supautils" "green" >}} {{< bg "17" "postgresql-17-supautils" "green" >}} {{< bg "16" "postgresql-16-supautils" "green" >}} {{< bg "15" "postgresql-15-supautils" "green" >}} {{< bg "14" "postgresql-14-supautils" "green" >}} {{< bg "13" "postgresql-13-supautils" "green" >}} | `postgresql-$v-supautils` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 2.10.0" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_13 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PIGSTY 2.10.0" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_13 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PIGSTY 2.10.0" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_13 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PIGSTY 2.10.0" "supautils_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "supautils_13 : AVAIL 1" "green" >}} |
|    `el10.x86_64`    |      {{< bg "MISS" "supautils_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "supautils_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "supautils_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-supautils : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-supautils : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-supautils : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-supautils : MISS 0" "red" >}}      | {{< bg "PIGSTY 2.10.0" "postgresql-17-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-16-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-15-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-14-supautils : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.10.0" "postgresql-13-supautils : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_18` | 2.10.0 | `el8.x86_64` | pigsty | 29.1 KiB | [supautils_18-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_18-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_18` | 2.10.0 | `el8.aarch64` | pigsty | 28.7 KiB | [supautils_18-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_18-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_18` | 2.10.0 | `el9.x86_64` | pigsty | 27.7 KiB | [supautils_18-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_18-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_18` | 2.10.0 | `el9.aarch64` | pigsty | 27.2 KiB | [supautils_18-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_18-2.10.0-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_17` | 2.10.0 | `el8.x86_64` | pigsty | 29.1 KiB | [supautils_17-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_17-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_17` | 2.10.0 | `el8.aarch64` | pigsty | 28.7 KiB | [supautils_17-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_17-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_17` | 2.10.0 | `el9.x86_64` | pigsty | 27.8 KiB | [supautils_17-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_17-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_17` | 2.10.0 | `el9.aarch64` | pigsty | 27.3 KiB | [supautils_17-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_17-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-supautils` | 2.10.0 | `d12.x86_64` | pigsty | 22.2 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-supautils` | 2.10.0 | `d12.aarch64` | pigsty | 21.7 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-supautils` | 2.10.0 | `u22.x86_64` | pigsty | 23.5 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-supautils` | 2.10.0 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-supautils` | 2.10.0 | `u24.x86_64` | pigsty | 23.3 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-supautils` | 2.10.0 | `u24.aarch64` | pigsty | 22.9 KiB | [postgresql-17-supautils_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-17-supautils_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_16` | 2.10.0 | `el8.x86_64` | pigsty | 29.2 KiB | [supautils_16-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_16-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_16` | 2.10.0 | `el8.aarch64` | pigsty | 28.7 KiB | [supautils_16-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_16-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_16` | 2.10.0 | `el9.x86_64` | pigsty | 27.8 KiB | [supautils_16-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_16-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_16` | 2.10.0 | `el9.aarch64` | pigsty | 27.4 KiB | [supautils_16-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_16-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-supautils` | 2.10.0 | `d12.x86_64` | pigsty | 22.3 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-supautils` | 2.10.0 | `d12.aarch64` | pigsty | 21.8 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-supautils` | 2.10.0 | `u22.x86_64` | pigsty | 23.5 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-supautils` | 2.10.0 | `u22.aarch64` | pigsty | 22.8 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-supautils` | 2.10.0 | `u24.x86_64` | pigsty | 23.3 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-supautils` | 2.10.0 | `u24.aarch64` | pigsty | 22.9 KiB | [postgresql-16-supautils_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-16-supautils_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_15` | 2.10.0 | `el8.x86_64` | pigsty | 29.8 KiB | [supautils_15-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_15-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_15` | 2.10.0 | `el8.aarch64` | pigsty | 29.2 KiB | [supautils_15-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_15-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_15` | 2.10.0 | `el9.x86_64` | pigsty | 29.2 KiB | [supautils_15-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_15-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_15` | 2.10.0 | `el9.aarch64` | pigsty | 28.5 KiB | [supautils_15-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_15-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-supautils` | 2.10.0 | `d12.x86_64` | pigsty | 22.8 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-supautils` | 2.10.0 | `d12.aarch64` | pigsty | 22.3 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-supautils` | 2.10.0 | `u22.x86_64` | pigsty | 24.4 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-supautils` | 2.10.0 | `u22.aarch64` | pigsty | 23.7 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-supautils` | 2.10.0 | `u24.x86_64` | pigsty | 23.8 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-supautils` | 2.10.0 | `u24.aarch64` | pigsty | 23.9 KiB | [postgresql-15-supautils_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-15-supautils_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_14` | 2.10.0 | `el8.x86_64` | pigsty | 29.8 KiB | [supautils_14-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_14-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_14` | 2.10.0 | `el8.aarch64` | pigsty | 29.2 KiB | [supautils_14-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_14-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_14` | 2.10.0 | `el9.x86_64` | pigsty | 29.2 KiB | [supautils_14-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_14-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_14` | 2.10.0 | `el9.aarch64` | pigsty | 28.5 KiB | [supautils_14-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_14-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-supautils` | 2.10.0 | `d12.x86_64` | pigsty | 22.8 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-supautils` | 2.10.0 | `d12.aarch64` | pigsty | 22.3 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-supautils` | 2.10.0 | `u22.x86_64` | pigsty | 24.4 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-supautils` | 2.10.0 | `u22.aarch64` | pigsty | 23.7 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-supautils` | 2.10.0 | `u24.x86_64` | pigsty | 23.8 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-supautils` | 2.10.0 | `u24.aarch64` | pigsty | 23.9 KiB | [postgresql-14-supautils_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-14-supautils_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `supautils_13` | 2.10.0 | `el8.x86_64` | pigsty | 29.2 KiB | [supautils_13-2.10.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/supautils_13-2.10.0-1PIGSTY.el8.x86_64.rpm) |
| `supautils_13` | 2.10.0 | `el8.aarch64` | pigsty | 29.2 KiB | [supautils_13-2.10.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/supautils_13-2.10.0-1PIGSTY.el8.aarch64.rpm) |
| `supautils_13` | 2.10.0 | `el9.x86_64` | pigsty | 28.8 KiB | [supautils_13-2.10.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/supautils_13-2.10.0-1PIGSTY.el9.x86_64.rpm) |
| `supautils_13` | 2.10.0 | `el9.aarch64` | pigsty | 28.3 KiB | [supautils_13-2.10.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/supautils_13-2.10.0-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-supautils` | 2.10.0 | `d12.x86_64` | pigsty | 22.8 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-supautils` | 2.10.0 | `d12.aarch64` | pigsty | 22.1 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-supautils` | 2.10.0 | `u22.x86_64` | pigsty | 24.1 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-supautils` | 2.10.0 | `u22.aarch64` | pigsty | 23.6 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-supautils` | 2.10.0 | `u24.x86_64` | pigsty | 23.8 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-supautils` | 2.10.0 | `u24.aarch64` | pigsty | 23.7 KiB | [postgresql-13-supautils_2.10.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/supautils/postgresql-13-supautils_2.10.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/supabase/supautils" title="Repository" icon="github" subtitle="github.com/supabase/supautils" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="supautils-3.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get supautils; # get supautils source code
pig build dep supautils; # install build dependencies
pig build pkg supautils; # build extension rpm or deb
pig build ext supautils; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install supautils; # install by extension name, for the current active PG version
pig ext install supautils; # install via package alias, for the active PG version
pig ext install supautils -v 18;   # install for PG 18
pig ext install supautils -v 17;   # install for PG 17
pig ext install supautils -v 16;   # install for PG 16
pig ext install supautils -v 15;   # install for PG 15
pig ext install supautils -v 14;   # install for PG 14
pig ext install supautils -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION supautils;
```

