---
title: "pg_tde"
linkTitle: "pg_tde"
description: "Percona pg_tde access method"
weight: 7500
categories: ["SEC"]
width: full
---

[**pg_tde**](https://github.com/Percona-Lab/pg_tde) : Percona pg_tde access method


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7500** | {{< badge content="pg_tde" link="https://github.com/Percona-Lab/pg_tde" >}} | {{< ext "pg_tde" >}} | `1.0` | {{< category "SEC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgsodium" >}} {{< ext "pgsmcrypto" >}} {{< ext "pgcrypto" >}} {{< ext "anon" >}} {{< ext "pgcryptokey" >}} {{< ext "faker" >}} {{< ext "sslutils" >}} {{< ext "uuid-ossp" >}} |

> [!Note] works on percona postgres fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_tde` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "percona-postgresql18" "red" >}} {{< bg "17" "percona-postgresql17" "green" >}} {{< bg "16" "percona-postgresql16" "red" >}} {{< bg "15" "percona-postgresql15" "red" >}} {{< bg "14" "percona-postgresql14" "red" >}} {{< bg "13" "percona-postgresql13" "red" >}} | `percona-postgresql$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "percona-postgresql-18" "red" >}} {{< bg "17" "percona-postgresql-17" "green" >}} {{< bg "16" "percona-postgresql-16" "red" >}} {{< bg "15" "percona-postgresql-15" "red" >}} {{< bg "14" "percona-postgresql-14" "red" >}} {{< bg "13" "percona-postgresql-13" "red" >}} | `percona-postgresql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "percona-postgresql18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql13 : THROW 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "percona-postgresql-18 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-17 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-16 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-15 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-14 : THROW 0" "red" >}}      |      {{< bg "MISS" "percona-postgresql-13 : THROW 0" "red" >}}      |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Percona-Lab/pg_tde" title="Repository" icon="github" subtitle="github.com/Percona-Lab/pg_tde" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tde-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tde;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tde;		# install via package name, for the active PG version

pig install pg_tde -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_tde';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tde;
```
