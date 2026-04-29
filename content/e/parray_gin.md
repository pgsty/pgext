---
title: "parray_gin"
linkTitle: "parray_gin"
description: "GIN index operator class and partial-match operators for text arrays"
weight: 4860
categories: ["FUNC"]
width: full
---

[**parray_gin**](https://github.com/theirix/parray_gin) : GIN index operator class and partial-match operators for text arrays


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4860** | {{< badge content="parray_gin" link="https://github.com/theirix/parray_gin" >}} | {{< ext "parray_gin" >}} | `1.5.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "intarray" >}} {{< ext "btree_gin" >}} {{< ext "btree_gist" >}} {{< ext "pg_trgm" >}} {{< ext "smlar" >}} {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "arraymath" >}} |

> [!Note] PGXN dist name and PostgreSQL extension name are both parray_gin; Pigsty packages are built for PG 14-18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `parray_gin` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "parray_gin_18" "green" >}} {{< bg "17" "parray_gin_17" "green" >}} {{< bg "16" "parray_gin_16" "green" >}} {{< bg "15" "parray_gin_15" "green" >}} {{< bg "14" "parray_gin_14" "green" >}} | `parray_gin_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.0` | {{< bg "18" "postgresql-18-parray-gin" "green" >}} {{< bg "17" "postgresql-17-parray-gin" "green" >}} {{< bg "16" "postgresql-16-parray-gin" "green" >}} {{< bg "15" "postgresql-15-parray-gin" "green" >}} {{< bg "14" "postgresql-14-parray-gin" "green" >}} | `postgresql-$v-parray-gin` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "parray_gin_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-18-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-17-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-parray-gin : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-parray-gin : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-parray-gin : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-parray-gin : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-parray-gin : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `parray_gin_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [parray_gin_18-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/parray_gin_18-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `parray_gin_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.4 KiB | [parray_gin_18-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/parray_gin_18-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `parray_gin_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.0 KiB | [parray_gin_18-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/parray_gin_18-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `parray_gin_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.8 KiB | [parray_gin_18-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/parray_gin_18-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `parray_gin_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.2 KiB | [parray_gin_18-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/parray_gin_18-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `parray_gin_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.1 KiB | [parray_gin_18-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/parray_gin_18-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-parray-gin` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.1 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.0 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.3 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.2 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 31.5 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 31.0 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.0 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-parray-gin` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.3 KiB | [postgresql-18-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-18-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `parray_gin_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [parray_gin_17-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/parray_gin_17-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `parray_gin_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [parray_gin_17-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/parray_gin_17-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `parray_gin_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.0 KiB | [parray_gin_17-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/parray_gin_17-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `parray_gin_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.8 KiB | [parray_gin_17-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/parray_gin_17-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `parray_gin_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.2 KiB | [parray_gin_17-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/parray_gin_17-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `parray_gin_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.2 KiB | [parray_gin_17-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/parray_gin_17-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-parray-gin` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.3 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.1 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.4 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.3 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.7 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.1 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-parray-gin` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.5 KiB | [postgresql-17-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-17-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `parray_gin_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [parray_gin_16-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/parray_gin_16-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `parray_gin_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [parray_gin_16-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/parray_gin_16-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `parray_gin_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.0 KiB | [parray_gin_16-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/parray_gin_16-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `parray_gin_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.8 KiB | [parray_gin_16-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/parray_gin_16-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `parray_gin_16` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.2 KiB | [parray_gin_16-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/parray_gin_16-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `parray_gin_16` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.2 KiB | [parray_gin_16-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/parray_gin_16-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-parray-gin` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.3 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 30.1 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.3 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.4 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.7 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.1 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-parray-gin` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.5 KiB | [postgresql-16-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-16-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `parray_gin_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [parray_gin_15-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/parray_gin_15-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `parray_gin_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [parray_gin_15-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/parray_gin_15-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `parray_gin_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.7 KiB | [parray_gin_15-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/parray_gin_15-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `parray_gin_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.2 KiB | [parray_gin_15-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/parray_gin_15-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `parray_gin_15` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.8 KiB | [parray_gin_15-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/parray_gin_15-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `parray_gin_15` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.6 KiB | [parray_gin_15-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/parray_gin_15-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-parray-gin` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.1 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.9 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.2 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.1 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.2 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-parray-gin` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.5 KiB | [postgresql-15-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-15-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `parray_gin_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 22.7 KiB | [parray_gin_14-1.4.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/parray_gin_14-1.4.0-1PIGSTY.el8.x86_64.rpm) |
| `parray_gin_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 22.5 KiB | [parray_gin_14-1.4.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/parray_gin_14-1.4.0-1PIGSTY.el8.aarch64.rpm) |
| `parray_gin_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 22.6 KiB | [parray_gin_14-1.4.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/parray_gin_14-1.4.0-1PIGSTY.el9.x86_64.rpm) |
| `parray_gin_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.2 KiB | [parray_gin_14-1.4.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/parray_gin_14-1.4.0-1PIGSTY.el9.aarch64.rpm) |
| `parray_gin_14` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 22.8 KiB | [parray_gin_14-1.4.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/parray_gin_14-1.4.0-1PIGSTY.el10.x86_64.rpm) |
| `parray_gin_14` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 22.6 KiB | [parray_gin_14-1.4.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/parray_gin_14-1.4.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-parray-gin` | `1.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 30.0 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 29.9 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 30.1 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 30.1 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 32.8 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 32.2 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 31.1 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-parray-gin` | `1.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 31.4 KiB | [postgresql-14-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/parray-gin/postgresql-14-parray-gin_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theirix/parray_gin" title="Repository" icon="github" subtitle="github.com/theirix/parray_gin" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="parray_gin-1.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg parray_gin;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install parray_gin;		# install via package name, for the active PG version

pig install parray_gin -v 18;   # install for PG 18
pig install parray_gin -v 17;   # install for PG 17
pig install parray_gin -v 16;   # install for PG 16
pig install parray_gin -v 15;   # install for PG 15
pig install parray_gin -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION parray_gin;
```

## Usage

- Sources: [README](https://github.com/theirix/parray_gin/blob/master/README.md), [reference doc](https://github.com/theirix/parray_gin/blob/master/doc/parray_gin.md)

`parray_gin` adds a GIN operator class for `text[]` plus strict and partial-match operators. Upstream describes it as trigram-based array indexing built on `pg_trgm` trigram decomposition.

### Create The Extension And Index

```sql
CREATE EXTENSION parray_gin;

CREATE TABLE test_table (
  id  bigserial,
  val text[]
);

CREATE INDEX test_tags_idx
ON test_table
USING gin (val parray_gin_ops);
```

### Indexed Operators

The reference doc says `parray_gin_ops` can support these operators:

- `@>`: strict contains.
- `<@`: strict contained-by.
- `@@>`: partial contains, where array elements may use `LIKE` patterns.
- `<@@`: partial contained-by.

Examples:

```sql
SELECT * FROM test_table WHERE val @> ARRAY['must','contain'];
SELECT * FROM test_table WHERE val @@> ARRAY['what%like%'];
SELECT * FROM test_table WHERE val <@ ARRAY['galaxy','ago','vader'];
SELECT * FROM test_table WHERE val <@@ ARRAY['%ar%','vader'];
```

### Matching Behavior

Strict matching requires array-item equality. Partial matching allows patterns such as `'foo%'` or `'%oo%'`. Because the trigram index can return false positives, the docs note that matches are rechecked after index lookup.

### Caveats

The README says support extends through PostgreSQL 18, while the reference document still says 9.1-14. The operator and opclass behavior is consistent across both docs, but the version note is not fully synchronized upstream.
