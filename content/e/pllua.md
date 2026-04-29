---
title: "pllua"
linkTitle: "pllua"
description: "Lua as a procedural language"
weight: 3020
categories: ["LANG"]
width: full
---

[**pllua**](https://github.com/pllua/pllua) : Lua as a procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3020** | {{< badge content="pllua" link="https://github.com/pllua/pllua" >}} | {{< ext "pllua" >}} | `2.0.12` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pg_catalog` |
|    **Need By**    | {{< ext "hstore_pllua" >}} |
|   **See Also**    | {{< ext "plperl" >}} {{< ext "plpgsql" >}} {{< ext "plpython3u" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pljava" >}} {{< ext "plperlu" >}} |
|    **Siblings**   | {{< ext "hstore_pllua" >}} {{< ext "plluau" >}} {{< ext "hstore_plluau" >}} |

> [!Note] missing pg12-15 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pllua` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "18" "pllua_18" "red" >}} {{< bg "17" "pllua_17" "green" >}} {{< bg "16" "pllua_16" "green" >}} {{< bg "15" "pllua_15" "green" >}} {{< bg "14" "pllua_14" "green" >}} | `pllua_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "18" "postgresql-18-pllua" "green" >}} {{< bg "17" "postgresql-17-pllua" "green" >}} {{< bg "16" "postgresql-16-pllua" "green" >}} {{< bg "15" "postgresql-15-pllua" "green" >}} {{< bg "14" "postgresql-14-pllua" "green" >}} | `postgresql-$v-pllua` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 2" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pllua` | `2.0.12` | [d12.x86_64](/os/d12.x86_64) | pgdg | 347.6 KiB | [postgresql-18-pllua_2.0.12-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg12+1_amd64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [d12.aarch64](/os/d12.aarch64) | pgdg | 336.3 KiB | [postgresql-18-pllua_2.0.12-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg12+1_arm64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [d13.x86_64](/os/d13.x86_64) | pgdg | 348.2 KiB | [postgresql-18-pllua_2.0.12-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg13+1_amd64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.2 KiB | [postgresql-18-pllua_2.0.12-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg13+1_arm64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u22.x86_64](/os/u22.x86_64) | pgdg | 354.5 KiB | [postgresql-18-pllua_2.0.12-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u22.aarch64](/os/u22.aarch64) | pgdg | 341.8 KiB | [postgresql-18-pllua_2.0.12-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u24.x86_64](/os/u24.x86_64) | pgdg | 347.6 KiB | [postgresql-18-pllua_2.0.12-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.4 KiB | [postgresql-18-pllua_2.0.12-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg24.04+1_arm64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u26.x86_64](/os/u26.x86_64) | pgdg | 345.1 KiB | [postgresql-18-pllua_2.0.12-7.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg26.04+1_amd64.deb) |
| `postgresql-18-pllua` | `2.0.12` | [u26.aarch64](/os/u26.aarch64) | pgdg | 332.1 KiB | [postgresql-18-pllua_2.0.12-7.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-18-pllua_2.0.12-7.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_17` | `2.0.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.4 KiB | [pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pllua_17-2.0.12-3PGDG.rhel8.x86_64.rpm) |
| `pllua_17` | `2.0.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 110.9 KiB | [pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pllua_17-2.0.12-3PGDG.rhel8.aarch64.rpm) |
| `pllua_17` | `2.0.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 120.4 KiB | [pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pllua_17-2.0.12-3PGDG.rhel9.x86_64.rpm) |
| `pllua_17` | `2.0.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.3 KiB | [pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pllua_17-2.0.12-3PGDG.rhel9.aarch64.rpm) |
| `pllua_17` | `2.0.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 122.4 KiB | [pllua_17-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pllua_17-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_17` | `2.0.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.8 KiB | [pllua_17-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pllua_17-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pllua` | `2.0.12` | [d12.x86_64](/os/d12.x86_64) | pgdg | 347.0 KiB | [postgresql-17-pllua_2.0.12-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg12+1_amd64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.6 KiB | [postgresql-17-pllua_2.0.12-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg12+1_arm64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [d13.x86_64](/os/d13.x86_64) | pgdg | 347.9 KiB | [postgresql-17-pllua_2.0.12-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg13+1_amd64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [d13.aarch64](/os/d13.aarch64) | pgdg | 336.0 KiB | [postgresql-17-pllua_2.0.12-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg13+1_arm64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u22.x86_64](/os/u22.x86_64) | pgdg | 391.0 KiB | [postgresql-17-pllua_2.0.12-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u22.aarch64](/os/u22.aarch64) | pgdg | 379.2 KiB | [postgresql-17-pllua_2.0.12-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u24.x86_64](/os/u24.x86_64) | pgdg | 347.0 KiB | [postgresql-17-pllua_2.0.12-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.0 KiB | [postgresql-17-pllua_2.0.12-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg24.04+1_arm64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u26.x86_64](/os/u26.x86_64) | pgdg | 344.7 KiB | [postgresql-17-pllua_2.0.12-7.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg26.04+1_amd64.deb) |
| `postgresql-17-pllua` | `2.0.12` | [u26.aarch64](/os/u26.aarch64) | pgdg | 331.5 KiB | [postgresql-17-pllua_2.0.12-7.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-17-pllua_2.0.12-7.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_16` | `2.0.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.2 KiB | [pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pllua_16-2.0.12-1PGDG.rhel8.x86_64.rpm) |
| `pllua_16` | `2.0.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 110.6 KiB | [pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pllua_16-2.0.12-1PGDG.rhel8.aarch64.rpm) |
| `pllua_16` | `2.0.12` | [el9.x86_64](/os/el9.x86_64) | pgdg | 120.1 KiB | [pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pllua_16-2.0.12-1PGDG.rhel9.x86_64.rpm) |
| `pllua_16` | `2.0.12` | [el9.aarch64](/os/el9.aarch64) | pgdg | 115.5 KiB | [pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pllua_16-2.0.12-1PGDG.rhel9.aarch64.rpm) |
| `pllua_16` | `2.0.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 122.7 KiB | [pllua_16-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pllua_16-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_16` | `2.0.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.7 KiB | [pllua_16-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pllua_16-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pllua` | `2.0.12` | [d12.x86_64](/os/d12.x86_64) | pgdg | 346.9 KiB | [postgresql-16-pllua_2.0.12-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg12+1_amd64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [d12.aarch64](/os/d12.aarch64) | pgdg | 335.7 KiB | [postgresql-16-pllua_2.0.12-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg12+1_arm64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [d13.x86_64](/os/d13.x86_64) | pgdg | 347.9 KiB | [postgresql-16-pllua_2.0.12-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg13+1_amd64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [d13.aarch64](/os/d13.aarch64) | pgdg | 335.8 KiB | [postgresql-16-pllua_2.0.12-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg13+1_arm64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u22.x86_64](/os/u22.x86_64) | pgdg | 389.6 KiB | [postgresql-16-pllua_2.0.12-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u22.aarch64](/os/u22.aarch64) | pgdg | 377.5 KiB | [postgresql-16-pllua_2.0.12-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u24.x86_64](/os/u24.x86_64) | pgdg | 347.1 KiB | [postgresql-16-pllua_2.0.12-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u24.aarch64](/os/u24.aarch64) | pgdg | 335.1 KiB | [postgresql-16-pllua_2.0.12-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg24.04+1_arm64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u26.x86_64](/os/u26.x86_64) | pgdg | 344.8 KiB | [postgresql-16-pllua_2.0.12-7.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg26.04+1_amd64.deb) |
| `postgresql-16-pllua` | `2.0.12` | [u26.aarch64](/os/u26.aarch64) | pgdg | 331.8 KiB | [postgresql-16-pllua_2.0.12-7.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-16-pllua_2.0.12-7.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_15` | `2.0.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.9 KiB | [pllua_15-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_15` | `2.0.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.9 KiB | [pllua_15-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pllua_15-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_15` | `2.0.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.5 KiB | [pllua_15-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_15` | `2.0.10` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.8 KiB | [pllua_15-2.0.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pllua_15-2.0.10-1.rhel9.x86_64.rpm) |
| `pllua_15` | `2.0.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 125.8 KiB | [pllua_15-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pllua_15-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_15` | `2.0.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.4 KiB | [pllua_15-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pllua_15-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pllua` | `2.0.12` | [d12.x86_64](/os/d12.x86_64) | pgdg | 348.7 KiB | [postgresql-15-pllua_2.0.12-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg12+1_amd64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [d12.aarch64](/os/d12.aarch64) | pgdg | 337.0 KiB | [postgresql-15-pllua_2.0.12-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg12+1_arm64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [d13.x86_64](/os/d13.x86_64) | pgdg | 349.6 KiB | [postgresql-15-pllua_2.0.12-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg13+1_amd64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [d13.aarch64](/os/d13.aarch64) | pgdg | 337.4 KiB | [postgresql-15-pllua_2.0.12-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg13+1_arm64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u22.x86_64](/os/u22.x86_64) | pgdg | 392.5 KiB | [postgresql-15-pllua_2.0.12-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u22.aarch64](/os/u22.aarch64) | pgdg | 380.0 KiB | [postgresql-15-pllua_2.0.12-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u24.x86_64](/os/u24.x86_64) | pgdg | 348.7 KiB | [postgresql-15-pllua_2.0.12-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u24.aarch64](/os/u24.aarch64) | pgdg | 337.2 KiB | [postgresql-15-pllua_2.0.12-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg24.04+1_arm64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u26.x86_64](/os/u26.x86_64) | pgdg | 346.1 KiB | [postgresql-15-pllua_2.0.12-7.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg26.04+1_amd64.deb) |
| `postgresql-15-pllua` | `2.0.12` | [u26.aarch64](/os/u26.aarch64) | pgdg | 333.7 KiB | [postgresql-15-pllua_2.0.12-7.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-15-pllua_2.0.12-7.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pllua_14` | `2.0.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 121.1 KiB | [pllua_14-2.0.11-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.11-1.rhel8.x86_64.rpm) |
| `pllua_14` | `2.0.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.9 KiB | [pllua_14-2.0.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pllua_14-2.0.10-1.rhel8.x86_64.rpm) |
| `pllua_14` | `2.0.11` | [el9.x86_64](/os/el9.x86_64) | pgdg | 123.2 KiB | [pllua_14-2.0.11-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pllua_14-2.0.11-1.rhel9.x86_64.rpm) |
| `pllua_14` | `2.0.12` | [el10.x86_64](/os/el10.x86_64) | pgdg | 125.8 KiB | [pllua_14-2.0.12-4PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pllua_14-2.0.12-4PGDG.rhel10.x86_64.rpm) |
| `pllua_14` | `2.0.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.6 KiB | [pllua_14-2.0.12-4PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pllua_14-2.0.12-4PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pllua` | `2.0.12` | [d12.x86_64](/os/d12.x86_64) | pgdg | 348.5 KiB | [postgresql-14-pllua_2.0.12-7.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg12+1_amd64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [d12.aarch64](/os/d12.aarch64) | pgdg | 336.3 KiB | [postgresql-14-pllua_2.0.12-7.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg12+1_arm64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [d13.x86_64](/os/d13.x86_64) | pgdg | 349.3 KiB | [postgresql-14-pllua_2.0.12-7.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg13+1_amd64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [d13.aarch64](/os/d13.aarch64) | pgdg | 337.2 KiB | [postgresql-14-pllua_2.0.12-7.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg13+1_arm64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u22.x86_64](/os/u22.x86_64) | pgdg | 388.1 KiB | [postgresql-14-pllua_2.0.12-7.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u22.aarch64](/os/u22.aarch64) | pgdg | 375.8 KiB | [postgresql-14-pllua_2.0.12-7.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u24.x86_64](/os/u24.x86_64) | pgdg | 348.9 KiB | [postgresql-14-pllua_2.0.12-7.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u24.aarch64](/os/u24.aarch64) | pgdg | 336.7 KiB | [postgresql-14-pllua_2.0.12-7.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg24.04+1_arm64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u26.x86_64](/os/u26.x86_64) | pgdg | 346.3 KiB | [postgresql-14-pllua_2.0.12-7.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg26.04+1_amd64.deb) |
| `postgresql-14-pllua` | `2.0.12` | [u26.aarch64](/os/u26.aarch64) | pgdg | 333.1 KiB | [postgresql-14-pllua_2.0.12-7.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-pllua/postgresql-14-pllua_2.0.12-7.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pllua/pllua" title="Repository" icon="github" subtitle="github.com/pllua/pllua" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pllua;		# install via package name, for the active PG version

pig install pllua -v 18;   # install for PG 18
pig install pllua -v 17;   # install for PG 17
pig install pllua -v 16;   # install for PG 16
pig install pllua -v 15;   # install for PG 15
pig install pllua -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pllua;
```




## Usage

> [pllua: Lua as a procedural language](https://github.com/pllua/pllua)

`pllua` enables writing PostgreSQL functions in Lua (5.3, 5.4, or LuaJIT 2.1).

```sql
CREATE EXTENSION pllua;
```

### Create Functions

```sql
CREATE FUNCTION lua_max(a integer, b integer) RETURNS integer LANGUAGE pllua AS $$
  if a > b then return a else return b end
$$;

CREATE FUNCTION hello(name text) RETURNS text LANGUAGE pllua AS $$
  return "Hello, " .. name .. "!"
$$;
```

### Data Type Handling

Arguments are automatically converted: integers/floats to Lua numbers, text/varchar to strings, booleans to Lua booleans, NULL to nil. Other types remain as datum objects.

Construct typed values with `pgtype`:

```lua
pgtype.numeric(1234)
pgtype.date('2017-12-01')
pgtype.array.integer(1, 2, 3, 4)
pgtype.numrange(1, 2)
```

### Composite Types (Rows)

```lua
row.columnname     -- access by name
row[3]             -- access by attribute number
for colname, value, attnum in pairs(row) do ... end
```

### Set-Returning Functions

```sql
CREATE FUNCTION generate_n(n integer) RETURNS SETOF integer LANGUAGE pllua AS $$
  for i = 1, n do
    coroutine.yield(i)
  end
$$;
```

### SPI Database Access

```lua
-- Simple query
local rows = spi.execute("SELECT * FROM mytable WHERE id = $1", 42)

-- Row iterator
for row in spi.rows("SELECT * FROM mytable") do
  print(row.name)
end

-- Prepared statements
local stmt = spi.prepare("SELECT * FROM users WHERE id = $1", {'integer'})
local result = stmt:execute(42)
for row in stmt:rows(42) do ... end
```

### Cursors

```lua
local cursor = spi.newcursor()
cursor:open("SELECT * FROM items")
local rows = cursor:fetch(10)
cursor:move(5)
cursor:close()
```

### Trigger Functions

```sql
CREATE FUNCTION my_trigger() RETURNS trigger LANGUAGE pllua AS $$
  function(trigger, old, new)
    trigger.row = new
    return trigger.row
  end
$$;
```

Trigger fields: `trigger.event` (INSERT/UPDATE/DELETE), `trigger.when` (BEFORE/AFTER), `trigger.level` (ROW/STATEMENT), `trigger.new`, `trigger.old`, `trigger.row`.

### Error Handling

```lua
spi.error('division_by_zero', 'Cannot divide by zero')
spi.notice('informational message')
spi.warning('warning message')

-- Subtransactions with pcall
local ok, err = pcall(function()
  spi.execute("INSERT INTO mytable VALUES ($1)", val)
end)
```

### Logging

```lua
print("info message")
spi.debug("debug")
spi.notice("notice")
spi.warning("warning")
spi.error("error")
```
