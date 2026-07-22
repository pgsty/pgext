---
title: "vectorscale"
linkTitle: "vectorscale"
description: "Advanced indexing for vector data with DiskANN"
weight: 1820
categories: ["RAG"]
width: full
---

[**pgvectorscale**](https://github.com/timescale/pgvectorscale) : Advanced indexing for vector data with DiskANN


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1820** | {{< badge content="vectorscale" link="https://github.com/timescale/pgvectorscale" >}} | {{< ext "vectorscale" "pgvectorscale" >}} | `0.9.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vchord" >}} {{< ext "vectorize" >}} {{< ext "pg_summarize" >}} {{< ext "pg_tiktoken" >}} {{< ext "pg4ml" >}} {{< ext "pgml" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_similarity" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgvectorscale` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.0` | {{< bg "18" "pgvectorscale_18" "green" >}} {{< bg "17" "pgvectorscale_17" "green" >}} {{< bg "16" "pgvectorscale_16" "green" >}} {{< bg "15" "pgvectorscale_15" "green" >}} {{< bg "14" "pgvectorscale_14" "green" >}} | `pgvectorscale_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.9.0` | {{< bg "18" "postgresql-18-pgvectorscale" "green" >}} {{< bg "17" "postgresql-17-pgvectorscale" "green" >}} {{< bg "16" "postgresql-16-pgvectorscale" "green" >}} {{< bg "15" "postgresql-15-pgvectorscale" "green" >}} {{< bg "14" "postgresql-14-pgvectorscale" "green" >}} | `postgresql-$v-pgvectorscale` | `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "pgvectorscale_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-18-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-17-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-16-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-15-pgvectorscale : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.9.0" "postgresql-14-pgvectorscale : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_18` | `0.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_18-0.9.0-3PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 977.8 KiB | [pgvectorscale_18-0.9.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_18-0.9.0-3PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_18` | `0.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_18-0.9.0-3PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pgvectorscale_18-0.9.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_18-0.9.0-3PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_18` | `0.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_18-0.9.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_18-0.9.0-3PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_18` | `0.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1018.5 KiB | [pgvectorscale_18-0.9.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_18-0.9.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 888.7 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 731.8 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 889.0 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 733.5 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 987.5 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 867.4 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 976.6 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 856.8 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 973.2 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgvectorscale` | `0.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 855.8 KiB | [postgresql-18-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-18-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_17` | `0.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_17-0.9.0-3PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 974.7 KiB | [pgvectorscale_17-0.9.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_17-0.9.0-3PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_17` | `0.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_17-0.9.0-3PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pgvectorscale_17-0.9.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_17-0.9.0-3PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_17` | `0.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_17-0.9.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_17-0.9.0-3PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_17` | `0.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1016.4 KiB | [pgvectorscale_17-0.9.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_17-0.9.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 887.0 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 731.4 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 886.8 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 731.7 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 984.9 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 865.5 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 973.2 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 854.4 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 969.8 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgvectorscale` | `0.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 853.1 KiB | [postgresql-17-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-17-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_16` | `0.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_16-0.9.0-3PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 973.2 KiB | [pgvectorscale_16-0.9.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_16-0.9.0-3PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_16` | `0.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_16-0.9.0-3PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pgvectorscale_16-0.9.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_16-0.9.0-3PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_16` | `0.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_16-0.9.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_16-0.9.0-3PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_16` | `0.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1016.7 KiB | [pgvectorscale_16-0.9.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_16-0.9.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 886.8 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 731.6 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 886.5 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 731.9 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 985.1 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 864.9 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 975.1 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 854.6 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 968.7 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgvectorscale` | `0.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 853.5 KiB | [postgresql-16-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-16-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_15` | `0.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_15-0.9.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_15-0.9.0-3PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 963.6 KiB | [pgvectorscale_15-0.9.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_15-0.9.0-3PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_15` | `0.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_15-0.9.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_15-0.9.0-3PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.0 MiB | [pgvectorscale_15-0.9.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_15-0.9.0-3PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_15` | `0.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_15-0.9.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_15-0.9.0-3PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_15` | `0.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1012.6 KiB | [pgvectorscale_15-0.9.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_15-0.9.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 879.9 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 726.2 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 879.2 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 727.0 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 975.4 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 857.7 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 966.1 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 849.4 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 962.9 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgvectorscale` | `0.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 846.2 KiB | [postgresql-15-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-15-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgvectorscale_14` | `0.9.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_14-0.9.0-3PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgvectorscale_14-0.9.0-3PIGSTY.el8.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 960.4 KiB | [pgvectorscale_14-0.9.0-3PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgvectorscale_14-0.9.0-3PIGSTY.el8.aarch64.rpm) |
| `pgvectorscale_14` | `0.9.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_14-0.9.0-3PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgvectorscale_14-0.9.0-3PIGSTY.el9.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1020.6 KiB | [pgvectorscale_14-0.9.0-3PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgvectorscale_14-0.9.0-3PIGSTY.el9.aarch64.rpm) |
| `pgvectorscale_14` | `0.9.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.1 MiB | [pgvectorscale_14-0.9.0-3PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgvectorscale_14-0.9.0-3PIGSTY.el10.x86_64.rpm) |
| `pgvectorscale_14` | `0.9.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1009.2 KiB | [pgvectorscale_14-0.9.0-3PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgvectorscale_14-0.9.0-3PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 876.1 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 723.6 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 876.8 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 723.4 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 970.8 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 854.3 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 960.7 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 845.1 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 958.6 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgvectorscale` | `0.9.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 842.6 KiB | [postgresql-14-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgvectorscale/postgresql-14-pgvectorscale_0.9.0-3PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pgvectorscale" title="Repository" icon="github" subtitle="github.com/timescale/pgvectorscale" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgvectorscale-0.9.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgvectorscale;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

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

`vectorscale` extends `pgvector` with the StreamingDiskANN index access method for approximate nearest-neighbor search. The example below creates a vector table, builds a `diskann` index, and runs a distance-ordered query.

```sql
CREATE EXTENSION vectorscale CASCADE;

CREATE TABLE IF NOT EXISTS document_embedding  (
    id BIGINT PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    metadata JSONB,
    contents TEXT,
    embedding VECTOR(1536)
);
  
CREATE INDEX document_embedding_idx ON document_embedding
USING diskann (embedding);

SELECT *
FROM document_embedding
ORDER BY embedding <=> $1
LIMIT 10
```
