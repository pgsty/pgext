---
title: "vectorscale"
linkTitle: "vectorscale"
description: "Advanced indexing for vector data with DiskANN"
weight: 1820
categories: ["RAG"]
languages: ["Rust"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pgvectorscale**](https://github.com/timescale/pgvectorscale) : Advanced indexing for vector data with DiskANN


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1820** | {{< badge content="vectorscale" link="https://github.com/timescale/pgvectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | `0.9.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vector" >}} {{< ext "vchord" >}} {{< ext "pgcontext" >}} {{< ext "vectorize" >}} {{< ext "pg_rrf" >}} {{< ext "pg_search" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_bestmatch" >}} {{< ext "pgml" >}} {{< ext "pg4ml" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgvectorscale` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "pgvectorscale_18" "green" >}} {{< bg "17" "pgvectorscale_17" "green" >}} {{< bg "16" "pgvectorscale_16" "green" >}} {{< bg "15" "pgvectorscale_15" "green" >}} {{< bg "14" "pgvectorscale_14" "green" >}} | `pgvectorscale_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.1` | {{< bg "18" "postgresql-18-pgvectorscale" "green" >}} {{< bg "17" "postgresql-17-pgvectorscale" "green" >}} {{< bg "16" "postgresql-16-pgvectorscale" "green" >}} {{< bg "15" "postgresql-15-pgvectorscale" "green" >}} {{< bg "14" "postgresql-14-pgvectorscale" "green" >}} | `postgresql-$v-pgvectorscale` | `postgresql-$v-pgvector` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.1" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_18` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_18-0.9.1-1PGSTY.el8.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 919.5 KiB | [pgvectorscale_18-0.9.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_18-0.9.1-1PGSTY.el8.aarch64.rpm) |
| `pgvectorscale_18` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_18-0.9.1-1PGSTY.el9.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 987.3 KiB | [pgvectorscale_18-0.9.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_18-0.9.1-1PGSTY.el9.aarch64.rpm) |
| `pgvectorscale_18` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_18-0.9.1-1PGSTY.el10.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 967.3 KiB | [pgvectorscale_18-0.9.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_18-0.9.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 903.4 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 743.9 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 903.9 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 744.2 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1001.4 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 879.1 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 992.3 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 869.8 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 988.3 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 868.1 KiB | [postgresql-18-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_17` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_17-0.9.1-1PGSTY.el8.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 916.9 KiB | [pgvectorscale_17-0.9.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_17-0.9.1-1PGSTY.el8.aarch64.rpm) |
| `pgvectorscale_17` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_17-0.9.1-1PGSTY.el9.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 983.1 KiB | [pgvectorscale_17-0.9.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_17-0.9.1-1PGSTY.el9.aarch64.rpm) |
| `pgvectorscale_17` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_17-0.9.1-1PGSTY.el10.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 966.9 KiB | [pgvectorscale_17-0.9.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_17-0.9.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 902.1 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 742.8 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 901.7 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 741.9 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1001.3 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 877.4 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 989.2 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 866.9 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 984.7 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 865.1 KiB | [postgresql-17-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_16` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_16-0.9.1-1PGSTY.el8.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 915.3 KiB | [pgvectorscale_16-0.9.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_16-0.9.1-1PGSTY.el8.aarch64.rpm) |
| `pgvectorscale_16` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_16-0.9.1-1PGSTY.el9.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 982.5 KiB | [pgvectorscale_16-0.9.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_16-0.9.1-1PGSTY.el9.aarch64.rpm) |
| `pgvectorscale_16` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_16-0.9.1-1PGSTY.el10.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 966.1 KiB | [pgvectorscale_16-0.9.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_16-0.9.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 900.2 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 740.6 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 900.5 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 741.4 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1000.0 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 876.5 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 991.2 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 866.2 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 984.6 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 864.7 KiB | [postgresql-16-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_15` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pgvectorscale_15-0.9.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_15-0.9.1-1PGSTY.el8.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 906.9 KiB | [pgvectorscale_15-0.9.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_15-0.9.1-1PGSTY.el8.aarch64.rpm) |
| `pgvectorscale_15` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_15-0.9.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_15-0.9.1-1PGSTY.el9.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 973.0 KiB | [pgvectorscale_15-0.9.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_15-0.9.1-1PGSTY.el9.aarch64.rpm) |
| `pgvectorscale_15` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_15-0.9.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_15-0.9.1-1PGSTY.el10.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 961.6 KiB | [pgvectorscale_15-0.9.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_15-0.9.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 895.3 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 736.7 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 895.5 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 736.9 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 991.6 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 871.0 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 981.9 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 861.7 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 978.2 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 858.2 KiB | [postgresql-15-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_14` | `0.9.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.0 MiB | [pgvectorscale_14-0.9.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_14-0.9.1-1PGSTY.el8.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 903.6 KiB | [pgvectorscale_14-0.9.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_14-0.9.1-1PGSTY.el8.aarch64.rpm) |
| `pgvectorscale_14` | `0.9.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_14-0.9.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_14-0.9.1-1PGSTY.el9.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 969.4 KiB | [pgvectorscale_14-0.9.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_14-0.9.1-1PGSTY.el9.aarch64.rpm) |
| `pgvectorscale_14` | `0.9.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_14-0.9.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_14-0.9.1-1PGSTY.el10.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 960.4 KiB | [pgvectorscale_14-0.9.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_14-0.9.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 892.3 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 733.8 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 892.0 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 734.9 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 988.4 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 868.2 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 978.7 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 858.4 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 977.1 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 855.8 KiB | [postgresql-14-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.1-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pgvectorscale" title="Repository" icon="github" subtitle="github.com/timescale/pgvectorscale" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgvectorscale-0.9.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgvectorscale;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgvectorscale;		# install via package name, for the active PG version
pig install vectorscale;		# install by extension name, for the current active PG version

pig install vectorscale -v 18;   # install for PG 18
pig install vectorscale -v 17;   # install for PG 17
pig install vectorscale -v 16;   # install for PG 16
pig install vectorscale -v 15;   # install for PG 15
pig install vectorscale -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vectorscale CASCADE; -- requires vector
```

## Usage

Sources:

- [0.9.1 README](https://github.com/timescale/pgvectorscale/blob/0.9.1/README.md)
- [Control and dependency](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/vectorscale.control)
- [0.9.1 migration SQL](https://github.com/timescale/pgvectorscale/blob/0.9.1/pgvectorscale/sql/vectorscale--0.9.0--0.9.1.sql)
- [0.9.1 security and upgrade notes](https://github.com/timescale/pgvectorscale/releases/tag/0.9.1)

`vectorscale` adds the StreamingDiskANN approximate vector index to pgvector. Its `diskann` access method supports L2, inner product, cosine distance, and label filtering. It requires `vector`; creating the extension requires superuser privileges. Version 0.9.1 validates vector types, dimensions, and stored datum layouts to address crashes, memory disclosure, and out-of-bounds writes.

### Core Workflow

Declare a concrete vector dimension and select the operator class matching the query distance:

```sql
CREATE EXTENSION vectorscale CASCADE;
CREATE TABLE documents (
    id bigserial PRIMARY KEY,
    contents text,
    embedding vector(3),
    labels smallint[]
);
INSERT INTO documents(contents, embedding, labels)
VALUES ('PostgreSQL search', '[1,2,3]', ARRAY[1,3]::smallint[]);
CREATE INDEX documents_diskann ON documents
USING diskann (embedding vector_cosine_ops, labels);

SELECT id, contents FROM documents
WHERE labels && ARRAY[1]::smallint[]
ORDER BY embedding <=> '[1,2,3]'::vector
LIMIT 10;
```

Use `vector_l2_ops` with `<->`, `vector_ip_ops` with `<#>`, or `vector_cosine_ops` with `<=>`. Labels use `smallint[]`, and `&&` means any requested label overlaps. Ordinary WHERE conditions are also supported, but selective filters can reduce the returned result count and need workload-specific testing.

### Tuning and Ordering

`diskann.query_search_list_size` controls extra graph-search candidates (default 100); `diskann.query_rescore` controls exact rescoring (default 50, 0 disables it):

```sql
SET diskann.query_search_list_size = 200;
SET diskann.query_rescore = 100;
```

Build options include `storage_layout`, `num_neighbors`, `search_list_size`, and `num_dimensions`. Compressed memory-optimized storage is the default. Increase maintenance memory only after considering concurrent builds and dataset size. Parallel builds require the supported compression layout and do not support label columns in this version.

DiskANN returns relaxed distance ordering. Sort a materialized result set when strict ordering is required; this reorders the retrieved candidates without making approximate search exhaustive. Null vectors are not indexed, null labels act as empty arrays, and null array elements are ignored. Index creation on UNLOGGED tables is unsupported.

### Upgrade to 0.9.1

Install matching extension files and update each database:

```sql
ALTER EXTENSION vectorscale UPDATE TO '0.9.1';
```

The upgrade binds operator classes to pgvector's actual installation schema and checks existing bindings. If an existing operator class references the wrong type or operators, the upgrade aborts; follow the release instructions to drop the affected operator class and recreate extension objects, accounting for dependent indexes.

**DiskANN now requires a valid `vector(N)` column type.** An unconstrained vector column cannot be indexed. Existing indexes with invalid persisted dimensions raise errors during scans, inserts, and vacuum. Correct the column type, then **drop and recreate the affected indexes**: `REINDEX` cannot repair this condition. Valid existing indexes do not need a blanket rebuild merely because 0.9.1 was installed. The SQL surface is not relocatable after extension creation.
