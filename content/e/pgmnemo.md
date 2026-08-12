---
title: "pgmnemo"
linkTitle: "pgmnemo"
description: "Single-plan multimodal agent memory for PostgreSQL"
weight: 1950
categories: ["RAG"]
width: full
---

[**pgmnemo**](https://github.com/pgmnemo/pgmnemo) : Single-plan multimodal agent memory for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1950** | {{< badge content="pgmnemo" link="https://github.com/pgmnemo/pgmnemo" >}} | {{< ext "pgmnemo" >}} | `0.16.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmnemo` |
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "pgcontext" >}} {{< ext "vector" >}} {{< ext "vectorize" >}} {{< ext "pgml" >}} {{< ext "pg4ml" >}} {{< ext "pg_summarize" >}} {{< ext "provsql" >}} {{< ext "pg_rrf" >}} {{< ext "pg_search" >}} |

> [!Note] SQL-only extension requiring pgvector 0.7.0 or newer; upstream 0.16.1 and PIGSTY packages support PostgreSQL 17 and 18; the control file lives under extension/.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pgmnemo` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.1` | {{< bg "18" "pgmnemo_18" "green" >}} {{< bg "17" "pgmnemo_17" "green" >}} {{< bg "16" "pgmnemo_16" "red" >}} {{< bg "15" "pgmnemo_15" "red" >}} {{< bg "14" "pgmnemo_14" "red" >}} | `pgmnemo_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.16.1` | {{< bg "18" "postgresql-18-pgmnemo" "green" >}} {{< bg "17" "postgresql-17-pgmnemo" "green" >}} {{< bg "16" "postgresql-16-pgmnemo" "red" >}} {{< bg "15" "postgresql-15-pgmnemo" "red" >}} {{< bg "14" "postgresql-14-pgmnemo" "red" >}} | `postgresql-$v-pgmnemo` | `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_18 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.16.1" "pgmnemo_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "pgmnemo_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-18-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.16.1" "postgresql-17-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-16-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-15-pgmnemo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.12.1" "postgresql-14-pgmnemo : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_18` | `0.16.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 195.3 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_18-0.16.1-1PIGSTY.el8.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 187.2 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_18-0.15.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_18` | `0.16.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 195.3 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_18-0.16.1-1PIGSTY.el8.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 187.2 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_18-0.15.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_18` | `0.16.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 176.1 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_18-0.16.1-1PIGSTY.el9.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 163.4 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_18-0.15.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_18` | `0.16.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 176.0 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_18-0.16.1-1PIGSTY.el9.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 163.3 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_18-0.15.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_18` | `0.16.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 176.3 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_18-0.16.1-1PIGSTY.el10.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 163.6 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_18-0.15.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_18` | `0.16.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 176.3 KiB | [pgmnemo_18-0.16.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_18-0.16.1-1PIGSTY.el10.noarch.rpm) |
| `pgmnemo_18` | `0.15.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.5 KiB | [pgmnemo_18-0.15.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_18-0.15.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmnemo` | `0.16.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 156.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 156.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 156.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 156.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 166.2 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.2 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 165.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 165.8 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 165.7 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pgmnemo` | `0.16.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 165.7 KiB | [postgresql-18-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-18-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmnemo_17` | `0.16.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 195.3 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_17-0.16.1-1PIGSTY.el8.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 187.2 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmnemo_17-0.15.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmnemo_17` | `0.16.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 195.3 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_17-0.16.1-1PIGSTY.el8.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 187.2 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmnemo_17-0.15.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmnemo_17` | `0.16.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 176.0 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_17-0.16.1-1PIGSTY.el9.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 163.4 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmnemo_17-0.15.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmnemo_17` | `0.16.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 176.0 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_17-0.16.1-1PIGSTY.el9.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 163.3 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmnemo_17-0.15.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmnemo_17` | `0.16.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 176.3 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_17-0.16.1-1PIGSTY.el10.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 163.6 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmnemo_17-0.15.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmnemo_17` | `0.16.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 176.3 KiB | [pgmnemo_17-0.16.1-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_17-0.16.1-1PIGSTY.el10.noarch.rpm) |
| `pgmnemo_17` | `0.15.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 163.5 KiB | [pgmnemo_17-0.15.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmnemo_17-0.15.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmnemo` | `0.16.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 156.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 156.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 156.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 156.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 166.2 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.2 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 165.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 165.8 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 165.6 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pgmnemo` | `0.16.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 165.6 KiB | [postgresql-17-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmnemo/postgresql-17-pgmnemo_0.16.1-1PIGSTY~resolute_all.deb) |

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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmnemo-0.16.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmnemo;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgmnemo;		# install via package name, for the active PG version

pig install pgmnemo -v 18;   # install for PG 18
pig install pgmnemo -v 17;   # install for PG 17

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmnemo CASCADE; -- requires vector
```

## Usage

Sources:

- [pgmnemo v0.16.1 README](https://github.com/pgmnemo/pgmnemo/blob/v0.16.1/README.md)
- [pgmnemo v0.16.1 usage guide](https://github.com/pgmnemo/pgmnemo/blob/v0.16.1/docs/USAGE.md)
- [pgmnemo v0.16.1 SQL reference](https://github.com/pgmnemo/pgmnemo/blob/v0.16.1/docs/SQL_REFERENCE.md)
- [pgmnemo v0.16.1 changelog](https://github.com/pgmnemo/pgmnemo/blob/v0.16.1/CHANGELOG.md)
- [pgmnemo v0.16.1 control file](https://github.com/pgmnemo/pgmnemo/blob/v0.16.1/extension/pgmnemo.control)

pgmnemo stores agent memory in PostgreSQL and retrieves it through vector, BM25-style text, graph, metadata, temporal, provenance, and outcome-confidence signals. It installs into schema pgmnemo, requires the vector extension, and expects 1024-dimensional embeddings in its current SQL API.

Version 0.16.1 retains the 0.14 corpus-maintenance surface and adds situation fingerprints, verified situation recall, entity-key extraction, and entity-centered recall.

### Install

    CREATE EXTENSION IF NOT EXISTS vector;
    CREATE EXTENSION IF NOT EXISTS pgmnemo CASCADE;

    SELECT pgmnemo.version();
    SELECT * FROM pgmnemo.stats();

The v0.16.1 control file marks pgmnemo as trusted, installs it in schema `pgmnemo`, requires `vector`, and is not relocatable.

### Ingest a Lesson

    SELECT pgmnemo.ingest(
      p_role        := 'developer',
      p_project_id  := 1,
      p_topic       := 'security',
      p_lesson_text := 'Rotate signing keys after a compromise.',
      p_importance  := 4,
      p_embedding   := NULL,
      p_commit_sha  := 'abc1234',
      p_metadata    := '{"source":"incident-runbook"}'::jsonb
    );

When pgmnemo.gate_strict is enforce, commit_sha or artifact_hash provenance is required. warn accepts an unverified write with an audit warning; off disables the gate.

### Recall with Confidence Filtering

Hybrid recall combines embedding and text signals:

    SELECT lesson_id, topic, score, match_confidence
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
      ARRAY['note', 'fact'],
      0.40
    );

The final p_min_score argument, added in 0.13.0, removes candidates whose match_confidence is below the threshold before LIMIT is applied. NULL preserves pre-0.13 behavior. The release notes suggest 0.40 as a starting point, not a universal value; calibrate it for the embedding model and feedback quality.

The same p_min_score concept is available in recall_fast, recall_lessons, and pooled recall entry points. recall_lessons routes to hybrid recall when both text and embedding are supplied and pgmnemo.disable_hybrid is off.

### Record Outcomes

    SELECT pgmnemo.reinforce(1001, 'success', true);
    SELECT pgmnemo.reinforce(
      ARRAY[1001, 1002]::bigint[],
      'failure',
      false
    );

The third p_used argument records whether the recalled memory was actually used. true or NULL increments use_count; false records the outcome without counting a use. Prefer an explicit value so analytics can distinguish ignored advice from used advice.

Under the default posterior mode, match confidence is:

    (success_count + alpha)
    / (success_count + failure_count + alpha + beta)

The default Beta prior is alpha 1 and beta 1. Set pgmnemo.confidence_prior_alpha and pgmnemo.confidence_prior_beta between 0.01 and 100 when a different prior is justified.

### Typed Memory and Navigation

Important write helpers include remember_fact, remember_event, remember_relation, add_edge, reembed, and recompute_content. remember_fact supersedes the active fact for an entity/property pair; events remain append-oriented; relations also populate the graph surface.

Use navigate_locate or navigate_locate_dispatch to select candidate IDs within a character budget, then navigate_expand_typed to fetch content and neighboring graph edges.

### Situation and Entity Recall

The 0.15 line adds deterministic situation fingerprints and a dedicated recall path:

```sql
SELECT pgmnemo.extract_sit_fp(
  'security',
  'failure_class=KEY_ROTATION outcome=COMPLETED'
);
SELECT *
FROM pgmnemo.recall_situation(
  pgmnemo.extract_sit_fp(
    'security',
    'failure_class=KEY_ROTATION outcome=COMPLETED'
  ),
  1,
  'developer',
  10
);
```

Starting with 0.15.1, `recall_situation` returns verified memories by default. Set `pgmnemo.include_unverified = on` only when the caller deliberately accepts memories without provenance verification.

The 0.16 line also extracts stable entity keys into `metadata.entity_keys` during ingestion and exposes entity-centered recall:

```sql
SELECT pgmnemo.extract_entity_keys('The run failed with INFRA_FAILURE.');
SELECT * FROM pgmnemo.recall_entity('failure:INFRA_FAILURE', 10);
```

These extractors are deterministic classifiers, not semantic entity resolution. Normalize application vocabulary and inspect the generated keys before relying on them for tenancy or authorization decisions.

### Configuration Index

- pgmnemo.confidence_mode: posterior by default; additive retains the legacy calculation.
- pgmnemo.confidence_prior_alpha and pgmnemo.confidence_prior_beta: Bayesian prior parameters.
- pgmnemo.confidence_boost_weight: contribution of confidence to ranking; defaults to 0, so confidence does not change rank unless enabled.
- pgmnemo.gate_strict and pgmnemo.include_unverified: provenance enforcement and retrieval.
- pgmnemo.disable_hybrid and pgmnemo.ef_search: recall strategy and HNSW search breadth.
- pgmnemo.track_recall_recency: whether recall updates last_recalled_at and recall_count.
- pgmnemo.max_query_text_chars, pgmnemo.tenant_id, and pgmnemo.test_project_floor: text, tenancy, and optional test-project controls.

The older confidence-delta settings are deprecated and ignored in posterior mode.

### Caveats

- Use PostgreSQL 17 or 18 for pgmnemo 0.16.1. The tagged changelog notes that syntax introduced in the 0.10 line makes older PostgreSQL 14-16 compatibility claims inaccurate; current Pigsty packages target 17-18.

Corpus-maintenance operations are read-only by default:

```sql
SELECT * FROM pgmnemo.reclassify_corpus();
SELECT * FROM pgmnemo.consolidate(
  p_threshold := 0.92,
  p_dry_run := true,
  p_role := NULL,
  p_limit := 100
);
SELECT * FROM pgmnemo.undo_consolidate(
  p_canonical_id := 42,
  p_dry_run := true
);
```

Set `p_dry_run := false` only after reviewing the result inside a transaction. In 0.14.2, reclassification touches only null or classifier-owned types and preserves curator-owned types such as event and relation. Consolidation marks noncanonical lessons superseded, writes edges, and accumulates evidence counts; `undo_consolidate` uses those edges to restore a selected cluster.

- Recall can write recency metadata. Disable pgmnemo.track_recall_recency for read-only analysis.
- The confidence model is only as reliable as reinforcement feedback. Avoid treating posterior values as calibrated probabilities without evaluation.
- HNSW, text, graph, and metadata indexes increase write and maintenance cost.
- The default confidence_boost_weight of 0 means p_min_score can filter results while confidence still contributes nothing to ranking.
- Classification is a deterministic keyword and regular-expression heuristic, not semantic review. Always inspect dry-run distributions and proposed duplicate clusters before applying corpus changes.
