---
title: "babelfishpg_tsql"
linkTitle: "babelfishpg_tsql"
description: "SQL Server Transact SQL compatibility"
weight: 9310
categories: ["Sim"]
width: full
---

SQL Server Transact SQL compatibility

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9310** | {{< badge content="babelfishpg_tsql" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_tsql" "babelfishpg_tsql" >}} | `3.3.1` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "babelfishpg_common" >}} {{< ext "uuid-ossp" >}} |
|    **Need By**    | {{< ext "babelfishpg_tds" >}} |
|   **See Also**    | {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "tds_fdw" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} {{< ext "db_migrator" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_tsql" >}} | `3.3.1` | {{< badge content="18" color="red" alt="babelfishpg-tsql*" >}} {{< badge content="17" color="red" alt="babelfishpg-tsql*" >}} {{< badge content="16" color="red" alt="babelfishpg-tsql*" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-tsql*" >}} | `babelfishpg-tsql*` | `babelfishpg-common`, `libantlr4-runtime` |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_tsql" >}} | `3.3.1` | {{< badge content="18" color="red" alt="babelfishpg-tsql" >}} {{< badge content="17" color="red" alt="babelfishpg-tsql" >}} {{< badge content="16" color="red" alt="babelfishpg-tsql" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-tsql" >}} | `babelfishpg-tsql` | `babelfishpg-common`, `libantlr4-runtime4.9.3` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `el8.aarch64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `el9.x86_64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `el9.aarch64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `d12.x86_64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `d12.aarch64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `u22.x86_64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `u22.aarch64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `u24.x86_64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |
|    `u24.aarch64`    |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |    {{< pkg "babelfishpg-tsql" >}}     |


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

