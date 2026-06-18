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
| **1900** | {{< badge content="pgmnemo" link="https://github.com/pgmnemo/pgmnemo" >}} | {{< ext "pgmnemo" >}} | `0.8.3` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


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
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmnemo` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "pgmnemo_18" "green" >}} {{< bg "17" "pgmnemo_17" "green" >}} {{< bg "16" "pgmnemo_16" "green" >}} {{< bg "15" "pgmnemo_15" "green" >}} {{< bg "14" "pgmnemo_14" "green" >}} | `pgmnemo_$v` | `pgvector` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.8.3` | {{< bg "18" "postgresql-18-pgmnemo" "green" >}} {{< bg "17" "postgresql-17-pgmnemo" "green" >}} {{< bg "16" "postgresql-16-pgmnemo" "green" >}} {{< bg "15" "postgresql-15-pgmnemo" "green" >}} {{< bg "14" "postgresql-14-pgmnemo" "green" >}} | `postgresql-$v-pgmnemo` | `pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_18` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 95.3 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_18-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_18` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.2 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_18-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_18` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_18-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_18` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.2 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_18-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_18` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.5 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_18-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_18` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.4 KiB | [pgmnemo_18-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_18-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmnemo` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.8 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 83.6 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.6 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 83.6 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pgmnemo` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-18-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_17` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 95.3 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_17-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_17` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.2 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_17-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_17` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_17-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_17` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.2 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_17-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_17` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.5 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_17-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_17` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.4 KiB | [pgmnemo_17-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_17-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmnemo` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.8 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 83.6 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.6 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 83.6 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pgmnemo` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-17-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_16` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 95.3 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_16-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_16` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.3 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_16-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_16` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_16-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_16` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.2 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_16-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_16` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.5 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_16-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_16` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.4 KiB | [pgmnemo_16-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_16-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmnemo` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.8 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 83.6 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.6 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 83.6 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pgmnemo` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-16-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-16-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_15` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 95.3 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_15-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_15` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.2 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_15-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_15` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.2 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_15-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_15` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.2 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_15-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_15` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.5 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_15-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_15` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.4 KiB | [pgmnemo_15-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_15-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmnemo` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.8 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.9 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.9 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 83.6 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.6 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 83.6 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pgmnemo` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-15-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-15-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_14` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 95.3 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_14-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_14` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 95.2 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_14-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_14` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 91.3 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_14-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_14` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 91.2 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_14-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_14` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 91.5 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_14-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_14` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 91.4 KiB | [pgmnemo_14-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_14-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmnemo` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 81.8 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 81.8 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 81.8 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 81.8 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 83.9 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 83.9 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 83.6 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 83.6 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 83.6 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pgmnemo` | `0.8.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 83.6 KiB | [postgresql-14-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-14-pgmnemo_0.8.3-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgmnemo/pgmnemo" title="Repository" icon="github" subtitle="github.com/pgmnemo/pgmnemo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmnemo-0.8.3.tar.gz" >}}
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

Source: [pgmnemo v0.8.3 README](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/README.md), [Usage Guide](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/docs/USAGE.md), [extension control file](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo.control), [SQL definition](https://github.com/pgmnemo/pgmnemo/blob/v0.8.3/extension/pgmnemo--0.8.3.sql).

## Usage

`pgmnemo` stores provenance-gated agent lessons in PostgreSQL and retrieves them through vector, BM25-style text, graph-edge, JSONB metadata, and relational filters. The extension control file requires `vector`, so `pgvector` must be available before creating `pgmnemo`. The local package metadata targets PostgreSQL 14-18.

### Create and Ingest Lessons

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

SELECT pgmnemo.ingest(
  p_role        := 'developer',
  p_project_id  := 1,
  p_topic       := 'security',
  p_lesson_text := 'Rotate JWT secrets after any key-compromise incident.',
  p_importance  := 4,
  p_embedding   := NULL,
  p_commit_sha  := 'abc1234',
  p_metadata    := '{"source":"incident-runbook"}'::jsonb
);
```

`pgmnemo.ingest()` is the preferred write path. It validates the 1024-dimensional embedding when supplied, stamps verified rows when provenance is present, and applies the provenance gate.

### Provenance Gate

```sql
SHOW pgmnemo.gate_strict;

SET pgmnemo.gate_strict = 'warn';
SET pgmnemo.gate_strict = 'enforce';
```

`pgmnemo.gate_strict` accepts `enforce`, `warn`, or `off`. In the default enforced mode, inserts fail when both `p_commit_sha` and `p_artifact_hash` are NULL. `pgmnemo.include_unverified` is separate: it controls whether unverified rows are eligible for recall, not whether writes are allowed.

### Recall

```sql
-- Text-only recall.
SELECT topic, lesson_text, score
FROM pgmnemo.recall_lessons(
  NULL::vector(1024),
  5,
  'developer',
  1,
  'JWT secret rotation'
);

-- Hybrid vector and text recall.
SELECT lesson_id, topic, score, vec_score, bm25_score, rrf_score
FROM pgmnemo.recall_hybrid(
  '<1024-dimensional vector literal>'::vector(1024),
  'JWT rotation key compromise',
  10,
  'developer',
  1
);
```

Hybrid routing in `recall_lessons()` requires `pgmnemo.disable_hybrid` to be off, non-empty `query_text`, and a non-NULL embedding. Use `recall_hybrid()` directly when you want explicit diagnostic scores.

### Navigation and Expansion

```sql
SELECT *
FROM pgmnemo.navigate_locate(
  NULL::vector(1024),
  'JWT rotation',
  10,
  'developer',
  1,
  '{"topic":"security"}'::jsonb,
  2000
);

SELECT *
FROM pgmnemo.navigate_expand(
  ARRAY[1001, 1002]::bigint[],
  include_edges := true
);
```

`navigate_locate()` returns ranked lesson IDs and short previews within a character budget. `navigate_expand()` fetches selected full lessons and optional graph neighbors after the caller chooses which IDs are worth expanding.

### Edges and Outcome Learning

```sql
SELECT pgmnemo.add_edge(1001, 1002, 'CAUSED_BY', 0.85, '{"run_id":7320}'::jsonb);

SELECT pgmnemo.reinforce(1001, 'success');
SELECT pgmnemo.reinforce(1002, 'failure');
```

`pgmnemo.add_edge()` is the idempotent helper for lesson relationships. `reinforce()` adjusts confidence after observed outcomes and feeds later match confidence.

### Maintenance and GUCs

```sql
SELECT * FROM pgmnemo.stats();

SELECT pgmnemo.reembed(
  p_lesson_id  := 1001,
  p_new_vector := '<1024-dimensional vector literal>'::vector(1024)
);

SELECT pgmnemo.recompute_content(
  p_lesson_id       := 1001,
  p_new_lesson_text := 'Rotate JWT secrets within 24 hours after compromise.'
);
```

Useful settings include `pgmnemo.gate_strict`, `pgmnemo.include_unverified`, `pgmnemo.ef_search`, `pgmnemo.disable_hybrid`, `pgmnemo.recency_weight`, `pgmnemo.importance_weight`, `pgmnemo.graph_proximity_weight`, `pgmnemo.temporal_boost`, and `pgmnemo.max_query_text_chars`.
