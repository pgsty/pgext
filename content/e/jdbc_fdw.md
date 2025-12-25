---
title: "jdbc_fdw"
linkTitle: "jdbc_fdw"
description: "foreign-data wrapper for remote servers available over JDBC"
weight: 8530
categories: ["FDW"]
width: full
---

[**jdbc_fdw**](https://github.com/pgspider/jdbc_fdw) : foreign-data wrapper for remote servers available over JDBC


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **8530** | {{< badge content="jdbc_fdw" link="https://github.com/pgspider/jdbc_fdw" >}} | {{< ext "jdbc_fdw" >}} | `0.4.0` | {{< category "FDW" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "wrappers" >}} {{< ext "multicorn" >}} {{< ext "odbc_fdw" >}} {{< ext "oracle_fdw" >}} {{< ext "mysql_fdw" >}} {{< ext "tds_fdw" >}} {{< ext "db2_fdw" >}} {{< ext "postgres_fdw" >}} |

> [!Note] missing el.aarch64


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.4.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `jdbc_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `0.4.0` | {{< bg "18" "jdbc_fdw_18" "red" >}} {{< bg "17" "jdbc_fdw_17" "red" >}} {{< bg "16" "jdbc_fdw_16" "green" >}} {{< bg "15" "jdbc_fdw_15" "green" >}} {{< bg "14" "jdbc_fdw_14" "green" >}} {{< bg "13" "jdbc_fdw_13" "green" >}} | `jdbc_fdw_$v` | `java-11-openjdk-headless` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 0.4.0" "jdbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      | {{< bg "PGDG 0.4.0" "jdbc_fdw_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.4.0" "jdbc_fdw_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |      {{< bg "MISS" "jdbc_fdw : MISS 0" "red" >}}      |


{{< tabs items="PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_16` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.8 KiB | [jdbc_fdw_16-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/jdbc_fdw_16-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_16` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [jdbc_fdw_16-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/jdbc_fdw_16-0.4.0-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_15` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [jdbc_fdw_15-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/jdbc_fdw_15-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_15` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [jdbc_fdw_15-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/jdbc_fdw_15-0.4.0-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_14` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.5 KiB | [jdbc_fdw_14-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/jdbc_fdw_14-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_14` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [jdbc_fdw_14-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/jdbc_fdw_14-0.4.0-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jdbc_fdw_13` | `0.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.0 KiB | [jdbc_fdw_13-0.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/jdbc_fdw_13-0.4.0-1PGDG.rhel8.x86_64.rpm) |
| `jdbc_fdw_13` | `0.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.8 KiB | [jdbc_fdw_13-0.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/jdbc_fdw_13-0.4.0-1PGDG.rhel9.x86_64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgspider/jdbc_fdw" title="Repository" icon="github" subtitle="github.com/pgspider/jdbc_fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install jdbc_fdw;		# install via package name, for the active PG version

pig install jdbc_fdw -v 16;   # install for PG 16
pig install jdbc_fdw -v 15;   # install for PG 15
pig install jdbc_fdw -v 14;   # install for PG 14
pig install jdbc_fdw -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jdbc_fdw;
```
