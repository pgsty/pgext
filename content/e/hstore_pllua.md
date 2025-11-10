---
title: "hstore_pllua"
linkTitle: "hstore_pllua"
description: "Hstore transform for Lua"
weight: 3021
categories: ["LANG"]
width: full
---

[**pllua**](https://github.com/pllua/pllua) : Hstore transform for Lua


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3021** | {{< badge content="hstore_pllua" link="https://github.com/pllua/pllua" >}} | {{< ext "hstore_pllua" "pllua" >}} | `2.0.12` | {{< category "LANG" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "hstore" >}} {{< ext "pllua" >}} |
|   **See Also**    | {{< ext "hstore_plperl" >}} {{< ext "hstore_plperlu" >}} {{< ext "hstore_plpython3u" >}} {{< ext "plpgsql" >}} |
|    **Siblings**   | {{< ext "pllua" >}} {{< ext "plluau" >}} {{< ext "hstore_plluau" >}} |

> [!Note] missing pg12-15 on el.aarch64


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pllua` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2.0.12` | {{< bg "18" "postgresql-18-pllua" "green" >}} {{< bg "17" "postgresql-17-pllua" "green" >}} {{< bg "16" "postgresql-16-pllua" "green" >}} {{< bg "15" "postgresql-15-pllua" "green" >}} {{< bg "14" "postgresql-14-pllua" "green" >}} {{< bg "13" "postgresql-13-pllua" "green" >}} | `postgresql-$v-pllua` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.9" "pllua_13 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.11" "pllua_13 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} |      {{< bg "MISS" "pllua_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pllua_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_13 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pllua_18 : MISS 0" "red" >}}      | {{< bg "PGDG 2.0.12" "pllua_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "pllua_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2.0.12" "postgresql-18-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-17-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-16-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-15-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-14-pllua : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.0.12" "postgresql-13-pllua : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pllua/pllua" title="Repository" icon="github" subtitle="github.com/pllua/pllua" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pllua;		# install via package name, for the active PG version
pig install hstore_pllua;		# install by extension name, for the current active PG version

pig install hstore_pllua -v 18;   # install for PG 18
pig install hstore_pllua -v 17;   # install for PG 17
pig install hstore_pllua -v 16;   # install for PG 16
pig install hstore_pllua -v 15;   # install for PG 15
pig install hstore_pllua -v 14;   # install for PG 14
pig install hstore_pllua -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION hstore_pllua CASCADE; -- requires hstore, pllua
```
