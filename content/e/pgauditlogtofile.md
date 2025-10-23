---
title: "pgauditlogtofile"
linkTitle: "pgauditlogtofile"
description: "pgAudit addon to redirect audit log to an independent file"
weight: 7090
categories: ["SEC"]
width: full
---

pgAudit addon to redirect audit log to an independent file


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7090** | {{< badge content="pgauditlogtofile" link="https://github.com/fmbiete/pgauditlogtofile" >}} | {{< ext "pgauditlogtofile" >}} | `1.7.1` | {{< category "SEC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgaudit" >}} {{< ext "pg_auth_mon" >}} {{< ext "logerrors" >}} {{< ext "pg_permissions" >}} {{< ext "login_hook" >}} {{< ext "set_user" >}} {{< ext "pg_drop_events" >}} {{< ext "table_log" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgauditlogtofile" >}} | `1.7.1` | {{< bg "18" "pgauditlogtofile_18*" "red" >}} {{< bg "17" "pgauditlogtofile_17*" "green" >}} {{< bg "16" "pgauditlogtofile_16*" "green" >}} {{< bg "15" "pgauditlogtofile_15*" "green" >}} {{< bg "14" "pgauditlogtofile_14*" "green" >}} | `pgauditlogtofile_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgauditlogtofile" >}} | `1.7.1` | {{< bg "18" "postgresql-18-pgauditlogtofile" "red" >}} {{< bg "17" "postgresql-17-pgauditlogtofile" "green" >}} {{< bg "16" "postgresql-16-pgauditlogtofile" "green" >}} {{< bg "15" "postgresql-15-pgauditlogtofile" "green" >}} {{< bg "14" "postgresql-14-pgauditlogtofile" "green" >}} | `postgresql-$v-pgauditlogtofile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.7.4" "pgauditlogtofile_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_14 : AVAIL 13" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.7.4" "pgauditlogtofile_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_15 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.7.4" "pgauditlogtofile_14 : AVAIL 9" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.7.3" "pgauditlogtofile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_14 : AVAIL 11" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.7.3" "pgauditlogtofile_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_17 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.3" "pgauditlogtofile_14 : AVAIL 9" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.7.4" "postgresql-18-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-17-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-16-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-15-pgauditlogtofile : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.7.4" "postgresql-14-pgauditlogtofile : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgauditlogtofile_18` | 1.7.4 | `el8.x86_64` | pgdg | 26.6 KiB | [pgauditlogtofile_18-1.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgauditlogtofile_18-1.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_18` | 1.7.4 | `el8.aarch64` | pgdg | 26.2 KiB | [pgauditlogtofile_18-1.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgauditlogtofile_18-1.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_18` | 1.7.3 | `el9.x86_64` | pgdg | 25.0 KiB | [pgauditlogtofile_18-1.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgauditlogtofile_18-1.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_18` | 1.7.3 | `el9.x86_64` | pgdg | 25.1 KiB | [pgauditlogtofile_18-1.7.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgauditlogtofile_18-1.7.3-2PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_18` | 1.7.3 | `el9.aarch64` | pgdg | 24.8 KiB | [pgauditlogtofile_18-1.7.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgauditlogtofile_18-1.7.3-2PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_18` | 1.7.3 | `el9.aarch64` | pgdg | 24.7 KiB | [pgauditlogtofile_18-1.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgauditlogtofile_18-1.7.3-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `d12.x86_64` | pgdg | 49.8 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `d12.aarch64` | pgdg | 49.0 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `u22.x86_64` | pgdg | 50.5 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `u22.aarch64` | pgdg | 49.6 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `u24.x86_64` | pgdg | 49.3 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pgauditlogtofile` | 1.7.4 | `u24.aarch64` | pgdg | 48.5 KiB | [postgresql-18-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-18-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgauditlogtofile_17` | 1.7.4 | `el8.x86_64` | pgdg | 26.5 KiB | [pgauditlogtofile_17-1.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgauditlogtofile_17-1.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.4 | `el8.x86_64` | pgdg | 23.5 KiB | [pgauditlogtofile_17-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgauditlogtofile_17-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.3 | `el8.x86_64` | pgdg | 23.3 KiB | [pgauditlogtofile_17-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgauditlogtofile_17-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.2 | `el8.x86_64` | pgdg | 23.2 KiB | [pgauditlogtofile_17-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgauditlogtofile_17-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.7.4 | `el8.aarch64` | pgdg | 26.2 KiB | [pgauditlogtofile_17-1.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgauditlogtofile_17-1.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.4 | `el8.aarch64` | pgdg | 23.2 KiB | [pgauditlogtofile_17-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgauditlogtofile_17-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.3 | `el8.aarch64` | pgdg | 23.1 KiB | [pgauditlogtofile_17-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgauditlogtofile_17-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.2 | `el8.aarch64` | pgdg | 22.9 KiB | [pgauditlogtofile_17-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgauditlogtofile_17-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.7.3 | `el9.x86_64` | pgdg | 25.1 KiB | [pgauditlogtofile_17-1.7.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.7.3-2PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.7.3 | `el9.x86_64` | pgdg | 25.0 KiB | [pgauditlogtofile_17-1.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.7.1 | `el9.x86_64` | pgdg | 24.6 KiB | [pgauditlogtofile_17-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.7.1 | `el9.x86_64` | pgdg | 24.4 KiB | [pgauditlogtofile_17-1.7.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.7.1-2PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.4 | `el9.x86_64` | pgdg | 22.5 KiB | [pgauditlogtofile_17-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.3 | `el9.x86_64` | pgdg | 22.5 KiB | [pgauditlogtofile_17-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.6.2 | `el9.x86_64` | pgdg | 22.5 KiB | [pgauditlogtofile_17-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgauditlogtofile_17-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_17` | 1.7.3 | `el9.aarch64` | pgdg | 24.8 KiB | [pgauditlogtofile_17-1.7.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.7.3-2PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.7.3 | `el9.aarch64` | pgdg | 24.7 KiB | [pgauditlogtofile_17-1.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.7.1 | `el9.aarch64` | pgdg | 24.3 KiB | [pgauditlogtofile_17-1.7.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.7.1-2PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.7.1 | `el9.aarch64` | pgdg | 24.2 KiB | [pgauditlogtofile_17-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.4 | `el9.aarch64` | pgdg | 22.3 KiB | [pgauditlogtofile_17-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.3 | `el9.aarch64` | pgdg | 22.3 KiB | [pgauditlogtofile_17-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_17` | 1.6.2 | `el9.aarch64` | pgdg | 22.3 KiB | [pgauditlogtofile_17-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgauditlogtofile_17-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `d12.x86_64` | pgdg | 49.6 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `d12.aarch64` | pgdg | 48.8 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `u22.x86_64` | pgdg | 53.9 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `u22.aarch64` | pgdg | 53.0 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `u24.x86_64` | pgdg | 49.2 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pgauditlogtofile` | 1.7.4 | `u24.aarch64` | pgdg | 48.4 KiB | [postgresql-17-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-17-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgauditlogtofile_16` | 1.7.4 | `el8.x86_64` | pgdg | 26.4 KiB | [pgauditlogtofile_16-1.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.4 | `el8.x86_64` | pgdg | 23.4 KiB | [pgauditlogtofile_16-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.3 | `el8.x86_64` | pgdg | 23.2 KiB | [pgauditlogtofile_16-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.2 | `el8.x86_64` | pgdg | 23.1 KiB | [pgauditlogtofile_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.0 | `el8.x86_64` | pgdg | 22.4 KiB | [pgauditlogtofile_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.5.12 | `el8.x86_64` | pgdg | 19.4 KiB | [pgauditlogtofile_16-1.5.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgauditlogtofile_16-1.5.12-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.7.4 | `el8.aarch64` | pgdg | 26.0 KiB | [pgauditlogtofile_16-1.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.4 | `el8.aarch64` | pgdg | 23.1 KiB | [pgauditlogtofile_16-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.3 | `el8.aarch64` | pgdg | 23.0 KiB | [pgauditlogtofile_16-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.2 | `el8.aarch64` | pgdg | 22.8 KiB | [pgauditlogtofile_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.0 | `el8.aarch64` | pgdg | 22.1 KiB | [pgauditlogtofile_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.5.12 | `el8.aarch64` | pgdg | 19.3 KiB | [pgauditlogtofile_16-1.5.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgauditlogtofile_16-1.5.12-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.7.3 | `el9.x86_64` | pgdg | 24.6 KiB | [pgauditlogtofile_16-1.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.7.3 | `el9.x86_64` | pgdg | 24.9 KiB | [pgauditlogtofile_16-1.7.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.7.3-2PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.7.1 | `el9.x86_64` | pgdg | 24.3 KiB | [pgauditlogtofile_16-1.7.1-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.7.1-2PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.7.1 | `el9.x86_64` | pgdg | 24.4 KiB | [pgauditlogtofile_16-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.4 | `el9.x86_64` | pgdg | 22.4 KiB | [pgauditlogtofile_16-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.3 | `el9.x86_64` | pgdg | 22.4 KiB | [pgauditlogtofile_16-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.2 | `el9.x86_64` | pgdg | 22.4 KiB | [pgauditlogtofile_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.6.0 | `el9.x86_64` | pgdg | 21.7 KiB | [pgauditlogtofile_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.5.12 | `el9.x86_64` | pgdg | 19.1 KiB | [pgauditlogtofile_16-1.5.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgauditlogtofile_16-1.5.12-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_16` | 1.7.3 | `el9.aarch64` | pgdg | 24.6 KiB | [pgauditlogtofile_16-1.7.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.7.3-2PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.7.3 | `el9.aarch64` | pgdg | 24.3 KiB | [pgauditlogtofile_16-1.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.7.1 | `el9.aarch64` | pgdg | 24.0 KiB | [pgauditlogtofile_16-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.7.1 | `el9.aarch64` | pgdg | 24.1 KiB | [pgauditlogtofile_16-1.7.1-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.7.1-2PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.4 | `el9.aarch64` | pgdg | 22.1 KiB | [pgauditlogtofile_16-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.3 | `el9.aarch64` | pgdg | 22.2 KiB | [pgauditlogtofile_16-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.2 | `el9.aarch64` | pgdg | 22.2 KiB | [pgauditlogtofile_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.6.0 | `el9.aarch64` | pgdg | 21.5 KiB | [pgauditlogtofile_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_16` | 1.5.12 | `el9.aarch64` | pgdg | 18.8 KiB | [pgauditlogtofile_16-1.5.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgauditlogtofile_16-1.5.12-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `d12.x86_64` | pgdg | 49.6 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `d12.aarch64` | pgdg | 48.8 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `u22.x86_64` | pgdg | 54.0 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `u22.aarch64` | pgdg | 53.0 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `u24.x86_64` | pgdg | 49.3 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pgauditlogtofile` | 1.7.4 | `u24.aarch64` | pgdg | 48.3 KiB | [postgresql-16-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-16-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgauditlogtofile_15` | 1.7.4 | `el8.x86_64` | pgdg | 27.2 KiB | [pgauditlogtofile_15-1.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.4 | `el8.x86_64` | pgdg | 24.3 KiB | [pgauditlogtofile_15-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.3 | `el8.x86_64` | pgdg | 24.2 KiB | [pgauditlogtofile_15-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.2 | `el8.x86_64` | pgdg | 24.0 KiB | [pgauditlogtofile_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.0 | `el8.x86_64` | pgdg | 23.2 KiB | [pgauditlogtofile_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.6 | `el8.x86_64` | pgdg | 17.9 KiB | [pgauditlogtofile_15-1.5.6-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.5.6-2.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.12 | `el8.x86_64` | pgdg | 19.8 KiB | [pgauditlogtofile_15-1.5.12-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.5.12-2PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.10 | `el8.x86_64` | pgdg | 19.3 KiB | [pgauditlogtofile_15-1.5.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgauditlogtofile_15-1.5.10-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.7.4 | `el8.aarch64` | pgdg | 26.7 KiB | [pgauditlogtofile_15-1.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.4 | `el8.aarch64` | pgdg | 24.0 KiB | [pgauditlogtofile_15-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.3 | `el8.aarch64` | pgdg | 23.8 KiB | [pgauditlogtofile_15-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.2 | `el8.aarch64` | pgdg | 23.6 KiB | [pgauditlogtofile_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.0 | `el8.aarch64` | pgdg | 22.9 KiB | [pgauditlogtofile_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.6 | `el8.aarch64` | pgdg | 17.9 KiB | [pgauditlogtofile_15-1.5.6-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.5.6-2.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.12 | `el8.aarch64` | pgdg | 19.7 KiB | [pgauditlogtofile_15-1.5.12-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.5.12-2PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.10 | `el8.aarch64` | pgdg | 19.2 KiB | [pgauditlogtofile_15-1.5.10-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgauditlogtofile_15-1.5.10-1.rhel8.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.7.3 | `el9.x86_64` | pgdg | 26.3 KiB | [pgauditlogtofile_15-1.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.7.1 | `el9.x86_64` | pgdg | 25.9 KiB | [pgauditlogtofile_15-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.4 | `el9.x86_64` | pgdg | 24.0 KiB | [pgauditlogtofile_15-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.3 | `el9.x86_64` | pgdg | 23.8 KiB | [pgauditlogtofile_15-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.2 | `el9.x86_64` | pgdg | 24.0 KiB | [pgauditlogtofile_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.6.0 | `el9.x86_64` | pgdg | 23.3 KiB | [pgauditlogtofile_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.6 | `el9.x86_64` | pgdg | 18.3 KiB | [pgauditlogtofile_15-1.5.6-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.5.6-2.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.12 | `el9.x86_64` | pgdg | 19.4 KiB | [pgauditlogtofile_15-1.5.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.5.12-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.5.10 | `el9.x86_64` | pgdg | 19.1 KiB | [pgauditlogtofile_15-1.5.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgauditlogtofile_15-1.5.10-1.rhel9.x86_64.rpm) |
| `pgauditlogtofile_15` | 1.7.3 | `el9.aarch64` | pgdg | 25.9 KiB | [pgauditlogtofile_15-1.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.7.1 | `el9.aarch64` | pgdg | 25.5 KiB | [pgauditlogtofile_15-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.4 | `el9.aarch64` | pgdg | 23.8 KiB | [pgauditlogtofile_15-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.3 | `el9.aarch64` | pgdg | 23.8 KiB | [pgauditlogtofile_15-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.2 | `el9.aarch64` | pgdg | 23.8 KiB | [pgauditlogtofile_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.6.0 | `el9.aarch64` | pgdg | 23.1 KiB | [pgauditlogtofile_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.6 | `el9.aarch64` | pgdg | 18.0 KiB | [pgauditlogtofile_15-1.5.6-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.5.6-2.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.12 | `el9.aarch64` | pgdg | 19.0 KiB | [pgauditlogtofile_15-1.5.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.5.12-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_15` | 1.5.10 | `el9.aarch64` | pgdg | 18.8 KiB | [pgauditlogtofile_15-1.5.10-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgauditlogtofile_15-1.5.10-1.rhel9.aarch64.rpm) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `d12.x86_64` | pgdg | 50.4 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `d12.aarch64` | pgdg | 49.5 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `u22.x86_64` | pgdg | 55.3 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `u22.aarch64` | pgdg | 54.3 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `u24.x86_64` | pgdg | 50.4 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pgauditlogtofile` | 1.7.4 | `u24.aarch64` | pgdg | 49.5 KiB | [postgresql-15-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-15-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgauditlogtofile_14` | 1.7.4 | `el8.x86_64` | pgdg | 27.2 KiB | [pgauditlogtofile_14-1.7.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.7.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.4 | `el8.x86_64` | pgdg | 24.3 KiB | [pgauditlogtofile_14-1.6.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.6.4-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.3 | `el8.x86_64` | pgdg | 24.2 KiB | [pgauditlogtofile_14-1.6.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.6.3-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.2 | `el8.x86_64` | pgdg | 24.0 KiB | [pgauditlogtofile_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.0 | `el8.x86_64` | pgdg | 23.2 KiB | [pgauditlogtofile_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.6 | `el8.x86_64` | pgdg | 17.9 KiB | [pgauditlogtofile_14-1.5.6-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.6-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.5 | `el8.x86_64` | pgdg | 32.9 KiB | [pgauditlogtofile_14-1.5.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.5-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el8.x86_64` | pgdg | 19.7 KiB | [pgauditlogtofile_14-1.5.12-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.12-2PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el8.x86_64` | pgdg | 19.6 KiB | [pgauditlogtofile_14-1.5.12-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.12-1PGDG.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.10 | `el8.x86_64` | pgdg | 19.2 KiB | [pgauditlogtofile_14-1.5.10-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.10-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.1 | `el8.x86_64` | pgdg | 32.3 KiB | [pgauditlogtofile_14-1.5.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.5.1-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.4 | `el8.x86_64` | pgdg | 31.5 KiB | [pgauditlogtofile_14-1.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.4-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.3 | `el8.x86_64` | pgdg | 32.8 KiB | [pgauditlogtofile_14-1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgauditlogtofile_14-1.3-1.rhel8.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.7.4 | `el8.aarch64` | pgdg | 26.6 KiB | [pgauditlogtofile_14-1.7.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.7.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.4 | `el8.aarch64` | pgdg | 23.9 KiB | [pgauditlogtofile_14-1.6.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.6.4-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.3 | `el8.aarch64` | pgdg | 23.7 KiB | [pgauditlogtofile_14-1.6.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.6.3-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.2 | `el8.aarch64` | pgdg | 23.5 KiB | [pgauditlogtofile_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.0 | `el8.aarch64` | pgdg | 22.8 KiB | [pgauditlogtofile_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.6 | `el8.aarch64` | pgdg | 17.8 KiB | [pgauditlogtofile_14-1.5.6-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.5.6-1.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el8.aarch64` | pgdg | 19.6 KiB | [pgauditlogtofile_14-1.5.12-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.5.12-2PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el8.aarch64` | pgdg | 19.5 KiB | [pgauditlogtofile_14-1.5.12-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.5.12-1PGDG.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.10 | `el8.aarch64` | pgdg | 19.1 KiB | [pgauditlogtofile_14-1.5.10-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgauditlogtofile_14-1.5.10-1.rhel8.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.7.3 | `el9.x86_64` | pgdg | 26.2 KiB | [pgauditlogtofile_14-1.7.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.7.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.7.1 | `el9.x86_64` | pgdg | 25.8 KiB | [pgauditlogtofile_14-1.7.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.7.1-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.4 | `el9.x86_64` | pgdg | 23.8 KiB | [pgauditlogtofile_14-1.6.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.6.4-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.3 | `el9.x86_64` | pgdg | 23.9 KiB | [pgauditlogtofile_14-1.6.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.6.3-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.2 | `el9.x86_64` | pgdg | 24.0 KiB | [pgauditlogtofile_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.6.0 | `el9.x86_64` | pgdg | 23.2 KiB | [pgauditlogtofile_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.6 | `el9.x86_64` | pgdg | 18.2 KiB | [pgauditlogtofile_14-1.5.6-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.5.6-1.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.5 | `el9.x86_64` | pgdg | 33.8 KiB | [pgauditlogtofile_14-1.5.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.5.5-1.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el9.x86_64` | pgdg | 19.3 KiB | [pgauditlogtofile_14-1.5.12-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.5.12-1PGDG.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.10 | `el9.x86_64` | pgdg | 19.0 KiB | [pgauditlogtofile_14-1.5.10-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.5.10-1.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.5.1 | `el9.x86_64` | pgdg | 33.3 KiB | [pgauditlogtofile_14-1.5.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgauditlogtofile_14-1.5.1-1.rhel9.x86_64.rpm) |
| `pgauditlogtofile_14` | 1.7.3 | `el9.aarch64` | pgdg | 26.0 KiB | [pgauditlogtofile_14-1.7.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.7.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.7.1 | `el9.aarch64` | pgdg | 25.5 KiB | [pgauditlogtofile_14-1.7.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.7.1-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.4 | `el9.aarch64` | pgdg | 23.6 KiB | [pgauditlogtofile_14-1.6.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.6.4-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.3 | `el9.aarch64` | pgdg | 23.7 KiB | [pgauditlogtofile_14-1.6.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.6.3-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.2 | `el9.aarch64` | pgdg | 23.7 KiB | [pgauditlogtofile_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.6.0 | `el9.aarch64` | pgdg | 23.0 KiB | [pgauditlogtofile_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.6 | `el9.aarch64` | pgdg | 17.9 KiB | [pgauditlogtofile_14-1.5.6-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.5.6-1.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.12 | `el9.aarch64` | pgdg | 19.0 KiB | [pgauditlogtofile_14-1.5.12-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.5.12-1PGDG.rhel9.aarch64.rpm) |
| `pgauditlogtofile_14` | 1.5.10 | `el9.aarch64` | pgdg | 18.6 KiB | [pgauditlogtofile_14-1.5.10-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgauditlogtofile_14-1.5.10-1.rhel9.aarch64.rpm) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `d12.x86_64` | pgdg | 50.1 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `d12.aarch64` | pgdg | 49.1 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `u22.x86_64` | pgdg | 55.0 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `u22.aarch64` | pgdg | 54.1 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `u24.x86_64` | pgdg | 50.1 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pgauditlogtofile` | 1.7.4 | `u24.aarch64` | pgdg | 49.3 KiB | [postgresql-14-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgauditlogtofile/postgresql-14-pgauditlogtofile_1.7.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fmbiete/pgauditlogtofile" title="Repository" icon="github" subtitle="github.com/fmbiete/pgauditlogtofile" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgauditlogtofile; # install by extension name, for the current active PG version
pig ext install pgauditlogtofile; # install via package alias, for the active PG version
pig ext install pgauditlogtofile -v 17;   # install for PG 17
pig ext install pgauditlogtofile -v 16;   # install for PG 16
pig ext install pgauditlogtofile -v 15;   # install for PG 15
pig ext install pgauditlogtofile -v 14;   # install for PG 14
pig ext install pgauditlogtofile -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgauditlogtofile;
```

