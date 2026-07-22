---
title: "pg_tiktoken_c"
linkTitle: "pg_tiktoken_c"
description: "Fast tiktoken BPE tokenizer for PostgreSQL implemented in C"
weight: 1880
categories: ["RAG"]
width: full
---

[**pg_tiktoken_c**](https://github.com/relytcloud/pg_tiktoken_c) : Fast tiktoken BPE tokenizer for PostgreSQL implemented in C


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1880** | {{< badge content="pg_tiktoken_c" link="https://github.com/relytcloud/pg_tiktoken_c" >}} | {{< ext "pg_tiktoken_c" >}} | `1.1` | {{< category "RAG" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_tiktoken" >}} {{< ext "pg_summarize" >}} {{< ext "vectorize" >}} {{< ext "pgml" >}} {{< ext "pg4ml" >}} {{< ext "pg_graphql" >}} |

> [!Note] Built from upstream main snapshot fa2957b; bundles five vocabularies and includes DESTDIR and correctness patches. Upstream README declares Apache-2.0, but the pinned snapshot omits the referenced LICENSE file.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_tiktoken_c` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "pg_tiktoken_c_18" "green" >}} {{< bg "17" "pg_tiktoken_c_17" "green" >}} {{< bg "16" "pg_tiktoken_c_16" "green" >}} {{< bg "15" "pg_tiktoken_c_15" "green" >}} {{< bg "14" "pg_tiktoken_c_14" "green" >}} | `pg_tiktoken_c_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "postgresql-18-pg-tiktoken-c" "green" >}} {{< bg "17" "postgresql-17-pg-tiktoken-c" "green" >}} {{< bg "16" "postgresql-16-pg-tiktoken-c" "green" >}} {{< bg "15" "postgresql-15-pg-tiktoken-c" "green" >}} {{< bg "14" "postgresql-14-pg-tiktoken-c" "green" >}} | `postgresql-$v-pg-tiktoken-c` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_tiktoken_c_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-tiktoken-c : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-tiktoken-c : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_c_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_c_18-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_c_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_c_18-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_c_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_c_18-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_c_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_c_18-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_c_18` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_c_18-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_c_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_18-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_c_18-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-tiktoken-c` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.1 MiB | [postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-18-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_c_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_c_17-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_c_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_c_17-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_c_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_c_17-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_c_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_c_17-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_c_17` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_c_17-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_c_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_17-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_c_17-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-tiktoken-c` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.1 MiB | [postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-17-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_c_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_c_16-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_c_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_c_16-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_c_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_c_16-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_c_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_c_16-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_c_16` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_c_16-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_c_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_16-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_c_16-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pg-tiktoken-c` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.1 MiB | [postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-16-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_c_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_c_15-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_c_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_c_15-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_c_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_c_15-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_c_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_c_15-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_c_15` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_c_15-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_c_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_15-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_c_15-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pg-tiktoken-c` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.1 MiB | [postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-15-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tiktoken_c_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.3 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tiktoken_c_14-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_tiktoken_c_14` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 2.3 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tiktoken_c_14-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_tiktoken_c_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tiktoken_c_14-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_tiktoken_c_14` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tiktoken_c_14-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_tiktoken_c_14` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.0 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tiktoken_c_14-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_tiktoken_c_14` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 2.0 MiB | [pg_tiktoken_c_14-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tiktoken_c_14-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.8 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.8 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.8 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.8 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pg-tiktoken-c` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.1 MiB | [postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tiktoken-c/postgresql-14-pg-tiktoken-c_1.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/relytcloud/pg_tiktoken_c" title="Repository" icon="github" subtitle="github.com/relytcloud/pg_tiktoken_c" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tiktoken_c-1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_tiktoken_c;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_tiktoken_c;		# install via package name, for the active PG version

pig install pg_tiktoken_c -v 18;   # install for PG 18
pig install pg_tiktoken_c -v 17;   # install for PG 17
pig install pg_tiktoken_c -v 16;   # install for PG 16
pig install pg_tiktoken_c -v 15;   # install for PG 15
pig install pg_tiktoken_c -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_tiktoken_c;
```

## Usage

Sources:

- [pg_tiktoken_c README at the packaged revision](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/README.md)
- [Version 1.1 SQL API](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/sql/pg_tiktoken_c--1.1.sql)
- [Extension control file](https://github.com/relytcloud/pg_tiktoken_c/blob/fa2957b6ece483322f4c4ce0c374b3b77e22b892/pg_tiktoken_c.control)
- [Bundled vocabulary data](https://github.com/relytcloud/pg_tiktoken_c/tree/fa2957b6ece483322f4c4ce0c374b3b77e22b892/data)

pg_tiktoken_c implements OpenAI-compatible tiktoken encoding in C inside PostgreSQL. Use it to count or materialize tokens near stored text and to split text into token-bounded chunks before embedding or model requests.

### Create the Extension

    CREATE EXTENSION pg_tiktoken_c;

The implementation depends on PCRE2 10.30 or later at build time. It does not require shared_preload_libraries; vocabulary data is loaded and cached per backend as encodings are used.

### Encode and Count

    SELECT tiktoken_encode('cl100k_base', 'PostgreSQL search');
    SELECT tiktoken_count('cl100k_base', 'PostgreSQL search');

tiktoken_encode returns a bigint array of token identifiers. tiktoken_count returns the token count without requiring the caller to retain the token array.

The bundled selectors include cl100k_base, o200k_base, r50k_base, p50k_base, and p50k_edit, together with aliases documented by the project. Choose the encoding required by the target model rather than assuming all models share a vocabulary.

### Chunk Text

Return chunks as an array:

    SELECT chunk_text(
      'long document text',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

Or return one row per chunk:

    SELECT *
    FROM chunk_text_table(
      'long document text',
      chunk_size => 512,
      chunk_overlap => 64,
      encoding => 'cl100k_base'
    );

chunk_text_table returns chunk_index, chunk, and token_count. The chunk index is zero-based. Overlap repeats boundary tokens between neighboring chunks and must be smaller than the chunk size.

### Function Index

- tiktoken_encode(selector, text) returns bigint[] token identifiers.
- tiktoken_count(selector, text) returns bigint token count.
- chunk_text(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) returns text[].
- chunk_text_table(input_text, chunk_size, chunk_overlap default 0, encoding default cl100k_base) returns one row per chunk with its index and token count.

The SQL functions are declared immutable and parallel safe. They can therefore be used in generated expressions or parallel plans only when the selected vocabulary files are deployed consistently across every server.

### Operational Notes

- Tokenization is model-encoding specific. Confirm both the encoding name and the model's current context limits in the application.
- Counting or chunking large text consumes backend CPU and memory; batch large corpora and monitor query latency.
- Backend-local caches avoid repeated parsing but increase memory use in sessions that touch several vocabularies.
- The upstream README's compatibility list can lag packaging. Test the exact pg_tiktoken_c build against the target PostgreSQL major version instead of inferring support from a different binary.
