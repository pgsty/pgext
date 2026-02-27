---
title: "lolor"
linkTitle: "lolor"
description: "Logical-replication-friendly replacement for PostgreSQL large objects"
weight: 9570
categories: ["ETL"]
width: full
---

[**lolor**](https://github.com/pgEdge/lolor) : Logical-replication-friendly replacement for PostgreSQL large objects


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9570** | {{< badge content="lolor" link="https://github.com/pgEdge/lolor" >}} | {{< ext "lolor" >}} | `1.2.2` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lolor` |
|   **See Also**    | {{< ext "spock" >}} {{< ext "snowflake" >}} |

> [!Note] works on pgedge kernel fork. Requires lolor.node


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `lolor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "lolor_18" "red" >}} {{< bg "17" "lolor_17" "green" >}} {{< bg "16" "lolor_16" "red" >}} {{< bg "15" "lolor_15" "red" >}} {{< bg "14" "lolor_14" "red" >}} {{< bg "13" "lolor_13" "red" >}} | `lolor_$v` | `pgedge_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "pgedge-18-lolor" "red" >}} {{< bg "17" "pgedge-17-lolor" "green" >}} {{< bg "16" "pgedge-16-lolor" "red" >}} {{< bg "15" "pgedge-15-lolor" "red" >}} {{< bg "14" "pgedge-14-lolor" "red" >}} {{< bg "13" "pgedge-13-lolor" "red" >}} | `pgedge-$v-lolor` | `pgedge-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "lolor_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "lolor_17 : FORK 1" >}}      |      {{< bg "MISS" "lolor_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "lolor_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "pgedge-18-lolor : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.2.2" "pgedge-17-lolor : FORK 1" >}}      |      {{< bg "MISS" "pgedge-16-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-15-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-14-lolor : FORK 0" "red" >}}      |      {{< bg "MISS" "pgedge-13-lolor : FORK 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `lolor_17` | `1.2.2` | [el8.x86_64](/os/el8.x86_64) | pigsty | 29.4 KiB | [lolor_17-1.2.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/lolor_17-1.2.2-1PIGSTY.el8.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el8.aarch64](/os/el8.aarch64) | pigsty | 28.5 KiB | [lolor_17-1.2.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/lolor_17-1.2.2-1PIGSTY.el8.aarch64.rpm) |
| `lolor_17` | `1.2.2` | [el9.x86_64](/os/el9.x86_64) | pigsty | 28.6 KiB | [lolor_17-1.2.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/lolor_17-1.2.2-1PIGSTY.el9.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.8 KiB | [lolor_17-1.2.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/lolor_17-1.2.2-1PIGSTY.el9.aarch64.rpm) |
| `lolor_17` | `1.2.2` | [el10.x86_64](/os/el10.x86_64) | pigsty | 28.8 KiB | [lolor_17-1.2.2-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/lolor_17-1.2.2-1PIGSTY.el10.x86_64.rpm) |
| `lolor_17` | `1.2.2` | [el10.aarch64](/os/el10.aarch64) | pigsty | 28.0 KiB | [lolor_17-1.2.2-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/lolor_17-1.2.2-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-17-lolor` | `1.2.2` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.6 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [d12.aarch64](/os/d12.aarch64) | pigsty | 15.9 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.7 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~trixie_amd64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [d13.aarch64](/os/d13.aarch64) | pigsty | 15.9 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~trixie_arm64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [u22.x86_64](/os/u22.x86_64) | pigsty | 18.1 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~jammy_amd64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.3 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~jammy_arm64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~noble_amd64.deb) |
| `pgedge-17-lolor` | `1.2.2` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.0 KiB | [pgedge-17-lolor_1.2.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/l/lolor/pgedge-17-lolor_1.2.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/lolor" title="Repository" icon="github" subtitle="github.com/pgEdge/lolor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lolor-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg lolor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install lolor;		# install via package name, for the active PG version

pig install lolor -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lolor;
```
