---
title: "pgrdf"
linkTitle: "pgrdf"
description: "RDF, SPARQL, SHACL, and OWL reasoning for PostgreSQL"
weight: 2640
categories: ["FEAT"]
width: full
---

[**pgrdf**](https://github.com/styk-tv/pgRDF) : RDF, SPARQL, SHACL, and OWL reasoning for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2640** | {{< badge content="pgrdf" link="https://github.com/styk-tv/pgRDF" >}} | {{< ext "pgrdf" >}} | `0.6.20` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgrdf` |
|   **See Also**    | {{< ext "rdf_fdw" >}} {{< ext "pg_sparql" >}} {{< ext "rdkit" >}} |

> [!Note] Production hook/cache deployments should preload pgrdf.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.20` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgrdf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.20` | {{< bg "18" "pgrdf_18" "green" >}} {{< bg "17" "pgrdf_17" "green" >}} {{< bg "16" "pgrdf_16" "green" >}} {{< bg "15" "pgrdf_15" "green" >}} {{< bg "14" "pgrdf_14" "green" >}} | `pgrdf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.20` | {{< bg "18" "postgresql-18-pgrdf" "green" >}} {{< bg "17" "postgresql-17-pgrdf" "green" >}} {{< bg "16" "postgresql-16-pgrdf" "green" >}} {{< bg "15" "postgresql-15-pgrdf" "green" >}} {{< bg "14" "postgresql-14-pgrdf" "green" >}} | `postgresql-$v-pgrdf` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pgrdf_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgrdf : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.6.4" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.4" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_17` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pgrdf_17-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_17-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_17` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_17-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_17-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_17` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgrdf_17-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_17-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_17` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.0 MiB | [pgrdf_17-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_17-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_17` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pgrdf_17-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_17-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_17` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [pgrdf_17-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_17-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgrdf` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.4 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.0 MiB | [postgresql-17-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_16` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pgrdf_16-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_16-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_16` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_16-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_16-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_16` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgrdf_16-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_16-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_16` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.0 MiB | [pgrdf_16-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_16-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_16` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pgrdf_16-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_16-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_16` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [pgrdf_16-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_16-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgrdf` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.4 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.0 MiB | [postgresql-16-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_15` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pgrdf_15-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_15-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_15` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.7 MiB | [pgrdf_15-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_15-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_15` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgrdf_15-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_15-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_15` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.0 MiB | [pgrdf_15-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_15-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_15` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pgrdf_15-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_15-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_15` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [pgrdf_15-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_15-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgrdf` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.4 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.0 MiB | [postgresql-15-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_14` | `0.6.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.4 MiB | [pgrdf_14-0.6.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_14-0.6.4-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_14` | `0.6.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.7 MiB | [pgrdf_14-0.6.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_14-0.6.4-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_14` | `0.6.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.1 MiB | [pgrdf_14-0.6.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_14-0.6.4-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_14` | `0.6.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.9 MiB | [pgrdf_14-0.6.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_14-0.6.4-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_14` | `0.6.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.1 MiB | [pgrdf_14-0.6.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_14-0.6.4-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_14` | `0.6.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.9 MiB | [pgrdf_14-0.6.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_14-0.6.4-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgrdf` | `0.6.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.0 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.4 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.0 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.4 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.0 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.4 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.0 MiB | [postgresql-14-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/styk-tv/pgRDF" title="Repository" icon="github" subtitle="github.com/styk-tv/pgRDF" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgrdf-0.6.20.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgrdf;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgrdf;		# install via package name, for the active PG version

pig install pgrdf -v 18;   # install for PG 18
pig install pgrdf -v 17;   # install for PG 17
pig install pgrdf -v 16;   # install for PG 16
pig install pgrdf -v 15;   # install for PG 15
pig install pgrdf -v 14;   # install for PG 14

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'pgrdf';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgrdf;
```




## Usage

> Sources: [pgRDF upstream README](https://github.com/styk-tv/pgRDF/blob/v0.6.4/README.md), [pgRDF user guide](https://github.com/styk-tv/pgRDF/tree/v0.6.4/guide), [v0.6.4 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.4).

`pgRDF` stores RDF data inside PostgreSQL and exposes SQL-callable helpers for Turtle/TriG/N-Quads loading, SPARQL query/update, named graphs, SHACL validation, and RDFS/OWL 2 RL materialization.

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### Preload And PostgreSQL Version Caveat

Upstream documents PostgreSQL 14-17 support and defers PostgreSQL 18 while pgRDF remains pinned to `pgrx` 0.16.

`pgrdf` must be present in `shared_preload_libraries` before PostgreSQL starts. Without preload, upstream documents that the shared-memory dictionary and plan-cache atomics are not initialized and the first pgRDF call can fail.

```conf
shared_preload_libraries = 'pgrdf'
```

Restart PostgreSQL after changing this setting, then verify:

```sql
SHOW shared_preload_libraries;

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:t a ex:T .',
  1::bigint,
  'http://example.org/'
);
```

### Load RDF

Use `parse_turtle` for inline Turtle payloads and `load_turtle` for server-side files. Graph ids are `bigint` values; named graph helpers map ids to IRIs. Version 0.6.x adds a parallel bulk loader path via `load_turtle(..., bulk_load => true)`.

```sql
SELECT pgrdf.add_graph(100::bigint, 'http://example.org/graph/main');

SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> .
   ex:alice ex:knows ex:bob .
   ex:alice ex:name "Alice" .',
  100::bigint,
  'http://example.org/graph/main'
);

SELECT pgrdf.load_turtle('/srv/rdf/foaf.ttl', 100::bigint);
SELECT pgrdf.load_turtle('/srv/rdf/bulk.ttl', 100::bigint, bulk_load => true);
SELECT pgrdf.count_quads(100::bigint);
```

Related ingest and graph-management functions documented upstream include `parse_trig`, `parse_nquads`, `add_graph`, `drop`, `clear`, `copy`, `move_graph`, `graph_id`, and `graph_iri`.

### Query With SPARQL

`pgrdf.sparql(text)` returns SPARQL results as SQL rows. The upstream v0.5 surface includes `SELECT` and `ASK`, filters, ordering, limits, `OPTIONAL`, `UNION`, `MINUS`, aggregates, `VALUES`, `BIND`, `CONSTRUCT`, `DESCRIBE`, named-graph `GRAPH` clauses, and property paths.

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?person ?name
   WHERE {
     ?person ex:name ?name .
     FILTER(REGEX(?name, "^A", "i"))
   }
   ORDER BY ?name
   LIMIT 20'
);
```

Named-graph queries can bind graph IRIs:

```sql
SELECT *
FROM pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   SELECT ?g (COUNT(*) AS ?n)
   WHERE { GRAPH ?g { ?s ex:name ?name } }
   GROUP BY ?g
   ORDER BY ?g'
);
```

### Update Graphs

The upstream v0.5 surface includes SPARQL Update forms such as `INSERT DATA`, `DELETE DATA`, `INSERT/DELETE WHERE`, `DELETE+INSERT WHERE`, and graph lifecycle statements.

```sql
SELECT pgrdf.sparql(
  'PREFIX ex: <http://example.org/>
   INSERT DATA {
     GRAPH <http://example.org/graph/main> {
       ex:bob ex:name "Bob" .
     }
   }'
);
```

### Reasoning And Validation

Use `pgrdf.materialize(graph_id, profile)` to write inferred triples for `rdfs` or `owl-rl` profiles. Materialization is intended to be repeatable; upstream documents that previous inferred rows are dropped before writing the new closure.

```sql
SELECT pgrdf.parse_turtle(
  '@prefix ex:   <http://example.com/> .
   @prefix rdf:  <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
   @prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
   ex:Engineer rdfs:subClassOf ex:Person .
   ex:Person   rdfs:subClassOf ex:Agent .
   ex:alice    rdf:type        ex:Engineer .',
  100::bigint
);

SELECT pgrdf.materialize(100::bigint, 'owl-rl');

SELECT *
FROM pgrdf.sparql(
  'PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
   PREFIX ex:  <http://example.com/>
   SELECT ?class WHERE { ex:alice rdf:type ?class }'
);
```

Use `pgrdf.validate(data, shapes, mode)` for SHACL validation; upstream documents JSONB `sh:ValidationReport` output and native SHACL Core support. SHACL-SPARQL constraint execution is documented upstream as gated by its RDF library dependency, so treat `mode => 'sparql'` as an advanced surface to verify against the exact installed version.

### Operational Helpers

Useful introspection and cache-management helpers documented upstream include:

| Function | Use |
|----------|-----|
| `pgrdf.stats()` | Inspect runtime counters and cache state |
| `pgrdf.shmem_reset()` | Reset shared-memory dictionary/cache state |
| `pgrdf.plan_cache_clear()` | Clear prepared SPARQL plan cache |
| `pgrdf.sparql_parse(text)` | Inspect parsed SPARQL without executing it |

The `pgrdf.path_max_depth` setting guards property-path expansion depth.

### Version Notes

`pgrdf` 0.6.4 improves the deferred-index bulk-load path: for fresh bulk loads above `pgrdf.bulk_defer_index_min`, `load_turtle(..., bulk_load => true)` also defers the dictionary `unique_term` constraint, then rebuilds and validates it in the same transaction. PostgreSQL 18 remains deferred upstream while pgRDF is pinned to `pgrx` 0.16.
