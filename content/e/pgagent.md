---
title: "pgagent"
linkTitle: "pgagent"
description: "A PostgreSQL job scheduler"
weight: 5880
categories: ["Admin"]
width: full
---

A PostgreSQL job scheduler

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5880** | {{< badge content="pgagent" link="https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html" >}} | {{< ext "pgagent" "pgagent" >}} | `4.2.3` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_jobmon" >}} {{< ext "pg_partman" >}} {{< ext "pglogical" >}} {{< ext "pg_background" >}} {{< ext "pg_repack" >}} {{< ext "pg_rewrite" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pgagent" >}} | `4.2.3` | {{< badge content="18" color="red" alt="pgagent_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgagent_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pgagent" >}} | `4.2.3` | {{< badge content="18" color="red" alt="pgagent" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pgagent` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pgagent_18" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgagent_18-4.2.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgagent_17" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgagent_17-4.2.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgagent_16" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgagent_16-4.2.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgagent_15" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgagent_15-4.2.3-1PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pgagent_14" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.3-1PGDG.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pgagent_18" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgagent_18-4.2.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgagent_17" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgagent_17-4.2.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgagent_16" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgagent_16-4.2.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgagent_15" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgagent_15-4.2.3-1PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pgagent_14" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgagent_14-4.2.3-1PGDG.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pgagent_18" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgagent_18-4.2.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgagent_17" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgagent_17-4.2.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgagent_16" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgagent_16-4.2.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgagent_15" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgagent_15-4.2.3-1PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pgagent_14" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.3-1PGDG.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pgagent_18" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgagent_18-4.2.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgagent_17" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgagent_17-4.2.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgagent_16" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgagent_16-4.2.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgagent_15" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgagent_15-4.2.3-1PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pgagent_14" "4.2.3" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgagent_14-4.2.3-1PGDG.rhel9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |
|    `d12.aarch64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |
|    `u22.x86_64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |
|    `u22.aarch64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |
|    `u24.x86_64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |
|    `u24.aarch64`    |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |    {{< pkg "pgagent" >}}     |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_18` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_18-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pgagent_18-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_18` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_18-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pgagent_18-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_18` | 4.2.3 | `el9.x86_64` | pgdg | 121.5 KiB | [pgagent_18-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pgagent_18-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_18` | 4.2.3 | `el9.aarch64` | pgdg | 117.8 KiB | [pgagent_18-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pgagent_18-4.2.3-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_17` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_17-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgagent_17-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_17` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_17-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgagent_17-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_17` | 4.2.2 | `el8.x86_64` | pgdg | 133.9 KiB | [pgagent_17-4.2.2-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pgagent_17-4.2.2-5PGDG.rhel8.x86_64.rpm) |
| `pgagent_17` | 4.2.2 | `el8.aarch64` | pgdg | 128.2 KiB | [pgagent_17-4.2.2-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pgagent_17-4.2.2-5PGDG.rhel8.aarch64.rpm) |
| `pgagent_17` | 4.2.3 | `el9.x86_64` | pgdg | 121.9 KiB | [pgagent_17-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgagent_17-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_17` | 4.2.3 | `el9.aarch64` | pgdg | 118.0 KiB | [pgagent_17-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgagent_17-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_17` | 4.2.2 | `el9.aarch64` | pgdg | 116.7 KiB | [pgagent_17-4.2.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pgagent_17-4.2.2-5PGDG.rhel9.aarch64.rpm) |
| `pgagent_17` | 4.2.2 | `el9.x86_64` | pgdg | 120.3 KiB | [pgagent_17-4.2.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pgagent_17-4.2.2-5PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_16` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_16-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgagent_16-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_16` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_16-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgagent_16-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_16` | 4.2.2 | `el8.aarch64` | pgdg | 127.9 KiB | [pgagent_16-4.2.2-3.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pgagent_16-4.2.2-3.rhel8.aarch64.rpm) |
| `pgagent_16` | 4.2.2 | `el8.x86_64` | pgdg | 133.5 KiB | [pgagent_16-4.2.2-3.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pgagent_16-4.2.2-3.rhel8.x86_64.rpm) |
| `pgagent_16` | 4.2.3 | `el9.x86_64` | pgdg | 122.0 KiB | [pgagent_16-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgagent_16-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_16` | 4.2.3 | `el9.aarch64` | pgdg | 118.0 KiB | [pgagent_16-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgagent_16-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_16` | 4.2.2 | `el9.aarch64` | pgdg | 116.4 KiB | [pgagent_16-4.2.2-3.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pgagent_16-4.2.2-3.rhel9.aarch64.rpm) |
| `pgagent_16` | 4.2.2 | `el9.x86_64` | pgdg | 120.0 KiB | [pgagent_16-4.2.2-3.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pgagent_16-4.2.2-3.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_15` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_15-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgagent_15-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_15` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_15-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgagent_15-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_15` | 4.2.2 | `el8.aarch64` | pgdg | 127.7 KiB | [pgagent_15-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pgagent_15-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_15` | 4.2.2 | `el8.x86_64` | pgdg | 133.4 KiB | [pgagent_15-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pgagent_15-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_15` | 4.2.3 | `el9.x86_64` | pgdg | 122.0 KiB | [pgagent_15-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgagent_15-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_15` | 4.2.3 | `el9.aarch64` | pgdg | 118.1 KiB | [pgagent_15-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgagent_15-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_15` | 4.2.2 | `el9.aarch64` | pgdg | 114.4 KiB | [pgagent_15-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pgagent_15-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_15` | 4.2.2 | `el9.x86_64` | pgdg | 119.1 KiB | [pgagent_15-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pgagent_15-4.2.2-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_14` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_14-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgagent_14-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_14` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_14-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_14` | 4.2.2 | `el8.aarch64` | pgdg | 127.7 KiB | [pgagent_14-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pgagent_14-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_14` | 4.2.2 | `el8.x86_64` | pgdg | 133.4 KiB | [pgagent_14-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_14` | 4.2.1 | `el8.x86_64` | pgdg | 153.1 KiB | [pgagent_14-4.2.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pgagent_14-4.2.1-1.rhel8.x86_64.rpm) |
| `pgagent_14` | 4.2.3 | `el9.aarch64` | pgdg | 118.0 KiB | [pgagent_14-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgagent_14-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_14` | 4.2.3 | `el9.x86_64` | pgdg | 122.0 KiB | [pgagent_14-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_14` | 4.2.2 | `el9.x86_64` | pgdg | 119.1 KiB | [pgagent_14-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.2-1.rhel9.x86_64.rpm) |
| `pgagent_14` | 4.2.2 | `el9.aarch64` | pgdg | 114.4 KiB | [pgagent_14-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pgagent_14-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_14` | 4.2.1 | `el9.x86_64` | pgdg | 138.6 KiB | [pgagent_14-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pgagent_14-4.2.1-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pgagent_13` | 4.2.3 | `el8.x86_64` | pgdg | 135.5 KiB | [pgagent_13-4.2.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.3-1PGDG.rhel8.x86_64.rpm) |
| `pgagent_13` | 4.2.3 | `el8.aarch64` | pgdg | 129.6 KiB | [pgagent_13-4.2.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgagent_13-4.2.3-1PGDG.rhel8.aarch64.rpm) |
| `pgagent_13` | 4.2.2 | `el8.aarch64` | pgdg | 127.6 KiB | [pgagent_13-4.2.2-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pgagent_13-4.2.2-1.rhel8.aarch64.rpm) |
| `pgagent_13` | 4.2.2 | `el8.x86_64` | pgdg | 133.4 KiB | [pgagent_13-4.2.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.2-1.rhel8.x86_64.rpm) |
| `pgagent_13` | 4.2.1 | `el8.x86_64` | pgdg | 153.1 KiB | [pgagent_13-4.2.1-0.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.2.1-0.rhel8.x86_64.rpm) |
| `pgagent_13` | 4.0.0 | `el8.x86_64` | pgdg | 151.1 KiB | [pgagent_13-4.0.0-4.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pgagent_13-4.0.0-4.rhel8.x86_64.rpm) |
| `pgagent_13` | 4.2.3 | `el9.x86_64` | pgdg | 121.9 KiB | [pgagent_13-4.2.3-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.3-1PGDG.rhel9.x86_64.rpm) |
| `pgagent_13` | 4.2.3 | `el9.aarch64` | pgdg | 118.1 KiB | [pgagent_13-4.2.3-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgagent_13-4.2.3-1PGDG.rhel9.aarch64.rpm) |
| `pgagent_13` | 4.2.2 | `el9.x86_64` | pgdg | 119.1 KiB | [pgagent_13-4.2.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.2-1.rhel9.x86_64.rpm) |
| `pgagent_13` | 4.2.2 | `el9.aarch64` | pgdg | 114.4 KiB | [pgagent_13-4.2.2-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pgagent_13-4.2.2-1.rhel9.aarch64.rpm) |
| `pgagent_13` | 4.2.1 | `el9.x86_64` | pgdg | 138.6 KiB | [pgagent_13-4.2.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pgagent_13-4.2.1-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://www.pgadmin.org/docs/pgadmin4/development/pgagent.html" title="Repository" icon="link" subtitle="www.pgadmin.org/docs/pgadmin4/development/pgagent.html" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pgagent; # install by extension name, for the current active PG version
pig ext install pgagent; # install via package alias, for the active PG version
pig ext install pgagent -v 17;   # install for PG 17
pig ext install pgagent -v 16;   # install for PG 16
pig ext install pgagent -v 15;   # install for PG 15
pig ext install pgagent -v 14;   # install for PG 14
pig ext install pgagent -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pgagent;
```

