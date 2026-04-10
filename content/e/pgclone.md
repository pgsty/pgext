---
title: "pgclone"
linkTitle: "pgclone"
description: "Clone PostgreSQL databases, schemas, tables, and functions across environments"
weight: 9590
categories: ["ETL"]
width: full
---

[**pgclone**](https://github.com/valehdba/pgclone) : Clone PostgreSQL databases, schemas, tables, and functions across environments


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9590** | {{< badge content="pgclone" link="https://github.com/valehdba/pgclone" >}} | {{< ext "pgclone" >}} | `2.2.0` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "db_migrator" >}} {{< ext "pglogical" >}} {{< ext "repmgr" >}} {{< ext "pgactive" >}} |

> [!Note] shared_preload_libraries = pgclone is required for async/progress features; synchronous clone functions work without it.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgclone` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "pgclone_18" "green" >}} {{< bg "17" "pgclone_17" "green" >}} {{< bg "16" "pgclone_16" "green" >}} {{< bg "15" "pgclone_15" "green" >}} {{< bg "14" "pgclone_14" "green" >}} | `pgclone_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.2.0` | {{< bg "18" "postgresql-18-pgclone" "green" >}} {{< bg "17" "postgresql-17-pgclone" "green" >}} {{< bg "16" "postgresql-16-pgclone" "green" >}} {{< bg "15" "postgresql-15-pgclone" "green" >}} {{< bg "14" "postgresql-14-pgclone" "green" >}} | `postgresql-$v-pgclone` | - |


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
{{< card link="https://github.com/valehdba/pgclone" title="Repository" icon="github" subtitle="github.com/valehdba/pgclone" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgclone-2.2.0.zip" >}}
{{< /cards >}}


```bash
pig build pkg pgclone;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgclone;		# install via package name, for the active PG version

pig install pgclone -v 18;   # install for PG 18
pig install pgclone -v 17;   # install for PG 17
pig install pgclone -v 16;   # install for PG 16
pig install pgclone -v 15;   # install for PG 15
pig install pgclone -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgclone';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgclone;
```
