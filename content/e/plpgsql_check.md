---
title: "plpgsql_check"
linkTitle: "plpgsql_check"
description: "Additional validation, profiling, and diagnostics for PL/pgSQL functions"
weight: 3060
categories: ["LANG"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PGDG"]
page_width: full
---

[**plpgsql_check**](https://github.com/okbob/plpgsql_check) : Additional validation, profiling, and diagnostics for PL/pgSQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3060** | {{< badge content="plpgsql_check" link="https://github.com/okbob/plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | `2.11` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pldbgapi" >}} {{< ext "plprofiler" >}} {{< ext "pglinter" >}} {{< ext "pg_stat_plans" >}} {{< ext "auto_explain" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_show_plans" >}} {{< ext "pg_store_plans" >}} {{< ext "hypopg" >}} {{< ext "index_advisor" >}} |

> [!Note] pkg 2.10.16; SQL 2.10; upstream PG14-20 includes development branches; packages PG14-18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.11` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `plpgsql_check` | `plpgsql` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.11.0` | {{< bg "19" "plpgsql_check_19" "green" >}} {{< bg "18" "plpgsql_check_18" "green" >}} {{< bg "17" "plpgsql_check_17" "green" >}} {{< bg "16" "plpgsql_check_16" "green" >}} {{< bg "15" "plpgsql_check_15" "green" >}} | `plpgsql_check_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.11.0` | {{< bg "19" "postgresql-19-plpgsql-check" "green" >}} {{< bg "18" "postgresql-18-plpgsql-check" "green" >}} {{< bg "17" "postgresql-17-plpgsql-check" "green" >}} {{< bg "16" "postgresql-16-plpgsql-check" "green" >}} {{< bg "15" "postgresql-15-plpgsql-check" "green" >}} | `postgresql-$v-plpgsql-check` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "plpgsql_check_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 26" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 40" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 48" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "plpgsql_check_19 : N/A 0" "gray" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 22" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 27" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 41" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 48" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 3" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_19 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_18 : AVAIL 30" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_17 : AVAIL 31" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_16 : AVAIL 31" "blue" >}} | {{< bg "PGDG 2.11.0" "plpgsql_check_15 : AVAIL 31" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2.10.16" "postgresql-19-plpgsql-check : AVAIL 3" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-18-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-17-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-16-plpgsql-check : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.10.16" "postgresql-15-plpgsql-check : AVAIL 4" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_19` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.6 KiB | [plpgsql_check_19-2.11.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/plpgsql_check_19-2.11.0-1PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_19` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 126.6 KiB | [plpgsql_check_19-2.10.16-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/plpgsql_check_19-2.10.16-2PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_19` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 128.2 KiB | [plpgsql_check_19-2.11.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/plpgsql_check_19-2.11.0-1PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_19` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 121.5 KiB | [plpgsql_check_19-2.10.16-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/plpgsql_check_19-2.10.16-2PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_19` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.7 KiB | [plpgsql_check_19-2.11.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/plpgsql_check_19-2.11.0-1PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_19` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.3 KiB | [plpgsql_check_19-2.10.16-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/plpgsql_check_19-2.10.16-2PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_19` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 130.1 KiB | [plpgsql_check_19-2.11.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/plpgsql_check_19-2.11.0-1PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_19` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.1 KiB | [plpgsql_check_19-2.10.16-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/plpgsql_check_19-2.10.16-2PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.7 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.7 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.8 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.6 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.7 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.5 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 330.3 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 330.2 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 330.1 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 318.0 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.8 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.6 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 340.2 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 340.3 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 340.3 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 327.1 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 326.8 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 327.0 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 329.1 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.9 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 329.2 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.2 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.1 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 315.9 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.9 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.9 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.7 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.5 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.6 KiB | [postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb) |
| `postgresql-19-plpgsql-check` | `2.10.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.6 KiB | [postgresql-19-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-19-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_18` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.3 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.11.0-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.2 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.16-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 346.8 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_18-2.10.16-1PGSTY.el8.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.2 KiB | [plpgsql_check_18-2.10.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.16-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 131.5 KiB | [plpgsql_check_18-2.10.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.14-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.3 KiB | [plpgsql_check_18-2.10.13-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.13-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.0 KiB | [plpgsql_check_18-2.10.12-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.12-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.3 KiB | [plpgsql_check_18-2.10.11-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.11-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.4 KiB | [plpgsql_check_18-2.10.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.7 KiB | [plpgsql_check_18-2.10.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.6 KiB | [plpgsql_check_18-2.10.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.2 KiB | [plpgsql_check_18-2.10.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.10.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.3 KiB | [plpgsql_check_18-2.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_18-2.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.7 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.8 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.0 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.2 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.11.0-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 123.0 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.16-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 334.8 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_18-2.10.16-1PGSTY.el8.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.6 KiB | [plpgsql_check_18-2.10.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.16-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 121.9 KiB | [plpgsql_check_18-2.10.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.14-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.7 KiB | [plpgsql_check_18-2.10.13-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.13-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.6 KiB | [plpgsql_check_18-2.10.12-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.12-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.1 KiB | [plpgsql_check_18-2.10.11-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.11-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.9 KiB | [plpgsql_check_18-2.10.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.7-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.2 KiB | [plpgsql_check_18-2.10.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.8 KiB | [plpgsql_check_18-2.10.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.6 KiB | [plpgsql_check_18-2.10.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 115.3 KiB | [plpgsql_check_18-2.10.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.10.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.4 KiB | [plpgsql_check_18-2.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.2 KiB | [plpgsql_check_18-2.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.2 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.9 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.5 KiB | [plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.4 KiB | [plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.4 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 133.8 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.11.0-1PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 126.8 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.10.16-2PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 348.1 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_18-2.10.16-1PGSTY.el9.x86_64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 128.3 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.11.0-1PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 121.6 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.10.16-2PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 340.4 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_18-2.10.16-1PGSTY.el9.aarch64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.1 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.11.0-1PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.0 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plpgsql_check_18-2.10.16-2PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 350.2 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_18-2.10.16-1PGSTY.el10.x86_64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 130.1 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.11.0-1PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_18` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 130.1 KiB | [plpgsql_check_18-2.11.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.11.0-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.2 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.16-2PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.2 KiB | [plpgsql_check_18-2.10.16-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.16-2PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 342.2 KiB | [plpgsql_check_18-2.10.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_18-2.10.16-1PGSTY.el10.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.9 KiB | [plpgsql_check_18-2.10.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.16-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.9 KiB | [plpgsql_check_18-2.10.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.15-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 121.3 KiB | [plpgsql_check_18-2.10.12-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.12-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.8 KiB | [plpgsql_check_18-2.10.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 118.3 KiB | [plpgsql_check_18-2.10.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.7-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.7 KiB | [plpgsql_check_18-2.10.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.4-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [plpgsql_check_18-2.10.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.2 KiB | [plpgsql_check_18-2.10.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 115.7 KiB | [plpgsql_check_18-2.10.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.10.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_18-2.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_18-2.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_18-2.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.1 KiB | [plpgsql_check_18-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.10-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.10-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.7 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_18-2.8.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.8-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.5-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.5-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.4-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.4-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_18` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.1 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.2 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 326.9 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.1 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 317.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.8 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 315.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.7 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.4 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 318.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 318.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 316.1 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.9 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 339.5 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 339.3 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 356.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 339.4 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 327.1 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 326.9 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 347.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 327.0 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.5 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.5 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.3 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.4 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 334.6 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.4 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 325.9 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.0 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 338.9 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.1 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.8 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 314.2 KiB | [postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 332.3 KiB | [postgresql-18-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-18-plpgsql-check` | `2.10.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.8 KiB | [postgresql-18-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_17` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.2 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.11.0-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.3 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.16-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 346.2 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_17-2.10.16-1PGSTY.el8.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.2 KiB | [plpgsql_check_17-2.10.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.16-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 131.4 KiB | [plpgsql_check_17-2.10.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.14-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.0 KiB | [plpgsql_check_17-2.10.13-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.13-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.9 KiB | [plpgsql_check_17-2.10.12-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.12-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.4 KiB | [plpgsql_check_17-2.10.11-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.11-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.3 KiB | [plpgsql_check_17-2.10.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.8 KiB | [plpgsql_check_17-2.10.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.6 KiB | [plpgsql_check_17-2.10.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.3 KiB | [plpgsql_check_17-2.10.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.10.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.2 KiB | [plpgsql_check_17-2.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.0 KiB | [plpgsql_check_17-2.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.1 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.3 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.1 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.0 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.8 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.3 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.0 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.4 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.0 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.11.0-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.16-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 334.5 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_17-2.10.16-1PGSTY.el8.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.4 KiB | [plpgsql_check_17-2.10.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.16-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 121.8 KiB | [plpgsql_check_17-2.10.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.14-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.6 KiB | [plpgsql_check_17-2.10.13-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.13-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.4 KiB | [plpgsql_check_17-2.10.12-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.12-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.9 KiB | [plpgsql_check_17-2.10.11-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.11-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.8 KiB | [plpgsql_check_17-2.10.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.7-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.1 KiB | [plpgsql_check_17-2.10.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.7 KiB | [plpgsql_check_17-2.10.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.4 KiB | [plpgsql_check_17-2.10.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 115.1 KiB | [plpgsql_check_17-2.10.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.10.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_17-2.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_17-2.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.3 KiB | [plpgsql_check_17-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.4 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.4 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.7 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.11.0-1PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 126.1 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.10.16-2PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 347.4 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_17-2.10.16-1PGSTY.el9.x86_64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 128.0 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.11.0-1PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 121.4 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.10.16-2PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 340.1 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_17-2.10.16-1PGSTY.el9.aarch64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 137.0 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.11.0-1PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.4 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plpgsql_check_17-2.10.16-2PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 349.5 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_17-2.10.16-1PGSTY.el10.x86_64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.9 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.11.0-1PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_17` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.9 KiB | [plpgsql_check_17-2.11.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.11.0-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.1 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.16-2PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.1 KiB | [plpgsql_check_17-2.10.16-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.16-2PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 342.2 KiB | [plpgsql_check_17-2.10.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_17-2.10.16-1PGSTY.el10.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_17-2.10.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.16-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_17-2.10.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.15-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 121.1 KiB | [plpgsql_check_17-2.10.12-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.12-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.6 KiB | [plpgsql_check_17-2.10.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 118.2 KiB | [plpgsql_check_17-2.10.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.7-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [plpgsql_check_17-2.10.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.4-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [plpgsql_check_17-2.10.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.2 KiB | [plpgsql_check_17-2.10.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 115.9 KiB | [plpgsql_check_17-2.10.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.10.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.7 KiB | [plpgsql_check_17-2.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.5 KiB | [plpgsql_check_17-2.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.8 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_17-2.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_17-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.7 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_17-2.8.10-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.10-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_17-2.8.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.8-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.5-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.5-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_17-2.8.4-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.4-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_17` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.3 KiB | [plpgsql_check_17-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plpgsql_check_17-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 328.9 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 328.9 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 326.7 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 328.7 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.5 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 314.7 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.5 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.4 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.4 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.5 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.3 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 318.0 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.7 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 315.9 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.9 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 418.0 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 417.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 436.9 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 417.7 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 405.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 405.7 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 428.4 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 405.4 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.3 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.6 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.5 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.3 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.2 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.2 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 334.1 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.0 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.0 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.0 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 338.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 325.8 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 312.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 312.8 KiB | [postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 332.0 KiB | [postgresql-17-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-17-plpgsql-check` | `2.10.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 312.7 KiB | [postgresql-17-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_16` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.1 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.11.0-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.3 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.16-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 346.2 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_16-2.10.16-1PGSTY.el8.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.1 KiB | [plpgsql_check_16-2.10.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.16-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 131.4 KiB | [plpgsql_check_16-2.10.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.14-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.0 KiB | [plpgsql_check_16-2.10.13-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.13-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.9 KiB | [plpgsql_check_16-2.10.12-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.12-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.4 KiB | [plpgsql_check_16-2.10.11-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.11-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.3 KiB | [plpgsql_check_16-2.10.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.9 KiB | [plpgsql_check_16-2.10.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.7 KiB | [plpgsql_check_16-2.10.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.3 KiB | [plpgsql_check_16-2.10.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.10.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.3 KiB | [plpgsql_check_16-2.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.2 KiB | [plpgsql_check_16-2.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 119.9 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.4 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.1 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.2 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.0 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.7 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.2 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.1 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 111.9 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.4 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.0 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.3 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.0 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.8 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.7 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.8 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.6 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.6 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 102.0 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.6 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.3 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 100.2 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 129.0 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.11.0-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.9 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.16-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 334.4 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_16-2.10.16-1PGSTY.el8.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.5 KiB | [plpgsql_check_16-2.10.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.16-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 121.7 KiB | [plpgsql_check_16-2.10.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.14-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.6 KiB | [plpgsql_check_16-2.10.13-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.13-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.4 KiB | [plpgsql_check_16-2.10.12-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.12-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.9 KiB | [plpgsql_check_16-2.10.11-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.11-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.8 KiB | [plpgsql_check_16-2.10.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.7-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.1 KiB | [plpgsql_check_16-2.10.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.7 KiB | [plpgsql_check_16-2.10.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.5 KiB | [plpgsql_check_16-2.10.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 115.1 KiB | [plpgsql_check_16-2.10.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.10.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.2 KiB | [plpgsql_check_16-2.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_16-2.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.3 KiB | [plpgsql_check_16-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 105.0 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.5 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.6 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.3 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.2 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.1 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.9 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.7 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.4 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.8 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.3 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.7 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.5 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.4 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.1 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 93.0 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 134.0 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.11.0-1PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 126.2 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.10.16-2PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 347.7 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_16-2.10.16-1PGSTY.el9.x86_64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 128.0 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.11.0-1PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 121.4 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.10.16-2PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 340.2 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_16-2.10.16-1PGSTY.el9.aarch64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 137.0 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.11.0-1PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.4 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plpgsql_check_16-2.10.16-2PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 349.5 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_16-2.10.16-1PGSTY.el10.x86_64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.8 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.11.0-1PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_16` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.9 KiB | [plpgsql_check_16-2.11.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.11.0-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.2 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.16-2PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.3 KiB | [plpgsql_check_16-2.10.16-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.16-2PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 342.1 KiB | [plpgsql_check_16-2.10.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_16-2.10.16-1PGSTY.el10.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_16-2.10.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.16-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_16-2.10.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.15-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 121.1 KiB | [plpgsql_check_16-2.10.12-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.12-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.7 KiB | [plpgsql_check_16-2.10.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 118.3 KiB | [plpgsql_check_16-2.10.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.7-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.5 KiB | [plpgsql_check_16-2.10.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.4-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.3 KiB | [plpgsql_check_16-2.10.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.1 KiB | [plpgsql_check_16-2.10.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 115.7 KiB | [plpgsql_check_16-2.10.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.10.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.7 KiB | [plpgsql_check_16-2.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.6 KiB | [plpgsql_check_16-2.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.8 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.8 KiB | [plpgsql_check_16-2.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.9 KiB | [plpgsql_check_16-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_16-2.8.10-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.10-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.5 KiB | [plpgsql_check_16-2.8.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.8-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.5-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.5-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.2 KiB | [plpgsql_check_16-2.8.4-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.4-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 105.7 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_16` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.3 KiB | [plpgsql_check_16-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plpgsql_check_16-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 329.1 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 328.7 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 326.7 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 328.7 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.7 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.6 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 314.6 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 316.7 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.1 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.2 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 327.3 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 329.3 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.9 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.7 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 315.8 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 317.8 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 411.9 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 411.6 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 430.0 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 411.4 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 399.7 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 399.2 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 421.2 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 399.1 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.5 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.0 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 341.4 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 328.4 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.2 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.2 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 344.4 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 316.0 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.0 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 325.9 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 338.9 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 326.0 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.5 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 313.1 KiB | [postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 331.8 KiB | [postgresql-16-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-16-plpgsql-check` | `2.10.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 312.9 KiB | [postgresql-16-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_15` | `2.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 139.0 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.11.0-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.4 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.16-2PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pigsty | 348.3 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/plpgsql_check_15-2.10.16-1PGSTY.el8.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.x86_64](/os/el8.x86_64) | pgdg | 132.2 KiB | [plpgsql_check_15-2.10.16-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.16-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 131.4 KiB | [plpgsql_check_15-2.10.14-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.14-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.13` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.2 KiB | [plpgsql_check_15-2.10.13-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.13-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 130.0 KiB | [plpgsql_check_15-2.10.12-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.12-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.5 KiB | [plpgsql_check_15-2.10.11-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.11-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.4 KiB | [plpgsql_check_15-2.10.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 126.0 KiB | [plpgsql_check_15-2.10.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 125.7 KiB | [plpgsql_check_15-2.10.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 124.4 KiB | [plpgsql_check_15-2.10.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.10.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.4 KiB | [plpgsql_check_15-2.9.3-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.9.3-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.3 KiB | [plpgsql_check_15-2.9.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.9.2-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 120.2 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.9.1-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.5 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.10-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.3 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.8-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 116.1 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.5-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.8 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.4-1PGDG.rhel8.10.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.7 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 115.2 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 114.1 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el8.x86_64](/os/el8.x86_64) | pgdg | 113.9 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el8.x86_64](/os/el8.x86_64) | pgdg | 106.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.9 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.1 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.6 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.5 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.4 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.2 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 104.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.4 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 103.2 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.9 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 101.7 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 98.6 KiB | [plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.4 KiB | [plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 97.3 KiB | [plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 96.7 KiB | [plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.6 KiB | [plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 94.8 KiB | [plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 128.8 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.11.0-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.16-2PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pigsty | 336.2 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/plpgsql_check_15-2.10.16-1PGSTY.el8.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.4 KiB | [plpgsql_check_15-2.10.16-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.16-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 121.7 KiB | [plpgsql_check_15-2.10.14-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.14-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.13` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.5 KiB | [plpgsql_check_15-2.10.13-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.13-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 120.3 KiB | [plpgsql_check_15-2.10.12-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.12-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 119.9 KiB | [plpgsql_check_15-2.10.11-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.11-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.8 KiB | [plpgsql_check_15-2.10.7-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.7-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 117.1 KiB | [plpgsql_check_15-2.10.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.8 KiB | [plpgsql_check_15-2.10.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 116.6 KiB | [plpgsql_check_15-2.10.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 115.2 KiB | [plpgsql_check_15-2.10.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.10.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.2 KiB | [plpgsql_check_15-2.9.3-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.9.3-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.1 KiB | [plpgsql_check_15-2.9.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.9.2-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 111.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el8.aarch64](/os/el8.aarch64) | pgdg | 108.1 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.9 KiB | [plpgsql_check_15-2.8.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.5-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.8 KiB | [plpgsql_check_15-2.8.4-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.4-1PGDG.rhel8.10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.6 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 106.1 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.15` | [el8.aarch64](/os/el8.aarch64) | pgdg | 104.8 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.14` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.12` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.8 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.11` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.8` | [el8.aarch64](/os/el8.aarch64) | pgdg | 97.0 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.5 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.4 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.8 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 99.7 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.6 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 96.0 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 95.8 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.7 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 94.5 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 91.8 KiB | [plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.7 KiB | [plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.6 KiB | [plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 90.1 KiB | [plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 89.0 KiB | [plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.2.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 88.7 KiB | [plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 133.5 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.11.0-1PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pgdg | 126.9 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.10.16-2PGDG.rhel9.9.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el9.x86_64](/os/el9.x86_64) | pigsty | 351.5 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/plpgsql_check_15-2.10.16-1PGSTY.el9.x86_64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 128.2 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.11.0-1PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pgdg | 121.5 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.10.16-2PGDG.rhel9.9.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el9.aarch64](/os/el9.aarch64) | pigsty | 343.5 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/plpgsql_check_15-2.10.16-1PGSTY.el9.aarch64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 136.6 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.11.0-1PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.3 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plpgsql_check_15-2.10.16-2PGDG.rhel10.3.x86_64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.x86_64](/os/el10.x86_64) | pigsty | 353.4 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/plpgsql_check_15-2.10.16-1PGSTY.el10.x86_64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.6 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.11.0-1PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_15` | `2.11.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 129.6 KiB | [plpgsql_check_15-2.11.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.11.0-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.0 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.16-2PGDG.rhel10.3.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 123.0 KiB | [plpgsql_check_15-2.10.16-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.16-2PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pigsty | 345.3 KiB | [plpgsql_check_15-2.10.16-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/plpgsql_check_15-2.10.16-1PGSTY.el10.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.16` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.9 KiB | [plpgsql_check_15-2.10.16-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.16-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.15` | [el10.aarch64](/os/el10.aarch64) | pgdg | 122.8 KiB | [plpgsql_check_15-2.10.15-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.15-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.12` | [el10.aarch64](/os/el10.aarch64) | pgdg | 121.3 KiB | [plpgsql_check_15-2.10.12-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.12-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 120.7 KiB | [plpgsql_check_15-2.10.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 118.4 KiB | [plpgsql_check_15-2.10.7-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.7-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.7 KiB | [plpgsql_check_15-2.10.4-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.4-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.4 KiB | [plpgsql_check_15-2.10.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 117.2 KiB | [plpgsql_check_15-2.10.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.10.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 115.7 KiB | [plpgsql_check_15-2.10.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.10.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.9 KiB | [plpgsql_check_15-2.9.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.3-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 111.7 KiB | [plpgsql_check_15-2.9.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.2-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 112.0 KiB | [plpgsql_check_15-2.9.1-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.9.1-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.11` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.9 KiB | [plpgsql_check_15-2.8.11-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.11-1PGDG.rhel10.2.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.10` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.8 KiB | [plpgsql_check_15-2.8.10-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.10-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.4 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.8` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.4 KiB | [plpgsql_check_15-2.8.8-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.8-1PGDG.rhel10.0.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.5-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.5-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.5-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.5-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.4-1PGDGrhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.4-1PGDGrhel10.1.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.2 KiB | [plpgsql_check_15-2.8.4-1PGDGrhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.4-1PGDGrhel10.0.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 108.6 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel10.aarch64.rpm) |
| `plpgsql_check_15` | `2.8.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 107.8 KiB | [plpgsql_check_15-2.8.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plpgsql_check_15-2.8.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 331.8 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg12+2_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pgdg | 332.0 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.x86_64](/os/d12.x86_64) | pigsty | 329.8 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [d12.x86_64](/os/d12.x86_64) | pgdg | 331.9 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 319.2 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg12+2_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pgdg | 319.1 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d12.aarch64](/os/d12.aarch64) | pigsty | 401.5 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [d12.aarch64](/os/d12.aarch64) | pgdg | 319.5 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 332.6 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg13+2_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pgdg | 332.2 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.x86_64](/os/d13.x86_64) | pigsty | 330.3 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [d13.x86_64](/os/d13.x86_64) | pgdg | 332.3 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg13+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 320.5 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg13+2_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pgdg | 320.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [d13.aarch64](/os/d13.aarch64) | pigsty | 319.1 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [d13.aarch64](/os/d13.aarch64) | pgdg | 320.7 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg13+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 415.3 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pgdg | 415.0 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.x86_64](/os/u22.x86_64) | pigsty | 434.2 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u22.x86_64](/os/u22.x86_64) | pgdg | 414.8 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 402.9 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pgdg | 402.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u22.aarch64](/os/u22.aarch64) | pigsty | 425.8 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u22.aarch64](/os/u22.aarch64) | pgdg | 403.0 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 331.9 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pgdg | 331.8 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.x86_64](/os/u24.x86_64) | pigsty | 345.5 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~noble_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u24.x86_64](/os/u24.x86_64) | pgdg | 331.8 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 320.0 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pgdg | 319.5 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u24.aarch64](/os/u24.aarch64) | pigsty | 338.8 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~noble_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u24.aarch64](/os/u24.aarch64) | pgdg | 319.6 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 329.9 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pgdg | 329.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.x86_64](/os/u26.x86_64) | pigsty | 342.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u26.x86_64](/os/u26.x86_64) | pgdg | 329.5 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 316.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+2_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pgdg | 316.9 KiB | [postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.16` | [u26.aarch64](/os/u26.aarch64) | pigsty | 335.7 KiB | [postgresql-15-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.16-1PGSTY~resolute_arm64.deb) |
| `postgresql-15-plpgsql-check` | `2.10.15` | [u26.aarch64](/os/u26.aarch64) | pgdg | 316.8 KiB | [postgresql-15-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.10.15-1.pgdg26.04+1_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/plpgsql_check" title="Repository" icon="github" subtitle="github.com/okbob/plpgsql_check" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plpgsql_check-2.11.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg plpgsql_check;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install plpgsql_check;		# install via package name, for the active PG version

pig install plpgsql_check -v 19;   # install for PG 19
pig install plpgsql_check -v 18;   # install for PG 18
pig install plpgsql_check -v 17;   # install for PG 17
pig install plpgsql_check -v 16;   # install for PG 16
pig install plpgsql_check -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION plpgsql_check CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [Current version source](https://api.pgxn.org/src/plpgsql_check/plpgsql_check-2.11.0/README.md)
- [Current version source](https://api.pgxn.org/src/plpgsql_check/plpgsql_check-2.11.0/plpgsql_check.control)
- [Installation SQL](https://api.pgxn.org/src/plpgsql_check/plpgsql_check-2.11.0/plpgsql_check--2.11.sql)

- [README.md](https://github.com/okbob/plpgsql_check/blob/9a5c0ec4babebd5eb8e0db104c54f7ddb43fcc07/README.md)
- [Control](https://github.com/okbob/plpgsql_check/blob/9a5c0ec4babebd5eb8e0db104c54f7ddb43fcc07/plpgsql_check.control)
- [Release](https://github.com/okbob/plpgsql_check/releases/tag/v2.10.16)
- [Changes](https://github.com/okbob/plpgsql_check/compare/2.10.15...v2.10.16)

`plpgsql_check` is a PL/pgSQL checker, linter, profiler, tracer, and coverage tool. It analyzes PL/pgSQL function bodies with PostgreSQL's own parser and executor infrastructure, so many problems that would otherwise appear only at runtime can be found during development or CI.

Upstream distribution 2.11.0 installs SQL extension version `2.11`. The README covers PostgreSQL 14–20, including development targets; this is not a binary-package availability claim.

```sql
CREATE EXTENSION IF NOT EXISTS plpgsql_check;
```

### Check a Function

```sql
SELECT *
FROM plpgsql_check_function('public.refresh_totals()');

SELECT *
FROM plpgsql_check_function('public.refresh_totals(int, text)', fatal_errors := false);
```

The table-returning variant is easier to filter, store, or use in CI output:

```sql
SELECT functionid, lineno, statement, sqlstate, message, level
FROM plpgsql_check_function_tb('public.refresh_totals()');
```

Output formats include text, JSON, and XML:

```sql
SELECT * FROM plpgsql_check_function('fx()', format := 'text');
SELECT * FROM plpgsql_check_function('fx()', format := 'json');
SELECT * FROM plpgsql_check_function('fx()', format := 'xml');
```

### Trigger Functions

Trigger functions need the relation they operate on:

```sql
SELECT *
FROM plpgsql_check_function('public.audit_trigger()', 'public.accounts');

SELECT *
FROM plpgsql_check_function(
  'public.audit_trigger()',
  'public.accounts',
  newtable := 'new_rows',
  oldtable := 'old_rows'
);
```

### Warning Levels

```sql
SELECT *
FROM plpgsql_check_function(
  'fx()',
  extra_warnings         := true,
  performance_warnings   := true,
  security_warnings      := true,
  compatibility_warnings := true
);
```

- `extra_warnings` covers missing returns, dead code, shadowed variables, and unused arguments.
- `performance_warnings` covers hidden casts, type modifiers, and patterns that can block index usage.
- `security_warnings` includes dynamic SQL and SQL-injection risk checks.
- `compatibility_warnings` reports obsolete or version-sensitive PL/pgSQL patterns.

### Batch Checks

```sql
SELECT n.nspname, p.proname, c.*
FROM pg_catalog.pg_namespace n
JOIN pg_catalog.pg_proc p ON p.pronamespace = n.oid
JOIN pg_catalog.pg_language l ON l.oid = p.prolang
CROSS JOIN LATERAL plpgsql_check_function_tb(p.oid) AS c
WHERE l.lanname = 'plpgsql'
  AND n.nspname NOT IN ('pg_catalog', 'information_schema');
```

Use this pattern in migration pipelines to catch changed dependencies, dropped columns, unsafe casts, and PL/pgSQL mistakes before release.

### Passive Checking

Passive mode checks functions when they start. It is useful in development and preproduction, but it adds overhead.

```sql
LOAD 'plpgsql_check';
SET plpgsql_check.mode = 'fresh_start';
```

Common settings:

```text
plpgsql_check.mode = disabled | by_function | fresh_start | every_start
plpgsql_check.fatal_errors = yes | no
plpgsql_check.show_nonperformance_warnings = false
plpgsql_check.show_performance_warnings = false
```

### Profiler

```sql
SELECT plpgsql_check_profiler(true);

SELECT public.refresh_totals();

SELECT lineno, exec_stmts, total_time, avg_time, source
FROM plpgsql_profiler_function_tb('public.refresh_totals()');

SELECT stmtid, parent_stmtid, lineno, exec_stmts, stmtname
FROM plpgsql_profiler_function_statements_tb('public.refresh_totals()');

SELECT * FROM plpgsql_profiler_functions_all();
SELECT plpgsql_profiler_reset_all();
```

For shared profiler statistics and reliable early initialization, preload `plpgsql` before `plpgsql_check`:

```ini
shared_preload_libraries = 'plpgsql,plpgsql_check'
```

Without shared preload, profiler data is limited to the active session.

### Tracer and Coverage

Tracing emits notices for function and statement entry/exit and can expose variable values. It is disabled by default and must be enabled by a superuser-controlled setting.

```sql
SET plpgsql_check.enable_tracer = on;
SELECT plpgsql_check_tracer(true, 'terse');

SELECT * FROM plpgsql_coverage_statements('public.refresh_totals()');
SELECT * FROM plpgsql_coverage_branches('public.refresh_totals()');
```

### Pragmas

Use pragma calls inside functions to describe dynamic SQL, temporary tables, inferred record types, or local check settings:

```sql
CREATE OR REPLACE FUNCTION fx(anyelement) RETURNS text AS $$
DECLARE
  r record;
BEGIN
  PERFORM plpgsql_check_pragma('type: r (id int, processed bool)');
  RETURN $1::text;
END;
$$ LANGUAGE plpgsql;
```

Version 2.10 adds `plpgsql_make_pragma(regprocedure)`, which plans temporary-table creation without executing it and returns table pragmas that can be supplied to a check. Release 2.10.2 expands this to additional `CREATE TABLE` statement forms:

```sql
SELECT *
FROM plpgsql_check_function(
  'public.refresh_stage()'::regprocedure,
  pragmas => ARRAY(
    SELECT plpgsql_make_pragma('public.refresh_stage()'::regprocedure)
  )
);
```

### Caveats

- `plpgsql_check` requires `plpgsql`.
- Preloading is optional for active checks, but required for shared profiler storage and robust tracer/profiler initialization.
- Tracer output can include function arguments and local variable values; do not enable it broadly on sensitive production workloads.
- The checker cannot perfectly understand every dynamic SQL string. Use pragmas to document expected dynamic objects and reduce false positives.
- Releases 2.10.2 and 2.10.3 fixed possible crashes in composite-parameter analysis and in coverage or profiler reports for functions with polymorphic arguments. Version 2.10.16 retains those earlier fixes and corrects further false alarms.

The distribution META labels the license BSD, while its shipped LICENSE contains the MIT text; the catalog follows the shipped license file.

### Changes in 2.10.16

Dynamic-SQL query IDs in profiler reports are now disabled by default. Enable `plpgsql_check.profiler_show_dynquery_query_id` only when needed. This release also improves format-string and polymorphic-type checking and fixes plan-retention and retry leaks. Branch coverage distinguishes failure while evaluating an IF/ELSIF condition from entering the implicit ELSE branch.

```sql
SET plpgsql_check.profiler_show_dynquery_query_id = on;
```

That historical patch retains SQL version `2.10` and does not itself require an extension migration. Reconnect sessions using the old library; shared-preloaded installations require a restart. Rehearse upgrades before changing a live server.

### 2.10.16 Correction

Distribution 2.10.16 keeps SQL extension version 2.10 and fixes access to composite-variable tuple descriptors in passive mode. This is an upstream library release; it does not assert that a matching binary package is available.


### 2.11.0 Upgrade

Current distribution 2.11.0 uses SQL extension version `2.11`. It adds warnings for labels in pure expressions and security checks for assigned SEARCH_PATH values, and fixes false positives. This distribution supplies installation SQL but no `ALTER EXTENSION UPDATE` migration scripts. When moving from `2.10`, review dependent objects and rehearse extension recreation; do not assume an in-place update exists or blindly use `DROP ... CASCADE`. After replacing the library, reconnect sessions that loaded it; shared-preloaded instances require a restart.
