---
title: "collection"
linkTitle: "collection"
description: "Memory optimized data type to be used inside of plpglsql func"
weight: 3630
categories: ["TYPE"]
width: full
---

[**pgcollection**](https://github.com/aws/pgcollection) : Memory optimized data type to be used inside of plpglsql func


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3630** | {{< badge content="collection" link="https://github.com/aws/pgcollection" >}} | {{< ext "collection" "pgcollection" >}} | `1.1.1` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pgcollection` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "pgcollection_18" "green" >}} {{< bg "17" "pgcollection_17" "green" >}} {{< bg "16" "pgcollection_16" "green" >}} {{< bg "15" "pgcollection_15" "green" >}} {{< bg "14" "pgcollection_14" "green" >}} {{< bg "13" "pgcollection_13" "red" >}} | `pgcollection_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.1` | {{< bg "18" "postgresql-18-collection" "green" >}} {{< bg "17" "postgresql-17-collection" "green" >}} {{< bg "16" "postgresql-16-collection" "green" >}} {{< bg "15" "postgresql-15-collection" "green" >}} {{< bg "14" "postgresql-14-collection" "green" >}} {{< bg "13" "postgresql-13-collection" "red" >}} | `postgresql-$v-collection` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.1" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_18` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.1 KiB | [pgcollection_18-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_18-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_18` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.2 KiB | [pgcollection_18-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_18-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_18` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.7 KiB | [pgcollection_18-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_18-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_18` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [pgcollection_18-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_18-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_18` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.2 KiB | [pgcollection_18-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_18-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_18` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.9 KiB | [pgcollection_18-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_18-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-collection` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-collection` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.3 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-collection` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.8 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-collection` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.2 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-collection` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 94.3 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-collection` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 93.1 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-collection` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.2 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-collection` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.6 KiB | [postgresql-18-collection_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_17` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.1 KiB | [pgcollection_17-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_17-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_17` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.2 KiB | [pgcollection_17-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_17-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_17` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.7 KiB | [pgcollection_17-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_17-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_17` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [pgcollection_17-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_17-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_17` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.2 KiB | [pgcollection_17-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_17-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_17` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.9 KiB | [pgcollection_17-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_17-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-collection` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 84.3 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-collection` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 82.8 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-collection` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 84.8 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-collection` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 82.6 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-collection` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 104.0 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-collection` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 103.2 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-collection` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 89.2 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-collection` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 88.0 KiB | [postgresql-17-collection_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_16` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 42.9 KiB | [pgcollection_16-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_16-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_16` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.0 KiB | [pgcollection_16-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_16-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_16` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.4 KiB | [pgcollection_16-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_16-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_16` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.6 KiB | [pgcollection_16-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_16-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_16` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.1 KiB | [pgcollection_16-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_16-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_16` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.7 KiB | [pgcollection_16-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_16-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-collection` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 82.9 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-collection` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.5 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-collection` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.4 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-collection` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.5 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-collection` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 102.3 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-collection` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 101.4 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-collection` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.7 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-collection` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 86.6 KiB | [postgresql-16-collection_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.3 KiB | [pgcollection_15-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_15-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.3 KiB | [pgcollection_15-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_15-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.5 KiB | [pgcollection_15-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_15-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [pgcollection_15-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_15-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_15` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.3 KiB | [pgcollection_15-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_15-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_15` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.7 KiB | [pgcollection_15-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_15-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-collection` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.3 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-collection` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.9 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-collection` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.9 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-collection` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.7 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-collection` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 102.7 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-collection` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 102.0 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-collection` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.8 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-collection` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.0 KiB | [postgresql-15-collection_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 43.3 KiB | [pgcollection_14-1.1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_14-1.1.1-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 42.3 KiB | [pgcollection_14-1.1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_14-1.1.1-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 42.6 KiB | [pgcollection_14-1.1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_14-1.1.1-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 41.9 KiB | [pgcollection_14-1.1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_14-1.1.1-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_14` | `1.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.3 KiB | [pgcollection_14-1.1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_14-1.1.1-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_14` | `1.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 41.6 KiB | [pgcollection_14-1.1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_14-1.1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-collection` | `1.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 83.3 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-collection` | `1.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-collection` | `1.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 83.8 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-collection` | `1.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-collection` | `1.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 102.4 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-collection` | `1.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 101.8 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-collection` | `1.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 87.8 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-collection` | `1.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 87.1 KiB | [postgresql-14-collection_1.1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_1.1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgcollection" title="Repository" icon="github" subtitle="github.com/aws/pgcollection" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcollection-1.1.1.tar.gz" >}}
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
