---
title: "babelfishpg_tds"
linkTitle: "babelfishpg_tds"
description: "SQL Server TDS protocol extension"
weight: 9320
categories: ["Sim"]
width: full
---

SQL Server TDS protocol extension

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9320** | {{< badge content="babelfishpg_tds" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tds" "babelfishpg_tds" >}} | `1.0.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "tds_fdw" >}} {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "uuid-ossp" >}} {{< ext "session_variable" >}} {{< ext "jdbc_fdw" >}} {{< ext "db_migrator" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_tds" >}} | `1.0.0` | {{< badge content="18" color="red" alt="babelfishpg-tds*" >}} {{< badge content="17" color="red" alt="babelfishpg-tds*" >}} {{< badge content="16" color="red" alt="babelfishpg-tds*" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-tds*" >}} | `babelfishpg-tds*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_tds" >}} | `1.0.0` | {{< badge content="18" color="red" alt="babelfishpg-tds" >}} {{< badge content="17" color="red" alt="babelfishpg-tds" >}} {{< badge content="16" color="red" alt="babelfishpg-tds" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-tds" >}} | `babelfishpg-tds` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `el8.aarch64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `el9.x86_64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `el9.aarch64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `d12.x86_64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `d12.aarch64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `u22.x86_64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `u22.aarch64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `u24.x86_64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |
|    `u24.aarch64`    |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |    {{< pkg "babelfishpg-tds" >}}     |


## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install babelfishpg_tds; # install by extension name, for the current active PG version
pig ext install babelfishpg_tds; # install via package alias, for the active PG version
pig ext install babelfishpg_tds -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION babelfishpg_tds;
```



--------

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