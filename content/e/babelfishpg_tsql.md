---
title: "babelfishpg_tsql"
linkTitle: "babelfishpg_tsql"
description: "SQL Server Transact SQL compatibility"
weight: 9310
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server Transact SQL compatibility


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9310** | {{< badge content="babelfishpg_tsql" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tsql" "babelfish" >}} | `5.5.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_common" >}} {{< ext "uuid-ossp" >}} |
|    **Need By**    | {{< ext "babelfishpg_tds" >}} |
|   **See Also**    | {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "tds_fdw" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "db_migrator" >}} |
|    **Siblings**   | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.5.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `babelfish` | `babelfishpg_common`, `uuid-ossp` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.5.0` | {{< bg "18" "babelfish_18" "red" >}} {{< bg "17" "babelfish_17" "green" >}} {{< bg "16" "babelfish_16" "red" >}} {{< bg "15" "babelfish_15" "red" >}} {{< bg "14" "babelfish_14" "red" >}} | `babelfish_$v` | `babelfishpg_$v`, `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.5.0` | {{< bg "18" "babelfishpg-18-babelfish" "red" >}} {{< bg "17" "babelfishpg-17-babelfish" "green" >}} {{< bg "16" "babelfishpg-16-babelfish" "red" >}} {{< bg "15" "babelfishpg-15-babelfish" "red" >}} {{< bg "14" "babelfishpg-14-babelfish" "red" >}} | `babelfishpg-$v-babelfish` | `babelfishpg-$v`, `libantlr4-runtime413` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfish_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfish_17 : FORK 1" >}}      |      {{< bg "MISS" "babelfish_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfish_14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-18-babelfish : FORK 0" "red" >}}      |      {{< bg "PIGSTY 5.5.0" "babelfishpg-17-babelfish : FORK 1" >}}      |      {{< bg "MISS" "babelfishpg-16-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-15-babelfish : FORK 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-14-babelfish : FORK 0" "red" >}}      |


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
pig install babelfishpg_tsql;		# install by extension name, for the current active PG version

pig install babelfishpg_tsql -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_tsql CASCADE; -- requires babelfishpg_common, uuid-ossp
```



## Usage

> [babelfishpg_tsql: SQL Server Transact SQL compatibility](https://babelfishpg.org/)

The `babelfishpg_tsql` extension provides Microsoft SQL Server Transact-SQL (T-SQL) compatibility for PostgreSQL as part of the Babelfish project. Applications written for SQL Server can connect and run queries against PostgreSQL with minimal changes.

### Enabling

```sql
CREATE EXTENSION babelfishpg_tsql;
```

### Key Features

- **T-SQL Language Support**: Understands T-SQL syntax including stored procedures, functions, triggers, and batches
- **SQL Server Wire Protocol**: Applications can connect using the TDS (Tabular Data Stream) protocol on port 1433
- **System Procedures**: Common `sp_` system stored procedures are available
- **System Views**: SQL Server catalog views (e.g., `sys.tables`, `sys.columns`, `sys.objects`)
- **Multi-Database Semantics**: Supports SQL Server-style database/schema separation

### Supported T-SQL Features

- `BEGIN...END` blocks, `IF...ELSE`, `WHILE` loops
- `TRY...CATCH` error handling
- Temporary tables (`#temp`, `##global_temp`)
- Table variables (`DECLARE @t TABLE (...)`)
- `IDENTITY` columns and `@@IDENTITY` / `SCOPE_IDENTITY()`
- `TOP` clause, `OUTPUT` clause
- `MERGE` statements
- Common Table Expressions (CTEs)
- Cross-database queries within the same instance
- `EXEC` / `EXECUTE` for dynamic SQL
- SQL Server-style string concatenation and NULL handling
- `PRINT` and `RAISERROR` statements

### Connecting via TDS Protocol

Applications can connect using SQL Server drivers (JDBC, ODBC, ADO.NET) to the TDS listener port (default 1433):

```
Server: hostname
Port: 1433
Database: mydb
```

### Notes

- Requires `babelfishpg_common` extension
- Part of the Babelfish for PostgreSQL project (Apache 2.0 / PostgreSQL license)
- Not all T-SQL features are supported; check the Babelfish compatibility reference
