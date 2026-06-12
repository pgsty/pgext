---
title: "jsonschema"
linkTitle: "jsonschema"
description: "JSON Schema validation functions for PostgreSQL"
weight: 2760
categories: ["FEAT"]
width: full
---

[**jsonschema**](https://github.com/theory/pg-jsonschema-boon) : JSON Schema validation functions for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2760** | {{< badge content="jsonschema" link="https://github.com/theory/pg-jsonschema-boon" >}} | {{< ext "jsonschema" >}} | `0.1.9` | {{< category "FEAT" >}} | {{< license "MIT" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_graphql" >}} {{< ext "plv8" >}} |

> [!Note] Distinct from Supabase pg_jsonschema.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `jsonschema` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "jsonschema_18" "green" >}} {{< bg "17" "jsonschema_17" "green" >}} {{< bg "16" "jsonschema_16" "green" >}} {{< bg "15" "jsonschema_15" "green" >}} {{< bg "14" "jsonschema_14" "green" >}} | `jsonschema_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.9` | {{< bg "18" "postgresql-18-jsonschema" "green" >}} {{< bg "17" "postgresql-17-jsonschema" "green" >}} {{< bg "16" "postgresql-16-jsonschema" "green" >}} {{< bg "15" "postgresql-15-jsonschema" "green" >}} {{< bg "14" "postgresql-14-jsonschema" "green" >}} | `postgresql-$v-jsonschema` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "jsonschema_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-18-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-17-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-16-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-15-jsonschema : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.9" "postgresql-14-jsonschema : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsonschema_18` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 924.8 KiB | [jsonschema_18-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jsonschema_18-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `jsonschema_18` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 806.8 KiB | [jsonschema_18-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jsonschema_18-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `jsonschema_18` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 936.5 KiB | [jsonschema_18-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jsonschema_18-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `jsonschema_18` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 875.2 KiB | [jsonschema_18-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jsonschema_18-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `jsonschema_18` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 933.5 KiB | [jsonschema_18-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jsonschema_18-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `jsonschema_18` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 875.7 KiB | [jsonschema_18-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jsonschema_18-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-jsonschema` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 785.2 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 664.6 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 786.4 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.6 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 873.1 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 797.3 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 862.5 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 782.2 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.4 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-jsonschema` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 778.5 KiB | [postgresql-18-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-18-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsonschema_17` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 924.9 KiB | [jsonschema_17-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jsonschema_17-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `jsonschema_17` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 806.8 KiB | [jsonschema_17-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jsonschema_17-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `jsonschema_17` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 936.7 KiB | [jsonschema_17-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jsonschema_17-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `jsonschema_17` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 875.0 KiB | [jsonschema_17-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jsonschema_17-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `jsonschema_17` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 940.4 KiB | [jsonschema_17-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jsonschema_17-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `jsonschema_17` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 875.3 KiB | [jsonschema_17-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jsonschema_17-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-jsonschema` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 787.1 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 664.5 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 786.6 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.6 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 873.2 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 792.2 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 866.3 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 787.8 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 859.1 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-jsonschema` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 778.4 KiB | [postgresql-17-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-17-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsonschema_16` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 924.9 KiB | [jsonschema_16-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jsonschema_16-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `jsonschema_16` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 806.7 KiB | [jsonschema_16-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jsonschema_16-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `jsonschema_16` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 936.6 KiB | [jsonschema_16-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jsonschema_16-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `jsonschema_16` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 875.2 KiB | [jsonschema_16-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jsonschema_16-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `jsonschema_16` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 936.4 KiB | [jsonschema_16-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jsonschema_16-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `jsonschema_16` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 875.5 KiB | [jsonschema_16-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jsonschema_16-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-jsonschema` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 786.5 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 664.6 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 785.3 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.9 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 873.1 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 792.4 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 866.7 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 782.7 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.3 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-jsonschema` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 780.0 KiB | [postgresql-16-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-16-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsonschema_15` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 924.5 KiB | [jsonschema_15-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jsonschema_15-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `jsonschema_15` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 807.1 KiB | [jsonschema_15-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jsonschema_15-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `jsonschema_15` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 937.5 KiB | [jsonschema_15-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jsonschema_15-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `jsonschema_15` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 875.2 KiB | [jsonschema_15-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jsonschema_15-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `jsonschema_15` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 936.1 KiB | [jsonschema_15-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jsonschema_15-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `jsonschema_15` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 875.6 KiB | [jsonschema_15-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jsonschema_15-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-jsonschema` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 786.7 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 665.4 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 786.2 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 665.1 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 873.6 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 797.6 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 864.6 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 786.6 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.2 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-jsonschema` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 781.5 KiB | [postgresql-15-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-15-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jsonschema_14` | `0.1.9` | [el8.x86_64](/os/el8.x86_64) | pigsty | 924.5 KiB | [jsonschema_14-0.1.9-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jsonschema_14-0.1.9-1PIGSTY.el8.x86_64.rpm) |
| `jsonschema_14` | `0.1.9` | [el8.aarch64](/os/el8.aarch64) | pigsty | 807.0 KiB | [jsonschema_14-0.1.9-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jsonschema_14-0.1.9-1PIGSTY.el8.aarch64.rpm) |
| `jsonschema_14` | `0.1.9` | [el9.x86_64](/os/el9.x86_64) | pigsty | 940.1 KiB | [jsonschema_14-0.1.9-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jsonschema_14-0.1.9-1PIGSTY.el9.x86_64.rpm) |
| `jsonschema_14` | `0.1.9` | [el9.aarch64](/os/el9.aarch64) | pigsty | 875.6 KiB | [jsonschema_14-0.1.9-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jsonschema_14-0.1.9-1PIGSTY.el9.aarch64.rpm) |
| `jsonschema_14` | `0.1.9` | [el10.x86_64](/os/el10.x86_64) | pigsty | 939.9 KiB | [jsonschema_14-0.1.9-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jsonschema_14-0.1.9-1PIGSTY.el10.x86_64.rpm) |
| `jsonschema_14` | `0.1.9` | [el10.aarch64](/os/el10.aarch64) | pigsty | 875.7 KiB | [jsonschema_14-0.1.9-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jsonschema_14-0.1.9-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-jsonschema` | `0.1.9` | [d12.x86_64](/os/d12.x86_64) | pigsty | 787.5 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [d12.aarch64](/os/d12.aarch64) | pigsty | 665.3 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [d13.x86_64](/os/d13.x86_64) | pigsty | 786.2 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [d13.aarch64](/os/d13.aarch64) | pigsty | 664.9 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u22.x86_64](/os/u22.x86_64) | pigsty | 872.9 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u22.aarch64](/os/u22.aarch64) | pigsty | 796.4 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u24.x86_64](/os/u24.x86_64) | pigsty | 860.6 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u24.aarch64](/os/u24.aarch64) | pigsty | 782.6 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u26.x86_64](/os/u26.x86_64) | pigsty | 856.4 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-jsonschema` | `0.1.9` | [u26.aarch64](/os/u26.aarch64) | pigsty | 779.0 KiB | [postgresql-14-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jsonschema/postgresql-14-jsonschema_0.1.9-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/theory/pg-jsonschema-boon" title="Repository" icon="github" subtitle="github.com/theory/pg-jsonschema-boon" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="jsonschema-0.1.9.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg jsonschema;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install jsonschema;		# install via package name, for the active PG version

pig install jsonschema -v 18;   # install for PG 18
pig install jsonschema -v 17;   # install for PG 17
pig install jsonschema -v 16;   # install for PG 16
pig install jsonschema -v 15;   # install for PG 15
pig install jsonschema -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jsonschema;
```

Source: [jsonschema v0.1.9 README](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/README.md), [documentation](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/doc/jsonschema.md), [control file](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/jsonschema.control), [SQL definition](https://github.com/theory/pg-jsonschema-boon/blob/v0.1.9/sql/jsonschema--0.1.9.sql).

## Usage

`jsonschema` validates JSON and JSONB values against JSON Schema inside PostgreSQL. It is the `theory/pg-jsonschema-boon` extension and is distinct from Supabase `pg_jsonschema`, although it provides compatibility wrappers named `json_matches_schema()` and `jsonb_matches_schema()`.

The extension supports JSON Schema draft 4, draft 6, draft 7, draft 2019-09, and draft 2020-12 through the Rust `boon` validator. It has no runtime dependency beyond PostgreSQL.

### Validate a Schema and a Document

```sql
CREATE EXTENSION IF NOT EXISTS jsonschema;

SELECT jsonschema_is_valid(
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "age":   { "type": "number", "minimum": 0 },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);

SELECT jsonschema_validates(
  '{"name":"Amos Burton","email":"amos@rocinante.ship"}'::json,
  '{
     "type": "object",
     "required": ["name", "email"],
     "properties": {
       "name":  { "type": "string" },
       "email": { "type": "string", "format": "email" }
     }
   }'::json
);
```

`jsonschema_is_valid(schema)` returns whether the schema itself compiles and validates against the selected draft. `jsonschema_validates(data, schema)` returns whether the JSON/JSONB value satisfies the schema.

### Check Constraints

```sql
CREATE TABLE customer_profile (
  id       bigserial PRIMARY KEY,
  profile  jsonb NOT NULL,
  CHECK (
    jsonschema_validates(
      profile,
      '{
         "type": "object",
         "required": ["email"],
         "properties": {
           "email": { "type": "string", "format": "email" },
           "tags":  {
             "type": "array",
             "items": { "type": "string", "maxLength": 16 }
           }
         }
       }'::jsonb
    )
  )
);
```

Use constraints when the database should reject malformed JSON documents at write time.

### Composed Schemas

```sql
SELECT jsonschema_validates(
  jsonb_build_object(
    'first_name', 'Naomi',
    'last_name', 'Nagata',
    'shipping_address', jsonb_build_object(
      'street_address', '1 Rocinante Way',
      'city', 'Ceres Station',
      'state', 'The Belt'
    )
  ),
  'https://example.com/schemas/customer',
  '{
     "$id": "https://example.com/schemas/address",
     "type": "object",
     "required": ["street_address", "city", "state"],
     "properties": {
       "street_address": { "type": "string" },
       "city": { "type": "string" },
       "state": { "type": "string" }
     }
   }'::jsonb,
  '{
     "$id": "https://example.com/schemas/customer",
     "type": "object",
     "required": ["first_name", "last_name", "shipping_address"],
     "properties": {
       "first_name": { "type": "string" },
       "last_name": { "type": "string" },
       "shipping_address": { "$ref": "/schemas/address" }
     }
   }'::jsonb
);
```

The `id` overloads let multiple schemas reference each other by `$id`, which is useful for componentized JSON Schema definitions.

### Compatibility Functions

```sql
SELECT json_matches_schema(
  '{"type":"string","maxLength":4}'::json,
  '"1234"'::json
);

SELECT jsonb_matches_schema(
  '{"type":"object","required":["id"]}'::json,
  '{"id":42}'::jsonb
);
```

These wrappers mirror the common `pg_jsonschema` argument order: schema first, instance second.

### Draft Selection and Caveats

```sql
SET jsonschema.default_draft = 'V2020';
SET jsonschema.default_draft = 'V7';
```

If a schema omits `$schema`, `jsonschema.default_draft` controls the default draft. Supported values are `V4`, `V6`, `V7`, `V2019`, and `V2020`.

- `jsonschema_validates(data, schema)` returns NULL if either argument is NULL.
- Invalid or uncompilable schemas can raise errors in validation calls; failed document validation returns `false` and logs details at `INFO`.
- `jsonschema_is_valid(id, VARIADIC schemas)` and `jsonschema_validates(data, id, VARIADIC schemas)` require matching `$id` values for reliable composed-schema resolution.
