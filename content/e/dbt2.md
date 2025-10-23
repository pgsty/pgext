---
title: "dbt2"
linkTitle: "dbt2"
description: "OSDL-DBT-2 test kit"
weight: 3220
categories: ["LANG"]
width: full
---

OSDL-DBT-2 test kit


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3220** | {{< badge content="dbt2" link="https://github.com/osdldbt/dbt2" >}} | {{< ext "dbt2" >}} | `0.45.0` | {{< category "LANG" >}} | {{< license "Artistic" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgtap" >}} {{< ext "faker" >}} {{< ext "plpgsql" >}} {{< ext "pg_stat_statements" >}} {{< ext "pg_tle" >}} {{< ext "plv8" >}} {{< ext "pllua" >}} {{< ext "hstore_pllua" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/dbt2" >}} | `0.61.7` | {{< bg "18" "dbt2-pg18-extensions*" "red" >}} {{< bg "17" "dbt2-pg17-extensions*" "green" >}} {{< bg "16" "dbt2-pg16-extensions*" "green" >}} {{< bg "15" "dbt2-pg15-extensions*" "green" >}} {{< bg "14" "dbt2-pg14-extensions*" "green" >}} | `dbt2-pg$v-extensions*` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |  {{< bg "MISS" "dbt2-pg18-extensions : HIDE 0" >}}   |  {{< bg "MISS" "dbt2-pg17-extensions : HIDE 0" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : HIDE 1" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : HIDE 5" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : HIDE 6" >}}   |
|    `el8.aarch64`    |  {{< bg "MISS" "dbt2-pg18-extensions : HIDE 0" >}}   |  {{< bg "MISS" "dbt2-pg17-extensions : HIDE 0" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg16-extensions : HIDE 1" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg15-extensions : HIDE 4" >}}   |  {{< bg "PGDG 0.53.7" "dbt2-pg14-extensions : HIDE 4" >}}   |
|    `el9.x86_64`    |  {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : HIDE 2" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : HIDE 4" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : HIDE 8" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : HIDE 7" >}}   |
|    `el9.aarch64`    |  {{< bg "PGDG 0.61.7" "dbt2-pg18-extensions : HIDE 1" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg17-extensions : HIDE 2" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg16-extensions : HIDE 4" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg15-extensions : HIDE 8" >}}   |  {{< bg "PGDG 0.61.7" "dbt2-pg14-extensions : HIDE 8" >}}   |
|    `d12.x86_64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |
|    `d12.aarch64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |
|    `u22.x86_64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |
|    `u22.aarch64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |
|    `u24.x86_64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |
|    `u24.aarch64`    |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |  {{< bg "MISS" "dbt2 : HIDE 0" >}}   |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg18-extensions` | 0.61.7 | `el9.x86_64` | pgdg | 29.9 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg18-extensions` | 0.61.7 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/dbt2-pg18-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg17-extensions` | 0.61.7 | `el9.x86_64` | pgdg | 29.9 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | 0.61.6 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg17-extensions` | 0.61.7 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg17-extensions` | 0.61.6 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/dbt2-pg17-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg16-extensions` | 0.53.7 | `el8.x86_64` | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg16-extensions` | 0.53.7 | `el8.aarch64` | pgdg | 29.6 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg16-extensions` | 0.61.7 | `el9.x86_64` | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | 0.61.6 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | 0.53.7 | `el9.x86_64` | pgdg | 30.5 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | 0.53.6 | `el9.x86_64` | pgdg | 30.4 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg16-extensions` | 0.61.7 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | 0.61.6 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | 0.53.7 | `el9.aarch64` | pgdg | 30.0 KiB | [dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg16-extensions` | 0.53.6 | `el9.aarch64` | pgdg | 29.9 KiB | [dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/dbt2-pg16-extensions-0.53.6-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg15-extensions` | 0.53.7 | `el8.x86_64` | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.53.4 | `el8.x86_64` | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.50.1 | `el8.x86_64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.49.1 | `el8.x86_64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.48.7 | `el8.x86_64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.53.7 | `el8.aarch64` | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.53.4 | `el8.aarch64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.50.1 | `el8.aarch64` | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.49.1 | `el8.aarch64` | pgdg | 29.2 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.61.7 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.61.6 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.53.7 | `el9.x86_64` | pgdg | 30.5 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.53.4 | `el9.x86_64` | pgdg | 30.3 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.50.1 | `el9.x86_64` | pgdg | 30.1 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.49.1 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.48.7 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.48.3 | `el9.x86_64` | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/dbt2-pg15-extensions-0.48.3-2.rhel9.x86_64.rpm) |
| `dbt2-pg15-extensions` | 0.61.7 | `el9.aarch64` | pgdg | 29.6 KiB | [dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.61.6 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.53.7 | `el9.aarch64` | pgdg | 29.9 KiB | [dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.53.4 | `el9.aarch64` | pgdg | 29.8 KiB | [dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.50.1 | `el9.aarch64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.49.1 | `el9.aarch64` | pgdg | 29.5 KiB | [dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.48.7 | `el9.aarch64` | pgdg | 29.4 KiB | [dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg15-extensions` | 0.48.3 | `el9.aarch64` | pgdg | 29.3 KiB | [dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/dbt2-pg15-extensions-0.48.3-2.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `dbt2-pg14-extensions` | 0.53.7 | `el8.x86_64` | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.53.4 | `el8.x86_64` | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.50.1 | `el8.x86_64` | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.49.1 | `el8.x86_64` | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.48.7 | `el8.x86_64` | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.48.3 | `el8.x86_64` | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/dbt2-pg14-extensions-0.48.3-2.rhel8.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.53.7 | `el8.aarch64` | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.53.4 | `el8.aarch64` | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.50.1 | `el8.aarch64` | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.49.1 | `el8.aarch64` | pgdg | 29.2 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel8.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.61.7 | `el9.x86_64` | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.61.6 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.53.7 | `el9.x86_64` | pgdg | 30.5 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.53.4 | `el9.x86_64` | pgdg | 30.3 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.50.1 | `el9.x86_64` | pgdg | 30.1 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.50.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.49.1 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.49.1-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.48.7 | `el9.x86_64` | pgdg | 30.0 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/dbt2-pg14-extensions-0.48.7-1.rhel9.x86_64.rpm) |
| `dbt2-pg14-extensions` | 0.61.7 | `el9.aarch64` | pgdg | 29.6 KiB | [dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.61.6 | `el9.aarch64` | pgdg | 29.7 KiB | [dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.61.6-2PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.53.7 | `el9.aarch64` | pgdg | 29.9 KiB | [dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.7-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.53.4 | `el9.aarch64` | pgdg | 29.8 KiB | [dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.53.4-1PGDG.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.50.1 | `el9.aarch64` | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.50.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.49.1 | `el9.aarch64` | pgdg | 29.5 KiB | [dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.49.1-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.48.7 | `el9.aarch64` | pgdg | 29.4 KiB | [dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.7-1.rhel9.aarch64.rpm) |
| `dbt2-pg14-extensions` | 0.48.3 | `el9.aarch64` | pgdg | 29.3 KiB | [dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/dbt2-pg14-extensions-0.48.3-2.rhel9.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/osdldbt/dbt2" title="Repository" icon="github" subtitle="github.com/osdldbt/dbt2" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install dbt2; # install by extension name, for the current active PG version
pig ext install dbt2; # install via package alias, for the active PG version
pig ext install dbt2 -v 17;   # install for PG 17
pig ext install dbt2 -v 16;   # install for PG 16
pig ext install dbt2 -v 15;   # install for PG 15
pig ext install dbt2 -v 14;   # install for PG 14
pig ext install dbt2 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION dbt2;
```

