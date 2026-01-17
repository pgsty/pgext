---
title: "mobilitydb_datagen"
linkTitle: "mobilitydb_datagen"
description: "MobilityDB random data generator functions"
weight: 1651
categories: ["GIS"]
width: full
---

[**mobilitydb**](https://github.com/MobilityDB/MobilityDB) : MobilityDB random data generator functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1651** | {{< badge content="mobilitydb_datagen" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< ext "mobilitydb_datagen" "mobilitydb" >}} | `1.3.0` | {{< category "GIS" >}} | {{< license "GPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "mobilitydb" >}} |
|   **See Also**    | {{< ext "mobilitydb" >}} {{< ext "postgis" >}} {{< ext "timescaledb" >}} {{< ext "pgrouting" >}} |
|    **Siblings**   | {{< ext "mobilitydb" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `mobilitydb` | `mobilitydb` |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.3.0` | {{< bg "18" "postgresql-18-mobilitydb" "green" >}} {{< bg "17" "postgresql-17-mobilitydb" "green" >}} {{< bg "16" "postgresql-16-mobilitydb" "green" >}} {{< bg "15" "postgresql-15-mobilitydb" "green" >}} {{< bg "14" "postgresql-14-mobilitydb" "green" >}} {{< bg "13" "postgresql-13-mobilitydb" "green" >}} | `postgresql-$v-mobilitydb` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |      {{< bg "MISS" "mobilitydb : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-mobilitydb : MISS 0" "red" >}}      | {{< bg "PGDG 1.2.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.3.0" "postgresql-18-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-17-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-16-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-15-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-14-mobilitydb : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "postgresql-13-mobilitydb : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MobilityDB/MobilityDB" title="Repository" icon="github" subtitle="github.com/MobilityDB/MobilityDB" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install mobilitydb;		# install via package name, for the active PG version
pig install mobilitydb_datagen;		# install by extension name, for the current active PG version

pig install mobilitydb_datagen -v 18;   # install for PG 18
pig install mobilitydb_datagen -v 17;   # install for PG 17
pig install mobilitydb_datagen -v 16;   # install for PG 16
pig install mobilitydb_datagen -v 15;   # install for PG 15
pig install mobilitydb_datagen -v 14;   # install for PG 14
pig install mobilitydb_datagen -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mobilitydb_datagen CASCADE; -- requires mobilitydb
```
