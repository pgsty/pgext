---
title: "pgautofailover"
linkTitle: "pgautofailover"
description: "pg_auto_failover"
weight: 5150
categories: ["ADMIN"]
width: full
---

[**pgautofailover**](https://github.com/hapostgres/pg_auto_failover)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5150** | {{< badge content="pgautofailover" link="https://github.com/hapostgres/pg_auto_failover" >}} | {{< ext "pgautofailover" >}} | `2.2` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "btree_gist" >}} |
|   **See Also**    | {{< ext "pglogical" >}} {{< ext "pglogical_origin" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgpool_recovery" >}} {{< ext "repmgr" >}} {{< ext "pg_checksums" >}} {{< ext "pgpool_adm" >}} {{< ext "bgw_replstatus" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgautofailover" >}} | `2.2` | {{< bg "18" "pg_auto_failover_18*" "red" >}} {{< bg "17" "pg_auto_failover_17*" "green" >}} {{< bg "16" "pg_auto_failover_16*" "green" >}} {{< bg "15" "pg_auto_failover_15*" "green" >}} {{< bg "14" "pg_auto_failover_14*" "green" >}} {{< bg "13" "pg_auto_failover_13*" "green" >}} | `pg_auto_failover_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgautofailover" >}} | `2.2` | {{< bg "18" "postgresql-18-auto-failover" "red" >}} {{< bg "17" "postgresql-17-auto-failover" "green" >}} {{< bg "16" "postgresql-16-auto-failover" "green" >}} {{< bg "15" "postgresql-15-auto-failover" "green" >}} {{< bg "14" "postgresql-14-auto-failover" "green" >}} {{< bg "13" "postgresql-13-auto-failover" "green" >}} | `postgresql-$v-auto-failover` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 11" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_auto_failover_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "pg_auto_failover_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_auto_failover_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-auto-failover : MISS 0" "red" >}}      | {{< bg "PGDG 2.2" "postgresql-17-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-16-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-15-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-14-auto-failover : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "postgresql-13-auto-failover : AVAIL 1" "blue" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_17` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 812.9 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 809.9 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 786.3 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 789.2 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 788.0 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_auto_failover_17-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_17` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 785.7 KiB | [pg_auto_failover_17-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_auto_failover_17-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 372.6 KiB | [postgresql-17-auto-failover_2.2-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg120+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 370.0 KiB | [postgresql-17-auto-failover_2.2-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg120+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 374.0 KiB | [postgresql-17-auto-failover_2.2-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg130+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 371.0 KiB | [postgresql-17-auto-failover_2.2-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg130+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 393.4 KiB | [postgresql-17-auto-failover_2.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.1 KiB | [postgresql-17-auto-failover_2.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 363.8 KiB | [postgresql-17-auto-failover_2.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 360.2 KiB | [postgresql-17-auto-failover_2.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-17-auto-failover_2.2-2.pgdg24.04+1_arm64.deb) |

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
| `postgresql-16-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 368.4 KiB | [postgresql-16-auto-failover_2.2-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg120+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 364.9 KiB | [postgresql-16-auto-failover_2.2-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg120+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 369.9 KiB | [postgresql-16-auto-failover_2.2-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg130+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 364.8 KiB | [postgresql-16-auto-failover_2.2-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg130+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 386.8 KiB | [postgresql-16-auto-failover_2.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 383.4 KiB | [postgresql-16-auto-failover_2.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 359.8 KiB | [postgresql-16-auto-failover_2.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.9 KiB | [postgresql-16-auto-failover_2.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-16-auto-failover_2.2-2.pgdg24.04+1_arm64.deb) |

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
| `postgresql-15-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 367.8 KiB | [postgresql-15-auto-failover_2.2-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg120+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 364.3 KiB | [postgresql-15-auto-failover_2.2-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg120+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 369.5 KiB | [postgresql-15-auto-failover_2.2-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg130+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 366.2 KiB | [postgresql-15-auto-failover_2.2-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg130+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 390.5 KiB | [postgresql-15-auto-failover_2.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.0 KiB | [postgresql-15-auto-failover_2.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 359.7 KiB | [postgresql-15-auto-failover_2.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 356.6 KiB | [postgresql-15-auto-failover_2.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-15-auto-failover_2.2-2.pgdg24.04+1_arm64.deb) |

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
| `postgresql-14-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.7 KiB | [postgresql-14-auto-failover_2.2-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg120+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 358.7 KiB | [postgresql-14-auto-failover_2.2-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg120+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 364.1 KiB | [postgresql-14-auto-failover_2.2-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg130+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 360.2 KiB | [postgresql-14-auto-failover_2.2-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg130+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 384.9 KiB | [postgresql-14-auto-failover_2.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 381.4 KiB | [postgresql-14-auto-failover_2.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 354.8 KiB | [postgresql-14-auto-failover_2.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 351.0 KiB | [postgresql-14-auto-failover_2.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-14-auto-failover_2.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_auto_failover_13` | `2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 810.5 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 840.6 KiB | [pg_auto_failover_13-2.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-2.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 839.5 KiB | [pg_auto_failover_13-2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-2.0-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.6.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 993.7 KiB | [pg_auto_failover_13-1.6.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.6.4-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 927.9 KiB | [pg_auto_failover_13-1.6.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.6.2-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.2 MiB | [pg_auto_failover_13-1.6.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.6.1-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 924.8 KiB | [pg_auto_failover_13-1.6.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.6.1-2.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.5.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.2 MiB | [pg_auto_failover_13-1.5.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.5.2-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [pg_auto_failover_13-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.5.1-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [pg_auto_failover_13-1.4.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.4.1-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 1.3 MiB | [pg_auto_failover_13-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_auto_failover_13-1.4.0-1.rhel8.x86_64.rpm) |
| `pg_auto_failover_13` | `2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 807.9 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_auto_failover_13-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_13` | `2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 841.4 KiB | [pg_auto_failover_13-2.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_auto_failover_13-2.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_auto_failover_13` | `2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 840.1 KiB | [pg_auto_failover_13-2.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_auto_failover_13-2.0-1.rhel8.aarch64.rpm) |
| `pg_auto_failover_13` | `2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 789.4 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_auto_failover_13-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_13` | `2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 811.7 KiB | [pg_auto_failover_13-2.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_auto_failover_13-2.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_auto_failover_13` | `2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 812.0 KiB | [pg_auto_failover_13-2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_auto_failover_13-2.0-1.rhel9.x86_64.rpm) |
| `pg_auto_failover_13` | `1.6.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 966.8 KiB | [pg_auto_failover_13-1.6.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_auto_failover_13-1.6.4-1.rhel9.x86_64.rpm) |
| `pg_auto_failover_13` | `2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 792.8 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_auto_failover_13-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_13` | `2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 820.6 KiB | [pg_auto_failover_13-2.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_auto_failover_13-2.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_auto_failover_13` | `2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 821.3 KiB | [pg_auto_failover_13-2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_auto_failover_13-2.0-1.rhel9.aarch64.rpm) |
| `pg_auto_failover_13` | `2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 792.4 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_auto_failover_13-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_auto_failover_13` | `2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 790.1 KiB | [pg_auto_failover_13-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/pg_auto_failover_13-2.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-auto-failover` | `2.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 359.0 KiB | [postgresql-13-auto-failover_2.2-2.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg120+1_amd64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 354.8 KiB | [postgresql-13-auto-failover_2.2-2.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg120+1_arm64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 360.3 KiB | [postgresql-13-auto-failover_2.2-2.pgdg130+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg130+1_amd64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.1 KiB | [postgresql-13-auto-failover_2.2-2.pgdg130+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg130+1_arm64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 382.7 KiB | [postgresql-13-auto-failover_2.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 377.9 KiB | [postgresql-13-auto-failover_2.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 351.1 KiB | [postgresql-13-auto-failover_2.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-13-auto-failover` | `2.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 347.3 KiB | [postgresql-13-auto-failover_2.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-auto-failover/postgresql-13-auto-failover_2.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hapostgres/pg_auto_failover" title="Repository" icon="github" subtitle="github.com/hapostgres/pg_auto_failover" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgautofailover; # install by extension name, for the current active PG version
pig ext install pgautofailover; # install via package alias, for the active PG version
pig ext install pgautofailover -v 17;   # install for PG 17
pig ext install pgautofailover -v 16;   # install for PG 16
pig ext install pgautofailover -v 15;   # install for PG 15
pig ext install pgautofailover -v 14;   # install for PG 14
pig ext install pgautofailover -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgautofailover;
```

