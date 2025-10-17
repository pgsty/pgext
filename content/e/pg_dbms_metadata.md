---
title: "pg_dbms_metadata"
linkTitle: "pg_dbms_metadata"
description: "Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL"
weight: 9240
categories: ["Sim"]
width: full
---

Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9240** | {{< badge content="pg_dbms_metadata" link="https://github.com/HexaCluster/pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" "pg_dbms_metadata" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "oracle_fdw" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "ddlx" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_dbms_metadata" >}} | `1.0.0` | {{< badge content="18" color="green" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pg_dbms_metadata_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_dbms_metadata_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el8.aarch64`    | {{< pkg "pg_dbms_metadata_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm" >}} |    {{< pkg "pg_dbms_metadata_15" >}}     | {{< pkg "pg_dbms_metadata_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm" >}} |
|    `el9.x86_64`    | {{< pkg "pg_dbms_metadata_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm" >}} |
|    `el9.aarch64`    | {{< pkg "pg_dbms_metadata_18" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_17" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_16" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_15" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm" >}} | {{< pkg "pg_dbms_metadata_14" "1.0.0" "pgdg" "https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm" >}} |
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
| `pg_dbms_metadata_18` | 1.0.0 | `el8.x86_64` | pgdg | 17.8 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | 1.0.0 | `el8.aarch64` | pgdg | 17.7 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | 1.0.0 | `el9.aarch64` | pgdg | 17.3 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_18` | 1.0.0 | `el9.x86_64` | pgdg | 17.4 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_metadata_17` | 1.0.0 | `el8.x86_64` | pgdg | 17.7 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | 1.0.0 | `el8.aarch64` | pgdg | 17.6 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | 1.0.0 | `el9.aarch64` | pgdg | 17.2 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_17` | 1.0.0 | `el9.x86_64` | pgdg | 17.3 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_metadata_16` | 1.0.0 | `el8.x86_64` | pgdg | 17.7 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | 1.0.0 | `el8.aarch64` | pgdg | 17.6 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | 1.0.0 | `el9.aarch64` | pgdg | 17.2 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_16` | 1.0.0 | `el9.x86_64` | pgdg | 17.3 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_metadata_15` | 1.0.0 | `el8.x86_64` | pgdg | 17.7 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_15` | 1.0.0 | `el9.x86_64` | pgdg | 17.3 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_15` | 1.0.0 | `el9.aarch64` | pgdg | 17.2 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_metadata_14` | 1.0.0 | `el8.x86_64` | pgdg | 17.7 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_14` | 1.0.0 | `el8.aarch64` | pgdg | 17.6 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_14` | 1.0.0 | `el9.aarch64` | pgdg | 17.2 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_14` | 1.0.0 | `el9.x86_64` | pgdg | 17.3 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_dbms_metadata_13` | 1.0.0 | `el8.aarch64` | pgdg | 17.6 KiB | [pg_dbms_metadata_13-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_dbms_metadata_13-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_13` | 1.0.0 | `el8.x86_64` | pgdg | 17.7 KiB | [pg_dbms_metadata_13-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_dbms_metadata_13-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_13` | 1.0.0 | `el9.x86_64` | pgdg | 17.3 KiB | [pg_dbms_metadata_13-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_dbms_metadata_13-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_13` | 1.0.0 | `el9.aarch64` | pgdg | 17.2 KiB | [pg_dbms_metadata_13-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_dbms_metadata_13-1.0.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_metadata" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_metadata" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_dbms_metadata; # install by extension name, for the current active PG version
pig ext install pg_dbms_metadata; # install via package alias, for the active PG version
pig ext install pg_dbms_metadata -v 18;   # install for PG 18
pig ext install pg_dbms_metadata -v 17;   # install for PG 17
pig ext install pg_dbms_metadata -v 16;   # install for PG 16
pig ext install pg_dbms_metadata -v 15;   # install for PG 15
pig ext install pg_dbms_metadata -v 14;   # install for PG 14
pig ext install pg_dbms_metadata -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_dbms_metadata;
```

