---
title: "addressing_dictionary"
linkTitle: "addressing_dictionary"
description: "Address-aware text search dictionaries and configurations"
weight: 2270
categories: ["FTS"]
languages: ["SQL"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**addressing_dictionary**](https://github.com/pramsey/pgsql-addressing-dictionary) : Address-aware text search dictionaries and configurations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2270** | {{< badge content="addressing_dictionary" link="https://github.com/pramsey/pgsql-addressing-dictionary" >}} | {{< ext "addressing_dictionary" >}} | `1.1` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} {{< ext "postal" >}} {{< ext "unaccent" >}} {{< ext "hunspell_en_us" >}} |

> [!Note] SQL and dictionary data only; no shared library. PostgreSQL 14-18 verified in PGSTY builds; upstream declares no current major-version matrix.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `addressing_dictionary` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "addressing_dictionary_18" "green" >}} {{< bg "17" "addressing_dictionary_17" "green" >}} {{< bg "16" "addressing_dictionary_16" "green" >}} {{< bg "15" "addressing_dictionary_15" "green" >}} {{< bg "14" "addressing_dictionary_14" "green" >}} | `addressing_dictionary_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "postgresql-18-addressing-dictionary" "green" >}} {{< bg "17" "postgresql-17-addressing-dictionary" "green" >}} {{< bg "16" "postgresql-16-addressing-dictionary" "green" >}} {{< bg "15" "postgresql-15-addressing-dictionary" "green" >}} {{< bg "14" "postgresql-14-addressing-dictionary" "green" >}} | `postgresql-$v-addressing-dictionary` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "addressing_dictionary_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-addressing-dictionary : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-addressing-dictionary : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `addressing_dictionary_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_18-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/addressing_dictionary_18-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_18-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/addressing_dictionary_18-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_18-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/addressing_dictionary_18-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_18-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/addressing_dictionary_18-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_18` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [addressing_dictionary_18-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/addressing_dictionary_18-1.1-1PGSTY.el10.noarch.rpm) |
| `addressing_dictionary_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [addressing_dictionary_18-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/addressing_dictionary_18-1.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-18-addressing-dictionary` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
| `postgresql-18-addressing-dictionary` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.5 KiB | [postgresql-18-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-18-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `addressing_dictionary_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_17-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/addressing_dictionary_17-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_17-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/addressing_dictionary_17-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_17-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/addressing_dictionary_17-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_17-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/addressing_dictionary_17-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_17` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [addressing_dictionary_17-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/addressing_dictionary_17-1.1-1PGSTY.el10.noarch.rpm) |
| `addressing_dictionary_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [addressing_dictionary_17-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/addressing_dictionary_17-1.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-addressing-dictionary` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
| `postgresql-17-addressing-dictionary` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.5 KiB | [postgresql-17-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-17-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `addressing_dictionary_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_16-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/addressing_dictionary_16-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_16-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/addressing_dictionary_16-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_16-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/addressing_dictionary_16-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_16-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/addressing_dictionary_16-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_16` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [addressing_dictionary_16-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/addressing_dictionary_16-1.1-1PGSTY.el10.noarch.rpm) |
| `addressing_dictionary_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [addressing_dictionary_16-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/addressing_dictionary_16-1.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-addressing-dictionary` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
| `postgresql-16-addressing-dictionary` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.5 KiB | [postgresql-16-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-16-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `addressing_dictionary_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_15-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/addressing_dictionary_15-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_15-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/addressing_dictionary_15-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_15-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/addressing_dictionary_15-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_15-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/addressing_dictionary_15-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_15` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [addressing_dictionary_15-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/addressing_dictionary_15-1.1-1PGSTY.el10.noarch.rpm) |
| `addressing_dictionary_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [addressing_dictionary_15-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/addressing_dictionary_15-1.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-addressing-dictionary` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
| `postgresql-15-addressing-dictionary` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.5 KiB | [postgresql-15-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-15-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `addressing_dictionary_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_14-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/addressing_dictionary_14-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_14` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_14-1.1-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/addressing_dictionary_14-1.1-1PGSTY.el8.noarch.rpm) |
| `addressing_dictionary_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.5 KiB | [addressing_dictionary_14-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/addressing_dictionary_14-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_14` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 KiB | [addressing_dictionary_14-1.1-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/addressing_dictionary_14-1.1-1PGSTY.el9.noarch.rpm) |
| `addressing_dictionary_14` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.7 KiB | [addressing_dictionary_14-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/addressing_dictionary_14-1.1-1PGSTY.el10.noarch.rpm) |
| `addressing_dictionary_14` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 KiB | [addressing_dictionary_14-1.1-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/addressing_dictionary_14-1.1-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-addressing-dictionary` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~bookworm_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~trixie_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~jammy_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~noble_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
| `postgresql-14-addressing-dictionary` | `1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.5 KiB | [postgresql-14-addressing-dictionary_1.1-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/a/addressing-dictionary/postgresql-14-addressing-dictionary_1.1-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-addressing-dictionary" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-addressing-dictionary" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="addressing_dictionary-1.1-bf3590e038d8.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg addressing_dictionary;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install addressing_dictionary;		# install via package name, for the active PG version

pig install addressing_dictionary -v 18;   # install for PG 18
pig install addressing_dictionary -v 17;   # install for PG 17
pig install addressing_dictionary -v 16;   # install for PG 16
pig install addressing_dictionary -v 15;   # install for PG 15
pig install addressing_dictionary -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION addressing_dictionary;
```

## Usage

Sources:

- [Official README](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary.control)
- [Version 1.1 extension SQL](https://github.com/pramsey/pgsql-addressing-dictionary/blob/master/addressing_dictionary--1.1.sql)
- [PostgreSQL text-search dictionary documentation](https://www.postgresql.org/docs/current/textsearch-dictionaries.html)

addressing_dictionary installs address-oriented full-text-search configurations and dictionaries. It is useful for approximate local address lookup where street-type, direction, and ordinal variants should normalize to comparable lexemes; it is not a geocoder or a substitute for address parsing and validation.

### Core Workflow

Install the extension once in the database, normalize address text with the supplied configuration, and index the same expression used by queries.

```sql
CREATE EXTENSION addressing_dictionary;

SELECT to_tsvector('addressing_en', '1234 n main st');

CREATE INDEX address_search_idx
ON address_book
USING gin (to_tsvector('addressing_en', full_address));

SELECT *
FROM address_book
WHERE to_tsvector('addressing_en', full_address)
      @@ plainto_tsquery('addressing_en', '1234 north main street');
```

Version 1.1 keeps an abbreviated cardinal direction and also emits its expanded form, so an input such as a single-letter north token can produce both forms. This deliberately preserves ambiguity rather than assuming every occurrence has only one meaning.

### Installed Objects

- The English configuration combines a thesaurus, synonym dictionary, and stop-word dictionary for address tokens.
- The French configuration installs the corresponding French text-search objects.
- Both configurations are copied from PostgreSQL's simple configuration, so they normalize through the extension's data files instead of applying a general-language stemmer.

### Operational Boundaries

The extension is relocatable and needs no preload or restart. Its dictionary, synonym, stop-word, and thesaurus files must be installed on every server that may execute queries or restore the extension.

Full-text matching is heuristic. Abbreviations can be ambiguous, and nationwide data can produce many false or duplicate candidates. Use it to improve candidate search—especially for a bounded city or region—and retain separate structured-address or geospatial checks when correctness matters.
