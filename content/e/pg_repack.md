---
title: "pg_repack"
linkTitle: "pg_repack"
description: "Reorganize tables in PostgreSQL databases with minimal locks"
weight: 5010
categories: ["ADMIN"]
width: full
---

[**pg_repack**](https://github.com/reorg/pg_repack) : Reorganize tables in PostgreSQL databases with minimal locks


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5010** | {{< badge content="pg_repack" link="https://github.com/reorg/pg_repack" >}} | {{< ext "pg_repack" >}} | `1.5.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_rewrite" >}} {{< ext "pg_squeeze" >}} {{< ext "pgfincore" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pgstattuple" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_repack` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.3` | {{< bg "18" "pg_repack_18" "green" >}} {{< bg "17" "pg_repack_17" "green" >}} {{< bg "16" "pg_repack_16" "green" >}} {{< bg "15" "pg_repack_15" "green" >}} {{< bg "14" "pg_repack_14" "green" >}} | `pg_repack_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.3` | {{< bg "18" "postgresql-18-repack" "green" >}} {{< bg "17" "postgresql-17-repack" "green" >}} {{< bg "16" "postgresql-16-repack" "green" >}} {{< bg "15" "postgresql-15-repack" "green" >}} {{< bg "14" "postgresql-14-repack" "green" >}} | `postgresql-$v-repack` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 5" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5.3" "pg_repack_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.3" "pg_repack_14 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.5.3" "postgresql-18-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-17-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-16-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-15-repack : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5.3" "postgresql-14-repack : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_18` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.4 KiB | [pg_repack_18-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_18` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.8 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_repack_18-1.5.2-5PGDG.rhel8.x86_64.rpm) |
| `pg_repack_18` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.3 KiB | [pg_repack_18-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_18` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_repack_18-1.5.2-5PGDG.rhel8.aarch64.rpm) |
| `pg_repack_18` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.9 KiB | [pg_repack_18-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_18` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.2 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_repack_18-1.5.2-5PGDG.rhel9.x86_64.rpm) |
| `pg_repack_18` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.3 KiB | [pg_repack_18-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_18` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.5 KiB | [pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_repack_18-1.5.2-5PGDG.rhel9.aarch64.rpm) |
| `pg_repack_18` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.9 KiB | [pg_repack_18-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_repack_18-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_18` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.1 KiB | [pg_repack_18-1.5.2-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_repack_18-1.5.2-5PGDG.rhel10.x86_64.rpm) |
| `pg_repack_18` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.8 KiB | [pg_repack_18-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_repack_18-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_18` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.1 KiB | [pg_repack_18-1.5.2-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_repack_18-1.5.2-5PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-repack` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.9 KiB | [postgresql-18-repack_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-repack` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.3 KiB | [postgresql-18-repack_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-repack` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 102.2 KiB | [postgresql-18-repack_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-repack` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.3 KiB | [postgresql-18-repack_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 100.9 KiB | [postgresql-18-repack_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 97.6 KiB | [postgresql-18-repack_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 98.9 KiB | [postgresql-18-repack_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.8 KiB | [postgresql-18-repack_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 99.4 KiB | [postgresql-18-repack_1.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-repack` | `1.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 97.2 KiB | [postgresql-18-repack_1.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-18-repack_1.5.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_17` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.7 KiB | [pg_repack_17-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.3 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.9 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_repack_17-1.5.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_17` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.5 KiB | [pg_repack_17-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.6 KiB | [pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_repack_17-1.5.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_17` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.0 KiB | [pg_repack_17-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.8 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.7 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_repack_17-1.5.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_17` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.3 KiB | [pg_repack_17-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.2 KiB | [pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.1 KiB | [pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_repack_17-1.5.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_17` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [pg_repack_17-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_repack_17-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_17` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.2 KiB | [pg_repack_17-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_repack_17-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_17` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.9 KiB | [pg_repack_17-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_repack_17-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_17` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.1 KiB | [pg_repack_17-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_repack_17-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-repack` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.9 KiB | [postgresql-17-repack_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-repack` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.4 KiB | [postgresql-17-repack_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-repack` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 102.3 KiB | [postgresql-17-repack_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-repack` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.3 KiB | [postgresql-17-repack_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.0 KiB | [postgresql-17-repack_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.5 KiB | [postgresql-17-repack_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.1 KiB | [postgresql-17-repack_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.8 KiB | [postgresql-17-repack_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 99.4 KiB | [postgresql-17-repack_1.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-repack` | `1.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 97.3 KiB | [postgresql-17-repack_1.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-17-repack_1.5.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_16` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.6 KiB | [pg_repack_16-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.3 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.5 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_repack_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_16` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.4 KiB | [pg_repack_16-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.0 KiB | [pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 79.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_repack_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_16` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.9 KiB | [pg_repack_16-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.8 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 68.0 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_repack_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_16` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.2 KiB | [pg_repack_16-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.2 KiB | [pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.4 KiB | [pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_repack_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_16` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [pg_repack_16-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_repack_16-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_16` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.1 KiB | [pg_repack_16-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_repack_16-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_16` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.8 KiB | [pg_repack_16-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_repack_16-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_16` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.0 KiB | [pg_repack_16-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_repack_16-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-repack` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 102.1 KiB | [postgresql-16-repack_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-repack` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.6 KiB | [postgresql-16-repack_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-repack` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 102.6 KiB | [postgresql-16-repack_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-repack` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.6 KiB | [postgresql-16-repack_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.9 KiB | [postgresql-16-repack_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.8 KiB | [postgresql-16-repack_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.2 KiB | [postgresql-16-repack_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 97.0 KiB | [postgresql-16-repack_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 99.5 KiB | [postgresql-16-repack_1.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-repack` | `1.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 97.9 KiB | [postgresql-16-repack_1.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-16-repack_1.5.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_15` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 76.8 KiB | [pg_repack_15-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.6 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 79.6 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_15` | `1.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.6 KiB | [pg_repack_15-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_repack_15-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_15` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 75.7 KiB | [pg_repack_15-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.3 KiB | [pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 79.2 KiB | [pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_15` | `1.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.0 KiB | [pg_repack_15-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_repack_15-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_15` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.3 KiB | [pg_repack_15-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.2 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 68.3 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_15` | `1.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 96.1 KiB | [pg_repack_15-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_repack_15-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_15` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.7 KiB | [pg_repack_15-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.7 KiB | [pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.7 KiB | [pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_15` | `1.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 95.0 KiB | [pg_repack_15-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_repack_15-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_15` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.3 KiB | [pg_repack_15-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_repack_15-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_15` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.5 KiB | [pg_repack_15-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_repack_15-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_15` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 68.2 KiB | [pg_repack_15-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_repack_15-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_15` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.5 KiB | [pg_repack_15-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_repack_15-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-repack` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 102.1 KiB | [postgresql-15-repack_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-repack` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.8 KiB | [postgresql-15-repack_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-repack` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 102.5 KiB | [postgresql-15-repack_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-repack` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.9 KiB | [postgresql-15-repack_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 105.6 KiB | [postgresql-15-repack_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 102.7 KiB | [postgresql-15-repack_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.5 KiB | [postgresql-15-repack_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 97.2 KiB | [postgresql-15-repack_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 99.9 KiB | [postgresql-15-repack_1.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-repack` | `1.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 97.8 KiB | [postgresql-15-repack_1.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-15-repack_1.5.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_repack_14` | `1.5.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 75.1 KiB | [pg_repack_14-1.5.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 73.9 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 78.0 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_repack_14` | `1.4.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.8-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | `1.4.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.7 KiB | [pg_repack_14-1.4.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_repack_14-1.4.7-1.rhel8.x86_64.rpm) |
| `pg_repack_14` | `1.5.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.3 KiB | [pg_repack_14-1.5.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | `1.5.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 77.8 KiB | [pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_repack_14` | `1.4.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.7 KiB | [pg_repack_14-1.4.8-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_repack_14-1.4.8-1.rhel8.aarch64.rpm) |
| `pg_repack_14` | `1.5.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 66.3 KiB | [pg_repack_14-1.5.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.3-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | `1.5.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.3 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 67.6 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_repack_14` | `1.4.8` | [el9.x86_64](/os/el9.x86_64) | pgdg | 94.2 KiB | [pg_repack_14-1.4.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.8-1.rhel9.x86_64.rpm) |
| `pg_repack_14` | `1.4.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 93.7 KiB | [pg_repack_14-1.4.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_repack_14-1.4.7-1.rhel9.x86_64.rpm) |
| `pg_repack_14` | `1.5.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 66.0 KiB | [pg_repack_14-1.5.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.3-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | `1.5.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 65.0 KiB | [pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 67.1 KiB | [pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_repack_14` | `1.4.8` | [el9.aarch64](/os/el9.aarch64) | pgdg | 93.4 KiB | [pg_repack_14-1.4.8-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_repack_14-1.4.8-1.rhel9.aarch64.rpm) |
| `pg_repack_14` | `1.5.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.6 KiB | [pg_repack_14-1.5.3-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_repack_14-1.5.3-1PGDG.rhel10.x86_64.rpm) |
| `pg_repack_14` | `1.5.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 66.8 KiB | [pg_repack_14-1.5.2-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_repack_14-1.5.2-4PGDG.rhel10.x86_64.rpm) |
| `pg_repack_14` | `1.5.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 67.6 KiB | [pg_repack_14-1.5.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_repack_14-1.5.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_repack_14` | `1.5.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 66.9 KiB | [pg_repack_14-1.5.2-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_repack_14-1.5.2-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-repack` | `1.5.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 101.7 KiB | [postgresql-14-repack_1.5.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-repack` | `1.5.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 99.5 KiB | [postgresql-14-repack_1.5.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-repack` | `1.5.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 101.8 KiB | [postgresql-14-repack_1.5.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-repack` | `1.5.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 100.3 KiB | [postgresql-14-repack_1.5.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 104.2 KiB | [postgresql-14-repack_1.5.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 101.2 KiB | [postgresql-14-repack_1.5.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 99.0 KiB | [postgresql-14-repack_1.5.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 96.8 KiB | [postgresql-14-repack_1.5.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u26.x86_64](/os/u26.x86_64) | pgdg | 98.9 KiB | [postgresql-14-repack_1.5.3-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-repack` | `1.5.3` | [u26.aarch64](/os/u26.aarch64) | pgdg | 97.6 KiB | [postgresql-14-repack_1.5.3-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-repack/postgresql-14-repack_1.5.3-1.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/reorg/pg_repack" title="Repository" icon="github" subtitle="github.com/reorg/pg_repack" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_repack;		# install via package name, for the active PG version

pig install pg_repack -v 18;   # install for PG 18
pig install pg_repack -v 17;   # install for PG 17
pig install pg_repack -v 16;   # install for PG 16
pig install pg_repack -v 15;   # install for PG 15
pig install pg_repack -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_repack;
```



## About

- GitHub Repo: [`reorg/pg_repack`](https://github.com/reorg/pg_repack)
- [**HomePage**](https://reorg.github.io/pg_repack/) 

pg_repack is a PostgreSQL extension which lets you remove bloat from
tables and indexes, and optionally restore the physical order of clustered
indexes. Unlike [CLUSTER](https://www.postgresql.org/docs/current/static/sql-cluster.html)
and [VACUUM FULL](https://www.postgresql.org/docs/current/static/sql-vacuum.html)
it works online, without holding an exclusive lock on the processed tables
during processing. pg_repack is efficient to boot, with performance comparable
to using CLUSTER directly.

pg_repack is a fork of the previous [pg_reorg](https://github.com/reorg/pg_reorg) project.

You can choose one of the following methods to reorganize:

- Online CLUSTER (ordered by cluster index)
- Ordered by specified columns
- Online VACUUM FULL (packing rows only)
- Rebuild or relocate only the indexes of a table

NOTICE:

- Only superusers or owners of tables and indexes can use the utility. To run  pg_repack as an owner you need to use the option `--no-superuser-check`.
- Target table must have a PRIMARY KEY, or at least a UNIQUE total index on a NOT NULL column.

### What about pg_reorg?

pg_repack is a fork of the [pg_reorg](https://github.com/reorg/pg_reorg) project, which has proven hugely
successful. Unfortunately new feature development on pg_reorg has slowed or stopped since late 2011.

pg_repack was initially released as a drop-in replacement for pg_reorg,
addressing some of the shortcomings of the last pg_reorg version (such as support for PostgreSQL 9.2 and EXTENSION packaging) and known bugs.

pg_repack 1.2 introduces further new features (parallel index builds,
ability to rebuild only indexes) and bugfixes. In some cases its behaviour
may be different from the 1.1.x release so it shouldn't be considered a
drop-in replacement: you are advised to check the documentation before
upgrading from previous versions.


## Requirements

PostgreSQL versions
: PostgreSQL 9.5, 9.6, 10, 11, 12, 13, 14, 15, 16, 17, 18. PostgreSQL 9.4 and before it are not supported.

Disks
: Performing a full-table repack requires free disk space about twice as large as the target table(s) and its indexes. For example, if the total
size of the tables and indexes to be reorganized is 1GB, an additional 2GB of disk space is required.


## Download & Installation

You can [download pg_repack](http://pgxn.org/dist/pg_repack/) from the PGXN website. Unpack the archive and follow the installation instructions below.

Alternatively you can use the [PGXN Client](https://pgxn.github.io/pgxnclient/) to download, compile and install the package; use:

```bash
$ pgxn install pg_repack
```

pg_repack can be built with `make` on UNIX or Linux. The PGXS build
framework is used automatically. Before building, you might need to install
the PostgreSQL development packages (`postgresql-devel`, etc.) and add the
directory containing `pg_config` to your `$PATH`. Then you can run:

```bash
$ cd pg_repack
$ make
$ sudo make install
```

You can also use Microsoft Visual C++ 2010 to build the program on Windows.
There are project files in the `msvc` folder.

After installation, load the pg_repack extension in the database you want to
process. pg_repack is packaged as an extension, so you can execute:

```bash
$ psql -c "CREATE EXTENSION pg_repack" -d your_database
```

You can remove pg_repack using `DROP EXTENSION pg_repack` or just dropping
the `repack` schema.

If you are upgrading from a previous version of pg_repack or pg_reorg, just
drop the old version from the database as explained above and install the new
version.


## Usage

```
pg_repack [OPTION]... [DBNAME]
```

The following options can be specified in `OPTIONS`.

```
Options:
  -a, --all                          repack all databases
  -t, --table=TABLE                  repack specific table only
  -I, --parent-table=TABLE           repack specific parent table and its inheritors
  -c, --schema=SCHEMA                repack tables in specific schema only
  -s, --tablespace=TBLSPC            move repacked tables to a new tablespace
  -S, --moveidx                      move repacked indexes to TBLSPC too
  -o, --order-by=COLUMNS             order by columns instead of cluster keys
  -n, --no-order                     do vacuum full instead of cluster
  -N, --dry-run                      print what would have been repacked and exit
  -j, --jobs=NUM                     Use this many parallel jobs for each table
  -i, --index=INDEX                  move only the specified index
  -x, --only-indexes                 move only indexes of the specified table
  -T, --wait-timeout=SECS            timeout to cancel other backends on conflict
  -D, --no-kill-backend              don't kill other backends when timed out
  -Z, --no-analyze                   don't analyze at end
  -k, --no-superuser-check           skip superuser checks in client
  -C, --exclude-extension            don't repack tables which belong to specific extension
      --no-error-on-invalid-index    repack even though invalid index is found
      --error-on-invalid-index       don't repack when invalid index is found (default)
      --apply-count                  number of tuples to apply in one transaction during replay
      --switch-threshold             switch tables when that many tuples are left to catchup

Connection options:
  -d, --dbname=DBNAME                database to connect
  -h, --host=HOSTNAME                database server host or socket directory
  -p, --port=PORT                    database server port
  -U, --username=USERNAME            user name to connect as
  -w, --no-password                  never prompt for password
  -W, --password                     force password prompt

Generic options:
  -e, --echo                         echo queries
  -E, --elevel=LEVEL                 set output message level
  --help                             show this help, then exit
  --version                          output version information, then exit
```


## Options

### Reorg Options

`-a`, `--all`
: Attempt to repack all the databases of the cluster. Databases where the
`pg_repack` extension is not installed will be skipped.

`-t TABLE`, `--table=TABLE`
: Reorganize the specified table(s) only. Multiple tables may be
reorganized by writing multiple `-t` switches. By default, all eligible
tables in the target databases are reorganized.

`-I TABLE`, `--parent-table=TABLE`
: Reorganize both the specified table(s) and its inheritors. Multiple
table hierarchies may be reorganized by writing multiple `-I` switches.

`-c`, `--schema`
: Repack the tables in the specified schema(s) only. Multiple schemas may
be repacked by writing multiple `-c` switches. May be used in
conjunction with `--tablespace` to move tables to a different tablespace.

`-o COLUMNS [,...]`, `--order-by=COLUMNS [,...]`
: Perform an online CLUSTER ordered by the specified columns.

`-n`, `--no-order`
: Perform an online VACUUM FULL. Since version 1.2 this is the default for
non-clustered tables.

`-N`, `--dry-run`
: List what would be repacked and exit.

`-j`, `--jobs`
: Create the specified number of extra connections to PostgreSQL, and
use these extra connections to parallelize the rebuild of indexes
on each table. Parallel index builds are only supported for full-table
repacks, not with `--index` or `--only-indexes` options. If your
PostgreSQL server has extra cores and disk I/O available, this can be a
useful way to speed up pg_repack.

`-s TBLSPC`, `--tablespace=TBLSPC`
: Move the repacked tables to the specified tablespace: essentially an
online version of `ALTER TABLE ... SET TABLESPACE`. The tables' indexes
are left in the original tablespace unless `--moveidx` is specified too.

`-S`, `--moveidx`
: Also move the indexes of the repacked tables to the tablespace specified
by the `--tablespace` option.

`-i`, `--index`
: Repack the specified index(es) only. Multiple indexes may be repacked
by writing multiple `-i` switches. May be used in conjunction with
`--tablespace` to move the index to a different tablespace.

`-x`, `--only-indexes`
: Repack only the indexes of the specified table(s), which must be specified
with the `--table` or `--parent-table` options.

`-T SECS`, `--wait-timeout=SECS`
: pg_repack needs to take one exclusive lock at the beginning as well as one
exclusive lock at the end of the repacking process. This setting controls
how many seconds pg_repack will wait to acquire this lock. If the lock
cannot be taken after this duration and `--no-kill-backend` option is
not specified, pg_repack will forcibly cancel the conflicting queries.
If you are using PostgreSQL version 8.4 or newer, pg_repack will fall
back to using `pg_terminate_backend()` to disconnect any remaining
backends after twice this timeout has passed.
The default is 60 seconds.

`-D`, `--no-kill-backend`
: Skip to repack table if the lock cannot be taken for duration specified
`--wait-timeout`, instead of cancelling conflicting queries. The default
is false.

`-Z`, `--no-analyze`
: Disable ANALYZE after a full-table reorganization. If not specified, run
ANALYZE after the reorganization.

`-k`, `--no-superuser-check`
: Skip the superuser checks in the client. This setting is useful for using
pg_repack on platforms that support running it as non-superusers.

`-C`, `--exclude-extension`
: Skip tables that belong to the specified extension(s). Some extensions
may heavily depend on such tables at planning time etc.

`--switch-threshold`
: Switch tables when that many tuples are left in log table.
This setting can be used to avoid the inability to catchup with write-heavy tables.


### Connection Options

Options to connect to servers. You cannot use `--all` and `--dbname` or
`--table` or `--parent-table` together.

`-d DBNAME`, `--dbname=DBNAME`
: Specifies the name of the database to be reorganized. If this is not
specified and `-a` (or `--all`) is not used, the database name is read
from the environment variable `PGDATABASE`. If that is not set, the user
name specified for the connection is used.

`-h HOSTNAME`, `--host=HOSTNAME`
: Specifies the host name of the machine on which the server is running. If
the value begins with a slash, it is used as the directory for the Unix
domain socket.

`-p PORT`, `--port=PORT`
: Specifies the TCP port or local Unix domain socket file extension on which
the server is listening for connections.

`-U USERNAME`, `--username=USERNAME`
: User name to connect as.

`-w`, `--no-password`
: Never issue a password prompt. If the server requires password
authentication and a password is not available by other means such as a
`.pgpass` file, the connection attempt will fail. This option can be
useful in batch jobs and scripts where no user is present to enter a
password.

`-W`, `--password`
: Force the program to prompt for a password before connecting to a
database. This option is never essential, since the program will automatically
prompt for a password if the server demands password authentication.
However, pg_repack will waste a connection attempt finding out that the
server wants a password. In some cases it is worth typing `-W` to avoid
the extra connection attempt.


### Generic Options

`-e`, `--echo`
: Echo commands sent to server.

`-E LEVEL`, `--elevel=LEVEL`
: Choose the output message level from `DEBUG`, `INFO`, `NOTICE`,
`WARNING`, `ERROR`, `LOG`, `FATAL`, and `PANIC`. The default is
`INFO`.

`--help`
: Show usage of the program.

`--version`
: Show the version number of the program.


## Environment

`PGDATABASE`, `PGHOST`, `PGPORT`, `PGUSER`
: Default connection parameters.

This utility, like most other PostgreSQL utilities, also uses the
environment variables supported by libpq (see [Environment Variables](http://www.postgresql.org/docs/current/static/libpq-envars.html)).


## Examples

Perform an online CLUSTER of all the clustered tables in the database
`test`, and perform an online VACUUM FULL of all the non-clustered tables:

```bash
$ pg_repack test
```

Perform an online VACUUM FULL on the tables `foo` and `bar` in the
database `test` (an eventual cluster index is ignored):

```bash
$ pg_repack --no-order --table foo --table bar test
```

Move all indexes of table `foo` to tablespace `tbs`:

```bash
$ pg_repack -d test --table foo --only-indexes --tablespace tbs
```

Move the specified index to tablespace `tbs`:

```bash
$ pg_repack -d test --index idx --tablespace tbs
```


## Diagnostics

Error messages are reported when pg_repack fails. The following list shows the
cause of errors.

You need to cleanup by hand after fatal errors. To cleanup, just remove
pg_repack from the database and install it again: execute
`DROP EXTENSION pg_repack CASCADE` in the database where the error occurred,
followed by `CREATE EXTENSION pg_repack`.

**INFO: database "db" skipped: pg_repack VER is not installed in the database**
: pg_repack is not installed in the database when the `--all` option is
specified. Create the pg_repack extension in the database.

**ERROR: pg_repack VER is not installed in the database**
: pg_repack is not installed in the database specified by `--dbname`.
Create the pg_repack extension in the database.

**ERROR: program 'pg_repack V1' does not match database library 'pg_repack V2'**
: There is a mismatch between the `pg_repack` binary and the database
library (`.so` or `.dll`). The mismatch could be due to the wrong binary in the `$PATH` or the
wrong database being addressed. Check the program directory and the
database; if they are what expected you may need to repeat pg_repack
installation.

**ERROR: extension 'pg_repack V1' required, found 'pg_repack V2'**
: The SQL extension found in the database does not match the version
required by the pg_repack program. You should drop the extension from the database and reload it.

**ERROR: relation "table" must have a primary key or not-null unique keys**
: The target table doesn't have a PRIMARY KEY or any UNIQUE constraints
defined. Define a PRIMARY KEY or a UNIQUE constraint on the table.

**ERROR: query failed: ERROR: column "col" does not exist**
: The target table doesn't have columns specified by `--order-by` option.
Specify existing columns.

**WARNING: the table "tbl" already has a trigger called repack_trigger**
: The trigger was probably installed during a previous attempt to run
pg_repack on the table which was interrupted and for some reason failed
to clean up the temporary objects. You can remove all the temporary objects by dropping and re-creating the
extension.

**ERROR: Another pg_repack command may be running on the table. Please try again later.**
: There is a chance of deadlock when two concurrent pg_repack commands are
run on the same table. So, try to run the command after some time.

**WARNING: Cannot create index "schema"."index_xxxxx", already exists**
: DETAIL: An invalid index may have been left behind by a previous pg_repack
on the table which was interrupted. Please use `DROP INDEX "schema"."index_xxxxx"`
to remove this index and try again. A temporary index apparently created by pg_repack has been left behind, and
we do not want to risk dropping this index ourselves. If the index was in
fact created by an old pg_repack job which didn't get cleaned up, you
should just use DROP INDEX and try the repack command again.


## Restrictions

### Temp tables

pg_repack cannot reorganize temp tables.

### GiST indexes

pg_repack cannot cluster tables by GiST indexes.

### DDL commands

You will not be able to perform DDL commands of the target table(s) while
pg_repack is working. pg_repack will hold a SHARE UPDATE EXCLUSIVE lock on the
target table during a full-table repack, to enforce this restriction.

If you are using version 1.1.8 or earlier, you must not attempt to perform any
DDL commands on the target table(s) while pg_repack is running. In many cases
pg_repack would fail and rollback correctly, but there were some cases in these
earlier versions which could result in data corruption.


## Details

### Full Table Repacks

To perform a full-table repack, pg_repack will:

1. create a log table to record changes made to the original table
2. add a trigger onto the original table, logging INSERTs, UPDATEs and DELETEs into our log table
3. create a new table containing all the rows in the old table
4. build indexes on this new table
5. apply all changes which have accrued in the log table to the new table
6. swap the tables, including indexes and toast tables, using the system catalogs
7. drop the original table

pg_repack will only hold an ACCESS EXCLUSIVE lock for a short period during
initial setup (steps 1 and 2 above) and during the final swap-and-drop phase
(steps 6 and 7). For the rest of its time, pg_repack only needs
to hold a SHARE UPDATE EXCLUSIVE lock on the original table, meaning INSERTs,
UPDATEs, and DELETEs may proceed as usual.


### Index Only Repacks

To perform an index-only repack, pg_repack will:

1. create new indexes on the table using CONCURRENTLY matching the definitions of the old indexes
2. swap out the old for the new indexes in the catalogs
3. drop the old indexes

Creating indexes concurrently comes with a few caveats, please see
[the PostgreSQL documentation](http://www.postgresql.org/docs/current/static/sql-createindex.html#SQL-CREATEINDEX-CONCURRENTLY) for details.


## Releases

- **pg_repack 1.5.3**
  - Cleanup temporary indexes on error when using `--only-indexes` to avoid artefacts after failures
  - Acquire SHARE UPDATE EXCLUSIVE lock on the target table during a full-table repack instead of ACCESS SHARE
  - Fix subtransactions cache overflow

- **pg_repack 1.5.2**
  - Allow to run pg_repack by non-superusers
  - Make `--error-on-invalid-index` the default behavior

- **pg_repack 1.5.1**
  - Added support for PostgreSQL 17
  - Fix wrong OID format type in `repack_trigger`
  - Fix check of NOT NULL by `repack.primary_keys`
  - Fixed processing of tablespace names requiring quoted identifiers
  - Replace `PQconnectdb()` by `PQconnectdbParams()`
  - Added `--apply-count` option
  - Do not include a declaratively partitioned table with option `--only-indexes`
  - Fix possible two vacuums concurrently processing the same relfilenode
  - Use savepoints when retrying to take AccessShareLock
  - Fix swap of relfrozenxid, relfrozenxid and relallvisible

- **pg_repack 1.5.0**
  - Added support for PostgreSQL 16
  - Fix possible SQL injection
  - Support longer password length
  - Fixed infinite loop on empty password
  - Added `--switch-threshold` option
  - Fixed crash in `get_order_by()` using invalid relations
  - Added support for tables that have been previously rewritten with VACUUM FULL and use `storage=plain` for all columns
  - More careful locks acquisition

- **pg_repack 1.4.8**
  - Added support for PostgreSQL 15
  - Fixed `--parent-table` on declarative partitioned tables
  - Removed connection info from error log

- **pg_repack 1.4.7**
  - Added support for PostgreSQL 14

- **pg_repack 1.4.6**
  - Added support for PostgreSQL 13
  - Dropped support for PostgreSQL before 9.4

- **pg_repack 1.4.5**
  - Added support for PostgreSQL 12
  - Fixed parallel processing for indexes with operators from public schema

- **pg_repack 1.4.4**
  - Added support for PostgreSQL 11
  - Remove duplicate password prompt

- **pg_repack 1.4.3**
  - Fixed possible CVE-2018-1058 attack paths
  - Fixed "unexpected index definition" after CVE-2018-1058 changes in PostgreSQL
  - Fixed build with recent Ubuntu packages

- **pg_repack 1.4.2**
  - Added PostgreSQL 10 support
  - Fixed error DROP INDEX CONCURRENTLY cannot run inside a transaction block

- **pg_repack 1.4.1**
  - Fixed broken `--order-by` option

- **pg_repack 1.4**
  - Added support for PostgreSQL 9.6, dropped support for versions before 9.1
  - Use `AFTER` trigger to solve concurrency problems with `INSERT CONFLICT`
  - Added `--no-kill-backend` option
  - Added `--no-superuser-check` option
  - Added `--exclude-extension` option
  - Added `--parent-table` option
  - Restore TOAST storage parameters on repacked tables
  - Restore columns storage types in repacked tables

- **pg_repack 1.3.4**
  - Grab exclusive lock before dropping original table
  - Do not attempt to repack unlogged tables

- **pg_repack 1.3.3**
  - Added support for PostgreSQL 9.5
  - Fixed possible deadlock when pg_repack command is interrupted
  - Fixed exit code for when pg_repack is invoked with `--help` and `--version`
  - Added Japanese language user manual

- **pg_repack 1.3.2**
  - Fixed to clean up temporary objects when pg_repack command is interrupted
  - Fixed possible crash when pg_repack shared library is loaded alongside pg_statsinfo

- **pg_repack 1.3.1**
  - Added support for PostgreSQL 9.4

- **pg_repack 1.3**
  - Added `--schema` to repack only the specified schema
  - Added `--dry-run` to do a dry run
  - Fixed advisory locking for >2B OID values
  - Avoid possible deadlock when other sessions lock a to-be-repacked table
  - Performance improvement for performing sql_pop DELETEs many-at-a-time
  - Attempt to avoid pg_repack taking forever when dealing with a constant heavy stream of changes to a table

- **pg_repack 1.2**
  - Support PostgreSQL 9.3
  - Added `--tablespace` and `--moveidx` options to perform online SET TABLESPACE
  - Added `--index` to repack the specified index only
  - Added `--only-indexes` to repack only the indexes of the specified table
  - Added `--jobs` option for parallel operation
  - Don't require `--no-order` to perform a VACUUM FULL on non-clustered tables
  - Don't wait for locks held in other databases
  - Bugfix: correctly handle key indexes with options such as DESC, NULL FIRST/LAST, COLLATE
  - Fixed data corruption bug on delete
  - More helpful program output and error messages

- **pg_repack 1.1.8**
  - Added support for PostgreSQL 9.2
  - Added support for CREATE EXTENSION on PostgreSQL 9.1 and following
  - Give user feedback while waiting for transactions to finish
  - Bugfix: Allow running on newly promoted streaming replication slaves
  - Bugfix: Fix interaction between pg_repack and Slony 2.0/2.1
  - Bugfix: Properly escape column names
  - Bugfix: Avoid recreating invalid indexes, or choosing them as key
  - Bugfix: Never choose a partial index as primary key

- **pg_reorg 1.1.7** (2011-08-07)
  - Bugfix: VIEWs and FUNCTIONs could be corrupted that used a reorganized table which has a dropped column
  - Supports PostgreSQL 9.1 and 9.2dev (but EXTENSION is not yet)


## See Also

- [clusterdb](http://www.postgresql.org/docs/current/static/app-clusterdb.html)
- [vacuumdb](http://www.postgresql.org/docs/current/static/app-vacuumdb.html)
