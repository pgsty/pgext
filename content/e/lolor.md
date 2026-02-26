---
title: "lolor"
linkTitle: "lolor"
description: "Logical-replication-friendly replacement for PostgreSQL large objects"
weight: 9570
categories: ["ETL"]
width: full
---

[**lolor**](https://github.com/pgEdge/lolor) : Logical-replication-friendly replacement for PostgreSQL large objects


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9570** | {{< badge content="lolor" link="https://github.com/pgEdge/lolor" >}} | {{< ext "lolor" >}} | `1.2.2` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lolor` |
|   **See Also**    | {{< ext "spock" >}} {{< ext "snowflake" >}} |

> [!Note] works on pgedge kernel fork. Requires lolor.node


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `lolor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "lolor_18" "red" >}} {{< bg "17" "lolor_17" "green" >}} {{< bg "16" "lolor_16" "red" >}} {{< bg "15" "lolor_15" "red" >}} {{< bg "14" "lolor_14" "red" >}} {{< bg "13" "lolor_13" "red" >}} | `lolor_$v` | `pgedge_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "postgresql-18-lolor" "red" >}} {{< bg "17" "postgresql-17-lolor" "green" >}} {{< bg "16" "postgresql-16-lolor" "red" >}} {{< bg "15" "postgresql-15-lolor" "red" >}} {{< bg "14" "postgresql-14-lolor" "red" >}} {{< bg "13" "postgresql-13-lolor" "red" >}} | `postgresql-$v-lolor` | `pgedge-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-lolor : FORK 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lolor_17` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 61.6 KiB | [lolor_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lolor_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 60.6 KiB | [lolor_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lolor_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `lolor_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 62.4 KiB | [lolor_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lolor_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 61.7 KiB | [lolor_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lolor_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `lolor_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 62.7 KiB | [lolor_17-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lolor_17-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 62.0 KiB | [lolor_17-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lolor_17-1.2.2-1PIGSTY.el10.aarch64.rpm) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/lolor" title="Repository" icon="github" subtitle="github.com/pgEdge/lolor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lolor-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg lolor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install lolor;		# install via package name, for the active PG version

pig install lolor -v 17;   # install for PG 17
pig install lolor -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lolor;
```
