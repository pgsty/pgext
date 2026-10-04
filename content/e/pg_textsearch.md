---
title: "pg_textsearch"
linkTitle: "pg_textsearch"
description: "Full-text search with BM25 ranking"
weight: 2190
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["MIXED"]
page_width: full
---

[**pg_textsearch**](https://github.com/timescale/pg_textsearch) : Full-text search with BM25 ranking


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2190** | {{< badge content="pg_textsearch" link="https://github.com/timescale/pg_textsearch" >}} | {{< ext "pg_textsearch" >}} | `1.5.1` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_search" >}} {{< ext "pg_bestmatch" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_fts" >}} {{< ext "pgroonga" >}} {{< ext "pg_rrf" >}} {{< ext "psql_bm25s" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} |

> [!Note] bm25 am conflicts with pg_search and vchord_bm25


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.5.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_textsearch` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.5.1` | {{< bg "18" "pg_textsearch_18" "green" >}} {{< bg "17" "pg_textsearch_17" "green" >}} {{< bg "16" "pg_textsearch_16" "red" >}} {{< bg "15" "pg_textsearch_15" "red" >}} {{< bg "14" "pg_textsearch_14" "red" >}} | `pg_textsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.5.1` | {{< bg "18" "postgresql-18-textsearch" "green" >}} {{< bg "17" "postgresql-17-textsearch" "green" >}} {{< bg "16" "postgresql-16-textsearch" "red" >}} {{< bg "15" "postgresql-15-textsearch" "red" >}} {{< bg "14" "postgresql-14-textsearch" "red" >}} | `postgresql-$v-textsearch` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.5.1" "pg_textsearch_17 : AVAIL 2" "blue" >}} | {{< bg "N/A" "pg_textsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_textsearch_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-18-textsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.5.1" "postgresql-17-textsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-textsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-textsearch : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_18` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 210.6 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_textsearch_18-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.5 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_textsearch_18-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_textsearch_18` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 194.6 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_textsearch_18-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.5 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_textsearch_18-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_textsearch_18` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.4 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_textsearch_18-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 125.8 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_textsearch_18-1.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_textsearch_18` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.2 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_textsearch_18-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 123.3 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_textsearch_18-1.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_textsearch_18` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 212.8 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_textsearch_18-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.5 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_textsearch_18-1.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_textsearch_18` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 200.0 KiB | [pg_textsearch_18-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_textsearch_18-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_textsearch_18` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 125.7 KiB | [pg_textsearch_18-1.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_textsearch_18-1.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-textsearch` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.9 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.9 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.0 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.0 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.0 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-textsearch` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.0 MiB | [postgresql-18-textsearch_1.5.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-18-textsearch_1.5.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_textsearch_17` | `1.5.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 210.3 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_textsearch_17-1.5.1-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 129.3 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_textsearch_17-1.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `pg_textsearch_17` | `1.5.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 194.4 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_textsearch_17-1.5.1-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 122.5 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_textsearch_17-1.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `pg_textsearch_17` | `1.5.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 206.3 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_textsearch_17-1.5.1-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 125.7 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_textsearch_17-1.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `pg_textsearch_17` | `1.5.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 196.3 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_textsearch_17-1.5.1-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 123.1 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_textsearch_17-1.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `pg_textsearch_17` | `1.5.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 212.7 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_textsearch_17-1.5.1-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 129.4 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_textsearch_17-1.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `pg_textsearch_17` | `1.5.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 199.9 KiB | [pg_textsearch_17-1.5.1-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_textsearch_17-1.5.1-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_textsearch_17` | `1.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 125.7 KiB | [pg_textsearch_17-1.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_textsearch_17-1.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-textsearch` | `1.5.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.2 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.0 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.0 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-textsearch` | `1.5.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.0 MiB | [postgresql-17-textsearch_1.5.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-textsearch/postgresql-17-textsearch_1.5.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_textsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_textsearch" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_textsearch-1.5.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_textsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [Version 1.5.1 README](https://github.com/timescale/pg_textsearch/blob/v1.5.1/README.md)
- [Control file](https://github.com/timescale/pg_textsearch/blob/v1.5.1/pg_textsearch.control)
- [Versioned SQL](https://github.com/timescale/pg_textsearch/blob/v1.5.1/sql/pg_textsearch--1.5.1.sql)
- [Version 1.5.1 release](https://github.com/timescale/pg_textsearch/releases/tag/v1.5.1)

`pg_textsearch` provides BM25-ranked full-text search with the `bm25` access method and `<@>` operator. Version 1.5.1 supports PostgreSQL 17 and 18 and requires preloading and restart. Preserve other entries when updating the preload list.

### Build and Query

```conf
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents(content) VALUES
    ('PostgreSQL is a database system'),
    ('BM25 ranks full text search results');
CREATE INDEX docs_idx ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT id, content <@> 'database system' AS score
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

Scores are negative BM25 values, so lower scores rank first. Use `ORDER BY` with `LIMIT` for top-k execution. Specify the index explicitly for standalone scoring, partial indexes, or PL/pgSQL:

```sql
SELECT id FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

### Index and Query Options

`text_config` is required and names a PostgreSQL text search configuration. `k1` defaults to 1.2 and `b` to 0.75. The `bm25query` type and `to_bm25query(text, text)` carry explicit query/index context. Standalone scoring requires SELECT permission on the table or indexed columns.

The extension supports `text[]`, `varchar[]` and `bpchar[]`, immutable text expression indexes, partial indexes and partitioned tables. Repeat an indexed expression in the query. Partial indexes need the matching predicate and an explicit index name. Version 1.5.0 improves filtered top-k execution, but restrictive post-filters can still return fewer rows than requested.

For Chinese tokenization, configure a parser such as `zhparser` and use that text search configuration. This is an optional workflow dependency. For very large texts without whitespace word boundaries, upstream recommends application-controlled chunks in a text array.

### Maintenance and Configuration

Starting with 1.3.0, the durable memtable is stored in index pages and uses standard PostgreSQL WAL replay. Queries can still use a shared-memory read cache, controlled by the cache and memory-limit settings below. Automatic compaction runs during spills, so write-heavy workloads can observe synchronous compaction latency.

```sql
SELECT bm25_spill_index('docs_idx');
SELECT bm25_force_merge('docs_idx');
```

Use force-merge after bulk loading rather than during steady write traffic. VACUUM also spills pending memtable pages. Relevant settings are:

| Setting | Default | Purpose |
| --- | --- | --- |
| `pg_textsearch.default_limit` | 1000 | Scoring bound without a query limit |
| `pg_textsearch.compress_segments` | on | Posting-block compression |
| `pg_textsearch.segments_per_level` | 8 | Compaction threshold |
| `pg_textsearch.bulk_load_threshold` | 100000 | Terms per transaction before spilling |
| `pg_textsearch.memtable_pages_threshold` | 64 | Chain-page count before spilling |
| `pg_textsearch.memtable_cache_enabled` | on | Shared-memory read cache |
| `pg_textsearch.memory_limit` | 2GB | Approximate cache admission budget; 0 removes the limit |

Parallel builds require at least 64 MB of `maintenance_work_mem` and available parallel maintenance workers. For upgrades, install the matching binary, restart PostgreSQL and run the extension update according to the release instructions.

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

### Boundaries

The `bm25` access-method name conflicts with `pg_search` and `vchord_bm25`; do not install conflicting providers in the same database. Phrase matching uses conservative heap rechecks where needed. Partition scores use partition-local statistics and may not be comparable across partitions. Very long tokens inherit PostgreSQL text-search limits. Fixed LWLock tranche IDs can also conflict with another extension and produce misleading wait-event names.

### Version 1.5.0 Queries and Maintenance

Boolean filtering now accepts PostgreSQL `tsquery`, including AND/OR/NOT, pure-negative, prefix, weight and phrase conditions. Phrase/weight conditions may require heap rechecks. PostgreSQL 19 support is beta. Optional managed background compaction requires `pg_durable` 0.2.8+ in the same configured database, preloaded and configured according to upstream; inline/manual modes do not require it. `pg_textsearch.allow_rls` defaults on: BM25 statistics include all indexed rows, including rows hidden by RLS. Turning it off prevents new/rebuilt BM25 indexes on protected tables and related RLS enablement; it does not disable existing indexes.

### Upgrade to 1.5.1

Version 1.5.1 fixes concurrent spill/force-merge truncation corruption, expression-index and array/domain handling, and maintenance in databases without the extension. Install the matching library, restart PostgreSQL, then run the extension update in each database. The 1.5.0-to-1.5.1 migration checks that the library was preloaded; it adds no SQL object or explicit index-format migration.

```sql
ALTER EXTENSION pg_textsearch UPDATE TO '1.5.1';
```

The cache budget is approximate and concurrent work can exceed it. Hot standbys serving BM25 queries require `hot_standby_feedback = on` to delay physical page reuse while snapshots are active. Mutating maintenance helpers require index ownership and cannot run during recovery; published physical merge replacements are not undone by transaction rollback.
