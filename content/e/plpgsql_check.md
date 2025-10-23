---
title: "plpgsql_check"
linkTitle: "plpgsql_check"
description: "extended check for plpgsql functions"
weight: 3060
categories: ["LANG"]
width: full
---

extended check for plpgsql functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3060** | {{< badge content="plpgsql_check" link="https://github.com/okbob/plpgsql_check" >}} | {{< ext "plpgsql_check" >}} | `2.8.2` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pldbgapi" >}} {{< ext "plprofiler" >}} {{< ext "pg_hint_plan" >}} {{< ext "pgtap" >}} {{< ext "auto_explain" >}} {{< ext "plv8" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/plpgsql_check" >}} | `2.8.2` | {{< bg "18" "plpgsql_check_18*" "green" >}} {{< bg "17" "plpgsql_check_17*" "green" >}} {{< bg "16" "plpgsql_check_16*" "green" >}} {{< bg "15" "plpgsql_check_15*" "green" >}} {{< bg "14" "plpgsql_check_14*" "green" >}} | `plpgsql_check_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/plpgsql_check" >}} | `2.8.2` | {{< bg "18" "postgresql-18-plpgsql-check" "green" >}} {{< bg "17" "postgresql-17-plpgsql-check" "green" >}} {{< bg "16" "postgresql-16-plpgsql-check" "green" >}} {{< bg "15" "postgresql-15-plpgsql-check" "green" >}} {{< bg "14" "postgresql-14-plpgsql-check" "green" >}} | `postgresql-$v-plpgsql-check` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.8.3" "plpgsql_check_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_15 : AVAIL 29" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_14 : AVAIL 39" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.8.3" "plpgsql_check_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_15 : AVAIL 28" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_14 : AVAIL 28" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.8.3" "plpgsql_check_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_15 : AVAIL 29" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_14 : AVAIL 36" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.8.3" "plpgsql_check_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_15 : AVAIL 28" "blue" >}} | {{< bg "PGDG 2.8.3" "plpgsql_check_14 : AVAIL 28" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.8.3" "postgresql-18-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-17-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-16-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-15-plpgsql-check : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.8.3" "postgresql-14-plpgsql-check : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_18` | 2.8.3 | `el8.x86_64` | pgdg | 113.8 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | 2.8.2 | `el8.x86_64` | pgdg | 113.0 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_18` | 2.8.3 | `el8.aarch64` | pgdg | 105.2 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | 2.8.2 | `el8.aarch64` | pgdg | 104.4 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_18` | 2.8.3 | `el9.x86_64` | pgdg | 109.0 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | 2.8.2 | `el9.x86_64` | pgdg | 108.6 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plpgsql_check_18-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_18` | 2.8.3 | `el9.aarch64` | pgdg | 103.7 KiB | [plpgsql_check_18-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_18` | 2.8.2 | `el9.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_18-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plpgsql_check_18-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `d12.x86_64` | pgdg | 292.5 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `d12.aarch64` | pgdg | 281.4 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `u22.x86_64` | pgdg | 301.2 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `u22.aarch64` | pgdg | 290.6 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `u24.x86_64` | pgdg | 291.8 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plpgsql-check` | 2.8.3 | `u24.aarch64` | pgdg | 280.2 KiB | [postgresql-18-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-18-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_17` | 2.8.3 | `el8.x86_64` | pgdg | 113.8 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.2 | `el8.x86_64` | pgdg | 113.3 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.0 | `el8.x86_64` | pgdg | 112.1 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.15 | `el8.x86_64` | pgdg | 112.0 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.14 | `el8.x86_64` | pgdg | 105.4 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.12 | `el8.x86_64` | pgdg | 105.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.11 | `el8.x86_64` | pgdg | 105.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.3 | `el8.aarch64` | pgdg | 104.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.8.2 | `el8.aarch64` | pgdg | 104.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.8.0 | `el8.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.15 | `el8.aarch64` | pgdg | 103.4 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.14 | `el8.aarch64` | pgdg | 97.7 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.12 | `el8.aarch64` | pgdg | 97.4 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.11 | `el8.aarch64` | pgdg | 97.2 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plpgsql_check_17-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_17` | 2.8.3 | `el9.x86_64` | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.2 | `el9.x86_64` | pgdg | 108.9 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.0 | `el9.x86_64` | pgdg | 107.6 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.15 | `el9.x86_64` | pgdg | 107.6 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.14 | `el9.x86_64` | pgdg | 102.9 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.12 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.7.11 | `el9.x86_64` | pgdg | 103.0 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plpgsql_check_17-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_17` | 2.8.3 | `el9.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.8.2 | `el9.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_17-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.8.0 | `el9.aarch64` | pgdg | 102.8 KiB | [plpgsql_check_17-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.15 | `el9.aarch64` | pgdg | 102.8 KiB | [plpgsql_check_17-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.14 | `el9.aarch64` | pgdg | 98.2 KiB | [plpgsql_check_17-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.12 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_17-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_17` | 2.7.11 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_17-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plpgsql_check_17-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `d12.x86_64` | pgdg | 292.4 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `d12.aarch64` | pgdg | 281.2 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `u22.x86_64` | pgdg | 371.4 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `u22.aarch64` | pgdg | 360.1 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `u24.x86_64` | pgdg | 292.0 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plpgsql-check` | 2.8.3 | `u24.aarch64` | pgdg | 280.3 KiB | [postgresql-17-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-17-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_16` | 2.8.3 | `el8.x86_64` | pgdg | 113.7 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.2 | `el8.x86_64` | pgdg | 113.2 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.0 | `el8.x86_64` | pgdg | 112.1 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.8 | `el8.x86_64` | pgdg | 104.3 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.7 | `el8.x86_64` | pgdg | 104.0 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.6 | `el8.x86_64` | pgdg | 103.8 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.5 | `el8.x86_64` | pgdg | 103.7 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.4 | `el8.x86_64` | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.2 | `el8.x86_64` | pgdg | 105.8 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.15 | `el8.x86_64` | pgdg | 111.9 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.14 | `el8.x86_64` | pgdg | 105.4 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.12 | `el8.x86_64` | pgdg | 105.1 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.11 | `el8.x86_64` | pgdg | 105.0 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.1 | `el8.x86_64` | pgdg | 105.6 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.0 | `el8.x86_64` | pgdg | 104.1 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.2 | `el8.x86_64` | pgdg | 102.6 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.1 | `el8.x86_64` | pgdg | 102.0 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.0 | `el8.x86_64` | pgdg | 101.7 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.4 | `el8.x86_64` | pgdg | 100.6 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.1 | `el8.x86_64` | pgdg | 100.3 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.0 | `el8.x86_64` | pgdg | 100.2 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plpgsql_check_16-2.5.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.3 | `el8.aarch64` | pgdg | 105.0 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.8.2 | `el8.aarch64` | pgdg | 104.5 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.8.0 | `el8.aarch64` | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.8 | `el8.aarch64` | pgdg | 96.5 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.7 | `el8.aarch64` | pgdg | 96.2 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.6 | `el8.aarch64` | pgdg | 96.1 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.5 | `el8.aarch64` | pgdg | 95.9 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.4 | `el8.aarch64` | pgdg | 95.7 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.2 | `el8.aarch64` | pgdg | 98.4 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.15 | `el8.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.14 | `el8.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.12 | `el8.aarch64` | pgdg | 97.3 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.11 | `el8.aarch64` | pgdg | 97.2 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.1 | `el8.aarch64` | pgdg | 98.1 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.0 | `el8.aarch64` | pgdg | 96.8 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.2 | `el8.aarch64` | pgdg | 95.3 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.1 | `el8.aarch64` | pgdg | 94.7 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.0 | `el8.aarch64` | pgdg | 94.5 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.4 | `el8.aarch64` | pgdg | 93.4 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.1 | `el8.aarch64` | pgdg | 93.1 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.0 | `el8.aarch64` | pgdg | 93.0 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plpgsql_check_16-2.5.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_16` | 2.8.3 | `el9.x86_64` | pgdg | 109.0 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.2 | `el9.x86_64` | pgdg | 108.9 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.0 | `el9.x86_64` | pgdg | 107.6 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.8 | `el9.x86_64` | pgdg | 102.4 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.7 | `el9.x86_64` | pgdg | 102.2 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.6 | `el9.x86_64` | pgdg | 102.0 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.5 | `el9.x86_64` | pgdg | 102.2 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.4 | `el9.x86_64` | pgdg | 102.0 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.2 | `el9.x86_64` | pgdg | 104.6 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.15 | `el9.x86_64` | pgdg | 107.7 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.14 | `el9.x86_64` | pgdg | 102.8 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.12 | `el9.x86_64` | pgdg | 103.0 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.11 | `el9.x86_64` | pgdg | 102.9 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.1 | `el9.x86_64` | pgdg | 104.6 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.7.0 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.2 | `el9.x86_64` | pgdg | 101.5 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.1 | `el9.x86_64` | pgdg | 100.9 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.6.0 | `el9.x86_64` | pgdg | 100.9 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.4 | `el9.x86_64` | pgdg | 99.8 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.1 | `el9.x86_64` | pgdg | 99.5 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.5.0 | `el9.x86_64` | pgdg | 99.4 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plpgsql_check_16-2.5.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_16` | 2.8.3 | `el9.aarch64` | pgdg | 103.6 KiB | [plpgsql_check_16-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.8.2 | `el9.aarch64` | pgdg | 103.5 KiB | [plpgsql_check_16-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.8.0 | `el9.aarch64` | pgdg | 102.7 KiB | [plpgsql_check_16-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.8 | `el9.aarch64` | pgdg | 97.8 KiB | [plpgsql_check_16-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.7 | `el9.aarch64` | pgdg | 97.5 KiB | [plpgsql_check_16-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.6 | `el9.aarch64` | pgdg | 97.4 KiB | [plpgsql_check_16-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.5 | `el9.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_16-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.4 | `el9.aarch64` | pgdg | 97.4 KiB | [plpgsql_check_16-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.2 | `el9.aarch64` | pgdg | 100.2 KiB | [plpgsql_check_16-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.15 | `el9.aarch64` | pgdg | 102.7 KiB | [plpgsql_check_16-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.14 | `el9.aarch64` | pgdg | 98.2 KiB | [plpgsql_check_16-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.12 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.11 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.1 | `el9.aarch64` | pgdg | 99.8 KiB | [plpgsql_check_16-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.7.0 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_16-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.2 | `el9.aarch64` | pgdg | 96.7 KiB | [plpgsql_check_16-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.1 | `el9.aarch64` | pgdg | 96.0 KiB | [plpgsql_check_16-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.6.0 | `el9.aarch64` | pgdg | 95.9 KiB | [plpgsql_check_16-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.4 | `el9.aarch64` | pgdg | 95.1 KiB | [plpgsql_check_16-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.1 | `el9.aarch64` | pgdg | 94.8 KiB | [plpgsql_check_16-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_16` | 2.5.0 | `el9.aarch64` | pgdg | 94.7 KiB | [plpgsql_check_16-2.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plpgsql_check_16-2.5.0-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `d12.x86_64` | pgdg | 292.5 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `d12.aarch64` | pgdg | 281.1 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `u22.x86_64` | pgdg | 365.6 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `u22.aarch64` | pgdg | 354.6 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `u24.x86_64` | pgdg | 291.6 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plpgsql-check` | 2.8.3 | `u24.aarch64` | pgdg | 280.3 KiB | [postgresql-16-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-16-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_15` | 2.8.3 | `el8.x86_64` | pgdg | 115.7 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.2 | `el8.x86_64` | pgdg | 115.2 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.0 | `el8.x86_64` | pgdg | 114.1 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.8 | `el8.x86_64` | pgdg | 105.1 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.7 | `el8.x86_64` | pgdg | 104.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.6 | `el8.x86_64` | pgdg | 104.6 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.5 | `el8.x86_64` | pgdg | 104.5 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.4 | `el8.x86_64` | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.2 | `el8.x86_64` | pgdg | 107.4 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.15 | `el8.x86_64` | pgdg | 113.9 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.14 | `el8.x86_64` | pgdg | 106.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.12 | `el8.x86_64` | pgdg | 105.9 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.11 | `el8.x86_64` | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.1 | `el8.x86_64` | pgdg | 107.2 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.0 | `el8.x86_64` | pgdg | 105.7 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.2 | `el8.x86_64` | pgdg | 104.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.1 | `el8.x86_64` | pgdg | 103.4 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.0 | `el8.x86_64` | pgdg | 103.2 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.5.4 | `el8.x86_64` | pgdg | 101.9 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.5.1 | `el8.x86_64` | pgdg | 101.7 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.4 | `el8.x86_64` | pgdg | 98.6 KiB | [plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.2 | `el8.x86_64` | pgdg | 97.4 KiB | [plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.1 | `el8.x86_64` | pgdg | 97.3 KiB | [plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.0 | `el8.x86_64` | pgdg | 96.7 KiB | [plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.3.0-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.6 | `el8.x86_64` | pgdg | 95.6 KiB | [plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.6-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.5 | `el8.x86_64` | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.4 | `el8.x86_64` | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.3 | `el8.x86_64` | pgdg | 95.3 KiB | [plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.2 | `el8.x86_64` | pgdg | 94.8 KiB | [plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plpgsql_check_15-2.2.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.3 | `el8.aarch64` | pgdg | 106.6 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.8.2 | `el8.aarch64` | pgdg | 106.1 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.8.0 | `el8.aarch64` | pgdg | 104.8 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.8 | `el8.aarch64` | pgdg | 97.0 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.7 | `el8.aarch64` | pgdg | 96.7 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.6 | `el8.aarch64` | pgdg | 96.5 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.5 | `el8.aarch64` | pgdg | 96.4 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.4 | `el8.aarch64` | pgdg | 96.2 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.2 | `el8.aarch64` | pgdg | 99.8 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.15 | `el8.aarch64` | pgdg | 104.8 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.14 | `el8.aarch64` | pgdg | 98.1 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.12 | `el8.aarch64` | pgdg | 97.8 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.11 | `el8.aarch64` | pgdg | 97.7 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.1 | `el8.aarch64` | pgdg | 99.7 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.0 | `el8.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.2 | `el8.aarch64` | pgdg | 96.6 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.1 | `el8.aarch64` | pgdg | 96.0 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.0 | `el8.aarch64` | pgdg | 95.8 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.5.4 | `el8.aarch64` | pgdg | 94.7 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.5.1 | `el8.aarch64` | pgdg | 94.5 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.4 | `el8.aarch64` | pgdg | 91.8 KiB | [plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.2 | `el8.aarch64` | pgdg | 90.7 KiB | [plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.2-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.1 | `el8.aarch64` | pgdg | 90.6 KiB | [plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.1-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.0 | `el8.aarch64` | pgdg | 90.1 KiB | [plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.3.0-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.6 | `el8.aarch64` | pgdg | 89.0 KiB | [plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.6-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.5 | `el8.aarch64` | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.5-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.4 | `el8.aarch64` | pgdg | 88.8 KiB | [plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.3 | `el8.aarch64` | pgdg | 88.7 KiB | [plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plpgsql_check_15-2.2.3-1.rhel8.aarch64.rpm) |
| `plpgsql_check_15` | 2.8.3 | `el9.x86_64` | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.2 | `el9.x86_64` | pgdg | 112.9 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.0 | `el9.x86_64` | pgdg | 112.1 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.8 | `el9.x86_64` | pgdg | 103.5 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.7 | `el9.x86_64` | pgdg | 103.2 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.6 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.5 | `el9.x86_64` | pgdg | 103.3 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.4 | `el9.x86_64` | pgdg | 103.0 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.2 | `el9.x86_64` | pgdg | 106.3 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.15 | `el9.x86_64` | pgdg | 112.1 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.14 | `el9.x86_64` | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.12 | `el9.x86_64` | pgdg | 104.2 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.11 | `el9.x86_64` | pgdg | 104.1 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.1 | `el9.x86_64` | pgdg | 107.2 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.7.0 | `el9.x86_64` | pgdg | 104.8 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.2 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.1 | `el9.x86_64` | pgdg | 102.2 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.6.0 | `el9.x86_64` | pgdg | 103.2 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.5.4 | `el9.x86_64` | pgdg | 100.9 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.5.1 | `el9.x86_64` | pgdg | 100.7 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.4 | `el9.x86_64` | pgdg | 97.9 KiB | [plpgsql_check_15-2.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.2 | `el9.x86_64` | pgdg | 96.8 KiB | [plpgsql_check_15-2.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.1 | `el9.x86_64` | pgdg | 96.8 KiB | [plpgsql_check_15-2.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.3.0 | `el9.x86_64` | pgdg | 96.4 KiB | [plpgsql_check_15-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.3.0-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.6 | `el9.x86_64` | pgdg | 95.7 KiB | [plpgsql_check_15-2.2.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.6-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.5 | `el9.x86_64` | pgdg | 95.4 KiB | [plpgsql_check_15-2.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.4 | `el9.x86_64` | pgdg | 95.4 KiB | [plpgsql_check_15-2.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.3 | `el9.x86_64` | pgdg | 95.5 KiB | [plpgsql_check_15-2.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.2.2 | `el9.x86_64` | pgdg | 94.9 KiB | [plpgsql_check_15-2.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plpgsql_check_15-2.2.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_15` | 2.8.3 | `el9.aarch64` | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.8.2 | `el9.aarch64` | pgdg | 107.4 KiB | [plpgsql_check_15-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.8.0 | `el9.aarch64` | pgdg | 106.5 KiB | [plpgsql_check_15-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.8 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.7 | `el9.aarch64` | pgdg | 98.1 KiB | [plpgsql_check_15-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.6 | `el9.aarch64` | pgdg | 98.0 KiB | [plpgsql_check_15-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.5 | `el9.aarch64` | pgdg | 98.3 KiB | [plpgsql_check_15-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.4 | `el9.aarch64` | pgdg | 97.9 KiB | [plpgsql_check_15-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.2 | `el9.aarch64` | pgdg | 101.6 KiB | [plpgsql_check_15-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.15 | `el9.aarch64` | pgdg | 106.8 KiB | [plpgsql_check_15-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.14 | `el9.aarch64` | pgdg | 98.8 KiB | [plpgsql_check_15-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.12 | `el9.aarch64` | pgdg | 98.9 KiB | [plpgsql_check_15-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.11 | `el9.aarch64` | pgdg | 98.9 KiB | [plpgsql_check_15-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.1 | `el9.aarch64` | pgdg | 101.3 KiB | [plpgsql_check_15-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.7.0 | `el9.aarch64` | pgdg | 99.9 KiB | [plpgsql_check_15-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.2 | `el9.aarch64` | pgdg | 98.1 KiB | [plpgsql_check_15-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.1 | `el9.aarch64` | pgdg | 97.5 KiB | [plpgsql_check_15-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.6.0 | `el9.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_15-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.5.4 | `el9.aarch64` | pgdg | 96.5 KiB | [plpgsql_check_15-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.5.1 | `el9.aarch64` | pgdg | 96.2 KiB | [plpgsql_check_15-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.4 | `el9.aarch64` | pgdg | 93.6 KiB | [plpgsql_check_15-2.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.2 | `el9.aarch64` | pgdg | 92.5 KiB | [plpgsql_check_15-2.3.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.2-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.1 | `el9.aarch64` | pgdg | 92.4 KiB | [plpgsql_check_15-2.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.1-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.3.0 | `el9.aarch64` | pgdg | 92.4 KiB | [plpgsql_check_15-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.3.0-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.6 | `el9.aarch64` | pgdg | 91.3 KiB | [plpgsql_check_15-2.2.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.6-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.5 | `el9.aarch64` | pgdg | 91.2 KiB | [plpgsql_check_15-2.2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.5-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.4 | `el9.aarch64` | pgdg | 91.1 KiB | [plpgsql_check_15-2.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_15` | 2.2.3 | `el9.aarch64` | pgdg | 91.0 KiB | [plpgsql_check_15-2.2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plpgsql_check_15-2.2.3-1.rhel9.aarch64.rpm) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `d12.x86_64` | pgdg | 295.7 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `d12.aarch64` | pgdg | 284.1 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `u22.x86_64` | pgdg | 369.0 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `u22.aarch64` | pgdg | 358.0 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `u24.x86_64` | pgdg | 295.7 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plpgsql-check` | 2.8.3 | `u24.aarch64` | pgdg | 284.4 KiB | [postgresql-15-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-15-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plpgsql_check_14` | 2.8.3 | `el8.x86_64` | pgdg | 115.7 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.3-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.2 | `el8.x86_64` | pgdg | 115.3 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.0 | `el8.x86_64` | pgdg | 113.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.8.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.8 | `el8.x86_64` | pgdg | 104.9 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.8-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.7 | `el8.x86_64` | pgdg | 104.6 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.7-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.6 | `el8.x86_64` | pgdg | 104.4 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.6-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.5 | `el8.x86_64` | pgdg | 104.2 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.5-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.4 | `el8.x86_64` | pgdg | 104.1 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.2 | `el8.x86_64` | pgdg | 107.2 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.15 | `el8.x86_64` | pgdg | 113.9 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.15-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.14 | `el8.x86_64` | pgdg | 105.9 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.14-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.12 | `el8.x86_64` | pgdg | 105.6 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.12-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.11 | `el8.x86_64` | pgdg | 105.5 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.11-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.1 | `el8.x86_64` | pgdg | 107.0 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.0 | `el8.x86_64` | pgdg | 105.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.7.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.2 | `el8.x86_64` | pgdg | 103.8 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.2-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.1 | `el8.x86_64` | pgdg | 103.2 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.0 | `el8.x86_64` | pgdg | 103.0 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.6.0-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.5.4 | `el8.x86_64` | pgdg | 101.7 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.5.4-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.5.1 | `el8.x86_64` | pgdg | 101.6 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.5.1-1PGDG.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.4 | `el8.x86_64` | pgdg | 98.4 KiB | [plpgsql_check_14-2.3.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.2 | `el8.x86_64` | pgdg | 97.3 KiB | [plpgsql_check_14-2.3.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.1 | `el8.x86_64` | pgdg | 97.2 KiB | [plpgsql_check_14-2.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.0 | `el8.x86_64` | pgdg | 96.7 KiB | [plpgsql_check_14-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.3.0-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.6 | `el8.x86_64` | pgdg | 95.6 KiB | [plpgsql_check_14-2.2.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.6-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.5 | `el8.x86_64` | pgdg | 95.4 KiB | [plpgsql_check_14-2.2.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.4 | `el8.x86_64` | pgdg | 95.3 KiB | [plpgsql_check_14-2.2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.4-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.3 | `el8.x86_64` | pgdg | 95.2 KiB | [plpgsql_check_14-2.2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.2 | `el8.x86_64` | pgdg | 94.7 KiB | [plpgsql_check_14-2.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.1 | `el8.x86_64` | pgdg | 94.2 KiB | [plpgsql_check_14-2.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.2.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.8 | `el8.x86_64` | pgdg | 89.5 KiB | [plpgsql_check_14-2.1.8-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.8-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.7 | `el8.x86_64` | pgdg | 89.4 KiB | [plpgsql_check_14-2.1.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.7-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.5 | `el8.x86_64` | pgdg | 89.3 KiB | [plpgsql_check_14-2.1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.3 | `el8.x86_64` | pgdg | 88.7 KiB | [plpgsql_check_14-2.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.2 | `el8.x86_64` | pgdg | 88.4 KiB | [plpgsql_check_14-2.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.2-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.10 | `el8.x86_64` | pgdg | 90.4 KiB | [plpgsql_check_14-2.1.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.1.10-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.0.5 | `el8.x86_64` | pgdg | 87.6 KiB | [plpgsql_check_14-2.0.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.0.5-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.0.3 | `el8.x86_64` | pgdg | 87.0 KiB | [plpgsql_check_14-2.0.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-2.0.3-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 1.17.1 | `el8.x86_64` | pgdg | 83.0 KiB | [plpgsql_check_14-1.17.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plpgsql_check_14-1.17.1-1.rhel8.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.3 | `el8.aarch64` | pgdg | 106.5 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.3-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.8.2 | `el8.aarch64` | pgdg | 106.1 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.8.0 | `el8.aarch64` | pgdg | 104.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.8.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.8 | `el8.aarch64` | pgdg | 96.8 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.8-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.7 | `el8.aarch64` | pgdg | 96.5 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.7-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.6 | `el8.aarch64` | pgdg | 96.4 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.6-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.5 | `el8.aarch64` | pgdg | 96.2 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.5-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.4 | `el8.aarch64` | pgdg | 96.1 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.2 | `el8.aarch64` | pgdg | 99.6 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.15 | `el8.aarch64` | pgdg | 104.8 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.15-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.14 | `el8.aarch64` | pgdg | 97.9 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.14-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.12 | `el8.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.12-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.11 | `el8.aarch64` | pgdg | 97.5 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.11-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.1 | `el8.aarch64` | pgdg | 99.4 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.0 | `el8.aarch64` | pgdg | 98.0 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.7.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.2 | `el8.aarch64` | pgdg | 96.4 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.2-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.1 | `el8.aarch64` | pgdg | 95.7 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.0 | `el8.aarch64` | pgdg | 95.6 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.6.0-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.5.4 | `el8.aarch64` | pgdg | 94.5 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.5.4-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.5.1 | `el8.aarch64` | pgdg | 94.3 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.5.1-1PGDG.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.4 | `el8.aarch64` | pgdg | 91.6 KiB | [plpgsql_check_14-2.3.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.2 | `el8.aarch64` | pgdg | 90.6 KiB | [plpgsql_check_14-2.3.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.2-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.1 | `el8.aarch64` | pgdg | 90.5 KiB | [plpgsql_check_14-2.3.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.1-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.0 | `el8.aarch64` | pgdg | 90.0 KiB | [plpgsql_check_14-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.3.0-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.6 | `el8.aarch64` | pgdg | 88.9 KiB | [plpgsql_check_14-2.2.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.6-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.5 | `el8.aarch64` | pgdg | 88.8 KiB | [plpgsql_check_14-2.2.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.5-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.4 | `el8.aarch64` | pgdg | 88.7 KiB | [plpgsql_check_14-2.2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.4-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.3 | `el8.aarch64` | pgdg | 88.6 KiB | [plpgsql_check_14-2.2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plpgsql_check_14-2.2.3-1.rhel8.aarch64.rpm) |
| `plpgsql_check_14` | 2.8.3 | `el9.x86_64` | pgdg | 112.9 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.3-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.2 | `el9.x86_64` | pgdg | 112.8 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.0 | `el9.x86_64` | pgdg | 111.9 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.8.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.8 | `el9.x86_64` | pgdg | 103.3 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.8-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.7 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.7-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.6 | `el9.x86_64` | pgdg | 102.9 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.6-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.5 | `el9.x86_64` | pgdg | 103.1 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.5-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.4 | `el9.x86_64` | pgdg | 102.9 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.2 | `el9.x86_64` | pgdg | 106.1 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.15 | `el9.x86_64` | pgdg | 112.1 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.15-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.14 | `el9.x86_64` | pgdg | 103.8 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.14-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.12 | `el9.x86_64` | pgdg | 103.9 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.12-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.11 | `el9.x86_64` | pgdg | 103.8 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.11-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.1 | `el9.x86_64` | pgdg | 105.9 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.7.0 | `el9.x86_64` | pgdg | 104.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.7.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.2 | `el9.x86_64` | pgdg | 102.7 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.2-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.1 | `el9.x86_64` | pgdg | 101.8 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.6.0 | `el9.x86_64` | pgdg | 101.8 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.6.0-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.5.4 | `el9.x86_64` | pgdg | 100.8 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.5.4-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.5.1 | `el9.x86_64` | pgdg | 100.6 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.5.1-1PGDG.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.4 | `el9.x86_64` | pgdg | 97.8 KiB | [plpgsql_check_14-2.3.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.2 | `el9.x86_64` | pgdg | 96.8 KiB | [plpgsql_check_14-2.3.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.1 | `el9.x86_64` | pgdg | 96.7 KiB | [plpgsql_check_14-2.3.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.3.0 | `el9.x86_64` | pgdg | 96.4 KiB | [plpgsql_check_14-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.3.0-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.6 | `el9.x86_64` | pgdg | 95.5 KiB | [plpgsql_check_14-2.2.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.6-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.5 | `el9.x86_64` | pgdg | 95.4 KiB | [plpgsql_check_14-2.2.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.4 | `el9.x86_64` | pgdg | 95.3 KiB | [plpgsql_check_14-2.2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.4-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.3 | `el9.x86_64` | pgdg | 95.2 KiB | [plpgsql_check_14-2.2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.2 | `el9.x86_64` | pgdg | 94.7 KiB | [plpgsql_check_14-2.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.2.1 | `el9.x86_64` | pgdg | 94.3 KiB | [plpgsql_check_14-2.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.2.1-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.8 | `el9.x86_64` | pgdg | 89.8 KiB | [plpgsql_check_14-2.1.8-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.8-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.7 | `el9.x86_64` | pgdg | 89.6 KiB | [plpgsql_check_14-2.1.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.7-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.5 | `el9.x86_64` | pgdg | 89.6 KiB | [plpgsql_check_14-2.1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.5-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.3 | `el9.x86_64` | pgdg | 88.8 KiB | [plpgsql_check_14-2.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.3-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.2 | `el9.x86_64` | pgdg | 88.7 KiB | [plpgsql_check_14-2.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.2-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.1.10 | `el9.x86_64` | pgdg | 90.4 KiB | [plpgsql_check_14-2.1.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plpgsql_check_14-2.1.10-1.rhel9.x86_64.rpm) |
| `plpgsql_check_14` | 2.8.3 | `el9.aarch64` | pgdg | 107.5 KiB | [plpgsql_check_14-2.8.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.3-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.8.2 | `el9.aarch64` | pgdg | 107.4 KiB | [plpgsql_check_14-2.8.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.8.0 | `el9.aarch64` | pgdg | 106.5 KiB | [plpgsql_check_14-2.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.8.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.8 | `el9.aarch64` | pgdg | 98.2 KiB | [plpgsql_check_14-2.7.8-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.8-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.7 | `el9.aarch64` | pgdg | 97.8 KiB | [plpgsql_check_14-2.7.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.7-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.6 | `el9.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.6-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.5 | `el9.aarch64` | pgdg | 97.8 KiB | [plpgsql_check_14-2.7.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.5-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.4 | `el9.aarch64` | pgdg | 97.6 KiB | [plpgsql_check_14-2.7.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.2 | `el9.aarch64` | pgdg | 101.1 KiB | [plpgsql_check_14-2.7.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.15 | `el9.aarch64` | pgdg | 106.7 KiB | [plpgsql_check_14-2.7.15-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.15-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.14 | `el9.aarch64` | pgdg | 98.5 KiB | [plpgsql_check_14-2.7.14-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.14-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.12 | `el9.aarch64` | pgdg | 98.7 KiB | [plpgsql_check_14-2.7.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.12-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.11 | `el9.aarch64` | pgdg | 98.6 KiB | [plpgsql_check_14-2.7.11-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.11-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.1 | `el9.aarch64` | pgdg | 100.9 KiB | [plpgsql_check_14-2.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.7.0 | `el9.aarch64` | pgdg | 99.6 KiB | [plpgsql_check_14-2.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.7.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.2 | `el9.aarch64` | pgdg | 98.0 KiB | [plpgsql_check_14-2.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.2-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.1 | `el9.aarch64` | pgdg | 97.2 KiB | [plpgsql_check_14-2.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.6.0 | `el9.aarch64` | pgdg | 97.2 KiB | [plpgsql_check_14-2.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.6.0-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.5.4 | `el9.aarch64` | pgdg | 96.2 KiB | [plpgsql_check_14-2.5.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.5.4-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.5.1 | `el9.aarch64` | pgdg | 96.0 KiB | [plpgsql_check_14-2.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.5.1-1PGDG.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.4 | `el9.aarch64` | pgdg | 93.3 KiB | [plpgsql_check_14-2.3.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.2 | `el9.aarch64` | pgdg | 92.5 KiB | [plpgsql_check_14-2.3.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.2-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.1 | `el9.aarch64` | pgdg | 92.4 KiB | [plpgsql_check_14-2.3.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.1-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.3.0 | `el9.aarch64` | pgdg | 92.3 KiB | [plpgsql_check_14-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.3.0-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.6 | `el9.aarch64` | pgdg | 91.3 KiB | [plpgsql_check_14-2.2.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.6-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.5 | `el9.aarch64` | pgdg | 91.1 KiB | [plpgsql_check_14-2.2.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.5-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.4 | `el9.aarch64` | pgdg | 91.0 KiB | [plpgsql_check_14-2.2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.4-1.rhel9.aarch64.rpm) |
| `plpgsql_check_14` | 2.2.3 | `el9.aarch64` | pgdg | 90.9 KiB | [plpgsql_check_14-2.2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plpgsql_check_14-2.2.3-1.rhel9.aarch64.rpm) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `d12.x86_64` | pgdg | 295.6 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `d12.aarch64` | pgdg | 284.2 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `u22.x86_64` | pgdg | 353.1 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `u22.aarch64` | pgdg | 342.1 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `u24.x86_64` | pgdg | 295.3 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plpgsql-check` | 2.8.3 | `u24.aarch64` | pgdg | 284.1 KiB | [postgresql-14-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plpgsql-check/postgresql-14-plpgsql-check_2.8.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/okbob/plpgsql_check" title="Repository" icon="github" subtitle="github.com/okbob/plpgsql_check" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="plpgsql_check-2.8.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get plpgsql_check; # get plpgsql_check source code
pig build dep plpgsql_check; # install build dependencies
pig build pkg plpgsql_check; # build extension rpm or deb
pig build ext plpgsql_check; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install plpgsql_check; # install by extension name, for the current active PG version
pig ext install plpgsql_check; # install via package alias, for the active PG version
pig ext install plpgsql_check -v 18;   # install for PG 18
pig ext install plpgsql_check -v 17;   # install for PG 17
pig ext install plpgsql_check -v 16;   # install for PG 16
pig ext install plpgsql_check -v 15;   # install for PG 15
pig ext install plpgsql_check -v 14;   # install for PG 14
pig ext install plpgsql_check -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION plpgsql_check;
```

