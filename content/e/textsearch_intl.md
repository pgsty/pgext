---
title: "textsearch_intl"
linkTitle: "textsearch_intl"
description: "Multilingual full text search with ICU segmentation, normalization and n-gram matching"
weight: 2290
categories: ["FTS"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**textsearch_intl**](https://github.com/benjaminsanborn/textsearch_intl) : Multilingual full text search with ICU segmentation, normalization and n-gram matching


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2290** | {{< badge content="textsearch_intl" link="https://github.com/benjaminsanborn/textsearch_intl" >}} | {{< ext "textsearch_intl" >}} | `1.0` | {{< category "FTS" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "unaccent" >}} |

> [!Note] Renamed from icu_parser; PGXN distribution 1.0.1 retains SQL/control 1.0. No direct upgrade from the old name. PostgreSQL 14-18, UTF8 and unaccent required; rebuild affected indexes after an ICU major-version change. Pigsty RPM packages are 1.0.1 on EL8-10, both architectures; DEB remains 1.0.0.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `textsearch_intl` | `unaccent` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "textsearch_intl_18" "green" >}} {{< bg "17" "textsearch_intl_17" "green" >}} {{< bg "16" "textsearch_intl_16" "green" >}} {{< bg "15" "textsearch_intl_15" "green" >}} {{< bg "14" "textsearch_intl_14" "green" >}} | `textsearch_intl_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-textsearch-intl" "green" >}} {{< bg "17" "postgresql-17-textsearch-intl" "green" >}} {{< bg "16" "postgresql-16-textsearch-intl" "green" >}} {{< bg "15" "postgresql-15-textsearch-intl" "green" >}} {{< bg "14" "postgresql-14-textsearch-intl" "green" >}} | `postgresql-$v-textsearch-intl` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "textsearch_intl_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-textsearch-intl : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-textsearch-intl : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `textsearch_intl_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.8 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/textsearch_intl_18-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `textsearch_intl_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.3 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/textsearch_intl_18-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `textsearch_intl_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.4 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/textsearch_intl_18-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `textsearch_intl_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.8 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/textsearch_intl_18-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `textsearch_intl_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.6 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/textsearch_intl_18-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `textsearch_intl_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.4 KiB | [textsearch_intl_18-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/textsearch_intl_18-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.2 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.1 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.2 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 47.6 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 46.6 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 46.6 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-textsearch-intl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 45.7 KiB | [postgresql-18-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-18-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `textsearch_intl_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.7 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/textsearch_intl_17-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `textsearch_intl_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.2 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/textsearch_intl_17-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `textsearch_intl_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.3 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/textsearch_intl_17-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `textsearch_intl_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.8 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/textsearch_intl_17-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `textsearch_intl_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.6 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/textsearch_intl_17-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `textsearch_intl_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.3 KiB | [textsearch_intl_17-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/textsearch_intl_17-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.1 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.1 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.2 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.3 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 46.6 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-textsearch-intl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 45.7 KiB | [postgresql-17-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-17-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `textsearch_intl_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.7 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/textsearch_intl_16-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `textsearch_intl_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.2 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/textsearch_intl_16-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `textsearch_intl_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/textsearch_intl_16-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `textsearch_intl_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.8 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/textsearch_intl_16-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `textsearch_intl_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.6 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/textsearch_intl_16-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `textsearch_intl_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.3 KiB | [textsearch_intl_16-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/textsearch_intl_16-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.1 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.1 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.3 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 46.6 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-textsearch-intl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 45.7 KiB | [postgresql-16-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-16-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `textsearch_intl_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.7 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/textsearch_intl_15-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `textsearch_intl_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.2 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/textsearch_intl_15-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `textsearch_intl_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.3 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/textsearch_intl_15-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `textsearch_intl_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.8 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/textsearch_intl_15-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `textsearch_intl_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.6 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/textsearch_intl_15-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `textsearch_intl_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.3 KiB | [textsearch_intl_15-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/textsearch_intl_15-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.6 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.1 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.3 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.4 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 46.6 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-textsearch-intl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 45.7 KiB | [postgresql-15-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-15-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `textsearch_intl_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 53.7 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/textsearch_intl_14-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `textsearch_intl_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 52.2 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/textsearch_intl_14-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `textsearch_intl_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 55.2 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/textsearch_intl_14-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `textsearch_intl_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 53.8 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/textsearch_intl_14-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `textsearch_intl_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 55.6 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/textsearch_intl_14-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `textsearch_intl_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 54.3 KiB | [textsearch_intl_14-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/textsearch_intl_14-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.1 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 44.1 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.2 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 44.1 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 49.2 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.2 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 46.8 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 46.2 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 46.5 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-textsearch-intl` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 45.7 KiB | [postgresql-14-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/t/textsearch-intl/postgresql-14-textsearch-intl_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/benjaminsanborn/textsearch_intl" title="Repository" icon="github" subtitle="github.com/benjaminsanborn/textsearch_intl" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="textsearch_intl-1.0.1.tar.gz textsearch_intl-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg textsearch_intl;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install textsearch_intl;		# install via package name, for the active PG version

pig install textsearch_intl -v 18;   # install for PG 18
pig install textsearch_intl -v 17;   # install for PG 17
pig install textsearch_intl -v 16;   # install for PG 16
pig install textsearch_intl -v 15;   # install for PG 15
pig install textsearch_intl -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION textsearch_intl CASCADE; -- requires unaccent
```

## Usage

Sources:

- [README.md](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/README.md)
- [textsearch_intl--1.0.sql](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/textsearch_intl--1.0.sql)
- [textsearch_intl.control](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/textsearch_intl.control)
- [META.json](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/META.json)
- [LICENSE](https://github.com/benjaminsanborn/textsearch_intl/blob/6efb0a28a5f6643b46ecd2e9b8de75c30fb82680/LICENSE)
- [Rename commit](https://github.com/benjaminsanborn/textsearch_intl/commit/b023b27aa16b82fc3c1b80ef0928e9f185e7f43e)
- [PGXN textsearch_intl 1.0.1](https://pgxn.org/dist/textsearch_intl/1.0.1/)

`textsearch_intl` combines ICU word segmentation, normalization and n-gram matching for multilingual text search. It requires PostgreSQL 14 or later, a UTF8 database, ICU and `unaccent`; upstream CI covers PostgreSQL 14–18.

### Name and Version

The project renamed `icu_parser` to `textsearch_intl` on 2026-10-09 and restarted versioning. The renamed PGXN distribution 1.0.1 was published on 2026-10-22. Its SQL/control version remains `1.0`; the distribution number is not a SQL upgrade target. There is no direct SQL upgrade script from the old name: inspect dependent objects and plan their recreation before migrating.

### Core Workflow

```sql
CREATE EXTENSION textsearch_intl CASCADE;
SELECT alias, token FROM ts_debug('intl_simple', '東京駅前ラーメン店 ร้านกาแฟ Café');
SELECT to_tsvector('intl_search', 'Café') @@ phraseto_tsquery('intl_search', 'cafe');
SELECT to_tsvector('intl_search_ngram', '행복한편의점')
       @@ phraseto_tsquery('intl_search_ngram', '편의점');
SELECT intl_icu_version();
```

### Search Configurations

- `intl_simple` uses the `intl` parser and preserves accents and character widths.
- `intl_search` combines `intl_normalize`, `unaccent` and `intl_join` for normalized whole-word search. Word joining is off by default.
- `intl_search_ngram` uses `intl_runs` and `intl_ngram` to match inside Korean/CJK tokens and unspaced Thai, Lao, Khmer and Burmese runs. N-grams can overmatch and increase index size.

### Operational Boundaries

Install with a superuser; the extension is relocatable, is not marked trusted and needs no preload. Prefer `phraseto_tsquery()` when word order matters. `intl_icu_version()` reports the loaded ICU version; rebuild affected indexes after an ICU major-version change. ICU segmentation is not Japanese morphological analysis, and partial CJK prefix queries remain unreliable.
