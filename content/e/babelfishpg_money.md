---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["SIM"]
width: full
---

SQL Server Money Data Type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_money" >}} | `1.1.0` | {{< bg "18" "babelfishpg-money*" "red" >}} {{< bg "17" "babelfishpg-money*" "red" >}} {{< bg "16" "babelfishpg-money*" "red" >}} {{< bg "15" "babelfishpg-money*" "green" >}} {{< bg "14" "babelfishpg-money*" "red" >}} {{< bg "13" "babelfishpg-money*" "red" >}} | `babelfishpg-money*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_money" >}} | `1.1.0` | {{< bg "18" "babelfishpg-money" "red" >}} {{< bg "17" "babelfishpg-money" "red" >}} {{< bg "16" "babelfishpg-money" "red" >}} {{< bg "15" "babelfishpg-money" "green" >}} {{< bg "14" "babelfishpg-money" "red" >}} {{< bg "13" "babelfishpg-money" "red" >}} | `babelfishpg-money` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |      {{< bg "MISS" "babelfishpg-money : MISS 0" "red" >}}      |


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
pig ext install babelfishpg_money; # install by extension name, for the current active PG version
pig ext install babelfishpg_money; # install via package alias, for the active PG version
pig ext install babelfishpg_money -v 15;   # install for PG 15

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION babelfishpg_money;
```

