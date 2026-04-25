---
title: "rdkit"
linkTitle: "rdkit"
description: "Cheminformatics functionality for PostgreSQL."
weight: 2930
categories: ["FEAT"]
width: full
---

[**rdkit**](https://github.com/rdkit/rdkit) : Cheminformatics functionality for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2930** | {{< badge content="rdkit" link="https://github.com/rdkit/rdkit" >}} | {{< ext "rdkit" >}} | `202503.6` | {{< category "FEAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] d13/u24/el10 rdkit build by pigsty, u24/el10 deps on inchi 


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `202503.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `rdkit` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `202503.6` | {{< bg "18" "rdkit_18" "green" >}} {{< bg "17" "rdkit_17" "green" >}} {{< bg "16" "rdkit_16" "green" >}} {{< bg "15" "rdkit_15" "green" >}} {{< bg "14" "rdkit_14" "green" >}} | `rdkit_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `202503.6` | {{< bg "18" "postgresql-18-rdkit" "green" >}} {{< bg "17" "postgresql-17-rdkit" "green" >}} {{< bg "16" "postgresql-16-rdkit" "green" >}} {{< bg "15" "postgresql-15-rdkit" "green" >}} {{< bg "14" "postgresql-14-rdkit" "green" >}} | `postgresql-$v-rdkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "rdkit_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "rdkit_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "rdkit_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "rdkit_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 202503.6" "rdkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 202503.6" "rdkit_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 202503.6" "rdkit_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 202503.6" "postgresql-18-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-17-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-16-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-15-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-14-rdkit : AVAIL 2" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 202503.6" "postgresql-18-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-17-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-16-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-15-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-14-rdkit : AVAIL 2" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 202503.6" "postgresql-18-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-17-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-16-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-15-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-14-rdkit : AVAIL 2" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 202503.6" "postgresql-18-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-17-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-16-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-15-rdkit : AVAIL 2" "green" >}} | {{< bg "PIGSTY 202503.6" "postgresql-14-rdkit : AVAIL 2" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdkit_18` | `202503.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 143.1 KiB | [rdkit_18-202503.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdkit_18-202503.6-1PIGSTY.el10.x86_64.rpm) |
| `rdkit_18` | `202503.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.2 KiB | [rdkit_18-202503.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdkit_18-202503.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-rdkit` | `202503.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 102.8 KiB | [postgresql-18-rdkit_202503.6-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-18-rdkit_202503.6-4PIGSTY~trixie_amd64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-18-rdkit` | `202503.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.2 KiB | [postgresql-18-rdkit_202503.6-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-18-rdkit_202503.6-4PIGSTY~trixie_arm64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.6 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-18-rdkit` | `202503.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 108.6 KiB | [postgresql-18-rdkit_202503.6-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-18-rdkit_202503.6-4PIGSTY~noble_amd64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rdkit` | `202503.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 105.8 KiB | [postgresql-18-rdkit_202503.6-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-18-rdkit_202503.6-4PIGSTY~noble_arm64.deb) |
| `postgresql-18-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdkit_17` | `202503.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 143.2 KiB | [rdkit_17-202503.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdkit_17-202503.6-1PIGSTY.el10.x86_64.rpm) |
| `rdkit_17` | `202503.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.2 KiB | [rdkit_17-202503.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdkit_17-202503.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-rdkit` | `202503.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 103.0 KiB | [postgresql-17-rdkit_202503.6-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-17-rdkit_202503.6-4PIGSTY~trixie_amd64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-17-rdkit` | `202503.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.5 KiB | [postgresql-17-rdkit_202503.6-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-17-rdkit_202503.6-4PIGSTY~trixie_arm64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-17-rdkit` | `202503.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 108.6 KiB | [postgresql-17-rdkit_202503.6-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-17-rdkit_202503.6-4PIGSTY~noble_amd64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rdkit` | `202503.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 105.8 KiB | [postgresql-17-rdkit_202503.6-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-17-rdkit_202503.6-4PIGSTY~noble_arm64.deb) |
| `postgresql-17-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdkit_16` | `202503.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 143.1 KiB | [rdkit_16-202503.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdkit_16-202503.6-1PIGSTY.el10.x86_64.rpm) |
| `rdkit_16` | `202503.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.2 KiB | [rdkit_16-202503.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdkit_16-202503.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 393.5 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 384.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-16-rdkit` | `202503.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 102.9 KiB | [postgresql-16-rdkit_202503.6-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-16-rdkit_202503.6-4PIGSTY~trixie_amd64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-16-rdkit` | `202503.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.3 KiB | [postgresql-16-rdkit_202503.6-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-16-rdkit_202503.6-4PIGSTY~trixie_arm64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.1 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rdkit` | `202503.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 108.5 KiB | [postgresql-16-rdkit_202503.6-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-16-rdkit_202503.6-4PIGSTY~noble_amd64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rdkit` | `202503.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 105.8 KiB | [postgresql-16-rdkit_202503.6-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-16-rdkit_202503.6-4PIGSTY~noble_arm64.deb) |
| `postgresql-16-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdkit_15` | `202503.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 143.1 KiB | [rdkit_15-202503.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdkit_15-202503.6-1PIGSTY.el10.x86_64.rpm) |
| `rdkit_15` | `202503.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.2 KiB | [rdkit_15-202503.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdkit_15-202503.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 394.5 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 385.2 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-15-rdkit` | `202503.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 103.1 KiB | [postgresql-15-rdkit_202503.6-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-15-rdkit_202503.6-4PIGSTY~trixie_amd64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-15-rdkit` | `202503.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.4 KiB | [postgresql-15-rdkit_202503.6-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-15-rdkit_202503.6-4PIGSTY~trixie_arm64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.5 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.8 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.0 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rdkit` | `202503.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 108.7 KiB | [postgresql-15-rdkit_202503.6-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-15-rdkit_202503.6-4PIGSTY~noble_amd64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 243.1 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rdkit` | `202503.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 105.8 KiB | [postgresql-15-rdkit_202503.6-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-15-rdkit_202503.6-4PIGSTY~noble_arm64.deb) |
| `postgresql-15-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `rdkit_14` | `202503.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 143.1 KiB | [rdkit_14-202503.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/rdkit_14-202503.6-1PIGSTY.el10.x86_64.rpm) |
| `rdkit_14` | `202503.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 135.1 KiB | [rdkit_14-202503.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/rdkit_14-202503.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-rdkit` | `202303.3` | [d12.x86_64](/os/d12.x86_64) | pgdg | 394.1 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [d12.aarch64](/os/d12.aarch64) | pgdg | 385.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-14-rdkit` | `202503.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 102.7 KiB | [postgresql-14-rdkit_202503.6-4PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-14-rdkit_202503.6-4PIGSTY~trixie_amd64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 245.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-14-rdkit` | `202503.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 94.2 KiB | [postgresql-14-rdkit_202503.6-4PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/r/rdkit/postgresql-14-rdkit_202503.6-4PIGSTY~trixie_arm64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 237.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [u22.x86_64](/os/u22.x86_64) | pgdg | 395.5 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-rdkit` | `202303.3` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-rdkit` | `202503.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 108.6 KiB | [postgresql-14-rdkit_202503.6-4PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-14-rdkit_202503.6-4PIGSTY~noble_amd64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 242.9 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-rdkit` | `202503.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 105.7 KiB | [postgresql-14-rdkit_202503.6-4PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/r/rdkit/postgresql-14-rdkit_202503.6-4PIGSTY~noble_arm64.deb) |
| `postgresql-14-rdkit` | `202503.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 237.0 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rdkit/rdkit" title="Repository" icon="github" subtitle="github.com/rdkit/rdkit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="rdkit_202503.6.orig.tar.xz" >}}
{{< /cards >}}


```bash
pig build pkg rdkit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install rdkit;		# install via package name, for the active PG version

pig install rdkit -v 18;   # install for PG 18
pig install rdkit -v 17;   # install for PG 17
pig install rdkit -v 16;   # install for PG 16
pig install rdkit -v 15;   # install for PG 15
pig install rdkit -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION rdkit;
```

## Usage

- Sources: [project README](https://github.com/rdkit/rdkit/blob/master/README.md), [cartridge docs](https://www.rdkit.org/docs/Cartridge.html), [2025.03.6 release](https://github.com/rdkit/rdkit/releases/tag/Release_2025.03.6)

RDKit ships a PostgreSQL cartridge for cheminformatics storage, search, fingerprints, and descriptors. The cartridge docs remain the main upstream usage reference; the 2025.03.6 release notes do not call out cartridge-specific user-facing changes.

### Create The Extension

```sql
CREATE EXTENSION rdkit;
```

The cartridge adds chemistry-specific types including `mol`, `bfp`, and `sfp`.

### Core Search Operators

The cartridge documentation covers:

- `@>` and `<@` for substructure matching.
- `@=` for exact molecular equality.
- `%`, `<%>`, and `<#>` style fingerprint similarity and KNN operators for similarity search.

These are typically combined with GiST indexes over fingerprint columns.

### Fingerprints And Similarity

Common fingerprint functions documented for SQL usage include `morgan_fp`, `morganbv_fp`, `featmorgan_fp`, `rdkit_fp`, `atompair_fp`, `torsion_fp`, `layered_fp`, and `maccs_fp`.

Example from the cartridge docs:

```sql
SELECT tanimoto_sml(
  morganbv_fp('c1ccccc1'::mol),
  morganbv_fp('c1ccccc1O'::mol)
);
```

### Descriptors And Validation

The cartridge docs also expose validation and descriptor helpers such as:

- `is_valid_smiles()`
- `is_valid_ctab()`
- `is_valid_smarts()`
- `mol_amw()`
- `mol_hba()`
- `mol_numrings()`

These functions are the main user-facing surface for SQL analytics on molecular structures.
