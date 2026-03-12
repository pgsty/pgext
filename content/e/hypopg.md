---
title: "hypopg"
linkTitle: "hypopg"
description: "Hypothetical indexes for PostgreSQL"
weight: 2790
categories: ["FEAT"]
width: full
---

[**hypopg**](https://github.com/HypoPG/hypopg) : Hypothetical indexes for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2790** | {{< badge content="hypopg" link="https://github.com/HypoPG/hypopg" >}} | {{< ext "hypopg" >}} | `1.4.2` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "index_advisor" >}} {{< ext "pg_qualstats" >}} {{< ext "powa" >}} {{< ext "pg_hint_plan" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "btree_gin" >}} {{< ext "pg_show_plans" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `hypopg` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.2` | {{< bg "18" "hypopg_18" "green" >}} {{< bg "17" "hypopg_17" "green" >}} {{< bg "16" "hypopg_16" "green" >}} {{< bg "15" "hypopg_15" "green" >}} {{< bg "14" "hypopg_14" "green" >}} | `hypopg_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4.2` | {{< bg "18" "postgresql-18-hypopg" "green" >}} {{< bg "17" "postgresql-17-hypopg" "green" >}} {{< bg "16" "postgresql-16-hypopg" "green" >}} {{< bg "15" "postgresql-15-hypopg" "green" >}} {{< bg "14" "postgresql-14-hypopg" "green" >}} | `postgresql-$v-hypopg` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4.2" "hypopg_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.1" "hypopg_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4.2" "postgresql-18-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-17-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-16-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-15-hypopg : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.2" "postgresql-14-hypopg : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hypopg_18` | `1.4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 31.1 KiB | [hypopg_18-1.4.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/hypopg_18-1.4.2-1PGDG.rhel8.x86_64.rpm) |
| `hypopg_18` | `1.4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.8 KiB | [hypopg_18-1.4.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/hypopg_18-1.4.2-1PGDG.rhel8.aarch64.rpm) |
| `hypopg_18` | `1.4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.7 KiB | [hypopg_18-1.4.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/hypopg_18-1.4.2-1PGDG.rhel9.x86_64.rpm) |
| `hypopg_18` | `1.4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.1 KiB | [hypopg_18-1.4.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/hypopg_18-1.4.2-1PGDG.rhel9.aarch64.rpm) |
| `hypopg_18` | `1.4.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.3 KiB | [hypopg_18-1.4.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/hypopg_18-1.4.2-1PGDG.rhel10.x86_64.rpm) |
| `hypopg_18` | `1.4.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.0 KiB | [hypopg_18-1.4.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/hypopg_18-1.4.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-hypopg` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.4 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg12+1_amd64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.9 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg12+1_arm64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.5 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg13+1_amd64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 58.3 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg13+1_arm64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 59.2 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 59.3 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.4 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-hypopg` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.7 KiB | [postgresql-18-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-18-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hypopg_17` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.2 KiB | [hypopg_17-1.4.1-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hypopg_17-1.4.1-2PGDG.rhel8.x86_64.rpm) |
| `hypopg_17` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [hypopg_17-1.4.1-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hypopg_17-1.4.1-2PGDG.rhel8.aarch64.rpm) |
| `hypopg_17` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.1 KiB | [hypopg_17-1.4.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hypopg_17-1.4.1-2PGDG.rhel9.x86_64.rpm) |
| `hypopg_17` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.6 KiB | [hypopg_17-1.4.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hypopg_17-1.4.1-2PGDG.rhel9.aarch64.rpm) |
| `hypopg_17` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.1 KiB | [hypopg_17-1.4.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/hypopg_17-1.4.1-3PGDG.rhel10.x86_64.rpm) |
| `hypopg_17` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.8 KiB | [hypopg_17-1.4.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/hypopg_17-1.4.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-hypopg` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.4 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg12+1_amd64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 57.9 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg12+1_arm64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.6 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg13+1_amd64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 58.3 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg13+1_arm64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 72.0 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 72.1 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.3 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-hypopg` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.6 KiB | [postgresql-17-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-17-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hypopg_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.1 KiB | [hypopg_16-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hypopg_16-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `hypopg_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.6 KiB | [hypopg_16-1.4.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hypopg_16-1.4.0-2PGDG.rhel8.x86_64.rpm) |
| `hypopg_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.9 KiB | [hypopg_16-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hypopg_16-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `hypopg_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.3 KiB | [hypopg_16-1.4.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hypopg_16-1.4.0-2PGDG.rhel8.aarch64.rpm) |
| `hypopg_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.9 KiB | [hypopg_16-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hypopg_16-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `hypopg_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 28.4 KiB | [hypopg_16-1.4.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hypopg_16-1.4.0-2PGDG.rhel9.x86_64.rpm) |
| `hypopg_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.3 KiB | [hypopg_16-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hypopg_16-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `hypopg_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 29.7 KiB | [hypopg_16-1.4.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hypopg_16-1.4.0-2PGDG.rhel9.aarch64.rpm) |
| `hypopg_16` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 30.2 KiB | [hypopg_16-1.4.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/hypopg_16-1.4.1-3PGDG.rhel10.x86_64.rpm) |
| `hypopg_16` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 31.9 KiB | [hypopg_16-1.4.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/hypopg_16-1.4.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-hypopg` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.4 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg12+1_amd64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 58.1 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg12+1_arm64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 57.7 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg13+1_amd64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 58.5 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg13+1_arm64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 72.0 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 72.0 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.3 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-hypopg` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 57.7 KiB | [postgresql-16-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-16-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hypopg_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.5 KiB | [hypopg_15-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hypopg_15-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `hypopg_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.8 KiB | [hypopg_15-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hypopg_15-1.4.0-1.rhel8.x86_64.rpm) |
| `hypopg_15` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.8 KiB | [hypopg_15-1.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hypopg_15-1.3.1-1.rhel8.x86_64.rpm) |
| `hypopg_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.1 KiB | [hypopg_15-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hypopg_15-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `hypopg_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.3 KiB | [hypopg_15-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hypopg_15-1.4.0-1.rhel8.aarch64.rpm) |
| `hypopg_15` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 74.8 KiB | [hypopg_15-1.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hypopg_15-1.3.1-1.rhel8.aarch64.rpm) |
| `hypopg_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.8 KiB | [hypopg_15-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hypopg_15-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `hypopg_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.1 KiB | [hypopg_15-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hypopg_15-1.4.0-1.rhel9.x86_64.rpm) |
| `hypopg_15` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 75.6 KiB | [hypopg_15-1.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hypopg_15-1.3.1-1.rhel9.x86_64.rpm) |
| `hypopg_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.2 KiB | [hypopg_15-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hypopg_15-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `hypopg_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.5 KiB | [hypopg_15-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hypopg_15-1.4.0-1.rhel9.aarch64.rpm) |
| `hypopg_15` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 76.2 KiB | [hypopg_15-1.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hypopg_15-1.3.1-1.rhel9.aarch64.rpm) |
| `hypopg_15` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.1 KiB | [hypopg_15-1.4.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/hypopg_15-1.4.1-3PGDG.rhel10.x86_64.rpm) |
| `hypopg_15` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.8 KiB | [hypopg_15-1.4.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/hypopg_15-1.4.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-hypopg` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.9 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg12+1_amd64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 58.1 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg12+1_arm64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.0 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg13+1_amd64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 58.7 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg13+1_arm64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 72.5 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 72.5 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.9 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-hypopg` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 58.2 KiB | [postgresql-15-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-15-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hypopg_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 30.6 KiB | [hypopg_14-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hypopg_14-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `hypopg_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 29.8 KiB | [hypopg_14-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hypopg_14-1.4.0-1.rhel8.x86_64.rpm) |
| `hypopg_14` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 74.4 KiB | [hypopg_14-1.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hypopg_14-1.3.1-1.rhel8.x86_64.rpm) |
| `hypopg_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 31.1 KiB | [hypopg_14-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hypopg_14-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `hypopg_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 30.3 KiB | [hypopg_14-1.4.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hypopg_14-1.4.0-1.rhel8.aarch64.rpm) |
| `hypopg_14` | `1.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 73.8 KiB | [hypopg_14-1.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hypopg_14-1.3.1-1.rhel8.aarch64.rpm) |
| `hypopg_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.8 KiB | [hypopg_14-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hypopg_14-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `hypopg_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 29.1 KiB | [hypopg_14-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hypopg_14-1.4.0-1.rhel9.x86_64.rpm) |
| `hypopg_14` | `1.3.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 74.5 KiB | [hypopg_14-1.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hypopg_14-1.3.1-1.rhel9.x86_64.rpm) |
| `hypopg_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 31.2 KiB | [hypopg_14-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hypopg_14-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `hypopg_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 30.4 KiB | [hypopg_14-1.4.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hypopg_14-1.4.0-1.rhel9.aarch64.rpm) |
| `hypopg_14` | `1.3.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 75.2 KiB | [hypopg_14-1.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hypopg_14-1.3.1-1.rhel9.aarch64.rpm) |
| `hypopg_14` | `1.4.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.0 KiB | [hypopg_14-1.4.1-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/hypopg_14-1.4.1-3PGDG.rhel10.x86_64.rpm) |
| `hypopg_14` | `1.4.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 32.8 KiB | [hypopg_14-1.4.1-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/hypopg_14-1.4.1-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-hypopg` | `1.4.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 57.8 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg12+1_amd64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 58.1 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg12+1_arm64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 58.0 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg13+1_amd64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 58.6 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg13+1_arm64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 71.6 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 71.7 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 57.9 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-hypopg` | `1.4.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 58.2 KiB | [postgresql-14-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/h/hypopg/postgresql-14-hypopg_1.4.2-2.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HypoPG/hypopg" title="Repository" icon="github" subtitle="github.com/HypoPG/hypopg" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hypopg-1.4.2.tar.gz" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install hypopg;		# install via package name, for the active PG version

pig install hypopg -v 18;   # install for PG 18
pig install hypopg -v 17;   # install for PG 17
pig install hypopg -v 16;   # install for PG 16
pig install hypopg -v 15;   # install for PG 15
pig install hypopg -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hypopg;
```




## Usage

> [hypopg: Hypothetical indexes for PostgreSQL](https://github.com/HypoPG/hypopg)

HypoPG lets you create hypothetical (virtual) indexes that exist only in the current session and are considered by `EXPLAIN` (without `ANALYZE`) for query planning. This enables testing the impact of indexes without the cost of actually creating them.

### Functions

| Function | Description |
|----------|-------------|
| `hypopg_create_index(query text)` | Create a hypothetical index using CREATE INDEX syntax |
| `hypopg_list_indexes()` | List all hypothetical indexes in the session |
| `hypopg_drop_index(oid)` | Drop a specific hypothetical index by OID |
| `hypopg_reset()` | Drop all hypothetical indexes |
| `hypopg()` | Return hypothetical indexes in pg_index-like format |

### Workflow

Create a test table and check the baseline plan:

```sql
CREATE TABLE hypo AS SELECT id, 'line ' || id AS val FROM generate_series(1, 10000) id;
ANALYZE hypo;
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Seq Scan on hypo (cost=0.00..170.00 rows=1 width=15)
```

Create a hypothetical index:

```sql
SELECT * FROM hypopg_create_index('CREATE INDEX ON hypo (id)');
--  indexrelid |      indexname
-- ------------+----------------------
--       13543 | <13543>btree_hypo_id
```

Check the plan with the hypothetical index:

```sql
EXPLAIN SELECT * FROM hypo WHERE id = 1;
-- Index Scan using <13543>btree_hypo_id on hypo (cost=0.04..8.06 rows=1 width=15)
```

List and manage hypothetical indexes:

```sql
SELECT * FROM hypopg_list_indexes();
SELECT * FROM hypopg_drop_index(13543);
SELECT * FROM hypopg_reset();
```

### Limitations

- Only `EXPLAIN` without `ANALYZE` will consider hypothetical indexes
- Hypothetical indexes exist only in the current backend session
- Other concurrent connections are not affected
- Index names and some CREATE INDEX options are ignored
