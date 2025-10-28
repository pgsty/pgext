---
title: "rdkit"
linkTitle: "rdkit"
description: "Cheminformatics functionality for PostgreSQL."
weight: 2940
categories: ["FEAT"]
width: full
---

Cheminformatics functionality for PostgreSQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2940** | {{< badge content="rdkit" link="https://github.com/rdkit/rdkit" >}} | {{< ext "rdkit" >}} | `202503.1` | {{< category "FEAT" >}} | {{< license "BSD 3-Clause" >}} | {{< language "C++" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "age" >}} {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_graphql" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsquery" >}} {{< ext "pg_hint_plan" >}} {{< ext "hypopg" >}} |

> [!Note] u24 has rdkit for pg17


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **Debian** | {{< badge content="PGDG" link="/e/rdkit" >}} | `202503.1` | {{< bg "18" "postgresql-18-rdkit" "green" >}} {{< bg "17" "postgresql-17-rdkit" "green" >}} {{< bg "16" "postgresql-16-rdkit" "green" >}} {{< bg "15" "postgresql-15-rdkit" "green" >}} {{< bg "14" "postgresql-14-rdkit" "green" >}} {{< bg "13" "postgresql-13-rdkit" "green" >}} | `postgresql-$v-rdkit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "rdkit : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-rdkit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-rdkit : MISS 0" "red" >}}      | {{< bg "PGDG 202303.3" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202303.3" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 202503.1" "postgresql-18-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-17-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-16-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-15-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-14-rdkit : AVAIL 1" "blue" >}} | {{< bg "PGDG 202503.1" "postgresql-13-rdkit : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-18-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.6 KiB | [postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-18-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 243.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-18-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.1 KiB | [postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-18-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-17-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.5 KiB | [postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-17-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 243.1 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-17-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.2 KiB | [postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-17-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-rdkit` | 202303.3 | `d12.x86_64` | pgdg | 393.5 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-16-rdkit` | 202303.3 | `d12.aarch64` | pgdg | 384.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-16-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-16-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.5 KiB | [postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-16-rdkit` | 202303.3 | `u22.x86_64` | pgdg | 395.8 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-rdkit` | 202303.3 | `u22.aarch64` | pgdg | 387.1 KiB | [postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 243.2 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-16-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.0 KiB | [postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-16-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-rdkit` | 202303.3 | `d12.x86_64` | pgdg | 394.5 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-15-rdkit` | 202303.3 | `d12.aarch64` | pgdg | 385.2 KiB | [postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-15-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.2 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-15-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.5 KiB | [postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-15-rdkit` | 202303.3 | `u22.x86_64` | pgdg | 395.8 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-rdkit` | 202303.3 | `u22.aarch64` | pgdg | 387.0 KiB | [postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 243.1 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-15-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.0 KiB | [postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-15-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-rdkit` | 202303.3 | `d12.x86_64` | pgdg | 394.1 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-14-rdkit` | 202303.3 | `d12.aarch64` | pgdg | 385.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-14-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-14-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.2 KiB | [postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-14-rdkit` | 202303.3 | `u22.x86_64` | pgdg | 395.5 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-rdkit` | 202303.3 | `u22.aarch64` | pgdg | 387.2 KiB | [postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 242.9 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-14-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.0 KiB | [postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-14-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-rdkit` | 202303.3 | `d12.x86_64` | pgdg | 394.1 KiB | [postgresql-13-rdkit_202303.3-3.pgdg120+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202303.3-3.pgdg120+1_amd64.deb) |
| `postgresql-13-rdkit` | 202303.3 | `d12.aarch64` | pgdg | 385.1 KiB | [postgresql-13-rdkit_202303.3-3.pgdg120+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202303.3-3.pgdg120+1_arm64.deb) |
| `postgresql-13-rdkit` | 202503.1 | `d13.x86_64` | pgdg | 245.1 KiB | [postgresql-13-rdkit_202503.1-5.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202503.1-5.pgdg13+1_amd64.deb) |
| `postgresql-13-rdkit` | 202503.1 | `d13.aarch64` | pgdg | 237.4 KiB | [postgresql-13-rdkit_202503.1-5.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202503.1-5.pgdg13+1_arm64.deb) |
| `postgresql-13-rdkit` | 202303.3 | `u22.x86_64` | pgdg | 395.4 KiB | [postgresql-13-rdkit_202303.3-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202303.3-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-rdkit` | 202303.3 | `u22.aarch64` | pgdg | 386.7 KiB | [postgresql-13-rdkit_202303.3-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202303.3-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-rdkit` | 202503.1 | `u24.x86_64` | pgdg | 242.8 KiB | [postgresql-13-rdkit_202503.1-5.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202503.1-5.pgdg24.04+1_amd64.deb) |
| `postgresql-13-rdkit` | 202503.1 | `u24.aarch64` | pgdg | 237.1 KiB | [postgresql-13-rdkit_202503.1-5.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/r/rdkit/postgresql-13-rdkit_202503.1-5.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rdkit/rdkit" title="Repository" icon="github" subtitle="github.com/rdkit/rdkit" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install rdkit; # install by extension name, for the current active PG version
pig ext install rdkit; # install via package alias, for the active PG version
pig ext install rdkit -v 18;   # install for PG 18
pig ext install rdkit -v 17;   # install for PG 17
pig ext install rdkit -v 16;   # install for PG 16
pig ext install rdkit -v 15;   # install for PG 15
pig ext install rdkit -v 14;   # install for PG 14
pig ext install rdkit -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION rdkit;
```

