---
title: "pg_dbms_lock"
linkTitle: "pg_dbms_lock"
description: "Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL"
weight: 9250
categories: ["SIM"]
width: full
---

Extension to add Oracle DBMS_LOCK full compatibility to PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9250** | {{< badge content="pg_dbms_lock" link="https://github.com/HexaCluster/pg_dbms_lock" >}} | {{< ext "pg_dbms_lock" >}} | `1.0` | {{< category "SIM" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "orafce" >}} {{< ext "session_variable" >}} {{< ext "pg_dbms_metadata" >}} {{< ext "pg_dbms_job" >}} {{< ext "oracle_fdw" >}} {{< ext "pgtt" >}} {{< ext "pg_statement_rollback" >}} {{< ext "mysql_fdw" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pg_dbms_lock" >}} | `1.0` | {{< bg "18" "pg_dbms_lock_18" "green" >}} {{< bg "17" "pg_dbms_lock_17" "green" >}} {{< bg "16" "pg_dbms_lock_16" "green" >}} {{< bg "15" "pg_dbms_lock_15" "green" >}} {{< bg "14" "pg_dbms_lock_14" "green" >}} | `pg_dbms_lock_$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.0" "pg_dbms_lock_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_14 : AVAIL 1" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.0" "pg_dbms_lock_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_14 : AVAIL 1" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.0" "pg_dbms_lock_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_14 : AVAIL 1" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.0" "pg_dbms_lock_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0" "pg_dbms_lock_14 : AVAIL 1" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_dbms_lock : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_18` | 1.0 | `el8.x86_64` | pgdg | 12.7 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_18` | 1.0 | `el8.aarch64` | pgdg | 12.7 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_18` | 1.0 | `el9.x86_64` | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_18` | 1.0 | `el9.aarch64` | pgdg | 12.6 KiB | [pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_dbms_lock_18-1.0-3PGDG.rhel9.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_17` | 1.0 | `el8.x86_64` | pgdg | 12.5 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_17` | 1.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_dbms_lock_17-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_17` | 1.0 | `el9.x86_64` | pgdg | 12.4 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_17` | 1.0 | `el9.aarch64` | pgdg | 12.4 KiB | [pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_dbms_lock_17-1.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_16` | 1.0 | `el8.x86_64` | pgdg | 12.5 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_16` | 1.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_dbms_lock_16-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_16` | 1.0 | `el9.x86_64` | pgdg | 12.4 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_16` | 1.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_dbms_lock_16-1.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_15` | 1.0 | `el8.x86_64` | pgdg | 12.5 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_15` | 1.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_dbms_lock_15-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_15` | 1.0 | `el9.x86_64` | pgdg | 12.4 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_15` | 1.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_dbms_lock_15-1.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_dbms_lock_14` | 1.0 | `el8.x86_64` | pgdg | 12.5 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_14` | 1.0 | `el8.aarch64` | pgdg | 12.5 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_dbms_lock_14-1.0-1PGDG.rhel8.noarch.rpm) |
| `pg_dbms_lock_14` | 1.0 | `el9.x86_64` | pgdg | 12.4 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm) |
| `pg_dbms_lock_14` | 1.0 | `el9.aarch64` | pgdg | 12.2 KiB | [pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_dbms_lock_14-1.0-1PGDG.rhel9.noarch.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HexaCluster/pg_dbms_lock" title="Repository" icon="github" subtitle="github.com/HexaCluster/pg_dbms_lock" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_dbms_lock; # install by extension name, for the current active PG version
pig ext install pg_dbms_lock; # install via package alias, for the active PG version
pig ext install pg_dbms_lock -v 18;   # install for PG 18
pig ext install pg_dbms_lock -v 17;   # install for PG 17
pig ext install pg_dbms_lock -v 16;   # install for PG 16
pig ext install pg_dbms_lock -v 15;   # install for PG 15
pig ext install pg_dbms_lock -v 14;   # install for PG 14
pig ext install pg_dbms_lock -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_dbms_lock;
```

