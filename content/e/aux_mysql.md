---
title: "aux_mysql"
linkTitle: "aux_mysql"
description: "MySQL Supplementary Extension"
weight: 9420
categories: ["SIM"]
width: full
---

[**openhalo**](https://github.com/HaloTech-Co-Ltd/openHalo) : MySQL Supplementary Extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9420** | {{< badge content="aux_mysql" link="https://github.com/HaloTech-Co-Ltd/openHalo" >}} | {{< ext "aux_mysql" "openhalo" >}} | `1.5` | {{< category "SIM" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `mysql` |

> [!Note] module_pathname=$libdir/mysm; openHalo 14.x only


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `openhalo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "openhalodb_18" "red" >}} {{< bg "17" "openhalodb_17" "red" >}} {{< bg "16" "openhalodb_16" "red" >}} {{< bg "15" "openhalodb_15" "red" >}} {{< bg "14" "openhalodb_14" "green" >}} {{< bg "13" "openhalodb_13" "red" >}} | `openhalodb_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "openhalodb-18" "red" >}} {{< bg "17" "openhalodb-17" "red" >}} {{< bg "16" "openhalodb-16" "red" >}} {{< bg "15" "openhalodb-15" "red" >}} {{< bg "14" "openhalodb-14" "green" >}} {{< bg "13" "openhalodb-13" "red" >}} | `openhalodb-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "openhalodb_18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb_15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb_14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "openhalodb-18 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-17 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-16 : FORK 0" "red" >}}      |      {{< bg "MISS" "openhalodb-15 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.0" "openhalodb-14 : FORK 1" >}}      |      {{< bg "MISS" "openhalodb-13 : FORK 0" "red" >}}      |


{{< tabs items="PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `openhalodb_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 MiB | [openhalodb_14-1.0-beta1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/openhalodb_14-1.0-beta1PIGSTY.el8.x86_64.rpm) |
| `openhalodb_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.4 MiB | [openhalodb_14-1.0-beta1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/openhalodb_14-1.0-beta1PIGSTY.el8.aarch64.rpm) |
| `openhalodb_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.0 MiB | [openhalodb_14-1.0-beta1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/openhalodb_14-1.0-beta1PIGSTY.el9.x86_64.rpm) |
| `openhalodb_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 9.8 MiB | [openhalodb_14-1.0-beta1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/openhalodb_14-1.0-beta1PIGSTY.el9.aarch64.rpm) |
| `openhalodb_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.1 MiB | [openhalodb_14-1.0-beta1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/openhalodb_14-1.0-beta1PIGSTY.el10.x86_64.rpm) |
| `openhalodb_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 9.9 MiB | [openhalodb_14-1.0-beta1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/openhalodb_14-1.0-beta1PIGSTY.el10.aarch64.rpm) |
| `openhalodb-14` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.4 MiB | [openhalodb-14_1.0-beta1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~bookworm_amd64.deb) |
| `openhalodb-14` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.8 MiB | [openhalodb-14_1.0-beta1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~bookworm_arm64.deb) |
| `openhalodb-14` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.6 MiB | [openhalodb-14_1.0-beta1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~trixie_amd64.deb) |
| `openhalodb-14` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.0 MiB | [openhalodb-14_1.0-beta1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~trixie_arm64.deb) |
| `openhalodb-14` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 21.2 MiB | [openhalodb-14_1.0-beta1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~jammy_amd64.deb) |
| `openhalodb-14` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 20.9 MiB | [openhalodb-14_1.0-beta1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~jammy_arm64.deb) |
| `openhalodb-14` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 19.5 MiB | [openhalodb-14_1.0-beta1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~noble_amd64.deb) |
| `openhalodb-14` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 19.3 MiB | [openhalodb-14_1.0-beta1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/openhalodb/openhalodb-14_1.0-beta1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HaloTech-Co-Ltd/openHalo" title="Repository" icon="github" subtitle="github.com/HaloTech-Co-Ltd/openHalo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="openhalodb-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg openhalo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install openhalo;		# install via package name, for the active PG version
pig install aux_mysql;		# install by extension name, for the current active PG version

pig install aux_mysql -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION aux_mysql;
```
