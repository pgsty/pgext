---
title: "orioledb"
linkTitle: "orioledb"
description: "OrioleDB, the next generation transactional engine"
weight: 2910
categories: ["FEAT"]
width: full
---

[**orioledb**](https://github.com/orioledb/orioledb) : OrioleDB, the next generation transactional engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2910** | {{< badge content="orioledb" link="https://github.com/orioledb/orioledb" >}} | {{< ext "orioledb" >}} | `1.5` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_mooncake" >}} {{< ext "citus_columnar" >}} {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "pg_strom" >}} |

> [!Note] special case: this extension only works on patched postgres kernel: oriolepg


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `orioledb` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "orioledb_18*" "red" >}} {{< bg "17" "orioledb_17*" "green" >}} {{< bg "16" "orioledb_16*" "red" >}} {{< bg "15" "orioledb_15*" "red" >}} {{< bg "14" "orioledb_14*" "red" >}} {{< bg "13" "orioledb_13*" "red" >}} | `orioledb_$v*` | `oriolepg_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5` | {{< bg "18" "oriolepg-18-orioledb" "red" >}} {{< bg "17" "oriolepg-17-orioledb" "green" >}} {{< bg "16" "oriolepg-16-orioledb" "red" >}} {{< bg "15" "oriolepg-15-orioledb" "red" >}} {{< bg "14" "oriolepg-14-orioledb" "red" >}} {{< bg "13" "oriolepg-13-orioledb" "red" >}} | `oriolepg-$v-orioledb` | `oriolepg-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "orioledb_18 : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : FORK 1" >}}      |      {{< bg "MISS" "orioledb_16 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : FORK 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : FORK 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "oriolepg-18-orioledb : FORK 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : FORK 1" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : FORK 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : FORK 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orioledb_17` | `1.5` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el8.x86_64.rpm) |
| `orioledb_17` | `1.5` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el8.aarch64.rpm) |
| `orioledb_17` | `1.5` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el9.x86_64.rpm) |
| `orioledb_17` | `1.5` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el9.aarch64.rpm) |
| `orioledb_17` | `1.5` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el10.x86_64.rpm) |
| `orioledb_17` | `1.5` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el10.aarch64.rpm) |
| `oriolepg-17-orioledb` | `1.5` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.5 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_amd64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.4 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_arm64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.2 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~trixie_amd64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~trixie_arm64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.6 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_amd64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.5 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_arm64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.4 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_amd64.deb) |
| `oriolepg-17-orioledb` | `1.5` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/orioledb/orioledb" title="Repository" icon="github" subtitle="github.com/orioledb/orioledb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="orioledb-beta12.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg orioledb;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install orioledb;		# install via package name, for the active PG version

pig install orioledb -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'orioledb';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION orioledb;
```
