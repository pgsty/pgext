---
title: "pgrdf"
linkTitle: "pgrdf"
description: "RDF, SPARQL, SHACL, and OWL reasoning for PostgreSQL"
weight: 2640
categories: ["FEAT"]
languages: ["Rust"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pgrdf**](https://github.com/styk-tv/pgRDF) : RDF, SPARQL, SHACL, and OWL reasoning for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2640** | {{< badge content="pgrdf" link="https://github.com/styk-tv/pgRDF" >}} | {{< ext "pgrdf" >}} | `0.6.39` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgrdf` |
|   **See Also**    | {{< ext "sparql" >}} {{< ext "rdf_fdw" >}} {{< ext "age" >}} {{< ext "pg_liquid" >}} {{< ext "onesparse" >}} {{< ext "graph" >}} {{< ext "ltree" >}} {{< ext "xml2" >}} {{< ext "plxslt" >}} {{< ext "omni_xml" >}} |

> [!Note] Production hook/cache deployments should preload pgrdf.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.39` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgrdf` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.39` | {{< bg "18" "pgrdf_18" "green" >}} {{< bg "17" "pgrdf_17" "green" >}} {{< bg "16" "pgrdf_16" "green" >}} {{< bg "15" "pgrdf_15" "green" >}} {{< bg "14" "pgrdf_14" "green" >}} | `pgrdf_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.6.39` | {{< bg "18" "postgresql-18-pgrdf" "green" >}} {{< bg "17" "postgresql-17-pgrdf" "green" >}} {{< bg "16" "postgresql-16-pgrdf" "green" >}} {{< bg "15" "postgresql-15-pgrdf" "green" >}} {{< bg "14" "postgresql-14-pgrdf" "green" >}} | `postgresql-$v-pgrdf` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "pgrdf_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-18-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-17-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-16-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-15-pgrdf : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.6.39" "postgresql-14-pgrdf : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_18` | `0.6.39` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.8 MiB | [pgrdf_18-0.6.39-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_18-0.6.39-1PGSTY.el8.x86_64.rpm) |
| `pgrdf_18` | `0.6.39` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgrdf_18-0.6.39-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_18-0.6.39-1PGSTY.el8.aarch64.rpm) |
| `pgrdf_18` | `0.6.39` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.6 MiB | [pgrdf_18-0.6.39-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_18-0.6.39-1PGSTY.el9.x86_64.rpm) |
| `pgrdf_18` | `0.6.39` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.2 MiB | [pgrdf_18-0.6.39-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_18-0.6.39-1PGSTY.el9.aarch64.rpm) |
| `pgrdf_18` | `0.6.39` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.6 MiB | [pgrdf_18-0.6.39-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_18-0.6.39-1PGSTY.el10.x86_64.rpm) |
| `pgrdf_18` | `0.6.39` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgrdf_18-0.6.39-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_18-0.6.39-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgrdf` | `0.6.39` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.8 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.8 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.8 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pgrdf` | `0.6.39` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-18-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-18-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_17` | `0.6.39` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.8 MiB | [pgrdf_17-0.6.39-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_17-0.6.39-1PGSTY.el8.x86_64.rpm) |
| `pgrdf_17` | `0.6.39` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgrdf_17-0.6.39-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_17-0.6.39-1PGSTY.el8.aarch64.rpm) |
| `pgrdf_17` | `0.6.39` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.6 MiB | [pgrdf_17-0.6.39-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_17-0.6.39-1PGSTY.el9.x86_64.rpm) |
| `pgrdf_17` | `0.6.39` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.2 MiB | [pgrdf_17-0.6.39-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_17-0.6.39-1PGSTY.el9.aarch64.rpm) |
| `pgrdf_17` | `0.6.39` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.6 MiB | [pgrdf_17-0.6.39-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_17-0.6.39-1PGSTY.el10.x86_64.rpm) |
| `pgrdf_17` | `0.6.39` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgrdf_17-0.6.39-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_17-0.6.39-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgrdf` | `0.6.39` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.8 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.8 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.8 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pgrdf` | `0.6.39` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-17-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-17-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_16` | `0.6.39` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.8 MiB | [pgrdf_16-0.6.39-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_16-0.6.39-1PGSTY.el8.x86_64.rpm) |
| `pgrdf_16` | `0.6.39` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgrdf_16-0.6.39-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_16-0.6.39-1PGSTY.el8.aarch64.rpm) |
| `pgrdf_16` | `0.6.39` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.6 MiB | [pgrdf_16-0.6.39-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_16-0.6.39-1PGSTY.el9.x86_64.rpm) |
| `pgrdf_16` | `0.6.39` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.2 MiB | [pgrdf_16-0.6.39-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_16-0.6.39-1PGSTY.el9.aarch64.rpm) |
| `pgrdf_16` | `0.6.39` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.6 MiB | [pgrdf_16-0.6.39-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_16-0.6.39-1PGSTY.el10.x86_64.rpm) |
| `pgrdf_16` | `0.6.39` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgrdf_16-0.6.39-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_16-0.6.39-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgrdf` | `0.6.39` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.8 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.8 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~noble_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.8 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~noble_arm64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-pgrdf` | `0.6.39` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-16-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-16-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_15` | `0.6.39` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.8 MiB | [pgrdf_15-0.6.39-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_15-0.6.39-1PGSTY.el8.x86_64.rpm) |
| `pgrdf_15` | `0.6.39` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgrdf_15-0.6.39-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_15-0.6.39-1PGSTY.el8.aarch64.rpm) |
| `pgrdf_15` | `0.6.39` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.6 MiB | [pgrdf_15-0.6.39-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_15-0.6.39-1PGSTY.el9.x86_64.rpm) |
| `pgrdf_15` | `0.6.39` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.2 MiB | [pgrdf_15-0.6.39-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_15-0.6.39-1PGSTY.el9.aarch64.rpm) |
| `pgrdf_15` | `0.6.39` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.5 MiB | [pgrdf_15-0.6.39-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_15-0.6.39-1PGSTY.el10.x86_64.rpm) |
| `pgrdf_15` | `0.6.39` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgrdf_15-0.6.39-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_15-0.6.39-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgrdf` | `0.6.39` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.8 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.8 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~noble_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~noble_arm64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-pgrdf` | `0.6.39` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-15-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-15-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgrdf_14` | `0.6.39` | [el8.x86_64](/os/el8.x86_64) | pigsty | 6.8 MiB | [pgrdf_14-0.6.39-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgrdf_14-0.6.39-1PGSTY.el8.x86_64.rpm) |
| `pgrdf_14` | `0.6.39` | [el8.aarch64](/os/el8.aarch64) | pigsty | 6.0 MiB | [pgrdf_14-0.6.39-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgrdf_14-0.6.39-1PGSTY.el8.aarch64.rpm) |
| `pgrdf_14` | `0.6.39` | [el9.x86_64](/os/el9.x86_64) | pigsty | 6.6 MiB | [pgrdf_14-0.6.39-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgrdf_14-0.6.39-1PGSTY.el9.x86_64.rpm) |
| `pgrdf_14` | `0.6.39` | [el9.aarch64](/os/el9.aarch64) | pigsty | 6.2 MiB | [pgrdf_14-0.6.39-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgrdf_14-0.6.39-1PGSTY.el9.aarch64.rpm) |
| `pgrdf_14` | `0.6.39` | [el10.x86_64](/os/el10.x86_64) | pigsty | 6.5 MiB | [pgrdf_14-0.6.39-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgrdf_14-0.6.39-1PGSTY.el10.x86_64.rpm) |
| `pgrdf_14` | `0.6.39` | [el10.aarch64](/os/el10.aarch64) | pigsty | 6.2 MiB | [pgrdf_14-0.6.39-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgrdf_14-0.6.39-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgrdf` | `0.6.39` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.8 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.1 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.8 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.1 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u22.x86_64](/os/u22.x86_64) | pigsty | 6.2 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.8 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u24.x86_64](/os/u24.x86_64) | pigsty | 6.1 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~noble_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.7 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~noble_arm64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u26.x86_64](/os/u26.x86_64) | pigsty | 6.1 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-pgrdf` | `0.6.39` | [u26.aarch64](/os/u26.aarch64) | pigsty | 5.7 MiB | [postgresql-14-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgrdf/postgresql-14-pgrdf_0.6.39-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/styk-tv/pgRDF" title="Repository" icon="github" subtitle="github.com/styk-tv/pgRDF" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgrdf-0.6.39.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgrdf;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

- [0.6.39 release](https://github.com/styk-tv/pgRDF/releases/tag/v0.6.39)
- [pgrdf.control](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/pgrdf.control)
- [README.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/README.md)
- [Cargo.toml](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/Cargo.toml)
- [guide/01-install.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/01-install.md)
- [guide/tour.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/tour.md)
- [guide/05-graphs.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/05-graphs.md)
- [guide/06-validation-recipes.md](https://github.com/styk-tv/pgRDF/blob/a1dce7bd7c4d1138f713c5b73e9d498aa279612b/guide/06-validation-recipes.md)

`pgrdf` 0.6.39 stores RDF graphs in PostgreSQL with Turtle/TriG/N-Quads ingestion, SPARQL queries and updates, SHACL validation, and RDFS/OWL reasoning. Preload `pgrdf` and restart before creating the extension as a superuser.

### Core Workflow

```ini
shared_preload_libraries = 'pgrdf'
```

```sql
CREATE EXTENSION pgrdf;
SELECT pgrdf.add_graph('http://example.org/people');
SELECT pgrdf.parse_turtle(
  '@prefix ex: <http://example.org/> . ex:alice ex:name "Alice" .',
  pgrdf.graph_id('http://example.org/people'));
SELECT * FROM pgrdf.sparql('SELECT ?s ?p ?o WHERE { ?s ?p ?o } LIMIT 10');
SELECT * FROM pgrdf.surface();
```

### Operational Boundaries

The fixed `pgrdf` schema provides `add_graph`, `graph_id`, `parse_turtle`, `load_turtle`, `sparql`, `materialize`, `validate`, `stats` and `surface`. File loaders read server-side paths; grant access deliberately. `can_clear_graphs()` reports graph-clear privileges. Non-owner graph writes require the documented underlying table grants; graph locks still apply. `shacl_capability()` reports the supported validation surface, and path-depth/truncation settings bound traversal. Source features cover PostgreSQL 14–18; published Linux binaries target PG18. Version 0.6.39 allocates graph IDs from a non-reusing sequence; the IRI is the graph identity, and gaps are normal. After installing the matching library, run `ALTER EXTENSION pgrdf UPDATE` in each database. A newer library rejects old SQL with SQLSTATE 55000. Do not use the broken 0.6.35–0.6.38 PGXN source archives. Preserve graph/dictionary data together in verified backups.
