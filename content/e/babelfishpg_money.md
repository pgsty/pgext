---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["Sim"]
width: full
---

SQL Server Money Data Type

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" "babelfishpg_money" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-dt-" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/babelfishpg_money" >}} | `1.1.0` | {{< badge content="18" color="red" alt="babelfishpg-money*" >}} {{< badge content="17" color="red" alt="babelfishpg-money*" >}} {{< badge content="16" color="red" alt="babelfishpg-money*" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-money*" >}} | `babelfishpg-money*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/babelfishpg_money" >}} | `1.1.0` | {{< badge content="18" color="red" alt="babelfishpg-money" >}} {{< badge content="17" color="red" alt="babelfishpg-money" >}} {{< badge content="16" color="red" alt="babelfishpg-money" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="red" alt="babelfishpg-money" >}} | `babelfishpg-money` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `el8.aarch64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `el9.x86_64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `el9.aarch64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `d12.x86_64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `d12.aarch64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `u22.x86_64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `u22.aarch64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `u24.x86_64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |
|    `u24.aarch64`    |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |    {{< pkg "babelfishpg-money" >}}     |


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

