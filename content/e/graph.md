---
title: "graph"
linkTitle: "graph"
description: "Graph database capabilities for PostgreSQL"
weight: 2630
categories: ["FEAT"]
width: full
---

[**pggraph**](https://github.com/evokoa/pggraph) : Graph database capabilities for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2630** | {{< badge content="graph" link="https://github.com/evokoa/pggraph" >}} | {{< ext "graph" "pggraph" >}} | `1.0.0` | {{< category "FEAT" >}} | {{< license "Apache-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "pg_graphql" >}} |

> [!Note] PGXN distribution and package are pggraph; installed extension name is graph.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pggraph` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pggraph_18" "green" >}} {{< bg "17" "pggraph_17" "green" >}} {{< bg "16" "pggraph_16" "green" >}} {{< bg "15" "pggraph_15" "green" >}} {{< bg "14" "pggraph_14" "green" >}} | `pggraph_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pggraph" "green" >}} {{< bg "17" "postgresql-17-pggraph" "green" >}} {{< bg "16" "postgresql-16-pggraph" "green" >}} {{< bg "15" "postgresql-15-pggraph" "green" >}} {{< bg "14" "postgresql-14-pggraph" "green" >}} | `postgresql-$v-pggraph` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pggraph_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pggraph : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pggraph : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pggraph_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.0 MiB | [pggraph_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pggraph_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pggraph_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pggraph_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pggraph_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pggraph_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.9 MiB | [pggraph_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pggraph_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pggraph_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pggraph_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pggraph_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pggraph_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.9 MiB | [pggraph_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pggraph_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pggraph_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pggraph_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pggraph_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pggraph` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.3 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.3 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.6 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.6 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pggraph` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-18-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-18-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pggraph_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.0 MiB | [pggraph_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pggraph_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pggraph_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pggraph_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pggraph_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pggraph_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.9 MiB | [pggraph_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pggraph_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pggraph_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pggraph_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pggraph_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pggraph_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.9 MiB | [pggraph_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pggraph_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pggraph_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pggraph_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pggraph_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pggraph` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.3 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.3 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.6 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.6 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pggraph` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-17-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-17-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pggraph_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.0 MiB | [pggraph_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pggraph_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pggraph_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pggraph_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pggraph_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pggraph_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.9 MiB | [pggraph_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pggraph_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pggraph_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pggraph_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pggraph_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pggraph_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.9 MiB | [pggraph_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pggraph_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pggraph_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pggraph_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pggraph_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pggraph` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.3 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.3 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.6 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.6 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pggraph` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-16-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-16-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pggraph_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.0 MiB | [pggraph_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pggraph_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pggraph_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pggraph_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pggraph_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pggraph_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.9 MiB | [pggraph_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pggraph_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pggraph_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pggraph_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pggraph_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pggraph_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.9 MiB | [pggraph_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pggraph_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pggraph_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pggraph_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pggraph_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pggraph` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.3 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.7 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.3 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.6 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.6 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pggraph` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-15-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-15-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pggraph_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 4.0 MiB | [pggraph_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pggraph_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pggraph_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.4 MiB | [pggraph_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pggraph_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pggraph_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.9 MiB | [pggraph_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pggraph_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pggraph_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.5 MiB | [pggraph_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pggraph_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pggraph_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.9 MiB | [pggraph_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pggraph_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pggraph_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.5 MiB | [pggraph_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pggraph_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pggraph` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.3 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.7 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.3 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.7 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.6 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.5 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pggraph` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.1 MiB | [postgresql-14-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pggraph/postgresql-14-pggraph_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/evokoa/pggraph" title="Repository" icon="github" subtitle="github.com/evokoa/pggraph" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pggraph-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pggraph;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pggraph;		# install via package name, for the active PG version
pig install graph;		# install by extension name, for the current active PG version

pig install graph -v 18;   # install for PG 18
pig install graph -v 17;   # install for PG 17
pig install graph -v 16;   # install for PG 16
pig install graph -v 15;   # install for PG 15
pig install graph -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION graph;
```

## Usage

Sources:

- [pgGraph v1.0.0 README](https://github.com/evokoa/pggraph/blob/v1.0.0/README.md)
- [v1.0.0 release notes](https://github.com/evokoa/pggraph/blob/v1.0.0/docs/release-notes.mdx)
- [SQL API Reference](https://github.com/evokoa/pggraph/blob/v1.0.0/docs/user_guide/api-reference.mdx)
- [Schema Registration](https://github.com/evokoa/pggraph/blob/v1.0.0/docs/user_guide/schema-registration.mdx)
- [Administration and Security](https://github.com/evokoa/pggraph/blob/v1.0.0/docs/user_guide/administration-and-security.mdx)
- [v0.1.8 to v1.0.0 migration guide](https://github.com/evokoa/pggraph/blob/v1.0.0/docs/user_guide/migration-1-0.mdx)

`pggraph` is the package and PGXN distribution name, but the installed PostgreSQL extension is `graph`. The extension builds derived graph artifacts from ordinary PostgreSQL tables, keeps those tables as the source of truth, and exposes graph search, traversal, shortest path, GQL-style reads, and selected mapped writes through the `graph` schema.

v1.0.0 is the first production release. It supports PostgreSQL 14-18, named graphs, graph-scoped grants and quotas, durable synchronization, bounded traversal and analytics, maintenance jobs, and selected GQL read/write profiles. It does not claim full ISO GQL, full openCypher, or a public SQL/PGQ `GRAPH_TABLE` surface. Standard PostgreSQL SQLSTATEs are paired with stable `PGxxx` details for application diagnostics.

### Basic Graph Build

```sql
CREATE EXTENSION IF NOT EXISTS graph;
SELECT graph.reset();

CREATE TABLE companies (
  id   text PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE people (
  id         text PRIMARY KEY,
  name       text NOT NULL,
  company_id text REFERENCES companies(id)
);

INSERT INTO companies VALUES
  ('c1', 'Acme Bank'),
  ('c2', 'Northwind Trading');

INSERT INTO people VALUES
  ('p1', 'Alice', 'c1'),
  ('p2', 'Bob', 'c1'),
  ('p3', 'Carol', 'c2');

SELECT * FROM graph.auto_discover('public');
SELECT * FROM graph.build();

SELECT node_count, edge_count, edge_types
FROM graph.status();
```

`graph.auto_discover('public')` scans primary keys and foreign keys in the selected schema, registers discovered source tables and edges, and prepares the graph for `graph.build()`. In production schemas, prefer explicit registration so labels, search columns, filter columns, weights, tenant behavior, and graph identity are deliberate.

### Manual and Named-Graph Registration

```sql
SELECT graph.create_graph('customer_360', namespace := 'analytics');
SELECT graph.set_current_graph('customer_360', namespace := 'analytics');

SELECT graph.add_table(
  table_name := 'public.people'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  tenant_column := NULL
);

SELECT graph.add_table_to_graph(
  graph_name := 'customer_360',
  table_name := 'public.companies'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  graph_namespace := 'analytics'
);

SELECT graph.add_edge_to_graph(
  graph_name := 'customer_360',
  from_table := 'public.people'::regclass,
  from_column := 'company_id',
  to_table := 'public.companies'::regclass,
  to_column := 'id',
  label := 'works_at',
  bidirectional := true,
  graph_namespace := 'analytics'
);

SELECT * FROM graph.build_graph('customer_360', graph_namespace := 'analytics');
```

Registration applies to the current graph selection unless you use the explicit `*_to_graph` and `*_from_graph` helpers. Node identifiers must match a primary key or a unique `NOT NULL` index. `columns` controls searchable and GQL-visible properties; traversal filter pushdown uses separate `graph.add_filter_column()` registrations. Edge-table and junction-table relationships are also supported, and `label_column` can provide dynamic edge labels within the documented public limit.

### Search, Traversal, and Paths

```sql
SELECT node_table_name, node_id, node
FROM graph.search(
  property_key   := 'name',
  property_value := 'Alice',
  table_filter   := 'public.people'::regclass,
  mode           := 'exact',
  hydrate        := true
);

SELECT depth, node_table_name, node_id, edge_path
FROM graph.traverse(
  'public.people'::regclass,
  'p1',
  2,
  hydrate := false
);

SELECT step, node_table_name, node_id, edge_label
FROM graph.shortest_path(
  'public.people'::regclass,
  'p1',
  'public.companies'::regclass,
  'c1',
  hydrate := false
);
```

With `hydrate := false`, graph functions return compact graph coordinates. With hydration enabled, PostgreSQL source-table ACLs and RLS still govern which source rows are visible. Stale coordinates fail closed rather than fabricating rows.

### GQL Queries and Relationship Writes

```sql
SELECT row
FROM graph.gql(
  'MATCH (p:people)-[:works_at]->(c:companies)
   WHERE p.name = $name
   RETURN p.id AS person_id, c.name AS company
   ORDER BY company',
  params  := '{"name":"Alice"}'::jsonb,
  hydrate := true
);
```

`graph.gql()` returns one `jsonb` object per SQL row. Node labels map to registered table names and relationship types map to registered edge labels. The supported mutable GQL profile includes registered relationship creation: mapped writes still go through PostgreSQL source-table DML, and source tables remain authoritative. Unsupported openCypher or SQL/PGQ shapes fail with explicit capability errors instead of partial behavior.

### Administration and Operations

```sql
GRANT USAGE, CREATE ON SCHEMA graph TO graph_admin;

SELECT * FROM graph.grant_graph('customer_360', 'app_reader', 'read', namespace := 'analytics');
SELECT * FROM graph.set_graph_quota('owner', 'max_named_graphs', 25, scope_key := 'app_owner');
SELECT * FROM graph.select_graph('customer_360', namespace := 'analytics');
SELECT * FROM graph.add_sync_policy('customer_360', schedule_interval_secs := 300, graph_namespace := 'analytics');
SELECT * FROM graph.run_due_jobs();
SELECT * FROM graph.projection_status();
```

Graph administration covers catalog mutation, builds, sync replay, maintenance, quotas, runtime graph loading, and global analytics. Named graph privileges are `read`, `write`, `build`, and `admin`, but graph `read` is not enough by itself: hydrated reads still require `SELECT` on source tables. A selected graph tenant also scopes traversal, search, GQL, and Cypher calls unless an explicit matching tenant is supplied.

### Migrating from the Alpha Release

The v0.1.8 to v1.0.0 transition is source-preserving but is not an in-place catalog or binary update. Back up and test a restore, inventory registrations and dependents, stop graph writers and schedulers, then preflight the drop in a transaction:

```sql
BEGIN;
DROP EXTENSION graph;
ROLLBACK;
```

After reviewing every dependent object, remove the alpha extension, install v1.0.0, reapply only reviewed public registration calls, and rebuild from the PostgreSQL source tables:

```sql
DROP EXTENSION graph CASCADE;
CREATE EXTENSION graph VERSION '1.0.0';

-- Reapply graph.add_table(...), graph.add_edge(...), and related calls.
SELECT * FROM graph.build();
SELECT * FROM graph.status();
```

`CASCADE` can remove application views, functions, generated synchronization objects, and other dependents. Alpha catalogs, `.pggraph` files, manifests, and projection segments are not v1.0.0 portable state. Rollback requires restoring the tested backup with the matching alpha package, then rebuilding its graph state.

### Caveats

- Source tables remain the source of truth. Graph artifacts, projection files, sync state, and runtime engines are derived and rebuildable.
- Use `graph.build()` or graph-scoped build helpers after registration changes, and use sync/maintenance APIs when relying on incremental projection state.
- Internal catalog tables such as `graph._graphs`, grants, quotas, jobs, sync logs, and projection metadata are implementation details; use public SQL functions instead.
- v1.0.0 uses Rust 1.96 and `cargo-pgrx` 0.19.1 for source builds. PostgreSQL 14 through 18 are supported upstream, with PostgreSQL 17 as the default release-gate target.
