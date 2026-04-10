---
title: "pg_liquid"
linkTitle: "pg_liquid"
description: "Liquid-inspired Datalog graph query extension for PostgreSQL"
weight: 2705
categories: ["FEAT"]
width: full
---

[**pg_liquid**](https://github.com/michael-golfi/pg_liquid) : Liquid-inspired Datalog graph query extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2705** | {{< badge content="pg_liquid" link="https://github.com/michael-golfi/pg_liquid" >}} | {{< ext "pg_liquid" >}} | `0.1.7` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `liquid` |
|   **See Also**    | {{< ext "age" >}} {{< ext "jsquery" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_search" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_liquid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "pg_liquid_18" "green" >}} {{< bg "17" "pg_liquid_17" "green" >}} {{< bg "16" "pg_liquid_16" "green" >}} {{< bg "15" "pg_liquid_15" "green" >}} {{< bg "14" "pg_liquid_14" "green" >}} | `pg_liquid_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.7` | {{< bg "18" "postgresql-18-pg-liquid" "green" >}} {{< bg "17" "postgresql-17-pg-liquid" "green" >}} {{< bg "16" "postgresql-16-pg-liquid" "green" >}} {{< bg "15" "postgresql-15-pg-liquid" "green" >}} {{< bg "14" "postgresql-14-pg-liquid" "green" >}} | `postgresql-$v-pg-liquid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "el8.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "el9.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "el9.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "el10.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "el10.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "d12.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "d12.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "d13.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "d13.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "u22.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "u22.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "u24.x86_64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |
| {{< os "u24.aarch64" >}} |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |    {{< bg "MISS" "N/A : MISS 0" "red" >}}     |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/michael-golfi/pg_liquid" title="Repository" icon="github" subtitle="github.com/michael-golfi/pg_liquid" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_liquid-0.1.7.zip" >}}
{{< /cards >}}


```bash
pig build pkg pg_liquid;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_liquid;		# install via package name, for the active PG version

pig install pg_liquid -v 18;   # install for PG 18
pig install pg_liquid -v 17;   # install for PG 17
pig install pg_liquid -v 16;   # install for PG 16
pig install pg_liquid -v 15;   # install for PG 15
pig install pg_liquid -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_liquid;
```
