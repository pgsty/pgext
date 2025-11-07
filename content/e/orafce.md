---
title: "orafce"
linkTitle: "orafce"
description: "Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS"
weight: 9100
categories: ["SIM"]
width: full
---

Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9100** | {{< badge content="orafce" link="https://github.com/orafce/orafce" >}} | {{< ext "orafce" >}} | `4.14.6` | {{< category "SIM" >}} | {{< license "BSD 0-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "pgtt" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "db_migrator" >}} |

> [!Note] el llvmjit deps break


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/orafce" >}} | `4.14.6` | {{< bg "18" "orafce_18" "green" >}} {{< bg "17" "orafce_17" "green" >}} {{< bg "16" "orafce_16" "green" >}} {{< bg "15" "orafce_15" "green" >}} {{< bg "14" "orafce_14" "green" >}} {{< bg "13" "orafce_13" "green" >}} | `orafce_$v` | - |
| **Debian** | {{< badge content="PGDG" link="/e/orafce" >}} | `4.14.6` | {{< bg "18" "postgresql-18-orafce" "green" >}} {{< bg "17" "postgresql-17-orafce" "green" >}} {{< bg "16" "postgresql-16-orafce" "green" >}} {{< bg "15" "postgresql-15-orafce" "green" >}} {{< bg "14" "postgresql-14-orafce" "green" >}} {{< bg "13" "postgresql-13-orafce" "green" >}} | `postgresql-$v-orafce` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.14.6" "orafce_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.14.6" "orafce_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_16 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_15 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_14 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_13 : AVAIL 19" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.14.6" "orafce_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 4.14.6" "orafce_17 : AVAIL 10" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_16 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_15 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_14 : AVAIL 19" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_13 : AVAIL 19" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.16.1" "orafce_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_16 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_14 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_13 : AVAIL 20" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.16.1" "orafce_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_17 : AVAIL 11" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_16 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_15 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_14 : AVAIL 20" "blue" >}} | {{< bg "PGDG 4.9.4" "orafce_13 : AVAIL 20" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.16.1" "orafce_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_13 : AVAIL 4" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.16.1" "orafce_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_14 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.1" "orafce_13 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.16.1" "postgresql-18-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-17-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-16-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-15-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-14-orafce : AVAIL 1" "blue" >}} | {{< bg "PGDG 4.16.1" "postgresql-13-orafce : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_18` | 4.14.6 | `el8.x86_64` | pgdg | 151.3 KiB | [orafce_18-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_18` | 4.14.5 | `el8.x86_64` | pgdg | 151.3 KiB | [orafce_18-4.14.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.14.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_18` | 4.14.6 | `el8.aarch64` | pgdg | 146.9 KiB | [orafce_18-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_18` | 4.14.5 | `el8.aarch64` | pgdg | 147.0 KiB | [orafce_18-4.14.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.14.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_18` | 4.16.1 | `el9.x86_64` | pgdg | 150.0 KiB | [orafce_18-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | 4.14.6 | `el9.x86_64` | pgdg | 148.9 KiB | [orafce_18-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | 4.14.5 | `el9.x86_64` | pgdg | 148.7 KiB | [orafce_18-4.14.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.14.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | 4.16.1 | `el9.aarch64` | pgdg | 147.7 KiB | [orafce_18-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | 4.14.6 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_18-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | 4.14.5 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_18-4.14.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.14.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | 4.16.1 | `el10.x86_64` | pgdg | 150.9 KiB | [orafce_18-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | 4.14.6 | `el10.x86_64` | pgdg | 150.1 KiB | [orafce_18-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | 4.14.5 | `el10.x86_64` | pgdg | 149.9 KiB | [orafce_18-4.14.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.14.5-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | 4.16.1 | `el10.aarch64` | pgdg | 149.2 KiB | [orafce_18-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_18` | 4.14.6 | `el10.aarch64` | pgdg | 148.3 KiB | [orafce_18-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_18` | 4.14.5 | `el10.aarch64` | pgdg | 148.3 KiB | [orafce_18-4.14.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.14.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 363.0 KiB | [postgresql-18-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 355.5 KiB | [postgresql-18-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 363.6 KiB | [postgresql-18-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 356.8 KiB | [postgresql-18-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 368.3 KiB | [postgresql-18-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 360.0 KiB | [postgresql-18-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 360.1 KiB | [postgresql-18-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 355.0 KiB | [postgresql-18-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_17` | 4.14.6 | `el8.x86_64` | pgdg | 151.4 KiB | [orafce_17-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.4 | `el8.x86_64` | pgdg | 150.9 KiB | [orafce_17-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.3 | `el8.x86_64` | pgdg | 150.4 KiB | [orafce_17-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.3 | `el8.x86_64` | pgdg | 150.7 KiB | [orafce_17-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.2 | `el8.x86_64` | pgdg | 150.2 KiB | [orafce_17-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.0 | `el8.x86_64` | pgdg | 148.5 KiB | [orafce_17-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.13.5 | `el8.x86_64` | pgdg | 148.1 KiB | [orafce_17-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.13.3 | `el8.x86_64` | pgdg | 147.8 KiB | [orafce_17-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.13.2 | `el8.x86_64` | pgdg | 147.6 KiB | [orafce_17-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.13.0 | `el8.x86_64` | pgdg | 147.4 KiB | [orafce_17-4.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | 4.14.6 | `el8.aarch64` | pgdg | 146.9 KiB | [orafce_17-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.14.4 | `el8.aarch64` | pgdg | 146.6 KiB | [orafce_17-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.14.3 | `el8.aarch64` | pgdg | 146.4 KiB | [orafce_17-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.14.3 | `el8.aarch64` | pgdg | 146.1 KiB | [orafce_17-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.14.2 | `el8.aarch64` | pgdg | 146.0 KiB | [orafce_17-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.14.0 | `el8.aarch64` | pgdg | 143.4 KiB | [orafce_17-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.13.5 | `el8.aarch64` | pgdg | 143.0 KiB | [orafce_17-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.13.3 | `el8.aarch64` | pgdg | 142.6 KiB | [orafce_17-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.13.2 | `el8.aarch64` | pgdg | 142.5 KiB | [orafce_17-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.13.0 | `el8.aarch64` | pgdg | 142.3 KiB | [orafce_17-4.13.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | 4.16.1 | `el9.x86_64` | pgdg | 150.0 KiB | [orafce_17-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.6 | `el9.x86_64` | pgdg | 148.8 KiB | [orafce_17-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.4 | `el9.x86_64` | pgdg | 148.9 KiB | [orafce_17-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.3 | `el9.x86_64` | pgdg | 148.5 KiB | [orafce_17-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.3 | `el9.x86_64` | pgdg | 148.6 KiB | [orafce_17-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.2 | `el9.x86_64` | pgdg | 148.4 KiB | [orafce_17-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.14.0 | `el9.x86_64` | pgdg | 143.8 KiB | [orafce_17-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.13.5 | `el9.x86_64` | pgdg | 143.5 KiB | [orafce_17-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.13.3 | `el9.x86_64` | pgdg | 143.4 KiB | [orafce_17-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.13.2 | `el9.x86_64` | pgdg | 143.1 KiB | [orafce_17-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.13.0 | `el9.x86_64` | pgdg | 143.1 KiB | [orafce_17-4.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | 4.16.1 | `el9.aarch64` | pgdg | 147.6 KiB | [orafce_17-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.6 | `el9.aarch64` | pgdg | 146.7 KiB | [orafce_17-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.4 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_17-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.3 | `el9.aarch64` | pgdg | 146.5 KiB | [orafce_17-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.3 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_17-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.2 | `el9.aarch64` | pgdg | 146.4 KiB | [orafce_17-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.14.0 | `el9.aarch64` | pgdg | 141.4 KiB | [orafce_17-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.13.5 | `el9.aarch64` | pgdg | 141.5 KiB | [orafce_17-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.13.3 | `el9.aarch64` | pgdg | 141.3 KiB | [orafce_17-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.13.2 | `el9.aarch64` | pgdg | 141.1 KiB | [orafce_17-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.13.0 | `el9.aarch64` | pgdg | 140.8 KiB | [orafce_17-4.13.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | 4.16.1 | `el10.x86_64` | pgdg | 150.8 KiB | [orafce_17-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | 4.14.6 | `el10.x86_64` | pgdg | 150.0 KiB | [orafce_17-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | 4.14.4 | `el10.x86_64` | pgdg | 149.7 KiB | [orafce_17-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | 4.14.3 | `el10.x86_64` | pgdg | 149.6 KiB | [orafce_17-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | 4.16.1 | `el10.aarch64` | pgdg | 149.0 KiB | [orafce_17-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | 4.14.6 | `el10.aarch64` | pgdg | 148.3 KiB | [orafce_17-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | 4.14.4 | `el10.aarch64` | pgdg | 148.1 KiB | [orafce_17-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | 4.14.3 | `el10.aarch64` | pgdg | 148.1 KiB | [orafce_17-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 363.1 KiB | [postgresql-17-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 355.3 KiB | [postgresql-17-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 363.8 KiB | [postgresql-17-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 356.9 KiB | [postgresql-17-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 398.1 KiB | [postgresql-17-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 389.8 KiB | [postgresql-17-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 360.3 KiB | [postgresql-17-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 354.9 KiB | [postgresql-17-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_16` | 4.9.4 | `el8.x86_64` | pgdg | 143.8 KiB | [orafce_16-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.9.3 | `el8.x86_64` | pgdg | 143.6 KiB | [orafce_16-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.9.2 | `el8.x86_64` | pgdg | 143.4 KiB | [orafce_16-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.9.1 | `el8.x86_64` | pgdg | 143.4 KiB | [orafce_16-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.9.0 | `el8.x86_64` | pgdg | 143.2 KiB | [orafce_16-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.6 | `el8.x86_64` | pgdg | 151.3 KiB | [orafce_16-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.4 | `el8.x86_64` | pgdg | 150.7 KiB | [orafce_16-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.3 | `el8.x86_64` | pgdg | 150.3 KiB | [orafce_16-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.3 | `el8.x86_64` | pgdg | 150.6 KiB | [orafce_16-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.2 | `el8.x86_64` | pgdg | 150.1 KiB | [orafce_16-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.14.0 | `el8.x86_64` | pgdg | 148.5 KiB | [orafce_16-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.13.5 | `el8.x86_64` | pgdg | 148.0 KiB | [orafce_16-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.13.3 | `el8.x86_64` | pgdg | 147.7 KiB | [orafce_16-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.13.2 | `el8.x86_64` | pgdg | 147.6 KiB | [orafce_16-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.12.0 | `el8.x86_64` | pgdg | 146.3 KiB | [orafce_16-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.11.0 | `el8.x86_64` | pgdg | 145.9 KiB | [orafce_16-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.10.3 | `el8.x86_64` | pgdg | 145.3 KiB | [orafce_16-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.10.2 | `el8.x86_64` | pgdg | 145.1 KiB | [orafce_16-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.10.0 | `el8.x86_64` | pgdg | 144.8 KiB | [orafce_16-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | 4.9.4 | `el8.aarch64` | pgdg | 139.0 KiB | [orafce_16-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.9.3 | `el8.aarch64` | pgdg | 138.9 KiB | [orafce_16-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.9.2 | `el8.aarch64` | pgdg | 138.5 KiB | [orafce_16-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.9.1 | `el8.aarch64` | pgdg | 138.4 KiB | [orafce_16-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.9.0 | `el8.aarch64` | pgdg | 138.2 KiB | [orafce_16-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.6 | `el8.aarch64` | pgdg | 147.0 KiB | [orafce_16-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.4 | `el8.aarch64` | pgdg | 146.7 KiB | [orafce_16-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.3 | `el8.aarch64` | pgdg | 146.5 KiB | [orafce_16-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.3 | `el8.aarch64` | pgdg | 146.3 KiB | [orafce_16-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.2 | `el8.aarch64` | pgdg | 146.1 KiB | [orafce_16-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.14.0 | `el8.aarch64` | pgdg | 143.3 KiB | [orafce_16-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.13.5 | `el8.aarch64` | pgdg | 142.9 KiB | [orafce_16-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.13.3 | `el8.aarch64` | pgdg | 142.6 KiB | [orafce_16-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.13.2 | `el8.aarch64` | pgdg | 142.4 KiB | [orafce_16-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.12.0 | `el8.aarch64` | pgdg | 141.3 KiB | [orafce_16-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.11.0 | `el8.aarch64` | pgdg | 140.9 KiB | [orafce_16-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.10.3 | `el8.aarch64` | pgdg | 140.3 KiB | [orafce_16-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.10.2 | `el8.aarch64` | pgdg | 140.1 KiB | [orafce_16-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.10.0 | `el8.aarch64` | pgdg | 139.8 KiB | [orafce_16-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | 4.9.4 | `el9.x86_64` | pgdg | 140.2 KiB | [orafce_16-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.9.3 | `el9.x86_64` | pgdg | 140.1 KiB | [orafce_16-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.9.2 | `el9.x86_64` | pgdg | 139.7 KiB | [orafce_16-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.9.1 | `el9.x86_64` | pgdg | 139.5 KiB | [orafce_16-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.9.0 | `el9.x86_64` | pgdg | 139.5 KiB | [orafce_16-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.16.1 | `el9.x86_64` | pgdg | 149.7 KiB | [orafce_16-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.6 | `el9.x86_64` | pgdg | 148.7 KiB | [orafce_16-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.4 | `el9.x86_64` | pgdg | 148.5 KiB | [orafce_16-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.3 | `el9.x86_64` | pgdg | 148.4 KiB | [orafce_16-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.3 | `el9.x86_64` | pgdg | 148.4 KiB | [orafce_16-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.2 | `el9.x86_64` | pgdg | 148.3 KiB | [orafce_16-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.14.0 | `el9.x86_64` | pgdg | 143.7 KiB | [orafce_16-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.13.5 | `el9.x86_64` | pgdg | 143.5 KiB | [orafce_16-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.13.3 | `el9.x86_64` | pgdg | 143.3 KiB | [orafce_16-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.13.2 | `el9.x86_64` | pgdg | 143.2 KiB | [orafce_16-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.12.0 | `el9.x86_64` | pgdg | 142.1 KiB | [orafce_16-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.11.0 | `el9.x86_64` | pgdg | 141.8 KiB | [orafce_16-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.10.3 | `el9.x86_64` | pgdg | 141.6 KiB | [orafce_16-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.10.2 | `el9.x86_64` | pgdg | 141.6 KiB | [orafce_16-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.10.0 | `el9.x86_64` | pgdg | 141.2 KiB | [orafce_16-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | 4.9.4 | `el9.aarch64` | pgdg | 137.6 KiB | [orafce_16-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.9.3 | `el9.aarch64` | pgdg | 137.5 KiB | [orafce_16-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.9.2 | `el9.aarch64` | pgdg | 137.3 KiB | [orafce_16-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.9.1 | `el9.aarch64` | pgdg | 137.2 KiB | [orafce_16-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.9.0 | `el9.aarch64` | pgdg | 137.1 KiB | [orafce_16-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.16.1 | `el9.aarch64` | pgdg | 147.6 KiB | [orafce_16-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.6 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_16-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.4 | `el9.aarch64` | pgdg | 146.7 KiB | [orafce_16-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.3 | `el9.aarch64` | pgdg | 146.5 KiB | [orafce_16-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.3 | `el9.aarch64` | pgdg | 146.7 KiB | [orafce_16-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.2 | `el9.aarch64` | pgdg | 146.4 KiB | [orafce_16-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.14.0 | `el9.aarch64` | pgdg | 141.4 KiB | [orafce_16-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.13.5 | `el9.aarch64` | pgdg | 141.5 KiB | [orafce_16-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.13.3 | `el9.aarch64` | pgdg | 141.3 KiB | [orafce_16-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.13.2 | `el9.aarch64` | pgdg | 141.2 KiB | [orafce_16-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.12.0 | `el9.aarch64` | pgdg | 140.0 KiB | [orafce_16-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.11.0 | `el9.aarch64` | pgdg | 139.5 KiB | [orafce_16-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.10.3 | `el9.aarch64` | pgdg | 139.3 KiB | [orafce_16-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.10.2 | `el9.aarch64` | pgdg | 139.1 KiB | [orafce_16-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.10.0 | `el9.aarch64` | pgdg | 138.3 KiB | [orafce_16-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | 4.16.1 | `el10.x86_64` | pgdg | 150.8 KiB | [orafce_16-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | 4.14.6 | `el10.x86_64` | pgdg | 149.8 KiB | [orafce_16-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | 4.14.4 | `el10.x86_64` | pgdg | 149.6 KiB | [orafce_16-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | 4.14.3 | `el10.x86_64` | pgdg | 149.6 KiB | [orafce_16-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | 4.16.1 | `el10.aarch64` | pgdg | 149.0 KiB | [orafce_16-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | 4.14.6 | `el10.aarch64` | pgdg | 148.2 KiB | [orafce_16-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | 4.14.4 | `el10.aarch64` | pgdg | 148.0 KiB | [orafce_16-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | 4.14.3 | `el10.aarch64` | pgdg | 148.0 KiB | [orafce_16-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 362.7 KiB | [postgresql-16-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 355.6 KiB | [postgresql-16-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 363.8 KiB | [postgresql-16-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 356.6 KiB | [postgresql-16-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 397.5 KiB | [postgresql-16-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 389.1 KiB | [postgresql-16-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 360.3 KiB | [postgresql-16-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 355.2 KiB | [postgresql-16-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_15` | 4.9.4 | `el8.x86_64` | pgdg | 145.5 KiB | [orafce_15-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.9.3 | `el8.x86_64` | pgdg | 145.4 KiB | [orafce_15-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.9.2 | `el8.x86_64` | pgdg | 145.1 KiB | [orafce_15-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.9.1 | `el8.x86_64` | pgdg | 145.0 KiB | [orafce_15-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.9.0 | `el8.x86_64` | pgdg | 144.9 KiB | [orafce_15-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.6 | `el8.x86_64` | pgdg | 151.5 KiB | [orafce_15-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.4 | `el8.x86_64` | pgdg | 150.9 KiB | [orafce_15-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.3 | `el8.x86_64` | pgdg | 150.8 KiB | [orafce_15-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.3 | `el8.x86_64` | pgdg | 150.5 KiB | [orafce_15-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.2 | `el8.x86_64` | pgdg | 150.3 KiB | [orafce_15-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.14.0 | `el8.x86_64` | pgdg | 149.8 KiB | [orafce_15-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.13.5 | `el8.x86_64` | pgdg | 149.3 KiB | [orafce_15-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.13.3 | `el8.x86_64` | pgdg | 149.0 KiB | [orafce_15-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.13.2 | `el8.x86_64` | pgdg | 148.9 KiB | [orafce_15-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.12.0 | `el8.x86_64` | pgdg | 147.6 KiB | [orafce_15-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.11.0 | `el8.x86_64` | pgdg | 147.2 KiB | [orafce_15-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.10.3 | `el8.x86_64` | pgdg | 146.9 KiB | [orafce_15-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.10.2 | `el8.x86_64` | pgdg | 146.8 KiB | [orafce_15-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.10.0 | `el8.x86_64` | pgdg | 146.4 KiB | [orafce_15-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | 4.9.4 | `el8.aarch64` | pgdg | 140.8 KiB | [orafce_15-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.9.3 | `el8.aarch64` | pgdg | 140.7 KiB | [orafce_15-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.9.2 | `el8.aarch64` | pgdg | 140.3 KiB | [orafce_15-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.9.1 | `el8.aarch64` | pgdg | 140.1 KiB | [orafce_15-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.9.0 | `el8.aarch64` | pgdg | 140.0 KiB | [orafce_15-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.6 | `el8.aarch64` | pgdg | 147.1 KiB | [orafce_15-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.4 | `el8.aarch64` | pgdg | 146.7 KiB | [orafce_15-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.3 | `el8.aarch64` | pgdg | 146.2 KiB | [orafce_15-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.3 | `el8.aarch64` | pgdg | 146.5 KiB | [orafce_15-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.2 | `el8.aarch64` | pgdg | 146.1 KiB | [orafce_15-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.14.0 | `el8.aarch64` | pgdg | 144.9 KiB | [orafce_15-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.13.5 | `el8.aarch64` | pgdg | 144.5 KiB | [orafce_15-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.13.3 | `el8.aarch64` | pgdg | 144.2 KiB | [orafce_15-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.13.2 | `el8.aarch64` | pgdg | 144.0 KiB | [orafce_15-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.12.0 | `el8.aarch64` | pgdg | 142.7 KiB | [orafce_15-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.11.0 | `el8.aarch64` | pgdg | 142.3 KiB | [orafce_15-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.10.3 | `el8.aarch64` | pgdg | 142.1 KiB | [orafce_15-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.10.2 | `el8.aarch64` | pgdg | 141.9 KiB | [orafce_15-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.10.0 | `el8.aarch64` | pgdg | 141.5 KiB | [orafce_15-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | 4.9.4 | `el9.x86_64` | pgdg | 144.6 KiB | [orafce_15-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.9.3 | `el9.x86_64` | pgdg | 144.6 KiB | [orafce_15-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.9.2 | `el9.x86_64` | pgdg | 144.2 KiB | [orafce_15-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.9.1 | `el9.x86_64` | pgdg | 144.2 KiB | [orafce_15-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.9.0 | `el9.x86_64` | pgdg | 144.2 KiB | [orafce_15-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.16.1 | `el9.x86_64` | pgdg | 150.0 KiB | [orafce_15-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.6 | `el9.x86_64` | pgdg | 148.9 KiB | [orafce_15-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.4 | `el9.x86_64` | pgdg | 148.6 KiB | [orafce_15-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.3 | `el9.x86_64` | pgdg | 148.7 KiB | [orafce_15-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.3 | `el9.x86_64` | pgdg | 148.5 KiB | [orafce_15-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.2 | `el9.x86_64` | pgdg | 148.5 KiB | [orafce_15-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.14.0 | `el9.x86_64` | pgdg | 148.6 KiB | [orafce_15-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.13.5 | `el9.x86_64` | pgdg | 148.0 KiB | [orafce_15-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.13.3 | `el9.x86_64` | pgdg | 148.1 KiB | [orafce_15-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.13.2 | `el9.x86_64` | pgdg | 147.9 KiB | [orafce_15-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.12.0 | `el9.x86_64` | pgdg | 146.5 KiB | [orafce_15-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.11.0 | `el9.x86_64` | pgdg | 146.1 KiB | [orafce_15-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.10.3 | `el9.x86_64` | pgdg | 146.1 KiB | [orafce_15-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.10.2 | `el9.x86_64` | pgdg | 146.0 KiB | [orafce_15-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.10.0 | `el9.x86_64` | pgdg | 145.9 KiB | [orafce_15-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | 4.9.4 | `el9.aarch64` | pgdg | 142.2 KiB | [orafce_15-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.9.3 | `el9.aarch64` | pgdg | 142.2 KiB | [orafce_15-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.9.2 | `el9.aarch64` | pgdg | 141.8 KiB | [orafce_15-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.9.1 | `el9.aarch64` | pgdg | 141.5 KiB | [orafce_15-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.9.0 | `el9.aarch64` | pgdg | 141.5 KiB | [orafce_15-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.16.1 | `el9.aarch64` | pgdg | 147.7 KiB | [orafce_15-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.6 | `el9.aarch64` | pgdg | 146.8 KiB | [orafce_15-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.4 | `el9.aarch64` | pgdg | 146.8 KiB | [orafce_15-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.3 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_15-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.3 | `el9.aarch64` | pgdg | 146.8 KiB | [orafce_15-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.2 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_15-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.14.0 | `el9.aarch64` | pgdg | 146.0 KiB | [orafce_15-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.13.5 | `el9.aarch64` | pgdg | 145.7 KiB | [orafce_15-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.13.3 | `el9.aarch64` | pgdg | 145.5 KiB | [orafce_15-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.13.2 | `el9.aarch64` | pgdg | 145.3 KiB | [orafce_15-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.12.0 | `el9.aarch64` | pgdg | 143.8 KiB | [orafce_15-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.11.0 | `el9.aarch64` | pgdg | 143.4 KiB | [orafce_15-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.10.3 | `el9.aarch64` | pgdg | 143.7 KiB | [orafce_15-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.10.2 | `el9.aarch64` | pgdg | 143.6 KiB | [orafce_15-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.10.0 | `el9.aarch64` | pgdg | 142.7 KiB | [orafce_15-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | 4.16.1 | `el10.x86_64` | pgdg | 150.9 KiB | [orafce_15-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | 4.14.6 | `el10.x86_64` | pgdg | 150.2 KiB | [orafce_15-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | 4.14.4 | `el10.x86_64` | pgdg | 150.1 KiB | [orafce_15-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | 4.14.3 | `el10.x86_64` | pgdg | 150.0 KiB | [orafce_15-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | 4.16.1 | `el10.aarch64` | pgdg | 149.4 KiB | [orafce_15-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | 4.14.6 | `el10.aarch64` | pgdg | 148.5 KiB | [orafce_15-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | 4.14.4 | `el10.aarch64` | pgdg | 148.3 KiB | [orafce_15-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | 4.14.3 | `el10.aarch64` | pgdg | 148.3 KiB | [orafce_15-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 365.1 KiB | [postgresql-15-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 357.6 KiB | [postgresql-15-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 365.8 KiB | [postgresql-15-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 359.2 KiB | [postgresql-15-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 403.9 KiB | [postgresql-15-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 395.6 KiB | [postgresql-15-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 364.9 KiB | [postgresql-15-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 358.5 KiB | [postgresql-15-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_14` | 4.9.4 | `el8.x86_64` | pgdg | 146.7 KiB | [orafce_14-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.9.3 | `el8.x86_64` | pgdg | 146.5 KiB | [orafce_14-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.9.2 | `el8.x86_64` | pgdg | 146.3 KiB | [orafce_14-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.9.1 | `el8.x86_64` | pgdg | 146.2 KiB | [orafce_14-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.9.0 | `el8.x86_64` | pgdg | 146.0 KiB | [orafce_14-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.6 | `el8.x86_64` | pgdg | 152.5 KiB | [orafce_14-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.4 | `el8.x86_64` | pgdg | 152.0 KiB | [orafce_14-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.3 | `el8.x86_64` | pgdg | 151.8 KiB | [orafce_14-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.3 | `el8.x86_64` | pgdg | 151.5 KiB | [orafce_14-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.2 | `el8.x86_64` | pgdg | 151.4 KiB | [orafce_14-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.14.0 | `el8.x86_64` | pgdg | 150.8 KiB | [orafce_14-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.13.5 | `el8.x86_64` | pgdg | 150.4 KiB | [orafce_14-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.13.3 | `el8.x86_64` | pgdg | 150.1 KiB | [orafce_14-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.13.2 | `el8.x86_64` | pgdg | 149.9 KiB | [orafce_14-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.12.0 | `el8.x86_64` | pgdg | 148.6 KiB | [orafce_14-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.11.0 | `el8.x86_64` | pgdg | 148.3 KiB | [orafce_14-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.10.3 | `el8.x86_64` | pgdg | 148.0 KiB | [orafce_14-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.10.2 | `el8.x86_64` | pgdg | 147.6 KiB | [orafce_14-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.10.0 | `el8.x86_64` | pgdg | 147.3 KiB | [orafce_14-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | 4.9.4 | `el8.aarch64` | pgdg | 141.7 KiB | [orafce_14-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.9.3 | `el8.aarch64` | pgdg | 141.5 KiB | [orafce_14-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.9.2 | `el8.aarch64` | pgdg | 141.2 KiB | [orafce_14-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.9.1 | `el8.aarch64` | pgdg | 140.9 KiB | [orafce_14-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.9.0 | `el8.aarch64` | pgdg | 140.8 KiB | [orafce_14-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.6 | `el8.aarch64` | pgdg | 147.9 KiB | [orafce_14-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.4 | `el8.aarch64` | pgdg | 147.5 KiB | [orafce_14-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.3 | `el8.aarch64` | pgdg | 147.0 KiB | [orafce_14-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.3 | `el8.aarch64` | pgdg | 147.3 KiB | [orafce_14-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.2 | `el8.aarch64` | pgdg | 146.9 KiB | [orafce_14-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.14.0 | `el8.aarch64` | pgdg | 145.8 KiB | [orafce_14-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.13.5 | `el8.aarch64` | pgdg | 145.4 KiB | [orafce_14-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.13.3 | `el8.aarch64` | pgdg | 145.1 KiB | [orafce_14-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.13.2 | `el8.aarch64` | pgdg | 145.0 KiB | [orafce_14-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.12.0 | `el8.aarch64` | pgdg | 143.6 KiB | [orafce_14-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.11.0 | `el8.aarch64` | pgdg | 143.2 KiB | [orafce_14-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.10.3 | `el8.aarch64` | pgdg | 143.0 KiB | [orafce_14-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.10.2 | `el8.aarch64` | pgdg | 142.8 KiB | [orafce_14-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.10.0 | `el8.aarch64` | pgdg | 142.3 KiB | [orafce_14-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | 4.9.4 | `el9.x86_64` | pgdg | 146.0 KiB | [orafce_14-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.9.3 | `el9.x86_64` | pgdg | 145.9 KiB | [orafce_14-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.9.2 | `el9.x86_64` | pgdg | 145.6 KiB | [orafce_14-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.9.1 | `el9.x86_64` | pgdg | 145.5 KiB | [orafce_14-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.9.0 | `el9.x86_64` | pgdg | 145.4 KiB | [orafce_14-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.16.1 | `el9.x86_64` | pgdg | 151.4 KiB | [orafce_14-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.6 | `el9.x86_64` | pgdg | 150.1 KiB | [orafce_14-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.4 | `el9.x86_64` | pgdg | 149.9 KiB | [orafce_14-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.3 | `el9.x86_64` | pgdg | 149.8 KiB | [orafce_14-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.3 | `el9.x86_64` | pgdg | 149.9 KiB | [orafce_14-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.2 | `el9.x86_64` | pgdg | 149.7 KiB | [orafce_14-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.14.0 | `el9.x86_64` | pgdg | 149.7 KiB | [orafce_14-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.13.5 | `el9.x86_64` | pgdg | 149.5 KiB | [orafce_14-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.13.3 | `el9.x86_64` | pgdg | 149.3 KiB | [orafce_14-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.13.2 | `el9.x86_64` | pgdg | 149.1 KiB | [orafce_14-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.12.0 | `el9.x86_64` | pgdg | 147.8 KiB | [orafce_14-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.11.0 | `el9.x86_64` | pgdg | 147.4 KiB | [orafce_14-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.10.3 | `el9.x86_64` | pgdg | 147.7 KiB | [orafce_14-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.10.2 | `el9.x86_64` | pgdg | 147.6 KiB | [orafce_14-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.10.0 | `el9.x86_64` | pgdg | 147.1 KiB | [orafce_14-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | 4.9.4 | `el9.aarch64` | pgdg | 142.8 KiB | [orafce_14-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.9.3 | `el9.aarch64` | pgdg | 142.8 KiB | [orafce_14-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.9.2 | `el9.aarch64` | pgdg | 142.5 KiB | [orafce_14-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.9.1 | `el9.aarch64` | pgdg | 142.4 KiB | [orafce_14-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.9.0 | `el9.aarch64` | pgdg | 142.4 KiB | [orafce_14-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.16.1 | `el9.aarch64` | pgdg | 148.7 KiB | [orafce_14-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.6 | `el9.aarch64` | pgdg | 147.7 KiB | [orafce_14-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.4 | `el9.aarch64` | pgdg | 147.7 KiB | [orafce_14-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.3 | `el9.aarch64` | pgdg | 147.5 KiB | [orafce_14-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.3 | `el9.aarch64` | pgdg | 147.7 KiB | [orafce_14-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.2 | `el9.aarch64` | pgdg | 147.5 KiB | [orafce_14-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.14.0 | `el9.aarch64` | pgdg | 146.8 KiB | [orafce_14-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.13.5 | `el9.aarch64` | pgdg | 146.5 KiB | [orafce_14-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.13.3 | `el9.aarch64` | pgdg | 146.2 KiB | [orafce_14-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.13.2 | `el9.aarch64` | pgdg | 146.1 KiB | [orafce_14-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.12.0 | `el9.aarch64` | pgdg | 145.1 KiB | [orafce_14-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.11.0 | `el9.aarch64` | pgdg | 144.5 KiB | [orafce_14-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.10.3 | `el9.aarch64` | pgdg | 144.6 KiB | [orafce_14-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.10.2 | `el9.aarch64` | pgdg | 144.4 KiB | [orafce_14-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.10.0 | `el9.aarch64` | pgdg | 143.6 KiB | [orafce_14-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | 4.16.1 | `el10.x86_64` | pgdg | 152.3 KiB | [orafce_14-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | 4.14.6 | `el10.x86_64` | pgdg | 151.4 KiB | [orafce_14-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | 4.14.4 | `el10.x86_64` | pgdg | 151.2 KiB | [orafce_14-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | 4.14.3 | `el10.x86_64` | pgdg | 151.1 KiB | [orafce_14-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | 4.16.1 | `el10.aarch64` | pgdg | 150.3 KiB | [orafce_14-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | 4.14.6 | `el10.aarch64` | pgdg | 149.6 KiB | [orafce_14-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | 4.14.4 | `el10.aarch64` | pgdg | 149.4 KiB | [orafce_14-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | 4.14.3 | `el10.aarch64` | pgdg | 149.3 KiB | [orafce_14-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 368.0 KiB | [postgresql-14-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 360.6 KiB | [postgresql-14-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 369.0 KiB | [postgresql-14-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 362.1 KiB | [postgresql-14-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 404.1 KiB | [postgresql-14-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 395.1 KiB | [postgresql-14-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 367.8 KiB | [postgresql-14-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 361.3 KiB | [postgresql-14-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_13` | 4.9.4 | `el8.x86_64` | pgdg | 144.5 KiB | [orafce_13-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.9.3 | `el8.x86_64` | pgdg | 144.4 KiB | [orafce_13-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.9.2 | `el8.x86_64` | pgdg | 144.3 KiB | [orafce_13-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.9.1 | `el8.x86_64` | pgdg | 144.1 KiB | [orafce_13-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.9.0 | `el8.x86_64` | pgdg | 144.0 KiB | [orafce_13-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.6 | `el8.x86_64` | pgdg | 150.6 KiB | [orafce_13-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.4 | `el8.x86_64` | pgdg | 150.1 KiB | [orafce_13-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.3 | `el8.x86_64` | pgdg | 149.7 KiB | [orafce_13-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.3 | `el8.x86_64` | pgdg | 150.0 KiB | [orafce_13-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.2 | `el8.x86_64` | pgdg | 149.5 KiB | [orafce_13-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.14.0 | `el8.x86_64` | pgdg | 148.9 KiB | [orafce_13-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.13.5 | `el8.x86_64` | pgdg | 148.4 KiB | [orafce_13-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.13.3 | `el8.x86_64` | pgdg | 148.0 KiB | [orafce_13-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.13.2 | `el8.x86_64` | pgdg | 147.9 KiB | [orafce_13-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.12.0 | `el8.x86_64` | pgdg | 146.7 KiB | [orafce_13-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.11.0 | `el8.x86_64` | pgdg | 146.3 KiB | [orafce_13-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.10.3 | `el8.x86_64` | pgdg | 146.0 KiB | [orafce_13-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.10.2 | `el8.x86_64` | pgdg | 145.9 KiB | [orafce_13-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.10.0 | `el8.x86_64` | pgdg | 145.5 KiB | [orafce_13-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/orafce_13-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_13` | 4.9.4 | `el8.aarch64` | pgdg | 141.2 KiB | [orafce_13-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.9.3 | `el8.aarch64` | pgdg | 141.2 KiB | [orafce_13-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.9.2 | `el8.aarch64` | pgdg | 140.9 KiB | [orafce_13-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.9.1 | `el8.aarch64` | pgdg | 140.7 KiB | [orafce_13-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.9.0 | `el8.aarch64` | pgdg | 140.5 KiB | [orafce_13-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.6 | `el8.aarch64` | pgdg | 147.7 KiB | [orafce_13-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.4 | `el8.aarch64` | pgdg | 147.2 KiB | [orafce_13-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.3 | `el8.aarch64` | pgdg | 147.0 KiB | [orafce_13-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.3 | `el8.aarch64` | pgdg | 146.8 KiB | [orafce_13-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.2 | `el8.aarch64` | pgdg | 146.6 KiB | [orafce_13-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.14.0 | `el8.aarch64` | pgdg | 145.5 KiB | [orafce_13-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.13.5 | `el8.aarch64` | pgdg | 145.1 KiB | [orafce_13-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.13.3 | `el8.aarch64` | pgdg | 144.8 KiB | [orafce_13-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.13.2 | `el8.aarch64` | pgdg | 144.6 KiB | [orafce_13-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.12.0 | `el8.aarch64` | pgdg | 143.3 KiB | [orafce_13-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.11.0 | `el8.aarch64` | pgdg | 142.9 KiB | [orafce_13-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.10.3 | `el8.aarch64` | pgdg | 142.6 KiB | [orafce_13-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.10.2 | `el8.aarch64` | pgdg | 142.5 KiB | [orafce_13-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.10.0 | `el8.aarch64` | pgdg | 142.0 KiB | [orafce_13-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/orafce_13-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_13` | 4.9.4 | `el9.x86_64` | pgdg | 145.8 KiB | [orafce_13-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.9.3 | `el9.x86_64` | pgdg | 145.8 KiB | [orafce_13-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.9.2 | `el9.x86_64` | pgdg | 145.0 KiB | [orafce_13-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.9.1 | `el9.x86_64` | pgdg | 144.6 KiB | [orafce_13-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.9.0 | `el9.x86_64` | pgdg | 144.6 KiB | [orafce_13-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.16.1 | `el9.x86_64` | pgdg | 150.4 KiB | [orafce_13-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.6 | `el9.x86_64` | pgdg | 149.6 KiB | [orafce_13-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.4 | `el9.x86_64` | pgdg | 149.4 KiB | [orafce_13-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.3 | `el9.x86_64` | pgdg | 149.4 KiB | [orafce_13-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.3 | `el9.x86_64` | pgdg | 149.3 KiB | [orafce_13-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.2 | `el9.x86_64` | pgdg | 149.3 KiB | [orafce_13-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.14.0 | `el9.x86_64` | pgdg | 149.0 KiB | [orafce_13-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.13.5 | `el9.x86_64` | pgdg | 148.9 KiB | [orafce_13-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.13.3 | `el9.x86_64` | pgdg | 149.0 KiB | [orafce_13-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.13.2 | `el9.x86_64` | pgdg | 148.4 KiB | [orafce_13-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.12.0 | `el9.x86_64` | pgdg | 147.5 KiB | [orafce_13-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.11.0 | `el9.x86_64` | pgdg | 147.1 KiB | [orafce_13-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.10.3 | `el9.x86_64` | pgdg | 147.1 KiB | [orafce_13-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.10.2 | `el9.x86_64` | pgdg | 147.1 KiB | [orafce_13-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.10.0 | `el9.x86_64` | pgdg | 146.4 KiB | [orafce_13-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/orafce_13-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_13` | 4.9.4 | `el9.aarch64` | pgdg | 142.6 KiB | [orafce_13-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.9.3 | `el9.aarch64` | pgdg | 142.5 KiB | [orafce_13-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.9.2 | `el9.aarch64` | pgdg | 142.3 KiB | [orafce_13-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.9.1 | `el9.aarch64` | pgdg | 142.1 KiB | [orafce_13-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.9.0 | `el9.aarch64` | pgdg | 142.2 KiB | [orafce_13-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.16.1 | `el9.aarch64` | pgdg | 148.4 KiB | [orafce_13-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.6 | `el9.aarch64` | pgdg | 147.5 KiB | [orafce_13-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.4 | `el9.aarch64` | pgdg | 147.6 KiB | [orafce_13-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.3 | `el9.aarch64` | pgdg | 147.6 KiB | [orafce_13-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.3 | `el9.aarch64` | pgdg | 147.4 KiB | [orafce_13-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.2 | `el9.aarch64` | pgdg | 147.4 KiB | [orafce_13-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.14.0 | `el9.aarch64` | pgdg | 146.6 KiB | [orafce_13-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.13.5 | `el9.aarch64` | pgdg | 146.4 KiB | [orafce_13-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.13.3 | `el9.aarch64` | pgdg | 145.9 KiB | [orafce_13-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.13.2 | `el9.aarch64` | pgdg | 145.8 KiB | [orafce_13-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.12.0 | `el9.aarch64` | pgdg | 144.7 KiB | [orafce_13-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.11.0 | `el9.aarch64` | pgdg | 144.3 KiB | [orafce_13-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.10.3 | `el9.aarch64` | pgdg | 144.3 KiB | [orafce_13-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.10.2 | `el9.aarch64` | pgdg | 144.2 KiB | [orafce_13-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.10.0 | `el9.aarch64` | pgdg | 143.5 KiB | [orafce_13-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/orafce_13-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_13` | 4.16.1 | `el10.x86_64` | pgdg | 151.8 KiB | [orafce_13-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/orafce_13-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_13` | 4.14.6 | `el10.x86_64` | pgdg | 151.0 KiB | [orafce_13-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/orafce_13-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_13` | 4.14.4 | `el10.x86_64` | pgdg | 150.8 KiB | [orafce_13-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/orafce_13-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_13` | 4.14.3 | `el10.x86_64` | pgdg | 150.8 KiB | [orafce_13-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/orafce_13-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_13` | 4.16.1 | `el10.aarch64` | pgdg | 150.3 KiB | [orafce_13-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/orafce_13-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_13` | 4.14.6 | `el10.aarch64` | pgdg | 149.3 KiB | [orafce_13-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/orafce_13-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_13` | 4.14.4 | `el10.aarch64` | pgdg | 149.1 KiB | [orafce_13-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/orafce_13-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_13` | 4.14.3 | `el10.aarch64` | pgdg | 149.1 KiB | [orafce_13-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-aarch64/orafce_13-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-orafce` | 4.16.1 | `d12.x86_64` | pgdg | 366.9 KiB | [postgresql-13-orafce_4.16.1-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg12+1_amd64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `d12.aarch64` | pgdg | 358.9 KiB | [postgresql-13-orafce_4.16.1-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg12+1_arm64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `d13.x86_64` | pgdg | 368.3 KiB | [postgresql-13-orafce_4.16.1-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg13+1_amd64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `d13.aarch64` | pgdg | 360.7 KiB | [postgresql-13-orafce_4.16.1-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg13+1_arm64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `u22.x86_64` | pgdg | 401.3 KiB | [postgresql-13-orafce_4.16.1-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg22.04+1_amd64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `u22.aarch64` | pgdg | 392.3 KiB | [postgresql-13-orafce_4.16.1-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg22.04+1_arm64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `u24.x86_64` | pgdg | 366.9 KiB | [postgresql-13-orafce_4.16.1-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg24.04+1_amd64.deb) |
| `postgresql-13-orafce` | 4.16.1 | `u24.aarch64` | pgdg | 360.2 KiB | [postgresql-13-orafce_4.16.1-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-13-orafce_4.16.1-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/orafce/orafce" title="Repository" icon="github" subtitle="github.com/orafce/orafce" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install orafce; # install by extension name, for the current active PG version
pig ext install orafce; # install via package alias, for the active PG version
pig ext install orafce -v 18;   # install for PG 18
pig ext install orafce -v 17;   # install for PG 17
pig ext install orafce -v 16;   # install for PG 16
pig ext install orafce -v 15;   # install for PG 15
pig ext install orafce -v 14;   # install for PG 14
pig ext install orafce -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION orafce;
```

