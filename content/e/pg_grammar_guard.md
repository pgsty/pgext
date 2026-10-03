---
title: "pg_grammar_guard"
linkTitle: "pg_grammar_guard"
description: "Catalog-derived grammars and approved grammar drift checks"
weight: 1890
categories: ["RAG"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_grammar_guard**](https://github.com/Manuelreyesbravo/pg_grammar_guard) : Catalog-derived grammars and approved grammar drift checks


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1890** | {{< badge content="pg_grammar_guard" link="https://github.com/Manuelreyesbravo/pg_grammar_guard" >}} | {{< ext "pg_grammar_guard" >}} | `0.4.1` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `grammar_guard` |
|   **Requires**    | {{< ext "pg_living_assertions" >}} |

> [!Note] Requires pg_living_assertions, despite older README text claiming no dependencies.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_grammar_guard` | `pg_living_assertions` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "pg_grammar_guard_18" "green" >}} {{< bg "17" "pg_grammar_guard_17" "green" >}} {{< bg "16" "pg_grammar_guard_16" "green" >}} {{< bg "15" "pg_grammar_guard_15" "green" >}} {{< bg "14" "pg_grammar_guard_14" "green" >}} | `pg_grammar_guard_$v` | `pg_living_assertions_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "postgresql-18-pg-grammar-guard" "green" >}} {{< bg "17" "postgresql-17-pg-grammar-guard" "green" >}} {{< bg "16" "postgresql-16-pg-grammar-guard" "green" >}} {{< bg "15" "postgresql-15-pg-grammar-guard" "green" >}} {{< bg "14" "postgresql-14-pg-grammar-guard" "green" >}} | `postgresql-$v-pg-grammar-guard` | `postgresql-$v-pg-living-assertions` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pg_grammar_guard_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pg-grammar-guard : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pg-grammar-guard : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_grammar_guard_18` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_grammar_guard_18-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_18` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_grammar_guard_18-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_18` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.3 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_grammar_guard_18-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_18` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_grammar_guard_18-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_18` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.4 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_grammar_guard_18-0.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_grammar_guard_18` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [pg_grammar_guard_18-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_grammar_guard_18-0.4.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.4 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.4 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.4 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.0 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
| `postgresql-18-pg-grammar-guard` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-18-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_grammar_guard_17` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_grammar_guard_17-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_17` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_grammar_guard_17-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_17` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.3 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_grammar_guard_17-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_17` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_grammar_guard_17-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_17` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.4 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_grammar_guard_17-0.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_grammar_guard_17` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [pg_grammar_guard_17-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_grammar_guard_17-0.4.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.4 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.4 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.4 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.0 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
| `postgresql-17-pg-grammar-guard` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-17-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_grammar_guard_16` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_grammar_guard_16-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_16` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_grammar_guard_16-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_16` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.3 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_grammar_guard_16-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_16` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_grammar_guard_16-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_16` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.4 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_grammar_guard_16-0.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_grammar_guard_16` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [pg_grammar_guard_16-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_grammar_guard_16-0.4.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.4 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.4 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.4 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.0 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
| `postgresql-16-pg-grammar-guard` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-16-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_grammar_guard_15` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_grammar_guard_15-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_15` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_grammar_guard_15-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_15` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.2 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_grammar_guard_15-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_15` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_grammar_guard_15-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_15` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.4 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_grammar_guard_15-0.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_grammar_guard_15` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [pg_grammar_guard_15-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_grammar_guard_15-0.4.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.4 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.4 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.4 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.0 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
| `postgresql-15-pg-grammar-guard` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-15-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_grammar_guard_14` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 27.6 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_grammar_guard_14-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_14` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 27.6 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_grammar_guard_14-0.4.1-1PGSTY.el8.noarch.rpm) |
| `pg_grammar_guard_14` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 27.2 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_grammar_guard_14-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_14` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 27.2 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_grammar_guard_14-0.4.1-1PGSTY.el9.noarch.rpm) |
| `pg_grammar_guard_14` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 27.4 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_grammar_guard_14-0.4.1-1PGSTY.el10.noarch.rpm) |
| `pg_grammar_guard_14` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 27.3 KiB | [pg_grammar_guard_14-0.4.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_grammar_guard_14-0.4.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 21.4 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 21.4 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 21.4 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 21.4 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 22.1 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 22.1 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 22.0 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 22.0 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~noble_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 22.0 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
| `postgresql-14-pg-grammar-guard` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 22.0 KiB | [postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-grammar-guard/postgresql-14-pg-grammar-guard_0.4.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Manuelreyesbravo/pg_grammar_guard" title="Repository" icon="github" subtitle="github.com/Manuelreyesbravo/pg_grammar_guard" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_grammar_guard-0.4.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_grammar_guard;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_grammar_guard;		# install via package name, for the active PG version

pig install pg_grammar_guard -v 18;   # install for PG 18
pig install pg_grammar_guard -v 17;   # install for PG 17
pig install pg_grammar_guard -v 16;   # install for PG 16
pig install pg_grammar_guard -v 15;   # install for PG 15
pig install pg_grammar_guard -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_grammar_guard CASCADE; -- requires pg_living_assertions
```

## Usage

Sources:

- [PGXN 0.4.1](https://pgxn.org/dist/pg_grammar_guard/0.4.1/)

`pg_grammar_guard` generates GBNF or JSON Schema from catalog identifiers and detects drift from approved grammars. It is pure SQL, needs no preload and is packaged for PostgreSQL 14–18. Its control file requires `pg_living_assertions`.

### Generate a grammar

```sql
CREATE EXTENSION pg_grammar_guard CASCADE;
CREATE TABLE public.grammar_demo (id integer, label text);
SELECT grammar_guard.grammar_for_json(ARRAY[
  ROW('column', 'enum',
      grammar_guard.catalog_columns('public.grammar_demo'), true)
]::grammar_guard.grammar_field[]);
```

The catalog helpers enumerate actual tables, columns and enum labels. Generation supports nested objects and bounded arrays; open-ended values such as arbitrary SQL or file paths need separate validation.

### Detect changes

`grammar_guard.watch()` stores the query that rebuilds a grammar, and `grammar_guard.check_grammar()` evaluates it against the live catalog. Results and their age are available through `living_assertions.status`.

Register baseline SQL only through trusted administrators. A grammar limits valid identifiers; it cannot prove that a selected table, join or answer is semantically correct. Baselines from the old 0.2 series lack the original generating query and require manual re-approval when upgrading.
