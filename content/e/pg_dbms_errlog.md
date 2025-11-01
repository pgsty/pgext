---
title: "pg_dbms_errlog"
linkTitle: "pg_dbms_errlog"
description: "Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table."
weight: 9270
categories: ["SIM"]
width: full
---

Emulate DBMS_ERRLOG Oracle module to log DML errors in a dedicated table.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9270** | {{< badge content="pg_dbms_errlog" link="https://github.com/HexaCluster/pg_dbms_errlog" >}} | {{< ext "pg_dbms_errlog" >}} | `2.2` | {{< category "SIM" >}} | {{< license "ISC" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_dbms_errlog" >}} | `2.2` | {{< bg "18" "pg_dbms_errlog_18*" "green" >}} {{< bg "17" "pg_dbms_errlog_17*" "green" >}} {{< bg "16" "pg_dbms_errlog_16*" "green" >}} {{< bg "15" "pg_dbms_errlog_15*" "green" >}} {{< bg "14" "pg_dbms_errlog_14*" "green" >}} {{< bg "13" "pg_dbms_errlog_13*" "green" >}} | `pg_dbms_errlog_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `el8.aarch64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `el9.x86_64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `el9.aarch64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `el10.x86_64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `el10.aarch64`    | {{< bg "PGDG 2.2" "pg_dbms_errlog_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.2" "pg_dbms_errlog_14 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_errlog_13 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_errlog : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_18` | 2.2 | `el8.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_18` | 2.2 | `el8.aarch64` | pgdg | 31.3 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_18` | 2.2 | `el9.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_18` | 2.2 | `el9.aarch64` | pgdg | 31.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_18` | 2.2 | `el10.x86_64` | pgdg | 32.6 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_18` | 2.2 | `el10.aarch64` | pgdg | 32.2 KiB | [pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_errlog_18-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_17` | 2.2 | `el8.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_17` | 2.2 | `el8.aarch64` | pgdg | 31.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_17` | 2.2 | `el9.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_17` | 2.2 | `el9.aarch64` | pgdg | 31.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_17` | 2.2 | `el10.x86_64` | pgdg | 32.6 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_17` | 2.2 | `el10.aarch64` | pgdg | 32.3 KiB | [pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_errlog_17-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_16` | 2.2 | `el8.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_16` | 2.2 | `el8.aarch64` | pgdg | 31.3 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_16` | 2.2 | `el9.x86_64` | pgdg | 32.0 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_16` | 2.2 | `el9.aarch64` | pgdg | 31.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_16` | 2.2 | `el10.x86_64` | pgdg | 32.6 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_16` | 2.2 | `el10.aarch64` | pgdg | 32.2 KiB | [pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_errlog_16-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_15` | 2.2 | `el8.x86_64` | pgdg | 32.6 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_15` | 2.2 | `el8.aarch64` | pgdg | 31.8 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_15` | 2.2 | `el9.x86_64` | pgdg | 33.1 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_15` | 2.2 | `el9.aarch64` | pgdg | 32.7 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_15` | 2.2 | `el10.x86_64` | pgdg | 33.5 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_15` | 2.2 | `el10.aarch64` | pgdg | 33.2 KiB | [pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_errlog_15-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_errlog_14` | 2.2 | `el8.x86_64` | pgdg | 32.5 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_errlog_14` | 2.2 | `el8.aarch64` | pgdg | 31.8 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_errlog_14` | 2.2 | `el9.x86_64` | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_errlog_14` | 2.2 | `el9.aarch64` | pgdg | 32.6 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_errlog_14` | 2.2 | `el10.x86_64` | pgdg | 33.4 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_errlog_14` | 2.2 | `el10.aarch64` | pgdg | 33.1 KiB | [pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_errlog_14-2.2-1PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_errlog" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_errlog" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_dbms_errlog; # install by extension name, for the current active PG version
pig ext install pg_dbms_errlog; # install via package alias, for the active PG version
pig ext install pg_dbms_errlog -v 18;   # install for PG 18
pig ext install pg_dbms_errlog -v 17;   # install for PG 17
pig ext install pg_dbms_errlog -v 16;   # install for PG 16
pig ext install pg_dbms_errlog -v 15;   # install for PG 15
pig ext install pg_dbms_errlog -v 14;   # install for PG 14
pig ext install pg_dbms_errlog -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_dbms_errlog CASCADE SCHEMA dbms_errlog;
```

