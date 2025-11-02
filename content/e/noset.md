---
title: "noset"
linkTitle: "noset"
description: "Module for blocking SET variables for non-super users."
weight: 7210
categories: ["SEC"]
width: full
---

Module for blocking SET variables for non-super users.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7210** | {{< badge content="noset" link="https://gitlab.com/ongresinc/extensions/noset" >}} | {{< ext "noset" "pg_noset" >}} | `0.3.0` | {{< category "SEC" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_permissions" >}} {{< ext "set_user" >}} {{< ext "pgaudit" >}} {{< ext "login_hook" >}} {{< ext "sepgsql" >}} {{< ext "safeupdate" >}} {{< ext "credcheck" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/noset" >}} | `0.3.0` | {{< bg "18" "noset_18*" "green" >}} {{< bg "17" "noset_17*" "green" >}} {{< bg "16" "noset_16*" "green" >}} {{< bg "15" "noset_15*" "green" >}} {{< bg "14" "noset_14*" "green" >}} {{< bg "13" "noset_13*" "green" >}} | `noset_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/noset" >}} | `0.3.0` | {{< bg "18" "postgresql-18-noset" "red" >}} {{< bg "17" "postgresql-17-noset" "green" >}} {{< bg "16" "postgresql-16-noset" "green" >}} {{< bg "15" "postgresql-15-noset" "green" >}} {{< bg "14" "postgresql-14-noset" "green" >}} {{< bg "13" "postgresql-13-noset" "green" >}} | `postgresql-$v-noset` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-noset : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-noset : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-13-noset : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_18` | 0.3.0 | `el8.x86_64` | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_18` | 0.3.0 | `el8.aarch64` | pigsty | 15.4 KiB | [noset_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_18` | 0.3.0 | `el9.x86_64` | pigsty | 15.9 KiB | [noset_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_18` | 0.3.0 | `el9.aarch64` | pigsty | 15.4 KiB | [noset_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_18` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_18` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_17` | 0.3.0 | `el8.x86_64` | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_17` | 0.3.0 | `el8.aarch64` | pigsty | 15.4 KiB | [noset_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_17` | 0.3.0 | `el9.x86_64` | pigsty | 16.0 KiB | [noset_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_17` | 0.3.0 | `el9.aarch64` | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_17` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_17` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-noset` | 0.3.0 | `d12.x86_64` | pigsty | 31.3 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-noset` | 0.3.0 | `d12.aarch64` | pigsty | 30.8 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-noset` | 0.3.0 | `u22.x86_64` | pigsty | 32.8 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-noset` | 0.3.0 | `u22.aarch64` | pigsty | 32.6 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-noset` | 0.3.0 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-noset` | 0.3.0 | `u24.aarch64` | pigsty | 27.8 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_16` | 0.3.0 | `el8.x86_64` | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_16` | 0.3.0 | `el8.aarch64` | pigsty | 15.4 KiB | [noset_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_16` | 0.3.0 | `el9.x86_64` | pigsty | 16.0 KiB | [noset_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_16` | 0.3.0 | `el9.aarch64` | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_16` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_16` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-noset` | 0.3.0 | `d12.x86_64` | pigsty | 30.7 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-noset` | 0.3.0 | `d12.aarch64` | pigsty | 30.2 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-noset` | 0.3.0 | `u22.x86_64` | pigsty | 32.2 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-noset` | 0.3.0 | `u22.aarch64` | pigsty | 32.0 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-noset` | 0.3.0 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-noset` | 0.3.0 | `u24.aarch64` | pigsty | 27.8 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_15` | 0.3.0 | `el8.x86_64` | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_15` | 0.3.0 | `el8.aarch64` | pigsty | 15.5 KiB | [noset_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_15` | 0.3.0 | `el9.x86_64` | pigsty | 16.0 KiB | [noset_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_15` | 0.3.0 | `el9.aarch64` | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_15` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_15` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-noset` | 0.3.0 | `d12.x86_64` | pigsty | 30.7 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-noset` | 0.3.0 | `d12.aarch64` | pigsty | 30.3 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-noset` | 0.3.0 | `u22.x86_64` | pigsty | 32.2 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-noset` | 0.3.0 | `u22.aarch64` | pigsty | 32.1 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-noset` | 0.3.0 | `u24.x86_64` | pigsty | 28.5 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-noset` | 0.3.0 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_14` | 0.3.0 | `el8.x86_64` | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_14` | 0.3.0 | `el8.aarch64` | pigsty | 15.4 KiB | [noset_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_14` | 0.3.0 | `el9.x86_64` | pigsty | 16.0 KiB | [noset_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_14` | 0.3.0 | `el9.aarch64` | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_14` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_14` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-noset` | 0.3.0 | `d12.x86_64` | pigsty | 29.6 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-noset` | 0.3.0 | `d12.aarch64` | pigsty | 29.1 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-noset` | 0.3.0 | `u22.x86_64` | pigsty | 31.0 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-noset` | 0.3.0 | `u22.aarch64` | pigsty | 30.8 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-noset` | 0.3.0 | `u24.x86_64` | pigsty | 28.4 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-noset` | 0.3.0 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_13` | 0.3.0 | `el8.x86_64` | pigsty | 15.5 KiB | [noset_13-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_13-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_13` | 0.3.0 | `el8.aarch64` | pigsty | 15.4 KiB | [noset_13-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_13-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_13` | 0.3.0 | `el9.x86_64` | pigsty | 15.9 KiB | [noset_13-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_13-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_13` | 0.3.0 | `el9.aarch64` | pigsty | 15.6 KiB | [noset_13-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_13-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_13` | 0.3.0 | `el10.x86_64` | pigsty | 15.6 KiB | [noset_13-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_13-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_13` | 0.3.0 | `el10.aarch64` | pigsty | 15.6 KiB | [noset_13-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_13-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-noset` | 0.3.0 | `d12.x86_64` | pigsty | 29.3 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-noset` | 0.3.0 | `d12.aarch64` | pigsty | 29.1 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-noset` | 0.3.0 | `u22.x86_64` | pigsty | 30.8 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-noset` | 0.3.0 | `u22.aarch64` | pigsty | 30.5 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-noset` | 0.3.0 | `u24.x86_64` | pigsty | 28.4 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-noset` | 0.3.0 | `u24.aarch64` | pigsty | 27.9 KiB | [postgresql-13-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-13-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/ongresinc/extensions/noset" title="Repository" icon="link" subtitle="gitlab.com/ongresinc/extensions/noset" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="noset-v0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get noset; # get noset source code
pig build dep noset; # install build dependencies
pig build pkg noset; # build extension rpm or deb
pig build ext noset; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install noset; # install by extension name, for the current active PG version
pig ext install pg_noset; # install via package alias, for the active PG version
pig ext install noset -v 18;   # install for PG 18
pig ext install noset -v 17;   # install for PG 17
pig ext install noset -v 16;   # install for PG 16
pig ext install noset -v 15;   # install for PG 15
pig ext install noset -v 14;   # install for PG 14
pig ext install noset -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION noset;
```

