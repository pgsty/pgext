---
title: "pg_dbms_metadata"
linkTitle: "pg_dbms_metadata"
description: "Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL"
weight: 9240
categories: ["SIM"]
width: full
---

[**pg_dbms_metadata**](https://github.com/HexaCluster/pg_dbms_metadata) : Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9240** | {{< badge content="pg_dbms_metadata" link="https://github.com/HexaCluster/pg_dbms_metadata" >}} | {{< ext "pg_dbms_metadata" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `dbms_metadata` |
|   **See Also**    | {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "pg_dbms_lock" >}} {{< ext "pg_dbms_job" >}} {{< ext "oracle_fdw" >}} {{< ext "session_variable" >}} {{< ext "pg_statement_rollback" >}} {{< ext "ddlx" >}} |

> [!Note] pgdg missing el8.aarch64.pg15


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_dbms_metadata` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.0` | {{< bg "18" "pg_dbms_metadata_18" "green" >}} {{< bg "17" "pg_dbms_metadata_17" "green" >}} {{< bg "16" "pg_dbms_metadata_16" "green" >}} {{< bg "15" "pg_dbms_metadata_15" "green" >}} {{< bg "14" "pg_dbms_metadata_14" "green" >}} | `pg_dbms_metadata_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pg_dbms_metadata_15 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.0" "pg_dbms_metadata_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_metadata : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.8 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.7 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.4 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.3 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_metadata_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_dbms_metadata_18-1.0.0-2PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_metadata_17-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_metadata_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_dbms_metadata_17-1.0.0-2PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_metadata_16-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_metadata_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_dbms_metadata_16-1.0.0-2PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_metadata_15-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_metadata_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_dbms_metadata_15-1.0.0-2PGDG.rhel10.noarch.rpm) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_metadata_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 17.7 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 17.6 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_metadata_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 17.3 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 17.2 KiB | [pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_metadata_14-1.0.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_metadata_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 17.9 KiB | [pg_dbms_metadata_14-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_dbms_metadata_14-1.0.0-2PGDG.rhel10.noarch.rpm) |
| `pg_dbms_metadata_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 17.8 KiB | [pg_dbms_metadata_14-1.0.0-2PGDG.rhel10.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/pg_dbms_metadata_14-1.0.0-2PGDG.rhel10.noarch.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_metadata" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_metadata" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_dbms_metadata;		# install via package name, for the active PG version

pig install pg_dbms_metadata -v 18;   # install for PG 18
pig install pg_dbms_metadata -v 17;   # install for PG 17
pig install pg_dbms_metadata -v 16;   # install for PG 16
pig install pg_dbms_metadata -v 15;   # install for PG 15
pig install pg_dbms_metadata -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_dbms_metadata;
```



## Usage

> [pg_dbms_metadata: Extension to add Oracle DBMS_METADATA compatibility to PostgreSQL](https://github.com/HexaCluster/pg_dbms_metadata)

### Enabling

```sql
CREATE EXTENSION pg_dbms_metadata;
```

### GET_DDL

Extract DDL of database objects. Supported types: TABLE, VIEW, SEQUENCE, PROCEDURE, FUNCTION, TRIGGER, INDEX, CONSTRAINT, CHECK_CONSTRAINT, REF_CONSTRAINT, TYPE, ENUM.

```sql
SELECT dbms_metadata.get_ddl('TABLE', 'employees', 'public');
SELECT dbms_metadata.get_ddl('VIEW', 'active_users', 'myschema');
SELECT dbms_metadata.get_ddl('FUNCTION', 'calculate_tax', 'public');
SELECT dbms_metadata.get_ddl('INDEX', 'idx_name', 'public');

-- Schema is optional; uses search_path when omitted
SELECT dbms_metadata.get_ddl('SEQUENCE', 'my_seq');
```

### GET_DEPENDENT_DDL

Extract DDL of all dependent objects of a specified type for a base object. Supported: SEQUENCE, TRIGGER, CONSTRAINT, REF_CONSTRAINT, INDEX, ENUM.

```sql
SELECT dbms_metadata.get_dependent_ddl('CONSTRAINT', 'employees', 'public');
SELECT dbms_metadata.get_dependent_ddl('INDEX', 'orders', 'sales');
SELECT dbms_metadata.get_dependent_ddl('TRIGGER', 'accounts');
```

### GET_GRANTED_DDL

Extract SQL statements to recreate granted privileges and roles. Supported: ROLE_GRANT.

```sql
SELECT dbms_metadata.get_granted_ddl('ROLE_GRANT', 'user_test');
```

### SET_TRANSFORM_PARAM

Customize DDL output with session-level transform parameters:

```sql
-- Append SQL terminator (;) to each DDL statement
CALL dbms_metadata.set_transform_param('SQLTERMINATOR', true);

-- Exclude constraints from TABLE DDL
CALL dbms_metadata.set_transform_param('CONSTRAINTS', false);

-- Exclude foreign keys from TABLE DDL
CALL dbms_metadata.set_transform_param('REF_CONSTRAINTS', false);

-- Exclude partitioning clauses
CALL dbms_metadata.set_transform_param('PARTITIONING', false);

-- Exclude storage parameters
CALL dbms_metadata.set_transform_param('STORAGE', false);

-- Reset all params to defaults
CALL dbms_metadata.set_transform_param('DEFAULT', true);
```
