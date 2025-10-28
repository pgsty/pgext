---
title: "pg_rrule"
linkTitle: "pg_rrule"
description: "RRULE field type for PostgreSQL"
weight: 3880
categories: ["TYPE"]
width: full
---

RRULE field type for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3880** | {{< badge content="pg_rrule" link="https://github.com/petropavel13/pg_rrule" >}} | {{< ext "pg_rrule" >}} | `0.2.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **Debian** | {{< badge content="PGDG" link="/e/pg_rrule" >}} | `0.2.0` | {{< bg "18" "postgresql-18-pg-rrule" "green" >}} {{< bg "17" "postgresql-17-pg-rrule" "green" >}} {{< bg "16" "postgresql-16-pg-rrule" "green" >}} {{< bg "15" "postgresql-15-pg-rrule" "green" >}} {{< bg "14" "postgresql-14-pg-rrule" "green" >}} {{< bg "13" "postgresql-13-pg-rrule" "green" >}} | `postgresql-$v-pg-rrule` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_rrule : MISS 0" "red" >}}      |
|    `d12.x86_64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `d13.x86_64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `d13.aarch64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 0.2.0" "postgresql-18-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-17-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-16-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-15-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-14-pg-rrule : AVAIL 1" "blue" >}} | {{< bg "PGDG 0.2.0" "postgresql-13-pg-rrule : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 23.9 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 23.9 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.1 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 24.1 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 23.4 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 23.9 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 23.9 KiB | [postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-18-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 24.0 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 23.9 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.1 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 24.8 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 24.7 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 24.0 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 23.9 KiB | [postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-17-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 24.0 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 23.9 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.1 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 24.8 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 24.8 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 23.9 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 23.9 KiB | [postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-16-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 23.9 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.1 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 23.9 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.1 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 25.0 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 24.3 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 24.1 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 24.0 KiB | [postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-15-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 24.1 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.0 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 24.0 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.1 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 24.9 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 24.3 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 24.2 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 24.0 KiB | [postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-14-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-13-pg-rrule` | 0.2.0 | `d12.x86_64` | pgdg | 23.8 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_amd64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `d12.aarch64` | pgdg | 24.1 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg12+1_arm64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `d13.x86_64` | pgdg | 23.8 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_amd64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `d13.aarch64` | pgdg | 24.2 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg13+1_arm64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `u22.x86_64` | pgdg | 24.6 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_amd64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `u22.aarch64` | pgdg | 24.2 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg22.04+1_arm64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `u24.x86_64` | pgdg | 23.9 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_amd64.deb) |
| `postgresql-13-pg-rrule` | 0.2.0 | `u24.aarch64` | pgdg | 24.1 KiB | [postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-rrule/postgresql-13-pg-rrule_0.2.0+git20211101.d7d10f2-4.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petropavel13/pg_rrule" title="Repository" icon="github" subtitle="github.com/petropavel13/pg_rrule" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_rrule-0.2.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_rrule; # get pg_rrule source code
pig build dep pg_rrule; # install build dependencies
pig build pkg pg_rrule; # build extension rpm or deb
pig build ext pg_rrule; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_rrule; # install by extension name, for the current active PG version
pig ext install pg_rrule; # install via package alias, for the active PG version
pig ext install pg_rrule -v 18;   # install for PG 18
pig ext install pg_rrule -v 17;   # install for PG 17
pig ext install pg_rrule -v 16;   # install for PG 16
pig ext install pg_rrule -v 15;   # install for PG 15
pig ext install pg_rrule -v 14;   # install for PG 14
pig ext install pg_rrule -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_rrule;
```

