---
title: "ivorysql_ora"
linkTitle: "ivorysql_ora"
description: "Oracle Compatible extension on Postgres Database"
weight: 9140
categories: ["SIM"]
width: full
---

[**ivorysql**](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora) : Oracle Compatible extension on Postgres Database


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9140** | {{< badge content="ivorysql_ora" link="https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" >}} | {{< ext "ivorysql_ora" "ivorysql" >}} | `1.0` | {{< category "SIM" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `sys` |
|    **Need By**    | {{< ext "ora_btree_gin" >}} {{< ext "ora_btree_gist" >}} |
|    **Siblings**   | {{< ext "ora_btree_gin" >}} {{< ext "ora_btree_gist" >}} {{< ext "pg_get_functiondef" >}} {{< ext "plisql" >}} {{< ext "gb18030_2022" >}} |

> [!Note] compatible with PostgreSQL 18.4


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `ivorysql` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4` | {{< bg "18" "ivorysql-18" "green" >}} {{< bg "17" "ivorysql-17" "red" >}} {{< bg "16" "ivorysql-16" "red" >}} {{< bg "15" "ivorysql-15" "red" >}} {{< bg "14" "ivorysql-14" "red" >}} | `ivorysql-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.4` | {{< bg "18" "ivorysql-18" "green" >}} {{< bg "17" "ivorysql-17" "red" >}} {{< bg "16" "ivorysql-16" "red" >}} {{< bg "15" "ivorysql-15" "red" >}} {{< bg "14" "ivorysql-14" "red" >}} | `ivorysql-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u26.x86_64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "PIGSTY 5.4" "ivorysql-18 : FORK 1" >}}      |      {{< bg "MISS" "ivorysql-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-15 : FORK 0" "red" >}}      |      {{< bg "MISS" "ivorysql-14 : FORK 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ivorysql-18` | `5.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 24.6 MiB | [ivorysql-18-5.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ivorysql-18-5.4-1PIGSTY.el8.x86_64.rpm) |
| `ivorysql-18` | `5.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.1 MiB | [ivorysql-18-5.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ivorysql-18-5.4-1PIGSTY.el8.aarch64.rpm) |
| `ivorysql-18` | `5.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.0 MiB | [ivorysql-18-5.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ivorysql-18-5.4-1PIGSTY.el9.x86_64.rpm) |
| `ivorysql-18` | `5.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 22.8 MiB | [ivorysql-18-5.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ivorysql-18-5.4-1PIGSTY.el9.aarch64.rpm) |
| `ivorysql-18` | `5.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.2 MiB | [ivorysql-18-5.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ivorysql-18-5.4-1PIGSTY.el10.x86_64.rpm) |
| `ivorysql-18` | `5.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.0 MiB | [ivorysql-18-5.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ivorysql-18-5.4-1PIGSTY.el10.aarch64.rpm) |
| `ivorysql-18` | `5.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 23.0 MiB | [ivorysql-18_5.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~bookworm_amd64.deb) |
| `ivorysql-18` | `5.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 22.4 MiB | [ivorysql-18_5.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~bookworm_arm64.deb) |
| `ivorysql-18` | `5.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 20.9 MiB | [ivorysql-18_5.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~trixie_amd64.deb) |
| `ivorysql-18` | `5.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 20.4 MiB | [ivorysql-18_5.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~trixie_arm64.deb) |
| `ivorysql-18` | `5.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 25.1 MiB | [ivorysql-18_5.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~jammy_amd64.deb) |
| `ivorysql-18` | `5.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 24.8 MiB | [ivorysql-18_5.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~jammy_arm64.deb) |
| `ivorysql-18` | `5.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 23.2 MiB | [ivorysql-18_5.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~noble_amd64.deb) |
| `ivorysql-18` | `5.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 23.0 MiB | [ivorysql-18_5.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~noble_arm64.deb) |
| `ivorysql-18` | `5.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.8 MiB | [ivorysql-18_5.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~resolute_amd64.deb) |
| `ivorysql-18` | `5.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.5 MiB | [ivorysql-18_5.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/i/ivorysql-18/ivorysql-18_5.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" title="Repository" icon="github" subtitle="github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ivorysql-5.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg ivorysql;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ivorysql;		# install via package name, for the active PG version
pig install ivorysql_ora;		# install by extension name, for the current active PG version

pig install ivorysql_ora -v 18;   # install for PG 18

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ivorysql_ora;
```




## Usage

> [ivorysql_ora: Oracle Compatible extension on Postgres Database](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora)

The `ivorysql_ora` extension provides Oracle compatibility features for PostgreSQL as part of the IvorySQL project. It adds Oracle-compatible data types, functions, and PL/SQL behavior.

### Enabling

```sql
CREATE EXTENSION ivorysql_ora;
```

### Oracle-Compatible Data Types

The extension adds Oracle-style data types including:

- `NUMBER` / `NUMBER(p,s)` - Oracle-compatible numeric type
- `VARCHAR2(n)` - Oracle-compatible variable-length string
- `DATE` - Oracle-style DATE with time component
- `BINARY_FLOAT` / `BINARY_DOUBLE` - IEEE floating point types

### Oracle-Compatible Functions

Provides Oracle-style built-in functions for string manipulation, date arithmetic, numeric operations, and type conversion that behave consistently with Oracle semantics.

### Compatibility Mode

IvorySQL supports an Oracle compatibility mode that changes parser behavior:

```sql
SET compatible_mode TO oracle;  -- enable Oracle compatibility
SET compatible_mode TO pg;      -- revert to standard PostgreSQL
```

In Oracle mode, the SQL parser accepts Oracle-style syntax including:

- Oracle-style outer joins (`(+)` syntax)
- `CONNECT BY` hierarchical queries
- Oracle-style sequences (`sequence.NEXTVAL`)
- Package-style object references
