---
title: "rum"
linkTitle: "rum"
description: "RUM index access method"
weight: 2710
categories: ["FEAT"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**rum**](https://github.com/postgrespro/rum) : RUM index access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2710** | {{< badge content="rum" link="https://github.com/postgrespro/rum" >}} | {{< ext "rum" >}} | `1.3.15` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_trgm" >}} {{< ext "pgfaceting" >}} {{< ext "roaringbitmap" >}} {{< ext "bloom" >}} {{< ext "parray_gin" >}} {{< ext "btree_gin" >}} {{< ext "intarray" >}} {{< ext "btree_gist" >}} |

> [!Note] 1.3.15 build pass on pg 16,17,18


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.3.15` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `rum` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.15` | {{< bg "19" "rum_19" "green" >}} {{< bg "18" "rum_18" "green" >}} {{< bg "17" "rum_17" "green" >}} {{< bg "16" "rum_16" "green" >}} {{< bg "15" "rum_15" "green" >}} | `rum_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.15` | {{< bg "19" "postgresql-19-rum" "red" >}} {{< bg "18" "postgresql-18-rum" "green" >}} {{< bg "17" "postgresql-17-rum" "green" >}} {{< bg "16" "postgresql-16-rum" "green" >}} {{< bg "15" "postgresql-15-rum" "green" >}} | `postgresql-$v-rum` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "rum_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "rum_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "MISS" "rum_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "MISS" "rum_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "MISS" "rum_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "MISS" "rum_19 : MISS 0" "red" >}} | {{< bg "PIGSTY 1.3.15" "rum_18 : AVAIL 1" "green" >}} | {{< bg "PGDG 1.3.15" "rum_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.3.15" "rum_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "MISS" "postgresql-19-rum : MISS 0" "red" >}} | {{< bg "PGDG 1.3.15" "postgresql-18-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-17-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-16-rum : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.15" "postgresql-15-rum : AVAIL 2" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_18` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pigsty | 104.0 KiB | [rum_18-1.3.15-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rum_18-1.3.15-1PIGSTY.el8.x86_64.rpm) |
| `rum_18` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pigsty | 97.5 KiB | [rum_18-1.3.15-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rum_18-1.3.15-1PIGSTY.el8.aarch64.rpm) |
| `rum_18` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.2 KiB | [rum_18-1.3.15-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rum_18-1.3.15-1PIGSTY.el9.x86_64.rpm) |
| `rum_18` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.7 KiB | [rum_18-1.3.15-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rum_18-1.3.15-1PIGSTY.el9.aarch64.rpm) |
| `rum_18` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.1 KiB | [rum_18-1.3.15-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rum_18-1.3.15-1PIGSTY.el10.x86_64.rpm) |
| `rum_18` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pigsty | 93.6 KiB | [rum_18-1.3.15-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rum_18-1.3.15-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 233.3 KiB | [postgresql-18-rum_1.3.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg12+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 233.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.5 KiB | [postgresql-18-rum_1.3.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg12+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.4 KiB | [postgresql-18-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 233.9 KiB | [postgresql-18-rum_1.3.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg13+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 233.5 KiB | [postgresql-18-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.2 KiB | [postgresql-18-rum_1.3.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg13+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.2 KiB | [postgresql-18-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 240.5 KiB | [postgresql-18-rum_1.3.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 240.5 KiB | [postgresql-18-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 232.4 KiB | [postgresql-18-rum_1.3.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 232.1 KiB | [postgresql-18-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 234.3 KiB | [postgresql-18-rum_1.3.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 234.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.5 KiB | [postgresql-18-rum_1.3.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.3 KiB | [postgresql-18-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 232.3 KiB | [postgresql-18-rum_1.3.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 232.5 KiB | [postgresql-18-rum_1.3.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 223.5 KiB | [postgresql-18-rum_1.3.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 223.8 KiB | [postgresql-18-rum_1.3.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-18-rum_1.3.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_17` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.9 KiB | [rum_17-1.3.15-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/rum_17-1.3.15-3PGDG.rhel8.10.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pigsty | 104.2 KiB | [rum_17-1.3.15-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rum_17-1.3.15-1PIGSTY.el8.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.2 KiB | [rum_17-1.3.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/rum_17-1.3.15-1PGDG.rhel8.10.x86_64.rpm) |
| `rum_17` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 93.1 KiB | [rum_17-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/rum_17-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.6 KiB | [rum_17-1.3.15-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/rum_17-1.3.15-3PGDG.rhel8.10.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pigsty | 97.7 KiB | [rum_17-1.3.15-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rum_17-1.3.15-1PIGSTY.el8.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 87.9 KiB | [rum_17-1.3.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/rum_17-1.3.15-1PGDG.rhel8.10.aarch64.rpm) |
| `rum_17` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.8 KiB | [rum_17-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/rum_17-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 93.1 KiB | [rum_17-1.3.15-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/rum_17-1.3.15-3PGDG.rhel9.9.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.5 KiB | [rum_17-1.3.15-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rum_17-1.3.15-1PIGSTY.el9.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.0 KiB | [rum_17-1.3.15-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/rum_17-1.3.15-3PGDG.rhel9.9.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.6 KiB | [rum_17-1.3.15-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rum_17-1.3.15-1PIGSTY.el9.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 94.1 KiB | [rum_17-1.3.15-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/rum_17-1.3.15-3PGDG.rhel10.3.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.3 KiB | [rum_17-1.3.15-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rum_17-1.3.15-1PIGSTY.el10.x86_64.rpm) |
| `rum_17` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.8 KiB | [rum_17-1.3.15-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/rum_17-1.3.15-3PGDG.rhel10.3.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.8 KiB | [rum_17-1.3.15-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/rum_17-1.3.15-3PGDG.rhel10.2.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pigsty | 94.0 KiB | [rum_17-1.3.15-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rum_17-1.3.15-1PIGSTY.el10.aarch64.rpm) |
| `rum_17` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.1 KiB | [rum_17-1.3.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/rum_17-1.3.15-1PGDG.rhel10.2.aarch64.rpm) |
| `rum_17` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.6 KiB | [rum_17-1.3.14-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/rum_17-1.3.14-4PGDG.rhel10.2.aarch64.rpm) |
| `rum_17` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.8 KiB | [rum_17-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/rum_17-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.3 KiB | [postgresql-17-rum_1.3.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg12+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.4 KiB | [postgresql-17-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 226.0 KiB | [postgresql-17-rum_1.3.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg12+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.9 KiB | [postgresql-17-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.7 KiB | [postgresql-17-rum_1.3.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg13+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.6 KiB | [postgresql-17-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 227.0 KiB | [postgresql-17-rum_1.3.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg13+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.7 KiB | [postgresql-17-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 265.2 KiB | [postgresql-17-rum_1.3.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 265.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 257.0 KiB | [postgresql-17-rum_1.3.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 257.4 KiB | [postgresql-17-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 235.0 KiB | [postgresql-17-rum_1.3.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 235.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.1 KiB | [postgresql-17-rum_1.3.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 227.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 232.8 KiB | [postgresql-17-rum_1.3.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 233.0 KiB | [postgresql-17-rum_1.3.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.1 KiB | [postgresql-17-rum_1.3.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.1 KiB | [postgresql-17-rum_1.3.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-17-rum_1.3.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_16` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.8 KiB | [rum_16-1.3.15-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.15-3PGDG.rhel8.10.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pigsty | 104.1 KiB | [rum_16-1.3.15-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/rum_16-1.3.15-1PIGSTY.el8.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.2 KiB | [rum_16-1.3.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.15-1PGDG.rhel8.10.x86_64.rpm) |
| `rum_16` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 93.0 KiB | [rum_16-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_16` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 92.7 KiB | [rum_16-1.3.13-2.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/rum_16-1.3.13-2.rhel8.1.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.6 KiB | [rum_16-1.3.15-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.15-3PGDG.rhel8.10.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pigsty | 97.7 KiB | [rum_16-1.3.15-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/rum_16-1.3.15-1PIGSTY.el8.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 87.9 KiB | [rum_16-1.3.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.15-1PGDG.rhel8.10.aarch64.rpm) |
| `rum_16` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.8 KiB | [rum_16-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_16` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 86.4 KiB | [rum_16-1.3.13-2.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/rum_16-1.3.13-2.rhel8.1.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 93.1 KiB | [rum_16-1.3.15-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/rum_16-1.3.15-3PGDG.rhel9.9.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pigsty | 96.5 KiB | [rum_16-1.3.15-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/rum_16-1.3.15-1PIGSTY.el9.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 89.0 KiB | [rum_16-1.3.15-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/rum_16-1.3.15-3PGDG.rhel9.9.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pigsty | 92.8 KiB | [rum_16-1.3.15-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/rum_16-1.3.15-1PIGSTY.el9.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 94.1 KiB | [rum_16-1.3.15-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/rum_16-1.3.15-3PGDG.rhel10.3.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pigsty | 97.3 KiB | [rum_16-1.3.15-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rum_16-1.3.15-1PIGSTY.el10.x86_64.rpm) |
| `rum_16` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.8 KiB | [rum_16-1.3.15-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/rum_16-1.3.15-3PGDG.rhel10.3.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.8 KiB | [rum_16-1.3.15-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/rum_16-1.3.15-3PGDG.rhel10.2.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pigsty | 93.9 KiB | [rum_16-1.3.15-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rum_16-1.3.15-1PIGSTY.el10.aarch64.rpm) |
| `rum_16` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 89.1 KiB | [rum_16-1.3.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/rum_16-1.3.15-1PGDG.rhel10.2.aarch64.rpm) |
| `rum_16` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.5 KiB | [rum_16-1.3.14-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/rum_16-1.3.14-4PGDG.rhel10.2.aarch64.rpm) |
| `rum_16` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 88.7 KiB | [rum_16-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/rum_16-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.3 KiB | [postgresql-16-rum_1.3.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg12+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 234.4 KiB | [postgresql-16-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.9 KiB | [postgresql-16-rum_1.3.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg12+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 225.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.4 KiB | [postgresql-16-rum_1.3.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg13+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 234.5 KiB | [postgresql-16-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.7 KiB | [postgresql-16-rum_1.3.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg13+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 226.6 KiB | [postgresql-16-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 265.0 KiB | [postgresql-16-rum_1.3.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 264.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 256.8 KiB | [postgresql-16-rum_1.3.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 256.7 KiB | [postgresql-16-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 235.2 KiB | [postgresql-16-rum_1.3.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 234.9 KiB | [postgresql-16-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.8 KiB | [postgresql-16-rum_1.3.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 226.9 KiB | [postgresql-16-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 232.6 KiB | [postgresql-16-rum_1.3.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 232.8 KiB | [postgresql-16-rum_1.3.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 223.6 KiB | [postgresql-16-rum_1.3.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 224.1 KiB | [postgresql-16-rum_1.3.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-16-rum_1.3.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rum_15` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.7 KiB | [rum_15-1.3.15-3PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.15-3PGDG.rhel8.10.x86_64.rpm) |
| `rum_15` | `1.3.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.7 KiB | [rum_15-1.3.15-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.15-1PGDG.rhel8.10.x86_64.rpm) |
| `rum_15` | `1.3.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.5 KiB | [rum_15-1.3.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.14-1PGDG.rhel8.x86_64.rpm) |
| `rum_15` | `1.3.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.1 KiB | [rum_15-1.3.13-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/rum_15-1.3.13-1.rhel8.x86_64.rpm) |
| `rum_15` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [rum_15-1.3.15-3PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.15-3PGDG.rhel8.10.aarch64.rpm) |
| `rum_15` | `1.3.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.8 KiB | [rum_15-1.3.15-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.15-1PGDG.rhel8.10.aarch64.rpm) |
| `rum_15` | `1.3.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.7 KiB | [rum_15-1.3.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.14-1PGDG.rhel8.aarch64.rpm) |
| `rum_15` | `1.3.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.0 KiB | [rum_15-1.3.13-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/rum_15-1.3.13-1.rhel8.aarch64.rpm) |
| `rum_15` | `1.3.15` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.6 KiB | [rum_15-1.3.15-3PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/rum_15-1.3.15-3PGDG.rhel9.9.x86_64.rpm) |
| `rum_15` | `1.3.15` | [el9.aarch64](/os/el9.aarch64) | pgdg | 108.8 KiB | [rum_15-1.3.15-3PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/rum_15-1.3.15-3PGDG.rhel9.9.aarch64.rpm) |
| `rum_15` | `1.3.15` | [el10.x86_64](/os/el10.x86_64) | pgdg | 113.1 KiB | [rum_15-1.3.15-3PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/rum_15-1.3.15-3PGDG.rhel10.3.x86_64.rpm) |
| `rum_15` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.1 KiB | [rum_15-1.3.15-3PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/rum_15-1.3.15-3PGDG.rhel10.3.aarch64.rpm) |
| `rum_15` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.2 KiB | [rum_15-1.3.15-3PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/rum_15-1.3.15-3PGDG.rhel10.2.aarch64.rpm) |
| `rum_15` | `1.3.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [rum_15-1.3.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/rum_15-1.3.15-1PGDG.rhel10.2.aarch64.rpm) |
| `rum_15` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.1 KiB | [rum_15-1.3.14-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/rum_15-1.3.14-4PGDG.rhel10.2.aarch64.rpm) |
| `rum_15` | `1.3.14` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.3 KiB | [rum_15-1.3.14-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/rum_15-1.3.14-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 288.9 KiB | [postgresql-15-rum_1.3.15-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg12+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 288.7 KiB | [postgresql-15-rum_1.3.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg12+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 278.2 KiB | [postgresql-15-rum_1.3.15-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg12+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 277.9 KiB | [postgresql-15-rum_1.3.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg12+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 288.9 KiB | [postgresql-15-rum_1.3.15-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg13+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 289.6 KiB | [postgresql-15-rum_1.3.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg13+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 278.5 KiB | [postgresql-15-rum_1.3.15-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg13+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 278.7 KiB | [postgresql-15-rum_1.3.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg13+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 327.2 KiB | [postgresql-15-rum_1.3.15-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 327.0 KiB | [postgresql-15-rum_1.3.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 316.5 KiB | [postgresql-15-rum_1.3.15-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 316.2 KiB | [postgresql-15-rum_1.3.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 289.2 KiB | [postgresql-15-rum_1.3.15-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 288.9 KiB | [postgresql-15-rum_1.3.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 279.0 KiB | [postgresql-15-rum_1.3.15-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 279.2 KiB | [postgresql-15-rum_1.3.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 287.1 KiB | [postgresql-15-rum_1.3.15-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 287.3 KiB | [postgresql-15-rum_1.3.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 276.9 KiB | [postgresql-15-rum_1.3.15-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-rum` | `1.3.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 276.8 KiB | [postgresql-15-rum_1.3.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-rum/postgresql-15-rum_1.3.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgrespro/rum" title="Repository" icon="github" subtitle="github.com/postgrespro/rum" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="rum-1.3.15.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg rum;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install rum;		# install via package name, for the active PG version

pig install rum -v 19;   # install for PG 19
pig install rum -v 18;   # install for PG 18
pig install rum -v 17;   # install for PG 17
pig install rum -v 16;   # install for PG 16
pig install rum -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION rum;
```




## Usage

> [rum: RUM index access method](https://github.com/postgrespro/rum)

RUM is an index access method that extends GIN by storing additional information in the posting tree. This enables direct access to positional data, avoiding extra heap scans for ranking, phrase searches, and timestamp ordering.

### Index Creation

```sql
CREATE INDEX idx ON table_name USING rum (column operator_class);
```

With addon operators (e.g., ordering by a timestamp alongside full-text search):

```sql
CREATE INDEX tsts_idx ON tsts USING rum (t rum_tsvector_addon_ops, d)
    WITH (attach = 'd', to = 't');
```

### Operator Classes

| Operator Class | Description |
|---------------|-------------|
| `rum_tsvector_ops` | Stores tsvector lexemes with positions. Supports `<=>` ordering and prefix search. |
| `rum_tsvector_hash_ops` | Stores hashed tsvector lexemes with positions. Supports `<=>` ordering, no prefix search. |
| `rum_tsvector_addon_ops` | Combines tsvector with additional fields (timestamps, integers, etc.) for filtering and ordering. |
| `rum_tsvector_hash_addon_ops` | Hashed variant supporting addon fields, no prefix search. |
| `rum_tsquery_ops` | Stores tsquery branches for fast query matching against indexed documents. |
| `rum_anyarray_ops` | Indexes array types. Supports `&&`, `@>`, `<@`, `=`, `%` and `<=>` ordering. |
| `rum_anyarray_addon_ops` | Combines array elements with additional fields. |
| `rum_TYPE_ops` | Generic ops for int2, int4, int8, float4, float8, money, oid, time, timetz, date, interval, macaddr, inet, cidr, text, varchar, char, bytea, bit, varbit, numeric, timestamp, timestamptz. |

### Ordering Operators

| Operator | Description |
|----------|-------------|
| `<=>` | Distance operator for tsvector, timestamp, numeric types, arrays |
| `<=\|` | Left-side distance for timestamp, int, float, money, oid |
| `\|=>` | Right-side distance for timestamp, int, float, money, oid |

### Examples

Full-text search with ranking:

```sql
SELECT t, a <=> to_tsquery('english', 'beautiful | place') AS rank
FROM test_rum
WHERE a @@ to_tsquery('english', 'beautiful | place')
ORDER BY a <=> to_tsquery('english', 'beautiful | place');
```

Timestamp-ordered full-text search:

```sql
SELECT id, d, d <=> '2016-05-16 14:21:25' FROM tsts
WHERE t @@ 'wr&qh'
ORDER BY d <=> '2016-05-16 14:21:25'
LIMIT 5;
```

Array matching with distance ordering:

```sql
SELECT * FROM test_array
WHERE i && '{1}'
ORDER BY i <=> '{1}' ASC;
```
