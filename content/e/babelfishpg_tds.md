---
title: "babelfishpg_tds"
linkTitle: "babelfishpg_tds"
description: "SQL Server TDS protocol extension"
weight: 9320
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server TDS protocol extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9320** | {{< badge content="babelfishpg_tds" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tds" "babelfish" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "tds_fdw" >}} {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "uuid-ossp" >}} {{< ext "session_variable" >}} {{< ext "jdbc_fdw" >}} {{< ext "db_migrator" >}} |
|    **Siblings**   | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_money" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `babelfish` | `babelfishpg_tsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "babelfish_18" "red" >}} {{< bg "17" "babelfish_17" "green" >}} {{< bg "16" "babelfish_16" "red" >}} {{< bg "15" "babelfish_15" "red" >}} {{< bg "14" "babelfish_14" "red" >}} {{< bg "13" "babelfish_13" "red" >}} | `babelfish_$v` | `babelfishpg_$v`, `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "babelfishpg-18-babelfish" "red" >}} {{< bg "17" "babelfishpg-17-babelfish" "green" >}} {{< bg "16" "babelfishpg-16-babelfish" "red" >}} {{< bg "15" "babelfishpg-15-babelfish" "red" >}} {{< bg "14" "babelfishpg-14-babelfish" "red" >}} {{< bg "13" "babelfishpg-13-babelfish" "red" >}} | `babelfishpg-$v-babelfish` | `babelfishpg-$v-babelfish`, `libantlr4-runtime413` |


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
pig install babelfishpg_tds;		# install by extension name, for the current active PG version

pig install babelfishpg_tds -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'babelfishpg_tds';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_tds CASCADE; -- requires babelfishpg_tsql
```


## Usage

Install `go-sqlcmd`:

```bash
curl -LO https://github.com/microsoft/go-sqlcmd/releases/download/v1.4.0/sqlcmd-v1.4.0-linux-amd64.tar.bz2
tar xjvf sqlcmd-v1.4.0-linux-amd64.tar.bz2
sudo mv sqlcmd* /usr/bin/
```

Try go-sqlcmd

```bash
$ sqlcmd -S 10.10.10.10,1433 -U dbuser_mssql -P DBUser.MSSQL
1> select @@version
2> go
version                                                                                                                                                                                                                                                         
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Babelfish for PostgreSQL with SQL Server Compatibility - 12.0.2000.8
Oct 22 2023 17:48:32
Copyright (c) Amazon Web Services
PostgreSQL 15.4 (EL 1:15.4.wiltondb3.3_2-2.el8) on x86_64-redhat-linux-gnu (Babelfish 3.3.0)                                        

(1 row affected)
```

Access pigsty exposed primary/replica service port

```bash 
sqlcmd -S 10.10.10.11,5433 -U dbuser_mssql -P DBUser.MSSQL

sqlcmd -S 10.10.10.11,5434 -U dbuser_mssql -P DBUser.MSSQL
```