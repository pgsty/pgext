---
title: "collection"
linkTitle: "collection"
description: "Memory optimized data type to be used inside of plpglsql func"
weight: 3630
categories: ["TYPE"]
width: full
---

Memory optimized data type to be used inside of plpglsql func


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3630** | {{< badge content="collection" link="https://github.com/aws/pgcollection" >}} | {{< ext "collection" "pgcollection" >}} | `1.1.0` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/collection" >}} | `1.1.0` | {{< bg "18" "pgcollection_18*" "green" >}} {{< bg "17" "pgcollection_17*" "green" >}} {{< bg "16" "pgcollection_16*" "green" >}} {{< bg "15" "pgcollection_15*" "green" >}} {{< bg "14" "pgcollection_14*" "green" >}} {{< bg "13" "pgcollection_13*" "red" >}} | `pgcollection_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/collection" >}} | `1.1.0` | {{< bg "18" "postgresql-18-collection" "green" >}} {{< bg "17" "postgresql-17-collection" "green" >}} {{< bg "16" "postgresql-16-collection" "green" >}} {{< bg "15" "postgresql-15-collection" "green" >}} {{< bg "14" "postgresql-14-collection" "green" >}} {{< bg "13" "postgresql-13-collection" "red" >}} | `postgresql-$v-collection` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    | {{< bg "PIGSTY 1.1.0" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1.0" "pgcollection_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pgcollection_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-collection : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-collection : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.0" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-collection : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-collection : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_18` | 1.1.0 | `el8.x86_64` | pigsty | 41.4 KiB | [pgcollection_18-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_18-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_18` | 1.1.0 | `el8.aarch64` | pigsty | 40.5 KiB | [pgcollection_18-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_18-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_18` | 1.1.0 | `el9.x86_64` | pigsty | 41.0 KiB | [pgcollection_18-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_18-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_18` | 1.1.0 | `el9.aarch64` | pigsty | 40.5 KiB | [pgcollection_18-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_18-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_18` | 1.1.0 | `el10.x86_64` | pigsty | 41.4 KiB | [pgcollection_18-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_18-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_18` | 1.1.0 | `el10.aarch64` | pigsty | 39.4 KiB | [pgcollection_18-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_18-1.1.0-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_17` | 1.1.0 | `el8.x86_64` | pigsty | 41.4 KiB | [pgcollection_17-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_17-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_17` | 1.0.0 | `el8.x86_64` | pigsty | 35.4 KiB | [pgcollection_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_17` | 1.1.0 | `el8.aarch64` | pigsty | 40.5 KiB | [pgcollection_17-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_17-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_17` | 1.0.0 | `el8.aarch64` | pigsty | 34.3 KiB | [pgcollection_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_17` | 1.1.0 | `el9.x86_64` | pigsty | 40.9 KiB | [pgcollection_17-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_17-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_17` | 1.0.0 | `el9.x86_64` | pigsty | 36.8 KiB | [pgcollection_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_17` | 1.1.0 | `el9.aarch64` | pigsty | 40.4 KiB | [pgcollection_17-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_17-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_17` | 1.0.0 | `el9.aarch64` | pigsty | 36.1 KiB | [pgcollection_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_17` | 1.1.0 | `el10.x86_64` | pigsty | 41.4 KiB | [pgcollection_17-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_17-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_17` | 1.1.0 | `el10.aarch64` | pigsty | 39.3 KiB | [pgcollection_17-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_17-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-collection` | 1.0.0 | `d12.x86_64` | pigsty | 70.1 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-collection` | 1.0.0 | `d12.aarch64` | pigsty | 68.8 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-collection` | 1.0.0 | `u22.x86_64` | pigsty | 86.1 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-collection` | 1.0.0 | `u22.aarch64` | pigsty | 86.3 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-collection` | 1.0.0 | `u24.x86_64` | pigsty | 74.2 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-collection` | 1.0.0 | `u24.aarch64` | pigsty | 73.1 KiB | [postgresql-17-collection_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_16` | 1.1.0 | `el8.x86_64` | pigsty | 41.2 KiB | [pgcollection_16-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_16-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_16` | 1.0.0 | `el8.x86_64` | pigsty | 35.3 KiB | [pgcollection_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_16` | 1.1.0 | `el8.aarch64` | pigsty | 40.3 KiB | [pgcollection_16-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_16-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_16` | 1.0.0 | `el8.aarch64` | pigsty | 34.2 KiB | [pgcollection_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_16` | 1.1.0 | `el9.x86_64` | pigsty | 40.8 KiB | [pgcollection_16-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_16-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_16` | 1.0.0 | `el9.x86_64` | pigsty | 36.5 KiB | [pgcollection_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_16` | 1.1.0 | `el9.aarch64` | pigsty | 40.2 KiB | [pgcollection_16-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_16-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_16` | 1.0.0 | `el9.aarch64` | pigsty | 35.8 KiB | [pgcollection_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_16` | 1.1.0 | `el10.x86_64` | pigsty | 41.3 KiB | [pgcollection_16-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_16-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_16` | 1.1.0 | `el10.aarch64` | pigsty | 39.1 KiB | [pgcollection_16-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_16-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-collection` | 1.0.0 | `d12.x86_64` | pigsty | 68.8 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-collection` | 1.0.0 | `d12.aarch64` | pigsty | 67.6 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-collection` | 1.0.0 | `u22.x86_64` | pigsty | 84.7 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-collection` | 1.0.0 | `u22.aarch64` | pigsty | 84.6 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-collection` | 1.0.0 | `u24.x86_64` | pigsty | 72.7 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-collection` | 1.0.0 | `u24.aarch64` | pigsty | 71.8 KiB | [postgresql-16-collection_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_15` | 1.1.0 | `el8.x86_64` | pigsty | 41.7 KiB | [pgcollection_15-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_15-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_15` | 1.0.0 | `el8.x86_64` | pigsty | 35.7 KiB | [pgcollection_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_15` | 1.1.0 | `el8.aarch64` | pigsty | 40.6 KiB | [pgcollection_15-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_15-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_15` | 1.0.0 | `el8.aarch64` | pigsty | 34.4 KiB | [pgcollection_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_15` | 1.1.0 | `el9.x86_64` | pigsty | 41.1 KiB | [pgcollection_15-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_15-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_15` | 1.0.0 | `el9.x86_64` | pigsty | 36.9 KiB | [pgcollection_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_15` | 1.1.0 | `el9.aarch64` | pigsty | 40.5 KiB | [pgcollection_15-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_15-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_15` | 1.0.0 | `el9.aarch64` | pigsty | 35.7 KiB | [pgcollection_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_15` | 1.1.0 | `el10.x86_64` | pigsty | 42.2 KiB | [pgcollection_15-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_15-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_15` | 1.1.0 | `el10.aarch64` | pigsty | 39.4 KiB | [pgcollection_15-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_15-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-collection` | 1.0.0 | `d12.x86_64` | pigsty | 69.3 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-collection` | 1.0.0 | `d12.aarch64` | pigsty | 67.9 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-collection` | 1.0.0 | `u22.x86_64` | pigsty | 85.4 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-collection` | 1.0.0 | `u22.aarch64` | pigsty | 84.8 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-collection` | 1.0.0 | `u24.x86_64` | pigsty | 72.9 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-collection` | 1.0.0 | `u24.aarch64` | pigsty | 71.5 KiB | [postgresql-15-collection_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_14` | 1.1.0 | `el8.x86_64` | pigsty | 41.7 KiB | [pgcollection_14-1.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_14-1.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_14` | 1.0.0 | `el8.x86_64` | pigsty | 35.6 KiB | [pgcollection_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pgcollection_14` | 1.1.0 | `el8.aarch64` | pigsty | 40.6 KiB | [pgcollection_14-1.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_14-1.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_14` | 1.0.0 | `el8.aarch64` | pigsty | 34.4 KiB | [pgcollection_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pgcollection_14` | 1.1.0 | `el9.x86_64` | pigsty | 41.0 KiB | [pgcollection_14-1.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_14-1.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_14` | 1.0.0 | `el9.x86_64` | pigsty | 36.9 KiB | [pgcollection_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pgcollection_14` | 1.1.0 | `el9.aarch64` | pigsty | 40.5 KiB | [pgcollection_14-1.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_14-1.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_14` | 1.0.0 | `el9.aarch64` | pigsty | 35.7 KiB | [pgcollection_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pgcollection_14` | 1.1.0 | `el10.x86_64` | pigsty | 41.6 KiB | [pgcollection_14-1.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_14-1.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgcollection_14` | 1.1.0 | `el10.aarch64` | pigsty | 39.4 KiB | [pgcollection_14-1.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_14-1.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-collection` | 1.0.0 | `d12.x86_64` | pigsty | 69.4 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-collection` | 1.0.0 | `d12.aarch64` | pigsty | 67.8 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-collection` | 1.0.0 | `u22.x86_64` | pigsty | 85.2 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-collection` | 1.0.0 | `u22.aarch64` | pigsty | 84.7 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-collection` | 1.0.0 | `u24.x86_64` | pigsty | 72.9 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-collection` | 1.0.0 | `u24.aarch64` | pigsty | 71.4 KiB | [postgresql-14-collection_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgcollection" title="Repository" icon="github" subtitle="github.com/aws/pgcollection" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcollection-1.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get collection; # get collection source code
pig build dep collection; # install build dependencies
pig build pkg collection; # build extension rpm or deb
pig build ext collection; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install collection; # install by extension name, for the current active PG version
pig ext install pgcollection; # install via package alias, for the active PG version
pig ext install collection -v 18;   # install for PG 18
pig ext install collection -v 17;   # install for PG 17
pig ext install collection -v 16;   # install for PG 16
pig ext install collection -v 15;   # install for PG 15
pig ext install collection -v 14;   # install for PG 14

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION collection;
```

