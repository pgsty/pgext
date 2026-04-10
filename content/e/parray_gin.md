---
title: "parray_gin"
linkTitle: "parray_gin"
description: "GIN index operator class and partial-match operators for text arrays"
weight: 4860
categories: ["FUNC"]
width: full
---

[**parray_gin**](https://github.com/theirix/parray_gin) : GIN index operator class and partial-match operators for text arrays


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4860** | {{< badge content="parray_gin" link="https://github.com/theirix/parray_gin" >}} | {{< ext "parray_gin" >}} | `1.4.0` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "intarray" >}} {{< ext "btree_gin" >}} {{< ext "btree_gist" >}} {{< ext "pg_trgm" >}} {{< ext "smlar" >}} {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "arraymath" >}} |

> [!Note] PGXN dist name and PostgreSQL extension name are both parray_gin; Pigsty packages are built for PG 14-18.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `parray_gin` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.0` | {{< bg "18" "parray_gin_18" "green" >}} {{< bg "17" "parray_gin_17" "green" >}} {{< bg "16" "parray_gin_16" "green" >}} {{< bg "15" "parray_gin_15" "green" >}} {{< bg "14" "parray_gin_14" "green" >}} | `parray_gin_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.4.0` | {{< bg "18" "postgresql-18-parray-gin" "green" >}} {{< bg "17" "postgresql-17-parray-gin" "green" >}} {{< bg "16" "postgresql-16-parray-gin" "green" >}} {{< bg "15" "postgresql-15-parray-gin" "green" >}} {{< bg "14" "postgresql-14-parray-gin" "green" >}} | `postgresql-$v-parray-gin` | - |


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
{{< card link="https://github.com/theirix/parray_gin" title="Repository" icon="github" subtitle="github.com/theirix/parray_gin" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="parray_gin-1.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg parray_gin;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install parray_gin;		# install via package name, for the active PG version

pig install parray_gin -v 18;   # install for PG 18
pig install parray_gin -v 17;   # install for PG 17
pig install parray_gin -v 16;   # install for PG 16
pig install parray_gin -v 15;   # install for PG 15
pig install parray_gin -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION parray_gin;
```
