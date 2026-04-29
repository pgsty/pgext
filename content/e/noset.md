---
title: "noset"
linkTitle: "noset"
description: "Module for blocking SET variables for non-super users."
weight: 7420
categories: ["SEC"]
width: full
---

[**pg_noset**](https://gitlab.com/ongresinc/extensions/noset) : Module for blocking SET variables for non-super users.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7420** | {{< badge content="noset" link="https://gitlab.com/ongresinc/extensions/noset" >}} | {{< ext "noset" "pg_noset" >}} | `0.3.0` | {{< category "SEC" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_readonly" >}} {{< ext "pg_permissions" >}} {{< ext "set_user" >}} {{< ext "pgaudit" >}} {{< ext "login_hook" >}} {{< ext "sepgsql" >}} {{< ext "safeupdate" >}} {{< ext "credcheck" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_noset` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "noset_18" "green" >}} {{< bg "17" "noset_17" "green" >}} {{< bg "16" "noset_16" "green" >}} {{< bg "15" "noset_15" "green" >}} {{< bg "14" "noset_14" "green" >}} | `noset_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.3.0` | {{< bg "18" "postgresql-18-noset" "green" >}} {{< bg "17" "postgresql-17-noset" "green" >}} {{< bg "16" "postgresql-16-noset" "green" >}} {{< bg "15" "postgresql-15-noset" "green" >}} {{< bg "14" "postgresql-14-noset" "green" >}} | `postgresql-$v-noset` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "noset_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "noset_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-18-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-17-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-16-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-15-noset : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.3.0" "postgresql-14-noset : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-noset : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-noset : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-noset : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_18` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_18-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_18` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [noset_18-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_18-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_18` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 15.9 KiB | [noset_18-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_18-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_18` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.4 KiB | [noset_18-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_18-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_18` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_18-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_18` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [noset_18-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_18-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-noset` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-noset` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.1 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-noset` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-noset` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-noset` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 28.9 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-noset` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 28.7 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-noset` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-noset` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-18-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-18-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_17` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_17-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_17` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [noset_17-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_17-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_17` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [noset_17-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_17-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_17` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_17-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_17` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_17-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_17` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [noset_17-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_17-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-noset` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-noset` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.0 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-noset` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-noset` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-noset` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-noset` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.6 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-noset` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-noset` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-17-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-17-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_16` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_16-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_16` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [noset_16-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_16-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_16` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [noset_16-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_16-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_16` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_16-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_16` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_16-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_16` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [noset_16-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_16-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-noset` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-noset` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.0 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-noset` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-noset` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-noset` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.2 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-noset` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.0 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-noset` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-noset` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-16-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-16-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_15` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_15-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_15` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.5 KiB | [noset_15-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_15-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_15` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [noset_15-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_15-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_15` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_15-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_15` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_15-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_15` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [noset_15-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_15-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-noset` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-noset` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.0 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-noset` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-noset` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-noset` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.2 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-noset` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.0 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-noset` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-noset` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-15-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-15-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `noset_14` | `0.3.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/noset_14-0.3.0-1PIGSTY.el8.x86_64.rpm) |
| `noset_14` | `0.3.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 15.4 KiB | [noset_14-0.3.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/noset_14-0.3.0-1PIGSTY.el8.aarch64.rpm) |
| `noset_14` | `0.3.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [noset_14-0.3.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/noset_14-0.3.0-1PIGSTY.el9.x86_64.rpm) |
| `noset_14` | `0.3.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/noset_14-0.3.0-1PIGSTY.el9.aarch64.rpm) |
| `noset_14` | `0.3.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/noset_14-0.3.0-1PIGSTY.el10.x86_64.rpm) |
| `noset_14` | `0.3.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 15.6 KiB | [noset_14-0.3.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/noset_14-0.3.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-noset` | `0.3.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 27.3 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-noset` | `0.3.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 27.0 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-noset` | `0.3.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 27.3 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-noset` | `0.3.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 27.0 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-noset` | `0.3.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.0 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-noset` | `0.3.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.8 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-noset` | `0.3.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 28.3 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-noset` | `0.3.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 27.9 KiB | [postgresql-14-noset_0.3.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/n/noset/postgresql-14-noset_0.3.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitlab.com/ongresinc/extensions/noset" title="Repository" icon="link" subtitle="gitlab.com/ongresinc/extensions/noset" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="noset-v0.3.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_noset;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_noset;		# install via package name, for the active PG version
pig install noset;		# install by extension name, for the current active PG version

pig install noset -v 18;   # install for PG 18
pig install noset -v 17;   # install for PG 17
pig install noset -v 16;   # install for PG 16
pig install noset -v 15;   # install for PG 15
pig install noset -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'noset';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION noset;
```




## Usage

> [noset: Prevent users from changing session parameters via SET/RESET](https://gitlab.com/ongresinc/extensions/noset)

`noset` is a loadable module that prevents specific users from using `SET` or `RESET` commands to change session parameters.

```sql
CREATE EXTENSION noset;
```

### Configuration

Add to `postgresql.conf`:

```ini
shared_preload_libraries = 'noset'
```

### GUC Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `noset.enabled` | `false` | Enable SET/RESET blocking for the role |
| `noset.parameters` | `*` | Parameters to block (comma-separated, `*` = all) |

### Setting Up Per-User Restrictions

```sql
-- Block ALL SET/RESET for a user
ALTER USER appuser SET noset.enabled = true;

-- Block only specific parameters
ALTER USER appuser SET noset.enabled = true;
ALTER USER appuser SET noset.parameters = 'work_mem,jit';
```

### Example

```sql
-- As appuser:
SET work_mem = '1GB';
-- ERROR: permission denied to set/reset parameter 'set work_mem = '1GB';'

SET maintenance_work_mem = '1GB';
-- SET (allowed, not in blocked list)
```

### Finding Restricted Users

```sql
SELECT usename, useconfig FROM pg_user
WHERE useconfig IS NOT NULL
  AND array['noset.enabled=on'] <@ useconfig;
```

### Notes

- Does not apply to superusers
- The extension revokes access to the `set_config` function from PUBLIC
