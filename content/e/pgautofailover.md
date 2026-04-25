---
title: "pgautofailover"
linkTitle: "pgautofailover"
description: "pg_auto_failover"
weight: 5150
categories: ["ADMIN"]
width: full
---

[**pgautofailover**](https://github.com/hapostgres/pg_auto_failover) : pg_auto_failover


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5150** | {{< badge content="pgautofailover" link="https://github.com/hapostgres/pg_auto_failover" >}} | {{< ext "pgautofailover" >}} | `2.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "pglogical_origin" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgpool_recovery" >}} {{< ext "repmgr" >}} {{< ext "pg_checksums" >}} {{< ext "pgpool_adm" >}} {{< ext "bgw_replstatus" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgautofailover` | `btree_gist` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "pg_auto_failover_18" "red" >}} {{< bg "17" "pg_auto_failover_17" "green" >}} {{< bg "16" "pg_auto_failover_16" "green" >}} {{< bg "15" "pg_auto_failover_15" "green" >}} {{< bg "14" "pg_auto_failover_14" "green" >}} | `pg_auto_failover_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.2` | {{< bg "18" "postgresql-18-auto-failover" "green" >}} {{< bg "17" "postgresql-17-auto-failover" "green" >}} {{< bg "16" "postgresql-16-auto-failover" "green" >}} {{< bg "15" "postgresql-15-auto-failover" "green" >}} {{< bg "14" "postgresql-14-auto-failover" "green" >}} | `postgresql-$v-auto-failover` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 5" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.2" "postgresql-18-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 377.7 KiB | [postgresql-18-auto-failover_2.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg12+1_amd64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 373.6 KiB | [postgresql-18-auto-failover_2.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg12+1_arm64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 379.1 KiB | [postgresql-18-auto-failover_2.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg13+1_amd64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 374.8 KiB | [postgresql-18-auto-failover_2.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg13+1_arm64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 378.3 KiB | [postgresql-18-auto-failover_2.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 374.1 KiB | [postgresql-18-auto-failover_2.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 369.2 KiB | [postgresql-18-auto-failover_2.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 364.7 KiB | [postgresql-18-auto-failover_2.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 368.7 KiB | [postgresql-18-auto-failover_2.2-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-auto-failover` | `2.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 364.4 KiB | [postgresql-18-auto-failover_2.2-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-18-auto-failover_2.2-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_17` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.9 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 809.9 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 786.3 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.2 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 788.0 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 785.7 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 374.2 KiB | [postgresql-17-auto-failover_2.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg12+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 370.7 KiB | [postgresql-17-auto-failover_2.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg12+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 375.0 KiB | [postgresql-17-auto-failover_2.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg13+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 372.4 KiB | [postgresql-17-auto-failover_2.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg13+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 393.9 KiB | [postgresql-17-auto-failover_2.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.7 KiB | [postgresql-17-auto-failover_2.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 364.9 KiB | [postgresql-17-auto-failover_2.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 362.5 KiB | [postgresql-17-auto-failover_2.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 366.9 KiB | [postgresql-17-auto-failover_2.2-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 361.7 KiB | [postgresql-17-auto-failover_2.2-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_16` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.5 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_auto_failover_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_16` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 844.0 KiB | [pg_auto_failover_16-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_auto_failover_16-2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_16` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 809.4 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_auto_failover_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_16` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 843.8 KiB | [pg_auto_failover_16-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_auto_failover_16-2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_16` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 786.3 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_auto_failover_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_16` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 807.5 KiB | [pg_auto_failover_16-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_auto_failover_16-2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_16` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.1 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_auto_failover_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_16` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 817.4 KiB | [pg_auto_failover_16-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_auto_failover_16-2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_16` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 788.1 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_auto_failover_16-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_16` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 785.7 KiB | [pg_auto_failover_16-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_auto_failover_16-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 368.8 KiB | [postgresql-16-auto-failover_2.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg12+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 365.9 KiB | [postgresql-16-auto-failover_2.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg12+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 369.8 KiB | [postgresql-16-auto-failover_2.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg13+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 366.4 KiB | [postgresql-16-auto-failover_2.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg13+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 387.5 KiB | [postgresql-16-auto-failover_2.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 383.4 KiB | [postgresql-16-auto-failover_2.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 358.7 KiB | [postgresql-16-auto-failover_2.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 356.6 KiB | [postgresql-16-auto-failover_2.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 360.7 KiB | [postgresql-16-auto-failover_2.2-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 357.7 KiB | [postgresql-16-auto-failover_2.2-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_15` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.6 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_auto_failover_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_15` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 843.6 KiB | [pg_auto_failover_15-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_auto_failover_15-2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_15` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 842.5 KiB | [pg_auto_failover_15-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_auto_failover_15-2.0-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_15` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 809.3 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_auto_failover_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_15` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 843.0 KiB | [pg_auto_failover_15-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_auto_failover_15-2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_15` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 841.6 KiB | [pg_auto_failover_15-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_auto_failover_15-2.0-1.rhel8.aarch64.rpm) |
| `pg_auto_failover_15` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 789.5 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_auto_failover_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_15` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 811.6 KiB | [pg_auto_failover_15-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_auto_failover_15-2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_15` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 812.1 KiB | [pg_auto_failover_15-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_auto_failover_15-2.0-1.rhel9.x86_64.rpm) |
| `pg_auto_failover_15` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 792.9 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_auto_failover_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_15` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 820.9 KiB | [pg_auto_failover_15-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_auto_failover_15-2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_15` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 821.5 KiB | [pg_auto_failover_15-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_auto_failover_15-2.0-1.rhel9.aarch64.rpm) |
| `pg_auto_failover_15` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 791.7 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_auto_failover_15-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_15` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 788.8 KiB | [pg_auto_failover_15-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_auto_failover_15-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 368.4 KiB | [postgresql-15-auto-failover_2.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg12+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 365.9 KiB | [postgresql-15-auto-failover_2.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg12+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 370.1 KiB | [postgresql-15-auto-failover_2.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg13+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 366.1 KiB | [postgresql-15-auto-failover_2.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg13+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 391.1 KiB | [postgresql-15-auto-failover_2.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.1 KiB | [postgresql-15-auto-failover_2.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.6 KiB | [postgresql-15-auto-failover_2.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 357.6 KiB | [postgresql-15-auto-failover_2.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 363.0 KiB | [postgresql-15-auto-failover_2.2-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 357.8 KiB | [postgresql-15-auto-failover_2.2-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_14` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 810.9 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auto_failover_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_14` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 841.9 KiB | [pg_auto_failover_14-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auto_failover_14-2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_14` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 840.6 KiB | [pg_auto_failover_14-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auto_failover_14-2.0-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_14` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 994.9 KiB | [pg_auto_failover_14-1.6.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auto_failover_14-1.6.4-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_14` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 929.2 KiB | [pg_auto_failover_14-1.6.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_auto_failover_14-1.6.2-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_14` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 808.0 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_auto_failover_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_14` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 841.4 KiB | [pg_auto_failover_14-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_auto_failover_14-2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_14` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.1 KiB | [pg_auto_failover_14-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_auto_failover_14-2.0-1.rhel8.aarch64.rpm) |
| `pg_auto_failover_14` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 789.0 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_auto_failover_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_14` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 811.2 KiB | [pg_auto_failover_14-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_auto_failover_14-2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_14` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 811.7 KiB | [pg_auto_failover_14-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_auto_failover_14-2.0-1.rhel9.x86_64.rpm) |
| `pg_auto_failover_14` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 966.2 KiB | [pg_auto_failover_14-1.6.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_auto_failover_14-1.6.4-1.rhel9.x86_64.rpm) |
| `pg_auto_failover_14` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 792.5 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_auto_failover_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_14` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 820.6 KiB | [pg_auto_failover_14-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_auto_failover_14-2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_14` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 821.0 KiB | [pg_auto_failover_14-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_auto_failover_14-2.0-1.rhel9.aarch64.rpm) |
| `pg_auto_failover_14` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 792.2 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_auto_failover_14-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_14` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 789.6 KiB | [pg_auto_failover_14-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_auto_failover_14-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.7 KiB | [postgresql-14-auto-failover_2.2-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg12+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 359.4 KiB | [postgresql-14-auto-failover_2.2-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg12+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 364.0 KiB | [postgresql-14-auto-failover_2.2-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg13+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 360.4 KiB | [postgresql-14-auto-failover_2.2-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg13+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 386.6 KiB | [postgresql-14-auto-failover_2.2-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 382.8 KiB | [postgresql-14-auto-failover_2.2-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 354.8 KiB | [postgresql-14-auto-failover_2.2-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 350.8 KiB | [postgresql-14-auto-failover_2.2-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 357.2 KiB | [postgresql-14-auto-failover_2.2-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 353.5 KiB | [postgresql-14-auto-failover_2.2-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hapostgres/pg_auto_failover" title="Repository" icon="github" subtitle="github.com/hapostgres/pg_auto_failover" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgautofailover;		# install via package name, for the active PG version

pig install pgautofailover -v 18;   # install for PG 18
pig install pgautofailover -v 17;   # install for PG 17
pig install pgautofailover -v 16;   # install for PG 16
pig install pgautofailover -v 15;   # install for PG 15
pig install pgautofailover -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgautofailover';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgautofailover CASCADE; -- requires btree_gist
```




## Usage

> [pgautofailover: pg_auto_failover](https://github.com/hapostgres/pg_auto_failover)

pg_auto_failover is an extension and service for automated PostgreSQL failover. It consists of a monitor node (running the `pgautofailover` extension), and keeper processes on each data node managed by the `pg_autoctl` CLI.

### Architecture

- **Monitor**: a PostgreSQL instance with the `pgautofailover` extension that implements a state machine for failover decisions
- **Keeper** (`pg_autoctl run`): runs on each data node, reports health to the monitor and executes state transitions
- Supports 2+ node setups with synchronous replication by default
- Supports Citus HA (since v2.0)

### Key CLI Commands

```bash
# Create the monitor
pg_autoctl create monitor --pgdata /path/to/monitor --pgport 5000

# Create a data node (primary or secondary auto-assigned)
pg_autoctl create postgres --pgdata /path/to/data --pgport 5001 --monitor postgres://monitor:5000/pg_auto_failover

# Run the keeper (foreground)
pg_autoctl run --pgdata /path/to/data

# Check cluster state
pg_autoctl show state --pgdata /path/to/monitor

# Perform a manual switchover
pg_autoctl perform switchover --pgdata /path/to/monitor

# Perform a manual failover
pg_autoctl perform failover --pgdata /path/to/monitor
```

### Failover Behavior

The monitor tracks node health. When a secondary becomes unavailable or its lag is too high, it is removed from `synchronous_standby_names` on the primary. Failover/switchover operations are blocked until the secondary is healthy again, preventing data loss.

### Documentation

Full documentation: [pg-auto-failover.readthedocs.io](https://pg-auto-failover.readthedocs.io/en/main/)
