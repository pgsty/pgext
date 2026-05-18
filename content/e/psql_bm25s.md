---
title: "psql_bm25s"
linkTitle: "psql_bm25s"
description: "PostgreSQL extension for BM25-family lexical retrieval"
weight: 2210
categories: ["FTS"]
width: full
---

[**psql_bm25s**](https://github.com/Intelligent-Internet/psql_bm25s) : PostgreSQL extension for BM25-family lexical retrieval


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2210** | {{< badge content="psql_bm25s" link="https://github.com/Intelligent-Internet/psql_bm25s" >}} | {{< ext "psql_bm25s" >}} | `0.4.13` | {{< category "FTS" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_textsearch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} |

> [!Note] Supports PostgreSQL 17-18; optional shared_preload_libraries arena is not required for normal use.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.13` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `psql_bm25s` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.13` | {{< bg "18" "psql_bm25s_18" "green" >}} {{< bg "17" "psql_bm25s_17" "green" >}} {{< bg "16" "psql_bm25s_16" "red" >}} {{< bg "15" "psql_bm25s_15" "red" >}} {{< bg "14" "psql_bm25s_14" "red" >}} | `psql_bm25s_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.13` | {{< bg "18" "postgresql-18-psql-bm25s" "green" >}} {{< bg "17" "postgresql-17-psql-bm25s" "green" >}} {{< bg "16" "postgresql-16-psql-bm25s" "red" >}} {{< bg "15" "postgresql-15-psql-bm25s" "red" >}} {{< bg "14" "postgresql-14-psql-bm25s" "red" >}} | `postgresql-$v-psql-bm25s` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "psql_bm25s_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "psql_bm25s_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-18-psql-bm25s : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.13" "postgresql-17-psql-bm25s : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-psql-bm25s : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-psql-bm25s : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-psql-bm25s` | `0.4.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 497.3 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 475.9 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 499.7 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 479.1 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 527.1 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 511.5 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 510.3 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 497.3 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u26.x86_64](/os/u26.x86_64) | pigsty | 506.9 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-psql-bm25s` | `0.4.13` | [u26.aarch64](/os/u26.aarch64) | pigsty | 492.8 KiB | [postgresql-18-psql-bm25s_0.4.13-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-18-psql-bm25s_0.4.13-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-psql-bm25s` | `0.4.13` | [d12.x86_64](/os/d12.x86_64) | pigsty | 497.2 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [d12.aarch64](/os/d12.aarch64) | pigsty | 475.6 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [d13.x86_64](/os/d13.x86_64) | pigsty | 499.3 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [d13.aarch64](/os/d13.aarch64) | pigsty | 479.1 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u22.x86_64](/os/u22.x86_64) | pigsty | 553.4 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u22.aarch64](/os/u22.aarch64) | pigsty | 538.1 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u24.x86_64](/os/u24.x86_64) | pigsty | 510.2 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u24.aarch64](/os/u24.aarch64) | pigsty | 497.3 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u26.x86_64](/os/u26.x86_64) | pigsty | 506.9 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-psql-bm25s` | `0.4.13` | [u26.aarch64](/os/u26.aarch64) | pigsty | 492.7 KiB | [postgresql-17-psql-bm25s_0.4.13-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/psql-bm25s/postgresql-17-psql-bm25s_0.4.13-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Intelligent-Internet/psql_bm25s" title="Repository" icon="github" subtitle="github.com/Intelligent-Internet/psql_bm25s" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="psql_bm25s-0.4.13.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg psql_bm25s;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install psql_bm25s;		# install via package name, for the active PG version

pig install psql_bm25s -v 18;   # install for PG 18
pig install psql_bm25s -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION psql_bm25s;
```
