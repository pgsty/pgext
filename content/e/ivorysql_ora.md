---
title: "ivorysql_ora"
linkTitle: "ivorysql_ora"
description: "Oracle Compatible extension on Postgres Database"
weight: 9140
categories: ["SIM"]
width: full
---

[**ivorysql**](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora) : Oracle Compatible extension on Postgres Database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9140** | {{< badge content="ivorysql_ora" link="https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" >}} | {{< ext "ivorysql_ora" "ivorysql" >}} | `1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `sys` |
|    **Siblings**   | {{< ext "ora_btree_gin" >}} {{< ext "ora_btree_gist" >}} {{< ext "pg_get_functiondef" >}} {{< ext "plisql" >}} {{< ext "gb18030_2022" >}} |

> [!Note] from contrib/ivorysql_ora/ivorysql_ora.control and package metadata


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `ivorysql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.1` | {{< bg "18" "ivorysql5" "green" >}} {{< bg "17" "ivorysql5" "red" >}} {{< bg "16" "ivorysql5" "red" >}} {{< bg "15" "ivorysql5" "red" >}} {{< bg "14" "ivorysql5" "red" >}} {{< bg "13" "ivorysql5" "red" >}} | `ivorysql5` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.1` | {{< bg "18" "ivorysql-5" "green" >}} {{< bg "17" "ivorysql-5" "red" >}} {{< bg "16" "ivorysql-5" "red" >}} {{< bg "15" "ivorysql-5" "red" >}} {{< bg "14" "ivorysql-5" "red" >}} {{< bg "13" "ivorysql-5" "red" >}} | `ivorysql-5` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql5 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-5 : FORK 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" title="Repository" icon="github" subtitle="github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ivorysql;		# install via package name, for the active PG version
pig install ivorysql_ora;		# install by extension name, for the current active PG version

pig install ivorysql_ora -v 18;   # install for PG 18

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ivorysql_ora;
```
