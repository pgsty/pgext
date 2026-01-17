---
title: "vchord"
linkTitle: "vchord"
description: "Vector database plugin for Postgres, written in Rust"
weight: 1810
categories: ["RAG"]
width: full
---

[**vchord**](https://github.com/tensorchord/VectorChord) : Vector database plugin for Postgres, written in Rust


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1810** | {{< badge content="vchord" link="https://github.com/tensorchord/VectorChord" >}} | {{< ext "vchord" >}} | `1.0.0` | {{< category "RAG" >}} | {{< license "AGPL-3.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "vector" >}} |
|   **See Also**    | {{< ext "vectorscale" >}} {{< ext "vectorize" >}} {{< ext "vchord_bm25" >}} {{< ext "pg_tiktoken" >}} {{< ext "pgml" >}} {{< ext "pg_bestmatch" >}} {{< ext "pg_similarity" >}} {{< ext "smlar" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `vchord` | `vector` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "vchord_18" "green" >}} {{< bg "17" "vchord_17" "green" >}} {{< bg "16" "vchord_16" "green" >}} {{< bg "15" "vchord_15" "green" >}} {{< bg "14" "vchord_14" "green" >}} {{< bg "13" "vchord_13" "red" >}} | `vchord_$v` | `pgvector_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-vchord" "green" >}} {{< bg "17" "postgresql-17-vchord" "green" >}} {{< bg "16" "postgresql-16-vchord" "green" >}} {{< bg "15" "postgresql-15-vchord" "green" >}} {{< bg "14" "postgresql-14-vchord" "green" >}} {{< bg "13" "postgresql-13-vchord" "red" >}} | `postgresql-$v-vchord` | `postgresql-$v-pgvector` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "vchord_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "vchord_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "vchord_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-vchord : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-vchord : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-vchord : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [vchord_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.9 MiB | [vchord_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [vchord_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [vchord_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.2 MiB | [vchord_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [vchord_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-vchord` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.2 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-vchord` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-18-vchord_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-18-vchord_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [vchord_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.9 MiB | [vchord_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [vchord_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [vchord_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.2 MiB | [vchord_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [vchord_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-vchord` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.2 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-vchord` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-17-vchord_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-17-vchord_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [vchord_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.8 MiB | [vchord_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [vchord_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [vchord_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [vchord_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [vchord_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-vchord` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-vchord` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-16-vchord_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-16-vchord_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [vchord_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.8 MiB | [vchord_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [vchord_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [vchord_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [vchord_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [vchord_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-vchord` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-vchord` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-15-vchord_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-15-vchord_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `vchord_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 2.2 MiB | [vchord_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/vchord_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `vchord_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.8 MiB | [vchord_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/vchord_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `vchord_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 2.3 MiB | [vchord_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/vchord_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `vchord_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 2.0 MiB | [vchord_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/vchord_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `vchord_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 2.1 MiB | [vchord_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/vchord_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `vchord_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.9 MiB | [vchord_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/vchord_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-vchord` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.0 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.6 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.0 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.6 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 2.1 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.9 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 2.1 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-vchord` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.9 MiB | [postgresql-14-vchord_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/v/vchord/postgresql-14-vchord_1.0.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tensorchord/VectorChord" title="Repository" icon="github" subtitle="github.com/tensorchord/VectorChord" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="VectorChord-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg vchord;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install vchord;		# install via package name, for the active PG version

pig install vchord -v 18;   # install for PG 18
pig install vchord -v 17;   # install for PG 17
pig install vchord -v 16;   # install for PG 16
pig install vchord -v 15;   # install for PG 15
pig install vchord -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'vchord';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION vchord CASCADE; -- requires vector
```


## Usage

- https://github.com/tensorchord/VectorChord
- Launch Blog: [VectorChord: Store 400k Vectors for $1 in PostgreSQL](https://blog.pgvecto.rs/vectorchord-store-400k-vectors-for-1-in-postgresql)

Add this extension to shared_preload_libraries in postgresql.conf

```sql
CREATE EXTENSION vchord CASCADE;
```

Create Index on embedding:

```sql
CREATE INDEX ON gist_train USING vchordrq (embedding vector_l2_ops) WITH (options = $$
residual_quantization = true
[build.internal]
lists = [4096]
spherical_centroids = false
$$);
```

--------

## Docs

### Query

The query statement is exactly the same as pgvector. VectorChord supports any filter operation and WHERE/JOIN clauses like pgvecto.rs with VBASE.

```sql
SELECT * FROM items ORDER BY embedding <-> '[3,1,2]' LIMIT 5;
```

Supported distance functions are:

- `<->` - L2 distance
- `<#>` - (negative) inner product
- `<=>` - cosine distance


> Due to the limitation of postgresql query planner, we cannot support the range query like `SELECT embedding <-> '[3,1,2]' as distance WHERE distance < 0.1 ORDER BY distance` directly.

To query vectors within a certain distance range, you can use the following syntax.

```sql
-- Query vectors within a certain distance range
-- sphere(center, radius) means the vectors within the sphere with the center and radius, aka range query
-- <<->> is L2 distance, <<#>> is inner product, <<=>> is cosine distance
SELECT vec FROM t WHERE vec <<->> sphere('[0.24, 0.24, 0.24]'::vector, 0.012) 
```

### Query Performance Tuning

You can fine-tune the search performance by adjusting the `probes` and `epsilon` parameters:

```sql
-- Set probes to control the number of lists scanned. 
-- Recommended range: 3%–10% of the total `lists` value.
SET vchordrq.probes = 100;

-- Set epsilon to control the reranking precision.
-- Larger value means more rerank for higher recall rate.
-- Don't change it unless you only have limited memory.
-- Recommended range: 1.0–1.9. Default value is 1.9.
SET vchordrq.epsilon = 1.9;

-- vchordrq relies on a projection matrix to optimize performance.
-- Add your vector dimensions to the `prewarm_dim` list to reduce latency.
-- If this is not configured, the first query will have higher latency as the matrix is generated on demand.
-- Default value: '64,128,256,384,512,768,1024,1536'
-- Note: This setting requires a database restart to take effect.
ALTER SYSTEM SET vchordrq.prewarm_dim = '64,128,256,384,512,768,1024,1536';
```

And for postgres's setting
```sql
-- If using SSDs, set `effective_io_concurrency` to 200 for faster disk I/O.
SET effective_io_concurrency = 200;

-- Disable JIT (Just-In-Time Compilation) as it offers minimal benefit (1–2%) 
-- and adds overhead for single-query workloads.
SET jit = off;

-- Allocate at least 25% of total memory to `shared_buffers`. 
-- For disk-heavy workloads, you can increase this to up to 90% of total memory. You may also want to disable swap with network storage to avoid io hang.
-- Note: A restart is required for this setting to take effect.
ALTER SYSTEM SET shared_buffers = '8GB';
```

### Indexing prewarm

To prewarm the index, you can use the following SQL. It will significantly improve performance when using limited memory.

```sql
-- vchordrq_prewarm(index_name::regclass) to prewarm the index into the shared buffer
SELECT vchordrq_prewarm('gist_train_embedding_idx'::regclass)"
```


### Index Build Time

Index building can parallelized, and with external centroid precomputation, the total time is primarily limited by disk speed. Optimize parallelism using the following settings:

```sql
-- Set this to the number of CPU cores available for parallel operations.
SET max_parallel_maintenance_workers = 8;
SET max_parallel_workers = 8;

-- Adjust the total number of worker processes. 
-- Note: A restart is required for this setting to take effect.
ALTER SYSTEM SET max_worker_processes = 8;
```

### Indexing Progress


You can check the indexing progress by querying the `pg_stat_progress_create_index` view.

```sql
SELECT phase, round(100.0 * blocks_done / nullif(blocks_total, 0), 1) AS "%" FROM pg_stat_progress_create_index;
```

### External Index Precomputation

Unlike pure SQL, an external index precomputation will first do clustering outside and insert centroids to a PostgreSQL table. Although it might be more complicated, external build is definitely much faster on larger dataset (>5M).

To get started, you need to do a clustering of vectors using `faiss`, `scikit-learn` or any other clustering library.

The centroids should be preset in a table of any name with 3 columns:
- id(integer): id of each centroid, should be unique
- parent(integer, nullable): parent id of each centroid, should be NULL for normal clustering
- vector(vector): representation of each centroid, `pgvector` vector type

And example could be like this:

```sql
-- Create table of centroids
CREATE TABLE public.centroids (id integer NOT NULL UNIQUE, parent integer, vector vector(768));
-- Insert centroids into it
INSERT INTO public.centroids (id, parent, vector) VALUES (1, NULL, '{0.1, 0.2, 0.3, ..., 0.768}');
INSERT INTO public.centroids (id, parent, vector) VALUES (2, NULL, '{0.4, 0.5, 0.6, ..., 0.768}');
INSERT INTO public.centroids (id, parent, vector) VALUES (3, NULL, '{0.7, 0.8, 0.9, ..., 0.768}');
-- ...

-- Create index using the centroid table
CREATE INDEX ON gist_train USING vchordrq (embedding vector_l2_ops) WITH (options = $$
[build.external]
table = 'public.centroids'
$$);
```

To simplify the workflow, we provide end-to-end scripts for external index pre-computation, see [scripts](./scripts/README.md#run-external-index-precomputation-toolkit).



------

## Limitations

- Data Type Support: Currently, only the `f32` data type is supported for vectors.
- Architecture Compatibility: The fast-scan kernel is optimized for x86_64 architectures. While it runs on aarch64, performance may be lower.
- KMeans Clustering: The built-in KMeans clustering is not yet fully optimized and may require substantial memory. We strongly recommend using external centroid precomputation for efficient index construction.


------

## Build

Building this extension requires [clang-17+](https://github.com/tensorchord/VectorChord/issues/188)

Which is available on EL 8/9, Ubuntu 24.04 directly, but require manual installation on Ubuntu 22.04 / Debian 12.

For example, install clang-18 on Ubuntu 22 / Debian 12 and set it as the default clang:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://apt.llvm.org/llvm.sh | bash -s -- 18
sudo update-alternatives --install /usr/bin/clang clang $(which clang-18) 255
```

