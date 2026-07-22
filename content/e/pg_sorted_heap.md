---
title: "pg_sorted_heap"
linkTitle: "pg_sorted_heap"
description: "Sorted heap table AM with zone map scan pruning and built-in vector search"
weight: 2550
categories: ["OLAP"]
width: full
---

[**pg_sorted_heap**](https://github.com/skuznetsov/pg_sorted_heap) : Sorted heap table AM with zone map scan pruning and built-in vector search


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2550** | {{< badge content="pg_sorted_heap" link="https://github.com/skuznetsov/pg_sorted_heap" >}} | {{< ext "pg_sorted_heap" >}} | `0.14.0` | {{< category "OLAP" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "storage_engine" >}} {{< ext "pg_ivm" >}} {{< ext "pgvector" >}} {{< ext "vchord" >}} {{< ext "pg_search" >}} |

> [!Note] sorted_hnsw.shared_cache requires shared_preload_libraries=pg_sorted_heap.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.14.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_sorted_heap` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.14.0` | {{< bg "18" "pg_sorted_heap_18" "green" >}} {{< bg "17" "pg_sorted_heap_17" "green" >}} {{< bg "16" "pg_sorted_heap_16" "green" >}} {{< bg "15" "pg_sorted_heap_15" "red" >}} {{< bg "14" "pg_sorted_heap_14" "red" >}} | `pg_sorted_heap_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.14.0` | {{< bg "18" "postgresql-18-pg-sorted-heap" "green" >}} {{< bg "17" "postgresql-17-pg-sorted-heap" "green" >}} {{< bg "16" "postgresql-16-pg-sorted-heap" "green" >}} {{< bg "15" "postgresql-15-pg-sorted-heap" "red" >}} {{< bg "14" "postgresql-14-pg-sorted-heap" "red" >}} | `postgresql-$v-pg-sorted-heap` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "pg_sorted_heap_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_sorted_heap_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_sorted_heap_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-18-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-17-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.14.0" "postgresql-16-pg-sorted-heap : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-pg-sorted-heap : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-sorted-heap : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sorted_heap_18` | `0.14.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 216.9 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sorted_heap_18-0.14.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_sorted_heap_18` | `0.14.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 200.7 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sorted_heap_18-0.14.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_sorted_heap_18` | `0.14.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 208.6 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sorted_heap_18-0.14.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_sorted_heap_18` | `0.14.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 199.7 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sorted_heap_18-0.14.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_sorted_heap_18` | `0.14.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 213.4 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sorted_heap_18-0.14.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_sorted_heap_18` | `0.14.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 206.6 KiB | [pg_sorted_heap_18-0.14.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sorted_heap_18-0.14.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 724.3 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 711.3 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 726.4 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 715.2 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 758.4 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 754.3 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 744.9 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 741.9 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 746.4 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-sorted-heap` | `0.14.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 736.7 KiB | [postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-18-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sorted_heap_17` | `0.14.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 216.9 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sorted_heap_17-0.14.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_sorted_heap_17` | `0.14.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 200.6 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sorted_heap_17-0.14.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_sorted_heap_17` | `0.14.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 208.5 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sorted_heap_17-0.14.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_sorted_heap_17` | `0.14.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 199.7 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sorted_heap_17-0.14.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_sorted_heap_17` | `0.14.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 213.5 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sorted_heap_17-0.14.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_sorted_heap_17` | `0.14.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 206.2 KiB | [pg_sorted_heap_17-0.14.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sorted_heap_17-0.14.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 723.6 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 710.9 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 725.9 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 714.3 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 805.8 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 802.1 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 743.8 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 741.4 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 745.3 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-sorted-heap` | `0.14.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 736.0 KiB | [postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-17-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_sorted_heap_16` | `0.14.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 217.0 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_sorted_heap_16-0.14.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_sorted_heap_16` | `0.14.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 200.8 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_sorted_heap_16-0.14.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_sorted_heap_16` | `0.14.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 208.6 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_sorted_heap_16-0.14.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_sorted_heap_16` | `0.14.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 199.8 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_sorted_heap_16-0.14.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_sorted_heap_16` | `0.14.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 213.5 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_sorted_heap_16-0.14.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_sorted_heap_16` | `0.14.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 206.2 KiB | [pg_sorted_heap_16-0.14.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_sorted_heap_16-0.14.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 723.5 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 711.1 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 725.7 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 714.4 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 802.1 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 798.2 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 743.8 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 741.4 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 745.4 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-sorted-heap` | `0.14.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 735.6 KiB | [postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-sorted-heap/postgresql-16-pg-sorted-heap_0.14.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/skuznetsov/pg_sorted_heap" title="Repository" icon="github" subtitle="github.com/skuznetsov/pg_sorted_heap" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_sorted_heap-0.14.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_sorted_heap;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_sorted_heap;		# install via package name, for the active PG version

pig install pg_sorted_heap -v 18;   # install for PG 18
pig install pg_sorted_heap -v 17;   # install for PG 17
pig install pg_sorted_heap -v 16;   # install for PG 16

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_sorted_heap';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_sorted_heap;
```




## Usage

Sources: [pg_sorted_heap README](https://github.com/skuznetsov/pg_sorted_heap), [stable API](https://github.com/skuznetsov/pg_sorted_heap/blob/main/docs/api-stable.md), [SQL API](https://github.com/skuznetsov/pg_sorted_heap/blob/main/docs/api.md), [control file](https://github.com/skuznetsov/pg_sorted_heap/blob/main/pg_sorted_heap.control).

`pg_sorted_heap` adds the `sorted_heap` table access method, per-page zone-map pruning, maintenance helpers, built-in `svec`/`hsvec` vector types, a planner-integrated `sorted_hnsw` index AM, and stable GraphRAG wrappers. Upstream documents PostgreSQL 16, 17, and 18 support for the current release surface.

### Sorted Heap Tables

Use `USING sorted_heap` on tables with a primary key. Bulk loads are sorted by primary key on the COPY path, and compaction globally sorts existing rows while rebuilding the zone map:

```sql
CREATE EXTENSION pg_sorted_heap;

CREATE TABLE events (
  ts timestamptz,
  src text,
  data jsonb,
  PRIMARY KEY (ts, src)
) USING sorted_heap;

COPY events FROM '/path/to/events.csv';

SELECT sorted_heap_compact('events'::regclass);

EXPLAIN (ANALYZE, BUFFERS)
SELECT *
FROM events
WHERE ts BETWEEN '2026-01-01' AND '2026-01-02'
  AND src = 'sensor-42';
```

The README describes planner-injected `SortedHeapScan` paths for primary-key predicates and zone-map pruning at the heap-block level.

### Maintenance And Observability

Stable maintenance functions include:

```sql
SELECT sorted_heap_compact('events'::regclass);
CALL sorted_heap_compact_online('events'::regclass);

SELECT sorted_heap_merge('events'::regclass);
CALL sorted_heap_merge_online('events'::regclass);

SELECT sorted_heap_rebuild_zonemap('events'::regclass);
SELECT sorted_heap_zonemap_stats('events'::regclass);
```

Partition helpers operate on concrete sorted-heap leaves under a parent:

```sql
SELECT * FROM sorted_heap_partition_status('events_parent'::regclass);
SELECT * FROM sorted_heap_partition_maintenance_plan('events_parent'::regclass, 'compact');
SELECT * FROM sorted_heap_compact_partitions('events_parent'::regclass);
```

### Vector Search

The stable vector API includes `svec(dim)` for float32 vectors, `hsvec(dim)` for float16 vectors, and the `sorted_hnsw` index AM:

```sql
CREATE TABLE documents (
  id bigserial PRIMARY KEY,
  embedding svec(384),
  content text
);

CREATE INDEX documents_embedding_idx
ON documents USING sorted_hnsw (embedding)
WITH (m = 16, ef_construction = 200);

SET sorted_hnsw.ef_search = 96;

SELECT id, content
FROM documents
ORDER BY embedding <=> '[0.1,0.2,0.3]'::svec
LIMIT 10;
```

For compact base-table storage, use `hsvec` and the matching operator class:

```sql
CREATE TABLE documents_compact (
  id bigserial PRIMARY KEY,
  embedding hsvec(384),
  content text
);

CREATE INDEX documents_compact_embedding_idx
ON documents_compact USING sorted_hnsw (embedding hsvec_cosine_ops)
WITH (m = 16, ef_construction = 200);
```

The shared decoded graph cache is controlled by `sorted_hnsw.shared_cache`. Upstream examples note that using it requires preloading the extension:

```conf
shared_preload_libraries = 'pg_sorted_heap'
```

```sql
SET sorted_hnsw.shared_cache = on;
```

### GraphRAG

The stable fact-shaped GraphRAG entry point expects facts clustered by `(entity_id, relation_id, target_id)` or a registered alias mapping:

```sql
CREATE TABLE facts (
  entity_id int4,
  relation_id int2,
  target_id int4,
  embedding svec(384),
  payload text,
  PRIMARY KEY (entity_id, relation_id, target_id)
) USING sorted_heap;

CREATE INDEX facts_embedding_idx
ON facts USING sorted_hnsw (embedding)
WITH (m = 24, ef_construction = 200);

SET sorted_hnsw.ef_search = 128;

SELECT *
FROM sorted_heap_graph_rag(
  'facts'::regclass,
  '[0.1,0.2,0.3]'::svec,
  relation_path := ARRAY[1, 2],
  ann_k := 64,
  top_k := 10,
  score_mode := 'path'
);
```

Register alternate fact column names once:

```sql
SELECT sorted_heap_graph_register(
  'facts_alias'::regclass,
  entity_column := 'src_id',
  relation_column := 'edge_type',
  target_column := 'dst_id',
  embedding_column := 'vec',
  payload_column := 'body'
);
```

For routed or tenant-sharded fact tables, use `sorted_heap_graph_route(...)` and inspect routing with `sorted_heap_graph_route_plan(...)`.

### Stable GUCs

- `sorted_heap.enable_scan_pruning`: enable sorted-heap custom scan pruning; default `on`.
- `sorted_heap.vacuum_rebuild_zonemap`: rebuild zone maps during `VACUUM`; default `off`.
- `sorted_heap.lazy_update`: defer eager zone-map update maintenance; default `off`.
- `sorted_hnsw.ef_search`: runtime HNSW search breadth; default `64`.
- `sorted_hnsw.shared_cache`: shared decoded graph cache when preloaded; default `on`.
- `sorted_hnsw.sq8`: SQ8 decoded cache representation; default `on`.
- `sorted_hnsw.build_sq8`: low-memory index build mode; default `off`.

### Caveats

- `sorted_heap.lazy_update = on` trades scan pruning for faster update-heavy workloads until compaction or merge restores pruning.
- `sorted_hnsw.shared_cache` should be used with `shared_preload_libraries = 'pg_sorted_heap'`.
- Planner-integrated `sorted_hnsw` ordered scans require `LIMIT`; the SQL API says they are not chosen when there is no limit or when `LIMIT > sorted_hnsw.ef_search`.
- The lower-level GraphRAG and legacy/manual ANN helpers remain documented, but the stable application-facing API is the compact surface in `docs/api-stable.md`.
