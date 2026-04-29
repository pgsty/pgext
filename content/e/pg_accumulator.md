---
title: "pg_accumulator"
linkTitle: "pg_accumulator"
description: "Accumulation registers for balance and turnover tracking in PostgreSQL"
weight: 4845
categories: ["FUNC"]
width: full
---

[**pg_accumulator**](https://github.com/Treedo/pg_accumulator) : Accumulation registers for balance and turnover tracking in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4845** | {{< badge content="pg_accumulator" link="https://github.com/Treedo/pg_accumulator" >}} | {{< ext "pg_accumulator" >}} | `1.1.3` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `accum` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "financial" >}} {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_accumulator` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "pg_accumulator_18" "green" >}} {{< bg "17" "pg_accumulator_17" "green" >}} {{< bg "16" "pg_accumulator_16" "green" >}} {{< bg "15" "pg_accumulator_15" "green" >}} {{< bg "14" "pg_accumulator_14" "green" >}} | `pg_accumulator_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.3` | {{< bg "18" "postgresql-18-pg-accumulator" "green" >}} {{< bg "17" "postgresql-17-pg-accumulator" "green" >}} {{< bg "16" "postgresql-16-pg-accumulator" "green" >}} {{< bg "15" "postgresql-15-pg-accumulator" "green" >}} {{< bg "14" "postgresql-14-pg-accumulator" "green" >}} | `postgresql-$v-pg-accumulator` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pg_accumulator_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_accumulator_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-accumulator : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-accumulator : MISS 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Treedo/pg_accumulator" title="Repository" icon="github" subtitle="github.com/Treedo/pg_accumulator" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_accumulator-1.1.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_accumulator;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_accumulator;		# install via package name, for the active PG version

pig install pg_accumulator -v 18;   # install for PG 18
pig install pg_accumulator -v 17;   # install for PG 17
pig install pg_accumulator -v 16;   # install for PG 16
pig install pg_accumulator -v 15;   # install for PG 15
pig install pg_accumulator -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_accumulator CASCADE; -- requires plpgsql
```
