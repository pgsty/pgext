---
title: "pg_vectorsearch"
linkTitle: "pg_vectorsearch"
description: "Approximate vector search with PRISM indexes and RaBitQ quantization"
weight: 1990
categories: ["RAG"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_vectorsearch**](https://github.com/timescale/pg_vectorsearch) : Approximate vector search with PRISM indexes and RaBitQ quantization


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1990** | {{< badge content="pg_vectorsearch" link="https://github.com/timescale/pg_vectorsearch" >}} | {{< ext "pg_vectorsearch" >}} | `0.2.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "vector" >}} {{< ext "vectorscale" >}} |

> [!Note] Alpha; PG18 packages; OpenBLAS required; no pgvector dependency; no in-place RC upgrade.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} | `pg_vectorsearch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "19" "pg_vectorsearch_19" "green" >}} {{< bg "18" "pg_vectorsearch_18" "green" >}} {{< bg "17" "pg_vectorsearch_17" "red" >}} {{< bg "16" "pg_vectorsearch_16" "red" >}} {{< bg "15" "pg_vectorsearch_15" "red" >}} | `pg_vectorsearch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "19" "postgresql-19-vectorsearch" "green" >}} {{< bg "18" "postgresql-18-vectorsearch" "green" >}} {{< bg "17" "postgresql-17-vectorsearch" "red" >}} {{< bg "16" "postgresql-16-vectorsearch" "red" >}} {{< bg "15" "postgresql-15-vectorsearch" "red" >}} | `postgresql-$v-vectorsearch` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_vectorsearch_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_vectorsearch_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "pg_vectorsearch_18 : AVAIL 2" "green" >}} | {{< bg "N/A" "pg_vectorsearch_17 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_vectorsearch_15 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-19-vectorsearch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-vectorsearch : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-17-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-16-vectorsearch : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-vectorsearch : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorsearch_19` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 579.4 KiB | [pg_vectorsearch_19-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorsearch_19-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_vectorsearch_19` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 510.5 KiB | [pg_vectorsearch_19-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorsearch_19-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_vectorsearch_19` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 585.7 KiB | [pg_vectorsearch_19-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorsearch_19-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_vectorsearch_19` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 518.5 KiB | [pg_vectorsearch_19-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorsearch_19-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-vectorsearch` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 549.2 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 477.5 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 548.9 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 479.9 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 583.4 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 522.8 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 574.7 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 509.5 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 575.0 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-vectorsearch` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 509.6 KiB | [postgresql-19-vectorsearch_0.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vectorsearch/postgresql-19-vectorsearch_0.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_vectorsearch_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 584.9 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorsearch_18-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 533.8 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_vectorsearch_18-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 505.4 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorsearch_18-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 453.7 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_vectorsearch_18-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `pg_vectorsearch_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 578.0 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorsearch_18-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 528.7 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_vectorsearch_18-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 510.5 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorsearch_18-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 458.6 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_vectorsearch_18-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `pg_vectorsearch_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 583.8 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorsearch_18-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 533.0 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_vectorsearch_18-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `pg_vectorsearch_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 518.1 KiB | [pg_vectorsearch_18-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorsearch_18-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `pg_vectorsearch_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 465.4 KiB | [pg_vectorsearch_18-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_vectorsearch_18-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-vectorsearch` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 547.6 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 477.2 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 547.1 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 480.5 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 581.7 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 522.6 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 572.7 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 509.5 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 572.5 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-vectorsearch` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 509.6 KiB | [postgresql-18-vectorsearch_0.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-vectorsearch/postgresql-18-vectorsearch_0.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/timescale/pg_vectorsearch" title="Repository" icon="github" subtitle="github.com/timescale/pg_vectorsearch" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_vectorsearch-0.2.0.tar.gz meson-0.63.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_vectorsearch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_vectorsearch;		# install via package name, for the active PG version

pig install pg_vectorsearch -v 19;   # install for PG 19
pig install pg_vectorsearch -v 18;   # install for PG 18

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_vectorsearch;
```

## Usage

Sources:

- [README.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/README.md)
- [sql/pg_vectorsearch.sql](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/sql/pg_vectorsearch.sql)
- [docs/maintenance.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/maintenance.md)
- [docs/tuning.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/docs/tuning.md)
- [CHANGELOG.md](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/CHANGELOG.md)
- [VERSION](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/VERSION)
- [meson.build](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/meson.build)
- [pg_vectorsearch.control.in](https://github.com/timescale/pg_vectorsearch/blob/v0.2.0/pg_vectorsearch.control.in)

`pg_vectorsearch` 0.2.0 is an alpha extension for approximate nearest-neighbor search with the PRISM index. It supports PostgreSQL 18 and 19, native `vec32` / `vec16` types and optional pgvector interoperability. It is intended for testing, with evolving behavior and maintenance requirements.

### Core Workflow

```sql
CREATE EXTENSION pg_vectorsearch;
CREATE TABLE vectorsearch_items (
    id serial PRIMARY KEY,
    embedding vec32(3) STORAGE PLAIN
);
INSERT INTO vectorsearch_items (embedding)
SELECT ARRAY[i::real, (i % 7)::real, (i % 5)::real]::vec32
FROM generate_series(1, 1000) AS i;
CREATE INDEX vectorsearch_items_prism ON vectorsearch_items
USING prism (embedding vec32_l2_ops);
SET prism.nprobe = 40;
SELECT id FROM vectorsearch_items
ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

### Operational Boundaries

Install as a superuser; no preload is required and the extension is not relocatable. The `prism` access method has `vec32_l2_ops`, `vec32_ip_ops`, `vec32_cosine_ops` and corresponding vec16 operator classes; `<->`, `<#>` and `<=>` express L2, negative inner product and cosine distance. `prism.nprobe` trades recall for query work; zero selects automatic tuning, not exact recall.

Load representative data before building the index. L2 and cosine indexes support up to 8192 dimensions: any dimension through 1968, then a power of two times an odd factor at most eight. Inner-product indexes remain limited to 1968 for vec32 and 3584 for vec16. Expression and unlogged indexes are unsupported. Keep `prism.leaf_refine_threshold` at zero. Filtering can return fewer rows than LIMIT, and candidate buffers bounded by `work_mem` can truncate large or unbounded results. Measure against an exact baseline.

Inspect `prism_index_settings` and follow the maintenance guide for churn. `prism_rebalance` supports only flat RaBitQ centroid-page indexes with `centroid_fastscan` off, not the default multilevel format. DELETE leaves tombstones, and retired pages are reclaimed only by REINDEX. There is no in-place upgrade from the earlier release candidates.

### Upgrade to 0.2.0

Install matching files, start a fresh session that has not loaded the old library, update the SQL extension and rebuild every PRISM index:

```sql
ALTER EXTENSION pg_vectorsearch UPDATE TO '0.2.0';
REINDEX INDEX vectorsearch_items_prism;
```

Old index format 0x06 is rejected by the 0x07 reader. The two library versions cannot share one session. `vec32_norm` is renamed to `l2_norm`; the vec16 overload of `vec32_dims` becomes `vec16_dims`, while `vector_dims` provides the pgvector-compatible interface. Update direct callers of renamed functions; dependent views retain function OIDs.

`posting_reference_format` defaults to float16. A reference element beyond ±65504 causes a build or split error; use float32 when needed. The vec16 type now rejects out-of-range input. Existing filter, buffer-size and maintenance limits still apply.
