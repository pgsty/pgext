---
title: "first_last_agg"
linkTitle: "first_last_agg"
description: "first() and last() aggregate functions"
weight: 4710
categories: ["FUNC"]
width: full
---

[**first_last_agg**](https://github.com/wulczer/first_last_agg) : first() and last() aggregate functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4710** | {{< badge content="first_last_agg" link="https://github.com/wulczer/first_last_agg" >}} | {{< ext "first_last_agg" >}} | `0.1.4` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "arraymath" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `0.1.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `first_last_agg` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.4` | {{< bg "18" "first_last_agg_18" "green" >}} {{< bg "17" "first_last_agg_17" "green" >}} {{< bg "16" "first_last_agg_16" "green" >}} {{< bg "15" "first_last_agg_15" "green" >}} {{< bg "14" "first_last_agg_14" "green" >}} | `first_last_agg_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.1.4` | {{< bg "18" "postgresql-18-first-last-agg" "green" >}} {{< bg "17" "postgresql-17-first-last-agg" "green" >}} {{< bg "16" "postgresql-16-first-last-agg" "green" >}} {{< bg "15" "postgresql-15-first-last-agg" "green" >}} {{< bg "14" "postgresql-14-first-last-agg" "green" >}} | `postgresql-$v-first-last-agg` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.4" "first_last_agg_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 0.1.4" "postgresql-18-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-17-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-16-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-15-first-last-agg : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.1.4" "postgresql-14-first-last-agg : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `first_last_agg_18` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/first_last_agg_18-0.1.4-1PIGSTY.el8.x86_64.rpm) |
| `first_last_agg_18` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.0 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/first_last_agg_18-0.1.4-1PIGSTY.el8.aarch64.rpm) |
| `first_last_agg_18` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.4 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/first_last_agg_18-0.1.4-1PIGSTY.el9.x86_64.rpm) |
| `first_last_agg_18` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/first_last_agg_18-0.1.4-1PIGSTY.el9.aarch64.rpm) |
| `first_last_agg_18` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/first_last_agg_18-0.1.4-1PIGSTY.el10.x86_64.rpm) |
| `first_last_agg_18` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [first_last_agg_18-0.1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/first_last_agg_18-0.1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-first-last-agg` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.6 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.6 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.4 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.5 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.6 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.5 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.5 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb) |
| `postgresql-18-first-last-agg` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-18-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `first_last_agg_17` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/first_last_agg_17-0.1.4-1PIGSTY.el8.x86_64.rpm) |
| `first_last_agg_17` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.0 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/first_last_agg_17-0.1.4-1PIGSTY.el8.aarch64.rpm) |
| `first_last_agg_17` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.4 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/first_last_agg_17-0.1.4-1PIGSTY.el9.x86_64.rpm) |
| `first_last_agg_17` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/first_last_agg_17-0.1.4-1PIGSTY.el9.aarch64.rpm) |
| `first_last_agg_17` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/first_last_agg_17-0.1.4-1PIGSTY.el10.x86_64.rpm) |
| `first_last_agg_17` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [first_last_agg_17-0.1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/first_last_agg_17-0.1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-first-last-agg` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.6 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.6 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.4 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.6 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.6 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.5 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.6 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb) |
| `postgresql-17-first-last-agg` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-17-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `first_last_agg_16` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/first_last_agg_16-0.1.4-1PIGSTY.el8.x86_64.rpm) |
| `first_last_agg_16` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.0 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/first_last_agg_16-0.1.4-1PIGSTY.el8.aarch64.rpm) |
| `first_last_agg_16` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.4 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/first_last_agg_16-0.1.4-1PIGSTY.el9.x86_64.rpm) |
| `first_last_agg_16` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/first_last_agg_16-0.1.4-1PIGSTY.el9.aarch64.rpm) |
| `first_last_agg_16` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/first_last_agg_16-0.1.4-1PIGSTY.el10.x86_64.rpm) |
| `first_last_agg_16` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [first_last_agg_16-0.1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/first_last_agg_16-0.1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-first-last-agg` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.6 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.6 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.4 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.6 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.6 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.5 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.6 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb) |
| `postgresql-16-first-last-agg` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-16-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `first_last_agg_15` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/first_last_agg_15-0.1.4-1PIGSTY.el8.x86_64.rpm) |
| `first_last_agg_15` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.0 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/first_last_agg_15-0.1.4-1PIGSTY.el8.aarch64.rpm) |
| `first_last_agg_15` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.4 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/first_last_agg_15-0.1.4-1PIGSTY.el9.x86_64.rpm) |
| `first_last_agg_15` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/first_last_agg_15-0.1.4-1PIGSTY.el9.aarch64.rpm) |
| `first_last_agg_15` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/first_last_agg_15-0.1.4-1PIGSTY.el10.x86_64.rpm) |
| `first_last_agg_15` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [first_last_agg_15-0.1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/first_last_agg_15-0.1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-first-last-agg` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.6 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.5 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.6 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.4 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.6 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.6 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.5 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.5 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb) |
| `postgresql-15-first-last-agg` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-15-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `first_last_agg_14` | `0.1.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.7 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/first_last_agg_14-0.1.4-1PIGSTY.el8.x86_64.rpm) |
| `first_last_agg_14` | `0.1.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.0 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/first_last_agg_14-0.1.4-1PIGSTY.el8.aarch64.rpm) |
| `first_last_agg_14` | `0.1.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.4 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/first_last_agg_14-0.1.4-1PIGSTY.el9.x86_64.rpm) |
| `first_last_agg_14` | `0.1.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/first_last_agg_14-0.1.4-1PIGSTY.el9.aarch64.rpm) |
| `first_last_agg_14` | `0.1.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.4 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/first_last_agg_14-0.1.4-1PIGSTY.el10.x86_64.rpm) |
| `first_last_agg_14` | `0.1.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 KiB | [first_last_agg_14-0.1.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/first_last_agg_14-0.1.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-first-last-agg` | `0.1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 8.5 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_amd64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 8.4 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg12+1_arm64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 8.5 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_amd64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 8.5 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg13+1_arm64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 8.4 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_amd64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 8.6 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg22.04+1_arm64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 8.6 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_amd64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 8.5 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg24.04+1_arm64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 8.5 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_amd64.deb) |
| `postgresql-14-first-last-agg` | `0.1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 8.9 KiB | [postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/f/first-last-agg/postgresql-14-first-last-agg_0.1.4-4-gd63ea3b-9.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/wulczer/first_last_agg" title="Repository" icon="github" subtitle="github.com/wulczer/first_last_agg" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="first_last_agg-0.1.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg first_last_agg;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install first_last_agg;		# install via package name, for the active PG version

pig install first_last_agg -v 18;   # install for PG 18
pig install first_last_agg -v 17;   # install for PG 17
pig install first_last_agg -v 16;   # install for PG 16
pig install first_last_agg -v 15;   # install for PG 15
pig install first_last_agg -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION first_last_agg;
```



## Usage

> [first_last_agg: first and last aggregate functions for PostgreSQL](https://github.com/wulczer/first_last_agg)

Provides `first` and `last` aggregate functions that return the first or last value in a group, operating on any element type.

```sql
CREATE EXTENSION first_last_agg;
```

### Functions

| Function | Description |
|---|---|
| `first(anyelement)` | Return the first value in the group |
| `last(anyelement)` | Return the last value in the group |

### Examples

```sql
-- Get the first and last values (use ORDER BY for deterministic results)
SELECT id, first(val ORDER BY ts), last(val ORDER BY ts)
FROM events
GROUP BY id;

-- Without ORDER BY, the ordering within groups is undefined
SELECT department, first(salary ORDER BY hire_date)
FROM employees
GROUP BY department;
```
