---
title: "babelfishpg_common"
linkTitle: "babelfishpg_common"
description: "SQL Server Transact SQL Datatype Support"
weight: 9300
categories: ["SIM"]
width: full
---

SQL Server Transact SQL Datatype Support


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9300** | {{< badge content="babelfishpg_common" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_common" >}} | `3.3.3` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "babelfishpg_tsql" >}} |
|   **See Also**    | {{< ext "tds_fdw" >}} {{< ext "babelfishpg_tds" >}} {{< ext "babelfishpg_money" >}} {{< ext "pg_hint_plan" >}} {{< ext "uuid-ossp" >}} {{< ext "session_variable" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_common" >}} | `3.3.3` | {{< bg "18" "babelfishpg-common*" "red" >}} {{< bg "17" "babelfishpg-common*" "red" >}} {{< bg "16" "babelfishpg-common*" "red" >}} {{< bg "15" "babelfishpg-common*" "green" >}} {{< bg "14" "babelfishpg-common*" "red" >}} {{< bg "13" "babelfishpg-common*" "red" >}} | `babelfishpg-common*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_common" >}} | `3.3.3` | {{< bg "18" "babelfishpg-common" "red" >}} {{< bg "17" "babelfishpg-common" "red" >}} {{< bg "16" "babelfishpg-common" "red" >}} {{< bg "15" "babelfishpg-common" "green" >}} {{< bg "14" "babelfishpg-common" "red" >}} {{< bg "13" "babelfishpg-common" "red" >}} | `babelfishpg-common` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-common : MISS 0" "red" >}}      |


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
pig ext install babelfishpg_common; # install by extension name, for the current active PG version
pig ext install babelfishpg_common; # install via package alias, for the active PG version
pig ext install babelfishpg_common -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION babelfishpg_common;
```

