---
title: "pgjwt"
linkTitle: "pgjwt"
description: "JSON Web Token API for Postgresql"
weight: 4160
categories: ["UTIL"]
width: full
---

[**pgjwt**](https://github.com/michelp/pgjwt) : JSON Web Token API for Postgresql


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4160** | {{< badge content="pgjwt" link="https://github.com/michelp/pgjwt" >}} | {{< ext "pgjwt" >}} | `0.2.0` | {{< category "UTIL" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "pgcrypto" >}} |
|   **See Also**    | {{< ext "http" >}} {{< ext "pg_net" >}} {{< ext "pg_curl" >}} {{< ext "pgjq" >}} {{< ext "sparql" >}} {{< ext "pgcrypto" >}} {{< ext "gzip" >}} {{< ext "bzip" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgjwt` | `pgcrypto` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "pgjwt_18" "green" >}} {{< bg "17" "pgjwt_17" "green" >}} {{< bg "16" "pgjwt_16" "green" >}} {{< bg "15" "pgjwt_15" "green" >}} {{< bg "14" "pgjwt_14" "green" >}} | `pgjwt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pgjwt" "green" >}} {{< bg "17" "postgresql-17-pgjwt" "green" >}} {{< bg "16" "postgresql-16-pgjwt" "green" >}} {{< bg "15" "postgresql-15-pgjwt" "green" >}} {{< bg "14" "postgresql-14-pgjwt" "green" >}} | `postgresql-$v-pgjwt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pgjwt_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-pgjwt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-pgjwt : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgjwt : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgjwt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgjwt : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgjwt_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.4 KiB | [pgjwt_18-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjwt_18-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjwt_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [pgjwt_18-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjwt_18-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjwt_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.4 KiB | [pgjwt_18-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjwt_18-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjwt_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [pgjwt_18-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjwt_18-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjwt_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 KiB | [pgjwt_18-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgjwt_18-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgjwt_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.4 KiB | [pgjwt_18-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgjwt_18-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgjwt` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.1 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.1 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgjwt` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-18-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgjwt_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.4 KiB | [pgjwt_17-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjwt_17-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjwt_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [pgjwt_17-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjwt_17-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjwt_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.4 KiB | [pgjwt_17-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjwt_17-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjwt_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [pgjwt_17-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjwt_17-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjwt_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 KiB | [pgjwt_17-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgjwt_17-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgjwt_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.4 KiB | [pgjwt_17-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgjwt_17-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgjwt` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.1 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.1 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgjwt` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-17-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgjwt_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.4 KiB | [pgjwt_16-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjwt_16-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjwt_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [pgjwt_16-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjwt_16-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjwt_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.4 KiB | [pgjwt_16-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjwt_16-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjwt_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [pgjwt_16-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjwt_16-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjwt_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 KiB | [pgjwt_16-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgjwt_16-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgjwt_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.4 KiB | [pgjwt_16-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgjwt_16-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgjwt` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.1 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.1 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgjwt` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-16-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgjwt_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.4 KiB | [pgjwt_15-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjwt_15-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjwt_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [pgjwt_15-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjwt_15-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjwt_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.4 KiB | [pgjwt_15-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjwt_15-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjwt_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [pgjwt_15-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjwt_15-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjwt_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 KiB | [pgjwt_15-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgjwt_15-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgjwt_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.4 KiB | [pgjwt_15-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgjwt_15-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgjwt` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.1 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.1 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgjwt` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-15-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgjwt_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 9.4 KiB | [pgjwt_14-0.2.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgjwt_14-0.2.0-1PIGSTY.el8.x86_64.rpm) |
| `pgjwt_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.3 KiB | [pgjwt_14-0.2.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgjwt_14-0.2.0-1PIGSTY.el8.aarch64.rpm) |
| `pgjwt_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 9.4 KiB | [pgjwt_14-0.2.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgjwt_14-0.2.0-1PIGSTY.el9.x86_64.rpm) |
| `pgjwt_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.3 KiB | [pgjwt_14-0.2.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgjwt_14-0.2.0-1PIGSTY.el9.aarch64.rpm) |
| `pgjwt_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 9.5 KiB | [pgjwt_14-0.2.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgjwt_14-0.2.0-1PIGSTY.el10.x86_64.rpm) |
| `pgjwt_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.4 KiB | [pgjwt_14-0.2.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgjwt_14-0.2.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgjwt` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.1 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.1 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.1 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.1 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgjwt` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgjwt/postgresql-14-pgjwt_0.2.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michelp/pgjwt" title="Repository" icon="github" subtitle="github.com/michelp/pgjwt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgjwt-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgjwt;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgjwt;		# install via package name, for the active PG version

pig install pgjwt -v 18;   # install for PG 18
pig install pgjwt -v 17;   # install for PG 17
pig install pgjwt -v 16;   # install for PG 16
pig install pgjwt -v 15;   # install for PG 15
pig install pgjwt -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgjwt CASCADE; -- requires pgcrypto
```




## Usage

> [pgjwt: JSON Web Tokens for PostgreSQL](https://github.com/michelp/pgjwt)

Requires the `pgcrypto` extension.

### Sign a Token

```sql
SELECT sign(
    '{"sub":"1234567890","name":"John Doe","admin":true}',
    'secret'
);
```

Returns a JWT string like: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOi...`

### Verify a Token

```sql
SELECT * FROM verify(
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ',
    'secret'
);
```

Returns:

| header | payload | valid |
|--------|---------|-------|
| `{"alg":"HS256","typ":"JWT"}` | `{"sub":"1234567890","name":"John Doe","admin":true}` | `t` |

### Algorithm Selection

`sign()` and `verify()` accept an optional third argument for the algorithm: `'HS256'` (default), `'HS384'`, or `'HS512'`.

```sql
SELECT sign('{"data":"value"}', 'secret', 'HS384');
```
