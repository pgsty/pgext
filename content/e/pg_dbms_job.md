---
title: "pg_dbms_job"
linkTitle: "pg_dbms_job"
description: "Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL"
weight: 9260
categories: ["Sim"]
width: full
---

Extension to add Oracle DBMS_JOB full compatibility to PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9260** | {{< badge content="pg_dbms_job" link="https://github.com/MigOpsRepos/pg_dbms_job" >}} | {{< ext "pg_dbms_job" "pg_dbms_job" >}} | `1.5` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_cron" >}} {{< ext "pg_task" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pgagent" >}} {{< ext "pg_jobmon" >}} {{< ext "oracle_fdw" >}} {{< ext "orafce" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_dbms_job" >}} | `1.5` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_dbms_job_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_dbms_job_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_17" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_16" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_15" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_job_15-1.5-1.rhel8.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_14" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.5-1.rhel8.x86_64.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_dbms_job_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_17" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_16" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_15" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_job_15-1.5-1.rhel8.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_14" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_job_14-1.5-1.rhel8.aarch64.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_dbms_job_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_17" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_16" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_15" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_job_15-1.5-1.rhel9.x86_64.rpm" >}} | {{< pkg "pg_dbms_job_14" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.5-1.rhel9.x86_64.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_dbms_job_18" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_17" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_16" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_15" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_job_15-1.5-1.rhel9.aarch64.rpm" >}} | {{< pkg "pg_dbms_job_14" "1.5" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_job_14-1.5-1.rhel9.aarch64.rpm" >}} |
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
| `pg_dbms_job_18` | 1.5 | `el8.aarch64` | pgdg | 26.7 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_18` | 1.5 | `el8.x86_64` | pgdg | 26.8 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_18` | 1.5 | `el9.aarch64` | pgdg | 26.0 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_job_18-1.5-5PGDG.rhel9.aarch64.rpm) |
| `pg_dbms_job_18` | 1.5 | `el9.x86_64` | pgdg | 26.0 KiB | [pg_dbms_job_18-1.5-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_job_18-1.5-5PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_job_17` | 1.5 | `el8.x86_64` | pgdg | 26.6 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_17` | 1.5 | `el8.aarch64` | pgdg | 26.5 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_17` | 1.5 | `el9.x86_64` | pgdg | 26.1 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_job_17-1.5-3PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_job_17` | 1.5 | `el9.aarch64` | pgdg | 26.1 KiB | [pg_dbms_job_17-1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_job_17-1.5-3PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_job_16` | 1.5 | `el8.x86_64` | pgdg | 26.6 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel8.x86_64.rpm) |
| `pg_dbms_job_16` | 1.5 | `el8.aarch64` | pgdg | 26.5 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel8.aarch64.rpm) |
| `pg_dbms_job_16` | 1.5 | `el9.x86_64` | pgdg | 26.1 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_job_16-1.5-3PGDG.rhel9.x86_64.rpm) |
| `pg_dbms_job_16` | 1.5 | `el9.aarch64` | pgdg | 26.0 KiB | [pg_dbms_job_16-1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_job_16-1.5-3PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_job_15` | 1.5 | `el8.aarch64` | pgdg | 26.2 KiB | [pg_dbms_job_15-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_job_15-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_15` | 1.5 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_dbms_job_15-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_job_15-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_15` | 1.5 | `el9.x86_64` | pgdg | 25.8 KiB | [pg_dbms_job_15-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_job_15-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_15` | 1.5 | `el9.aarch64` | pgdg | 25.7 KiB | [pg_dbms_job_15-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_job_15-1.5-1.rhel9.aarch64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_job_14` | 1.5 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_dbms_job_14-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | 1.5 | `el8.aarch64` | pgdg | 26.2 KiB | [pg_dbms_job_14-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_job_14-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_14` | 1.4.0 | `el8.x86_64` | pgdg | 26.0 KiB | [pg_dbms_job_14-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.4.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | 1.2.0 | `el8.x86_64` | pgdg | 25.2 KiB | [pg_dbms_job_14-1.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_job_14-1.2.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_14` | 1.5 | `el9.aarch64` | pgdg | 25.7 KiB | [pg_dbms_job_14-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_job_14-1.5-1.rhel9.aarch64.rpm) |
| `pg_dbms_job_14` | 1.5 | `el9.x86_64` | pgdg | 25.8 KiB | [pg_dbms_job_14-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_14` | 1.4.0 | `el9.x86_64` | pgdg | 25.6 KiB | [pg_dbms_job_14-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.4.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_14` | 1.2.0 | `el9.x86_64` | pgdg | 24.8 KiB | [pg_dbms_job_14-1.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_job_14-1.2.0-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_job_13` | 1.5 | `el8.aarch64` | pgdg | 26.2 KiB | [pg_dbms_job_13-1.5-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_dbms_job_13-1.5-1.rhel8.aarch64.rpm) |
| `pg_dbms_job_13` | 1.5 | `el8.x86_64` | pgdg | 26.3 KiB | [pg_dbms_job_13-1.5-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.5-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | 1.4.0 | `el8.x86_64` | pgdg | 26.0 KiB | [pg_dbms_job_13-1.4.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.4.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | 1.2.0 | `el8.x86_64` | pgdg | 25.2 KiB | [pg_dbms_job_13-1.2.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_job_13-1.2.0-1.rhel8.x86_64.rpm) |
| `pg_dbms_job_13` | 1.5 | `el9.x86_64` | pgdg | 25.8 KiB | [pg_dbms_job_13-1.5-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.5-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_13` | 1.5 | `el9.aarch64` | pgdg | 25.7 KiB | [pg_dbms_job_13-1.5-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_dbms_job_13-1.5-1.rhel9.aarch64.rpm) |
| `pg_dbms_job_13` | 1.4.0 | `el9.x86_64` | pgdg | 25.6 KiB | [pg_dbms_job_13-1.4.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.4.0-1.rhel9.x86_64.rpm) |
| `pg_dbms_job_13` | 1.2.0 | `el9.x86_64` | pgdg | 24.8 KiB | [pg_dbms_job_13-1.2.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_job_13-1.2.0-1.rhel9.x86_64.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MigOpsRepos/pg_dbms_job" title="Repository" icon="github" subtitle="github.com/MigOpsRepos/pg_dbms_job" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_dbms_job; # install by extension name, for the current active PG version
pig ext install pg_dbms_job; # install via package alias, for the active PG version
pig ext install pg_dbms_job -v 18;   # install for PG 18
pig ext install pg_dbms_job -v 17;   # install for PG 17
pig ext install pg_dbms_job -v 16;   # install for PG 16
pig ext install pg_dbms_job -v 15;   # install for PG 15
pig ext install pg_dbms_job -v 14;   # install for PG 14
pig ext install pg_dbms_job -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_dbms_job;
```

