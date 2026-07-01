---
title: "pgmnemo"
linkTitle: "pgmnemo"
description: "Provenance-gated vector memory for LLM agents in PostgreSQL"
weight: 1900
categories: ["RAG"]
width: full
---

[**pgmnemo**](https://github.com/pgmnemo/pgmnemo) : Provenance-gated vector memory for LLM agents in PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1900** | {{< badge content="pgmnemo" link="https://github.com/pgmnemo/pgmnemo" >}} | {{< ext "pgmnemo" >}} | `0.12.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmnemo` |
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vector" >}} {{< ext "pg_search" >}} {{< ext "pg_ai_query" >}} {{< ext "pg_later" >}} |

> [!Note] SQL-only extension; requires pgvector.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmnemo` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.1` | {{< bg "18" "pgmnemo_18" "green" >}} {{< bg "17" "pgmnemo_17" "green" >}} {{< bg "16" "pgmnemo_16" "green" >}} {{< bg "15" "pgmnemo_15" "green" >}} {{< bg "14" "pgmnemo_14" "green" >}} | `pgmnemo_$v` | `pgvector` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.12.1` | {{< bg "18" "postgresql-18-pgmnemo" "green" >}} {{< bg "17" "postgresql-17-pgmnemo" "green" >}} {{< bg "16" "postgresql-16-pgmnemo" "green" >}} {{< bg "15" "postgresql-15-pgmnemo" "green" >}} {{< bg "14" "postgresql-14-pgmnemo" "green" >}} | `postgresql-$v-pgmnemo` | `pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_18` | `0.12.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 158.4 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_18-0.12.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_18` | `0.12.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 158.4 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_18-0.12.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_18` | `0.12.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.8 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_18-0.12.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_18` | `0.12.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 133.8 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_18-0.12.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_18` | `0.12.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_18-0.12.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_18` | `0.12.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.9 KiB | [pgmnemo_18-0.12.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_18-0.12.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmnemo` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.2 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 124.2 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.2 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 124.2 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.1 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 124.8 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 124.9 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pgmnemo` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-18-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_17` | `0.12.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 158.4 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_17-0.12.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_17` | `0.12.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 158.4 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_17-0.12.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_17` | `0.12.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.8 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_17-0.12.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_17` | `0.12.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 133.8 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_17-0.12.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_17` | `0.12.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_17-0.12.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_17` | `0.12.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.9 KiB | [pgmnemo_17-0.12.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_17-0.12.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmnemo` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.2 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 124.2 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.2 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 124.2 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.1 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 124.8 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 124.9 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pgmnemo` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-17-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_16` | `0.12.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 158.4 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_16-0.12.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_16` | `0.12.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 158.4 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_16-0.12.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_16` | `0.12.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.8 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_16-0.12.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_16` | `0.12.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 133.8 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_16-0.12.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_16` | `0.12.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_16-0.12.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_16` | `0.12.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.9 KiB | [pgmnemo_16-0.12.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_16-0.12.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmnemo` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.2 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 124.2 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.2 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 124.2 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.1 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 124.8 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 124.9 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pgmnemo` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-16-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_15` | `0.12.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 158.4 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_15-0.12.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_15` | `0.12.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 158.4 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_15-0.12.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_15` | `0.12.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.8 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_15-0.12.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_15` | `0.12.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 133.8 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_15-0.12.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_15` | `0.12.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_15-0.12.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_15` | `0.12.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.9 KiB | [pgmnemo_15-0.12.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_15-0.12.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmnemo` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.2 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 124.2 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.2 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 124.2 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.1 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 124.8 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 124.9 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pgmnemo` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-15-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_14` | `0.12.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 158.4 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_14-0.12.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_14` | `0.12.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 158.4 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_14-0.12.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_14` | `0.12.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 133.8 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_14-0.12.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_14` | `0.12.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 133.7 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_14-0.12.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_14` | `0.12.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 134.0 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_14-0.12.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_14` | `0.12.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 133.9 KiB | [pgmnemo_14-0.12.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_14-0.12.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmnemo` | `0.12.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 124.2 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 124.2 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 124.2 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 124.2 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 125.1 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 125.1 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 124.8 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 124.8 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 124.9 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pgmnemo` | `0.12.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 124.9 KiB | [postgresql-14-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.12.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmnemo/pgmnemo" title="Repository" icon="github" subtitle="github.com/pgmnemo/pgmnemo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmnemo-0.12.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmnemo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmnemo;		# install via package name, for the active PG version

pig install pgmnemo -v 18;   # install for PG 18
pig install pgmnemo -v 17;   # install for PG 17
pig install pgmnemo -v 16;   # install for PG 16
pig install pgmnemo -v 15;   # install for PG 15
pig install pgmnemo -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmnemo CASCADE; -- requires vector
```




## Usage

Sources:

- [PGXN pgmnemo 0.12.1](https://pgxn.org/dist/pgmnemo/0.12.1/)
- [pgmnemo README](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/README.md)
- [pgmnemo CHANGELOG](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/CHANGELOG.md)
- [pgmnemo control file](https://github.com/pgmnemo/pgmnemo/blob/v0.12.1/extension/pgmnemo.control)
- [Local package metadata](../db/extension.csv)

`pgmnemo` stores agent memory in PostgreSQL and retrieves it through a single multimodal plan that combines pgvector HNSW search, BM25-style text matching, graph-edge proximity, JSONB metadata filtering, temporal filters, and outcome confidence. It is a SQL-only extension that requires `vector`, installs into schema `pgmnemo`, is marked trusted, and does not require superuser privileges in the v0.12.1 control file.

v0.12.1 keeps the v0.12.0 typed write API and changes `guard_no_test_project()` so project-id floor enforcement is opt-in through `pgmnemo.test_project_floor`; the default `0` disables that floor check.

### Install

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

SELECT pgmnemo.version();
SELECT * FROM pgmnemo.stats();
```

### Ingest Lessons

```sql
SELECT pgmnemo.ingest(
  p_role        := 'developer',
  p_project_id  := 1,
  p_topic       := 'security',
  p_lesson_text := 'Rotate JWT signing keys after a key-compromise incident.',
  p_importance  := 4,
  p_embedding   := NULL,
  p_commit_sha  := 'abc1234',
  p_metadata    := '{"source":"incident-runbook"}'::jsonb
);
```

`pgmnemo.ingest()` is the baseline write path. It applies the provenance gate, validates 1024-dimensional embeddings when supplied, truncates long lesson text according to `pgmnemo.max_query_text_chars`, and stamps `verified_at` when provenance exists.

### Provenance Gate

```sql
SHOW pgmnemo.gate_strict;

SET pgmnemo.gate_strict = 'enforce';
SET pgmnemo.gate_strict = 'warn';
SET pgmnemo.gate_strict = 'off';
```

When `pgmnemo.gate_strict = 'enforce'`, writes without either `commit_sha` or `artifact_hash` are rejected. `warn` accepts the write with an audit warning, and `off` disables the gate.

The v0.12.1 test-project guard is opt-in:

```sql
SHOW pgmnemo.test_project_floor;  -- default 0
SET pgmnemo.test_project_floor = '1000000';
```

With the default `0`, the production-floor check is disabled. Set a positive value in test harnesses that reserve low project IDs.

### Recall

```sql
SELECT lesson_id, topic, lesson_text, score
FROM pgmnemo.recall_fast(
  '<1024-dimensional vector literal>'::vector(1024),
  10,
  'developer',
  1,
  'dag-2026-abc',
  ARRAY['note', 'fact']
);

SELECT lesson_id, topic, score, vec_score, bm25_score, rrf_score, match_confidence
FROM pgmnemo.recall_hybrid(
  '<1024-dimensional vector literal>'::vector(1024),
  'JWT rotation key compromise',
  10,
  'developer',
  1,
  0.4,
  0.4,
  60,
  'dag-2026-abc',
  ARRAY['note', 'fact']
);
```

`recall_lessons()` routes to hybrid recall when both text and embedding are present and `pgmnemo.disable_hybrid` is false. `exclude_dag_id` avoids retrieving records from the same workflow run, and `p_content_types` filters typed memories.

Recall functions update `last_recalled_at` and `recall_count` by default. Disable that side effect for read-only analysis:

```sql
SET pgmnemo.track_recall_recency = 'off';
```

### Navigate and Expand

```sql
SELECT *
FROM pgmnemo.navigate_locate(
  NULL::vector(1024),
  'JWT rotation',
  2000,
  '{"topic":"security"}'::jsonb,
  1
);

SELECT *
FROM pgmnemo.navigate_locate_dispatch(
  query_text            := 'JWT rotation',
  content_type_dispatch := 'entity',
  project_id_filter     := 1,
  token_budget_chars    := 2000
);

SELECT *
FROM pgmnemo.navigate_expand_typed(
  ARRAY[1001, 1002]::bigint[],
  graph_expand_depth := 1,
  relation_types     := ARRAY['CAUSED_BY', 'DERIVED_FROM']
);
```

Use `navigate_locate()` to find candidate IDs within a character budget, then use expand functions to fetch full content and graph neighbors only for the selected IDs.

### Typed Writes

```sql
SELECT pgmnemo.canonical_slug('concept', 'JWT Rotation');

SELECT *
FROM pgmnemo.remember_fact(
  p_role            := 'developer',
  p_entity_key      := 'concept:jwt_rotation',
  p_property        := 'runbook',
  p_value           := 'Rotate signing secrets within 24 hours after compromise.',
  p_confidence      := 0.82,
  p_has_contact_pii := false,
  p_source_type     := 'agent_authored',
  p_project_id      := 1,
  p_commit_sha      := 'abc1234'
);

SELECT pgmnemo.remember_event(
  p_role        := 'developer',
  p_entity_key  := 'concept:jwt_rotation',
  p_event_label := 'incident_response',
  p_event_body  := 'Rotation procedure validated during the July drill.',
  p_project_id  := 1,
  p_commit_sha  := 'abc1234'
);

SELECT pgmnemo.remember_relation(
  p_role          := 'developer',
  p_from_key      := 'concept:jwt_rotation',
  p_to_key        := 'concept:key_compromise',
  p_relation_type := 'MITIGATES',
  p_project_id    := 1,
  p_commit_sha    := 'abc1234'
);
```

`remember_fact()` supersedes prior active facts for the same entity/property, `remember_event()` is append-only, and `remember_relation()` writes relation memories and graph edges.

### Edges, Reinforcement, and Maintenance

```sql
SELECT pgmnemo.add_edge(1001, 1002, 'CAUSED_BY', 0.85, '{"run_id":7320}'::jsonb);

SELECT pgmnemo.reinforce(1001, 'success');
SELECT pgmnemo.reinforce(ARRAY[1001, 1002]::bigint[], 'failure');

SELECT pgmnemo.reembed(1001, '<1024-dimensional vector literal>'::vector(1024));
SELECT pgmnemo.recompute_content(1001, 'Updated lesson text.');
```

Useful settings include `pgmnemo.gate_strict`, `pgmnemo.include_unverified`, `pgmnemo.ef_search`, `pgmnemo.disable_hybrid`, `pgmnemo.recency_weight`, `pgmnemo.importance_weight`, `pgmnemo.graph_proximity_weight`, `pgmnemo.temporal_boost`, `pgmnemo.confidence_boost_weight`, `pgmnemo.track_recall_recency`, `pgmnemo.max_query_text_chars`, `pgmnemo.tenant_id`, and `pgmnemo.test_project_floor`.

### Caveats

- `pgmnemo` requires `vector >= 0.7.0` according to PGXN metadata.
- Embeddings are expected to be 1024-dimensional in the current SQL definition.
- The default provenance gate is intentional. Prefer `warn` over `off` when migrating legacy memory rows.
- Recall functions can write recency metadata; turn off `pgmnemo.track_recall_recency` for read-only diagnostics.
