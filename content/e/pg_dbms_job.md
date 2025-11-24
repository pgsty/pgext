---
title: "pg_dbms_job"
linkTitle: "pg_dbms_job"
description: "Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL"
weight: 9260
categories: ["SIM"]
width: full
---

[**pg_dbms_job**](https://github.com/MigOpsRepos/pg_dbms_job) : Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9260** | {{< badge content="pg_dbms_job" link="https://github.com/MigOpsRepos/pg_dbms_job" >}} | {{< ext "pg_dbms_job" >}} | `1.5` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} {{< ext "oracle_fdw" >}} {{< ext "orafce" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_dbms_job` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5` | {{< bg "18" "pg_dbms_job_18" "green" >}} {{< bg "17" "pg_dbms_job_17" "green" >}} {{< bg "16" "pg_dbms_job_16" "green" >}} {{< bg "15" "pg_dbms_job_15" "green" >}} {{< bg "14" "pg_dbms_job_14" "green" >}} {{< bg "13" "pg_dbms_job_13" "green" >}} | `pg_dbms_job_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "PGDG 1.5" "pg_dbms_job_18 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_17 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_16 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_15 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_14 : BREAK 3" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_13 : BREAK 3" "orange" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "PGDG 1.5" "pg_dbms_job_18 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_17 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_16 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_15 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_14 : BREAK 1" "orange" >}}      |      {{< bg "PGDG 1.5" "pg_dbms_job_13 : BREAK 1" "orange" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_14 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_13 : AVAIL 3" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.5" "pg_dbms_job_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_job : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_18` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.8 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_18` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.7 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_18` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.0 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_job_18` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.0 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_job_18` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_18` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.6 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_17` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_17` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.5 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_job_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.1 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_job_17` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_17-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_job_17-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_17` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pg_dbms_job_17-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_dbms_job_17-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_16` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_16` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.5 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_16` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 26.1 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_job_16` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 26.0 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_job_16` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_16-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_job_16-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_16` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.6 KiB | [pg_dbms_job_16-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_dbms_job_16-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_15` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_dbms_job_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_job_15-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_15` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.2 KiB | [pg_dbms_job_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_job_15-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_15` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [pg_dbms_job_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_job_15-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_15` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [pg_dbms_job_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_job_15-1.5-1.rhel9.aarch64.rpm) |
| `pg_dbms_job_15` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_15-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_job_15-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_15` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.5 KiB | [pg_dbms_job_15-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_dbms_job_15-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_14` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_dbms_job_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.0 KiB | [pg_dbms_job_14-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.4.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.2 KiB | [pg_dbms_job_14-1.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.2.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.2 KiB | [pg_dbms_job_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_job_14-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_14` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [pg_dbms_job_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.6 KiB | [pg_dbms_job_14-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.4.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [pg_dbms_job_14-1.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.2.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_14` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [pg_dbms_job_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_job_14-1.5-1.rhel9.aarch64.rpm) |
| `pg_dbms_job_14` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_14-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_job_14-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_14` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.6 KiB | [pg_dbms_job_14-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_dbms_job_14-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_job_13` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.3 KiB | [pg_dbms_job_13-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 26.0 KiB | [pg_dbms_job_13-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.4.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 25.2 KiB | [pg_dbms_job_13-1.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.2.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 26.2 KiB | [pg_dbms_job_13-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_dbms_job_13-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_13` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.8 KiB | [pg_dbms_job_13-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_13` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 25.6 KiB | [pg_dbms_job_13-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.4.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_13` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 24.8 KiB | [pg_dbms_job_13-1.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.2.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_13` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 25.7 KiB | [pg_dbms_job_13-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_dbms_job_13-1.5-1.rhel9.aarch64.rpm) |
| `pg_dbms_job_13` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pgdg | 26.6 KiB | [pg_dbms_job_13-1.5-5PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_dbms_job_13-1.5-5PGDG.rhel10.x86_64.rpm) |
| `pg_dbms_job_13` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pgdg | 26.6 KiB | [pg_dbms_job_13-1.5-5PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_dbms_job_13-1.5-5PGDG.rhel10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MigOpsRepos/pg_dbms_job" title="Repository" icon="github" subtitle="github.com/MigOpsRepos/pg_dbms_job" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dbms_job;		# install via package name, for the active PG version

pig install pg_dbms_job -v 18;   # install for PG 18
pig install pg_dbms_job -v 17;   # install for PG 17
pig install pg_dbms_job -v 16;   # install for PG 16
pig install pg_dbms_job -v 15;   # install for PG 15
pig install pg_dbms_job -v 14;   # install for PG 14
pig install pg_dbms_job -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dbms_job;
```
