---
title: "istore"
linkTitle: "istore"
description: "Integer key/value types with arithmetic, aggregation, and GIN indexing"
weight: 3710
categories: ["TYPE"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**istore**](https://github.com/adjust/istore) : Integer key/value types with arithmetic, aggregation, and GIN indexing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3710** | {{< badge content="istore" link="https://github.com/adjust/istore" >}} | {{< ext "istore" >}} | `0.1.12` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "hstore" >}} {{< ext "collection" >}} {{< ext "intarray" >}} {{< ext "aggs_for_arrays" >}} |

> [!Note] PGXN 0.1.12 source requires the PG16+ varatt/Datum compatibility patch. PostgreSQL 14-18 verified across all 16 PGSTY target platforms. Does not require hstore.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.12` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `istore` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.12` | {{< bg "18" "istore_18" "green" >}} {{< bg "17" "istore_17" "green" >}} {{< bg "16" "istore_16" "green" >}} {{< bg "15" "istore_15" "green" >}} {{< bg "14" "istore_14" "green" >}} | `istore_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.12` | {{< bg "18" "postgresql-18-istore" "green" >}} {{< bg "17" "postgresql-17-istore" "green" >}} {{< bg "16" "postgresql-16-istore" "green" >}} {{< bg "15" "postgresql-15-istore" "green" >}} {{< bg "14" "postgresql-14-istore" "green" >}} | `postgresql-$v-istore` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "istore_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "istore_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-18-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-17-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-16-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-15-istore : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.12" "postgresql-14-istore : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `istore_18` | `0.1.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 129.8 KiB | [istore_18-0.1.12-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/istore_18-0.1.12-1PGSTY.el8.x86_64.rpm) |
| `istore_18` | `0.1.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 126.9 KiB | [istore_18-0.1.12-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/istore_18-0.1.12-1PGSTY.el8.aarch64.rpm) |
| `istore_18` | `0.1.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.3 KiB | [istore_18-0.1.12-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/istore_18-0.1.12-1PGSTY.el9.x86_64.rpm) |
| `istore_18` | `0.1.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 131.6 KiB | [istore_18-0.1.12-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/istore_18-0.1.12-1PGSTY.el9.aarch64.rpm) |
| `istore_18` | `0.1.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 133.9 KiB | [istore_18-0.1.12-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/istore_18-0.1.12-1PGSTY.el10.x86_64.rpm) |
| `istore_18` | `0.1.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 132.2 KiB | [istore_18-0.1.12-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/istore_18-0.1.12-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-istore` | `0.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 109.2 KiB | [postgresql-18-istore_0.1.12-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-istore` | `0.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 106.7 KiB | [postgresql-18-istore_0.1.12-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-istore` | `0.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 109.4 KiB | [postgresql-18-istore_0.1.12-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-istore` | `0.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 106.9 KiB | [postgresql-18-istore_0.1.12-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 121.9 KiB | [postgresql-18-istore_0.1.12-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 119.9 KiB | [postgresql-18-istore_0.1.12-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 118.2 KiB | [postgresql-18-istore_0.1.12-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~noble_amd64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 117.3 KiB | [postgresql-18-istore_0.1.12-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~noble_arm64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u26.x86_64](/os/u26.x86_64) | pigsty | 116.9 KiB | [postgresql-18-istore_0.1.12-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-istore` | `0.1.12` | [u26.aarch64](/os/u26.aarch64) | pigsty | 115.3 KiB | [postgresql-18-istore_0.1.12-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-18-istore_0.1.12-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `istore_17` | `0.1.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 128.8 KiB | [istore_17-0.1.12-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/istore_17-0.1.12-1PGSTY.el8.x86_64.rpm) |
| `istore_17` | `0.1.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 127.0 KiB | [istore_17-0.1.12-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/istore_17-0.1.12-1PGSTY.el8.aarch64.rpm) |
| `istore_17` | `0.1.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 132.2 KiB | [istore_17-0.1.12-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/istore_17-0.1.12-1PGSTY.el9.x86_64.rpm) |
| `istore_17` | `0.1.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 131.7 KiB | [istore_17-0.1.12-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/istore_17-0.1.12-1PGSTY.el9.aarch64.rpm) |
| `istore_17` | `0.1.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 132.9 KiB | [istore_17-0.1.12-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/istore_17-0.1.12-1PGSTY.el10.x86_64.rpm) |
| `istore_17` | `0.1.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 132.2 KiB | [istore_17-0.1.12-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/istore_17-0.1.12-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-istore` | `0.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.2 KiB | [postgresql-17-istore_0.1.12-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-istore` | `0.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.5 KiB | [postgresql-17-istore_0.1.12-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-istore` | `0.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.3 KiB | [postgresql-17-istore_0.1.12-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-istore` | `0.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.8 KiB | [postgresql-17-istore_0.1.12-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.6 KiB | [postgresql-17-istore_0.1.12-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.3 KiB | [postgresql-17-istore_0.1.12-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.2 KiB | [postgresql-17-istore_0.1.12-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~noble_amd64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 116.1 KiB | [postgresql-17-istore_0.1.12-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~noble_arm64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u26.x86_64](/os/u26.x86_64) | pigsty | 116.0 KiB | [postgresql-17-istore_0.1.12-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-istore` | `0.1.12` | [u26.aarch64](/os/u26.aarch64) | pigsty | 114.4 KiB | [postgresql-17-istore_0.1.12-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-17-istore_0.1.12-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `istore_16` | `0.1.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 128.8 KiB | [istore_16-0.1.12-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/istore_16-0.1.12-1PGSTY.el8.x86_64.rpm) |
| `istore_16` | `0.1.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 127.0 KiB | [istore_16-0.1.12-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/istore_16-0.1.12-1PGSTY.el8.aarch64.rpm) |
| `istore_16` | `0.1.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 132.2 KiB | [istore_16-0.1.12-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/istore_16-0.1.12-1PGSTY.el9.x86_64.rpm) |
| `istore_16` | `0.1.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 131.7 KiB | [istore_16-0.1.12-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/istore_16-0.1.12-1PGSTY.el9.aarch64.rpm) |
| `istore_16` | `0.1.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 132.9 KiB | [istore_16-0.1.12-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/istore_16-0.1.12-1PGSTY.el10.x86_64.rpm) |
| `istore_16` | `0.1.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 132.2 KiB | [istore_16-0.1.12-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/istore_16-0.1.12-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-istore` | `0.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 108.1 KiB | [postgresql-16-istore_0.1.12-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-istore` | `0.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.5 KiB | [postgresql-16-istore_0.1.12-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-istore` | `0.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.3 KiB | [postgresql-16-istore_0.1.12-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-istore` | `0.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.9 KiB | [postgresql-16-istore_0.1.12-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.5 KiB | [postgresql-16-istore_0.1.12-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.2 KiB | [postgresql-16-istore_0.1.12-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.3 KiB | [postgresql-16-istore_0.1.12-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~noble_amd64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 116.2 KiB | [postgresql-16-istore_0.1.12-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~noble_arm64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u26.x86_64](/os/u26.x86_64) | pigsty | 116.0 KiB | [postgresql-16-istore_0.1.12-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-istore` | `0.1.12` | [u26.aarch64](/os/u26.aarch64) | pigsty | 114.4 KiB | [postgresql-16-istore_0.1.12-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-16-istore_0.1.12-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `istore_15` | `0.1.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 128.3 KiB | [istore_15-0.1.12-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/istore_15-0.1.12-1PGSTY.el8.x86_64.rpm) |
| `istore_15` | `0.1.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 126.6 KiB | [istore_15-0.1.12-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/istore_15-0.1.12-1PGSTY.el8.aarch64.rpm) |
| `istore_15` | `0.1.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 129.1 KiB | [istore_15-0.1.12-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/istore_15-0.1.12-1PGSTY.el9.x86_64.rpm) |
| `istore_15` | `0.1.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 128.2 KiB | [istore_15-0.1.12-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/istore_15-0.1.12-1PGSTY.el9.aarch64.rpm) |
| `istore_15` | `0.1.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 129.5 KiB | [istore_15-0.1.12-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/istore_15-0.1.12-1PGSTY.el10.x86_64.rpm) |
| `istore_15` | `0.1.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.5 KiB | [istore_15-0.1.12-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/istore_15-0.1.12-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-istore` | `0.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 107.9 KiB | [postgresql-15-istore_0.1.12-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-istore` | `0.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 105.0 KiB | [postgresql-15-istore_0.1.12-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-istore` | `0.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 108.0 KiB | [postgresql-15-istore_0.1.12-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-istore` | `0.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 105.4 KiB | [postgresql-15-istore_0.1.12-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 126.1 KiB | [postgresql-15-istore_0.1.12-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 123.3 KiB | [postgresql-15-istore_0.1.12-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 113.8 KiB | [postgresql-15-istore_0.1.12-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~noble_amd64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 112.3 KiB | [postgresql-15-istore_0.1.12-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~noble_arm64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u26.x86_64](/os/u26.x86_64) | pigsty | 112.6 KiB | [postgresql-15-istore_0.1.12-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-istore` | `0.1.12` | [u26.aarch64](/os/u26.aarch64) | pigsty | 110.8 KiB | [postgresql-15-istore_0.1.12-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-15-istore_0.1.12-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `istore_14` | `0.1.12` | [el8.x86_64](/os/el8.x86_64) | pigsty | 128.2 KiB | [istore_14-0.1.12-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/istore_14-0.1.12-1PGSTY.el8.x86_64.rpm) |
| `istore_14` | `0.1.12` | [el8.aarch64](/os/el8.aarch64) | pigsty | 126.5 KiB | [istore_14-0.1.12-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/istore_14-0.1.12-1PGSTY.el8.aarch64.rpm) |
| `istore_14` | `0.1.12` | [el9.x86_64](/os/el9.x86_64) | pigsty | 129.0 KiB | [istore_14-0.1.12-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/istore_14-0.1.12-1PGSTY.el9.x86_64.rpm) |
| `istore_14` | `0.1.12` | [el9.aarch64](/os/el9.aarch64) | pigsty | 128.1 KiB | [istore_14-0.1.12-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/istore_14-0.1.12-1PGSTY.el9.aarch64.rpm) |
| `istore_14` | `0.1.12` | [el10.x86_64](/os/el10.x86_64) | pigsty | 129.4 KiB | [istore_14-0.1.12-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/istore_14-0.1.12-1PGSTY.el10.x86_64.rpm) |
| `istore_14` | `0.1.12` | [el10.aarch64](/os/el10.aarch64) | pigsty | 128.5 KiB | [istore_14-0.1.12-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/istore_14-0.1.12-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-istore` | `0.1.12` | [d12.x86_64](/os/d12.x86_64) | pigsty | 111.6 KiB | [postgresql-14-istore_0.1.12-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-istore` | `0.1.12` | [d12.aarch64](/os/d12.aarch64) | pigsty | 108.8 KiB | [postgresql-14-istore_0.1.12-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-istore` | `0.1.12` | [d13.x86_64](/os/d13.x86_64) | pigsty | 111.6 KiB | [postgresql-14-istore_0.1.12-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-istore` | `0.1.12` | [d13.aarch64](/os/d13.aarch64) | pigsty | 109.1 KiB | [postgresql-14-istore_0.1.12-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u22.x86_64](/os/u22.x86_64) | pigsty | 129.8 KiB | [postgresql-14-istore_0.1.12-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u22.aarch64](/os/u22.aarch64) | pigsty | 127.1 KiB | [postgresql-14-istore_0.1.12-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u24.x86_64](/os/u24.x86_64) | pigsty | 117.3 KiB | [postgresql-14-istore_0.1.12-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~noble_amd64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u24.aarch64](/os/u24.aarch64) | pigsty | 115.9 KiB | [postgresql-14-istore_0.1.12-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~noble_arm64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u26.x86_64](/os/u26.x86_64) | pigsty | 116.2 KiB | [postgresql-14-istore_0.1.12-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-istore` | `0.1.12` | [u26.aarch64](/os/u26.aarch64) | pigsty | 114.4 KiB | [postgresql-14-istore_0.1.12-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/istore/postgresql-14-istore_0.1.12-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/adjust/istore" title="Repository" icon="github" subtitle="github.com/adjust/istore" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="istore-0.1.12.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg istore;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install istore;		# install via package name, for the active PG version

pig install istore -v 18;   # install for PG 18
pig install istore -v 17;   # install for PG 17
pig install istore -v 16;   # install for PG 16
pig install istore -v 15;   # install for PG 15
pig install istore -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION istore;
```

## Usage

Sources:

- [Official README](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/README.md)
- [Official extension SQL](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore--0.1.12.sql)
- [Official extension control file](https://github.com/adjust/istore/blob/46c1cfeceeea193b75fd3fa11bc6959d8ac4d26f/istore.control)

`istore` provides compact integer-key maps for PostgreSQL. The `istore` type stores integer values, while `bigistore` stores bigint values; both are useful for sparse counters, distributions, and analytical aggregates keyed by integer identifiers.

### Core Workflow

Text input resembles hstore, and array constructors can aggregate repeated keys:

```sql
CREATE EXTENSION istore;

SELECT istore(ARRAY[1, 2, 1, 3, 2, 2]);
SELECT '1=>4,2=>5'::istore -> 1;
SELECT '1=>4,2=>5'::istore + '1=>4,3=>6'::istore;

CREATE INDEX metrics_keys_idx
  ON metrics USING gin (values istore_key_ops);
```

Key lookup uses `->`; `?`, `?&`, and `?|` test key presence. `||` concatenates stores, while arithmetic operators combine matching keys. The `istore_key_ops` GIN operator class accelerates key-existence predicates.

### Important Objects and Caveats

Constructors and transforms include `istore(...)`, `bigistore(...)`, `akeys`, `avals`, `each`, `slice`, `compact`, `fill_gaps`, `accumulate`, `clamp`, and delete helpers. `sum_up` summarizes values, and aggregate SUM returns `bigistore` so that totals have a wider value type. Per-key minimum and maximum aggregates are also provided.

The cast from `istore` to `bigistore` is implicit; narrowing assignments can overflow. Arithmetic and division still require application-level checks for overflow and zero divisors. Large values can be TOASTed, and GIN indexes add write and maintenance cost, so benchmark the actual distribution and update workload instead of treating upstream examples as performance guarantees.
