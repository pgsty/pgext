---
title: "pgproto"
linkTitle: "pgproto"
description: "Native Protobuf parsing, mutation, indexing, and JSON conversion support"
weight: 4130
categories: ["UTIL"]
width: full
---

[**pgproto**](https://github.com/Apaezmx/pgproto) : Native Protobuf parsing, mutation, indexing, and JSON conversion support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4130** | {{< badge content="pgproto" link="https://github.com/Apaezmx/pgproto" >}} | {{< ext "pgproto" >}} | `0.2.1` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_protobuf" >}} {{< ext "pg_jsonschema" >}} {{< ext "pg_csv" >}} |

> [!Note] Extension SQL default_version is 1.0 while the packaged upstream release version is 0.2.1.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgproto` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.1` | {{< bg "18" "pgproto_18" "green" >}} {{< bg "17" "pgproto_17" "green" >}} {{< bg "16" "pgproto_16" "green" >}} {{< bg "15" "pgproto_15" "green" >}} {{< bg "14" "pgproto_14" "green" >}} | `pgproto_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.1` | {{< bg "18" "postgresql-18-pgproto" "green" >}} {{< bg "17" "postgresql-17-pgproto" "green" >}} {{< bg "16" "postgresql-16-pgproto" "green" >}} {{< bg "15" "postgresql-15-pgproto" "green" >}} {{< bg "14" "postgresql-14-pgproto" "green" >}} | `postgresql-$v-pgproto` | - |


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
{{< card link="https://github.com/Apaezmx/pgproto" title="Repository" icon="github" subtitle="github.com/Apaezmx/pgproto" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgproto-0.2.1-vendor.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgproto;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgproto;		# install via package name, for the active PG version

pig install pgproto -v 18;   # install for PG 18
pig install pgproto -v 17;   # install for PG 17
pig install pgproto -v 16;   # install for PG 16
pig install pgproto -v 15;   # install for PG 15
pig install pgproto -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgproto;
```
