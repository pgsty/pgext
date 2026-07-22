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
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.20" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_18` | `0.6.20` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.7 MiB | [pgrdf_18-0.6.20-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_18-0.6.20-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_18` | `0.6.20` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_18-0.6.20-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_18-0.6.20-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_18` | `0.6.20` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.4 MiB | [pgrdf_18-0.6.20-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_18-0.6.20-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_18` | `0.6.20` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.1 MiB | [pgrdf_18-0.6.20-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_18-0.6.20-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_18` | `0.6.20` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.3 MiB | [pgrdf_18-0.6.20-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_18-0.6.20-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_18` | `0.6.20` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pgrdf_18-0.6.20-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_18-0.6.20-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgrdf` | `0.6.20` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.3 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.1 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.6 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.20` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.1 MiB | [postgresql-18-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_17` | `0.6.20` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.7 MiB | [pgrdf_17-0.6.20-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_17-0.6.20-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_17` | `0.6.20` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_17-0.6.20-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_17-0.6.20-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_17` | `0.6.20` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.4 MiB | [pgrdf_17-0.6.20-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_17-0.6.20-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_17` | `0.6.20` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.1 MiB | [pgrdf_17-0.6.20-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_17-0.6.20-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_17` | `0.6.20` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.3 MiB | [pgrdf_17-0.6.20-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_17-0.6.20-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_17` | `0.6.20` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pgrdf_17-0.6.20-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_17-0.6.20-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgrdf` | `0.6.20` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.3 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.1 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.6 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.20` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.1 MiB | [postgresql-17-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_16` | `0.6.20` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.7 MiB | [pgrdf_16-0.6.20-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_16-0.6.20-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_16` | `0.6.20` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_16-0.6.20-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_16-0.6.20-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_16` | `0.6.20` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.4 MiB | [pgrdf_16-0.6.20-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_16-0.6.20-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_16` | `0.6.20` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.1 MiB | [pgrdf_16-0.6.20-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_16-0.6.20-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_16` | `0.6.20` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.3 MiB | [pgrdf_16-0.6.20-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_16-0.6.20-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_16` | `0.6.20` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pgrdf_16-0.6.20-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_16-0.6.20-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgrdf` | `0.6.20` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.3 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.1 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.6 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.20` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.1 MiB | [postgresql-16-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_15` | `0.6.20` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.7 MiB | [pgrdf_15-0.6.20-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_15-0.6.20-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_15` | `0.6.20` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_15-0.6.20-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_15-0.6.20-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_15` | `0.6.20` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.3 MiB | [pgrdf_15-0.6.20-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_15-0.6.20-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_15` | `0.6.20` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.1 MiB | [pgrdf_15-0.6.20-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_15-0.6.20-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_15` | `0.6.20` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.3 MiB | [pgrdf_15-0.6.20-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_15-0.6.20-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_15` | `0.6.20` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pgrdf_15-0.6.20-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_15-0.6.20-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgrdf` | `0.6.20` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.3 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.1 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.6 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.20` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.1 MiB | [postgresql-15-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_14` | `0.6.20` | [el8.x86_64](/os/el8.x86_64) | pigsty | 7.7 MiB | [pgrdf_14-0.6.20-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_14-0.6.20-1PIGSTY.el8.x86_64.rpm) |
| `pgrdf_14` | `0.6.20` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.8 MiB | [pgrdf_14-0.6.20-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_14-0.6.20-1PIGSTY.el8.aarch64.rpm) |
| `pgrdf_14` | `0.6.20` | [el9.x86_64](/os/el9.x86_64) | pigsty | 7.3 MiB | [pgrdf_14-0.6.20-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_14-0.6.20-1PIGSTY.el9.x86_64.rpm) |
| `pgrdf_14` | `0.6.20` | [el9.aarch64](/os/el9.aarch64) | pigsty | 7.0 MiB | [pgrdf_14-0.6.20-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_14-0.6.20-1PIGSTY.el9.aarch64.rpm) |
| `pgrdf_14` | `0.6.20` | [el10.x86_64](/os/el10.x86_64) | pigsty | 7.3 MiB | [pgrdf_14-0.6.20-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_14-0.6.20-1PIGSTY.el10.x86_64.rpm) |
| `pgrdf_14` | `0.6.20` | [el10.aarch64](/os/el10.aarch64) | pigsty | 7.0 MiB | [pgrdf_14-0.6.20-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_14-0.6.20-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgrdf` | `0.6.20` | [d12.x86_64](/os/d12.x86_64) | pigsty | 6.2 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.3 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [d13.x86_64](/os/d13.x86_64) | pigsty | 6.2 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.3 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.6 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u22.aarch64](/os/u22.aarch64) | pigsty | 6.1 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.6 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u24.aarch64](/os/u24.aarch64) | pigsty | 6.1 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.6 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.20` | [u26.aarch64](/os/u26.aarch64) | pigsty | 6.1 MiB | [postgresql-14-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.20-1PIGSTY~resolute_arm64.deb) |

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

Sources:

- [pgRDF 0.6.20 README](https://github.com/styk-tv/pgRDF/blob/v0.6.20/README.md)
- [pgRDF 0.6.20 user guide](https://github.com/styk-tv/pgRDF/tree/v0.6.20/guide)
- [pgRDF 0.6.20 changelog](https://github.com/styk-tv/pgRDF/blob/v0.6.20/CHANGELOG.md)
- [pgRDF 0.6.20 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.20)

`pgRDF` stores RDF data inside PostgreSQL and exposes SQL-callable helpers for Turtle/TriG/N-Quads loading, SPARQL query/update, named graphs, SHACL validation, and RDFS/OWL 2 RL materialization.

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.version();
```

### Preload And PostgreSQL Version Caveat

pgRDF 0.6.20 supports PostgreSQL 14-18 and moves from `pgrx` 0.16.1 to 0.19.1. Upstream describes 0.6.20 as a build/runtime migration with no schema or query-surface change; PostgreSQL 19 remains a tracked follow-up.

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

### Carve Graph Slices

The 0.6.x series adds `carve_graph` overloads for copying a predicate-defined slice or a bounded neighbourhood into another graph without decoding and re-encoding the shared dictionary:

```sql
SELECT pgrdf.carve_graph(
  100::bigint,
  'http://example.org/type'::text,
  200::bigint
);

SELECT pgrdf.carve_graph(
  100::bigint,
  ARRAY['http://example.org/alice', 'http://example.org/bob']::text[],
  201::bigint,
  2
);
```

The neighbourhood form uses `max_hops` as a graph-distance boundary. Set `pgrdf.on_path_truncation` to `warn` or `error` when a truncated property-path walk must not be silently accepted.

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

The `pgrdf.path_max_depth` setting guards property-path expansion depth, while `pgrdf.on_path_truncation = count | warn | error` controls how callers learn that the guard was reached.

### Version Notes

The releases between 0.6.4 and 0.6.20 materially improve large RDF ingestion and query correctness: streaming/windowed bulk ingestion, a staged multi-backend loader, safe repeated loads into populated dictionaries, graph-carving helpers, dictionary inclusion in `pg_dump`, SPARQL expression/aggregate additions, and fail-closed path-truncation handling. The 0.6.20 release itself changes only the build/runtime layer to `pgrx` 0.19.1 and adds PostgreSQL 18 support.

For very large fresh N-Triples loads, upstream documents `pgrdf.load_turtle_staged_run` as the resumable, phase-oriented path. It commits parsing, dictionary, resolution, and index phases separately and is operationally different from the transactional `load_turtle()` call; validate staging tablespaces, disk headroom, and recovery procedures before using it for production-scale imports.
