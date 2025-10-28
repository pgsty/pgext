---
title: "debversion"
linkTitle: "debversion"
description: "Debian version number data type"
weight: 3870
categories: ["TYPE"]
width: full
---

Debian version number data type


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3870** | {{< badge content="debversion" link="https://github.com/ATIX-AG/postgresql-debversion-evr" >}} | {{< ext "debversion" >}} | `1.2.0` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **Debian** | {{< badge content="PGDG" link="/e/debversion" >}} | `1.2.0` | {{< bg "18" "postgresql-18-debversion" "green" >}} {{< bg "17" "postgresql-17-debversion" "green" >}} {{< bg "16" "postgresql-16-debversion" "green" >}} {{< bg "15" "postgresql-15-debversion" "green" >}} {{< bg "14" "postgresql-14-debversion" "green" >}} {{< bg "13" "postgresql-13-debversion" "green" >}} | `postgresql-$v-debversion` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |      {{< bg "MISS" "debversion : MISS 0" "red" >}}      |
|    `d12.x86_64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 1.2.0" "postgresql-18-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-17-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-16-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-15-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-14-debversion : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.2.0" "postgresql-13-debversion : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-18-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.7 KiB | [postgresql-18-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 14.0 KiB | [postgresql-18-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 14.0 KiB | [postgresql-18-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.0 KiB | [postgresql-18-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 15.6 KiB | [postgresql-18-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 14.0 KiB | [postgresql-18-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-18-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-18-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-17-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.6 KiB | [postgresql-17-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 13.9 KiB | [postgresql-17-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 13.9 KiB | [postgresql-17-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.4 KiB | [postgresql-17-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 16.1 KiB | [postgresql-17-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 14.0 KiB | [postgresql-17-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-17-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-17-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-16-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.6 KiB | [postgresql-16-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 13.9 KiB | [postgresql-16-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 13.9 KiB | [postgresql-16-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.4 KiB | [postgresql-16-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 16.1 KiB | [postgresql-16-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 14.0 KiB | [postgresql-16-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-16-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-16-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-15-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.7 KiB | [postgresql-15-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 13.9 KiB | [postgresql-15-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 13.9 KiB | [postgresql-15-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.4 KiB | [postgresql-15-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 16.1 KiB | [postgresql-15-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 14.0 KiB | [postgresql-15-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-15-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-15-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-14-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.6 KiB | [postgresql-14-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 13.9 KiB | [postgresql-14-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 13.9 KiB | [postgresql-14-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.2 KiB | [postgresql-14-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 15.8 KiB | [postgresql-14-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 13.9 KiB | [postgresql-14-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-14-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-14-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-debversion` | 1.2.0 | `d12.x86_64` | pgdg | 14.6 KiB | [postgresql-13-debversion_1.2.0-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg12+1_amd64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `d12.aarch64` | pgdg | 14.6 KiB | [postgresql-13-debversion_1.2.0-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg12+1_arm64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `d13.x86_64` | pgdg | 13.9 KiB | [postgresql-13-debversion_1.2.0-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg13+1_amd64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `d13.aarch64` | pgdg | 13.9 KiB | [postgresql-13-debversion_1.2.0-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg13+1_arm64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `u22.x86_64` | pgdg | 16.2 KiB | [postgresql-13-debversion_1.2.0-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `u22.aarch64` | pgdg | 15.8 KiB | [postgresql-13-debversion_1.2.0-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `u24.x86_64` | pgdg | 13.9 KiB | [postgresql-13-debversion_1.2.0-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-debversion` | 1.2.0 | `u24.aarch64` | pgdg | 13.9 KiB | [postgresql-13-debversion_1.2.0-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgresql-debversion/postgresql-13-debversion_1.2.0-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/ATIX-AG/postgresql-debversion-evr" title="Repository" icon="github" subtitle="github.com/ATIX-AG/postgresql-debversion-evr" >}}
{{< /cards >}}


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install debversion; # install by extension name, for the current active PG version
pig ext install debversion; # install via package alias, for the active PG version
pig ext install debversion -v 18;   # install for PG 18
pig ext install debversion -v 17;   # install for PG 17
pig ext install debversion -v 16;   # install for PG 16
pig ext install debversion -v 15;   # install for PG 15
pig ext install debversion -v 14;   # install for PG 14
pig ext install debversion -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION debversion;
```

