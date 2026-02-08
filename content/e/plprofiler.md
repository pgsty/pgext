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
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pldbgapi" >}} {{< ext "plpgsql_check" >}} {{< ext "plpgsql" >}} {{< ext "pgtap" >}} {{< ext "pg_profile" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_store_plans" >}} {{< ext "auto_explain" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `plprofiler` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "plprofiler_18" "green" >}} {{< bg "17" "plprofiler_17" "green" >}} {{< bg "16" "plprofiler_16" "green" >}} {{< bg "15" "plprofiler_15" "green" >}} {{< bg "14" "plprofiler_14" "green" >}} {{< bg "13" "plprofiler_13" "green" >}} | `plprofiler_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.2.5` | {{< bg "18" "postgresql-18-plprofiler" "green" >}} {{< bg "17" "postgresql-17-plprofiler" "green" >}} {{< bg "16" "postgresql-16-plprofiler" "green" >}} {{< bg "15" "postgresql-15-plprofiler" "green" >}} {{< bg "14" "postgresql-14-plprofiler" "green" >}} {{< bg "13" "postgresql-13-plprofiler" "green" >}} | `postgresql-$v-plprofiler` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.2" "plprofiler_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.2" "plprofiler_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.2" "plprofiler_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.2.2" "plprofiler_13 : AVAIL 3" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.2.5" "plprofiler_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "plprofiler_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "plprofiler_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.2.5" "postgresql-18-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-17-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-16-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-15-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-14-plprofiler : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.2.5" "postgresql-13-plprofiler : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
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

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `plprofiler_13` | `4.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.8 KiB | [plprofiler_13-4.2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plprofiler_13-4.2.2-1PGDG.rhel8.x86_64.rpm) |
| `plprofiler_13` | `4.2.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.6 KiB | [plprofiler_13-4.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plprofiler_13-4.2.1-1.rhel8.x86_64.rpm) |
| `plprofiler_13` | `4.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 6.4 KiB | [plprofiler_13-4.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/plprofiler_13-4.2-1.rhel8.x86_64.rpm) |
| `plprofiler_13` | `4.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.7 KiB | [plprofiler_13-4.2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plprofiler_13-4.2.2-1PGDG.rhel8.aarch64.rpm) |
| `plprofiler_13` | `4.2.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.5 KiB | [plprofiler_13-4.2.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plprofiler_13-4.2.1-1.rhel8.aarch64.rpm) |
| `plprofiler_13` | `4.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 6.4 KiB | [plprofiler_13-4.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/plprofiler_13-4.2-1.rhel8.aarch64.rpm) |
| `plprofiler_13` | `4.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.8 KiB | [plprofiler_13-4.2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plprofiler_13-4.2.2-1PGDG.rhel9.x86_64.rpm) |
| `plprofiler_13` | `4.2.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.6 KiB | [plprofiler_13-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plprofiler_13-4.2.1-1.rhel9.x86_64.rpm) |
| `plprofiler_13` | `4.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 6.5 KiB | [plprofiler_13-4.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/plprofiler_13-4.2-1.rhel9.x86_64.rpm) |
| `plprofiler_13` | `4.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.7 KiB | [plprofiler_13-4.2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plprofiler_13-4.2.2-1PGDG.rhel9.aarch64.rpm) |
| `plprofiler_13` | `4.2.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.5 KiB | [plprofiler_13-4.2.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plprofiler_13-4.2.1-1.rhel9.aarch64.rpm) |
| `plprofiler_13` | `4.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 6.3 KiB | [plprofiler_13-4.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/plprofiler_13-4.2-1.rhel9.aarch64.rpm) |
| `plprofiler_13` | `4.2.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 7.3 KiB | [plprofiler_13-4.2.5-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/plprofiler_13-4.2.5-2PGDG.rhel10.x86_64.rpm) |
| `postgresql-13-plprofiler` | `4.2.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 45.1 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg12+1_amd64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 44.3 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg12+1_arm64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 45.0 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg13+1_amd64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 44.4 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg13+1_arm64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 49.9 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 49.0 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 45.2 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-plprofiler` | `4.2.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 44.4 KiB | [postgresql-13-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/plprofiler/postgresql-13-plprofiler_4.2.5-4.pgdg24.04+1_arm64.deb) |

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
pig install plprofiler -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION plprofiler;
```
