---
title: "hdfs_fdw"
linkTitle: "hdfs_fdw"
description: "foreign-data wrapper for remote hdfs servers"
weight: 8740
categories: ["FDW"]
width: full
---

foreign-data wrapper for remote hdfs servers


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8740** | {{< badge content="hdfs_fdw" link="https://github.com/EnterpriseDB/hdfs_fdw" >}} | {{< ext "hdfs_fdw" >}} | `2.3.2` | {{< category "FDW" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_parquet" >}} {{< ext "mongo_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} {{< ext "aws_s3" >}} {{< ext "duckdb_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/hdfs_fdw" >}} | `2.3.2` | {{< bg "18" "hdfs_fdw_18*" "red" >}} {{< bg "17" "hdfs_fdw_17*" "green" >}} {{< bg "16" "hdfs_fdw_16*" "green" >}} {{< bg "15" "hdfs_fdw_15*" "green" >}} {{< bg "14" "hdfs_fdw_14*" "green" >}} | `hdfs_fdw_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.3.3" "hdfs_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_14 : AVAIL 7" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.3.3" "hdfs_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_14 : AVAIL 4" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.3.3" "hdfs_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_14 : AVAIL 6" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.3.3" "hdfs_fdw_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 2.3.3" "hdfs_fdw_14 : AVAIL 5" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "hdfs_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdfs_fdw_18` | 2.3.3 | `el8.x86_64` | pgdg | 116.2 KiB | [hdfs_fdw_18-2.3.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/hdfs_fdw_18-2.3.3-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_18` | 2.3.3 | `el8.aarch64` | pgdg | 113.2 KiB | [hdfs_fdw_18-2.3.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/hdfs_fdw_18-2.3.3-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_18` | 2.3.3 | `el9.x86_64` | pgdg | 116.4 KiB | [hdfs_fdw_18-2.3.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/hdfs_fdw_18-2.3.3-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_18` | 2.3.3 | `el9.aarch64` | pgdg | 114.2 KiB | [hdfs_fdw_18-2.3.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/hdfs_fdw_18-2.3.3-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdfs_fdw_17` | 2.3.3 | `el8.x86_64` | pgdg | 116.1 KiB | [hdfs_fdw_17-2.3.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hdfs_fdw_17-2.3.3-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_17` | 2.3.2 | `el8.x86_64` | pgdg | 117.9 KiB | [hdfs_fdw_17-2.3.2-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/hdfs_fdw_17-2.3.2-3PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_17` | 2.3.3 | `el8.aarch64` | pgdg | 113.1 KiB | [hdfs_fdw_17-2.3.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hdfs_fdw_17-2.3.3-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_17` | 2.3.2 | `el8.aarch64` | pgdg | 115.3 KiB | [hdfs_fdw_17-2.3.2-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/hdfs_fdw_17-2.3.2-3PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_17` | 2.3.3 | `el9.x86_64` | pgdg | 116.3 KiB | [hdfs_fdw_17-2.3.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hdfs_fdw_17-2.3.3-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_17` | 2.3.2 | `el9.x86_64` | pgdg | 118.8 KiB | [hdfs_fdw_17-2.3.2-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/hdfs_fdw_17-2.3.2-3PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_17` | 2.3.3 | `el9.aarch64` | pgdg | 113.9 KiB | [hdfs_fdw_17-2.3.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hdfs_fdw_17-2.3.3-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_17` | 2.3.2 | `el9.aarch64` | pgdg | 116.4 KiB | [hdfs_fdw_17-2.3.2-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/hdfs_fdw_17-2.3.2-3PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdfs_fdw_16` | 2.3.3 | `el8.x86_64` | pgdg | 116.1 KiB | [hdfs_fdw_16-2.3.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hdfs_fdw_16-2.3.3-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_16` | 2.3.1 | `el8.x86_64` | pgdg | 129.6 KiB | [hdfs_fdw_16-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/hdfs_fdw_16-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_16` | 2.3.3 | `el8.aarch64` | pgdg | 113.2 KiB | [hdfs_fdw_16-2.3.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hdfs_fdw_16-2.3.3-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_16` | 2.3.1 | `el8.aarch64` | pgdg | 126.5 KiB | [hdfs_fdw_16-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/hdfs_fdw_16-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_16` | 2.3.3 | `el9.x86_64` | pgdg | 116.2 KiB | [hdfs_fdw_16-2.3.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hdfs_fdw_16-2.3.3-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_16` | 2.3.1 | `el9.x86_64` | pgdg | 130.6 KiB | [hdfs_fdw_16-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/hdfs_fdw_16-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_16` | 2.3.3 | `el9.aarch64` | pgdg | 114.3 KiB | [hdfs_fdw_16-2.3.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hdfs_fdw_16-2.3.3-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_16` | 2.3.1 | `el9.aarch64` | pgdg | 128.4 KiB | [hdfs_fdw_16-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/hdfs_fdw_16-2.3.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdfs_fdw_15` | 2.3.3 | `el8.x86_64` | pgdg | 115.9 KiB | [hdfs_fdw_15-2.3.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hdfs_fdw_15-2.3.3-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.2 | `el8.x86_64` | pgdg | 117.4 KiB | [hdfs_fdw_15-2.3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hdfs_fdw_15-2.3.2-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.1 | `el8.x86_64` | pgdg | 129.0 KiB | [hdfs_fdw_15-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hdfs_fdw_15-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.0 | `el8.x86_64` | pgdg | 127.9 KiB | [hdfs_fdw_15-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hdfs_fdw_15-2.3.0-1.rhel8.x86_64.rpm) |
| `hdfs_fdw_15` | 2.2.0 | `el8.x86_64` | pgdg | 117.1 KiB | [hdfs_fdw_15-2.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/hdfs_fdw_15-2.2.0-1.rhel8.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.3 | `el8.aarch64` | pgdg | 112.7 KiB | [hdfs_fdw_15-2.3.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hdfs_fdw_15-2.3.3-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.2 | `el8.aarch64` | pgdg | 114.4 KiB | [hdfs_fdw_15-2.3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hdfs_fdw_15-2.3.2-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.1 | `el8.aarch64` | pgdg | 113.6 KiB | [hdfs_fdw_15-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hdfs_fdw_15-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.0 | `el8.aarch64` | pgdg | 124.9 KiB | [hdfs_fdw_15-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/hdfs_fdw_15-2.3.0-1.rhel8.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.3 | `el9.x86_64` | pgdg | 116.9 KiB | [hdfs_fdw_15-2.3.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hdfs_fdw_15-2.3.3-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.2 | `el9.x86_64` | pgdg | 119.0 KiB | [hdfs_fdw_15-2.3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hdfs_fdw_15-2.3.2-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.1 | `el9.x86_64` | pgdg | 130.9 KiB | [hdfs_fdw_15-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hdfs_fdw_15-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.0 | `el9.x86_64` | pgdg | 130.0 KiB | [hdfs_fdw_15-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hdfs_fdw_15-2.3.0-1.rhel9.x86_64.rpm) |
| `hdfs_fdw_15` | 2.2.0 | `el9.x86_64` | pgdg | 119.7 KiB | [hdfs_fdw_15-2.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/hdfs_fdw_15-2.2.0-1.rhel9.x86_64.rpm) |
| `hdfs_fdw_15` | 2.3.3 | `el9.aarch64` | pgdg | 114.2 KiB | [hdfs_fdw_15-2.3.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hdfs_fdw_15-2.3.3-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.2 | `el9.aarch64` | pgdg | 116.6 KiB | [hdfs_fdw_15-2.3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hdfs_fdw_15-2.3.2-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.1 | `el9.aarch64` | pgdg | 115.9 KiB | [hdfs_fdw_15-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hdfs_fdw_15-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_15` | 2.3.0 | `el9.aarch64` | pgdg | 127.6 KiB | [hdfs_fdw_15-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hdfs_fdw_15-2.3.0-1.rhel9.aarch64.rpm) |
| `hdfs_fdw_15` | 2.2.0 | `el9.aarch64` | pgdg | 117.3 KiB | [hdfs_fdw_15-2.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/hdfs_fdw_15-2.2.0-1.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdfs_fdw_14` | 2.3.3 | `el8.x86_64` | pgdg | 115.9 KiB | [hdfs_fdw_14-2.3.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.3.3-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.2 | `el8.x86_64` | pgdg | 117.4 KiB | [hdfs_fdw_14-2.3.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.3.2-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.1 | `el8.x86_64` | pgdg | 128.8 KiB | [hdfs_fdw_14-2.3.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.3.1-1PGDG.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.0 | `el8.x86_64` | pgdg | 127.7 KiB | [hdfs_fdw_14-2.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.3.0-1.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.2.0 | `el8.x86_64` | pgdg | 117.0 KiB | [hdfs_fdw_14-2.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.2.0-1.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.1.0 | `el8.x86_64` | pgdg | 112.0 KiB | [hdfs_fdw_14-2.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.1.0-1.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.0.9 | `el8.x86_64` | pgdg | 94.6 KiB | [hdfs_fdw_14-2.0.9-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/hdfs_fdw_14-2.0.9-2.rhel8.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.3 | `el8.aarch64` | pgdg | 112.5 KiB | [hdfs_fdw_14-2.3.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hdfs_fdw_14-2.3.3-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.2 | `el8.aarch64` | pgdg | 114.4 KiB | [hdfs_fdw_14-2.3.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hdfs_fdw_14-2.3.2-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.1 | `el8.aarch64` | pgdg | 113.5 KiB | [hdfs_fdw_14-2.3.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hdfs_fdw_14-2.3.1-1PGDG.rhel8.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.0 | `el8.aarch64` | pgdg | 124.5 KiB | [hdfs_fdw_14-2.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/hdfs_fdw_14-2.3.0-1.rhel8.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.3 | `el9.x86_64` | pgdg | 116.9 KiB | [hdfs_fdw_14-2.3.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.3.3-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.2 | `el9.x86_64` | pgdg | 119.0 KiB | [hdfs_fdw_14-2.3.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.3.2-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.1 | `el9.x86_64` | pgdg | 130.7 KiB | [hdfs_fdw_14-2.3.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.3.1-1PGDG.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.0 | `el9.x86_64` | pgdg | 129.7 KiB | [hdfs_fdw_14-2.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.3.0-1.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.2.0 | `el9.x86_64` | pgdg | 119.5 KiB | [hdfs_fdw_14-2.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.2.0-1.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.1.0 | `el9.x86_64` | pgdg | 114.3 KiB | [hdfs_fdw_14-2.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/hdfs_fdw_14-2.1.0-1.rhel9.x86_64.rpm) |
| `hdfs_fdw_14` | 2.3.3 | `el9.aarch64` | pgdg | 114.1 KiB | [hdfs_fdw_14-2.3.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hdfs_fdw_14-2.3.3-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.2 | `el9.aarch64` | pgdg | 116.5 KiB | [hdfs_fdw_14-2.3.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hdfs_fdw_14-2.3.2-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.1 | `el9.aarch64` | pgdg | 115.9 KiB | [hdfs_fdw_14-2.3.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hdfs_fdw_14-2.3.1-1PGDG.rhel9.aarch64.rpm) |
| `hdfs_fdw_14` | 2.3.0 | `el9.aarch64` | pgdg | 127.4 KiB | [hdfs_fdw_14-2.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hdfs_fdw_14-2.3.0-1.rhel9.aarch64.rpm) |
| `hdfs_fdw_14` | 2.2.0 | `el9.aarch64` | pgdg | 117.1 KiB | [hdfs_fdw_14-2.2.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/hdfs_fdw_14-2.2.0-1.rhel9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/hdfs_fdw" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/hdfs_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hdfs_fdw-2.0.5.tar.gz" >}}
{{< /cards >}}


```bash
pig build get hdfs_fdw; # get hdfs_fdw source code
pig build dep hdfs_fdw; # install build dependencies
pig build pkg hdfs_fdw; # build extension rpm or deb
pig build ext hdfs_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install hdfs_fdw; # install by extension name, for the current active PG version
pig ext install hdfs_fdw; # install via package alias, for the active PG version
pig ext install hdfs_fdw -v 17;   # install for PG 17
pig ext install hdfs_fdw -v 16;   # install for PG 16
pig ext install hdfs_fdw -v 15;   # install for PG 15
pig ext install hdfs_fdw -v 14;   # install for PG 14
pig ext install hdfs_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION hdfs_fdw;
```

