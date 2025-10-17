---
title: "multicorn"
linkTitle: "multicorn"
description: "Fetch foreign data in Python in your PostgreSQL server."
weight: 8510
categories: ["Fdw"]
width: full
---

Fetch foreign data in Python in your PostgreSQL server.

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8510** | {{< badge content="multicorn" link="https://github.com/pgsql-io/multicorn2" >}} | {{< ext "multicorn" "multicorn" >}} | `3.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "odbc_fdw" >}} {{< ext "jdbc_fdw" >}} {{< ext "pgspider_ext" >}} {{< ext "mysql_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "mongo_fdw" >}} {{< ext "redis_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/multicorn" >}} | `3.0` | {{< badge content="18" color="red" alt="multicorn2_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `multicorn2_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "multicorn2_18" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/multicorn2_18-3.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "multicorn2_17" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "multicorn2_16" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "multicorn2_15" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.2-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "multicorn2_14" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.2-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "multicorn2_18" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/multicorn2_18-3.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "multicorn2_17" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "multicorn2_16" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "multicorn2_15" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.2-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "multicorn2_14" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.2-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "multicorn2_18" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "multicorn2_17" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "multicorn2_16" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "multicorn2_15" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.2-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "multicorn2_14" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.2-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "multicorn2_18" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/multicorn2_18-3.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "multicorn2_17" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/multicorn2_17-3.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "multicorn2_16" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/multicorn2_16-3.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "multicorn2_15" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-3.2-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "multicorn2_14" "3.2" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.2-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |
|    `d12.aarch64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |
|    `u22.x86_64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |
|    `u22.aarch64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |
|    `u24.x86_64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |
|    `u24.aarch64`    |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |    {{< pkg "None" >}}     |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_18` | 3.2 | `el8.aarch64` | pgdg | 136.3 KiB | [multicorn2_18-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/multicorn2_18-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_18` | 3.2 | `el8.x86_64` | pgdg | 138.0 KiB | [multicorn2_18-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/multicorn2_18-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_18` | 3.1 | `el8.x86_64` | pgdg | 135.5 KiB | [multicorn2_18-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/multicorn2_18-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_18` | 3.1 | `el8.aarch64` | pgdg | 133.9 KiB | [multicorn2_18-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/multicorn2_18-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_18` | 3.2 | `el9.x86_64` | pgdg | 134.6 KiB | [multicorn2_18-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_18` | 3.2 | `el9.aarch64` | pgdg | 133.5 KiB | [multicorn2_18-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/multicorn2_18-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_18` | 3.1 | `el9.aarch64` | pgdg | 110.0 KiB | [multicorn2_18-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/multicorn2_18-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_18` | 3.1 | `el9.x86_64` | pgdg | 111.0 KiB | [multicorn2_18-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/multicorn2_18-3.1-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_17` | 3.2 | `el8.aarch64` | pgdg | 136.3 KiB | [multicorn2_17-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | 3.2 | `el8.x86_64` | pgdg | 138.0 KiB | [multicorn2_17-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | 3.1 | `el8.x86_64` | pgdg | 135.4 KiB | [multicorn2_17-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | 3.1 | `el8.aarch64` | pgdg | 133.9 KiB | [multicorn2_17-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | 3.0 | `el8.x86_64` | pgdg | 114.6 KiB | [multicorn2_17-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/multicorn2_17-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_17` | 3.0 | `el8.aarch64` | pgdg | 113.1 KiB | [multicorn2_17-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/multicorn2_17-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_17` | 3.2 | `el9.x86_64` | pgdg | 134.6 KiB | [multicorn2_17-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | 3.2 | `el9.aarch64` | pgdg | 133.5 KiB | [multicorn2_17-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/multicorn2_17-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_17` | 3.1 | `el9.aarch64` | pgdg | 109.9 KiB | [multicorn2_17-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/multicorn2_17-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_17` | 3.1 | `el9.x86_64` | pgdg | 111.0 KiB | [multicorn2_17-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | 3.0 | `el9.x86_64` | pgdg | 110.5 KiB | [multicorn2_17-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/multicorn2_17-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_17` | 3.0 | `el9.aarch64` | pgdg | 109.6 KiB | [multicorn2_17-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/multicorn2_17-3.0-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_16` | 3.2 | `el8.aarch64` | pgdg | 136.4 KiB | [multicorn2_16-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | 3.2 | `el8.x86_64` | pgdg | 138.0 KiB | [multicorn2_16-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | 3.1 | `el8.x86_64` | pgdg | 135.4 KiB | [multicorn2_16-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | 3.1 | `el8.aarch64` | pgdg | 134.0 KiB | [multicorn2_16-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | 3.0 | `el8.x86_64` | pgdg | 114.6 KiB | [multicorn2_16-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/multicorn2_16-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_16` | 3.0 | `el8.aarch64` | pgdg | 113.2 KiB | [multicorn2_16-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/multicorn2_16-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_16` | 3.2 | `el9.x86_64` | pgdg | 134.6 KiB | [multicorn2_16-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | 3.2 | `el9.aarch64` | pgdg | 133.6 KiB | [multicorn2_16-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/multicorn2_16-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_16` | 3.1 | `el9.aarch64` | pgdg | 110.1 KiB | [multicorn2_16-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/multicorn2_16-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_16` | 3.1 | `el9.x86_64` | pgdg | 111.0 KiB | [multicorn2_16-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | 3.0 | `el9.x86_64` | pgdg | 110.6 KiB | [multicorn2_16-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/multicorn2_16-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_16` | 3.0 | `el9.aarch64` | pgdg | 109.8 KiB | [multicorn2_16-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/multicorn2_16-3.0-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_15` | 3.2 | `el8.x86_64` | pgdg | 139.3 KiB | [multicorn2_15-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | 3.2 | `el8.aarch64` | pgdg | 137.7 KiB | [multicorn2_15-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | 3.1 | `el8.aarch64` | pgdg | 135.2 KiB | [multicorn2_15-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | 3.1 | `el8.x86_64` | pgdg | 136.8 KiB | [multicorn2_15-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | 3.0 | `el8.aarch64` | pgdg | 114.2 KiB | [multicorn2_15-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_15` | 3.0 | `el8.x86_64` | pgdg | 115.9 KiB | [multicorn2_15-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_15` | 2.4 | `el8.aarch64` | pgdg | 35.4 KiB | [multicorn2_15-2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-2.4-1.rhel8.aarch64.rpm) |
| `multicorn2_15` | 2.4 | `el8.x86_64` | pgdg | 111.5 KiB | [multicorn2_15-2.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-2.4-2.rhel8.x86_64.rpm) |
| `multicorn2_15` | 2.4 | `el8.x86_64` | pgdg | 36.8 KiB | [multicorn2_15-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/multicorn2_15-2.4-1.rhel8.x86_64.rpm) |
| `multicorn2_15` | 2.4 | `el8.aarch64` | pgdg | 110.0 KiB | [multicorn2_15-2.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/multicorn2_15-2.4-2.rhel8.aarch64.rpm) |
| `multicorn2_15` | 3.2 | `el9.x86_64` | pgdg | 138.5 KiB | [multicorn2_15-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | 3.2 | `el9.aarch64` | pgdg | 136.7 KiB | [multicorn2_15-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_15` | 3.1 | `el9.x86_64` | pgdg | 114.8 KiB | [multicorn2_15-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | 3.1 | `el9.aarch64` | pgdg | 113.4 KiB | [multicorn2_15-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_15` | 3.0 | `el9.aarch64` | pgdg | 112.9 KiB | [multicorn2_15-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-3.0-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_15` | 3.0 | `el9.x86_64` | pgdg | 114.6 KiB | [multicorn2_15-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_15` | 2.4 | `el9.x86_64` | pgdg | 37.4 KiB | [multicorn2_15-2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-2.4-1.rhel9.x86_64.rpm) |
| `multicorn2_15` | 2.4 | `el9.aarch64` | pgdg | 108.2 KiB | [multicorn2_15-2.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-2.4-2.rhel9.aarch64.rpm) |
| `multicorn2_15` | 2.4 | `el9.aarch64` | pgdg | 35.8 KiB | [multicorn2_15-2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/multicorn2_15-2.4-1.rhel9.aarch64.rpm) |
| `multicorn2_15` | 2.4 | `el9.x86_64` | pgdg | 109.9 KiB | [multicorn2_15-2.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/multicorn2_15-2.4-2.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_14` | 3.2 | `el8.aarch64` | pgdg | 137.7 KiB | [multicorn2_14-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | 3.2 | `el8.x86_64` | pgdg | 139.4 KiB | [multicorn2_14-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | 3.1 | `el8.x86_64` | pgdg | 136.8 KiB | [multicorn2_14-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | 3.1 | `el8.aarch64` | pgdg | 135.2 KiB | [multicorn2_14-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | 3.0 | `el8.x86_64` | pgdg | 115.9 KiB | [multicorn2_14-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_14` | 3.0 | `el8.aarch64` | pgdg | 114.2 KiB | [multicorn2_14-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_14` | 2.4 | `el8.aarch64` | pgdg | 35.4 KiB | [multicorn2_14-2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.4-1.rhel8.aarch64.rpm) |
| `multicorn2_14` | 2.4 | `el8.aarch64` | pgdg | 110.0 KiB | [multicorn2_14-2.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.4-2.rhel8.aarch64.rpm) |
| `multicorn2_14` | 2.4 | `el8.x86_64` | pgdg | 36.8 KiB | [multicorn2_14-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.4-1.rhel8.x86_64.rpm) |
| `multicorn2_14` | 2.4 | `el8.x86_64` | pgdg | 111.5 KiB | [multicorn2_14-2.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.4-2.rhel8.x86_64.rpm) |
| `multicorn2_14` | 2.3 | `el8.x86_64` | pgdg | 115.3 KiB | [multicorn2_14-2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/multicorn2_14-2.3-1.rhel8.x86_64.rpm) |
| `multicorn2_14` | 2.3 | `el8.aarch64` | pgdg | 113.9 KiB | [multicorn2_14-2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/multicorn2_14-2.3-1.rhel8.aarch64.rpm) |
| `multicorn2_14` | 3.2 | `el9.x86_64` | pgdg | 138.5 KiB | [multicorn2_14-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | 3.2 | `el9.aarch64` | pgdg | 136.7 KiB | [multicorn2_14-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | 3.1 | `el9.x86_64` | pgdg | 114.8 KiB | [multicorn2_14-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | 3.1 | `el9.aarch64` | pgdg | 113.3 KiB | [multicorn2_14-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | 3.0 | `el9.aarch64` | pgdg | 112.8 KiB | [multicorn2_14-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-3.0-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_14` | 3.0 | `el9.x86_64` | pgdg | 114.2 KiB | [multicorn2_14-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_14` | 2.4 | `el9.aarch64` | pgdg | 108.3 KiB | [multicorn2_14-2.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.4-2.rhel9.aarch64.rpm) |
| `multicorn2_14` | 2.4 | `el9.x86_64` | pgdg | 109.9 KiB | [multicorn2_14-2.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.4-2.rhel9.x86_64.rpm) |
| `multicorn2_14` | 2.4 | `el9.aarch64` | pgdg | 35.7 KiB | [multicorn2_14-2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.4-1.rhel9.aarch64.rpm) |
| `multicorn2_14` | 2.4 | `el9.x86_64` | pgdg | 37.6 KiB | [multicorn2_14-2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.4-1.rhel9.x86_64.rpm) |
| `multicorn2_14` | 2.3 | `el9.aarch64` | pgdg | 113.1 KiB | [multicorn2_14-2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/multicorn2_14-2.3-1.rhel9.aarch64.rpm) |
| `multicorn2_14` | 2.3 | `el9.x86_64` | pgdg | 114.0 KiB | [multicorn2_14-2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/multicorn2_14-2.3-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `multicorn2_13` | 3.2 | `el8.aarch64` | pgdg | 137.3 KiB | [multicorn2_13-3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-3.2-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_13` | 3.2 | `el8.x86_64` | pgdg | 138.6 KiB | [multicorn2_13-3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-3.2-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_13` | 3.1 | `el8.x86_64` | pgdg | 136.0 KiB | [multicorn2_13-3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-3.1-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_13` | 3.1 | `el8.aarch64` | pgdg | 134.7 KiB | [multicorn2_13-3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-3.1-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_13` | 3.0 | `el8.aarch64` | pgdg | 113.8 KiB | [multicorn2_13-3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-3.0-1PGDG.rhel8.aarch64.rpm) |
| `multicorn2_13` | 3.0 | `el8.x86_64` | pgdg | 115.1 KiB | [multicorn2_13-3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-3.0-1PGDG.rhel8.x86_64.rpm) |
| `multicorn2_13` | 2.4 | `el8.x86_64` | pgdg | 111.3 KiB | [multicorn2_13-2.4-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-2.4-2.rhel8.x86_64.rpm) |
| `multicorn2_13` | 2.4 | `el8.aarch64` | pgdg | 35.4 KiB | [multicorn2_13-2.4-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-2.4-1.rhel8.aarch64.rpm) |
| `multicorn2_13` | 2.4 | `el8.x86_64` | pgdg | 36.6 KiB | [multicorn2_13-2.4-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-2.4-1.rhel8.x86_64.rpm) |
| `multicorn2_13` | 2.4 | `el8.aarch64` | pgdg | 110.0 KiB | [multicorn2_13-2.4-2.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-2.4-2.rhel8.aarch64.rpm) |
| `multicorn2_13` | 2.3 | `el8.aarch64` | pgdg | 113.9 KiB | [multicorn2_13-2.3-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/multicorn2_13-2.3-1.rhel8.aarch64.rpm) |
| `multicorn2_13` | 2.3 | `el8.x86_64` | pgdg | 115.1 KiB | [multicorn2_13-2.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/multicorn2_13-2.3-1.rhel8.x86_64.rpm) |
| `multicorn2_13` | 3.2 | `el9.x86_64` | pgdg | 137.8 KiB | [multicorn2_13-3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-3.2-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_13` | 3.2 | `el9.aarch64` | pgdg | 136.2 KiB | [multicorn2_13-3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-3.2-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_13` | 3.1 | `el9.x86_64` | pgdg | 114.2 KiB | [multicorn2_13-3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-3.1-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_13` | 3.1 | `el9.aarch64` | pgdg | 112.9 KiB | [multicorn2_13-3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-3.1-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_13` | 3.0 | `el9.x86_64` | pgdg | 113.7 KiB | [multicorn2_13-3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-3.0-1PGDG.rhel9.x86_64.rpm) |
| `multicorn2_13` | 3.0 | `el9.aarch64` | pgdg | 112.5 KiB | [multicorn2_13-3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-3.0-1PGDG.rhel9.aarch64.rpm) |
| `multicorn2_13` | 2.4 | `el9.aarch64` | pgdg | 108.3 KiB | [multicorn2_13-2.4-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-2.4-2.rhel9.aarch64.rpm) |
| `multicorn2_13` | 2.4 | `el9.aarch64` | pgdg | 35.8 KiB | [multicorn2_13-2.4-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-2.4-1.rhel9.aarch64.rpm) |
| `multicorn2_13` | 2.4 | `el9.x86_64` | pgdg | 109.8 KiB | [multicorn2_13-2.4-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-2.4-2.rhel9.x86_64.rpm) |
| `multicorn2_13` | 2.4 | `el9.x86_64` | pgdg | 37.4 KiB | [multicorn2_13-2.4-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-2.4-1.rhel9.x86_64.rpm) |
| `multicorn2_13` | 2.3 | `el9.aarch64` | pgdg | 113.0 KiB | [multicorn2_13-2.3-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/multicorn2_13-2.3-1.rhel9.aarch64.rpm) |
| `multicorn2_13` | 2.3 | `el9.x86_64` | pgdg | 114.0 KiB | [multicorn2_13-2.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/multicorn2_13-2.3-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgsql-io/multicorn2" title="Repository" icon="github" subtitle="github.com/pgsql-io/multicorn2" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install multicorn; # install by extension name, for the current active PG version
pig ext install multicorn; # install via package alias, for the active PG version
pig ext install multicorn -v 17;   # install for PG 17
pig ext install multicorn -v 16;   # install for PG 16
pig ext install multicorn -v 15;   # install for PG 15
pig ext install multicorn -v 14;   # install for PG 14
pig ext install multicorn -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION multicorn;
```

