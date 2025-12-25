---
title: "babelfishpg_common"
linkTitle: "babelfishpg_common"
description: "SQL Server Transact SQL Datatype Support"
weight: 9300
categories: ["SIM"]
width: full
---

[**babelfishpg_common**](https://babelfishpg.org/) : SQL Server Transact SQL Datatype Support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9300** | {{< badge content="babelfishpg_common" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_common" >}} | `3.3.3` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "tds_fdw" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "uuid-ossp" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.3.3` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `babelfishpg_common` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.3.3` | {{< bg "18" "babelfishpg-common" "red" >}} {{< bg "17" "babelfishpg-common" "red" >}} {{< bg "16" "babelfishpg-common" "red" >}} {{< bg "15" "babelfishpg-common" "green" >}} {{< bg "14" "babelfishpg-common" "red" >}} {{< bg "13" "babelfishpg-common" "red" >}} | `babelfishpg-common` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.3.3` | {{< bg "18" "babelfishpg-common" "red" >}} {{< bg "17" "babelfishpg-common" "red" >}} {{< bg "16" "babelfishpg-common" "red" >}} {{< bg "15" "babelfishpg-common" "green" >}} {{< bg "14" "babelfishpg-common" "red" >}} {{< bg "13" "babelfishpg-common" "red" >}} | `babelfishpg-common` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : FORK 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install babelfishpg_common;		# install via package name, for the active PG version

pig install babelfishpg_common -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_common;
```
