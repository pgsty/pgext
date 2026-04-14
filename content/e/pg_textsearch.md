---
title: "pg_textsearch"
linkTitle: "pg_textsearch"
description: "Full-text search with BM25 ranking"
weight: 2180
categories: ["FTS"]
width: full
---

[**pg_textsearch**](https://github.com/timescale/pg_textsearch) : Full-text search with BM25 ranking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2180** | {{< badge content="pg_textsearch" link="https://github.com/timescale/pg_textsearch" >}} | {{< ext "pg_textsearch" >}} | `1.0.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "biscuit" >}} {{< ext "fuzzystrmatch" >}} |

> [!Note] bm25 am conflicts with pg_search; must be preloaded via shared_preload_libraries.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_textsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pg_textsearch_18" "green" >}} {{< bg "17" "pg_textsearch_17" "green" >}} {{< bg "16" "pg_textsearch_16" "red" >}} {{< bg "15" "pg_textsearch_15" "red" >}} {{< bg "14" "pg_textsearch_14" "red" >}} | `pg_textsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-textsearch" "green" >}} {{< bg "17" "postgresql-17-textsearch" "green" >}} {{< bg "16" "postgresql-16-textsearch" "red" >}} {{< bg "15" "postgresql-15-textsearch" "red" >}} {{< bg "14" "postgresql-14-textsearch" "red" >}} | `postgresql-$v-textsearch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 110.6 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 105.3 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 101.9 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 99.3 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 105.1 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.7 KiB | [pg_textsearch_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-textsearch` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 894.9 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 887.0 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 895.9 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 888.3 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 991.8 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 990.0 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 953.6 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-textsearch` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 949.3 KiB | [postgresql-18-textsearch_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 110.5 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 105.1 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 101.8 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 99.0 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 104.9 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 100.6 KiB | [pg_textsearch_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-textsearch` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 881.3 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 874.8 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 881.6 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 876.6 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.0 MiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 942.8 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-textsearch` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 938.6 KiB | [postgresql-17-textsearch_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_textsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_textsearch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_textsearch-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_textsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_textsearch;		# install via package name, for the active PG version

pig install pg_textsearch -v 18;   # install for PG 18
pig install pg_textsearch -v 17;   # install for PG 17

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pg_textsearch';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_textsearch;
```



## Usage

> [pg_textsearch: Modern ranked text search for PostgreSQL with BM25](https://github.com/timescale/pg_textsearch)

Modern ranked text search using BM25 scoring with Block-Max WAND optimization. Simple syntax, fast top-k queries, parallel index builds, and partitioned table support.

Add to `shared_preload_libraries`:

```
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
```

### Quick Start

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents (content) VALUES
    ('PostgreSQL is a powerful database system'),
    ('BM25 is an effective ranking function'),
    ('Full text search with custom scoring');

-- Create a BM25 index
CREATE INDEX docs_idx ON documents USING bm25(content) WITH (text_config='english');

-- Query using the <@> operator (returns negative BM25 score, lower = better match)
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

### Querying

```sql
-- Auto-detect index from column
SELECT * FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;

-- Explicit index specification
SELECT * FROM documents
WHERE content <@> to_bm25query('database system', 'docs_idx') < -1.0;
```

### Filtering

**Pre-filtering** reduces rows before scoring (best with selective filters):

```sql
CREATE INDEX ON documents (category_id);
SELECT * FROM documents
WHERE category_id = 123
ORDER BY content <@> 'search terms'
LIMIT 10;
```

**Post-filtering** applies BM25 scan first, then filters:

```sql
SELECT * FROM documents
WHERE content <@> to_bm25query('search terms', 'docs_idx') < -5.0
ORDER BY content <@> 'search terms'
LIMIT 10;
```

### Index Options

| Option | Default | Description |
|--------|---------|-------------|
| `text_config` | (required) | PostgreSQL text search configuration |
| `k1` | 1.2 | Term frequency saturation parameter |
| `b` | 0.75 | Length normalization parameter |

```sql
CREATE INDEX ON documents USING bm25(content)
  WITH (text_config='english', k1=1.5, b=0.8);

-- Language-specific configurations
CREATE INDEX ON french_docs USING bm25(content) WITH (text_config='french');
CREATE INDEX ON german_docs USING bm25(content) WITH (text_config='german');
```

### Data Types

**bm25query** — represents queries for BM25 scoring:

```sql
SELECT to_bm25query('search query text', 'docs_idx');
-- docs_idx:search query text
```
