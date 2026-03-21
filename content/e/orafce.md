---
title: "orafce"
linkTitle: "orafce"
description: "Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS"
weight: 9100
categories: ["SIM"]
width: full
---

[**orafce**](https://github.com/orafce/orafce) : Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9100** | {{< badge content="orafce" link="https://github.com/orafce/orafce" >}} | {{< ext "orafce" >}} | `4.16.5` | {{< category "SIM" >}} | {{< license "BSD 0-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "oracle_fdw" >}} {{< ext "pgtt" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "db_migrator" >}} |

> [!Note] el llvmjit deps break


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.16.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `orafce` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.16.2` | {{< bg "18" "orafce_18" "green" >}} {{< bg "17" "orafce_17" "green" >}} {{< bg "16" "orafce_16" "green" >}} {{< bg "15" "orafce_15" "green" >}} {{< bg "14" "orafce_14" "green" >}} | `orafce_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `4.16.5` | {{< bg "18" "postgresql-18-orafce" "green" >}} {{< bg "17" "postgresql-17-orafce" "green" >}} {{< bg "16" "postgresql-16-orafce" "green" >}} {{< bg "15" "postgresql-15-orafce" "green" >}} {{< bg "14" "postgresql-14-orafce" "green" >}} | `postgresql-$v-orafce` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 21" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 12" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 21" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 21" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 21" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 22" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 22" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 22" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 13" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 22" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 22" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 22" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 4.16.5" "orafce_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 4.16.5" "orafce_14 : AVAIL 6" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 4.16.5" "postgresql-18-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-17-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-16-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-15-orafce : AVAIL 3" "blue" >}} | {{< bg "PGDG 4.16.5" "postgresql-14-orafce : AVAIL 3" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_18` | `4.16.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.1 KiB | [orafce_18-4.16.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.16.5-1PGDG.rhel8.10.x86_64.rpm) |
| `orafce_18` | `4.16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.6 KiB | [orafce_18-4.16.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.16.2-2PGDG.rhel8.x86_64.rpm) |
| `orafce_18` | `4.14.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.3 KiB | [orafce_18-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_18` | `4.14.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.3 KiB | [orafce_18-4.14.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/orafce_18-4.14.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_18` | `4.16.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.7 KiB | [orafce_18-4.16.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.16.5-1PGDG.rhel8.10.aarch64.rpm) |
| `orafce_18` | `4.16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.2 KiB | [orafce_18-4.16.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.16.2-2PGDG.rhel8.aarch64.rpm) |
| `orafce_18` | `4.14.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.9 KiB | [orafce_18-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_18` | `4.14.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.0 KiB | [orafce_18-4.14.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/orafce_18-4.14.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_18` | `4.16.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.2 KiB | [orafce_18-4.16.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.16.5-1PGDG.rhel9.7.x86_64.rpm) |
| `orafce_18` | `4.16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [orafce_18-4.16.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.16.2-2PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | `4.16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.0 KiB | [orafce_18-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | `4.14.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.9 KiB | [orafce_18-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | `4.14.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.7 KiB | [orafce_18-4.14.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/orafce_18-4.14.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_18` | `4.16.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 148.0 KiB | [orafce_18-4.16.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.16.5-1PGDG.rhel9.7.aarch64.rpm) |
| `orafce_18` | `4.16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 148.0 KiB | [orafce_18-4.16.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.16.2-2PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | `4.16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.7 KiB | [orafce_18-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | `4.14.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_18-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | `4.14.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_18-4.14.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/orafce_18-4.14.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_18` | `4.16.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.9 KiB | [orafce_18-4.16.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.16.5-1PGDG.rhel10.1.x86_64.rpm) |
| `orafce_18` | `4.16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_18-4.16.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.16.2-2PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | `4.16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.9 KiB | [orafce_18-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | `4.14.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.1 KiB | [orafce_18-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | `4.14.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.9 KiB | [orafce_18-4.14.5-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/orafce_18-4.14.5-1PGDG.rhel10.x86_64.rpm) |
| `orafce_18` | `4.16.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.1 KiB | [orafce_18-4.16.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.16.5-1PGDG.rhel10.1.aarch64.rpm) |
| `orafce_18` | `4.16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.0 KiB | [orafce_18-4.16.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.16.2-2PGDG.rhel10.aarch64.rpm) |
| `orafce_18` | `4.16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.2 KiB | [orafce_18-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_18` | `4.14.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.3 KiB | [orafce_18-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_18` | `4.14.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.3 KiB | [orafce_18-4.14.5-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/orafce_18-4.14.5-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-orafce` | `4.16.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.6 KiB | [postgresql-18-orafce_4.16.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg12+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.4 KiB | [postgresql-18-orafce_4.16.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.9 KiB | [postgresql-18-orafce_4.16.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg12+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.5 KiB | [postgresql-18-orafce_4.16.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg12+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.6 KiB | [postgresql-18-orafce_4.16.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.5 KiB | [postgresql-18-orafce_4.16.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg12+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.6 KiB | [postgresql-18-orafce_4.16.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg13+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 364.4 KiB | [postgresql-18-orafce_4.16.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.7 KiB | [postgresql-18-orafce_4.16.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg13+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.8 KiB | [postgresql-18-orafce_4.16.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg13+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.5 KiB | [postgresql-18-orafce_4.16.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.8 KiB | [postgresql-18-orafce_4.16.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg13+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 368.3 KiB | [postgresql-18-orafce_4.16.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 368.6 KiB | [postgresql-18-orafce_4.16.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 368.2 KiB | [postgresql-18-orafce_4.16.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 360.4 KiB | [postgresql-18-orafce_4.16.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 360.2 KiB | [postgresql-18-orafce_4.16.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 360.2 KiB | [postgresql-18-orafce_4.16.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.4 KiB | [postgresql-18-orafce_4.16.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.1 KiB | [postgresql-18-orafce_4.16.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.3 KiB | [postgresql-18-orafce_4.16.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-orafce` | `4.16.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.9 KiB | [postgresql-18-orafce_4.16.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.8 KiB | [postgresql-18-orafce_4.16.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-orafce` | `4.16.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 355.0 KiB | [postgresql-18-orafce_4.16.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-18-orafce_4.16.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_17` | `4.16.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.0 KiB | [orafce_17-4.16.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.16.5-1PGDG.rhel8.10.x86_64.rpm) |
| `orafce_17` | `4.16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.6 KiB | [orafce_17-4.16.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.16.2-2PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.4 KiB | [orafce_17-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.9 KiB | [orafce_17-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.7 KiB | [orafce_17-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.4 KiB | [orafce_17-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.2 KiB | [orafce_17-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.14.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.5 KiB | [orafce_17-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.13.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.1 KiB | [orafce_17-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.13.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.8 KiB | [orafce_17-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.13.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.6 KiB | [orafce_17-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.13.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.4 KiB | [orafce_17-4.13.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/orafce_17-4.13.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_17` | `4.16.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.7 KiB | [orafce_17-4.16.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.16.5-1PGDG.rhel8.10.aarch64.rpm) |
| `orafce_17` | `4.16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.2 KiB | [orafce_17-4.16.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.16.2-2PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.9 KiB | [orafce_17-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.6 KiB | [orafce_17-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.4 KiB | [orafce_17-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.1 KiB | [orafce_17-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.0 KiB | [orafce_17-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.14.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.4 KiB | [orafce_17-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.13.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.0 KiB | [orafce_17-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.13.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.6 KiB | [orafce_17-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.13.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.5 KiB | [orafce_17-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.13.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.3 KiB | [orafce_17-4.13.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/orafce_17-4.13.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_17` | `4.16.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [orafce_17-4.16.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.16.5-1PGDG.rhel9.7.x86_64.rpm) |
| `orafce_17` | `4.16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [orafce_17-4.16.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.16.2-2PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.0 KiB | [orafce_17-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.8 KiB | [orafce_17-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.9 KiB | [orafce_17-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.6 KiB | [orafce_17-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.5 KiB | [orafce_17-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.4 KiB | [orafce_17-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.14.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.8 KiB | [orafce_17-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.13.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.5 KiB | [orafce_17-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.13.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.4 KiB | [orafce_17-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.13.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.1 KiB | [orafce_17-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.13.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.1 KiB | [orafce_17-4.13.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/orafce_17-4.13.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_17` | `4.16.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.9 KiB | [orafce_17-4.16.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.16.5-1PGDG.rhel9.7.aarch64.rpm) |
| `orafce_17` | `4.16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.9 KiB | [orafce_17-4.16.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.16.2-2PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.6 KiB | [orafce_17-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.7 KiB | [orafce_17-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_17-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_17-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.5 KiB | [orafce_17-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.4 KiB | [orafce_17-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.14.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.4 KiB | [orafce_17-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.13.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.5 KiB | [orafce_17-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.13.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.3 KiB | [orafce_17-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.13.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.1 KiB | [orafce_17-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.13.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 140.8 KiB | [orafce_17-4.13.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/orafce_17-4.13.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_17` | `4.16.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_17-4.16.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.16.5-1PGDG.rhel10.1.x86_64.rpm) |
| `orafce_17` | `4.16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.6 KiB | [orafce_17-4.16.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.16.2-2PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | `4.16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_17-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | `4.14.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.0 KiB | [orafce_17-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | `4.14.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.7 KiB | [orafce_17-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | `4.14.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.6 KiB | [orafce_17-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/orafce_17-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_17` | `4.16.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.8 KiB | [orafce_17-4.16.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.16.5-1PGDG.rhel10.1.aarch64.rpm) |
| `orafce_17` | `4.16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.9 KiB | [orafce_17-4.16.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.16.2-2PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | `4.16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.0 KiB | [orafce_17-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | `4.14.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.3 KiB | [orafce_17-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | `4.14.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.1 KiB | [orafce_17-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_17` | `4.14.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.1 KiB | [orafce_17-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/orafce_17-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-orafce` | `4.16.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.2 KiB | [postgresql-17-orafce_4.16.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg12+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.2 KiB | [postgresql-17-orafce_4.16.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.9 KiB | [postgresql-17-orafce_4.16.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg12+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.0 KiB | [postgresql-17-orafce_4.16.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg12+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 354.7 KiB | [postgresql-17-orafce_4.16.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.3 KiB | [postgresql-17-orafce_4.16.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg12+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.1 KiB | [postgresql-17-orafce_4.16.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg13+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.0 KiB | [postgresql-17-orafce_4.16.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 364.0 KiB | [postgresql-17-orafce_4.16.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg13+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.1 KiB | [postgresql-17-orafce_4.16.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg13+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.3 KiB | [postgresql-17-orafce_4.16.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.6 KiB | [postgresql-17-orafce_4.16.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg13+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 398.2 KiB | [postgresql-17-orafce_4.16.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 398.2 KiB | [postgresql-17-orafce_4.16.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 398.3 KiB | [postgresql-17-orafce_4.16.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 390.0 KiB | [postgresql-17-orafce_4.16.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.9 KiB | [postgresql-17-orafce_4.16.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 390.2 KiB | [postgresql-17-orafce_4.16.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 359.8 KiB | [postgresql-17-orafce_4.16.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 359.7 KiB | [postgresql-17-orafce_4.16.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.5 KiB | [postgresql-17-orafce_4.16.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-orafce` | `4.16.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.4 KiB | [postgresql-17-orafce_4.16.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.4 KiB | [postgresql-17-orafce_4.16.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-orafce` | `4.16.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 355.0 KiB | [postgresql-17-orafce_4.16.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-17-orafce_4.16.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_16` | `4.16.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.0 KiB | [orafce_16-4.16.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.16.5-1PGDG.rhel8.10.x86_64.rpm) |
| `orafce_16` | `4.16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.6 KiB | [orafce_16-4.16.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.16.2-2PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.3 KiB | [orafce_16-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.7 KiB | [orafce_16-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.6 KiB | [orafce_16-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.3 KiB | [orafce_16-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.1 KiB | [orafce_16-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.14.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.5 KiB | [orafce_16-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.13.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.0 KiB | [orafce_16-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.13.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.7 KiB | [orafce_16-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.13.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.6 KiB | [orafce_16-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.3 KiB | [orafce_16-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.9 KiB | [orafce_16-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.3 KiB | [orafce_16-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.1 KiB | [orafce_16-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 144.8 KiB | [orafce_16-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.8 KiB | [orafce_16-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.6 KiB | [orafce_16-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.4 KiB | [orafce_16-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.4 KiB | [orafce_16-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 143.2 KiB | [orafce_16-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/orafce_16-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_16` | `4.16.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.7 KiB | [orafce_16-4.16.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.16.5-1PGDG.rhel8.10.aarch64.rpm) |
| `orafce_16` | `4.16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.3 KiB | [orafce_16-4.16.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.16.2-2PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.0 KiB | [orafce_16-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.7 KiB | [orafce_16-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.5 KiB | [orafce_16-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.3 KiB | [orafce_16-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.1 KiB | [orafce_16-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.14.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.3 KiB | [orafce_16-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.13.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.9 KiB | [orafce_16-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.13.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.6 KiB | [orafce_16-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.13.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.4 KiB | [orafce_16-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.3 KiB | [orafce_16-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.9 KiB | [orafce_16-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.3 KiB | [orafce_16-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.1 KiB | [orafce_16-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 139.8 KiB | [orafce_16-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 139.0 KiB | [orafce_16-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 138.9 KiB | [orafce_16-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 138.5 KiB | [orafce_16-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 138.4 KiB | [orafce_16-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 138.2 KiB | [orafce_16-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/orafce_16-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_16` | `4.16.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.9 KiB | [orafce_16-4.16.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.16.5-1PGDG.rhel9.7.x86_64.rpm) |
| `orafce_16` | `4.16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.8 KiB | [orafce_16-4.16.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.16.2-2PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.7 KiB | [orafce_16-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.7 KiB | [orafce_16-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.5 KiB | [orafce_16-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.4 KiB | [orafce_16-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.4 KiB | [orafce_16-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.3 KiB | [orafce_16-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.14.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.7 KiB | [orafce_16-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.13.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.5 KiB | [orafce_16-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.13.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.3 KiB | [orafce_16-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.13.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 143.2 KiB | [orafce_16-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 142.1 KiB | [orafce_16-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 141.8 KiB | [orafce_16-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.10.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 141.6 KiB | [orafce_16-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.10.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 141.6 KiB | [orafce_16-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.10.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 141.2 KiB | [orafce_16-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.2 KiB | [orafce_16-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 140.1 KiB | [orafce_16-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.7 KiB | [orafce_16-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.5 KiB | [orafce_16-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 139.5 KiB | [orafce_16-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/orafce_16-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_16` | `4.16.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.9 KiB | [orafce_16-4.16.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.16.5-1PGDG.rhel9.7.aarch64.rpm) |
| `orafce_16` | `4.16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.9 KiB | [orafce_16-4.16.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.16.2-2PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.6 KiB | [orafce_16-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_16-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.7 KiB | [orafce_16-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.7 KiB | [orafce_16-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.5 KiB | [orafce_16-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.4 KiB | [orafce_16-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.14.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.4 KiB | [orafce_16-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.13.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.5 KiB | [orafce_16-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.13.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.3 KiB | [orafce_16-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.13.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.2 KiB | [orafce_16-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 140.0 KiB | [orafce_16-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 139.5 KiB | [orafce_16-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.10.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 139.3 KiB | [orafce_16-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.10.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 139.1 KiB | [orafce_16-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 138.3 KiB | [orafce_16-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 137.6 KiB | [orafce_16-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 137.5 KiB | [orafce_16-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 137.3 KiB | [orafce_16-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 137.2 KiB | [orafce_16-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 137.1 KiB | [orafce_16-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/orafce_16-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_16` | `4.16.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_16-4.16.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.16.5-1PGDG.rhel10.1.x86_64.rpm) |
| `orafce_16` | `4.16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.7 KiB | [orafce_16-4.16.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.16.2-2PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | `4.16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_16-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | `4.14.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.8 KiB | [orafce_16-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | `4.14.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.6 KiB | [orafce_16-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | `4.14.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 149.6 KiB | [orafce_16-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/orafce_16-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_16` | `4.16.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.8 KiB | [orafce_16-4.16.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.16.5-1PGDG.rhel10.1.aarch64.rpm) |
| `orafce_16` | `4.16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.8 KiB | [orafce_16-4.16.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.16.2-2PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | `4.16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.0 KiB | [orafce_16-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | `4.14.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.2 KiB | [orafce_16-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | `4.14.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.0 KiB | [orafce_16-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_16` | `4.14.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.0 KiB | [orafce_16-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/orafce_16-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-orafce` | `4.16.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.2 KiB | [postgresql-16-orafce_4.16.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg12+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.2 KiB | [postgresql-16-orafce_4.16.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 362.7 KiB | [postgresql-16-orafce_4.16.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg12+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 354.5 KiB | [postgresql-16-orafce_4.16.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg12+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 354.5 KiB | [postgresql-16-orafce_4.16.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 355.3 KiB | [postgresql-16-orafce_4.16.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg12+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 362.9 KiB | [postgresql-16-orafce_4.16.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg13+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.0 KiB | [postgresql-16-orafce_4.16.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 363.6 KiB | [postgresql-16-orafce_4.16.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg13+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.3 KiB | [postgresql-16-orafce_4.16.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg13+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 355.9 KiB | [postgresql-16-orafce_4.16.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 356.5 KiB | [postgresql-16-orafce_4.16.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg13+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 397.4 KiB | [postgresql-16-orafce_4.16.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 397.2 KiB | [postgresql-16-orafce_4.16.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 397.0 KiB | [postgresql-16-orafce_4.16.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.1 KiB | [postgresql-16-orafce_4.16.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.3 KiB | [postgresql-16-orafce_4.16.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 389.4 KiB | [postgresql-16-orafce_4.16.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.0 KiB | [postgresql-16-orafce_4.16.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 359.9 KiB | [postgresql-16-orafce_4.16.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 360.8 KiB | [postgresql-16-orafce_4.16.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-orafce` | `4.16.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.4 KiB | [postgresql-16-orafce_4.16.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.5 KiB | [postgresql-16-orafce_4.16.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-orafce` | `4.16.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 354.9 KiB | [postgresql-16-orafce_4.16.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-16-orafce_4.16.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_15` | `4.16.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.2 KiB | [orafce_15-4.16.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.16.5-1PGDG.rhel8.10.x86_64.rpm) |
| `orafce_15` | `4.16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.7 KiB | [orafce_15-4.16.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.16.2-2PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.5 KiB | [orafce_15-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.9 KiB | [orafce_15-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.8 KiB | [orafce_15-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.5 KiB | [orafce_15-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.3 KiB | [orafce_15-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.14.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.8 KiB | [orafce_15-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.13.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.3 KiB | [orafce_15-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.13.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.0 KiB | [orafce_15-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.13.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.9 KiB | [orafce_15-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.6 KiB | [orafce_15-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.2 KiB | [orafce_15-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.9 KiB | [orafce_15-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.8 KiB | [orafce_15-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.4 KiB | [orafce_15-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.5 KiB | [orafce_15-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.4 KiB | [orafce_15-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.1 KiB | [orafce_15-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 145.0 KiB | [orafce_15-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 144.9 KiB | [orafce_15-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/orafce_15-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_15` | `4.16.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.8 KiB | [orafce_15-4.16.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.16.5-1PGDG.rhel8.10.aarch64.rpm) |
| `orafce_15` | `4.16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 148.3 KiB | [orafce_15-4.16.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.16.2-2PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.1 KiB | [orafce_15-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.7 KiB | [orafce_15-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.5 KiB | [orafce_15-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.2 KiB | [orafce_15-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.1 KiB | [orafce_15-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.14.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.9 KiB | [orafce_15-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.13.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.5 KiB | [orafce_15-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.13.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.2 KiB | [orafce_15-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.13.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 144.0 KiB | [orafce_15-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.7 KiB | [orafce_15-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.3 KiB | [orafce_15-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.1 KiB | [orafce_15-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.9 KiB | [orafce_15-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.5 KiB | [orafce_15-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.8 KiB | [orafce_15-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.7 KiB | [orafce_15-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.3 KiB | [orafce_15-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.1 KiB | [orafce_15-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.0 KiB | [orafce_15-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/orafce_15-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_15` | `4.16.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.2 KiB | [orafce_15-4.16.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.16.5-1PGDG.rhel9.7.x86_64.rpm) |
| `orafce_15` | `4.16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [orafce_15-4.16.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.16.2-2PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.0 KiB | [orafce_15-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.9 KiB | [orafce_15-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.6 KiB | [orafce_15-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.7 KiB | [orafce_15-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.5 KiB | [orafce_15-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.5 KiB | [orafce_15-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.14.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.6 KiB | [orafce_15-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.13.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.0 KiB | [orafce_15-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.13.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 148.1 KiB | [orafce_15-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.13.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.9 KiB | [orafce_15-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.5 KiB | [orafce_15-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.1 KiB | [orafce_15-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.10.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.1 KiB | [orafce_15-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.10.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.0 KiB | [orafce_15-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.10.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 145.9 KiB | [orafce_15-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 144.6 KiB | [orafce_15-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 144.6 KiB | [orafce_15-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 144.2 KiB | [orafce_15-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 144.2 KiB | [orafce_15-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 144.2 KiB | [orafce_15-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/orafce_15-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_15` | `4.16.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 148.0 KiB | [orafce_15-4.16.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.16.5-1PGDG.rhel9.7.aarch64.rpm) |
| `orafce_15` | `4.16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 148.0 KiB | [orafce_15-4.16.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.16.2-2PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.7 KiB | [orafce_15-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.8 KiB | [orafce_15-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.8 KiB | [orafce_15-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.8 KiB | [orafce_15-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_15-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.6 KiB | [orafce_15-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.14.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.0 KiB | [orafce_15-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.13.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.7 KiB | [orafce_15-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.13.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.5 KiB | [orafce_15-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.13.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.3 KiB | [orafce_15-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.8 KiB | [orafce_15-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.4 KiB | [orafce_15-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.10.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.7 KiB | [orafce_15-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.10.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.6 KiB | [orafce_15-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.7 KiB | [orafce_15-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.2 KiB | [orafce_15-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.2 KiB | [orafce_15-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.8 KiB | [orafce_15-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.5 KiB | [orafce_15-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 141.5 KiB | [orafce_15-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/orafce_15-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_15` | `4.16.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.7 KiB | [orafce_15-4.16.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.16.5-1PGDG.rhel10.1.x86_64.rpm) |
| `orafce_15` | `4.16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.8 KiB | [orafce_15-4.16.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.16.2-2PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | `4.16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.9 KiB | [orafce_15-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | `4.14.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.2 KiB | [orafce_15-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | `4.14.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.1 KiB | [orafce_15-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | `4.14.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 150.0 KiB | [orafce_15-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/orafce_15-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_15` | `4.16.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.3 KiB | [orafce_15-4.16.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.16.5-1PGDG.rhel10.1.aarch64.rpm) |
| `orafce_15` | `4.16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.3 KiB | [orafce_15-4.16.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.16.2-2PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | `4.16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.4 KiB | [orafce_15-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | `4.14.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.5 KiB | [orafce_15-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | `4.14.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.3 KiB | [orafce_15-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_15` | `4.14.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 148.3 KiB | [orafce_15-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/orafce_15-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-orafce` | `4.16.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 364.7 KiB | [postgresql-15-orafce_4.16.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg12+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 364.4 KiB | [postgresql-15-orafce_4.16.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 364.9 KiB | [postgresql-15-orafce_4.16.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg12+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 357.2 KiB | [postgresql-15-orafce_4.16.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg12+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 357.1 KiB | [postgresql-15-orafce_4.16.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 357.6 KiB | [postgresql-15-orafce_4.16.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg12+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 365.7 KiB | [postgresql-15-orafce_4.16.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg13+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 366.1 KiB | [postgresql-15-orafce_4.16.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 366.1 KiB | [postgresql-15-orafce_4.16.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg13+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 358.8 KiB | [postgresql-15-orafce_4.16.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg13+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 358.6 KiB | [postgresql-15-orafce_4.16.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 358.9 KiB | [postgresql-15-orafce_4.16.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg13+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 404.0 KiB | [postgresql-15-orafce_4.16.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 404.1 KiB | [postgresql-15-orafce_4.16.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 403.9 KiB | [postgresql-15-orafce_4.16.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 395.6 KiB | [postgresql-15-orafce_4.16.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 395.3 KiB | [postgresql-15-orafce_4.16.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 395.5 KiB | [postgresql-15-orafce_4.16.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 364.4 KiB | [postgresql-15-orafce_4.16.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 364.1 KiB | [postgresql-15-orafce_4.16.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 364.6 KiB | [postgresql-15-orafce_4.16.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-orafce` | `4.16.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 358.5 KiB | [postgresql-15-orafce_4.16.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 358.2 KiB | [postgresql-15-orafce_4.16.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-orafce` | `4.16.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 358.5 KiB | [postgresql-15-orafce_4.16.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-15-orafce_4.16.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orafce_14` | `4.16.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 154.3 KiB | [orafce_14-4.16.5-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.16.5-1PGDG.rhel8.10.x86_64.rpm) |
| `orafce_14` | `4.16.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 153.8 KiB | [orafce_14-4.16.2-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.16.2-2PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.5 KiB | [orafce_14-4.14.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.6-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 152.0 KiB | [orafce_14-4.14.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.8 KiB | [orafce_14-4.14.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.3-2PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.5 KiB | [orafce_14-4.14.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 151.4 KiB | [orafce_14-4.14.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.14.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.8 KiB | [orafce_14-4.14.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.14.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.13.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.4 KiB | [orafce_14-4.13.5-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.5-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.13.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 150.1 KiB | [orafce_14-4.13.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.13.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 149.9 KiB | [orafce_14-4.13.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.13.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.12.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.6 KiB | [orafce_14-4.12.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.12.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.11.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.3 KiB | [orafce_14-4.11.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.11.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.10.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 148.0 KiB | [orafce_14-4.10.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.10.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.6 KiB | [orafce_14-4.10.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.10.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 147.3 KiB | [orafce_14-4.10.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.10.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.9.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.7 KiB | [orafce_14-4.9.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.4-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.9.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.5 KiB | [orafce_14-4.9.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.3-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.9.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.3 KiB | [orafce_14-4.9.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.2-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.2 KiB | [orafce_14-4.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.1-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.9.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 146.0 KiB | [orafce_14-4.9.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/orafce_14-4.9.0-1PGDG.rhel8.x86_64.rpm) |
| `orafce_14` | `4.16.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 149.6 KiB | [orafce_14-4.16.5-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.16.5-1PGDG.rhel8.10.aarch64.rpm) |
| `orafce_14` | `4.16.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 149.2 KiB | [orafce_14-4.16.2-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.16.2-2PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.9 KiB | [orafce_14-4.14.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.6-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.5 KiB | [orafce_14-4.14.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.3 KiB | [orafce_14-4.14.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.3-2PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 147.0 KiB | [orafce_14-4.14.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 146.9 KiB | [orafce_14-4.14.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.14.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.8 KiB | [orafce_14-4.14.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.14.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.13.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.4 KiB | [orafce_14-4.13.5-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.5-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.13.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.1 KiB | [orafce_14-4.13.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.13.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 145.0 KiB | [orafce_14-4.13.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.13.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.12.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.6 KiB | [orafce_14-4.12.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.12.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.11.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.2 KiB | [orafce_14-4.11.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.11.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.10.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 143.0 KiB | [orafce_14-4.10.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.10.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.8 KiB | [orafce_14-4.10.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.10.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 142.3 KiB | [orafce_14-4.10.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.10.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.9.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.7 KiB | [orafce_14-4.9.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.4-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.9.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.5 KiB | [orafce_14-4.9.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.3-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.9.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 141.2 KiB | [orafce_14-4.9.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.2-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.9 KiB | [orafce_14-4.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.1-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.9.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 140.8 KiB | [orafce_14-4.9.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/orafce_14-4.9.0-1PGDG.rhel8.aarch64.rpm) |
| `orafce_14` | `4.16.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 151.3 KiB | [orafce_14-4.16.5-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.16.5-1PGDG.rhel9.7.x86_64.rpm) |
| `orafce_14` | `4.16.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 151.4 KiB | [orafce_14-4.16.2-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.16.2-2PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.16.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 151.4 KiB | [orafce_14-4.16.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.16.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 150.1 KiB | [orafce_14-4.14.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.6-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.9 KiB | [orafce_14-4.14.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.9 KiB | [orafce_14-4.14.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.3-2PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.8 KiB | [orafce_14-4.14.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.7 KiB | [orafce_14-4.14.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.14.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.7 KiB | [orafce_14-4.14.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.14.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.13.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.5 KiB | [orafce_14-4.13.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.5-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.13.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.3 KiB | [orafce_14-4.13.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.13.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 149.1 KiB | [orafce_14-4.13.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.13.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.12.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.8 KiB | [orafce_14-4.12.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.12.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.11.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.4 KiB | [orafce_14-4.11.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.11.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.10.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.7 KiB | [orafce_14-4.10.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.10.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.6 KiB | [orafce_14-4.10.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.10.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 147.1 KiB | [orafce_14-4.10.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.10.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.9.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 146.0 KiB | [orafce_14-4.9.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.4-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.9.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 145.9 KiB | [orafce_14-4.9.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.3-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.9.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 145.6 KiB | [orafce_14-4.9.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.2-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 145.5 KiB | [orafce_14-4.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.1-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.9.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 145.4 KiB | [orafce_14-4.9.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/orafce_14-4.9.0-1PGDG.rhel9.x86_64.rpm) |
| `orafce_14` | `4.16.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 149.0 KiB | [orafce_14-4.16.5-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.16.5-1PGDG.rhel9.7.aarch64.rpm) |
| `orafce_14` | `4.16.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 149.0 KiB | [orafce_14-4.16.2-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.16.2-2PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.16.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 148.7 KiB | [orafce_14-4.16.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.16.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.7 KiB | [orafce_14-4.14.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.6-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.7 KiB | [orafce_14-4.14.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.7 KiB | [orafce_14-4.14.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.3-2PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.5 KiB | [orafce_14-4.14.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 147.5 KiB | [orafce_14-4.14.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.14.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.8 KiB | [orafce_14-4.14.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.14.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.13.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.5 KiB | [orafce_14-4.13.5-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.5-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.13.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.2 KiB | [orafce_14-4.13.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.13.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 146.1 KiB | [orafce_14-4.13.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.13.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.12.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 145.1 KiB | [orafce_14-4.12.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.12.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.11.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 144.5 KiB | [orafce_14-4.11.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.11.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.10.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 144.6 KiB | [orafce_14-4.10.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.10.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 144.4 KiB | [orafce_14-4.10.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.10.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 143.6 KiB | [orafce_14-4.10.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.10.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.9.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.8 KiB | [orafce_14-4.9.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.4-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.9.3` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.8 KiB | [orafce_14-4.9.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.3-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.9.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.5 KiB | [orafce_14-4.9.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.2-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [orafce_14-4.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.1-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.9.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 142.4 KiB | [orafce_14-4.9.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/orafce_14-4.9.0-1PGDG.rhel9.aarch64.rpm) |
| `orafce_14` | `4.16.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 152.2 KiB | [orafce_14-4.16.5-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.16.5-1PGDG.rhel10.1.x86_64.rpm) |
| `orafce_14` | `4.16.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 152.2 KiB | [orafce_14-4.16.2-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.16.2-2PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | `4.16.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 152.3 KiB | [orafce_14-4.16.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.16.1-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | `4.14.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.4 KiB | [orafce_14-4.14.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.6-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | `4.14.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.2 KiB | [orafce_14-4.14.4-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.4-1PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | `4.14.3` | [el10.x86_64](/os/el10.x86_64) | pgdg | 151.1 KiB | [orafce_14-4.14.3-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/orafce_14-4.14.3-2PGDG.rhel10.x86_64.rpm) |
| `orafce_14` | `4.16.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 150.2 KiB | [orafce_14-4.16.5-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.16.5-1PGDG.rhel10.1.aarch64.rpm) |
| `orafce_14` | `4.16.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 150.2 KiB | [orafce_14-4.16.2-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.16.2-2PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | `4.16.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 150.3 KiB | [orafce_14-4.16.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.16.1-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | `4.14.6` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.6 KiB | [orafce_14-4.14.6-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.6-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | `4.14.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.4 KiB | [orafce_14-4.14.4-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.4-1PGDG.rhel10.aarch64.rpm) |
| `orafce_14` | `4.14.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 149.3 KiB | [orafce_14-4.14.3-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/orafce_14-4.14.3-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-orafce` | `4.16.5` | [d12.x86_64](/os/d12.x86_64) | pgdg | 367.7 KiB | [postgresql-14-orafce_4.16.5-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg12+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 368.0 KiB | [postgresql-14-orafce_4.16.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 367.9 KiB | [postgresql-14-orafce_4.16.3-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg12+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [d12.aarch64](/os/d12.aarch64) | pgdg | 360.3 KiB | [postgresql-14-orafce_4.16.5-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg12+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 360.1 KiB | [postgresql-14-orafce_4.16.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 360.8 KiB | [postgresql-14-orafce_4.16.3-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg12+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [d13.x86_64](/os/d13.x86_64) | pgdg | 368.8 KiB | [postgresql-14-orafce_4.16.5-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg13+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 369.7 KiB | [postgresql-14-orafce_4.16.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [d13.x86_64](/os/d13.x86_64) | pgdg | 369.0 KiB | [postgresql-14-orafce_4.16.3-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg13+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [d13.aarch64](/os/d13.aarch64) | pgdg | 361.6 KiB | [postgresql-14-orafce_4.16.5-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg13+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 361.5 KiB | [postgresql-14-orafce_4.16.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [d13.aarch64](/os/d13.aarch64) | pgdg | 362.1 KiB | [postgresql-14-orafce_4.16.3-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg13+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [u22.x86_64](/os/u22.x86_64) | pgdg | 404.0 KiB | [postgresql-14-orafce_4.16.5-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 404.1 KiB | [postgresql-14-orafce_4.16.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 404.0 KiB | [postgresql-14-orafce_4.16.3-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [u22.aarch64](/os/u22.aarch64) | pgdg | 395.3 KiB | [postgresql-14-orafce_4.16.5-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 394.9 KiB | [postgresql-14-orafce_4.16.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 395.5 KiB | [postgresql-14-orafce_4.16.3-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [u24.x86_64](/os/u24.x86_64) | pgdg | 367.3 KiB | [postgresql-14-orafce_4.16.5-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 367.2 KiB | [postgresql-14-orafce_4.16.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [u24.x86_64](/os/u24.x86_64) | pgdg | 368.0 KiB | [postgresql-14-orafce_4.16.3-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-orafce` | `4.16.5` | [u24.aarch64](/os/u24.aarch64) | pgdg | 361.0 KiB | [postgresql-14-orafce_4.16.5-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.5-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 361.3 KiB | [postgresql-14-orafce_4.16.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.4-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-orafce` | `4.16.3` | [u24.aarch64](/os/u24.aarch64) | pgdg | 361.4 KiB | [postgresql-14-orafce_4.16.3-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/o/orafce/postgresql-14-orafce_4.16.3-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/orafce/orafce" title="Repository" icon="github" subtitle="github.com/orafce/orafce" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install orafce;		# install via package name, for the active PG version

pig install orafce -v 18;   # install for PG 18
pig install orafce -v 17;   # install for PG 17
pig install orafce -v 16;   # install for PG 16
pig install orafce -v 15;   # install for PG 15
pig install orafce -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION orafce;
```



## Usage

> [orafce: Functions and operators that emulate a subset of functions and packages from the Oracle RDBMS](https://github.com/orafce/orafce)

### Date Functions

```sql
SELECT add_months(date '2005-05-31', 1);        -- 2005-06-30
SELECT last_day(date '2005-05-24');              -- 2005-05-31
SELECT next_day(date '2005-05-24', 'monday');    -- 2005-05-30
SELECT months_between(date '1995-02-02', date '1995-01-01'); -- 1.032...
SELECT trunc(date '2005-07-12', 'iw');           -- 2005-07-11
SELECT round(date '2005-07-12', 'yyyy');         -- 2006-01-01
```

### Oracle DATE Data Type

```sql
SET search_path TO oracle, "$user", public, pg_catalog;
CREATE TABLE t (col1 date);
INSERT INTO t VALUES('2014-06-24 12:12:11'::date);  -- includes time component
```

### String Functions (NVL, DECODE, etc.)

```sql
SELECT nvl('A', 'B');            -- A
SELECT nvl(NULL, 'B');           -- B
SELECT decode(1, 1, 'one', 2, 'two', 'other');  -- one
SELECT lnnvl(true);              -- false
SELECT nanvl(0.0/0.0, 999);     -- 999
```

### DUAL Table

```sql
SELECT * FROM dual;
```

### Package DBMS_OUTPUT

```sql
SELECT dbms_output.enable();
SELECT dbms_output.put_line('Hello');
SELECT dbms_output.get_line(line, status);  -- retrieves output
```

### Package DBMS_PIPE

```sql
SELECT dbms_pipe.create_pipe('my_pipe');
SELECT dbms_pipe.pack_message('message text');
SELECT dbms_pipe.send_message('my_pipe');
-- In another session:
SELECT dbms_pipe.receive_message('my_pipe');
SELECT dbms_pipe.unpack_message_text();
```

### Package DBMS_ALERT

```sql
CALL dbms_alert.register('my_alert');
-- In another session:
CALL dbms_alert.signal('my_alert', 'Alert message');
-- Back in first session:
CALL dbms_alert.waitone('my_alert', name, message, status, 60);
```

### Package DBMS_UTILITY

```sql
SELECT dbms_utility.format_call_stack();
```

### Package UTL_FILE

```sql
CALL utl_file.fopen('/tmp', 'test.txt', 'w');
CALL utl_file.put_line(f, 'Hello World');
CALL utl_file.fclose(f);
```

### Package PLVstr / PLVchr

```sql
SELECT plvstr.left('Hello World', 5);     -- Hello
SELECT plvstr.right('Hello World', 5);    -- World
SELECT plvstr.rvrs('Hello');              -- olleH
SELECT plvchr.nth('Hello', 3);            -- l
SELECT plvchr.first('Hello');             -- H
SELECT plvchr.last('Hello');              -- o
```

### Package PLVsubst

```sql
SELECT plvsubst.string('My name is %s %s.', ARRAY['Pavel','Stehule']);
-- My name is Pavel Stehule.
```

### DBMS_ASSERT (SQL Injection Protection)

```sql
SELECT dbms_assert.enquote_literal('some value');
SELECT dbms_assert.schema_name('public');
SELECT dbms_assert.object_name('my_table');
```

### VARCHAR2 and NVARCHAR2 Types

The extension provides Oracle-compatible `varchar2` and `nvarchar2` data types that enforce the declared length in bytes (varchar2) or characters (nvarchar2).
