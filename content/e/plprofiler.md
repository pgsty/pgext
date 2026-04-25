---
title: "plprofiler"
linkTitle: "plprofiler"
description: "server-side support for profiling PL/pgSQL functions"
weight: 3070
categories: ["LANG"]
width: full
---

[**plprofiler**](https://github.com/bigsql/plprofiler) : server-side support for profiling PL/pgSQL functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3070** | {{< badge content="plprofiler" link="https://github.com/bigsql/plprofiler" >}} | {{< ext "plprofiler" >}} | `4.2.5` | {{< category "LANG" >}} | {{< license "Artistic" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pldbgapi" >}} {{< ext "plpgsql_check" >}} {{< ext "plpgsql" >}} {{< ext "pgtap" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `plprofiler` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "plprofiler_18" "green" >}} {{< bg "17" "plprofiler_17" "green" >}} {{< bg "16" "plprofiler_16" "green" >}} {{< bg "15" "plprofiler_15" "green" >}} {{< bg "14" "plprofiler_14" "green" >}} | `plprofiler_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "postgresql-18-plprofiler" "green" >}} {{< bg "17" "postgresql-17-plprofiler" "green" >}} {{< bg "16" "postgresql-16-plprofiler" "green" >}} {{< bg "15" "postgresql-15-plprofiler" "green" >}} {{< bg "14" "postgresql-14-plprofiler" "green" >}} | `postgresql-$v-plprofiler` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_18` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.5 KiB | [plprofiler_18-4.2.5-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/plprofiler_18-4.2.5-5PGDG.rhel8.10.x86_64.rpm) |
| `plprofiler_18` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.5 KiB | [plprofiler_18-4.2.5-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/plprofiler_18-4.2.5-5PGDG.rhel8.10.aarch64.rpm) |
| `plprofiler_18` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_18-4.2.5-5PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plprofiler_18-4.2.5-5PGDG.rhel9.7.x86_64.rpm) |
| `plprofiler_18` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.7 KiB | [plprofiler_18-4.2.5-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/plprofiler_18-4.2.5-2PGDG.rhel9.x86_64.rpm) |
| `plprofiler_18` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.9 KiB | [plprofiler_18-4.2.5-5PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plprofiler_18-4.2.5-5PGDG.rhel9.7.aarch64.rpm) |
| `plprofiler_18` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.7 KiB | [plprofiler_18-4.2.5-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/plprofiler_18-4.2.5-2PGDG.rhel9.aarch64.rpm) |
| `plprofiler_18` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.1 KiB | [plprofiler_18-4.2.5-5PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plprofiler_18-4.2.5-5PGDG.rhel10.1.x86_64.rpm) |
| `plprofiler_18` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_18-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/plprofiler_18-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `plprofiler_18` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.0 KiB | [plprofiler_18-4.2.5-5PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plprofiler_18-4.2.5-5PGDG.rhel10.1.aarch64.rpm) |
| `plprofiler_18` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.2 KiB | [plprofiler_18-4.2.5-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/plprofiler_18-4.2.5-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.5 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.3 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.5 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 46.2 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 45.4 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.3 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.4 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.9 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-18-plprofiler` | `4.2.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 43.9 KiB | [postgresql-18-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-18-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_17` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.5 KiB | [plprofiler_17-4.2.5-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plprofiler_17-4.2.5-5PGDG.rhel8.10.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.0 KiB | [plprofiler_17-4.2.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/plprofiler_17-4.2.5-1PGDG.rhel8.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.5 KiB | [plprofiler_17-4.2.5-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plprofiler_17-4.2.5-5PGDG.rhel8.10.aarch64.rpm) |
| `plprofiler_17` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.0 KiB | [plprofiler_17-4.2.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/plprofiler_17-4.2.5-1PGDG.rhel8.aarch64.rpm) |
| `plprofiler_17` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_17-4.2.5-5PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plprofiler_17-4.2.5-5PGDG.rhel9.7.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 7.1 KiB | [plprofiler_17-4.2.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/plprofiler_17-4.2.5-1PGDG.rhel9.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.9 KiB | [plprofiler_17-4.2.5-5PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plprofiler_17-4.2.5-5PGDG.rhel9.7.aarch64.rpm) |
| `plprofiler_17` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 7.0 KiB | [plprofiler_17-4.2.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/plprofiler_17-4.2.5-1PGDG.rhel9.aarch64.rpm) |
| `plprofiler_17` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.1 KiB | [plprofiler_17-4.2.5-5PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plprofiler_17-4.2.5-5PGDG.rhel10.1.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_17-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/plprofiler_17-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `plprofiler_17` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.0 KiB | [plprofiler_17-4.2.5-5PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/plprofiler_17-4.2.5-5PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-17-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.4 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.5 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.1 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.2 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.4 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.6 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.8 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-17-plprofiler` | `4.2.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 43.9 KiB | [postgresql-17-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-17-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_16` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.5 KiB | [plprofiler_16-4.2.5-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plprofiler_16-4.2.5-5PGDG.rhel8.10.x86_64.rpm) |
| `plprofiler_16` | `4.2.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.8 KiB | [plprofiler_16-4.2.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/plprofiler_16-4.2.4-1PGDG.rhel8.x86_64.rpm) |
| `plprofiler_16` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.5 KiB | [plprofiler_16-4.2.5-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plprofiler_16-4.2.5-5PGDG.rhel8.10.aarch64.rpm) |
| `plprofiler_16` | `4.2.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.8 KiB | [plprofiler_16-4.2.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/plprofiler_16-4.2.4-1PGDG.rhel8.aarch64.rpm) |
| `plprofiler_16` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_16-4.2.5-5PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plprofiler_16-4.2.5-5PGDG.rhel9.7.x86_64.rpm) |
| `plprofiler_16` | `4.2.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_16-4.2.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/plprofiler_16-4.2.4-1PGDG.rhel9.x86_64.rpm) |
| `plprofiler_16` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.9 KiB | [plprofiler_16-4.2.5-5PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plprofiler_16-4.2.5-5PGDG.rhel9.7.aarch64.rpm) |
| `plprofiler_16` | `4.2.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.7 KiB | [plprofiler_16-4.2.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/plprofiler_16-4.2.4-1PGDG.rhel9.aarch64.rpm) |
| `plprofiler_16` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.1 KiB | [plprofiler_16-4.2.5-5PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plprofiler_16-4.2.5-5PGDG.rhel10.1.x86_64.rpm) |
| `plprofiler_16` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_16-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/plprofiler_16-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `plprofiler_16` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.0 KiB | [plprofiler_16-4.2.5-5PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/plprofiler_16-4.2.5-5PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-16-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.3 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.4 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.2 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.5 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.1 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.1 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.4 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.5 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.8 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-16-plprofiler` | `4.2.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 43.9 KiB | [postgresql-16-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-16-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_15` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.5 KiB | [plprofiler_15-4.2.5-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plprofiler_15-4.2.5-5PGDG.rhel8.10.x86_64.rpm) |
| `plprofiler_15` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.8 KiB | [plprofiler_15-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plprofiler_15-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `plprofiler_15` | `4.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.6 KiB | [plprofiler_15-4.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plprofiler_15-4.2.1-1.rhel8.x86_64.rpm) |
| `plprofiler_15` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.4 KiB | [plprofiler_15-4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/plprofiler_15-4.2-1.rhel8.x86_64.rpm) |
| `plprofiler_15` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.5 KiB | [plprofiler_15-4.2.5-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plprofiler_15-4.2.5-5PGDG.rhel8.10.aarch64.rpm) |
| `plprofiler_15` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.7 KiB | [plprofiler_15-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plprofiler_15-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `plprofiler_15` | `4.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.5 KiB | [plprofiler_15-4.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plprofiler_15-4.2.1-1.rhel8.aarch64.rpm) |
| `plprofiler_15` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.4 KiB | [plprofiler_15-4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/plprofiler_15-4.2-1.rhel8.aarch64.rpm) |
| `plprofiler_15` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_15-4.2.5-5PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plprofiler_15-4.2.5-5PGDG.rhel9.7.x86_64.rpm) |
| `plprofiler_15` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 KiB | [plprofiler_15-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plprofiler_15-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `plprofiler_15` | `4.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.6 KiB | [plprofiler_15-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plprofiler_15-4.2.1-1.rhel9.x86_64.rpm) |
| `plprofiler_15` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.5 KiB | [plprofiler_15-4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/plprofiler_15-4.2-1.rhel9.x86_64.rpm) |
| `plprofiler_15` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.9 KiB | [plprofiler_15-4.2.5-5PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plprofiler_15-4.2.5-5PGDG.rhel9.7.aarch64.rpm) |
| `plprofiler_15` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.7 KiB | [plprofiler_15-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plprofiler_15-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `plprofiler_15` | `4.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.5 KiB | [plprofiler_15-4.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plprofiler_15-4.2.1-1.rhel9.aarch64.rpm) |
| `plprofiler_15` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 KiB | [plprofiler_15-4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/plprofiler_15-4.2-1.rhel9.aarch64.rpm) |
| `plprofiler_15` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.1 KiB | [plprofiler_15-4.2.5-5PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plprofiler_15-4.2.5-5PGDG.rhel10.1.x86_64.rpm) |
| `plprofiler_15` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_15-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/plprofiler_15-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `plprofiler_15` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.0 KiB | [plprofiler_15-4.2.5-5PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/plprofiler_15-4.2.5-5PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-15-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.5 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.5 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.5 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.7 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 52.3 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 51.3 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.6 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.7 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 45.0 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-15-plprofiler` | `4.2.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 44.1 KiB | [postgresql-15-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-15-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_14` | `4.2.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 7.5 KiB | [plprofiler_14-4.2.5-5PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plprofiler_14-4.2.5-5PGDG.rhel8.10.x86_64.rpm) |
| `plprofiler_14` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.8 KiB | [plprofiler_14-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plprofiler_14-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `plprofiler_14` | `4.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.6 KiB | [plprofiler_14-4.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plprofiler_14-4.2.1-1.rhel8.x86_64.rpm) |
| `plprofiler_14` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.4 KiB | [plprofiler_14-4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/plprofiler_14-4.2-1.rhel8.x86_64.rpm) |
| `plprofiler_14` | `4.2.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 7.5 KiB | [plprofiler_14-4.2.5-5PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plprofiler_14-4.2.5-5PGDG.rhel8.10.aarch64.rpm) |
| `plprofiler_14` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.7 KiB | [plprofiler_14-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plprofiler_14-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `plprofiler_14` | `4.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.5 KiB | [plprofiler_14-4.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plprofiler_14-4.2.1-1.rhel8.aarch64.rpm) |
| `plprofiler_14` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.4 KiB | [plprofiler_14-4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/plprofiler_14-4.2-1.rhel8.aarch64.rpm) |
| `plprofiler_14` | `4.2.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.9 KiB | [plprofiler_14-4.2.5-5PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plprofiler_14-4.2.5-5PGDG.rhel9.7.x86_64.rpm) |
| `plprofiler_14` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 KiB | [plprofiler_14-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plprofiler_14-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `plprofiler_14` | `4.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.6 KiB | [plprofiler_14-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plprofiler_14-4.2.1-1.rhel9.x86_64.rpm) |
| `plprofiler_14` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.5 KiB | [plprofiler_14-4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/plprofiler_14-4.2-1.rhel9.x86_64.rpm) |
| `plprofiler_14` | `4.2.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.9 KiB | [plprofiler_14-4.2.5-5PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plprofiler_14-4.2.5-5PGDG.rhel9.7.aarch64.rpm) |
| `plprofiler_14` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.7 KiB | [plprofiler_14-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plprofiler_14-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `plprofiler_14` | `4.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.5 KiB | [plprofiler_14-4.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plprofiler_14-4.2.1-1.rhel9.aarch64.rpm) |
| `plprofiler_14` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 KiB | [plprofiler_14-4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/plprofiler_14-4.2-1.rhel9.aarch64.rpm) |
| `plprofiler_14` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.1 KiB | [plprofiler_14-4.2.5-5PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plprofiler_14-4.2.5-5PGDG.rhel10.1.x86_64.rpm) |
| `plprofiler_14` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_14-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/plprofiler_14-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `plprofiler_14` | `4.2.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 7.0 KiB | [plprofiler_14-4.2.5-5PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/plprofiler_14-4.2.5-5PGDG.rhel10.1.aarch64.rpm) |
| `postgresql-14-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.2 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.2 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.3 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.3 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 50.4 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.4 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.4 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.3 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u26.x86_64](/os/u26.x86_64) | pgdg | 44.8 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg26.04+1_amd64.deb) |
| `postgresql-14-plprofiler` | `4.2.5` | [u26.aarch64](/os/u26.aarch64) | pgdg | 43.8 KiB | [postgresql-14-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-14-plprofiler_4.2.5-4.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsql/plprofiler" title="Repository" icon="github" subtitle="github.com/bigsql/plprofiler" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install plprofiler;		# install via package name, for the active PG version

pig install plprofiler -v 18;   # install for PG 18
pig install plprofiler -v 17;   # install for PG 17
pig install plprofiler -v 16;   # install for PG 16
pig install plprofiler -v 15;   # install for PG 15
pig install plprofiler -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'plprofiler';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plprofiler;
```




## Usage

> [plprofiler: server-side support for profiling PL/pgSQL functions](https://github.com/bigsql/plprofiler)

`plprofiler` is a profiling tool for PL/pgSQL functions that identifies performance bottlenecks and generates interactive flame graph reports.

```sql
CREATE EXTENSION plprofiler;
```

### Configuration

Add to `postgresql.conf`:

```
shared_preload_libraries = 'plprofiler'
```

### Command-Line Tool

The `plprofiler` command-line utility controls profiling and generates reports:

```bash
# Profile a specific SQL command
plprofiler run -d mydb --command "SELECT my_function()" --output report.html

# Monitor profiling in real-time
plprofiler monitor -d mydb

# Save profiling data for later analysis
plprofiler save -d mydb --name "my_profile"

# Generate HTML report with flame graphs
plprofiler report -d mydb --from-data "my_profile" --output report.html

# List saved profiling datasets
plprofiler list -d mydb

# Reset profiling data
plprofiler reset-data -d mydb

# Export/import profiling data
plprofiler export -d mydb --from-data "my_profile" > profile.json
plprofiler import -d mydb --into-data "imported" < profile.json
```

### SQL Interface

```sql
-- Enable profiling for the current session
SELECT pl_profiler_set_enabled_local(true);

-- Execute functions to be profiled
SELECT my_function();

-- Collect profiling data into shared hash tables
SELECT pl_profiler_collect_data();

-- Disable profiling
SELECT pl_profiler_set_enabled_local(false);

-- Enable profiling globally (for all sessions)
SELECT pl_profiler_set_enabled_global(true);

-- Reset local/shared profiling data
SELECT pl_profiler_reset_local();
SELECT pl_profiler_reset_shared();
```

### Report Output

Generated HTML reports include:

- **Interactive flame graphs** showing wall-clock time spent in PL/pgSQL code
- **Per-function statistics** with self-time (total minus children)
- **Top functions** ranked by time consumption (default: top 10)
- Self-contained HTML requiring no external dependencies

### Profiling Methods

- **Direct profiling**: Run specific SQL while collecting data
- **Timed collection**: Interval-based statistics gathering
- **Per-user profiling**: Enable profiling for specific database users via `ALTER USER`
- **Production monitoring**: Low-overhead profiling on live systems
