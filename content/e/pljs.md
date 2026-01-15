---
title: "pljs"
linkTitle: "pljs"
description: "PL/JS trusted procedural language"
weight: 3011
categories: ["LANG"]
width: full
---

[**pljs**](https://github.com/plv8/pljs) : PL/JS trusted procedural language


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3011** | {{< badge content="pljs" link="https://github.com/plv8/pljs" >}} | {{< ext "pljs" >}} | `1.0.4` | {{< category "LANG" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "plv8" >}} {{< ext "jsquery" >}} {{< ext "pllua" >}} {{< ext "pg_tle" >}} {{< ext "plpgsql" >}} {{< ext "pg_jsonschema" >}} {{< ext "plperl" >}} {{< ext "plpython3u" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pljs` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.0.4` | {{< bg "18" "postgresql-18-pljs" "green" >}} {{< bg "17" "postgresql-17-pljs" "green" >}} {{< bg "16" "postgresql-16-pljs" "green" >}} {{< bg "15" "postgresql-15-pljs" "green" >}} {{< bg "14" "postgresql-14-pljs" "green" >}} {{< bg "13" "postgresql-13-pljs" "red" >}} | `postgresql-$v-pljs` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |      {{< bg "MISS" "pljs : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.0.4" "postgresql-18-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-17-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-16-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-15-pljs : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.4" "postgresql-14-pljs : AVAIL 1" "blue" >}} |      {{< bg "MISS" "postgresql-13-pljs : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-18-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-18-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 374.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.6 KiB | [postgresql-18-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 406.8 KiB | [postgresql-18-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 372.9 KiB | [postgresql-18-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.7 KiB | [postgresql-18-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.4 KiB | [postgresql-18-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-18-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-17-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-17-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.1 KiB | [postgresql-17-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 428.8 KiB | [postgresql-17-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.6 KiB | [postgresql-17-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 423.1 KiB | [postgresql-17-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.6 KiB | [postgresql-17-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.3 KiB | [postgresql-17-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 375.9 KiB | [postgresql-17-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-17-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-16-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 428.9 KiB | [postgresql-16-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.8 KiB | [postgresql-16-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.6 KiB | [postgresql-16-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.3 KiB | [postgresql-16-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.6 KiB | [postgresql-16-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.0 KiB | [postgresql-16-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-16-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-15-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-15-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.7 KiB | [postgresql-15-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 382.8 KiB | [postgresql-15-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 421.6 KiB | [postgresql-15-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 387.9 KiB | [postgresql-15-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.5 KiB | [postgresql-15-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-15-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgresql-14-pljs` | `1.0.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 410.1 KiB | [postgresql-14-pljs_1.0.4-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg12+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 375.3 KiB | [postgresql-14-pljs_1.0.4-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg12+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 429.3 KiB | [postgresql-14-pljs_1.0.4-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg13+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 381.5 KiB | [postgresql-14-pljs_1.0.4-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg13+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 422.0 KiB | [postgresql-14-pljs_1.0.4-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 388.1 KiB | [postgresql-14-pljs_1.0.4-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 408.2 KiB | [postgresql-14-pljs_1.0.4-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-pljs` | `1.0.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 376.4 KiB | [postgresql-14-pljs_1.0.4-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pljs/postgresql-14-pljs_1.0.4-1.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/plv8/pljs" title="Repository" icon="github" subtitle="github.com/plv8/pljs" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pljs;		# install via package name, for the active PG version

pig install pljs -v 18;   # install for PG 18
pig install pljs -v 17;   # install for PG 17
pig install pljs -v 16;   # install for PG 16
pig install pljs -v 15;   # install for PG 15
pig install pljs -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pljs;
```
