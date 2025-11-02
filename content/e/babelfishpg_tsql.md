---
title: "babelfishpg_tsql"
linkTitle: "babelfishpg_tsql"
description: "SQL Server Transact SQL compatibility"
weight: 9310
categories: ["SIM"]
width: full
---

SQL Server Transact SQL compatibility


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9310** | {{< badge content="babelfishpg_tsql" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tsql" >}} | `3.3.1` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_common" >}} {{< ext "uuid-ossp" >}} |
|    **Need By**    | {{< ext "babelfishpg_tds" >}} |
|   **See Also**    | {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "tds_fdw" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "db_migrator" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_tsql" >}} | `3.3.1` | {{< bg "18" "babelfishpg-tsql*" "red" >}} {{< bg "17" "babelfishpg-tsql*" "red" >}} {{< bg "16" "babelfishpg-tsql*" "red" >}} {{< bg "15" "babelfishpg-tsql*" "green" >}} {{< bg "14" "babelfishpg-tsql*" "red" >}} {{< bg "13" "babelfishpg-tsql*" "red" >}} | `babelfishpg-tsql*` | `babelfishpg-common`, `libantlr4-runtime` |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_tsql" >}} | `3.3.1` | {{< bg "18" "babelfishpg-tsql" "red" >}} {{< bg "17" "babelfishpg-tsql" "red" >}} {{< bg "16" "babelfishpg-tsql" "red" >}} {{< bg "15" "babelfishpg-tsql" "green" >}} {{< bg "14" "babelfishpg-tsql" "red" >}} {{< bg "13" "babelfishpg-tsql" "red" >}} | `babelfishpg-tsql` | `babelfishpg-common`, `libantlr4-runtime4.9.3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-tsql : MISS 0" "red" >}}      |


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
pig ext install babelfishpg_tsql; # install by extension name, for the current active PG version
pig ext install babelfishpg_tsql; # install via package alias, for the active PG version
pig ext install babelfishpg_tsql -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION babelfishpg_tsql;
```

