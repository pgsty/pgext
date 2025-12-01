---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["SIM"]
width: full
---

[**babelfishpg_money**](https://babelfishpg.org/) : SQL Server Money Data Type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `babelfishpg_money` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfishpg-money*" "red" >}} {{< bg "17" "babelfishpg-money*" "red" >}} {{< bg "16" "babelfishpg-money*" "red" >}} {{< bg "15" "babelfishpg-money*" "green" >}} {{< bg "14" "babelfishpg-money*" "red" >}} {{< bg "13" "babelfishpg-money*" "red" >}} | `babelfishpg-money*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfishpg-money" "red" >}} {{< bg "17" "babelfishpg-money" "red" >}} {{< bg "16" "babelfishpg-money" "red" >}} {{< bg "15" "babelfishpg-money" "green" >}} {{< bg "14" "babelfishpg-money" "red" >}} {{< bg "13" "babelfishpg-money" "red" >}} | `babelfishpg-money` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : FORK 0" "red" >}}      |


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
pig install babelfishpg_money;		# install via package name, for the active PG version

pig install babelfishpg_money -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_money;
```
