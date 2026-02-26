---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server Money Data Type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" "babelfish" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |
|    **Siblings**   | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `babelfish` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfish_18" "red" >}} {{< bg "17" "babelfish_17" "green" >}} {{< bg "16" "babelfish_16" "red" >}} {{< bg "15" "babelfish_15" "red" >}} {{< bg "14" "babelfish_14" "red" >}} {{< bg "13" "babelfish_13" "red" >}} | `babelfish_$v` | `babelfishpg_$v`, `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfishpg-18-babelfish" "red" >}} {{< bg "17" "babelfishpg-17-babelfish" "green" >}} {{< bg "16" "babelfishpg-16-babelfish" "red" >}} {{< bg "15" "babelfishpg-15-babelfish" "red" >}} {{< bg "14" "babelfishpg-14-babelfish" "red" >}} {{< bg "13" "babelfishpg-13-babelfish" "red" >}} | `babelfishpg-$v-babelfish` | `babelfishpg-$v-babelfish`, `libantlr4-runtime413` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-13-babelfish : FORK 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="babelfishpg-17.8-5.5.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg babelfish;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install babelfish;		# install via package name, for the active PG version
pig install babelfishpg_money;		# install by extension name, for the current active PG version

pig install babelfishpg_money -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_money;
```
