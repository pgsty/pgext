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
| **2180** | {{< badge content="pg_textsearch" link="https://github.com/timescale/pg_textsearch" >}} | {{< ext "pg_textsearch" >}} | `0.1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pgroonga" >}} {{< ext "pg_bigm" >}} {{< ext "zhparser" >}} {{< ext "pg_trgm" >}} {{< ext "rum" >}} {{< ext "biscuit" >}} {{< ext "fuzzystrmatch" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} {{< bg "13" "" "red" >}} | `pg_textsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pg_textsearch_18*" "green" >}} {{< bg "17" "pg_textsearch_17*" "green" >}} {{< bg "16" "pg_textsearch_16*" "red" >}} {{< bg "15" "pg_textsearch_15*" "red" >}} {{< bg "14" "pg_textsearch_14*" "red" >}} {{< bg "13" "pg_textsearch_13*" "red" >}} | `pg_textsearch_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-textsearch" "green" >}} {{< bg "17" "postgresql-17-textsearch" "green" >}} {{< bg "16" "postgresql-16-textsearch" "red" >}} {{< bg "15" "postgresql-15-textsearch" "red" >}} {{< bg "14" "postgresql-14-textsearch" "red" >}} {{< bg "13" "postgresql-13-textsearch" "red" >}} | `postgresql-$v-textsearch` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pg_textsearch_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_textsearch_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_textsearch_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-textsearch : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-textsearch : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-textsearch : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.6 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.5 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.3 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 66.5 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.4 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 67.0 KiB | [pg_textsearch_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-textsearch` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 576.3 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 569.8 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 577.4 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 570.7 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 623.7 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 620.0 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 617.2 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-textsearch` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 613.8 KiB | [postgresql-18-textsearch_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 72.3 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_textsearch_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_textsearch_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.2 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_textsearch_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_textsearch_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 69.2 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_textsearch_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_textsearch_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 66.1 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_textsearch_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_textsearch_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 69.1 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_textsearch_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_textsearch_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 66.7 KiB | [pg_textsearch_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_textsearch_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-textsearch` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 572.3 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 564.8 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 572.1 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 565.1 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 657.4 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 653.1 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 612.1 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-textsearch` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 608.0 KiB | [postgresql-17-textsearch_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_0.1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_textsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_textsearch" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_textsearch-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_textsearch;		# build rpm / deb with pig
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


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_textsearch;
```
