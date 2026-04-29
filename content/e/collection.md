---
title: "collection"
linkTitle: "collection"
description: "Memory optimized data type to be used inside of plpglsql func"
weight: 3690
categories: ["TYPE"]
width: full
---

[**pgcollection**](https://github.com/aws/pgcollection) : Memory optimized data type to be used inside of plpglsql func


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3690** | {{< badge content="collection" link="https://github.com/aws/pgcollection" >}} | {{< ext "collection" "pgcollection" >}} | `2.0.0` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgcollection` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "pgcollection_18" "green" >}} {{< bg "17" "pgcollection_17" "green" >}} {{< bg "16" "pgcollection_16" "green" >}} {{< bg "15" "pgcollection_15" "green" >}} {{< bg "14" "pgcollection_14" "green" >}} | `pgcollection_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.0.0` | {{< bg "18" "postgresql-18-collection" "green" >}} {{< bg "17" "postgresql-17-collection" "green" >}} {{< bg "16" "postgresql-16-collection" "green" >}} {{< bg "15" "postgresql-15-collection" "green" >}} {{< bg "14" "postgresql-14-collection" "green" >}} | `postgresql-$v-collection` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-collection : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-collection : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_18` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.1 KiB | [pgcollection_18-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_18-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_18` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.1 KiB | [pgcollection_18-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_18-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_18` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.1 KiB | [pgcollection_18-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_18-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_18` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.3 KiB | [pgcollection_18-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_18-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_18` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pgcollection_18-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_18-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_18` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.6 KiB | [pgcollection_18-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_18-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-collection` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.8 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-collection` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.6 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-collection` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 132.7 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-collection` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 129.0 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-collection` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 149.4 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-collection` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 148.0 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-collection` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 139.8 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-collection` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 137.7 KiB | [postgresql-18-collection_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_17` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.0 KiB | [pgcollection_17-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_17-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_17` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.1 KiB | [pgcollection_17-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_17-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_17` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 54.1 KiB | [pgcollection_17-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_17-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_17` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.3 KiB | [pgcollection_17-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_17-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_17` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.4 KiB | [pgcollection_17-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_17-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_17` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.6 KiB | [pgcollection_17-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_17-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-collection` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.5 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-collection` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.3 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-collection` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 132.4 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-collection` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 128.6 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-collection` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 164.1 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-collection` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.8 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-collection` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 139.6 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-collection` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 137.5 KiB | [postgresql-17-collection_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_16` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 55.9 KiB | [pgcollection_16-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_16-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_16` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 53.9 KiB | [pgcollection_16-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_16-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_16` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.8 KiB | [pgcollection_16-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_16-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_16` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.0 KiB | [pgcollection_16-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_16-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_16` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.3 KiB | [pgcollection_16-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_16-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_16` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 53.2 KiB | [pgcollection_16-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_16-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-collection` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 130.4 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-collection` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 127.2 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-collection` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 131.4 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-collection` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 127.6 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-collection` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 162.7 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-collection` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 161.1 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-collection` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 138.2 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-collection` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 136.1 KiB | [postgresql-16-collection_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_15` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.7 KiB | [pgcollection_15-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_15-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_15` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.7 KiB | [pgcollection_15-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_15-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_15` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.7 KiB | [pgcollection_15-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_15-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_15` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 54.0 KiB | [pgcollection_15-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_15-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_15` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 56.1 KiB | [pgcollection_15-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_15-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_15` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.1 KiB | [pgcollection_15-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_15-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-collection` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.4 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-collection` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.1 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-collection` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 131.9 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-collection` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 128.7 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-collection` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 162.3 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-collection` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.0 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-collection` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 138.7 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-collection` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 136.9 KiB | [postgresql-15-collection_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_14` | `2.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 56.7 KiB | [pgcollection_14-2.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_14-2.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_14` | `2.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 54.7 KiB | [pgcollection_14-2.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_14-2.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_14` | `2.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 53.7 KiB | [pgcollection_14-2.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_14-2.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_14` | `2.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.9 KiB | [pgcollection_14-2.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_14-2.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_14` | `2.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 56.0 KiB | [pgcollection_14-2.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_14-2.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_14` | `2.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.1 KiB | [pgcollection_14-2.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_14-2.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-collection` | `2.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 131.2 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-collection` | `2.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 128.1 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-collection` | `2.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 131.9 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-collection` | `2.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 128.6 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-collection` | `2.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 162.5 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-collection` | `2.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 162.1 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-collection` | `2.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 138.6 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-collection` | `2.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 136.8 KiB | [postgresql-14-collection_2.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_2.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgcollection" title="Repository" icon="github" subtitle="github.com/aws/pgcollection" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcollection-2.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgcollection;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgcollection;		# install via package name, for the active PG version
pig install collection;		# install by extension name, for the current active PG version

pig install collection -v 18;   # install for PG 18
pig install collection -v 17;   # install for PG 17
pig install collection -v 16;   # install for PG 16
pig install collection -v 15;   # install for PG 15
pig install collection -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION collection;
```



## Usage

> [collection: key-value collection data types for PL/pgSQL](https://github.com/aws/pgcollection)

The `collection` extension provides two memory-optimized collection data types for use within PL/pgSQL functions.

```sql
CREATE EXTENSION collection;
```

### Data Types

- **`collection`**: Key-value pairs with text keys (max 32,767 chars), stored in creation order
- **`icollection`**: Key-value pairs with 64-bit integer keys, enabling sparse arrays

Both types support type modifiers to specify element types:

```sql
DECLARE
    c1  collection('date');
    ic1 icollection('int4');
```

### Subscript Access

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['Japan'] := 'Tokyo';
    RAISE NOTICE '%', t_capital['USA'];  -- Washington, D.C.
END $$;
```

### Core Functions

| Function | Description |
|----------|-------------|
| `add(coll, key, value)` | Add element |
| `count(coll)` | Element count |
| `delete(coll, key)` | Remove element |
| `exist(coll, key)` | Check key existence |
| `find(coll, key)` | Retrieve value |
| `first(coll)` | Move iterator to start |
| `last(coll)` | Move iterator to end |
| `next(coll)` | Advance iterator |
| `prev(coll)` | Reverse iterator |
| `key(coll)` | Current key |
| `value(coll)` | Current value |
| `copy(coll)` | Create copy |
| `sort(coll)` | Sort by keys |
| `keys_to_table(coll)` | All keys as set |
| `values_to_table(coll)` | All values as set |
| `to_table(coll)` | Keys and values as table |

### Iterator Example

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['United Kingdom'] := 'London';
    t_capital['Japan'] := 'Tokyo';

    t_capital := first(t_capital);
    WHILE NOT isnull(t_capital) LOOP
        RAISE NOTICE 'Capital of % is %', key(t_capital), value(t_capital);
        t_capital := next(t_capital);
    END LOOP;
END $$;
```

### Sparse Arrays (icollection)

`icollection` supports non-contiguous integer indices and distinguishes between NULL values and uninitialized keys:

```sql
DECLARE sparse icollection;
BEGIN
    sparse[1] := 'first';
    sparse[1000000] := 'millionth';  -- no memory wasted on gaps
END;
```
