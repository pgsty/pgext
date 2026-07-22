---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["SIM"]
width: full
---

[**babelfish**](https://babelfishpg.org/) : SQL Server Money Data Type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" "babelfish" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |
|    **Siblings**   | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `babelfish` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfish-18" "green" >}} {{< bg "17" "babelfish-17" "green" >}} {{< bg "16" "babelfish-16" "red" >}} {{< bg "15" "babelfish-15" "red" >}} {{< bg "14" "babelfish-14" "red" >}} | `babelfish-$v` | `libantlr4-runtime413` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 6.0.0" "babelfish-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 5.4.0" "babelfish-17 : AVAIL 1" "green" >}} | {{< bg "N/A" "babelfish-16 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-15 : N/A 0" "gray" >}} | {{< bg "N/A" "babelfish-14 : N/A 0" "gray" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="babelfish-17-17.7-5.4.0.tar.gz" >}}
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
pig install babelfishpg_money;		# install by extension name, for the current active PG version

pig install babelfishpg_money -v 18;   # install for PG 18
pig install babelfishpg_money -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_money;
```




## Usage

> [babelfishpg_money: SQL Server Money Data Type](https://babelfishpg.org/)

The `babelfishpg_money` extension provides SQL Server-compatible `MONEY` and `SMALLMONEY` data type implementations for PostgreSQL as part of the Babelfish project.

### Enabling

```sql
CREATE EXTENSION babelfishpg_money;
```

### Data Types

- **MONEY** - 8-byte monetary value ranging from -922,337,203,685,477.5808 to 922,337,203,685,477.5807, with fixed 4 decimal places
- **SMALLMONEY** - 4-byte monetary value ranging from -214,748.3648 to 214,748.3647, with fixed 4 decimal places

### Behavior

The extension implements SQL Server's monetary arithmetic rules:

- Fixed-point representation with exactly 4 decimal digits
- SQL Server-compatible rounding behavior for monetary calculations
- Proper casting between MONEY and other numeric types
- Arithmetic operations follow SQL Server semantics (e.g., money / money = money, not float)

### Notes

- Part of the Babelfish for PostgreSQL project
- Works in conjunction with `babelfishpg_common` which provides the base type infrastructure
- The PostgreSQL built-in `money` type has different precision and behavior; this extension provides the SQL Server-compatible variant
