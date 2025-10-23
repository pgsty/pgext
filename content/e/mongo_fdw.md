---
title: "mongo_fdw"
linkTitle: "mongo_fdw"
description: "foreign data wrapper for MongoDB access"
weight: 8700
categories: ["FDW"]
width: full
---

foreign data wrapper for MongoDB access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8700** | {{< badge content="mongo_fdw" link="https://github.com/EnterpriseDB/mongo_fdw" >}} | {{< ext "mongo_fdw" >}} | `1.1` | {{< category "FDW" >}} | {{< license "LGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "redis_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "hdfs_fdw" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} {{< ext "multicorn" >}} {{< ext "jdbc_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/mongo_fdw" >}} | `5.5.1` | {{< bg "18" "mongo_fdw_18*" "red" >}} {{< bg "17" "mongo_fdw_17*" "red" >}} {{< bg "16" "mongo_fdw_16*" "green" >}} {{< bg "15" "mongo_fdw_15*" "green" >}} {{< bg "14" "mongo_fdw_14*" "green" >}} | `mongo_fdw_$v*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "mongo_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.1" "mongo_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_14 : AVAIL 5" "blue" >}} |
|    `el8.aarch64`    |      {{< bg "MISS" "mongo_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.1" "mongo_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_14 : AVAIL 1" "blue" >}} |
|    `el9.x86_64`    |      {{< bg "MISS" "mongo_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.1" "mongo_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_14 : AVAIL 4" "blue" >}} |
|    `el9.aarch64`    |      {{< bg "MISS" "mongo_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 5.5.1" "mongo_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 5.5.1" "mongo_fdw_14 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "mongo_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_16` | 5.5.1 | `el8.x86_64` | pgdg | 74.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/mongo_fdw_16-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_16` | 5.5.1 | `el8.aarch64` | pgdg | 70.7 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/mongo_fdw_16-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_16` | 5.5.1 | `el9.x86_64` | pgdg | 65.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/mongo_fdw_16-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_16` | 5.5.1 | `el9.aarch64` | pgdg | 63.2 KiB | [mongo_fdw_16-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/mongo_fdw_16-5.5.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_15` | 5.5.1 | `el8.x86_64` | pgdg | 77.5 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | 5.5.0 | `el8.x86_64` | pgdg | 74.5 KiB | [mongo_fdw_15-5.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.5.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | 5.4.0 | `el8.x86_64` | pgdg | 74.3 KiB | [mongo_fdw_15-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/mongo_fdw_15-5.4.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_15` | 5.5.1 | `el8.aarch64` | pgdg | 73.8 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/mongo_fdw_15-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_15` | 5.5.1 | `el9.x86_64` | pgdg | 79.2 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | 5.5.0 | `el9.x86_64` | pgdg | 75.9 KiB | [mongo_fdw_15-5.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.5.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | 5.4.0 | `el9.x86_64` | pgdg | 76.6 KiB | [mongo_fdw_15-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/mongo_fdw_15-5.4.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_15` | 5.5.1 | `el9.aarch64` | pgdg | 75.7 KiB | [mongo_fdw_15-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/mongo_fdw_15-5.5.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `mongo_fdw_14` | 5.5.1 | `el8.x86_64` | pgdg | 77.5 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.5.1-1PGDG.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | 5.5.0 | `el8.x86_64` | pgdg | 74.4 KiB | [mongo_fdw_14-5.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.5.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | 5.4.0 | `el8.x86_64` | pgdg | 74.3 KiB | [mongo_fdw_14-5.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.4.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | 5.3.0 | `el8.x86_64` | pgdg | 70.3 KiB | [mongo_fdw_14-5.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.3.0-1.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | 5.2.10 | `el8.x86_64` | pgdg | 63.7 KiB | [mongo_fdw_14-5.2.10-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/mongo_fdw_14-5.2.10-2.rhel8.x86_64.rpm) |
| `mongo_fdw_14` | 5.5.1 | `el8.aarch64` | pgdg | 73.8 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/mongo_fdw_14-5.5.1-1PGDG.rhel8.aarch64.rpm) |
| `mongo_fdw_14` | 5.5.1 | `el9.x86_64` | pgdg | 79.0 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.1-1PGDG.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | 5.5.0 | `el9.x86_64` | pgdg | 75.9 KiB | [mongo_fdw_14-5.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.5.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | 5.4.0 | `el9.x86_64` | pgdg | 76.9 KiB | [mongo_fdw_14-5.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.4.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | 5.3.0 | `el9.x86_64` | pgdg | 72.9 KiB | [mongo_fdw_14-5.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/mongo_fdw_14-5.3.0-1.rhel9.x86_64.rpm) |
| `mongo_fdw_14` | 5.5.1 | `el9.aarch64` | pgdg | 75.6 KiB | [mongo_fdw_14-5.5.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/mongo_fdw_14-5.5.1-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/EnterpriseDB/mongo_fdw" title="Repository" icon="github" subtitle="github.com/EnterpriseDB/mongo_fdw" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mongo_fdw-REL-5_5_1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get mongo_fdw; # get mongo_fdw source code
pig build dep mongo_fdw; # install build dependencies
pig build pkg mongo_fdw; # build extension rpm or deb
pig build ext mongo_fdw; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install mongo_fdw; # install by extension name, for the current active PG version
pig ext install mongo_fdw; # install via package alias, for the active PG version
pig ext install mongo_fdw -v 16;   # install for PG 16
pig ext install mongo_fdw -v 15;   # install for PG 15
pig ext install mongo_fdw -v 14;   # install for PG 14
pig ext install mongo_fdw -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION mongo_fdw;
```

