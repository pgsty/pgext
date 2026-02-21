---
title: "babelfishpg_money"
linkTitle: "babelfishpg_money"
description: "SQL Server Money Data Type"
weight: 9330
categories: ["SIM"]
width: full
---

[**babelfishpg_money**](https://babelfishpg.org/) : SQL Server Money Data Type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9330** | {{< badge content="babelfishpg_money" link="https://babelfishpg.org/" >}} | {{< ext "babelfishpg_money" >}} | `1.1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "babelfishpg_common" >}} {{< ext "babelfishpg_tsql" >}} {{< ext "babelfishpg_tds" >}} {{< ext "financial" >}} {{< ext "tds_fdw" >}} {{< ext "numeral" >}} {{< ext "orafce" >}} {{< ext "pgtt" >}} |

> [!Note] special case: this extension only works on wiltondb kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `babelfishpg_money` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfish_18" "red" >}} {{< bg "17" "babelfish_17" "green" >}} {{< bg "16" "babelfish_16" "red" >}} {{< bg "15" "babelfish_15" "red" >}} {{< bg "14" "babelfish_14" "red" >}} {{< bg "13" "babelfish_13" "red" >}} | `babelfish_$v` | `babelfishpg_$v`, `antlr4-runtime413` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1.0` | {{< bg "18" "babelfishpg-18-babelfish" "red" >}} {{< bg "17" "babelfishpg-17-babelfish" "green" >}} {{< bg "16" "babelfishpg-16-babelfish" "red" >}} {{< bg "15" "babelfishpg-15-babelfish" "red" >}} {{< bg "14" "babelfishpg-14-babelfish" "red" >}} {{< bg "13" "babelfishpg-13-babelfish" "red" >}} | `babelfishpg-$v-babelfish` | `babelfishpg-$v-babelfish`, `libantlr4-runtime413` |


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


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `babelfish_17` | `5.5.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.5 MiB | [babelfish_17-5.5.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/babelfish_17-5.5.0-1PIGSTY.el8.x86_64.rpm) |
| `babelfish_17` | `5.5.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [babelfish_17-5.5.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/babelfish_17-5.5.0-1PIGSTY.el8.aarch64.rpm) |
| `babelfish_17` | `5.5.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.4 MiB | [babelfish_17-5.5.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/babelfish_17-5.5.0-1PIGSTY.el9.x86_64.rpm) |
| `babelfish_17` | `5.5.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.3 MiB | [babelfish_17-5.5.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/babelfish_17-5.5.0-1PIGSTY.el9.aarch64.rpm) |
| `babelfish_17` | `5.5.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.2 MiB | [babelfish_17-5.5.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/babelfish_17-5.5.0-1PIGSTY.el10.x86_64.rpm) |
| `babelfish_17` | `5.5.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.2 MiB | [babelfish_17-5.5.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/babelfish_17-5.5.0-1PIGSTY.el10.aarch64.rpm) |
| `babelfishpg-17-babelfish` | `5.5.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~bookworm_amd64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~bookworm_arm64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~trixie_amd64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.9 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~trixie_arm64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~jammy_amd64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~jammy_arm64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.2 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~noble_amd64.deb) |
| `babelfishpg-17-babelfish` | `5.5.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [babelfishpg-17-babelfish_5.5.0-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/b/babelfishpg-17-babelfish/babelfishpg-17-babelfish_5.5.0-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://babelfishpg.org/" title="Repository" icon="link" subtitle="babelfishpg.org/" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install babelfishpg_money;		# install via package name, for the active PG version

pig install babelfishpg_money -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION babelfishpg_money;
```
