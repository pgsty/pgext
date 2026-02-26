---
title: "snowflake"
linkTitle: "snowflake"
description: "Snowflake-style 64-bit ID generator and sequence utilities for PostgreSQL"
weight: 4590
categories: ["FUNC"]
width: full
---

[**snowflake**](https://github.com/pgEdge/snowflake) : Snowflake-style 64-bit ID generator and sequence utilities for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4590** | {{< badge content="snowflake" link="https://github.com/pgEdge/snowflake" >}} | {{< ext "snowflake" >}} | `2.4` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `snowflake` |
|   **See Also**    | {{< ext "spock" >}} {{< ext "lolor" >}} |

> [!Note] works on pgedge kernel fork. Set snowflake.node (1..1023) before using snowflake.nextval().


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `snowflake` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4` | {{< bg "18" "snowflake_18" "red" >}} {{< bg "17" "snowflake_17" "green" >}} {{< bg "16" "snowflake_16" "red" >}} {{< bg "15" "snowflake_15" "red" >}} {{< bg "14" "snowflake_14" "red" >}} {{< bg "13" "snowflake_13" "red" >}} | `snowflake_$v` | `pgedge_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4` | {{< bg "18" "pgedge-18-snowflake" "red" >}} {{< bg "17" "pgedge-17-snowflake" "green" >}} {{< bg "16" "pgedge-16-snowflake" "red" >}} {{< bg "15" "pgedge-15-snowflake" "red" >}} {{< bg "14" "pgedge-14-snowflake" "red" >}} {{< bg "13" "pgedge-13-snowflake" "red" >}} | `pgedge-$v-snowflake` | `pgedge-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "snowflake_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "snowflake_17 : FORK 1" >}}      |      {{< bg "MISS" "snowflake_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "snowflake_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pgedge-18-snowflake : FORK 0" "red" >}}      |      {{< bg "PIGSTY 2.4" "pgedge-17-snowflake : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-snowflake : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-snowflake : FORK 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `snowflake_17` | `2.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 20.7 KiB | [snowflake_17-2.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/snowflake_17-2.4-1PIGSTY.el8.x86_64.rpm) |
| `snowflake_17` | `2.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 20.7 KiB | [snowflake_17-2.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/snowflake_17-2.4-1PIGSTY.el8.aarch64.rpm) |
| `snowflake_17` | `2.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 20.6 KiB | [snowflake_17-2.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/snowflake_17-2.4-1PIGSTY.el9.x86_64.rpm) |
| `snowflake_17` | `2.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 20.4 KiB | [snowflake_17-2.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/snowflake_17-2.4-1PIGSTY.el9.aarch64.rpm) |
| `snowflake_17` | `2.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 20.7 KiB | [snowflake_17-2.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/snowflake_17-2.4-1PIGSTY.el10.x86_64.rpm) |
| `snowflake_17` | `2.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 20.6 KiB | [snowflake_17-2.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/snowflake_17-2.4-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-17-snowflake` | `2.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.2 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-17-snowflake` | `2.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.1 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-17-snowflake` | `2.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.2 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~trixie_amd64.deb) |
| `pgedge-17-snowflake` | `2.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~trixie_arm64.deb) |
| `pgedge-17-snowflake` | `2.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.7 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~jammy_amd64.deb) |
| `pgedge-17-snowflake` | `2.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.4 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~jammy_arm64.deb) |
| `pgedge-17-snowflake` | `2.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.6 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~noble_amd64.deb) |
| `pgedge-17-snowflake` | `2.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.4 KiB | [pgedge-17-snowflake_2.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/snowflake/pgedge-17-snowflake_2.4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/snowflake" title="Repository" icon="github" subtitle="github.com/pgEdge/snowflake" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="snowflake-2.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg snowflake;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install snowflake;		# install via package name, for the active PG version

pig install snowflake -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION snowflake;
```
